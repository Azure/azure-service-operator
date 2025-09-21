/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"regexp"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

type ReferenceType string

const (
	ReferenceTypeString       = ReferenceType("string")        // A simple string
	ReferenceTypeARM          = ReferenceType("arm")           // A full ARM reference
	ReferenceTypeARMWellknown = ReferenceType("arm+wellknown") // A full ARM reference OR a well known identifier
	ReferenceTypeARMCompat    = ReferenceType("arm+compat")    // A full ARM reference, but backward compatible
	ReferenceTypeUnspecified  = ReferenceType("unspecified")   // We don't know what to do with this
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

			isCrossResourceReference := func(typeName astmodel.InternalTypeName, prop *astmodel.PropertyDefinition) ReferenceType {
				// First check if we know that this property is an ARMID already
				referenceType, ok := configuration.ObjectModelConfiguration.ReferenceType.Lookup(typeName, prop.PropertyName())
				isSwaggerARMID := isTypeARMID(prop.PropertyType())

				// If we've got a Swagger ARM ID entry AND an entry in our config, that might be a problem
				if ok && isSwaggerARMID {
					switch referenceType {
					case config.ReferenceTypeSimple:
						// We permit "simple" in the config to force a reference to be handled as a string
						return ReferenceTypeString

					case config.ReferenceTypeARM:
						// Swagger has marked this field as a reference, and we also have it marked in our
						// config. Record an error saying that the config entry is no longer needed
						crossResourceReferenceErrs = append(
							crossResourceReferenceErrs,
							eris.Errorf("%s.%s marked with reference type ARM, but value is not needed because Swagger already says it is an ARM reference",
								typeName.String(),
								prop.PropertyName().String()),
						)
					case config.ReferenceTypeWellKnown:
						// We permit "arm+wellknown" in config to accommodate references that are MORE than a simple ARM ID
						return ReferenceTypeARMWellknown

					case config.ReferenceTypeCompatible:
						// We permit "arm+compat" in config to accommodate references that need to be specially handled for backward compatibility
						return ReferenceTypeARMCompat
					}
				}

				if DoesPropertyLookLikeARMReference(prop) && !ok {
					// This is an error to ensure that we don't accidentally miss adding references.
					crossResourceReferenceErrs = append(
						crossResourceReferenceErrs,
						eris.Errorf(
							"%s.%s looks like a resource reference but was not labelled as one; You may need to add it to the 'objectModelConfiguration' section of the config file",
							typeName,
							prop.PropertyName()),
					)
				}

				// Convert from configuration reference types to internal reference types
				switch referenceType {
				case config.ReferenceTypeARM:
					return ReferenceTypeARM

				case config.ReferenceTypeSimple:
					return ReferenceTypeString

				case config.ReferenceTypeWellKnown:
					return ReferenceTypeARMWellknown

				case config.ReferenceTypeCompatible:
					return ReferenceTypeARMCompat

				default:
					return ReferenceTypeUnspecified
				}
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
					return nil, eris.Wrapf(err, "visiting %q", def.Name())
				}

				typesWithARMIDs.Add(def.WithType(t))

				// TODO: Remove types that have only a single field ID and pull things up a level? Will need to wait for George's
				// TODO: Properties collapsing work for this.
			}

			var err error = kerrors.NewAggregate(crossResourceReferenceErrs)
			if err != nil {
				return nil, err
			}

			err = configuration.ObjectModelConfiguration.ReferenceType.VerifyConsumed()
			if err != nil {
				return nil, eris.Wrap(
					err,
					"Found unused $armReference configurations; these need to be fixed or removed.")
			}

			return state.WithDefinitions(typesWithARMIDs), nil
		})
}

type crossResourceReferenceChecker func(typeName astmodel.InternalTypeName, prop *astmodel.PropertyDefinition) ReferenceType

type ARMIDPropertyTypeVisitor struct {
	astmodel.TypeVisitor[astmodel.InternalTypeName]
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
		_ *astmodel.TypeVisitor[astmodel.InternalTypeName],
		it *astmodel.ObjectType,
		ctx astmodel.InternalTypeName,
	) (astmodel.Type, error) {
		var newProps []*astmodel.PropertyDefinition
		it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
			classification := visitor.isPropertyAnARMReference(ctx, prop)

			wasType := astmodel.DebugDescription(prop.PropertyType())

			switch classification {
			case ReferenceTypeARM:
				prop = makeARMReferenceProperty(prop)

			case ReferenceTypeString:
				prop = makeStringReferenceProperty(prop)

			case ReferenceTypeARMWellknown:
				prop = makeWellKnownARMReferenceProperty(prop)

			case ReferenceTypeARMCompat:
				prop = makeCompatibleARMReferenceProperty(prop)

			case ReferenceTypeUnspecified:
				// Do nothing, we don't know what this is
			}

			nowType := astmodel.DebugDescription(prop.PropertyType())

			if wasType != nowType {
				log.V(1).Info(
					"Transformed property",
					"definition", ctx,
					"property", prop.PropertyName(),
					"was", wasType,
					"now", nowType)
			}

			newProps = append(newProps, prop)
		})

		it = it.WithoutProperties()
		result := it.WithProperties(newProps...)
		return result, nil
	}

	visitor.TypeVisitor = astmodel.TypeVisitorBuilder[astmodel.InternalTypeName]{
		VisitObjectType: transformResourceReferenceProperties,
	}.Build()

	return visitor
}

