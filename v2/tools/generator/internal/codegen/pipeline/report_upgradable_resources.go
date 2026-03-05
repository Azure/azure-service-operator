/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"cmp"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
)

// ReportUpgradableResourcesStageID is the unique identifier of this stage
const ReportUpgradableResourcesStageID = "reportUpgradableResources"

// armVersionDateRegex matches date-based ARM API versions like 2023-01-01 or 2023-01-01-preview
var armVersionDateRegex = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})(-[a-z]+)?$`)

// ReportUpgradableResources creates a pipeline stage that generates a report identifying resources
// with available newer versions in the azure-rest-api-specs.
func ReportUpgradableResources(configuration *config.Configuration) *Stage {
	stage := NewStage(
		ReportUpgradableResourcesStageID,
		"Generate a report identifying resources with available upgrades",
		func(ctx context.Context, state *State) (*State, error) {
			reportCfg := configuration.UpgradableResourcesReport
			if reportCfg.OutputFolder == "" {
				// No output folder configured, skip this stage
				return state, nil
			}

			allKnownResources, err := GetStateData[map[string]astmodel.TypeNameSet](state, AllKnownResources)
			if err != nil {
				return nil, eris.Wrap(err, "retrieving known resources from pipeline state")
			}

			report := NewUpgradableResourcesReport(allKnownResources, state.Definitions(), configuration.UpgradableResourcesReport)

			outputFile := reportCfg.FullOutputPath()
			err = report.SaveTo(outputFile)
			if err != nil {
				return nil, eris.Wrapf(err, "writing upgradable resources report to %s", outputFile)
			}

			return state, nil
		})

	stage.RequiresPrerequisiteStages(CatalogKnownResourcesStageID)
	return stage
}

// versionEntry holds both the ASO package name (for display) and the ARM version (for date comparison).
type versionEntry struct {
	pkgName    string // ASO package name for display, e.g. "v20230101" or "v1api20230101"
	armVersion string // ARM date string for comparison, e.g. "2023-01-01" or "2023-01-01-preview"
}

// isEmpty returns true if this entry has no version information.
func (v versionEntry) isEmpty() bool {
	return v.pkgName == ""
}

// resourceVersionInfo tracks the latest stable and preview versions for a resource.
type resourceVersionInfo struct {
	stable  versionEntry
	preview versionEntry
}

// resourceKey identifies a resource uniquely within the pipeline state.
type resourceKey struct {
	group string
	name  string
}

// upgradableResourcesReportItem captures the information for a single resource that has a newer version available.
// Version strings use ASO package naming, e.g. "v20230101" or "v1api20230101preview".
type upgradableResourcesReportItem struct {
	group            string
	resource         string
	supportedStable  versionEntry
	supportedPreview versionEntry
	availableStable  versionEntry
	availablePreview versionEntry
}

// hasStableUpgrade returns true if there is a newer stable version available.
func (i upgradableResourcesReportItem) hasStableUpgrade() bool {
	return !i.availableStable.isEmpty() && isVersionNewer(i.availableStable.pkgName, i.supportedStable.pkgName)
}

// hasPreviewUpgrade returns true if there is a newer preview version available.
func (i upgradableResourcesReportItem) hasPreviewUpgrade() bool {
	return !i.availablePreview.isEmpty() && isVersionNewer(i.availablePreview.pkgName, i.supportedPreview.pkgName)
}

// isStableUpgradeRecommended returns true if a stable upgrade is recommended based on configured thresholds.
func (i upgradableResourcesReportItem) isStableUpgradeRecommended(cfg *config.UpgradableResourcesReport, now time.Time) bool {
	if !i.hasStableUpgrade() {
		return false
	}

	supportedDate, err := parseARMVersionDate(i.supportedStable.armVersion)
	if err != nil {
		// If we can't parse the supported version, be conservative
		return false
	}

	availableDate, err := parseARMVersionDate(i.availableStable.armVersion)
	if err != nil {
		return false
	}

	stableUpgradeThreshold := cfg.StableVersionsExpiry
	if stableUpgradeThreshold == 0 {
		stableUpgradeThreshold = 12
	}

	latestThreshold := cfg.LatestVersionThreshold
	if latestThreshold == 0 {
		latestThreshold = 24
	}

	// Condition: available is more than stableUpgradeThreshold months newer than supported
	if monthsBetween(supportedDate, availableDate) > stableUpgradeThreshold {
		return true
	}

	// Condition: the currently supported version is more than latestThreshold months older than today
	if monthsBetween(supportedDate, now) > latestThreshold {
		return true
	}

	return false
}

// isPreviewUpgradeRecommended returns true if a preview upgrade is recommended based on configured thresholds.
func (i upgradableResourcesReportItem) isPreviewUpgradeRecommended(cfg *config.UpgradableResourcesReport) bool {
	if !i.hasPreviewUpgrade() {
		return false
	}

	supportedDate, err := parseARMVersionDate(i.supportedPreview.armVersion)
	if err != nil {
		// If we can't parse the supported version (e.g., empty), use zero time (infinitely old)
		supportedDate = time.Time{}
	}

	availableDate, err := parseARMVersionDate(i.availablePreview.armVersion)
	if err != nil {
		return false
	}

	previewUpgradeThreshold := cfg.PreviewVersionsExpiry
	if previewUpgradeThreshold == 0 {
		previewUpgradeThreshold = 6
	}

	return monthsBetween(supportedDate, availableDate) > previewUpgradeThreshold
}

// UpgradableResourcesReport is a report of resources that have available upgrades
type UpgradableResourcesReport struct {
	cfg   *config.UpgradableResourcesReport
	items []upgradableResourcesReportItem
}

// NewUpgradableResourcesReport creates a new UpgradableResourcesReport.
// allKnownResources is the full catalog of resources before export filters (from CatalogKnownResources stage).
// supported is the set of type definitions that passed through the export filter.
// cfg is the report configuration.
func NewUpgradableResourcesReport(
	allKnownResources map[string]astmodel.TypeNameSet,
	supported astmodel.TypeDefinitionSet,
	cfg *config.UpgradableResourcesReport,
) *UpgradableResourcesReport {
	report := &UpgradableResourcesReport{cfg: cfg}

	// Build a map of (group, name) → latest available stable/preview versions from AllKnownResources.
	// This represents all versions present in the specs, before any export filtering.
	available := make(map[resourceKey]resourceVersionInfo)
	for group, names := range allKnownResources {
		for typeName := range names {
			internalName, ok := astmodel.AsInternalTypeName(typeName)
			if !ok {
				continue
			}
			recordLatestVersions(available, group, internalName, internalName.InternalPackageReference().APIVersion())
		}
	}

	// Build a map of (group, name) → latest supported stable/preview versions from filtered definitions.
	// Use APIVersionEnumValue() to obtain the precise ARM API version for date comparisons,
	// falling back to the package's APIVersion() when the enum value is not available.
	supportedMap := make(map[resourceKey]resourceVersionInfo)
	for _, rsrc := range supported.AllResources() {
		name := rsrc.Name()
		defType := astmodel.MustBeResourceType(rsrc.Type())
		armVersion := strings.Trim(defType.APIVersionEnumValue().Value, "\"")
		if armVersion == "" {
			armVersion = name.InternalPackageReference().APIVersion()
		}
		recordLatestVersions(supportedMap, name.InternalPackageReference().Group(), name, armVersion)
	}

	// For each supported resource, include it in the report if any newer version is available.
	for key, supportInfo := range supportedMap {
		availInfo, ok := available[key]
		if !ok {
			continue
		}

		rec := upgradableResourcesReportItem{
			group:            key.group,
			resource:         key.name,
			supportedStable:  supportInfo.stable,
			supportedPreview: supportInfo.preview,
			availableStable:  availInfo.stable,
			availablePreview: availInfo.preview,
		}

		if rec.hasStableUpgrade() || rec.hasPreviewUpgrade() {
			report.items = append(report.items, rec)
		}
	}

	// Sort by group, then resource name
	slices.SortFunc(report.items, func(a, b upgradableResourcesReportItem) int {
		result := cmp.Compare(a.group, b.group)
		if result == 0 {
			result = cmp.Compare(a.resource, b.resource)
		}
		return result
	})

	return report
}

// recordLatestVersions updates the version map with the latest version for the given resource.
// It is called for both the AllKnownResources catalog (type names only) and the supported definitions.
func recordLatestVersions(m map[resourceKey]resourceVersionInfo, group string, name astmodel.InternalTypeName, armVersion string) {
	pkg := name.InternalPackageReference()
	if astmodel.IsStoragePackageReference(pkg) {
		return
	}

	key := resourceKey{group: group, name: name.Name()}
	info := m[key]
	entry := versionEntry{pkgName: pkg.PackageName(), armVersion: armVersion}
	if pkg.IsPreview() {
		if isVersionNewer(entry.pkgName, info.preview.pkgName) {
			info.preview = entry
		}
	} else {
		if isVersionNewer(entry.pkgName, info.stable.pkgName) {
			info.stable = entry
		}
	}
	m[key] = info
}

// SaveTo writes the report to the specified file.
func (r *UpgradableResourcesReport) SaveTo(outputFile string) error {
	var buffer strings.Builder
	r.WriteTo(&buffer)

	outputFolder := filepath.Dir(outputFile)
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		if err = os.MkdirAll(outputFolder, 0o700); err != nil {
			return eris.Wrapf(err, "unable to create directory %q", outputFolder)
		}
	}

	return os.WriteFile(outputFile, []byte(buffer.String()), 0o600)
}

// WriteTo writes the report content to the provided buffer.
func (r *UpgradableResourcesReport) WriteTo(buffer *strings.Builder) {
	r.writeTo(buffer, time.Now())
}

// writeTo writes the report content to the provided buffer, using now for recommendation date comparisons.
// This internal method accepts a fixed time to allow deterministic tests.
func (r *UpgradableResourcesReport) writeTo(buffer *strings.Builder, now time.Time) {
	buffer.WriteString("# Upgradable Resources\n\n")

	if len(r.items) == 0 {
		buffer.WriteString("No resources with available upgrades were found.\n")
		return
	}

	buffer.WriteString("The following resources have newer versions available in the Azure REST API specifications." +
		" Versions in **bold** are recommended for upgrade.\n\n")

	table := reporting.NewMarkdownTable(
		"Group",
		"Resource",
		"Supported Stable",
		"Supported Preview",
		"Available Stable",
		"Available Preview")

	for _, item := range r.items {
		stableAvail := orDash(item.availableStable.pkgName)
		if item.isStableUpgradeRecommended(r.cfg, now) {
			stableAvail = bold(stableAvail)
		}

		previewAvail := orDash(item.availablePreview.pkgName)
		if item.isPreviewUpgradeRecommended(r.cfg) {
			previewAvail = bold(previewAvail)
		}

		table.AddRow(
			item.group,
			item.resource,
			orDash(item.supportedStable.pkgName),
			orDash(item.supportedPreview.pkgName),
			stableAvail,
			previewAvail)
	}

	table.WriteTo(buffer)
	buffer.WriteString("\n")
}

// orDash returns the value if non-empty, otherwise "-"
func orDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

// bold wraps a non-dash value in Markdown bold markers.
func bold(s string) string {
	if s == "-" {
		return s
	}
	return "**" + s + "**"
}

// isVersionNewer returns true if candidate is strictly newer than current.
// An empty current means any non-empty candidate is newer.
func isVersionNewer(candidate, current string) bool {
	if candidate == "" {
		return false
	}
	if current == "" {
		return true
	}
	return astmodel.ComparePathAndVersion(candidate, current) > 0
}

// parseARMVersionDate parses the date part of an ARM API version string.
// For example, "2023-01-01" → 2023-01-01, "2023-01-01-preview" → 2023-01-01.
func parseARMVersionDate(version string) (time.Time, error) {
	if version == "" {
		return time.Time{}, eris.New("empty version")
	}

	m := armVersionDateRegex.FindStringSubmatch(version)
	if m == nil {
		return time.Time{}, eris.Errorf("version %q is not in expected ARM date format (YYYY-MM-DD)", version)
	}

	return time.Parse("2006-01-02", m[1])
}

// monthsBetween returns the approximate number of months between two dates.
func monthsBetween(from time.Time, to time.Time) int {
	years := to.Year() - from.Year()
	months := int(to.Month()) - int(from.Month())
	total := years*12 + months
	if total < 0 {
		return 0
	}
	return total
}

// Ensure UpgradableResourcesReport implements fmt.Stringer for debugging
var _ fmt.Stringer = &UpgradableResourcesReport{}

// String returns a string representation for debugging.
func (r *UpgradableResourcesReport) String() string {
	return fmt.Sprintf("UpgradableResourcesReport{items: %d}", len(r.items))
}
