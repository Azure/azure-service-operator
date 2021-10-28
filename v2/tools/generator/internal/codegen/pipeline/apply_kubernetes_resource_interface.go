/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/interfaces"
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
					newResource, err := interfaces.AddKubernetesResourceInterfaceImpls(typeName, resource, idFactory, types)
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
