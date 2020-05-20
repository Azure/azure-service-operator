/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
)

// MapType is used to define fields that contain additional property values
type MapType struct {
	key   Type
	value Type
}

// NewMap creates a new map with the specified key and value types
func NewMap(key Type, value Type) *MapType {
	return &MapType{key, value}
}

// NewStringMap creates a new map with string keys and the specified value type
func NewStringMap(value Type) *MapType {
	return NewMap(StringType, value)
}

// assert that we implemented Type correctly
var _ Type = (*MapType)(nil)

// AsType implements Type for MapType to create the abstract syntax tree for a map
func (m *MapType) AsType() ast.Expr {
	return &ast.MapType{
		Key:   m.key.AsType(),
		Value: m.value.AsType(),
	}
}

// RequiredImports returns a list of packages required by this
func (m *MapType) RequiredImports() []PackageReference {
	var result []PackageReference
	result = append(result, m.key.RequiredImports()...)
	result = append(result, m.value.RequiredImports()...)
	return result
}

// References this type has to the given type
func (m *MapType) References(d *DefinitionName) bool {
	return m.key.References(d) || m.value.References(d)
}

// Equals returns true if the passed type is a map type with the same kinds of keys and elements, false otherwise
func (m *MapType) Equals(t Type) bool {
	if m == t {
		return true
	}

	if mt, ok := t.(*MapType); ok {
		return (m.key.Equals(mt.key) && m.value.Equals(mt.value))
	}

	return false
}

// CreateRelatedDefinitions returns any definitions required by our key or value types
func (m *MapType) CreateRelatedDefinitions(ref PackageReference, namehint string, idFactory IdentifierFactory) []Definition {
	result := m.key.CreateRelatedDefinitions(ref, namehint, idFactory)

	otherDefinitions := m.value.CreateRelatedDefinitions(ref, namehint, idFactory)
	result = append(result, otherDefinitions...)

	return result
}
