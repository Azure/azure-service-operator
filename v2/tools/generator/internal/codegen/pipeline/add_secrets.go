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
func AddSecrets(config *config.Configuration) Stage {
	stage := MakeStage(
		AddSecretsStageID,
		"Replace properties flagged as secret with genruntime.SecretReference",
		func(ctx context.Context, state *State) (*State, error) {

			types, err := applyConfigSecretOverrides(config, state.Types())
			if err != nil {
				return nil, errors.Wrap(err, "applying config secret overrides")
			}

			updatedSpecs, err := transformSpecSecrets(types)
			if err != nil {
				return nil, errors.Wrap(err, "transforming spec secrets")
			}

			updatedStatuses, err := removeStatusSecrets(types)
			if err != nil {
				return nil, errors.Wrap(err, "removing status secrets")
			}

			result := types.OverlayWith(astmodel.TypesDisjointUnion(updatedSpecs, updatedStatuses))
			return state.WithTypes(result), nil
		})

	return stage.
		RequiresPrerequisiteStages(AugmentSpecWithStatusStageID).
		RequiresPostrequisiteStages(CreateARMTypesStageID)
}

func applyConfigSecretOverrides(config *config.Configuration, types astmodel.Types) (astmodel.Types, error) {
	result := make(astmodel.Types)

	applyConfigSecrets := func(_ *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
		typeName := ctx.(astmodel.TypeName)
		for _, prop := range it.Properties() {
			isSecret, _ := config.IsSecret(typeName, prop.PropertyName())
			if isSecret {
				it = it.WithProperty(prop.WithIsSecret(true))
			}
		}

		return it, nil
	}

	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: applyConfigSecrets,
	}.Build()

	for _, def := range types {
		updatedDef, err := visitor.VisitDefinition(def, def.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "visiting type %q", def.Name())
		}

		result.Add(updatedDef)
	}

	return result, nil
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
