/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"

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
			definitions := make(astmodel.TypeDefinitionSet, len(state.Definitions()))
			for name, def := range state.Definitions() {
				// Apply type transformation, if any
				newDef, err := transformDefinition(def, config, log)
				if err != nil {
					return nil, eris.Wrapf(err, "unable to transform type %q", name)
				}

				definitions.Add(newDef)

				// If we renamed the type, add an alias from the old name to the new;
				// When we remove type aliases later in the pipeline, these will result in any references being updated
				if newDef.Name() != def.Name() {
					alias := astmodel.MakeTypeDefinition(
						def.Name(), newDef.Name())
					definitions.Add(alias)
				}
			}

			// Ensure the type transformers had no errors
			if err := config.GetTransformersError(); err != nil {
				return nil, err
			}

			return state.WithDefinitions(definitions), nil
		})

	stage.RequiresPrerequisiteStages("nameTypes", "allof-anyof-objects")

	return stage
}

// TransformDefinition applies a type transformation to the definition if appropriate.
// If a transformation is applied, the new definition is returned along with a reason for the transformation.
// If something goes wrong, only an error is returned.
// If no transformation is found, returns nil, ""
func transformDefinition(
	def astmodel.TypeDefinition,
	config *config.Configuration,
	log logr.Logger,
) (astmodel.TypeDefinition, error) {
	name := def.Name()
	for _, transformer := range config.Transformers {
		if transformer.AppliesToDefinition(def) {
			result, err := transformer.TransformDefinition(def)
			if err != nil {
				return astmodel.TypeDefinition{}, err
			}

			log.V(2).Info(
				"Transforming type",
				"type", name,
				"because", transformer.Because,
				"transformation", result.Type())

			def = result
		}
	}

	return def, nil
}
