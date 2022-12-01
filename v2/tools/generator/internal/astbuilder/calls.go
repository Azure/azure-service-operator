/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"github.com/dave/dst"
)

// CallFunc creates an expression to call a function with specified arguments
func CallFunc(funcName string, arguments ...dst.Expr) *dst.CallExpr {
	/*
	 * Sample output:
	 *
	 * <funcName>(<arguments>...)
	 *
	 */
	return createCallExpr(dst.NewIdent(funcName), arguments...)
}

// CallQualifiedFunc creates an expression to call a qualified function with the specified
// arguments.
func CallQualifiedFunc(qualifier string, funcName string, arguments ...dst.Expr) *dst.CallExpr {
	/*
	 * Sample output:
	 *
	 * <qualifier>.<funcName>(arguments...)
	 *
	 */
	return createCallExpr(
		&dst.SelectorExpr{
			X:   dst.NewIdent(qualifier),
			Sel: dst.NewIdent(funcName),
		},
		arguments...)
}

// CallExpr creates an expression to call the named function with the specified arguments
func CallExpr(expr dst.Expr, funcName string, arguments ...dst.Expr) *dst.CallExpr {
	/*
	 * Sample output:
	 *
	 * <expr>.<funcName>(arguments...)
	 *
	 */
	var receiver dst.Expr = expr
	if star, ok := expr.(*dst.StarExpr); ok {
		// We don't need to dereference the expression - even value methods are available from pointer receivers
		receiver = star.X
	}

	return createCallExpr(
		&dst.SelectorExpr{
			X:   receiver,
			Sel: dst.NewIdent(funcName),
		},
		arguments...)
}

func createCallExpr(expr dst.Expr, arguments ...dst.Expr) *dst.CallExpr {
	// Check to see how many of our arguments are nested function calls
	nestedCalls := 0
	for _, e := range arguments {
		if _, ok := e.(*dst.CallExpr); ok {
			nestedCalls++
			break
		}
	}

	// Create our result expression
	result := &dst.CallExpr{
		Fun:  dst.Clone(expr).(dst.Expr),
		Args: Expressions(arguments),
	}

	// If we have more than one nested call, break our arguments out onto separate lines
	// Ditto if we have more than three arguments
	// (heuristics to try and make some of the complex method calls more readable)
	if nestedCalls > 1 || len(result.Args) > 3 {
		for _, e := range result.Args {
			e.Decorations().Before = dst.NewLine
		}
	}

	return result
}

// InvokeFunc creates a statement to invoke a function with specified arguments.
// If you want to use the result of the function call as a value, use CallFunc() instead.
func InvokeFunc(funcName string, arguments ...dst.Expr) dst.Stmt {
	/*
	 * Sample output:
	 *
	 * <funcName>(arguments...)
	 *
	 */
	return &dst.ExprStmt{
		X: CallFunc(funcName, arguments...),
	}
}

// InvokeQualifiedFunc creates a statement to invoke a qualified function with specified
// arguments.
// If you want to use the result of the function call as a value, use CallQualifiedFunc() instead.
func InvokeQualifiedFunc(qualifier string, funcName string, arguments ...dst.Expr) dst.Stmt {
	/*
	 * Sample output:
	 *
	 * <qualifier>.<funcName>(arguments...)
	 *
	 */
	return &dst.ExprStmt{
		X: CallQualifiedFunc(qualifier, funcName, arguments...),
	}
}

// InvokeExpr creates a statement to invoke the named function with the specified arguments.
// If you want to use the result of the function call as a value, use CallExpr() instead.
func InvokeExpr(expr dst.Expr, funcName string, arguments ...dst.Expr) dst.Stmt {
	/*
	 * Sample output:
	 *
	 * <expr>.<funcName>(arguments...)
	 *
	 */
	return &dst.ExprStmt{
		X: CallExpr(expr, funcName, arguments...),
	}
}
