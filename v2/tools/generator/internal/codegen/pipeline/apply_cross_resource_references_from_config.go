/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

type ARMIDPropertyClassification string

const (
	ARMIDPropertyClassificationUnset       = ARMIDPropertyClassification("unset")
	ARMIDPropertyClassificationSet         = ARMIDPropertyClassification("set")
	ARMIDPropertyClassificationUnspecified = ARMIDPropertyClassification("unspecified")
)

// ApplyCrossResourceReferencesFromConfigStageID is the unique identifier for this pipeline stage
const ApplyCrossResourceReferencesFromConfigStageID = "applyCrossResourceReferencesFromConfig"

// ApplyCrossResourceReferencesFromConfig replaces cross resource references from the configuration with astmodel.ARMID.
func ApplyCrossResourceReferencesFromConfig(
	configuration *config.Configuration,
	log logr.Logger,
) *Stage {
	return NewStage(
		ApplyCrossResourceReferencesFromConfigStageID,
		"Replace cross-resource references in the config with astmodel.ARMID",
		func(ctx context.Context, state *State) (*State, error) {
			typesWithARMIDs := make(astmodel.TypeDefinitionSet)

			var crossResourceReferenceErrs []error

			isCrossResourceReference := func(typeName astmodel.TypeName, prop *astmodel.PropertyDefinition) ARMIDPropertyClassification {
				// First check if we know that this property is an ARMID already
				isReference, err := configuration.ARMReference(typeName, prop.PropertyName())
				isSwaggerARMID := isTypeARMID(prop.PropertyType())

				// If we've got a Swagger ARM ID entry AND an entry in our config, that might be a problem
				if isSwaggerARMID && err == nil {
					if !isReference {
						// We allow overriding the ARM ID status of a property to false in our config
						return ARMIDPropertyClassificationUnset
					} else {
						// Swagger has marked this field as a reference, and we also have it marked in our
						// config. Record an error saying that the config entry is no longer needed
						crossResourceReferenceErrs = append(
							crossResourceReferenceErrs,
							errors.Errorf("%s.%s marked as ARM reference, but value is not needed because Swagger already says it is an ARM reference",
								typeName.String(),
								prop.PropertyName().String()))
					}
				}

				if DoesPropertyLookLikeARMReference(prop) && err != nil {
					if config.IsNotConfiguredError(err) {
						// This is an error for now to ensure that we don't accidentally miss adding references.
						// If/when we move to using an upstream marker for cross resource refs, we can remove this and just
						// trust the Swagger.
						crossResourceReferenceErrs = append(
							crossResourceReferenceErrs,
							errors.Wrapf(
								err,
								"%s.%s looks like a resource reference but was not labelled as one; You may need to add it to the 'objectModelConfiguration' section of the config file",
								typeName,
								prop.PropertyName()))
					} else {
						// Something else went wrong checking our configuration
						crossResourceReferenceErrs = append(
							crossResourceReferenceErrs,
							errors.Wrapf(
								err,
								"%s.%s looks like a resource reference but something went wrong checking for configuration",
								typeName,
								prop.PropertyName()))
					}
				}

				if isReference {
					return ARMIDPropertyClassificationSet
				}
				return ARMIDPropertyClassificationUnspecified
			}

			visitor := MakeARMIDPropertyTypeVisitor(isCrossResourceReference, log)
			for _, def := range state.Definitions() {
				// Skip Status types
				// TODO: we need flags
				if def.Name().IsStatus() {
					typesWithARMIDs.Add(def)
					continue
				}

				t, err := visitor.Visit(def.Type(), def.Name())
				if err != nil {
					return nil, errors.Wrapf(err, "visiting %q", def.Name())
				}

				typesWithARMIDs.Add(def.WithType(t))

				// TODO: Remove types that have only a single field ID and pull things up a level? Will need to wait for George's
				// TODO: Properties collapsing work for this.
			}

			var err error = kerrors.NewAggregate(crossResourceReferenceErrs)
			if err != nil {
				return nil, err
			}

			err = configuration.VerifyARMReferencesConsumed()
			if err != nil {
				return nil, errors.Wrap(
					err,
					"Found unused $armReference configurations; these need to be fixed or removed.")
			}

			return state.WithDefinitions(typesWithARMIDs), nil
		})
}

type crossResourceReferenceChecker func(typeName astmodel.TypeName, prop *astmodel.PropertyDefinition) ARMIDPropertyClassification

