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

// NewGetStatusFunction returns a function to generate GetStatus() on resource types
func NewGetStatusFunction(idFactory astmodel.IdentifierFactory) *ObjectFunction {
	result := NewObjectFunction("GetStatus", idFactory, createGetStatusFunction)
	result.AddReferencedTypes(astmodel.ConvertibleStatusInterfaceType)
	return result
}

func createGetStatusFunction(
	f *ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.InternalTypeName,
	_ string) *dst.FuncDecl {
	receiverIdent := f.IdFactory().CreateReceiver(receiver.Name())
	receiverType := astmodel.NewOptionalType(receiver)

	fn := &astbuilder.FuncDetails{
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverType.AsType(genContext),
		Name:          "GetStatus",
		Body: astbuilder.Statements(
			astbuilder.Returns(
				astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(receiverIdent), "Status")))),
	}

	fn.AddReturn(astmodel.ConvertibleStatusInterfaceType.AsType(genContext))
	fn.AddComments("returns the status of this resource")

	return fn.DefineFunc()
}
