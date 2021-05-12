/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// addCrossplaneEmbeddedResourceStatus puts an embedded runtimev1alpha1.ResourceStatus on every spec type
func addCrossplaneEmbeddedResourceStatus(idFactory astmodel.IdentifierFactory) PipelineStage {

	return MakePipelineStage(
		"addCrossplaneEmbeddedResourceStatus",
		"Adds an embedded runtimev1alpha1.ResourceStatus to every status type",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			statusTypeName := astmodel.MakeTypeName(
				CrossplaneRuntimeV1Alpha1Package,
				idFactory.CreateIdentifier("ResourceStatus", astmodel.Exported))
			embeddedStatus := astmodel.NewPropertyDefinition("", ",inline", statusTypeName)

			result := make(astmodel.Types)
			for _, typeDef := range types {
				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {

					if astmodel.IgnoringErrors(resource.StatusType()) == nil {
						continue
					}

					statusDef, err := types.ResolveResourceStatusDefinition(resource)
					if err != nil {
						return nil, errors.Wrapf(err, "getting resource status definition")
					}

					// The assumption here is that specs are all Objects
					updatedDef, err := statusDef.ApplyObjectTransformation(func(o *astmodel.ObjectType) (astmodel.Type, error) {
						return o.WithEmbeddedProperty(embeddedStatus)
					})
					if err != nil {
						return nil, errors.Wrapf(err, "adding embedded crossplane status")
					}

					result.Add(typeDef)
					result.Add(updatedDef)
				}
			}

			for _, typeDef := range types {
				if !result.Contains(typeDef.Name()) {
					result.Add(typeDef)
				}
			}

			return result, nil
		})
}
