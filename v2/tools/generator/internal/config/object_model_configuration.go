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

// ObjectModelConfiguration contains additional information about entire object model, allowing fine-tuning of the
// information loaded from JSON schema and Swagger specs. There is a hierarchy of types involved, as follows:
//
// ╔══════════════════════════╗       ┌────────────────────┐       ┌──────────────────────┐       ┌───────────────────┐       ┌───────────────────────┐
// ║                          ║       │                    │       │                      │       │                   │       │                       │
// ║ ObjectModelConfiguration ║───────│ GroupConfiguration │───────│ VersionConfiguration │───────│ TypeConfiguration │───────│ PropertyConfiguration │
// ║                          ║1  1..n│                    │1  1..n│                      │1  1..n│                   │1  1..n│                       │
// ╚══════════════════════════╝       └────────────────────┘       └──────────────────────┘       └───────────────────┘       └───────────────────────┘
//
type ObjectModelConfiguration struct {
	groups map[string]*GroupConfiguration
}

// TypeRename looks up a rename for the specified type, returning the new name and true if found, or empty string
// and false if not.
func (o *ObjectModelConfiguration) TypeRename(name astmodel.TypeName) (string, bool) {
	localref, ok := name.PackageReference.AsLocalPackage()
	if !ok {
		return "", false
	}

	group := strings.ToLower(localref.Group())
	if g, ok := o.groups[group]; ok {
		return g.TypeRename(name)
	}

	return "", false
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (o *ObjectModelConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	o.groups = make(map[string]*GroupConfiguration)
	var lastId string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = c.Value
			continue
		}

		// Handle nested version metadata
		if c.Kind == yaml.MappingNode && lastId != "" {
			var g GroupConfiguration
			err := c.Decode(&g)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			// store the group name using lowercase,
			// so we can do case-insensitive lookups later
			o.groups[strings.ToLower(lastId)] = &g
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf("unexpected yaml value %s (line %d col %d)", c.Value, c.Line, c.Column)
	}

	return nil
}
