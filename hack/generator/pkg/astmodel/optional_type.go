/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	ast "github.com/dave/dst"
)

// OptionalType is used for items that may or may not be present
type OptionalType struct {
	element Type
}

// NewOptionalType creates a new optional type that may or may not have the specified 'element' type
func NewOptionalType(element Type) Type {
	if isTypeOptional(element) {
		return element
	}

	return &OptionalType{element}
}

func isTypeOptional(t Type) bool {
	// Arrays and Maps are already "optional" as far as Go is concerned,
	// so don't wrap them. Optional is also obviously already optional.
	switch t.(type) {
	case *ArrayType:
		return true
	case *MapType:
		return true
	case *OptionalType:
		return true
	default:
		return false
	}
}

// Element returns the type which is optional
func (optional *OptionalType) Element() Type {
	return optional.element
}

// assert we implemented Type correctly
var _ Type = (*OptionalType)(nil)

func (optional *OptionalType) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []ast.Decl {
	return AsSimpleDeclarations(codeGenerationContext, declContext, optional)
}

// AsType renders the Go abstract syntax tree for an optional type
func (optional *OptionalType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {
	// Special case interface{} as it shouldn't be a pointer
	if optional.element == AnyType {
		return optional.element.AsType(codeGenerationContext)
	}

	return &ast.StarExpr{
		X: optional.element.AsType(codeGenerationContext),
	}
}

// RequiredPackageReferences returns the imports required by the 'element' type
func (optional *OptionalType) RequiredPackageReferences() *PackageReferenceSet {
	return optional.element.RequiredPackageReferences()
}

// References returns the set of types that the underlying type refers to directly.
func (optional *OptionalType) References() TypeNameSet {
	return optional.element.References()
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

// BaseType returns the underlying type
func (optional *OptionalType) BaseType() Type {
	return optional.element
}

// String implements fmt.Stringer
func (optional *OptionalType) String() string {
	return fmt.Sprintf("(optional: %s)", optional.element.String())
}
