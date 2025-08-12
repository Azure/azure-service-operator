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

	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()

	g.Expect(
		omc.ModifyType(
			typeName,
			func(tc *TypeConfiguration) error {
				tc.NameInNextVersion.Set("Party")
				return nil
			})).
		To(Succeed())

	nextName, ok := omc.TypeNameInNextVersion.Lookup(typeName)
	g.Expect(ok).To(BeTrue())
	g.Expect(nextName).To(Equal("Party"))
}

func TestObjectModelConfiguration_TypeRename_WhenTypeNotFound_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Address")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			typeName,
			func(tc *TypeConfiguration) error {
				tc.NameInNextVersion.Set("Party")
				return nil
			})).
		To(Succeed())

	otherName := astmodel.MakeInternalTypeName(test.Pkg2020, "Location")
	nextName, ok := omc.TypeNameInNextVersion.Lookup(otherName)
	g.Expect(ok).To(BeFalse())
	g.Expect(nextName).To(Equal(""))
}

func TestObjectModelConfiguration_VerifyTypeRenamesConsumed_WhenRenameUsed_ReturnsEmptySlice(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			typeName,
			func(tc *TypeConfiguration) error {
				tc.NameInNextVersion.Set("Party")
				return nil
			})).
		To(Succeed())

	_, ok := omc.TypeNameInNextVersion.Lookup(typeName)
	g.Expect(ok).To(BeTrue())
	g.Expect(omc.TypeNameInNextVersion.VerifyConsumed()).To(Succeed())
}

func TestObjectModelConfiguration_VerifyTypeRenamesConsumed_WhenRenameUnused_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
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

	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			typeName,
			"Spouse",
			func(pc *PropertyConfiguration) error {
				pc.ReferenceType.Set(ReferenceTypeARM)
				return nil
			})).
		To(Succeed())

	referenceType, ok := omc.ReferenceType.Lookup(typeName, "Spouse")
	g.Expect(ok).To(BeTrue())
	g.Expect(referenceType).To(Equal(ReferenceTypeARM))
}

func TestObjectModelConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")

	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			typeName,
			"FullName",
			func(pc *PropertyConfiguration) error {
				pc.ReferenceType.Set(ReferenceTypeSimple)
				return nil
			})).
		To(Succeed())

	referenceType, ok := omc.ReferenceType.Lookup(typeName, "FullName")
	g.Expect(ok).To(BeTrue())
	g.Expect(referenceType).To(Equal(ReferenceTypeSimple))
}

func TestObjectModelConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")

	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			typeName,
			"Spouse",
			func(pc *PropertyConfiguration) error {
				pc.ReferenceType.Set(ReferenceTypeARM)
				return nil
			})).
		To(Succeed())

	_, ok := omc.ReferenceType.Lookup(typeName, "KnownAs")
	g.Expect(ok).To(BeFalse())
}

func TestObjectModelConfiguration_VerifyARMReferencesConsumed_WhenReferenceUsed_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")

	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			typeName,
			"Spouse",
			func(pc *PropertyConfiguration) error {
				pc.ReferenceType.Set(ReferenceTypeARM)
				return nil
			})).
		To(Succeed())

	referenceType, ok := omc.ReferenceType.Lookup(typeName, "Spouse")
	g.Expect(ok).To(BeTrue())
	g.Expect(referenceType).To(Equal(ReferenceTypeARM))
	g.Expect(omc.ReferenceType.VerifyConsumed()).To(Succeed())
}

func TestObjectModelConfiguration_VerifyARMReferencesConsumed_WhenReferenceNotUsed_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			typeName,
			"Spouse",
			func(pc *PropertyConfiguration) error {
				pc.ReferenceType.Set(ReferenceTypeARM)
				return nil
			})).
		To(Succeed())

	g.Expect(
		omc.ReferenceType.VerifyConsumed()).NotTo(Succeed())
}

/*
 * Export As Tests
 */

func TestObjectModelConfiguration_LookupExportAs_AfterConsumption_CanLookupUsingNewName(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "People")
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
	exportAs, ok := omc.ExportAs.Lookup(typeName)
	g.Expect(ok).To(BeTrue())

	// Lookup the name in next version using the new name of the type
	omc.AddTypeAlias(typeName, exportAs)
	newTypeName := typeName.WithName(exportAs)
	nextName, ok := omc.TypeNameInNextVersion.Lookup(newTypeName)
	g.Expect(ok).To(BeTrue())
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
	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
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
	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
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
	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
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
	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
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

	name := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			name,
			func(tc *TypeConfiguration) error {
				tc.SupportedFrom.Set("beta.5")
				return nil
			})).
		To(Succeed())

	supportedFrom, ok := omc.SupportedFrom.Lookup(name)
	g.Expect(ok).To(BeTrue())
	g.Expect(supportedFrom).To(Equal("beta.5"))
}

