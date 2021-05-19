/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/dave/dst"
	"strings"
)

// PrimitiveType represents a Go primitive type
type PrimitiveType struct {
	name string
	zero string
}

// IntType represents a Go integer type
var IntType = &PrimitiveType{"int", "0"}

// UInt64Type represents a Go uint64 type
var UInt64Type = &PrimitiveType{"uint64", "0"}

// UInt32Type represents a Go uint32 type
var UInt32Type = &PrimitiveType{"uint32", "0"}

// StringType represents the Go string type
var StringType = &PrimitiveType{"string", "\"\""}

// FloatType represents the Go float64 type
var FloatType = &PrimitiveType{"float64", "0"}

// BoolType represents the Go bool type
var BoolType = &PrimitiveType{"bool", "false"}

// AnyType represents the root Go interface type, permitting any object
var AnyType = &PrimitiveType{"interface{}", "nil"}

// assert that we implemented Type correctly
var _ Type = (*PrimitiveType)(nil)

// AsType implements Type for PrimitiveType returning an abstract syntax tree
func (prim *PrimitiveType) AsType(_ *CodeGenerationContext) dst.Expr {
	return dst.NewIdent(prim.name)
}

func (prim *PrimitiveType) AsDeclarations(genContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	return AsSimpleDeclarations(genContext, declContext, prim)
}

// RequiredPackageReferences returns a list of package required by this
func (prim *PrimitiveType) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet()
}

// AsZero renders an expression for the "zero" value of the type
// types allows TypeName to resolve to the underlying type
// ctx allows current imports to be correctly identified where needed
func (prim *PrimitiveType) AsZero(types Types, ctx *CodeGenerationContext) dst.Expr {
	return &dst.BasicLit{
		Value: prim.zero,
	}
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

// WriteDebugDescription adds a description of this primitive type to the passed builder
func (prim *PrimitiveType) WriteDebugDescription(builder *strings.Builder, _ Types) {
	builder.WriteString(prim.name)
}
