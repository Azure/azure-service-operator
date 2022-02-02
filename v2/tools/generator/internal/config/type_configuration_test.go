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
	g.Expect(*typeConfig.nameInNextVersion).To(Equal("Demo"))
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
	typeConfig := NewTypeConfiguration("Person").SetNameInNextVersion("Address")

	name, err := typeConfig.LookupNameInNextVersion()

	g.Expect(name).To(Equal("Address"))
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.nameInNextVersionConsumed).To(BeTrue())
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

	typeConfig := NewTypeConfiguration("Person").SetNameInNextVersion("Party")
	_, err := typeConfig.LookupNameInNextVersion()
	g.Expect(err).To(Succeed())
	g.Expect(typeConfig.VerifyNameInNextVersionConsumed()).To(Succeed())
}

func TestTypeConfiguration_VerifyTypeRenameConsumed_WhenRenameUnused_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	typeConfig := NewTypeConfiguration("Person").SetNameInNextVersion("Party")

	err := typeConfig.VerifyNameInNextVersionConsumed()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(typeConfig.name))
}
