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

var CrossplaneRuntimeV1Alpha1Package = astmodel.MakeExternalPackageReference("github.com/crossplane/crossplane-runtime/apis/core/v1alpha1")

// addCrossplaneEmbeddedResourceSpec puts an embedded runtimev1alpha1.ResourceSpec on every spec type
func addCrossplaneEmbeddedResourceSpec(idFactory astmodel.IdentifierFactory) PipelineStage {

	return MakePipelineStage(
		"addCrossplaneEmbeddedResourceSpec",
		"Adds an embedded runtimev1alpha1.ResourceSpec to every spec type",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			specTypeName := astmodel.MakeTypeName(
				CrossplaneRuntimeV1Alpha1Package,
				idFactory.CreateIdentifier("ResourceSpec", astmodel.Exported))
			embeddedSpec := astmodel.NewPropertyDefinition("", ",inline", specTypeName)

			result := make(astmodel.Types)
			for _, typeDef := range types {

				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {

					specDef, err := types.ResolveResourceSpecDefinition(resource)
					if err != nil {
						return nil, errors.Wrapf(err, "getting resource spec definition")
					}

					// The assumption here is that specs are all Objects
					updatedDef, err := specDef.ApplyObjectTransformation(func(o *astmodel.ObjectType) (astmodel.Type, error) {
						return o.WithEmbeddedProperty(embeddedSpec)
					})
					if err != nil {
						return nil, errors.Wrapf(err, "adding embedded crossplane spec")
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
