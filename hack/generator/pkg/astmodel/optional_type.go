/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
)

// OptionalType is used for items that may or may not be present
type OptionalType struct {
	element Type
}

// NewOptionalType creates a new optional type that may or may not have the specified 'element' type
func NewOptionalType(element Type) *OptionalType {
	return &OptionalType{element}
}

// assert we implemented Type correctly
var _ Type = (*OptionalType)(nil)

// AsType renders the Go abstract syntax tree for an optional type
func (optional *OptionalType) AsType() ast.Expr {
	return &ast.StarExpr{
		X: optional.element.AsType(),
	}
}

// RequiredImports returns the imports required by the 'element' type
func (optional *OptionalType) RequiredImports() []PackageReference {
	return optional.element.RequiredImports()
}

// References is true if it is this type or the 'element' type references it
func (optional *OptionalType) References(d *DefinitionName) bool {
	return optional.element.References(d)
}

// Equals returns true if this type is equal to the other type
func (optional *OptionalType) Equals(t Type) bool {
	if optional == t {
		return true // reference equality short-cut
	}

	if otherOptional, ok := t.(*OptionalType); ok {
		return optional.element.Equals(otherOptional.element)
	}

	return false
}

// CreateRelatedDefinitions returns any additional definitions that need to be created
func (optional *OptionalType) CreateRelatedDefinitions(ref PackageReference, namehint string, idFactory IdentifierFactory) []Definition {
	return optional.element.CreateRelatedDefinitions(ref, namehint, idFactory)
}

