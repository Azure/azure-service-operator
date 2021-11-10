/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
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
func (vc *VersionConfiguration) TypeRename(name string) (string, bool) {
	tc, ok := vc.findType(name)
	if !ok {
		return "", false
	}

	return tc.TypeRename()
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (vc *VersionConfiguration) ARMReference(name string, property astmodel.PropertyName) (bool, bool) {
	tc, ok := vc.findType(name)
	if !ok {
		return false, false
	}

	return tc.ARMReference(property)
}

// FindUnusedARMReferences returns a slice listing any unused ARMReference configuration
func (vc *VersionConfiguration) FindUnusedARMReferences() []string {
	var result []string
	for _, tc := range vc.types {
		for _, s := range tc.FindUnusedARMReferences() {
			msg := fmt.Sprintf("version %s %s", vc.name, s)
			result = append(result, msg)
		}
	}

	return result
}

// Add includes configuration for the specified type as a part of this version configuration
func (vc *VersionConfiguration) Add(tc *TypeConfiguration) *VersionConfiguration {
	// Indexed by lowercase name of the type to allow case insensitive lookups
	vc.types[strings.ToLower(tc.name)] = tc
	return vc
}

// findtype uses the provided name to work out which nested TypeConfiguration should be used
func (vc *VersionConfiguration) findType(name string) (*TypeConfiguration, bool) {
	n := strings.ToLower(name)
	t, ok := vc.types[n]
	return t, ok
}

// UnmarshalYAML populates our instance from the YAML.
// Nested objects are handled by a *pair* of nodes (!!), one for the id and one for the content of the object
func (vc *VersionConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	vc.types = make(map[string]*TypeConfiguration)
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

			vc.Add(tc)
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"version configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}
