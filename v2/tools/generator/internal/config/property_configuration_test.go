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
	g.Expect(property.renamedTo).To(Equal("DemoProperty"))
	g.Expect(property.haveArmReference).To(BeTrue())
	g.Expect(property.armReference).To(BeTrue())
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

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).To(Succeed())

	isReference, ok := property.ARMReference()
	g.Expect(ok).To(BeTrue())
	g.Expect(isReference).To(BeTrue())
}

func TestPropertyConfiguration_ARMReference_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	g := NewGomegaWithT(t)

	yamlBytes := loadTestData(t)

	var property PropertyConfiguration
	err := yaml.Unmarshal(yamlBytes, &property)
	g.Expect(err).To(Succeed())

	_, ok := property.ARMReference()
	g.Expect(ok).To(BeFalse())
}
