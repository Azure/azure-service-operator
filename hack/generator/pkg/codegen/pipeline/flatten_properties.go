/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

func FlattenProperties() Stage {
	return MakeStage("flattenProperties", "Apply flattening to properties marked for flattening", applyPropertyFlattening)
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
			name := ctx.(astmodel.TypeName)

			newIt, err := astmodel.IdentityVisitOfObjectType(this, it, ctx)
			if err != nil {
				return nil, err
			}

			it = newIt.(*astmodel.ObjectType)

			newProps, err := collectAndFlattenProperties(it, defs)
			if err != nil {
				return nil, err
			}

			// safety check:
			if err := checkForDuplicateNames(newProps); err != nil {
				klog.Warningf("Skipping flattening for %s: %s", name, err)
				return removeFlattenFromObject(it), nil
			}

			if len(newProps) != len(it.Properties()) {
				klog.V(4).Infof("Flattened properties in %s", name)
			}

			result := it.WithoutProperties().WithProperties(newProps...)

			return result, nil
		},
	}.Build()
}

func removeFlattenFromObject(tObj *astmodel.ObjectType) *astmodel.ObjectType {
	var props []*astmodel.PropertyDefinition
	for _, prop := range tObj.Properties() {
		prop = prop.WithType(removeFlatten(prop.PropertyType()))
		prop = prop.SetFlatten(false)
		props = append(props, prop)
	}

	return tObj.WithProperties(props...)
}

func removeFlatten(t astmodel.Type) astmodel.Type {
	if tObj, ok := t.(*astmodel.ObjectType); ok {
		return removeFlattenFromObject(tObj)
	}

	return t
}

func checkForDuplicateNames(props []*astmodel.PropertyDefinition) error {
	names := make(map[astmodel.PropertyName]struct{})
	for _, p := range props {
		if _, ok := names[p.PropertyName()]; ok {
			return errors.Errorf("flattening caused duplicate property name %q", p.PropertyName())
		}

		names[p.PropertyName()] = struct{}{}
	}

	return nil
}

// collectAndFlattenProperties walks the object type and extracts all properties, flattening any properties that require flattening
func collectAndFlattenProperties(objectType *astmodel.ObjectType, defs astmodel.Types) ([]*astmodel.PropertyDefinition, error) {
	var flattenedProps []*astmodel.PropertyDefinition

	props := objectType.Properties()
	for _, prop := range props {
		if prop.Flatten() {
			innerProps, err := flattenPropType(prop.PropertyType(), defs)
			if err != nil {
				return nil, err
			}

			for _, innerProp := range innerProps {
				flattenedProps = append(flattenedProps, innerProp.AddFlattenedFrom(prop.PropertyName()))
			}
		} else {
			flattenedProps = append(flattenedProps, prop)
		}
	}

	return flattenedProps, nil
}

// flattenPropType is invoked on a property marked for flattening to collect all inner properties
func flattenPropType(propType astmodel.Type, defs astmodel.Types) ([]*astmodel.PropertyDefinition, error) {
	switch propType := propType.(type) {
	// "base case"
	case *astmodel.ObjectType:
		return collectAndFlattenProperties(propType, defs)

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
		return nil, errors.Errorf("flatten applied to non-object type: %s", propType.String())
	}
}
