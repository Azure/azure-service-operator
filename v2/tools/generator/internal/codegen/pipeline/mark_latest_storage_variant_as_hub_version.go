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

// MarkLatestStorageVariantAsHubVersionID is the unique identifier for this pipeline stage
const MarkLatestStorageVariantAsHubVersionID = "markLatestStorageVariantAsHubVersion"

// MarkLatestStorageVariantAsHubVersion creates a Stage to mark the latest non-preview storage variant of a resource
// as the hub version of that resource for persistence
func MarkLatestStorageVariantAsHubVersion() *Stage {
	stage := NewStage(
		MarkLatestStorageVariantAsHubVersionID,
		"Mark the latest GA storage variant of each resource as the hub version",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs, err := astmodel.FindResourceDefinitions(state.Definitions()).Process(
				func(def astmodel.TypeDefinition) (*astmodel.TypeDefinition, error) {
					rsrc := astmodel.MustBeResourceType(def.Type())
					hub, err := state.ConversionGraph().FindHub(def.Name(), state.Definitions())
					if err != nil {
						return nil, errors.Wrapf(err, "finding hub type for %s", def.Name())
					}

					if astmodel.TypeEquals(def.Name(), hub) {
						// We have the hub type, modify it and return for Process() to accumulate into updatedDefs
						def = def.WithType(rsrc.MarkAsStorageVersion())
						return &def, nil
					}

					// Nothing to modify, nothing to return
					return nil, nil
				})
			if err != nil {
				return nil, errors.Wrap(err, "marking storage versions")
			}

			defs := state.Definitions().OverlayWith(updatedDefs)
			return state.WithDefinitions(defs), nil
		})

	return stage.WithRequiredPrerequisites(CreateConversionGraphStageId)
}
