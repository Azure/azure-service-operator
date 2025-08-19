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

func TestPropertyConfiguration_WhenYAMLWellFormed_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).To(Succeed())
	g.Expect(*property.NameInNextVersion.value).To(Equal("DemoProperty"))
	g.Expect(*property.ReferenceType.value).To(Equal(ReferenceTypeARM))
}

func TestPropertyConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).NotTo(Succeed())
}

func TestPropertyConfiguration_ARMReference_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.ReferenceType.Set(ReferenceTypeARM)

	referenceType, ok := property.ReferenceType.Lookup()
	g.Expect(ok).To(BeTrue())
	g.Expect(referenceType).To(Equal(ReferenceTypeARM))
}

func TestPropertyConfiguration_ARMReference_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	_, ok := property.ReferenceType.Lookup()
	g.Expect(ok).To(BeFalse())
}

func TestPropertyConfiguration_VerifyARMReferenceConsumed_WhenNotConfigured_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	_, _ = property.ReferenceType.Lookup()

	g.Expect(property.ReferenceType.VerifyConsumed()).To(Succeed())
}

func TestPropertyConfiguration_VerifyARMReferenceConsumed_WhenReferenceUsed_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.ReferenceType.Set(ReferenceTypeARM)

	_, _ = property.ReferenceType.Lookup()

	g.Expect(property.ReferenceType.VerifyConsumed()).To(Succeed())
}

func TestPropertyConfiguration_VerifyARMReferenceConsumed_WhenReferenceNotUsed_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.ReferenceType.Set(ReferenceTypeARM)

	err := property.ReferenceType.VerifyConsumed()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(property.name))
}

func TestPropertyConfiguration_IsSecret_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.IsSecret.Set(true)

	isSecret, ok := property.IsSecret.Lookup()
	g.Expect(ok).To(BeTrue())
	g.Expect(isSecret).To(BeTrue())
}

func TestPropertyConfiguration_IsSecret_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	isSecret, ok := property.IsSecret.Lookup()

	g.Expect(isSecret).To(BeFalse())
	g.Expect(ok).To(BeFalse())
}

func TestPropertyConfiguration_Merge_WhenNil_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewPropertyConfiguration("Property")
	err := base.Merge(nil)
	g.Expect(err).To(Succeed())
}

func TestPropertyConfiguration_Merge_WhenEmpty_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewPropertyConfiguration("Property")
	other := NewPropertyConfiguration("Property")
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
}

func TestPropertyConfiguration_Merge_WhenNoConflicts_MergesSuccessfully(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewPropertyConfiguration("Property")
	base.Description.Set("Base description")
	
	other := NewPropertyConfiguration("Property")
	other.ImportConfigMapMode.Set(ImportConfigMapMode(ImportConfigMapModeOptional))
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
	
	// Verify both values are present
	desc, ok := base.Description.Lookup()
	g.Expect(ok).To(BeTrue())
	g.Expect(desc).To(Equal("Base description"))
	
	mode, ok := base.ImportConfigMapMode.Lookup()
	g.Expect(ok).To(BeTrue())
	g.Expect(mode).To(Equal(ImportConfigMapMode(ImportConfigMapModeOptional)))
}

func TestPropertyConfiguration_Merge_WhenConflicts_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewPropertyConfiguration("Property")
	base.Description.Set("Base description")
	
	other := NewPropertyConfiguration("Property")
	other.Description.Set("Other description")
	
	err := base.Merge(other)
	g.Expect(err).To(MatchError(ContainSubstring("conflict")))
}
