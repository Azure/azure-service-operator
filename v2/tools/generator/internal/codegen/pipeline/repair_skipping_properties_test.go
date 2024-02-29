/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestSkippingPropertyRepairer_AddProperty_CreatesExpectedChain(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// API versions
	person2020 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	person2021 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty)
	person2022 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty)

	// Storage versions
	person2020s := test.CreateSpec(test.Pkg2020s, "Person", test.FullNameProperty)
	person2021s := test.CreateSpec(test.Pkg2021s, "Person", test.FullNameProperty)
	person2022s := test.CreateSpec(test.Pkg2022s, "Person", test.FullNameProperty)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, person2021, person2022)
	defs.AddAll(person2020s, person2021s, person2022s)

	cfg := config.NewObjectModelConfiguration()
	builder := storage.NewConversionGraphBuilder(cfg, "v")
	builder.Add(person2020.Name(), person2020s.Name())
	builder.Add(person2021.Name(), person2021s.Name())
	builder.Add(person2022.Name(), person2022s.Name())
	graph, err := builder.Build()
	g.Expect(err).NotTo(HaveOccurred())

	repairer := newSkippingPropertyRepairer(defs, graph)
	err = repairer.AddProperties(person2020.Name(), test.FullNameProperty)
	g.Expect(err).NotTo(HaveOccurred())

	// We expect to have four links:
	// Pkg2020.Person.FullName => Person2020storage.Person.FullName
	// Pkg2020storage.Person.FullName => Person2021storage.Person.FullName
	// Pkg2021storage.Person.FullName => Person2022storage.Person.FullName
	// Pkg2022storage.Person.FullName => EmptyReference
	g.Expect(repairer.links).To(HaveLen(4))
	AssertLinkExists(g, repairer, person2020, test.FullNameProperty.PropertyName())
	AssertLinkExists(g, repairer, person2020s, test.FullNameProperty.PropertyName())
	AssertLinkExists(g, repairer, person2021s, test.FullNameProperty.PropertyName())
	AssertLinkExists(g, repairer, person2022s, test.FullNameProperty.PropertyName())
}

func TestSkippingPropertyRepairer_findBreak_returnsExpectedResults(t *testing.T) {
	t.Parallel()

	// These property references all have different names, for ease of inspection
	alphaSeen := createPropertyRef("v1", "alpha")
	betaSeen := createPropertyRef("v2", "beta")
	gammaSeen := createPropertyRef("v3", "gamma")
	deltaMissing := createPropertyRef("v4", "delta")
	epsilonMissing := createPropertyRef("v5", "epsilon")
	zetaMissing := createPropertyRef("v6", "zeta")
	thetaSeen := createPropertyRef("v7", "theta")
	iotaSeen := createPropertyRef("v8", "iota")
	kappaSeen := createPropertyRef("v9", "kappa")
	empty := astmodel.EmptyPropertyReference

	defs := make(astmodel.TypeDefinitionSet)
	cfg := config.NewObjectModelConfiguration()
	graph, _ := storage.NewConversionGraphBuilder(cfg, "v").Build()
	repairer := newSkippingPropertyRepairer(defs, graph)

	repairer.addLink(alphaSeen, betaSeen)
	repairer.addLink(betaSeen, gammaSeen)
	repairer.addLink(gammaSeen, deltaMissing)
	repairer.addLink(deltaMissing, epsilonMissing)
	repairer.addLink(epsilonMissing, zetaMissing)
	repairer.addLink(zetaMissing, thetaSeen)
	repairer.addLink(thetaSeen, iotaSeen)
	repairer.addLink(iotaSeen, kappaSeen)

	repairer.propertyObserved(alphaSeen)
	repairer.propertyObserved(betaSeen)
	repairer.propertyObserved(gammaSeen)
	repairer.propertyObserved(thetaSeen)
	repairer.propertyObserved(iotaSeen)
	repairer.propertyObserved(kappaSeen)

	cases := []struct {
		name           string
		ref            astmodel.PropertyReference
		expectedBefore astmodel.PropertyReference
		expectedAfter  astmodel.PropertyReference
	}{
		{"At end of chain of observedProperties properties", kappaSeen, kappaSeen, empty},
		{"When rest of chain all observedProperties", thetaSeen, kappaSeen, empty},
		{"At end of chain of missing properties", zetaMissing, zetaMissing, thetaSeen},
		{"In midst of missing properties", deltaMissing, zetaMissing, thetaSeen},
		{"At end of first run of observedProperties properties", gammaSeen, gammaSeen, deltaMissing},
		{"In midst of first run of observedProperties properties", betaSeen, gammaSeen, deltaMissing},
		{"At start of chain of observedProperties properties", alphaSeen, gammaSeen, deltaMissing},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actualBefore, actualAfter := repairer.findBreak(c.ref, repairer.wasPropertyObserved)
			g.Expect(actualBefore).To(Equal(c.expectedBefore))
			g.Expect(actualAfter).To(Equal(c.expectedAfter))
		})
	}
}

