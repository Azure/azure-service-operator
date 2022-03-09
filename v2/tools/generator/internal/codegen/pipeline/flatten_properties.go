/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func FlattenProperties() *Stage {
	return NewLegacyStage("flattenProperties", "Apply flattening to properties marked for flattening", applyPropertyFlattening)
}

func applyPropertyFlattening(
	ctx context.Context,
	defs astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	visitor := makeFlatteningVisitor(defs)

	result := make(astmodel.TypeDefinitionSet)
	for name, def := range defs {
		newDef, err := visitor.VisitDefinition(def, name)
		if err != nil {
			return nil, err
		}

		result.Add(newDef)
	}

	return result, nil
}

func makeFlatteningVisitor(defs astmodel.TypeDefinitionSet) astmodel.TypeVisitor {
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

			// fix any colliding names:
			newProps = fixCollisions(newProps)

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

// fixCollisions renames properties to avoid collisions created while flattening.
//
// For example:
//
// If we have a structure like:
// - `Type`
// - `Properties` (marked to flatten)
//   - `Type`
//
// We will flatten it and rename the second property, resulting in:
// - `Type`
// - `PropertiesType`
//
func fixCollisions(props []*astmodel.PropertyDefinition) []*astmodel.PropertyDefinition {
	names := make(map[astmodel.PropertyName]int)

	// collect all names
	for _, p := range props {
		n := p.PropertyName()
		names[n] = names[n] + 1
	}

	result := make([]*astmodel.PropertyDefinition, len(props))
	for ix, p := range props {

		// check if there will be a collision
		if names[p.PropertyName()] > 1 {
			// rename the flattened one, not the top-level one
			if p.WasFlattened() {
				names := p.FlattenedFrom()
				// FFS
				stringNames := make([]string, len(names))
				for i := range names {
					if i < len(names)-1 { // the last entry in FlattenedFrom is the source property which might have a different name
						stringNames[i] = string(names[i])
					}
				}

				// disambiguate by prefixing with properties
				newName := astmodel.PropertyName(strings.Join(stringNames, "") + string(p.PropertyName()))
				newJsonName := strings.ToLower(strings.Join(stringNames, "_") + string(p.PropertyName()))
				p = p.WithName(newName).WithJsonName(newJsonName)
			}
		}

		result[ix] = p
	}

	return result
}

// collectAndFlattenProperties walks the object type and extracts all properties, flattening any properties that require flattening
func collectAndFlattenProperties(objectType *astmodel.ObjectType, defs astmodel.TypeDefinitionSet) ([]*astmodel.PropertyDefinition, error) {
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
func flattenPropType(propType astmodel.Type, defs astmodel.TypeDefinitionSet) ([]*astmodel.PropertyDefinition, error) {
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
			innerProps[ix] = innerProps[ix].MakeTypeOptional()
		}

		return innerProps, nil

	default:
		return nil, errors.Errorf("flatten applied to non-object type: %s", propType.String())
	}
}
