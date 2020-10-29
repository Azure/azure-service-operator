/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"
)

// ArmType wraps an existing type to indicate that it is an ARM targeted variation
type ArmType struct {
	objectType ObjectType
}

// ArmType is a Type
var _ Type = &ArmType{}

// NewArmType wraps an object type to indicate it's an ARM targeted variation
func NewArmType(object ObjectType) *ArmType {
	return &ArmType{
		objectType: object,
	}
}

// RequiredPackageReferences returns a list of packages required by this type
func (at ArmType) RequiredPackageReferences() []PackageReference {
	return at.objectType.RequiredPackageReferences()
}

// References returns the names of all types that this type references
func (at *ArmType) References() TypeNameSet {
	return at.objectType.References()
}

// AsType renders as a Go abstract syntax tree for a type
func (at *ArmType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {
	return at.objectType.AsType(codeGenerationContext)
}

// AsDeclarations renders as a Go abstract syntax tree for a declaration
func (at *ArmType) AsDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description []string) []ast.Decl {
	return at.objectType.AsDeclarations(codeGenerationContext, name, description)
}

// Equals decides if the types are the same
func (at *ArmType) Equals(t Type) bool {
	if ost, ok := t.(*ArmType); ok {
		return TypeEquals(&at.objectType, ost)
	}

	return false
}

// String returns a string representation of our ARM type
func (at *ArmType) String() string {
	return fmt.Sprintf("ARM(%v)", at.objectType)
}

// ObjectType returns the underlying object type of the arm type
func (at *ArmType) ObjectType() ObjectType {
	return at.objectType
}
