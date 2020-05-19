/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "go/ast"

// HasImports describes an entity that has required imports
type HasImports interface {
	// RequiredImports returns a list of packages required by this
	RequiredImports() []PackageReference
}

// ReferenceChecker is used to check for references to a specific definition
type ReferenceChecker interface {
	// References determines if this type has a direct reference to the given definition name
	// For example, a struct references its field
	References(d *DefinitionName) bool
}

// Definition represents models that can render into Go code
type Definition interface {
	// FileNameHint returns what a file that contains this definition (if any) should be called
	// this is not always used as we might combine multiple definitions into one file
	FileNameHint() string

	// AsDecalarations() generates the Go code representing this definition
	AsDeclarations() []ast.Decl

	// DefinitionName: How do you refer to this definition?
	Reference() *DefinitionName

	// Type: What is the type associated to this definition?
	Type() Type

	// Tidy cleans up the definition prior to code generation
	Tidy()
}

// A HasRelatedDefinitions is capable of creating additional definitions
type HasRelatedDefinitions interface {
	RelatedDefinitions(ref PackageReference, namehint string, idFactory IdentifierFactory) []Definition
}

// Type represents something that is a Go type
type Type interface {
	HasImports
	ReferenceChecker

	// AsType renders the current instance as a Go abstract syntax tree
	AsType() ast.Expr

	// Equals returns true if the passed type is the same as this one, false otherwise
	Equals(t Type) bool
}
