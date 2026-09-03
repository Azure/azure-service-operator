/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-beta-sdk-go"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/groups"
	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
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

	// Work out which owners we want for the group
	desiredOwners, err := group.Spec.ResolveOwnerObjectIDs(resolvedConfigMaps)
	if err != nil {
		return eris.Wrapf(err, "failed resolving desired owners for group %s", group.Name)
	}

	// Reconcile the owners relationship for the group
	if err := r.reconcileRelationship(
		ctx,
		"owners",
		group.Status.Owners,
		desiredOwners,
		r.addOwner(groupRequestBuilder),
		r.removeOwner(groupRequestBuilder, log),
		log,
	); err != nil {
		return eris.Wrapf(err, "reconciling owners for group %s", id)
	}

	// Work out which members we want for the group
	desiredMembers, err := group.Spec.ResolveMemberObjectIDs(resolvedConfigMaps)
	if err != nil {
		return eris.Wrapf(err, "failed resolving desired members for group %s", group.Name)
	}

	// Reconcile the members relationship for the group
	if err := r.reconcileRelationship(
		ctx,
		"members",
		group.Status.Members,
		desiredMembers,
		r.addMember(groupRequestBuilder),
		r.removeMember(groupRequestBuilder),
		log,
	); err != nil {
		return eris.Wrapf(err, "reconciling members for group %s", id)
	}

	return nil
}

// addOwner returns a function that can be used to add an owner to a group using the provided GroupItemRequestBuilder.
func (r *EntraSecurityGroupReconciler) addOwner(
	groupRequestBuilder *groups.GroupItemRequestBuilder,
) func(ctx context.Context, objectID string) error {
	ownersRefBuilder := groupRequestBuilder.Owners().Ref()
	return func(ctx context.Context, objectID string) error {
		ref := msgraphmodels.NewReferenceCreate()
		ref.SetOdataId(to.Ptr(asoentrav1.DirectoryObjectRefURI(objectID)))
		err := ownersRefBuilder.Post(ctx, ref, nil)
		if err != nil {
			return eris.Wrapf(err, "failed adding owner %s to group", objectID)
		}

		return nil
	}
}

// removeOwner returns a function that can be used to remove an owner from a group using the provided GroupItemRequestBuilder.
func (r *EntraSecurityGroupReconciler) removeOwner(
	groupRequestBuilder *groups.GroupItemRequestBuilder,
	log logr.Logger,
) func(ctx context.Context, objectID string) error {
	ownersBuilder := groupRequestBuilder.Owners()
	return func(ctx context.Context, objectID string) error {
		err := ownersBuilder.ByDirectoryObjectId(objectID).Ref().Delete(ctx, nil)
		// Entra doesn't allow removal of the last user from Owners, even if there are other owners listed.
		// If we encounter that situation, we just leave the user there as the alternative is to have the ASO resource
		// in a permanent failed state.
		if err != nil {
			if strings.Contains(err.Error(), "this owner cannot be removed") {
				// Ignore this error as it indicates we are trying to remove the last owner
				log.V(1).Info("entra does not permit removal of the last owner", "objectID", objectID)
				return nil
			}

			return err
		}

		return nil
	}
}

// addMember returns a function that can be used to add a member to a group using the provided GroupItemRequestBuilder.
func (r *EntraSecurityGroupReconciler) addMember(
	groupRequestBuilder *groups.GroupItemRequestBuilder,
) func(ctx context.Context, objectID string) error {
	membersRefBuilder := groupRequestBuilder.Members().Ref()
	return func(ctx context.Context, objectID string) error {
		ref := msgraphmodels.NewReferenceCreate()
		ref.SetOdataId(to.Ptr(asoentrav1.DirectoryObjectRefURI(objectID)))
		return membersRefBuilder.Post(ctx, ref, nil)
	}
}

// removeMember returns a function that can be used to remove a member from a group using the provided GroupItemRequestBuilder.
func (r *EntraSecurityGroupReconciler) removeMember(
	groupRequestBuilder *groups.GroupItemRequestBuilder,
) func(ctx context.Context, objectID string) error {
	membersBuilder := groupRequestBuilder.Members()
	return func(ctx context.Context, objectID string) error {
		return membersBuilder.ByDirectoryObjectId(objectID).Ref().Delete(ctx, nil)
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
