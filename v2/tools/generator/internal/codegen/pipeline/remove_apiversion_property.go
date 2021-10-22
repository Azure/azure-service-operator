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

// RemoveAPIVersionPropertyStageID is the unique identifier for this pipeline stage
const RemoveAPIVersionPropertyStageID = "removeAPIVersionProperty"

func RemoveAPIVersionProperty() Stage {
	return MakeStage(
		RemoveAPIVersionPropertyStageID,
		"Remove the ARM API version property and instead augment the ResourceType with it",
		func(ctx context.Context, state *State) (*State, error) {

			newDefs := make(astmodel.Types)

			defs := state.Types()
			resources := astmodel.FindResourceTypes(defs)

			for _, resource := range resources {
				resolved, err := defs.ResolveResourceSpecAndStatus(resource)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to resolve resource spec and status types")
				}

				apiVersionProp, ok := resolved.SpecType.Property(astmodel.APIVersionProperty)
				if !ok {
					return nil, errors.Errorf("resource %s is missing type property", resolved.ResourceDef.Name())
				}

				apiVersionEnumTypeName, ok := astmodel.AsTypeName(apiVersionProp.PropertyType())
				if !ok {
					return nil, errors.Errorf("resource %s APIVersion property is not of type TypeName", resolved.ResourceDef.Name())
				}

				apiVersion, err := extractPropertySingleEnumValue(defs, apiVersionProp)
				if err != nil {
					return nil, errors.Wrapf(err, "error extracting %s type property", resolved.SpecDef.Name())
				}

				resourceType := resolved.ResourceType.WithAPIVersion(apiVersionEnumTypeName, apiVersion)
				specType := resolved.SpecType.WithoutProperty(astmodel.APIVersionProperty)

				newDefs.Add(resolved.ResourceDef.WithType(resourceType))
				newDefs.Add(resolved.SpecDef.WithType(specType))
			}

			return state.WithTypes(defs.OverlayWith(newDefs)), nil
		})
}
