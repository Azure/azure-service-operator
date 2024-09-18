/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"regexp"
	"strings"

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

			return state.WithOverlaidDefinitions(astmodel.TypesDisjointUnion(updatedSpecs, updatedStatuses)), nil
		})

	stage.RequiresPostrequisiteStages(CreateARMTypesStageID)

	return stage
}

func applyConfigSecretOverrides(
	config *config.Configuration,
	definitions astmodel.TypeDefinitionSet,
) (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)

	applyConfigSecrets := func(
		_ *astmodel.TypeVisitor[astmodel.InternalTypeName],
		it *astmodel.ObjectType,
		ctx astmodel.InternalTypeName,
	) (astmodel.Type, error) {
		strippedTypeName := ctx.WithName(strings.TrimSuffix(ctx.Name(), astmodel.StatusSuffix))

		for _, prop := range it.Properties().Copy() {
			maybeSecret := mightBeSecretProperty(prop, definitions)

			isSecret, isSecretConfigured := config.ObjectModelConfiguration.IsSecret.Lookup(ctx, prop.PropertyName())
			if ctx.IsStatus() && !isSecretConfigured {
				isSecret, isSecretConfigured = config.ObjectModelConfiguration.IsSecret.Lookup(strippedTypeName, prop.PropertyName())
			}

			if maybeSecret && !prop.IsSecret() && !isSecretConfigured {
				// Property might be a secret, but isn't already configured as one,
				// and we don't have config to tell us for sure
				return nil, errors.Errorf(
					"property %q might be a secret and must to be configured with $isSecret",
					prop.PropertyName())
			}

			if isSecretConfigured {
				propWithSecret := prop.WithIsSecret(isSecret)
				it = it.WithProperty(propWithSecret)
			}
		}

		return it, nil
	}

	visitor := astmodel.TypeVisitorBuilder[astmodel.InternalTypeName]{
		VisitObjectType: applyConfigSecrets,
	}.Build()

	for _, def := range definitions {
		if def.Name().IsARMType() {
			// No need to process ARM types
			continue
		}

		updatedDef, err := visitor.VisitDefinition(def, def.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "visiting type %q", def.Name())
		}

		result.Add(updatedDef)
	}

	// Verify that all 'isSecret' modifiers are consumed before returning the result
	err := config.ObjectModelConfiguration.IsSecret.VerifyConsumed()
	if err != nil {
		return nil, errors.Wrap(
			err,
			"Found unused $isSecret configurations; these need to be fixed or removed.")
	}

	return result, nil
}

// mightBeSecret returns true if the given name might represent a secret that shouldn't be present
// in plain text in the resource. This is a heuristic used to require the presence of $isSecret
// configuration so that we know for sure.
// property is the property to check
// definitions is the set of all definitions so we can look up a typename
func mightBeSecretProperty(
	prop *astmodel.PropertyDefinition,
	definitions astmodel.TypeDefinitionSet,
) bool {
	// Only properties that are strings can be secrets
	propertyType := prop.PropertyType()

	// Look through a reference if we have one
	if tn, ok := astmodel.AsInternalTypeName(propertyType); ok {
		if def, ok := definitions[tn]; ok {
			propertyType = def.Type()
		}
	}

	// If not a string type, can't be a secret
	pt, ok := astmodel.AsPrimitiveType(propertyType)
	if !ok || pt != astmodel.StringType {
		return false
	}

	// If the property name matches any of our detectors, it might be a secret
	propertyName := string(prop.PropertyName())
	for _, detector := range secretDetectors {
		if detector.MatchString(propertyName) {
			return true
		}
	}

	return false
}

// A list of regexes for detecting potentially secret properties
// Lowercase only, please.
var secretDetectors = []regexp.Regexp{
	// Look for the word `password` in any position
	*regexp.MustCompile(`(?i)password`),
	// Look for the word `token` in any position
	*regexp.MustCompile(`(?i)token`),
}

func transformSpecSecrets(definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	specVisitor := astmodel.TypeVisitorBuilder[any]{
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
	specVisitor := astmodel.TypeVisitorBuilder[any]{
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

func isTypeSecretReferenceCandidate(t astmodel.Type) bool {
	isStringOrOptionalString := astmodel.TypeEquals(astmodel.Unwrap(t), astmodel.StringType)

	isStringSlice := isTypeSecretSliceCandidate(t)
	isStringMap := isTypeSecretMapCandidate(t)

	return isStringOrOptionalString || isStringSlice || isStringMap
}

func isTypeSecretSliceCandidate(t astmodel.Type) bool {
	return astmodel.TypeEquals(t, astmodel.NewArrayType(astmodel.StringType))
}

func isTypeSecretMapCandidate(t astmodel.Type) bool {
	return astmodel.TypeEquals(t, astmodel.MapOfStringStringType)
}

func removeSecretProperties(_ *astmodel.TypeVisitor[any], it *astmodel.ObjectType, _ any) (astmodel.Type, error) {
	for _, prop := range it.Properties().Copy() {
		if prop.IsSecret() {
			propType := prop.PropertyType()

			// We only remove pure secret references here. For the case of secret maps, different services seem to treat them
			// differently. Some services (such as Microsoft.KubernetesConfiguration/extensions) will return the keys of the map
			// but not the values. Other services (such as APIM) will return certain keys and values that it knows are non-secret, but
			// redact the ones that are secret. Since it's hard to know statically what will be returned for any given service, we
			// default to having the map[string]string on the Status type and letting the service return what it wants.
			if isTypeSecretMapCandidate(propType) {
				it = it.WithProperty(prop.WithIsSecret(false))
				continue
			}

			if !isTypeSecretReferenceCandidate(propType) {
				return nil, errors.Errorf("expected property %q to be a string, optional string, map[string]string, or []string, but was: %q", prop.PropertyName(), astmodel.DebugDescription(propType))
			}

			it = it.WithoutProperty(prop.PropertyName())
		}
	}

	return it, nil
}

func transformSecretProperties(_ *astmodel.TypeVisitor[any], it *astmodel.ObjectType, _ any) (astmodel.Type, error) {
	for _, prop := range it.Properties().Copy() {
		if prop.IsSecret() {
			propType := prop.PropertyType()

			if !isTypeSecretReferenceCandidate(propType) {
				return nil, errors.Errorf("expected property %q to be a string, optional string, map[string]string, or []string, but was: %T", prop.PropertyName(), astmodel.DebugDescription(propType))
			}

			var newType astmodel.Type
			if isTypeSecretSliceCandidate(prop.PropertyType()) {
				newType = astmodel.NewArrayType(astmodel.SecretReferenceType)
			} else if isTypeSecretMapCandidate(prop.PropertyType()) {
				newType = astmodel.OptionalSecretMapReferenceType
			} else if _, ok := astmodel.AsOptionalType(prop.PropertyType()); ok {
				newType = astmodel.NewOptionalType(astmodel.SecretReferenceType)
			} else {
				newType = astmodel.SecretReferenceType
			}

			updatedProp := prop.WithType(newType)
			it = it.WithProperty(updatedProp)
		}
	}

	return it, nil
}
