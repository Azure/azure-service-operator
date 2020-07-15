/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"
	"go/token"
)

// OneOfJSONMarshalFunction is a function for marshalling discriminated unions
// (types with only mutually exclusive properties) to JSON
type OneOfJSONMarshalFunction struct {
	oneOfStruct *StructType
	idFactory   IdentifierFactory // TODO: It's this or pass it in the AsFunc method
}

// NewOneOfJSONMarshalFunction creates a new OneOfJSONMarshalFunction struct
func NewOneOfJSONMarshalFunction(oneOfStruct *StructType, idFactory IdentifierFactory) *OneOfJSONMarshalFunction {
	return &OneOfJSONMarshalFunction{oneOfStruct, idFactory}
}

// Ensure OneOfJSONMarshalFunction implements Function interface correctly
var _ Function = (*OneOfJSONMarshalFunction)(nil)

// Equals determines if this function is equal to the passed in function
func (f *OneOfJSONMarshalFunction) Equals(other Function) bool {
	if o, ok := other.(*OneOfJSONMarshalFunction); ok {
		return f.oneOfStruct.Equals(o.oneOfStruct)
	}

	return false
}

// References returns the set of references for the underlying struct.
func (f *OneOfJSONMarshalFunction) References() TypeNameSet {
	// Defer this check to the owning struct as we only refer to its properties and it
	return f.oneOfStruct.References()
}

// AsFunc returns the function as a go ast
func (f *OneOfJSONMarshalFunction) AsFunc(
	codeGenerationContext *CodeGenerationContext,
	receiver *TypeName,
	methodName string) *ast.FuncDecl {

	receiverName := f.idFactory.CreateIdentifier(receiver.name, NotExported)

	header, _ := createComments(
		fmt.Sprintf(
			"%s defers JSON marshaling to the first non-nil property, because %s represents a discriminated union (JSON OneOf)",
			methodName,
			receiver.name))

	result := &ast.FuncDecl{
		Doc: &ast.CommentGroup{
			List: header,
		},
		Name: ast.NewIdent(methodName),
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Type:  receiver.AsType(codeGenerationContext),
					Names: []*ast.Ident{ast.NewIdent(receiverName)},
				},
			},
		},
		Type: &ast.FuncType{
			Func: token.HighestPrec,
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("[]byte"),
					},
					{
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
	}

	var statements []ast.Stmt

	for _, property := range f.oneOfStruct.Properties() {
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

	result.Body = &ast.BlockStmt{
		List: statements,
	}

	return result
}

// RequiredImports returns a list of packages required by this
func (f *OneOfJSONMarshalFunction) RequiredImports() []*PackageReference {
	return []*PackageReference{
		NewPackageReference("encoding/json"),
	}
}
