/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/go-logr/logr"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/rotisserie/eris"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra"
	asoentrav1 "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
)

// EntraSecurityGroupReconciler reconciles an Entra security group.
// TODO: Factor out common code shared with other Entra resources into entraGenericReconciler
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
		return ctrl.Result{}, eris.Wrapf(err, "creating or updating security group %s", obj.GetName())
	}

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

		if id != "" {
			// We found an existing group to adopt
			setEntraID(obj, id)
			return r.update(ctx, id, group, log)
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
	log.V(Status).Info("Deleting Entra security group")

	group, err := r.asSecurityGroup(obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "deleting security group %s", obj.GetName())
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
	group *asoentrav1.SecurityGroup,
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

	if err := r.reconcileOwnersAndMembers(ctx, group, client.Client(), log); err != nil {
		return classifyRelationshipError(err)
	}

	group.Status.AssignFromGroup(g)

	return ctrl.Result{}, nil
}

func (r *EntraSecurityGroupReconciler) reconcileOwnersAndMembers(
	ctx context.Context,
	group *asoentrav1.SecurityGroup,
	graphClient *msgraphsdkgo.GraphServiceClient,
	log logr.Logger,
) error {
	id, ok := getEntraID(group)
	if !ok || id == "" {
		return eris.Errorf("missing Entra ID annotation for security group %s", group.Name)
	}

	manageOwners, manageMembers := relationshipSidesToManage(group.Spec)
	if !manageOwners && !manageMembers {
		return nil
	}

	resolvedConfigMaps, err := r.ResourceResolver.ResolveResourceConfigMapReferences(ctx, group)
	if err != nil {
		return eris.Wrapf(err, "failed resolving config map references for group %s", group.Name)
	}

	groupRequestBuilder := graphClient.Groups().ByGroupId(id)

	var sides []relationshipSide
	if manageOwners {
		desired, err := group.Spec.ResolveOwnerObjectIDs(resolvedConfigMaps)
		if err != nil {
			return eris.Wrapf(err, "failed resolving desired owners for group %s", group.Name)
		}
		sides = append(sides, ownersSide(groupRequestBuilder, desired))
	}
	if manageMembers {
		desired, err := group.Spec.ResolveMemberObjectIDs(resolvedConfigMaps)
		if err != nil {
			return eris.Wrapf(err, "failed resolving desired members for group %s", group.Name)
		}
		sides = append(sides, membersSide(groupRequestBuilder, desired))
	}

	// Each side reconciles independently so an outage on one side (typically a
	// permissions issue) does not block the other from converging.
	var sideErrors []error
	for _, side := range sides {
		current, err := side.list(ctx)
		if err != nil {
			sideErrors = append(sideErrors, eris.Wrapf(err, "%s list for group %s", side.name, id))
			continue
		}
		if err := r.reconcileRelationshipSide(ctx, side, current, log); err != nil {
			sideErrors = append(sideErrors, eris.Wrapf(err, "reconciling %s for group %s", side.name, id))
		}
	}

	if len(sideErrors) > 0 {
		return errors.Join(sideErrors...)
	}

	return nil
}

// ownersSide adapts the msgraph SDK's owners endpoint into a relationshipSide.
func ownersSide(
	groupBuilder *groups.GroupItemRequestBuilder,
	desired []string,
) relationshipSide {
	ownersBuilder := groupBuilder.Owners()
	refBuilder := ownersBuilder.Ref()
	return relationshipSide{
		name:    "owners",
		desired: desired,
		list: func(ctx context.Context) ([]string, error) {
			return collectDirectoryObjectIDs(
				ctx,
				func(ctx context.Context) (msgraphmodels.DirectoryObjectCollectionResponseable, error) {
					return ownersBuilder.Get(ctx, nil)
				},
				func(nextLink string) (msgraphmodels.DirectoryObjectCollectionResponseable, error) {
					return ownersBuilder.WithUrl(nextLink).Get(ctx, nil)
				},
			)
		},
		add: func(ctx context.Context, objectID string) error {
			ref := msgraphmodels.NewReferenceCreate()
			ref.SetOdataId(to.Ptr(asoentra.DirectoryObjectRefURI(objectID)))
			return refBuilder.Post(ctx, ref, nil)
		},
		remove: func(ctx context.Context, objectID string) error {
			deleteID := asoentra.DirectoryObjectRefURI(objectID)
			return refBuilder.Delete(ctx, &groups.ItemOwnersRefRequestBuilderDeleteRequestConfiguration{
				QueryParameters: &groups.ItemOwnersRefRequestBuilderDeleteQueryParameters{
					Id: &deleteID,
				},
			})
		},
	}
}

