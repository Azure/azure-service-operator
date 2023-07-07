/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

func ApplyIsResourceOverrides(configuration *config.Configuration) *Stage {
	stage := NewStage(
		"applyIsResourceOverrides",
		"Apply $isResource overrides to objects",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs := make(astmodel.TypeDefinitionSet)
			var errs []error
			for name, def := range state.Definitions() {
				objectType, ok := def.Type().(*astmodel.ObjectType)
				if !ok {
					continue
				}

				isResource, err := configuration.ObjectModelConfiguration.IsResource.Lookup(name)
				if err != nil {
					if config.IsNotConfiguredError(err) {
						// $isResource is not configured, skip this object
						continue
					}

					// If something else went wrong, keep details
					errs = append(errs, err)
					continue
				}

				objectType = objectType.WithIsResource(isResource)
				def = def.WithType(objectType)
				updatedDefs.Add(def)
			}

			var err error
			err = kerrors.NewAggregate(errs)
			if err != nil {
				return nil, err
			}

			// Ensure that all the $isResource properties were used
			err = configuration.ObjectModelConfiguration.IsResource.VerifyConsumed()
			if err != nil {
				return nil, err
			}

			state = state.WithDefinitions(state.Definitions().OverlayWith(updatedDefs))
			return state, nil
		})

	return stage
}
