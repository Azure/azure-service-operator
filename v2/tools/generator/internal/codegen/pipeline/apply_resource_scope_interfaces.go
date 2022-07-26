/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/interfaces"
)

const ApplyResourceScopeInterfacesStageID = "applyResourceScopeInterfaces"

// ApplyResourceScopeInterfaces ensures that resources implement the correct resource scope interfaces
func ApplyResourceScopeInterfaces(idFactory astmodel.IdentifierFactory) *Stage {
	return NewStage(
		ApplyResourceScopeInterfacesStageID,
		"Ensures that resources implement the correct resource scope interfaces",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs := make(astmodel.TypeDefinitionSet)

			for typeName, typeDef := range astmodel.FindResourceDefinitions(state.Definitions()) {
				resource := typeDef.Type().(*astmodel.ResourceType)
				if resource.Kind() == astmodel.ResourceKindTenant {
					updatedDef, err := interfaces.AddTenantResourceInterface(typeDef, idFactory)
					if err != nil {
						return nil, errors.Wrapf(err, "failed to add tenant interface to %s", typeName)
					}
					updatedDefs.Add(updatedDef)
				}
			}

			return state.WithDefinitions(state.Definitions().OverlayWith(updatedDefs)), nil
		})
}
