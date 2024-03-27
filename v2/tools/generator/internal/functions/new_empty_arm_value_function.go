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
		receiverType := astbuilder.PointerTo(receiver.AsType(genContext))

		// return &<pkg.instanceType>{}
		instance := astbuilder.NewCompositeLiteralBuilder(instanceType.AsType(genContext))
		returnInstance := astbuilder.Returns(astbuilder.AddrOf(instance.Build()))

		details := &astbuilder.FuncDetails{
			Name:          "NewEmptyARMValue",
			ReceiverIdent: receiverName,
			ReceiverType:  receiverType,
			Body:          astbuilder.Statements(returnInstance),
		}

		details.AddReturn(astmodel.ARMResourceStatusType.AsType(genContext))
		details.AddComments("returns an empty ARM value suitable for deserializing into")

		return details.DefineFunc(), nil
	}
}
