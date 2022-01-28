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

const (
	ExtendedResourcesFunctionName = "GetExtendedResources"
)

// NewGetExtendedResourcesFunction returns a function to generate GetExtendedResource() on resource types
func NewGetExtendedResourcesFunction(idFactory astmodel.IdentifierFactory) *ObjectFunction {
	result := NewObjectFunction(ExtendedResourcesFunctionName, idFactory, createGetExtendedResourcesFunction)
	result.AddReferencedTypes(astmodel.KubernetesResourceType)
	return result
}

func createGetExtendedResourcesFunction(
	f *ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	_ string) *dst.FuncDecl {

	krType := astmodel.NewArrayType(astmodel.KubernetesResourceType).AsType(genContext)
	receiverName := f.idFactory.CreateReceiver(receiver.Name())

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  astbuilder.Dereference(receiver.AsType(genContext)),
		Name:          f.Name(),
		Body:          astbuilder.Statements(astbuilder.Returns(astbuilder.AddrOf(astbuilder.NewCompositeLiteralDetails(krType).Build()))),
	}

	funcDetails.AddComments("returns the original API version used to create the resource.")
	funcDetails.AddReturn(astbuilder.Dereference(krType))

	return funcDetails.DefineFunc()
}
