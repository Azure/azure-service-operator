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

const (
	// ReportUpgradableResourcesStageID is the unique identifier of this stage
	ReportUpgradableResourcesStageID = "reportUpgradableResources"
	// defaultStableVersionExpiryMonths is the default number of months before a stable version is considered expired
	defaultStableVersionExpiryMonths = 12
	// defaultLatestVersionThresholdMonths is the default number of months before a latest version is considered outdated
	defaultLatestVersionThresholdMonths = 24
	// defaultPreviewVersionExpiryMonths is the default number of months before a preview version is considered expired
	defaultPreviewVersionExpiryMonths = 6
)

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

// resourceVersionInfo tracks the latest stable and preview package references for a resource.
type resourceVersionInfo struct {
	stable  astmodel.InternalPackageReference
	preview astmodel.InternalPackageReference
}

// resourceKey identifies a resource uniquely within the pipeline state.
type resourceKey struct {
	group string
	name  string
}

// latestVersionIndex maps each (group, resource) pair to the latest stable and preview package references seen so far.
type latestVersionIndex map[resourceKey]resourceVersionInfo

// Add records the package reference of name, updating the index to track the latest stable and preview versions.
// Storage package references are silently ignored.
func (idx latestVersionIndex) Add(name astmodel.InternalTypeName) {
	pkg := name.InternalPackageReference()
	if astmodel.IsStoragePackageReference(pkg) {
		return
	}

	key := resourceKey{group: pkg.Group(), name: name.Name()}
	info := idx[key]
	if pkg.IsPreview() {
		if isVersionNewer(pkg, info.preview) {
			info.preview = pkg
		}
	} else {
		if isVersionNewer(pkg, info.stable) {
			info.stable = pkg
		}
	}
	idx[key] = info
}

// upgradableResourcesReportItem captures the information for a single resource that has a newer version available.
type upgradableResourcesReportItem struct {
	group            string
	resource         string
	supportedStable  astmodel.InternalPackageReference
	supportedPreview astmodel.InternalPackageReference
	availableStable  astmodel.InternalPackageReference
	availablePreview astmodel.InternalPackageReference
}

// hasStableUpgrade returns true if there is a newer stable version available.
// If the resource has no supported stable version, stable upgrades are not shown.
func (i upgradableResourcesReportItem) hasStableUpgrade() bool {
	return i.supportedStable != nil &&
		i.availableStable != nil &&
		isVersionNewer(i.availableStable, i.supportedStable)
}

// hasPreviewUpgrade returns true if there is a newer preview version available.
// If the resource has no supported preview version, preview upgrades are not shown.
func (i upgradableResourcesReportItem) hasPreviewUpgrade() bool {
	return i.supportedPreview != nil &&
		i.availablePreview != nil &&
		isVersionNewer(i.availablePreview, i.supportedPreview)
}

// isStableUpgradeRecommended returns true if a stable upgrade is recommended based on configured thresholds.
func (i upgradableResourcesReportItem) isStableUpgradeRecommended(
	cfg *config.UpgradableResourcesReport,
	now time.Time,
) bool {
	if !i.hasStableUpgrade() {
		return false
	}

	if i.supportedStable == nil {
		return false
	}

	supportedDate, err := parseARMVersionDate(i.supportedStable.APIVersion())
	if err != nil {
		// If we can't parse the supported version, be conservative
		return false
	}

	availableDate, err := parseARMVersionDate(i.availableStable.APIVersion())
	if err != nil {
		return false
	}

	stableUpgradeThreshold := cfg.StableVersionsExpiry
	if stableUpgradeThreshold == 0 {
		stableUpgradeThreshold = defaultStableVersionExpiryMonths
	}

	latestThreshold := cfg.LatestVersionThreshold
	if latestThreshold == 0 {
		latestThreshold = defaultLatestVersionThresholdMonths
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

	var supportedDate time.Time
	if i.supportedPreview != nil {
		d, err := parseARMVersionDate(i.supportedPreview.APIVersion())
		if err == nil {
			supportedDate = d
		}
		// If we can't parse, keep zero time (infinitely old), so gap will always exceed threshold
	}

	availableDate, err := parseARMVersionDate(i.availablePreview.APIVersion())
	if err != nil {
		return false
	}

	previewUpgradeThreshold := cfg.PreviewVersionsExpiry
	if previewUpgradeThreshold == 0 {
		previewUpgradeThreshold = defaultPreviewVersionExpiryMonths
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

	// Build an index of (group, name) → latest available stable/preview versions from AllKnownResources.
	// This represents all versions present in the specs, before any export filtering.
	available := make(latestVersionIndex)
	for _, names := range allKnownResources {
		for typeName := range names {
			if internalName, ok := astmodel.AsInternalTypeName(typeName); ok {
				available.Add(internalName)
			}
		}
	}

	// Build an index of (group, name) → latest supported stable/preview versions from filtered definitions.
	supportedIdx := make(latestVersionIndex)
	for _, rsrc := range supported.AllResources() {
		supportedIdx.Add(rsrc.Name())
	}

	// For each supported resource, include it in the report if any newer version is available.
	for key, supportInfo := range supportedIdx {
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
		" Resources with a 💡 have a recommended or overdue update.\n\n")

	// Items are sorted by group then resource; iterate and flush a new table for each group.
	var currentGroup string
	var table *reporting.MarkdownTable

	flushTable := func() {
		if table != nil {
			table.WriteTo(buffer)
			buffer.WriteString("\n")
		}
	}

	for _, item := range r.items {
		if item.group != currentGroup {
			flushTable()
			currentGroup = item.group
			fmt.Fprintf(buffer, "## %s\n\n", currentGroup)
			table = reporting.NewMarkdownTable(
				"",
				"Resource",
				"Available Stable",
				"Supported Stable",
				"Available Preview",
				"Supported Preview")
		}

		stableSupported := pkgRefVersion(item.supportedStable)

		stableAvail := orDash(pkgRefVersion(item.availableStable))
		if item.isStableUpgradeRecommended(r.cfg, now) {
			stableAvail = bold(stableAvail)
		}

		if stableAvail == stableSupported {
			// No actual upgrade
			stableAvail = ""
		}

		previewSupported := pkgRefVersion(item.supportedPreview)

		previewAvail := orDash(pkgRefVersion(item.availablePreview))
		if item.isPreviewUpgradeRecommended(r.cfg) {
			previewAvail = bold(previewAvail)
		}

		if previewAvail == previewSupported {
			// No actual upgrade
			previewAvail = ""
		}

		indicator := ""
		if item.isStableUpgradeRecommended(r.cfg, now) || item.isPreviewUpgradeRecommended(r.cfg) {
			indicator = "💡"
		}

		table.AddRow(
			indicator,
			item.resource,
			stableAvail,
			orDash(stableSupported),
			previewAvail,
			orDash(previewSupported))
	}

	flushTable()
}

// pkgRefVersion returns the API version of the given reference, or "" if the reference is nil.
func pkgRefVersion(ref astmodel.InternalPackageReference) string {
	if ref == nil {
		return ""
	}
	return ref.APIVersion()
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
// A nil current means any non-nil candidate is newer.
func isVersionNewer(candidate, current astmodel.InternalPackageReference) bool {
	if candidate == nil {
		return false
	}

	if current == nil {
		return true
	}

	return candidate.APIVersion() > current.APIVersion()
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
