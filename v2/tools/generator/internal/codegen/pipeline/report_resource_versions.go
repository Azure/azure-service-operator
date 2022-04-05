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
	groups []string                              // A sorted slice of all our groups
	kinds  map[string]astmodel.TypeDefinitionSet // For each group, the set of all available resources
	// A separate list of resources for each package
	lists map[astmodel.PackageReference][]astmodel.TypeDefinition
}

func NewResourceVersionsReport(definitions astmodel.TypeDefinitionSet) *ResourceVersionsReport {
	result := &ResourceVersionsReport{
		kinds: make(map[string]astmodel.TypeDefinitionSet),
		lists: make(map[astmodel.PackageReference][]astmodel.TypeDefinition),
	}

	result.summarize(definitions)
	return result
}

// summarize collates a list of all resources, grouped by package
func (r *ResourceVersionsReport) summarize(definitions astmodel.TypeDefinitionSet) {
	resources := astmodel.FindResourceDefinitions(definitions)
	for _, rsrc := range resources {
		name := rsrc.Name()
		pkg := name.PackageReference
		grp, _, ok := pkg.GroupVersion()
		if !ok {
			panic(fmt.Sprintf("external package reference not expected on definition"))
		}

		r.lists[pkg] = append(r.lists[pkg], rsrc)

		set, ok := r.kinds[grp]
		if !ok {
			set = make(astmodel.TypeDefinitionSet)
			r.kinds[grp]=set
		}

		set.Add(rsrc)
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

	buffer.WriteString("# Supported Resources\n\n")
	buffer.WriteString("These are the resources with Azure Service Operator support committed to our **main** branch, ")
	buffer.WriteString("grouped by the originating ARM service. ")
	buffer.WriteString("(Newly supported resources will appear in this list prior to inclusion in any ASO release.)\n\n")

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
			buffer.WriteString(fmt.Sprintf("## %s\n\n", svc))
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
		armVersion := strings.Trim(firstResource.APIVersionEnumValue().Value, "\"")
		crdVersion := firstDef.Name().PackageReference.PackageName()
		if armVersion == "" {
			armVersion = crdVersion
		}

		buffer.WriteString(
			fmt.Sprintf(
				"\n### ARM version %s\n\n",
				armVersion))

		// write an alphabetical list of resources

		for _, rsrc := range resources {
			rsrcName := rsrc.Name().Name()
			if samplesURL != "" {
				// Note: These links are guaranteed to work because of the Taskfile 'controller:verify-samples' target
				samplePath := fmt.Sprintf("%s/%s/%s_%s.yaml", samplesURL, svc, pkg.PackageName(), strings.ToLower(rsrcName))
				buffer.WriteString(fmt.Sprintf("- %s ([sample](%s))\n", rsrcName, samplePath))
			} else {
				buffer.WriteString(fmt.Sprintf("- %s\n", rsrc.Name()))
			}
		}

		buffer.WriteString(fmt.Sprintf("\nUse CRD version `%s`\n", crdVersion))
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
