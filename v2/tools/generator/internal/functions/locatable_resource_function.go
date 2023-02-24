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

func NewLocatableResource(
	idFactory astmodel.IdentifierFactory,
	resourceType *astmodel.ResourceType,
) *astmodel.InterfaceImplementation {

	f := NewResourceFunction(
		"Location",
		resourceType,
		idFactory,
		locatableResourceLocationFunc,
		astmodel.NewPackageReferenceSet())

	return astmodel.NewInterfaceImplementation(astmodel.LocatableResourceInterfaceName, f)
}

// validateDelete returns a function that performs validation of deletion for the resource
func locatableResourceLocationFunc(k *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.idFactory.CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	locationSelector := astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec", "Location")
	returnIfLocationNil := astbuilder.ReturnIfNil(locationSelector, astbuilder.StringLiteral(""))
	returnLocation := astbuilder.Returns(astbuilder.Dereference(locationSelector))

	body := astbuilder.Statements(
		returnIfLocationNil,
		returnLocation)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiverType,
		},
		Returns: []*dst.Field{
			{
				Type: astmodel.StringType.AsType(codeGenerationContext),
			},
		},
		Body: body,
	}

	fn.AddComments("returns the location of the resource")
	return fn.DefineFunc()
}
