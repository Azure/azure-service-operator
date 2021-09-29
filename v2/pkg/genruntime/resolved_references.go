/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import "github.com/pkg/errors"

// ResolvedReferences is a set of references which have been resolved for a particular resource.
// The special self field is the fully qualified ARM ID of the resource that this ResolvedReferences applies to.
type ResolvedReferences struct {
	// references is a map of ResourceReference to ARM ID.
	references map[ResourceReference]string
}

// MakeResolvedReferences creates a ResolvedReferences from the fully qualified ARM ID of the resource and
// and ARM IDs that the resource refers to.
func MakeResolvedReferences(references map[ResourceReference]string) ResolvedReferences {
	return ResolvedReferences{
		references: references,
	}
}

// ARMID looks up the fully qualified ARM ID for the given reference. If it cannot be found, false is returned for the second parameter.
func (r ResolvedReferences) ARMID(ref ResourceReference) (string, bool) {
	result, ok := r.references[ref]
	return result, ok
}

// ARMIDOrErr looks up the fully qualified ARM ID for the given reference. If it cannot be found, an error is returned.
func (r ResolvedReferences) ARMIDOrErr(ref ResourceReference) (string, error) {
	result, ok := r.references[ref]
	if !ok {
		return "", errors.Errorf("couldn't find resolved reference %s", ref.String())
	}
	return result, nil
}
