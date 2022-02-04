/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
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

	// Original Approach

	newDefinitions := make(astmodel.Types)
	filterer := configuration.BuildExportFilterer(state.types)

	oldRenames := make(map[astmodel.TypeName]astmodel.TypeName)
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
				oldRenames[defName] = *newName
			}
		default:
			panic(fmt.Sprintf("unhandled shouldExport case %q", shouldExport))
		}
	}

	// New Approach

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

	// Verify consistency
	missing := astmodel.FindResourceTypes(newDefinitions.Except(defs))
	if len(missing) > 0 {
		// Some expected types missing
		for name := range missing {
			klog.Errorf("Type %s is missing", name)
		}

		return nil, errors.Errorf("Found %d missing types", len(missing))
	}

	extras := astmodel.FindResourceTypes(defs.Except(newDefinitions))
	if len(extras) > 0 {
		// Some unexpected types present
		for name := range extras {
			klog.Errorf("Type %s is unexpected", name)
		}

		return nil, errors.Errorf("Found %d unexpected types", len(extras))
	}

	count := 0
	for o, n := range oldRenames {
		if _, ok := renames[o]; !ok {
			klog.Errorf("Missing rename of %s to %s", o, n)
			count++
		}
	}

	if count > 0 {
		return nil, errors.Errorf("Found %d unsupported renames", count)
	}

	if err := configuration.ObjectModelConfiguration.VerifyExportConsumed(); err != nil {
		return nil, err
	}

	if err := configuration.ObjectModelConfiguration.VerifyExportAsConsumed(); err != nil {
		return nil, err
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
