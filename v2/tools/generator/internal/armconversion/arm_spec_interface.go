/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

func checkPropertyPresence(o *astmodel.ObjectType, name astmodel.PropertyName) error {
	_, ok := o.Property(name)
	if !ok {
		return errors.Errorf("resource spec doesn't have %q property", name)
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

	getNameFunc := functions.NewObjectFunction("Get"+astmodel.NameProperty, idFactory, getNameFunction)
	getNameFunc.AddPackageReference(astmodel.GenRuntimeReference)

	getTypeFunc := functions.NewGetTypeFunction(resource.ARMType(), idFactory, functions.ReceiverTypeStruct)

	getAPIVersionFunc := functions.NewGetAPIVersionFunction(
		resource.APIVersionTypeName(),
		resource.APIVersionEnumValue(),
		idFactory)

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
) *dst.FuncDecl {
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
) *dst.FuncDecl {
	receiverIdent := fn.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

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
		// TODO: We're too loosey-goosey here with ptr vs value receiver.
		// TODO: We basically need to use a value receiver right now because
		// TODO: ConvertToARM always returns a value, but for other interface impls
		// TODO: for example on resource we use ptr receiver... the inconsistency is
		// TODO: awkward...
		ReceiverType: receiverType,
		Body:         astbuilder.Statements(retResult),
	}

	details.AddComments(fmt.Sprintf("returns the %s of the resource", propertyName))
	details.AddReturns("string")

	return details.DefineFunc()
}
