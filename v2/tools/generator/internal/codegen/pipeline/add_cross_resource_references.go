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

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// TransformCrossResourceReferencesStageID is the unique identifier for this pipeline stage
const TransformCrossResourceReferencesStageID = "transformCrossResourceReferences"

var armIDDescriptionRegex = regexp.MustCompile("(?i)(.*/subscriptions/.*?/resourceGroups/.*|ARM ID|Resource ID|resourceId)")
var idRegex = regexp.MustCompile("^(.*)I[d|D]s?$")

// TransformCrossResourceReferences replaces cross resource references with genruntime.ResourceReference.
func TransformCrossResourceReferences(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) *Stage {
	return NewStage(
		TransformCrossResourceReferencesStageID,
		"Replace cross-resource references with genruntime.ResourceReference",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs := make(astmodel.TypeDefinitionSet)

			visitor := MakeARMIDToResourceReferenceTypeVisitor(idFactory)
			for _, def := range state.Definitions() {
				// Skip Status types
				// TODO: we need flags
				if def.Name().IsStatus() {
					updatedDefs.Add(def)
					continue
				}

				t, err := visitor.Visit(def.Type(), def.Name())
				if err != nil {
					return nil, errors.Wrapf(err, "visiting %q", def.Name())
				}

				updatedDefs.Add(def.WithType(t))
			}

			resultDefs, err := stripARMIDPrimitiveTypes(updatedDefs)
			if err != nil {
				return nil, errors.Wrap(err, "failed to strip ARM ID primitive types")
			}

			return state.WithDefinitions(resultDefs), nil
		})
}

func makeReferencePropertyName(existing *astmodel.PropertyDefinition, isSlice bool, isMap bool) string {
	propertyNameSuffix := "Reference"
	if isSlice || isMap {
		propertyNameSuffix = "References"
	}

	var referencePropertyName string
	// Special case for "Id" and properties that end in "Id", which are quite common in the specs. This is primarily
	// because it's awkward to have a field called "Id" not just be a string and instead but a complex type describing
	// a reference.
	s := existing.PropertyName().String()

	if strings.ToLower(s) == "id" {
		referencePropertyName = propertyNameSuffix
	} else if idRegex.MatchString(s) {
		referencePropertyName = idRegex.ReplaceAllString(s, "${1}"+propertyNameSuffix)
	} else {
		referencePropertyName = s + propertyNameSuffix
	}

	return referencePropertyName
}

// makeLegacyReferencePropertyName does not correctly deal with properties with "ID" suffix, but exists
// to ensure backward compatibility with old versions.
// See https://github.com/Azure/azure-service-operator/issues/2501#issuecomment-1251650714
func makeLegacyReferencePropertyName(existing *astmodel.PropertyDefinition, isSlice bool, isMap bool) string {
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

	return referencePropertyName
}

type ARMIDToGenruntimeReferenceTypeVisitor struct {
	astmodel.TypeVisitor
	idFactory astmodel.IdentifierFactory
}

func (v *ARMIDToGenruntimeReferenceTypeVisitor) transformARMIDToGenruntimeResourceReference(
	_ *astmodel.TypeVisitor,
	it *astmodel.PrimitiveType,
	_ interface{},
) (astmodel.Type, error) {
	if it == astmodel.ARMIDType {
		return astmodel.ResourceReferenceType, nil
	}

	return it, nil
}

func (v *ARMIDToGenruntimeReferenceTypeVisitor) renamePropertiesWithARMIDReferences(
	this *astmodel.TypeVisitor,
	it *astmodel.ObjectType,
	ctx interface{},
) (astmodel.Type, error) {
	// First, we visit this type like normal
	result, err := astmodel.IdentityVisitOfObjectType(this, it, ctx)
	if err != nil {
		return nil, err
	}

	typeName := ctx.(astmodel.TypeName)
	ot, ok := result.(*astmodel.ObjectType)
	if !ok {
		return nil, errors.Errorf("result for visitObjectType of %s was not expected ObjectType, instead %T", typeName, result)
	}

	// Now, check if any properties have been updated and if they have change their names to match
	var newProps []*astmodel.PropertyDefinition
	var errs []error
	ot.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
		origProp, ok := it.Property(prop.PropertyName())
		if !ok {
			errs = append(errs, errors.Errorf("expected to find property %q on %s", prop.PropertyName(), typeName))
			return
		}

		if origProp.Equals(prop, astmodel.EqualityOverrides{}) {
			newProps = append(newProps, prop)
			return
		}

		newProp := makeResourceReferenceProperty(typeName, v.idFactory, prop)
		newProps = append(newProps, newProp)
	})

	err = kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	ot = ot.WithoutProperties()
	ot = ot.WithProperties(newProps...)
	return ot, nil
}

func (v *ARMIDToGenruntimeReferenceTypeVisitor) stripValidationForResourceReferences(
	this *astmodel.TypeVisitor,
	it *astmodel.ValidatedType,
	ctx interface{},
) (astmodel.Type, error) {
	result, err := astmodel.IdentityVisitOfValidatedType(this, it, ctx)
	if err != nil {
		return nil, err
	}

	if astmodel.TypeEquals(it, result) {
		return result, nil
	}

	// If they aren't equal, just return the element
	validated, ok := result.(*astmodel.ValidatedType)
	if !ok {
		return nil, errors.Errorf("expected IdentityVisitOfValidatedType to return a ValidatedType, but it instead returned %T", result)
	}
	return validated.ElementType(), nil
}

func MakeARMIDToResourceReferenceTypeVisitor(idFactory astmodel.IdentifierFactory) ARMIDToGenruntimeReferenceTypeVisitor {
	result := ARMIDToGenruntimeReferenceTypeVisitor{
		idFactory: idFactory,
	}
	result.TypeVisitor = astmodel.TypeVisitorBuilder{
		VisitPrimitive:     result.transformARMIDToGenruntimeResourceReference,
		VisitObjectType:    result.renamePropertiesWithARMIDReferences,
		VisitValidatedType: result.stripValidationForResourceReferences,
	}.Build()

	return result
}

func makeResourceReferenceProperty(
	typeName astmodel.TypeName,
	idFactory astmodel.IdentifierFactory,
	existing *astmodel.PropertyDefinition) *astmodel.PropertyDefinition {

	_, isSlice := astmodel.AsArrayType(existing.PropertyType())
	_, isMap := astmodel.AsMapType(existing.PropertyType())
	var referencePropertyName string
	// This is hacky but works
	if group, version, ok := typeName.PackageReference().TryGroupVersion(); ok && group == "containerservice" && strings.Contains(version, "20210501") {
		referencePropertyName = makeLegacyReferencePropertyName(existing, isSlice, isMap)
	} else {
		referencePropertyName = makeReferencePropertyName(existing, isSlice, isMap)
	}

	originalName := string(existing.PropertyName())

	newProp := astmodel.NewPropertyDefinition(
		idFactory.CreatePropertyName(referencePropertyName, astmodel.Exported),
		idFactory.CreateStringIdentifier(referencePropertyName, astmodel.NotExported),
		existing.PropertyType())
	// TODO: We could pass this information forward some other way?
	// Add tag so that we remember what this field was before
	newProp = newProp.WithTag(astmodel.ARMReferenceTag, originalName)

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
