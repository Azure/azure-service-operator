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

// CreateConversionGraphStageId is the unique identifier for this stage
const CreateConversionGraphStageId = "linkConversionGraph"

// CreateConversionGraph walks the set of available types and creates a graph of conversions that will be used to
// convert resources to/from the designated storage (or hub) version
func CreateConversionGraph(conversionGraph *storage.ConversionGraph) Stage {
	stage := MakeStage(
		CreateConversionGraphStageId,
		"Create the graph of conversions between versions of each resource group",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			// Collect all distinct references
			allReferences := astmodel.NewPackageReferenceSet()
			for _, def := range types {
				allReferences.AddReference(def.Name().PackageReference)
			}

			// For each one, create a storage reference to match
			sortedApiReferences := allReferences.AsSlice()
			sortedStorageReferences := make([]astmodel.PackageReference, len(sortedApiReferences))
			astmodel.SortPackageReferencesByPathAndVersion(sortedApiReferences)
			for index, ref := range sortedApiReferences {
				localRef, ok := ref.AsLocalPackage()
				if !ok {
					// Shouldn't have any non-local references, if we do, abort
					return nil, errors.Errorf("expected all API references to be local references, but %s was not", ref)
				}

				storageRef := astmodel.MakeStoragePackageReference(localRef)
				sortedStorageReferences[index] = storageRef
				conversionGraph.AddLink(localRef, storageRef)
			}

			// For each Preview API version, the link goes from the associated storage version to the immediately prior
			// storage version, no matter whether it's preview or GA.
			// We use the API reference to work out whether it's preview or not, because the suffix used for storage
			// versions means IsPreview() always returns true
			for i, ref := range sortedApiReferences {
				if i == 0 || !ref.IsPreview() {
					continue
				}

				conversionGraph.AddLink(sortedStorageReferences[i], sortedStorageReferences[i-1])
			}

			// For each GA (non-Preview) API version, the link goes from the associated storage version to the next GA
			// release
			var gaRelease astmodel.PackageReference
			for index, ref := range sortedApiReferences {
				if index == 0 {
					// Always treat the very earliest version as GA
					gaRelease = sortedStorageReferences[index]
					continue
				}

				if !ref.IsPreview() {
					// Found a GA release, link to the prior GA release
					conversionGraph.AddLink(gaRelease, sortedStorageReferences[index])
					gaRelease = sortedStorageReferences[index]
				}
			}

			return types, nil
		})

	return stage
}
