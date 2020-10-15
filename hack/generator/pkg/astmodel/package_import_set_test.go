/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	. "github.com/onsi/gomega"
	"testing"
)

var (
	simpleTestRef PackageReference = MakeExternalPackageReference("simple")
	pathTestRef   PackageReference = MakeExternalPackageReference("package/path")

	simpleTestImport         = NewPackageImport(simpleTestRef)
	pathTestImport           = NewPackageImport(pathTestRef)
	simpleTestImportWithName = simpleTestImport.WithName("simple")
)

/*
 * NewPackageImportSet() tests
 */

func TestEmptyPackageImportSet_ReturnsEmptySet(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	g.Expect(set.imports).To(HaveLen(0))
}

/*
 * AddImport() tests
 */

func TestAddImport_WhenImportMissing_IncreasesSizeOfSet(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	g.Expect(set.imports).To(HaveLen(1))
}

func TestAddImport_WhenImportPresent_LeavesSetSameSize(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	set.AddImport(simpleTestImport)
	g.Expect(set.imports).To(HaveLen(1))
}

func TestAddImport_WhenAddingNamedImportAndUnnamedExists_PrefersNamedImport(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	set.AddImport(simpleTestImportWithName)
	imp, ok := set.ImportFor(simpleTestRef)
	g.Expect(ok).To(BeTrue())
	g.Expect(imp).To(Equal(simpleTestImportWithName))
}

func TestAddImport_WhenAddingUnnamedImportAndNamedExists_PrefersNamedImport(t *testing.T) {
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

func TestAddReference_WhenReferenceMissing_IncreasesSizeOfSet(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImportOfReference(simpleTestRef)
	g.Expect(set.imports).To(HaveLen(1))
}

func TestAddImport_WhenReferencePresent_LeavesSetSameSize(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImportOfReference(simpleTestRef)
	set.AddImportOfReference(simpleTestRef)
	g.Expect(set.imports).To(HaveLen(1))
}

/*
 * Merge() tests
 */

func TestMerge_GivenEmptySet_LeavesSetUnchanged(t *testing.T) {
	g := NewGomegaWithT(t)
	setA := NewPackageImportSet()
	setA.AddImportOfReference(simpleTestRef)
	setA.AddImportOfReference(pathTestRef)
	setB := NewPackageImportSet()
	setA.Merge(setB)
	g.Expect(setA.imports).To(HaveLen(2))
}

func TestMerge_GivenIdenticalSet_LeavesSetUnchanged(t *testing.T) {
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
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	g.Expect(set.ContainsImport(simpleTestImport)).To(BeTrue())
}

func TestContains_GivenNonMemberOfSet_ReturnsFalse(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	g.Expect(set.ContainsImport(pathTestImport)).To(BeFalse())
}

/*
 * Remove() tests
 */

func TestRemove_WhenItemInSet_RemovesIt(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	set.Remove(simpleTestImport)
	g.Expect(set.ContainsImport(simpleTestImport)).To(BeFalse())
}

func TestRemove_WhenItemNotInSet_LeavesSetWithoutIt(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImport(simpleTestImport)
	set.Remove(pathTestImport)
	g.Expect(set.ContainsImport(pathTestImport)).To(BeFalse())
}

/*
 * ByNameInGroups() tests
 */

func TestByNameInGroups_AppliesExpectedOrdering(t *testing.T) {

	fmtRef := MakeExternalPackageReference("fmt")
	testingRef := MakeExternalPackageReference("testing")
	gomegaRef := MakeExternalPackageReference("github.com/onsi/gomega")

	bareFmtImport := NewPackageImport(fmtRef)
	namedFmtImport := bareFmtImport.WithName("f")

	bareTestingImport := NewPackageImport(testingRef)
	namedTestingImport := bareTestingImport.WithName("tst")

	gomegaImport := NewPackageImport(gomegaRef)
	implicitGomegaImport := gomegaImport.WithName(".")

	localRef := MakeLocalPackageReference("this", "v1")
	localImport := NewPackageImport(localRef)

	cases := []struct {
		name  string
		left  PackageImport
		right PackageImport
		less  bool
	}{
		// Comparison with self
		{"fmt not less than self", bareFmtImport, bareFmtImport, false},
		{"testing not less than self", bareTestingImport, bareTestingImport, false},
		{"named fmt not less than self", namedTestingImport, namedTestingImport, false},
		{"implicit gomega not less than self", implicitGomegaImport, implicitGomegaImport, false},
		// named imports before others
		{"named testing before bare testing", namedTestingImport, bareTestingImport, true},
		{"bare testing after named testing", bareTestingImport, namedTestingImport, false},
		// named imports are ordered by name
		{"named fmt before named testing", namedFmtImport, namedTestingImport, true},
		{"named testing after named fmt", namedTestingImport, namedFmtImport, false},
		// local imports before external ones
		{"external gomega after local import", gomegaImport, localImport, false},
		{"local import before external gomega", localImport, gomegaImport, true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			less := ByNameInGroups(c.left, c.right)
			g.Expect(less).To(Equal(c.less))
		})
	}
}
