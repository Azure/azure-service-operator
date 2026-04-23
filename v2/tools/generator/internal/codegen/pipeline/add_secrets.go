/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"regexp"
	"strings"

	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

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
				return nil, eris.Wrap(err, "applying config secret overrides")
			}

			updatedSpecs, err := transformSpecSecrets(types)
			if err != nil {
				return nil, eris.Wrap(err, "transforming spec secrets")
			}

			updatedStatuses, err := removeStatusSecrets(types)
			if err != nil {
				return nil, eris.Wrap(err, "removing status secrets")
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

			secrecy, secrecyConfigured := config.ObjectModelConfiguration.Secrecy.Lookup(ctx, prop.PropertyName())
			if ctx.IsStatus() && !secrecyConfigured {
				secrecy, secrecyConfigured = config.ObjectModelConfiguration.Secrecy.Lookup(strippedTypeName, prop.PropertyName())
			}

			// If it's not a secret, but it looks like a secret, and we don't have any configuration to tell us for
			// sure, request configuration so we know for sure.
			if prop.Secrecy() != astmodel.ImportSecretModeRequired && prop.Secrecy() != astmodel.ImportSecretModeOptional && maybeSecret && !secrecyConfigured {
				// Property might be a secret, but isn't already configured as one,
				// and we don't have config to tell us for sure
				return nil, eris.Errorf(
					"property %s might be a secret and must be configured with $importSecretMode",
					prop.PropertyName())
			}

			if secrecyConfigured {
				it = it.WithProperty(prop.WithSecrecy(secrecy))
			}
		}

		return it, nil
	}

	visitor := astmodel.TypeVisitorBuilder[astmodel.InternalTypeName]{
		VisitObjectType: applyConfigSecrets,
	}.Build()

	var errs []error
	for _, def := range definitions {
		if def.Name().IsARMType() {
			// No need to process ARM types
			continue
		}

		updatedDef, err := visitor.VisitDefinition(def, def.Name())
		if err != nil {
			errs = append(errs, eris.Wrapf(err, "visiting type %q", def.Name()))
			continue
		}

		result.Add(updatedDef)
	}

	if len(errs) > 0 {
		return nil, eris.Wrap(
			kerrors.NewAggregate(errs),
			"encountered errors while applying config secrets")
	}

	// Verify that all 'isSecret' modifiers are consumed before returning the result
	err := config.ObjectModelConfiguration.Secrecy.VerifyConsumed()
	if err != nil {
		return nil, eris.Wrap(
			err,
			"Found unused $importSecretMode configurations; these need to be fixed or removed.")
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

	// If the property name matches a detector, that tells us
	// whether to expect it's a secret or not
	propertyName := string(prop.PropertyName())
	for _, detector := range secretDetectors {
		if detector.regex.MatchString(propertyName) {
			return detector.isSecret
		}
	}

	return false
}

// Rules for detecting potentially secret properties.
// These are processed in order, with the first matching rule being used.
var secretDetectors = []struct {
	regex    regexp.Regexp // Regular expression to match
	isSecret bool          // Whether to treat the property as a secret
}{
	{
		// Look for the word `password` in any position
		regex:    *regexp.MustCompile(`(?i)password`),
		isSecret: true,
	},
	{
		// Look for the word `token` in any position
		regex:    *regexp.MustCompile(`(?i)token`),
		isSecret: true,
	},
	{
		// a PublicKey is not a secret
		regex:    *regexp.MustCompile(`(?i)publickey`),
		isSecret: false,
	},
	{
		// KeyData is always a secret
		regex:    *regexp.MustCompile(`(?i)keydata`),
		isSecret: true,
	},
	{
		// URLs and URIs are not secrets
		regex:    *regexp.MustCompile(`(?i)url|uri`),
		isSecret: false,
	},
	{
		// IDs, Identifiers, and Names are not secrets, they're used to look them up
		// (must match at the end)
		regex:    *regexp.MustCompile(`(?i)(id|identifier|Identity|name)$`),
		isSecret: false,
	},
}

