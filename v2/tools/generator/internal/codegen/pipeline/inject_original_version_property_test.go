/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestInjectOriginalVersionProperty_InjectsIntoSpec(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec)

	initialState := NewState().WithDefinitions(defs)

	finalState, err := RunTestPipeline(
		initialState,
		InjectOriginalVersionProperty())
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions(), test.DiffWithTypes(initialState.Definitions()))
}

func TestInjectOriginalVersionProperty_WhenOriginalVersionFunctionFound_DoesNotInjectIntoSpec(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	fnInjector := astmodel.NewFunctionInjector()

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	spec, err := fnInjector.Inject(spec, functions.NewOriginalVersionFunction(idFactory))
	g.Expect(err).To(Succeed())

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec)

	initialState, err := RunTestPipeline(
		NewState(defs),
		InjectOriginalVersionFunction(idFactory))
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(
		initialState,
		InjectOriginalVersionProperty())
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions(), test.DiffWithTypes(initialState.Definitions()))
}
