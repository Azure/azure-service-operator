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

// NilGuardAndReturn checks if a variable is nil and if it is returns, like:
// 	if <toCheck> == nil {
// 		return <returns...>
//	}
func NilGuardAndReturn(toCheck *ast.Ident, returns ...ast.Expr) ast.Stmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  toCheck,
			Op: token.EQL,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: returns,
				},
			},
		},
	}
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

// AppendList returns an assignment statement for a list append, like:
//	<lhs> = append(<lhs>, <rhs>)
func AppendList(lhs ast.Expr, rhs ast.Expr) *ast.AssignStmt {
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
