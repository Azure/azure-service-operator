/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/rotisserie/eris"
	ctrl "sigs.k8s.io/controller-runtime"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra"
	asoentrav1 "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

type relationshipDelta struct {
	ToAdd    []string
	ToRemove []string
}

// planRelationshipDelta returns which ids need to be added or removed to move
// current to desired. Both inputs are expected to be pre-deduplicated by their
// caller (collectDirectoryObjectIDs for current, ResolveOwnerObjectIDs /
// ResolveMemberObjectIDs for desired); duplicates in the inputs will appear
// duplicated in the output.
func planRelationshipDelta(current []string, desired []string) relationshipDelta {
	currentSet := make(map[string]struct{}, len(current))
	desiredSet := make(map[string]struct{}, len(desired))

	for _, value := range current {
		currentSet[value] = struct{}{}
	}

	for _, value := range desired {
		desiredSet[value] = struct{}{}
	}

	toAdd := make([]string, 0)
	for _, value := range desired {
		if _, ok := currentSet[value]; !ok {
			toAdd = append(toAdd, value)
		}
	}

	toRemove := make([]string, 0)
	for _, value := range current {
		if _, ok := desiredSet[value]; !ok {
			toRemove = append(toRemove, value)
		}
	}

	return relationshipDelta{
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
}

func orderedUnique(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))

	for _, value := range values {
		if _, ok := seen[value]; ok {
			continue
		}

		seen[value] = struct{}{}
		result = append(result, value)
	}

	return result
}

// relationshipSide bundles everything reconcileRelationshipSide needs to bring one
// side (owners or members) of a group's directory-object relationships to the
// desired state. The msgraph SDK generates distinct types per side, so we hide the
// divergence behind these closures and let the reconciler treat both sides the same.
type relationshipSide struct {
	name    string
	desired []string
	list    func(context.Context) ([]string, error)
	add     func(context.Context, string) error
	remove  func(context.Context, string) error
}

// reconcileRelationshipSide brings a single side (owners or members) to its desired
// state. We bias toward availability: adds run before removes and, if an add fails,
// we return without touching removes so the group cannot end up transiently empty
// while we still cannot restore the intended members.
func (r *EntraSecurityGroupReconciler) reconcileRelationshipSide(
	ctx context.Context,
	side relationshipSide,
	current []string,
	log logr.Logger,
) error {
	delta := planRelationshipDelta(current, side.desired)

	for _, id := range delta.ToAdd {
		if err := side.add(ctx, id); err != nil {
			return eris.Wrapf(err, "%s add %s", side.name, id)
		}
	}

	for _, id := range delta.ToRemove {
		if err := side.remove(ctx, id); err != nil {
			return eris.Wrapf(err, "%s remove %s", side.name, id)
		}
	}

	log.V(1).Info(
		"Reconciled relationship side",
		"side", side.name,
		"added", len(delta.ToAdd),
		"removed", len(delta.ToRemove),
	)

	return nil
}

func (r *EntraSecurityGroupReconciler) reconcileOwnersAndMembers(
	ctx context.Context,
	group *asoentrav1.SecurityGroup,
	graphClient *msgraphsdkgo.GraphServiceClient,
	log logr.Logger,
) (ctrl.Result, error) {
	id, ok := getEntraID(group)
	if !ok || id == "" {
		return ctrl.Result{}, eris.Errorf("missing Entra ID annotation for security group %s", group.Name)
	}

	manageOwners, manageMembers := relationshipSidesToManage(group.Spec)
	if !manageOwners && !manageMembers {
		return ctrl.Result{}, nil
	}

	resolvedConfigMaps, err := r.ResourceResolver.ResolveResourceConfigMapReferences(ctx, group)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "failed resolving config map references for group %s", group.Name)
	}

	groupRequestBuilder := graphClient.Groups().ByGroupId(id)

	var sides []relationshipSide
	if manageOwners {
		desired, err := group.Spec.ResolveOwnerObjectIDs(resolvedConfigMaps)
		if err != nil {
			return ctrl.Result{}, eris.Wrapf(err, "failed resolving desired owners for group %s", group.Name)
		}
		sides = append(sides, ownersSide(groupRequestBuilder, desired))
	}
	if manageMembers {
		desired, err := group.Spec.ResolveMemberObjectIDs(resolvedConfigMaps)
		if err != nil {
			return ctrl.Result{}, eris.Wrapf(err, "failed resolving desired members for group %s", group.Name)
		}
		sides = append(sides, membersSide(groupRequestBuilder, desired))
	}

	// Each side reconciles independently so an outage on one side (typically a
	// permissions issue) does not block the other from converging. We extract any
	// HTTP 429 Retry-After per side at the point of failure so classifyRelationshipError
	// does not have to walk the joined-error tree, and take the largest across sides.
	var (
		sideErrors []error
		throttle   ctrl.Result
	)
	recordFailure := func(err error) {
		sideErrors = append(sideErrors, err)
		throttle = maxThrottleResult(throttle, retryAfterResult(err))
	}
	for _, side := range sides {
		current, err := side.list(ctx)
		if err != nil {
			recordFailure(eris.Wrapf(err, "%s list for group %s", side.name, id))
			continue
		}
		if err := r.reconcileRelationshipSide(ctx, side, current, log); err != nil {
			recordFailure(eris.Wrapf(err, "reconciling %s for group %s", side.name, id))
		}
	}

	if len(sideErrors) > 0 {
		return throttle, errors.Join(sideErrors...)
	}

	return ctrl.Result{}, nil
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
