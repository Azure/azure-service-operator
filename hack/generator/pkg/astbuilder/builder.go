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

// CheckErrorAndReturn checks if the err is non-nil, and if it is returns.
//
// 	if err != nil {
// 		return <otherReturns...>, err
//	}
func CheckErrorAndReturn(otherReturns ...dst.Expr) dst.Stmt {

	returnValues := append([]dst.Expr{}, cloneExprSlice(otherReturns)...)
	returnValues = append(returnValues, dst.NewIdent("err"))

	return &dst.IfStmt{
		Cond: &dst.BinaryExpr{
			X:  dst.NewIdent("err"),
			Op: token.NEQ,
			Y:  dst.NewIdent("nil"),
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{
					Results: returnValues,
				},
			},
		},
	}
}

// NewVariableQualified creates a new declaration statement where a variable is declared
// with its default value.
//
// For example:
//     var <varName> <packageRef>.<structName>
//
// Note that it does *not* do:
//     <varName> := <packageRef>.<structName>{}
//
// …as that does not work for enum types.
func NewVariableQualified(varName string, qualifier string, structName string) dst.Stmt {
	return &dst.DeclStmt{
		Decl: &dst.GenDecl{
			Tok: token.VAR,
			Specs: []dst.Spec{
				&dst.TypeSpec{
					Name: dst.NewIdent(varName),
					Type: &dst.SelectorExpr{
						X:   dst.NewIdent(qualifier),
						Sel: dst.NewIdent(structName),
					},
				},
			},
		},
	}
}

// NewVariable creates a new declaration statement where a variable is declared
// with its default value.
//
// For example:
//     var <varName> <structName>
//
// Note that it does *not* do:
//     <varName> := <structName>{}
//
// …as that does not work for enum types.
func NewVariable(varName string, structName string) dst.Stmt {
	return &dst.DeclStmt{
		Decl: &dst.GenDecl{
			Tok: token.VAR,
			Specs: []dst.Spec{
				&dst.TypeSpec{
					Name: dst.NewIdent(varName),
					Type: dst.NewIdent(structName),
				},
			},
		},
	}
}

// LocalVariableDeclaration performs a local variable declaration for use within a method
//
// 	var <ident> <typ>
//
func LocalVariableDeclaration(ident string, typ dst.Expr, comment string) dst.Stmt {
	return &dst.DeclStmt{
		Decl: VariableDeclaration(ident, typ, comment),
	}
}

// VariableDeclaration performs a global variable declaration
//
//  // <comment>
// 	var <ident> <typ>
//
// For a LocalVariable within a method, use LocalVariableDeclaration() to create an ast.Stmt instead
func VariableDeclaration(ident string, typ dst.Expr, comment string) *dst.GenDecl {
	decl := &dst.GenDecl{
		Tok: token.VAR,
		Specs: []dst.Spec{
			&dst.ValueSpec{
				Names: []*dst.Ident{
					dst.NewIdent(ident),
				},
				Type: dst.Clone(typ).(dst.Expr),
			},
		},
	}

	AddWrappedComment(&decl.Decs.Start, comment, 80)

	return decl
}

// TypeAssert returns an assignment statement with a type assertion
//
// 	<lhs>, ok := <rhs>.(<type>)
//
func TypeAssert(lhs dst.Expr, rhs dst.Expr, typ dst.Expr) *dst.AssignStmt {

	return &dst.AssignStmt{
		Lhs: []dst.Expr{
			dst.Clone(lhs).(dst.Expr),
			dst.NewIdent("ok"),
		},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.TypeAssertExpr{
				X:    dst.Clone(rhs).(dst.Expr),
				Type: dst.Clone(typ).(dst.Expr),
			},
		},
	}
}

// ReturnIfOk checks a boolean ok variable and if it is ok returns the specified values
//
//	if ok {
//		return <returns>
//	}
//
func ReturnIfOk(returns ...dst.Expr) *dst.IfStmt {
	return ReturnIfExpr(dst.NewIdent("ok"), returns...)
}

// ReturnIfNotOk checks a boolean ok variable and if it is not ok returns the specified values
//
//	if !ok {
//		return <returns>
//	}
//
func ReturnIfNotOk(returns ...dst.Expr) *dst.IfStmt {
	return ReturnIfExpr(
		&dst.UnaryExpr{
			Op: token.NOT,
			X:  dst.NewIdent("ok"),
		},
		returns...)
}

// ReturnIfNil checks if a variable is nil and if it is returns
//
// 	if <toCheck> == nil {
// 		return <returns...>
//	}
//
func ReturnIfNil(toCheck dst.Expr, returns ...dst.Expr) dst.Stmt {
	return ReturnIfExpr(
		&dst.BinaryExpr{
			X:  dst.Clone(toCheck).(dst.Expr),
			Op: token.EQL,
			Y:  dst.NewIdent("nil"),
		},
		returns...)
}

// ReturnIfNotNil checks if a variable is not nil and if it is returns
//
// 	if <toCheck> != nil {
// 		return <returns...>
//	}
//
func ReturnIfNotNil(toCheck dst.Expr, returns ...dst.Expr) dst.Stmt {
	return ReturnIfExpr(
		&dst.BinaryExpr{
			X:  dst.Clone(toCheck).(dst.Expr),
			Op: token.NEQ,
			Y:  dst.NewIdent("nil"),
		},
		returns...)
}

