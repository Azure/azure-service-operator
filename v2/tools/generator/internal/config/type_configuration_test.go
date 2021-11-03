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

func TestTypeConfigurationLookupTypeRename_WhenRenameConfigured_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	newName := "Demo"
	config := NewTypeConfiguration(newName)

	name, ok := config.LookupTypeRename()

	g.Expect(name).To(Equal(newName))
	g.Expect(ok).To(BeTrue())
}

func TestTypeConfigurationLookupTypeRename_WhenRenameConfigured_FlagsRenameAsObserved(t *testing.T) {
	g := NewGomegaWithT(t)
	newName := "Demo"
	config := NewTypeConfiguration(newName)

	config.LookupTypeRename()
	g.Expect(config.usedRenamedTo).To(BeTrue())
}

func TestTypeConfiguration_LookupTypeRename_WhenRenameNotConfigured_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	config := &TypeConfiguration{}

	name, ok := config.LookupTypeRename()
	g.Expect(name).To(Equal(""))
	g.Expect(ok).To(BeFalse())
}
