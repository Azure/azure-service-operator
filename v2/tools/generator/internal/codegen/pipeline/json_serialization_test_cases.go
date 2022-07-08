/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/testcases"

	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// InjectJsonSerializationTestsID is the unique identifier for this pipeline stage
const InjectJsonSerializationTestsID = "injectJSONTestCases"

func InjectJsonSerializationTests(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectJsonSerializationTestsID,
		"Add test cases to verify JSON serialization",
		func(ctx context.Context, state *State) (*State, error) {
			factory := makeObjectSerializationTestCaseFactory(idFactory)
			modifiedDefinitions := make(astmodel.TypeDefinitionSet)
			var errs []error
			for _, def := range state.Definitions() {
				if factory.NeedsTest(def) {
					updated, err := factory.AddTestTo(def)
					if err != nil {
						errs = append(errs, err)
					} else {
						modifiedDefinitions[updated.Name()] = updated
					}
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			return state.WithDefinitions(state.Definitions().OverlayWith(modifiedDefinitions)), nil
		})

	stage.RequiresPostrequisiteStages("simplifyDefinitions" /* needs flags */)

	return stage
}

type objectSerializationTestCaseFactory struct {
	injector     *astmodel.TestCaseInjector
	idFactory    astmodel.IdentifierFactory
	suppressions []string
}

func makeObjectSerializationTestCaseFactory(idFactory astmodel.IdentifierFactory) objectSerializationTestCaseFactory {
	result := objectSerializationTestCaseFactory{
		injector:     astmodel.NewTestCaseInjector(),
		idFactory:    idFactory,
		suppressions: []string{},
	}

	return result
}

// NeedsTest returns true if we should generate a testcase for the specified definition
func (s *objectSerializationTestCaseFactory) NeedsTest(def astmodel.TypeDefinition) bool {
	pc, ok := astmodel.AsPropertyContainer(def.Type())
	if !ok {
		// Can only generate tests for property containers
		return false
	}

	// No test needed for types with no properties
	if len(pc.Properties().AsSlice()) == 0 {
		return false
	}

	// Check for types that we need to suppress - these are ARM types that don't currently round trip because they're
	// OneOf implementations that are only used in one direction.
	//
	// See https://github.com/Azure/azure-service-operator/issues/1721 for more information
	//
	result := true
	for _, s := range s.suppressions {
		if def.Name().Name() == s {
			result = false
			break
		}
	}

	return result
}

func (s *objectSerializationTestCaseFactory) AddTestTo(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	container, ok := astmodel.AsPropertyContainer(def.Type())
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("expected %s to be a property container", def.Name())
	}

	isOneOf := astmodel.OneOfFlag.IsOn(def.Type()) // this is ugly but can’t do much better right now

	// How many tests we want to run for resources
	const resourceTestCount = 20

	// How many tests to run just on spec types (gopter default is 100)
	const specTestCount = 100 - resourceTestCount

	// How many tests to run just on status types (gopter default is 100)
	const statusTestCount = 100 - resourceTestCount

	testcase := testcases.NewJSONSerializationTestCase(def.Name(), container, isOneOf, s.idFactory)
	if _, isResource := astmodel.AsResourceType(def.Type()); isResource {
		// Don't need to test resources many times, the spec and status types are tested independently
		testcase.SetMinSuccessfulTests(resourceTestCount)
	} else if astmodel.IsSpec(def.Name()) {
		// Reduce count of Spec and Status tests to reflect those done by the resource tests
		testcase.SetMinSuccessfulTests(specTestCount)
	} else if astmodel.IsStatus(def.Name()) {
		// Reduce count of Spec and Status tests to reflect those done by the resource tests
		testcase.SetMinSuccessfulTests(statusTestCount)
	}

	return s.injector.Inject(def, testcase)
}
