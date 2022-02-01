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

func TestObjectModelConfiguration_WhenYAMLWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var model ObjectModelConfiguration
	err := yaml.Unmarshal(yamlBytes, &model)
	g.Expect(err).To(Succeed())
	g.Expect(model.groups).To(HaveLen(1))
}

func TestObjectModelConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var model ObjectModelConfiguration
	err := yaml.Unmarshal(yamlBytes, &model)
	g.Expect(err).NotTo(Succeed())
}

func TestObjectModelConfiguration_TypeRename_WhenTypeFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person").SetTypeRename("Party")
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	name, err := modelConfig.TypeRename(typeName)
	g.Expect(err).To(Succeed())
	g.Expect(name).To(Equal("Party"))
}

func TestObjectModelConfiguration_TypeRename_WhenTypeNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person").SetTypeRename("Party")
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Address")
	name, err := modelConfig.TypeRename(typeName)
	g.Expect(err).NotTo(Succeed())
	g.Expect(name).To(Equal(""))
	g.Expect(err.Error()).To(ContainSubstring(typeName.Name()))
}

func TestObjectModelConfiguration_FindUnusedTypeRenames_WhenRenameUsed_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person").SetTypeRename("Party")
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	_, err := modelConfig.TypeRename(typeName)
	g.Expect(err).To(Succeed())
	g.Expect(modelConfig.VerifyTypeRenamesConsumed()).To(Succeed())
}

func TestObjectModelConfiguration_FindUnusedTypeRenames_WhenRenameUnused_ReturnsExpectedMessage(t *testing.T) {
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person").SetTypeRename("Party")
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	g.Expect(modelConfig.VerifyTypeRenamesConsumed()).NotTo(Succeed())
}

func TestObjectModelConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	person := NewTypeConfiguration("Person").Add(spouse)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, err := modelConfig.ARMReference(typeName, "Spouse")
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeTrue())
}

func TestObjectModelConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	fullName := NewPropertyConfiguration("FullName").SetARMReference(false)
	person := NewTypeConfiguration("Person").Add(fullName)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, err := modelConfig.ARMReference(typeName, "FullName")
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeFalse())
}

func TestObjectModelConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	person := NewTypeConfiguration("Person").Add(spouse)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	_, err := modelConfig.ARMReference(typeName, "KnownAs")
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring("KnownAs"))
}

func TestObjectModelConfiguration_FindUnusedARMReferences_WhenReferenceUsed_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	person := NewTypeConfiguration("Person").Add(spouse)
	version := NewVersionConfiguration("2015-01-01").Add(person)
	group := NewGroupConfiguration("microsoft.demo").Add(version)
	modelConfig := NewObjectModelConfiguration().Add(group)

	ref, err := spouse.ARMReference()
	g.Expect(ref).To(BeTrue())
	g.Expect(err).To(Succeed())
	g.Expect(modelConfig.FindUnusedARMReferences()).To(BeEmpty())
}

func TestObjectModelConfiguration_FindUnusedARMReferences_WhenReferenceNotUsed_ReturnsExpectedMessage(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	person := NewTypeConfiguration("Person").Add(spouse)
	version := NewVersionConfiguration("2015-01-01").Add(person)
	group := NewGroupConfiguration("microsoft.demo").Add(version)
	modelConfig := NewObjectModelConfiguration().Add(group)

	unused := modelConfig.FindUnusedARMReferences()
	g.Expect(unused).To(HaveLen(1))
}
