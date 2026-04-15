/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
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
	for _, def := range state.Definitions().AllResources() {
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
		return nil, eris.Wrap(err, "finding types connected to resources marked for export")
	}

	// Find and apply renames.
	// If we rename a resource, we use that name for its spec and status as well.
	renames := make(astmodel.TypeAssociation)
	addRename := func(name astmodel.InternalTypeName, newName string) error {
		if name.Name() == newName {
			// Nothing to do
			return nil
		}

		n := name.WithName(newName)
		if existing, ok := state.Definitions()[n]; ok {
			// Can't rename, as we'd create a name collision
			return eris.Errorf(
				"can't rename %s to %s, name is already used by %s",
				name,
				newName,
				existing.Name())
		}

		renames[name] = n
		// Add an alias to the configuration so that we can use the new name to access the rest of the config
		if configuration.ObjectModelConfiguration.IsTypeConfigured(name) {
			configuration.ObjectModelConfiguration.AddTypeAlias(name, newName)
		}

		return nil
	}

	for n, def := range typesToExport {
		newName := ""
		if as, ok := configuration.ObjectModelConfiguration.ExportAs.Lookup(n); ok {
			newName = as
		} else if to, ok := configuration.ObjectModelConfiguration.RenameTo.Lookup(n); ok {
			newName = to
		}

		if newName != "" {
			if err := addRename(n, newName); err != nil {
				return nil, err
			}

			// If this is a resource, we also need to rename the spec and status.
			// These renames are optional — if there's a collision we log a warning
			// but don't fail, as the resource rename is the important one.
			if rt, ok := astmodel.AsResourceType(def.Type()); ok {
				if spec, ok := astmodel.AsInternalTypeName(rt.SpecType()); ok {
					if err := addRename(spec, newName+astmodel.SpecSuffix); err != nil {
						log.V(1).Info("Warning: could not rename spec type", "type", spec, "error", err)
					}
				}
				if status, ok := astmodel.AsInternalTypeName(rt.StatusType()); ok {
					if err := addRename(status, newName+astmodel.StatusSuffix); err != nil {
						log.V(1).Info("Warning: could not rename status type", "type", status, "error", err)
					}
				}
			}
		}
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
	if _, ok := configuration.ObjectModelConfiguration.ExportAs.Lookup(defName); ok {
		// $exportAs is configured, we DO want to export
		return true
	}

	// Default is to not export
	return false
}
