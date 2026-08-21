/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"context"

	"github.com/go-logr/logr"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/rotisserie/eris"

	asoentrav1 "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/set"
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
	currentSet := set.Make[string](current...)
	desiredSet := set.Make[string](desired...)

	return relationshipDelta{
		ToAdd:    desiredSet.Except(currentSet).Values(),
		ToRemove: currentSet.Except(desiredSet).Values(),
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

// relationshipDefinition bundles everything reconcileRelationshipSide needs to bring one
// side (owners or members) of a group's directory-object relationships to the
// desired state. The msgraph SDK generates distinct types per side, so we hide the
// divergence behind these closures and let the reconciler treat both sides the same.
type relationshipDefinition struct {
	name    string
	desired []string
	list    func(context.Context) ([]string, error)
	add     func(context.Context, string) error
	remove  func(context.Context, string) error
}

// reconcileRelationship brings a single side (owners or members) to its desired
// state. We bias toward availability: adds run before removes and, if an add fails,
// we return without touching removes so the group cannot end up transiently empty
// while we still cannot restore the intended members.
func (r *EntraSecurityGroupReconciler) reconcileRelationship(
	ctx context.Context,
	def relationshipDefinition,
	current []string,
	log logr.Logger,
) error {
	delta := planRelationshipDelta(current, def.desired)

	for _, id := range delta.ToAdd {
		if err := def.add(ctx, id); err != nil {
			return eris.Wrapf(err, "%s add %s", def.name, id)
		}
	}

	for _, id := range delta.ToRemove {
		if err := def.remove(ctx, id); err != nil {
			return eris.Wrapf(err, "%s remove %s", def.name, id)
		}
	}

	log.V(1).Info(
		"Reconciled relationship definition",
		"definition", def.name,
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
) error {
	id, ok := getEntraID(group)
	if !ok || id == "" {
		return eris.Errorf("missing Entra ID annotation for security group %s", group.Name)
	}

	resolvedConfigMaps, err := r.ResourceResolver.ResolveResourceConfigMapReferences(ctx, group)
	if err != nil {
		return eris.Wrapf(err, "failed resolving config map references for group %s", group.Name)
	}

	groupRequestBuilder := graphClient.Groups().ByGroupId(id)

	var definitions []relationshipDefinition

	// Set up for Reconcile Owners
	desired, err := group.Spec.ResolveOwnerObjectIDs(resolvedConfigMaps)
	if err != nil {
		return eris.Wrapf(err, "failed resolving desired owners for group %s", group.Name)
	}

	definitions = append(definitions, ownersRelationshipDefinition(groupRequestBuilder, desired))

	// Set up for reconciling Members
	desired, err = group.Spec.ResolveMemberObjectIDs(resolvedConfigMaps)
	if err != nil {
		return eris.Wrapf(err, "failed resolving desired members for group %s", group.Name)
	}

	definitions = append(definitions, membersRelationshipDefinition(groupRequestBuilder, desired))

	for _, def := range definitions {
		current, err := def.list(ctx)
		if err != nil {
			return eris.Wrapf(err, "%s list for group %s", def.name, id)
		}

		if err := r.reconcileRelationship(ctx, def, current, log); err != nil {
			return eris.Wrapf(err, "reconciling %s for group %s", def.name, id)
		}
	}

	return nil
}

// ownersRelationshipDefinition provides a relationshipDefinition for updating SecurityGroup owners.
func ownersRelationshipDefinition(
	groupBuilder *groups.GroupItemRequestBuilder,
	desired []string,
) relationshipDefinition {
	ownersBuilder := groupBuilder.Owners()
	refBuilder := ownersBuilder.Ref()
	return relationshipDefinition{
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
			ref.SetOdataId(to.Ptr(asoentrav1.DirectoryObjectRefURI(objectID)))
			return refBuilder.Post(ctx, ref, nil)
		},
		remove: func(ctx context.Context, objectID string) error {
			deleteID := asoentrav1.DirectoryObjectRefURI(objectID)
			return refBuilder.Delete(ctx, &groups.ItemOwnersRefRequestBuilderDeleteRequestConfiguration{
				QueryParameters: &groups.ItemOwnersRefRequestBuilderDeleteQueryParameters{
					Id: &deleteID,
				},
			})
		},
	}
}

// membersRelationshipDefinition provides a relationshipDefinition for updating SecurityGroup members.
func membersRelationshipDefinition(
	groupBuilder *groups.GroupItemRequestBuilder,
	desired []string,
) relationshipDefinition {
	membersBuilder := groupBuilder.Members()
	refBuilder := membersBuilder.Ref()
	return relationshipDefinition{
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
			ref.SetOdataId(to.Ptr(asoentrav1.DirectoryObjectRefURI(objectID)))
			return refBuilder.Post(ctx, ref, nil)
		},
		remove: func(ctx context.Context, objectID string) error {
			deleteID := asoentrav1.DirectoryObjectRefURI(objectID)
			return refBuilder.Delete(ctx, &groups.ItemMembersRefRequestBuilderDeleteRequestConfiguration{
				QueryParameters: &groups.ItemMembersRefRequestBuilderDeleteQueryParameters{
					Id: &deleteID,
				},
			})
		},
	}
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
