/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"github.com/Azure/k8s-infra/hack/generator/pkg/codegen/storage"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// createStorageTypes returns a pipeline stage that creates dedicated storage types for each resource and nested object.
// Storage versions are created for *all* API versions to allow users of older versions of the operator to easily
// upgrade. This is of course a bit odd for the first release, but defining the approach from day one is useful.
func createStorageTypes() PipelineStage {
	return MakePipelineStage(
		"createStorage",
		"Create storage versions of CRD types",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			storageFactory := storage.NewStorageTypeFactory(types)

				ref, ok := name.PackageReference.AsLocalPackage()
				if !ok {
					// Skip definitions from non-local packages
					// (should never happen)
					continue
				}

				def, err := visitor.VisitDefinition(d, vc)
				if err != nil {
					errs = append(errs, err)
					continue
				}

				finalDef := def.WithDescription(storage.DescriptionForStorageVariant(d))
				storageFactory.Add(finalDef)
			}

			if len(errs) > 0 {
				err := kerrors.NewAggregate(errs)
				return nil, err
			}

			unmodified := types.Except(result)
			result.AddTypes(unmodified)
			return result, nil
		})
}
