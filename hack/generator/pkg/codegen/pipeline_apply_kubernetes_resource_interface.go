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

			result := make(astmodel.Types)
			for _, typeDef := range types {

				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {

					specName, ok := resource.SpecType().(astmodel.TypeName)
					if !ok {
						return nil, errors.Errorf("resource %q spec was not of type TypeName, instead: %T", typeDef.Name(), typeDef.Type())
					}

					spec, ok := types[specName]
					if !ok {
						return nil, errors.Errorf("couldn't find resource spec %q", specName)
					}

					specObj, err := astmodel.TypeAsObjectType(spec.Type())
					if err != nil {
						return nil, errors.Wrapf(err, "resource spec %q did not contain an object", specName)
					}

					iface, err := astmodel.NewKubernetesResourceInterfaceImpl(idFactory, specObj)
					if err != nil {
						return nil, errors.Wrapf(err, "Couldn't implement Kubernetes resource interface for %q", spec.Name())
					}

					newDef := typeDef.WithType(resource.WithInterface(iface))
					result.Add(newDef)
				} else {
					result.Add(typeDef)
				}
			}

			return result, nil
		})
}
