/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog"
)

// simplifyDefinitions creates a pipeline stage that removes any wrapper types prior to actual code generation
func simplifyDefinitions() PipelineStage {
	return MakePipelineStage(
		"simplifyDefinitions",
		"Flatten definitions by removing wrapper types",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			visitor := createSimplifyingVisitor()
			var errs []error
			result := make(astmodel.Types)
			for _, def := range defs {
				d, err := visitor.VisitDefinition(def, nil)
				if err != nil {
					errs = append(errs, err)
				} else {
					result.Add(*d)
					if !def.Type().Equals(d.Type()) {
						klog.V(3).Infof("Simplified %v from %v to %v", def.Name(), def.Type(), d.Type())
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
	result := astmodel.MakeTypeVisitor()

	// Unwrap ArmTypes, promoting the object within
	result.VisitArmType = func(tv *astmodel.TypeVisitor, at *astmodel.ArmType, ctx interface{}) (astmodel.Type, error) {
		ot := at.ObjectType()
		return tv.Visit(&ot, ctx)
	}

	// Don't need to waste time iterating within complex objects
	result.VisitObjectType = func(_ *astmodel.TypeVisitor, ot *astmodel.ObjectType, _ interface{}) (astmodel.Type, error) {
		return ot, nil
	}

	return result
}
