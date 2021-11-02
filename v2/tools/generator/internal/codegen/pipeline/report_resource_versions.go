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

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// ReportResourceVersionsStageID is the unique identifier of this stage
const ReportResourceVersionsStageID = "reportResourceVersions"

// ReportResourceVersions creates a pipeline stage that generates a report listing all generated resources.
func ReportResourceVersions(configuration *config.Configuration) Stage {
	return MakeStage(
		ReportResourceVersionsStageID,
		"Generate a report listing all the resources generated",
		func(ctx context.Context, state *State) (*State, error) {
			report := NewResourceVersionsReport(state.Types())
			err := report.WriteTo(configuration.FullTypesOutputPath(), configuration.SamplesURL)
			return state, err
		})
}

type ResourceVersionsReport struct {
	// A separate list of resources for each package
	lists map[astmodel.PackageReference][]astmodel.TypeDefinition
}

func NewResourceVersionsReport(types astmodel.Types) *ResourceVersionsReport {
	result := &ResourceVersionsReport{
		lists: make(map[astmodel.PackageReference][]astmodel.TypeDefinition),
	}

	result.summarize(types)
	return result
}

// summarize collates a list of all resources, grouped by package
func (r *ResourceVersionsReport) summarize(types astmodel.Types) {
	resources := astmodel.FindResourceTypes(types)
	for _, rsrc := range resources {
		name := rsrc.Name()
		pkg := name.PackageReference
		r.lists[pkg] = append(r.lists[pkg], rsrc)
	}
}

// WriteTo creates a file containing the generated report
func (r *ResourceVersionsReport) WriteTo(outputPath string, samplesURL string) error {
	var buffer strings.Builder
	r.WriteToBuffer(&buffer, samplesURL)

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		err = os.MkdirAll(outputPath, 0700)
		if err != nil {
			return errors.Wrapf(err, "Unable to create directory %q", outputPath)
		}
	}

	destination := path.Join(outputPath, "resources.md")
	return ioutil.WriteFile(destination, []byte(buffer.String()), 0600)
}

// WriteToBuffer creates the report in the provided buffer
func (r *ResourceVersionsReport) WriteToBuffer(buffer *strings.Builder, samplesURL string) {
	// Sort packages into increasing order
	// Skip storage versions
	var packages []astmodel.PackageReference
	for pkg := range r.lists {
		if !astmodel.IsStoragePackageReference(pkg) {
			packages = append(packages, pkg)
		}
	}

	astmodel.SortPackageReferencesByPathAndVersion(packages)

	lastService := ""
	for _, pkg := range packages {

		// Write a header for each service
		svc := r.serviceName(pkg)
		if lastService != svc {
			buffer.WriteString(fmt.Sprintf("### %s\n\n", svc))
			lastService = svc
		}

		// For each version, write a header
		// We use the API version of the first resource in each set, as this reflects the ARM API Version
		resources := r.lists[pkg]
		sort.Slice(
			resources,
			func(i, j int) bool {
				return resources[i].Name().Name() < resources[j].Name().Name()
			})
		firstDef := resources[0]
		firstResource := astmodel.MustBeResourceType(firstDef.Type())
		buffer.WriteString(
			fmt.Sprintf(
				"\n#### ARM version %s\n\n",
				strings.Trim(firstResource.APIVersionEnumValue().Value, "\"")))

		// write an alphabetical list of resources

		for _, rsrc := range resources {
			if samplesURL != "" {
				// Note: These links are guaranteed to work because of the Taskfile 'controller:verify-samples' target
				samplePath := fmt.Sprintf("%s/%s/%s_%s.yaml", samplesURL, svc, pkg.PackageName(), strings.ToLower(rsrc))
				buffer.WriteString(fmt.Sprintf("- %s ([sample](%s))\n", rsrc, samplePath))
			} else {
				buffer.WriteString(fmt.Sprintf("- %s\n", rsrc))
			}
		}

		buffer.WriteString(fmt.Sprintf("\nUse CRD version `%s`\n", firstDef.Name().PackageReference.PackageName()))
		buffer.WriteString("\n")
	}
}

func (r *ResourceVersionsReport) serviceName(ref astmodel.PackageReference) string {
	pathBits := strings.Split(ref.PackagePath(), "/")
	index := len(pathBits) - 1
	if index > 0 {
		index--
	}

	return pathBits[index]
}
