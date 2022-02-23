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

// TestGolden_AddStatusConditions checks that the Add Status Conditions pipeline stage does what we expect
func TestGolden_AddStatusConditions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resourceV1 := test.CreateResource(test.Pkg2020, "Person", spec, status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resourceV1, spec, status)

	initialState := NewState().WithDefinitions(defs)
	finalState, err := RunTestPipeline(
		initialState,
		AddStatusConditions(idFactory))
	g.Expect(err).To(Succeed())

	// When verifying the golden file, check to ensure that the Conditions property on the Status type looks correct,
	// and that the conditions.Conditioner interface is properly implemented on the resource.
	test.AssertPackagesGenerateExpectedCode(t, finalState.definitions, test.DiffWithTypes(defs))
}