// ReturnIfExpr returns if the expression evaluates as true
//
//	if <cond> {
// 		return <returns...>
//	}
//
func ReturnIfExpr(cond dst.Expr, returns ...dst.Expr) *dst.IfStmt {
	if len(returns) == 0 {
		panic("Expected at least 1 return for ReturnIfOk")
	}

	return &dst.IfStmt{
		Cond: cond,
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{
					Results: cloneExprSlice(returns),
				},
			},
		},
	}
}

// FormatError produces a call to fmt.Errorf with the given format string and args
//
//	fmt.Errorf(<formatString>, <args>)
//
func FormatError(fmtPackage string, formatString string, args ...dst.Expr) dst.Expr {
	var callArgs []dst.Expr
	callArgs = append(
		callArgs,
		StringLiteral(formatString))
	callArgs = append(callArgs, args...)
	return CallQualifiedFunc(fmtPackage, "Errorf", callArgs...)
}

// AddrOf returns a statement that gets the address of the provided expression.
//
//	&<expr>
//
func AddrOf(expr dst.Expr) *dst.UnaryExpr {
	return &dst.UnaryExpr{
		Op: token.AND,
		X:  dst.Clone(expr).(dst.Expr),
	}
}

// Dereference returns a statement that dereferences the pointer returned by the provided expression
//
// *<expr>
//
func Dereference(expr dst.Expr) *dst.UnaryExpr {
	return &dst.UnaryExpr{
		Op: token.MUL,
		X:  dst.Clone(expr).(dst.Expr),
	}
}

// Returns creates a return statement with one or more expressions, of the form
//
//    return <expr>
// or return <expr>, <expr>, ...
//
func Returns(returns ...dst.Expr) dst.Stmt {
	return &dst.ReturnStmt{
		Decs: dst.ReturnStmtDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.NewLine,
			},
		},
		Results: cloneExprSlice(returns),
	}
}

// ReturnNoError creates a return nil statement for when no error occurs
//
//    // No error
//    return nil
//
func ReturnNoError() dst.Stmt {
	result := Returns(dst.NewIdent("nil"))
	result.Decorations().Before = dst.EmptyLine
	result.Decorations().Start.Append("// No error")
	return result
}

// WrappedErrorf returns the err local, wrapped with additional information
//
// errors.Wrap(err, <message>)
//
func WrappedErrorf(template string, args ...interface{}) dst.Expr {
	return CallQualifiedFunc(
		"errors",
		"Wrap",
		dst.NewIdent("err"),
		StringLiteralf(template, args...))
}

// QualifiedTypeName generates a reference to a type within an imported package
//
// <pkg>.<name>
//
func QualifiedTypeName(pkg string, name string) *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X:   dst.NewIdent(pkg),
		Sel: dst.NewIdent(name),
	}
}

// Selector generates a field reference into an existing expression
//
// <expr>.<name>
//
func Selector(expr dst.Expr, name string) *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X:   dst.Clone(expr).(dst.Expr),
		Sel: dst.NewIdent(name),
	}
}

// NotEqual generates a != comparison between the two expressions
//
// <lhs> != <rhs>
//
func NotEqual(lhs dst.Expr, rhs dst.Expr) *dst.BinaryExpr {
	return &dst.BinaryExpr{
		X:  dst.Clone(lhs).(dst.Expr),
		Op: token.NEQ,
		Y:  dst.Clone(rhs).(dst.Expr),
	}
}

// StatementBlock generates a block containing the supplied statements
func StatementBlock(statements ...dst.Stmt) *dst.BlockStmt {
	return &dst.BlockStmt{
		List: cloneStmtSlice(statements),
	}
}

// EnsureStatementBlock wraps any statement into a block safely
// (without double wrapping an existing block)
func EnsureStatementBlock(statement dst.Stmt) *dst.BlockStmt {
	if block, ok := statement.(*dst.BlockStmt); ok {
		return block
	}

	return StatementBlock(statement)
}

// Statements creates a sequence of statements from the provided values, each of which may be a
// single dst.Stmt or a slice of multiple []dst.Stmts
func Statements(statements ...interface{}) []dst.Stmt {
	var result []dst.Stmt
	for _, s := range statements {
		switch s := s.(type) {
		case nil:
			// Skip nils
			continue
		case dst.Stmt:
			// Add a single statement
			result = append(result, s)
		case []dst.Stmt:
			// Add many statements
			result = append(result, s...)
		default:
			panic(fmt.Sprintf("expected dst.Stmt or []dst.Stmt, but found %T", s))
		}
	}

	return result
}

// cloneExprSlice is a utility method to clone a slice of expressions
func cloneExprSlice(exprs []dst.Expr) []dst.Expr {
	var result []dst.Expr
	for _, exp := range exprs {
		result = append(result, dst.Clone(exp).(dst.Expr))
	}

	return result
}

// cloneStmtSlice is a utility method to clone a slice of statements
func cloneStmtSlice(stmts []dst.Stmt) []dst.Stmt {
	var result []dst.Stmt
	for _, st := range stmts {
		result = append(result, dst.Clone(st).(dst.Stmt))
	}

	return result
}
