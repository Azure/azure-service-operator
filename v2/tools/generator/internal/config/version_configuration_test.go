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
	g.Expect(version.kinds).To(HaveLen(3))
}

func TestVersionConfiguration_WhenYamlIllformed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var version VersionConfiguration
	err := yaml.Unmarshal(yamlBytes, &version)
	g.Expect(err).NotTo(Succeed())
}

