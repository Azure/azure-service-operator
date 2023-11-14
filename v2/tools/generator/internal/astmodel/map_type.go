/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
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

func (m *MapType) WithKeyType(t Type) *MapType {
	if TypeEquals(m.key, t) {
		return m
	}

	result := *m
	result.key = t
	return &result
}

func (m *MapType) WithValueType(t Type) *MapType {
	if TypeEquals(m.value, t) {
		return m
	}

	result := *m
	result.value = t
	return &result
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

func (m *MapType) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) ([]dst.Decl, error) {
	return AsSimpleDeclarations(codeGenerationContext, declContext, m), nil
}

// AsType implements Type for MapType to create the abstract syntax tree for a map
func (m *MapType) AsType(codeGenerationContext *CodeGenerationContext) dst.Expr {
	return &dst.MapType{
		Key:   m.key.AsType(codeGenerationContext),
		Value: m.value.AsType(codeGenerationContext),
	}
}

// AsZero renders an expression for the "zero" value of a map by calling make()
func (m *MapType) AsZero(_ TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	return astbuilder.Nil()
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
func (m *MapType) Equals(t Type, overrides EqualityOverrides) bool {
	if m == t {
		return true // short-circuit
	}

	if mt, ok := t.(*MapType); ok {
		return m.key.Equals(mt.key, overrides) &&
			m.value.Equals(mt.value, overrides)
	}

	return false
}

// String implements fmt.Stringer
func (m *MapType) String() string {
	return fmt.Sprintf("map[%s]%s", m.key.String(), m.value.String())
}

// WriteDebugDescription adds a description of the current type to the passed builder
// builder receives the full description, including nested types
// definitions is for resolving named types
func (m *MapType) WriteDebugDescription(builder *strings.Builder, currentPackage InternalPackageReference) {
	if m == nil {
		builder.WriteString("<nilMap>")
		return
	}

	builder.WriteString("Map[")
	m.key.WriteDebugDescription(builder, currentPackage)
	builder.WriteString("]")
	m.value.WriteDebugDescription(builder, currentPackage)
}
