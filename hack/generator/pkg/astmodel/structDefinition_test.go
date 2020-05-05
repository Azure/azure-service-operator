/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_NewStructDefinition_GivenValues_InitializesFields(t *testing.T) {
	g := NewGomegaWithT(t)

	const name = "demo"
	const group = "group"
	const version = "2020-01-01"
	fullNameField := createStringField("fullName", "Full legal name")
	familyNameField := createStringField("familiyName", "Shared family name")
	knownAsField := createStringField("knownAs", "Commonly known as")

	ref := NewStructReference(name, group, version)
	definition := NewStructDefinition(ref, fullNameField, familyNameField, knownAsField)

	g.Expect(definition.name).To(Equal(name))
	g.Expect(definition.groupName).To(Equal(group))
	g.Expect(definition.packageName).To(Equal(version))
	g.Expect(definition.fields).To(HaveLen(3))
}

func Test_StructDefinitionAsAst_GivenValidStruct_ReturnsNonNilResult(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := NewStructReference("name", "group", "2020-01-01")
	field := NewStructDefinition(ref)
	node := field.AsDeclaration()

	g.Expect(node).NotTo(BeNil())
}

func createStringField(name string, description string) *FieldDefinition {
	return NewFieldDefinition(name, name, StringType).WithDescription(&description)
}
