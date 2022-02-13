/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGolden_CreateStorageTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Test Resource V1

	specV1 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	statusV1 := test.CreateStatus(test.Pkg2020, "Person")
	resourceV1 := test.CreateResource(test.Pkg2020, "Person", specV1, statusV1)

	// Test Resource V2

	specV2 := test.CreateSpec(
		test.Pkg2021,
		"Person",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty,
		test.ResidentialAddress2021,
		test.PostalAddress2021)
	statusV2 := test.CreateStatus(test.Pkg2021, "Person")
	resourceV2 := test.CreateResource(test.Pkg2021, "Person", specV2, statusV2)

	types := make(astmodel.TypeDefinitionSet)
	types.AddAll(resourceV1, specV1, statusV1, resourceV2, specV2, statusV2, test.Address2021)

	initialState, err := RunTestPipeline(
		NewState().WithTypes(types),
		CreateConversionGraph(config.NewConfiguration()))
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(initialState, CreateStorageTypes())
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Types())
}
