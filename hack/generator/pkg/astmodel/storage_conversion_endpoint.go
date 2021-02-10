/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"strconv"
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

// CreateSingularLocal creates an identifier for a local variable that represents a single item
// Each call will return a unique identifier
func (endpoint *StorageConversionEndpoint) CreateSingularLocal() string {
	singular := flect.Singularize(endpoint.name)
	return endpoint.knownLocals.createLocal(singular)
}

// CreatePluralLocal creates an identifier for a local variable that represents multiple items
// Each call will return a unique identifier
func (endpoint *StorageConversionEndpoint) CreatePluralLocal(suffix string) string {
	plural := flect.Singularize(endpoint.name)
	return endpoint.knownLocals.createLocal(plural + suffix)
}

// WithType creates a new endpoint with a different type
func (endpoint *StorageConversionEndpoint) WithType(theType Type) *StorageConversionEndpoint {
	return NewStorageConversionEndpoint(
		theType,
		endpoint.name,
		endpoint.knownLocals)
}

type KnownLocalsSet struct {
	names     map[string]struct{}
	idFactory IdentifierFactory
}

// NewKnownLocalsSet returns a new empty set of locals
// idFactory is a reference to an identifier factory for creating valid Go identifiers
func NewKnownLocalsSet(idFactory IdentifierFactory) *KnownLocalsSet {
	return &KnownLocalsSet{
		names:     make(map[string]struct{}),
		idFactory: idFactory,
	}
}

// createLocal creates a new unique Go local variable with the specified suffix
// Has to be deterministic, so we use an incrementing number to make them unique
func (locals KnownLocalsSet) createLocal(nameHint string) string {
	baseName := locals.idFactory.CreateIdentifier(nameHint, NotExported)
	id := baseName
	index := 0
	for {
		_, found := locals.names[id]
		if !found {
			break
		}

		index++
		id = baseName + strconv.Itoa(index)
	}

	locals.names[id] = struct{}{}

	return id
}

// Add allows identifiers that have already been used to be registered, avoiding duplicates
func (locals KnownLocalsSet) Add(local string) {
	name := locals.idFactory.CreateIdentifier(local, NotExported)
	locals.names[name] = struct{}{}
}
