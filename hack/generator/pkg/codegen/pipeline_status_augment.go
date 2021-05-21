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

			newTypes := make(astmodel.Types)
			newTypes.AddTypes(types)

			// build the augmenter we will use:
			augmenter := fuseAugmenters(
				flattenAugmenter(newTypes), // note that this can alter newTypes as a side-effect, so it must be preloaded with all the types
			)

			for typeName, typeDef := range newTypes {
				resource, ok := typeDef.Type().(*astmodel.ResourceType)
				if !ok {
					continue
				}

				// augment spec with any bits needed from status
				newSpec, err := augmenter(resource.SpecType(), resource.StatusType())
				if err != nil {
					return nil, err
				}

				newTypes[typeName] = typeDef.WithType(resource.WithSpec(newSpec))
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
func flattenAugmenter(allTypes astmodel.Types) augmenter {
	return func(main astmodel.Type, swagger astmodel.Type) (astmodel.Type, error) {
		var merger = astmodel.NewTypeMerger(func(ctx interface{}, left, right astmodel.Type) (astmodel.Type, error) {
			// return left (= main) if the two sides are not ObjectTypes
			return left, nil
		})

		merger.Add(func(main, swagger astmodel.TypeName) (astmodel.Type, error) {
			// when we merge two typenames we always return the main type name
			// however, *as a side effect*, we augment this named type with the information
			// from the corresponding swagger type.

			// this allows us to handle cases where names differ greatly from JSON schema to Swagger,
			// we rely on the structure of the types to tell us which types are the same despite having different names

			// TODO: do we need to resolve these until we hit a non-TypeName?
			newType, err := merger.Merge(allTypes[main].Type(), allTypes[swagger].Type())
			if err != nil {
				return nil, err
			}

			allTypes[main] = allTypes[main].WithType(newType)

			return main, nil
		})

		// need to resolve main type
		merger.Add(func(main astmodel.TypeName, swagger astmodel.Type) (astmodel.Type, error) {
			newMain, err := merger.Merge(allTypes[main].Type(), swagger)
			if err != nil {
				return nil, err
			}

			return newMain, nil
		})

		// need to resolve swagger type
		merger.Add(func(main astmodel.Type, swagger astmodel.TypeName) (astmodel.Type, error) {
			result, err := merger.Merge(main, allTypes[swagger].Type())
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
