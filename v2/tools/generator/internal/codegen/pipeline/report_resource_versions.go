/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/rotisserie/eris"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/typo"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
)

const (
	// ReportResourceVersionsStageID is the unique identifier of this stage
	ReportResourceVersionsStageID = "reportResourceVersions"
	v1VersionPrefix               = "v1api"
)

// ReportResourceVersions creates a pipeline stage that generates a report listing all generated resources.
func ReportResourceVersions(configuration *config.Configuration) *Stage {
	return NewStage(
		ReportResourceVersionsStageID,
		"Generate a report listing all the resources generated",
		func(ctx context.Context, state *State) (*State, error) {
			report, err := NewResourceVersionsReport(state.Definitions(), configuration)
			if err != nil {
				return nil, err
			}

			err = report.SaveAllResourcesReportTo(configuration.SupportedResourcesReport.FullOutputPath())
			if err != nil {
				return nil, err
			}

			var errs []error
			for grp := range report.groups {
				outputFile := configuration.SupportedResourcesReport.GroupFullOutputPath(grp)
				err = report.SaveGroupResourcesReportTo(grp, outputFile)
				if err != nil {
					errs = append(errs, eris.Wrapf(err, "writing versions report to %s for group %s", outputFile, grp))
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			err = configuration.ObjectModelConfiguration.SupportedFrom.VerifyConsumed()
			return state, err
		})
}

type ResourceVersionsReport struct {
	reportConfiguration      *config.SupportedResourcesReport
	objectModelConfiguration *config.ObjectModelConfiguration
	rootURL                  string
	samplesPath              string
	availableFragments       map[string]string                                      // A collection of the fragments to use in the report
	groups                   set.Set[string]                                        // A set of all our groups
	items                    map[string]set.Set[ResourceVersionsReportResourceItem] // For each group, the set of all available items
	typoAdvisor              *typo.Advisor                                          // Advisor used to troubleshoot unused fragments
	titleCase                cases.Caser                                            // Helper for title casing
	localPath                string                                                 // Prefix to use for local packages
}

type ResourceVersionsReportResourceItem struct {
	name          astmodel.InternalTypeName
	armType       string
	armVersion    string
	supportedFrom string
}

type ResourceVersionsReportGroupInfo struct {
	Group    string
	Provider string
	Title    string
}

func NewResourceVersionsReport(
	definitions astmodel.TypeDefinitionSet,
	cfg *config.Configuration,
) (*ResourceVersionsReport, error) {
	result := &ResourceVersionsReport{
		reportConfiguration:      cfg.SupportedResourcesReport,
		objectModelConfiguration: cfg.ObjectModelConfiguration,
		rootURL:                  cfg.RootURL,
		samplesPath:              cfg.FullSamplesPath(),
		availableFragments:       make(map[string]string),
		groups:                   set.Make[string](),
		items:                    make(map[string]set.Set[ResourceVersionsReportResourceItem]),
		typoAdvisor:              typo.NewAdvisor(),
		titleCase:                cases.Title(language.English),
		localPath:                cfg.LocalPathPrefix(),
	}

	err := result.loadFragments()
	if err != nil {
		return nil, eris.Wrapf(err, "Unable to load report fragments")
	}

	err = result.summarize(definitions)
	if err != nil {
		return nil, eris.Wrapf(err, "Unable to summarize resources")
	}

	return result, nil
}

// loadFragments scans the files in the fragments directory and loads them into the availableFragments map
func (report *ResourceVersionsReport) loadFragments() error {
	if report.reportConfiguration.FragmentPath == "" {
		// No fragments to load
		return nil
	}

	fragmentsPath := report.reportConfiguration.FullFragmentPath()
	err := filepath.WalkDir(
		fragmentsPath,
		func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				// Failed to walk into path, abort early and propagate error
				return err
			}

			// Skip subdirectories
			if info.IsDir() {
				return nil
			}

			// Load the file contents
			content, err := os.ReadFile(path)
			if err != nil {
				return eris.Wrapf(err, "Unable to read fragment file %q", info.Name())
			}

			// Strip the extension from the filename
			name := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

			report.availableFragments[name] = string(content)
			return nil
		})
	if err != nil {
		return eris.Wrapf(err, "Unable to load fragments from %q", fragmentsPath)
	}

	// We don't want our README to trigger an error
	report.typoAdvisor.AddTerm("README")

	return nil
}

