/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "sort"

// PackageReferenceSet represents a set of distinct PackageReferences
type PackageReferenceSet struct {
	references map[PackageReference]struct{}
}

// NewPackageReferenceSet creates a new empty set of PackageReferences
func NewPackageReferenceSet(refs ...PackageReference) *PackageReferenceSet {
	result := &PackageReferenceSet{
		references: make(map[PackageReference]struct{}),
	}

	for _, ref := range refs {
		result.AddReference(ref)
	}

	return result
}

// AddReference ensures the set includes an specified Reference
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

// Contains allows checking to see if an import is included
func (set *PackageReferenceSet) Contains(ref PackageReference) bool {
	_, ok := set.references[ref]
	return ok
}

// AsSlice returns a slice containing all the imports
func (set *PackageReferenceSet) AsSlice() []PackageReference {
	var result []PackageReference
	for ref := range set.references {
		result = append(result, ref)
	}

	return result
}

// AsSortedSlice return a sorted slice containing all the references
// less specifies how to order the imports
func (set *PackageReferenceSet) AsSortedSlice(less func(i PackageReference, j PackageReference) bool) []PackageReference {
	result := set.AsSlice()

	sort.Slice(result, func(i int, j int) bool {
		return less(result[i], result[j])
	})

	return result
}

// Length returns the number of unique imports in this set
func (set *PackageReferenceSet) Length() int {
	return len(set.references)
}
