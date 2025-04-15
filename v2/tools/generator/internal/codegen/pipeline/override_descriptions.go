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

// OverrideDescriptionsStageID is the unique identifier for this pipeline stage
const OverrideDescriptionsStageID = "overrideDescriptions"

// OverrideDescriptions overrides the specified property descriptions
func OverrideDescriptions(configuration *config.Configuration) *Stage {
	stage := NewStage(
		OverrideDescriptionsStageID,
		"Applies the configured description overrides",
		func(ctx context.Context, state *State) (*State, error) {
			visitor := createDescriptionOverrideVisitor(configuration)
			result := make(astmodel.TypeDefinitionSet)
			for _, def := range state.Definitions() {
				newDef, err := visitor.VisitDefinition(def, def.Name())
				if err != nil {
					return nil, err
				}

				result.Add(newDef)
			}

			if err := configuration.ObjectModelConfiguration.Description.VerifyConsumed(); err != nil {
				return nil, eris.Wrap(err, "verifying description augmentation")
			}

			return state.WithDefinitions(result), nil
		})

	stage.RequiresPrerequisiteStages(FlattenPropertiesStageID, RenamePropertiesStageID)
	return stage
}

func createDescriptionOverrideVisitor(
	configuration *config.Configuration,
) astmodel.TypeVisitor[astmodel.InternalTypeName] {
	visitor := astmodel.TypeVisitorBuilder[astmodel.InternalTypeName]{
		VisitObjectType: func(
			this *astmodel.TypeVisitor[astmodel.InternalTypeName],
			it *astmodel.ObjectType,
			ctx astmodel.InternalTypeName,
		) (astmodel.Type, error) {
			result := it
			name := ctx
			for _, prop := range it.Properties().AsSlice() {
				// If the property has an overridden description, use it
				description, ok := configuration.ObjectModelConfiguration.Description.Lookup(name, prop.PropertyName())
				if !ok {
					continue
				}
				prop = prop.WithDescription(description)
				result = result.WithProperty(prop)
			}

			return result, nil
		},
	}

	return visitor.Build()
}
