/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "go/ast"

// Definition represents models that can render into Go code
type Definition interface {
	// AsAst() renders a definition into a Go abstract syntax tree
	AsAst() ast.Node
}

// Type represents something that is a Go type
type Type interface {
	// AsType renders the current instance as a Go abstract syntax tree
	AsType() ast.Expr
}
