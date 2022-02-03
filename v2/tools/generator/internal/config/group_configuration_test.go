/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

func TestGroupConfiguration_WhenYAMLWellFormed_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var group GroupConfiguration
	err := yaml.Unmarshal(yamlBytes, &group)
	g.Expect(err).To(Succeed())
	// Check for exact versions present in the YAML
	g.Expect(group.versions).To(HaveKey("2021-01-01"))
	g.Expect(group.versions).To(HaveKey("2021-05-15"))
	// Check for local package name equivalents
	g.Expect(group.versions).To(HaveKey("v1alpha1api20210101"))
	g.Expect(group.versions).To(HaveKey("v1alpha1api20210515"))
}

func TestGroupConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	yamlBytes := loadTestData(t)

	var group GroupConfiguration
	err := yaml.Unmarshal(yamlBytes, &group)
	g.Expect(err).NotTo(Succeed())
}

func loadTestData(t *testing.T) []byte {
	testName := t.Name()
	index := strings.Index(testName, "_")

	folder := string(testName[0:index])
	file := string(testName[index+1:]) + ".yaml"
	yamlPath := filepath.Join("testdata", folder, file)

	yamlBytes, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		// If the file doesn't exist we fail the test
		t.Fatalf("unable to load %s (%s)", yamlPath, err)
	}

	return yamlBytes
}