// Test_RepairSkippingProperties checks that the pipeline stage correctly skips when properties of identical types skip versions.
func Test_RepairSkippingProperties_WhenPropertyTypesIdentical_DoesNotChangeDefinitions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange - create multiple versions of person, with KnownAs missing
	personV1 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	personV2 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.FamilyNameProperty)
	personV3 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(personV1, personV2, personV3)

	// Arrange - create our initial state, prior to running the Repairer
	cfg := config.NewConfiguration()
	initialState, err := RunTestPipeline(
		NewState(defs),
		CreateStorageTypes(),            // First create the storage types
		CreateConversionGraph(cfg, "v"), // Then, create the conversion graph showing relationships
	)
	g.Expect(err).To(Succeed())

	// Act - run the Repairer stage
	finalState, err := RunTestPipeline(
		initialState,
		RepairSkippingProperties(), // and then we get to run the stage we're testing
	)

	// Assert - we expect no error, and no new definitions
	g.Expect(err).To(BeNil())
	g.Expect(finalState.definitions).To(HaveLen(6)) // Our three original resources, plus the storage variants
}

func Test_RepairSkippingProperties_WhenPropertyStructurlyIdentical_DoesNotChangeDefinitions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange - create multiple versions of Person, where the Residence goes missing but returns with the same shape
	residenceV1 := test.CreateObjectDefinition(test.Pkg2020, "Residence", test.FullAddressProperty, test.CityProperty)
	residenceV3 := test.CreateObjectDefinition(test.Pkg2022, "Residence", test.FullAddressProperty, test.CityProperty)

	residenceV1Prop := astmodel.NewPropertyDefinition("Residence", "residence", residenceV1.Name())
	residenceV3Prop := astmodel.NewPropertyDefinition("Residence", "residence", residenceV3.Name())

	// Create multiple versions of person, with Residence missing
	personV1 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, residenceV1Prop)
	personV2 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.FamilyNameProperty)
	personV3 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty, test.FamilyNameProperty, residenceV3Prop)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(personV1, personV2, personV3, residenceV1, residenceV3)

	// Arrange - create our initial state, prior to running the Repairer
	cfg := config.NewConfiguration()
	initialState, err := RunTestPipeline(
		NewState(defs),
		CreateStorageTypes(),            // First create the storage types
		CreateConversionGraph(cfg, "v"), // Then, create the conversion graph showing relationships
	)
	g.Expect(err).To(Succeed())

	// Act - run the Repairer stage
	finalState, err := RunTestPipeline(
		initialState,
		RepairSkippingProperties(), // and then we get to run the stage we're testing
	)
	g.Expect(err).To(BeNil())

	// Assert - we expect no error, and no new definitions
	g.Expect(err).To(BeNil())
	g.Expect(finalState.definitions).To(HaveLen(10)) // Our five original resources, plus the storage variants
}

// Test_RepairSkippingProperties checks that the pipeline stage correctly injects new types when properties of different types skip versions.
func Test_RepairSkippingProperties_WhenPropertyTypesDiffer_InjectsExpectedAdditionalDefinition(t *testing.T) {
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
	initialState, err := RunTestPipeline(
		NewState(defs),
		CreateStorageTypes(),            // First create the storage types
		CreateConversionGraph(cfg, "v"), // Then, create the conversion graph showing relationships
	)
	g.Expect(err).To(Succeed())

	// Act - run the Repairer stage
	finalState, err := RunTestPipeline(
		initialState,
		RepairSkippingProperties(), // and then we get to run the stage we're testing
	)

	// Assert - we expect no error, and one new definition
	g.Expect(err).To(BeNil())
	g.Expect(finalState.definitions).To(HaveLen(11)) // Our five original resources, plus the storage variants, plus one more

	expected := astmodel.MakeInternalTypeName(
		astmodel.MakeCompatPackageReference(test.Pkg2021s),
		"Residence")
	g.Expect(finalState.definitions).To(HaveKey(expected))
}

func AssertLinkExists(
	g *WithT,
	repairer *skippingPropertyRepairer,
	def astmodel.TypeDefinition,
	property astmodel.PropertyName,
) {
	ref := astmodel.MakePropertyReference(def.Name(), property)
	g.Expect(repairer.links).To(HaveKey(ref))
}

func createPropertyRef(version string, name astmodel.PropertyName) astmodel.PropertyReference {
	tn := astmodel.MakeInternalTypeName(test.MakeLocalPackageReference("greek", version), "test")
	return astmodel.MakePropertyReference(tn, name)
}
