/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
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
	renames := make(map[astmodel.TypeName]astmodel.TypeName)

	for _, def := range state.types {
		defName := def.Name()
		shouldExport, newName, reason := filterer(defName)

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
			if newName != nil {
				renames[defName] = *newName
			}
		default:
			panic(fmt.Sprintf("unhandled shouldExport case %q", shouldExport))
		}
	}

	// Now apply all the renames
	renamingVisitor := astmodel.NewRenamingVisitor(renames)
	result, err := renamingVisitor.RenameAll(newDefinitions)
	if err != nil {
		return nil, err
	}

	// Ensure that the export filters had no issues
	err = configuration.GetExportFiltersError()
	if err != nil {
		return nil, err
	}

	return state.WithTypes(result), nil
}
