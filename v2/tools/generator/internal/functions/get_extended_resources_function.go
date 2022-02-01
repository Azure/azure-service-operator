/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/dave/dst"
	"sort"
)

const (
	ExtendedResourcesFunctionName = "GetExtendedResources"
)

type GetExtendedResourcesFunction struct {
	idFactory  astmodel.IdentifierFactory
	resources  []astmodel.TypeName
	packageRef *astmodel.PackageReferenceSet
}

// Ensure GetExtendedResources properly implements Function
var _ astmodel.Function = &GetExtendedResourcesFunction{}

// NewGetExtendedResourcesFunction creates a new GetExtendedResources
func NewGetExtendedResourcesFunction(idFactory astmodel.IdentifierFactory, resources []astmodel.TypeName, references []astmodel.PackageReference) *GetExtendedResourcesFunction {

	sortResources(resources)
	result := &GetExtendedResourcesFunction{
		idFactory:  idFactory,
		resources:  sortResources(resources),
		packageRef: astmodel.NewPackageReferenceSet(append(references, astmodel.KubernetesResourceType.PackageReference)...),
	}
	return result
}

//Sort resources according to the package name and resource names
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
	return ext.packageRef
}

// References shows that GetExtendedResources() references nextother generated types
func (ext *GetExtendedResourcesFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet()
}

// AddPackageReference adds one or more required package references
func (ext *GetExtendedResourcesFunction) AddPackageReference(refs ...astmodel.PackageReference) {
	for _, ref := range refs {
		ext.packageRef.AddReference(ref)
	}
}

// AsFunc returns the generated code for the GetExtendedResources() function
func (ext *GetExtendedResourcesFunction) AsFunc(
	generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {

	krType := astmodel.NewArrayType(astmodel.KubernetesResourceType).AsType(generationContext)
	krLiteral := astbuilder.NewCompositeLiteralDetails(krType).Build()

	//iterate through the resourceType versions and add them to the KubernetesResource literal slice
	for _, resource := range ext.resources {
		expr := astbuilder.AddrOf(astbuilder.NewCompositeLiteralDetails(resource.AsType(generationContext)).Build())
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
	_, ok := f.(*GetExtendedResourcesFunction)
	// Equality is just based on Type for now
	return ok
}
