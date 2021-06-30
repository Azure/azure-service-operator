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

			// Partition the four kinds of types for which we create storage variants
			resourceTypes := storage.FindResourceTypes(types)
			specTypes := storage.FindSpecTypes(types)
			statusTypes := storage.FindStatusTypes(types)
			otherTypes := types.Where(func(def astmodel.TypeDefinition) bool {
				_, isObject := astmodel.AsObjectType(def.Type())
				return isObject &&
					!specTypes.Contains(def.Name()) &&
					!statusTypes.Contains(def.Name())
			})

			storageTypes := make(astmodel.Types)

			typeConverter := storage.NewTypeConverter(types, idFactory)

			// Create storage variants of Resource types
			for name, def := range resourceTypes {
				storageDef, err := typeConverter.ConvertResourceDefinition(def)
				if err != nil {
					return nil, errors.Wrapf(err, "creating storage variant for resource %q", name)
				}

				storageTypes.Add(storageDef)
				conversionGraph.AddLink(name.PackageReference, storageDef.Name().PackageReference)
			}

			// Create storage variants of Spec types
			for name, def := range specTypes {
				storageDef, err := typeConverter.ConvertSpecDefinition(def)
				if err != nil {
					return nil, errors.Wrapf(err, "creating storage variant for spec %q", name)
				}

				storageTypes.Add(storageDef)
			}

			// Create storage variants of Status types
			for name, def := range statusTypes {
				storageDef, err := typeConverter.ConvertStatusDefinition(def)
				if err != nil {
					return nil, errors.Wrapf(err, "creating storage variant for status %q", name)
				}

				storageTypes.Add(storageDef)
			}

			// Create storage variants of other Object types
			for name, def := range otherTypes {
				storageDef, err := typeConverter.ConvertObjectDefinition(def)
				if err != nil {
					return nil, errors.Wrapf(err, "creating storage variant for status %q", name)
				}

				storageTypes.Add(storageDef)
			}

			result := types.Copy()
			result.AddTypes(storageTypes)
			return result, nil
		})

	result.RequiresPrerequisiteStages(injectOriginalVersionFunctionId)
	return result
}
