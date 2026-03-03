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

func TestUpgradableResourcesReport_WriteTo_WithItems(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := config.NewUpgradableResourcesReport(config.NewConfiguration())
	report := &UpgradableResourcesReport{
		cfg: cfg,
		items: []upgradableResourcesReportItem{
			{
				group:           "storage",
				resource:        "StorageAccount",
				supportedStable: "v20230101",
				availableStable: "v20250601",
			},
			{
				group:            "compute",
				resource:         "VirtualMachine",
				supportedPreview: "v1api20211101preview",
				availablePreview: "v1api20250401preview",
			},
		},
	}

	var buf strings.Builder
	report.WriteTo(&buf)

	output := buf.String()
	g.Expect(output).To(ContainSubstring("# Upgradable Resources"))
	g.Expect(output).To(ContainSubstring("StorageAccount"))
	g.Expect(output).To(ContainSubstring("VirtualMachine"))
	g.Expect(output).To(ContainSubstring("v20230101"))
	g.Expect(output).To(ContainSubstring("v20250601"))
	g.Expect(output).To(ContainSubstring("compute"))
	g.Expect(output).To(ContainSubstring("storage"))
}

func TestNewUpgradableResourcesReport_RecommendStableUpgrade(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// storage is a hybrid group so version prefix is "v" (not "v1api")
	// Supported: v20230101 (2023-01-01)
	// Available: v20250601 (2025-06-01) — more than 12 months newer
	storagePkg2023 := test.MakeLocalPackageReference("storage", "v20230101")
	storagePkg2025 := test.MakeLocalPackageReference("storage", "v20250601")

	// Supported definitions contain only the 2023 version
	supported := make(astmodel.TypeDefinitionSet)
	supported.Add(makeTestResource(storagePkg2023, "StorageAccount"))

	// AllKnownResources contains both versions
	allKnown := makeAllKnownResources(
		astmodel.MakeInternalTypeName(storagePkg2023, "StorageAccount"),
		astmodel.MakeInternalTypeName(storagePkg2025, "StorageAccount"),
	)

	cfg := config.NewUpgradableResourcesReport(config.NewConfiguration())
	report := NewUpgradableResourcesReport(allKnown, supported, cfg)

	g.Expect(report.items).To(HaveLen(1))
	g.Expect(report.items[0].resource).To(Equal("StorageAccount"))
	g.Expect(report.items[0].supportedStable).To(Equal(storagePkg2023.PackageName()))
	g.Expect(report.items[0].availableStable).To(Equal(storagePkg2025.PackageName()))
}

func TestNewUpgradableResourcesReport_NoUpgradeNeeded(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Use dates relative to now so the test remains correct regardless of when it runs.
	// supported is 11 months ago, available is 3 months newer → gap (3 months) is below the
	// 12-month stableVersionsExpiry threshold, and supported is < 24 months old.
	now := time.Now()
	supportedDate := now.AddDate(0, -11, 0)
	availableDate := now.AddDate(0, -8, 0) // 3 months newer than supported

	supportedVersion := supportedDate.Format("v20060102")
	availableVersion := availableDate.Format("v20060102")

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

	g.Expect(report.items).To(BeEmpty())
}

func TestNewUpgradableResourcesReport_RecommendPreviewUpgrade(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// containerservice is a legacy group so version prefix is "v1api".
	// The argument to MakeLocalPackageReference is the base API version (without the group prefix);
	// the resulting PackageName() will be "v1api20220101preview" and "v1api20220901preview".
	csPkg2022 := test.MakeLocalPackageReference("containerservice", "20220101preview")
	csPkg2022newer := test.MakeLocalPackageReference("containerservice", "20220901preview")

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
	g.Expect(report.items[0].supportedPreview).To(Equal(csPkg2022.PackageName()))
	g.Expect(report.items[0].availablePreview).To(Equal(csPkg2022newer.PackageName()))
}

func TestNewUpgradableResourcesReport_SortsByGroupThenResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Old-enough packages to be above the latestVersionThreshold (> 24 months)
	batchPkg2020 := test.MakeLocalPackageReference("batch", "v20200101")
	batchPkg2025 := test.MakeLocalPackageReference("batch", "v20250101")
	storagePkg2020 := test.MakeLocalPackageReference("storage", "v20200101")
	storagePkg2025 := test.MakeLocalPackageReference("storage", "v20250101")

	supported := make(astmodel.TypeDefinitionSet)
	supported.Add(makeTestResource(batchPkg2020, "BatchAccount"))
	supported.Add(makeTestResource(storagePkg2020, "StorageAccount"))
	supported.Add(makeTestResource(storagePkg2020, "BlobContainer"))

	allKnown := makeAllKnownResources(
		astmodel.MakeInternalTypeName(batchPkg2020, "BatchAccount"),
		astmodel.MakeInternalTypeName(batchPkg2025, "BatchAccount"),
		astmodel.MakeInternalTypeName(storagePkg2020, "StorageAccount"),
		astmodel.MakeInternalTypeName(storagePkg2025, "StorageAccount"),
		astmodel.MakeInternalTypeName(storagePkg2020, "BlobContainer"),
		astmodel.MakeInternalTypeName(storagePkg2025, "BlobContainer"),
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
	r := &UpgradableResourcesReport{cfg: cfg}

	// Supported version is more than 24 months old; available is only 6 months newer.
	// Should still recommend because the supported version is > 24 months old.
	now := time.Now()
	oldDate := now.AddDate(-3, 0, 0) // 3 years ago
	supported := oldDate.Format("v20060102")
	available := oldDate.AddDate(0, 6, 0).Format("v20060102")

	g.Expect(r.isStableUpgradeRecommended(supported, available, now)).To(BeTrue())
}

func TestIsPreviewUpgradeRecommended_NewerThanThreshold(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := config.NewConfiguration()
	cfg := config.NewUpgradableResourcesReport(c)
	r := &UpgradableResourcesReport{cfg: cfg}

	// Available is 8 months newer than supported (above 6-month threshold)
	supported := "v1api20220101preview"
	available := "v1api20220901preview"

	g.Expect(r.isPreviewUpgradeRecommended(supported, available)).To(BeTrue())
}

func TestIsPreviewUpgradeRecommended_NoCurrentSupport(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := config.NewConfiguration()
	cfg := config.NewUpgradableResourcesReport(c)
	r := &UpgradableResourcesReport{cfg: cfg}

	// No supported preview version, available is 2023
	g.Expect(r.isPreviewUpgradeRecommended("", "v20230601preview")).To(BeTrue())
}

func TestParseASOVersionDate(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	d, err := parseASOVersionDate("v20230101")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(d.Year()).To(Equal(2023))

	d, err = parseASOVersionDate("v1api20230615preview")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(d.Month()).To(Equal(time.June))
	g.Expect(d.Day()).To(Equal(15))

	_, err = parseASOVersionDate("")
	g.Expect(err).To(HaveOccurred())

	_, err = parseASOVersionDate("customizations")
	g.Expect(err).To(HaveOccurred())
}

func TestMonthsBetween(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

	g.Expect(monthsBetween(from, to)).To(Equal(29))
}

