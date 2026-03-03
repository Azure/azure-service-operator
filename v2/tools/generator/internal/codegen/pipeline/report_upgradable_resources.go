/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"io/fs"
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
	return NewStage(
		ReportUpgradableResourcesStageID,
		"Generate a report identifying resources with available upgrades",
		func(ctx context.Context, state *State) (*State, error) {
			reportCfg := configuration.UpgradableResourcesReport
			if reportCfg.OutputFolder == "" {
				// No output folder configured, skip this stage
				return state, nil
			}

			report, err := NewUpgradableResourcesReport(state.Definitions(), configuration)
			if err != nil {
				return nil, err
			}

			outputFile := reportCfg.FullOutputPath()
			err = report.SaveTo(outputFile)
			if err != nil {
				return nil, eris.Wrapf(err, "writing upgradable resources report to %s", outputFile)
			}

			return state, nil
		})
}

// upgradableResourcesReportItem captures the information for a single resource that has an available upgrade
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

// NewUpgradableResourcesReport creates a new UpgradableResourcesReport.
func NewUpgradableResourcesReport(
	definitions astmodel.TypeDefinitionSet,
	cfg *config.Configuration,
) (*UpgradableResourcesReport, error) {
	report := &UpgradableResourcesReport{
		cfg: cfg.UpgradableResourcesReport,
	}

	// Build a map of provider (lowercase) -> available spec versions
	specVersions, err := scanSpecVersions(cfg.SchemaRoot)
	if err != nil {
		return nil, eris.Wrapf(err, "scanning spec versions from %s", cfg.SchemaRoot)
	}

	// Collect the latest supported stable/preview version for each (group, resourceName)
	type resourceKey struct {
		group    string
		name     string
		provider string // lowercase provider from armType
	}

	type versionInfo struct {
		stable  string
		preview string
	}

	supported := make(map[resourceKey]versionInfo)

	for _, rsrc := range definitions.AllResources() {
		name := rsrc.Name()
		pkg := name.InternalPackageReference()

		if astmodel.IsStoragePackageReference(pkg) {
			continue
		}

		defType := astmodel.MustBeResourceType(rsrc.Type())
		armType := defType.ARMType()
		if armType == "" {
			continue
		}

		// Extract provider from ARM type (e.g., "Microsoft.Storage" from "Microsoft.Storage/storageAccounts")
		provider := armTypeProvider(armType)
		if provider == "" {
			continue
		}

		armVersion := strings.Trim(defType.APIVersionEnumValue().Value, "\"")
		if armVersion == "" {
			continue
		}

		grp := pkg.Group()
		key := resourceKey{group: grp, name: name.Name(), provider: strings.ToLower(provider)}

		info := supported[key]
		if pkg.IsPreview() {
			if isVersionNewer(armVersion, info.preview) {
				info.preview = armVersion
			}
		} else {
			if isVersionNewer(armVersion, info.stable) {
				info.stable = armVersion
			}
		}
		supported[key] = info
	}

	now := time.Now()

	// For each resource, check if an upgrade is recommended
	for key, info := range supported {
		specVers, ok := specVersions[key.provider]
		if !ok {
			continue
		}

		var rec upgradableResourcesReportItem
		rec.group = key.group
		rec.resource = key.name
		rec.supportedStable = info.stable
		rec.supportedPreview = info.preview
		rec.availableStable = specVers.latestStable
		rec.availablePreview = specVers.latestPreview

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

	return report, nil
}

// isStableUpgradeRecommended returns true if a stable upgrade is recommended based on configured thresholds.
func (r *UpgradableResourcesReport) isStableUpgradeRecommended(supported, available string, now time.Time) bool {
	supportedDate, err := parseARMVersionDate(supported)
	if err != nil {
		// If we can't parse the supported version, be conservative
		return false
	}

	availableDate, err := parseARMVersionDate(available)
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
	supportedDate, err := parseARMVersionDate(supported)
	if err != nil {
		// If we can't parse the supported version (e.g., empty), use zero time
		supportedDate = time.Time{}
	}

	availableDate, err := parseARMVersionDate(available)
	if err != nil {
		return false
	}

	previewUpgradeThreshold := r.cfg.PreviewVersionsExpiry
	if previewUpgradeThreshold == 0 {
		previewUpgradeThreshold = 6
	}

	// If there's no currently supported preview version, use a zero time (infinitely old)
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

// specProviderVersions holds the latest stable and preview versions for a provider
type specProviderVersions struct {
	latestStable  string
	latestPreview string
}

// scanSpecVersions scans the spec root directory and returns a map of
// provider (lowercase) -> latest available stable and preview versions.
func scanSpecVersions(schemaRoot string) (map[string]specProviderVersions, error) {
	result := make(map[string]specProviderVersions)

	// If schemaRoot doesn't exist, return empty map (not an error)
	if _, err := os.Stat(schemaRoot); os.IsNotExist(err) {
		return result, nil
	}

	// Walk the spec directory looking for stable/preview directories
	err := filepath.WalkDir(schemaRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil // Skip errors
		}

		if !d.IsDir() {
			return nil
		}

		dirName := d.Name()
		isStable := dirName == "stable"
		isPreview := dirName == "preview"

		if !isStable && !isPreview {
			return nil
		}

		// Extract the provider from the path (the Microsoft.XXX part)
		provider := extractProviderFromPath(path, schemaRoot)
		if provider == "" {
			return nil
		}

		providerKey := strings.ToLower(provider)

		// List subdirectories of this stable/preview dir - these are the versions
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil
		}

		current := result[providerKey]

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			versionDir := entry.Name()
			if !armVersionDateRegex.MatchString(versionDir) {
				continue
			}

			// Determine if it's a preview version
			isPreviewVersion := isPreview || armVersionDateRegex.FindStringSubmatch(versionDir)[2] != ""

			if isPreviewVersion {
				if isVersionNewer(versionDir, current.latestPreview) {
					current.latestPreview = versionDir
				}
			} else {
				if isVersionNewer(versionDir, current.latestStable) {
					current.latestStable = versionDir
				}
			}
		}

		result[providerKey] = current
		return nil
	})

	return result, err
}

// extractProviderFromPath finds the Microsoft.XXX provider from a file path.
// The provider is the first path component matching Microsoft.XXX found after the schemaRoot.
func extractProviderFromPath(path, schemaRoot string) string {
	// Normalize slashes
	path = filepath.ToSlash(path)
	schemaRoot = filepath.ToSlash(schemaRoot)
	rel := strings.TrimPrefix(path, schemaRoot)

	parts := strings.Split(strings.Trim(rel, "/"), "/")
	for _, part := range parts {
		if strings.HasPrefix(strings.ToLower(part), "microsoft.") {
			return part
		}
	}

	return ""
}

// armTypeProvider extracts the provider (e.g., "Microsoft.Storage") from an ARM type
// like "Microsoft.Storage/storageAccounts".
func armTypeProvider(armType string) string {
	idx := strings.Index(armType, "/")
	if idx < 0 {
		return armType
	}
	return armType[:idx]
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

	// Extract just the date part (YYYY-MM-DD)
	m := armVersionDateRegex.FindStringSubmatch(version)
	if m == nil {
		return time.Time{}, eris.Errorf("version %q is not in expected date format", version)
	}

	return time.Parse("2006-01-02", m[1])
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
