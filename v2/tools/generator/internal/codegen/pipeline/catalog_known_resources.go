/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// CatalogKnownResourcesStageID is the unique identifier for this pipeline stage
const CatalogKnownResourcesStageID = "catalogKnownResources"

func CatalogKnownResources() *Stage {
	stage := NewStage(
		CatalogKnownResourcesStageID,
		"Catalog known resources",
		func(ctx context.Context, state *State) (*State, error) {
			// Index by ARM type so resources can be matched after export renaming.
			catalog := make(map[string]astmodel.TypeNameSet)
			for _, rsrc := range state.Definitions().AllResources() {
				armType := astmodel.MustBeResourceType(rsrc.Type()).ARMType()
				resources, ok := catalog[armType]
				if !ok {
					resources = make(astmodel.TypeNameSet)
					catalog[armType] = resources
				}

				resources.Add(rsrc.Name())
			}

			return StateWithData(state, AllKnownResources, catalog), nil
		},
	)

	// We're cataloging all known resources, so we have to do this before we reduce the set of
	// resources we're processing by applying the export filters.
	stage.RequiresPostrequisiteStages(ApplyExportFiltersStageID)

	return stage
}
