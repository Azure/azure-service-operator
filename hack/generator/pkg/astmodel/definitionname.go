/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "go/ast"

// DefinitionName is a reference to a name
type DefinitionName struct {
	PackageReference
	name string
}

// NewDefinitionName creates a new definition name
func NewDefinitionName(pr PackageReference, name string) DefinitionName {
	return DefinitionName{pr, name}
}

// Name returns the name
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

// References this type has to the given type
func (dn *DefinitionName) References(t Type) bool {
	return dn.Equals(t)
}

// RequiredImports returns a list of packages required by this
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
