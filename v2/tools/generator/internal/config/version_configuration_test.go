/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestVersionConfiguration_WhenYAMLWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var versionConfig VersionConfiguration
	err := yaml.Unmarshal(yamlBytes, &versionConfig)
	g.Expect(err).To(Succeed())
	g.Expect(versionConfig.types).To(HaveLen(3))
}

func TestVersionConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var versionConfig VersionConfiguration
	err := yaml.Unmarshal(yamlBytes, &versionConfig)
	g.Expect(err).NotTo(Succeed())
}

func TestVersionConfiguration_TypeRename_WhenTypeFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person").SetTypeRename("Party")
	versionConfig := NewVersionConfiguration("2015-01-01").Add(person)

	name, ok := versionConfig.TypeRename("Person")
	g.Expect(ok).To(BeTrue())
	g.Expect(name).To(Equal("Party"))
}

func TestVersionConfiguration_TypeRename_WhenTypeNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person").SetTypeRename("Party")
	versionConfig := NewVersionConfiguration("2015-01-01").Add(person)

	name, ok := versionConfig.TypeRename("Address")
	g.Expect(ok).To(BeFalse())
	g.Expect(name).To(Equal(""))
}

func TestVersionConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	person := NewTypeConfiguration("Person").Add(spouse)
	versionConfig := NewVersionConfiguration("2015-01-01").Add(person)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, ok := versionConfig.ARMReference(typeName.Name(), "Spouse")
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeTrue())
}

func TestVersionConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	fullName := NewPropertyConfiguration("FullName").SetARMReference(false)
	person := NewTypeConfiguration("Person").Add(fullName)
	versionConfig := NewVersionConfiguration("2015-01-01").Add(person)

	isReference, ok := versionConfig.ARMReference("Person", "FullName")
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeFalse())
}

func TestVersionConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	fullName := NewPropertyConfiguration("FullName").SetARMReference(false)
	person := NewTypeConfiguration("Person").Add(fullName)
	versionConfig := NewVersionConfiguration("2015-01-01").Add(person)

	_, ok := versionConfig.ARMReference("Person", "KnownAs")
	g.Expect(ok).To(BeFalse())
}

func TestVersionConfiguration_FindUnusedARMReferences_WhenReferenceUsed_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	spouse.ARMReference()
	person := NewTypeConfiguration("Person").Add(spouse)
	versionConfig := NewVersionConfiguration("2015-01-01").Add(person)

	g.Expect(versionConfig.FindUnusedARMReferences()).To(BeEmpty())
}

func TestVersionConfiguration_FindUnusedARMReferences_WhenReferenceNotUsed_ReturnsExpectedMessage(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	person := NewTypeConfiguration("Person").Add(spouse)
	versionConfig := NewVersionConfiguration("2015-01-01").Add(person)

	unused := versionConfig.FindUnusedARMReferences()
	g.Expect(unused).To(HaveLen(1))
	g.Expect(unused[0]).To(ContainSubstring(versionConfig.name))
}
