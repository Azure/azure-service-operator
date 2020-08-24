/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// createResourceLists creates the <resource>List types expected by controller-runtime
func createResourceLists() PipelineStage {

	return MakePipelineStage(
		"createResourceLists",
		"Create the ResourceList types for each resource",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			result := make(astmodel.Types)
			for _, typeDef := range types {
				result.Add(typeDef)
				if _, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					newDef := astmodel.MakeTypeDefinition(
						astmodel.MakeTypeName(typeDef.Name().PackageReference, typeDef.Name().Name()+"List"),
						astmodel.NewResourceListType(typeDef.Name()))
					result.Add(newDef)
				}
			}

			return result, nil
		})
}
