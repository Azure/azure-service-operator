/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

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
	var errs []error
	for _, def := range astmodel.FindResourceDefinitions(state.Definitions()) {
		defName := def.Name()

		export, err := shouldExport(defName, configuration)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if !export {
			log.V(1).Info("Skipping resource", "resource", defName)
			continue
		}

		log.V(1).Info("Exporting resource", "resource", defName)
		resourcesToExport.Add(def)
	}

	if err := kerrors.NewAggregate(errs); err != nil {
		return nil, err
	}

	typesToExport, err := astmodel.FindConnectedDefinitions(state.Definitions(), resourcesToExport)
	if err != nil {
		return nil, errors.Wrap(err, "finding types connected to resources marked for export")
	}

	// Find and apply renames
	renames := make(map[astmodel.TypeName]astmodel.TypeName)
	for n := range typesToExport {
		if as, asErr := configuration.ObjectModelConfiguration.LookupExportAs(n); asErr == nil {
			renames[n] = n.WithName(as)
		}
	}

	if err = configuration.ObjectModelConfiguration.VerifyExportConsumed(); err != nil {
		return nil, err
	}

	if err = configuration.ObjectModelConfiguration.VerifyExportAsConsumed(); err != nil {
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
func shouldExport(defName astmodel.TypeName, configuration *config.Configuration) (bool, error) {
	export, err := configuration.ObjectModelConfiguration.LookupExport(defName)
	if err == nil {
		// $export is configured, return that value
		return export, nil
	}

	if !config.IsNotConfiguredError(err) {
		// Problem isn't lack of configuration, it's something else
		return false, errors.Wrapf(err, "looking up export config for %s", defName)
	}

	_, err = configuration.ObjectModelConfiguration.LookupExportAs(defName)
	if err == nil {
		// $exportAs is configured, we DO want to export
		return true, nil
	}

	if !config.IsNotConfiguredError(err) {
		// Problem isn't lack of configuration, it's something else
		return false, errors.Wrapf(err, "looking up exportAs config for %s", defName)
	}

	// Default is to not export
	return false, nil
}
