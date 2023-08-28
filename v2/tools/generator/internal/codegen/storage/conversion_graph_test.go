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

	person2020 := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	person2020s := astmodel.MakeInternalTypeName(test.Pkg2020s, "Person")

	account2020 := astmodel.MakeInternalTypeName(test.Pkg2020, "Account")
	account2020s := astmodel.MakeInternalTypeName(test.Pkg2020s, "Account")

	omc := config.NewObjectModelConfiguration()
	builder := NewConversionGraphBuilder(omc, "v")
	builder.Add(person2020, person2020s)
	builder.Add(account2020, account2020s)
	graph, err := builder.Build()

	// Check size of graph
	g.Expect(err).To(Succeed())
	g.Expect(graph.TransitionCount()).To(Equal(2))

	// Check for the expected transition from Person2020
	next := graph.LookupTransition(person2020)
	g.Expect(next).NotTo(BeNil())
	g.Expect(astmodel.IsStoragePackageReference(next.PackageReference())).To(BeTrue())

	// Check for the expected transition from Account2020
	next = graph.LookupTransition(account2020)
	g.Expect(next).NotTo(BeNil())
	g.Expect(astmodel.IsStoragePackageReference(next.PackageReference())).To(BeTrue())
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

	// Need two versions of a student resource, ensuring they skip an intermediate version
	student2020 := test.CreateSimpleResource(test.Pkg2020, "Student")
	student2020s := test.CreateSimpleResource(test.Pkg2020s, "Student")
	student2022 := test.CreateSimpleResource(test.Pkg2022, "Student")
	student2022s := test.CreateSimpleResource(test.Pkg2022s, "Student")

	// Create our set of definitions
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, person2021, person2022, address2020, address2021)
	defs.AddAll(person2020s, person2021s, person2022s, address2020s, address2021s)
	defs.AddAll(student2020, student2020s, student2022, student2022s)

	// Create a builder, and use it to configure a graph to test
	omc := config.NewObjectModelConfiguration()
	builder := NewConversionGraphBuilder(omc, "v")
	builder.Add(person2020.Name(), person2020s.Name(), person2021.Name(), person2021s.Name(), person2022.Name(), person2022s.Name())
	builder.Add(address2020.Name(), address2020s.Name(), address2021.Name(), address2021s.Name())
	builder.Add(student2020.Name(), student2020s.Name(), student2022.Name(), student2022s.Name())

	graph, err := builder.Build()
	g.Expect(err).To(Succeed())

	// We expect all person resources to use this hub
	personHub := person2022s.Name()

	// Address resources use this as the hub because Address doesn't exist in Pkg2022
	addressHub := address2021s.Name()

	// Student resources use this even though Student doesn't exist in Pkg2021
	studentHub := student2022s.Name()

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
		{"Directly linked API resolves when resource doesn't exist in intermediate package", student2022.Name(), studentHub},
		{"Indirectly linked API resolves when resource doesn't exist in intermediate package", student2020.Name(), studentHub},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			t.Parallel()

			actual, err := graph.FindHub(c.start, defs)
			g.Expect(err).To(Succeed())
			g.Expect(actual).To(Equal(c.expectedName))
		})
	}
}

func Test_ConversionGraph_WhenRenameConfigured_FindsRenamedType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create some resources to use for testing.
	// Need both definitions and the storage variations
	person2020 := test.CreateSimpleResource(test.Pkg2020, "Person")
	person2020s := test.CreateSimpleResource(test.Pkg2020s, "Person")

	party2021 := test.CreateSimpleResource(test.Pkg2021, "Party")
	party2021s := test.CreateSimpleResource(test.Pkg2021s, "Party")

	// Create our set of definitions
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, party2021)
	defs.AddAll(person2020s, party2021s)

	// Create configuration for our rename
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			person2020.Name(),
			func(tc *config.TypeConfiguration) error {
				tc.NameInNextVersion.Set(party2021.Name().Name())
				return nil
			})).
		To(Succeed())

	// Create a builder use it to configure a graph to test
	builder := NewConversionGraphBuilder(omc, "v")
	builder.Add(person2020.Name(), person2020s.Name())
	builder.Add(party2021.Name(), party2021s.Name())

	graph, err := builder.Build()
	g.Expect(err).To(Succeed())

	name, err := graph.FindNextType(person2020s.Name(), defs)
	g.Expect(err).To(Succeed())
	g.Expect(name).To(Equal(party2021s.Name()))
}

