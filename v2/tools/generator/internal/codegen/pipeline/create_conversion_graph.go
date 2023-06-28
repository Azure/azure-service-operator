/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// CreateConversionGraphStageId is the unique identifier for this stage
const CreateConversionGraphStageId = "createConversionGraph"

// CreateConversionGraph walks the set of available types and creates a graph of conversions that will be used to
// convert resources to/from the designated storage (or hub) version
func CreateConversionGraph(
	configuration *config.Configuration,
	generatorPrefix string) *Stage {
	stage := NewStage(
		CreateConversionGraphStageId,
		"Create the graph of conversions between versions of each resource group",
		func(ctx context.Context, state *State) (*State, error) {
			// Collect all distinct references
			allNames := astmodel.NewTypeNameSet()
			for _, def := range state.Definitions() {
				if def.Name().IsARMType() {
					// ARM types don't participate in the conversion graph
					continue
				}

				allNames.Add(def.Name())
			}

			builder := storage.NewConversionGraphBuilder(
				configuration.ObjectModelConfiguration, generatorPrefix)
			builder.AddAll(allNames)
			graph, err := builder.Build()
			if err != nil {
				// Shouldn't have any non-local references, if we do, abort
				return nil, errors.Wrapf(err, "creating conversion graph")
			}

			return state.WithConversionGraph(graph), nil
		})

	stage.RequiresPrerequisiteStages(CreateStorageTypesStageID)
	return stage
}
