/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const RemoveResourceScopeStageID = "removeResourceScope"

// RemoveResourceScope removes the "Scope" property from resource Specs, as it only applies in the URL
// of requests, not in the body.
func RemoveResourceScope() Stage {
	return MakeStage(
		RemoveResourceScopeStageID,
		"Remove scope from all resources",
		func(ctx context.Context, state *State) (*State, error) {
			newDefs := make(astmodel.TypeDefinitionSet)
			scopePropertyRemovalVisitor := makeScopePropertyRemovalVisitor()

			resources := astmodel.FindResourceDefinitions(state.Definitions())
			for _, resource := range resources {
				resolved, err := state.Definitions().ResolveResourceSpecAndStatus(resource)
				if err != nil {
					return nil, errors.Wrapf(err, "unable to find resource %s spec and status", resource.Name())
				}

				updatedDef, err := scopePropertyRemovalVisitor.VisitDefinition(resolved.SpecDef, nil)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to remove scope property from %s", updatedDef.Name())
				}
				newDefs.Add(updatedDef)
			}

			result := state.Definitions().OverlayWith(newDefs)

			return state.WithDefinitions(result), nil
		})
}

func makeScopePropertyRemovalVisitor() astmodel.TypeVisitor {
	return astmodel.TypeVisitorBuilder{
		VisitObjectType: removeScopeProperty,
	}.Build()
}

func removeScopeProperty(this *astmodel.TypeVisitor, ot *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	ot = ot.WithoutProperty(astmodel.ScopeProperty)

	return astmodel.IdentityVisitOfObjectType(this, ot, ctx)
}
