/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

func augmentSpecWithStatus() PipelineStage {
	return MakePipelineStage(
		"augmentSpecWithStatus",
		"Merges information from Status into Spec",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			// build the augmenter we will use:
			augmenter := fuseAugmenters(
				flattenAugmenter(types),
			)

			newTypes := make(astmodel.Types)

			for _, typeDef := range types {
				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					// augment spec with any bits needed from status
					newSpec, err := augmenter(resource.SpecType(), resource.StatusType())
					if err != nil {
						return nil, err
					}

					newTypes.Add(typeDef.WithType(resource.WithSpec(newSpec)))
				} else {
					newTypes.Add(typeDef)
				}
			}

			return newTypes, nil
		})
}

// an augmenter adds information from the Swagger-derived type to
// the main JSON schema-derived type, and returns the new type
type augmenter func(main astmodel.Type, swagger astmodel.Type) (astmodel.Type, error)

// fuseAugmenters merges multiple augments into one by applying each
// augmenter in order to the result of the previous augmenter
func fuseAugmenters(augments ...augmenter) augmenter {
	return func(main astmodel.Type, swagger astmodel.Type) (astmodel.Type, error) {
		var err error
		for _, augment := range augments {
			main, err = augment(main, swagger)
			if err != nil {
				return nil, err
			}
		}

		return main, nil
	}
}

// flattenAugmenter copies across the "flatten" property from Swagger
func flattenAugmenter(allTypes astmodel.ReadonlyTypes) augmenter {
	return func(main astmodel.Type, swagger astmodel.Type) (astmodel.Type, error) {
		merger := astmodel.NewTypeMerger(func(ctx interface{}, left, right astmodel.Type) (astmodel.Type, error) {
			// as a fallback, return left (= main) if we have nothing else to do
			return left, nil
		})

		merger.Add(func(main, swagger astmodel.TypeName) (astmodel.Type, error) {
			// when we merge two typenames we know that (structurally) they must
			// be the ‘same’ type, even if they have different names

			// this allows us to handle cases where names differ greatly from JSON schema to Swagger,
			// we rely on the structure of the types to tell us which types are the same

			newType, err := merger.Merge(allTypes.Get(main).Type(), allTypes.Get(swagger).Type())
			if err != nil {
				return nil, err
			}

			return newType, nil
		})

		// need to resolve main type
		merger.Add(func(main astmodel.TypeName, swagger astmodel.Type) (astmodel.Type, error) {
			newMain, err := merger.Merge(allTypes.Get(main).Type(), swagger)
			if err != nil {
				return nil, err
			}

			return newMain, nil
		})

		// need to resolve swagger type
		merger.Add(func(main astmodel.Type, swagger astmodel.TypeName) (astmodel.Type, error) {
			result, err := merger.Merge(main, allTypes.Get(swagger).Type())
			if err != nil {
				return nil, err
			}

			return result, nil
		})

		merger.Add(func(main, swagger *astmodel.ObjectType) (astmodel.Type, error) {
			props := main.Properties()
			for ix, mainProp := range props {
				// find a matching property in the swagger spec
				if swaggerProp, ok := swagger.Property(mainProp.PropertyName()); ok {
					// first copy over flatten property
					mainProp = mainProp.SetFlatten(swaggerProp.Flatten())

					// now recursively merge property types
					newType, err := merger.Merge(mainProp.PropertyType(), swaggerProp.PropertyType())
					if err != nil {
						return nil, err
					}

					props[ix] = mainProp.WithType(newType)
				}
			}

			return main.WithProperties(props...), nil
		})

		return merger.Merge(main, swagger)
	}
}
