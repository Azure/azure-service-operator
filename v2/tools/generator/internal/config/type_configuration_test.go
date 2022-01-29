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
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var typeConfig TypeConfiguration
	err := yaml.Unmarshal(yamlBytes, &typeConfig)
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.properties).To(HaveLen(4))
	g.Expect(*typeConfig.renamedTo).To(Equal("Demo"))
}

func TestTypeConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var typeConfig TypeConfiguration
	err := yaml.Unmarshal(yamlBytes, &typeConfig)
	g.Expect(err).NotTo(Succeed())
}

func TestTypeConfiguration_TypeRename_WhenRenameConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person").SetTypeRename("Address")

	name, err := typeConfig.TypeRename()

	g.Expect(name).To(Equal("Address"))
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.usedRenamedTo).To(BeTrue())
}

func TestTypeConfiguration_TypeRename_WhenRenameNotConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person")

	name, err := typeConfig.TypeRename()
	g.Expect(name).To(Equal(""))
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring(typeConfig.name))
}

func TestTypeConfiguration_FindUnusedTypeRenames_WhenRenameUsed_ReturnsEmptySlice(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeConfig := NewTypeConfiguration("Person").SetTypeRename("Party")
	_, err := typeConfig.TypeRename()
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.FindUnusedTypeRenames()).To(BeEmpty())
}

func TestTypeConfiguration_FindUnusedTypeRenames_WhenRenameUnused_ReturnsExpectedMessage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeConfig := NewTypeConfiguration("Person").SetTypeRename("Party")

	unused := typeConfig.FindUnusedTypeRenames()
	g.Expect(unused).To(HaveLen(1))
	g.Expect(unused[0]).To(ContainSubstring(typeConfig.name))
}

func TestTypeConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	typeConfig := NewTypeConfiguration("Person").Add(spouse)

	isReference, err := typeConfig.ARMReference("Spouse")
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeTrue())
}

func TestTypeConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	fullName := NewPropertyConfiguration("FullName").SetARMReference(false)
	typeConfig := NewTypeConfiguration("Person").Add(fullName)

	isReference, err := typeConfig.ARMReference("FullName")
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeFalse())
}

func TestTypeConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person")

	_, err := typeConfig.ARMReference("KnownAs")
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring("KnownAs"))
}

func TestTypeConfiguration_FindUnusedARMReferences_WhenReferenceUsed_ReturnsEmptySlice(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	typeConfig := NewTypeConfiguration("Person").Add(spouse)

	ref, err := spouse.ARMReference()
	g.Expect(ref).To(BeTrue())
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.FindUnusedARMReferences()).To(BeEmpty())
}

func TestTypeConfiguration_FindUnusedARMReferences_WhenReferenceNotUsed_ReturnsExpectedMessage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	typeConfig := NewTypeConfiguration("Person").Add(spouse)

	unused := typeConfig.FindUnusedARMReferences()
	g.Expect(unused).To(HaveLen(1))
	g.Expect(unused[0]).To(ContainSubstring(typeConfig.name))
}
