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

	return func(spec astmodel.Type, status astmodel.Type) (astmodel.Type, error) {
		// reminder: merger is invoked with a pair of Types and invokes the first
		// Added function that matches those types
		//
		// in this incarnation we are using it to handle the pair of values
		// (SpecType, StatusType),  to copy the StatusType’s “flatten” property
		// across, if present (this is only provided on the Swagger = Status side)
		//
		// most of the cases are only recursively invoking the merger/matcher on
		// “inner” types; the (Object, Object) case is where the work is done
		//
		// note that we do a small optimization where we don’t return a new type
		// if nothing was changed; this allows us to keep the same TypeNames
		// if no alterations were made

		merger := astmodel.NewTypeMerger(func(ctx interface{}, left, right astmodel.Type) (astmodel.Type, error) {
			// as a fallback, return left (= main type) if we have nothing else to do
			return left, nil
		})

		// this is the main part of the merging, we want to match up (Object, Object)
		// based on their properties and copy across “flatten” when present, and
		// also invoke the merger on the (Spec, Status) types of each property recursively
		merger.Add(func(spec, status *astmodel.ObjectType) (astmodel.Type, error) {
			props := spec.Properties()

			changed := false
			for ix, specProp := range props {
				// find a matching property in the swagger spec
				if swaggerProp, ok := status.Property(specProp.PropertyName()); ok {
					// first copy over flatten property
					if specProp.Flatten() != swaggerProp.Flatten() {
						changed = true
						specProp = specProp.SetFlatten(swaggerProp.Flatten())
					}

					// now recursively merge property types
					newType, err := merger.Merge(specProp.PropertyType(), swaggerProp.PropertyType())
					if err != nil {
						return nil, err
					}

					if !newType.Equals(specProp.PropertyType()) {
						changed = true
					}

					props[ix] = specProp.WithType(newType)
				}
			}

			if !changed {
				return spec, nil
			}

			return spec.WithProperties(props...), nil
		})

		// handle (TypeName, Type) by resolving LHS
		merger.Add(func(spec astmodel.TypeName, status astmodel.Type) (astmodel.Type, error) {
			// Connascence alert! This is in coöperation with the code in
			// determine_resource_ownership.go:updateChildResourceDefinitionsWithOwner
			// which requires the typenames to identify which ChildResources match
			// which Resources. The ChildResource types should never actually be
			// used, so it's okay that we don't visit them here.
			if strings.HasSuffix(spec.Name(), ChildResourceNameSuffix) {
				// don't touch child resources!
				return spec, nil
			}

			specType := allTypes.Get(spec).Type()

			newSpec, err := merger.Merge(specType, status)
			if err != nil {
				return nil, err
			}

			// return original typename if not changed
			if newSpec.Equals(specType) {
				return spec, nil
			}

			return newSpec, nil
		})

		// handle (Type, TypeName) by resolving RHS
		merger.Add(func(spec astmodel.Type, status astmodel.TypeName) (astmodel.Type, error) {
			return merger.Merge(spec, allTypes.Get(status).Type())
		})

		// handle (Optional, Type)
		merger.Add(func(spec *astmodel.OptionalType, status astmodel.Type) (astmodel.Type, error) {
			// we ignore optionality when matching things up, since there are
			// discordances between the JSON Schema/Swagger as some teams have handcrafted them
			newSpecElement, err := merger.Merge(spec.Element(), status)
			if err != nil {
				return nil, err
			}

			// return original type if not changed
			if newSpecElement.Equals(spec.Element()) {
				return spec, nil
			}

			return astmodel.NewOptionalType(newSpecElement), nil
		})

		// handle (Type, Optional)
		merger.Add(func(spec astmodel.Type, status *astmodel.OptionalType) (astmodel.Type, error) {
			// we ignore optionality when matching things up, since there are
			// discordances between the JSON Schema/Swagger as some teams have handcrafted them
			return merger.Merge(spec, status.Element())
		})

		// handle (FlaggedType, Type); there is no need to handle the opposite direction
		// as the swagger types won’t be flagged
		merger.Add(func(spec *astmodel.FlaggedType, status astmodel.Type) (astmodel.Type, error) {
			newSpec, err := merger.Merge(spec.Element(), status)
			if err != nil {
				return nil, err
			}

			// return original type if not changed
			if newSpec.Equals(spec.Element()) {
				return spec, nil
			}

			return spec.WithElement(newSpec), nil
		})

		// handle (Map, Map) by matching up the keys and values
		merger.Add(func(spec, status *astmodel.MapType) (astmodel.Type, error) {
			keyResult, err := merger.Merge(spec.KeyType(), status.KeyType())
			if err != nil {
				return nil, err
			}

			valueResult, err := merger.Merge(spec.ValueType(), status.ValueType())
			if err != nil {
				return nil, err
			}

			// return original type if not changed
			if keyResult.Equals(spec.KeyType()) && valueResult.Equals(spec.ValueType()) {
				return spec, nil
			}

			return astmodel.NewMapType(keyResult, valueResult), nil
		})

		// handle (Array, Array) by matching up the inner types
		merger.Add(func(spec, status *astmodel.ArrayType) (astmodel.Type, error) {
			newSpecElement, err := merger.Merge(spec.Element(), status.Element())
			if err != nil {
				return nil, err
			}

			// return original type if not changed
			if newSpecElement.Equals(spec.Element()) {
				return spec, nil
			}

			return astmodel.NewArrayType(newSpecElement), nil
		})

		// safety check, (AllOf, AllOf) should not occur
		merger.Add(func(_, _ *astmodel.AllOfType) (astmodel.Type, error) {
			panic("allofs should have been removed")
		})

		// safety check, (OneOf, OneOf) should not occur
		merger.Add(func(_, _ *astmodel.OneOfType) (astmodel.Type, error) {
			panic("oneofs should have been removed")
		})

		// this shouldn't ever happen, I think, (Resource, Resource)
		merger.Add(func(_, _ *astmodel.ResourceType) (astmodel.Type, error) {
			return astmodel.NewErroredType(nil, []string{"unsupported"}, nil), nil
		})

		// now go!
		return merger.Merge(spec, status)
	}
}
