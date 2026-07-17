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

func (r *EntraSecurityGroupReconciler) reconcileRelationshipSide(
	ctx context.Context,
	side string,
	current []string,
	desired []string,
	add func(context.Context, string) error,
	remove func(context.Context, string) error,
	log logr.Logger,
) error {
	delta := planRelationshipDelta(current, desired)

	for _, id := range delta.ToAdd {
		if err := add(ctx, id); err != nil {
			// Add failures intentionally skip remove for this side in this pass.
			return eris.Wrapf(err, "%s add %s", side, id)
		}
	}

	for _, id := range delta.ToRemove {
		if err := remove(ctx, id); err != nil {
			return eris.Wrapf(err, "%s remove %s", side, id)
		}
	}

	log.V(1).Info(
		"Reconciled relationship side",
		"side", side,
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
