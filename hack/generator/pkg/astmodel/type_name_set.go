/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "fmt"

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

// Remove removes the specified item if it is in the set. If it is not in
// the set this is a no-op.
func (ts TypeNameSet) Remove(val TypeName) TypeNameSet {
	if ts == nil {
		return ts
	}

	delete(ts, val)

	return ts
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

// Single returns the single TypeName in the set. This panics if there is not a single item in the set.
func (ts TypeNameSet) Single() TypeName {
	if len(ts) != 1 {
		panic(fmt.Sprintf("Single() cannot be called with %d types in the set", len(ts)))
	}

	for name := range ts {
		return name
	}

	panic("Reached unreachable code")
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
