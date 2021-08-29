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

// MarkLatestStorageVariantAsHubVersionID creates a Stage to mark the latest non-preview storage variant of a resource
// as the hub version of that resource for persistence
func MarkLatestStorageVariantAsHubVersion() Stage {
	stage := MakeStage(
		MarkLatestStorageVariantAsHubVersionID,
		"Mark the latest GA storage variant of each resource as the hub version",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs, err := MarkLatestStorageVariantsAsHubs(state)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to mark latest resource version as storage version")
			}

			types := state.Types().OverlayWith(updatedDefs)
			return state.WithTypes(types), nil
		})

	stage.RequiresPrerequisiteStages(CreateConversionGraphStageId)

	return stage
}

// MarkLatestStorageVariantsAsHubs marks the latest version of each resource as the storage version, returning
// a set that contains only the modified types
func MarkLatestStorageVariantsAsHubs(state *State) (astmodel.Types, error) {
	modified := make(astmodel.Types)
	for _, def := range state.Types() {
		// see if it is a resource
		if resourceType, ok := def.Type().(*astmodel.ResourceType); ok {

			// Check whether this is the hub type for this resource
			name := def.Name()
			hubType := state.ConversionGraph().FindHubTypeName(name)
			if name.Equals(hubType) {
				def = def.WithType(resourceType.MarkAsStorageVersion())
				modified.Add(def)
			}
		}
	}

	return modified, nil
}
