/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package set

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type StringSet = Set[string]

// Set provides a standard way to have a set of distinct things
type Set[t comparable] map[t]struct{}

// MakeStringSet creates a new set with the supplied strings
func MakeStringSet(strings ...string) StringSet {
	return Make(strings...)
}

// Make creates a new set with the given values
func Make[t comparable](ts ...t) Set[t] {
	result := make(Set[t], len(ts))

	for _, x := range ts {
		result.Add(x)
	}

	return result
}

// Contains does a check to see if the provided value is in the set
func (set Set[t]) Contains(x t) bool {
	_, ok := set[x]
	return ok
}

// Add adds the provided value into the set
// Nothing happens if the value is already present
func (set Set[t]) Add(x t) {
	set[x] = struct{}{}
}

// Remove deletes the provided value from the set
// Nothing happens if the value is not present
func (set Set[t]) Remove(x t) {
	delete(set, x)
}

// Copy returns an independent copy of the set
func (set Set[t]) Copy() Set[t] {
	return maps.Clone(set)
}

func (set Set[t]) Equals(other Set[t]) bool {
	return maps.Equal(set, other)
}

// Values returns a slice of all the values in the set
func (set Set[t]) Values() []t {
	return maps.Keys(set)
}

// AsSortedSlice returns a sorted slice of values from this set
func AsSortedSlice[t constraints.Ordered](set Set[t]) []t {
	result := maps.Keys(set)
	slices.Sort(result)
	return result
}
