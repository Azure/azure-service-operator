/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// MarkLatestStorageVariantAsHubVersionID is the unique identifier for this pipeline stage
const MarkLatestStorageVariantAsHubVersionID = "markLatestStorageVariantAsHubVersion"

// MarkLatestStorageVariantAsHubVersion creates a Stage to mark the latest non-preview storage variant of a resource
// as the hub version of that resource for persistence
func MarkLatestStorageVariantAsHubVersion() Stage {
	stage := MakeStage(
		MarkLatestStorageVariantAsHubVersionID,
		"Mark the latest GA storage variant of each resource as the hub version",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs, err := astmodel.FindResourceTypes(state.Types()).Process(
				func(def astmodel.TypeDefinition) (*astmodel.TypeDefinition, error) {
					rsrc, _ := astmodel.AsResourceType(def.Type())
					hub := state.ConversionGraph().FindHub(def.Name(), state.Types())
					if astmodel.TypeEquals(def.Name(), hub) {
						// We have the hub type
						def = def.WithType(rsrc.MarkAsStorageVersion())
						return &def, nil
					}

					return nil, nil
				})

			if err != nil {
				return nil, errors.Wrap(err, "marking storage versions")
			}

			types := state.Types().OverlayWith(updatedDefs)
			return state.WithTypes(types), nil
		})

	stage.RequiresPrerequisiteStages(CreateConversionGraphStageId)

	return stage
}
