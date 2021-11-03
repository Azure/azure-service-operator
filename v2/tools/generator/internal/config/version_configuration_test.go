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

func TestVersionConfiguration_WhenYamlWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var version VersionConfiguration
	err := yaml.Unmarshal(yamlBytes, &version)
	g.Expect(err).To(Succeed())
	g.Expect(version.types).To(HaveLen(3))
}

func TestVersionConfiguration_WhenYamlIllformed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var version VersionConfiguration
	err := yaml.Unmarshal(yamlBytes, &version)
	g.Expect(err).NotTo(Succeed())
}

func TestVersionConfiguration_LookupTypeRename_WhenTypeFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	newName := "Party"

	config := &VersionConfiguration{
		types: map[string]*TypeConfiguration{
			"person": {
				renamedTo: newName,
			},
		},
	}

	name, ok := config.LookupTypeRename("Person")
	g.Expect(ok).To(BeTrue())
	g.Expect(name).To(Equal(newName))
}

func TestVersionConfiguration_LookupTypeRename_WhenTypeNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	newName := "Party"

	config := &VersionConfiguration{
		types: map[string]*TypeConfiguration{
			"person": {
				renamedTo: newName,
			},
		},
	}

	name, ok := config.LookupTypeRename("Address")
	g.Expect(ok).To(BeFalse())
	g.Expect(name).To(Equal(""))
}
