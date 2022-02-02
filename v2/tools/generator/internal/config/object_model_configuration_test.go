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
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var model ObjectModelConfiguration
	err := yaml.Unmarshal(yamlBytes, &model)
	g.Expect(err).To(Succeed())
	g.Expect(model.groups).To(HaveLen(1))
}

func TestObjectModelConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var model ObjectModelConfiguration
	err := yaml.Unmarshal(yamlBytes, &model)
	g.Expect(err).NotTo(Succeed())
}

func TestObjectModelConfiguration_TypeRename_WhenTypeFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person")
	person.nameInNextVersion.write("Party")

	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	name, err := modelConfig.LookupNameInNextVersion(typeName)
	g.Expect(err).To(Succeed())
	g.Expect(name).To(Equal("Party"))
}

func TestObjectModelConfiguration_TypeRename_WhenTypeNotFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person")
	person.nameInNextVersion.write("Party")

	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Address")
	name, err := modelConfig.LookupNameInNextVersion(typeName)
	g.Expect(err).NotTo(Succeed())
	g.Expect(name).To(Equal(""))
	g.Expect(err.Error()).To(ContainSubstring(typeName.Name()))
}

func TestObjectModelConfiguration_VerifyTypeRenamesConsumed_WhenRenameUsed_ReturnsEmptySlice(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person")
	person.nameInNextVersion.write("Party")

	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	_, err := modelConfig.LookupNameInNextVersion(typeName)
	g.Expect(err).To(Succeed())
	g.Expect(modelConfig.VerifyNameInNextVersionConsumed()).To(Succeed())
}

func TestObjectModelConfiguration_VerifyTypeRenamesConsumed_WhenRenameUnused_ReturnsExpectedMessage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person")
	person.nameInNextVersion.write("Party")

	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	g.Expect(modelConfig.VerifyNameInNextVersionConsumed()).NotTo(Succeed())
}

func TestObjectModelConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse")
	person := NewTypeConfiguration("Person").Add(spouse)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	spouse.armReference.write(true)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, err := modelConfig.ARMReference(typeName, "Spouse")
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeTrue())
}

func TestObjectModelConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fullName := NewPropertyConfiguration("FullName")
	person := NewTypeConfiguration("Person").Add(fullName)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	fullName.armReference.write(false)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, err := modelConfig.ARMReference(typeName, "FullName")
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeFalse())
}

func TestObjectModelConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse")
	person := NewTypeConfiguration("Person").Add(spouse)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)
	modelConfig := NewObjectModelConfiguration().Add(group)

	spouse.armReference.write(true)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	_, err := modelConfig.ARMReference(typeName, "KnownAs")
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring("KnownAs"))
}

func TestObjectModelConfiguration_VerifyARMReferencesConsumed_WhenReferenceUsed_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse")
	person := NewTypeConfiguration("Person").Add(spouse)
	version := NewVersionConfiguration("2015-01-01").Add(person)
	group := NewGroupConfiguration("microsoft.demo").Add(version)
	modelConfig := NewObjectModelConfiguration().Add(group)

	spouse.armReference.write(true)

	ref, err := spouse.ARMReference()
	g.Expect(ref).To(BeTrue())
	g.Expect(err).To(Succeed())
	g.Expect(modelConfig.VerifyARMReferencesConsumed()).To(Succeed())
}

func TestObjectModelConfiguration_VerifyARMReferencesConsumed_WhenReferenceNotUsed_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse")
	person := NewTypeConfiguration("Person").Add(spouse)
	version := NewVersionConfiguration("2015-01-01").Add(person)
	group := NewGroupConfiguration("microsoft.demo").Add(version)
	modelConfig := NewObjectModelConfiguration().Add(group)

	spouse.armReference.write(true)

	g.Expect(modelConfig.VerifyARMReferencesConsumed()).NotTo(Succeed())
}
