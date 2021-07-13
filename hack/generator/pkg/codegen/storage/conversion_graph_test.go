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

	builder := NewConversionGraphBuilder()
	builder.Add(test.Pkg2020)
	builder.Add(test.Pkg2021)
	builder.Add(test.Pkg2022)

	graph, err := builder.Build()
	g.Expect(err).To(Succeed())

	pkg2020storage, ok := graph.LookupTransition(test.Pkg2020)
	g.Expect(ok).To(BeTrue())

	pkg2021storage, ok := graph.LookupTransition(test.Pkg2021)
	g.Expect(ok).To(BeTrue())

	pkg2022storage, ok := graph.LookupTransition(test.Pkg2022)
	g.Expect(ok).To(BeTrue())

	cases := []struct {
		name     string
		start astmodel.PackageReference
		expected astmodel.PackageReference
	}{
		{"Hub type resolves to self", pkg2022storage, pkg2022storage},
		{"Directly linked api resolves", test.Pkg2021, pkg2022storage},
		{"Directly linked storage resolves", pkg2021storage, pkg2022storage},
		{"Doubly linked api resolves", test.Pkg2021, pkg2022storage},
		{"Indirectly linked api resolves", test.Pkg2020, pkg2022storage},
		{"Indirectly linked storage resolves", pkg2020storage, pkg2022storage},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			startType := astmodel.MakeTypeName(c.start, "Person")
			expectedType := astmodel.MakeTypeName(c.expected, "Person")

			actual := graph.FindHubTypeName(startType)
			g.Expect(actual).To(Equal(expectedType))
		})
	}
}
