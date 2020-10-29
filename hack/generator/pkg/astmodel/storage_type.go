/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"
)

// StorageType wraps an existing type to indicate that it is a storage focussed variation
type StorageType struct {
	objectType ObjectType
}

func (st *StorageType) String() string {
	return fmt.Sprintf("Storage(%v)", st.objectType)
}

// StorageType is a Type
var _ Type = &StorageType{}

// NewStorageType wraps an object type to indicate it's a dedicated storage version
func NewStorageType(objectType ObjectType) *StorageType {
	return &StorageType{
		objectType: objectType,
	}
}

// IsStorageType returns true if the passed type is a storage type; false otherwise.
func IsStorageType(t Type) bool {
	_, ok := t.(*StorageType)
	return ok
}

// IsStorageDefinition returns true if the passed definition is for a storage type; false otherwise.
func IsStorageDefinition(definition TypeDefinition) bool {
	return IsStorageType(definition.Type())
}

// RequiredImports returns a list of packages required by this type
func (st *StorageType) RequiredPackageReferences() []PackageReference {
	return st.objectType.RequiredPackageReferences()
}

// References returns the names of all types that this type references
func (st *StorageType) References() TypeNameSet {
	return st.objectType.References()
}

// AsType renders as a Go abstract syntax tree for a type
func (st *StorageType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {
	return st.objectType.AsType(codeGenerationContext)
}

// AsDeclarations renders as a Go abstract syntax tree for a declaration
func (st *StorageType) AsDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description []string) []ast.Decl {
	return st.objectType.AsDeclarations(codeGenerationContext, name, description)
}

// Equals decides if the types are the same
func (st *StorageType) Equals(t Type) bool {
	if other, ok := t.(*StorageType); ok {
		return TypeEquals(&st.objectType, other)
	}

	return false
}
