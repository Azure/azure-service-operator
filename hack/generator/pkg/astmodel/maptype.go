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

func (m *MapType) RequiredImports() []PackageReference {
	var result []PackageReference
	result = append(result, m.key.RequiredImports()...)
	result = append(result, m.value.RequiredImports()...)
	return result
}

func (m *MapType) References(t Type) bool {
	return m == t || m.key.References(t) || m.value.References(t)
}