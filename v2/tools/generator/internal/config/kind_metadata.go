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

// KindMetaData contains additional information about a specific kind of resource within a version of a group and forms
// part of a heirarchy containing information to supplement the schema and swagger sources consumed by the generator.
//
// ┌──────────────────┐       ┌──────────────────┐       ╔══════════════════╗       ┌──────────────────┐
// │                  │       │                  │       ║                  ║       │                  │
// │  GroupMetaData   │───────│ VersionMetaData  │───────║   KindMetadata   ║───────│ PropertyMetaData │
// │                  │1  1..n│                  │1  1..n║                  ║1  1..n│                  │
// └──────────────────┘       └──────────────────┘       ╚══════════════════╝       └──────────────────┘
//
type KindMetaData struct {
	renamedTo  string
	properties map[string]*PropertyMetaData
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (k *KindMetaData) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	k.properties = make(map[string]*PropertyMetaData)
	var lastId string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = c.Value
			continue
		}

		// Handle nested property metadata
		if c.Kind == yaml.MappingNode {
			var p PropertyMetaData
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