func transformSpecSecrets(definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	specVisitor := astmodel.TypeVisitorBuilder[any]{
		VisitObjectType: transformSecretProperties,
	}.Build()

	specTypes, err := astmodel.FindSpecConnectedDefinitions(definitions)
	if err != nil {
		return nil, eris.Wrap(err, "couldn't find all spec definitions")
	}

	result := make(astmodel.TypeDefinitionSet)

	for _, def := range specTypes {
		updatedDef, err := specVisitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, eris.Wrapf(err, "visiting type %q", def.Name())
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
		return nil, eris.Wrap(err, "couldn't find all status definitions")
	}

	result := make(astmodel.TypeDefinitionSet)

	for _, def := range statusTypes {
		updatedDef, err := specVisitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, eris.Wrapf(err, "visiting type %q", def.Name())
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
		switch prop.Secrecy() {
		case astmodel.ImportSecretModeRequired, astmodel.ImportSecretModeOptional:
			propType := prop.PropertyType()

			// We only remove pure secret references here. For the case of secret maps, different services seem to treat them
			// differently. Some services (such as Microsoft.KubernetesConfiguration/extensions) will return the keys of the map
			// but not the values. Other services (such as APIM) will return certain keys and values that it knows are non-secret, but
			// redact the ones that are secret. Since it's hard to know statically what will be returned for any given service, we
			// default to having the map[string]string on the Status type and letting the service return what it wants.
			if isTypeSecretMapCandidate(propType) {
				it = it.WithProperty(prop.WithSecrecy(astmodel.ImportSecretModeNever))
				continue
			}

			if !isTypeSecretReferenceCandidate(propType) {
				return nil, eris.Errorf(
					"expected property %q to be a string, optional string, map[string]string, or []string, but was: %s",
					prop.PropertyName(),
					astmodel.DebugDescription(propType))
			}

			it = it.WithoutProperty(prop.PropertyName())
		case astmodel.ImportSecretModeNever:
			// Not a secret, nothing to do
		}
	}

	return it, nil
}

func transformSecretProperties(_ *astmodel.TypeVisitor[any], it *astmodel.ObjectType, _ any) (astmodel.Type, error) {
	for _, prop := range it.Properties().Copy() {
		switch prop.Secrecy() {
		case astmodel.ImportSecretModeRequired, astmodel.ImportSecretModeOptional:
			propType := prop.PropertyType()

			if !isTypeSecretReferenceCandidate(propType) {
				return nil, eris.Errorf(
					"expected property %q to be a string, optional string, map[string]string, or []string, but was: %s",
					prop.PropertyName(),
					astmodel.DebugDescription(propType))
			}

			// Work out the secret reference type
			var newType astmodel.Type
			if isTypeSecretSliceCandidate(propType) {
				newType = astmodel.NewArrayType(astmodel.SecretReferenceType)
			} else if isTypeSecretMapCandidate(propType) {
				newType = astmodel.OptionalSecretMapReferenceType
			} else if _, ok := astmodel.AsOptionalType(propType); ok {
				newType = astmodel.NewOptionalType(astmodel.SecretReferenceType)
			} else {
				newType = astmodel.SecretReferenceType
			}

			if prop.Secrecy() == astmodel.ImportSecretModeOptional {
				// For optional secrets, create a dual-field pair: original value + new SecretReference
				updatedProp, newProp, err := createNewSecretReference(prop, newType)
				if err != nil {
					return nil, eris.Wrapf(err, "failed to create optional secret pair for property %s", prop.PropertyName())
				}

				it = it.WithProperty(updatedProp)

				// If the property we're about to add already exists, that's bad!
				if _, ok := it.Property(newProp.PropertyName()); ok {
					return nil, eris.Errorf(
						"property %q already exists on type, can't create optional secret pair",
						newProp.PropertyName())
				}

				it = it.WithProperty(newProp)
			} else {
				// For always-secret properties, replace with the secret reference type
				it = it.WithProperty(prop.WithType(newType))
			}
		case astmodel.ImportSecretModeNever:
			// Not a secret, nothing to do
		}
	}

	return it, nil
}

func createNewSecretReference(
	prop *astmodel.PropertyDefinition,
	newType astmodel.Type,
) (*astmodel.PropertyDefinition, *astmodel.PropertyDefinition, error) {
	jsonName, ok := prop.JSONName()
	if !ok {
		return nil, nil, eris.Errorf("property %s didn't have a JSON name", prop.PropertyName())
	}

	// Neither property can be required anymore.
	updatedProp := prop.
		WithTag(astmodel.OptionalSecretPairTag, string(prop.PropertyName())).
		WithSecrecy(astmodel.ImportSecretModeNever). // Clear secret flag on the plain value property
		MakeOptional().
		MakeTypeOptional()

	newProp := prop.
		WithName(prop.PropertyName()+astmodel.OptionalSecretReferenceSuffix).
		WithType(newType).
		WithJSONName(jsonName+astmodel.OptionalSecretReferenceSuffix).
		WithTag(astmodel.OptionalSecretPairTag, string(prop.PropertyName())).
		WithSecrecy(astmodel.ImportSecretModeNever). // The FromSecret property is a reference, not itself a secret
		MakeOptional().
		MakeTypeOptional()

	return updatedProp, newProp, nil
}
