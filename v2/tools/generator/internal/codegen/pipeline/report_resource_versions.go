/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/internal/set"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
)

// ReportResourceVersionsStageID is the unique identifier of this stage
const ReportResourceVersionsStageID = "reportResourceVersions"

// ReportResourceVersions creates a pipeline stage that generates a report listing all generated resources.
func ReportResourceVersions(configuration *config.Configuration) *Stage {
	return NewStage(
		ReportResourceVersionsStageID,
		"Generate a report listing all the resources generated",
		func(ctx context.Context, state *State) (*State, error) {
			report := NewResourceVersionsReport(state.Definitions(), configuration)

			err := report.WriteTo(configuration.SupportedResourcesReport.FullOutputPath())
			if err != nil {
				return nil, err
			}

			err = configuration.ObjectModelConfiguration.VerifySupportedFromConsumed()
			return state, err
		})
}

type ResourceVersionsReport struct {
	reportConfiguration      *config.SupportedResourcesReport
	objectModelConfiguration *config.ObjectModelConfiguration
	rootUrl                  string
	samplesPath              string
	apiDocsUrlTemplate       string                                // Template for generating API URLs
	groups                   set.Set[string]                       // A set of all our groups
	kinds                    map[string]astmodel.TypeDefinitionSet // For each group, the set of all available resources
	frontMatter              string                                // Front matter to be inserted at the top of the report
	// A separate list of resources for each package
	lists map[astmodel.PackageReference][]astmodel.TypeDefinition
}

func NewResourceVersionsReport(
	definitions astmodel.TypeDefinitionSet,
	cfg *config.Configuration,
) *ResourceVersionsReport {
	result := &ResourceVersionsReport{
		reportConfiguration:      cfg.SupportedResourcesReport,
		objectModelConfiguration: cfg.ObjectModelConfiguration,
		rootUrl:                  cfg.RootURL,
		apiDocsUrlTemplate:       cfg.ApiDocsUrlTemplate,
		samplesPath:              cfg.SamplesPath,
		groups:                   set.Make[string](),
		kinds:                    make(map[string]astmodel.TypeDefinitionSet),
		lists:                    make(map[astmodel.PackageReference][]astmodel.TypeDefinition),
	}

	result.summarize(definitions)
	return result
}

// summarize collates a list of all resources, grouped by package
func (report *ResourceVersionsReport) summarize(definitions astmodel.TypeDefinitionSet) {
	resources := astmodel.FindResourceDefinitions(definitions)
	for _, rsrc := range resources {
		name := rsrc.Name()
		pkg := name.PackageReference
		if astmodel.IsStoragePackageReference(pkg) {
			// Skip storage versions - they're an implementation detail
			continue
		}

		grp, _ := pkg.GroupVersion()
		report.groups.Add(grp)
		report.lists[pkg] = append(report.lists[pkg], rsrc)

		defs, ok := report.kinds[grp]
		if !ok {
			defs = make(astmodel.TypeDefinitionSet)
			report.kinds[grp] = defs
		}

		defs.Add(rsrc)
	}
}

// WriteTo creates a file containing the generated report
func (report *ResourceVersionsReport) WriteTo(outputFile string) error {

	klog.V(1).Infof("Writing report to %s", outputFile)
	report.frontMatter = report.readFrontMatter(outputFile)

	var buffer strings.Builder
	err := report.WriteToBuffer(&buffer)
	if err != nil {
		return errors.Wrapf(err, "writing versions report to %s", outputFile)
	}

	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		err = os.MkdirAll(outputFile, 0o700)
		if err != nil {
			return errors.Wrapf(err, "Unable to create directory %q", outputFile)
		}
	}

	return ioutil.WriteFile(outputFile, []byte(buffer.String()), 0o600)
}

// WriteToBuffer creates the report in the provided buffer
func (report *ResourceVersionsReport) WriteToBuffer(buffer *strings.Builder) error {

	if report.frontMatter != "" {
		buffer.WriteString(report.frontMatter)
	} else {
		buffer.WriteString(report.defaultFrontMatter())
	}

	buffer.WriteString(report.reportConfiguration.Introduction)
	buffer.WriteString("\n\n")

	// Sort groups into alphabetical order
	groups := set.AsSortedSlice(report.groups)

	var errs []error
	for _, svc := range groups {
		buffer.WriteString(fmt.Sprintf("## %s\n\n", strings.Title(svc)))
		table, err := report.createTable(report.kinds[svc])
		if err != nil {
			errs = append(errs, err)
			continue
		}

		table.WriteTo(buffer)
		buffer.WriteString("\n")
	}

	return kerrors.NewAggregate(errs)
}

