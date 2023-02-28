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
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
)

// ReportOnTypesAndVersionsStageID is the unique identifier of this stage
const ReportOnTypesAndVersionsStageID = "reportTypesAndVersions"

// ReportOnTypesAndVersions creates a pipeline stage that generates a report for each group showing a matrix of all
// types and versions
func ReportOnTypesAndVersions(configuration *config.Configuration) *Stage {
	return NewLegacyStage(
		ReportOnTypesAndVersionsStageID,
		"Generate reports on types and versions in each package",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			report := NewPackagesMatrixReport()
			report.Summarize(definitions)
			err := report.WriteTo(configuration.FullTypesOutputPath())
			return definitions, err
		})
}

type PackagesMatrixReport struct {
	// A separate table for each package
	tables map[string]*reporting.SparseTable
}

func NewPackagesMatrixReport() *PackagesMatrixReport {
	return &PackagesMatrixReport{
		tables: make(map[string]*reporting.SparseTable),
	}
}

func (report *PackagesMatrixReport) Summarize(definitions astmodel.TypeDefinitionSet) {
	for _, t := range definitions {
		typeName := t.Name().Name()
		packageName := report.ServiceName(t.Name().PackageReference)
		packageVersion := t.Name().PackageReference.PackageName()
		table, ok := report.tables[packageName]
		if !ok {
			table = reporting.NewSparseTable(fmt.Sprintf("Type Definitions in package %q", packageName))
			report.tables[packageName] = table
		}

		table.SetCell(typeName, packageVersion, packageVersion)
	}
}

func (report *PackagesMatrixReport) WriteTo(outputPath string) error {
	var errs []error
	for pkg, table := range report.tables {
		err := report.WriteTableTo(table, pkg, outputPath)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

func (report *PackagesMatrixReport) ServiceName(ref astmodel.PackageReference) string {
	pathBits := strings.Split(ref.PackagePath(), "/")
	index := len(pathBits) - 1
	if index > 0 {
		index--
	}

	return pathBits[index]
}

func (report *PackagesMatrixReport) WriteTableTo(table *reporting.SparseTable, pkg string, outputPath string) error {
	table.SortColumns(func(left string, right string) bool {
		return left < right
	})
	table.SortRows(func(top string, bottom string) bool {
		return top < bottom
	})

	var buffer strings.Builder
	table.WriteTo(&buffer)

	outputFolder := filepath.Join(outputPath, pkg)
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err = os.MkdirAll(outputFolder, 0700)
		if err != nil {
			return errors.Wrapf(err, "Unable to create directory %q", outputFolder)
		}
	}

	destination := filepath.Join(outputFolder, "versions_matrix.md")
	return ioutil.WriteFile(destination, []byte(buffer.String()), 0600)
}
