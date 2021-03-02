/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// TypeNameSet stores type names in no particular order without
// duplicates.
type TypeNameSet map[TypeName]struct{}

// NewTypeNameSet makes a TypeNameSet containing the specified
// names. If no elements are passed it might be nil.
func NewTypeNameSet(initial ...TypeName) TypeNameSet {
	var result TypeNameSet
	for _, name := range initial {
		result = result.Add(name)
	}
	return result
}

// Add includes the passed name in the set and returns the updated
// set, so that adding can work for a nil set - this makes it more
// convenient to add to sets kept in a map (in the way you might with
// a map of slices).
func (ts TypeNameSet) Add(val TypeName) TypeNameSet {
	if ts == nil {
		ts = make(TypeNameSet)
	}
	ts[val] = struct{}{}
	return ts
}

// Contains returns whether this name is in the set. Works for nil
// sets too.
func (ts TypeNameSet) Contains(val TypeName) bool {
	if ts == nil {
		return false
	}
	_, found := ts[val]
	return found
}

func (ts TypeNameSet) Equals(set TypeNameSet) bool {
	if len(ts) != len(set) {
		// Different sizes, not equal
		return false
	}

	for k := range set {
		if _, ok := ts[k]; !ok {
			// Missing key, not equal
			return false
		}
	}

	return true
}

// AddAll adds the provided TypeNameSet to the set
func (ts TypeNameSet) AddAll(other TypeNameSet) TypeNameSet {
	if ts == nil {
		ts = make(TypeNameSet)
	}

	for val := range other {
		ts[val] = struct{}{}
	}

	return ts
}

// SetUnion returns a new set with all of the names in s1 or s2.
func SetUnion(s1, s2 TypeNameSet) TypeNameSet {
	var result TypeNameSet
	for val := range s1 {
		result = result.Add(val)
	}
	for val := range s2 {
		result = result.Add(val)
	}
	return result
}
