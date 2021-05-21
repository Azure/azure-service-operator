/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/pkg/errors"
)

func flattenProperties() PipelineStage {
	return MakePipelineStage("flatten-properties", "Apply flattening to properties marked for flattening", applyPropertyFlattening)
}

func applyPropertyFlattening(
	ctx context.Context,
	defs astmodel.Types) (astmodel.Types, error) {
	visitor := makeFlatteningVisitor(defs)

	result := make(astmodel.Types)
	for name, def := range defs {
		newDef, err := visitor.VisitDefinition(def, name)
		if err != nil {
			return nil, err
		}

		result.Add(newDef)
	}

	return result, nil
}

func makeFlatteningVisitor(defs astmodel.Types) astmodel.TypeVisitor {
	return astmodel.TypeVisitorBuilder{
		VisitObjectType: func(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
			newIt, err := astmodel.IdentityVisitOfObjectType(this, it, ctx)
			if err != nil {
				return nil, err
			}

			it = newIt.(*astmodel.ObjectType)

			newProps, err := flattenObjectType(it, defs)
			if err != nil {
				return nil, err
			}

			// safety check:
			if err := checkForDuplicateNames(newProps); err != nil {
				return nil, err
			}

			result := it.WithoutProperties().WithProperties(newProps...)

			return result, nil
		},
		VisitFlaggedType: func(this *astmodel.TypeVisitor, it *astmodel.FlaggedType, ctx interface{}) (astmodel.Type, error) {
			// skip ARM types, do not flatten
			if it.HasFlag(astmodel.ARMFlag) {
				return it, nil
			}

			return astmodel.IdentityVisitOfFlaggedType(this, it, ctx)
		},
	}.Build()
}

func checkForDuplicateNames(props []*astmodel.PropertyDefinition) error {
	names := make(map[astmodel.PropertyName]struct{})
	for _, p := range props {
		if _, ok := names[p.PropertyName()]; ok {
			return errors.Errorf("Flattening caused duplicate property name %q", p.PropertyName())
		}

		names[p.PropertyName()] = struct{}{}
	}

	return nil
}

// containerName is the type that will hold the flattened properties, i.e. act as a container for them
func flattenObjectType(objectType *astmodel.ObjectType, defs astmodel.Types) ([]*astmodel.PropertyDefinition, error) {
	var flattenedProps []*astmodel.PropertyDefinition

	props := objectType.Properties()
	for _, prop := range props {
		if prop.Flatten() {
			innerProps, err := flattenPropType(prop.PropertyType(), defs)
			if err != nil {
				return nil, err
			}

			for ix := range innerProps {
				innerProps[ix] = innerProps[ix].AddFlattenedFrom(prop.PropertyName())
			}

			flattenedProps = append(flattenedProps, innerProps...)
		} else {
			flattenedProps = append(flattenedProps, prop)
		}
	}

	return flattenedProps, nil
}

func flattenPropType(propType astmodel.Type, defs astmodel.Types) ([]*astmodel.PropertyDefinition, error) {
	switch propType := propType.(type) {
	// "base case"
	case *astmodel.ObjectType:
		return flattenObjectType(propType, defs)

	// typename must be resolved
	case astmodel.TypeName:
		resolved, err := defs.FullyResolve(propType)
		if err != nil {
			return nil, err
		}

		props, err := flattenPropType(resolved, defs)
		if err != nil {
			return nil, err
		}

		return props, nil

	// flattening something that is optional makes everything inside it optional
	case *astmodel.OptionalType:
		innerProps, err := flattenPropType(propType.Element(), defs)
		if err != nil {
			return nil, err
		}

		for ix := range innerProps {
			innerProps[ix] = innerProps[ix].MakeOptional()
		}

		return innerProps, nil

	default:
		return nil, errors.Errorf("Flatten applied to non-object type: %s", propType.String())
	}
}
