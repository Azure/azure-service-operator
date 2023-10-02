/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// ApplyTypeRewrites applies any typeTransformers.
func ApplyTypeRewrites(
	config *config.Configuration,
	log logr.Logger,
) *Stage {
	stage := NewStage(
		"typeRewrites",
		"Modify types using configured type transforms",
		func(ctx context.Context, state *State) (*State, error) {
			definitions := state.Definitions()

			modifiedDefinitions := make(astmodel.TypeDefinitionSet, len(definitions))
			for name, def := range definitions {

				// First apply entire type transformation, if any
				if newType, because := config.TransformType(name); newType != nil {
					log.V(2).Info(
						"Transforming type",
						"type", name,
						"because", because,
						"transformation", newType)
					modifiedDefinitions.Add(def.WithType(newType))
					continue
				}

				// Apply property transformations, if any
				if objectType, ok := def.Type().(*astmodel.ObjectType); ok {
					transformations := config.TransformTypeProperties(name, objectType)
					for _, transformation := range transformations {
						transformation.LogTo(log)
						objectType = transformation.NewType
					}

					modifiedDefinitions.Add(def.WithType(objectType))
					continue
				}
			}

			// Ensure the type transformers had no errors
			if err := config.GetTypeTransformersError(); err != nil {
				return nil, err
			}

			// Ensure that the property transformers had no errors
			if err := config.GetPropertyTransformersError(); err != nil {
				return nil, err
			}

			return state.WithDefinitions(definitions.OverlayWith(modifiedDefinitions)), nil
		})

	stage.RequiresPrerequisiteStages("nameTypes", "allof-anyof-objects")

	return stage
}
