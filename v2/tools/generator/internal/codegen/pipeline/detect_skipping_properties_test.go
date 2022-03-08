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

func TestSkippingPropertyDetector_AddProperty_CreatesExpectedChain(t *testing.T) {
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
	builder.Add(test.Pkg2020, test.Pkg2020s)
	builder.Add(test.Pkg2021, test.Pkg2021s)
	builder.Add(test.Pkg2022, test.Pkg2022s)
	graph, err := builder.Build()
	g.Expect(err).NotTo(HaveOccurred())

	detector := newSkippingPropertyDetector(defs, graph)
	err = detector.AddProperties(person2020.Name(), test.FullNameProperty)
	g.Expect(err).NotTo(HaveOccurred())

	// We expect to have four links:
	// Pkg2020.Person.FullName => Person2020storage.Person.FullName
	// Pkg2020storage.Person.FullName => Person2021storage.Person.FullName
	// Pkg2021storage.Person.FullName => Person2022storage.Person.FullName
	// Pkg2022storage.Person.FullName => EmptyReference
	g.Expect(detector.links).To(HaveLen(4))
	AssertLinkExists(g, detector, person2020, test.FullNameProperty.PropertyName())
	AssertLinkExists(g, detector, person2020s, test.FullNameProperty.PropertyName())
	AssertLinkExists(g, detector, person2021s, test.FullNameProperty.PropertyName())
	AssertLinkExists(g, detector, person2022s, test.FullNameProperty.PropertyName())
}

func TestSkippingPropertyDetector_findBreak_returnsExpectedResults(t *testing.T) {
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
	detector := newSkippingPropertyDetector(defs, graph)

	detector.addLink(alphaSeen, betaSeen)
	detector.addLink(betaSeen, gammaSeen)
	detector.addLink(gammaSeen, deltaMissing)
	detector.addLink(deltaMissing, epsilonMissing)
	detector.addLink(epsilonMissing, zetaMissing)
	detector.addLink(zetaMissing, thetaSeen)
	detector.addLink(thetaSeen, iotaSeen)
	detector.addLink(iotaSeen, kappaSeen)

	detector.propertyObserved(alphaSeen)
	detector.propertyObserved(betaSeen)
	detector.propertyObserved(gammaSeen)
	detector.propertyObserved(thetaSeen)
	detector.propertyObserved(iotaSeen)
	detector.propertyObserved(kappaSeen)

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

			actualBefore, actualAfter := detector.findBreak(c.ref, detector.wasPropertyObserved)
			g.Expect(actualBefore).To(Equal(c.expectedBefore))
			g.Expect(actualAfter).To(Equal(c.expectedAfter))
		})
	}
}

// Test_DetectSkippingProperties checks that the pipeline stage correctly skips when properties of identical types skip versions.
func Test_DetectSkippingProperties_WhenPropertyTypesIdentical_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create multiple versions of person, with KnownAs missing
	personV1 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	personV2 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.FamilyNameProperty)
	personV3 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(personV1, personV2, personV3)

	cfg := config.NewConfiguration()
	initialState, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		CreateStorageTypes(),            // First create the storage types
		CreateConversionGraph(cfg, "v"), // Then, create the conversion graph showing relationships
	)
	g.Expect(err).To(Succeed())

	_, err = RunTestPipeline(
		initialState,
		DetectSkippingProperties(), // and then we get to run the stage we're testing
	)

	g.Expect(err).To(Succeed())
}

// Test_DetectSkippingProperties checks that the pipeline stage correctly detects when properties of different types skip versions.
func Test_DetectSkippingProperties_WhenPropertyTypesDiffer_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vipNumber := astmodel.NewPropertyDefinition("VIP", "vip", astmodel.IntType)
	vipID := astmodel.NewPropertyDefinition("VIP", "vip", astmodel.StringType)

	// Create multiple versions of person, with KnownAs missing
	personV1 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, vipNumber)
	personV2 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.FamilyNameProperty)
	personV3 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty, test.FamilyNameProperty, vipID)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(personV1, personV2, personV3)

	cfg := config.NewConfiguration()
	initialState, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		CreateStorageTypes(),            // First create the storage types
		CreateConversionGraph(cfg, "v"), // Then, create the conversion graph showing relationships
	)
	g.Expect(err).To(Succeed())

	_, err = RunTestPipeline(
		initialState,
		DetectSkippingProperties(), // and then we get to run the stage we're testing
	)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("person/v20200101storage/Person_Spec.VIP was discontinued"))
	g.Expect(err.Error()).To(ContainSubstring("reintroduced as person/v20220630storage/Person_Spec.VIP"))
	// Make sure the error message links to GitHub
	g.Expect(err.Error()).To(ContainSubstring("https://github.com/Azure/azure-service-operator/issues/1776"))
}

func AssertLinkExists(
	g *WithT,
	detector *skippingPropertyDetector,
	def astmodel.TypeDefinition,
	property astmodel.PropertyName) {
	ref := astmodel.MakePropertyReference(def.Name(), property)
	g.Expect(detector.links).To(HaveKey(ref))
}

func createPropertyRef(version string, name astmodel.PropertyName) astmodel.PropertyReference {
	tn := astmodel.MakeTypeName(test.MakeLocalPackageReference("greek", version), "test")
	return astmodel.MakePropertyReference(tn, name)
}
