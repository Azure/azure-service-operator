/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// CreateConversionGraphStageID is the unique identifier for this stage
const CreateConversionGraphStageID = "createConversionGraph"

// CreateConversionGraph walks the set of available types and creates a graph of conversions that will be used to
// convert resources to/from the designated storage (or hub) version
func CreateConversionGraph(
	configuration *config.Configuration,
	generatorPrefix string,
) *Stage {
	stage := NewStage(
		CreateConversionGraphStageID,
		"Create the graph of conversions between versions of each resource group",
		func(ctx context.Context, state *State) (*State, error) {
			// Collect all distinct references
			allNames := astmodel.NewInternalTypeNameSet()
			for _, def := range state.Definitions() {
				if def.Name().IsARMType() {
					// ARM types don't participate in the conversion graph
					continue
				}

				if astmodel.IsWebhookPackageReference(def.Name().PackageReference()) {
					// Webhook types also don't participate in the conversion graph
					continue
				}

				allNames.Add(def.Name())
			}

			builder := storage.NewConversionGraphBuilder(
				configuration.ObjectModelConfiguration, generatorPrefix)
			builder.AddAll(allNames)
			graph, err := builder.Build()
			if err != nil {
				// Shouldn't have any non-local references, if we do, abort
				return nil, eris.Wrapf(err, "creating conversion graph")
			}

			return StateWithData(state, ConversionGraphInfo, graph), nil
		})

	stage.AddDiagnostic(exportConversionGraph)

	stage.RequiresPrerequisiteStages(CreateStorageTypesStageID)
	return stage
}

func exportConversionGraph(settings *DebugSettings, index int, state *State) error {
	graph, err := GetStateData[*storage.ConversionGraph](state, ConversionGraphInfo)
	if err != nil {
		return eris.Wrapf(err, "conversion graph not found")
	}

	// Create our output folder
	outputFolder := settings.CreateFileName(fmt.Sprintf("conversion-graph-%d", index))
	err = os.Mkdir(outputFolder, 0o700)
	if err != nil {
		return eris.Wrapf(err, "creating output folder for conversion graph diagnostic")
	}

	done := set.Make[string]()
	for name, def := range state.Definitions() {
		if name.IsARMType() {
			// ARM types don't participate in the conversion graph
			continue
		}

		_, isResource := astmodel.AsResourceType(def.Type())
		_, isObject := astmodel.AsObjectType(def.Type())
		if !isResource && !isObject {
			// Not a resource or object, so not a type we're interested in
			continue
		}

		if !settings.MatchesGroup(name.InternalPackageReference()) {
			// Not a group/version we're interested in
			continue
		}

		grp := name.InternalPackageReference().Group()
		name := name.Name()
		key := fmt.Sprintf("%s-%s.gv", grp, name)
		if done.Contains(key) {
			// We've already exported this graph
			continue
		}

		done.Add(key)

		filename := filepath.Join(outputFolder, key)
		err := graph.SaveTo(grp, name, filename)
		if err != nil {
			return eris.Wrapf(err, "writing conversion graph for %s", name)
		}
	}

	return nil
}
