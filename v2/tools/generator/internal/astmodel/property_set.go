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
	ForEach(func(def *PropertyDefinition))
	Len() int
	Copy() PropertySet
	First() *PropertyDefinition
	AsSlice() []*PropertyDefinition
	IsEmpty() bool
}

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

func (p PropertySet) IsEmpty() bool {
	return len(p) == 0
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
