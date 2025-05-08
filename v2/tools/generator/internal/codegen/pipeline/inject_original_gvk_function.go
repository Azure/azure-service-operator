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

// InjectOriginalGVKFunctionStageID is the unique identifier for this pipeline stage
const InjectOriginalGVKFunctionStageID = "injectOriginalGVKFunction"

// InjectOriginalGVKFunction injects the function OriginalGVK() into each Resource type
// This function allows us to recover the original version used to create each custom resource, giving the operator the
// information needed to interact with ARM using the correct API version.
func InjectOriginalGVKFunction(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectOriginalGVKFunctionStageID,
		"Inject the function OriginalGVK() into each Resource type",
		func(ctx context.Context, state *State) (*State, error) {
			definitions := state.Definitions()
			injector := astmodel.NewFunctionInjector()
			result := definitions.Copy()

			for name, def := range definitions.AllResources() {
				var fn *functions.OriginalGVKFunction
				if astmodel.IsStoragePackageReference(name.PackageReference()) {
					fn = functions.NewOriginalGVKFunction(functions.ReadProperty, idFactory)
				} else {
					fn = functions.NewOriginalGVKFunction(functions.ReadFunction, idFactory)
				}

				defWithFn, err := injector.Inject(def, fn)
				if err != nil {
					return nil, eris.Wrapf(err, "injecting OriginalGVK() into %s", name)
				}

				result[defWithFn.Name()] = defWithFn
			}

			return state.WithDefinitions(result), nil
		})

	stage.RequiresPrerequisiteStages(InjectOriginalVersionFunctionStageID)
	return stage
}
