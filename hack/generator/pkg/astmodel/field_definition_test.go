/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_NewFieldDefinition_GivenValues_InitializesFields(t *testing.T) {
	g := NewGomegaWithT(t)

	fieldName := FieldName("FullName")
	fieldtype := StringType
	jsonName := "family-name"

	field := NewFieldDefinition(fieldName, jsonName, fieldtype)

	g.Expect(field.fieldName).To(Equal(fieldName))
	g.Expect(field.fieldType).To(Equal(fieldtype))
	g.Expect(field.jsonName).To(Equal(jsonName))
	g.Expect(field.description).To(BeEmpty())
}

func Test_FieldDefinitionWithDescription_GivenDescription_SetsField(t *testing.T) {
	g := NewGomegaWithT(t)

	description := "description"
	field := NewFieldDefinition("FullName", "fullname", StringType).WithDescription(&description)

	g.Expect(field.description).To(Equal(description))
}

func Test_FieldDefinitionWithDescription_GivenDescription_DoesNotModifyOriginal(t *testing.T) {
	g := NewGomegaWithT(t)
	description := "description"
	original := NewFieldDefinition("FullName", "fullName", StringType)

	field := original.WithDescription(&description)

	g.Expect(field.description).NotTo(Equal(original.description))
}

func Test_FieldDefinitionAsAst_GivenValidField_ReturnsNonNilResult(t *testing.T) {
	g := NewGomegaWithT(t)

	field := NewFieldDefinition("FullName", "fullName", StringType)
	node := field.AsField(nil)

	g.Expect(node).NotTo(BeNil())
}

func TestFieldDefinition_Equals_WhenGivenFieldDefinition_ReturnsExpectedResult(t *testing.T) {

	strField := createStringField("FullName", "Full Legal Name")
	otherStrField := createStringField("FullName", "Full Legal Name")

	intField := createIntField("Age", "Age at last birthday")

	differentName := createStringField("Name", "Full Legal Name")
	differentType := createIntField("FullName", "Full Legal Name")
	differentDescription := createIntField("FullName", "The whole thing")

	cases := []struct {
		name       string
		thisField  *FieldDefinition
		otherField *FieldDefinition
		expected   bool
	}{
		// Expect equal to self
		{"Equal to self", strField, strField, true},
		{"Equal to self", intField, intField, true},
		// Expect equal to same
		{"Equal to same", strField, otherStrField, true},
		{"Equal to same", otherStrField, strField, true},
		// Expect not-equal when properties are different
		{"Not-equal if names are different", strField, differentName, false},
		{"Not-equal if types are different", strField, differentType, false},
		{"Not-equal if descriptions are different", strField, differentDescription, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := c.thisField.Equals(c.otherField)

			g.Expect(areEqual).To(Equal(c.expected))
		})
	}
}
