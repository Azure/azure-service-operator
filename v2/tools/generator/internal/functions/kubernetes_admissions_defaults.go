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

func NewDefaultAzureNameFunction(resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *ResourceFunction {
	return &ResourceFunction{
		name:             "defaultAzureName",
		resource:         resource,
		idFactory:        idFactory,
		asFunc:           defaultAzureNameFunction,
		requiredPackages: astmodel.NewPackageReferenceSet(astmodel.GenRuntimeReference),
	}
}

// defaultAzureNameFunction returns a function that defaults the AzureName property of the resource spec
// to the Name property of the resource spec
func defaultAzureNameFunction(k *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.idFactory.CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	specSelector := &dst.SelectorExpr{
		X:   dst.NewIdent(receiverIdent),
		Sel: dst.NewIdent("Spec"),
	}

	azureNameProp := &dst.SelectorExpr{
		X:   specSelector,
		Sel: dst.NewIdent(astmodel.AzureNameProperty),
	}

	nameProp := &dst.SelectorExpr{
		X:   dst.NewIdent(receiverIdent),
		Sel: dst.NewIdent("Name"), // this comes from ObjectMeta
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Body: astbuilder.Statements(
			astbuilder.IfEqual(
				azureNameProp,
				astbuilder.StringLiteral(""),
				astbuilder.SimpleAssignment(azureNameProp, nameProp))),
	}

	fn.AddComments("defaults the Azure name of the resource to the Kubernetes name")
	return fn.DefineFunc()
}
