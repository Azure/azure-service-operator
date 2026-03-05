/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

// makeTestResource creates a simple resource TypeDefinition suitable for use in upgradable report tests.
func makeTestResource(pkg astmodel.InternalPackageReference, name string) astmodel.TypeDefinition {
	spec := test.CreateSpec(pkg, name)
	status := test.CreateStatus(pkg, name)
	return test.CreateResource(pkg, name, spec, status)
}

// makeAllKnownResources builds a map[group]TypeNameSet in the format expected by NewUpgradableResourcesReport.
func makeAllKnownResources(names ...astmodel.InternalTypeName) map[string]astmodel.TypeNameSet {
	result := make(map[string]astmodel.TypeNameSet)
	for _, name := range names {
		group := name.InternalPackageReference().Group()
		if _, ok := result[group]; !ok {
			result[group] = make(astmodel.TypeNameSet)
		}
		result[group].Add(name)
	}
	return result
}

// TestGolden_ReportUpgradableResources verifies the markdown output of the upgradable resources report.
func TestGolden_ReportUpgradableResources(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	gold := goldie.New(t)

	// storage is a hybrid group (prefix "v")
	// Supported: 2020-01-01 stable. Available: 2023-06-01 stable (42 months newer) → recommended (bold)
	storagePkg2020 := test.MakeLocalPackageReference("storage", "2020-01-01")
	storagePkg2023 := test.MakeLocalPackageReference("storage", "2023-06-01")

	// compute is a legacy group (prefix "v1api")
	// Supported: 2020-09-01 stable. Available: 2021-03-01 stable (6 months newer, not enough for 12-month threshold)
	// But supported is also > 24 months old → still recommended (bold)
	computePkg2020 := test.MakeLocalPackageReference("compute", "2020-09-01")
	computePkg2021 := test.MakeLocalPackageReference("compute", "2021-03-01")

	// containerservice is a legacy group (prefix "v1api")
	// Supported preview: 2022-01-01-preview. Available preview: 2022-09-01-preview (8 months newer > 6-month threshold) → recommended (bold)
	csPkg2022 := test.MakeLocalPackageReference("containerservice", "2022-01-01-preview")
	csPkg2022newer := test.MakeLocalPackageReference("containerservice", "2022-09-01-preview")

	// batch is a legacy group (prefix "v1api")
	// Supported: 2021-01-01 stable. Available newer: 2021-06-01 (5 months, below threshold).
	// But supported is > 24 months old from now (2026-03-03) → still recommended (bold)
	batchPkg2021 := test.MakeLocalPackageReference("batch", "2021-01-01")
	batchPkg2021newer := test.MakeLocalPackageReference("batch", "2021-06-01")

	supported := make(astmodel.TypeDefinitionSet)
	supported.Add(makeTestResource(storagePkg2020, "StorageAccount"))
	supported.Add(makeTestResource(computePkg2020, "VirtualMachine"))
	supported.Add(makeTestResource(csPkg2022, "ManagedCluster"))
	supported.Add(makeTestResource(batchPkg2021, "BatchAccount"))

	allKnown := makeAllKnownResources(
		astmodel.MakeInternalTypeName(storagePkg2020, "StorageAccount"),
		astmodel.MakeInternalTypeName(storagePkg2023, "StorageAccount"),
		astmodel.MakeInternalTypeName(computePkg2020, "VirtualMachine"),
		astmodel.MakeInternalTypeName(computePkg2021, "VirtualMachine"),
		astmodel.MakeInternalTypeName(csPkg2022, "ManagedCluster"),
		astmodel.MakeInternalTypeName(csPkg2022newer, "ManagedCluster"),
		astmodel.MakeInternalTypeName(batchPkg2021, "BatchAccount"),
		astmodel.MakeInternalTypeName(batchPkg2021newer, "BatchAccount"),
	)

	cfg := config.NewUpgradableResourcesReport(config.NewConfiguration())
	report := NewUpgradableResourcesReport(allKnown, supported, cfg)

	var buffer strings.Builder
	// Use a fixed time so the golden file is deterministic.
	// 2026-03-03 means all supported versions from 2020-2022 are > 24 months old.
	report.writeTo(&buffer, time.Date(2026, 3, 3, 0, 0, 0, 0, time.UTC))

	gold.Assert(t, t.Name(), []byte(buffer.String()))
	g.Expect(buffer.String()).To(ContainSubstring("# Upgradable Resources"))
}

