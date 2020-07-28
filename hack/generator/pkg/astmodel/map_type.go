/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
)

// MapType is used to define properties that contain additional property values
type MapType struct {
	key   Type
	value Type
}

// KeyType returns the type of keys in the type represented by this MapType
func (m *MapType) KeyType() Type {
	return m.key
}

// ValueType returns the type of values in the type represented by this MapType
func (m *MapType) ValueType() Type {
	return m.value
}

// NewMapType creates a new map with the specified key and value types
func NewMapType(key Type, value Type) *MapType {
	return &MapType{key, value}
}

// NewStringMapType creates a new map with string keys and the specified value type
func NewStringMapType(value Type) *MapType {
	return NewMapType(StringType, value)
}

// assert that we implemented Type correctly
var _ Type = (*MapType)(nil)

func (m *MapType) AsDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description *string) []ast.Decl {
	return AsSimpleDeclarations(codeGenerationContext, name, description, m)
}

// AsType implements Type for MapType to create the abstract syntax tree for a map
func (m *MapType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {
	return &ast.MapType{
		Key:   m.key.AsType(codeGenerationContext),
		Value: m.value.AsType(codeGenerationContext),
	}
}

// RequiredImports returns a list of packages required by this
func (m *MapType) RequiredImports() []PackageReference {
	var result []PackageReference
	result = append(result, m.key.RequiredImports()...)
	result = append(result, m.value.RequiredImports()...)
	return result
}

// References returns all of the types referenced by either the the key or value types.
func (m *MapType) References() TypeNameSet {
	return SetUnion(m.key.References(), m.value.References())
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
