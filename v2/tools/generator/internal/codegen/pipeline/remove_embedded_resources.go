/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/embeddedresources"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// RemoveEmbeddedResourcesStageID is the unique identifier for this pipeline stage
const RemoveEmbeddedResourcesStageID = "removeEmbeddedResources"

func RemoveEmbeddedResources(
	configuration *config.Configuration,
	log logr.Logger,
) *Stage {
	return NewStage(
		RemoveEmbeddedResourcesStageID,
		// Only removes structural aspects of embedded resources, Id/ARMId references are retained.
		"Remove properties that point to embedded resources.",
		func(ctx context.Context, state *State) (*State, error) {
			definitions := state.Definitions()
			remover, err := embeddedresources.MakeEmbeddedResourceRemover(configuration, definitions)
			if err != nil {
				return nil, err
			}

			result, err := remover.RemoveEmbeddedResources(log)
			if err != nil {
				return nil, err
			}

			return state.WithDefinitions(result), nil
		})
}
