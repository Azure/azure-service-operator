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
func AddSecrets(config *config.Configuration) *Stage {
	stage := NewStage(
		AddSecretsStageID,
		"Replace properties flagged as secret with genruntime.SecretReference",
		func(ctx context.Context, state *State) (*State, error) {
			types, err := applyConfigSecretOverrides(config, state.Definitions())
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
			return state.WithDefinitions(result), nil
		})

	stage.RequiresPostrequisiteStages(CreateARMTypesStageID)

	return stage
}

func applyConfigSecretOverrides(config *config.Configuration, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)

	applyConfigSecrets := func(_ *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
		typeName := ctx.(astmodel.TypeName)
		for _, prop := range it.Properties().Copy() {
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

	for _, def := range definitions {
		updatedDef, err := visitor.VisitDefinition(def, def.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "visiting type %q", def.Name())
		}

		result.Add(updatedDef)
	}

	// Verify that all 'isSecret' modifiers are consumed before returning the result
	err := config.VerifyIsSecretConsumed()
	if err != nil {
		return nil, errors.Wrap(
			err,
			"Found unused $isSecret configurations; these need to be fixed or removed.")
	}

	return result, nil
}

func transformSpecSecrets(definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	specVisitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: transformSecretProperties,
	}.Build()

	specTypes, err := astmodel.FindSpecConnectedDefinitions(definitions)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find all spec definitions")
	}

	result := make(astmodel.TypeDefinitionSet)

	for _, def := range specTypes {
		updatedDef, err := specVisitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "visiting type %q", def.Name())
		}

		result.Add(updatedDef)
	}

	return result, nil
}

func removeStatusSecrets(definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	specVisitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: removeSecretProperties,
	}.Build()

	statusTypes, err := astmodel.FindStatusConnectedDefinitions(definitions)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find all status definitions")
	}

	result := make(astmodel.TypeDefinitionSet)

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
	for _, prop := range it.Properties().Copy() {
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
	for _, prop := range it.Properties().Copy() {
		if prop.IsSecret() {
			// The expectation is that this is a string
			propType := prop.PropertyType()
			if !astmodel.Unwrap(propType).Equals(astmodel.StringType, astmodel.EqualityOverrides{}) {
				return nil, errors.Errorf("expected property %q to be a string, but was: %T", prop.PropertyName(), propType)
			}

			// check if it's optional
			required := prop.IsRequired()

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
