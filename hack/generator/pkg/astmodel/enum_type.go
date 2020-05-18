/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "go/ast"

// EnumType represents a set of mutually exclusive predefined options
type EnumType struct {
	DefinitionName
	// BaseType is the underlying type used to define the values
	BaseType *PrimitiveType
	// Options is the set of all unique values
	Options []EnumValue
}

// EnumValue captures a single value of the enumeration
type EnumValue struct {
	// Identifer is a Go identifer for the value
	Identifier string
	// Value is the actual value expected by ARM
	Value string
}

// EnumType must implement the Type interface correctly
var _ Type = (*EnumType)(nil)

// EnumType must implement the HasRelatedDefinitions interface correctly
var _ HasRelatedDefinitions = (*EnumType)(nil)

// NewEnumType defines a new enumeration including the legal values
func NewEnumType(baseType *PrimitiveType, options []EnumValue) *EnumType {
	return &EnumType{BaseType: baseType, Options: options}
}

// AsType implements Type for EnumType
func (enum *EnumType) AsType() ast.Expr {
	return ast.NewIdent(enum.name)
}

// References indicates whether this Type includes any direct references to the given Type?
func (enum *EnumType) References(t Type) bool {
	return enum.DefinitionName.References(t)
}

// RelatedDefinitions implements the HasRelatedDefinitions interface for EnumType
func (enum *EnumType) RelatedDefinitions(ref PackageReference, namehint string, idFactory IdentifierFactory) []Definition {
	identifier := idFactory.CreateEnumIdentifier(namehint)
	enum.DefinitionName = DefinitionName{PackageReference: ref, name: identifier}
	var definition Definition = &EnumDefinition{EnumType: *enum}
	return []Definition{definition}
}
