/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// VersionConfiguration contains additional information about a specific version of a group and forms part of a
// hierarchy containing information to supplement the schema and swagger sources consumed by the generator.
//
// ┌──────────────────────────┐       ┌────────────────────┐       ╔══════════════════════╗       ┌───────────────────┐       ┌───────────────────────┐
// │                          │       │                    │       ║                      ║       │                   │       │                       │
// │ ObjectModelConfiguration │───────│ GroupConfiguration │───────║ VersionConfiguration ║───────│ TypeConfiguration │───────│ PropertyConfiguration │
// │                          │1  1..n│                    │1  1..n║                      ║1  1..n│                   │1  1..n│                       │
// └──────────────────────────┘       └────────────────────┘       ╚══════════════════════╝       └───────────────────┘       └───────────────────────┘
//
type VersionConfiguration struct {
	name  string
	types map[string]*TypeConfiguration
}

// NewVersionConfiguration returns a new (empty) VersionConfiguration
func NewVersionConfiguration(name string) *VersionConfiguration {
	return &VersionConfiguration{
		name:  name,
		types: make(map[string]*TypeConfiguration),
	}
}

// TypeRename looks up a rename for the specified type, returning the new name and true if found, or empty string
// and false if not.
func (v *VersionConfiguration) TypeRename(name string) (string, bool) {
	tc, ok := v.findType(name)
	if !ok {
		return "", false
	}

	return tc.TypeRename()
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (v *VersionConfiguration) ARMReference(name string, property astmodel.PropertyName) (bool, bool) {
	tc, ok := v.findType(name)
	if !ok {
		return false, false
	}

	return tc.ARMReference(property)
}

// Add includes configuration for the specified type as a part of this version configuration
func (v *VersionConfiguration) Add(tc *TypeConfiguration) *VersionConfiguration {
	// Indexed by lowercase name of the type to allow case insensitive lookups
	v.types[strings.ToLower(tc.name)] = tc
	return v
}

// findtype uses the provided name to work out which nested TypeConfiguration should be used
func (v *VersionConfiguration) findType(name string) (*TypeConfiguration, bool) {
	n := strings.ToLower(name)
	t, ok := v.types[n]
	return t, ok
}

// UnmarshalYAML populates our instance from the YAML.
// Nested objects are handled by a *pair* of nodes (!!), one for the id and one for the content of the object
func (v *VersionConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	v.types = make(map[string]*TypeConfiguration)
	var lastId string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = c.Value
			continue
		}

		// Handle nested kind metadata
		if c.Kind == yaml.MappingNode {
			tc := NewTypeConfiguration(lastId)
			err := c.Decode(&tc)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			v.Add(tc)
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"version configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}
