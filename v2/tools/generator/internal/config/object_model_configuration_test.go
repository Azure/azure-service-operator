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

/*
 * YAML tests
 */

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

/*
 * Type Rename Tests
 */

func TestObjectModelConfiguration_TypeRename_WhenTypeFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()

	g.Expect(
		omc.ModifyType(
			typeName,
			func(tc *TypeConfiguration) error {
				tc.NameInNextVersion.Set("Party")
				return nil
			})).
		To(Succeed())

	nextName, err := omc.TypeNameInNextVersion.Lookup(typeName)
	g.Expect(err).To(Succeed())
	g.Expect(nextName).To(Equal("Party"))
}

func TestObjectModelConfiguration_TypeRename_WhenTypeNotFound_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Address")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			typeName,
			func(tc *TypeConfiguration) error {
				tc.NameInNextVersion.Set("Party")
				return nil
			})).
		To(Succeed())

	otherName := astmodel.MakeTypeName(test.Pkg2020, "Location")
	nextName, err := omc.TypeNameInNextVersion.Lookup(otherName)

	g.Expect(err).NotTo(Succeed())
	g.Expect(nextName).To(Equal(""))
	g.Expect(err.Error()).To(ContainSubstring(typeName.Name()))
}

func TestObjectModelConfiguration_VerifyTypeRenamesConsumed_WhenRenameUsed_ReturnsEmptySlice(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			typeName,
			func(tc *TypeConfiguration) error {
				tc.NameInNextVersion.Set("Party")
				return nil
			})).
		To(Succeed())

	_, err := omc.TypeNameInNextVersion.Lookup(typeName)
	g.Expect(err).To(Succeed())
	g.Expect(omc.TypeNameInNextVersion.VerifyConsumed()).To(Succeed())
}

func TestObjectModelConfiguration_VerifyTypeRenamesConsumed_WhenRenameUnused_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			typeName,
			func(tc *TypeConfiguration) error {
				tc.NameInNextVersion.Set("Party")
				return nil
			})).
		To(Succeed())

	g.Expect(omc.TypeNameInNextVersion.VerifyConsumed()).NotTo(Succeed())
}

/*
 * ARM Reference Tests
 */

func TestObjectModelConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			typeName,
			"Spouse",
			func(pc *PropertyConfiguration) error {
				pc.ARMReference.Set(true)
				return nil
			})).
		To(Succeed())

	isReference, err := omc.ARMReference.Lookup(typeName, "Spouse")
	g.Expect(err).To(BeNil())
	g.Expect(isReference).To(BeTrue())
}

func TestObjectModelConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")

	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			typeName,
			"FullName",
			func(pc *PropertyConfiguration) error {
				pc.ARMReference.Set(false)
				return nil
			})).
		To(Succeed())

	isReference, err := omc.ARMReference.Lookup(typeName, "FullName")
	g.Expect(err).To(BeNil())
	g.Expect(isReference).To(BeFalse())
}

func TestObjectModelConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")

	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			typeName,
			"Spouse",
			func(pc *PropertyConfiguration) error {
				pc.ARMReference.Set(true)
				return nil
			})).
		To(Succeed())

	_, err := omc.ARMReference.Lookup(typeName, "KnownAs")
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring("KnownAs"))
}

func TestObjectModelConfiguration_VerifyARMReferencesConsumed_WhenReferenceUsed_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")

	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			typeName,
			"Spouse",
			func(pc *PropertyConfiguration) error {
				pc.ARMReference.Set(true)
				return nil
			})).
		To(Succeed())

	ref, err := omc.ARMReference.Lookup(typeName, "Spouse")
	g.Expect(ref).To(BeTrue())
	g.Expect(err).To(Succeed())
	g.Expect(omc.ARMReference.VerifyConsumed()).To(Succeed())
}

func TestObjectModelConfiguration_VerifyARMReferencesConsumed_WhenReferenceNotUsed_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			typeName,
			"Spouse",
			func(pc *PropertyConfiguration) error {
				pc.ARMReference.Set(true)
				return nil
			})).
		To(Succeed())

	g.Expect(
		omc.ARMReference.VerifyConsumed()).NotTo(Succeed())
}

/*
 * Export As Tests
 */

