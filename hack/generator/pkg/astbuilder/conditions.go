/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"go/token"

	"github.com/dave/dst"
)

// SimpleIfElse creates a simple if statement with a single statement in each branch
//
// if <condition> {
//     <trueBranch>
// } else {
//     <falseBranch>
// }
//
func SimpleIfElse(condition dst.Expr, trueBranch []dst.Stmt, falseBranch []dst.Stmt) *dst.IfStmt {
	result := &dst.IfStmt{
		Cond: condition,
		Body: StatementBlock(trueBranch...),
		Else: StatementBlock(falseBranch...),
	}

	return result
}

// IfEqual executes a series of statements if the supplied expressions are
//
// if <left> == <right> {
//     <statements>
// }
//
func IfEqual(left dst.Expr, right dst.Expr, statements ...dst.Stmt) *dst.IfStmt {
	return &dst.IfStmt{
		Cond: Equal(left, right),
		Body: StatementBlock(statements...),
	}
}

// IfNotNil executes a series of statements if the supplied expression is not nil
//
// if <source> != nil {
//     <statements>
// }
//
func IfNotNil(toCheck dst.Expr, statements ...dst.Stmt) *dst.IfStmt {
	return &dst.IfStmt{
		Cond: &dst.BinaryExpr{
			X:  dst.Clone(toCheck).(dst.Expr),
			Op: token.NEQ,
			Y:  Nil(),
		},
		Body: StatementBlock(statements...),
	}
}

// IfOk checks a boolean ok variable and if it is ok runs the given statements
//
//	if ok {
//		<statements>
//	}
//
func IfOk(statements ...dst.Stmt) *dst.IfStmt {
	return &dst.IfStmt{
		Cond: dst.NewIdent("ok"),
		Body: StatementBlock(statements...),
	}
}

// IfNotOk checks a boolean ok variable and if it is not ok runs the given statements
//
//	if !ok {
//		<statements>
//	}
//
func IfNotOk(statements ...dst.Stmt) *dst.IfStmt {
	return &dst.IfStmt{
		Cond: &dst.UnaryExpr{
			Op: token.NOT,
			X:  dst.NewIdent("ok"),
		},
		Body: StatementBlock(statements...),
	}
}

