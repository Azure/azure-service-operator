/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

// makeARMResource creates a resource definition with the given ARM type and API version for tests.
func makeARMResource(
	pkg astmodel.InternalPackageReference,
	name string,
	armType string,
	armVersion string,
) astmodel.TypeDefinition {
	spec := test.CreateSpec(pkg, name)
	status := test.CreateStatus(pkg, name)
	resourceType := astmodel.NewResourceType(spec.Name(), status.Name())
	if armType != "" {
		resourceType = resourceType.WithARMType(armType)
	}
	if armVersion != "" {
		enumType := astmodel.NewEnumType(astmodel.StringType, astmodel.MakeEnumValue("v", `"`+armVersion+`"`))
		apiVersionName := astmodel.MakeInternalTypeName(pkg, "APIVersion")
		apiVersionDef := astmodel.MakeTypeDefinition(apiVersionName, enumType)
		resourceType = resourceType.WithAPIVersion(apiVersionDef.Name(), enumType.Options()[0])
	}
	return astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(pkg, name), resourceType)
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
				supportedStable: "2023-01-01",
				availableStable: "2025-06-01",
			},
			{
				group:            "compute",
				resource:         "VirtualMachine",
				supportedPreview: "2021-11-01-preview",
				availablePreview: "2025-04-01-preview",
			},
		},
	}

	var buf strings.Builder
	report.WriteTo(&buf)

	output := buf.String()
	g.Expect(output).To(ContainSubstring("# Upgradable Resources"))
	g.Expect(output).To(ContainSubstring("StorageAccount"))
	g.Expect(output).To(ContainSubstring("VirtualMachine"))
	g.Expect(output).To(ContainSubstring("2023-01-01"))
	g.Expect(output).To(ContainSubstring("2025-06-01"))
	g.Expect(output).To(ContainSubstring("compute"))
	g.Expect(output).To(ContainSubstring("storage"))
}

func TestScanSpecVersions_WithTempDir(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create a temporary spec directory structure
	tmpDir := t.TempDir()

	// Create stable versions for microsoft.storage
	mkdirAll(t, filepath.Join(tmpDir, "storage", "resource-manager", "Microsoft.Storage", "stable", "2023-01-01"))
	mkdirAll(t, filepath.Join(tmpDir, "storage", "resource-manager", "Microsoft.Storage", "stable", "2025-06-01"))

	// Create preview versions for microsoft.storage
	mkdirAll(t, filepath.Join(tmpDir, "storage", "resource-manager", "Microsoft.Storage", "preview", "2020-08-01-preview"))

	// Create stable versions for microsoft.compute
	mkdirAll(t, filepath.Join(tmpDir, "compute", "resource-manager", "Microsoft.Compute", "ComputeRP", "stable", "2023-09-01"))
	mkdirAll(t, filepath.Join(tmpDir, "compute", "resource-manager", "Microsoft.Compute", "ComputeRP", "stable", "2025-04-01"))

	versions, err := scanSpecVersions(tmpDir)
	g.Expect(err).ToNot(HaveOccurred())

	storageVers, ok := versions["microsoft.storage"]
	g.Expect(ok).To(BeTrue())
	g.Expect(storageVers.latestStable).To(Equal("2025-06-01"))
	g.Expect(storageVers.latestPreview).To(Equal("2020-08-01-preview"))

	computeVers, ok := versions["microsoft.compute"]
	g.Expect(ok).To(BeTrue())
	g.Expect(computeVers.latestStable).To(Equal("2025-04-01"))
}

func TestScanSpecVersions_NonexistentDir(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	versions, err := scanSpecVersions("/nonexistent/path/that/does/not/exist")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(versions).To(BeEmpty())
}

