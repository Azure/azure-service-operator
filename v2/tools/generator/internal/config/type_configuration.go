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

// TypeConfiguration contains additional information about a specific kind of resource within a version of a group and forms
// part of a hierarchy containing information to supplement the schema and swagger sources consumed by the generator.
//
// ┌──────────────────────────┐       ┌────────────────────┐       ┌──────────────────────┐       ╔═══════════════════╗       ┌───────────────────────┐
// │                          │       │                    │       │                      │       ║                   ║       │                       │
// │ ObjectModelConfiguration │───────│ GroupConfiguration │───────│ VersionConfiguration │───────║ TypeConfiguration ║───────│ PropertyConfiguration │
// │                          │1  1..n│                    │1  1..n│                      │1  1..n║                   ║1  1..n│                       │
// └──────────────────────────┘       └────────────────────┘       └──────────────────────┘       ╚═══════════════════╝       └───────────────────────┘
//
type TypeConfiguration struct {
	renamedTo     string
	usedRenamedTo bool
	properties    map[string]*PropertyConfiguration
}

func NewTypeConfiguration() *TypeConfiguration {
	return &TypeConfiguration{
		properties: make(map[string]*PropertyConfiguration),
	}
}

// TypeRename returns a new name (and true) if one is configured for this type, or empty string and false if not.
func (tc *TypeConfiguration) TypeRename() (string, bool) {
	if tc.renamedTo != "" {
		tc.usedRenamedTo = true
		return tc.renamedTo, true
	}

	return "", false
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (tc *TypeConfiguration) ARMReference(property astmodel.PropertyName) (bool, bool) {
	pc, ok := tc.findProperty(property)
	if !ok {
		return false, false
	}

	return pc.ARMReference()
}

// findProperty uses the provided property name to work out which nested PropertyConfiguration should be used
func (tc *TypeConfiguration) findProperty(property astmodel.PropertyName) (*PropertyConfiguration, bool) {
	p := strings.ToLower(string(property))
	pc, ok := tc.properties[p]
	return pc, ok
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (tc *TypeConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	tc.properties = make(map[string]*PropertyConfiguration)
	var lastId string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = c.Value
			continue
		}

		// Handle nested property metadata
		if c.Kind == yaml.MappingNode {
			var p PropertyConfiguration
			err := c.Decode(&p)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			// Store the property id using lowercase,
			// so we can do case-insensitive lookups later
			tc.properties[strings.ToLower(lastId)] = &p
			continue
		}

		if strings.ToLower(lastId) == "$renamedto" && c.Kind == yaml.ScalarNode {
			tc.renamedTo = c.Value
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf("type configuration, unexpected yaml value %s (line %d col %d)", c.Value, c.Line, c.Column)
	}

	return nil
}
