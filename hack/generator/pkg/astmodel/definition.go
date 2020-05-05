/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "go/ast"

type HasImports interface {
	// RequiredImports returns a list of packages required by this
	RequiredImports() []PackageReference
}

// Definition represents models that can render into Go code
type Definition interface {
	HasImports

	// FileNameHint returns what a file that contains this definition (if any) should be called
	// this is not always used as we might combine multiple definitions into one file
	FileNameHint() string

	// AsDecalaration() renders a definition into a Go abstract syntax tree
	AsDeclaration() ast.Decl
}

// Type represents something that is a Go type
type Type interface {
	HasImports

	// AsType renders the current instance as a Go abstract syntax tree
	AsType() ast.Expr
}
