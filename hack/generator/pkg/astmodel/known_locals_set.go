/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"
)

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

// CreateLocal creates a new unique Go local variable with one of the specified suffixes.
// Has to be deterministic, so we use an incrementing number to make them unique if necessary.
func (locals *KnownLocalsSet) CreateLocal(nameHint string, suffixes ...string) string {

	// Ensure we have a safe base case
	if len(suffixes) == 0 {
		suffixes = []string{""}
	}

	// Try to use the suffixes as supplied if we can
	for _, s := range suffixes {
		if strings.HasSuffix(strings.ToLower(nameHint), strings.ToLower(s)) {
			s = ""
		}
		if id, ok := locals.tryCreateLocal(nameHint + " " + s); ok {
			return id
		}
	}

	// Use numeric suffixes if we must
	index := 1
	for {
		for _, s := range suffixes {
			// TODO: I think this isn't right -- we need to trim all the suffixes from the end first
			if strings.HasSuffix(strings.ToLower(nameHint), strings.ToLower(s)) {
				s = ""
			}
			if id, ok := locals.tryCreateLocal(fmt.Sprintf("%s %s%d", nameHint, s, index)); ok {
				return id
			}
		}

		index++
	}
}

// tryCreateLocal attempts to create a new local, returning the new identifier and true if
// successful (local hasn't been used before) or "" and false if not (local already exists)
func (locals *KnownLocalsSet) tryCreateLocal(name string) (string, bool) {
	id := locals.idFactory.CreateIdentifier(name, NotExported)
	if _, found := locals.names[id]; found {
		// Failed to create the name
		return "", false
	}

	locals.names[id] = struct{}{}
	return id, true
}

// Add allows identifiers that have already been used to be registered, avoiding duplicates
func (locals *KnownLocalsSet) Add(local string) {
	name := locals.idFactory.CreateIdentifier(local, NotExported)
	locals.names[name] = struct{}{}
}

func (locals *KnownLocalsSet) HasName(name string) bool {
	_, ok := locals.names[name]
	return ok
}

// Clone clones the KnownLocalsSet
func (locals *KnownLocalsSet) Clone() *KnownLocalsSet {
	names := make(map[string]struct{}, len(locals.names))
	for name := range locals.names {
		names[name] = struct{}{}
	}

	return &KnownLocalsSet{
		idFactory: locals.idFactory,
		names:     names,
	}
}
