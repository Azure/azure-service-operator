/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestArrayType_Equals_WhenGivenType_ReturnsExpectedResult(t *testing.T) {
	strArray := NewArrayType(StringType)
	intArray := NewArrayType(IntType)
	otherStrArray := NewArrayType(StringType)
	otherType := StringType

	cases := []struct {
		name string
		thisType  Type
		otherType Type
		expected  bool
	}{
		// Expect equal to self
		{"Equal to self", strArray, strArray, true},
		// Expect equal to same
		{"Equal to same", strArray, otherStrArray, true},
		{"Equal to same", otherStrArray, strArray, true},
		// Expect not-equal when different element type
		{"Not equal to different", strArray, intArray, false},
		{"Not equal to different", intArray, strArray, false},
		// Expect not-equal when completely different type
		{"Not equal to different", strArray, otherType, false},
		{"Not equal to different", otherType, strArray, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := c.thisType.Equals(c.otherType)

			g.Expect(areEqual).To(Equal(c.expected))
		})
	}
}
