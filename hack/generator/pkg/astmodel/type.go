/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/dave/dst"
)

// Type represents something that is a Go type
type Type interface {
	// RequiredPackageReferences returns a set of packages imports required by this type
	RequiredPackageReferences() *PackageReferenceSet

	// References returns the names of all types that this type
	// references. For example, an Array of Persons references a
	// Person.
	References() TypeNameSet

	// AsType renders as a Go abstract syntax tree for a type
	// (yes this says dst.Expr but that is what the Go 'dst' package uses for types)
	AsType(codeGenerationContext *CodeGenerationContext) dst.Expr

	// AsDeclarations renders as a Go abstract syntax tree for a declaration
	AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl

	// AsZero renders an expression for the "zero" value of the type
	// types allows TypeName to resolve to the underlying type
	// ctx allows current imports to be correctly identified where needed
	AsZero(types Types, ctx *CodeGenerationContext) dst.Expr

	// Equals returns true if the passed type is the same as this one, false otherwise
	Equals(t Type) bool

	// Make sure all Types have a printable version for debugging/user info.
	// This doesn't need to be a full representation of the type.
	fmt.Stringer
}

// IgnoringErrors returns the type stripped of any ErroredType wrapper
func IgnoringErrors(t Type) Type {
	if errored, ok := t.(*ErroredType); ok {
		return errored.InnerType()
	}

	return t
}

// DeclarationContext represents some metadata about a specific declaration
type DeclarationContext struct {
	Name        TypeName
	Description []string
	Validations []KubeBuilderValidation
}

// TypeEquals decides if the types are the same and handles the `nil` case
func TypeEquals(left, right Type) bool {
	if left == nil {
		return right == nil
	}

	return left.Equals(right)
}
