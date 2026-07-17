/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
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

func directoryObjectIDFromRef(value string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return ""
	}

	lastSlash := strings.LastIndex(trimmed, "/")
	if lastSlash < 0 || lastSlash == len(trimmed)-1 {
		return trimmed
	}

	trimmed = trimmed[lastSlash+1:]

	if questionMark := strings.Index(trimmed, "?"); questionMark >= 0 {
		trimmed = trimmed[:questionMark]
	}

	if hash := strings.Index(trimmed, "#"); hash >= 0 {
		trimmed = trimmed[:hash]
	}

	return trimmed
}
