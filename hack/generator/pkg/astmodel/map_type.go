/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"github.com/dave/dst"
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

func (m *MapType) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	return AsSimpleDeclarations(codeGenerationContext, declContext, m)
}

// AsType implements Type for MapType to create the abstract syntax tree for a map
func (m *MapType) AsType(codeGenerationContext *CodeGenerationContext) dst.Expr {
	return &dst.MapType{
		Key:   m.key.AsType(codeGenerationContext),
		Value: m.value.AsType(codeGenerationContext),
	}
}

// AsZero renders an expression for the "zero" value of a map by calling make()
func (m *MapType) AsZero(_ Types, ctx *CodeGenerationContext) dst.Expr {
	return dst.NewIdent("nil")
}

// RequiredPackageReferences returns a list of packages required by this
func (m *MapType) RequiredPackageReferences() *PackageReferenceSet {
	result := NewPackageReferenceSet()
	result.Merge(m.key.RequiredPackageReferences())
	result.Merge(m.value.RequiredPackageReferences())
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
		return m.key.Equals(mt.key) && m.value.Equals(mt.value)
	}

	return false
}

// String implements fmt.Stringer
func (m *MapType) String() string {
	return fmt.Sprintf("map[%s]%s", m.key.String(), m.value.String())
}
