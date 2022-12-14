/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGolden_InjectOriginalGVKFunction(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec)

	initialState, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		InjectOriginalVersionFunction(idFactory))
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(
		initialState,
		InjectOriginalGVKFunction(idFactory))
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions(), test.DiffWithTypes(initialState.Definitions()))
}
