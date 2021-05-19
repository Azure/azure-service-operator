package storage

import (
	"fmt"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/pkg/errors"
)

// PropertyConverter is used to convert the properties of object types as required for storage variants
type PropertyConverter struct {
	visitor             astmodel.TypeVisitor // Visitor used to achieve the required modification
	types               astmodel.Types       // All the types for this group
	propertyConversions []propertyConversion // Conversion rules to use for properties when creating storage variants
}

// NewPropertyConverter creates a new property converter for modifying object properties
func NewPropertyConverter(types astmodel.Types) *PropertyConverter {
	result := &PropertyConverter{
		types: types,
	}

	result.propertyConversions = []propertyConversion{
		result.preserveKubernetesResourceStorageProperties,
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
func (p PropertyConverter) ConvertProperty(property *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error) {

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

	return nil, fmt.Errorf("failed to find a conversion for property %v", property.PropertyName())
}

// stripAllValidations removes all validations
func (p PropertyConverter) stripAllValidations(
	this *astmodel.TypeVisitor, v *astmodel.ValidatedType, ctx interface{}) (astmodel.Type, error) {
	// strip all type validations from storage property types
	// act as if they do not exist
	return this.Visit(v.ElementType(), ctx)
}

// useBaseTypeForEnumerations replaces an enumeration with its underlying base type
func (p PropertyConverter) useBaseTypeForEnumerations(
	tv *astmodel.TypeVisitor, et *astmodel.EnumType, ctx interface{}) (astmodel.Type, error) {
	return tv.Visit(et.BaseType(), ctx)
}

// shortCircuitNamesOfSimpleTypes redirects TypeNames that reference resources or objects into our
// storage namespace, and replaces TypeNames that point to simple types (enumerations or
// primitives) with their underlying type.
func (p PropertyConverter) shortCircuitNamesOfSimpleTypes(
	tv *astmodel.TypeVisitor, tn astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {

	actualType, err := p.types.FullyResolve(tn)
	if err != nil {
		// Can't resolve to underlying type, give up
		return nil, err
	}

	_, isObject := astmodel.AsObjectType(actualType)
	_, isResource := astmodel.AsResourceType(actualType)

	if isObject || isResource {
		// We have an object or a resource, redirect to our storage namespace if we can
		if storageName, ok := p.tryConvertToStorageNamespace(tn); ok {
			return storageName, nil
		}

		// Otherwise just keep the name
		return tn, nil
	}

	// Replace the name with the underlying type
	return tv.Visit(actualType, ctx)
}

func (_ PropertyConverter) tryConvertToStorageNamespace(name astmodel.TypeName) (astmodel.TypeName, bool) {
	// Map the type name into our storage namespace
	localRef, ok := name.PackageReference.AsLocalPackage()
	if !ok {
		return astmodel.TypeName{}, false
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
func (p PropertyConverter) preserveKubernetesResourceStorageProperties(
	prop *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error) {

	if astmodel.IsKubernetesResourceProperty(prop.PropertyName()) {
		// Keep these unchanged
		return prop, nil
	}

	// Not a kubernetes type, defer to another conversion
	return nil, nil
}

func (p PropertyConverter) defaultPropertyConversion(
	property *astmodel.PropertyDefinition) (*astmodel.PropertyDefinition, error) {

	propertyType, err := p.visitor.Visit(property.PropertyType(), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "converting property %q", property.PropertyName())
	}

	newProperty := property.WithType(propertyType).
		MakeOptional().
		WithDescription("")

	return newProperty, nil
}
