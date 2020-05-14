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
}

// Type represents something that is a Go type
type Type interface {
	HasImports

	// AsType renders the current instance as a Go abstract syntax tree
	AsType() ast.Expr

	// References this type has to the given type
	// Does this Type include any direct references to the given Type?
	// "Direct" means we don't walk into any StructReferences (nor could we with these arguments),
	// but we do walk into included StructTypes.
	References(t Type) bool
}
