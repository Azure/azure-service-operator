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

	fieldName := "FullName"
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
	node := field.AsAst()

	g.Expect(node).NotTo(BeNil())
}
