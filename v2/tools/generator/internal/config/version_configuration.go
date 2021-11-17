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

// VersionConfiguration contains additional information about a specific version of a group and forms part of a
// hierarchy containing information to supplement the schema and swagger sources consumed by the generator.
//
// ┌──────────────────────────┐       ┌────────────────────┐       ╔══════════════════════╗       ┌───────────────────┐       ┌───────────────────────┐
// │                          │       │                    │       ║                      ║       │                   │       │                       │
// │ ObjectModelConfiguration │───────│ GroupConfiguration │───────║ VersionConfiguration ║───────│ TypeConfiguration │───────│ PropertyConfiguration │
// │                          │1  1..n│                    │1  1..n║                      ║1  1..n│                   │1  1..n│                       │
// └──────────────────────────┘       └────────────────────┘       ╚══════════════════════╝       └───────────────────┘       └───────────────────────┘
//
type VersionConfiguration struct {
	name  string
	types map[string]*TypeConfiguration
}

// NewVersionConfiguration returns a new (empty) VersionConfiguration
func NewVersionConfiguration(name string) *VersionConfiguration {
	return &VersionConfiguration{
		name:  name,
		types: make(map[string]*TypeConfiguration),
	}
}

// TypeRename looks up a rename for the specified type, returning the new name and true if found, or empty string
// and false if not.
func (vc *VersionConfiguration) TypeRename(name string) (string, error) {
	tc, err := vc.findType(name)
	if err != nil {
		return "", err
	}

	rename, err := tc.TypeRename()
	if err != nil {
		return "", errors.Wrapf(
			err,
			"configuration of version %s",
			vc.name)
	}

	return rename, nil
}

// FindUnusedTypeRenames returns a slice listing any unused type rename configuration
func (vc *VersionConfiguration) FindUnusedTypeRenames() []string {
	var result []string
	for _, tc := range vc.types {
		result = appendWithPrefix(result, fmt.Sprintf("version %s ", vc.name), tc.FindUnusedTypeRenames()...)
	}

	return result
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (vc *VersionConfiguration) ARMReference(name string, property astmodel.PropertyName) (bool, error) {
	tc, err := vc.findType(name)
	if err != nil {
		return false, err
	}

	armReference, err := tc.ARMReference(property)
	if err != nil {
		return false, errors.Wrapf(
			err,
			"configuration of version %s",
			vc.name)
	}

	return armReference, nil
}

// FindUnusedARMReferences returns a slice listing any unused ARMReference configuration
func (vc *VersionConfiguration) FindUnusedARMReferences() []string {
	var result []string
	for _, tc := range vc.types {
		result = appendWithPrefix(result, fmt.Sprintf("version %s ", vc.name), tc.FindUnusedARMReferences()...)
	}

	return result
}

// Add includes configuration for the specified type as a part of this version configuration
func (vc *VersionConfiguration) Add(tc *TypeConfiguration) *VersionConfiguration {
	// Indexed by lowercase name of the type to allow case-insensitive lookups
	vc.types[strings.ToLower(tc.name)] = tc
	return vc
}

// findType uses the provided name to work out which nested TypeConfiguration should be used
func (vc *VersionConfiguration) findType(name string) (*TypeConfiguration, error) {
	n := strings.ToLower(name)
	if t, ok := vc.types[n]; ok {
		return t, nil
	}

	msg := fmt.Sprintf("configuration of version %s has no detail for type %s",
		vc.name,
		name)
	return nil, NewNotConfiguredError(msg).WithOptions("types", vc.configuredTypes())
}

// UnmarshalYAML populates our instance from the YAML.
// Nested objects are handled by a *pair* of nodes (!!), one for the id and one for the content of the object
func (vc *VersionConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	vc.types = make(map[string]*TypeConfiguration)
	var lastId string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = c.Value
			continue
		}

		// Handle nested kind metadata
		if c.Kind == yaml.MappingNode {
			tc := NewTypeConfiguration(lastId)
			err := c.Decode(&tc)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			vc.Add(tc)
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"version configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}

// configuredTypes returns a sorted slice containing all the properties configured on this type
func (vc *VersionConfiguration) configuredTypes() []string {
	var result []string
	for _, t := range vc.types {
		// Use the actual names of the types, not the lower-cased keys of the map
		result = append(result, t.name)
	}

	return result
}
