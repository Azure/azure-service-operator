/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"os"
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

			v := groupConfiguration.findVersion(c.ref)
			if c.expectedFound {
				g.Expect(v).To(Equal(versionConfig))
			} else {
				g.Expect(v).To(BeNil())
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

	yamlBytes, err := os.ReadFile(yamlPath)
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

	// Lookup $supportedFrom for our type - version is from 2021 but our config has 2022, so it won't be found
	tn := astmodel.MakeInternalTypeName(
		test.MakeLocalPackageReference(groupConfig.name, "2021-01-01"),
		"Person")

	_, ok := omConfig.SupportedFrom.Lookup(tn)
	g.Expect(ok).To(BeFalse())

	err := omConfig.SupportedFrom.VerifyConsumed()
	g.Expect(err).NotTo(BeNil())                                   // We expect an error, config hasn't been used
	g.Expect(err.Error()).To(ContainSubstring("did you mean"))     // We want to receive a tip
	g.Expect(err.Error()).To(ContainSubstring(versionConfig.name)) // and we want the correct version to be suggested
}

/*
 * PayloadType tests
 */

func TestGroupConfiguration_LookupPayloadType_WhenConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	groupConfig := NewGroupConfiguration("Network")
	groupConfig.PayloadType.Set(ExplicitProperties)

	payloadType, ok := groupConfig.PayloadType.Lookup()
	g.Expect(payloadType).To(Equal(ExplicitProperties))
	g.Expect(ok).To(BeTrue())
}

func TestGroupConfiguration_LookupPayloadType_WhenNotConfigured_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	groupConfig := NewGroupConfiguration("Network")

	name, ok := groupConfig.PayloadType.Lookup()
	g.Expect(name).To(Equal(PayloadType("")))
	g.Expect(ok).To(BeFalse())
}

func TestGroupConfiguration_VerifyPayloadTypeConsumed_WhenConsumed_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	groupConfig := NewGroupConfiguration("Network")
	groupConfig.PayloadType.Set(OmitEmptyProperties)

	_, ok := groupConfig.PayloadType.Lookup()
	g.Expect(ok).To(BeTrue())
	g.Expect(groupConfig.PayloadType.VerifyConsumed()).To(Succeed())
}

func TestGroupConfiguration_VerifyPayloadTypeConsumed_WhenNotConsumed_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	groupConfig := NewGroupConfiguration("Network")
	groupConfig.PayloadType.Set(ExplicitProperties)

	err := groupConfig.PayloadType.VerifyConsumed()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(groupConfig.name))
}

func TestGroupConfiguration_Merge_WhenNil_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewGroupConfiguration("Network")
	err := base.Merge(nil)
	g.Expect(err).To(Succeed())
}

func TestGroupConfiguration_Merge_WhenEmpty_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewGroupConfiguration("Network")
	other := NewGroupConfiguration("Network")
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
}

func TestGroupConfiguration_Merge_WhenNoConflicts_MergesSuccessfully(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewGroupConfiguration("Network")
	base.PayloadType.Set(ExplicitProperties)
	baseVersion := NewVersionConfiguration("v1api20210101")
	base.addVersion("v1api20210101", baseVersion)
	
	other := NewGroupConfiguration("Network")
	otherVersion := NewVersionConfiguration("v1api20210515")
	other.addVersion("v1api20210515", otherVersion)
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
	
	// Should have both versions
	g.Expect(base.versions).To(HaveKey("v1api20210101"))
	g.Expect(base.versions).To(HaveKey("v1api20210515"))
	
	// PayloadType should be preserved
	val, ok := base.PayloadType.read()
	g.Expect(ok).To(BeTrue())
	g.Expect(val).To(Equal(ExplicitProperties))
}

func TestGroupConfiguration_Merge_WhenVersionConflicts_MergesVersions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewGroupConfiguration("Network")
	baseVersion := NewVersionConfiguration("v1api20210101")
	base.addVersion("v1api20210101", baseVersion)
	
	other := NewGroupConfiguration("Network")
	otherVersion := NewVersionConfiguration("v1api20210101")
	otherTypeConfig := NewTypeConfiguration("VirtualNetwork")
	otherVersion.addType("VirtualNetwork", otherTypeConfig)
	other.addVersion("v1api20210101", otherVersion)
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
	
	// Should have merged the version configurations
	mergedVersion := base.versions["v1api20210101"]
	g.Expect(mergedVersion).NotTo(BeNil())
	g.Expect(mergedVersion.types).To(HaveKey("virtualnetwork")) // lowercase key
}
