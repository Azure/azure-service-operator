/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const VerifyNoErroredTypesStageID = "verifyNoErroredTypes"

// VerifyNoErroredTypes creates a Stage that verifies that no types contain an ErroredType with errors
func VerifyNoErroredTypes() *Stage {
	stage := NewStage(
		VerifyNoErroredTypesStageID,
		"Verify there are no ErroredType's containing errors",
		func(ctx context.Context, state *State) (*State, error) {
			visitor := newErrorCollectingVisitor()

			for _, def := range state.definitions {

				_, err := visitor.visitor.Visit(def.Type(), erroredTypeVisitorContext{name: def.Name()})
				if err != nil {
					return nil, errors.Wrapf(err, "failed while visiting %q", def.Name())
				}
			}

			err := kerrors.NewAggregate(visitor.errs)
			if err != nil {
				return nil, err
			}

			// This stage doesn't change the generated types at all - if the verification
			// has passed, just return the same defs we started with
			return state, nil
		})

	return stage
}

type errorCollectingVisitor struct {
	errs    []error
	visitor astmodel.TypeVisitor[erroredTypeVisitorContext]
}

func newErrorCollectingVisitor() *errorCollectingVisitor {
	includePropertyContext := astmodel.MakeIdentityVisitOfObjectType(
		func(_ *astmodel.ObjectType, prop *astmodel.PropertyDefinition, ctx erroredTypeVisitorContext) (erroredTypeVisitorContext, error) {
			return ctx.WithProperty(prop.PropertyName()), nil
		})

	result := &errorCollectingVisitor{}
	result.visitor = astmodel.TypeVisitorBuilder[erroredTypeVisitorContext]{
		VisitErroredType:  result.catalogErrors,
		VisitObjectType:   includePropertyContext,
		VisitResourceType: includeResourcePropertyContext,
	}.Build()

	return result
}

func (v *errorCollectingVisitor) catalogErrors(
	this *astmodel.TypeVisitor[erroredTypeVisitorContext],
	it *astmodel.ErroredType,
	ctx erroredTypeVisitorContext,
) (astmodel.Type, error) {
	if len(it.Errors()) > 0 {
		errStrings := strings.Join(it.Errors(), ", ")
		v.errs = append(v.errs, errors.Errorf("%q has property %q with errors: %q", ctx.name, ctx.property, errStrings))
	}

	return astmodel.IdentityVisitOfErroredType(this, it, ctx)
}

func includeResourcePropertyContext(
	this *astmodel.TypeVisitor[erroredTypeVisitorContext],
	it *astmodel.ResourceType,
	ctx erroredTypeVisitorContext,
) (astmodel.Type, error) {
	_, err := this.Visit(it.SpecType(), ctx.WithProperty("Spec"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit resource spec type %q", it.SpecType())
	}

	_, err = this.Visit(it.StatusType(), ctx.WithProperty("Status"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit resource status type %q", it.StatusType())
	}

	return it, nil
}

type erroredTypeVisitorContext struct {
	name     astmodel.TypeName
	property *astmodel.PropertyName
}

func (e erroredTypeVisitorContext) WithProperty(prop astmodel.PropertyName) erroredTypeVisitorContext {
	e.property = &prop
	return e
}
