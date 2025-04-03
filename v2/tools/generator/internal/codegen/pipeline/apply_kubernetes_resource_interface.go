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
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/interfaces"
)

const ApplyKubernetesResourceInterfaceStageID = "applyKubernetesResourceInterface"

// ApplyKubernetesResourceInterface ensures that every Resource implements the KubernetesResource interface
func ApplyKubernetesResourceInterface(
	idFactory astmodel.IdentifierFactory,
	log logr.Logger,
) *Stage {
	return NewStage(
		ApplyKubernetesResourceInterfaceStageID,
		"Add the KubernetesResource interface to every resource",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs := make(astmodel.TypeDefinitionSet)

			for typeName, typeDef := range state.Definitions().AllResources() {
				newDefs, err := interfaces.AddKubernetesResourceInterfaceImpls(
					typeDef,
					idFactory,
					state.Definitions(),
					log)
				if err != nil {
					return nil, eris.Wrapf(err, "couldn't implement Kubernetes resource interface for %q", typeName)
				}

				updatedDefs.AddTypes(newDefs)
			}

			return state.WithOverlaidDefinitions(updatedDefs), nil
		})
}
