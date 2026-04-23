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

func TestGroupConfigurationFile_MergesTypeFilters_Prepended(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := NewConfiguration()
	globalFilter := &TypeFilter{Action: TypeFilterPrune}
	cfg.TypeFilters = []*TypeFilter{globalFilter}

	gcf := &GroupConfigurationFile{
		TypeFilters: []*TypeFilter{
			{Action: TypeFilterInclude},
		},
	}

	err := gcf.mergeInto(cfg, "testgroup")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cfg.TypeFilters).To(HaveLen(2))
	// Group filter should be prepended (first)
	g.Expect(cfg.TypeFilters[0].Action).To(Equal(TypeFilterInclude))
	// Global filter should be second
	g.Expect(cfg.TypeFilters[1].Action).To(Equal(TypeFilterPrune))
}

func TestGroupConfigurationFile_MergesTransformers_Appended(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := NewConfiguration()
	globalTransformer := &TypeTransformer{}
	globalTransformer.RenameTo = "GlobalRename"
	cfg.Transformers = []*TypeTransformer{globalTransformer}

	gcf := &GroupConfigurationFile{
		Transformers: []*TypeTransformer{
			{RenameTo: "GroupRename"},
		},
	}

	err := gcf.mergeInto(cfg, "testgroup")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cfg.Transformers).To(HaveLen(2))
	g.Expect(cfg.Transformers[0].RenameTo).To(Equal("GlobalRename"))
	g.Expect(cfg.Transformers[1].RenameTo).To(Equal("GroupRename"))
}

func TestGroupConfigurationFile_MergesGroupModelConfiguration(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := NewConfiguration()
	gc := NewGroupConfiguration("testgroup")
	gcf := &GroupConfigurationFile{
		GroupModelConfiguration: gc,
	}

	err := gcf.mergeInto(cfg, "testgroup")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cfg.ObjectModelConfiguration.groups).To(HaveKey("testgroup"))
}

func TestGroupConfigurationFile_DuplicateGroup_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := NewConfiguration()
	existing := NewGroupConfiguration("testgroup")
	err := cfg.ObjectModelConfiguration.addGroup("testgroup", existing)
	g.Expect(err).ToNot(HaveOccurred())

	gc := NewGroupConfiguration("testgroup")
	gcf := &GroupConfigurationFile{
		GroupModelConfiguration: gc,
	}

	err = gcf.mergeInto(cfg, "testgroup")
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("duplicate"))
}
