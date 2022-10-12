/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// TypeSet represents a set of types
type TypeSet struct {
	types []Type
}

// ReadonlyTypeSet is a readonly version of the TypeSet API.
type ReadonlyTypeSet interface {
	Contains(Type, EqualityOverrides) bool
	ForEach(func(t Type, ix int))
	ForEachError(func(t Type, ix int) error) error
	Len() int
}

var _ ReadonlyTypeSet = TypeSet{}

// MakeTypeSet makes a new TypeSet containing the given types
func MakeTypeSet(types ...Type) TypeSet {
	var result TypeSet
	for _, t := range types {
		result.Add(t)
	}

	return result
}

// ForEach executes the action for each type in the set
// this works around not having `range` on custom types
func (ts TypeSet) ForEach(action func(t Type, ix int)) {
	for ix, t := range ts.types {
		action(t, ix)
	}
}

// ForEachError executes the action for each type in the set,
// with the possibility to fail. This works around not having
// `range` on custom types.
func (ts TypeSet) ForEachError(action func(t Type, ix int) error) error {
	for ix, t := range ts.types {
		if err := action(t, ix); err != nil {
			return err
		}
	}

	return nil
}

// Add adds the type to the set if it does not already exist
// and returns if it was added or not
func (ts *TypeSet) Add(t Type) bool {
	if ts.Contains(t, EqualityOverrides{}) {
		return false
	}

	ts.types = append(ts.types, t)
	return true
}

// Contains checks if the set already contains the type
func (ts TypeSet) Contains(t Type, overrides EqualityOverrides) bool {
	// this is slow, but what can you do?
	for _, other := range ts.types {
		if t.Equals(other, overrides) {
			return true
		}
	}

	return false
}

// Equals returns true if both sets contain the same types
func (ts TypeSet) Equals(other TypeSet, overrides ...EqualityOverrides) bool {
	if len(ts.types) != len(other.types) {
		return false
	}

	override := EqualityOverrides{}
	if len(overrides) > 0 {
		if len(overrides) > 1 {
			panic("can only pass one EqualityOverrides")
		}

		override = overrides[0]
	}

	for _, t := range ts.types {
		if !other.Contains(t, override) {
			return false
		}
	}

	return true
}

func (ts TypeSet) Len() int {
	return len(ts.types)
}
