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

func TestKindMetaDataReadsFromYaml(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var kind KindMetaData
	err := yaml.Unmarshal(yamlBytes, &kind)
	if err != nil {
		t.Fatalf("unable to deserialize yaml testdata: %s", err)
	}

	g.Expect(kind.properties).To(HaveLen(4))
	g.Expect(kind.rename).To(Equal("Demo"))
}

