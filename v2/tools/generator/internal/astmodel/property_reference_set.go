/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// PropertyReferenceSet is a set of property references
type PropertyReferenceSet struct {
	references map[PropertyReference]bool
}

// NewPropertyReferenceSet returns a new empty property reference set
func NewPropertyReferenceSet() *PropertyReferenceSet {
	return &PropertyReferenceSet{
		references: make(map[PropertyReference]bool),
	}
}

// IsEmpty returns true if the set is empty, false otherwise.
func (set *PropertyReferenceSet) IsEmpty() bool {
	return len(set.references) == 0
}

// Add ensures the set contains the specified references.
func (set *PropertyReferenceSet) Add(refs ...PropertyReference) {
	for _, r := range refs {
		set.references[r] = true
	}
}

// Contains returns true if the set contains the specified reference, false otherwise.
func (set *PropertyReferenceSet) Contains(ref PropertyReference) bool {
	_, ok := set.references[ref]
	return ok
}

// Except returns a new PropertyReferenceSet containing only the items in this set that are not in the provided set
func (set *PropertyReferenceSet) Except(otherSet *PropertyReferenceSet) *PropertyReferenceSet {
	result := NewPropertyReferenceSet()
	for r := range set.references {
		if !otherSet.Contains(r) {
			result.Add(r)
		}
	}

	return result
}

func (set *PropertyReferenceSet) AsSlice() []PropertyReference {
	result := make([]PropertyReference, 0, len(set.references))
	for r := range set.references {
		result = append(result, r)
	}

	return result
}