// summarize collates a list of all resources, grouped by package
func (report *ResourceVersionsReport) summarize(definitions astmodel.TypeDefinitionSet) error {
	for _, rsrc := range definitions.AllResources() {
		name := rsrc.Name()
		pkg := name.PackageReference()

		if astmodel.IsStoragePackageReference(pkg) {
			// Skip storage versions - they're an implementation detail
			continue
		}

		defType := astmodel.MustBeResourceType(rsrc.Type())
		armVersion := strings.Trim(defType.APIVersionEnumValue().Value, "\"")
		item := report.createItem(name, defType.ARMType(), armVersion)

		report.addItem(item)
	}

	handcraftedTypes, err := report.objectModelConfiguration.FindHandCraftedTypeNames(report.localPath)
	if err != nil {
		return err
	}

	for name := range handcraftedTypes {
		item := report.createItem(name, "", "")
		report.addItem(item)
	}

	return nil
}

func (report *ResourceVersionsReport) addItem(item ResourceVersionsReportResourceItem) {
	grp := item.name.InternalPackageReference().Group()
	report.groups.Add(grp)

	items, ok := report.items[grp]
	if !ok {
		items = set.Make[ResourceVersionsReportResourceItem]()
		report.items[grp] = items
	}

	items.Add(item)
}

// SaveAllResourcesReportTo creates a file containing a report listing all supported resources
// outputFile is the path to the file to create
func (report *ResourceVersionsReport) SaveAllResourcesReportTo(outputFile string) error {
	frontMatter := report.readFrontMatter(outputFile)

	var buffer strings.Builder
	err := report.WriteAllResourcesReportToBuffer(frontMatter, &buffer)
	if err != nil {
		return eris.Wrapf(err, "writing versions report to %s", outputFile)
	}

	err = report.ensureFolderExists(outputFile)
	if err != nil {
		return eris.Wrapf(err, "writing versions report to %s", outputFile)
	}

	return os.WriteFile(outputFile, []byte(buffer.String()), 0o600)
}

func (report *ResourceVersionsReport) ensureFolderExists(outputFile string) error {
	outputFolder := filepath.Dir(outputFile)
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err = os.MkdirAll(outputFolder, 0o700)
		if err != nil {
			return eris.Wrapf(err, "unable to create directory %q", outputFile)
		}
	}

	return nil
}

// SaveGroupResourcesReportTo creates a file containing a report listing supported resources in the specified group.
// group identifies the set of resources to include.
// outputFile is the path to the file to create.
func (report *ResourceVersionsReport) SaveGroupResourcesReportTo(group string, outputFile string) error {
	frontMatter := report.readFrontMatter(outputFile)

	var buffer strings.Builder
	err := report.WriteGroupResourcesReportToBuffer(group, frontMatter, &buffer)
	if err != nil {
		return eris.Wrapf(err, "writing versions report to %s for group %s", outputFile, group)
	}

	err = report.ensureFolderExists(outputFile)
	if err != nil {
		return eris.Wrapf(err, "writing versions report to %s", outputFile)
	}

	return os.WriteFile(outputFile, []byte(buffer.String()), 0o600)
}

// WriteAllResourcesReportToBuffer creates the report in the provided buffer
func (report *ResourceVersionsReport) WriteAllResourcesReportToBuffer(
	frontMatter string,
	buffer *strings.Builder,
) error {
	if frontMatter != "" {
		buffer.WriteString(frontMatter)
	} else {
		buffer.WriteString(report.defaultAllResourcesFrontMatter())
	}

	// Include file header if found
	err := report.writeFragment("all-resources-header", nil, buffer)
	if err != nil {
		return eris.Wrapf(err, "writing all-resources-header fragment")
	}

	// Sort groups into alphabetical order
	groups := set.AsSortedSlice(report.groups)

	errs := make([]error, 0, len(groups)) // Preallocate maximum size
	for _, grp := range groups {
		items := report.items[grp]

		info := report.groupInfo(grp, items)
		fmt.Fprintf(buffer, "## %s\n\n", info.Title)

		// Include our group fragment
		err := report.writeFragment("group-header", info, buffer)
		if err != nil {
			return eris.Wrapf(err, "writing group-header fragment for group %s", info.Group)
		}

		// Include a custom fragment for this group if we have one
		err = report.writeFragment(grp, info, buffer)
		if err != nil {
			errs = append(errs, err) // Don't need to wrap, will already specify the group
			continue
		}

		err = report.writeGroupSections(info, items, buffer)
		if err != nil {
			errs = append(errs, err) // Don't need to wrap, will already specify the group
		}
	}

	// Check to see whether all fragments were used
	for name := range report.availableFragments {
		if report.typoAdvisor.HasTerm(name) {
			continue
		}

		err := report.typoAdvisor.Errorf(name, "Fragment %q was not used", name)
		errs = append(errs, err)
	}

	return kerrors.NewAggregate(errs)
}

