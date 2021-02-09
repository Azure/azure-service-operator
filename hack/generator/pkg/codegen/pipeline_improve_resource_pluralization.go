/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// improveResourcePluralization improves pluralization for resources
func improveResourcePluralization() PipelineStage {

	return MakePipelineStage(
		"pluralizeNames",
		"Improve resource pluralization",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			result := make(astmodel.Types)

			for _, typeDef := range types {
				if resourceType, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					newTypeName := typeDef.Name().Singular()
					// check if there is already a resource with this name
					if _, ok := types[newTypeName]; !ok {
						// not found: rename the resource
						typeDef = typeDef.WithName(newTypeName)
					}

					// Need to update owner ref too if applicable
					if resourceType.Owner() != nil {
						owner := resourceType.Owner().Singular()
						resourceType = resourceType.WithOwner(&owner)
						typeDef = typeDef.WithType(resourceType)
					}

					result.Add(typeDef)
				} else {
					result.Add(typeDef)
				}
			}

			return result, nil
		})
}
