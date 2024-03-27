/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// NewNewEmptyARMValueFunc returns a function that creates an empty value suitable for using with PopulateFromARM.
// It should be equivalent to ConvertToARM("") on a default struct value.
func NewNewEmptyARMValueFunc(
	armType astmodel.InternalTypeName,
	idFactory astmodel.IdentifierFactory,
) astmodel.Function {
	result := NewObjectFunction(
		"NewEmptyARMValue",
		idFactory,
		newEmptyARMValueBody(armType))

	result.AddPackageReference(armType.InternalPackageReference())

	return result
}

func newEmptyARMValueBody(instanceType astmodel.InternalTypeName) ObjectFunctionHandler {
	return func(
		fn *ObjectFunction,
		genContext *astmodel.CodeGenerationContext,
		receiver astmodel.TypeName,
		methodName string,
	) (*dst.FuncDecl, error) {
		receiverName := fn.IdFactory().CreateReceiver(receiver.Name())
		receiverTypeExpr, err := receiver.AsTypeExpr(genContext)
		if err != nil {
			return nil, errors.Wrapf(err, "creating type expression for %s", receiver)
		}
		receiverTypeExpr = astbuilder.PointerTo(receiverTypeExpr)

		// return &<pkg.instanceType>{}
		instanceTypeExpr, err := instanceType.AsTypeExpr(genContext)
		if err != nil {
			return nil, errors.Wrapf(err, "creating type expression for %s", instanceType)
		}

		instance := astbuilder.NewCompositeLiteralBuilder(instanceTypeExpr)
		returnInstance := astbuilder.Returns(astbuilder.AddrOf(instance.Build()))

		details := &astbuilder.FuncDetails{
			Name:          "NewEmptyARMValue",
			ReceiverIdent: receiverName,
			ReceiverType:  receiverTypeExpr,
			Body:          astbuilder.Statements(returnInstance),
		}

		armResourceStatusTypeExpr, err := astmodel.ARMResourceStatusType.AsTypeExpr(genContext)
		if err != nil {
			return nil, errors.Wrap(err, "creating ARM resource status type expression")
		}

		details.AddReturn(armResourceStatusTypeExpr)
		details.AddComments("returns an empty ARM value suitable for deserializing into")

		return details.DefineFunc(), nil
	}
}
