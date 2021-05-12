/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestMapType_Equals_WhenGivenType_ReturnsExpectedResult(t *testing.T) {

	strToStr := NewMapType(StringType, StringType)
	strToBool := NewMapType(StringType, BoolType)
	intToBool := NewMapType(IntType, BoolType)
	otherStrToStr := NewMapType(StringType, StringType)
	otherType := NewArrayType(StringType)

	cases := []struct {
		name      string
		thisType  Type
		otherType Type
		expected  bool
	}{
		// Expect equal to self
		{"Equal to self", strToStr, strToStr, true},
		{"Equal to self", strToBool, strToBool, true},
		// Expect equal to same
		{"Equal to same", strToStr, otherStrToStr, true},
		{"Equal to same", otherStrToStr, strToStr, true},
		// Expect not-equal when different key types
		{"Not-equal if keys are different", intToBool, strToBool, false},
		{"Not-equal if keys are different", strToBool, intToBool, false},
		// Expect not-equal when different value types
		{"Not-equal if values are different", strToStr, strToBool, false},
		{"Not-equal if values are different", strToBool, strToStr, false},
		// Expect not-equal when completely different type
		{"Not-equal if values are different", strToStr, otherType, false},
		{"Not-equal if values are different", otherType, strToStr, false},
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
