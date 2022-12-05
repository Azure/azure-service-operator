/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

type FuncDetails struct {
	ReceiverIdent string
	ReceiverType  dst.Expr
	Name          string
	Comments      []string
	Params        []*dst.Field
	Returns       []*dst.Field
	Body          []dst.Stmt
}

// NewTestFuncDetails returns a FuncDetails for a test method
// Tests require a particular signature, so this makes it simpler to create test functions
func NewTestFuncDetails(testingPackage string, testName string, body ...dst.Stmt) *FuncDetails {

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
		&dst.StarExpr{
			X: &dst.SelectorExpr{
				X:   dst.NewIdent(testingPackage),
				Sel: dst.NewIdent("T"),
			}},
	)

	return result
}

// DefineFunc defines a function (header, body, etc), like:
//
//	<comment>
//	func (<receiverIdent> <receiverType>) <name>(<params...>) (<returns...>) {
//		<body...>
//	}
func (fn *FuncDetails) DefineFunc() *dst.FuncDecl {

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
	var body []dst.Stmt
	for _, s := range fn.Body {
		if s != nil {
			body = append(body, s)
		}
	}

	var comment dst.Decorations
	if len(fn.Comments) > 0 {
		fn.Comments[0] = fmt.Sprintf("// %s %s", fn.Name, fn.Comments[0])
		AddComments(&comment, fn.Comments)
	}

	result := &dst.FuncDecl{
		Name: dst.NewIdent(fn.Name),
		Decs: dst.FuncDeclDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.EmptyLine,
				After:  dst.EmptyLine,
				Start:  comment,
			},
		},
		Type: &dst.FuncType{
			Params: &dst.FieldList{
				List: fn.Params,
			},
			Results: &dst.FieldList{
				List: fn.Returns,
			},
		},
		Body: &dst.BlockStmt{
			List: body,
		},
	}

	if fn.ReceiverIdent != "" {
		// We have a receiver, so include it

		field := &dst.Field{
			Names: []*dst.Ident{
				dst.NewIdent(fn.ReceiverIdent),
			},
			Type: fn.ReceiverType,
		}

		recv := dst.FieldList{
			List: []*dst.Field{field},
		}

		result.Recv = &recv
	}

	return result
}

// AddStatements adds additional statements to the function
func (fn *FuncDetails) AddStatements(statements ...dst.Stmt) {
	fn.Body = append(fn.Body, statements...)
}

// AddParameter adds another parameter to the function definition
func (fn *FuncDetails) AddParameter(id string, parameterType dst.Expr) {
	field := &dst.Field{
		Names: []*dst.Ident{dst.NewIdent(id)},
		Type:  parameterType,
	}
	fn.Params = append(fn.Params, field)
}

// AddReturns adds (possibly many) return values to the function definition
func (fn *FuncDetails) AddReturns(types ...string) {
	for _, t := range types {
		fn.AddReturn(dst.NewIdent(t))
	}
}

// AddReturn adds a single return value to the function definition
func (fn *FuncDetails) AddReturn(expr dst.Expr) {
	field := &dst.Field{
		Type: expr,
	}
	fn.Returns = append(fn.Returns, field)
}

// AddComments adds multiple comments to the function declaration
func (fn *FuncDetails) AddComments(comment ...string) {
	fn.Comments = append(fn.Comments, comment...)
}
