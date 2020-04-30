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

// AsType implements Type for MapType to create the abstract syntax tree for a map
func (m *MapType) AsType() ast.Expr {
	return &ast.MapType{
		Key:   m.key.AsType(),
		Value: m.value.AsType(),
	}
}
