/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
	"io/ioutil"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/pipeline"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// debugReporter is a helper for generating debug logs during the code generation process.
type debugReporter struct {
	outputFolder  string
	groupSelector config.StringMatcher
}

// newDebugReporter creates a new debugReporter.
// groupSelector specifies which groups to include (may include wildcards).
// outputFolder specifies where to write the debug output.
func newDebugReporter(groupSelector string, outputFolder string) *debugReporter {
	return &debugReporter{
		groupSelector: config.NewStringMatcher(groupSelector),
		outputFolder:  outputFolder,
	}
}

func (dr *debugReporter) ReportStage(stage int, description string, state *pipeline.State) error {
	report, err := dr.createReport(state)
	if err != nil {
		return errors.Wrapf(err, "failed to create debug report for stage %d", stage)
	}

	// Save the template to the debug directory
	filename := dr.createFileName(stage, description)
	err = ioutil.WriteFile(filename, []byte(report), 0600)
	return errors.Wrapf(err, "failed to write debug log to %s", filename)
}

func (dr *debugReporter) createReport(state *pipeline.State) (string, error) {
	var buffer strings.Builder

	packages := dr.findPackages(state.Definitions())
	for _, pkg := range packages {
		grp, _ := pkg.GroupVersion()
		if !dr.groupSelector.Matches(grp) {
			// Skip this package
			continue
		}

		rpt := reporting.NewStructureReport(pkg.PackagePath())
		dr.writeDefinitions(rpt, dr.inPackage(pkg, state.Definitions()))

		err := rpt.SaveTo(&buffer)
		if err != nil {
			return "", errors.Wrapf(err, "failed to create debug report for %s", pkg.PackagePath())
		}
	}

	return buffer.String(), nil
}

// writeDefinitions writes the definitions to the debug report.
// rpt is the debug report to write to.
// definitions is the set of definitions to write.
// Definitions are written in alphabetical order.
func (dr *debugReporter) writeDefinitions(rpt *reporting.StructureReport, definitions astmodel.TypeDefinitionSet) {
	defs := definitions.AsSlice()
	sort.Slice(defs, func(i, j int) bool {
		return defs[i].Name().Name() < defs[j].Name().Name()
	})

	for _, d := range defs {
		dr.writeDefinition(rpt, d)
	}
}

// writeDefinition writes the definition to the debug report.
// rpt is the debug report to write to.
// definition is the definition to write.
func (dr *debugReporter) writeDefinition(rpt *reporting.StructureReport, definition astmodel.TypeDefinition) {
	name := definition.Name()
	t := definition.Type()
	sub := rpt.Add(
		fmt.Sprintf("%s: %s", name.Name(), astmodel.DebugDescription(t, name.PackageReference)))
	dr.writeType(sub, definition.Type(), name.PackageReference)
}

// writeType writes the type to the debug report.
// rpt is the debug report to write to.
// t is the type to write.
// currentPackage is the package that the type is defined in (used to simplify type descriptions).
// Only complex types where astmodel.DebugDescription is insufficient are written.
func (dr *debugReporter) writeType(
	rpt *reporting.StructureReport,
	t astmodel.Type,
	currentPackage astmodel.PackageReference,
) {
	if rsrc, ok := astmodel.AsResourceType(t); ok {
		dr.writeResource(rpt, rsrc, currentPackage)
	} else if obj, ok := astmodel.AsObjectType(t); ok {
		dr.writeObject(rpt, obj, currentPackage)
	} else if obj, ok := astmodel.AsEnumType(t); ok {
		dr.writeEnum(rpt, obj, currentPackage)
	}
}

// writeResource writes the resource to the debug report.
// rpt is the debug report to write to.
// resource is the resource to write.
// currentPackage is the package that the resource is defined in (used to simplify type descriptions).
func (dr *debugReporter) writeResource(
	rpt *reporting.StructureReport,
	resource *astmodel.ResourceType,
	currentPackage astmodel.PackageReference,
) {
	for _, prop := range resource.Properties().AsSlice() {
		dr.writeProperty(rpt, prop, currentPackage)
	}

	for _, fn := range resource.Functions() {
		dr.writeFunction(rpt, fn)
	}
}

// writeObject writes the object to the debug report.
func (dr *debugReporter) writeObject(
	rpt *reporting.StructureReport,
	obj *astmodel.ObjectType,
	currentPackage astmodel.PackageReference,
) {
	for _, prop := range obj.Properties().AsSlice() {
		dr.writeProperty(rpt, prop, currentPackage)
	}

	for _, fn := range obj.Functions() {
		dr.writeFunction(rpt, fn)
	}
}

func (dr *debugReporter) writeProperty(
	rpt *reporting.StructureReport,
	prop *astmodel.PropertyDefinition,
	currentPackage astmodel.PackageReference,
) {
	rpt.Add(fmt.Sprintf(
		"%s: %s",
		prop.PropertyName(),
		astmodel.DebugDescription(prop.PropertyType(), currentPackage)))
}

func (dr *debugReporter) writeFunction(rpt *reporting.StructureReport, fn astmodel.Function) {
	rpt.Add(fmt.Sprintf("%s()", fn.Name()))
}

func (dr *debugReporter) writeEnum(rpt *reporting.StructureReport, enum *astmodel.EnumType, currentPackage astmodel.PackageReference) {
	dr.writeType(rpt, enum.BaseType(), currentPackage)
	for _, v := range enum.Options() {
		rpt.Add(fmt.Sprintf("%s: %s", v.Identifier, v.Value))
	}
}

var dashMatcher = regexp.MustCompile("-+")

// createFileName creates the filename for the debug report from the name of the stage, filtering out any characters
// that are unsafe in filenames
func (dr *debugReporter) createFileName(stage int, description string) string {
	// filter symbols and other unsafe characters from the stage name to generate a safe filename for the debug log
	stageName := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			return r
		}

		return '-'
	}, description)

	// Replace any sequence of dashes with a single one using a regular expression
	stageName = dashMatcher.ReplaceAllString(stageName, "-")

	// Create a filename using the description and the stage number.
	filename := strconv.Itoa(stage+1) + "-" + stageName + ".txt"

	return path.Join(dr.outputFolder, filename)
}

func (dr *debugReporter) findPackages(def astmodel.TypeDefinitionSet) []astmodel.PackageReference {
	packages := astmodel.NewPackageReferenceSet()
	for _, def := range def {
		packages.AddReference(def.Name().PackageReference)
	}

	result := packages.AsSortedSlice(func(left astmodel.PackageReference, right astmodel.PackageReference) bool {
		return astmodel.ComparePathAndVersion(left.PackagePath(), right.PackagePath())
	})

	return result
}

func (dr *debugReporter) inPackage(
	ref astmodel.PackageReference,
	def astmodel.TypeDefinitionSet,
) astmodel.TypeDefinitionSet {
	result := make(astmodel.TypeDefinitionSet)
	for _, def := range def {
		if def.Name().PackageReference == ref {
			result.Add(def)
		}
	}

	return result
}
