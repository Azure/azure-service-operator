/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
)

const RemoveResourceScopeStageID = "removeResourceScope"

// RemoveResourceScope removes the "Scope" property from resource Specs except for on Extension
// resources, which support scope.
func RemoveResourceScope() Stage {
	return MakeStage(
		RemoveResourceScopeStageID,
		"Removes scope from all resources except extension resources",
		func(ctx context.Context, state *State) (*State, error) {
			newDefs := make(astmodel.Types)
			scopePropertyRemovalVisitor := makeScopePropertyRemovalVisitor()

			for _, def := range astmodel.FindResourceTypes(state.Types()) {
				resource, ok := astmodel.AsResourceType(def.Type())
				if !ok {
					// panic here because this must be a bug
					panic("FindResourceTypes returned a type that wasn't a resource")
				}

				// If the resource is an extension resource, we don't do anything to it
				if resource.Kind() == astmodel.ResourceKindExtension {
					continue
				}

				specDef, err := state.Types().ResolveResourceSpecDefinition(resource)
				if err != nil {
					return nil, err
				}

				updatedDef, err := scopePropertyRemovalVisitor.VisitDefinition(specDef, nil)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to remove scope property from %s", updatedDef.Name())
				}
				newDefs.Add(updatedDef)
			}

			result := state.Types().OverlayWith(newDefs)

			return state.WithTypes(result), nil
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
