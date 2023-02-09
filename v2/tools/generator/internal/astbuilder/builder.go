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
//	if err != nil {
//		return <otherReturns...>, err
//	}
func CheckErrorAndReturn(otherReturns ...dst.Expr) dst.Stmt {
	retStmt := &dst.ReturnStmt{
		Results: Expressions(otherReturns, dst.NewIdent("err")),
	}

	return CheckErrorAndSingleStatement(retStmt)
}

// CheckErrorAndWrap checks if the err is non-nil, and if it is returns it, wrapped with additional information
// If no arguments are provided, will generate
//
//	if err != nil {
//	     return errors.Wrap(err, <message>)
//	}
//
// otherwise will generate
//
//	if err != nil {
//	     return errors.Wrapf(err, <message>, <args>)
//	}
func CheckErrorAndWrap(errorsPackage string, message string, args ...dst.Expr) dst.Stmt {
	wrap := WrapError(errorsPackage, "err", message, args...)
	return CheckErrorAndSingleStatement(Returns(wrap))
}

// CheckErrorAndSingleStatement checks if the err is non-nil, and if it is executes the provided statement.
//
//	if err != nil {
//		<stmt>
//	}
func CheckErrorAndSingleStatement(stmt dst.Stmt) dst.Stmt {
	return &dst.IfStmt{
		Cond: NotNil(dst.NewIdent("err")),
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				stmt,
			},
		},
	}
}

// WrapError wraps the specified err.
// If no arguments are provided, will generate
//
// errors.Wrap(err, <message>)
//
// otherwise will generate
//
// errors.Wrapf(err, <message>, <args>)
func WrapError(errorsPackage string, err string, message string, args ...dst.Expr) dst.Expr {
	funcName := "Wrap"
	if len(args) > 0 {
		funcName = "Wrapf"
	}
	return CallQualifiedFunc(
		errorsPackage,
		funcName,
		Expressions(dst.NewIdent(err), StringLiteral(message), args)...)
}

// NewVariableQualified creates a new declaration statement where a variable is declared
// with its default value.
//
// For example:
//
//	var <varName> <packageRef>.<structName>
//
// Note that it does *not* do:
//
//	<varName> := <packageRef>.<structName>{}
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
//
//	var <varName> <structName>
//
// Note that it does *not* do:
//
//	<varName> := <structName>{}
//
// …as that does not work for enum types.
func NewVariable(varName string, structName string) dst.Stmt {
	return NewVariableWithType(varName, dst.NewIdent(structName))
}

func NewVariableWithType(varName string, varType dst.Expr) dst.Stmt {
	return &dst.DeclStmt{
		Decl: &dst.GenDecl{
			Tok: token.VAR,
			Specs: []dst.Spec{
				&dst.TypeSpec{
					Name: dst.NewIdent(varName),
					Type: varType,
				},
			},
		},
	}
}

// LocalVariableDeclaration performs a local variable declaration for use within a method
//
//	var <ident> <typ>
func LocalVariableDeclaration(ident string, typ dst.Expr, comment string) *dst.DeclStmt {
	return &dst.DeclStmt{
		Decl: VariableDeclaration(ident, typ, comment),
	}
}

// VariableDeclaration performs a global variable declaration
//
//	 // <comment>
//		var <ident> <typ>
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

	AddWrappedComment(&decl.Decs.Start, comment)

	return decl
}

// TypeAssert returns an assignment statement with a type assertion
//
//	<lhs>, ok := <rhs>.(<type>)
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
func ReturnIfOk(returns ...dst.Expr) *dst.IfStmt {
	return ReturnIfExpr(dst.NewIdent("ok"), returns...)
}

// ReturnIfNotOk checks a boolean ok variable and if it is not ok returns the specified values
//
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

// ReturnIfNil checks if a variable is nil and if it is returns
//
//	if <toCheck> == nil {
//		return <returns...>
//	}
func ReturnIfNil(toCheck dst.Expr, returns ...dst.Expr) dst.Stmt {
	return ReturnIfExpr(
		AreEqual(toCheck, Nil()),
		returns...)
}

// ReturnIfNotNil checks if a variable is not nil and if it is returns
//
//	if <toCheck> != nil {
//		return <returns...>
//	}
func ReturnIfNotNil(toCheck dst.Expr, returns ...dst.Expr) dst.Stmt {
	return ReturnIfExpr(
		AreNotEqual(toCheck, Nil()),
		returns...)
}

