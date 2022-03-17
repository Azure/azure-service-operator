/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

var (
	simpleTestRef            PackageReference = MakeExternalPackageReference("simple")
	pathTestRef              PackageReference = MakeExternalPackageReference("package/path")
	emailTestRef             PackageReference = MakeExternalPackageReference("email/v20180801")
	simpleTestImport                          = NewPackageImport(simpleTestRef)
	pathTestImport                            = NewPackageImport(pathTestRef)
	simpleTestImportWithName                  = simpleTestImport.WithName("simple")
)

/*
 * NewPackageImportSet() tests
 */

func TestEmptyPackageImportSet_ReturnsEmptySet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	g.Expect(set.imports).To(HaveLen(0))
}

/*
 * AddImport() tests
 */

func TestAddImport_WhenImportMissing_IncreasesSizeOfSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	g.Expect(set.imports).To(HaveLen(1))
}

func TestAddImport_WhenImportPresent_LeavesSetSameSize(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	set.AddImport(simpleTestImport)
	g.Expect(set.imports).To(HaveLen(1))
}

func TestAddImport_WhenAddingNamedImportAndUnnamedExists_PrefersNamedImport(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	set.AddImport(simpleTestImportWithName)
	imp, ok := set.ImportFor(simpleTestRef)
	g.Expect(ok).To(BeTrue())
	g.Expect(imp).To(Equal(simpleTestImportWithName))
}

func TestAddImport_WhenAddingUnnamedImportAndNamedExists_PrefersNamedImport(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImportWithName)
	set.AddImport(simpleTestImport)
	imp, ok := set.ImportFor(simpleTestRef)
	g.Expect(ok).To(BeTrue())
	g.Expect(imp).To(Equal(simpleTestImportWithName))
}

/*
 * AddImportOfReference() tests
 */

func TestAddImportOfReference_WhenReferenceMissing_IncreasesSizeOfSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImportOfReference(simpleTestRef)
	g.Expect(set.imports).To(HaveLen(1))
}

func TestAddImportOfReference_WhenReferencePresent_LeavesSetSameSize(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImportOfReference(simpleTestRef)
	set.AddImportOfReference(simpleTestRef)
	g.Expect(set.imports).To(HaveLen(1))
}

/*
 * AddImportsOfReferences() Tests
 */

func TestAddImportsOfReferences_WhenAddingMultipleReferences_AddsExpectedReferences(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImportsOfReferences(simpleTestRef, pathTestRef)
	g.Expect(set.imports).To(HaveLen(2))
}

func TestAddImportsOfReferences_WhenAddingReferencesAlreadyPresent_LeavesSetSameSize(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImportOfReference(simpleTestRef)
	set.AddImportOfReference(pathTestRef)
	size := len(set.imports)
	set.AddImportsOfReferences(simpleTestRef, pathTestRef)
	g.Expect(set.imports).To(HaveLen(size))
}

/*
 * Merge() tests
 */

func TestMerge_GivenEmptySet_LeavesSetUnchanged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	setA := NewPackageImportSet()
	setA.AddImportOfReference(simpleTestRef)
	setA.AddImportOfReference(pathTestRef)
	setB := NewPackageImportSet()
	setA.Merge(setB)
	g.Expect(setA.imports).To(HaveLen(2))
}

func TestMerge_GivenIdenticalSet_LeavesSetUnchanged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	setA := NewPackageImportSet()
	setA.AddImportOfReference(simpleTestRef)
	setA.AddImportOfReference(pathTestRef)
	setB := NewPackageImportSet()
	setB.AddImportOfReference(simpleTestRef)
	setB.AddImportOfReference(pathTestRef)
	setA.Merge(setB)
	g.Expect(setA.imports).To(HaveLen(2))
}

func TestMerge_GivenDisjointSets_MergesSets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	setA := NewPackageImportSet()
	setA.AddImportOfReference(simpleTestRef)
	setB := NewPackageImportSet()
	setB.AddImportOfReference(pathTestRef)
	setA.Merge(setB)
	g.Expect(setA.imports).To(HaveLen(2))
}

/*
 * Contains() tests
 */

func TestContains_GivenMemberOfSet_ReturnsTrue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	g.Expect(set.ContainsImport(simpleTestImport)).To(BeTrue())
}

func TestContains_GivenNonMemberOfSet_ReturnsFalse(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	g.Expect(set.ContainsImport(pathTestImport)).To(BeFalse())
}

/*
 * Remove() tests
 */

