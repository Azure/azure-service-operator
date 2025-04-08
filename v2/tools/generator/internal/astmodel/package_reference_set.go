/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"iter"
)

// PackageReferenceSet represents a set of distinct PackageReferences
type PackageReferenceSet struct {
	references map[PackageReference]struct{}
}

// NewPackageReferenceSet creates a new empty set of PackageReferences
func NewPackageReferenceSet(refs ...PackageReference) *PackageReferenceSet {
	result := &PackageReferenceSet{
		references: make(map[PackageReference]struct{}, len(refs)),
	}

	for _, ref := range refs {
		result.AddReference(ref)
	}

	return result
}

// AddReference ensures the set includes a specified Reference
func (set *PackageReferenceSet) AddReference(ref PackageReference) {
	set.references[ref] = struct{}{}
}

// Merge ensures that all references specified in other are included
func (set *PackageReferenceSet) Merge(other *PackageReferenceSet) {
	for ref := range other.references {
		set.AddReference(ref)
	}
}

// Remove ensures the specified item is not present
// Removing an item not in the set is not an error.
func (set *PackageReferenceSet) Remove(ref PackageReference) {
	delete(set.references, ref)
}

// Clear removes everything from the set
func (set *PackageReferenceSet) Clear() {
	set.references = make(map[PackageReference]struct{})
}

// Contains allows checking to see if an import is included
func (set *PackageReferenceSet) Contains(ref PackageReference) bool {
	_, ok := set.references[ref]
	return ok
}

// All returns an iterator over all the references in the set
func (set *PackageReferenceSet) All() iter.Seq[PackageReference] {
	return func(yield func(ref PackageReference) bool) {
		for ref := range set.references {
			if !yield(ref) {
				break
			}
		}
	}
}

// Length returns the number of unique imports in this set
func (set *PackageReferenceSet) Length() int {
	return len(set.references)
}
