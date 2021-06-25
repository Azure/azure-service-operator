/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"github.com/dave/dst"
	"go/token"
)

// SimpleIfElse creates a simple if statement with a single statement in each branch
//
// if <condition> {
//     <trueBranch>
// } else {
//     <falseBranch>
// }
//
func SimpleIfElse(condition dst.Expr, trueBranch dst.Stmt, falseBranch dst.Stmt) *dst.IfStmt {
	result := &dst.IfStmt{
		Cond: condition,
		Body: EnsureStatementBlock(trueBranch),
		Else: EnsureStatementBlock(falseBranch),
	}

	return result
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
			Y:  dst.NewIdent("nil"),
		},
		Body: StatementBlock(statements...),
	}
}
