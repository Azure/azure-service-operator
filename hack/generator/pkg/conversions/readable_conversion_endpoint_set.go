package conversions

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// ReadableConversionEndpointSet is a set of uniquely named readable conversion endpoints
type ReadableConversionEndpointSet map[string]ReadableConversionEndpoint

// NewReadableConversionEndpointSet is a set of all the readable endpoints found on a type
func NewReadableConversionEndpointSet() ReadableConversionEndpointSet {
	return make(ReadableConversionEndpointSet)
}

func (set ReadableConversionEndpointSet) AddPropertyEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) {
	// Add an endpoint for each property we can read
	set.addForEachProperty(instance, func(prop *astmodel.PropertyDefinition) ReadableConversionEndpoint {
		return MakeReadableConversionEndpointReadingProperty(prop.PropertyName(), prop.PropertyType(), knownLocals)
	})
}

func (set ReadableConversionEndpointSet) AddValueFunctionEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) {
	// Add more endpoints for any value functions we can read
	set.addForEachValueFunction(instance, func(fn astmodel.ValueFunction) ReadableConversionEndpoint {
		return MakeReadableConversionEndpointReadingValueFunction(fn.Name(), fn.ReturnType(), knownLocals)
	})
}

func (set ReadableConversionEndpointSet) AddBagItemEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) {
	// Add a property bag item endpoint for each property we don't already support
	set.addForEachProperty(instance, func(prop *astmodel.PropertyDefinition) ReadableConversionEndpoint {
		name := string(prop.PropertyName())
		return MakeReadableConversionEndpointReadingBagItem(name, prop.PropertyType(), knownLocals)
	})
}

// addForEachProperty iterates over the properties defined by the instance and uses the supplied factory func to
// create an endpoint for each one. Existing endpoints will NOT be overwritten. If a property containing a PropertyBag
// is found, it will be skipped as property bags are special cased elsewhere.
func (set ReadableConversionEndpointSet) addForEachProperty(
	instance astmodel.Type,
	factory func(definition *astmodel.PropertyDefinition) ReadableConversionEndpoint) {
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
		}
	}
}

// addForEachValueFunction iterates over the functions defined by the instance and uses the supplied factory func to
// create an endpoint for each one. Existing endpoints will NOT be overwritten.
func (set ReadableConversionEndpointSet) addForEachValueFunction(
	instance astmodel.Type,
	factory func(definition astmodel.ValueFunction) ReadableConversionEndpoint) {
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
			}
		}
	}
}
