/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const SimplifyDefinitionsStageId = "simplifyDefinitions"

// SimplifyDefinitions creates a pipeline stage that removes any wrapper types prior to actual code generation
func SimplifyDefinitions() *Stage {
	return NewLegacyStage(
		SimplifyDefinitionsStageId,
		"Flatten definitions by removing wrapper types",
		func(ctx context.Context, defs astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			visitor := createSimplifyingVisitor()
			var errs []error
			result := make(astmodel.TypeDefinitionSet)
			for _, def := range defs {
				visited, err := visitor.VisitDefinition(def, nil)
				if err != nil {
					errs = append(errs, err)
				} else {
					result.Add(visited)
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			return result, nil
		})
}

func createSimplifyingVisitor() astmodel.TypeVisitor[any] {
	removeFlags := func(tv *astmodel.TypeVisitor[any], ft *astmodel.FlaggedType, ctx any) (astmodel.Type, error) {
		element := ft.Element()
		return tv.Visit(element, ctx)
	}

	skipComplexTypes := func(ot *astmodel.ObjectType) astmodel.Type {
		return ot
	}

	result := astmodel.TypeVisitorBuilder[any]{
		VisitFlaggedType: removeFlags,
		VisitObjectType:  skipComplexTypes,
	}.Build()

	return result
}
