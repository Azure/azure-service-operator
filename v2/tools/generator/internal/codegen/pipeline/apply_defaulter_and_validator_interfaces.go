/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/interfaces"
)

const ApplyDefaulterAndValidatorInterfaceStageID = "applyDefaulterAndValidatorInterfaces"

// ApplyDefaulterAndValidatorInterfaces add the admission.Defaulter and admission.Validator interfaces to each resource that requires them
func ApplyDefaulterAndValidatorInterfaces(idFactory astmodel.IdentifierFactory) Stage {
	stage := MakeStage(
		ApplyDefaulterAndValidatorInterfaceStageID,
		"Add the admission.Defaulter and admission.Validator interfaces to each resource that requires them",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			updatedDefs := make(astmodel.TypeDefinitionSet)

			for _, typeDef := range astmodel.FindResourceDefinitions(defs) {
				resource, err := interfaces.AddDefaulterInterface(typeDef, idFactory, defs)
				if err != nil {
					return nil, err
				}

				resource, err = interfaces.AddValidatorInterface(resource, idFactory, defs)
				if err != nil {
					return nil, err
				}

				updatedDefs.Add(resource)
			}

			return state.WithDefinitions(defs.OverlayWith(updatedDefs)), nil
		})

	return stage.RequiresPrerequisiteStages(ApplyKubernetesResourceInterfaceStageID)
}
