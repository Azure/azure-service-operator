/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPlanRelationshipDelta_ComputesAddAndRemoveSets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	delta := planRelationshipDelta(
		[]string{"A", "B", "C"},
		[]string{"B", "C", "D"},
	)

	g.Expect(delta.ToAdd).To(Equal([]string{"D"}))
	g.Expect(delta.ToRemove).To(Equal([]string{"A"}))
}

func TestPlanRelationshipDelta_UsesInputOrderingAndDeduplicates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	delta := planRelationshipDelta(
		[]string{"A", "C", "A", "E", "F", "E"},
		[]string{"C", "B", "B", "F", "D", "D"},
	)

	g.Expect(delta.ToAdd).To(Equal([]string{"B", "D"}))
	g.Expect(delta.ToRemove).To(Equal([]string{"A", "E"}))
}

func TestPlanRelationshipDelta_NoChangesProducesEmptyDelta(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	delta := planRelationshipDelta(
		[]string{"A", "B", "C"},
		[]string{"A", "B", "C"},
	)

	g.Expect(delta.ToAdd).To(BeEmpty())
	g.Expect(delta.ToRemove).To(BeEmpty())
}

func TestOrderedUnique_DeduplicatesAndPreservesFirstSeenOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	actual := orderedUnique([]string{"B", "A", "B", "C", "A", "D", "D"})

	g.Expect(actual).To(Equal([]string{"B", "A", "C", "D"}))
}
