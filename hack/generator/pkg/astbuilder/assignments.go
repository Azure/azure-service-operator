/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"go/token"

	ast "github.com/dave/dst"
)

// SimpleAssignment performs a simple assignment like:
//     <lhs> := <rhs>       // tok = token.DEFINE
// or  <lhs> = <rhs>        // tok = token.ASSIGN
func SimpleAssignment(lhs ast.Expr, tok token.Token, rhs ast.Expr) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			lhs,
		},
		Tok: tok,
		Rhs: []ast.Expr{
			rhs,
		},
	}
}

// SimpleAssignmentWithErr performs a simple assignment like:
// 	    <lhs>, err := <rhs>       // tok = token.DEFINE
// 	or  <lhs>, err = <rhs>        // tok = token.ASSIGN
func SimpleAssignmentWithErr(lhs ast.Expr, tok token.Token, rhs ast.Expr) *ast.AssignStmt {
	errId := ast.NewIdent("err")
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			lhs,
			errId,
		},
		Tok: tok,
		Rhs: []ast.Expr{
			rhs,
		},
	}
}