// WriteGroupResourcesReportToBuffer creates the report in the provided buffer
func (report *ResourceVersionsReport) WriteGroupResourcesReportToBuffer(
	group string,
	frontMatter string,
	buffer *strings.Builder,
) error {
	items := report.items[group]
	info := report.groupInfo(group, items)

	// Reuse existing front-matter if available, else generate a default one
	if frontMatter != "" {
		buffer.WriteString(frontMatter)
	} else {
		buffer.WriteString(report.defaultGroupResourcesFrontMatter(info.Title))
	}

	// Include our group fragment
	err := report.writeFragment("group-header", info, buffer)
	if err != nil {
		return eris.Wrapf(err, "writing group-header fragment for group %s", group)
	}

	// Include a fragment for this group if we have one
	err = report.writeFragment(group, info, buffer)
	if err != nil {
		return eris.Wrapf(err, "writing fragment for group %s", group)
	}

	return report.writeGroupSections(info, items, buffer)
}

type resourceVersionsReportSection struct {
	id    string
	title string
	kinds set.Set[ResourceVersionsReportResourceItem]
}

func (report *ResourceVersionsReport) writeGroupSections(
	group *ResourceVersionsReportGroupInfo,
	kinds set.Set[ResourceVersionsReportResourceItem],
	buffer *strings.Builder,
) error {
	// By default, we treat everything as released
	releasedResources := kinds

	createSection := func(
		id string,
		title string,
		kinds set.Set[ResourceVersionsReportResourceItem],
	) resourceVersionsReportSection {
		releasedResources = releasedResources.Except(kinds)

		return resourceVersionsReportSection{
			id:    id,
			title: title,
			kinds: kinds,
		}
	}

	// Create a section for all deprecated resources (this is commonly empty)
	deprecatedSection := createSection(
		"deprecated",
		"Deprecated",
		releasedResources.Where(report.isDeprecatedResource))

	// Create a section for all prerelease resources (those not yet released)
	prereleaseSection := createSection(
		"prerelease",
		"Next Release",
		releasedResources.Where(report.isUnreleasedResource))

	// Create a section for the latest versions of all supported resources
	latestSection := createSection(
		"latest",
		"Latest Released Versions",
		findRecommendedReleases(releasedResources))

	// Create a section for all other supported versions
	otherSection := createSection(
		"other",
		"Other Supported Versions",
		releasedResources)

	// Create an empty section for all released resources
	// (This will initially be empty, but )
	releasedSection := createSection(
		"released",
		"Released",
		set.Make[ResourceVersionsReportResourceItem]())

	// If every resource is the latest version, move them to the released section
	if len(otherSection.kinds) == 0 {
		releasedSection.kinds = latestSection.kinds
		latestSection.kinds = nil
	}

	sections := []resourceVersionsReportSection{
		prereleaseSection,
		latestSection,
		otherSection,
		releasedSection,
		deprecatedSection,
	}

	for _, section := range sections {
		err := report.writeSection(
			group,
			section.id,
			"### "+section.title,
			section.kinds,
			buffer)
		if err != nil {
			return eris.Wrapf(err, "writing section %s for group %s", section.id, group)
		}
	}

	// Always flag the prerelease and deprecated fragments as used
	report.typoAdvisor.AddTerm("prerelease")
	report.typoAdvisor.AddTerm("deprecated")

	return nil
}

// isUnreleasedResource returns true if the type definition is for an unreleased resource
func (report *ResourceVersionsReport) isUnreleasedResource(item ResourceVersionsReportResourceItem) bool {
	currentRelease := report.reportConfiguration.CurrentRelease
	if item.supportedFrom == currentRelease {
		return false
	}

	result := astmodel.ComparePathAndVersion(item.supportedFrom, currentRelease)
	return result >= 0
}

// isDeprecatedResource returns true if the type definition is for a deprecated resource
func (report *ResourceVersionsReport) isDeprecatedResource(item ResourceVersionsReportResourceItem) bool {
	_, ver := item.name.InternalPackageReference().GroupVersion()

	// Handcrafted versions are never deprecated
	// (reusing the regex from config to ensure consistency)
	if config.VersionRegex.MatchString(ver) {
		return false
	}

	result := !strings.HasPrefix(ver, astmodel.GeneratorVersion) && len(ver) > 3
	return result
}

