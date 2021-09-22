/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

/*
 * NewPackageImportSet() tests
 */

func TestNewPackageReferenceSet_ReturnsEmptySet(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet()
	g.Expect(set.references).To(HaveLen(0))
}

func TestNewPackageReferenceSet_GivenSingleReference_ReturnsPopulatedSet(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet(simpleTestRef)
	g.Expect(set.references).To(HaveLen(1))
}

func TestNewPackageReferenceSet_GivenMultipleReferences_ReturnsPopulatedSet(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet(simpleTestRef, pathTestRef)
	g.Expect(set.Contains(simpleTestRef)).To(BeTrue())
	g.Expect(set.Contains(pathTestRef)).To(BeTrue())
}

/*
 * AddReference() tests
 */

func TestPackageReferenceSet_AddReference_WhenReferenceMissing_IncreasesSizeOfSet(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet()
	set.AddReference(simpleTestRef)
	g.Expect(set.references).To(HaveLen(1))
}

func TestPackageReferenceSet_AddReference_WhenReferencePresent_LeavesSetSameSize(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet()
	set.AddReference(simpleTestRef)
	set.AddReference(simpleTestRef)
	g.Expect(set.references).To(HaveLen(1))
}

/*
 * Merge() tests
 */

func TestPackageReferenceSet_Merge_GivenEmptySet_LeavesSetUnchanged(t *testing.T) {
	g := NewGomegaWithT(t)
	setA := NewPackageReferenceSet(simpleTestRef, pathTestRef)
	setB := NewPackageReferenceSet()
	setA.Merge(setB)
	g.Expect(setA.references).To(HaveLen(2))
}

func TestPackageReferenceSet_Merge_GivenIdenticalSet_LeavesSetUnchanged(t *testing.T) {
	g := NewGomegaWithT(t)
	setA := NewPackageReferenceSet(simpleTestRef, pathTestRef)
	setB := NewPackageReferenceSet(simpleTestRef, pathTestRef)
	setA.Merge(setB)
	g.Expect(setA.references).To(HaveLen(2))
}

func TestPackageReferenceSet_Merge_GivenDisjointSets_MergesSets(t *testing.T) {
	g := NewGomegaWithT(t)
	setA := NewPackageReferenceSet(simpleTestRef)
	setB := NewPackageReferenceSet(pathTestRef)
	setA.Merge(setB)
	g.Expect(setA.references).To(HaveLen(2))
}

/*
 * Contains() tests
 */

func TestPackageReferenceSet_Contains_GivenMemberOfSet_ReturnsTrue(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet(simpleTestRef)
	g.Expect(set.Contains(simpleTestRef)).To(BeTrue())
}

func TestPackageReferenceSet_Contains_GivenNonMemberOfSet_ReturnsFalse(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet(simpleTestRef)
	g.Expect(set.Contains(pathTestRef)).To(BeFalse())
}

/*
 * Remove() tests
 */

func TestPackageReferenceSet_Remove_WhenItemInSet_RemovesIt(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet(simpleTestRef)
	set.Remove(simpleTestRef)
	g.Expect(set.Contains(simpleTestRef)).To(BeFalse())
}

func TestPackageReferenceSet_Remove_WhenItemNotInSet_LeavesSetWithoutIt(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet(simpleTestRef)
	set.Remove(pathTestRef)
	g.Expect(set.Contains(pathTestRef)).To(BeFalse())
}

/*
 * AsSlice() tests
 */

func TestPackageReferenceSet_AsSlice_WhenEmpty_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet()
	slice := set.AsSlice()
	g.Expect(slice).To(HaveLen(0))
}

func TestPackageReferenceSet_AsSlice_WhenSetPopulated_ReturnsExpectedSlice(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet(simpleTestRef, pathTestRef)
	slice := set.AsSlice()
	g.Expect(slice).To(HaveLen(2))
}

/*
 * AsSortedSlice() tests
 */

func TestPackageReferenceSet_AsSortedSlice_WhenEmpty_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet()
	slice := set.AsSortedSlice(func(left PackageReference, right PackageReference) bool {
		return left.PackageName() < right.PackageName()
	})
	g.Expect(slice).To(HaveLen(0))
}

func TestPackageReferenceSet_AsSortedSlice_WhenSetPopulated_ReturnsExpectedSlice(t *testing.T) {
	g := NewGomegaWithT(t)
	set := NewPackageReferenceSet(simpleTestRef, pathTestRef)
	slice := set.AsSortedSlice(func(left PackageReference, right PackageReference) bool {
		return left.PackageName() < right.PackageName()
	})
	g.Expect(slice).To(HaveLen(2))
}
