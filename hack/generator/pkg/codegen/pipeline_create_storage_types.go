/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"
)

// createStorageTypes returns a pipeline stage that creates dedicated storage types for each resource and nested object.
// Storage versions are created for *all* API versions to allow users of older versions of the operator to easily
// upgrade. This is of course a bit odd for the first release, but defining the approach from day one is useful.
func createStorageTypes(idFactory astmodel.IdentifierFactory) PipelineStage {
	return MakePipelineStage(
		"createStorage",
		"Create storage versions of CRD types",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			// Create a factory for each group (aka service) and divvy up the types
			factories := make(map[string]*storage.StorageTypeFactory)
			for name, def := range types {

				ref, ok := name.PackageReference.AsLocalPackage()
				if !ok {
					// Skip definitions from non-local packages
					// (should never happen)
					klog.Warningf("Skipping storage type generation for unexpected non-local package reference %q", name.PackageReference)
					continue
				}

				factory, ok := factories[ref.Group()]
				if !ok {
					klog.V(3).Infof("Creating storage factory for %s", ref.Group())
					factory = storage.NewStorageTypeFactory(ref.Group(), idFactory)
					factories[ref.Group()] = factory
				}

				if astmodel.ARMFlag.IsOn(def.Type()) {
					// skip ARM types as they don't need storage variants
					continue
				}

				factory.Add(def)
			}

			// Collect up all the results
			result := make(astmodel.Types)
			var errs []error
			for _, factory := range factories {
				t, err := factory.Types()
				if err != nil {
					errs = append(errs, err)
					continue
				}

				result.AddTypes(t)
			}

			err := kerrors.NewAggregate(errs)
			if err != nil {
				return nil, err
			}

			unmodified := types.Except(result)
			result.AddTypes(unmodified)
			return result, nil
		})
}
