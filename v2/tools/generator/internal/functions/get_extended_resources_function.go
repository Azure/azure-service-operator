/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"reflect"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"
	"golang.org/x/exp/slices"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const (
	ExtendedResourcesFunctionName = "GetExtendedResources"
)

type GetExtendedResourcesFunction struct {
	idFactory astmodel.IdentifierFactory
	resources []astmodel.InternalTypeName
}

// Ensure GetExtendedResources properly implements Function
var _ astmodel.Function = &GetExtendedResourcesFunction{}

// NewGetExtendedResourcesFunction creates a new GetExtendedResources
func NewGetExtendedResourcesFunction(
	idFactory astmodel.IdentifierFactory,
	resources []astmodel.InternalTypeName,
) *GetExtendedResourcesFunction {
	result := &GetExtendedResourcesFunction{
		idFactory: idFactory,
		resources: sortResources(resources),
	}

	return result
}

// getPackageRefs iterates through the resources and returns package references
func getPackageRefs(resources []astmodel.InternalTypeName) []astmodel.PackageReference {
	packageRefs := make([]astmodel.PackageReference, 0, len(resources)+1)
	// Package reference for return type
	packageRefs = append(packageRefs, astmodel.KubernetesResourceType.PackageReference())

	for _, typeDef := range resources {
		packageRefs = append(packageRefs, typeDef.PackageReference())
	}

	return packageRefs
}

// Sort resources according to the package name and resource names
func sortResources(resources []astmodel.InternalTypeName) []astmodel.InternalTypeName {
	slices.SortFunc(
		resources,
		func(left astmodel.InternalTypeName, right astmodel.InternalTypeName) int {
			if left.InternalPackageReference().FolderPath() < right.InternalPackageReference().FolderPath() {
				return -1
			} else if left.InternalPackageReference().FolderPath() > right.InternalPackageReference().FolderPath() {
				return 1
			}

			if left.Name() < right.Name() {
				return -1
			} else if left.Name() > right.Name() {
				return 1
			}

			return 0
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

// References shows that GetExtendedResources() references other generated types
func (ext *GetExtendedResourcesFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet()
}

// AsFunc returns the generated code for the GetExtendedResources() function
func (ext *GetExtendedResourcesFunction) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.InternalTypeName,
) (*dst.FuncDecl, error) {
	krType, err := astmodel.NewArrayType(astmodel.KubernetesResourceType).AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", astmodel.KubernetesResourceType)
	}

	krLiteral := astbuilder.NewCompositeLiteralBuilder(krType).Build()

	// Iterate through the resourceType versions and add them to the KubernetesResource literal slice
	for _, resource := range ext.resources {
		var resourceExpr dst.Expr
		resourceExpr, err = resource.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating type expression for %s", resource)
		}

		expr := astbuilder.AddrOf(astbuilder.NewCompositeLiteralBuilder(resourceExpr).Build())
		expr.Decs.Before = dst.NewLine
		krLiteral.Elts = append(krLiteral.Elts, expr)
	}

	receiverName := ext.idFactory.CreateReceiver(receiver.Name())
	receiverType, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", receiver)
	}

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Name:          ExtendedResourcesFunctionName,
		Body:          astbuilder.Statements(astbuilder.Returns(krLiteral)),
	}

	funcDetails.AddComments("Returns the KubernetesResource slice for Resource versions")
	funcDetails.AddReturn(krType)

	return funcDetails.DefineFunc(), nil
}

// Equals returns true if the passed function is equal to us, or false otherwise
func (ext *GetExtendedResourcesFunction) Equals(f astmodel.Function, _ astmodel.EqualityOverrides) bool {
	obj, ok := f.(*GetExtendedResourcesFunction)
	if ok {
		// Resources are sorted before being populated, so DeepEqual would be fine here for comparison.
		ok = reflect.DeepEqual(obj.resources, ext.resources)
	}

	return ok
}
