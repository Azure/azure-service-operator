/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
)

// InternalTypeNameSet stores type names in no particular order without
// duplicates.
type InternalTypeNameSet map[InternalTypeName]struct{}

// NewInternalTypeNameSet makes a InternalTypeNameSet containing the specified
// names. If no elements are passed it might be nil.
func NewInternalTypeNameSet(initial ...InternalTypeName) InternalTypeNameSet {
	result := make(InternalTypeNameSet)
	for _, name := range initial {
		result.Add(name)
	}

	return result
}

// Add includes the passed name in the set
func (ts InternalTypeNameSet) Add(val InternalTypeName) {
	ts[val] = struct{}{}
}

// Contains returns whether this name is in the set. Works for nil
// sets too.
func (ts InternalTypeNameSet) Contains(val InternalTypeName) bool {
	if ts == nil {
		return false
	}
	_, found := ts[val]
	return found
}

// ContainsAll returns whether all the names are in the set. Works for nil
// sets too.
func (ts InternalTypeNameSet) ContainsAll(other InternalTypeNameSet) bool {
	if ts == nil {
		return false
	}
	for val := range other {
		if !ts.Contains(val) {
			return false
		}
	}
	return true
}

// ContainsAny returns whether any item of other is contained in the set. Works for nil
// sets too.
func (ts InternalTypeNameSet) ContainsAny(other InternalTypeNameSet) bool {
	if ts == nil {
		return false
	}
	for val := range other {
		if ts.Contains(val) {
			return true
		}
	}
	return false
}

// Remove removes the specified item if it is in the set. If it is not in the set this is a no-op.
func (ts InternalTypeNameSet) Remove(val InternalTypeName) {
	delete(ts, val)
}

func (ts InternalTypeNameSet) Equals(set InternalTypeNameSet) bool {
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

// AddAll adds the provided InternalTypeNameSet to the set
func (ts InternalTypeNameSet) AddAll(other InternalTypeNameSet) {
	for val := range other {
		ts[val] = struct{}{}
	}
}

// Single returns the single InternalTypeName in the set. This panics if there is not a single item in the set.
func (ts InternalTypeNameSet) Single() InternalTypeName {
	if len(ts) == 1 {
		for name := range ts {
			return name
		}
	}

	panic(fmt.Sprintf("Single() cannot be called with %d types in the set", len(ts)))
}

func (ts InternalTypeNameSet) Copy() InternalTypeNameSet {
	result := make(InternalTypeNameSet, len(ts))
	for k := range ts {
		result.Add(k)
	}

	return result
}
