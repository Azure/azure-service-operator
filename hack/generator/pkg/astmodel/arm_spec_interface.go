/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/pkg/errors"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
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

// NewARMSpecInterfaceImpl creates a new interface implementation with the functions required to implement the
// genruntime.ARMResourceSpec interface
func NewARMSpecInterfaceImpl(
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
		MakeTypeName(GenRuntimeReference, "ARMResourceSpec"),
		getNameFunc,
		getTypeFunc,
		getApiVersionFunc)

	return result, nil
}

func getNameFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl {
	return armSpecInterfaceSimpleGetFunction(
		k,
		codeGenerationContext,
		receiver,
		methodName,
		"Name",
		false)
}

func getTypeFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl {
	return armSpecInterfaceSimpleGetFunction(
		k,
		codeGenerationContext,
		receiver,
		methodName,
		"Type",
		true)
}

func getApiVersionFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl {
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
	castToString bool) *dst.FuncDecl {

	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	var result dst.Expr
	result = &dst.SelectorExpr{
		X:   dst.NewIdent(receiverIdent),
		Sel: dst.NewIdent(propertyName),
	}

	// This is not the most beautiful thing but it saves some code.
	if castToString {
		result = &dst.CallExpr{
			Fun: dst.NewIdent("string"),
			Args: []dst.Expr{
				result,
			},
		}
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		// TODO: We're too loosey-goosey here with ptr vs value receiver.
		// TODO: We basically need to use a value receiver right now because
		// TODO: ConvertToARM always returns a value, but for other interface impls
		// TODO: for example on resource we use ptr receiver... the inconsistency is
		// TODO: awkward...
		ReceiverType: receiverType,
		Body: []dst.Stmt{
			&dst.ReturnStmt{
				Decs: dst.ReturnStmtDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.NewLine,
						After:  dst.NewLine,
					},
				},
				Results: []dst.Expr{
					result,
				},
			},
		},
	}

	fn.AddComments(fmt.Sprintf("returns the %s of the resource", propertyName))
	fn.AddReturns("string")

	return fn.DefineFunc()
}
