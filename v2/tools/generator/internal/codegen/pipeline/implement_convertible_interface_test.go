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

// TestGolden_InjectConvertibleInterface checks that the pipeline stage does what we expect when run in relative isolation,
// with only a few expected (and closely related) stages in operation
func TestGolden_InjectConvertibleInterface(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	// Test Resource V1

	specV1 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	statusV1 := test.CreateStatus(test.Pkg2020, "Person")
	resourceV1 := test.CreateResource(test.Pkg2020, "Person", specV1, statusV1)

	// Test Resource V2

	specV2 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty)
	statusV2 := test.CreateStatus(test.Pkg2021, "Person")
	resourceV2 := test.CreateResource(test.Pkg2021, "Person", specV2, statusV2)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resourceV1, specV1, statusV1, resourceV2, specV2, statusV2)

	cfg := config.NewConfiguration()
	initialState, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		CreateStorageTypes(),                              // First create the storage types
		CreateConversionGraph(cfg, "v"),                   // Then, create the conversion graph showing relationships
		InjectPropertyAssignmentFunctions(cfg, idFactory), // After which we inject property assignment functions
	)
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(
		initialState,
		ImplementConvertibleInterface(idFactory), // And then we get to run the stage we're testing
	)
	g.Expect(err).To(Succeed())

	// When verifying the golden file, check that the implementations of ConvertTo() and ConvertFrom() are correctly
	// injected on the resources, but not on the other types. Verify that the code does what you expect. If you don't
	// know what to expect, check that they do the right thing. :-)
	test.AssertPackagesGenerateExpectedCode(t, finalState.definitions, test.DiffWithTypes(initialState.Definitions()))
}
