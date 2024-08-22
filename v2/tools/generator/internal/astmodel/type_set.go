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
	Single() (Type, bool)
}

var _ ReadonlyTypeSet = TypeSet{}

// MakeTypeSet makes a new TypeSet containing the given types
func MakeTypeSet(types ...Type) TypeSet {
	var result TypeSet
	result.types = make([]Type, 0, len(types))

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

// Remove removes the type from the set if it exists,
// and returns whether it was removed or not
func (ts *TypeSet) Remove(t Type) bool {
	for i, other := range ts.types {
		if t.Equals(other, EqualityOverrides{}) {
			ts.types = append(ts.types[:i], ts.types[i+1:]...)
			return true
		}
	}

	return false
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

// Len returns the number of items in the set
func (ts TypeSet) Len() int {
	return len(ts.types)
}

// Single returns the only item in the set, and true, if it contains exactly one; otherwise it returns nil and false.
func (ts TypeSet) Single() (Type, bool) {
	if len(ts.types) != 1 {
		return nil, false
	}

	return ts.types[0], true
}

// Copy returns a copy of this set that can then be modified.
func (ts TypeSet) Copy() TypeSet {
	return MakeTypeSet(ts.types...)
}

// AsSlice returns the content of this set as a slice
func (ts TypeSet) AsSlice() []Type {
	result := make([]Type, 0, len(ts.types))
	result = append(result, ts.types...)
	return result
}
