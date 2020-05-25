/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
)

// ArrayType is used for fields that contain an array of values
type ArrayType struct {
	element Type
}

// NewArrayType creates a new array with elements of the specified type
func NewArrayType(element Type) *ArrayType {
	return &ArrayType{element}
}

// assert we implemented Type correctly
var _ Type = (*ArrayType)(nil)

// AsType renders the Go abstract syntax tree for an array type
func (array *ArrayType) AsType() ast.Expr {
	return &ast.ArrayType{
		Elt: array.element.AsType(),
	}
}

// RequiredImports returns a list of packages required by this
func (array *ArrayType) RequiredImports() []PackageReference {
	return array.element.RequiredImports()
}

// References this type has to the given type
func (array *ArrayType) References(d *TypeName) bool {
	return array.element.References(d)
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

// CreateInternalDefinitions invokes CreateInternalDefinitions on the inner 'element' type
func (array *ArrayType) CreateInternalDefinitions(name *TypeName, idFactory IdentifierFactory) (Type, []TypeDefiner) {
	newElementType, otherTypes := array.element.CreateInternalDefinitions(name, idFactory)
	return NewArrayType(newElementType), otherTypes
}

// CreateDefinitions defines a named type for this array type
func (array *ArrayType) CreateDefinitions(name *TypeName, _ IdentifierFactory, _ bool) (TypeDefiner, []TypeDefiner) {
	return NewSimpleTypeDefiner(name, array), nil
}
