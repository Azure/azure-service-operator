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

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/pkg/errors"

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
			report := NewResourceVersionsReport(state.Definitions())
			err := report.WriteTo(configuration.FullTypesOutputPath(), configuration.SamplesURL)
			return state, err
		})
}

type ResourceVersionsReport struct {
	groups set.Set[string]                       // A set of all our groups
	kinds  map[string]astmodel.TypeDefinitionSet // For each group, the set of all available resources
	// A separate list of resources for each package
	lists map[astmodel.PackageReference][]astmodel.TypeDefinition
}

func NewResourceVersionsReport(definitions astmodel.TypeDefinitionSet) *ResourceVersionsReport {
	result := &ResourceVersionsReport{
		groups: set.Make[string](),
		kinds:  make(map[string]astmodel.TypeDefinitionSet),
		lists:  make(map[astmodel.PackageReference][]astmodel.TypeDefinition),
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

		set, ok := report.kinds[grp]
		if !ok {
			set = make(astmodel.TypeDefinitionSet)
			report.kinds[grp] = set
		}

		set.Add(rsrc)
	}
}

// WriteTo creates a file containing the generated report
func (report *ResourceVersionsReport) WriteTo(outputPath string, samplesURL string) error {
	var buffer strings.Builder
	report.WriteToBuffer(&buffer, samplesURL)

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
func (report *ResourceVersionsReport) WriteToBuffer(buffer *strings.Builder, samplesURL string) {
	buffer.WriteString("---\n")
	buffer.WriteString("title: Supported Resources\n")
	buffer.WriteString("---\n\n")
	buffer.WriteString("These are the resources with Azure Service Operator support committed to our **main** branch, ")
	buffer.WriteString("grouped by the originating ARM service. ")
	buffer.WriteString("(Newly supported resources will appear in this list prior to inclusion in any ASO release.)\n\n")

	// Sort groups into increasing order
	groups := set.AsSortedSlice(report.groups)

	for _, svc := range groups {
		buffer.WriteString(fmt.Sprintf("## %s\n\n", strings.Title(svc)))
		table := report.createTable(report.kinds[svc], svc, samplesURL)
		table.WriteTo(buffer)
		buffer.WriteString("\n")
	}
}

func (report *ResourceVersionsReport) createTable(
	resources astmodel.TypeDefinitionSet,
	group string,
	samplesURL string,
) *reporting.MarkdownTable {
	const (
		name       = "Resource"
		armVersion = "ARM Version"
		crdVersion = "CRD Version"
		sample     = "Sample"
	)

	result := reporting.NewMarkdownTable(
		name,
		armVersion,
		crdVersion,
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

	for _, rsrc := range toIterate {
		resourceType := astmodel.MustBeResourceType(rsrc.Type())

		crdVersion := rsrc.Name().PackageReference.PackageName()
		armVersion := strings.Trim(resourceType.APIVersionEnumValue().Value, "\"")
		if armVersion == "" {
			armVersion = crdVersion
		}

		var sample string
		if samplesURL != "" {
			// Note: These links are guaranteed to work because of the Taskfile 'controller:verify-samples' target
			samplePath := fmt.Sprintf("%s/%s/%s_%s.yaml", samplesURL, group, crdVersion, strings.ToLower(rsrc.Name().Name()))
			sample = fmt.Sprintf("[View](%s)", samplePath)
		} else {
			sample = "-"
		}

		result.AddRow(
			rsrc.Name().Name(),
			armVersion,
			crdVersion,
			sample)
	}

	return result
}
