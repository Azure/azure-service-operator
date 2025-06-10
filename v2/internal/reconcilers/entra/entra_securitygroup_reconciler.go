/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"context"
	"fmt"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/go-logr/logr"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/rotisserie/eris"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

type EntraSecurityGroupReconciler struct {
	reconcilers.ReconcilerCommon
	ResourceResolver   *resolver.Resolver
	CredentialProvider identity.CredentialProvider
	Config             config.Values
	EntraClientFactory EntraConnectionFactory
}

var _ genruntime.Reconciler = &EntraSecurityGroupReconciler{}

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

	// TODO: Respect annotation serviceoperator.azure.com/reconcile-policy

	// If we already know the Entra ID of the group (captured in an annotation), we can update it directly
	if id, ok := getEntraID(obj); ok {
		return r.update(ctx, id, group, log)
	}

	// If we're allowed to adopt the group, we can try to find it
	if r.canAdopt(group) {
		id, err := r.tryAdopt(ctx, group, log)
		if err != nil {
			return ctrl.Result{}, eris.Wrapf(err, "trying to adopt security group %s", group.Name)
		}

		if id != nil {
			// We found an existing group to adopt
			setEntraID(obj, *id)
			return r.update(ctx, *id, group, log)
		}
	}

	// If we can't adopt, or didn't find a group to adopt, we can create a new one
	if r.canCreate(group) {
		return r.create(ctx, group, log)
	}

	// Nothing to do
	return ctrl.Result{}, nil
}

func (r *EntraSecurityGroupReconciler) Delete(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) (ctrl.Result, error) {
	log.V(Status).Info("Updating Entra security group")

	group, err := r.asSecurityGroup(obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "creating or updating security group %s", group.Name)
	}

	// If don't know the Entra ID of the group (captured in an annotation), there's nothing to do.
	id, ok := getEntraID(obj)
	if !ok {
		return ctrl.Result{}, nil
	}

	client, err := r.EntraClientFactory(ctx, obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrap(err, "creating entra client")
	}

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

// update completes our reconciliation by updating an existing Entra security group.
// ctx is the context for the operation.
// id is the Entra ID of the group to update.
// group is the security group to update.
// eventRecorder is used to record events for the group.
// log is the logger to use for logging.
func (r *EntraSecurityGroupReconciler) update(
	ctx context.Context,
	id string,
	group *asoentra.SecurityGroup,
	log logr.Logger,
) (ctrl.Result, error) {
	log = log.WithValues("id", id)
	log.V(Status).Info("Updating Entra security group")

	// Create our Entra Client
	client, err := r.EntraClientFactory(ctx, group)
	if err != nil {
		return ctrl.Result{}, eris.Wrap(err, "creating entra client prior to update")
	}

	// Load the existing group by ID
	g, err := r.loadGroupByID(ctx, id, client.Client())
	if err != nil {
		if r.isNotFound(err) {
			// Group used to exist, but no longer does - it's probably been deleted
			// Remove the existing annotation and requeue the reconciliation to create a replacement
			log.V(Status).Info("Group no longer exists")
			setEntraID(group, "")
			return ctrl.Result{
				Requeue: true,
			}, nil
		}

		return ctrl.Result{}, eris.Wrapf(err, "getting group by ID %s", id)
	}

	// Update - PATCH
	group.Spec.AssignToGroup(g)

	_, err = client.Client().Groups().ByGroupId(id).Patch(ctx, g, nil)
	if err != nil {
		// Failed to update
		return ctrl.Result{}, eris.Wrapf(err, "failed to update group %s", id)
	}

	group.Status.AssignFromGroup(g)

	return ctrl.Result{}, nil
}

// tryAdopt tries to find an existing Entra security group to adopt.
// ctx is the context for the operation.
// id is the Entra ID of the group to update.
// group is the security group to update.
// eventRecorder is used to record events for the group.
// log is the logger to use for logging.
// Returns the Entra ID of the group to adopt, and nil if found;
// nil, nil if no group was found to adopt;
// or nil, error if there was a problem finding the group.
func (r *EntraSecurityGroupReconciler) tryAdopt(
	ctx context.Context,
	group *asoentra.SecurityGroup,
	log logr.Logger,
) (*string, error) {
	log.V(Status).Info("Searching for existing Entra security group to adopt", "group", group.Name)

	// Create our Entra Client
	client, err := r.EntraClientFactory(ctx, group)
	if err != nil {
		return nil, eris.Wrap(err, "creating entra client prior to adoption search")
	}

	// Try to find the group by DisplayName
	displayName := group.Spec.DisplayName
	if displayName == nil || *displayName == "" {
		// Can't adopt without a display name
		return nil, nil
	}

	log.V(Status).Info("Searching for existing Entra security group by display name", "displayName", *displayName)
	groups, err := r.loadGroupsByDisplayName(ctx, *displayName, client.Client())
	if err != nil {
		if r.isNotFound(err) {
			// No group to adopt
			return nil, nil
		}

		return nil, eris.Wrapf(err, "getting group by display name %s", *displayName)
	}

	if len(groups) == 0 {
		// No group to adopt
		log.V(Status).Info("No existing Entra security group found by display name", "displayName", *displayName)
		return nil, nil
	}

	if len(groups) > 1 {
		// Multiple groups found with the same display name
		log.V(Status).Info("Multiple existing Entra security groups found by display name", "displayName", *displayName)
		return nil, eris.Errorf("multiple existing Entra security groups found with display name %s", *displayName)
	}

	// We found a single group to adopt
	g := groups[0]
	return g.GetId(), nil
}

