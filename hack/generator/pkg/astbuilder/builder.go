/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"fmt"
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

// NewQualifiedStruct creates a new assignment statement where a struct is constructed and stored in a variable of the given name. For example:
// 	<varName> := <packageRef>.<structName>{}
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

// NewStruct creates a new assignment statement where a struct is constructed and stored in a variable of the given name. For example:
// 	<varName> := <structName>{}
func NewStruct(varName *ast.Ident, structName *ast.Ident) ast.Stmt {
	return SimpleAssignment(
		varName,
		token.DEFINE,
		&ast.CompositeLit{
			Type: structName,
		})
}

type FuncDetails struct {
	ReceiverIdent *ast.Ident
	ReceiverType  ast.Expr
	Name          *ast.Ident
	Comment       string
	Params        []*ast.Field
	Returns       []*ast.Field
	Body          []ast.Stmt
}

// DefineFunc defines a function (header, body, etc), like:
// 	<comment>
//	func (<receiverIdent> <receiverType>) <name>(<params...>) (<returns...>) {
// 		<body...>
//	}
func DefineFunc(funcDetails FuncDetails) *ast.FuncDecl {

	var comment []*ast.Comment
	if funcDetails.Comment != "" {
		comment = []*ast.Comment{
			{
				Text: fmt.Sprintf("// %s %s", funcDetails.Name, funcDetails.Comment),
			},
		}
	}

	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						funcDetails.ReceiverIdent,
					},
					Type: funcDetails.ReceiverType,
				},
			},
		},
		Name: funcDetails.Name,
		Doc: &ast.CommentGroup{
			List: comment,
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: funcDetails.Params,
			},
			Results: &ast.FieldList{
				List: funcDetails.Returns,
			},
		},
		Body: &ast.BlockStmt{
			List: funcDetails.Body,
		},
	}
}

// SimpleAssignment performs a simple assignment like:
// 	<lhs> <tok> <rhs>
func SimpleAssignment(lhs ast.Expr, tok token.Token, rhs ast.Expr) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			lhs,
		},
		Tok: tok,
		Rhs: []ast.Expr{
			rhs,
		},
	}
}

// SimpleVariableDeclaration performs a simple variable declaration like:
// 	var <ident> <typ>
func SimpleVariableDeclaration(ident *ast.Ident, typ ast.Expr) *ast.DeclStmt {
	return &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						ident,
					},
					Type: typ,
				},
			},
		},
	}
}

// AppendList returns a statement for a list append, like:
//	<lhs> = append(<lhs>, <rhs>)
func AppendList(lhs ast.Expr, rhs ast.Expr) ast.Stmt {
	return SimpleAssignment(
		lhs,
		token.ASSIGN,
		&ast.CallExpr{
			Fun: ast.NewIdent("append"),
			Args: []ast.Expr{
				lhs,
				rhs,
			},
		})
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
func FormatError(formatString string, args ...ast.Expr) *ast.CallExpr {
	var callArgs []ast.Expr
	callArgs = append(
		callArgs,
		&ast.BasicLit{
			Kind:  token.STRING,
			Value: formatString,
		})
	callArgs = append(callArgs, args...)
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent("fmt"),
			Sel: ast.NewIdent("Errorf"),
		},
		Args: callArgs,
	}
}

// AddrOf returns a statement that gets the address of the provided expression.
//	&<expr>
func AddrOf(exp ast.Expr) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		Op: token.AND,
		X:  exp,
	}
}