// findRecommendedReleases selects a single version of each resource to recommend.
// Stable versions are preferred over preview.
// Later versions are preferred over earlier.
func findRecommendedReleases(
	kinds set.Set[ResourceVersionsReportResourceItem],
) set.Set[ResourceVersionsReportResourceItem] {
	index := make(map[string]ResourceVersionsReportResourceItem)

	for _, item := range kinds.Values() {
		known, ok := index[item.name.Name()]
		if !ok {
			// First time we've seen this resource, use this version
			index[item.name.Name()] = item
			continue
		}

		knownVersion := known.name.InternalPackageReference()
		itemVersion := item.name.InternalPackageReference()

		if knownVersion.IsPreview() && !itemVersion.IsPreview() {
			// Prefer stable versions over preview, so replace the preview version
			index[item.name.Name()] = item
			continue
		}

		if !knownVersion.IsPreview() && itemVersion.IsPreview() {
			// Prefer stable versions over preview, so keep what we have
			continue
		}

		// Both are either preview or stable, so compare versions
		if astmodel.ComparePathAndVersion(itemVersion.Version(), knownVersion.Version()) > 0 {
			// Prefer later versions
			index[item.name.Name()] = item
		}
	}

	result := set.Make[ResourceVersionsReportResourceItem]()

	for _, item := range index {
		result.Add(item)
	}

	return result
}

// writeSection writes a section to the buffer, consisting of a header, description, and a table listing resources.
func (report *ResourceVersionsReport) writeSection(
	info *ResourceVersionsReportGroupInfo,
	id string,
	heading string,
	items set.Set[ResourceVersionsReportResourceItem],
	buffer *strings.Builder,
) error {
	// Skip if nothing to do
	if len(items) == 0 {
		return nil
	}

	buffer.WriteString(heading)
	buffer.WriteString("\n\n")

	// Write a generic description, then a specific one
	err := report.writeFragment(fmt.Sprintf("%s-%s", info.Group, id), nil, buffer)
	if err != nil {
		return eris.Wrapf(err, "writing fragment for group %s", info.Group)
	}

	err = report.writeFragment(id, nil, buffer)
	if err != nil {
		return eris.Wrapf(err, "writing fragment for group %s", info.Group)
	}

	table, err := report.createTable(info, items)
	if err != nil {
		return eris.Wrapf(err, "creating table for group %s", info.Group)
	}

	table.WriteTo(buffer)
	buffer.WriteString("\n")

	return nil
}

func (report *ResourceVersionsReport) createTable(
	info *ResourceVersionsReportGroupInfo,
	items set.Set[ResourceVersionsReportResourceItem],
) (*reporting.MarkdownTable, error) {
	const (
		name          = "Resource"
		armVersion    = "ARM Version"
		crdVersion    = "CRD Version"
		sample        = "Sample"
		supportedFrom = "Supported From"
	)

	result := reporting.NewMarkdownTable(
		name,
		armVersion,
		crdVersion,
		supportedFrom,
		sample)

	toIterate := items.Values()
	slices.SortFunc(
		toIterate,
		func(i ResourceVersionsReportResourceItem, j ResourceVersionsReportResourceItem) int {
			left := i.name
			right := j.name
			if left.Name() < right.Name() {
				return -1
			} else if left.Name() > right.Name() {
				return 1
			}

			// Reversed parameters because we want more recent versions listed first
			return astmodel.ComparePathAndVersion(right.PackageReference().ImportPath(), left.PackageReference().ImportPath())
		})

	sampleLinks, err := report.FindSampleLinks(info.Group)
	if err != nil {
		return nil, err
	}

	for _, item := range toIterate {
		name := item.name

		crdVersion := name.PackageReference().PackageName()
		armVersion := item.armVersion
		if armVersion == "" {
			armVersion = crdVersion
		}

		api := report.generateAPILink(name)
		sample := report.generateSampleLink(name, sampleLinks)
		supportedFrom := report.supportedFrom(name)

		result.AddRow(
			api,
			armVersion,
			crdVersion,
			supportedFrom,
			sample)
	}

	return result, nil
}