func Test_ConversionGraph_WhenRenameSpecifiesMissingType_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create some resources to use for testing.
	// Need both definitions and the storage variations
	person2020 := test.CreateSimpleResource(test.Pkg2020, "Person")
	person2020s := test.CreateSimpleResource(test.Pkg2020s, "Person")

	party2021 := test.CreateSimpleResource(test.Pkg2021, "Party")
	party2021s := test.CreateSimpleResource(test.Pkg2021s, "Party")

	// Create our set of definitions
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, party2021)
	defs.AddAll(person2020s, party2021s)

	// Create mis-configuration for our rename specifying a type that doesn't exist
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			person2020.Name(),
			func(tc *config.TypeConfiguration) error {
				tc.NameInNextVersion.Set("Phantom")
				return nil
			})).
		To(Succeed())

	// Create a builder use it to configure a graph to test
	builder := NewConversionGraphBuilder(omc, "v")
	builder.Add(person2020.Name(), person2020s.Name())
	builder.Add(party2021.Name(), party2021s.Name())

	graph, err := builder.Build()
	g.Expect(err).To(Succeed())

	_, err = graph.FindNextType(person2020s.Name(), defs)
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring("Phantom"))
}

func Test_ConversionGraph_WhenRenameSpecifiesConflictingType_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create some resources to use for testing.
	// Need both definitions and the storage variations
	person2020 := test.CreateSimpleResource(test.Pkg2020, "Person")
	person2020s := test.CreateSimpleResource(test.Pkg2020s, "Person")

	person2021 := test.CreateSimpleResource(test.Pkg2021, "Person")
	person2021s := test.CreateSimpleResource(test.Pkg2021s, "Person")

	party2021 := test.CreateSimpleResource(test.Pkg2021, "Party")
	party2021s := test.CreateSimpleResource(test.Pkg2021s, "Party")

	// Create our set of definitions
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, person2021, party2021)
	defs.AddAll(person2020s, person2021s, party2021s)

	// Create mis-configuration for our rename that conflicts with the second type
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			person2020.Name(),
			func(tc *config.TypeConfiguration) error {
				tc.NameInNextVersion.Set(party2021.Name().Name())
				return nil
			})).
		To(Succeed())

	// Create a builder use it to configure a graph to test
	builder := NewConversionGraphBuilder(omc, "v")
	builder.Add(person2020.Name(), person2020s.Name())
	builder.Add(person2021.Name(), person2021s.Name())
	builder.Add(party2021.Name(), party2021s.Name())

	graph, err := builder.Build()
	g.Expect(err).To(Succeed())

	_, err = graph.FindNextType(person2020s.Name(), defs)
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring(person2020.Name().Name()))
	g.Expect(err.Error()).To(ContainSubstring(party2021.Name().Name()))
}

func TestConversionGraph_WithAResourceOnlyInPreviewVersions_HasExpectedTransitions(t *testing.T) {
	/*
	 *  Test that a conversion graph where one type is defined only in preview versions still has the expected
	 *  transitions for that type, as well as for the other types.
	 */

	t.Parallel()
	g := NewGomegaWithT(t)

	pkg2020p := test.MakeLocalPackageReference(test.Group, "v20200101preview")
	pkg2020ps := astmodel.MakeStoragePackageReference(pkg2020p)

	pkg2021p := test.MakeLocalPackageReference(test.Group, "v20211231preview")
	pkg2021ps := astmodel.MakeStoragePackageReference(pkg2021p)

	person2020 := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	person2020s := astmodel.MakeInternalTypeName(test.Pkg2020s, "Person")

	person2021 := astmodel.MakeInternalTypeName(test.Pkg2021, "Person")
	person2021s := astmodel.MakeInternalTypeName(test.Pkg2021s, "Person")

	address2020p := astmodel.MakeInternalTypeName(pkg2020p, "Address")
	address2020ps := astmodel.MakeInternalTypeName(pkg2020ps, "Address")

	address2021p := astmodel.MakeInternalTypeName(pkg2021p, "Address")
	address2021ps := astmodel.MakeInternalTypeName(pkg2021ps, "Address")

	omc := config.NewObjectModelConfiguration()
	builder := NewConversionGraphBuilder(omc, "v")
	builder.Add(person2020, person2020s)
	builder.Add(person2021, person2021s)
	builder.Add(address2020p, address2020ps)
	builder.Add(address2021p, address2021ps)

	graph, err := builder.Build()

	// Check size of graph
	g.Expect(err).To(Succeed())
	g.Expect(graph.TransitionCount()).To(Equal(6))

	expectedTransitions := []struct {
		from astmodel.TypeName
		to   astmodel.TypeName
	}{
		{person2020, person2020s},
		{person2021, person2021s},
		{person2020s, person2021s},

		{address2020p, address2020ps},
		{address2021p, address2021ps},
		{address2021ps, address2020ps}, // Preview versions always convert backwards
	}

	for _, expected := range expectedTransitions {
		g.Expect(graph.LookupTransition(expected.from)).To(Equal(expected.to))
	}
}
