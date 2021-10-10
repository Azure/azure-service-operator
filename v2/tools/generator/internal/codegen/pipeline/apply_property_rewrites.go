/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// ApplyPropertyRewrites applies any typeTransformers for properties.
// It is its own pipeline stage so that we can apply it after the allOf/oneOf types have
// been "lowered" to objects.
func ApplyPropertyRewrites(config *config.Configuration) Stage {
	return MakeLegacyStage(
		"propertyRewrites",
		"Modify property types using configured transforms",
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
