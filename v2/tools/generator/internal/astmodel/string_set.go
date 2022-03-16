/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"sort"
)

// StringSet provides a standard way to have a set of distinct strings with no sort applied
type StringSet map[string]struct{}

// MakeStringSet creates a new set with the supplied strings
func MakeStringSet(strings ...string) StringSet {
	result := make(StringSet)

	for _, s := range strings {
		result.Add(s)
	}

	return result
}

// Contains does a case-sensitive check to see if the provided string is in the set
func (set StringSet) Contains(s string) bool {
	_, ok := set[s]
	return ok
}

// Add adds the provided string into the set
// Nothing happens if the string is already present
func (set StringSet) Add(s string) {
	set[s] = struct{}{}
}

// Remove deletes the provided string from the set
// Nothing happens if the string is not present
func (set StringSet) Remove(s string) {
	delete(set, s)
}

// Copy returns an independent copy of the set
func (set StringSet) Copy() StringSet {
	result := MakeStringSet()
	for s := range set {
		result.Add(s)
	}

	return result
}

// AsSlice returns a sorted slice of strings from this set
func (set StringSet) AsSortedSlice() []string {
	result := make([]string, 0, len(set))
	for s := range set {
		result = append(result, s)
	}

	sort.Strings(result)
	return result
}
