/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// ReadableConversionEndpointSet is a set of uniquely named readable conversion endpoints
type ReadableConversionEndpointSet map[string]ReadableConversionEndpoint

// NewReadableConversionEndpointSet returns a new set of readable conversion endpoints
func NewReadableConversionEndpointSet() ReadableConversionEndpointSet {
	return make(ReadableConversionEndpointSet)
}

// CreatePropertyEndpoints will create readable conversion endpoints for any properties found on the passed instance
// type. Existing endpoints won't be overwritten. Returns the count of new endpoints created
func (set ReadableConversionEndpointSet) CreatePropertyEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) int {
	// Add an endpoint for each property we can read
	return set.addForEachProperty(instance, func(prop *astmodel.PropertyDefinition) ReadableConversionEndpoint {
		return MakeReadableConversionEndpointReadingProperty(prop.PropertyName(), prop.PropertyType(), knownLocals)
	})
}

// CreateValueFunctionEndpoints will create additional readable conversion endpoints for any compatible functions found
// on the passed instance type that don't collide with the names of existing endpoints. Returns the count of new
// endpoints created
func (set ReadableConversionEndpointSet) CreateValueFunctionEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) int {
	// Add more endpoints for any value functions we can read
	return set.addForEachValueFunction(instance, func(fn astmodel.ValueFunction) ReadableConversionEndpoint {
		return MakeReadableConversionEndpointReadingValueFunction(fn.Name(), fn.ReturnType(), knownLocals)
	})
}

// CreateBagItemEndpoints will create additional property bag item endpoints for any property on the passed instance
// type that doesn't already have one. Returns the count of new endpoints created.
func (set ReadableConversionEndpointSet) CreateBagItemEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) int {
	// Add a property bag item endpoint for each property we don't already support
	return set.addForEachProperty(instance, func(prop *astmodel.PropertyDefinition) ReadableConversionEndpoint {
		name := string(prop.PropertyName())
		return MakeReadableConversionEndpointReadingBagItem(name, prop.PropertyType(), knownLocals)
	})
}

// addForEachProperty iterates over the properties defined by the instance and uses the supplied factory func to
// create an endpoint for each one. Existing endpoints will NOT be overwritten. If a property containing a PropertyBag
// is found, it will be skipped as property bags are special-cased elsewhere.
// Returns the count of new endpoints created.
func (set ReadableConversionEndpointSet) addForEachProperty(
	instance astmodel.Type,
	factory func(definition *astmodel.PropertyDefinition) ReadableConversionEndpoint) int {
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
			count++
		}
	}

	return count
}

// addForEachValueFunction iterates over the functions defined by the instance and uses the supplied factory func to
// create an endpoint for each one. Existing endpoints will NOT be overwritten.
// Returns the count of new endpoints created.
func (set ReadableConversionEndpointSet) addForEachValueFunction(
	instance astmodel.Type,
	factory func(definition astmodel.ValueFunction) ReadableConversionEndpoint) int {
	count := 0
	if container, ok := astmodel.AsFunctionContainer(instance); ok {
		for _, fn := range container.Functions() {
			name := fn.Name()
			if _, defined := set[name]; defined {
				// Don't overwrite any existing endpoints
				continue
			}

			valueFn, ok := fn.(astmodel.ValueFunction)
			if ok {
				endpoint := factory(valueFn)
				set[name] = endpoint
				count++
			}
		}
	}

	return count
}
