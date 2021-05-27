/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/gobuffalo/flect"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// PropertyConversionEndpoint represents either a source or a destination target for a storage conversion
// In simple cases these will be fields, but can also represent indexed members of slices and maps.
type PropertyConversionEndpoint struct {
	// theType is the Type of the value accessible via this endpoint
	theType astmodel.Type
	// name is the name of the underlying property, used to generate useful local identifiers
	name string
	// knownLocals is a shared map of locals that have already been created within a given function, to prevent duplicates
	knownLocals *astmodel.KnownLocalsSet
}

func NewStorageConversionEndpoint(
	theType astmodel.Type,
	name string,
	knownLocals *astmodel.KnownLocalsSet) *PropertyConversionEndpoint {
	return &PropertyConversionEndpoint{
		theType:     theType,
		name:        name,
		knownLocals: knownLocals,
	}
}

func (endpoint *PropertyConversionEndpoint) Name() string {
	return endpoint.name
}

// Type returns the type of this endpoint
//TODO: Remove ?
func (endpoint *PropertyConversionEndpoint) Type() astmodel.Type {
	return endpoint.theType
}

// CreateLocal creates an identifier for a local variable using one of the supplied suffixes if
// possible. If all of those suffixes have been used, integer suffixes will be used
// Each call will return a unique identifier
func (endpoint *PropertyConversionEndpoint) CreateLocal(suffix ...string) string {
	singular := flect.Singularize(endpoint.name)
	return endpoint.knownLocals.CreateLocal(singular, suffix...)
}

// WithType creates a new endpoint with a different type
func (endpoint *PropertyConversionEndpoint) WithType(theType astmodel.Type) *PropertyConversionEndpoint {
	return &PropertyConversionEndpoint{
		theType:     theType,
		name:        endpoint.name,
		knownLocals: endpoint.knownLocals,
	}
}
