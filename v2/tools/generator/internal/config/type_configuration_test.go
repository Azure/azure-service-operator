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

func TestTypeConfiguration_WhenYamlWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var kind TypeConfiguration
	err := yaml.Unmarshal(yamlBytes, &kind)
	g.Expect(err).To(Succeed())
	g.Expect(kind.properties).To(HaveLen(4))
	g.Expect(kind.renamedTo).To(Equal("Demo"))
}

func TestTypeConfiguration_WhenYamlIllformed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var kind TypeConfiguration
	err := yaml.Unmarshal(yamlBytes, &kind)
	g.Expect(err).NotTo(Succeed())
}

func TestTypeConfiguration_TypeRename_WhenRenameConfigured_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	config := loadTestTypeConfiguration(t)
	name, ok := config.TypeRename()

	g.Expect(name).To(Equal("Demo"))
	g.Expect(ok).To(BeTrue())
}

func TestTypeConfiguration_TypeRename_WhenRenameConfigured_FlagsRenameAsObserved(t *testing.T) {
	g := NewGomegaWithT(t)
	config := loadTestTypeConfiguration(t)

	config.TypeRename()
	g.Expect(config.usedRenamedTo).To(BeTrue())
}

func TestTypeConfiguration_TypeRename_WhenRenameNotConfigured_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	config := loadTestTypeConfiguration(t)

	name, ok := config.TypeRename()
	g.Expect(name).To(Equal(""))
	g.Expect(ok).To(BeFalse())
}

func TestTypeConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	config := loadTestTypeConfiguration(t)
	isReference, ok := config.ARMReference("Spouse")
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeTrue())
}

func TestTypeConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	config := loadTestTypeConfiguration(t)

	isReference, ok := config.ARMReference("FullName")
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeFalse())
}

func TestTypeConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	config := NewTypeConfiguration()

	_, ok := config.ARMReference("KnownAs")
	g.Expect(ok).To(BeFalse())
}

func loadTestTypeConfiguration(t *testing.T) *TypeConfiguration {
	yamlBytes := loadTestData(t)
	var model TypeConfiguration
	err := yaml.Unmarshal(yamlBytes, &model)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return &model
}
