/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGroupConfiguration_WhenYamlWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var group GroupConfiguration
	err := yaml.Unmarshal(yamlBytes, &group)
	g.Expect(err).To(Succeed())
	g.Expect(group.versions).To(HaveLen(2))
}

func TestGroupConfiguration_WhenYamlIllformed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)
	yamlBytes := loadTestData(t)

	var group GroupConfiguration
	err := yaml.Unmarshal(yamlBytes, &group)
	g.Expect(err).NotTo(Succeed())
}

func TestGroupConfiguration_TypeRename_WhenTypeFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	yamlBytes := loadTestData(t)
	var group GroupConfiguration
	err := yaml.Unmarshal(yamlBytes, &group)
	g.Expect(err).To(Succeed())

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	name, ok := group.TypeRename(typeName)
	g.Expect(ok).To(BeTrue())
	g.Expect(name).To(Equal("Address"))
}

func loadTestData(t *testing.T) []byte {
	yamlPath := filepath.Join("testdata", t.Name()+".yaml")
	yamlBytes, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		// If the file doesn't exist we fail the test
		t.Fatalf("unable to load %s", yamlPath)
	}

	return yamlBytes
}

