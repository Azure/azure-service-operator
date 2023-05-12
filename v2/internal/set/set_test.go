/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package set

import (
	. "github.com/onsi/gomega"
	"testing"
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
