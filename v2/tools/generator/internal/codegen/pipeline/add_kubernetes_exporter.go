/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

const AddKubernetesExporterStageID = "addKubernetesExporter"

func AddKubernetesExporter(idFactory astmodel.IdentifierFactory) *Stage {
	return NewStage(
		AddKubernetesExporterStageID,
		"Adds the KubernetesExporter interface to resources that need it",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			updatedDefs := make(astmodel.TypeDefinitionSet)

			mappings := state.GeneratedConfigMaps()

			for _, def := range astmodel.FindResourceDefinitions(defs) {
				resourceType, ok := astmodel.AsResourceType(def.Type())
				if !ok {
					return nil, errors.Errorf("%s definition type wasn't a resource", def.Name())
				}

				configMapMappings, ok := mappings.Get(def.Name())
				if !ok {
					continue
				}

				builder := functions.NewKubernetesExporterBuilder(def.Name(), resourceType, idFactory, configMapMappings)

				resourceType = resourceType.WithInterface(builder.ToInterfaceImplementation())
				updatedDef := def.WithType(resourceType)
				updatedDefs.Add(updatedDef)
			}

			result := defs.OverlayWith(updatedDefs)

			return state.WithDefinitions(result), nil
		})
}
