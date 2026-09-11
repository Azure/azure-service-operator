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
	"sort"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
)

// ReportPropertyChangesStageID is the unique identifier for this stage
const ReportPropertyChangesStageID = "reportPropertyChanges"

// ReportPropertyChanges creates a pipeline stage that reports, for each generated resource, the
// property-level changes between it and its "next" version, as determined by the version-conversion
// graph. One property-changes.md file is generated per package alongside structure.txt.
// ARM, webhook, and compatibility packages are excluded.
func ReportPropertyChanges(configuration *config.Configuration) *Stage {
	stage := NewStage(
		ReportPropertyChangesStageID,
		"Reports property changes between each resource and its next version",
		func(ctx context.Context, state *State) (*State, error) {
			graph, err := GetStateData[*storage.ConversionGraph](state, ConversionGraphInfo)
			if err != nil {
				return nil, eris.Wrapf(err, "couldn't find conversion graph")
			}

			reporter := NewPropertyChangesReporter(state.Definitions(), graph, configuration.ObjectModelConfiguration)
			err = reporter.SaveReports(configuration.FullTypesOutputPath())
			return state, err
		},
	)

	stage.RequiresPrerequisiteStages(CreateConversionGraphStageID)
	return stage.UsedFor(ARMTarget)
}

// PropertyChangesReporter identifies the resources needing a property changes report, resolves each
// one's "next" version via the conversion graph, and writes the resulting reports.
type PropertyChangesReporter struct {
	definitions   astmodel.TypeDefinitionSet
	graph         *storage.ConversionGraph
	configuration *config.ObjectModelConfiguration
}

// NewPropertyChangesReporter creates a new PropertyChangesReporter.
func NewPropertyChangesReporter(
	definitions astmodel.TypeDefinitionSet,
	graph *storage.ConversionGraph,
	configuration *config.ObjectModelConfiguration,
) *PropertyChangesReporter {
	return &PropertyChangesReporter{
		definitions:   definitions,
		graph:         graph,
		configuration: configuration,
	}
}

// SaveReports writes one property-changes.md file per package containing non-hub resources.
func (r *PropertyChangesReporter) SaveReports(baseFolder string) error {
	pairsByPackage := make(map[astmodel.InternalPackageReference][]reporting.ResourceVersionPair)
	for _, resource := range r.findResources() {
		nextResource, err := r.graph.FindNextType(resource, r.definitions)
		if err != nil {
			return eris.Wrapf(err, "finding next version of %s", resource)
		}

		if nextResource.IsEmpty() {
			// resource has no next version (it's the hub) - no report to generate
			continue
		}

		pkg := resource.InternalPackageReference()
		pairsByPackage[pkg] = append(
			pairsByPackage[pkg],
			reporting.ResourceVersionPair{This: resource, Next: nextResource},
		)
	}

	packages := make([]astmodel.InternalPackageReference, 0, len(pairsByPackage))
	for pkg := range pairsByPackage {
		packages = append(packages, pkg)
	}

	sort.Slice(packages, func(i, j int) bool {
		return packages[i].PackagePath() < packages[j].PackagePath()
	})

	for _, pkg := range packages {
		filePath := filepath.Join(baseFolder, pkg.FolderPath(), "property-changes.md")
		err := r.saveReport(filePath, pairsByPackage[pkg])
		if err != nil {
			return err
		}
	}

	return nil
}

// findResources returns every resource except those in ARM, webhook, or compatibility packages,
// sorted by package path and then by name for deterministic output.
func (r *PropertyChangesReporter) findResources() []astmodel.InternalTypeName {
	var result []astmodel.InternalTypeName
	for name, def := range r.definitions {
		pkg := name.PackageReference()
		if astmodel.IsARMPackageReference(pkg) ||
			astmodel.IsWebhookPackageReference(pkg) ||
			astmodel.IsCompatPackageReference(pkg) {
			continue
		}

		if _, ok := astmodel.AsResourceType(def.Type()); !ok {
			continue
		}

		result = append(result, name)
	}

	sort.Slice(result, func(i, j int) bool {
		leftPath := result[i].InternalPackageReference().PackagePath()
		rightPath := result[j].InternalPackageReference().PackagePath()
		if leftPath != rightPath {
			return leftPath < rightPath
		}

		return result[i].Name() < result[j].Name()
	})

	return result
}

func (r *PropertyChangesReporter) saveReport(
	filePath string,
	resources []reporting.ResourceVersionPair,
) error {
	rpt := reporting.NewPropertyChangesReport(
		resources,
		r.definitions,
		r.configuration.TypeNameInNextVersion.Lookup,
		r.configuration.PropertyNameInNextVersion.Lookup,
	)

	for _, line := range astmodel.CodeGenerationComments {
		// Wrapped as an HTML comment so it's invisible when the Markdown is rendered, while still
		// containing the exact text the generated-file cleanup mechanism looks for.
		rpt.AddHeader(fmt.Sprintf("<!-- %s -->", line))
	}

	// The package folder should already exist (it holds the generated types), but create it
	// defensively in case this stage ever runs before anything else has written to it.
	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		return eris.Wrapf(err, "creating folder for %q", filePath)
	}

	err := rpt.SaveTo(filePath)
	return eris.Wrapf(err, "unable to save property changes report to %q", filePath)
}
