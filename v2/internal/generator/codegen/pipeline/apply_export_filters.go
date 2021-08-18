/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
	"github.com/Azure/azure-service-operator/v2/internal/generator/config"
)

const ApplyExportFiltersStageID = "filterTypes"

// ApplyExportFilters creates a Stage to reduce our set of types for export
func ApplyExportFilters(configuration *config.Configuration) Stage {
	return MakeStage(
		ApplyExportFiltersStageID,
		"Apply export filters to reduce the number of generated types",
		func(ctx context.Context, state *State) (*State, error) {
			return filterTypes(configuration, state)
		})
}

// filterTypes applies the configuration include/exclude filters to the generated definitions
func filterTypes(
	configuration *config.Configuration,
	state *State) (*State, error) {

	newDefinitions := make(astmodel.Types)

	filterer := configuration.BuildExportFilterer(state.types)

	for _, def := range state.types {
		defName := def.Name()
		shouldExport, reason := filterer(defName)

		switch shouldExport {
		case config.Skip:
			klog.V(3).Infof("Skipping %s because %s", defName, reason)

		case config.Export:
			if reason == "" {
				klog.V(3).Infof("Exporting %s", defName)
			} else {
				klog.V(2).Infof("Exporting %s because %s", defName, reason)
			}

			newDefinitions[def.Name()] = def
		default:
			panic(fmt.Sprintf("unhandled shouldExport case %q", shouldExport))
		}
	}

	// Ensure that the export filters had no issues
	err := configuration.GetExportFiltersError()
	if err != nil {
		return nil, err
	}

	return state.WithTypes(newDefinitions), nil
}
