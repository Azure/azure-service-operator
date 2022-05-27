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

// Set provides a standard way to have a set of distinct things
type Set[T comparable] map[T]struct{}

// Make creates a new set with the given values
func Make[T comparable](ts ...T) Set[T] {
	result := make(Set[T], len(ts))

	for _, x := range ts {
		result.Add(x)
	}

	return result
}

// Contains does a check to see if the provided value is in the set
func (set Set[T]) Contains(x T) bool {
	_, ok := set[x]
	return ok
}

// Add adds the provided value into the set
// Nothing happens if the value is already present
func (set Set[T]) Add(x T) {
	set[x] = struct{}{}
}

// Remove deletes the provided value from the set
// Nothing happens if the value is not present
func (set Set[T]) Remove(x T) {
	delete(set, x)
}

// Copy returns an independent copy of the set
func (set Set[T]) Copy() Set[T] {
	return maps.Clone(set)
}

// Clear removes all the itesm from this set
func (set Set[T]) Clear() {
	// TODO: Once the generics bug in Go 1.18.2 is fixed, revert to this implementation
	// maps.Clear(set)
	for k := range set {
		delete(set, k)
	}
}

/* compiler crashes at the moment: https://github.com/golang/go/issues/51840

func (set Set[T]) Equals(other Set[T]) bool {
	return maps.Equal(set, other)
}

*/

func AreEqual[T comparable](left, right Set[T]) bool {
	return maps.Equal(left, right)
}

// Values returns a slice of all the values in the set
func (set Set[T]) Values() []T {
	return maps.Keys(set)
}

// AsSortedSlice returns a sorted slice of values from this set
func AsSortedSlice[T constraints.Ordered](set Set[T]) []T {
	result := maps.Keys(set)
	slices.Sort(result)
	return result
}
