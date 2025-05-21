/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"context"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"k8s.io/client-go/tools/record"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

var _ genruntime.Reconciler = &EntraSecurityGroupReconciler{}

type EntraSecurityGroupReconciler struct {
	reconcilers.ReconcilerCommon
	ResourceResolver   *resolver.Resolver
	CredentialProvider identity.CredentialProvider
	Config             config.Values
	EntraClientFactory EntraConnectionFactory
}

func NewEntraSecurityGroupReconciler(
	kubeClient kubeclient.Client,
	entraClientFactory EntraConnectionFactory,
	resourceResolver *resolver.Resolver,
	positiveConditions *conditions.PositiveConditionBuilder,
	cfg config.Values,
) *EntraSecurityGroupReconciler {
	return &EntraSecurityGroupReconciler{
		ResourceResolver:   resourceResolver,
		Config:             cfg,
		EntraClientFactory: entraClientFactory,
		ReconcilerCommon: reconcilers.ReconcilerCommon{
			KubeClient:         kubeClient,
			PositiveConditions: positiveConditions,
		},
	}
}

func (r *EntraSecurityGroupReconciler) asSecurityGroup(obj genruntime.MetaObject) (*asoentra.SecurityGroup, error) {
	typedObj, ok := obj.(*asoentra.SecurityGroup)
	if !ok {
		return nil, eris.Errorf("cannot modify resource that is not of type *entra.SecurityGroup. Type is %T", obj)
	}

	return typedObj, nil
}

func (r *EntraSecurityGroupReconciler) CreateOrUpdate(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) (ctrl.Result, error) {
	group, err := r.asSecurityGroup(obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "creating or updating security group %s", group.Name)
	}

	client, err := r.EntraClientFactory(ctx, obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrap(err, "creating entra client")
	}

	// Augment Log
	azureName := group.Spec.AzureName
	log = log.WithValues("azureName", azureName)

	groupable, err := r.loadGroup(ctx, azureName, client.Client(), log)
	if err != nil {
		return ctrl.Result{}, eris.Wrap(err, "getting group")
	}

	var mode string
	if groupable == nil {
		// Creation
		log.V(Status).Info("Creating Entra security group")
		mode = "create"
		groupable = msgraphmodels.NewGroup()
	} else {
		// Update
		log.V(Status).Info("Updating Entra security group")
		mode = "update"
	}

	group.Spec.AssignToGroup(groupable)
	_, err = client.Client().Groups().Post(ctx, groupable, nil)
	if err != nil {
		// Failed to update
		return ctrl.Result{}, eris.Wrapf(err, "failed to %s group", mode)
	}

	/*
		// Resolve the secrets
		secrets, err := r.ResourceResolver.ResolveResourceSecretReferences(ctx, user)
		if err != nil {
			return ctrl.Result{}, reconcilers.ClassifyResolverError(err)
		}
	*/

	return ctrl.Result{}, err
}

func (r *EntraSecurityGroupReconciler) Delete(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) (ctrl.Result, error) {
	group, err := r.asSecurityGroup(obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "creating or updating security group %s", group.Name)
	}

	client, err := r.EntraClientFactory(ctx, obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrap(err, "creating entra client")
	}

	// Augment Log
	azureName := group.Spec.AzureName
	log = log.WithValues("azureName", azureName)

	err = client.Client().Groups().ByGroupId(*group.Status.ID).Delete(ctx, nil)
	if err != nil {
		// If the group doesn't exist, return nil and nil as we've successfully ensured that it doesn't exist
		if r.isNotFound(err) {
			return ctrl.Result{}, nil
		}

		// If the error is not a 404, return the error
		return ctrl.Result{}, eris.Wrapf(err, "failed to delete group %s", azureName)
	}

	return ctrl.Result{}, nil
}

func (r *EntraSecurityGroupReconciler) Claim(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) error {
	// Nothing to do
	//!! Confirm this is true
	return nil
}

func (r *EntraSecurityGroupReconciler) UpdateStatus(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) error {
	group, err := r.asSecurityGroup(obj)
	if err != nil {
		return eris.Wrapf(err, "updating status of security group %s", group.Name)
	}

	client, err := r.EntraClientFactory(ctx, obj)
	if err != nil {
		return eris.Wrap(err, "creating entra client")
	}

	groupable, err := r.loadGroup(ctx, group.Spec.AzureName, client.Client(), log)
	if err != nil {
		// If the group doesn't exist, return nil and nil as we're probably in the midst of deleting it
		if r.isNotFound(err) {
			return nil
		}

		// If the error is not a 404, return the error
		return eris.Wrapf(err, "failed to update status of security group %s", group.Spec.AzureName)
	}

	group.Status.AssignFromGroup(groupable)

	return nil
}

// loadGroup loads the group from Entra, if it exists.
// Returns the group and nil if it exists, nil and nil if it doesn't, or nil and an error if there was a problem.
func (r *EntraSecurityGroupReconciler) loadGroup(
	ctx context.Context,
	name string,
	client *msgraphsdkgo.GraphServiceClient,
	_ logr.Logger,
) (msgraphmodels.Groupable, error) {
	groupable, err := client.Groups().ByGroupId(name).Get(ctx, nil)
	if err != nil {
		// If the only problem is that the group doesn't exist, return nil and nil
		if r.isNotFound(err) {
			return nil, nil
		}

		return nil, err
	}

	return groupable, nil
}

// isNotFound returns true if the error is a 404 error.
func (r *EntraSecurityGroupReconciler) isNotFound(_ error) bool {
	//!! Fix
	return false
}
