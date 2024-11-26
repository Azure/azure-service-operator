/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

const AddSerializationTypeTagStageID = "addSerializationTypeTag"

// AddSerializationTypeTag adds a property tag to properties with special serialization instructions to initialize empty collections
// when serializing the payload to Azure.
// This uses a property tag for a few reasons:
//  1. Some types are flattened and other approaches are easily
//     lost when flattening occurs. Putting the tag onto the property
//     preserves it even through flattening.
//  2. In many ways this behavior is like an augmented `json:omitempty`,
//     so it (IMO) makes sense to have a tag just like for JSON. It makes
//     it clearer when looking at the Go object that the serialization behavior
//     of these fields is special.
func AddSerializationTypeTag(configuration *config.Configuration) *Stage {
	return NewStage(
		AddSerializationTypeTagStageID,
		"Adds a property tag to properties with special serialization instructions to initialize empty collections when serializing the payload to Azure",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs := make(astmodel.TypeDefinitionSet)

			visitor := makePayloadTypeVisitor()
			for _, def := range state.Definitions() {
				t, err := visitor.Visit(def.Type(), &serializationVisitorContext{
					name:   def.Name(),
					config: configuration,
				})
				if err != nil {
					return nil, eris.Wrapf(err, "visiting %q", def.Name())
				}

				updatedDefs.Add(def.WithType(t))
			}

			return state.WithDefinitions(updatedDefs), nil
		})
}

type serializationVisitorContext struct {
	name   astmodel.InternalTypeName
	config *config.Configuration
}

func applySerializationTypeTag(it *astmodel.TypeVisitor[*serializationVisitorContext], ot *astmodel.ObjectType, ctx *serializationVisitorContext) (astmodel.Type, error) {
	var updatedProps []*astmodel.PropertyDefinition

	ot.Properties().ForEach(
		func(prop *astmodel.PropertyDefinition) {
			payloadType, ok := ctx.config.ObjectModelConfiguration.PayloadType.Lookup(ctx.name, prop.PropertyName())
			if !ok {
				return // continue
			}

			// Don't bother annotating the default behavior
			if payloadType != config.ExplicitEmptyCollections {
				return // continue
			}

			prop = prop.WithTag(astmodel.SerializationType, astmodel.SerializationTypeExplicitEmptyCollection)
			updatedProps = append(updatedProps, prop)
		})

	ot = ot.WithProperties(updatedProps...).WithProperties(updatedProps...)

	return astmodel.IdentityVisitOfObjectType(it, ot, ctx)
}

func makePayloadTypeVisitor() astmodel.TypeVisitor[*serializationVisitorContext] {
	visitor := astmodel.TypeVisitorBuilder[*serializationVisitorContext]{
		VisitObjectType: applySerializationTypeTag,
	}.Build()

	return visitor
}
