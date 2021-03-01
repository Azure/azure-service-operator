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

			storageTypes := make(astmodel.Types)
			visitor := storage.MakeStorageTypesVisitor(types)
			vc := storage.MakeStorageTypesVisitorContext()
			var errs []error
			for _, d := range types {
				d := d

				if astmodel.ARMFlag.IsOn(d.Type()) {
					// Skip ARM definitions, we don't need to create storage variants of those
					continue
				}

				if _, ok := types.ResolveEnumDefinition(&d); ok {
					// Skip Enum definitions as we use the base type for storage
					continue
				}

				def, err := visitor.VisitDefinition(d, vc)
				if err != nil {
					errs = append(errs, err)
					continue
				}

				finalDef := def.WithDescription(storage.DescriptionForStorageVariant(d))
				storageTypes[finalDef.Name()] = finalDef
			}

			if len(errs) > 0 {
				err := kerrors.NewAggregate(errs)
				return nil, err
			}

			types.AddTypes(storageTypes)

			return types, nil
		})
}
