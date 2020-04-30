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

// AsType renders the Go abstract syntax tree for an array type
func (array *ArrayType) AsType() ast.Expr {
	return &ast.ArrayType{
		Elt: array.element.AsType(),
	}
}
