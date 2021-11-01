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

func TestVersionMetaData_WhenYamlWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var version VersionMetaData
	err := yaml.Unmarshal(yamlBytes, &version)
	g.Expect(err).To(Succeed())
	g.Expect(version.kinds).To(HaveLen(3))
}

func TestVersionMetaData_WhenYamlIllformed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var version VersionMetaData
	err := yaml.Unmarshal(yamlBytes, &version)
	g.Expect(err).NotTo(Succeed())
}

