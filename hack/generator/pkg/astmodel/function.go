/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/dave/dst"
)

// Function represents something that is an (unnamed) Go function
type Function interface {
	// The unique name of this function
	// (You can't have two functions with the same name on the same object or resource)
	Name() string

	RequiredPackageReferences() *PackageReferenceSet

	// References returns the set of types to which this function refers.
	// Should *not* include the receiver of this function
	References() TypeNameSet

	// AsFunc renders the current instance as a Go abstract syntax tree
	AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *dst.FuncDecl

	// Equals determines if this Function is equal to another one
	Equals(f Function) bool
}