func TestUpgradableResourcesReport_WriteTo_NoItems(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := config.NewUpgradableResourcesReport(config.NewConfiguration())
	report := &UpgradableResourcesReport{cfg: cfg}

	var buf strings.Builder
	report.WriteTo(&buf)

	output := buf.String()
	g.Expect(output).To(ContainSubstring("# Upgradable Resources"))
	g.Expect(output).To(ContainSubstring("No resources with available upgrades"))
}

func TestUpgradableResourcesReport_WriteTo_WithItems_BoldsRecommended(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	cfg := config.NewUpgradableResourcesReport(config.NewConfiguration())
	report := &UpgradableResourcesReport{
		cfg: cfg,
		items: []upgradableResourcesReportItem{
			{
				// Recommended: available is >12 months newer
				group:    "storage",
				resource: "StorageAccount",
				supportedStable: versionEntry{pkgName: "v20230101", armVersion: "2023-01-01"},
				availableStable: versionEntry{pkgName: "v20250601", armVersion: "2025-06-01"},
			},
			{
				// Not recommended: supported preview is only 2 months old relative to available
				group:    "compute",
				resource: "VirtualMachine",
				supportedPreview: versionEntry{pkgName: "v1api20250101preview", armVersion: "2025-11-01-preview"},
				availablePreview: versionEntry{pkgName: "v1api20260101preview", armVersion: "2026-01-01-preview"},
			},
		},
	}

	var buf strings.Builder
	report.writeTo(&buf, now)

	output := buf.String()
	g.Expect(output).To(ContainSubstring("# Upgradable Resources"))
	g.Expect(output).To(ContainSubstring("StorageAccount"))
	g.Expect(output).To(ContainSubstring("VirtualMachine"))
	// StorageAccount stable upgrade should be bolded (recommended)
	g.Expect(output).To(ContainSubstring("**v20250601**"))
	// VirtualMachine preview upgrade should NOT be bolded (only 2 months gap, below 6-month threshold)
	g.Expect(output).To(ContainSubstring("v1api20260101preview"))
	g.Expect(output).ToNot(ContainSubstring("**v1api20260101preview**"))
}

func TestNewUpgradableResourcesReport_RecommendStableUpgrade(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// storage is a hybrid group so version prefix is "v" (not "v1api")
	// Supported: 2023-01-01 stable. Available: 2025-06-01 stable — more than 12 months newer.
	storagePkg2023 := test.MakeLocalPackageReference("storage", "2023-01-01")
	storagePkg2025 := test.MakeLocalPackageReference("storage", "2025-06-01")

	supported := make(astmodel.TypeDefinitionSet)
	supported.Add(makeTestResource(storagePkg2023, "StorageAccount"))

	allKnown := makeAllKnownResources(
		astmodel.MakeInternalTypeName(storagePkg2023, "StorageAccount"),
		astmodel.MakeInternalTypeName(storagePkg2025, "StorageAccount"),
	)

	cfg := config.NewUpgradableResourcesReport(config.NewConfiguration())
	report := NewUpgradableResourcesReport(allKnown, supported, cfg)

	g.Expect(report.items).To(HaveLen(1))
	g.Expect(report.items[0].resource).To(Equal("StorageAccount"))
	g.Expect(report.items[0].supportedStable.pkgName).To(Equal(storagePkg2023.PackageName()))
	g.Expect(report.items[0].availableStable.pkgName).To(Equal(storagePkg2025.PackageName()))
}

