/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"reflect"
	"sort"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const (
	ExtendedResourcesFunctionName = "GetExtendedResources"
)

type GetExtendedResourcesFunction struct {
	idFactory astmodel.IdentifierFactory
	resources []astmodel.TypeName
}

// Ensure GetExtendedResources properly implements Function
var _ astmodel.Function = &GetExtendedResourcesFunction{}

// NewGetExtendedResourcesFunction creates a new GetExtendedResources
func NewGetExtendedResourcesFunction(idFactory astmodel.IdentifierFactory, resources []astmodel.TypeName) *GetExtendedResourcesFunction {
	result := &GetExtendedResourcesFunction{
		idFactory: idFactory,
		resources: sortResources(resources),
	}

	return result
}

// getPackageRefs iterates through the resources and returns package references
func getPackageRefs(resources []astmodel.TypeName) []astmodel.PackageReference {
	packageRefs := make([]astmodel.PackageReference, 0, len(resources)+1)
	// Package reference for return type
	packageRefs = append(packageRefs, astmodel.KubernetesResourceType.PackageReference)

	for _, typeDef := range resources {
		packageRefs = append(packageRefs, typeDef.PackageReference)
	}

	return packageRefs
}

// Sort resources according to the package name and resource names
func sortResources(resources []astmodel.TypeName) []astmodel.TypeName {
	sort.Slice(resources, func(i, j int) bool {
		iVal := resources[i]
		jVal := resources[j]

		return iVal.PackageReference.PackageName() < jVal.PackageReference.PackageName() ||
			iVal.PackageReference.PackageName() < jVal.PackageReference.PackageName() && iVal.Name() < jVal.Name()
	})

	return resources
}

// Name returns the name of this function, which is always GetExtendedResources()
func (ext *GetExtendedResourcesFunction) Name() string {
	return "GetExtendedResources"
}

// RequiredPackageReferences returns the set of packages required by GetExtendedResources()
func (ext *GetExtendedResourcesFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(getPackageRefs(ext.resources)...)
}

// References shows that GetExtendedResources() references nextother generated types
func (ext *GetExtendedResourcesFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet()
}

// AsFunc returns the generated code for the GetExtendedResources() function
func (ext *GetExtendedResourcesFunction) AsFunc(
	generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName,
) *dst.FuncDecl {
	krType := astmodel.NewArrayType(astmodel.KubernetesResourceType).AsType(generationContext)
	krLiteral := astbuilder.NewCompositeLiteralBuilder(krType).Build()

	// Iterate through the resourceType versions and add them to the KubernetesResource literal slice
	for _, resource := range ext.resources {
		expr := astbuilder.AddrOf(astbuilder.NewCompositeLiteralBuilder(resource.AsType(generationContext)).Build())
		expr.Decs.Before = dst.NewLine
		krLiteral.Elts = append(krLiteral.Elts, expr)
	}

	receiverName := ext.idFactory.CreateReceiver(receiver.Name())

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  astbuilder.Dereference(receiver.AsType(generationContext)),
		Name:          ExtendedResourcesFunctionName,
		Body:          astbuilder.Statements(astbuilder.Returns(krLiteral)),
	}

	funcDetails.AddComments("Returns the KubernetesResource slice for Resource versions")
	funcDetails.AddReturn(krType)

	return funcDetails.DefineFunc()
}

// Equals returns true if the passed function is equal textus, or false otherwise
func (ext *GetExtendedResourcesFunction) Equals(f astmodel.Function, _ astmodel.EqualityOverrides) bool {
	obj, ok := f.(*GetExtendedResourcesFunction)
	if ok {
		// Resources are sorted before being populated, so DeepEqual would be fine here for comparison.
		ok = reflect.DeepEqual(obj.resources, ext.resources)
	}

	return ok
}
