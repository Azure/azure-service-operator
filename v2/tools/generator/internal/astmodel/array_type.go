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

// ArrayType is used for properties that contain an array of values
type ArrayType struct {
	element Type
}

// NewArrayType creates a new array with elements of the specified type
func NewArrayType(element Type) *ArrayType {
	return &ArrayType{element}
}

// Element returns the element type of the array
func (array *ArrayType) Element() Type {
	return array.element
}

// WithElement returns an ArrayType with the specified element.
// the benefit of this is it allows reusing the same value if the
// element type is the same
func (array *ArrayType) WithElement(t Type) *ArrayType {
	if TypeEquals(array.element, t) {
		return array
	}

	result := *array
	result.element = t
	return &result
}

// assert we implemented Type correctly
var _ Type = (*ArrayType)(nil)

func (array *ArrayType) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	return AsSimpleDeclarations(codeGenerationContext, declContext, array)
}

// AsType renders the Go abstract syntax tree for an array type
func (array *ArrayType) AsType(codeGenerationContext *CodeGenerationContext) dst.Expr {
	return &dst.ArrayType{
		Elt: array.element.AsType(codeGenerationContext),
	}
}

// AsZero renders an expression for the "zero" value of the array by calling make()
func (array *ArrayType) AsZero(_ TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	return astbuilder.Nil()
}

// RequiredPackageReferences returns a list of packages required by this
func (array *ArrayType) RequiredPackageReferences() *PackageReferenceSet {
	return array.element.RequiredPackageReferences()
}

// References returns the references of the type this array contains.
func (array *ArrayType) References() TypeNameSet[TypeName] {
	return array.element.References()
}

// Equals returns true if the passed type is an array type with the same kind of elements, false otherwise
func (array *ArrayType) Equals(t Type, overrides EqualityOverrides) bool {
	if array == t {
		return true
	}

	if et, ok := t.(*ArrayType); ok {
		return array.element.Equals(et.element, overrides)
	}

	return false
}

// String implements fmt.Stringer
func (array *ArrayType) String() string {
	return fmt.Sprintf("[]%s", array.element.String())
}

// WriteDebugDescription adds a description of the current array type to the passed builder
// builder receives the full description, including nested types
// definitions is a dictionary for resolving named types
func (array *ArrayType) WriteDebugDescription(builder *strings.Builder, currentPackage PackageReference) {
	if array == nil {
		builder.WriteString("<nilArray>")
		return
	}

	builder.WriteString("Array[")
	array.element.WriteDebugDescription(builder, currentPackage)
	builder.WriteString("]")
}
