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
