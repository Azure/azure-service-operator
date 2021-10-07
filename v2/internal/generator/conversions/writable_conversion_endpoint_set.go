/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
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

// CreatePropertyBagMemberEndpoints will create additional property bag item endpoints for any property on the passed instance
// type that doesn't already have one. Returns the count of new endpoints created.
//
// Background: When our destination instance has a property bag, that bag can be used to stash properties from the
// source where there is no matching destination property. We therefore iterate through each property on the *source*
// type and create a WritableConversionEndpoint for each one so the value is stashed in the property bag.
//
func (set WritableConversionEndpointSet) CreatePropertyBagMemberEndpoints(sourceType astmodel.Type) int {
	// Add a property bag member endpoint for each property we don't already support
	return set.addForEachProperty(sourceType, func(prop *astmodel.PropertyDefinition) *WritableConversionEndpoint {
		name := string(prop.PropertyName())
		return NewWritableConversionEndpointWritingPropertyBagMember(name, prop.PropertyType())
	})
}

// addForEachProperty iterates over the properties defined by the instance and uses the supplied factory func to
// create an endpoint for each one. Existing endpoints will NOT be overwritten. If a property containing a PropertyBag
// is found, it will be skipped as property bags are special cased elsewhere.
// Returns the count of writable endpoints added.
func (set WritableConversionEndpointSet) addForEachProperty(
	instance astmodel.Type,
	factory func(definition *astmodel.PropertyDefinition) *WritableConversionEndpoint) int {
	count := 0
	if container, ok := astmodel.AsPropertyContainer(instance); ok {
		for _, prop := range container.Properties() {
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