func TestObjectModelConfiguration_LookupExportAs_AfterConsumption_CanLookupUsingNewName(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "People")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			typeName,
			func(tc *TypeConfiguration) error {
				tc.ExportAs.Set("Person")
				tc.NameInNextVersion.Set("Party")
				return nil
			})).
		To(Succeed())

	// Lookup the new name for the type
	exportAs, err := omc.ExportAs.Lookup(typeName)
	g.Expect(err).To(BeNil())

	// Lookup the name in next version using the new name of the type
	omc.AddTypeAlias(typeName, exportAs)
	newTypeName := typeName.WithName(exportAs)
	nextName, err := omc.TypeNameInNextVersion.Lookup(newTypeName)
	g.Expect(err).To(BeNil())
	g.Expect(nextName).To(Equal("Party"))
}

/*
 * Modify Group Tests
 */

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

/*
 * Modify Version Tests
 */

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

/*
 * Modify Type Tests
 */

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

/*
 * Modify Property Tests
 */

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

/*
 * SupportedFrom Tests
 */

func TestObjectModelConfiguration_LookupSupportedFrom_WhenConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			name,
			func(tc *TypeConfiguration) error {
				tc.SupportedFrom.Set("beta.5")
				return nil
			})).
		To(Succeed())

	supportedFrom, err := omc.SupportedFrom.Lookup(name)
	g.Expect(err).To(Succeed())
	g.Expect(supportedFrom).To(Equal("beta.5"))
}

func TestObjectModelConfiguration_LookupSupportedFrom_WhenNotConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			name,
			func(tc *TypeConfiguration) error {
				// No change, just provoking creation
				return nil
			})).
		To(Succeed())

	_, err := omc.SupportedFrom.Lookup(name)
	g.Expect(err).NotTo(Succeed())
}

func TestObjectModelConfiguration_LookupSupportedFrom_WhenConsumed_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			name,
			func(tc *TypeConfiguration) error {
				tc.SupportedFrom.Set("beta.5")
				return nil
			})).
		To(Succeed())

	_, err := omc.SupportedFrom.Lookup(name)
	g.Expect(err).To(Succeed())

	err = omc.SupportedFrom.VerifyConsumed()
	g.Expect(err).To(Succeed())
}

func TestObjectModelConfiguration_LookupSupportedFrom_WhenUnconsumed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			name,
			func(tc *TypeConfiguration) error {
				tc.SupportedFrom.Set("beta.5")
				return nil
			})).
		To(Succeed())

	err := omc.SupportedFrom.VerifyConsumed()
	g.Expect(err).NotTo(Succeed())
}

/*
 * PayloadType Tests
 */

func TestObjectModelConfiguration_LookupPayloadType_WhenConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyGroup(
			name.PackageReference,
			func(gc *GroupConfiguration) error {
				gc.PayloadType.Set(ExplicitProperties)
				return nil
			})).
		To(Succeed())

	payloadType, err := omc.PayloadType.Lookup(name.PackageReference)
	g.Expect(err).To(Succeed())
	g.Expect(payloadType).To(Equal(ExplicitProperties))
}

func TestObjectModelConfiguration_LookupPayloadType_WhenNotConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyGroup(
			name.PackageReference,
			func(_ *GroupConfiguration) error {
				// No change, just provoking creation
				return nil
			})).
		To(Succeed())

	_, err := omc.PayloadType.Lookup(name.PackageReference)
	g.Expect(err).NotTo(Succeed())
}

func TestObjectModelConfiguration_VerifyPayloadTypeConsumed_WhenConsumed_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyGroup(
			name.PackageReference,
			func(gc *GroupConfiguration) error {
				gc.PayloadType.Set(OmitEmptyProperties)
				return nil
			})).
		To(Succeed())

	_, err := omc.PayloadType.Lookup(name.PackageReference)
	g.Expect(err).To(Succeed())

	err = omc.PayloadType.VerifyConsumed()
	g.Expect(err).To(Succeed())
}

func TestObjectModelConfiguration_VerifyPayloadTypeConsumed_WhenUnconsumed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyGroup(
			name.PackageReference,
			func(gc *GroupConfiguration) error {
				gc.PayloadType.Set(ExplicitProperties)
				return nil
			})).
		To(Succeed())

	err := omc.PayloadType.VerifyConsumed()
	g.Expect(err).NotTo(Succeed())
}
