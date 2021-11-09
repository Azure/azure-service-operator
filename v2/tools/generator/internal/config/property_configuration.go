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

// PropertyConfiguration contains additional information about a specific property and forms part of a hierarchy
// containing information to supplement the schema and swagger sources consumed by the generator.
//
// ┌──────────────────────────┐       ┌────────────────────┐       ┌──────────────────────┐       ┌───────────────────┐       ╔═══════════════════════╗
// │                          │       │                    │       │                      │       │                   │       ║                       ║
// │ ObjectModelConfiguration │───────│ GroupConfiguration │───────│ VersionConfiguration │───────│ TypeConfiguration │───────║ PropertyConfiguration ║
// │                          │1  1..n│                    │1  1..n│                      │1  1..n│                   │1  1..n║                       ║
// └──────────────────────────┘       └────────────────────┘       └──────────────────────┘       └───────────────────┘       ╚═══════════════════════╝
//
type PropertyConfiguration struct {
	name         string
	renamedTo    *string
	armReference *bool
}

// NewPropertyConfiguration returns a new (empty) property configuration
func NewPropertyConfiguration(name string) *PropertyConfiguration {
	return &PropertyConfiguration{
		name: name,
	}
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (pc *PropertyConfiguration) ARMReference() (bool, bool) {
	if pc.armReference == nil {
		return false, false
	}

	return *pc.armReference, true
}

// SetARMReference configures specifies whether this property is an ARM reference or not
func (pc *PropertyConfiguration) SetARMReference(isARMRef bool) *PropertyConfiguration {
	pc.armReference = &isARMRef
	return pc
}

// SetRenamedTo configures what this property is renamed to in the next version of the containing type
func (pc *PropertyConfiguration) SetRenamedTo(renamedTo string) *PropertyConfiguration {
	pc.renamedTo = &renamedTo
	return pc
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (pc *PropertyConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	var lastId string
	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = strings.ToLower(c.Value)
			continue
		}

		if lastId == "$renamedto" && c.Kind == yaml.ScalarNode {
			pc.SetRenamedTo(c.Value)
			continue
		}

		if lastId == "$armreference" && c.Kind == yaml.ScalarNode {
			var isARMRef bool
			err := c.Decode(&isARMRef)
			if err != nil {
				return errors.Wrap(err, "decoding $armReference")
			}

			pc.SetARMReference(isARMRef)
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"property configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}
