/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
)

const createStorageTypesStageId = "createStorageTypes"

// CreateStorageTypes returns a pipeline stage that creates dedicated storage types for each resource and nested object.
// Storage versions are created for *all* API versions to allow users of older versions of the operator to easily
// upgrade. This is of course a bit odd for the first release, but defining the approach from day one is useful.
func CreateStorageTypes(conversionGraph *storage.ConversionGraph, idFactory astmodel.IdentifierFactory) Stage {
	result := MakeStage(
		createStorageTypesStageId,
		"Create storage versions of CRD types",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			// Isolate both resources and complex objects
			typesToConvert := types.Where(func(def astmodel.TypeDefinition) bool {
				_, ok := astmodel.AsPropertyContainer(def.Type())
				return ok
			})

			storageTypes := make(astmodel.Types)
			typeConverter := storage.NewTypeConverter(types, idFactory)

			// Create storage variants
			for name, def := range typesToConvert {
				storageDef, err := typeConverter.ConvertDefinition(def)
				if err != nil {
					return nil, errors.Wrapf(err, "creating storage variant of %q", name)
				}

				storageTypes.Add(storageDef)
				conversionGraph.AddLink(name.PackageReference, storageDef.Name().PackageReference)
			}

			result := types.Copy()
			result.AddTypes(storageTypes)
			return result, nil
		})

	result.RequiresPrerequisiteStages(injectOriginalVersionFunctionStageId)
	return result
}
