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

// NewEmptyStatusFunction creates a new function to generate NewEmptyStatus() on resource types
func NewEmptyStatusFunction(
	status astmodel.TypeName,
	idFactory astmodel.IdentifierFactory) *ObjectFunction {
	result := NewObjectFunction("NewEmptyStatus", idFactory, createNewEmptyStatusFunction(status))
	result.AddReferencedTypes(astmodel.ConvertibleStatusInterfaceType)
	return result
}

func createNewEmptyStatusFunction(
	statusType astmodel.TypeName) func(
	f *ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	_ string) *dst.FuncDecl {
	return func(f *ObjectFunction, genContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, _ string) *dst.FuncDecl {
		receiverIdent := f.IdFactory().CreateReceiver(receiver.Name())
		receiverType := astmodel.NewOptionalType(receiver)

		// When Storage variants are created from resources, any existing functions are copied across - which means
		// the statusType we've been passed in above may be from the wrong package. We always want to use the status
		// type that's in the current package, never from a different one, so we just use the name, bypassing any
		// potential imports.

		literal := astbuilder.NewCompositeLiteralBuilder(dst.NewIdent(statusType.Name())).Build()

		fn := &astbuilder.FuncDetails{
			ReceiverIdent: receiverIdent,
			ReceiverType:  receiverType.AsType(genContext),
			Name:          "NewEmptyStatus",
			Body: astbuilder.Statements(
				astbuilder.Returns(
					astbuilder.AddrOf(literal))),
		}

		fn.AddReturn(astmodel.ConvertibleStatusInterfaceType.AsType(genContext))
		fn.AddComments("returns a new empty (blank) status")

		return fn.DefineFunc()
	}
}