type ARMIDPropertyTypeVisitor struct {
	astmodel.TypeVisitor[astmodel.TypeName]
	// isPropertyAnARMReference is a function describing what a cross resource reference looks like. It is overridable so that
	// we can use a more simplistic criteria for tests.
	isPropertyAnARMReference crossResourceReferenceChecker
}

func MakeARMIDPropertyTypeVisitor(
	referenceChecker crossResourceReferenceChecker,
	log logr.Logger,
) ARMIDPropertyTypeVisitor {
	visitor := ARMIDPropertyTypeVisitor{
		isPropertyAnARMReference: referenceChecker,
	}

	transformResourceReferenceProperties := func(
		_ *astmodel.TypeVisitor[astmodel.TypeName],
		it *astmodel.ObjectType,
		ctx astmodel.TypeName,
	) (astmodel.Type, error) {
		var newProps []*astmodel.PropertyDefinition
		it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
			classification := visitor.isPropertyAnARMReference(ctx, prop)
			if classification == ARMIDPropertyClassificationSet {
				log.V(1).Info(
					"Transforming property",
					"definition", ctx,
					"property", prop.PropertyName(),
					"was", prop.PropertyType(),
					"now", "astmodel.ARMID")
				prop = makeARMIDProperty(prop)
			} else if classification == ARMIDPropertyClassificationUnset {
				log.V(1).Info(
					"Transforming property",
					"definition", ctx,
					"property", prop.PropertyName(),
					"was", prop.PropertyType(),
					"now", "string")
				prop = unsetARMIDProperty(prop)
			}

			newProps = append(newProps, prop)
		})

		it = it.WithoutProperties()
		result := it.WithProperties(newProps...)
		return result, nil
	}

	visitor.TypeVisitor = astmodel.TypeVisitorBuilder[astmodel.TypeName]{
		VisitObjectType: transformResourceReferenceProperties,
	}.Build()

	return visitor
}

// DoesPropertyLookLikeARMReference uses a simple heuristic to determine if a property looks like it might be an ARM reference.
// This can be used for logging/reporting purposes to discover references which we missed.
func DoesPropertyLookLikeARMReference(prop *astmodel.PropertyDefinition) bool {
	// The property must be a string, optional string, list of strings, or map[string]string
	isString := astmodel.TypeEquals(prop.PropertyType(), astmodel.StringType)
	isOptionalString := astmodel.TypeEquals(prop.PropertyType(), astmodel.OptionalStringType)
	isStringSlice := astmodel.TypeEquals(prop.PropertyType(), astmodel.NewArrayType(astmodel.StringType))
	isStringMap := astmodel.TypeEquals(prop.PropertyType(), astmodel.NewMapType(astmodel.StringType, astmodel.StringType))

	if !isString && !isOptionalString && !isStringSlice && !isStringMap {
		return false
	}

	hasMatchingDescription := armIDDescriptionRegex.MatchString(prop.Description())
	namedID := prop.HasName("Id")
	if hasMatchingDescription || namedID {
		return true
	}

	return false
}

func makeARMIDProperty(existing *astmodel.PropertyDefinition) *astmodel.PropertyDefinition {
	return makeARMIDPropertyImpl(existing, astmodel.ARMIDType)
}

func makeARMIDPropertyImpl(existing *astmodel.PropertyDefinition, newType astmodel.Type) *astmodel.PropertyDefinition {
	_, isSlice := astmodel.AsArrayType(existing.PropertyType())
	_, isMap := astmodel.AsMapType(existing.PropertyType())

	var newPropType astmodel.Type

	if isSlice {
		newPropType = astmodel.NewArrayType(newType)
	} else if isMap {
		newPropType = astmodel.NewMapType(astmodel.StringType, newType)
	} else {
		newPropType = astmodel.NewOptionalType(newType)
	}

	newProp := existing.WithType(newPropType)

	return newProp
}

func unsetARMIDProperty(existing *astmodel.PropertyDefinition) *astmodel.PropertyDefinition {
	return makeARMIDPropertyImpl(existing, astmodel.StringType)
}

// isTypeARMID determines if the type has an ARM ID somewhere inside of it
func isTypeARMID(aType astmodel.Type) bool {
	if primitiveType, ok := astmodel.AsPrimitiveType(aType); ok {
		return primitiveType == astmodel.ARMIDType
	}

	if arrayType, ok := astmodel.AsArrayType(aType); ok {
		return isTypeARMID(arrayType.Element())
	}

	if mapType, ok := astmodel.AsMapType(aType); ok {
		return isTypeARMID(mapType.KeyType()) || isTypeARMID(mapType.ValueType())
	}

	return false
}
