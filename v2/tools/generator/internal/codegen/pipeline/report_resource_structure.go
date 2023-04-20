/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
)

// ReportResourceStructureStageId is the unique identifier for this stage
const ReportResourceStructureStageId = "reportResourceStructure"

// ReportResourceStructure creates a pipeline stage that reports the structure of resources in each package
func ReportResourceStructure(configuration *config.Configuration) *Stage {
	return NewStage(
		ReportResourceStructureStageId,
		"Reports the structure of resources in each package",
		func(ctx context.Context, state *State) (*State, error) {
			report := NewResourceStructureReport(state.Definitions())
			err := report.SaveReports(configuration.FullTypesOutputPath())
			return state, err
		})
}

type ResourceStructureReport struct {
	lists map[astmodel.PackageReference]astmodel.TypeDefinitionSet // A separate list of resources for each package
}

func NewResourceStructureReport(defs astmodel.TypeDefinitionSet) *ResourceStructureReport {
	result := &ResourceStructureReport{
		lists: make(map[astmodel.PackageReference]astmodel.TypeDefinitionSet),
	}

	result.summarize(defs)
	return result
}

// SaveReports writes the reports to the specified files
func (report *ResourceStructureReport) SaveReports(baseFolder string) error {
	for pkg, defs := range report.lists {
		filePath := report.getFilePath(baseFolder, pkg)
		err := report.saveReport(filePath, defs)
		if err != nil {
			return err
		}
	}

	return nil
}

// summarize collates a list of all definitions, grouped by package
func (report *ResourceStructureReport) summarize(definitions astmodel.TypeDefinitionSet) {
	for name, def := range definitions {
		pkg := name.PackageReference

		if _, ok := report.lists[pkg]; !ok {
			report.lists[pkg] = make(astmodel.TypeDefinitionSet)
		}

		report.lists[pkg].Add(def)
	}
}

func (report *ResourceStructureReport) getFilePath(baseFolder string, pkg astmodel.PackageReference) string {
	grp, ver := pkg.GroupVersion()
	return filepath.Join(baseFolder, grp, ver, "structure.txt")
}

func (report *ResourceStructureReport) saveReport(filePath string, defs astmodel.TypeDefinitionSet) error {
	rpt := reporting.NewTypeCatalogReport(defs)
	rpt.InlineTypes()
	rpt.AddHeader(astmodel.CodeGenerationComments...)
	err := rpt.SaveTo(filePath)
	return errors.Wrapf(err, "unable to save type catalog report to %q", filePath)
}
