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

func TestResourceConversionGraph_WithSingleReference_HasExpectedTransition(t *testing.T) {
	/*
	 * Test that a graph that contains just one API version and the associate Storage version ends up with a single
	 * transition, between the original API version and the storage variant.
	 */
	t.Parallel()
	g := NewGomegaWithT(t)

	person2020 := astmodel.MakeTypeName(test.Pkg2020, "Person")
	person2020s := astmodel.MakeTypeName(test.Pkg2020s, "Person")

	builder := NewResourceConversionGraphBuilder("demo", "v")
	builder.Add(person2020, person2020s)
	graph, err := builder.Build()

	// Check size of graph
	g.Expect(err).To(Succeed())
	g.Expect(graph.TransitionCount()).To(Equal(1))

	// Check for single expected transition
	next := graph.LookupTransition(person2020)
	g.Expect(next).NotTo(Equal(astmodel.EmptyTypeName))
	g.Expect(astmodel.IsStoragePackageReference(next.PackageReference())).To(BeTrue())
}

func TestResourceConversionGraph_WithTwoGAReferences_HasExpectedTransitions(t *testing.T) {
	/*
	 * Test that a graph that contains two GA API releases and the matching storage versions ends up with three
	 * transitions. Each API version should have a transition to its matching storage variant, and there should be a
	 * transition from the older storage variant to the newer one.
	 */
	t.Parallel()
	g := NewGomegaWithT(t)

	person2020 := astmodel.MakeTypeName(test.Pkg2020, "Person")
	person2020s := astmodel.MakeTypeName(test.Pkg2020s, "Person")
	person2021 := astmodel.MakeTypeName(test.Pkg2021, "Person")
	person2021s := astmodel.MakeTypeName(test.Pkg2021s, "Person")

	builder := NewResourceConversionGraphBuilder("demo", "v")
	builder.Add(person2020, person2020s)
	builder.Add(person2021, person2021s)
	graph, err := builder.Build()

	// Check size of graph
	g.Expect(err).To(Succeed())
	g.Expect(graph.TransitionCount()).To(Equal(3))

	// Check for expected transition for Person2020
	after2020 := graph.LookupTransition(person2020)
	g.Expect(after2020).NotTo(Equal(astmodel.EmptyTypeName))

	// Check for expected transition for Person2021
	after2021 := graph.LookupTransition(person2021)
	g.Expect(after2021).NotTo(Equal(astmodel.EmptyTypeName))

	// Check for expected transition for the storage variant of Person2020s
	after2020s := graph.LookupTransition(after2020)
	g.Expect(after2020s).NotTo(Equal(astmodel.EmptyTypeName))
}
func TestResourceConversionGraph_WithGAAndPreviewReferences_HasExpectedTransitions(t *testing.T) {
	/*
	 * Test that a graph containing two GA and one *Preview* API release (and matching storage versions) ends up with
	 * five transitions. Each API version should have a transition to a matching storage variant, and there should be a
	 * transition from the preview storage variant *back* to the prior GA storage variant. This test only checks for
	 * cases not already covered by other tests, above.
	 */
	t.Parallel()
	g := NewGomegaWithT(t)

	person2020 := astmodel.MakeTypeName(test.Pkg2020, "Person")
	person2020s := astmodel.MakeTypeName(test.Pkg2020s, "Person")
	person2021 := astmodel.MakeTypeName(test.Pkg2021, "Person")
	person2021s := astmodel.MakeTypeName(test.Pkg2021s, "Person")
	person2021p := astmodel.MakeTypeName(test.Pkg2021Preview, "Person")
	person2021ps := astmodel.MakeTypeName(test.Pkg2021PreviewStorage, "Person")

	builder := NewResourceConversionGraphBuilder("demo", "v")
	builder.Add(person2020, person2020s)
	builder.Add(person2021, person2021s)
	builder.Add(person2021p, person2021ps)
	graph, err := builder.Build()

	// Check size of graph
	g.Expect(err).To(Succeed())
	g.Expect(graph.TransitionCount()).To(Equal(5))

	// Check for expected transition for Person2020
	after2020 := graph.LookupTransition(person2020)
	g.Expect(after2020).NotTo(Equal(astmodel.EmptyTypeName))

	// Check for expected transition for Person2021Preview
	after2021p := graph.LookupTransition(person2021p)
	g.Expect(after2021p).NotTo(Equal(astmodel.EmptyTypeName))

	// Check for expected transition for the storage variant of Person2021Preview - it goes BACK to Person2020
	ref := graph.LookupTransition(after2021p)
	g.Expect(ref).To(Equal(after2020))
}

func TestResourceConversionGraph_WithCompatibilityReferences_HasExpectedTransitions(t *testing.T) {
	/*
	 * Test that a graph containing two GA versions and one *backward compatibility* version ends up with
	 * five transitions. Each API version should have a transition to a matching storage variant, and there should be a
	 * transition from the compatibility storage variant *forward* to the earlier GA storage variant. There will also
	 * be a spurious (unused) transition from the API type underlying the compatibility storage variant.
	 * This test only checks for cases not already covered by other tests, above.
	 */
	t.Parallel()
	g := NewGomegaWithT(t)

	compatApi := test.Pkg2022.WithVersionPrefix("c")
	compatStorage := astmodel.MakeStoragePackageReference(compatApi)

	person2020 := astmodel.MakeTypeName(test.Pkg2020, "Person")
	person2020s := astmodel.MakeTypeName(test.Pkg2020s, "Person")
	person2021 := astmodel.MakeTypeName(test.Pkg2021, "Person")
	person2021s := astmodel.MakeTypeName(test.Pkg2021s, "Person")
	person2020a := astmodel.MakeTypeName(compatApi, "Person")
	person2020as := astmodel.MakeTypeName(compatStorage, "Person")

	builder := NewResourceConversionGraphBuilder("demo", "v")
	builder.Add(person2020a, person2020as)
	builder.Add(person2020, person2020s)
	builder.Add(person2021, person2021s)

	graph, err := builder.Build()

	// Check size of graph
	g.Expect(err).To(Succeed())
	g.Expect(graph.TransitionCount()).To(Equal(5))

	// Check for expected transition for compatApi
	after2020a := graph.LookupTransition(person2020a)
	g.Expect(after2020a).To(Equal(person2020as))

	// Check for expected transition for compatStorage
	after2020as := graph.LookupTransition(person2020as)
	g.Expect(after2020as).To(Equal(person2020s))
}
