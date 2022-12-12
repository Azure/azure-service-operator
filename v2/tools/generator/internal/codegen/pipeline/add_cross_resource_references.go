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
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// TransformCrossResourceReferencesStageID is the unique identifier for this pipeline stage
const TransformCrossResourceReferencesStageID = "transformCrossResourceReferences"

var armIDDescriptionRegex = regexp.MustCompile("(?i)(.*/subscriptions/.*?/resourceGroups/.*|ARM ID|Resource ID|resourceId)")

// TransformCrossResourceReferences replaces cross resource references with genruntime.ResourceReference.
func TransformCrossResourceReferences(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) *Stage {
	return NewStage(
		TransformCrossResourceReferencesStageID,
		"Replace cross-resource references with genruntime.ResourceReference",
		func(ctx context.Context, state *State) (*State, error) {
			typesWithARMIDs := make(astmodel.TypeDefinitionSet)

			var crossResourceReferenceErrs []error

			isCrossResourceReference := func(typeName astmodel.TypeName, prop *astmodel.PropertyDefinition) bool {
				// First check if we know that this property is an ARMID already
				isReference, err := configuration.ARMReference(typeName, prop.PropertyName())
				if primitive, ok := astmodel.AsPrimitiveType(prop.PropertyType()); ok {
					if primitive == astmodel.ARMIDType {
						if err == nil {
							if !isReference {
								// We've overridden the ARM spec details
								return false
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

						return true
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

				return isReference
			}

			visitor := MakeCrossResourceReferenceTypeVisitor(idFactory, isCrossResourceReference)
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

			updatedDefs, err := stripARMIDPrimitiveTypes(typesWithARMIDs)
			if err != nil {
				return nil, errors.Wrap(err, "failed to strip ARM ID primitive types")
			}

			err = configuration.VerifyARMReferencesConsumed()
			if err != nil {
				klog.Error(err)

				return nil, errors.Wrap(
					err,
					"Found unused $armReference configurations; these need to be fixed or removed.")
			}

			return state.WithDefinitions(updatedDefs), nil
		})
}

type crossResourceReferenceChecker func(typeName astmodel.TypeName, prop *astmodel.PropertyDefinition) bool

type CrossResourceReferenceTypeVisitor struct {
	astmodel.TypeVisitor
	// referenceChecker is a function describing what a cross resource reference looks like. It is overridable so that
	// we can use a more simplistic criteria for tests.
	isPropertyAnARMReference crossResourceReferenceChecker
}

func MakeCrossResourceReferenceTypeVisitor(idFactory astmodel.IdentifierFactory, referenceChecker crossResourceReferenceChecker) CrossResourceReferenceTypeVisitor {
	visitor := CrossResourceReferenceTypeVisitor{
		isPropertyAnARMReference: referenceChecker,
	}

	transformResourceReferenceProperties := func(_ *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
		typeName := ctx.(astmodel.TypeName)

		var newProps []*astmodel.PropertyDefinition
		it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
			if visitor.isPropertyAnARMReference(typeName, prop) {
				klog.V(4).Infof("Transforming \"%s.%s\" field into genruntime.ResourceReference", typeName, prop.PropertyName())
				originalName := string(prop.PropertyName())
				prop = makeResourceReferenceProperty(idFactory, prop)

				// TODO: We could pass this information forward some other way?
				// Add tag so that we remember what this field was before
				prop = prop.WithTag(astmodel.ARMReferenceTag, originalName)
			}

			newProps = append(newProps, prop)
		})

		it = it.WithoutProperties()
		result := it.WithProperties(newProps...)
		return result, nil
	}

	visitor.TypeVisitor = astmodel.TypeVisitorBuilder{
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

func makeResourceReferenceProperty(idFactory astmodel.IdentifierFactory, existing *astmodel.PropertyDefinition) *astmodel.PropertyDefinition {
	_, isSlice := astmodel.AsArrayType(existing.PropertyType())
	_, isMap := astmodel.AsMapType(existing.PropertyType())
	propertyNameSuffix := "Reference"
	if isSlice || isMap {
		propertyNameSuffix = "References"
	}

	var referencePropertyName string
	// Special case for "Id" and properties that end in "Id", which are quite common in the specs. This is primarily
	// because it's awkward to have a field called "Id" not just be a string and instead but a complex type describing
	// a reference.
	if existing.PropertyName() == "Id" {
		referencePropertyName = propertyNameSuffix
	} else if strings.HasSuffix(string(existing.PropertyName()), "Id") {
		referencePropertyName = strings.TrimSuffix(string(existing.PropertyName()), "Id") + propertyNameSuffix
	} else if strings.HasSuffix(string(existing.PropertyName()), "Ids") {
		referencePropertyName = strings.TrimSuffix(string(existing.PropertyName()), "Ids") + propertyNameSuffix
	} else {
		referencePropertyName = string(existing.PropertyName()) + propertyNameSuffix
	}

	var newPropType astmodel.Type

	if isSlice {
		newPropType = astmodel.NewArrayType(astmodel.ResourceReferenceType)
	} else if isMap {
		newPropType = astmodel.NewMapType(astmodel.StringType, astmodel.ResourceReferenceType)
	} else {
		newPropType = astmodel.NewOptionalType(astmodel.ResourceReferenceType)
	}

	newProp := astmodel.NewPropertyDefinition(
		idFactory.CreatePropertyName(referencePropertyName, astmodel.Exported),
		idFactory.CreateStringIdentifier(referencePropertyName, astmodel.NotExported),
		newPropType)

	newProp = newProp.WithDescription(existing.Description())
	if existing.IsRequired() {
		newProp = newProp.MakeRequired()
	} else {
		newProp = newProp.MakeOptional()
	}

	return newProp
}

func stripARMIDPrimitiveTypes(types astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	// Any astmodel.ARMReference's which have escaped need to be turned into
	// strings
	armIDStrippingVisitor := astmodel.TypeVisitorBuilder{
		VisitPrimitive: func(_ *astmodel.TypeVisitor, it *astmodel.PrimitiveType, ctx interface{}) (astmodel.Type, error) {
			if astmodel.TypeEquals(it, astmodel.ARMIDType) {
				return astmodel.StringType, nil
			}
			return it, nil
		},
	}.Build()

	result := make(astmodel.TypeDefinitionSet)
	for _, def := range types {
		newDef, err := armIDStrippingVisitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, err
		}
		result.Add(newDef)
	}

	return result, nil
}
