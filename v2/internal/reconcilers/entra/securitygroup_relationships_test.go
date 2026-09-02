/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestOrderedUnique_DeduplicatesAndPreservesFirstSeenOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	actual := orderedUnique([]string{"B", "A", "B", "C", "A", "D", "D"})

	g.Expect(actual).To(Equal([]string{"B", "A", "C", "D"}))
}
