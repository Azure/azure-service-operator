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

func TestGolden_ResoureExtension(t *testing.T) {
	g := NewGomegaWithT(t)

	// Test Resource V1

	specV1 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	statusV1 := test.CreateStatus(test.Pkg2020, "Person")
	resourceV1 := test.CreateResource(test.Pkg2020, "Person", specV1, statusV1)

	types := make(astmodel.Types)
	types.AddAll(resourceV1, specV1, statusV1)

	initialState, err := RunTestPipeline(
		NewState().WithTypes(types))
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(initialState, CreateResourceExtensions("testPath", astmodel.NewIdentifierFactory()))
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Types())
}
