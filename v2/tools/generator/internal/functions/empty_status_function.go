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

// NewEmptyStatusFunction creates a new function to generate NewEmptyStatus() on resource types
func NewEmptyStatusFunction(
	status astmodel.TypeName,
	idFactory astmodel.IdentifierFactory,
) *ObjectFunction {
	result := NewObjectFunction(
		"NewEmptyStatus",
		idFactory,
		createNewEmptyStatusFunction(status))
	result.AddReferencedTypes(astmodel.ConvertibleStatusInterfaceType)
	return result
}

func createNewEmptyStatusFunction(
	statusType astmodel.TypeName) func(
	f *ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	_ string) (*dst.FuncDecl, error) {
	return func(f *ObjectFunction, genContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, _ string) (*dst.FuncDecl, error) {
		receiverIdent := f.IDFactory().CreateReceiver(receiver.Name())
		receiverType := astmodel.NewOptionalType(receiver)
		receiverTypeExpr, err := receiverType.AsTypeExpr(genContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating receiver expression for %s", receiverType)
		}

		// When Storage variants are created from resources, any existing functions are copied across - which means
		// the statusType we've been passed in above may be from the wrong package. We always want to use the status
		// type that's in the current package, never from a different one, so we just use the name, bypassing any
		// potential imports.

		literal := astbuilder.NewCompositeLiteralBuilder(dst.NewIdent(statusType.Name())).Build()

		fn := &astbuilder.FuncDetails{
			ReceiverIdent: receiverIdent,
			ReceiverType:  receiverTypeExpr,
			Name:          "NewEmptyStatus",
			Body: astbuilder.Statements(
				astbuilder.Returns(
					astbuilder.AddrOf(literal))),
		}

		convertibleStatusInterfaceExpr, err := astmodel.ConvertibleStatusInterfaceType.AsTypeExpr(genContext)
		if err != nil {
			return nil, eris.Wrap(err, "unable to create ConvertibleStatusInterfaceType expression")
		}

		fn.AddReturn(convertibleStatusInterfaceExpr)
		fn.AddComments("returns a new empty (blank) status")

		return fn.DefineFunc(), nil
	}
}
