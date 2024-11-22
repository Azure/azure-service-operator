/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

func checkPropertyPresence(o *astmodel.ObjectType, name astmodel.PropertyName) error {
	_, ok := o.Property(name)
	if !ok {
		return eris.Errorf("resource spec doesn't have %q property", name)
	}

	return nil
}

// NewARMSpecInterfaceImpl creates a new interface implementation with the functions required to implement the
// genruntime.ARMResourceSpec interface
func NewARMSpecInterfaceImpl(
	idFactory astmodel.IdentifierFactory,
	resource *astmodel.ResourceType,
	spec *astmodel.ObjectType,
) (*astmodel.InterfaceImplementation, error) {
	nameProperty := idFactory.CreatePropertyName(astmodel.NameProperty, astmodel.Exported)
	err := checkPropertyPresence(spec, nameProperty)
	if err != nil {
		return nil, err
	}

	getNameFunc := functions.NewObjectFunction(
		"Get"+astmodel.NameProperty,
		idFactory,
		getNameFunction)
	getNameFunc.AddPackageReference(astmodel.GenRuntimeReference)

	getTypeFunc := functions.NewGetTypeFunction(resource.ARMType(), idFactory, functions.ReceiverTypePtr)

	getAPIVersionFunc := functions.NewGetAPIVersionFunction(resource.APIVersionEnumValue(), idFactory)

	result := astmodel.NewInterfaceImplementation(
		astmodel.ARMResourceSpecType,
		getNameFunc,
		getTypeFunc,
		getAPIVersionFunc)

	return result, nil
}

func getNameFunction(
	fn *functions.ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	return armSpecInterfaceSimpleGetFunction(
		fn,
		genContext,
		receiver,
		methodName,
		"Name",
		false)
}

func armSpecInterfaceSimpleGetFunction(
	fn *functions.ObjectFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
	propertyName string,
	castToString bool,
) (*dst.FuncDecl, error) {
	receiverIdent := fn.IdFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", receiver.Name())
	}

	var result dst.Expr = astbuilder.Selector(dst.NewIdent(receiverIdent), propertyName)

	// This is not the most beautiful thing, but it saves some code.
	if castToString {
		result = astbuilder.CallFunc("string", result)
	}

	retResult := astbuilder.Returns(result)
	retResult.Decorations().Before = dst.NewLine
	retResult.Decorations().After = dst.NewLine

	details := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.Dereference(receiverExpr),
		Body:          astbuilder.Statements(retResult),
	}

	details.AddComments(fmt.Sprintf("returns the %s of the resource", propertyName))
	details.AddReturns("string")

	return details.DefineFunc(), nil
}
