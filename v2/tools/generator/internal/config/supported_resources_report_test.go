/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestSupportedResourcesReport_Merge_WhenNil_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := NewConfiguration()
	base := NewSupportedResourcesReport(cfg)
	err := base.Merge(nil)
	g.Expect(err).To(Succeed())
}

func TestSupportedResourcesReport_Merge_WhenEmpty_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := NewConfiguration()
	base := NewSupportedResourcesReport(cfg)
	other := NewSupportedResourcesReport(cfg)
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
}

func TestSupportedResourcesReport_Merge_WhenNoConflicts_MergesSuccessfully(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := NewConfiguration()
	base := NewSupportedResourcesReport(cfg)
	base.OutputFolder = "docs/reports"
	base.FragmentPath = "fragments"
	
	other := NewSupportedResourcesReport(cfg)
	other.ResourceURLTemplate = "https://docs.microsoft.com/{group}/{version}/{kind}"
	other.CurrentRelease = "v2.5.0"
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
	g.Expect(base.OutputFolder).To(Equal("docs/reports"))
	g.Expect(base.FragmentPath).To(Equal("fragments"))
	g.Expect(base.ResourceURLTemplate).To(Equal("https://docs.microsoft.com/{group}/{version}/{kind}"))
	g.Expect(base.CurrentRelease).To(Equal("v2.5.0"))
}

func TestSupportedResourcesReport_Merge_WhenConflicts_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := NewConfiguration()
	base := NewSupportedResourcesReport(cfg)
	base.OutputFolder = "docs/reports"
	
	other := NewSupportedResourcesReport(cfg)
	other.OutputFolder = "other/reports"
	
	err := base.Merge(other)
	g.Expect(err).To(MatchError(ContainSubstring("conflict in field OutputFolder")))
	g.Expect(err).To(MatchError(ContainSubstring("base value \"docs/reports\" cannot be overwritten with \"other/reports\"")))
}

func TestSupportedResourcesReport_Merge_PreservesWhenOtherIsEmpty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := NewConfiguration()
	base := NewSupportedResourcesReport(cfg)
	base.FragmentPath = "base/fragments"
	base.ResourcePathTemplate = "docs/{group}/{version}/{kind}.md"
	
	other := NewSupportedResourcesReport(cfg)
	// other has empty fields
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
	g.Expect(base.FragmentPath).To(Equal("base/fragments"))
	g.Expect(base.ResourcePathTemplate).To(Equal("docs/{group}/{version}/{kind}.md"))
}

func TestSupportedResourcesReport_Merge_MultipleFields_WhenConflicts_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := NewConfiguration()
	base := NewSupportedResourcesReport(cfg)
	base.ResourceURLTemplate = "https://example.com/{group}"
	base.CurrentRelease = "v2.4.0"
	
	other := NewSupportedResourcesReport(cfg)
	other.ResourceURLTemplate = "https://different.com/{group}"
	other.CurrentRelease = "v2.5.0"
	
	err := base.Merge(other)
	g.Expect(err).To(HaveOccurred())
	// Should fail on the first conflict (ResourceURLTemplate)
	g.Expect(err).To(MatchError(ContainSubstring("conflict in field ResourceURLTemplate")))
}
