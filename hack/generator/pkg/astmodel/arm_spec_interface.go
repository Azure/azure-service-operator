/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	ast "github.com/dave/dst"
	"github.com/pkg/errors"
)

const (
	ApiVersionProperty = "ApiVersion"
	TypeProperty       = "Type"
	NameProperty       = "Name"
)

func checkPropertyPresence(o *ObjectType, name PropertyName) error {
	_, ok := o.Property(name)
	if !ok {
		return errors.Errorf("Resource spec doesn't have %q property", name)
	}

	return nil
}

// NewArmSpecInterfaceImpl creates a new interface implementation with the functions required to implement the
// genruntime.ArmResourceSpec interface
func NewArmSpecInterfaceImpl(
	idFactory IdentifierFactory,
	spec *ObjectType) (*InterfaceImplementation, error) {

	// Check the spec first to ensure it looks how we expect
	apiVersionProperty := idFactory.CreatePropertyName(ApiVersionProperty, Exported)
	err := checkPropertyPresence(spec, apiVersionProperty)
	if err != nil {
		return nil, err
	}
	typeProperty := idFactory.CreatePropertyName(TypeProperty, Exported)
	err = checkPropertyPresence(spec, typeProperty)
	if err != nil {
		return nil, err
	}
	nameProperty := idFactory.CreatePropertyName(NameProperty, Exported)
	err = checkPropertyPresence(spec, nameProperty)
	if err != nil {
		return nil, err
	}

	getNameFunc := &objectFunction{
		name:      "GetName",
		o:         spec,
		idFactory: idFactory,
		asFunc:    getNameFunction,
	}

	getTypeFunc := &objectFunction{
		name:      "GetType",
		o:         spec,
		idFactory: idFactory,
		asFunc:    getTypeFunction,
	}

	getApiVersionFunc := &objectFunction{
		name:      "GetApiVersion",
		o:         spec,
		idFactory: idFactory,
		asFunc:    getApiVersionFunction,
	}

	result := NewInterfaceImplementation(
		MakeTypeName(MakeGenRuntimePackageReference(), "ArmResourceSpec"),
		getNameFunc,
		getTypeFunc,
		getApiVersionFunc)

	return result, nil
}

func getNameFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
	return armSpecInterfaceSimpleGetFunction(
		k,
		codeGenerationContext,
		receiver,
		methodName,
		"Name",
		false)
}

func getTypeFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
	return armSpecInterfaceSimpleGetFunction(
		k,
		codeGenerationContext,
		receiver,
		methodName,
		"Type",
		true)
}

func getApiVersionFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
	return armSpecInterfaceSimpleGetFunction(
		k,
		codeGenerationContext,
		receiver,
		methodName,
		"ApiVersion",
		true)
}

func armSpecInterfaceSimpleGetFunction(
	k *objectFunction,
	codeGenerationContext *CodeGenerationContext,
	receiver TypeName,
	methodName string,
	propertyName string,
	castToString bool) *ast.FuncDecl {

	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	var result ast.Expr
	result = &ast.SelectorExpr{
		X:   ast.NewIdent(receiverIdent),
		Sel: ast.NewIdent(propertyName),
	}

	// This is not the most beautiful thing but it saves some code.
	if castToString {
		result = &ast.CallExpr{
			Fun: ast.NewIdent("string"),
			Args: []ast.Expr{
				result,
			},
		}
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		// TODO: We're too loosey-goosey here with ptr vs value receiver.
		// TODO: We basically need to use a value receiver right now because
		// TODO: ConvertToArm always returns a value, but for other interface impls
		// TODO: for example on resource we use ptr receiver... the inconsistency is
		// TODO: awkward...
		ReceiverType: receiverType,
		Body: []ast.Stmt{
			&ast.ReturnStmt{
				Decs: ast.ReturnStmtDecorations{
					NodeDecs: ast.NodeDecs{
						Before: ast.NewLine,
						After:  ast.NewLine,
					},
				},
				Results: []ast.Expr{
					result,
				},
			},
		},
	}

	fn.AddComments(fmt.Sprintf("returns the %s of the resource", propertyName))
	fn.AddReturns("string")

	return fn.DefineFunc()
}
