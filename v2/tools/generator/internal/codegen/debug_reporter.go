/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"strconv"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/pipeline"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
)

// debugReporter is a helper for generating debug logs during the code generation process.
type debugReporter struct {
	settings *pipeline.DebugSettings
}

// newDebugReporter creates a new debugReporter.
// settings specifies the debug settings to use as configuration.
func newDebugReporter(settings *pipeline.DebugSettings) *debugReporter {
	return &debugReporter{
		settings: settings,
	}
}

func (dr *debugReporter) ReportStage(stage int, description string, state *pipeline.State) error {
	included := state.Definitions().Where(
		func(def astmodel.TypeDefinition) bool {
			// Allow matching just the group (e.g. network)
			return dr.settings.MatchesGroup(def.Name().InternalPackageReference())
		})

	tcr := reporting.NewTypeCatalogReport(included, reporting.IncludeFunctions)
	name := strconv.Itoa(stage) + "-" + description + ".txt"
	filename := dr.settings.CreateFileName(name)

	err := tcr.SaveTo(filename)
	return eris.Wrapf(err, "failed to save type catalog to %s", filename)
}
