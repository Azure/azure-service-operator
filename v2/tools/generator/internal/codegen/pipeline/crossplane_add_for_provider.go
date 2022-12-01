/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// AddCrossplaneForProvider adds a "ForProvider" property as the sole property in every resource spec
// and moves everything that was at the spec level down a level into the ForProvider type
func AddCrossplaneForProvider(idFactory astmodel.IdentifierFactory) *Stage {
	return NewLegacyStage(
		"addCrossplaneForProviderProperty",
		"Add a 'ForProvider' property on every spec",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			result := make(astmodel.TypeDefinitionSet)
			for _, typeDef := range definitions {
				if _, ok := astmodel.AsResourceType(typeDef.Type()); ok {
					forProviderTypes, err := nestSpecIntoForProvider(
						idFactory, definitions, typeDef)
					if err != nil {
						return nil, errors.Wrapf(err, "creating 'ForProvider' definitions")
					}

					result.AddAll(forProviderTypes...)
				}
			}

			unmodified := definitions.Except(result)
			result.AddTypes(unmodified)

			return result, nil
		})
}

// nestSpecIntoForProvider returns the type definitions required to nest the contents of the "Spec" type
// into a property named "ForProvider" whose type is "<name>Parameters"
func nestSpecIntoForProvider(
	idFactory astmodel.IdentifierFactory,
	definitions astmodel.TypeDefinitionSet,
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

	// In the case where a spec type is reused across multiple resource definitions, we need to make sure
	// to generate the same names for all of their nested properties, so base the nested type name off the
	// spec type name
	nestedTypeName := strings.TrimSuffix(specName.Name(), astmodel.SpecSuffix) + "Parameters"
	nestedPropertyName := "ForProvider"
	return nestType(idFactory, definitions, specName, nestedTypeName, nestedPropertyName)
}

// nestType nests the contents of the provided outerType into a property with the given nestedPropertyName whose
// type is the given nestedTypeName.
func nestType(
	idFactory astmodel.IdentifierFactory,
	definitions astmodel.TypeDefinitionSet,
	outerTypeName astmodel.TypeName,
	nestedTypeName string,
	nestedPropertyName string) ([]astmodel.TypeDefinition, error) {
	/*
	 * Sample output:
	 *
	 * type <outerTypeName> struct {
	 *     <nestedPropertyName> <nestedTypeName> `yaml:"<nestedPropertyName>"`
	 * }
	 *
	 */

	outerType, ok := definitions[outerTypeName]
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
