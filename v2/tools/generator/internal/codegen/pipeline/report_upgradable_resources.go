/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
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

// asoVersionDateRegex extracts the YYYYMMDD date from an ASO package version like "v20230101" or "v1api20230101"
var asoVersionDateRegex = regexp.MustCompile(`(\d{8})`)

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

// upgradableResourcesReportItem captures the information for a single resource that has an available upgrade.
// Version strings use ASO package naming, e.g. "v20230101" or "v1api20230101".
type upgradableResourcesReportItem struct {
	group            string
	resource         string
	supportedStable  string
	supportedPreview string
	availableStable  string
	availablePreview string
}

// UpgradableResourcesReport is a report of resources that have available upgrades
type UpgradableResourcesReport struct {
	cfg   *config.UpgradableResourcesReport
	items []upgradableResourcesReportItem
}

// resourceVersionInfo tracks the latest stable and preview package versions for a resource
type resourceVersionInfo struct {
	stable  string // ASO package name, e.g. "v20230101" or "v1api20230101"
	preview string // ASO package name, e.g. "v20230101preview" or "v1api20230101preview"
}

// resourceKey identifies a resource uniquely within the pipeline state
type resourceKey struct {
	group string
	name  string
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
		for name := range names {
			// TypeNameSet stores TypeName (interface), but the values are InternalTypeName (struct)
			internalName, ok := name.(astmodel.InternalTypeName)
			if !ok {
				continue
			}

			pkg := internalName.InternalPackageReference()
			if astmodel.IsStoragePackageReference(pkg) {
				continue
			}

			key := resourceKey{group: group, name: internalName.Name()}
			info := available[key]
			pkgName := pkg.PackageName()
			if pkg.IsPreview() {
				if isVersionNewer(pkgName, info.preview) {
					info.preview = pkgName
				}
			} else {
				if isVersionNewer(pkgName, info.stable) {
					info.stable = pkgName
				}
			}
			available[key] = info
		}
	}

	// Build a map of (group, name) → latest supported stable/preview versions from filtered definitions.
	supportedMap := make(map[resourceKey]resourceVersionInfo)
	for _, rsrc := range supported.AllResources() {
		name := rsrc.Name()
		pkg := name.InternalPackageReference()
		if astmodel.IsStoragePackageReference(pkg) {
			continue
		}

		key := resourceKey{group: pkg.Group(), name: name.Name()}
		info := supportedMap[key]
		pkgName := pkg.PackageName()
		if pkg.IsPreview() {
			if isVersionNewer(pkgName, info.preview) {
				info.preview = pkgName
			}
		} else {
			if isVersionNewer(pkgName, info.stable) {
				info.stable = pkgName
			}
		}
		supportedMap[key] = info
	}

	now := time.Now()

	// For each supported resource, check if an upgrade is recommended
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

		shouldInclude := false

		// Check stable upgrade
		if rec.availableStable != "" && isVersionNewer(rec.availableStable, rec.supportedStable) {
			if report.isStableUpgradeRecommended(rec.supportedStable, rec.availableStable, now) {
				shouldInclude = true
			}
		}

		// Check preview upgrade
		if rec.availablePreview != "" && isVersionNewer(rec.availablePreview, rec.supportedPreview) {
			if report.isPreviewUpgradeRecommended(rec.supportedPreview, rec.availablePreview) {
				shouldInclude = true
			}
		}

		if shouldInclude {
			report.items = append(report.items, rec)
		}
	}

	// Sort by group, then resource name
	slices.SortFunc(report.items, func(a, b upgradableResourcesReportItem) int {
		if a.group != b.group {
			if a.group < b.group {
				return -1
			}
			return 1
		}
		if a.resource < b.resource {
			return -1
		}
		if a.resource > b.resource {
			return 1
		}
		return 0
	})

	return report
}

// isStableUpgradeRecommended returns true if a stable upgrade is recommended based on configured thresholds.
func (r *UpgradableResourcesReport) isStableUpgradeRecommended(supported, available string, now time.Time) bool {
	supportedDate, err := parseASOVersionDate(supported)
	if err != nil {
		// If we can't parse the supported version, be conservative
		return false
	}

	availableDate, err := parseASOVersionDate(available)
	if err != nil {
		return false
	}

	stableUpgradeThreshold := r.cfg.StableVersionsExpiry
	if stableUpgradeThreshold == 0 {
		stableUpgradeThreshold = 12
	}

	latestThreshold := r.cfg.LatestVersionThreshold
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
func (r *UpgradableResourcesReport) isPreviewUpgradeRecommended(supported, available string) bool {
	supportedDate, err := parseASOVersionDate(supported)
	if err != nil {
		// If we can't parse the supported version (e.g., empty), use zero time (infinitely old)
		supportedDate = time.Time{}
	}

	availableDate, err := parseASOVersionDate(available)
	if err != nil {
		return false
	}

	previewUpgradeThreshold := r.cfg.PreviewVersionsExpiry
	if previewUpgradeThreshold == 0 {
		previewUpgradeThreshold = 6
	}

	return monthsBetween(supportedDate, availableDate) > previewUpgradeThreshold
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
	buffer.WriteString("# Upgradable Resources\n\n")

	if len(r.items) == 0 {
		buffer.WriteString("No resources with available upgrades were found.\n")
		return
	}

	buffer.WriteString("The following resources have newer versions available in the Azure REST API specifications.\n\n")

	table := reporting.NewMarkdownTable(
		"Group",
		"Resource",
		"Supported Stable",
		"Supported Preview",
		"Available Stable",
		"Available Preview")

	for _, item := range r.items {
		table.AddRow(
			item.group,
			item.resource,
			orDash(item.supportedStable),
			orDash(item.supportedPreview),
			orDash(item.availableStable),
			orDash(item.availablePreview))
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

// parseASOVersionDate parses the date embedded in an ASO package version string.
// It extracts the YYYYMMDD component from versions like "v20230101", "v1api20230101",
// or "v1api20230101preview".
func parseASOVersionDate(version string) (time.Time, error) {
	if version == "" {
		return time.Time{}, eris.New("empty version")
	}

	m := asoVersionDateRegex.FindString(version)
	if m == "" {
		return time.Time{}, eris.Errorf("version %q does not contain expected YYYYMMDD date", version)
	}

	return time.Parse("20060102", m)
}

// monthsBetween returns the approximate number of months between two dates.
func monthsBetween(from, to time.Time) int {
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
