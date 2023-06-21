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

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
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
	g.Expect(group.versions).To(HaveKey("v1api20210101"))
	g.Expect(group.versions).To(HaveKey("v1api20210515"))
}

func TestGroupConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	yamlBytes := loadTestData(t)

	var group GroupConfiguration
	err := yaml.Unmarshal(yamlBytes, &group)
	g.Expect(err).NotTo(Succeed())
}

func TestGroupConfiguration_FindVersion_GivenTypeName_ReturnsExpectedVersion(t *testing.T) {
	t.Parallel()

	ver := "2021-01-01"
	refTest := test.MakeLocalPackageReference("demo", ver)
	refOther := test.MakeLocalPackageReference("demo", "2022-12-31")
	refAlpha := astmodel.MakeLocalPackageReference("prefix", "demo", "v1alpha1api", ver)
	refBeta := astmodel.MakeLocalPackageReference("prefix", "demo", "v1beta", ver)

	groupConfiguration := NewGroupConfiguration("demo")
	versionConfig := NewVersionConfiguration("2021-01-01")
	groupConfiguration.addVersion(versionConfig.name, versionConfig)

	cases := []struct {
		name          string
		ref           astmodel.PackageReference
		expectedFound bool
	}{
		{"Lookup by version", refTest, true},
		{"Lookup by alpha version", refAlpha, true},
		{"Lookup by beta version", refBeta, true},
		{"Lookup by other version", refOther, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			v, err := groupConfiguration.findVersion(c.ref)
			if c.expectedFound {
				g.Expect(err).To(BeNil())
				g.Expect(v).To(Equal(versionConfig))
			} else {
				g.Expect(err).NotTo(BeNil())
			}
		})
	}

}

func loadTestData(t *testing.T) []byte {
	testName := t.Name()
	index := strings.Index(testName, "_")

	folder := testName[0:index]
	file := string(testName[index+1:]) + ".yaml"
	yamlPath := filepath.Join("testdata", folder, file)

	yamlBytes, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		// If the file doesn't exist we fail the test
		t.Fatalf("unable to load %s (%s)", yamlPath, err)
	}

	return yamlBytes
}

func TestGroupConfiguration_WhenVersionConfigurationNotConsumed_ReturnsErrorWithExpectedTip(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create configuration with the wrong version
	typeConfig := NewTypeConfiguration("Person")
	typeConfig.SupportedFrom.Set("vNext")

	versionConfig := NewVersionConfiguration("2022-01-01")
	versionConfig.addType(typeConfig.name, typeConfig)

	groupConfig := NewGroupConfiguration("demo")
	groupConfig.addVersion(versionConfig.name, versionConfig)

	omConfig := NewObjectModelConfiguration()
	omConfig.addGroup(groupConfig.name, groupConfig)

	// Lookup $supportedFrom for our type - version is from 2021 but our config has 2022
	tn := astmodel.MakeTypeName(
		test.MakeLocalPackageReference(groupConfig.name, "2021-01-01"),
		"Person")

	_, err := omConfig.LookupSupportedFrom(tn)
	g.Expect(err).NotTo(BeNil()) // We expect this error

	err = omConfig.VerifySupportedFromConsumed()
	g.Expect(err).NotTo(BeNil())                                   // We expect an error, config hasn't been used
	g.Expect(err.Error()).To(ContainSubstring("did you mean"))     // We want to receive a tip
	g.Expect(err.Error()).To(ContainSubstring(versionConfig.name)) // and we want the correct version to be suggested
}
