/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// TypeConfiguration contains additional information about a specific kind of resource within a version of a group and forms
// part of a heirarchy containing information to supplement the schema and swagger sources consumed by the generator.
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

func NewTypeConfiguration(renamedTo string) *TypeConfiguration {
	return &TypeConfiguration{
		renamedTo:  renamedTo,
		properties: make(map[string]*PropertyConfiguration),
	}
}

// LookupTypeRename returns a new name (and true) if one is configured for this type, or empty string and false if not.
func (t *TypeConfiguration) LookupTypeRename() (string, bool) {
	if t.renamedTo != "" {
		t.usedRenamedTo = true
		return t.renamedTo, true
	}

	return "", false
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (k *TypeConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	k.properties = make(map[string]*PropertyConfiguration)
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
			err := c.Decode(p)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			// Store the property id using lowercase
			// so we can do case insensitive lookups later
			k.properties[strings.ToLower(lastId)] = &p
			continue
		}

		if strings.ToLower(lastId) == "$renamedto" && c.Kind == yaml.ScalarNode {
			k.renamedTo = c.Value
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf("unexpected yaml value %s line %d col %d)", c.Value, c.Line, c.Column)
	}

	return nil
}
