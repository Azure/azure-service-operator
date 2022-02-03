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

	name, ok := typeConfig.nameInNextVersion.read()
	g.Expect(name).To(Equal("Demo"))
	g.Expect(ok).To(BeTrue())

	export, ok := typeConfig.export.read()
	g.Expect(export).To(BeTrue())
	g.Expect(ok).To(BeTrue())
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
	typeConfig := NewTypeConfiguration("Person")
	typeConfig.nameInNextVersion.write("Address")

	name, err := typeConfig.LookupNameInNextVersion()

	g.Expect(name).To(Equal("Address"))
	g.Expect(err).To(Succeed())
}

func TestTypeConfiguration_TypeRename_WhenRenameNotConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	typeConfig := NewTypeConfiguration("Person")

	name, err := typeConfig.LookupNameInNextVersion()
	g.Expect(name).To(Equal(""))
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring(typeConfig.name))
}

func TestTypeConfiguration_VerifyTypeRenameConsumed_WhenRenameUsed_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeConfig := NewTypeConfiguration("Person")
	typeConfig.nameInNextVersion.write("Party")

	_, err := typeConfig.LookupNameInNextVersion()
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.VerifyNameInNextVersionConsumed()).To(Succeed())
}

func TestTypeConfiguration_VerifyTypeRenameConsumed_WhenRenameUnused_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeConfig := NewTypeConfiguration("Person")
	typeConfig.nameInNextVersion.write("Party")

	err := typeConfig.VerifyNameInNextVersionConsumed()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(typeConfig.name))
}