// membersSide adapts the msgraph SDK's members endpoint into a relationshipSide.
func membersSide(
	groupBuilder *groups.GroupItemRequestBuilder,
	desired []string,
) relationshipSide {
	membersBuilder := groupBuilder.Members()
	refBuilder := membersBuilder.Ref()
	return relationshipSide{
		name:    "members",
		desired: desired,
		list: func(ctx context.Context) ([]string, error) {
			return collectDirectoryObjectIDs(
				ctx,
				func(ctx context.Context) (msgraphmodels.DirectoryObjectCollectionResponseable, error) {
					return membersBuilder.Get(ctx, nil)
				},
				func(nextLink string) (msgraphmodels.DirectoryObjectCollectionResponseable, error) {
					return membersBuilder.WithUrl(nextLink).Get(ctx, nil)
				},
			)
		},
		add: func(ctx context.Context, objectID string) error {
			ref := msgraphmodels.NewReferenceCreate()
			ref.SetOdataId(to.Ptr(asoentra.DirectoryObjectRefURI(objectID)))
			return refBuilder.Post(ctx, ref, nil)
		},
		remove: func(ctx context.Context, objectID string) error {
			deleteID := asoentra.DirectoryObjectRefURI(objectID)
			return refBuilder.Delete(ctx, &groups.ItemMembersRefRequestBuilderDeleteRequestConfiguration{
				QueryParameters: &groups.ItemMembersRefRequestBuilderDeleteQueryParameters{
					Id: &deleteID,
				},
			})
		},
	}
}

func relationshipSidesToManage(spec asoentrav1.SecurityGroupSpec) (bool, bool) {
	// Nil means omitted (unmanaged); explicit empty means managed-to-empty.
	return spec.Owners != nil, spec.Members != nil
}

func collectDirectoryObjectIDs(
	ctx context.Context,
	firstPage func(context.Context) (msgraphmodels.DirectoryObjectCollectionResponseable, error),
	nextPage func(string) (msgraphmodels.DirectoryObjectCollectionResponseable, error),
) ([]string, error) {
	response, err := firstPage(ctx)
	if err != nil {
		return nil, err
	}

	iterations := 0
	result := make([]string, 0)
	for response != nil {
		for _, entry := range response.GetValue() {
			id := to.Value(entry.GetId())
			if id == "" {
				continue
			}
			result = append(result, id)
		}

		nextLink := to.Value(response.GetOdataNextLink())
		if nextLink == "" {
			break
		}

		response, err = nextPage(nextLink)
		if err != nil {
			return nil, err
		}

		// Protect against infinite loops in case we're talking to a malicious server
		iterations++
		if iterations > 100 {
			return nil, eris.New("too many iterations while collecting directory object IDs")
		}

		// Stop if our context is cancelled or times out
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
	}

	return orderedUnique(result), nil
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
	group *asoentrav1.SecurityGroup,
	log logr.Logger,
) (string, error) {
	log.V(Status).Info("Searching for existing Entra security group to adopt", "group", group.Name)

	// Create our Entra Client
	client, err := r.EntraClientFactory(ctx, group)
	if err != nil {
		return "", eris.Wrap(err, "creating entra client prior to adoption search")
	}

	// Try to find the group by DisplayName
	displayName := group.Spec.DisplayName
	if to.Value(displayName) == "" {
		// Can't adopt without a display name
		return "", nil
	}

	log.V(Status).Info("Searching for existing Entra security group by display name", "displayName", *displayName)
	groups, err := r.loadGroupsByDisplayName(ctx, *displayName, client.Client())
	if err != nil {
		if r.isNotFound(err) {
			// No group to adopt
			return "", nil
		}

		return "", eris.Wrapf(err, "getting group by display name %s", *displayName)
	}

	if len(groups) == 0 {
		// No group to adopt
		log.V(Status).Info("No existing Entra security group found by display name", "displayName", *displayName)
		return "", nil
	}

	if len(groups) > 1 {
		// Multiple groups found with the same display name
		log.V(Status).Info("Multiple existing Entra security groups found by display name", "displayName", *displayName)
		return "", eris.Errorf("multiple existing Entra security groups found with display name %s", *displayName)
	}

	// We found a single group to adopt
	g := groups[0]
	id := g.GetId()
	if id == nil {
		return "", nil
	}

	return *id, nil
}