// ReturnIfExpr returns if the expression evaluates as true
//
//	if <cond> {
//		return <returns...>
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
					Results: Expressions(returns),
				},
			},
		},
	}
}

// FormatError produces a call to fmt.Errorf with the given format string and args
//
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
//
//	&<expr>
func AddrOf(expr dst.Expr) *dst.UnaryExpr {
	return &dst.UnaryExpr{
		Op: token.AND,
		X:  dst.Clone(expr).(dst.Expr),
	}
}

// AsReference returns a statement that is a reference to the supplied expression
//
//	&<expr>
//
// If the expression given is a StarExpr dereference, we unwrap it instead of taking its address
func AsReference(expr dst.Expr) dst.Expr {
	if star, ok := expr.(*dst.StarExpr); ok {
		return star.X
	}

	return AddrOf(expr)
}

// Dereference returns a statement that dereferences the pointer returned by the provided expression
//
// *<expr>
func Dereference(expr dst.Expr) dst.Expr {
	return &dst.StarExpr{
		X: dst.Clone(expr).(dst.Expr),
	}
}

// Returns creates a return statement with one or more expressions, of the form
//
//	return <expr>
//
// or return <expr>, <expr>, ...
func Returns(returns ...dst.Expr) dst.Stmt {
	return &dst.ReturnStmt{
		Decs: dst.ReturnStmtDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.NewLine,
			},
		},
		Results: Expressions(returns),
	}
}

// ReturnNoError creates a return nil statement for when no error occurs
//
//	// No error
//	return nil
func ReturnNoError() dst.Stmt {
	result := Returns(Nil())
	result.Decorations().Before = dst.EmptyLine
	result.Decorations().Start.Append("// No error")
	return result
}

// WrappedErrorf returns the err local, wrapped with additional information
//
// errors.Wrap(err, <message>)
//
// (actual package name will be used, which will usually be 'errors')
func WrappedErrorf(errorsPackage string, template string, args ...interface{}) dst.Expr {
	return CallQualifiedFunc(
		errorsPackage,
		"Wrap",
		dst.NewIdent("err"),
		StringLiteralf(template, args...))
}

// WrappedError returns the err local, wrapped with additional information
//
// errors.Wrap(err, <message>)
//
// (actual package name will be used, which will usually be 'errors')
func WrappedError(errorsPackage string, str string) dst.Expr {
	return CallQualifiedFunc(
		errorsPackage,
		"Wrap",
		dst.NewIdent("err"),
		StringLiteral(str))
}

// QualifiedTypeName generates a reference to a type within an imported package
//
// <pkg>.<name>
func QualifiedTypeName(pkg string, name string) *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X:   dst.NewIdent(pkg),
		Sel: dst.NewIdent(name),
	}
}

// Selector generates a field reference into an existing expression
//
// <expr>.<name0>.(<name1>.<name2>…)
func Selector(expr dst.Expr, names ...string) *dst.SelectorExpr {
	exprs := []dst.Expr{dst.Clone(expr).(dst.Expr)}
	for _, name := range names {
		exprs = append(exprs, dst.NewIdent(name))
	}

	return Reduce(
		func(l, r dst.Expr) dst.Expr {
			return &dst.SelectorExpr{X: l, Sel: r.(*dst.Ident)}
		},
		dst.None,
		exprs...).(*dst.SelectorExpr)
}

// AreEqual generates a == comparison between the two expressions
//
// <lhs> == <rhs>
func AreEqual(lhs dst.Expr, rhs dst.Expr) *dst.BinaryExpr {
	return BinaryExpr(lhs, token.EQL, rhs)
}

// AreNotEqual generates a != comparison between the two expressions
//
// <lhs> != <rhs>
func AreNotEqual(lhs dst.Expr, rhs dst.Expr) *dst.BinaryExpr {
	return BinaryExpr(lhs, token.NEQ, rhs)
}

// NotExpr generates a `!x` expression
func NotExpr(expr dst.Expr) *dst.UnaryExpr {
	return &dst.UnaryExpr{
		Op: token.NOT,
		X:  expr,
	}
}

// NotNil generates an `x != nil` condition
func NotNil(x dst.Expr) *dst.BinaryExpr {
	return AreNotEqual(x, Nil())
}

// Nil returns the nil identifier (not keyword!)
func Nil() *dst.Ident {
	return dst.NewIdent("nil")
}

// Continue returns the continue keyword
func Continue() dst.Stmt {
	return &dst.BranchStmt{
		Tok: token.CONTINUE,
	}
}

