/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// AddConfigMapsStageID is the unique identifier for this pipeline stage
const AddConfigMapsStageID = "addConfigMaps"

// AddConfigMaps replaces properties flagged as a config map with genruntime.ConfigMapReference or genruntime.OptionalConfigMapReference
func AddConfigMaps(config *config.Configuration) *Stage {
	stage := NewStage(
		AddConfigMapsStageID,
		"Replace properties flagged as a configMap with genruntime.ConfigMapReference. For properties "+
			"flagged as an optional configMap, add a new <property>FromConfig property.",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()

			updatedDefs, err := transformConfigMaps(config, defs)
			if err != nil {
				return nil, errors.Wrap(err, "transforming spec secrets")
			}

			return state.WithDefinitions(updatedDefs), nil
		})

	stage.RequiresPostrequisiteStages(CreateARMTypesStageID)

	return stage
}

func transformPropertyToConfigMapReference(prop *astmodel.PropertyDefinition, newType astmodel.Type) (*astmodel.PropertyDefinition, error) {
	// The expectation is that this is a string
	propType := prop.PropertyType()
	if !astmodel.Unwrap(propType).Equals(astmodel.StringType, astmodel.EqualityOverrides{}) {
		return nil, errors.Errorf("expected property %q to be a string, but was: %T", prop.PropertyName(), propType)
	}

	// check if it's optional
	_, optional := astmodel.AsOptionalType(propType)

	if optional {
		newType = astmodel.NewOptionalType(newType)
	}

	return prop.WithType(newType), nil
}

func createNewConfigMapReference(prop *astmodel.PropertyDefinition, newType astmodel.Type) (*astmodel.PropertyDefinition, *astmodel.PropertyDefinition, error) {
	// The expectation is that this is a string
	propType := prop.PropertyType()
	if !astmodel.TypeEquals(astmodel.Unwrap(propType), astmodel.StringType) {
		return nil, nil, errors.Errorf("expected property %q to be a string, but was: %s", prop.PropertyName(), astmodel.DebugDescription(propType))
	}

	jsonName, ok := prop.JSONName()
	if !ok {
		return nil, nil, errors.Errorf("property %s didn't have a JSON name", prop.PropertyName())
	}

	// Neither property can be required anymore. That's a bit unfortunate from a fail-fast perspective.
	updatedProp := prop.
		WithTag(astmodel.OptionalConfigMapPairTag, string(prop.PropertyName())).
		MakeOptional().
		MakeTypeOptional()

	newProp := prop.
		WithName(prop.PropertyName()+astmodel.OptionalConfigMapReferenceSuffix).
		WithType(newType).
		WithJsonName(jsonName+astmodel.OptionalConfigMapReferenceSuffix).
		WithTag(astmodel.OptionalConfigMapPairTag, string(prop.PropertyName())).
		MakeOptional().
		MakeTypeOptional()

	return updatedProp, newProp, nil
}

func transformConfigMaps(cfg *config.Configuration, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)

	applyConfigMaps := func(_ *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
		typeName := ctx.(astmodel.TypeName)
		for _, prop := range it.Properties().Copy() {
			mode, _ := cfg.ImportConfigMapMode(typeName, prop.PropertyName())
			switch mode {
			case config.ImportConfigMapModeRequired:
				newProp, err := transformPropertyToConfigMapReference(prop, astmodel.ConfigMapReferenceType)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to transform property to configmap on type %s", typeName)
				}
				it = it.WithProperty(newProp)
			case config.ImportConfigMapModeOptional:
				updatedProp, newProp, err := createNewConfigMapReference(prop, astmodel.ConfigMapReferenceType)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to transform property to optional configmap on type %s", typeName)
				}

				// If the property we're about to add already exists, that's bad!
				if _, ok := it.Property(newProp.PropertyName()); ok {
					return nil, errors.Errorf(
						"failed to transform property to optional configmap on type %s. Property %s already exists",
						typeName,
						newProp.PropertyName())
				}

				it = it.WithProperties(updatedProp, newProp)
			default:
				continue
			}
		}

		return it, nil
	}

	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: applyConfigMaps,
	}.Build()

	for _, def := range definitions {
		updatedDef, err := visitor.VisitDefinition(def, def.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "visiting type %q", def.Name())
		}

		result.Add(updatedDef)
	}

	// Verify that all 'importConfigMapMode' modifiers are consumed before returning the result
	err := cfg.VerifyImportConfigMapModeConsumed()
	if err != nil {
		return nil, err
	}

	return result, nil
}
