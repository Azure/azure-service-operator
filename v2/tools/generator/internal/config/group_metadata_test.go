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
)

func TestGroupMetaDataReadsFromYaml(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var group GroupMetaData
	err := yaml.Unmarshal(yamlBytes, &group)
	if err != nil {
		t.Fatalf("unable to deserialize yaml testdata: %s", err)
	}

	g.Expect(group.versions).To(HaveLen(2))
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