func (report *ResourceVersionsReport) createTable(
	resources astmodel.TypeDefinitionSet,
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

	toIterate := resources.AsSlice()
	sort.Slice(toIterate, func(i, j int) bool {
		left := toIterate[i].Name()
		right := toIterate[j].Name()
		if left.Name() != right.Name() {
			return left.Name() < right.Name()
		}

		// Reversed parameters because we want more recent versions listed first
		return astmodel.ComparePathAndVersion(right.PackageReference.PackagePath(), left.PackageReference.PackagePath())
	})

	sampleLinks := make(map[string]string)
	if report.rootUrl != "" {
		parsedRootURL, err := url.Parse(report.rootUrl)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing rootUrl %s", report.rootUrl)
		}
		err = filepath.WalkDir(report.samplesPath, func(filePath string, d fs.DirEntry, err error) error {
			// We don't include 'refs' directory here, as it contains dependency references for the group and is purely for
			// samples testing.
			if !d.IsDir() && filepath.Base(filepath.Dir(filePath)) != "refs" {
				filePath = filepath.ToSlash(filePath)
				filePathURL := url.URL{Path: filePath}
				sampleLink := parsedRootURL.ResolveReference(&filePathURL).String()
				sampleFile := filepath.Base(filePath)
				sampleLinks[sampleFile] = sampleLink
			}

			return nil
		})
		if err != nil {
			return nil, errors.Wrapf(err, "walking through samples directory %s", report.samplesPath)
		}
	}

	errs := make([]error, 0, len(toIterate))
	for _, rsrc := range toIterate {
		resourceType := astmodel.MustBeResourceType(rsrc.Type())

		crdVersion := rsrc.Name().PackageReference.PackageName()
		armVersion := strings.Trim(resourceType.APIVersionEnumValue().Value, "\"")
		if armVersion == "" {
			armVersion = crdVersion
		}

		api := report.generateApiLink(rsrc)
		sample := report.generateSampleLink(rsrc, sampleLinks)
		supportedFrom, err := report.generateSupportedFrom(rsrc.Name())
		errs = append(errs, err)

		result.AddRow(
			api,
			armVersion,
			crdVersion,
			supportedFrom,
			sample)
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, errors.Wrap(err, "generating versions report")
	}

	return result, nil
}

// generateApiLink returns a link to the API definition for the given resource
func (report *ResourceVersionsReport) generateApiLink(rsrc astmodel.TypeDefinition) string {

	crdKind := rsrc.Name().Name()
	if report.apiDocsUrlTemplate == "" {
		return crdKind
	}

	crdGroup, crdVersion := rsrc.Name().PackageReference.GroupVersion()

	link := report.apiDocsUrlTemplate
	link = strings.Replace(link, "{group}", crdGroup, -1)
	link = strings.Replace(link, "{version}", crdVersion, -1)
	link = strings.Replace(link, "{kind}", crdKind, -1)

	return fmt.Sprintf("[%s](%s)", crdKind, link)
}

func (report *ResourceVersionsReport) generateSampleLink(rsrc astmodel.TypeDefinition, sampleLinks map[string]string) string {
	crdVersion := rsrc.Name().PackageReference.PackageName()
	key := fmt.Sprintf("%s_%s.yaml", crdVersion, strings.ToLower(rsrc.Name().Name()))
	sampleLink, ok := sampleLinks[key]

	if ok {
		// Note: These links are guaranteed to work because of the Taskfile 'controller:verify-samples' target
		return fmt.Sprintf("[View](%s)", sampleLink)
	}

	return "-"
}

func (report *ResourceVersionsReport) generateSupportedFrom(typeName astmodel.TypeName) (string, error) {
	supportedFrom, err := report.objectModelConfiguration.LookupSupportedFrom(typeName)
	if err != nil {
		return "", err
	}

	_, ver := typeName.PackageReference.GroupVersion()

	// Special case for resources that existed prior to beta.0
	// the `v1beta` versions of those resources are only available from "beta.0"
	if strings.Contains(ver, v1betaVersionPrefix) && strings.HasPrefix(supportedFrom, "v2.0.0-alpha") {
		return "v2.0.0-beta.0", nil
	}

	return supportedFrom, nil
}

// Read in any front matter present in our output file, so we preserve it when writing out the new file.
// Returns an empty string if the file doesn't exist
func (report *ResourceVersionsReport) readFrontMatter(outputPath string) string {

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return ""
	}

	klog.V(2).Infof("Reading front matter from %s", outputPath)
	data, err := ioutil.ReadFile(outputPath)
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

// defaultFrontMatter returns the default front-matter for the report if no existing file is present,
// or if it has no front-matter present.
func (report *ResourceVersionsReport) defaultFrontMatter() string {
	var buffer strings.Builder
	buffer.WriteString("---\n")
	buffer.WriteString("title: Supported Resources\n")
	buffer.WriteString("---\n\n")
	return buffer.String()
}
