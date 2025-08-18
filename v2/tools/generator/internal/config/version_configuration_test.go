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

func TestVersionConfiguration_WhenYAMLWellFormed_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var versionConfig VersionConfiguration
	err := yaml.Unmarshal(yamlBytes, &versionConfig)
	g.Expect(err).To(Succeed())
	g.Expect(versionConfig.types).To(HaveLen(3))
}

func TestVersionConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var versionConfig VersionConfiguration
	err := yaml.Unmarshal(yamlBytes, &versionConfig)
	g.Expect(err).NotTo(Succeed())
}

func TestVersionConfiguration_AddTypeAlias_WhenTypeKnown_AddsAlias(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vc := NewVersionConfiguration("1")
	g.Expect(vc.addType("Person", NewTypeConfiguration("Person"))).To(Succeed())
	g.Expect(vc.addTypeAlias("Person", "Party")).To(Succeed())

	party := vc.findType("Party")
	g.Expect(party).To(Equal(NewTypeConfiguration("Person")))
}

func TestVersionConfiguration_AddTypeAlias_WhenTypeUnknown_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vc := NewVersionConfiguration("1")

	g.Expect(vc.addTypeAlias("Person", "Party")).NotTo(Succeed())
}

func TestVersionConfiguration_AddTypeAlias_WhenTypeClashes_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vc := NewVersionConfiguration("1")
	g.Expect(vc.addType("Person", NewTypeConfiguration("Person"))).To(Succeed())
	g.Expect(vc.addType("Party", NewTypeConfiguration("Party"))).To(Succeed())

	g.Expect(vc.addTypeAlias("Person", "Party")).NotTo(Succeed())
}

/*
 * Duplicate Key Detection Tests
 */

func TestVersionConfiguration_UnmarshalYAML_WhenDuplicateTypes_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlContent := `
Person:
  $export: true
Person:
  $renameTo: "Individual"
`

	var vc VersionConfiguration
	err := yaml.Unmarshal([]byte(yamlContent), &vc)
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring("duplicate type configuration"))
	g.Expect(err.Error()).To(ContainSubstring("Person"))
}

func TestVersionConfiguration_UnmarshalYAML_WhenDuplicateTypesCaseInsensitive_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlContent := `
person:
  $export: true
PERSON:
  $renameTo: "Individual"
`

	var vc VersionConfiguration
	err := yaml.Unmarshal([]byte(yamlContent), &vc)
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring("duplicate type configuration"))
	g.Expect(err.Error()).To(ContainSubstring("PERSON"))
}