// NotEmpty generates an `len(x) > 0` condition
func NotEmpty(x dst.Expr) *dst.BinaryExpr {
	return &dst.BinaryExpr{
		X:  CallFunc("len", x),
		Op: token.GTR,
		Y:  dst.NewIdent("0"),
	}
}

// StatementBlock generates a block containing the supplied statements
// If we're given a single statement that's already a block, we won't double wrap it
func StatementBlock(statements ...dst.Stmt) *dst.BlockStmt {
	stmts := Statements(statements)

	if len(stmts) == 1 {
		if block, ok := stmts[0].(*dst.BlockStmt); ok {
			return block
		}
	}

	return &dst.BlockStmt{
		List: stmts,
	}
}

func BinaryExpr(lhs dst.Expr, op token.Token, rhs dst.Expr) *dst.BinaryExpr {
	return &dst.BinaryExpr{
		X:  dst.Clone(lhs).(dst.Expr),
		Op: op,
		Y:  dst.Clone(rhs).(dst.Expr),
	}
}

// Statements creates a sequence of statements from the provided values, each of which may be a
// single dst.Stmt or a slice of multiple []dst.Stmts
func Statements(statements ...interface{}) []dst.Stmt {
	var stmts []dst.Stmt
	for _, s := range statements {
		switch s := s.(type) {
		case nil:
			// Skip nils
			continue
		case dst.Stmt:
			// Add a single statement
			stmts = append(stmts, s)
		case []dst.Stmt:
			// Add many statements
			stmts = append(stmts, s...)
		default:
			panic(fmt.Sprintf("expected dst.Stmt or []dst.Stmt, but found %T", s))
		}
	}

	result := make([]dst.Stmt, 0, len(stmts))
	for _, st := range stmts {
		result = append(result, dst.Clone(st).(dst.Stmt))
	}

	return result
}

// Expression creates a sequence of expressions from the provided values, each of which may be a
// single dst.Expr or a slice of multiple []dst.Expr
func Expressions(statements ...interface{}) []dst.Expr {
	var exprs []dst.Expr
	for _, e := range statements {
		switch s := e.(type) {
		case nil:
			// Skip nils
			continue
		case dst.Expr:
			// Add a single expression
			exprs = append(exprs, s)
		case []dst.Expr:
			// Add many expressions
			exprs = append(exprs, s...)
		default:
			panic(fmt.Sprintf("expected dst.Expr or []dst.Expr, but found %T", s))
		}
	}

	result := make([]dst.Expr, 0, len(exprs))
	for _, ex := range exprs {
		result = append(result, dst.Clone(ex).(dst.Expr))
	}

	return result
}

// JoinOr combines a sequence of expressions with the OR (||) operator
// If there are more than two expressions to combine, they're broken out on separate lines
func JoinOr(exprs ...dst.Expr) dst.Expr {
	spaceType := dst.None
	if len(exprs) > 2 {
		spaceType = dst.NewLine
	}

	return JoinBinaryOp(token.LOR, spaceType, exprs...)
}

// JoinAnd combines a sequence of expressiosn with the AND (&&) operator
func JoinAnd(exprs ...dst.Expr) dst.Expr {
	return JoinBinaryOp(token.LAND, dst.None, exprs...)
}

// JoinBinaryOp combines a sequence of expressions using a given operator
// op defines the operator to use to combine two adjacent expressions.
// spaceType specifies what kind of whitespace is needed after each one.
// exprs is the sequence of expressions to combine.
func JoinBinaryOp(op token.Token, spaceType dst.SpaceType, exprs ...dst.Expr) dst.Expr {
	return Reduce(
		func(x, y dst.Expr) dst.Expr {
			return &dst.BinaryExpr{X: x, Op: op, Y: y}
		},
		spaceType,
		exprs...)
}

// Reduce combines a sequence of expressions using the provided function.
// operator defines how to combine two adjacent expressions.
// spaceType specifies what kind of whitespace is needed after each one.
// exprs is the sequence of expressions to combine.
func Reduce(operator func(l, r dst.Expr) dst.Expr, spaceType dst.SpaceType, exprs ...dst.Expr) dst.Expr {
	if len(exprs) == 0 {
		panic("must provide at least one expression to reduce")
	}

	result := exprs[0]
	for _, e := range exprs[1:] {
		e.Decorations().Before = spaceType
		result = operator(result, e)
	}

	return result
}
