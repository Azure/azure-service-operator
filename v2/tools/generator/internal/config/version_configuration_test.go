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

	person := NewTypeConfiguration("Person")
	vc := NewVersionConfiguration("1")
	vc.add(person)

	g.Expect(vc.addTypeAlias("Person", "Party")).To(Succeed())

	party, err := vc.findType("Party")
	g.Expect(party).To(Equal(person))
	g.Expect(err).To(BeNil())
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

	person := NewTypeConfiguration("Person")
	party := NewTypeConfiguration("Party")
	vc := NewVersionConfiguration("1")
	vc.add(person)
	vc.add(party)

	g.Expect(vc.addTypeAlias("Person", "Party")).NotTo(Succeed())
}
