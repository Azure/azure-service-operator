/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"github.com/dave/dst"
)

// CallFunc creates an expression to call a function with specified arguments
//
// <funcName>(<arguments>...)
//
func CallFunc(funcName string, arguments ...dst.Expr) dst.Expr {
	return &dst.CallExpr{
		Fun:  dst.NewIdent(funcName),
		Args: cloneExprSlice(arguments),
	}
}

// CallQualifiedFunc creates an expression to call a qualified function with the specified
// arguments
//
// <qualifier>.<funcName>(arguments...)
//
func CallQualifiedFunc(qualifier string, funcName string, arguments ...dst.Expr) dst.Expr {
	return &dst.CallExpr{
		Fun: &dst.SelectorExpr{
			X:   dst.NewIdent(qualifier),
			Sel: dst.NewIdent(funcName),
		},
		Args: cloneExprSlice(arguments),
	}
}

// CallExpr creates an expression to call the named function with the specified arguments
//
// <expr>.<funcName>(arguments...)
//
func CallExpr(expr dst.Expr, funcName string, arguments ...dst.Expr) dst.Expr {
	return &dst.CallExpr{
		Fun: &dst.SelectorExpr{
			X:   expr,
			Sel: dst.NewIdent(funcName),
		},
		Args: cloneExprSlice(arguments),
	}
}

// InvokeFunc creates a statement to invoke a function with specified arguments
//
// <funcName>(arguments...)
//
// If you want to use the result of the function call as a value, use CallFunc() instead
func InvokeFunc(funcName string, arguments ...dst.Expr) dst.Stmt {
	return &dst.ExprStmt{
		X: CallFunc(funcName, arguments...),
	}
}

// InvokeQualifiedFunc creates a statement to invoke a qualified function with specified
// arguments
//
// <qualifier>.<funcName>(arguments...)
//
// If you want to use the result of the function call as a value, use CallQualifiedFunc() instead
func InvokeQualifiedFunc(qualifier string, funcName string, arguments ...dst.Expr) dst.Stmt {
	return &dst.ExprStmt{
		X: CallQualifiedFunc(qualifier, funcName, arguments...),
	}
}