func TestRemove_WhenItemInSet_RemovesIt(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	set.Remove(simpleTestImport)
	g.Expect(set.ContainsImport(simpleTestImport)).To(BeFalse())
}

func TestRemove_WhenItemNotInSet_LeavesSetWithoutIt(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	set.Remove(pathTestImport)
	g.Expect(set.ContainsImport(pathTestImport)).To(BeFalse())
}

/*
 * orderImports() tests
 */

func Test_PackageSet_OrderImports(t *testing.T) {
	t.Parallel()

	alphaRef := MakeExternalPackageReference("alpha")
	betaRef := MakeExternalPackageReference("beta")

	alphaImport := NewPackageImport(alphaRef)
	betaImport := NewPackageImport(betaRef)

	alphaImportWithName := alphaImport.WithName("alpha")
	betaImportWithName := betaImport.WithName("beta")

	cases := []struct {
		name  string
		left  PackageImport
		right PackageImport
		less  bool
	}{
		{"Anonymous imports are alphabetical (i)", alphaImport, betaImport, true},
		{"Anonymous imports are alphabetical (ii)", betaImport, alphaImport, false},
		{"Named imports are alphabetical (i)", alphaImportWithName, betaImportWithName, true},
		{"Named imports are alphabetical (ii)", betaImportWithName, alphaImportWithName, false},
		{"Named imports come after anonymous of the same name (i)", alphaImportWithName, alphaImport, false},
		{"Named imports come after anonymous of the same name (ii)", betaImportWithName, betaImport, false},
		{"Named imports are alphabetical with anonymous imports (i)", betaImportWithName, alphaImport, false},
		{"Named imports are alphabetical with anonymous imports (ii)", alphaImportWithName, betaImport, true},
		{"Anonymous imports come before named import of same package (i)", alphaImport, alphaImportWithName, true},
		{"Anonymous imports come before named import of same package (ii)", betaImport, betaImportWithName, true},
		{"Anonymous imports are alphabetical with named imports (i)", alphaImport, betaImportWithName, true},
		{"Anonymous imports are alphabetical with named imports (i)", betaImport, alphaImportWithName, false},
	}

	var set PackageImportSet

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			less := set.orderImports(c.left, c.right)
			g.Expect(less).To(Equal(c.less))
		})
	}
}

/*
 * Import name assignment tests
 */

func TestPackageImportSet_GivenSet_AssignsExpectedAliases(t *testing.T) {

	batch_v2020 := makeTestLocalPackageReference("batch", "2020-01-01")
	batch_v2021 := makeTestLocalPackageReference("batch", "2021-01-01")
	batch_v2022 := makeTestLocalPackageReference("batch", "2022-01-01")

	compute_v2020 := makeTestLocalPackageReference("compute", "2020-01-01")
	compute_v2021 := makeTestLocalPackageReference("compute", "2021-01-01")
	compute_v2022 := makeTestLocalPackageReference("compute", "2022-01-01")

	network_v2020 := makeTestLocalPackageReference("network", "2020-01-01")
	network_v2021 := makeTestLocalPackageReference("network", "2021-01-01")
	network_v2022 := makeTestLocalPackageReference("network", "2022-01-01")

	cases := []struct {
		name       string
		references map[PackageReference]string
	}{
		{
			"Single version of Batch and nothing else",
			map[PackageReference]string{
				batch_v2020: "v20200101",
			},
		},
		{
			"Single version of Compute and nothing else",
			map[PackageReference]string{
				compute_v2022: "v20220101",
			},
		},
		{
			"Multiple versions of Batch",
			map[PackageReference]string{
				batch_v2020: "v20200101",
				batch_v2021: "v20210101",
				batch_v2022: "v20220101",
			},
		},
		{
			"One import from each group",
			map[PackageReference]string{
				batch_v2020:   "batch",
				compute_v2020: "compute",
				network_v2020: "network",
			},
		},
		{
			"Groups of different sizes",
			map[PackageReference]string{
				batch_v2020:   "batch",
				compute_v2020: "compute_v20200101",
				compute_v2021: "compute_v20210101",
				network_v2020: "network_v20200101",
				network_v2021: "network_v20210101",
				network_v2022: "network_v20220101",
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			set := NewPackageImportSet()
			for ref := range c.references {
				set.AddImportOfReference(ref)
			}

			for _, imp := range set.AsSlice() {
				g.Expect(imp.name).To(Equal(c.references[imp.packageReference]))
			}
		})
	}
}
