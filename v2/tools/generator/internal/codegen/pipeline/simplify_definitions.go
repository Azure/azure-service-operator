/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

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
					if !astmodel.TypeEquals(def.Type(), visited.Type()) {
						klog.V(3).Infof("Simplified %s from %s to %s", def.Name(), def.Type(), visited.Type())
					}
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			return result, nil
		})
}

func createSimplifyingVisitor() astmodel.TypeVisitor {
	removeFlags := func(tv *astmodel.TypeVisitor, ft *astmodel.FlaggedType, ctx interface{}) (astmodel.Type, error) {
		element := ft.Element()
		return tv.Visit(element, ctx)
	}

	skipComplexTypes := func(ot *astmodel.ObjectType) astmodel.Type {
		return ot
	}

	result := astmodel.TypeVisitorBuilder{
		VisitFlaggedType: removeFlags,
		VisitObjectType:  skipComplexTypes,
	}.Build()

	return result
}
