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
	// Name is the unique name of this function
	// (You can't have two functions with the same name on the same object or resource)
	Name() string

	RequiredPackageReferences() *PackageReferenceSet

	// References returns the set of types to which this function refers.
	// SHOULD include any types which this function references but its receiver doesn't.
	// SHOULD NOT include the receiver of this function.
	References() TypeNameSet[TypeName]

	// AsFunc renders the current instance as a Go abstract syntax tree
	AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *dst.FuncDecl

	// Equals determines if this Function is equal to another one
	Equals(f Function, overrides EqualityOverrides) bool
}
