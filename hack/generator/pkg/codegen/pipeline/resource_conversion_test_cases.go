/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/testcases"
)

// InjectPropertyAssignmentTestsID is the unique identifier for this stage
const InjectResourceConversionTestsID = "injectResourceConversionTestCases"

func InjectResourceConversionTestCases(idFactory astmodel.IdentifierFactory) Stage {

	stage := MakeStage(
		InjectResourceConversionTestsID,
		"Add test cases to verify Resource Conversion functions",
		func(ctx context.Context, state *State) (*State, error) {
			factory := makeResourceConversionTestCaseFactory(idFactory)
			modifiedTypes := make(astmodel.Types)
			var errs []error
			for _, d := range state.Types() {
				if factory.NeedsTest(d) {
					updated, err := factory.AddTestTo(d)
					if err != nil {
						errs = append(errs, err)
					} else {
						modifiedTypes[updated.Name()] = updated
					}
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			return state.WithTypes(state.Types().OverlayWith(modifiedTypes)), nil
		})

	stage.RequiresPrerequisiteStages(
		InjectPropertyAssignmentFunctionsStageID, // Need PropertyAssignmentFunctions to test
		ImplementConvertibleInterfaceStageId,     // Need the conversions.Convertible interface to be present
		InjectJsonSerializationTestsID,           // We reuse the generators from the JSON tests
	)

	return stage
}

type resourceConversionTestCaseFactory struct {
	injector  *astmodel.TestCaseInjector
	idFactory astmodel.IdentifierFactory
}

func makeResourceConversionTestCaseFactory(idFactory astmodel.IdentifierFactory) resourceConversionTestCaseFactory {
	return resourceConversionTestCaseFactory{
		injector:  astmodel.NewTestCaseInjector(),
		idFactory: idFactory,
	}
}

func (s *resourceConversionTestCaseFactory) NeedsTest(def astmodel.TypeDefinition) bool {
	resourceType, ok := astmodel.AsResourceType(def.Type())
	if !ok {
		return false
	}

	_, found := resourceType.FindInterface(astmodel.ConvertibleInterface)
	return found
}

func (s *resourceConversionTestCaseFactory) AddTestTo(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	_, ok := astmodel.AsResourceType(def.Type())
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("expected %s to be a resourceType", def.Name())
	}

	testCase := testcases.NewResourceConversionTestCase(def, s.idFactory)
	return s.injector.Inject(def, testCase)
}
