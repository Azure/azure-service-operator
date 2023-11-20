/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
)

const ApplyExportFiltersStageID = "filterTypes"

// ApplyExportFilters creates a Stage to reduce our set of types for export
func ApplyExportFilters(
	configuration *config.Configuration,
	log logr.Logger,
) *Stage {
	stage := NewStage(
		ApplyExportFiltersStageID,
		"Apply export filters to reduce the number of generated types",
		func(ctx context.Context, state *State) (*State, error) {
			return filterTypes(configuration, state, log)
		})

	stage.RequiresPostrequisiteStages(VerifyNoErroredTypesStageID)
	return stage
}

// filterTypes applies the configuration include/exclude filters to the generated definitions
func filterTypes(
	configuration *config.Configuration,
	state *State,
	log logr.Logger,
) (*State, error) {
	resourcesToExport := make(astmodel.TypeDefinitionSet)
	for _, def := range astmodel.FindResourceDefinitions(state.Definitions()) {
		defName := def.Name()

		export := shouldExport(defName, configuration)
		if !export {
			log.V(1).Info("Skipping resource", "resource", defName)
			continue
		}

		log.V(1).Info("Exporting resource", "resource", defName)
		resourcesToExport.Add(def)
	}

	typesToExport, err := astmodel.FindConnectedDefinitions(state.Definitions(), resourcesToExport)
	if err != nil {
		return nil, errors.Wrap(err, "finding types connected to resources marked for export")
	}

	// Find and apply renames
	renames := make(astmodel.TypeAssociation)
	for n := range typesToExport {
		newName := ""
		if as, ok := configuration.ObjectModelConfiguration.ExportAs.Lookup(n); ok {
			newName = as
		} else if to, ok := configuration.ObjectModelConfiguration.RenameTo.Lookup(n); ok {
			newName = to
		}

		if newName != "" {
			// Add an alias to the configuration so that we can use the new name to access the rest of the config
			configuration.ObjectModelConfiguration.AddTypeAlias(n, newName)
			renames[n] = n.WithName(newName)
		}
	}

	if err = configuration.ObjectModelConfiguration.Export.VerifyConsumed(); err != nil {
		return nil, err
	}

	if err = configuration.ObjectModelConfiguration.ExportAs.VerifyConsumed(); err != nil {
		return nil, err
	}

	if err = configuration.ObjectModelConfiguration.RenameTo.VerifyConsumed(); err != nil {
		return nil, err
	}

	// Now apply all the renames
	renamingVisitor := astmodel.NewRenamingVisitor(renames)
	result, err := renamingVisitor.RenameAll(typesToExport)
	if err != nil {
		return nil, err
	}

	return state.WithDefinitions(result), nil
}

// shouldExport works out whether the specified Resource should be exported or not
func shouldExport(defName astmodel.InternalTypeName, configuration *config.Configuration) bool {
	if export, ok := configuration.ObjectModelConfiguration.Export.Lookup(defName); ok {
		// $export is configured, return that value
		return export
	}

	if _, ok := configuration.ObjectModelConfiguration.ExportAs.Lookup(defName); ok {
		// $exportAs is configured, we DO want to export
		return true
	}

	// Default is to not export
	return false
}