func (report *ResourceVersionsReport) FindSampleLinks(group string) (map[string]string, error) {
	result := make(map[string]string)
	if report.rootURL != "" {
		parsedRootURL, err := url.Parse(report.rootURL)
		if err != nil {
			return nil, eris.Wrapf(err, "parsing rootUrl %s", report.rootURL)
		}

		// We look for samples within only the subfolder for this group - this avoids getting sample links wrong
		// if there are identically named resources in different groups
		basePath := filepath.Join(report.samplesPath, group)
		if _, err = os.Stat(basePath); os.IsNotExist(err) {
			// No samples for this group
			return result, nil
		}

		err = filepath.WalkDir(basePath, func(filePath string, d fs.DirEntry, err error) error {
			if err != nil {
				// Failed to walk into path, abort early and propagate error
				return err
			}

			// We don't include 'refs' directory here, as it contains dependency references for the group and is purely for
			// samples testing.
			if !d.IsDir() && filepath.Base(filepath.Dir(filePath)) != "refs" {
				filePath, err = filepath.Rel(filepath.Dir(report.samplesPath), filePath)
				if err != nil {
					return eris.Wrapf(err, "getting relative path for %s", filePath)
				}

				filePath = filepath.ToSlash(filePath)
				filePathURL := url.URL{Path: filePath}
				sampleLink := parsedRootURL.ResolveReference(&filePathURL).String()
				sampleFile := filepath.Base(filePath)
				result[sampleFile] = sampleLink
			}

			return nil
		})
		if err != nil {
			return nil, eris.Wrapf(err, "walking through samples directory %s", report.samplesPath)
		}
	}

	return result, nil
}

func (report *ResourceVersionsReport) createItem(
	name astmodel.InternalTypeName,
	armType string,
	armVersion string,
) ResourceVersionsReportResourceItem {
	return ResourceVersionsReportResourceItem{
		name:          name,
		armType:       armType,
		armVersion:    armVersion,
		supportedFrom: report.supportedFrom(name),
	}
}

// generateAPILink returns a link to the API definition for the given resource
func (report *ResourceVersionsReport) generateAPILink(
	name astmodel.InternalTypeName,
) string {
	crdKind := name.Name()
	linkTemplate := report.reportConfiguration.ResourceURLTemplate
	pathTemplate := report.reportConfiguration.ResourcePathTemplate
	if linkTemplate == "" || pathTemplate == "" {
		// One or both of LinkTemplate and PathTemplate are not set, so we can't generate a link
		return crdKind
	}

	docFile := report.resourceDocFile(name)
	if _, err := os.Stat(docFile); eris.Is(err, fs.ErrNotExist) {
		// docFile does not exist, don't build a link
		return crdKind
	}

	link := report.expandPlaceholders(linkTemplate, name)
	return fmt.Sprintf("[%s](%s)", crdKind, link)
}

func (report *ResourceVersionsReport) resourceDocFile(name astmodel.InternalTypeName) string {
	relativePath := report.expandPlaceholders(report.reportConfiguration.ResourcePathTemplate, name)
	baseDir := filepath.Dir(report.reportConfiguration.FullOutputPath())
	return filepath.Join(baseDir, relativePath)
}

func (report *ResourceVersionsReport) expandPlaceholders(template string, rsrc astmodel.InternalTypeName) string {
	crdKind := rsrc.Name()
	crdGroup, crdVersion := rsrc.InternalPackageReference().GroupVersion()

	result := template
	result = strings.ReplaceAll(result, "{group}", crdGroup)
	result = strings.ReplaceAll(result, "{version}", crdVersion)
	result = strings.ReplaceAll(result, "{kind}", crdKind)
	return result
}

func (report *ResourceVersionsReport) generateSampleLink(
	name astmodel.InternalTypeName,
	sampleLinks map[string]string,
) string {
	crdVersion := name.InternalPackageReference().PackageName()
	key := fmt.Sprintf("%s_%s.yaml", crdVersion, strings.ToLower(name.Name()))
	sampleLink, ok := sampleLinks[key]

	if ok {
		// Note: These links are guaranteed to work because of the Taskfile 'controller:verify-samples' target
		return fmt.Sprintf("[View](%s)", sampleLink)
	}

	return "-"
}

func (report *ResourceVersionsReport) supportedFrom(typeName astmodel.InternalTypeName) string {
	supportedFrom, ok := report.objectModelConfiguration.SupportedFrom.Lookup(typeName)
	if !ok {
		// Leave it blank
		return ""
	}

	ver := typeName.InternalPackageReference().Version()

	// Special case for resources that existed prior to GA
	// the `v1api` versions of those resources are only available from "v2.0.0"
	if strings.Contains(ver, v1VersionPrefix) && strings.HasPrefix(supportedFrom, "v2.0.0-") {
		return "v2.0.0"
	}

	return supportedFrom
}

