/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

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

			for typeName, typeDef := range astmodel.FindResourceDefinitions(state.Definitions()) {
				resource := typeDef.Type().(*astmodel.ResourceType)
				newResource, err := interfaces.AddKubernetesResourceInterfaceImpls(
					typeDef,
					idFactory,
					state.Definitions(),
					log)
				if err != nil {
					return nil, errors.Wrapf(err, "couldn't implement Kubernetes resource interface for %q", typeName)
				}

				// this is really very ugly; a better way?
				if _, ok := newResource.SpecType().(astmodel.TypeName); !ok {
					// the resource Spec was replaced with a new definition; update it
					// by replacing the named definition:
					specName := resource.SpecType().(astmodel.InternalTypeName)
					updatedDefs.Add(astmodel.MakeTypeDefinition(specName, newResource.SpecType()))
					newResource = newResource.WithSpec(specName)
				}

				newDef := typeDef.WithType(newResource)
				updatedDefs.Add(newDef)
			}

			return state.WithDefinitions(state.Definitions().OverlayWith(updatedDefs)), nil
		})
}
