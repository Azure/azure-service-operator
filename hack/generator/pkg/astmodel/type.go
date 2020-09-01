/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"
)

// Type represents something that is a Go type
type Type interface {
	// RequiredImports returns a list of packages required by this type
	RequiredImports() []PackageReference

	// References returns the names of all types that this type
	// references. For example, an Array of Persons references a
	// Person.
	References() TypeNameSet

	// AsType renders as a Go abstract syntax tree for a type
	// (yes this says ast.Expr but that is what the Go 'ast' package uses for types)
	AsType(codeGenerationContext *CodeGenerationContext) ast.Expr

	// AsDeclarations renders as a Go abstract syntax tree for a declaration
	AsDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description []string) []ast.Decl

	// Equals returns true if the passed type is the same as this one, false otherwise
	Equals(t Type) bool

	// Make sure all Types have a printable version for debugging/user info.
	// This doesn't need to be a full representation of the type.
	fmt.Stringer
}

// TypeEquals decides if the types are the same and handles the `nil` case
func TypeEquals(left, right Type) bool {
	if left == nil {
		return right == nil
	}

	return left.Equals(right)
}
