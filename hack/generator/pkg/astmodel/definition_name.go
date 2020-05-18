/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "go/ast"

// DefinitionName encapsulates all the information required to uniquely identify a definition
type DefinitionName struct {
	PackageReference
	name string
}

// NewDefinitionName is a factory method for creating a DefinitionName
func NewDefinitionName(pr PackageReference, name string) DefinitionName {
	return DefinitionName{pr, name}
}

// Name returns the unique name for this definition
func (dn *DefinitionName) Name() string {
	return dn.name
}

// A DefinitionName can be used as a Type,
// it is simply a reference to the name.
var _ Type = (*DefinitionName)(nil)

// AsType implements Type for DefinitionName
func (dn *DefinitionName) AsType() ast.Expr {
	return ast.NewIdent(dn.name)
}

// References indicates whether this Type includes any direct references to the given Type?
func (dn *DefinitionName) References(t Type) bool {
	return dn.Equals(t)
}

// RequiredImports returns all the imports required for this definition
func (dn *DefinitionName) RequiredImports() []PackageReference {
	return []PackageReference{dn.PackageReference}
}

// Equals returns true if the passed type references the same definition, false otherwise
func (dn *DefinitionName) Equals(t Type) bool {
	if d, ok := t.(*DefinitionName); ok {
		return dn.name == d.name && dn.PackageReference.Equals(&d.PackageReference)
	}

	return false
}
