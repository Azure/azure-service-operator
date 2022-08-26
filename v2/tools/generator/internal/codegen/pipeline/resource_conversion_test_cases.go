/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/testcases"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// InjectResourceConversionTestsID is the unique identifier for this stage
const InjectResourceConversionTestsID = "injectResourceConversionTestCases"

// InjectResourceConversionTestCases is a pipeline stage to inject test cases
func InjectResourceConversionTestCases(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectResourceConversionTestsID,
		"Add test cases to verify Resource implementations of conversion.Convertible (funcs ConvertTo & ConvertFrom) behave as expected",
		func(ctx context.Context, state *State) (*State, error) {
			factory := makeResourceConversionTestCaseFactory(idFactory)
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
		ImplementConvertibleInterfaceStageId,     // Need the conversions.Convertible interface to be present
		InjectJsonSerializationTestsID)           // We reuse the generators from the JSON tests

	return stage
}

// resourceConversionTestCaseFactory is a factory for injecting test cases where needed
type resourceConversionTestCaseFactory struct {
	injector  *astmodel.TestCaseInjector
	idFactory astmodel.IdentifierFactory
}

// makeResourceConversionTestCaseFactory creates a new factory to inject the test cases
func makeResourceConversionTestCaseFactory(idFactory astmodel.IdentifierFactory) resourceConversionTestCaseFactory {
	return resourceConversionTestCaseFactory{
		injector:  astmodel.NewTestCaseInjector(),
		idFactory: idFactory,
	}
}

// NeedsTest will return true if the passed TypeDefinition is a resource implementing conversion.Convertible
func (_ *resourceConversionTestCaseFactory) NeedsTest(def astmodel.TypeDefinition) bool {
	resourceType, ok := astmodel.AsResourceType(def.Type())
	if !ok {
		return false
	}

	_, found := resourceType.FindInterface(astmodel.ConvertibleInterface)
	return found
}

// AddTestTo modifies the passed TypeDefinition by injecting the required test case
func (factory *resourceConversionTestCaseFactory) AddTestTo(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	resource, ok := astmodel.AsResourceType(def.Type())
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("expected %s to be a resourceType", def.Name())
	}

	testCase, err := testcases.NewResourceConversionTestCase(def.Name(), resource, factory.idFactory)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "adding resource conversion test case to %s", def.Name())
	}

	return factory.injector.Inject(def, testCase)
}