func TestNewUpgradableResourcesReport_ListsResourcesEvenWhenNotRecommended(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// supported is 11 months ago, available is 3 months newer → gap (3 months) is below the
	// 12-month stableVersionsExpiry threshold, and supported is < 24 months old.
	// The resource should still be LISTED because a newer version exists.
	now := time.Now()
	supportedDate := now.AddDate(0, -11, 0)
	availableDate := now.AddDate(0, -8, 0) // 3 months newer than supported

	supportedVersion := supportedDate.Format("2006-01-02")
	availableVersion := availableDate.Format("2006-01-02")

	storagePkgSupported := test.MakeLocalPackageReference("storage", supportedVersion)
	storagePkgAvailable := test.MakeLocalPackageReference("storage", availableVersion)

	supported := make(astmodel.TypeDefinitionSet)
	supported.Add(makeTestResource(storagePkgSupported, "StorageAccount"))

	allKnown := makeAllKnownResources(
		astmodel.MakeInternalTypeName(storagePkgSupported, "StorageAccount"),
		astmodel.MakeInternalTypeName(storagePkgAvailable, "StorageAccount"),
	)

	cfg := config.NewUpgradableResourcesReport(config.NewConfiguration())
	report := NewUpgradableResourcesReport(allKnown, supported, cfg)

	// The resource has a newer version so it should be listed
	g.Expect(report.items).To(HaveLen(1))

	// But the upgrade should not be recommended (gap < 12 months and version < 24 months old)
	item := report.items[0]
	g.Expect(item.isStableUpgradeRecommended(cfg, now)).To(BeFalse())
}

func TestNewUpgradableResourcesReport_NoUpgradeAvailable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Only one version available – no upgrade possible
	storagePkg2023 := test.MakeLocalPackageReference("storage", "2023-01-01")

	supported := make(astmodel.TypeDefinitionSet)
	supported.Add(makeTestResource(storagePkg2023, "StorageAccount"))

	allKnown := makeAllKnownResources(
		astmodel.MakeInternalTypeName(storagePkg2023, "StorageAccount"),
	)

	cfg := config.NewUpgradableResourcesReport(config.NewConfiguration())
	report := NewUpgradableResourcesReport(allKnown, supported, cfg)

	g.Expect(report.items).To(BeEmpty())
}

func TestNewUpgradableResourcesReport_RecommendPreviewUpgrade(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// containerservice is a legacy group so version prefix is "v1api".
	// Supported preview: 2022-01-01-preview. Available: 2022-09-01-preview (8 months newer > 6-month threshold).
	csPkg2022 := test.MakeLocalPackageReference("containerservice", "2022-01-01-preview")
	csPkg2022newer := test.MakeLocalPackageReference("containerservice", "2022-09-01-preview")

	supported := make(astmodel.TypeDefinitionSet)
	supported.Add(makeTestResource(csPkg2022, "ManagedCluster"))

	allKnown := makeAllKnownResources(
		astmodel.MakeInternalTypeName(csPkg2022, "ManagedCluster"),
		astmodel.MakeInternalTypeName(csPkg2022newer, "ManagedCluster"),
	)

	cfg := config.NewUpgradableResourcesReport(config.NewConfiguration())
	report := NewUpgradableResourcesReport(allKnown, supported, cfg)

	g.Expect(report.items).To(HaveLen(1))
	g.Expect(report.items[0].resource).To(Equal("ManagedCluster"))
	g.Expect(report.items[0].supportedPreview.pkgName).To(Equal(csPkg2022.PackageName()))
	g.Expect(report.items[0].availablePreview.pkgName).To(Equal(csPkg2022newer.PackageName()))
}

