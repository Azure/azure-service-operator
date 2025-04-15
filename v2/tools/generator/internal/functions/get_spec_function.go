/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// NewGetSpecFunction returns a new function to GetSpec() on resource types
func NewGetSpecFunction(idFactory astmodel.IdentifierFactory) *ObjectFunction {
	result := NewObjectFunction("GetSpec", idFactory, createGetSpecFunction)
	result.AddReferencedTypes(astmodel.ConvertibleSpecInterfaceType)
	return result
}

func createGetSpecFunction(
	f *ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	_ string,
) (*dst.FuncDecl, error) {
	receiverIdent := f.IDFactory().CreateReceiver(receiver.Name())
	receiverType := astmodel.NewOptionalType(receiver)
	receiverTypeExpr, err := receiverType.AsTypeExpr(genContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	ret := astbuilder.Returns(astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec")))

	fn := &astbuilder.FuncDetails{
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverTypeExpr,
		Name:          "GetSpec",
		Body:          astbuilder.Statements(ret),
	}

	convertibleSpecInterfaceExpr, err := astmodel.ConvertibleSpecInterfaceType.AsTypeExpr(genContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating type expression for ConvertibleSpecInterface")
	}

	fn.AddReturn(convertibleSpecInterfaceExpr)
	fn.AddComments("returns the specification of this resource")

	return fn.DefineFunc(), nil
}