func TestNewUpgradableResourcesReport_RecommendStableUpgrade(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	tmpDir := t.TempDir()
	// Stable version much newer (>12 months)
	mkdirAll(t, filepath.Join(tmpDir, "storage", "resource-manager", "Microsoft.Storage", "stable", "2025-06-01"))

	// Create a resource at 2023-01-01 stable (more than 12 months older than 2025-06-01)
	storagePkg := test.MakeLocalPackageReference("storage", "v20230101")
	storageAccount := makeARMResource(storagePkg, "StorageAccount", "Microsoft.Storage/storageAccounts", "2023-01-01")

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(storageAccount)

	c := config.NewConfiguration()
	c.SchemaRoot = tmpDir
	upgradeCfg := config.NewUpgradableResourcesReport(c)
	c.UpgradableResourcesReport = upgradeCfg

	report, err := NewUpgradableResourcesReport(defs, c)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(report.items).To(HaveLen(1))
	g.Expect(report.items[0].resource).To(Equal("StorageAccount"))
	g.Expect(report.items[0].supportedStable).To(Equal("2023-01-01"))
	g.Expect(report.items[0].availableStable).To(Equal("2025-06-01"))
}

func TestNewUpgradableResourcesReport_NoUpgradeNeeded(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Use dates relative to now so the test remains correct regardless of when it runs.
	// We want: supported is < 24 months old, and available is only 3 months newer than supported.
	// This means neither the stableVersionsExpiry (12 months) nor latestVersionThreshold (24 months)
	// condition is satisfied.
	now := time.Now()
	supported := now.AddDate(0, -11, 0).Format("2006-01-02")
	available := now.AddDate(0, -8, 0).Format("2006-01-02") // 3 months newer

	tmpDir := t.TempDir()
	mkdirAll(t, filepath.Join(tmpDir, "storage", "resource-manager", "Microsoft.Storage", "stable", available))

	storagePkg := test.MakeLocalPackageReference("storage", "v20230101")
	storageAccount := makeARMResource(storagePkg, "StorageAccount", "Microsoft.Storage/storageAccounts", supported)

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(storageAccount)

	c := config.NewConfiguration()
	c.SchemaRoot = tmpDir
	c.UpgradableResourcesReport = config.NewUpgradableResourcesReport(c)

	report, err := NewUpgradableResourcesReport(defs, c)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(report.items).To(BeEmpty())
}

func TestIsStableUpgradeRecommended_OldSupportedVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := config.NewConfiguration()
	cfg := config.NewUpgradableResourcesReport(c)
	r := &UpgradableResourcesReport{cfg: cfg}

	// Supported version is more than 24 months old
	now := time.Now()
	oldDate := now.AddDate(-3, 0, 0) // 3 years ago
	supported := oldDate.Format("2006-01-02")
	// Available is only 6 months newer than supported (less than 12 months)
	available := oldDate.AddDate(0, 6, 0).Format("2006-01-02")

	// Should still recommend because the supported version is > 24 months old
	g.Expect(r.isStableUpgradeRecommended(supported, available, now)).To(BeTrue())
}

func TestIsPreviewUpgradeRecommended_NewerThanThreshold(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := config.NewConfiguration()
	cfg := config.NewUpgradableResourcesReport(c)
	r := &UpgradableResourcesReport{cfg: cfg}

	// Available is 8 months newer than supported (above 6 month threshold)
	supported := "2022-01-01-preview"
	available := "2022-09-01-preview"

	g.Expect(r.isPreviewUpgradeRecommended(supported, available)).To(BeTrue())
}

func TestIsPreviewUpgradeRecommended_NoCurrentSupport(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	c := config.NewConfiguration()
	cfg := config.NewUpgradableResourcesReport(c)
	r := &UpgradableResourcesReport{cfg: cfg}

	// No supported preview version, available is 2023
	g.Expect(r.isPreviewUpgradeRecommended("", "2023-06-01-preview")).To(BeTrue())
}

func TestArmTypeProvider(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(armTypeProvider("Microsoft.Storage/storageAccounts")).To(Equal("Microsoft.Storage"))
	g.Expect(armTypeProvider("Microsoft.Compute/virtualMachines")).To(Equal("Microsoft.Compute"))
	g.Expect(armTypeProvider("Microsoft.Network")).To(Equal("Microsoft.Network"))
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

	_, err = parseARMVersionDate("")
	g.Expect(err).To(HaveOccurred())
}

func TestMonthsBetween(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

	g.Expect(monthsBetween(from, to)).To(Equal(29))
}

// mkdirAll is a helper that creates a directory and fails the test on error.
func mkdirAll(t *testing.T, path string) {
	t.Helper()
	err := os.MkdirAll(path, 0o700)
	if err != nil {
		t.Fatalf("failed to create directory %q: %v", path, err)
	}
}
