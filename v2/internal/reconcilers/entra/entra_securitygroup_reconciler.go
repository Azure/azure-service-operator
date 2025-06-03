/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"context"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/go-logr/logr"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/rotisserie/eris"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
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

	// If we know the group ID, we can try to load it for an update
	var existingGroup msgraphmodels.Groupable
	id := group.Spec.EntraID
	if id != nil && *id != "" {
		g, err := r.loadGroup(ctx, *id, client.Client(), log)
		if err != nil && !r.isNotFound(err) {
			return ctrl.Result{}, eris.Wrap(err, "getting group")
		}

		existingGroup = g
	}

	if existingGroup == nil {
		// Creation - PUT
		log.V(Status).Info("Creating Entra security group")
		newGroup := msgraphmodels.NewGroup()
		group.Spec.AssignToGroup(newGroup)

		statusGroup, err := client.Client().Groups().Post(ctx, newGroup, nil)
		if err != nil {
			// Failed to create
			return ctrl.Result{}, eris.Wrapf(err, "failed to create group %s", group.Name)
		}

		group.Status.AssignFromGroup(statusGroup)
	} else {
		// Update - PATCH
		log.V(Status).Info("Updating Entra security group")
		group.Spec.AssignToGroup(existingGroup)

		statusGroup, err := client.Client().Groups().ByGroupId(*id).Patch(ctx, existingGroup, nil)
		if err != nil {
			// Failed to update
			return ctrl.Result{}, eris.Wrapf(err, "failed to update group %s", group.Name)
		}

		group.Status.AssignFromGroup(statusGroup)
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

	// If the group doesn't have an ID, something went wrong with our webhooks
	// but we know there's nothing to delete from Entra so there's nothing to do
	if group.Spec.EntraID == nil || *group.Spec.EntraID == "" {
		log.V(Status).Info("No Entra ID found for security group, nothing to delete")
		return ctrl.Result{}, nil
	}

	id := *group.Spec.EntraID
	err = client.Client().Groups().ByGroupId(id).Delete(ctx, nil)
	if err != nil {
		// If the group doesn't exist, return nil and nil as we've successfully ensured that it doesn't exist
		if r.isNotFound(err) {
			return ctrl.Result{}, nil
		}

		// If the error is not a 404, return the error
		var displayName string
		if group.Spec.DisplayName != nil {
			displayName = *group.Spec.DisplayName
		}

		return ctrl.Result{}, eris.Wrapf(err, "failed to delete group %q (%s)", displayName, id)
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

	id := *group.Spec.EntraID
	groupable, err := r.loadGroup(ctx, id, client.Client(), log)
	if err != nil {
		// If the group doesn't exist, return nil and nil as we're probably in the midst of deleting it
		if r.isNotFound(err) {
			return nil
		}

		// If the error is not a 404, return the error
		return eris.Wrapf(err, "failed to update status of security group %s", id)
	}

	group.Status.AssignFromGroup(groupable)

	return nil
}

// loadGroup loads the group from Entra, if it exists.
// Returns the group and nil if it exists, nil and nil if it doesn't, or nil and an error if there was a problem.
func (r *EntraSecurityGroupReconciler) loadGroup(
	ctx context.Context,
	id string,
	client *msgraphsdkgo.GraphServiceClient,
	_ logr.Logger,
) (msgraphmodels.Groupable, error) {
	groupable, err := client.Groups().ByGroupId(id).Get(ctx, nil)
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
func (r *EntraSecurityGroupReconciler) isNotFound(err error) bool {
	var odataError *odataerrors.ODataError
	if eris.As(err, &odataError) {
		if odataError.ResponseStatusCode == 404 {
			return true
		}
	}

	return false
}
