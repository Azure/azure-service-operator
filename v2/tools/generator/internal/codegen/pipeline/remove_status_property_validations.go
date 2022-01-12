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

const RemoveStatusPropertyValidationsStageID = "removeStatusPropertyValidation"

// RemoveStatusValidations removes property validations from all Status types.
// This is required because Status is retrieved directly from the ARM API, and there are
// cases where ARM might return something that isn't actually "valid" according to the validation,
// but makes sense in context. Some examples:
// 1. Status has a modelAsString enum with 2 values, but in a future API version, a 3rd value is added.
//    The fact that the enum is modelAsString allows the service to return the new 3rd value even in old API
//    versions.
// 2. Status has an int that must be between 10 and 20. In a future API version, that restriction is relaxed and
//    the same int can now be between 0 and 50.
// 3. A bug in the services Swagger specification causes the service to accept enums with any case, but always
//    return the enum all uppercase
// In the above cases, if we left validation on the Status types, we would be unable to persist the content
// returned by the service (apiserver will reject it as not matching the OpenAPI schema). This could be a problem
// in cases where the resource was created via some other means and then imported into
func RemoveStatusValidations() Stage {
	return MakeStage(
		RemoveStatusPropertyValidationsStageID,
		"Remove validation from all status properties",
		func(ctx context.Context, state *State) (*State, error) {
			result, err := removeStatusTypeValidations(state.Types())
			if err != nil {
				return nil, err
			}

			err = errorIfSpecStatusOverlap(result, state.Types())
			if err != nil {
				return nil, err
			}

			remaining := state.Types().Except(result)
			result.AddTypes(remaining)

			return state.WithTypes(result), nil
		})
}

func removeStatusTypeValidations(types astmodel.Types) (astmodel.Types, error) {
	statusTypes := astmodel.FindStatusTypes(types)

	walker := astmodel.NewTypeWalker(
		types,
		astmodel.TypeVisitorBuilder{
			VisitEnumType:      removeEnumValidations,
			VisitValidatedType: removeValidatedType,
		}.Build())

	var errs []error

	result := make(astmodel.Types)
	for _, def := range statusTypes {
		updatedTypes, err := walker.Walk(def)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "failed walking types"))
		}

		err = result.AddTypesAllowDuplicates(updatedTypes)
		if err != nil {
			errs = append(errs, err)
		}
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	return result, err
}

func errorIfSpecStatusOverlap(statusTypes astmodel.Types, types astmodel.Types) error {
	allSpecTypes, err := astmodel.FindSpecConnectedTypes(types)
	if err != nil {
		return errors.Wrap(err, "couldn't find all spec types")
	}

	// Verify that the set of spec types and the set of modified status types is totally disjoint
	intersection := allSpecTypes.Intersect(statusTypes)
	if len(intersection) > 0 {
		var nameStrings []string
		for name := range intersection {
			nameStrings = append(nameStrings, name.String())
		}

		return errors.Errorf("expected 0 overlapping spec/status types but there were %d. Overlapping: %s", len(intersection), strings.Join(nameStrings, ", "))
	}

	return nil
}

// removeValidatedType returns the validated types element. This assumes that there aren't deeply nested validations.
func removeValidatedType(this *astmodel.TypeVisitor, vt *astmodel.ValidatedType, _ interface{}) (astmodel.Type, error) {
	return vt.ElementType(), nil
}

func removeEnumValidations(this *astmodel.TypeVisitor, et *astmodel.EnumType, _ interface{}) (astmodel.Type, error) {
	return et.WithoutValidation(), nil
}