func TestNewUpgradableResourcesReport_SortsByGroupThenResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// All versions old enough (> 24 months from any reasonable test date) to be recommended
	batchPkg2020 := test.MakeLocalPackageReference("batch", "2020-01-01")
	batchPkg2021 := test.MakeLocalPackageReference("batch", "2021-01-01")
	storagePkg2020 := test.MakeLocalPackageReference("storage", "2020-01-01")
	storagePkg2021 := test.MakeLocalPackageReference("storage", "2021-01-01")

	supported := make(astmodel.TypeDefinitionSet)
	supported.Add(makeTestResource(batchPkg2020, "BatchAccount"))
	supported.Add(makeTestResource(storagePkg2020, "StorageAccount"))
	supported.Add(makeTestResource(storagePkg2020, "BlobContainer"))

	allKnown := makeAllKnownResources(
		astmodel.MakeInternalTypeName(batchPkg2020, "BatchAccount"),
		astmodel.MakeInternalTypeName(batchPkg2021, "BatchAccount"),
		astmodel.MakeInternalTypeName(storagePkg2020, "StorageAccount"),
		astmodel.MakeInternalTypeName(storagePkg2021, "StorageAccount"),
		astmodel.MakeInternalTypeName(storagePkg2020, "BlobContainer"),
		astmodel.MakeInternalTypeName(storagePkg2021, "BlobContainer"),
	)

	cfg := config.NewUpgradableResourcesReport(config.NewConfiguration())
	report := NewUpgradableResourcesReport(allKnown, supported, cfg)

	g.Expect(report.items).To(HaveLen(3))
	// Sorted: batch/BatchAccount, storage/BlobContainer, storage/StorageAccount
	g.Expect(report.items[0].group).To(Equal("batch"))
	g.Expect(report.items[0].resource).To(Equal("BatchAccount"))
	g.Expect(report.items[1].group).To(Equal("storage"))
	g.Expect(report.items[1].resource).To(Equal("BlobContainer"))
	g.Expect(report.items[2].group).To(Equal("storage"))
	g.Expect(report.items[2].resource).To(Equal("StorageAccount"))
}

func TestIsStableUpgradeRecommended_OldSupportedVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := config.NewConfiguration()
	cfg := config.NewUpgradableResourcesReport(c)

	// Supported version is more than 24 months old; available is only 6 months newer.
	// Should still recommend because the supported version is > 24 months old.
	now := time.Now()
	oldDate := now.AddDate(-3, 0, 0) // 3 years ago
	item := upgradableResourcesReportItem{
		supportedStable: versionEntry{pkgName: "v" + oldDate.Format("20060102"), armVersion: oldDate.Format("2006-01-02")},
		availableStable: versionEntry{pkgName: "v" + oldDate.AddDate(0, 6, 0).Format("20060102"), armVersion: oldDate.AddDate(0, 6, 0).Format("2006-01-02")},
	}

	g.Expect(item.isStableUpgradeRecommended(cfg, now)).To(BeTrue())
}

func TestIsPreviewUpgradeRecommended_NewerThanThreshold(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := config.NewConfiguration()
	cfg := config.NewUpgradableResourcesReport(c)

	// Available is 8 months newer than supported (above 6-month threshold)
	item := upgradableResourcesReportItem{
		supportedPreview: versionEntry{pkgName: "v1api20220101preview", armVersion: "2022-01-01-preview"},
		availablePreview: versionEntry{pkgName: "v1api20220901preview", armVersion: "2022-09-01-preview"},
	}

	g.Expect(item.isPreviewUpgradeRecommended(cfg)).To(BeTrue())
}

func TestIsPreviewUpgradeRecommended_NoCurrentSupport(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := config.NewConfiguration()
	cfg := config.NewUpgradableResourcesReport(c)

	// No supported preview version, available is 2023
	item := upgradableResourcesReportItem{
		// supportedPreview is zero-value (empty)
		availablePreview: versionEntry{pkgName: "v20230601preview", armVersion: "2023-06-01-preview"},
	}

	g.Expect(item.isPreviewUpgradeRecommended(cfg)).To(BeTrue())
}

func TestParseARMVersionDate(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	d, err := parseARMVersionDate("2023-01-01")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(d.Year()).To(Equal(2023))

	d, err = parseARMVersionDate("2023-06-15-preview")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(d.Month()).To(Equal(time.June))
	g.Expect(d.Day()).To(Equal(15))

	_, err = parseARMVersionDate("")
	g.Expect(err).To(HaveOccurred())

	_, err = parseARMVersionDate("v20230101")
	g.Expect(err).To(HaveOccurred())
}

func TestMonthsBetween(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

	g.Expect(monthsBetween(from, to)).To(Equal(29))
}

