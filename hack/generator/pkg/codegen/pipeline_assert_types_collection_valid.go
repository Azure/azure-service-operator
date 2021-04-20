/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// assertTypesCollectionValid creates a PipelineStage that ensures that each reachable type in the types collection
// has TypeName's that are all reachable as well. This check fails if there is any TypeName that refers to a type that doesn't
// exist.
func assertTypesCollectionValid() PipelineStage {
	return MakePipelineStage(
		"assertTypesStructureValid",
		"Asserts that the types collection is valid",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			visitor := astmodel.MakeTypeVisitor()
			typeWalker := astmodel.NewTypeWalker(types, visitor)

			for _, def := range types {
				if astmodel.IsResourceDefinition(def) {
					_, err := typeWalker.Walk(def)
					if err != nil {
						return nil, err
					}
				}
			}

			return types, nil
		})
}
