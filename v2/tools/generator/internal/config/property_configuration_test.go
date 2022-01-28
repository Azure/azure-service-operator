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

func TestPropertyConfiguration_WhenYAMLWellFormed_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).To(Succeed())
	g.Expect(*property.renamedTo).To(Equal("DemoProperty"))
	g.Expect(*property.armReference).To(BeTrue())
}

func TestPropertyConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).NotTo(Succeed())
}

func TestPropertyConfiguration_ARMReference_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property").SetARMReference(true)

	isReference, err := property.ARMReference()
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeTrue())
}

func TestPropertyConfiguration_ARMReference_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	_, err := property.ARMReference()
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring(property.name))
}

func TestPropertyConfiguration_FindUnusedARMReferences_WhenNotConfigured_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	_, _ = property.ARMReference()

	g.Expect(property.FindUnusedARMReferences()).To(BeEmpty())
}

func TestPropertyConfiguration_FindUnusedARMReferences_WhenReferenceUsed_ReturnsEmptySlice(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property").SetARMReference(true)
	_, _ = property.ARMReference()

	g.Expect(property.FindUnusedARMReferences()).To(BeEmpty())
}

func TestPropertyConfiguration_FindUnusedARMReferences_WhenReferenceNotUsed_ReturnsExpectedMessage(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property").SetARMReference(true)

	unused := property.FindUnusedARMReferences()
	g.Expect(unused).To(HaveLen(1))
	g.Expect(unused[0]).To(ContainSubstring(property.name))
}

func TestPropertyConfiguration_IsSecret_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.SetIsSecret(true)

	g.Expect(property.IsSecret()).To(BeTrue())
}

func TestPropertyConfiguration_IsSecret_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	isSecret, err := property.IsSecret()

	g.Expect(err.Error()).To(ContainSubstring(property.name))
	g.Expect(isSecret).To(BeFalse())
}