var (
	armIDNameRegex        = regexp.MustCompile("(?i)(^Id$|ResourceID|ARMID)")
	armIDDescriptionRegex = regexp.MustCompile("(?i)(.*/subscriptions/.*?/resourceGroups/.*|ARMID|ARM ID|Resource ID|resourceId)")
)

// DoesPropertyLookLikeARMReference uses a simple heuristic to determine if a property looks like it might be an ARM reference.
// This can be used for logging/reporting purposes to discover references which we missed.
func DoesPropertyLookLikeARMReference(prop *astmodel.PropertyDefinition) bool {
	// The property must be a string, optional string, list of strings, or map[string]string
	mightBeReference := false
	if pt, ok := astmodel.AsPrimitiveType(prop.PropertyType()); ok {
		// Might be a reference if we have a primitive type that's a string
		mightBeReference = pt == astmodel.StringType
	}

	if at, ok := astmodel.AsArrayType(prop.PropertyType()); ok {
		// Might be references if we have an array of strings
		elementType, elementTypeIsPrimitive := astmodel.AsPrimitiveType(at.Element())
		mightBeReference = elementTypeIsPrimitive &&
			elementType == astmodel.StringType
	}

	if mt, ok := astmodel.AsMapType(prop.PropertyType()); ok {
		// Might be references if we have a map of strings to strings
		keyType, keyTypeIsPrimitive := astmodel.AsPrimitiveType(mt.KeyType())
		valueType, valueTypeIsPrimitive := astmodel.AsPrimitiveType(mt.ValueType())
		mightBeReference = keyTypeIsPrimitive && valueTypeIsPrimitive &&
			keyType == astmodel.StringType && valueType == astmodel.StringType
	}

	hasMatchingName := armIDNameRegex.MatchString(prop.PropertyName().String())
	hasMatchingDescription := armIDDescriptionRegex.MatchString(prop.Description())

	if mightBeReference && (hasMatchingName || hasMatchingDescription) {
		return true
	}

	return false
}

// makeARMReferenceProperty modifies an existing property definition into an ARM reference property.
func makeARMReferenceProperty(existing *astmodel.PropertyDefinition) *astmodel.PropertyDefinition {
	return makeReferenceProperty(existing, astmodel.ARMIDType)
}

// makeWellKnownARMReferenceProperty modifies an existing property definition into a well-known ARM reference property.
func makeWellKnownARMReferenceProperty(existing *astmodel.PropertyDefinition) *astmodel.PropertyDefinition {
	return makeReferenceProperty(existing, astmodel.ARMIDType, astmodel.WellKnownReferenceFlag)
}

// makeCompatibleARMReferenceProperty modifies an existing property definition into a ARM reference property with backward compatibility characteristics.
func makeCompatibleARMReferenceProperty(existing *astmodel.PropertyDefinition) *astmodel.PropertyDefinition {
	return makeReferenceProperty(existing, astmodel.ARMIDType, astmodel.CompatibleReferenceFlag)
}

// makeStringReferenceProperty modifies an existing property definition into a string reference property.
func makeStringReferenceProperty(existing *astmodel.PropertyDefinition) *astmodel.PropertyDefinition {
	return makeReferenceProperty(existing, astmodel.StringType)
}

// makeReferenceProperty modifies an existing property definition to use the specified type and any associated flags.
// property is the existing property definition.
// newType is the new type to use for the property.
// flags is an optional set of flags to apply.
func makeReferenceProperty(
	property *astmodel.PropertyDefinition,
	newType astmodel.Type,
	flags ...astmodel.TypeFlag,
) *astmodel.PropertyDefinition {
	_, isSlice := astmodel.AsArrayType(property.PropertyType())
	_, isMap := astmodel.AsMapType(property.PropertyType())

	var newPropType astmodel.Type

	if isSlice {
		newPropType = astmodel.NewArrayType(newType)
	} else if isMap {
		newPropType = astmodel.NewMapType(astmodel.StringType, newType)
	} else {
		newPropType = astmodel.NewOptionalType(newType)
	}

	// Apply flags
	for _, f := range flags {
		newPropType = f.ApplyTo(newPropType)
	}

	newProp := property.WithType(newPropType)

	return newProp
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
