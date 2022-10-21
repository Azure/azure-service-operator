/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"sort"

	"golang.org/x/exp/maps"
)

// PropertySet wraps a set of property definitions, indexed by name, along with some convenience methods
type PropertySet map[PropertyName]*PropertyDefinition

type ReadOnlyPropertySet interface {
	AsSlice() []*PropertyDefinition
	ContainsProperty(name PropertyName) bool
	Get(name PropertyName) (*PropertyDefinition, bool)
	Find(predicate func(*PropertyDefinition) bool) (*PropertyDefinition, bool)
	FindAll(predicate func(*PropertyDefinition) bool) []*PropertyDefinition
	Copy() PropertySet
	Equals(other ReadOnlyPropertySet, overrides EqualityOverrides) bool
	First() *PropertyDefinition
	ForEach(func(def *PropertyDefinition))
	IsEmpty() bool
	Len() int
}

var _ ReadOnlyPropertySet = PropertySet{} // assert satisfies

// NewPropertySet creates a new set of properties
func NewPropertySet(properties ...*PropertyDefinition) PropertySet {
	result := make(PropertySet, len(properties))
	for _, prop := range properties {
		result[prop.PropertyName()] = prop
	}

	return result
}

func (p PropertySet) First() *PropertyDefinition {
	for _, v := range p {
		return v
	}

	return nil
}

func (p PropertySet) ContainsProperty(name PropertyName) bool {
	_, ok := p[name]
	return ok
}

func (p PropertySet) Find(predicate func(*PropertyDefinition) bool) (*PropertyDefinition, bool) {
	for _, prop := range p {
		if predicate(prop) {
			return prop, true
		}
	}

	return nil, false
}

func (p PropertySet) FindAll(predicate func(*PropertyDefinition) bool) []*PropertyDefinition {
	var result []*PropertyDefinition
	for _, prop := range p {
		if predicate(prop) {
			result = append(result, prop)
		}
	}

	return result
}

func (p PropertySet) Get(name PropertyName) (*PropertyDefinition, bool) {
	result, ok := p[name]
	return result, ok
}

func (p PropertySet) IsEmpty() bool {
	return len(p) == 0
}

func (p PropertySet) Equals(other ReadOnlyPropertySet, overrides EqualityOverrides) bool {
	if p.Len() != other.Len() {
		return false
	}

	for k, v := range p {
		v2, ok := other.Get(k)
		if !ok || !v.Equals(v2, overrides) {
			return false
		}
	}

	return true
}

func (p PropertySet) ForEach(f func(*PropertyDefinition)) {
	for _, v := range p {
		f(v)
	}
}

func (p PropertySet) Len() int {
	return len(p)
}

// AsSlice returns all the properties in a slice, sorted alphabetically by name
func (p PropertySet) AsSlice() []*PropertyDefinition {
	result := maps.Values(p)

	// Sort it so that it's always consistent
	sort.Slice(result, func(left int, right int) bool {
		return result[left].propertyName < result[right].propertyName
	})

	return result
}

// Add updates the set by including the provided property
// Any existing definition by that name will be overwritten if present
func (p PropertySet) Add(property *PropertyDefinition) {
	p[property.propertyName] = property
}

// Copy returns a new property set with the same properties as this one
func (p PropertySet) Copy() PropertySet {
	result := make(PropertySet, len(p))
	for name, prop := range p {
		result[name] = prop
	}

	return result
}
