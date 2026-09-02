/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"context"

	"github.com/go-logr/logr"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/rotisserie/eris"

	asoentrav1 "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

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
	relationshipName string,
	current []string,
	desired []string,
	add func(context.Context, string) error,
	remove func(context.Context, string) error,
	log logr.Logger,
) error {
	currentSet := set.Make(current...)
	desiredSet := set.Make(desired...)

	toAdd := desiredSet.Except(currentSet).Values()
	toRemove := currentSet.Except(desiredSet).Values()

	for _, id := range toAdd {
		if err := add(ctx, id); err != nil {
			return eris.Wrapf(err, "add %s to %s", id, relationshipName)
		}
	}

	for _, id := range toRemove {
		if err := remove(ctx, id); err != nil {
			return eris.Wrapf(err, "remove %s from %s", id, relationshipName)
		}
	}

	log.V(1).Info(
		"Reconciled relationship",
		"relationship", relationshipName,
		"added", len(toAdd),
		"removed", len(toRemove),
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

	// Set up for Reconcile Owners
	desiredOwners, err := group.Spec.ResolveOwnerObjectIDs(resolvedConfigMaps)
	if err != nil {
		return eris.Wrapf(err, "failed resolving desired owners for group %s", group.Name)
	}

	ownersBuilder := groupRequestBuilder.Owners()
	ownersRefBuilder := ownersBuilder.Ref()
	ownersDef := relationshipDefinition{
		name:    "owners",
		desired: desiredOwners,
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
			return ownersRefBuilder.Post(ctx, ref, nil)
		},
		remove: func(ctx context.Context, objectID string) error {
			return ownersBuilder.ByDirectoryObjectId(objectID).Ref().Delete(ctx, nil)
		},
	}

	currentOwners := group.Status.Owners

	if err := r.reconcileRelationship(
		ctx,
		"owners",
		currentOwners,
		ownersDef.desired,
		ownersDef.add,
		ownersDef.remove,
		log,
	); err != nil {
		return eris.Wrapf(err, "reconciling %s for group %s", ownersDef.name, id)
	}

	// Set up for reconciling Members
	desiredMembers, err := group.Spec.ResolveMemberObjectIDs(resolvedConfigMaps)
	if err != nil {
		return eris.Wrapf(err, "failed resolving desired members for group %s", group.Name)
	}

	membersBuilder := groupRequestBuilder.Members()
	membersRefBuilder := membersBuilder.Ref()
	membersDef := relationshipDefinition{
		name:    "members",
		desired: desiredMembers,
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
			return membersRefBuilder.Post(ctx, ref, nil)
		},
		remove: func(ctx context.Context, objectID string) error {
			return membersBuilder.ByDirectoryObjectId(objectID).Ref().Delete(ctx, nil)
		},
	}
	currentMembers := group.Status.Members

	if err := r.reconcileRelationship(
		ctx,
		"members",
		currentMembers,
		membersDef.desired,
		membersDef.add,
		membersDef.remove,
		log,
	); err != nil {
		return eris.Wrapf(err, "reconciling %s for group %s", membersDef.name, id)
	}

	return nil
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
