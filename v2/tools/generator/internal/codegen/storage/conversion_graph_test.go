/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestConversionGraph_WithTwoUnrelatedReferences_HasExpectedTransitions(t *testing.T) {
	/*
	 *  Test that a conversion graph that contains two API package references from different groups ends up with just
	 *  one transition for each group, each linking from the provided API version to the matching storage variant.
	 */
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := config.NewObjectModelConfiguration()
	builder := NewConversionGraphBuilder(omc)
	builder.Add(test.Pkg2020)
	builder.Add(test.BatchPkg2020)
	graph, err := builder.Build()

	// Check size of graph
	g.Expect(err).To(Succeed())
	g.Expect(graph.TransitionCount()).To(Equal(2))

	// Check for the expected transition from Pkg2020
	pkg, ok := graph.LookupTransition(test.Pkg2020)
	g.Expect(ok).To(BeTrue())
	g.Expect(pkg).NotTo(BeNil())
	g.Expect(astmodel.IsStoragePackageReference(pkg)).To(BeTrue())

	// Check for the expected transition from BatchPkg2020
	pkg, ok = graph.LookupTransition(test.BatchPkg2020)
	g.Expect(ok).To(BeTrue())
	g.Expect(pkg).NotTo(BeNil())
	g.Expect(astmodel.IsStoragePackageReference(pkg)).To(BeTrue())
}

func TestConversionGraph_GivenTypeName_ReturnsExpectedHubTypeName(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create some resources to use for testing.
	// Need three versions of the Person resource, plus the storage variants of those
	person2020 := test.CreateSimpleResource(test.Pkg2020, "Person")
	person2020s := test.CreateSimpleResource(test.Pkg2020s, "Person")
	person2021 := test.CreateSimpleResource(test.Pkg2021, "Person")
	person2021s := test.CreateSimpleResource(test.Pkg2021s, "Person")
	person2022 := test.CreateSimpleResource(test.Pkg2022, "Person")
	person2022s := test.CreateSimpleResource(test.Pkg2022s, "Person")

	// Need two versions of the Address resource, plus the storage variants of those
	address2020 := test.CreateSimpleResource(test.Pkg2020, "Address")
	address2020s := test.CreateSimpleResource(test.Pkg2020s, "Address")
	address2021 := test.CreateSimpleResource(test.Pkg2021, "Address")
	address2021s := test.CreateSimpleResource(test.Pkg2021s, "Address")

	// Create our set of types
	types := make(astmodel.Types)
	types.AddAll(person2020, person2021, person2022, address2020, address2021)
	types.AddAll(person2020s, person2021s, person2022s, address2020s, address2021s)

	// Create a builder use it to configure a graph to test
	omc := config.NewObjectModelConfiguration()
	builder := NewConversionGraphBuilder(omc)
	builder.Add(person2020.Name().PackageReference)
	builder.Add(person2021.Name().PackageReference)
	builder.Add(person2022.Name().PackageReference)

	graph, err := builder.Build()
	g.Expect(err).To(Succeed())

	// We expect all person resources to use this hub
	personHub := person2022s.Name()

	// Address resources use this as the hub because Address doesn't exist in Pkg2022
	addressHub := address2021s.Name()

	cases := []struct {
		name         string
		start        astmodel.TypeName
		expectedName astmodel.TypeName
	}{
		{"Hub type returns self", personHub, personHub},
		{"Directly linked api resolves", person2022.Name(), personHub},
		{"Directly linked storage resolves", person2021s.Name(), personHub},
		{"Doubly linked api resolves", person2021.Name(), personHub},
		{"Indirectly linked api resolves", person2020.Name(), personHub},
		{"Indirectly linked storage resolves", person2020s.Name(), personHub},
		{"Hub type returns self even when resource does not exist in latest package", addressHub, addressHub},
		{"Directly linked api resolves when resource does not exist in latest package", address2021.Name(), addressHub},
		{"Indirectly linked api resolves when resource does not exist in latest package", address2020.Name(), addressHub},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			t.Parallel()

			actual, err := graph.FindHub(c.start, types)
			g.Expect(err).To(Succeed())
			g.Expect(actual).To(Equal(c.expectedName))
		})
	}
}

