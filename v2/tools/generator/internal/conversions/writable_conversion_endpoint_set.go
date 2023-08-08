/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// WritableConversionEndpointSet is a set of uniquely named writable conversion endpoints
type WritableConversionEndpointSet map[string]*WritableConversionEndpoint

// NewWritableConversionEndpointSet returns a new empty set of writable conversion endpoints
func NewWritableConversionEndpointSet() WritableConversionEndpointSet {
	return make(WritableConversionEndpointSet)
}

// CreatePropertyEndpoints will create writable conversion endpoints for any properties found on the passed instance
// type. Existing endpoints won't be overwritten. Returns the count of new endpoints created
func (set WritableConversionEndpointSet) CreatePropertyEndpoints(destinationType astmodel.Type) int {
	// Add an endpoint for each property we can read
	return set.addForEachProperty(destinationType, func(prop *astmodel.PropertyDefinition) *WritableConversionEndpoint {
		return NewWritableConversionEndpointWritingProperty(prop.PropertyName(), prop.PropertyType())
	})
}

// addForEachProperty iterates over the properties defined by the instance and uses the supplied factory func to
// create an endpoint for each one. Existing endpoints will NOT be overwritten. If a property containing a PropertyBag
// is found, it will be skipped as property bags are special cased elsewhere.
// Returns the count of writable endpoints added.
func (set WritableConversionEndpointSet) addForEachProperty(
	instance astmodel.Type,
	factory func(definition *astmodel.PropertyDefinition) *WritableConversionEndpoint,
) int {
	count := 0
	if container, ok := astmodel.AsPropertyContainer(instance); ok {

		// Construct a set containing the properties we can assign
		// This is made up of all regular properties, plus specific kinds of embedded properties

		properties := container.Properties().Copy()
		typesToCopy := astmodel.NewTypeNameSet[astmodel.TypeName](astmodel.ObjectMetaType)
		for _, prop := range container.EmbeddedProperties() {
			name, ok := astmodel.AsTypeName(prop.PropertyType())
			if !ok {
				// We only expect to get embedded type names, but skip any others just in case
				continue
			}

			if !typesToCopy.Contains(name) {
				// Not a type we need to copy
				continue
			}

			properties.Add(prop.WithName(astmodel.PropertyName(name.Name())))
		}

		for _, prop := range properties {
			name := string(prop.PropertyName())
			if _, defined := set[name]; defined {
				// Don't overwrite any existing endpoints
				continue
			}

			if prop.PropertyType().Equals(astmodel.PropertyBagType, astmodel.EqualityOverrides{}) {
				// We don't create endpoints for property bag properties, they're special cased elsewhere
				continue
			}

			endpoint := factory(prop)
			set[name] = endpoint
			count++
		}
	}

	return count
}

// Delete removes a specific endpoint from the set
func (set WritableConversionEndpointSet) Delete(name string) {
	delete(set, name)
}
