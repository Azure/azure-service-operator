/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
	"github.com/Azure/azure-service-operator/v2/internal/generator/interfaces"
)

// ApplyKubernetesResourceInterface ensures that every Resource implements the KubernetesResource interface
func ApplyKubernetesResourceInterface(idFactory astmodel.IdentifierFactory) Stage {
	return MakeLegacyStage(
		"applyKubernetesResourceInterface",
		"Add the KubernetesResource interface to every resource",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			skip := make(map[astmodel.TypeName]struct{})

			result := make(astmodel.Types)
			for typeName, typeDef := range types {
				if _, ok := skip[typeName]; ok {
					continue
				}

				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					typeValue, err := extractType(types, resource, typeName)
					if err != nil {
						return nil, err
					}

					newResource, err := interfaces.AddKubernetesResourceInterfaceImpls(typeName, resource, idFactory, types, typeValue)
					if err != nil {
						return nil, errors.Wrapf(err, "couldn't implement Kubernetes resource interface for %q", typeName)
					}

					// this is really very ugly; a better way?
					if _, ok := newResource.SpecType().(astmodel.TypeName); !ok {
						// the resource Spec was replaced with a new definition; update it
						// by replacing the named definition:
						specName := resource.SpecType().(astmodel.TypeName)
						result[specName] = astmodel.MakeTypeDefinition(specName, newResource.SpecType())
						skip[specName] = struct{}{}
						newResource = newResource.WithSpec(specName)
					}

					newDef := typeDef.WithType(newResource)
					result.Add(newDef)
				} else {
					result.Add(typeDef)
				}
			}

			return result, nil
		})
}

func extractType(types astmodel.Types, resource *astmodel.ResourceType, resourceName astmodel.TypeName) (string, error) {
	// Find the ARM type associated with this resource and get it's "Type" property, which we need
	specTypeName, ok := astmodel.AsTypeName(resource.SpecType())
	if !ok {
		return "", errors.Errorf("Resource %q doesn't have spec with Type TypeName", resourceName)
	}

	armTypeDefinition, err := GetARMTypeDefinition(types, specTypeName)
	if err != nil {
		return "", err
	}

	// Look up the magical "type" property - we'll need this to implement GetType()
	ot, ok := astmodel.AsObjectType(armTypeDefinition.Type())
	if !ok {
		return "", errors.Errorf("expected %q to be astmodel.ObjectType but was %T", armTypeDefinition.Name(), armTypeDefinition.Type())
	}

	typeProperty, ok := ot.Property(astmodel.TypeProperty)
	if !ok {
		return "", errors.Errorf("expected %q to have a Type property", armTypeDefinition.Name())
	}

	typePropertyTypeName, ok := astmodel.AsTypeName(typeProperty.PropertyType())
	if !ok {
		return "", errors.Errorf("expected %s.Type to be of type astmodel.TypeName", armTypeDefinition.Name())
	}

	t, ok := types[typePropertyTypeName]
	if !ok {
		return "", errors.Errorf("couldn't find type %q", typePropertyTypeName)
	}

	enumType, ok := astmodel.AsEnumType(t.Type())
	if !ok {
		return "", errors.Errorf("Type field with type %s definition was not of type EnumType", typePropertyTypeName)
	}

	return enumType.Options()[0].Value, nil
}