func Test_ConversionGraph_WhenRenameConfigured_FindsRenamedType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create some resources to use for testing.
	// Need both types and the storage variations
	person2020 := test.CreateSimpleResource(test.Pkg2020, "Person")
	person2020s := test.CreateSimpleResource(test.Pkg2020s, "Person")

	party2021 := test.CreateSimpleResource(test.Pkg2021, "Party")
	party2021s := test.CreateSimpleResource(test.Pkg2021s, "Party")

	// Create our set of types
	types := make(astmodel.Types)
	types.AddAll(person2020, party2021)
	types.AddAll(person2020s, party2021s)

	// Create configuration for our rename
	tc := config.NewTypeConfiguration(person2020.Name().Name()).SetTypeRename(party2021.Name().Name())
	vc := config.NewVersionConfiguration("v20200101").Add(tc)
	gc := config.NewGroupConfiguration(test.Pkg2020.Group()).Add(vc)
	omc := config.NewObjectModelConfiguration().Add(gc)

	// Create a builder use it to configure a graph to test
	builder := NewConversionGraphBuilder(omc)
	builder.Add(person2020.Name().PackageReference)
	builder.Add(party2021.Name().PackageReference)

	graph, err := builder.Build()
	g.Expect(err).To(Succeed())

	name, err := graph.FindNextType(person2020s.Name(), types)
	g.Expect(err).To(Succeed())
	g.Expect(name).To(Equal(party2021s.Name()))
}

func Test_ConversionGraph_WhenRenameSpecifiesMissingType_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	// Create some resources to use for testing.
	// Need both types and the storage variations
	person2020 := test.CreateSimpleResource(test.Pkg2020, "Person")
	person2020s := test.CreateSimpleResource(test.Pkg2020s, "Person")

	party2021 := test.CreateSimpleResource(test.Pkg2021, "Party")
	party2021s := test.CreateSimpleResource(test.Pkg2021s, "Party")

	// Create our set of types
	types := make(astmodel.Types)
	types.AddAll(person2020, party2021)
	types.AddAll(person2020s, party2021s)

	// Create mis-configuration for our rename specifying a type that doesn't exist
	tc := config.NewTypeConfiguration(person2020.Name().Name()).SetTypeRename("Phantom")
	vc := config.NewVersionConfiguration("v20200101").Add(tc)
	gc := config.NewGroupConfiguration(test.Pkg2020.Group()).Add(vc)
	omc := config.NewObjectModelConfiguration().Add(gc)

	// Create a builder use it to configure a graph to test
	builder := NewConversionGraphBuilder(omc)
	builder.Add(person2020.Name().PackageReference)
	builder.Add(party2021.Name().PackageReference)

	graph, err := builder.Build()
	g.Expect(err).To(Succeed())

	_, err = graph.FindNextType(person2020s.Name(), types)
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring("Phantom"))
}

func Test_ConversionGraph_WhenRenameSpecifiesConflictingType_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create some resources to use for testing.
	// Need both types and the storage variations
	person2020 := test.CreateSimpleResource(test.Pkg2020, "Person")
	person2020s := test.CreateSimpleResource(test.Pkg2020s, "Person")

	person2021 := test.CreateSimpleResource(test.Pkg2021, "Person")
	person2021s := test.CreateSimpleResource(test.Pkg2021s, "Person")

	party2021 := test.CreateSimpleResource(test.Pkg2021, "Party")
	party2021s := test.CreateSimpleResource(test.Pkg2021s, "Party")

	// Create our set of types
	types := make(astmodel.Types)
	types.AddAll(person2020, person2021, party2021)
	types.AddAll(person2020s, person2021s, party2021s)

	// Create mis-configuration for our rename that conflicts with the second type
	tc := config.NewTypeConfiguration(person2020.Name().Name()).SetTypeRename(party2021.Name().Name())
	vc := config.NewVersionConfiguration("v20200101").Add(tc)
	gc := config.NewGroupConfiguration(test.Pkg2020.Group()).Add(vc)
	omc := config.NewObjectModelConfiguration().Add(gc)

	// Create a builder use it to configure a graph to test
	builder := NewConversionGraphBuilder(omc)
	builder.Add(person2020.Name().PackageReference)
	builder.Add(person2021.Name().PackageReference)

	graph, err := builder.Build()
	g.Expect(err).To(Succeed())

	_, err = graph.FindNextType(person2020s.Name(), types)
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring(person2020.Name().Name()))
	g.Expect(err.Error()).To(ContainSubstring(party2021.Name().Name()))
}
