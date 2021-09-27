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
	names     knownNameSet
	idFactory IdentifierFactory
}

// NewKnownLocalsSet returns a new empty set of locals
// idFactory is a reference to an identifier factory for creating valid Go identifiers
func NewKnownLocalsSet(idFactory IdentifierFactory) *KnownLocalsSet {
	return &KnownLocalsSet{
		names:     make(knownNameSet),
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

// Add allows one or more identifiers that have already been used to be registered, avoiding duplicates
func (locals *KnownLocalsSet) Add(identifiers ...string) {
	for _, id := range identifiers {
		name := locals.idFactory.CreateIdentifier(id, NotExported)
		locals.names[name] = struct{}{}
	}
}

// HasName returns true if the specified name exists in the set, false otherwise
func (locals *KnownLocalsSet) HasName(name string) bool {
	_, ok := locals.names[name]
	return ok
}

// Clone clones the KnownLocalsSet
func (locals *KnownLocalsSet) Clone() *KnownLocalsSet {
	return &KnownLocalsSet{
		idFactory: locals.idFactory,
		names:     locals.names.clone(),
	}
}

// Checkpoint returns an opaque checkpoint for the collection allowing it to be reset later on
func (locals *KnownLocalsSet) Checkpoint() KnownNameSet {
	return locals.names.clone()
}

// Restore returns the state of the set to what it was when checkpointed
func (locals *KnownLocalsSet) Restore(set KnownNameSet) {
	locals.names = set.clone()
}

// KnownNameSet is an interface with no public members used for checkpointing
type KnownNameSet interface {
	clone() knownNameSet
}

// knownNameSet is the internal type of a set of names
type knownNameSet map[string]struct{}

// clone returns a new, independent, copy of the knownNameSet
func (s knownNameSet) clone() knownNameSet {
	result := make(knownNameSet, len(s))
	for name := range s {
		result[name] = struct{}{}
	}

	return result
}
