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
	renamedTo        string
	armReference     bool
	haveArmReference bool
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (pc *PropertyConfiguration) ARMReference() (bool, bool) {
	if !pc.haveArmReference {
		return false, false
	}

	return pc.armReference, true
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
			pc.renamedTo = c.Value
			continue
		}

		if lastId == "$armreference" && c.Kind == yaml.ScalarNode {
			err := c.Decode(&pc.armReference)
			if err != nil {
				return errors.Wrap(err, "decoding $armReference")
			}

			pc.haveArmReference = true
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf("property configuration, unexpected yaml value %s line %d col %d)", c.Value, c.Line, c.Column)
	}

	return nil
}
