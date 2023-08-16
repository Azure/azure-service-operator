/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const FlattenPropertiesStageId = "flattenProperties"

func FlattenProperties(log logr.Logger) *Stage {
	return NewStage(
		FlattenPropertiesStageId,
		"Apply flattening to properties marked for flattening",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			visitor := makeFlatteningVisitor(defs, log)

			result := make(astmodel.TypeDefinitionSet)
			for name, def := range defs {
				newDef, err := visitor.VisitDefinition(def, name)
				if err != nil {
					return nil, err
				}

				result.Add(newDef)
			}

			return state.WithDefinitions(result), nil
		})
}

func makeFlatteningVisitor(defs astmodel.TypeDefinitionSet, log logr.Logger) astmodel.TypeVisitor[astmodel.TypeName] {
	return astmodel.TypeVisitorBuilder[astmodel.TypeName]{
		VisitObjectType: func(
			this *astmodel.TypeVisitor[astmodel.TypeName],
			it *astmodel.ObjectType,
			ctx astmodel.TypeName,
		) (astmodel.Type, error) {
			newIt, err := astmodel.IdentityVisitOfObjectType(this, it, ctx)
			if err != nil {
				return nil, err
			}

			it = newIt.(*astmodel.ObjectType)

			newProps, err := collectAndFlattenProperties(ctx, it, defs, log)
			if err != nil {
				return nil, err
			}

			// fix any colliding names:
			newProps = fixCollisions(newProps)

			result := it.WithoutProperties().WithProperties(newProps...)

			return result, nil
		},
	}.Build()
}

func removeFlattenFromObject(tObj *astmodel.ObjectType) *astmodel.ObjectType {
	objProps := tObj.Properties()
	props := make([]*astmodel.PropertyDefinition, 0, objProps.Len())
	objProps.ForEach(func(prop *astmodel.PropertyDefinition) {
		prop = prop.WithType(removeFlatten(prop.PropertyType()))
		prop = prop.SetFlatten(false)
		props = append(props, prop)
	})

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
func collectAndFlattenProperties(
	container astmodel.TypeName,
	objectType *astmodel.ObjectType,
	defs astmodel.TypeDefinitionSet,
	log logr.Logger,
) ([]*astmodel.PropertyDefinition, error) {
	var flattenedProps []*astmodel.PropertyDefinition

	objectType.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
		if !prop.Flatten() {
			// Property doesn't need to be flattened, move along
			flattenedProps = append(flattenedProps, prop)
			return // continue
		}

		innerProps, err := flattenProperty(container, prop, defs, logr.Discard())
		if err != nil {
			log.V(2).Info(
				"Skipping flatten",
				"property", prop.PropertyName(),
				"container", container,
				"reason", err)

			innerProps = []*astmodel.PropertyDefinition{
				prop.SetFlatten(false),
			}
		}

		flattenedProps = append(flattenedProps, innerProps...)
	})

	return flattenedProps, nil
}

func flattenProperty(
	container astmodel.TypeName,
	prop *astmodel.PropertyDefinition,
	defs astmodel.TypeDefinitionSet,
	log logr.Logger,
) ([]*astmodel.PropertyDefinition, error) {
	props, err := flattenPropType(container, prop.PropertyType(), defs, log)
	if err != nil {
		return nil, errors.Wrapf(err, "flattening property %s", prop.PropertyName())
	}

	for i, p := range props {
		props[i] = p.AddFlattenedFrom(prop.PropertyName())
	}

	return props, nil
}

// flattenPropType is invoked on a property marked for flattening to collect all inner properties
func flattenPropType(
	container astmodel.TypeName,
	propType astmodel.Type,
	defs astmodel.TypeDefinitionSet,
	log logr.Logger,
) ([]*astmodel.PropertyDefinition, error) {
	switch propType := propType.(type) {
	// "base case"
	case *astmodel.ObjectType:
		return collectAndFlattenProperties(container, propType, defs, log)

	// typename must be resolved
	case astmodel.TypeName:
		resolved, err := defs.FullyResolve(propType)
		if err != nil {
			return nil, err
		}

		props, err := flattenPropType(container, resolved, defs, log)
		if err != nil {
			return nil, err
		}

		return props, nil

	// flattening something that is optional makes everything inside it optional
	case *astmodel.OptionalType:
		innerProps, err := flattenPropType(container, propType.Element(), defs, log)
		if err != nil {
			return nil, errors.Wrap(err, "wrapping optional type")
		}

		for ix := range innerProps {
			innerProps[ix] = innerProps[ix].MakeTypeOptional()
		}

		return innerProps, nil

	default:
		desc := astmodel.DebugDescription(propType)
		return nil, errors.Errorf("flatten applied to non-object type: %s", desc)
	}
}
