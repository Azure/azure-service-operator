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
	g := NewGomegaWithT(t)
	yamlBytes := loadTestData(t)

	var group GroupConfiguration
	err := yaml.Unmarshal(yamlBytes, &group)
	g.Expect(err).NotTo(Succeed())
}

func TestGroupConfiguration_TypeRename_WhenTypeFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	person := NewTypeConfiguration("Person").SetTypeRename("Party")
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	name, err := group.TypeRename(typeName)
	g.Expect(err).To(Succeed())
	g.Expect(name).To(Equal("Party"))
}

func TestGroupConfiguration_TypeRename_WhenTypeNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	person := NewTypeConfiguration("Person").SetTypeRename("Party")
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Address")
	name, err := group.TypeRename(typeName)
	g.Expect(err).NotTo(Succeed())
	g.Expect(name).To(Equal(""))
	g.Expect(err.Error()).To(ContainSubstring(typeName.Name()))
}

func TestGroupConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	person := NewTypeConfiguration("Person").Add(spouse)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, err := group.ARMReference(typeName, "Spouse")
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeTrue())
}

func TestGroupConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	fullNAme := NewPropertyConfiguration("FullName").SetARMReference(false)
	person := NewTypeConfiguration("Person").Add(fullNAme)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, err:= group.ARMReference(typeName, "FullName")
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeFalse())
}

func TestGroupConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	fullNAme := NewPropertyConfiguration("FullName").SetARMReference(false)
	person := NewTypeConfiguration("Person").Add(fullNAme)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	_, err := group.ARMReference(typeName, "KnownAs")
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring("KnownAs"))
}

func TestGroupConfiguration_FindUnusedARMReferences_WhenReferenceUsed_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	_, _ = spouse.ARMReference()
	person := NewTypeConfiguration("Person").Add(spouse)
	version := NewVersionConfiguration("2015-01-01").Add(person)
	groupConfig := NewGroupConfiguration("microsoft.demo").Add(version)

	g.Expect(groupConfig.FindUnusedARMReferences()).To(BeEmpty())
}

func TestGroupConfiguration_FindUnusedARMReferences_WhenReferenceNotUsed_ReturnsExpectedMessage(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	person := NewTypeConfiguration("Person").Add(spouse)
	version := NewVersionConfiguration("2015-01-01").Add(person)
	groupConfig := NewGroupConfiguration("microsoft.demo").Add(version)

	unused := groupConfig.FindUnusedARMReferences()
	g.Expect(unused).To(HaveLen(1))
	g.Expect(unused[0]).To(ContainSubstring(groupConfig.name))
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
