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

func TestGolden_InjectJsonSerializationTests(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	// Test Resource V1

	specV1 := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)
	statusV1 := test.CreateStatus(test.Pkg2020, "Person")
	resourceV1 := test.CreateResource(test.Pkg2020, "Person", specV1, statusV1)

	// Test Resource V2
	// We need two resources to trigger creation of two storage types
	// The v1 storage type will be a non-hub, and the v2 storage type will be a hub

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

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resourceV1, specV1, statusV1, resourceV2, specV2, statusV2, test.Address2021)

	state := NewState().WithDefinitions(defs)

	cfg := config.NewConfiguration()
	finalState, err := RunTestPipeline(
		state,
		CreateStorageTypes(),       // First create the storage types
		CreateConversionGraph(cfg), // Then, create the conversion graph showing relationships
		InjectPropertyAssignmentFunctions(cfg, idFactory),
		InjectJsonSerializationTests(idFactory))
	g.Expect(err).To(Succeed())

	/*
	 * When verifying these golden tests, check the _test files contain the expected tests, test runners, and generator
	 * factory methods.
	 */
	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions(), test.IncludeTestFiles())
}
