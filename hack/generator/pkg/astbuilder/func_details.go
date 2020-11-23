/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"fmt"
	"strings"

	ast "github.com/dave/dst"
)

type FuncDetails struct {
	ReceiverIdent string
	ReceiverType  ast.Expr
	Name          string
	Comments      []string
	Params        []*ast.Field
	Returns       []*ast.Field
	Body          []ast.Stmt
}

// NewTestFuncDetails() returns a FuncDetails for a test method
// Tests require a particular signature, so this makes it simpler to create test functions
func NewTestFuncDetails(testName string, body ...ast.Stmt) *FuncDetails {

	// Ensure the method name starts with `Test` as required
	var name string
	if strings.HasPrefix(testName, "Test") {
		name = testName
	} else {
		name = "Test_" + testName
	}

	result := &FuncDetails{
		Name: name,
		Body: body,
	}

	result.AddParameter("t",
		&ast.StarExpr{
			X: &ast.SelectorExpr{
				X:   ast.NewIdent("testing"),
				Sel: ast.NewIdent("T"),
			}},
	)

	return result
}

// DefineFunc defines a function (header, body, etc), like:
// 	<comment>
//	func (<receiverIdent> <receiverType>) <name>(<params...>) (<returns...>) {
// 		<body...>
//	}
func (fn *FuncDetails) DefineFunc() *ast.FuncDecl {

	// Safety check that we are making something valid
	if (fn.ReceiverIdent == "") != (fn.ReceiverType == nil) {
		reason := fmt.Sprintf(
			"ReceiverIdent and ReceiverType must both be specified, or both omitted. ReceiverIdent: %q, ReceiverType: %q",
			fn.ReceiverIdent,
			fn.ReceiverType)
		panic(reason)
	}

	// Filter out any nil statements
	// this helps creation of the fn go simpler
	var body []ast.Stmt
	for _, s := range fn.Body {
		if s != nil {
			body = append(body, s)
		}
	}

	var comment ast.Decorations
	if len(fn.Comments) > 0 {
		fn.Comments[0] = fmt.Sprintf("// %s %s", fn.Name, fn.Comments[0])
		AddComments(&comment, fn.Comments)
	}

	result := &ast.FuncDecl{
		Name: ast.NewIdent(fn.Name),
		Decs: ast.FuncDeclDecorations{
			NodeDecs: ast.NodeDecs{
				Before: ast.EmptyLine,
				After:  ast.EmptyLine,
				Start:  comment,
			},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: fn.Params,
			},
			Results: &ast.FieldList{
				List: fn.Returns,
			},
		},
		Body: &ast.BlockStmt{
			List: body,
		},
	}

	if fn.ReceiverIdent != "" {
		// We have a receiver, so include it

		field := &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent(fn.ReceiverIdent),
			},
			Type: fn.ReceiverType,
		}

		recv := ast.FieldList{
			List: []*ast.Field{field},
		}

		result.Recv = &recv
	}

	return result
}

// AddStatements adds additional statements to the function
func (fn *FuncDetails) AddStatements(statements ...ast.Stmt) {
	fn.Body = append(fn.Body, statements...)
}

// AddParameter adds another parameter to the function definition
func (fn *FuncDetails) AddParameter(id string, parameterType ast.Expr) {
	field := &ast.Field{
		Names: []*ast.Ident{ast.NewIdent(id)},
		Type:  parameterType,
	}
	fn.Params = append(fn.Params, field)
}

// AddReturns adds (possibly many) return values to the function definition
func (fn *FuncDetails) AddReturns(types ...string) {
	for _, t := range types {
		field := &ast.Field{
			Type: ast.NewIdent(t),
		}
		fn.Returns = append(fn.Returns, field)
	}
}

// AddComments() adds multiple comments to the function declaration
func (fn *FuncDetails) AddComments(comment ...string) {
	fn.Comments = append(fn.Comments, comment...)
}
