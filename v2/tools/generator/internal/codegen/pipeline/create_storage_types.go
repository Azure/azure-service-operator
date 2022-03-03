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
)

const CreateStorageTypesStageID = "createStorageTypes"

// CreateStorageTypes returns a pipeline stage that creates dedicated storage types for each resource and nested object.
// Storage versions are created for *all* API versions to allow users of older versions of the operator to easily
// upgrade. This is of course a bit odd for the first release, but defining the approach from day one is useful.
func CreateStorageTypes() *Stage {
	stage := NewStage(
		CreateStorageTypesStageID,
		"Create storage versions of CRD types",
		func(ctx context.Context, state *State) (*State, error) {
			// Predicate to isolate both resources and complex objects
			isResourceOrObject := func(def astmodel.TypeDefinition) bool {
				_, isResource := astmodel.AsResourceType(def.Type())
				_, isObject := astmodel.AsObjectType(def.Type())
				return isResource || isObject
			}

			// Predicate to filter out ARM types
			isNotARMType := func(def astmodel.TypeDefinition) bool {
				return !astmodel.ARMFlag.IsOn(def.Type())
			}

			// Filter to the types we want to process
			typesToConvert := state.Definitions().Where(isResourceOrObject).Where(isNotARMType)

			storageDefs := make(astmodel.TypeDefinitionSet)
			typeConverter := storage.NewTypeConverter(state.Definitions(), state.ConversionGraph())

			// Create storage variants
			for name, def := range typesToConvert {
				storageDef, err := typeConverter.ConvertDefinition(def)
				if err != nil {
					return nil, errors.Wrapf(err, "creating storage variant of %q", name)
				}

				storageDefs.Add(storageDef)
			}

			defs := state.Definitions().Copy()
			defs.AddTypes(storageDefs)

			return state.WithDefinitions(defs), nil
		})

	stage.RequiresPrerequisiteStages(InjectOriginalVersionFunctionStageID, CreateConversionGraphStageId)

	return stage
}
