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


