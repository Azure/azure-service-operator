/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// RemoveTypePropertyStageID is the unique identifier for this pipeline stage
const RemoveTypePropertyStageID = "removeTypeProperty"

func RemoveTypeProperty() *Stage {
	return NewStage(
		RemoveTypePropertyStageID,
		"Remove the ARM type property and instead augment the ResourceType with it",
		func(ctx context.Context, state *State) (*State, error) {

			newDefs := make(astmodel.TypeDefinitionSet)

			defs := state.Definitions()
			resources := astmodel.FindResourceDefinitions(defs)

			for _, resource := range resources {
				resolved, err := state.Definitions().ResolveResourceSpecAndStatus(resource)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to resolve resource spec and status types")
				}

				typeProp, ok := resolved.SpecType.Property(astmodel.TypeProperty)
				if !ok {
					return nil, errors.Errorf("resource %s is missing type property", resolved.ResourceDef.Name())
				}

				armType, err := extractPropertySingleEnumValue(defs, typeProp)
				if err != nil {
					return nil, errors.Wrapf(err, "error extracting %s type property", resolved.SpecDef.Name())
				}

				resourceType := resolved.ResourceType.WithARMType(armType.Value)
				// Remove the Type property from the resource now, we don't need it anymore
				specType := resolved.SpecType.WithoutProperty(astmodel.TypeProperty)

				newDefs.Add(resolved.ResourceDef.WithType(resourceType))
				newDefs.Add(resolved.SpecDef.WithType(specType))
			}

			return state.WithDefinitions(defs.OverlayWith(newDefs)), nil
		})
}

// extractPropertySingleEnumValue returns the enum id and value for a property that is an enum with a single value.
// Any other type of property results in an error. An enum with more than a single value results in an error.
func extractPropertySingleEnumValue(definitions astmodel.TypeDefinitionSet, prop *astmodel.PropertyDefinition) (astmodel.EnumValue, error) {
	propertyTypeName, ok := astmodel.AsTypeName(prop.PropertyType())
	if !ok {
		return astmodel.EnumValue{}, errors.Errorf("property %s was not of type astmodel.TypeName", prop.PropertyName())
	}

	t, ok := definitions[propertyTypeName]
	if !ok {
		return astmodel.EnumValue{}, errors.Errorf("couldn't find type %q", propertyTypeName)
	}

	enumType, ok := astmodel.AsEnumType(t.Type())
	if !ok {
		return astmodel.EnumValue{}, errors.Errorf("%s field with type %s definition was not of type EnumType", prop.PropertyName(), propertyTypeName)
	}

	if len(enumType.Options()) != 1 {
		return astmodel.EnumValue{}, errors.Errorf("enum %s used on property %s has more than one possible value", propertyTypeName, prop.PropertyName())
	}

	return enumType.Options()[0], nil
}
