/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// AddSecretsStageID is the unique identifier for this pipeline stage
const AddSecretsStageID = "addSecrets"

// AddSecrets replaces properties flagged as secret with genruntime.SecretReference
func AddSecrets(_ *config.Configuration) Stage { // TODO: Will need this configuration parameter eventually to allow config overrides
	stage := MakeStage(
		AddSecretsStageID,
		"Replace properties flagged as secret with genruntime.SecretReference",
		func(ctx context.Context, state *State) (*State, error) {

			updatedSpecs, err := transformSpecSecrets(state.Types())
			if err != nil {
				return nil, errors.Wrap(err, "transforming spec secrets")
			}

			updatedStatuses, err := removeStatusSecrets(state.Types())
			if err != nil {
				return nil, errors.Wrap(err, "removing status secrets")
			}

			result := state.Types().OverlayWith(astmodel.TypesDisjointUnion(updatedSpecs, updatedStatuses))
			return state.WithTypes(result), nil
		})

	return stage.
		RequiresPrerequisiteStages(AugmentSpecWithStatusStageID).
		RequiresPostrequisiteStages(CreateARMTypesStageID)
}

func transformSpecSecrets(types astmodel.Types) (astmodel.Types, error) {
	specVisitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: transformSecretProperties,
	}.Build()

	specTypes, err := astmodel.FindSpecConnectedTypes(types)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find all spec types")
	}

	result := make(astmodel.Types)

	for _, def := range specTypes {
		updatedDef, err := specVisitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "visiting type %q", def.Name())
		}

		result.Add(updatedDef)
	}

	return result, nil
}

func removeStatusSecrets(types astmodel.Types) (astmodel.Types, error) {
	specVisitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: removeSecretProperties,
	}.Build()

	statusTypes, err := astmodel.FindStatusConnectedTypes(types)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find all status types")
	}

	result := make(astmodel.Types)

	for _, def := range statusTypes {
		updatedDef, err := specVisitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "visiting type %q", def.Name())
		}

		result.Add(updatedDef)
	}

	return result, nil
}

func removeSecretProperties(_ *astmodel.TypeVisitor, it *astmodel.ObjectType, _ interface{}) (astmodel.Type, error) {
	for _, prop := range it.Properties() {
		if prop.IsSecret() {
			// The expectation is that this is a string
			propType := prop.PropertyType()
			if !astmodel.Unwrap(propType).Equals(astmodel.StringType, astmodel.EqualityOverrides{}) {
				return nil, errors.Errorf("expected property %q to be a string, but was: %T", prop.PropertyName(), propType)
			}

			it = it.WithoutProperty(prop.PropertyName())
		}
	}

	return it, nil
}

func transformSecretProperties(_ *astmodel.TypeVisitor, it *astmodel.ObjectType, _ interface{}) (astmodel.Type, error) {
	for _, prop := range it.Properties() {
		if prop.IsSecret() {
			// The expectation is that this is a string
			propType := prop.PropertyType()
			if !astmodel.Unwrap(propType).Equals(astmodel.StringType, astmodel.EqualityOverrides{}) {
				return nil, errors.Errorf("expected property %q to be a string, but was: %T", prop.PropertyName(), propType)
			}

			// check if it's optional
			required := prop.HasKubebuilderRequiredValidation()

			var newType astmodel.Type
			if required {
				newType = astmodel.SecretReferenceType
			} else {
				newType = astmodel.NewOptionalType(astmodel.SecretReferenceType)
			}
			it = it.WithProperty(prop.WithType(newType))
		}
	}

	return it, nil
}
