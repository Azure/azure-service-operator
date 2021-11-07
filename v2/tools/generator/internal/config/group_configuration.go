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

// GroupConfiguration contains additional information about an entire group and forms the top of a hierarchy containing
// information to supplement the schema and swagger sources consumed by the generator.
//
// ┌──────────────────────────┐       ╔════════════════════╗       ┌──────────────────────┐       ┌───────────────────┐       ┌───────────────────────┐
// │                          │       ║                    ║       │                      │       │                   │       │                       │
// │ ObjectModelConfiguration │───────║ GroupConfiguration ║───────│ VersionConfiguration │───────│ TypeConfiguration │───────│ PropertyConfiguration │
// │                          │1  1..n║                    ║1  1..n│                      │1  1..n│                   │1  1..n│                       │
// └──────────────────────────┘       ╚════════════════════╝       └──────────────────────┘       └───────────────────┘       └───────────────────────┘
//
type GroupConfiguration struct {
	versions map[string]*VersionConfiguration
}

// TypeRename looks up a rename for the specified type, returning the new name and true if found, or empty string
// and false if not.
func (g *GroupConfiguration) TypeRename(name astmodel.TypeName) (string, bool) {
	version, ok := g.findVersion(name)
	if !ok {
		return "", false
	}

	return version.TypeRename(name.Name())
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (g *GroupConfiguration) ARMReference(name astmodel.TypeName, property astmodel.PropertyName) (bool, bool) {
	version, ok := g.findVersion(name)
	if !ok {
		return false, false
	}

	return version.ARMReference(name.Name(), property)
}

// findVersion uses the provided TypeName to work out which nested VersionConfiguration should be used
func (g *GroupConfiguration) findVersion(name astmodel.TypeName) (*VersionConfiguration, bool) {
	v := strings.ToLower(name.PackageReference.PackageName())
	version, ok := g.versions[v]
	return version, ok
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (g *GroupConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	g.versions = make(map[string]*VersionConfiguration)
	var lastId string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = c.Value
			continue
		}

		// Handle nested version metadata
		if c.Kind == yaml.MappingNode && lastId != "" {
			var v VersionConfiguration
			err := c.Decode(&v)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			// store the version id using lowercase,
			// so we can do case-insensitive lookups later
			g.versions[strings.ToLower(lastId)] = &v
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf("unexpected yaml value %s (line %d col %d)", c.Value, c.Line, c.Column)
	}

	return nil
}
