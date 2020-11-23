/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	ast "github.com/dave/dst"
)

// PrimitiveType represents a Go primitive type
type PrimitiveType struct {
	name string
}

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
func (prim *PrimitiveType) AsType(_ *CodeGenerationContext) ast.Expr {
	return ast.NewIdent(prim.name)
}

func (prim *PrimitiveType) AsDeclarations(genContext *CodeGenerationContext, name TypeName, description []string) []ast.Decl {
	return AsSimpleDeclarations(genContext, name, description, prim)
}

// RequiredPackageReferences returns a list of package required by this
func (prim *PrimitiveType) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet()
}

// References always returns nil because primitive types don't refer to
// any other types.
func (prim *PrimitiveType) References() TypeNameSet {
	return nil
}

// Equals returns true if the passed type is another primitive type the same name, false otherwise
func (prim *PrimitiveType) Equals(t Type) bool {
	if p, ok := t.(*PrimitiveType); ok {
		return prim.name == p.name
	}

	return false
}

// Name returns the name of the primitive type
func (prim *PrimitiveType) Name() string {
	return prim.name
}

// String implements fmt.Stringer
func (prim *PrimitiveType) String() string {
	return prim.name
}
