/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// ReportResourceStructureStageID is the unique identifier of this stage
const ReportResourceStructureStageID = "reportResourceStructure"

// ReportResourceVersions creates a pipeline stage that generates reports showing the structure of each generated resource
// These reports allow a PR diff to be used to see what has changed in the structure of each resource between versions
func ReportResourceStructure(configuration *config.Configuration) *Stage {
	return NewStage(
		ReportResourceStructureStageID,
		"Generate reports showing the structure of each generated resource",
		func(ctx context.Context, state *State) (*State, error) {
			//report := NewResourceStructureReport(state.Definitions(), configuration)
			//
			//err := report.WriteTo(configuration.SupportedResourcesReport.FullOutputPath())
			//if err != nil {
			//	return nil, err
			//}
			//
			//err = configuration.ObjectModelConfiguration.VerifySupportedFromConsumed()
			//return state, err
			return nil, nil
		})
}
