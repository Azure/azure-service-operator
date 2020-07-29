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

	return PipelineStage{
		Name: "Improve resource pluralization",
		Action: func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			result := make(astmodel.Types)
			for _, typeDef := range types {
				if _, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					newTypeName := typeDef.Name().Singular()
					// check if there is already a resource with this name
					if _, ok := types[newTypeName]; !ok {
						// not found: rename the resource
						result.Add(typeDef.WithName(newTypeName))
					} else {
						// resource with singular name already exists,
						// so output the resource without depluralizing
						result.Add(typeDef)
					}
				} else {
					result.Add(typeDef)
				}
			}

			return result, nil
		},
	}
}
