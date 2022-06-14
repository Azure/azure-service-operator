/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

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
			report := NewResourceVersionsReport(state.Definitions(), configuration.ObjectModelConfiguration)

			err := report.WriteTo(configuration.FullTypesOutputPath(), configuration.SamplesURL)
			if err != nil {
				return nil, err
			}

			err = configuration.ObjectModelConfiguration.VerifySupportedFromConsumed()
			return state, err
		})
}

type ResourceVersionsReport struct {
	objectModelConfiguration *config.ObjectModelConfiguration
	groups                   set.Set[string]                       // A set of all our groups
	kinds                    map[string]astmodel.TypeDefinitionSet // For each group, the set of all available resources
	// A separate list of resources for each package
	lists map[astmodel.PackageReference][]astmodel.TypeDefinition
}

func NewResourceVersionsReport(
	definitions astmodel.TypeDefinitionSet,
	cfg *config.ObjectModelConfiguration,
) *ResourceVersionsReport {
	result := &ResourceVersionsReport{
		objectModelConfiguration: cfg,
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
func (report *ResourceVersionsReport) WriteTo(outputPath string, samplesURL string) error {
	var buffer strings.Builder
	err := report.WriteToBuffer(&buffer, samplesURL)
	if err != nil {
		return errors.Wrapf(err, "writing versions report to %s", outputPath)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		err = os.MkdirAll(outputPath, 0o700)
		if err != nil {
			return errors.Wrapf(err, "Unable to create directory %q", outputPath)
		}
	}

	destination := path.Join(outputPath, "resources.md")
	return ioutil.WriteFile(destination, []byte(buffer.String()), 0o600)
}

// WriteToBuffer creates the report in the provided buffer
func (report *ResourceVersionsReport) WriteToBuffer(buffer *strings.Builder, samplesURL string) error {
	buffer.WriteString("---\n")
	buffer.WriteString("title: Supported Resources\n")
	buffer.WriteString("---\n\n")
	buffer.WriteString("These are the resources with Azure Service Operator support committed to our **main** branch, ")
	buffer.WriteString("grouped by the originating ARM service. ")
	buffer.WriteString("(Newly supported resources will appear in this list prior to inclusion in any ASO release.)\n\n")

	// Sort groups into alphabetical order
	groups := set.AsSortedSlice(report.groups)

	var errs []error
	for _, svc := range groups {
		buffer.WriteString(fmt.Sprintf("## %s\n\n", strings.Title(svc)))
		table, err := report.createTable(report.kinds[svc], svc, samplesURL)
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
	group string,
	samplesURL string,
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

	errs := make([]error, 0, len(toIterate))
	for _, rsrc := range toIterate {
		resourceType := astmodel.MustBeResourceType(rsrc.Type())

		crdVersion := rsrc.Name().PackageReference.PackageName()
		armVersion := strings.Trim(resourceType.APIVersionEnumValue().Value, "\"")
		if armVersion == "" {
			armVersion = crdVersion
		}

		sample := report.generateSampleLink(samplesURL, group, rsrc)
		supportedFrom, err := report.generateSupportedFrom(rsrc.Name())
		errs = append(errs, err)

		result.AddRow(
			rsrc.Name().Name(),
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

func (report *ResourceVersionsReport) generateSampleLink(samplesURL string, group string, rsrc astmodel.TypeDefinition) string {
	crdVersion := rsrc.Name().PackageReference.PackageName()
	if samplesURL != "" {
		// Note: These links are guaranteed to work because of the Taskfile 'controller:verify-samples' target
		samplePath := fmt.Sprintf("%s/%s/%s_%s.yaml", samplesURL, group, crdVersion, strings.ToLower(rsrc.Name().Name()))
		return fmt.Sprintf("[View](%s)", samplePath)
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
	if strings.Contains(ver, "v1beta") && strings.HasPrefix(supportedFrom, "v2.0.0-alpha") {
		return "v2.0.0-beta.0", nil
	}

	return supportedFrom, nil
}
