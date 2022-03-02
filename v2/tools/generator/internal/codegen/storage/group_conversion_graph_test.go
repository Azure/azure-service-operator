/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGroupConversionGraph_WithSingleReference_HasExpectedTransition(t *testing.T) {
	/*
	 * Test that a group that contains just one API version and the associate Storage version ends up with a single
	 * transition, between the original API version and the storage variant.
	 */
	t.Parallel()
	g := NewGomegaWithT(t)

	builder := NewGroupConversionGraphBuilder("demo")
	builder.Add(test.Pkg2020, test.Pkg2020s)
	graph, err := builder.Build()

	// Check size of graph
	g.Expect(err).To(Succeed())
	g.Expect(graph.TransitionCount()).To(Equal(1))

	// Check for single expected transition
	pkg, ok := graph.LookupTransition(test.Pkg2020)
	g.Expect(ok).To(BeTrue())
	g.Expect(pkg).NotTo(BeNil())
	g.Expect(astmodel.IsStoragePackageReference(pkg)).To(BeTrue())
}

func TestGroupConversionGraph_WithTwoGAReferences_HasExpectedTransitions(t *testing.T) {
	/*
	 * Test that a group that contains two GA API releases and the matching storage versions ends up with three
	 * transitions. Each API version should have a transition to its matching storage variant, and there should be a
	 * transition from the older storage variant to the newer one.
	 */
	t.Parallel()
	g := NewGomegaWithT(t)

	builder := NewGroupConversionGraphBuilder("demo")
	builder.Add(test.Pkg2020, test.Pkg2020s)
	builder.Add(test.Pkg2021, test.Pkg2021s)
	graph, err := builder.Build()

	// Check size of graph
	g.Expect(err).To(Succeed())
	g.Expect(graph.TransitionCount()).To(Equal(3))

	// Check for expected transition for Pkg2020
	pkg2020storage, ok := graph.LookupTransition(test.Pkg2020)
	g.Expect(ok).To(BeTrue())
	g.Expect(pkg2020storage).NotTo(BeNil())

	// Check for expected transition for Pkg2021
	pkg2021storage, ok := graph.LookupTransition(test.Pkg2021)
	g.Expect(ok).To(BeTrue())
	g.Expect(pkg2021storage).NotTo(BeNil())

	// Check for expected transition for the storage variant of Pkg2020
	ref, ok := graph.LookupTransition(pkg2020storage)
	g.Expect(ok).To(BeTrue())
	g.Expect(ref).To(Equal(pkg2021storage))
}

func TestGroupConversionGraph_WithGAAndPreviewReferences_HasExpectedTransitions(t *testing.T) {
	/*
	 * Test that a group containing two GA and one *Preview* API release (and matching storage versions) ends up with
	 * five transitions. Each API version should have a transition to a matching storage variant, and there should be a
	 * transition from the preview storage variant *back* to the prior GA storage variant. This test only checks for
	 * cases not already covered by other tests, above.
	 */
	t.Parallel()
	g := NewGomegaWithT(t)

	builder := NewGroupConversionGraphBuilder("demo")
	builder.Add(test.Pkg2020, test.Pkg2020s)
	builder.Add(test.Pkg2021Preview, test.Pkg2021PreviewStorage)
	builder.Add(test.Pkg2021, test.Pkg2021s)
	graph, err := builder.Build()

	// Check size of graph
	g.Expect(err).To(Succeed())
	g.Expect(graph.TransitionCount()).To(Equal(5))

	// Check for expected transition for Pkg2020
	pkg2020storage, ok := graph.LookupTransition(test.Pkg2020)
	g.Expect(ok).To(BeTrue())
	g.Expect(pkg2020storage).NotTo(BeNil())

	// Check for expected transition for Pkg2021Preview
	pkg2021previewStorage, ok := graph.LookupTransition(test.Pkg2021Preview)
	g.Expect(ok).To(BeTrue())
	g.Expect(pkg2021previewStorage).NotTo(BeNil())

	// Check for expected transition for the storage variant of Pkg2021Preview - it goes BACK to Pkg2020
	ref, ok := graph.LookupTransition(pkg2021previewStorage)
	g.Expect(ok).To(BeTrue())
	g.Expect(ref).To(Equal(pkg2020storage))
}
