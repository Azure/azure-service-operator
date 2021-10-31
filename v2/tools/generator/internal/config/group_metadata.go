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

// GroupMetaData contains additional information about an entire group
type GroupMetaData struct {
	versions map[string]*VersionMetaData
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (g *GroupMetaData) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	g.versions = make(map[string]*VersionMetaData)
	var lastId string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = c.Value
			continue
		}

		// Handle nested version metadata
		if c.Kind == yaml.MappingNode && lastId != "" {
			var vmd VersionMetaData
			err := c.Decode(vmd)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			// store the version id using lowercase
			// so we can do case insensitive lookups later
			g.versions[strings.ToLower(lastId)] = &v
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf("unexpected yaml value %s line %d col %d)", c.Value, c.Line, c.Column)
	}

	return nil
}
