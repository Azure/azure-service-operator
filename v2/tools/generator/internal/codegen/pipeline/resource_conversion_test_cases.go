/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/rotisserie/eris"
	"golang.org/x/net/context"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/testcases"
)

// InjectResourceConversionTestsID is the unique identifier for this stage
const InjectResourceConversionTestsID = "injectResourceConversionTestCases"

// InjectResourceConversionTestCases is a pipeline stage to inject test cases
func InjectResourceConversionTestCases(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectResourceConversionTestsID,
		"Add test cases to verify Resource implementations of conversion.Convertible (funcs ConvertTo & ConvertFrom) behave as expected",
		func(ctx context.Context, state *State) (*State, error) {
			gopterFactory := makeResourceConversionTestCaseFactory(idFactory)
			rapidFactory := makeRapidResourceConversionTestCaseFactory(idFactory)
			modifiedDefs := make(astmodel.TypeDefinitionSet)
			var errs []error
			for _, d := range state.Definitions() {
				useRapid := false
				if ref, ok := d.Name().PackageReference().(astmodel.InternalPackageReference); ok {
					useRapid = testcases.UseRapidForGroup(ref.Group())
				}

				if useRapid {
					if rapidFactory.NeedsTest(d) {
						updated, err := rapidFactory.AddTestTo(d)
						if err != nil {
							errs = append(errs, err)
						} else {
							modifiedDefs[updated.Name()] = updated
						}
					}
				} else {
					if gopterFactory.NeedsTest(d) {
						updated, err := gopterFactory.AddTestTo(d)
						if err != nil {
							errs = append(errs, err)
						} else {
							modifiedDefs[updated.Name()] = updated
						}
					}
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			return state.WithOverlaidDefinitions(modifiedDefs), nil
		})

	stage.RequiresPrerequisiteStages(
		InjectPropertyAssignmentFunctionsStageID, // Need PropertyAssignmentFunctions to test
		ImplementConvertibleInterfaceStageID,     // Need the conversions.Convertible interface to be present
		InjectJSONSerializationTestsID,           // We reuse the generators from the JSON tests
		InjectRapidSerializationTestsStageID)     // We reuse the generators from the rapid JSON tests

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
func (*resourceConversionTestCaseFactory) NeedsTest(def astmodel.TypeDefinition) bool {
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
		return astmodel.TypeDefinition{}, eris.Errorf("expected %s to be a resourceType", def.Name())
	}

	testCase, err := testcases.NewResourceConversionTestCase(def.Name(), resource, factory.idFactory)
	if err != nil {
		return astmodel.TypeDefinition{}, eris.Wrapf(err, "adding resource conversion test case to %s", def.Name())
	}

	return factory.injector.Inject(def, testCase)
}

// rapidResourceConversionTestCaseFactory is a factory for injecting rapid-based resource conversion test cases
type rapidResourceConversionTestCaseFactory struct {
	injector  *astmodel.TestCaseInjector
	idFactory astmodel.IdentifierFactory
}

func makeRapidResourceConversionTestCaseFactory(idFactory astmodel.IdentifierFactory) rapidResourceConversionTestCaseFactory {
	return rapidResourceConversionTestCaseFactory{
		injector:  astmodel.NewTestCaseInjector(),
		idFactory: idFactory,
	}
}

func (*rapidResourceConversionTestCaseFactory) NeedsTest(def astmodel.TypeDefinition) bool {
	resourceType, ok := astmodel.AsResourceType(def.Type())
	if !ok {
		return false
	}

	_, found := resourceType.FindInterface(astmodel.ConvertibleInterface)
	return found
}

func (factory *rapidResourceConversionTestCaseFactory) AddTestTo(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	resource, ok := astmodel.AsResourceType(def.Type())
	if !ok {
		return astmodel.TypeDefinition{}, eris.Errorf("expected %s to be a resourceType", def.Name())
	}

	testCase, err := testcases.NewRapidResourceConversionTestCase(def.Name(), resource, factory.idFactory)
	if err != nil {
		return astmodel.TypeDefinition{}, eris.Wrapf(err, "adding rapid resource conversion test case to %s", def.Name())
	}

	return factory.injector.Inject(def, testCase)
}
