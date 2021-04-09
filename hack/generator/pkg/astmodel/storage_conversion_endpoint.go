/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
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
	return endpoint.knownLocals.createLocal(singular, suffix...)
}

// WithType creates a new endpoint with a different type
func (endpoint *StorageConversionEndpoint) WithType(theType Type) *StorageConversionEndpoint {
	return &StorageConversionEndpoint{
		theType:     theType,
		name:        endpoint.name,
		knownLocals: endpoint.knownLocals,
	}
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

// createLocal creates a new unique Go local variable with one of the specified suffixes.
// Has to be deterministic, so we use an incrementing number to make them unique if necessary.
func (locals KnownLocalsSet) createLocal(nameHint string, suffixes ...string) string {

	// Ensure we have a safe base case
	if len(suffixes) == 0 {
		suffixes = []string{""}
	}

	// Try to use the suffixes as supplied if we can
	for _, s := range suffixes {
		if id, ok := locals.tryCreateLocal(nameHint + s); ok {
			return id
		}
	}

	// Use numeric suffixes if we must
	index := 1
	for {
		for _, s := range suffixes {
			if id, ok := locals.tryCreateLocal(fmt.Sprintf("%s%s%d", nameHint, s, index)); ok {
				return id
			}
		}

		index++
	}
}

// tryCreateLocal attempts to create a new local, returning the new identifier and true if
// successful (local hasn't been used before) or "" and false if not (local already exists)
func (locals KnownLocalsSet) tryCreateLocal(name string) (string, bool) {
	id := locals.idFactory.CreateIdentifier(name, NotExported)
	if _, found := locals.names[id]; found {
		// Failed to create the name
		return "", false
	}

	locals.names[id] = struct{}{}
	return id, true
}

// Add allows identifiers that have already been used to be registered, avoiding duplicates
func (locals KnownLocalsSet) Add(local string) {
	name := locals.idFactory.CreateIdentifier(local, NotExported)
	locals.names[name] = struct{}{}
}
