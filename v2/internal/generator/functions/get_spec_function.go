/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astbuilder"
	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
)

func NewGetSpecFunction(idFactory astmodel.IdentifierFactory) *ObjectFunction {
	result := NewObjectFunction("GetSpec", idFactory, createGetSpecFunction)
	result.AddReferencedTypes(astmodel.ConvertibleSpecInterfaceType)
	return result
}

func createGetSpecFunction(
	f *ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	_ string) *dst.FuncDecl {
	receiverIdent := f.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
	receiverType := astmodel.NewOptionalType(receiver)

	ret := astbuilder.Returns(astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec")))

	fn := &astbuilder.FuncDetails{
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverType.AsType(genContext),
		Name:          "GetSpec",
		Body:          astbuilder.Statements(ret),
	}

	fn.AddReturn(astmodel.ConvertibleSpecInterfaceType.AsType(genContext))
	fn.AddComments("returns the specification of this resource")

	return fn.DefineFunc()
}
