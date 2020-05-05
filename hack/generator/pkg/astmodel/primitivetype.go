/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
)

// PrimitiveType represents a Go primitive type
type PrimitiveType struct {
	name string
}

// Enum represents an enumeration of predefined values
var Enum = &PrimitiveType{"ENUM"} // TODO

// IntType represents a Go integer type
var IntType = &PrimitiveType{"int"}

// StringType represents the Go string type
var StringType = &PrimitiveType{"string"}

// FloatType represents the Go float64 type
var FloatType = &PrimitiveType{"float64"}

// BoolType represents the Go bool type
var BoolType = &PrimitiveType{"bool"}

// AnyType represents the root Go interface type, permitting any object
var AnyType = &PrimitiveType{"interface{}"}

// assert that we implemented Type correctly
var _ Type = (*PrimitiveType)(nil)

// AsType implements Type for PrimitiveType returning an abstract syntax tree
func (prim *PrimitiveType) AsType() ast.Expr {
	return ast.NewIdent(prim.name)
}

func (prim *PrimitiveType) RequiredImports() []PackageReference {
	return nil
}