// create completes our reconciliation by creating a new Entra security group.
// ctx is the context for the operation.
// group is the security group to create.
// eventRecorder is used to record events for the group.
// log is the logger to use for logging.
func (r *EntraSecurityGroupReconciler) create(
	ctx context.Context,
	group *asoentrav1.SecurityGroup,
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

	// Resolve config map references for this resource so we can populate any
	// ObjectIDFromConfig values used in owners/members.
	resolvedConfigMaps, err := r.ResourceResolver.ResolveResourceConfigMapReferences(ctx, group)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "failed resolving config map references for group %s", group.Name)
	}

	if err := group.Spec.AssignODataBindOnCreate(g, resolvedConfigMaps); err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "failed preparing create payload for group %s", group.Name)
	}

	status, err := client.Client().Groups().Post(ctx, g, nil)
	if err != nil {
		// Failed to create
		return ctrl.Result{}, eris.Wrapf(err, "failed to create group %s", group.Name)
	}

	group.Status.AssignFromGroup(status)

	if id := status.GetId(); id != nil {
		setEntraID(group, *id)
	}

	err = r.saveAssociatedKubernetesResources(ctx, group, log)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "failed to save associated Kubernetes resources for group %s", group.Name)
	}

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
		return eris.Wrapf(err, "updating status of security group %s", obj.GetName())
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

	err = r.saveAssociatedKubernetesResources(ctx, group, log)
	if err != nil {
		return eris.Wrapf(err, "failed to save associated Kubernetes resources for group %s", group.Name)
	}

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
	// Try to get the group by display name.
	// Escape single quotes in the display name per the OData v4 spec: a single quote
	// within a string literal is represented as two consecutive single quotes.
	escapedDisplayName := strings.ReplaceAll(displayName, "'", "''")
	filterStr := fmt.Sprintf("displayName eq '%s'", escapedDisplayName)

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
		if odataError.ResponseStatusCode == http.StatusNotFound {
			return true
		}
	}

	return false
}

func (r *EntraSecurityGroupReconciler) asSecurityGroup(
	obj genruntime.MetaObject,
) (*asoentrav1.SecurityGroup, error) {
	typedObj, ok := obj.(*asoentrav1.SecurityGroup)
	if !ok {
		return nil, eris.Errorf("cannot modify resource that is not of type *entra.SecurityGroup. Type is %T", obj)
	}

	return typedObj, nil
}

func (r *EntraSecurityGroupReconciler) canAdopt(group *asoentrav1.SecurityGroup) bool {
	if group.Spec.OperatorSpec == nil {
		// Default is AdoptOrCreate
		return true
	}

	return group.Spec.OperatorSpec.AdoptionAllowed()
}

func (r *EntraSecurityGroupReconciler) canCreate(group *asoentrav1.SecurityGroup) bool {
	if group.Spec.OperatorSpec == nil {
		// Default is AdoptOrCreate
		return true
	}

	return group.Spec.OperatorSpec.CreationAllowed()
}

// saveAssociatedKubernetesResources retrieves Kubernetes resources to create and saves them to Kubernetes.
// If there are no resources to save this method is a no-op.
// TODO: Currently hard coded, but we should extract common features for reuse across other Entra resources
func (r *EntraSecurityGroupReconciler) saveAssociatedKubernetesResources(
	ctx context.Context,
	group *asoentrav1.SecurityGroup,
	log logr.Logger,
) error {
	if group == nil ||
		group.Spec.OperatorSpec == nil {
		// No OperatorSpec, nothing to do
		return nil
	}

	// Accumulate all the resources we need to save
	var resources []client.Object

	operatorSpec := group.Spec.OperatorSpec
	if operatorSpec.ConfigMaps != nil {
		// If we have secrets to export, we need to collect them
		collector := configmaps.NewCollector(group.Namespace)

		if operatorSpec.ConfigMaps.EntraID != nil && group.Status.EntraID != nil {
			// If we have an Entra ID secret, we need to collect it
			collector.AddValue(operatorSpec.ConfigMaps.EntraID, *group.Status.EntraID)
		}

		values, err := collector.Values()
		if err != nil {
			return eris.Wrap(err, "failed to collect configmaps for Entra security group")
		}

		if len(values) > 0 {
			resources = append(resources, configmaps.SliceToClientObjectSlice(values)...)
		}
	}

	if len(resources) == 0 {
		// No resources to save, nothing to do
		return nil
	}

	// Save the resources to Kubernetes
	results, err := genruntime.ApplyObjsAndEnsureOwner(ctx, r.KubeClient, group, resources)
	if err != nil {
		return err
	}

	for i := 0; i < len(resources); i++ {
		resource := resources[i]
		result := results[i]

		log.V(Debug).Info("Successfully created resource",
			"namespace", resource.GetNamespace(),
			"name", resource.GetName(),
			"type", fmt.Sprintf("%T", resource),
			"action", result)
	}

	return nil
}
