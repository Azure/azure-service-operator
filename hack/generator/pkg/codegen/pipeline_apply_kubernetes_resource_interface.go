/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
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
						return nil, errors.Errorf("Resource %q spec was not of type TypeName, instead: %T", typeDef.Name(), typeDef.Type())
					}

					spec, ok := types[specName]
					if !ok {
						return nil, errors.Errorf("Couldn't find resource spec %q", specName)
					}

					specObject, ok := spec.Type().(*astmodel.ObjectType)
					if !ok {
						return nil, errors.Errorf("Spec %q was not of type ObjectType, instead %T", specName, spec.Type())
					}

					iface, err := astmodel.NewKubernetesResourceInterfaceImpl(idFactory, specObject)
					if err != nil {
						// TODO: This should be changed to an error once we handle oneOf/allOf better
						// return nil, errors.Wrapf(err, "Couldn't implement Kubernetes resource interface for %q", spec.Name())
						klog.Warningf("Couldn't implement Kubernetes resource interface for %q, due to: %v", spec.Name(), err)
						result.Add(typeDef)
						continue
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
