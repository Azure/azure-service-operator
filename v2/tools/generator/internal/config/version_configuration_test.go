/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestVersionConfiguration_WhenYamlWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var version VersionConfiguration
	err := yaml.Unmarshal(yamlBytes, &version)
	g.Expect(err).To(Succeed())
	g.Expect(version.types).To(HaveLen(3))
}

func TestVersionConfiguration_WhenYamlIllformed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var version VersionConfiguration
	err := yaml.Unmarshal(yamlBytes, &version)
	g.Expect(err).NotTo(Succeed())
}

func TestVersionConfiguration_TypeRename_WhenTypeFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	newName := "Party"

	config := &VersionConfiguration{
		types: map[string]*TypeConfiguration{
			"person": {
				renamedTo: newName,
			},
		},
	}

	name, ok := config.TypeRename("Person")
	g.Expect(ok).To(BeTrue())
	g.Expect(name).To(Equal(newName))
}

func TestVersionConfiguration_TypeRename_WhenTypeNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	config := loadVersionConfiguration(t)
	name, ok := config.TypeRename("Address")
	g.Expect(ok).To(BeFalse())
	g.Expect(name).To(Equal(""))
}

func TestVersionConfiguration_ARMReference_WhenSpousePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	config := loadVersionConfiguration(t)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, ok := config.ARMReference(typeName.Name(), "Spouse")
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeTrue())
}

func TestVersionConfiguration_ARMReference_WhenFullNamePropertyFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	config := loadVersionConfiguration(t)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	isReference, ok := config.ARMReference(typeName.Name(), "FullName")
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeFalse())
}

func TestVersionConfiguration_ARMReference_WhenPropertyNotFound_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)
	config := loadVersionConfiguration(t)

	typeName := astmodel.MakeTypeName(test.Pkg2020, "Person")
	_, ok := config.ARMReference(typeName.Name(), "KnownAs")
	g.Expect(ok).To(BeFalse())
}

func loadVersionConfiguration(t *testing.T) *VersionConfiguration {
	yamlBytes := loadTestData(t)
	var config VersionConfiguration
	err := yaml.Unmarshal(yamlBytes, &config)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return &config
}
