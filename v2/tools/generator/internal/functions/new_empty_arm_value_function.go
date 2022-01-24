/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// NewNewEmptyARMValueFunc returns a function that creates an empty value suitable for using with PopulateFromARM.
// It should be equivalent to ConvertToARM("") on a default struct value.
func NewNewEmptyARMValueFunc(
	armType astmodel.TypeName,
	idFactory astmodel.IdentifierFactory) astmodel.Function {
	result := NewObjectFunction(
		"NewEmptyARMValue",
		idFactory,
		newEmptyARMValueBody(armType))

	return result
}

func newEmptyARMValueBody(instanceType astmodel.TypeName) func(fn *ObjectFunction, genContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	return func(fn *ObjectFunction, genContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverName := fn.IdFactory().CreateReceiver(receiver.Name())
		receiverType := astbuilder.Dereference(receiver.AsType(genContext))
		instance := astbuilder.NewCompositeLiteralBuilder(dst.NewIdent(instanceType.Name()))
		returnInstance := astbuilder.Returns(astbuilder.AddrOf(instance.Build()))
		details := &astbuilder.FuncDetails{
			Name:          "NewEmptyARMValue",
			ReceiverIdent: receiverName,
			ReceiverType:  receiverType,
			Body:          astbuilder.Statements(returnInstance),
		}

		details.AddReturn(astmodel.ARMResourceStatusType.AsType(genContext))
		details.AddComments("returns an empty ARM value suitable for deserializing into")

		return details.DefineFunc()
	}
}
