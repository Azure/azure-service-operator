/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
)

// Function represents something that is an (unnamed) Go function
type Function interface {
	RequiredImports() []*PackageReference

	// References returns the set of types to which this function refers.
	References() TypeNameSet

	// AsFunc renders the current instance as a Go abstract syntax tree
	AsFunc(codeGenerationContext *CodeGenerationContext, receiver *TypeName, methodName string) *ast.FuncDecl

	// Equals determines if this Function is equal to another one
	Equals(f Function) bool
}
