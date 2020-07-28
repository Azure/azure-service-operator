/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
)

// ArrayType is used for properties that contain an array of values
type ArrayType struct {
	element Type
}

// NewArrayType creates a new array with elements of the specified type
func NewArrayType(element Type) *ArrayType {
	return &ArrayType{element}
}

// assert we implemented Type correctly
var _ Type = (*ArrayType)(nil)

func (array *ArrayType) AsDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description *string) []ast.Decl {
	return AsSimpleDeclarations(codeGenerationContext, name, description, array)
}

// AsType renders the Go abstract syntax tree for an array type
func (array *ArrayType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {
	return &ast.ArrayType{
		Elt: array.element.AsType(codeGenerationContext),
	}
}

// RequiredImports returns a list of packages required by this
func (array *ArrayType) RequiredImports() []PackageReference {
	return array.element.RequiredImports()
}

// References returns the references of the type this array contains.
func (array *ArrayType) References() TypeNameSet {
	return array.element.References()
}

// Equals returns true if the passed type is an array type with the same kind of elements, false otherwise
func (array *ArrayType) Equals(t Type) bool {
	if array == t {
		return true
	}

	if et, ok := t.(*ArrayType); ok {
		return array.element.Equals(et.element)
	}

	return false
}
