/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/google/uuid"
)

func TestOrderedUnique_DeduplicatesAndPreservesFirstSeenOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	a := uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	b := uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb")
	c := uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc")
	d := uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd")
	actual := orderedUnique([]uuid.UUID{b, a, b, c, a, d, d})

	g.Expect(actual).To(Equal([]uuid.UUID{b, a, c, d}))
}
