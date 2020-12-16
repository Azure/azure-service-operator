/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"github.com/Azure/k8s-infra/hack/generator/pkg/reporting"
	"github.com/pkg/errors"
	"io/ioutil"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"os"
	"path"
	"strings"
)

// reportOnTypesAndVersions creates a pipeline stage that removes any wrapper types prior to actual code generation
func reportOnTypesAndVersions(configuration *config.Configuration) PipelineStage {

	return MakePipelineStage(
		"reportTypesAndVersions",
		"Generate reports on types and versions in each package",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			report := NewPackagesMatrixReport()
			report.Summarize(types)
			err := report.WriteTo(configuration.OutputPath)
			return types, err
		})
}

type PackagesMatrixReport struct {
	// A separate table for each package
	tables map[string]*reporting.Table
}

func NewPackagesMatrixReport() *PackagesMatrixReport {
	return &PackagesMatrixReport{
		tables: make(map[string]*reporting.Table),
	}
}

func (report *PackagesMatrixReport) Summarize(types astmodel.Types) {
	for _, t := range types {
		typeName := t.Name().Name()
		packageName := report.ServiceName(t.Name().PackageReference)
		packageVersion := t.Name().PackageReference.PackageName()
		table, ok := report.tables[packageName]
		if !ok {
			table = reporting.NewTable()
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

func (report *PackagesMatrixReport) WriteTableTo(table *reporting.Table, pkg string, outputPath string) error {
	table.SortColumns(func(left string, right string) bool {
		return left < right
	})
	table.SortRows(func(top string, bottom string) bool {
		return top < bottom
	})

	var buffer strings.Builder
	table.WriteTo(&buffer)

	outputFolder := path.Join(outputPath, pkg)
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err = os.MkdirAll(outputFolder, 0700)
		if err != nil {
			return errors.Wrapf(err, "Unable to create directory %q", outputFolder)
		}
	}

	destination := path.Join(outputFolder, "versions_matrix.md")
	return ioutil.WriteFile(destination, []byte(buffer.String()), 0600)
}
