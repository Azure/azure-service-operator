/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"sort"

	"k8s.io/klog/v2"
)

// EnumType represents a set of mutually exclusive predefined options
type EnumType struct {
	// BaseType is the underlying type used to define the values
	BaseType *PrimitiveType
	// Options is the set of all unique values
	options []EnumValue
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
func (enum *EnumType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {
	// this should "never" happen as we name all enums; warn about it if it does
	klog.Warning("Emitting unnamed enum, something’s awry")
	return enum.BaseType.AsType(codeGenerationContext)
}

// References indicates whether this Type includes any direct references to the given Type
func (enum *EnumType) References(tn *TypeName) bool {
	return enum.BaseType.References(tn)
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
func (enum *EnumType) RequiredImports() []*PackageReference {
	return nil
}

// CreateInternalDefinitions defines a named type for this enum and returns that type to be used in place
// of this "raw" enum type
func (enum *EnumType) CreateInternalDefinitions(nameHint *TypeName, idFactory IdentifierFactory) (Type, []TypeDefiner) {
	// an internal enum must always be named:
	definedEnum, otherTypes := enum.CreateDefinitions(nameHint, idFactory, false)
	return definedEnum.Name(), append(otherTypes, definedEnum)
}

// CreateDefinitions defines a named type for this "raw" enum type
func (enum *EnumType) CreateDefinitions(name *TypeName, idFactory IdentifierFactory, _ bool) (TypeDefiner, []TypeDefiner) {
	identifier := idFactory.CreateEnumIdentifier(name.name)
	canonicalName := NewTypeName(name.PackageReference, identifier)
	return NewEnumDefinition(canonicalName, enum), nil
}

// Options returns all the enum options
// A copy of the slice is returned to preserve immutability
func (enum *EnumType) Options() []EnumValue {
	return append(enum.options[:0:0], enum.options...)
}
