/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

func TestTypeConfiguration_WhenYAMLWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var typeConfig TypeConfiguration
	err := yaml.Unmarshal(yamlBytes, &typeConfig)
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.properties).To(HaveLen(4))
	g.Expect(*typeConfig.renamedTo).To(Equal("Demo"))
}

func TestTypeConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var typeConfig TypeConfiguration
	err := yaml.Unmarshal(yamlBytes, &typeConfig)
	g.Expect(err).NotTo(Succeed())
}

func TestTypeConfiguration_TypeRename_WhenRenameConfigured_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person").SetTypeRename("Address")

	name, ok := typeConfig.TypeRename()

	g.Expect(name).To(Equal("Address"))
	g.Expect(ok).To(BeTrue())
	g.Expect(typeConfig.usedRenamedTo).To(BeTrue())
}

func TestTypeConfiguration_TypeRename_WhenRenameNotConfigured_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person")

	name, ok := typeConfig.TypeRename()
	g.Expect(name).To(Equal(""))
	g.Expect(ok).To(BeFalse())
}

func TestTypeConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	typeConfig := NewTypeConfiguration("Person").Add(spouse)

	isReference, ok := typeConfig.ARMReference("Spouse")
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeTrue())
}

func TestTypeConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	fullName := NewPropertyConfiguration("FullName").SetARMReference(false)
	typeConfig := NewTypeConfiguration("Person").Add(fullName)

	isReference, ok := typeConfig.ARMReference("FullName")
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeFalse())
}

func TestTypeConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person")

	_, ok := typeConfig.ARMReference("KnownAs")
	g.Expect(ok).To(BeFalse())
}

func TestTypeConfiguration_FindUnusedARMReferences_WhenReferenceUsed_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	spouse.ARMReference()
	typeConfig := NewTypeConfiguration("Person").Add(spouse)

	g.Expect(typeConfig.FindUnusedARMReferences()).To(BeEmpty())
}

func TestTypeConfiguration_FindUnusedARMReferences_WhenReferenceNotUsed_ReturnsExpectedMessage(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	typeConfig := NewTypeConfiguration("Person").Add(spouse)

	unused := typeConfig.FindUnusedARMReferences()
	g.Expect(unused).To(HaveLen(1))
	g.Expect(unused[0]).To(ContainSubstring(typeConfig.name))
}
