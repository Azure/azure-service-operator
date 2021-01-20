/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"go/token"

	"github.com/dave/dst"
)

// SimpleAssignment performs a simple assignment like:
//     <lhs> := <rhs>       // tok = token.DEFINE
// or  <lhs> = <rhs>        // tok = token.ASSIGN
func SimpleAssignment(lhs dst.Expr, tok token.Token, rhs dst.Expr) *dst.AssignStmt {
	return &dst.AssignStmt{
		Lhs: []dst.Expr{
			lhs,
		},
		Tok: tok,
		Rhs: []dst.Expr{
			rhs,
		},
	}
}

// SimpleAssignmentWithErr performs a simple assignment like:
// 	    <lhs>, err := <rhs>       // tok = token.DEFINE
// 	or  <lhs>, err = <rhs>        // tok = token.ASSIGN
func SimpleAssignmentWithErr(lhs dst.Expr, tok token.Token, rhs dst.Expr) *dst.AssignStmt {
	errId := dst.NewIdent("err")
	return &dst.AssignStmt{
		Lhs: []dst.Expr{
			lhs,
			errId,
		},
		Tok: tok,
		Rhs: []dst.Expr{
			rhs,
		},
	}
}
