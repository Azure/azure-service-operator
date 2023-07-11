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
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).To(Succeed())
	g.Expect(*property.NameInNextVersion.value).To(Equal("DemoProperty"))
	g.Expect(*property.ARMReference.value).To(BeTrue())
}

func TestPropertyConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).NotTo(Succeed())
}

func TestPropertyConfiguration_ARMReference_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.ARMReference.Set(true)

	isReference, err := property.ARMReference.Lookup()
	g.Expect(err).To(Succeed())
	g.Expect(isReference).To(BeTrue())
}

func TestPropertyConfiguration_ARMReference_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	_, err := property.ARMReference.Lookup()
	g.Expect(err).NotTo(Succeed())
	g.Expect(err.Error()).To(ContainSubstring(property.name))
}

func TestPropertyConfiguration_VerifyARMReferenceConsumed_WhenNotConfigured_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	_, _ = property.ARMReference.Lookup()

	g.Expect(property.ARMReference.VerifyConsumed()).To(Succeed())
}

func TestPropertyConfiguration_VerifyARMReferenceConsumed_WhenReferenceUsed_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.ARMReference.Set(true)

	_, _ = property.ARMReference.Lookup()

	g.Expect(property.ARMReference.VerifyConsumed()).To(Succeed())
}

func TestPropertyConfiguration_VerifyARMReferenceConsumed_WhenReferenceNotUsed_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.ARMReference.Set(true)

	err := property.ARMReference.VerifyConsumed()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(property.name))
}

func TestPropertyConfiguration_IsSecret_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.IsSecret.Set(true)

	g.Expect(property.IsSecret.Lookup()).To(BeTrue())
}

func TestPropertyConfiguration_IsSecret_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	isSecret, err := property.IsSecret.Lookup()

	g.Expect(err.Error()).To(ContainSubstring(property.name))
	g.Expect(isSecret).To(BeFalse())
}
