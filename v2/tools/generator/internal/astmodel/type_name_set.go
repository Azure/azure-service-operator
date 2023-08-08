/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
)

type TypeNameSet[N comparableTypeName] map[N]struct{}

type comparableTypeName interface {
	comparable
	TypeName
}

// NewTypeNameSet makes a TypeNameSet containing the specified
// names. If no elements are passed it might be nil.
func NewTypeNameSet[N comparableTypeName](initial ...N) TypeNameSet[N] {
	result := make(TypeNameSet[N])
	for _, name := range initial {
		result.Add(name)
	}

	return result
}

// Add includes the passed name in the set
func (ts TypeNameSet[N]) Add(val N) {
	ts[val] = struct{}{}
}

// Contains returns whether this name is in the set. Works for nil
// sets too.
func (ts TypeNameSet[N]) Contains(val N) bool {
	if ts == nil {
		return false
	}
	_, found := ts[val]
	return found
}

// ContainsAll returns whether all the names are in the set. Works for nil
// sets too.
func (ts TypeNameSet[N]) ContainsAll(other TypeNameSet[N]) bool {
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
func (ts TypeNameSet[N]) ContainsAny(other TypeNameSet[N]) bool {
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
func (ts TypeNameSet[N]) Remove(val N) {
	delete(ts, val)
}

func (ts TypeNameSet[N]) Equals(set TypeNameSet[N]) bool {
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
func (ts TypeNameSet[N]) AddAll(other TypeNameSet[N]) {
	for val := range other {
		ts[val] = struct{}{}
	}
}

// Single returns the single TypeName in the set. This panics if there is not a single item in the set.
func (ts TypeNameSet[N]) Single() TypeName {
	if len(ts) == 1 {
		for name := range ts {
			return name
		}
	}

	panic(fmt.Sprintf("Single() cannot be called with %d types in the set", len(ts)))
}

func (ts TypeNameSet[N]) Copy() TypeNameSet[N] {
	result := make(TypeNameSet[N], len(ts))
	for k := range ts {
		result.Add(k)
	}

	return result
}

// SetUnion returns a new set with all of the names in s1 or s2.
func SetUnion[N comparableTypeName](s1 TypeNameSet[N], s2 TypeNameSet[N]) TypeNameSet[N] {
	result := NewTypeNameSet[N]()
	for val := range s1 {
		result.Add(val)
	}
	for val := range s2 {
		result.Add(val)
	}
	return result
}
