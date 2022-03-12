/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/testcases"
)

// InjectPropertyAssignmentTestsID is the unique identifier for this stage
const InjectPropertyAssignmentTestsID = "injectPropertyAssignmentTestCases"

func InjectPropertyAssignmentTests(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectPropertyAssignmentTestsID,
		"Add test cases to verify PropertyAssignment functions",
		func(ctx context.Context, state *State) (*State, error) {
			factory := makePropertyAssignmentTestCaseFactory(idFactory)
			modifiedDefs := make(astmodel.TypeDefinitionSet)
			var errs []error
			for _, d := range state.Definitions() {
				if factory.NeedsTest(d) {
					updated, err := factory.AddTestTo(d)
					if err != nil {
						errs = append(errs, err)
					} else {
						modifiedDefs[updated.Name()] = updated
					}
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			return state.WithDefinitions(state.Definitions().OverlayWith(modifiedDefs)), nil
		})

	stage.RequiresPrerequisiteStages(
		InjectPropertyAssignmentFunctionsStageID, // Need PropertyAssignmentFunctions to test
		InjectJsonSerializationTestsID)           // We reuse the generators from the JSON tests

	return stage
}

type propertyAssignmentTestCaseFactory struct {
	injector  *astmodel.TestCaseInjector
	idFactory astmodel.IdentifierFactory
}

func makePropertyAssignmentTestCaseFactory(idFactory astmodel.IdentifierFactory) propertyAssignmentTestCaseFactory {
	return propertyAssignmentTestCaseFactory{
		injector:  astmodel.NewTestCaseInjector(),
		idFactory: idFactory,
	}
}

func (s *propertyAssignmentTestCaseFactory) NeedsTest(def astmodel.TypeDefinition) bool {
	container, ok := astmodel.AsFunctionContainer(def.Type())
	if !ok {
		return false
	}

	for _, fn := range container.Functions() {
		if _, ok := fn.(*functions.PropertyAssignmentFunction); ok {
			return true
		}
	}

	return false
}

func (s *propertyAssignmentTestCaseFactory) AddTestTo(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	container, ok := astmodel.AsFunctionContainer(def.Type())
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("expected %s to be a function container", def.Name())
	}

	testCase := testcases.NewPropertyAssignmentTestCase(def.Name(), container, s.idFactory)
	return s.injector.Inject(def, testCase)
}
