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
	g.Expect(*property.ReferenceType.value).To(Equal(ReferenceTypeARM))
}

func TestPropertyConfiguration_WhenYAMLBadlyFormed_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).NotTo(Succeed())
}

func TestPropertyConfiguration_ReferenceType_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.ReferenceType.Set(ReferenceTypeARM)

	referenceType, ok := property.ReferenceType.Lookup()
	g.Expect(ok).To(BeTrue())
	g.Expect(referenceType).To(Equal(ReferenceTypeARM))
}

func TestPropertyConfiguration_ReferenceType_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	_, ok := property.ReferenceType.Lookup()
	g.Expect(ok).To(BeFalse())
}

func TestPropertyConfiguration_VerifyReferenceTypeConsumed_WhenNotConfigured_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	_, _ = property.ReferenceType.Lookup()

	g.Expect(property.ReferenceType.VerifyConsumed()).To(Succeed())
}

func TestPropertyConfiguration_VerifyReferenceTypeConsumed_WhenReferenceUsed_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.ReferenceType.Set(ReferenceTypeARM)

	_, _ = property.ReferenceType.Lookup()

	g.Expect(property.ReferenceType.VerifyConsumed()).To(Succeed())
}

func TestPropertyConfiguration_VerifyReferenceTypeConsumed_WhenReferenceNotUsed_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.ReferenceType.Set(ReferenceTypeARM)

	err := property.ReferenceType.VerifyConsumed()
	g.Expect(err).To(MatchError(ContainSubstring(property.name)))
}

func TestPropertyConfiguration_IsSecret_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.IsSecret.Set(true)

	isSecret, ok := property.IsSecret.Lookup()
	g.Expect(ok).To(BeTrue())
	g.Expect(isSecret).To(BeTrue())
}

func TestPropertyConfiguration_IsSecret_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	isSecret, ok := property.IsSecret.Lookup()

	g.Expect(isSecret).To(BeFalse())
	g.Expect(ok).To(BeFalse())
}
