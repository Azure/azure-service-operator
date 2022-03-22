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

func TestGolden_InjectConvertibleStatusInterface(t *testing.T) {
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

	initialState := NewState().WithDefinitions(defs)

	cfg := config.NewConfiguration()
	finalState, err := RunTestPipeline(
		initialState,
		CreateStorageTypes(),            // First create the storage types
		CreateConversionGraph(cfg, "v"), // Then, create the conversion graph showing relationships
		InjectPropertyAssignmentFunctions(cfg, idFactory), // After which we inject property assignment functions
		ImplementConvertibleStatusInterface(idFactory),    // And then we get to run the stage we're testing
	)
	g.Expect(err).To(Succeed())

	// When verifying the golden file, check that the implementations of ConvertStatusTo() and ConvertStatusFrom() are
	// correctly injected on the specs, but not on the other types. Verify that the code does what you expect. If you
	// don't know what to expect, check that the non-hub status types have chained conversions that use the property
	// assignment functions, and that the hub spec type has a pivot conversion.
	test.AssertPackagesGenerateExpectedCode(t, finalState.definitions)
}
