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

// NewGetStatusFunction returns a function to generate GetStatus() on resource types
func NewGetStatusFunction(idFactory astmodel.IdentifierFactory) *ObjectFunction {
	result := NewObjectFunction("GetStatus", idFactory, createGetStatusFunction)
	result.AddReferencedTypes(astmodel.ConvertibleStatusInterfaceType)
	return result
}

func createGetStatusFunction(
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

	fn := &astbuilder.FuncDetails{
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverTypeExpr,
		Name:          "GetStatus",
		Body: astbuilder.Statements(
			astbuilder.Returns(
				astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(receiverIdent), "Status")))),
	}

	convertibleStatusInterfaceExpr, err := astmodel.ConvertibleStatusInterfaceType.AsTypeExpr(genContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating type expression for ConvertibleStatusInterface")
	}

	fn.AddReturn(convertibleStatusInterfaceExpr)
	fn.AddComments("returns the status of this resource")

	return fn.DefineFunc(), nil
}