func TestObjectModelConfiguration_LookupSupportedFrom_WhenNotConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			name,
			func(tc *TypeConfiguration) error {
				// No change, just provoking creation
				return nil
			})).
		To(Succeed())

	_, ok := omc.SupportedFrom.Lookup(name)
	g.Expect(ok).To(BeFalse())
}

func TestObjectModelConfiguration_LookupSupportedFrom_WhenConsumed_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			name,
			func(tc *TypeConfiguration) error {
				tc.SupportedFrom.Set("beta.5")
				return nil
			})).
		To(Succeed())

	_, ok := omc.SupportedFrom.Lookup(name)
	g.Expect(ok).To(BeTrue())

	err := omc.SupportedFrom.VerifyConsumed()
	g.Expect(err).To(Succeed())
}

func TestObjectModelConfiguration_LookupSupportedFrom_WhenUnconsumed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
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

	name := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyGroup(
			name.InternalPackageReference(),
			func(gc *GroupConfiguration) error {
				gc.PayloadType.Set(ExplicitProperties)
				return nil
			})).
		To(Succeed())

	payloadType, ok := omc.PayloadType.Lookup(name, "")
	g.Expect(ok).To(BeTrue())
	g.Expect(payloadType).To(Equal(ExplicitProperties))
}

func TestObjectModelConfiguration_LookupPayloadType_WhenNotConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyGroup(
			name.InternalPackageReference(),
			func(_ *GroupConfiguration) error {
				// No change, just provoking creation
				return nil
			})).
		To(Succeed())

	_, ok := omc.PayloadType.Lookup(name, "")
	g.Expect(ok).To(BeFalse())
}

func TestObjectModelConfiguration_VerifyPayloadTypeConsumed_WhenConsumed_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyGroup(
			name.InternalPackageReference(),
			func(gc *GroupConfiguration) error {
				gc.PayloadType.Set(OmitEmptyProperties)
				return nil
			})).
		To(Succeed())

	_, ok := omc.PayloadType.Lookup(name, "")
	g.Expect(ok).To(BeTrue())

	err := omc.PayloadType.VerifyConsumed()
	g.Expect(err).To(Succeed())
}

func TestObjectModelConfiguration_VerifyPayloadTypeConsumed_WhenUnconsumed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	name := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	omc := NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyGroup(
			name.InternalPackageReference(),
			func(gc *GroupConfiguration) error {
				gc.PayloadType.Set(ExplicitProperties)
				return nil
			})).
		To(Succeed())

	err := omc.PayloadType.VerifyConsumed()
	g.Expect(err).NotTo(Succeed())
}

func TestObjectModelConfiguration_Merge_WhenNil_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewObjectModelConfiguration()
	err := base.Merge(nil)
	g.Expect(err).To(Succeed())
}

func TestObjectModelConfiguration_Merge_WhenEmpty_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewObjectModelConfiguration()
	other := NewObjectModelConfiguration()
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
}

func TestObjectModelConfiguration_Merge_WhenNoConflicts_MergesSuccessfully(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewObjectModelConfiguration()
	baseGroup := NewGroupConfiguration("Network")
	base.addGroup("Network", baseGroup)
	
	other := NewObjectModelConfiguration()
	otherGroup := NewGroupConfiguration("Compute")
	other.addGroup("Compute", otherGroup)
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
	
	// Should have both groups (stored with lowercase keys)
	g.Expect(base.groups).To(HaveKey("network"))
	g.Expect(base.groups).To(HaveKey("compute"))
}

func TestObjectModelConfiguration_Merge_WhenGroupConflicts_MergesGroups(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewObjectModelConfiguration()
	baseGroup := NewGroupConfiguration("Network")
	baseVersion := NewVersionConfiguration("v1api20210101")
	baseGroup.addVersion("v1api20210101", baseVersion)
	base.addGroup("Network", baseGroup)
	
	other := NewObjectModelConfiguration()
	otherGroup := NewGroupConfiguration("Network")
	otherVersion := NewVersionConfiguration("v1api20210515")
	otherGroup.addVersion("v1api20210515", otherVersion)
	other.addGroup("Network", otherGroup)
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
	
	// Should have merged the group configurations (stored with lowercase key)
	mergedGroup := base.groups["network"] // lowercase key
	g.Expect(mergedGroup).NotTo(BeNil())
	g.Expect(mergedGroup.versions).To(HaveKey("v1api20210101"))
	g.Expect(mergedGroup.versions).To(HaveKey("v1api20210515"))
}
