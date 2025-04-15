/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// RenamePropertiesStageID is the unique identifier for this pipeline stage
const RenamePropertiesStageID = "renameProperties"

func RenameProperties(cfg *config.ObjectModelConfiguration) *Stage {
	stage := NewStage(
		RenamePropertiesStageID,
		"Rename properties",
		func(_ context.Context, state *State) (*State, error) {
			visitor := createPropertyRenamingVisitor(cfg)
			var errs []error
			definitions := state.Definitions()
			modified := make(astmodel.TypeDefinitionSet)
			for name, def := range definitions {
				_, ok := astmodel.AsObjectType(def.Type())
				if !ok {
					// Not an object type, so no properties to rename
					continue
				}

				if !cfg.IsTypeConfigured(name) {
					// If this definition isn't configured, don't do anything
					continue
				}

				updated, err := visitor.VisitDefinition(def, name)
				if err != nil {
					errs = append(errs, err)
				} else {
					modified.Add(updated)
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			if err := cfg.RenamePropertyTo.VerifyConsumed(); err != nil {
				return nil, eris.Wrap(err, "verifying property rename configuration")
			}

			return state.WithOverlaidDefinitions(modified), nil
		})

	stage.RequiresPostrequisiteStages(
		AddStatusConditionsStageID, // Must rename other properties before we try to introduce the `Conditions` property
		AddOperatorSpecStageID,     // Must rename other properties before we try to introduce the `OperatorSpec` property
	)

	return stage
}

func createPropertyRenamingVisitor(
	cfg *config.ObjectModelConfiguration,
) astmodel.TypeVisitor[astmodel.InternalTypeName] {
	builder := astmodel.TypeVisitorBuilder[astmodel.InternalTypeName]{
		VisitObjectType: func(
			this *astmodel.TypeVisitor[astmodel.InternalTypeName],
			it *astmodel.ObjectType,
			ctx astmodel.InternalTypeName,
		) (astmodel.Type, error) {
			return renamePropertiesInObjectType(ctx, it, cfg)
		},
	}

	return builder.Build()
}

func renamePropertiesInObjectType(
	typeName astmodel.InternalTypeName,
	ot *astmodel.ObjectType,
	cfg *config.ObjectModelConfiguration,
) (astmodel.Type, error) {
	properties := make([]*astmodel.PropertyDefinition, 0, ot.Properties().Len())
	err := ot.Properties().ForEachError(
		func(prop *astmodel.PropertyDefinition) error {
			name, ok := cfg.RenamePropertyTo.Lookup(typeName, prop.PropertyName())
			if !ok {
				properties = append(properties, prop)
				return nil
			}

			updated := prop.
				WithName(astmodel.PropertyName(name)).
				WithJSONName(strings.ToLower(name))
			properties = append(properties, updated)

			return nil
		})
	if err != nil {
		return nil, eris.Wrapf(err, "renaming properties for %s", typeName)
	}

	return ot.WithoutProperties().WithProperties(properties...), nil
}
