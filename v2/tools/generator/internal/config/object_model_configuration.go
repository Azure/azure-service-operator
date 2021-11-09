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

// NewObjectModelConfiguration returns a new (empty) ObjectModelConfiguration
func NewObjectModelConfiguration() *ObjectModelConfiguration {
	return &ObjectModelConfiguration{
		groups: make(map[string]*GroupConfiguration),
	}
}

// TypeRename looks up a rename for the specified type, returning the new name and true if found, or empty string
// and false if not.
func (o *ObjectModelConfiguration) TypeRename(name astmodel.TypeName) (string, bool) {
	group, ok := o.findGroup(name)
	if !ok {
		return "", false
	}

	return group.TypeRename(name)
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (o *ObjectModelConfiguration) ARMReference(name astmodel.TypeName, property astmodel.PropertyName) (bool, bool) {
	group, ok := o.findGroup(name)
	if !ok {
		return false, false
	}

	return group.ARMReference(name, property)
}

func (o *ObjectModelConfiguration) Add(group *GroupConfiguration) *ObjectModelConfiguration {
	if o.groups == nil {
		// Initialize the map just-in-time
		o.groups = make(map[string]*GroupConfiguration)
	}
	// store the name name using lowercase,
	// so we can do case-insensitive lookups later
	o.groups[strings.ToLower(group.name)] = group
	return o
}

// findGroup uses the provided TypeName to work out which nested GroupConfiguration should be used
func (o *ObjectModelConfiguration) findGroup(name astmodel.TypeName) (*GroupConfiguration, bool) {
	localRef, ok := name.PackageReference.AsLocalPackage()
	if !ok {
		return nil, false
	}

	group := strings.ToLower(localRef.Group())
	g, ok := o.groups[group]
	if !ok {
		return nil, false
	}

	return g, true
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (o *ObjectModelConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	var lastId string
	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = c.Value
			continue
		}

		// Handle nested name metadata
		if c.Kind == yaml.MappingNode && lastId != "" {
			g := NewGroupConfiguration(lastId)
			err := c.Decode(&g)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			o.Add(g)
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"object model configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}
