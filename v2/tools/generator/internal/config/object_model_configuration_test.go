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

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc, tc := CreateTestObjectModelConfiguration(name)
	tc.nameInNextVersion.write("Party")

	nextName, err := omc.LookupNameInNextVersion(name)
	g.Expect(err).To(Succeed())
	g.Expect(nextName).To(Equal("Party"))
}

func TestObjectModelConfiguration_TypeRename_WhenTypeNotFound_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Address")
	omc, tc := CreateTestObjectModelConfiguration(name)
	tc.nameInNextVersion.write("Party")

	otherName := astmodel.MakeTypeName(test.Pkg2020, "Location")
	nextName, err := omc.LookupNameInNextVersion(otherName)

	g.Expect(err).NotTo(Succeed())
	g.Expect(nextName).To(Equal(""))
	g.Expect(err.Error()).To(ContainSubstring(name.Name()))
}

func TestObjectModelConfiguration_VerifyTypeRenamesConsumed_WhenRenameUsed_ReturnsEmptySlice(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc, tc := CreateTestObjectModelConfiguration(name)
	tc.nameInNextVersion.write("Party")

	_, err := omc.LookupNameInNextVersion(name)
	g.Expect(err).To(Succeed())
	g.Expect(omc.VerifyNameInNextVersionConsumed()).To(Succeed())
}

func TestObjectModelConfiguration_VerifyTypeRenamesConsumed_WhenRenameUnused_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc, tc := CreateTestObjectModelConfiguration(name)
	tc.nameInNextVersion.write("Party")

	g.Expect(omc.VerifyNameInNextVersionConsumed()).NotTo(Succeed())
}

func TestObjectModelConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse")
	spouse.armReference.write(true)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc, tc := CreateTestObjectModelConfiguration(name)
	tc.add(spouse)

	isReference, err := omc.ARMReference(name, "Spouse")
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeTrue())
}

func TestObjectModelConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fullName := NewPropertyConfiguration("FullName")
	fullName.armReference.write(false)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc, tc := CreateTestObjectModelConfiguration(name)
	tc.add(fullName)

	isReference, err := omc.ARMReference(name, "FullName")
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeFalse())
}

func TestObjectModelConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse")
	spouse.armReference.write(true)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc, tc := CreateTestObjectModelConfiguration(name)
	tc.add(spouse)

	_, err := omc.ARMReference(name, "KnownAs")
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring("KnownAs"))
}

func TestObjectModelConfiguration_VerifyARMReferencesConsumed_WhenReferenceUsed_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse")
	spouse.armReference.write(true)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc, tc := CreateTestObjectModelConfiguration(name)
	tc.add(spouse)

	ref, err := omc.ARMReference(name, "Spouse")
	g.Expect(ref).To(BeTrue())
	g.Expect(err).To(Succeed())
	g.Expect(omc.VerifyARMReferencesConsumed()).To(Succeed())
}

func TestObjectModelConfiguration_VerifyARMReferencesConsumed_WhenReferenceNotUsed_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse")
	spouse.armReference.write(true)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc, tc := CreateTestObjectModelConfiguration(name)
	tc.add(spouse)

	g.Expect(omc.VerifyARMReferencesConsumed()).NotTo(Succeed())
}

func TestObjectModelConfiguration_ModifyGroup_WhenGroupDoesNotExist_CallsActionWithNewGroup(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := NewObjectModelConfiguration()
	var cfg *GroupConfiguration

	g.Expect(
		omc.ModifyGroup(
			test.Pkg2020,
			func(configuration *GroupConfiguration) error {
				cfg = configuration
				return nil
			})).To(Succeed())

	g.Expect(cfg).NotTo(BeNil())
}

func TestObjectModelConfiguration_ModifyGroup_WhenGroupExists_CallsActionWithExistingGroup(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := NewObjectModelConfiguration()
	var first *GroupConfiguration
	var second *GroupConfiguration

	g.Expect(
		omc.ModifyGroup(
			test.Pkg2020,
			func(configuration *GroupConfiguration) error {
				first = configuration
				return nil
			})).To(Succeed())

	g.Expect(
		omc.ModifyGroup(
			test.Pkg2020,
			func(configuration *GroupConfiguration) error {
				second = configuration
				return nil
			})).To(Succeed())

	g.Expect(first).To(Equal(second))
}

func TestObjectModelConfiguration_ModifyVersion_WhenVersionDoesNotExist_CallsActionWithNewVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := NewObjectModelConfiguration()
	var cfg *VersionConfiguration

	g.Expect(
		omc.ModifyVersion(
			test.Pkg2020,
			func(configuration *VersionConfiguration) error {
				cfg = configuration
				return nil
			})).To(Succeed())

	g.Expect(cfg).NotTo(BeNil())
}

func TestObjectModelConfiguration_ModifyVersion_WhenVersionExists_CallsActionWithExistingVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := NewObjectModelConfiguration()
	var first *VersionConfiguration
	var second *VersionConfiguration

	g.Expect(
		omc.ModifyVersion(
			test.Pkg2020,
			func(configuration *VersionConfiguration) error {
				first = configuration
				return nil
			})).To(Succeed())

	g.Expect(
		omc.ModifyVersion(
			test.Pkg2020,
			func(configuration *VersionConfiguration) error {
				second = configuration
				return nil
			})).To(Succeed())

	g.Expect(first).To(Equal(second))
}

func TestObjectModelConfiguration_ModifyType_WhenTypeDoesNotExist_CallsActionWithNewType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := NewObjectModelConfiguration()
	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	var cfg *TypeConfiguration

	g.Expect(
		omc.ModifyType(
			typeName,
			func(configuration *TypeConfiguration) error {
				cfg = configuration
				return nil
			})).To(Succeed())

	g.Expect(cfg).NotTo(BeNil())
}

func TestObjectModelConfiguration_ModifyType_WhenTypeExists_CallsActionWithExistingType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := NewObjectModelConfiguration()
	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	var first *TypeConfiguration
	var second *TypeConfiguration

	g.Expect(
		omc.ModifyType(
			typeName,
			func(configuration *TypeConfiguration) error {
				first = configuration
				return nil
			})).To(Succeed())

	g.Expect(
		omc.ModifyType(
			typeName,
			func(configuration *TypeConfiguration) error {
				second = configuration
				return nil
			})).To(Succeed())

	g.Expect(first).To(Equal(second))
}

func TestObjectModelConfiguration_ModifyProperty_WhenPropertyDoesNotExist_CallsActionWithNewProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := NewObjectModelConfiguration()
	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	var cfg *PropertyConfiguration

	g.Expect(
		omc.ModifyProperty(
			typeName,
			"FullName",
			func(configuration *PropertyConfiguration) error {
				cfg = configuration
				return nil
			})).To(Succeed())

	g.Expect(cfg).NotTo(BeNil())
}

func TestObjectModelConfiguration_ModifyProperty_WhenPropertyExists_CallsActionWithExistingProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := NewObjectModelConfiguration()
	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	var first *PropertyConfiguration
	var second *PropertyConfiguration

	g.Expect(
		omc.ModifyProperty(
			typeName,
			"FullName",
			func(configuration *PropertyConfiguration) error {
				first = configuration
				return nil
			})).To(Succeed())

	g.Expect(
		omc.ModifyProperty(
			typeName,
			"FullName",
			func(configuration *PropertyConfiguration) error {
				second = configuration
				return nil
			})).To(Succeed())

	g.Expect(first).To(Equal(second))
}
