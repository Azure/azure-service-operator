/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

// TestGolden_AddStatusConditions checks that the Add Status Conditions pipeline stage does what we expect
func TestGolden_AddStatusConditions(t *testing.T) {
	g := NewWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resourceV1 := test.CreateResource(test.Pkg2020, "Person", spec, status)

	types := make(astmodel.Types)
	types.AddAll(resourceV1, spec, status)

	initialState := NewState().WithTypes(types)
	finalState, err := RunTestPipeline(
		initialState,
		AddStatusConditions(idFactory))
	g.Expect(err).To(Succeed())

	// When verifying the golden file, check to ensure that the Conditions property on the Status type looks correct,
	// and that the conditions.Conditioner interface is properly implemented on the resource.
	test.AssertPackagesGenerateExpectedCode(t, finalState.types, test.DiffWithTypes(types))
}
