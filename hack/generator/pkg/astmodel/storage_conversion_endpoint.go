/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"strings"

	"github.com/gobuffalo/flect"
)

// StorageConversionEndpoint represents either a source or a destination target for a storage conversion
// In simple cases these will be fields, but can also represent indexed members of slices and maps.
type StorageConversionEndpoint struct {
	// theType is the Type of the value accessible via this endpoint
	theType Type
	// name is the name of the underlying property, used to generate useful local identifiers
	name string
	// knownLocals is a shared map of locals that have already been created within a given function, to prevent duplicates
	knownLocals *KnownLocalsSet
}

func NewStorageConversionEndpoint(
	theType Type,
	name string,
	knownLocals *KnownLocalsSet) *StorageConversionEndpoint {
	return &StorageConversionEndpoint{
		theType:     theType,
		name:        strings.ToLower(name),
		knownLocals: knownLocals,
	}
}

// Type returns the type of this endpoint
func (endpoint *StorageConversionEndpoint) Type() Type {
	return endpoint.theType
}

// CreateLocal creates an identifier for a local variable using one of the supplied suffixes if
// possible. If all of those suffixes have been used, integer suffixes will be used
// Each call will return a unique identifier
func (endpoint *StorageConversionEndpoint) CreateLocal(suffix ...string) string {
	singular := flect.Singularize(endpoint.name)
	return endpoint.knownLocals.CreateLocal(singular, suffix...)
}

// WithType creates a new endpoint with a different type
func (endpoint *StorageConversionEndpoint) WithType(theType Type) *StorageConversionEndpoint {
	return &StorageConversionEndpoint{
		theType:     theType,
		name:        endpoint.name,
		knownLocals: endpoint.knownLocals,
	}
}
