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

func TestPropertyConfiguration_WhenYamlWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).To(Succeed())
	g.Expect(*property.renamedTo).To(Equal("DemoProperty"))
	g.Expect(*property.armReference).To(BeTrue())
}

func TestPropertyConfiguration_WhenYamlIllFormed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).NotTo(Succeed())
}

func TestPropertyConfiguration_ARMReference_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property").SetARMReference(true)

	isReference, ok := property.ARMReference()
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeTrue())
}

func TestPropertyConfiguration_ARMReference_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	_, ok := property.ARMReference()
	g.Expect(ok).To(BeFalse())
}

func TestPropertyConfiguration_FindUnusedARMReferences_WhenNotConfigured_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.ARMReference()

	g.Expect(property.FindUnusedARMReferences()).To(BeEmpty())
}

func TestPropertyConfiguration_FindUnusedARMReferences_WhenReferenceUsed_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property").SetARMReference(true)
	property.ARMReference()

	g.Expect(property.FindUnusedARMReferences()).To(BeEmpty())
}

func TestPropertyConfiguration_FindUnusedARMReferences_WhenReferenceNotUsed_ReturnsExpectedMessage(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property").SetARMReference(true)

	unused := property.FindUnusedARMReferences()
	g.Expect(unused).To(HaveLen(1))
	g.Expect(unused[0]).To(ContainSubstring(property.name))
}
