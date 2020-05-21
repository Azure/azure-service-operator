/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"sort"
)

// EnumType represents a set of mutually exclusive predefined options
type EnumType struct {
	// BaseType is the underlying type used to define the values
	BaseType *PrimitiveType
	// Options is the set of all unique values
	options []EnumValue
	// canonicalName is our actual name, only available once generated, assigned by CreateRelatedDefinitions()
	canonicalName DefinitionName
}

// EnumType must implement the Type interface correctly
var _ Type = (*EnumType)(nil)

// NewEnumType defines a new enumeration including the legal values
func NewEnumType(baseType *PrimitiveType, options []EnumValue) *EnumType {
	sort.Slice(options, func(left int, right int) bool {
		return options[left].Identifier < options[right].Identifier
	})

	return &EnumType{BaseType: baseType, options: options}
}

// AsType implements Type for EnumType
func (enum *EnumType) AsType() ast.Expr {
	return ast.NewIdent(enum.canonicalName.name)
}

// References indicates whether this Type includes any direct references to the given Type?
func (enum *EnumType) References(d *DefinitionName) bool {
	return enum.canonicalName.References(d)
}

// CreateRelatedDefinitions returns a definition for our enumeration, with a name based on the referencing property
func (enum *EnumType) CreateRelatedDefinitions(ref PackageReference, namehint string, idFactory IdentifierFactory) []Definition {
	identifier := idFactory.CreateEnumIdentifier(namehint)
	enum.canonicalName = DefinitionName{PackageReference: ref, name: identifier}
	definition := NewEnumDefinition(enum.canonicalName, enum)
	return []Definition{definition}
}

// Equals will return true if the supplied type has the same base type and options
func (enum *EnumType) Equals(t Type) bool {
	if e, ok := t.(*EnumType); ok {
		if !enum.BaseType.Equals(e.BaseType) {
			return false
		}

		if len(enum.options) != len(e.options) {
			// Different number of fields, not equal
			return false
		}

		for i := range enum.options {
			if !enum.options[i].Equals(&e.options[i]) {
				return false
			}
		}

		// All options match, equal
		return true
	}

	return false
}

// RequiredImports indicates that Enums never need additional imports
func (enum *EnumType) RequiredImports() []PackageReference {
	return nil
}

// Options returns all our options
// A copy of the slice is returned to preserve immutability
func (enum *EnumType) Options() []EnumValue {
	return append(enum.options[:0:0], enum.options...)
}
