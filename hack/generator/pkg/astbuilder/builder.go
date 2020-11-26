/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"go/token"

	ast "github.com/dave/dst"
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
func NewVariableQualified(varName string, qualifier string, structName string) ast.Stmt {
	return &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: ast.NewIdent(varName),
					Type: &ast.SelectorExpr{
						X:   ast.NewIdent(qualifier),
						Sel: ast.NewIdent(structName),
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
func NewVariable(varName string, structName string) ast.Stmt {
	return &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: ast.NewIdent(varName),
					Type: ast.NewIdent(structName),
				},
			},
		},
	}
}

// LocalVariableDeclaration performs a local variable declaration for use within a method like:
// 	var <ident> <typ>
func LocalVariableDeclaration(ident string, typ ast.Expr, comment string) ast.Stmt {
	return &ast.DeclStmt{
		Decl: VariableDeclaration(ident, typ, comment),
	}
}

// VariableDeclaration performs a global variable declaration like:
//  // <comment>
// 	var <ident> <typ>
// For a LocalVariable within a method, use LocalVariableDeclaration() to create an ast.Stmt instead
func VariableDeclaration(ident string, typ ast.Expr, comment string) *ast.GenDecl {
	decl := &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{
					ast.NewIdent(ident),
				},
				Type: typ,
			},
		},
	}

	AddWrappedComment(&decl.Decs.Start, comment, 80)

	return decl
}

// AppendList returns a statement for a list append, like:
//     <lhs> = append(<lhs>, <rhs>)
func AppendList(lhs ast.Expr, rhs ast.Expr) ast.Stmt {
	return SimpleAssignment(
		ast.Clone(lhs).(ast.Expr),
		token.ASSIGN,
		CallFunc("append", ast.Clone(lhs).(ast.Expr), ast.Clone(rhs).(ast.Expr)))
}

// InsertMap returns an assignment statement for inserting an item into a map, like:
// 	<m>[<key>] = <rhs>
func InsertMap(m ast.Expr, key ast.Expr, rhs ast.Expr) *ast.AssignStmt {
	return SimpleAssignment(
		&ast.IndexExpr{
			X:     ast.Clone(m).(ast.Expr),
			Index: ast.Clone(key).(ast.Expr),
		},
		token.ASSIGN,
		ast.Clone(rhs).(ast.Expr))
}

// MakeMap returns the call expression for making a map, like:
// 	make(map[<key>]<value>)
func MakeMap(key ast.Expr, value ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: ast.NewIdent("make"),
		Args: []ast.Expr{
			&ast.MapType{
				Key:   ast.Clone(key).(ast.Expr),
				Value: ast.Clone(value).(ast.Expr),
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
	return CallQualifiedFunc("fmt", "Errorf", callArgs...)
}

// AddrOf returns a statement that gets the address of the provided expression.
//	&<expr>
func AddrOf(exp ast.Expr) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		Op: token.AND,
		X:  exp,
	}
}

// Returns creates a return statement with one or more expressions, of the form
//    return <expr>
// or return <expr>, <expr>, ...
func Returns(returns ...ast.Expr) ast.Stmt {
	return &ast.ReturnStmt{
		Decs: ast.ReturnStmtDecorations{
			NodeDecs: ast.NodeDecs{
				Before: ast.NewLine,
			},
		},
		Results: returns,
	}
}

// QualifiedTypeName generates a reference to a type within an imported package
// of the form <pkg>.<name>
func QualifiedTypeName(pkg string, name string) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   ast.NewIdent(pkg),
		Sel: ast.NewIdent(name),
	}
}
