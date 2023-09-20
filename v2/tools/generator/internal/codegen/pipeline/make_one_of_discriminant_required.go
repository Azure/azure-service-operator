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

// MakeOneOfDiscriminantRequiredStageId is the unique identifier for this pipeline stage
const MakeOneOfDiscriminantRequiredStageId = "makeOneOfDiscriminantRequired"

// MakeOneOfDiscriminantRequired walks the type graph and builds new types for communicating
// with ARM
func MakeOneOfDiscriminantRequired() *Stage {
	return NewStage(
		MakeOneOfDiscriminantRequiredStageId,
		"Fix one of types to a discriminator which is not omitempty/optional",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs := make(astmodel.TypeDefinitionSet)
			for _, def := range state.Definitions() {
				isOneOf := astmodel.OneOfFlag.IsOn(def.Type())
				isARM := def.Name().IsARMType()
				if !isOneOf || !isARM {
					continue
				}

				updated, err := makeOneOfDiscriminantTypeRequired(def, state.Definitions())
				if err != nil {
					return nil, err
				}
				updatedDefs.AddTypes(updated)
			}

			return state.WithDefinitions(state.Definitions().OverlayWith(updatedDefs)), nil
		})
}

type propertyModifier struct {
	visitor astmodel.TypeVisitor[any]
	json    string
}

func newPropertyModifier(json string) *propertyModifier {
	modifier := &propertyModifier{
		json: json,
	}

	modifier.visitor = astmodel.TypeVisitorBuilder[any]{
		VisitObjectType: modifier.makeDiscriminatorPropertiesRequired,
	}.Build()

	return modifier
}

func (r *propertyModifier) makeDiscriminatorPropertiesRequired(
	this *astmodel.TypeVisitor[any],
	ot *astmodel.ObjectType,
	ctx any,
) (astmodel.Type, error) {
	ot.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
		if json, ok := prop.JSONName(); ok && r.json == json {
			ot = ot.WithProperty(prop.MakeTypeRequired())
		}
	})

	return astmodel.IdentityVisitOfObjectType(this, ot, ctx)
}

func makeOneOfDiscriminantTypeRequired(
	oneOf astmodel.TypeDefinition,
	defs astmodel.TypeDefinitionSet,
) (astmodel.TypeDefinitionSet, error) {
	objectType, ok := astmodel.AsObjectType(oneOf.Type())
	if !ok {
		return nil, errors.Errorf(
			"OneOf %s was not of type Object, instead was: %s",
			oneOf.Name(),
			astmodel.DebugDescription(oneOf.Type()))
	}

	result := make(astmodel.TypeDefinitionSet)
	discriminantJson, values, err := astmodel.DetermineDiscriminantAndValues(objectType, defs)
	if err != nil {
		return nil, err
	}

	astmodel.NewPropertyInjector()
	remover := newPropertyModifier(discriminantJson)

	for _, value := range values {
		def, err := defs.GetDefinition(value.TypeName)
		if err != nil {
			return nil, err
		}
		updatedDef, err := remover.visitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "error updating definition %s", def.Name())
		}

		result.Add(updatedDef)

	}

	return result, nil
}
