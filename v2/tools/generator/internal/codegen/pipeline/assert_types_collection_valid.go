/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// AssertTypesCollectionValid creates a Stage that ensures that each reachable type in the types collection
// has TypeName's that are all reachable as well. This check fails if there is any TypeName that refers to a type that doesn't
// exist.
func AssertTypesCollectionValid() *Stage {
	return NewLegacyStage(
		"assertTypesStructureValid",
		"Verify that all local TypeNames refer to a type",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			visitor := astmodel.TypeVisitorBuilder{}.Build()
			typeWalker := astmodel.NewTypeWalker(definitions, visitor)

			for _, def := range definitions {
				if astmodel.IsResourceDefinition(def) {
					_, err := typeWalker.Walk(def)
					if err != nil {
						return nil, err
					}
				}
			}

			return definitions, nil
		})
}
