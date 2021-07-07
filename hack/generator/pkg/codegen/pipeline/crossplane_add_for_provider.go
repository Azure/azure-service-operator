/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// AddCrossplaneForProvider adds a "ForProvider" property as the sole property in every resource spec
// and moves everything that was at the spec level down a level into the ForProvider type
func AddCrossplaneForProvider(idFactory astmodel.IdentifierFactory) Stage {

	return MakeStage(
		"addCrossplaneForProviderProperty",
		"Add a 'ForProvider' property on every spec",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			result := make(astmodel.Types)
			for _, typeDef := range types {
				if _, ok := astmodel.AsResourceType(typeDef.Type()); ok {
					forProviderTypes, err := nestSpecIntoForProvider(
						idFactory, types, typeDef)
					if err != nil {
						return nil, errors.Wrapf(err, "creating 'ForProvider' types")
					}

					result.AddAll(forProviderTypes...)
				}
			}

			unmodified := types.Except(result)
			result.AddTypes(unmodified)

			return result, nil
		})
}

// nestSpecIntoForProvider returns the type definitions required to nest the contents of the "Spec" type
// into a property named "ForProvider" whose type is "<name>Parameters"
func nestSpecIntoForProvider(
	idFactory astmodel.IdentifierFactory,
	types astmodel.Types,
	typeDef astmodel.TypeDefinition) ([]astmodel.TypeDefinition, error) {

	resource, ok := astmodel.AsResourceType(typeDef.Type())
	if !ok {
		return nil, errors.Errorf("provided typeDef was not a resourceType, instead %T", typeDef.Type())
	}
	resourceName := typeDef.Name()

	specName, ok := astmodel.AsTypeName(resource.SpecType())
	if !ok {
		return nil, errors.Errorf("resource %q spec was not of type TypeName, instead: %T", resourceName, resource.SpecType())
	}

	// In the case where a spec type is reused across multiple resource types, we need to make sure
	// to generate the same names for all of their nested properties, so base the nested type name off the
	// spec type name
	nestedTypeName := strings.Split(specName.Name(), "_")[0] + "Parameters"
	nestedPropertyName := "ForProvider"
	return nestType(idFactory, types, specName, nestedTypeName, nestedPropertyName)
}

// nestType nests the contents of the provided outerType into a property with the given nestedPropertyName whose
// type is the given nestedTypeName. The result is a type that looks something like the following:
//
// type <outerTypeName> struct {
//     <nestedPropertyName> <nestedTypeName> `yaml:"<nestedPropertyName>"`
// }
func nestType(
	idFactory astmodel.IdentifierFactory,
	types astmodel.Types,
	outerTypeName astmodel.TypeName,
	nestedTypeName string,
	nestedPropertyName string) ([]astmodel.TypeDefinition, error) {

	outerType, ok := types[outerTypeName]
	if !ok {
		return nil, errors.Errorf("couldn't find type %q", outerTypeName)
	}

	outerObject, ok := astmodel.AsObjectType(outerType.Type())
	if !ok {
		return nil, errors.Errorf("type %q was not of type ObjectType, instead %T", outerTypeName, outerType.Type())
	}

	var result []astmodel.TypeDefinition

	// Copy outer type properties onto new "nesting type" with name nestedTypeName
	nestedDef := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(outerTypeName.PackageReference, nestedTypeName),
		outerObject)
	result = append(result, nestedDef)

	nestedProperty := astmodel.NewPropertyDefinition(
		idFactory.CreatePropertyName(nestedPropertyName, astmodel.Exported),
		idFactory.CreateIdentifier(nestedPropertyName, astmodel.NotExported),
		nestedDef.Name())

	// Change existing object type to have a single property pointing to the above nested type
	updatedObject := outerObject.WithoutProperties().WithProperty(nestedProperty)
	result = append(result, astmodel.MakeTypeDefinition(outerTypeName, updatedObject))

	return result, nil
}
