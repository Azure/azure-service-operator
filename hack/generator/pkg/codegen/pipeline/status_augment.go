/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

func AugmentSpecWithStatus() Stage {
	return MakeStage(
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
	/* TODO: there is a lot here that should be pulled into a "default augmenter" value that can be reused,
	but at the moment we only have one augmenter */

	return func(main astmodel.Type, swagger astmodel.Type) (astmodel.Type, error) {
		merger := astmodel.NewTypeMerger(func(ctx interface{}, left, right astmodel.Type) (astmodel.Type, error) {
			// as a fallback, return left (= main) if we have nothing else to do
			return left, nil
		})

		// need to resolve main type
		merger.Add(func(main astmodel.TypeName, swagger astmodel.Type) (astmodel.Type, error) {
			// Connascence alert! This is in coöperation with the code in
			// determine_resource_ownership.go:updateChildResourceDefinitionsWithOwner
			// which requires the typenames to identify which ChildResources match
			// which Resources. The ChildResource types should never actually be
			// used, so it's okay that we don't visit them here.
			if strings.HasSuffix(main.Name(), "ChildResource") {
				// don't touch child resources!
				return main, nil
			}

			mainType := allTypes.Get(main).Type()

			newMain, err := merger.Merge(mainType, swagger)
			if err != nil {
				return nil, err
			}

			// return original typename if not changed
			if newMain == mainType {
				return main, nil
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

		merger.Add(func(main *astmodel.OptionalType, swagger astmodel.Type) (astmodel.Type, error) {
			// we ignore optionality when matching things up, since there are
			// discordances between the JSON Schema/Swagger as some teams have handcrafted them
			result, err := merger.Merge(main.Element(), swagger)
			if err != nil {
				return nil, err
			}

			// return original type if not changed
			if result == main.Element() {
				return main, nil
			}

			return astmodel.NewOptionalType(result), nil
		})

		merger.Add(func(main astmodel.Type, swagger *astmodel.OptionalType) (astmodel.Type, error) {
			// we ignore optionality when matching things up, since there are
			// discordances between the JSON Schema/Swagger as some teams have handcrafted them
			return merger.Merge(main, swagger.Element())
		})

		merger.Add(func(main *astmodel.FlaggedType, swagger astmodel.Type) (astmodel.Type, error) {
			result, err := merger.Merge(main.Element(), swagger)
			if err != nil {
				return nil, err
			}

			// return original type if not changed
			if result == main.Element() {
				return main, nil
			}

			return main.WithElement(result), nil
		})

		merger.Add(func(main, swagger *astmodel.MapType) (astmodel.Type, error) {
			keyResult, err := merger.Merge(main.KeyType(), swagger.KeyType())
			if err != nil {
				return nil, err
			}

			valueResult, err := merger.Merge(main.ValueType(), swagger.ValueType())
			if err != nil {
				return nil, err
			}

			// return original type if not changed
			if keyResult == main.KeyType() && valueResult == main.ValueType() {
				return main, nil
			}

			return astmodel.NewMapType(keyResult, valueResult), nil
		})

		merger.Add(func(main, swagger *astmodel.ArrayType) (astmodel.Type, error) {
			result, err := merger.Merge(main.Element(), swagger.Element())
			if err != nil {
				return nil, err
			}

			// return original type if not changed
			if result == main.Element() {
				return main, nil
			}

			return astmodel.NewArrayType(result), nil
		})

		merger.Add(func(main, swagger *astmodel.AllOfType) (astmodel.Type, error) {
			panic("allofs should have been removed")
		})

		merger.Add(func(main, swagger *astmodel.OneOfType) (astmodel.Type, error) {
			panic("oneofs should have been removed")
		})

		// this shouldn't ever happen, I think
		merger.Add(func(main, swagger *astmodel.ResourceType) (astmodel.Type, error) {
			return astmodel.NewErroredType(nil, []string{"unsupported"}, nil), nil
		})

		merger.Add(func(main, swagger *astmodel.ObjectType) (astmodel.Type, error) {
			props := main.Properties()

			changed := false
			for ix, mainProp := range props {
				// find a matching property in the swagger spec
				if swaggerProp, ok := swagger.Property(mainProp.PropertyName()); ok {
					// first copy over flatten property
					if mainProp.Flatten() != swaggerProp.Flatten() {
						changed = true
						mainProp = mainProp.SetFlatten(swaggerProp.Flatten())
					}

					// now recursively merge property types
					newType, err := merger.Merge(mainProp.PropertyType(), swaggerProp.PropertyType())
					if err != nil {
						return nil, err
					}

					if newType != mainProp.PropertyType() {
						changed = true
					}

					props[ix] = mainProp.WithType(newType)
				}
			}

			if !changed {
				return main, nil
			}

			return main.WithProperties(props...), nil
		})

		return merger.Merge(main, swagger)
	}
}
