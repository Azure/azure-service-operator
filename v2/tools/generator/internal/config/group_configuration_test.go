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

	person := NewTypeConfiguration("Person").SetTypeRename("Party")
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	name, ok := group.TypeRename(typeName)
	g.Expect(ok).To(BeTrue())
	g.Expect(name).To(Equal("Party"))
}

func TestGroupConfiguration_TypeRename_WhenTypeNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	person := NewTypeConfiguration("Person").SetTypeRename("Party")
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Address")
	name, ok := group.TypeRename(typeName)
	g.Expect(ok).To(BeFalse())
	g.Expect(name).To(Equal(""))
}

func TestGroupConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	spouse := NewPropertyConfiguration("Spouse").SetARMReference(true)
	person := NewTypeConfiguration("Person").Add(spouse)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, ok := group.ARMReference(typeName, "Spouse")
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeTrue())
}

func TestGroupConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	fullNAme := NewPropertyConfiguration("FullName").SetARMReference(false)
	person := NewTypeConfiguration("Person").Add(fullNAme)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, ok := group.ARMReference(typeName, "FullName")
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeFalse())
}

func TestGroupConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	fullNAme := NewPropertyConfiguration("FullName").SetARMReference(false)
	person := NewTypeConfiguration("Person").Add(fullNAme)
	version2015 := NewVersionConfiguration("v20200101").Add(person)
	group := NewGroupConfiguration(test.Group).Add(version2015)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	_, ok := group.ARMReference(typeName, "KnownAs")
	g.Expect(ok).To(BeFalse())
}

func loadTestGroup(t *testing.T) *GroupConfiguration {
	yamlBytes := loadTestData(t)
	var group GroupConfiguration
	err := yaml.Unmarshal(yamlBytes, &group)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return &group
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
		t.Fatalf("unable to load %s", yamlPath)
	}

	return yamlBytes
}
