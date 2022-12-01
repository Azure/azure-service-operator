/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"go/token"

	"github.com/dave/dst"
)

// SimpleIf creates a simple if statement with multiple statements
func SimpleIf(condition dst.Expr, statements ...dst.Stmt) *dst.IfStmt {
	/*
	 * Sample output:
	 *
	 * if <condition> {
	 *     <trueBranch>
	 * }
	 *
	 */
	return &dst.IfStmt{
		Cond: condition,
		Body: StatementBlock(statements...),
	}
}

// SimpleIfElse creates a simple if else statement. Each branch may contain multiple statements.
func SimpleIfElse(condition dst.Expr, trueBranch []dst.Stmt, falseBranch []dst.Stmt) *dst.IfStmt {
	/*
	 * Sample output:
	 *
	 * if <condition> {
	 *      <trueBranch>
	 * } else {
	 *     <falseBranch>
	 * }
	 *
	 */
	result := &dst.IfStmt{
		Cond: condition,
		Body: StatementBlock(trueBranch...),
		Else: StatementBlock(falseBranch...),
	}

	return result
}

// IfEqual executes a series of statements if the supplied expressions are
func IfEqual(left dst.Expr, right dst.Expr, statements ...dst.Stmt) *dst.IfStmt {
	/*
	 * Sample output:
	 *
	 * if <left> == <right> {
	 *     <statements>
	 * }
	 *
	 */
	return &dst.IfStmt{
		Cond: AreEqual(left, right),
		Body: StatementBlock(statements...),
	}
}

// IfNotNil executes a series of statements if the supplied expression is not nil
func IfNotNil(toCheck dst.Expr, statements ...dst.Stmt) *dst.IfStmt {
	/*
	 * Sample output:
	 *
	 * if <source> != nil {
	 *     <statements>
	 * }
	 *
	 */
	return &dst.IfStmt{
		Cond: NotNil(toCheck),
		Body: StatementBlock(statements...),
	}
}

// IfNil executes a series of statements if the supplied expression is nil
func IfNil(toCheck dst.Expr, statements ...dst.Stmt) *dst.IfStmt {
	/*
	 * Sample output:
	 *
	 * if <source> != nil {
	 *     <statements>
	 * }
	 *
	 */
	return &dst.IfStmt{
		Cond: AreEqual(toCheck, Nil()),
		Body: StatementBlock(statements...),
	}
}

// IfOk checks a boolean ok variable and if it is ok runs the given statements
func IfOk(statements ...dst.Stmt) *dst.IfStmt {
	/*
	 * Sample output:
	 *
	 * if ok {
	 *     <statements>
	 * }
	 *
	 */
	return &dst.IfStmt{
		Cond: dst.NewIdent("ok"),
		Body: StatementBlock(statements...),
	}
}

// IfNotOk checks a boolean ok variable and if it is not ok runs the given statements
func IfNotOk(statements ...dst.Stmt) *dst.IfStmt {
	/*
	 * Sample output:
	 *
	 * if !ok {
	 *     <statements>
	 * }
	 *
	 */
	return &dst.IfStmt{
		Cond: &dst.UnaryExpr{
			Op: token.NOT,
			X:  dst.NewIdent("ok"),
		},
		Body: StatementBlock(statements...),
	}
}

// IfType does a type assertion and executes the provided statements if it is true.
// expr is the expression to cast;
// typeExpr is the type we want to cast to;
// local is the name of the local variable to initialize;
// statements form the body of the if statement,
func IfType(expr dst.Expr, typeExpr dst.Expr, local string, statements ...dst.Stmt) *dst.IfStmt {
	/*
	 * Sample output:
	 *
	 * if <local>, ok := <expr>.(<typeExpr>); ok {
	 *     <statements>
	 * }
	 *
	 */
	return &dst.IfStmt{
		Init: TypeAssert(dst.NewIdent(local), expr, typeExpr),
		Cond: dst.NewIdent("ok"),
		Body: StatementBlock(statements...),
	}
}
