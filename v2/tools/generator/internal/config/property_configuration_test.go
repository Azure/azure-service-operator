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

func TestPropertyConfiguration_Secrecy_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.Secrecy.Set(astmodel.ImportSecretModeRequired)

	secrecy, ok := property.Secrecy.Lookup()
	g.Expect(ok).To(BeTrue())
	g.Expect(secrecy).To(Equal(astmodel.ImportSecretModeRequired))
}

func TestPropertyConfiguration_Secrecy_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	secrecy, ok := property.Secrecy.Lookup()

	g.Expect(secrecy).To(Equal(astmodel.ImportSecretMode("")))
	g.Expect(ok).To(BeFalse())
}