// Read in any front matter present in our output file, so we preserve it when writing out the new file.
// Returns an empty string if the file doesn't exist
func (report *ResourceVersionsReport) readFrontMatter(outputPath string) string {
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return ""
	}

	data, err := os.ReadFile(outputPath)
	if err != nil {
		return ""
	}

	var buffer strings.Builder

	// Copy until we find the end of the frontmatter
	for index, line := range strings.Split(string(data), "\n") {
		buffer.WriteString(line)
		buffer.WriteString("\n")
		if index >= 2 && strings.HasPrefix(line, "---") {
			break
		}
	}

	return buffer.String()
}

// defaultAllResourcesFrontMatter returns the default front-matter for the report if no existing file is present,
// or if it has no front-matter present.
func (report *ResourceVersionsReport) defaultAllResourcesFrontMatter() string {
	var buffer strings.Builder
	buffer.WriteString("---\n")
	buffer.WriteString("title: Supported Resources\n")
	buffer.WriteString("---\n\n")
	return buffer.String()
}

// defaultGroupResourcesFrontMatter returns the default front-matter for the report if no existing file is present,
// or if it has no front-matter present.
func (report *ResourceVersionsReport) defaultGroupResourcesFrontMatter(title string) string {
	var buffer strings.Builder
	buffer.WriteString("---\n")
	buffer.WriteString(fmt.Sprintf("title: %s Supported Resources\n", title))
	buffer.WriteString(fmt.Sprintf("linktitle: %s\n", title))
	buffer.WriteString("no_list: true\n")
	buffer.WriteString("---\n\n")
	return buffer.String()
}

// writeFragment will find the named fragment and write the rendered result to the buffer if it exists.
func (report *ResourceVersionsReport) writeFragment(name string, data any, buffer io.Writer) error {
	report.typoAdvisor.AddTerm(name)

	// Look up the fragement
	fragment, ok := report.availableFragments[name]
	if !ok {
		return nil
	}

	// Create a template based on the fragment
	tmpl, err := template.New(name).Parse(fragment)
	if err != nil {
		return eris.Wrapf(err, "unable to parse template for %s", name)
	}

	// Render the template
	err = tmpl.Execute(buffer, data)
	if err != nil {
		return eris.Wrapf(err, "unable to render template for %s", name)
	}

	return nil
}

type reportMetadataQuality int

const (
	poor   reportMetadataQuality = 10
	good   reportMetadataQuality = 20
	better reportMetadataQuality = 30
	best   reportMetadataQuality = 40
)

// groupTitle returns metadata to use for the given group.
// group is the name of the group to generate metadata for.
// items is the set of items in the group.
// Some resource-providers (incl alertsmanagement) are inconsistent with letter casing, so we
// look for the best information we have.
func (report *ResourceVersionsReport) groupInfo(
	group string,
	items set.Set[ResourceVersionsReportResourceItem],
) *ResourceVersionsReportGroupInfo {
	caser := cases.Title(language.English, cases.NoLower)
	result := &ResourceVersionsReportGroupInfo{
		Group:    group,
		Title:    caser.String(group),
		Provider: caser.String(group),
	}
	currentQuality := poor

	saveCandidate := func(
		title string,
		provider string,
		quality reportMetadataQuality,
	) {
		if quality > currentQuality {
			result.Title = title
			result.Provider = provider
			currentQuality = quality
		}
	}

	// Scan through available resources to find better values for Title & Provider
	for item := range items {
		if item.armType == "" {
			// Didn't find a resource, keep looking
			continue
		}

		// Slice off the part before the first "/" to get the Provider name
		parts := strings.Split(item.armType, "/")
		provider := parts[0]

		const prefix = "Microsoft."
		if strings.HasPrefix(provider, prefix) {
			// If the provider starts with exactly "Microsoft." (with a case-sensitive check),
			// that's our best bet for a properly cased provider name
			saveCandidate(
				strings.TrimPrefix(provider, prefix),
				provider,
				best)
		} else if strings.HasPrefix(strings.ToLower(provider), strings.ToLower(prefix)) {
			// If the provider starts with "Microsoft." (but doesn't match the case exactly),
			// that's not so good, but better than the defaults
			saveCandidate(
				caser.String(provider[len(prefix):]),
				caser.String(provider),
				better)
		} else {
			// Just use what we have
			saveCandidate(provider, provider, good)
		}
	}

	return result
}
