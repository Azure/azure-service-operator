/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// PropertyConverter is used to convert the properties of object definitions as required for storage variants
type PropertyConverter struct {
	// visitor is used to apply the modification
	visitor astmodel.TypeVisitor
	// definitions contains all the definitions for this group
	definitions astmodel.TypeDefinitionSet
	// propertyConversions is an ordered list of all our conversion rules for creating storage variants
	propertyConversions []propertyConversion
}

// NewPropertyConverter creates a new property converter for modifying object properties
func NewPropertyConverter(definitions astmodel.TypeDefinitionSet) *PropertyConverter {
	result := &PropertyConverter{
		definitions: definitions,
	}

	result.propertyConversions = []propertyConversion{
		result.preserveKubernetesResourceStorageProperties,
		result.preserveResourceReferenceProperties,
		result.defaultPropertyConversion,
	}

	result.visitor = astmodel.TypeVisitorBuilder{
		VisitEnumType:      result.useBaseTypeForEnumerations,
		VisitValidatedType: result.stripAllValidations,
		VisitTypeName:      result.shortCircuitNamesOfSimpleTypes,
	}.Build()

	return result
}

// ConvertProperty applies our conversion rules to a specific property
func (p *PropertyConverter) ConvertProperty(property *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error) {
	for _, conv := range p.propertyConversions {
		prop, err := conv(property)
		if err != nil {
			// Something went wrong, return the error
			return nil, err
		}
		if prop != nil {
			// We have the conversion we need, return it promptly
			return prop, nil
		}
	}

	// No conversion found
	return nil, fmt.Errorf(
		"failed to find a conversion for property %s (%s)",
		property.PropertyName(),
		astmodel.DebugDescription(property.PropertyType()))
}

// stripAllValidations removes all validations
func (p *PropertyConverter) stripAllValidations(
	this *astmodel.TypeVisitor, v *astmodel.ValidatedType, ctx interface{}) (astmodel.Type, error) {
	// strip all type validations from storage properties
	// act as if they do not exist
	return this.Visit(v.ElementType(), ctx)
}

// useBaseTypeForEnumerations replaces an enumeration with its underlying base type
func (p *PropertyConverter) useBaseTypeForEnumerations(
	tv *astmodel.TypeVisitor, et *astmodel.EnumType, ctx interface{}) (astmodel.Type, error) {
	return tv.Visit(et.BaseType(), ctx)
}

// shortCircuitNamesOfSimpleTypes redirects or replaces TypeNames
//
//	o  If a TypeName points into an API package, it is redirected into the appropriate storage package
//	o  If a TypeName references an enumeration, it is replaced with the underlying type of the enumeration as our
//	   storage definitions don't use enumerations, they use primitive definitions
//	o  If a TypeName references an alias for a primitive type (these are used to specify validations), it is replaced
//	   with the primitive type
func (p *PropertyConverter) shortCircuitNamesOfSimpleTypes(
	tv *astmodel.TypeVisitor, tn astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {

	// for nonlocal packages, preserve the name as is
	if astmodel.IsExternalPackageReference(tn.PackageReference()) {
		return tn, nil
	}

	actualType, err := p.definitions.FullyResolve(tn)
	if err != nil {
		// Can't resolve to underlying type, give up
		return nil, err
	}

	_, isObject := astmodel.AsObjectType(actualType)
	_, isResource := astmodel.AsResourceType(actualType)

	if isObject || isResource {
		// We have an object or a resource, redirect to our storage package if we can
		if storageName, ok := p.tryConvertToStoragePackage(tn); ok {
			return storageName, nil
		}

		// Otherwise just keep the name
		return tn, nil
	}

	// Replace the name with the underlying type
	return tv.Visit(actualType, ctx)
}

func (_ *PropertyConverter) tryConvertToStoragePackage(name astmodel.TypeName) (astmodel.TypeName, bool) {
	// Map the type name into our storage package
	localRef, ok := name.PackageReference().(astmodel.LocalPackageReference)
	if !ok {
		return nil, false
	}

	storageRef := astmodel.MakeStoragePackageReference(localRef)
	visitedName := astmodel.MakeTypeName(storageRef, name.Name())
	return visitedName, true
}

// A property conversion accepts a property definition and optionally applies a conversion to make
// the property suitable for use on a storage type. Conversions return nil if they decline to
// convert, deferring the conversion to another.
type propertyConversion = func(property *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error)

// preserveKubernetesResourceStorageProperties preserves properties required by the
// KubernetesResource interface as they're always required exactly as declared
func (p *PropertyConverter) preserveKubernetesResourceStorageProperties(
	prop *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error) {

	if astmodel.IsKubernetesResourceProperty(prop.PropertyName()) {
		// Keep these unchanged
		return prop, nil
	}

	// Not a kubernetes type, defer to another conversion
	return nil, nil
}

// preserveResourceReferenceProperties preserves properties required by the
// KubernetesResource interface as they're always required exactly as declared
func (p *PropertyConverter) preserveResourceReferenceProperties(
	prop *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error) {

	propertyType := prop.PropertyType()
	if opt, ok := astmodel.AsOptionalType(propertyType); ok {
		if astmodel.TypeEquals(opt.Element(), astmodel.ResourceReferenceType) {
			// Keep these unchanged
			return prop, nil
		}
	}

	if astmodel.TypeEquals(propertyType, astmodel.ResourceReferenceType) {
		// Keep these unchanged
		return prop, nil
	}

	// Not a resource reference property, defer to another conversion
	return nil, nil
}

func (p *PropertyConverter) defaultPropertyConversion(
	property *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error) {

	propertyType, err := p.visitor.Visit(property.PropertyType(), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "converting property %q", property.PropertyName())
	}

	newProperty := property.WithType(propertyType).
		MakeOptional().
		MakeTypeOptional().
		WithDescription("")

	return newProperty, nil
}
