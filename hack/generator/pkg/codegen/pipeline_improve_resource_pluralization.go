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
		Action: func(ctx context.Context, types Types) (Types, error) {

			result := make(Types)
			for typeName, typeDef := range types {
				if _, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					newTypeName := typeName.Singular()
					typeDef = typeDef.WithName(newTypeName)
					result[*newTypeName] = typeDef
				} else {
					result[typeName] = typeDef
				}
			}

			return result, nil
		},
	}
}
