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

// VersionMetaData contains additional information about a specific version of a group
type VersionMetaData struct {
	kinds map[string]*KindMetaData
}

// UnmarshalYAML populates our instance from the YAML.
// Nested objects are handled by a *pair* of nodes (!!), one for the id and one for the content of the object
func (v *VersionMetaData) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	v.kinds = make(map[string]*KindMetaData)
	var lastId string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = c.Value
			continue
		}

		// Handle nested kind metadata
		if c.Kind == yaml.MappingNode {
			var k KindMetaData
			err := c.Decode(k)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			// Store the kind name using a lowercase
			// so we can do case insensitive lookups later
			v.kinds[strings.ToLower(lastId)] = &k
			lastId = "" // hedge against reusing the id
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf("unexpected yaml value %s line %d col %d)", c.Value, c.Line, c.Column)
	}

	return nil
}
