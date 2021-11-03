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

func TestObjectModelConfiguration_WhenYamlWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var model ObjectModelConfiguration
	err := yaml.Unmarshal(yamlBytes, &model)
	g.Expect(err).To(Succeed())
	g.Expect(model.groups).To(HaveLen(1))
}

func TestObjectModelConfiguration_WhenYamlIllformed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var model ObjectModelConfiguration
	err := yaml.Unmarshal(yamlBytes, &model)
	g.Expect(err).NotTo(Succeed())
}
