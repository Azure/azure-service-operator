/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

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

	defs := make(astmodel.Types)
	types := state.Types()
	renames := make(map[astmodel.TypeName]astmodel.TypeName)
	for _, def := range astmodel.FindResourceTypes(types) {
		defName := def.Name()

		export, err := configuration.ObjectModelConfiguration.LookupExport(defName)
		if err != nil {
			if config.IsNotConfiguredError(err) {
				export = false
			}
		}

		if _, err = configuration.ObjectModelConfiguration.LookupExportAs(defName); err == nil {
			// $exportAs is configured, we must export
			export = true
		}

		if !export {
			klog.V(3).Infof("Skipping resource %s", defName)
			continue
		}

		klog.V(3).Infof("Exporting resource %s and related types", defName)
		typesToExport := findReferencedTypes(defName, types)
		for n := range typesToExport {
			defs[n] = types[n]

			if as, err := configuration.ObjectModelConfiguration.LookupExportAs(n); err == nil {
				renames[n] = n.WithName(as)
			}
		}
	}

	if err := configuration.ObjectModelConfiguration.VerifyExportConsumed(); err != nil {
		return nil, err
	}

	if err := configuration.ObjectModelConfiguration.VerifyExportAsConsumed(); err != nil {
		return nil, err
	}

	// Now apply all the renames
	renamingVisitor := astmodel.NewRenamingVisitor(renames)
	result, err := renamingVisitor.RenameAll(defs)
	if err != nil {
		return nil, err
	}

	return state.WithTypes(result), nil
}

func findReferencedTypes(root astmodel.TypeName, types astmodel.Types) astmodel.TypeNameSet {
	result := astmodel.NewTypeNameSet(root)
	collectReferencedTypes(root, types, result)
	return result
}

func collectReferencedTypes(root astmodel.TypeName, types astmodel.Types, referenced astmodel.TypeNameSet) {
	referenced.Add(root)
	for ref := range types[root].Type().References() {
		if !referenced.Contains(ref) {
			collectReferencedTypes(ref, types, referenced)
		}
	}
}
