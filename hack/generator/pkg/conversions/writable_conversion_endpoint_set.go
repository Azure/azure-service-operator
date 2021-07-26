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

func (set WritableConversionEndpointSet) CreatePropertyEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) {
	// Add an endpoint for each property we can read
	set.addForEachProperty(instance, func(prop *astmodel.PropertyDefinition) WritableConversionEndpoint {
		return MakeWritableConversionEndpointWritingProperty(prop.PropertyName(), prop.PropertyType(), knownLocals)
	})
}

func (set WritableConversionEndpointSet) AddBagItemEndpoints(
	instance astmodel.Type,
	knownLocals *astmodel.KnownLocalsSet) {
	// Add a property bag item endpoint for each property we don't already support
	set.addForEachProperty(instance, func(prop *astmodel.PropertyDefinition) WritableConversionEndpoint {
		name := string(prop.PropertyName())
		return MakeWritableConversionEndpointWritingBagItem(name, prop.PropertyType(), knownLocals)
	})
}

// addForEachProperty iterates over the properties defined by the instance and uses the supplied factory func to
// create an endpoint for each one. Existing endpoints will NOT be overwritten. If a property containing a PropertyBag
// is found, it will be skipped as property bags are special cased elsewhere.
func (set WritableConversionEndpointSet) addForEachProperty(
	instance astmodel.Type,
	factory func(definition *astmodel.PropertyDefinition) WritableConversionEndpoint) {
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
