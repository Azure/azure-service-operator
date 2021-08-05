/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// WritableConversionEndpointSet is a set of uniquely named writable conversion endpoints
type WritableConversionEndpointSet map[string]WritableConversionEndpoint

// NewWritableConversionEndpointSet returns a new empty set of writable conversion endpoints
func NewWritableConversionEndpointSet() WritableConversionEndpointSet {
	return make(WritableConversionEndpointSet)
}

// CreatePropertyEndpoints will create writable conversion endpoints for any properties found on the passed instance
// type. Existing endpoints won't be overwritten. Returns the count of new endpoints created
func (set WritableConversionEndpointSet) CreatePropertyEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) int {
	// Add an endpoint for each property we can read
	return set.addForEachProperty(instance, func(prop *astmodel.PropertyDefinition) WritableConversionEndpoint {
		return MakeWritableConversionEndpointWritingProperty(prop.PropertyName(), prop.PropertyType(), knownLocals)
	})
}

// CreateBagItemEndpoints will create additional property bag item endpoints for any property on the passed instance
// type that doesn't already have one. Returns the count of new endpoints created.
func (set WritableConversionEndpointSet) CreateBagItemEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) int {
	// Add a property bag item endpoint for each property we don't already support
	return set.addForEachProperty(instance, func(prop *astmodel.PropertyDefinition) WritableConversionEndpoint {
		name := string(prop.PropertyName())
		return MakeWritableConversionEndpointWritingBagItem(name, prop.PropertyType(), knownLocals)
	})
}

// addForEachProperty iterates over the properties defined by the instance and uses the supplied factory func to
// create an endpoint for each one. Existing endpoints will NOT be overwritten. If a property containing a PropertyBag
// is found, it will be skipped as property bags are special cased elsewhere.
// Returns the count of writable endpoints added.
func (set WritableConversionEndpointSet) addForEachProperty(
	instance astmodel.Type,
	factory func(definition *astmodel.PropertyDefinition) WritableConversionEndpoint) int {
	count := 0
	if container, ok := astmodel.AsPropertyContainer(instance); ok {
		for _, prop := range container.Properties() {
			name := string(prop.PropertyName())
			if _, defined := set[name]; defined {
				// Don't overwrite any existing endpoints
				continue
			}

			if prop.PropertyType().Equals(astmodel.PropertyBagType) {
				// We don't create endpoints for property bag properties, they're special cased elsewhere
				continue
			}

			endpoint := factory(prop)
			set[name] = endpoint
			count ++
		}
	}

	return count
}
