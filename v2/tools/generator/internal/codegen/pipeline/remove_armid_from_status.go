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

func FixIDFields() *Stage {
	stage := NewStage(
		"fixIdFields",
		"Remove ARM ID annotations from status, and Id from Spec types",
		func(ctx context.Context, state *State) (*State, error) {

			updatedStatusDefs, err := replaceStatusARMIDWithString(state.Definitions())
			if err != nil {
				return nil, err
			}

			updatedSpecDefs, err := removeSpecIDField(state.Definitions())
			if err != nil {
				return nil, err
			}

			updatedDefs := astmodel.TypesDisjointUnion(updatedStatusDefs, updatedSpecDefs)

			state = state.WithDefinitions(state.Definitions().OverlayWith(updatedDefs))
			return state, nil
		})

	return stage
}

func removeSpecIDField(defs astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	removeIDVisitor := astmodel.TypeVisitorBuilder[any]{
		VisitObjectType: func(this *astmodel.TypeVisitor[any], it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
			it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
				prim, isPrimitive := astmodel.AsPrimitiveType(prop.PropertyType())
				if prop.HasName("Id") && isPrimitive && prim == astmodel.ARMIDType {
					it = it.WithoutSpecificProperties(prop.PropertyName())
				}
			})

			return astmodel.IdentityVisitOfObjectType(this, it, ctx)
		},
	}.Build()

	updatedDefs, err := removeIDVisitor.VisitDefinitions(astmodel.FindSpecDefinitions(defs), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to remove spec.Id")
	}

	return updatedDefs, nil
}

func replaceStatusARMIDWithString(defs astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	replaceARMIDWithStringVisitor := astmodel.TypeVisitorBuilder[any]{
		VisitPrimitive: func(_ *astmodel.TypeVisitor[any], it *astmodel.PrimitiveType, _ interface{}) (astmodel.Type, error) {
			if it == astmodel.ARMIDType {
				return astmodel.StringType, nil
			}

			return it, nil
		},
	}.Build()

	updatedDefs, err := replaceARMIDWithStringVisitor.VisitDefinitions(astmodel.FindStatusDefinitions(defs), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to replace ARM ID with String on status types")
	}

	return updatedDefs, nil
}
