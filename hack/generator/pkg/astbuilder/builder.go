/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"go/ast"
	"go/token"
)

// CheckErrorAndReturn checks if the err is non-nil, and if it is returns. For example:
// 	if err != nil {
// 		return <otherReturns...>, err
//	}
func CheckErrorAndReturn(otherReturns ...ast.Expr) ast.Stmt {

	returnValues := append([]ast.Expr{}, otherReturns...)
	returnValues = append(returnValues, ast.NewIdent("err"))

	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  ast.NewIdent("err"),
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: returnValues,
				},
			},
		},
	}
}

// NewQualifiedStruct creates a new assignment statement where a struct is constructed and stored in a variable of the given name.
// For example:
//     <varName> := <packageRef>.<structName>{}
func NewQualifiedStruct(varName *ast.Ident, qualifier *ast.Ident, structName *ast.Ident) ast.Stmt {
	return SimpleAssignment(
		varName,
		token.DEFINE,
		&ast.CompositeLit{
			Type: &ast.SelectorExpr{
				X:   qualifier,
				Sel: structName,
			},
		})
}

// NewStruct creates a new assignment statement where a struct is constructed and stored in a variable of the given name.
// For example:
//     <varName> := <structName>{}
func NewStruct(varName *ast.Ident, structName *ast.Ident) ast.Stmt {
	return SimpleAssignment(
		varName,
		token.DEFINE,
		&ast.CompositeLit{
			Type: structName,
		})
}

// LocalVariableDeclaration performs a local variable declaration for use within a method like:
// 	var <ident> <typ>
func LocalVariableDeclaration(ident *ast.Ident, typ ast.Expr, comment string) ast.Stmt {
	return &ast.DeclStmt{
		Decl: VariableDeclaration(ident, typ, comment),
	}
}

// VariableDeclaration performs a global variable declaration like:
//  // <comment>
// 	var <ident> <typ>
// For a LocalVariable within a method, use LocalVariableDeclaration() to create an ast.Stmt instead
func VariableDeclaration(ident *ast.Ident, typ ast.Expr, comment string) *ast.GenDecl {
	decl := &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{
					ident,
				},
				Type: typ,
			},
		},
		Doc: &ast.CommentGroup{},
	}

	AddWrappedComment(&decl.Doc.List, comment, 80)

	return decl
}

// AppendList returns a statement for a list append, like:
//     <lhs> = append(<lhs>, <rhs>)
func AppendList(lhs ast.Expr, rhs ast.Expr) ast.Stmt {
	return SimpleAssignment(
		lhs,
		token.ASSIGN,
		CallFuncByName("append", lhs, rhs))
}

// InsertMap returns an assignment statement for inserting an item into a map, like:
// 	<m>[<key>] = <rhs>
func InsertMap(m ast.Expr, key ast.Expr, rhs ast.Expr) *ast.AssignStmt {
	return SimpleAssignment(
		&ast.IndexExpr{
			X:     m,
			Index: key,
		},
		token.ASSIGN,
		rhs)
}

// MakeMap returns the call expression for making a map, like:
// 	make(map[<key>]<value>)
func MakeMap(key ast.Expr, value ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: ast.NewIdent("make"),
		Args: []ast.Expr{
			&ast.MapType{
				Key:   key,
				Value: value,
			},
		},
	}
}

// TypeAssert returns an assignment statement with a type assertion
// 	<lhs>, ok := <rhs>.(<type>)
func TypeAssert(lhs ast.Expr, rhs ast.Expr, typ ast.Expr) *ast.AssignStmt {

	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			lhs,
			ast.NewIdent("ok"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.TypeAssertExpr{
				X:    rhs,
				Type: typ,
			},
		},
	}
}

// ReturnIfOk checks a boolean ok variable and if it is ok returns the specified values
//	if ok {
//		return <returns>
//	}
func ReturnIfOk(returns ...ast.Expr) *ast.IfStmt {
	return ReturnIfExpr(ast.NewIdent("ok"), returns...)
}

// ReturnIfNotOk checks a boolean ok variable and if it is not ok returns the specified values
//	if !ok {
//		return <returns>
//	}
func ReturnIfNotOk(returns ...ast.Expr) *ast.IfStmt {
	return ReturnIfExpr(
		&ast.UnaryExpr{
			Op: token.NOT,
			X:  ast.NewIdent("ok"),
		},
		returns...)
}

// ReturnIfNil checks if a variable is nil and if it is returns, like:
// 	if <toCheck> == nil {
// 		return <returns...>
//	}
func ReturnIfNil(toCheck ast.Expr, returns ...ast.Expr) ast.Stmt {
	return ReturnIfExpr(
		&ast.BinaryExpr{
			X:  toCheck,
			Op: token.EQL,
			Y:  ast.NewIdent("nil"),
		},
		returns...)
}

// ReturnIfNotNil checks if a variable is not nil and if it is returns, like:
// 	if <toCheck> != nil {
// 		return <returns...>
//	}
func ReturnIfNotNil(toCheck ast.Expr, returns ...ast.Expr) ast.Stmt {
	return ReturnIfExpr(
		&ast.BinaryExpr{
			X:  toCheck,
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		returns...)
}

// ReturnIfExpr returns if the expression evaluates as true.
//	if <cond> {
// 		return <returns...>
//	}
func ReturnIfExpr(cond ast.Expr, returns ...ast.Expr) *ast.IfStmt {
	if len(returns) == 0 {
		panic("Expected at least 1 return for ReturnIfOk")
	}

	return &ast.IfStmt{
		Cond: cond,
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: returns,
				},
			},
		},
	}
}

// FormatError produces a call to fmt.Errorf with the given format string and args
//	fmt.Errorf(<formatString>, <args>)
func FormatError(formatString string, args ...ast.Expr) ast.Expr {
	var callArgs []ast.Expr
	callArgs = append(
		callArgs,
		StringLiteral(formatString))
	callArgs = append(callArgs, args...)
	return CallQualifiedFuncByName("fmt", "Errorf", callArgs...)
}

// AddrOf returns a statement that gets the address of the provided expression.
//	&<expr>
func AddrOf(exp ast.Expr) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		Op: token.AND,
		X:  exp,
	}
}

func Returns(returns ...ast.Expr) ast.Stmt {
	return &ast.ReturnStmt{
		Results: returns,
	}
}
