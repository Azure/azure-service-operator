/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"go/token"

	"github.com/dave/dst"
)

// CheckErrorAndReturn checks if the err is non-nil, and if it is returns. For example:
// 	if err != nil {
// 		return <otherReturns...>, err
//	}
func CheckErrorAndReturn(otherReturns ...dst.Expr) dst.Stmt {

	returnValues := append([]dst.Expr{}, otherReturns...)
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

// LocalVariableDeclaration performs a local variable declaration for use within a method like:
// 	var <ident> <typ>
func LocalVariableDeclaration(ident string, typ dst.Expr, comment string) dst.Stmt {
	return &dst.DeclStmt{
		Decl: VariableDeclaration(ident, typ, comment),
	}
}

// VariableDeclaration performs a global variable declaration like:
//  // <comment>
// 	var <ident> <typ>
// For a LocalVariable within a method, use LocalVariableDeclaration() to create an ast.Stmt instead
func VariableDeclaration(ident string, typ dst.Expr, comment string) *dst.GenDecl {
	decl := &dst.GenDecl{
		Tok: token.VAR,
		Specs: []dst.Spec{
			&dst.ValueSpec{
				Names: []*dst.Ident{
					dst.NewIdent(ident),
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
func AppendList(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
	return SimpleAssignment(
		dst.Clone(lhs).(dst.Expr),
		token.ASSIGN,
		CallFunc("append", dst.Clone(lhs).(dst.Expr), dst.Clone(rhs).(dst.Expr)))
}

// InsertMap returns an assignment statement for inserting an item into a map, like:
// 	<m>[<key>] = <rhs>
func InsertMap(m dst.Expr, key dst.Expr, rhs dst.Expr) *dst.AssignStmt {
	return SimpleAssignment(
		&dst.IndexExpr{
			X:     dst.Clone(m).(dst.Expr),
			Index: dst.Clone(key).(dst.Expr),
		},
		token.ASSIGN,
		dst.Clone(rhs).(dst.Expr))
}

// MakeMap returns the call expression for making a map, like:
// 	make(map[<key>]<value>)
func MakeMap(key dst.Expr, value dst.Expr) *dst.CallExpr {
	return &dst.CallExpr{
		Fun: dst.NewIdent("make"),
		Args: []dst.Expr{
			&dst.MapType{
				Key:   dst.Clone(key).(dst.Expr),
				Value: dst.Clone(value).(dst.Expr),
			},
		},
	}
}

// TypeAssert returns an assignment statement with a type assertion
// 	<lhs>, ok := <rhs>.(<type>)
func TypeAssert(lhs dst.Expr, rhs dst.Expr, typ dst.Expr) *dst.AssignStmt {

	return &dst.AssignStmt{
		Lhs: []dst.Expr{
			lhs,
			dst.NewIdent("ok"),
		},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.TypeAssertExpr{
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
func ReturnIfOk(returns ...dst.Expr) *dst.IfStmt {
	return ReturnIfExpr(dst.NewIdent("ok"), returns...)
}

// ReturnIfNotOk checks a boolean ok variable and if it is not ok returns the specified values
//	if !ok {
//		return <returns>
//	}
func ReturnIfNotOk(returns ...dst.Expr) *dst.IfStmt {
	return ReturnIfExpr(
		&dst.UnaryExpr{
			Op: token.NOT,
			X:  dst.NewIdent("ok"),
		},
		returns...)
}

// ReturnIfNil checks if a variable is nil and if it is returns, like:
// 	if <toCheck> == nil {
// 		return <returns...>
//	}
func ReturnIfNil(toCheck dst.Expr, returns ...dst.Expr) dst.Stmt {
	return ReturnIfExpr(
		&dst.BinaryExpr{
			X:  toCheck,
			Op: token.EQL,
			Y:  dst.NewIdent("nil"),
		},
		returns...)
}

// ReturnIfNotNil checks if a variable is not nil and if it is returns, like:
// 	if <toCheck> != nil {
// 		return <returns...>
//	}
func ReturnIfNotNil(toCheck dst.Expr, returns ...dst.Expr) dst.Stmt {
	return ReturnIfExpr(
		&dst.BinaryExpr{
			X:  toCheck,
			Op: token.NEQ,
			Y:  dst.NewIdent("nil"),
		},
		returns...)
}

// ReturnIfExpr returns if the expression evaluates as true.
//	if <cond> {
// 		return <returns...>
//	}
func ReturnIfExpr(cond dst.Expr, returns ...dst.Expr) *dst.IfStmt {
	if len(returns) == 0 {
		panic("Expected at least 1 return for ReturnIfOk")
	}

	return &dst.IfStmt{
		Cond: cond,
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{
					Results: returns,
				},
			},
		},
	}
}

// FormatError produces a call to fmt.Errorf with the given format string and args
//	fmt.Errorf(<formatString>, <args>)
func FormatError(fmtPackage string, formatString string, args ...dst.Expr) dst.Expr {
	var callArgs []dst.Expr
	callArgs = append(
		callArgs,
		StringLiteral(formatString))
	callArgs = append(callArgs, args...)
	return CallQualifiedFunc(fmtPackage, "Errorf", callArgs...)
}

// AddrOf returns a statement that gets the address of the provided expression.
//	&<expr>
func AddrOf(exp dst.Expr) *dst.UnaryExpr {
	return &dst.UnaryExpr{
		Op: token.AND,
		X:  exp,
	}
}

// Returns creates a return statement with one or more expressions, of the form
//    return <expr>
// or return <expr>, <expr>, ...
func Returns(returns ...dst.Expr) dst.Stmt {
	return &dst.ReturnStmt{
		Decs: dst.ReturnStmtDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.NewLine,
			},
		},
		Results: returns,
	}
}

// QualifiedTypeName generates a reference to a type within an imported package
// of the form <pkg>.<name>
func QualifiedTypeName(pkg string, name string) *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X:   dst.NewIdent(pkg),
		Sel: dst.NewIdent(name),
	}
}
