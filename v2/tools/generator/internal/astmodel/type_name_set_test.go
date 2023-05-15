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
	oneTypeName = MakeTypeName(GenRuntimeReference, "One")
	twoTypeName = MakeTypeName(GenRuntimeReference, "Two")
)

/*
 * Empty Set tests
 */

func Test_TypeNameSet_WhenEmpty_HasLengthZero(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	emptySet := NewTypeNameSet()
	g.Expect(len(emptySet)).To(Equal(0))
}

func Test_TypeNameSet_RemoveFromEmptySet_SetUnchanged(t *testing.T) {
	g := NewGomegaWithT(t)
	emptySet := NewTypeNameSet()
	emptySet.Remove(oneTypeName)
	g.Expect(len(emptySet)).To(Equal(0))
}

func Test_TypeNameSet_Equals_ReturnsTrueForEmptySets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	emptySet1 := NewTypeNameSet()
	emptySet2 := NewTypeNameSet()
	g.Expect(emptySet1.Equals(emptySet2)).To(BeTrue())
}

/*
 * Add() Tests
 */

func Test_TypeNameSet_AfterAddingFirstItem_ContainsItem(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewTypeNameSet()
	set.Add(oneTypeName)
	g.Expect(set.Contains(oneTypeName)).To(BeTrue())
}

func Test_TypeNameSet_AfterAddingSecondItem_ContainsItem(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewTypeNameSet()
	set.Add(oneTypeName)
	set.Add(twoTypeName)
	g.Expect(set.Contains(twoTypeName)).To(BeTrue())
}

/*
 * Remove() Tests
 */
func Test_TypeNameSet_AfterRemovingOnlyItem_SetIsEmpty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewTypeNameSet(oneTypeName)
	set.Remove(oneTypeName)
	g.Expect(len(set)).To(Equal(0))
}

func Test_TypeNameSet_RemoveNonExistentItem_SetUnchanged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewTypeNameSet(oneTypeName)
	set.Remove(twoTypeName)
	g.Expect(set.Contains(oneTypeName)).To(BeTrue())
}

func Test_TypeNameSet_AfterRemovingItem_SetDoesNotContainItem(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewTypeNameSet(oneTypeName, twoTypeName)
	set.Remove(oneTypeName)
	g.Expect(set.Contains(oneTypeName)).To(BeFalse())
	g.Expect(set.Contains(twoTypeName)).To(BeTrue())
}

/*
 * Equal() Tests
 */
func Test_TypeNameSet_Equals_ReturnsTrueForEqualSets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	// Constructing two sets with different order to check if it internally checks for equality or not
	set1 := NewTypeNameSet(oneTypeName, twoTypeName)
	set2 := NewTypeNameSet(twoTypeName, oneTypeName)
	g.Expect(set1.Equals(set2)).To(BeTrue())
}

func Test_TypeNameSet_Equals_ReturnsFalseForSetsOfDifferentLengths(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set1 := NewTypeNameSet(oneTypeName, twoTypeName)
	set2 := NewTypeNameSet(oneTypeName)
	g.Expect(set1.Equals(set2)).To(BeFalse())
}

func Test_TypeNameSet_Equals_ReturnsFalseForSetsWithDifferentContents(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set1 := NewTypeNameSet(oneTypeName, twoTypeName)
	set2 := NewTypeNameSet(twoTypeName)
	g.Expect(set1.Equals(set2)).To(BeFalse())
}
