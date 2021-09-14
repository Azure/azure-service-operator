/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// NewCreateEmptyARMValueFunc returns a function that creates an empty value suitable for using with PopulateFromARM.
// It should be equivalent to ConvertToARM("") on a default struct value.
func NewCreateEmptyARMValueFunc(
	armType astmodel.TypeName,
	idFactory astmodel.IdentifierFactory) astmodel.Function {
	result := NewObjectFunction(
		"CreateEmptyARMValue",
		idFactory,
		createEmptyARMValueBody(armType))

	return result
}

func createEmptyARMValueBody(instanceType astmodel.TypeName) func(fn *ObjectFunction, genContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	return func(fn *ObjectFunction, genContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverName := fn.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
		receiverType := astbuilder.Dereference(receiver.AsType(genContext))
		instance := astbuilder.NewCompositeLiteralDetails(dst.NewIdent(instanceType.Name()))
		returnInstance := astbuilder.Returns(instance.Build())
		details := &astbuilder.FuncDetails{
			Name:          "CreateEmptyARMValue",
			ReceiverIdent: receiverName,
			ReceiverType:  receiverType,
			Body:          astbuilder.Statements(returnInstance),
		}

		details.AddReturns("interface{}")
		details.AddComments("returns an empty ARM value suitable for deserializing into")

		return details.DefineFunc()
	}
}
