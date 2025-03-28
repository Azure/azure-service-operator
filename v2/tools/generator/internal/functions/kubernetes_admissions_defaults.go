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

// NewDefaultAzureNameFunction returns a function that defaults the AzureName property of the resource spec
// to the Name property of the resource spec.
//
//	func (r *<reciever>) defaultAzureName(ctx context.Context, resource *<resourceType>) error {
//		if resource.Spec.AzureName == "" {
//			resource.Spec.AzureName = resource.Name
//		}
//		return nil
//	}
func NewDefaultAzureNameFunction(resource astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory) *DefaultFunction {
	return NewDefaultFunction(
		"defaultAzureName",
		resource.Name(),
		idFactory,
		defaultAzureNameFunction,
		astmodel.GenRuntimeReference)
}

func defaultAzureNameFunction(
	k *DefaultFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	objIdent := "obj"
	contextIdent := "ctx"

	receiverIdent := k.idFactory.CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	azureNameProp := astbuilder.Selector(dst.NewIdent(objIdent), "Spec", astmodel.AzureNameProperty)
	nameProp := astbuilder.Selector(dst.NewIdent(objIdent), "Name")

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body: astbuilder.Statements(
			astbuilder.IfEqual(
				azureNameProp,
				astbuilder.StringLiteral(""),
				astbuilder.SimpleAssignment(azureNameProp, nameProp)),
			astbuilder.Returns(astbuilder.Nil())),
	}

	contextTypeExpr, err := astmodel.ContextType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating context type expression")
	}
	fn.AddParameter(contextIdent, contextTypeExpr)

	resourceTypeExpr, err := k.data.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating resource type expression")
	}
	fn.AddParameter(objIdent, astbuilder.PointerTo(resourceTypeExpr))
	fn.AddReturn(dst.NewIdent("error"))

	fn.AddComments("defaults the Azure name of the resource to the Kubernetes name")
	return fn.DefineFunc(), nil
}
