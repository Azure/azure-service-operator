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
			dst.Clone(lhs).(dst.Expr),
		},
		Tok: tok,
		Rhs: []dst.Expr{
			dst.Clone(rhs).(dst.Expr),
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
			dst.Clone(lhs).(dst.Expr),
			errId,
		},
		Tok: tok,
		Rhs: []dst.Expr{
			dst.Clone(rhs).(dst.Expr),
		},
	}
}

// AssignToInterface performs an assignment of a well-typed variable to an interface{}. This is usually used to
// perform a type assertion on a concrete type in a subsequent statement (which Go doesn't allow, it only allows type
// assertions on interface types).
//	var <lhsVar> interface{} = <rhs>
func AssignToInterface(lhsVar string, rhs dst.Expr) *dst.DeclStmt {
	return &dst.DeclStmt{
		Decl: &dst.GenDecl{
			Tok: token.VAR,
			Specs: []dst.Spec{
				&dst.ValueSpec{
					Names: []*dst.Ident{
						dst.NewIdent(lhsVar),
					},
					Type: dst.NewIdent("interface{}"),
					Values: []dst.Expr{
						dst.Clone(rhs).(dst.Expr),
					},
				},
			},
		},
	}
}
