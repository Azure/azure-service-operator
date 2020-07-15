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
 * NewStructType() tests
 */

func TestNewStuctType_ReturnsEmptyType(t *testing.T) {
	g := NewGomegaWithT(t)

	st := NewStructType()

	g.Expect(st.properties).To(HaveLen(0))
	g.Expect(st.functions).To(HaveLen(0))
}

func TestStructType_Equals_WhenGivenType_ReturnsExpectedResult(t *testing.T) {

	fullNameField := NewPropertyDefinition("FullName", "full-name", StringType)
	familyNameField := NewPropertyDefinition("FamilyName", "family-name", StringType)
	knownAsField := NewPropertyDefinition("KnownAs", "known-as", StringType)
	genderField := NewPropertyDefinition("Gender", "gender", StringType)

	personType := NewStructType().WithProperties(fullNameField, familyNameField, knownAsField)
	otherPersonType := NewStructType().WithProperties(fullNameField, familyNameField, knownAsField)
	reorderedType := NewStructType().WithProperties(knownAsField, familyNameField, fullNameField)
	shorterType := NewStructType().WithProperties(knownAsField, fullNameField)
	longerType := NewStructType().WithProperties(fullNameField, familyNameField, knownAsField, genderField)
	mapType := NewMapType(StringType, personType)

	cases := []struct {
		name      string
		thisType  Type
		otherType Type
		expected  bool
	}{
		// Expect equal to self
		{"Equal to self", personType, personType, true},
		{"Equal to self", otherPersonType, otherPersonType, true},
		// Expect equal to same
		{"Equal to same", personType, otherPersonType, true},
		{"Equal to same", otherPersonType, personType, true},
		// Expect equal when properties are reordered
		{"Equal when properties reordered", personType, reorderedType, true},
		{"Equal when properties reordered", reorderedType, personType, true},
		// Expect not-equal when properties missing
		{"Not-equal when properties missing", personType, shorterType, false},
		{"Not-equal when properties missing", longerType, personType, false},
		// Expect not-equal when properties added
		{"Not-equal when properties added", personType, longerType, false},
		{"Not-equal when properties added", shorterType, personType, false},
		// Expect not-equal for different type
		{"Not-equal when different type", personType, mapType, false},
		{"Not-equal when different type", mapType, personType, false},
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
