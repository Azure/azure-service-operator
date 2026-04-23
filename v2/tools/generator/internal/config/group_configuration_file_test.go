/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGroupConfigurationFile_WhenYAMLWellFormed_LoadsCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gcf, err := loadGroupConfigurationFile(
		"testdata/TestGroupConfigurationFile/WellFormed.yaml",
		"testgroup",
	)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(gcf).ToNot(BeNil())
	g.Expect(gcf.TypeFilters).To(HaveLen(1))
	g.Expect(gcf.Transformers).To(HaveLen(1))
	g.Expect(gcf.GroupModelConfiguration).ToNot(BeNil())
	g.Expect(gcf.GroupModelConfiguration.name).To(Equal("testgroup"))
}

func TestGroupConfigurationFile_WhenMinimal_LoadsCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gcf, err := loadGroupConfigurationFile(
		"testdata/TestGroupConfigurationFile/WellFormed_MinimalNoGroupModel.yaml",
		"testgroup",
	)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(gcf.TypeFilters).To(HaveLen(1))
	g.Expect(gcf.Transformers).To(BeEmpty())
	g.Expect(gcf.GroupModelConfiguration).To(BeNil())
}

func TestGroupConfigurationFile_WhenGroupModelOnly_LoadsCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gcf, err := loadGroupConfigurationFile(
		"testdata/TestGroupConfigurationFile/WellFormed_GroupModelOnly.yaml",
		"testgroup",
	)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(gcf.TypeFilters).To(BeEmpty())
	g.Expect(gcf.GroupModelConfiguration).ToNot(BeNil())
}

func TestGroupConfigurationFile_WhenFilterHasWrongGroup_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	_, err := loadGroupConfigurationFile(
		"testdata/TestGroupConfigurationFile/WrongGroup_Filter.yaml",
		"testgroup",
	)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("wronggroup"))
}

func TestGroupConfigurationFile_WhenTransformerHasWrongGroup_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	_, err := loadGroupConfigurationFile(
		"testdata/TestGroupConfigurationFile/WrongGroup_Transformer.yaml",
		"testgroup",
	)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("wronggroup"))
}

func TestGroupConfigurationFile_WhenFilterHasMatchingGroup_Succeeds(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gcf, err := loadGroupConfigurationFile(
		"testdata/TestGroupConfigurationFile/MatchingGroup_Filter.yaml",
		"testgroup",
	)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(gcf.TypeFilters).To(HaveLen(1))
}

func TestGroupConfigurationFile_WhenFilterGroupOmitted_AutoFillsGroup(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gcf, err := loadGroupConfigurationFile(
		"testdata/TestGroupConfigurationFile/OmittedGroup_Filter.yaml",
		"testgroup",
	)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(gcf.TypeFilters).To(HaveLen(1))
	g.Expect(gcf.TypeFilters[0].Group.IsRestrictive()).To(BeTrue())
}
