/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/util/match"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/pipeline"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
)

// debugReporter is a helper for generating debug logs during the code generation process.
type debugReporter struct {
	outputFolder  string
	groupSelector match.StringMatcher
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
	included := state.Definitions().Where(
		func(def astmodel.TypeDefinition) bool {
			grp := def.Name().PackageReference().Group()
			return dr.groupSelector.Matches(grp).Matched
		})

	tcr := reporting.NewTypeCatalogReport(included)
	tcr.IncludeFunctions()
	filename := dr.createFileName(stage, description)
	err := tcr.SaveTo(filename)
	return errors.Wrapf(err, "failed to save type catalog to %s", filename)
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

	return filepath.Join(dr.outputFolder, filename)
}
