/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/pipeline"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/config"
)

// applyPropertyRewrites applies any typeTransformers for properties.
// It is its own pipeline stage so that we can apply it after the allOf/oneOf types have
// been "lowered" to objects.
func applyPropertyRewrites(config *config.Configuration) pipeline.Stage {

	return pipeline.MakeStage(
		"propertyRewrites",
		"Applying type transformers to properties",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			newTypes := make(astmodel.Types, len(types))
			for name, t := range types {

				objectType, ok := t.Type().(*astmodel.ObjectType)
				if !ok {
					newTypes.Add(t)
					continue
				}

				transformations := config.TransformTypeProperties(name, objectType)
				for _, transformation := range transformations {
					klog.V(2).Infof("Transforming %s", transformation)
					objectType = transformation.NewType
				}

				newTypes.Add(t.WithType(objectType))
			}

			// Ensure that the property transformers had no errors
			err := config.GetPropertyTransformersError()
			if err != nil {
				return nil, err
			}

			return newTypes, nil
		})
}
