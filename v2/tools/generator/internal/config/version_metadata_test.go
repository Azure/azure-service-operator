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

func TestVersionMetaDataReadsFromYaml(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var version VersionMetaData
	err := yaml.Unmarshal(yamlBytes, &version)
	if err != nil {
		t.Fatalf("unable to deserialize yaml testdata: %s", err)
	}

	g.Expect(version.kinds).To(HaveLen(3))
}

