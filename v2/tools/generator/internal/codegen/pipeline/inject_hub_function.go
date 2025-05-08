/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// InjectHubFunctionStageID is the unique identifier for this pipeline stage
const InjectHubFunctionStageID = "injectHubFunction"

// InjectHubFunction modifies the nominates storage version (aka hub version) of each resource by injecting a Hub()
// function so that it satisfies the required interface.
func InjectHubFunction(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectHubFunctionStageID,
		"Inject the function Hub() into each hub resource",
		func(ctx context.Context, state *State) (*State, error) {
			definitions := state.Definitions()
			injector := astmodel.NewFunctionInjector()
			result := definitions.Copy()

			for name, def := range definitions.AllResources() {
				rt, ok := astmodel.AsResourceType(def.Type())
				if !ok {
					return nil, eris.Errorf("expected %s to be a resource type (should never happen)", name)
				}

				if rt.IsStorageVersion() {
					fn := functions.NewHubFunction(idFactory)
					defWithFn, err := injector.Inject(def, fn)
					if err != nil {
						return nil, eris.Wrapf(err, "injecting Hub() into %s", name)
					}

					result[name] = defWithFn
				}
			}

			return state.WithDefinitions(result), nil
		})

	stage.RequiresPrerequisiteStages(MarkLatestStorageVariantAsHubVersionID)
	return stage
}
