/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ReadableConversionEndpointSet is a set of uniquely named readable conversion endpoints
type ReadableConversionEndpointSet map[string]*ReadableConversionEndpoint

// NewReadableConversionEndpointSet returns a new set of readable conversion endpoints
func NewReadableConversionEndpointSet() ReadableConversionEndpointSet {
	return make(ReadableConversionEndpointSet)
}

// CreatePropertyEndpoints will create readable conversion endpoints for any properties found on the passed instance
// type. Existing endpoints won't be overwritten. Returns the count of new endpoints created
func (set ReadableConversionEndpointSet) CreatePropertyEndpoints(sourceType astmodel.Type) int {
	// Add an endpoint for each property we can read
	return set.addForEachProperty(sourceType, func(prop *astmodel.PropertyDefinition) *ReadableConversionEndpoint {
		return NewReadableConversionEndpointReadingProperty(prop.PropertyName(), prop.PropertyType())
	})
}

// CreateValueFunctionEndpoints will create additional readable conversion endpoints for any compatible functions found
// on the passed instance type that don't collide with the names of existing endpoints. Returns the count of new
// endpoints created
func (set ReadableConversionEndpointSet) CreateValueFunctionEndpoints(sourceType astmodel.Type) int {
	// Add more endpoints for any value functions we can read
	return set.addForEachValueFunction(sourceType, func(fn astmodel.ValueFunction) *ReadableConversionEndpoint {
		return NewReadableConversionEndpointReadingValueFunction(fn.Name(), fn.ReturnType())
	})
}

// addForEachProperty iterates over the properties defined by the instance and uses the supplied factory func to
// create an endpoint for each one. Existing endpoints will NOT be overwritten. If a property containing a PropertyBag
// is found, it will be skipped as property bags are special-cased elsewhere.
// Returns the count of new endpoints created.
func (set ReadableConversionEndpointSet) addForEachProperty(
	instance astmodel.Type,
	factory func(definition *astmodel.PropertyDefinition) *ReadableConversionEndpoint,
) int {
	count := 0
	if container, ok := astmodel.AsPropertyContainer(instance); ok {

		// Construct a set containing the properties we can assign
		// This is made up of all regular properties, plus specific kinds of embedded properties

		properties := container.Properties().Copy()
		typesToCopy := astmodel.NewTypeNameSet(astmodel.ObjectMetaType)
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

// addForEachValueFunction iterates over the functions defined by the instance and uses the supplied factory func to
// create an endpoint for each one. Existing endpoints will NOT be overwritten.
// Returns the count of new endpoints created.
func (set ReadableConversionEndpointSet) addForEachValueFunction(
	instance astmodel.Type,
	factory func(definition astmodel.ValueFunction) *ReadableConversionEndpoint,
) int {
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

// Delete removes a specific endpoint from the set
func (set ReadableConversionEndpointSet) Delete(name string) {
	delete(set, name)
}
