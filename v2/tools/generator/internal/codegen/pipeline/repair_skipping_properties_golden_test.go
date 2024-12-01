/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

// TestGolden_RepairSkippingProperties_WhenPropertyTypesDiffer_GeneratesExpectedCode checks that the pipeline stage
// correctly injects new types with the right functions when properties of different types skip versions.
func TestGolden_RepairSkippingProperties_WhenPropertyTypesDiffer_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange - create multiple versions of Person, where the Residence goes missing and returns with a different shape
	residenceV1 := test.CreateObjectDefinition(test.Pkg2020, "Residence", test.FullAddressProperty, test.CityProperty)
	residenceV3 := test.CreateObjectDefinition(test.Pkg2022, "Residence", test.FullAddressProperty, test.SuburbProperty, test.CityProperty)

	residenceV1Prop := astmodel.NewPropertyDefinition("Residence", "residence", residenceV1.Name())
	residenceV3Prop := astmodel.NewPropertyDefinition("Residence", "residence", residenceV3.Name())

	personV1 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, residenceV1Prop)
	personV2 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.FamilyNameProperty)
	personV3 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty, test.FamilyNameProperty, residenceV3Prop)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(personV1, personV2, personV3, residenceV1, residenceV3)

	// Arrange - create our initial state, prior to running the Repairer
	cfg := config.NewConfiguration()
	idFactory := astmodel.NewIdentifierFactory()
	initialState, err := RunTestPipeline(
		NewState(defs),
		CreateStorageTypes(),            // First create the storage types
		CreateConversionGraph(cfg, "v"), // Then, create the conversion graph showing relationships
	)
	g.Expect(err).To(Succeed())

	// Assert - Create a comparison state that has the same code, but without a repair
	comparisonState, err := RunTestPipeline(
		initialState,
		InjectPropertyAssignmentFunctions(cfg, idFactory, logr.Discard()),
	)
	g.Expect(err).To(BeNil())

	// Act - run the Repairer stage
	finalState, err := RunTestPipeline(
		initialState,
		RepairSkippingProperties(),      // and then we get to run the stage we're testing
		CreateConversionGraph(cfg, "v"), // Then, RECREATE the conversion graph showing relationships
		InjectPropertyAssignmentFunctions(cfg, idFactory, logr.Discard()),
	)
	g.Expect(err).To(BeNil())

	// Assert - check that we get the expected code changes
	// When reviewing the golden file, look for the code that's removed ('-') or added ('+') as a result of
	// including the repair stage in the run.
	test.AssertPackagesGenerateExpectedCode(
		t,
		finalState.Definitions(),
		test.DiffWith(comparisonState.Definitions().AsSlice()...))
}
