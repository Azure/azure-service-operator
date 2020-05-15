/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPrimitiveType_Equals_WhenGivenType_ReturnsExpectedResult(t *testing.T) {
	otherType := NewArrayType(IntType)

	cases := []struct {
		name string
		thisType  Type
		otherType Type
		expected  bool
	}{
		// Expect equal to self
		{"Equal to self", StringType, StringType, true},
		{"Equal to self", IntType, IntType, true},
		// Expect not-equal when different
		{"Not equal to different", StringType, IntType, false},
		{"Not equal to different", IntType, BoolType, false},
		// Expect not-equal when different type
		{"Not equal to different", IntType, otherType, false},
		{"Not equal to different", otherType, BoolType, false},
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
