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
	simpleTestRef PackageReference = MakeExternalPackageReference("simple")
	pathTestRef   PackageReference = MakeExternalPackageReference("package/path")

	// Important for these two package references to have the same version so that they conflict
	emailTestRef   PackageReference = MakeExternalPackageReference("microsoft.email/v20180801")
	networkTestRef PackageReference = MakeExternalPackageReference("microsoft.network/v20180801")

	// Important for these two package references to have the same version as each other,
	// AND each service name must conflict with the references above
	emailTestAltRef   PackageReference = MakeExternalPackageReference("microsoft.email/v20200801")
	networkTestAltRef PackageReference = MakeExternalPackageReference("microsoft.network/v20200801")

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

func TestAddImportOfReference_WhenReferenceMissing_IncreasesSizeOfSet(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImportOfReference(simpleTestRef)
	g.Expect(set.imports).To(HaveLen(1))
}

func TestAddImportOfReference_WhenReferencePresent_LeavesSetSameSize(t *testing.T) {
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
	g := NewGomegaWithT(t)
	set := NewPackageImportSet()
	set.AddImportsOfReferences(simpleTestRef, pathTestRef)
	g.Expect(set.imports).To(HaveLen(2))
}

func TestAddImportsOfReferences_WhenAddingReferencesAlreadyPresent_LeavesSetSameSize(t *testing.T) {
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

	localRef := makeTestLocalPackageReference("this", "v1")
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

/*
 * Resolve Conflict Tests
 */

func TestPackageImportSet_ResolveConflicts_GivenExplicitlyNamedConflicts_ReturnsErrors(t *testing.T) {
	g := NewGomegaWithT(t)
	importA := NewPackageImport(emailTestRef).WithName("collide")
	importB := NewPackageImport(networkTestRef).WithName("collide")

	set := NewPackageImportSet()
	set.AddImport(importA)
	set.AddImport(importB)

	err := set.ResolveConflicts()
	g.Expect(err).NotTo(BeNil())
}

func TestPackageImportSet_ResolveConflicts_GivenImplicityNamedConflicts_AssignsExpectedNames(t *testing.T) {

	createSet := func(refs ...PackageReference) *PackageImportSet {
		result := NewPackageImportSet()
		for _, ref := range refs {
			result.AddImportOfReference(ref)
		}

		return result
	}

	cases := []struct {
		name         string
		set          *PackageImportSet
		testRef      PackageReference
		expectedName string
	}{
		{
			"Import conflicts with simple resolution (i)",
			createSet(emailTestRef, networkTestRef),
			emailTestRef,
			"email",
		},
		{
			"Import conflicts with simple resolution (ii)",
			createSet(emailTestRef, networkTestRef),
			networkTestRef,
			"network",
		},
		{
			"Import conflicts with versioned resolution (i)",
			createSet(emailTestRef, networkTestRef, emailTestAltRef, networkTestAltRef),
			emailTestRef,
			"emailv20180801",
		},
		{
			"Import conflicts with versioned resolution (ii)",
			createSet(emailTestRef, networkTestRef, emailTestAltRef, networkTestAltRef),
			networkTestRef,
			"networkv20180801"},
		{
			"Import conflicts with versioned resolution (iii)",
			createSet(emailTestRef, networkTestRef, emailTestAltRef, networkTestAltRef),
			emailTestAltRef,
			"emailv20200801",
		},
		{
			"Import conflicts with versioned resolution (iv)",
			createSet(emailTestRef, networkTestRef, emailTestAltRef, networkTestAltRef),
			networkTestAltRef,
			"networkv20200801",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			err := c.set.ResolveConflicts()
			g.Expect(err).To(BeNil())

			imp, ok := c.set.ImportFor(c.testRef)
			g.Expect(ok).To(BeTrue())
			g.Expect(imp.name).To(Equal(c.expectedName))
		})
	}
}

/*
 * orderImports() tests
 */

func Test_PackageSet_OrderImports(t *testing.T) {

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
