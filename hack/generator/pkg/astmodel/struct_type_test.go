/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestStructType_Equals_WhenGivenType_ReturnsExpectedResult(t *testing.T) {

	fullNameField := NewFieldDefinition("FullName", "full-name", StringType)
	familyNameField := NewFieldDefinition("FamilyName", "family-name", StringType)
	knownAsField := NewFieldDefinition("KnownAs", "known-as", StringType)
	genderField := NewFieldDefinition("Gender", "gender", StringType)

	personType := NewStructType().WithFields(fullNameField, familyNameField, knownAsField)
	otherPersonType := NewStructType().WithFields(fullNameField, familyNameField, knownAsField)
	reorderedType := NewStructType().WithFields(knownAsField, familyNameField, fullNameField)
	shorterType := NewStructType().WithFields(knownAsField, fullNameField)
	longerType := NewStructType().WithFields(fullNameField, familyNameField, knownAsField, genderField)
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
		// Expect equal when fields are reordered
		{"Equal when fields reordered", personType, reorderedType, true},
		{"Equal when fields reordered", reorderedType, personType, true},
		// Expect not-equal when fields missing
		{"Not-equal when fields missing", personType, shorterType, false},
		{"Not-equal when fields missing", longerType, personType, false},
		// Expect not-equal when fields added
		{"Not-equal when fields added", personType, longerType, false},
		{"Not-equal when fields added", shorterType, personType, false},
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
