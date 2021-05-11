/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/pkg/errors"
)

// applyKubernetesResourceInterface ensures that every Resource implements the KubernetesResource interface
func applyKubernetesResourceInterface(idFactory astmodel.IdentifierFactory) PipelineStage {

	return MakePipelineStage(
		"applyKubernetesResourceInterface",
		"Ensures that every resource implements the KubernetesResource interface",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			skip := make(map[astmodel.TypeName]struct{})

			result := make(astmodel.Types)
			for typeName, typeDef := range types {
				if _, ok := skip[typeName]; ok {
					continue
				}

				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					newResource, err := astmodel.AddKubernetesResourceInterfaceImpls(typeName, resource, idFactory, types)
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
