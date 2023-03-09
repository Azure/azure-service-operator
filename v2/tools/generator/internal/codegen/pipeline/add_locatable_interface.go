/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

const AddLocatableInterfaceStageID = "addLocatableInterface"

func AddLocatableInterface(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		AddLocatableInterfaceStageID,
		"Add the Locatable interface for Location based resources such as ResourceGroup",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs := make(astmodel.TypeDefinitionSet)

			for _, def := range astmodel.FindResourceDefinitions(state.Definitions()) {
				rt := def.Type().(*astmodel.ResourceType)

				if rt.Scope() == astmodel.ResourceScopeLocation {
					rt = rt.WithInterface(functions.NewLocatableResource(idFactory, rt))
					updatedDefs.Add(def.WithType(rt))
				}
			}

			return state.WithDefinitions(state.Definitions().OverlayWith(updatedDefs)), nil
		},
	)

	return stage
}
