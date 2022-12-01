/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"
)

// SimpleAssignment performs a simple assignment.
// See also ShortDeclaration.
func SimpleAssignment(lhs dst.Expr, rhs dst.Expr) *dst.AssignStmt {
	/*
	 * Sample output:
	 *
	 * <lhs> = <rhs>
	 *
	 */
	return &dst.AssignStmt{
		Lhs: []dst.Expr{
			dst.Clone(lhs).(dst.Expr),
		},
		Tok: token.ASSIGN,
		Rhs: []dst.Expr{
			dst.Clone(rhs).(dst.Expr),
		},
	}
}

// ShortDeclaration performs a simple assignment to a new variable.
// Method naming inspired by https://tour.golang.org/basics/10.
func ShortDeclaration(id string, rhs dst.Expr) *dst.AssignStmt {
	/*
	 * Sample output:
	 *
	 * <id> := <rhs>
	 *
	 */
	return &dst.AssignStmt{
		Lhs: []dst.Expr{
			dst.NewIdent(id),
		},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			dst.Clone(rhs).(dst.Expr),
		},
	}
}

// AssignmentStatement allows for either variable declaration or assignment by passing the required token
// Only token.DEFINE and token.ASSIGN are supported, other values will panic.
// Use SimpleAssignment or ShortDeclaration if possible; use this method only if you must.
func AssignmentStatement(lhs dst.Expr, tok token.Token, rhs dst.Expr) *dst.AssignStmt {
	if tok != token.ASSIGN && tok != token.DEFINE {
		panic(fmt.Sprintf("token %q not supported in VariableAssignment", tok))
	}

	return &dst.AssignStmt{
		Lhs: Expressions(lhs),
		Tok: tok,
		Rhs: Expressions(rhs),
	}
}

// QualifiedAssignment performs a simple assignment.
func QualifiedAssignment(lhs dst.Expr, lhsSel string, tok token.Token, rhs dst.Expr) *dst.AssignStmt {
	/*
	 * Sample output:
	 *
	 * <lhs>.<lhsSel> := <rhs>       // tok = token.DEFINE
	 * or
	 * <lhs>.<lhsSel> = <rhs>        // tok = token.ASSIGN
	 *
	 */
	return &dst.AssignStmt{
		Lhs: []dst.Expr{Selector(lhs, lhsSel)},
		Tok: tok,
		Rhs: []dst.Expr{dst.Clone(rhs).(dst.Expr)},
	}
}

// SimpleAssignmentWithErr performs a simple assignment like:
func SimpleAssignmentWithErr(lhs dst.Expr, tok token.Token, rhs dst.Expr) *dst.AssignStmt {
	/*
	 * Sample output:
	 *
	 * <lhs>, err := <rhs>       // tok = token.DEFINE
	 * or
	 * <lhs>, err = <rhs>        // tok = token.ASSIGN
	 *
	 */
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
func AssignToInterface(lhsVar string, rhs dst.Expr) *dst.DeclStmt {
	/*
	 * Sample output:
	 *
	 * var <lhsVar> interface{} = <rhs>
	 *
	 */
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
