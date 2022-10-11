/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/flect"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

type KnownLocalsSet struct {
	names     set.Set[string]
	idFactory IdentifierFactory
}

// NewKnownLocalsSet returns a new empty set of locals
// idFactory is a reference to an identifier factory for creating valid Go identifiers
func NewKnownLocalsSet(idFactory IdentifierFactory) *KnownLocalsSet {
	return &KnownLocalsSet{
		names:     set.Make[string](),
		idFactory: idFactory,
	}
}

// CreateSingularLocal creates a new unique Go local variable for a single value with one of the specified suffixes.
func (locals *KnownLocalsSet) CreateSingularLocal(nameHint string, suffixes ...string) string {
	hint := flect.Singularize(nameHint)
	return locals.CreateLocal(hint, suffixes...)
}

// CreatePluralLocal creates a new unique Go local variable for multiple values with one of the specified suffixes.
func (locals *KnownLocalsSet) CreatePluralLocal(nameHint string, suffixes ...string) string {
	hint := flect.Pluralize(nameHint)
	return locals.CreateLocal(hint, suffixes...)
}

// CreateLocal creates a new unique Go local variable with one of the specified suffixes.
// Has to be deterministic, so we use an incrementing number to make them unique if necessary.
func (locals *KnownLocalsSet) CreateLocal(nameHint string, suffixes ...string) string {
	// Ensure we have a safe base case
	if len(suffixes) == 0 {
		suffixes = []string{""}
	}

	// Remove underscores from nameHint - we get them from type names, but don't want them in locals
	nameHint = strings.ReplaceAll(nameHint, "_", "")

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

// TryCreateLocal returns true if the specified local was successfully created, false if it already existed
func (locals *KnownLocalsSet) TryCreateLocal(local string) bool {
	if locals.HasName(local) {
		return false
	}

	locals.Add(local)
	return true
}

// tryCreateLocal attempts to create a new local, returning the new identifier and true if
// successful (local hasn't been used before) or "" and false if not (local already exists)
func (locals *KnownLocalsSet) tryCreateLocal(name string) (string, bool) {
	id := locals.idFactory.CreateLocal(name)
	if locals.names.Contains(id) {
		// Failed to create the name
		return "", false
	}

	locals.names.Add(id)
	return id, true
}

// Add allows one or more identifiers that have already been used to be registered, avoiding duplicates
func (locals *KnownLocalsSet) Add(identifiers ...string) {
	for _, id := range identifiers {
		name := locals.idFactory.CreateLocal(id)
		locals.names.Add(name)
	}
}

// HasName returns true if the specified name exists in the set, false otherwise
func (locals *KnownLocalsSet) HasName(name string) bool {
	return locals.names.Contains(name)
}

// Clone clones the KnownLocalsSet
func (locals *KnownLocalsSet) Clone() *KnownLocalsSet {
	return &KnownLocalsSet{
		idFactory: locals.idFactory,
		names:     locals.names.Copy(),
	}
}
