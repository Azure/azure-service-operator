/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_NewStructDefinition_GivenValues_InitializesProperties(t *testing.T) {
	g := NewGomegaWithT(t)

	const name = "demo"
	const group = "group"
	const version = "2020-01-01"
	fullNameProperty := createStringProperty("fullName", "Full legal name")
	familyNameProperty := createStringProperty("familyName", "Shared family name")
	knownAsProperty := createStringProperty("knownAs", "Commonly known as")

	ref := NewTypeName(*NewLocalPackageReference(group, version), name)
	definition := NewStructDefinition(ref, NewStructType().WithProperties(fullNameProperty, familyNameProperty, knownAsProperty))

	definitionGroup, definitionPackage, err := definition.Name().PackageReference.GroupAndPackage()
	g.Expect(err).To(BeNil())

	g.Expect(definition.Name().name).To(Equal(name))
	g.Expect(definitionGroup).To(Equal(group))
	g.Expect(definitionPackage).To(Equal(version))
	g.Expect(definition.structType.properties).To(HaveLen(3))
}

func Test_StructDefinitionAsAst_GivenValidStruct_ReturnsNonNilResult(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := NewTypeName(*NewLocalPackageReference("group", "2020-01-01"), "name")
	definition := NewStructDefinition(ref, NewStructType())
	node := definition.AsDeclarations(nil)

	g.Expect(node).NotTo(BeNil())
}

func createStringProperty(name string, description string) *PropertyDefinition {
	return NewPropertyDefinition(PropertyName(name), name, StringType).WithDescription(&description)
}

func createIntProperty(name string, description string) *PropertyDefinition {
	return NewPropertyDefinition(PropertyName(name), name, IntType).WithDescription(&description)
}
