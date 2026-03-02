/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/testcases"
)

// InjectRapidSerializationTestsStageID is the unique identifier for this pipeline stage
const InjectRapidSerializationTestsStageID = "injectRapidSerializationTests"

func InjectRapidSerializationTests(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectRapidSerializationTestsStageID,
		"Add rapid-based test cases to verify JSON serialization",
		func(ctx context.Context, state *State) (*State, error) {
			factory := makeRapidSerializationTestCaseFactory(idFactory)
			modifiedDefinitions := make(astmodel.TypeDefinitionSet)
			var errs []error
			for _, def := range state.Definitions() {
				ref, ok := def.Name().PackageReference().(astmodel.InternalPackageReference)
				if !ok || !testcases.UseRapidForGroup(ref.Group()) {
					continue
				}

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

			return state.WithOverlaidDefinitions(modifiedDefinitions), nil
		})

	stage.RequiresPostrequisiteStages("simplifyDefinitions")

	return stage
}

type rapidSerializationTestCaseFactory struct {
	injector  *astmodel.TestCaseInjector
	idFactory astmodel.IdentifierFactory
}

func makeRapidSerializationTestCaseFactory(idFactory astmodel.IdentifierFactory) rapidSerializationTestCaseFactory {
	return rapidSerializationTestCaseFactory{
		injector:  astmodel.NewTestCaseInjector(),
		idFactory: idFactory,
	}
}

// NeedsTest returns true if we should generate a testcase for the specified definition
func (s *rapidSerializationTestCaseFactory) NeedsTest(def astmodel.TypeDefinition) bool {
	_, ok := astmodel.AsPropertyContainer(def.Type())
	if !ok {
		return false
	}

	if astmodel.IsWebhookPackageReference(def.Name().PackageReference()) {
		return false
	}

	return true
}

func (s *rapidSerializationTestCaseFactory) AddTestTo(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	container, ok := astmodel.AsPropertyContainer(def.Type())
	if !ok {
		return astmodel.TypeDefinition{}, eris.Errorf("expected %s to be a property container", def.Name())
	}

	isOneOf := astmodel.OneOfFlag.IsOn(def.Type())

	testcase := testcases.NewRapidJSONSerializationTestCase(def.Name(), container, isOneOf, s.idFactory)
	return s.injector.Inject(def, testcase)
}
