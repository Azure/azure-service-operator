/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"go/ast"
	"go/token"
)

const JSONMarshalFunctionName string = "MarshalJSON"

// OneOfJSONMarshalFunction is a function for marshalling discriminated unions
// (types with only mutually exclusive properties) to JSON
type OneOfJSONMarshalFunction struct {
	oneOfObject *ObjectType
	idFactory   IdentifierFactory // TODO: It's this or pass it in the AsFunc method
}

// NewOneOfJSONMarshalFunction creates a new OneOfJSONMarshalFunction struct
func NewOneOfJSONMarshalFunction(oneOfObject *ObjectType, idFactory IdentifierFactory) *OneOfJSONMarshalFunction {
	return &OneOfJSONMarshalFunction{oneOfObject, idFactory}
}

// Ensure OneOfJSONMarshalFunction implements Function interface correctly
var _ Function = (*OneOfJSONMarshalFunction)(nil)

func (f *OneOfJSONMarshalFunction) Name() string {
	return JSONMarshalFunctionName
}

// Equals determines if this function is equal to the passed in function
func (f *OneOfJSONMarshalFunction) Equals(other Function) bool {
	if o, ok := other.(*OneOfJSONMarshalFunction); ok {
		return f.oneOfObject.Equals(o.oneOfObject)
	}

	return false
}

// References returns the set of references for the underlying object.
func (f *OneOfJSONMarshalFunction) References() TypeNameSet {
	// Defer this check to the owning object as we only refer to its properties and it
	return f.oneOfObject.References()
}

// AsFunc returns the function as a go ast
func (f *OneOfJSONMarshalFunction) AsFunc(
	codeGenerationContext *CodeGenerationContext,
	receiver TypeName) *ast.FuncDecl {

	receiverName := f.idFactory.CreateIdentifier(receiver.name, NotExported)

	var statements []ast.Stmt

	for _, property := range f.oneOfObject.Properties() {
		fieldSelectorExpr := &ast.SelectorExpr{
			X:   ast.NewIdent(receiverName),
			Sel: ast.NewIdent(string(property.propertyName)),
		}

		ifStatement := ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  fieldSelectorExpr,
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("json"),
									Sel: ast.NewIdent("Marshal"),
								},
								Args: []ast.Expr{
									fieldSelectorExpr,
								},
							},
						},
					},
				},
			},
		}

		statements = append(statements, &ifStatement)
	}

	finalReturnStatement := &ast.ReturnStmt{
		Results: []ast.Expr{
			ast.NewIdent("nil"),
			ast.NewIdent("nil"),
		},
	}
	statements = append(statements, finalReturnStatement)

	fn := &astbuilder.FuncDetails{
		Name:          ast.NewIdent(f.Name()),
		ReceiverIdent: ast.NewIdent(receiverName),
		ReceiverType:  receiver.AsType(codeGenerationContext),
		Body:          statements,
	}

	fn.AddComments(fmt.Sprintf(
		"defers JSON marshaling to the first non-nil property, because %s represents a discriminated union (JSON OneOf)",
		receiver.name))
	fn.AddReturns("[]byte", "error")
	return fn.DefineFunc()
}

// RequiredImports returns a list of packages required by this
func (f *OneOfJSONMarshalFunction) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet(MakeExternalPackageReference("encoding/json"))
}
