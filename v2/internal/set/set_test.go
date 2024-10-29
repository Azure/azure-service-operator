/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package set

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestSet_WhenConstructedWithItems_HasExpectedSize(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(1, 2, 3)
	g.Expect(set).To(HaveLen(3))
}

func TestSet_WhenAddingItem_AddsToSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(1, 2, 3)
	set.Add(4)

	g.Expect(set).To(HaveLen(4))
}

func TestSet_WhenAddingAllItemsFromAnotherSet_AddsOnlyNewItems(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(1, 2, 3)
	otherSet := Make(3, 4, 5)
	set.AddAll(otherSet)
	g.Expect(set).To(HaveLen(5))
}

func TestSet_WhenRemovingItem_RemovesFromSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(1, 2, 3)
	set.Remove(2)
	g.Expect(set).To(HaveLen(2))
}

func TestSet_WhenCopying_ReturnsEqualSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(1, 2, 3)
	copySet := set.Copy()
	g.Expect(copySet).To(Equal(set))
}

func TestSet_WhenCleared_IsEmpty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(1, 2, 3)
	set.Clear()
	g.Expect(set).To(BeEmpty())
}

func TestSet_WhenCheckingEquality_ReturnsTrueForEqualSets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set1 := Make(1, 2, 3)
	set2 := Make(3, 1, 2)
	g.Expect(set1).To(Equal(set2))
}

func TestSet_WhenGettingValues_ReturnsAllValues(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make("foo", "bar", "baz")
	values := set.Values()
	g.Expect(values).To(ConsistOf("foo", "bar", "baz"))
}

func TestSet_WhenGettingSortedSlice_ReturnsSortedValues(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(3, 1, 2)
	sortedValues := AsSortedSlice(set)
	g.Expect(sortedValues).To(Equal([]int{1, 2, 3}))
}

func TestSet_Where_GivenPredicate_ReturnsSetOfMatchingItems(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(1, 2, 3, 4, 5)
	evenSet := set.Where(func(x int) bool {
		return x%2 == 0
	})
	g.Expect(evenSet).To(Equal(Make(2, 4)))
}

func TestSet_Except_GivenSetOfValues_ReturnsRemainingValues(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(1, 2, 3, 4, 5)
	remainingSet := set.Except(Make(2, 4))
	g.Expect(remainingSet).To(Equal(Make(1, 3, 5)))
}

func TestSet_Union_ReturnsUnionOfTwoSets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(1, 2, 3, 4, 5)
	set2 := Make(5, 6, 7, 8, 9, 10)

	g.Expect(Union(set, set2)).To(Equal(Make(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)))
}

func TestSet_Union_ReturnsUnionOfFourSets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	set := Make(1, 2, 3, 4, 5)
	set2 := Make(5, 6, 7, 8, 9, 10)
	set3 := Make(5, 8, 1, 11, 15, 2)
	set4 := Make(12, 2, 2, 13, 7, 14)

	g.Expect(Union(set, set2, set3, set4)).To(Equal(Make(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)))
}