// create completes our reconciliation by creating a new Entra security group.
// ctx is the context for the operation.
// group is the security group to create.
// eventRecorder is used to record events for the group.
// log is the logger to use for logging.
func (r *EntraSecurityGroupReconciler) create(
	ctx context.Context,
	group *asoentra.SecurityGroup,
	log logr.Logger,
) (ctrl.Result, error) {
	log.V(Status).Info("Creating Entra security group")

	// Create our Entra Client
	client, err := r.EntraClientFactory(ctx, group)
	if err != nil {
		return reconcile.Result{}, eris.Wrap(err, "creating entra client prior to adoption search")
	}

	g := msgraphmodels.NewGroup()
	group.Spec.AssignToGroup(g)

	status, err := client.Client().Groups().Post(ctx, g, nil)
	if err != nil {
		// Failed to create
		return ctrl.Result{}, eris.Wrapf(err, "failed to create group %s", group.Name)
	}

	group.Status.AssignFromGroup(status)
	setEntraID(group, *status.GetId())

	return ctrl.Result{}, nil
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

	id, ok := getEntraID(obj)
	if !ok {
		// If we don't know the Entra ID of the group, there's nothing to do.
		log.V(Status).Info("No Entra ID found for security group, skipping status update")
		return nil
	}

	groupable, err := r.loadGroupByID(ctx, id, client.Client())
	if err != nil {
		// If the group doesn't exist, nothing to do as we're probably in the midst of deleting it
		if r.isNotFound(err) {
			return nil
		}

		// If the error is not a 404, return the error
		return eris.Wrapf(err, "failed to update status of security group %s", id)
	}

	group.Status.AssignFromGroup(groupable)

	return nil
}

// loadGroupByID loads the group from Entra, if it exists.
// Returns the group and nil if it exists, nil and nil if it doesn't, or nil and an error if there was a problem.
func (r *EntraSecurityGroupReconciler) loadGroupByID(
	ctx context.Context,
	id string,
	client *msgraphsdkgo.GraphServiceClient,
) (msgraphmodels.Groupable, error) {
	// Try to get the group by ID
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

// loadGroupsByDisplayName loads groups from Entra by display name.
func (r *EntraSecurityGroupReconciler) loadGroupsByDisplayName(
	ctx context.Context,
	displayName string,
	client *msgraphsdkgo.GraphServiceClient,
) ([]msgraphmodels.Groupable, error) {
	// Try to get the group by display name
	filterStr := fmt.Sprintf("displayName eq '%s'", displayName)

	query := &groups.GroupsRequestBuilderGetQueryParameters{
		Filter: &filterStr,
	}

	options := &groups.GroupsRequestBuilderGetRequestConfiguration{
		QueryParameters: query,
	}

	result, err := client.Groups().Get(ctx, options)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to search for group with name %s", displayName)
	}

	groups := result.GetValue()
	return groups, nil
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

func (r *EntraSecurityGroupReconciler) asSecurityGroup(obj genruntime.MetaObject) (*asoentra.SecurityGroup, error) {
	typedObj, ok := obj.(*asoentra.SecurityGroup)
	if !ok {
		return nil, eris.Errorf("cannot modify resource that is not of type *entra.SecurityGroup. Type is %T", obj)
	}

	return typedObj, nil
}

func (r *EntraSecurityGroupReconciler) canAdopt(group *asoentra.SecurityGroup) bool {
	if group.Spec.OperatorSpec == nil {
		// Default is AdoptOrCreate
		return true
	}

	return group.Spec.OperatorSpec.AdoptionAllowed()
}

func (r *EntraSecurityGroupReconciler) canCreate(group *asoentra.SecurityGroup) bool {
	if group.Spec.OperatorSpec == nil {
		// Default is AdoptOrCreate
		return true
	}

	return group.Spec.OperatorSpec.CreationAllowed()
}
