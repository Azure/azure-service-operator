/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
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
	name          string
	renamedTo     *string
	usedRenamedTo bool
	properties    map[string]*PropertyConfiguration
}

func NewTypeConfiguration(name string) *TypeConfiguration {
	return &TypeConfiguration{
		name:       name,
		properties: make(map[string]*PropertyConfiguration),
	}
}

// TypeRename returns a new name (and true) if one is configured for this type, or empty string and false if not.
func (tc *TypeConfiguration) TypeRename() (string, bool) {
	if tc.renamedTo != nil {
		tc.usedRenamedTo = true
		return *tc.renamedTo, true
	}

	return "", false
}

// SetTypeRename sets the name this type is renamed to
func (tc *TypeConfiguration) SetTypeRename(renameTo string) *TypeConfiguration {
	tc.renamedTo = &renameTo
	return tc
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (tc *TypeConfiguration) ARMReference(property astmodel.PropertyName) (bool, bool) {
	pc, ok := tc.findProperty(property)
	if !ok {
		return false, false
	}

	return pc.ARMReference()
}

// FindUnusedARMReferences returns a slice listing any unused ARMReference configuration
func (tc *TypeConfiguration) FindUnusedARMReferences() []string {
	var result []string
	for _, pc := range tc.properties {
		for _, s := range pc.FindUnusedARMReferences() {
			msg := fmt.Sprintf("type %s %s", tc.name, s)
			result = append(result, msg)
		}
	}

	return result
}

// Add includes configuration for the specified property as a part of this type configuration
func (tc *TypeConfiguration) Add(property *PropertyConfiguration) *TypeConfiguration {
	// Indexed by lowercase name of the property to allow case insensitive lookups
	tc.properties[strings.ToLower(property.name)] = property
	return tc
}

// findProperty uses the provided property name to work out which nested PropertyConfiguration should be used
func (tc *TypeConfiguration) findProperty(property astmodel.PropertyName) (*PropertyConfiguration, bool) {
	// Store the property id using lowercase,
	// so we can do case-insensitive lookups later
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
			p := NewPropertyConfiguration(lastId)
			err := c.Decode(p)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			tc.Add(p)
			continue
		}

		if strings.ToLower(lastId) == "$renamedto" && c.Kind == yaml.ScalarNode {
			tc.SetTypeRename(c.Value)
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"type configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}
