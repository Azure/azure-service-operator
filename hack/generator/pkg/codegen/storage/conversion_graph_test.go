/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func TestConversionGraph_WithTwoUnrelatedReferences_HasExpectedTransitions(t *testing.T) {
	/*
	 *  Test that a conversion graph that contains two API package references from different groups ends up with just
	 *  one transition for each group, each linking from the provided API version to the matching storage variant.
	 */
	g := NewGomegaWithT(t)

	builder := NewConversionGraphBuilder()
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
	builder := NewConversionGraphBuilder()
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
		expectedOk   bool
	}{
		{"Hub type returns nothing", personHub, astmodel.TypeName{}, false},
		{"Directly linked api resolves", person2022.Name(), personHub, true},
		{"Directly linked storage resolves", person2021s.Name(), personHub, true},
		{"Doubly linked api resolves", person2021.Name(), personHub, true},
		{"Indirectly linked api resolves", person2020.Name(), personHub, true},
		{"Indirectly linked storage resolves", person2020s.Name(), personHub, true},
		{"Hub type returns nothing even when resource does not exist in latest package", addressHub, astmodel.TypeName{}, false},
		{"Directly linked api resolves when resource does not exist in latest package", address2021.Name(), addressHub, true},
		{"Indirectly linked api resolves when resource does not exist in latest package", address2020.Name(), addressHub, true},
	}

	t.Parallel()
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			t.Parallel()

			actual, ok := graph.FindHub(c.start, types)
			g.Expect(ok).To(Equal(c.expectedOk))
			if ok {
				g.Expect(actual).To(Equal(c.expectedName))
			}
		})
	}
}
