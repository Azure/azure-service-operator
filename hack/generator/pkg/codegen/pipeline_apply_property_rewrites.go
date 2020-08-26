/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"k8s.io/klog/v2"
)

// applyPropertyRewrites applies any typeTransformers for properties.
// It is its own pipeline stage so that we can apply it after the allOf/oneOf types have
// been "lowered" to objects.
func applyPropertyRewrites(config *config.Configuration) PipelineStage {

	return MakePipelineStage(
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

				transformation := config.TransformTypeProperties(name, objectType)
				if transformation != nil {
					klog.V(2).Infof("Transforming %s", transformation)
					objectType = transformation.NewType
				}

				newTypes.Add(t.WithType(objectType))
			}

			return newTypes, nil
		})
}
