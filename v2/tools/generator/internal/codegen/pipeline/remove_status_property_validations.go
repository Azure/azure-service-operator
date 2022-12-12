/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
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
//  1. Status has a modelAsString enum with 2 values, but in a future API version, a 3rd value is added.
//     The fact that the enum is modelAsString allows the service to return the new 3rd value even in old API
//     versions.
//  2. Status has an int that must be between 10 and 20. In a future API version, that restriction is relaxed and
//     the same int can now be between 0 and 50.
//  3. A bug in the services Swagger specification causes the service to accept enums with any case, but always
//     return the enum all uppercase
//
// In the above cases, if we left validation on the Status types, we would be unable to persist the content
// returned by the service (apiserver will reject it as not matching the OpenAPI schema). This could be a problem
// in cases where the resource was created via some other means and then imported into
func RemoveStatusValidations() *Stage {
	return NewStage(
		RemoveStatusPropertyValidationsStageID,
		"Remove validation from all status properties",
		func(ctx context.Context, state *State) (*State, error) {
			result, err := removeStatusTypeValidations(state.Definitions())
			if err != nil {
				return nil, err
			}

			/* TODO(donotmerge)
			err = errorIfSpecStatusOverlap(result, state.Definitions())
			if err != nil {
				return nil, err
			}
			*/

			remaining := state.Definitions().Except(result)
			result.AddTypes(remaining)

			return state.WithDefinitions(result), nil
		})
}

func removeStatusTypeValidations(definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	statusDefinitions := astmodel.FindStatusDefinitions(definitions)

	walker := astmodel.NewTypeWalker(
		definitions,
		astmodel.TypeVisitorBuilder{
			VisitEnumType:      removeEnumValidations,
			VisitValidatedType: removeValidatedType,
			VisitObjectType:    removeKubebuilderRequired,
		}.Build())

	var errs []error

	result := make(astmodel.TypeDefinitionSet)
	for _, def := range statusDefinitions {
		updatedTypes, err := walker.Walk(def)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "failed walking definitions"))
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

type overlapError struct {
	name       astmodel.TypeName
	specRefs   []astmodel.TypeName
	statusRefs []astmodel.TypeName
}

// TODO: Remove nolint below
func errorIfSpecStatusOverlap(statusDefinitions astmodel.TypeDefinitionSet, definitions astmodel.TypeDefinitionSet) error { // nolint:deadcode
	allSpecTypes, err := astmodel.FindSpecConnectedDefinitions(definitions)
	if err != nil {
		return errors.Wrap(err, "couldn't find all spec definitions")
	}

	// Verify that the set of spec definitions and the set of modified status definitions is totally disjoint
	intersection := allSpecTypes.Intersect(statusDefinitions)
	if len(intersection) > 0 {
		var problems []overlapError
		for name := range intersection {
			specRefs := []astmodel.TypeName{}
			for tname, t := range definitions {
				if t.References().Contains(name) {
					specRefs = append(specRefs, tname)
				}
			}

			statusRefs := []astmodel.TypeName{}
			for tname, t := range statusDefinitions {
				if t.References().Contains(name) {
					statusRefs = append(statusRefs, tname)
				}
			}

			problems = append(problems, overlapError{
				name:       name,
				specRefs:   specRefs,
				statusRefs: statusRefs,
			})
		}

		result := strings.Builder{}
		result.WriteString(fmt.Sprintf("expected 0 overlapping spec/status definitions but there were %d.\n", len(intersection)))
		for _, problem := range problems {
			result.WriteString(fmt.Sprintf("- %s, referenced by:\n", problem.name))
			for _, referencer := range problem.specRefs {
				result.WriteString(fmt.Sprintf("\t- [spec] %s\n", referencer))
			}
			for _, referencer := range problem.statusRefs {
				result.WriteString(fmt.Sprintf("\t- [status] %s\n", referencer))
			}
		}

		return errors.Errorf(result.String())
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

// removeKubebuilderRequired removes kubebuilder:validation:Required from all properties
func removeKubebuilderRequired(this *astmodel.TypeVisitor, ot *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	ot.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
		ot = ot.WithProperty(prop.MakeOptional())
	})

	return astmodel.IdentityVisitOfObjectType(this, ot, ctx)
}
