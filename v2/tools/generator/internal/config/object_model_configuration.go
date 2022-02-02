/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"sort"
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
func (omc *ObjectModelConfiguration) TypeRename(name astmodel.TypeName) (string, error) {
	group, err := omc.findGroup(name)
	if err != nil {
		return "", err
	}

	return group.TypeRename(name)
}

// FindUnusedTypeRenames returns a slice listing any unused type renaming configuration
func (omc *ObjectModelConfiguration) FindUnusedTypeRenames() []string {
	var result []string
	for _, gc := range omc.groups {
		result = append(result, gc.FindUnusedTypeRenames()...)
	}

	sort.Strings(result)

	return result
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (omc *ObjectModelConfiguration) ARMReference(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	group, err := omc.findGroup(name)
	if err != nil {
		return false, err
	}

	return group.ARMReference(name, property)
}

// FindUnusedARMReferences returns a slice listing any unused ARMReference configuration
func (omc *ObjectModelConfiguration) FindUnusedARMReferences() []string {
	var result []string
	for _, gc := range omc.groups {
		result = append(result, gc.FindUnusedARMReferences()...)
	}

	sort.Strings(result)

	return result
}

// IsSecret looks up a property to determine whether it is a secret.
func (omc *ObjectModelConfiguration) IsSecret(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	group, err := omc.findGroup(name)
	if err != nil {
		return false, err
	}

	return group.IsSecret(name, property)
}

// Add includes the provided GroupConfiguration in this model configuration
func (omc *ObjectModelConfiguration) Add(group *GroupConfiguration) *ObjectModelConfiguration {
	if omc.groups == nil {
		// Initialize the map just-in-time
		omc.groups = make(map[string]*GroupConfiguration)
	}

	// store the group name using lowercase,
	// so we can do case-insensitive lookups later
	omc.groups[strings.ToLower(group.name)] = group
	return omc
}

// findGroup uses the provided TypeName to work out which nested GroupConfiguration should be used
func (omc *ObjectModelConfiguration) findGroup(name astmodel.TypeName) (*GroupConfiguration, error) {
	group, _, ok := name.PackageReference.GroupVersion()
	if !ok {
		return nil, errors.Errorf(
			"external package reference %s not supported",
			name.PackageReference)
	}

	if g, ok := omc.groups[group]; ok {
		return g, nil
	}

	msg := fmt.Sprintf("no configuration for group %s", group)
	return nil, NewNotConfiguredError(msg).WithOptions("groups", omc.configuredGroups())
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (omc *ObjectModelConfiguration) UnmarshalYAML(value *yaml.Node) error {
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

			omc.Add(g)
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"object model configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}

// configuredGroups returns a sorted slice containing all the groups configured in this group
func (omc *ObjectModelConfiguration) configuredGroups() []string {
	var result []string

	for _, g := range omc.groups {
		// Use the actual names of the groups, not the lower-cased keys of the map
		result = append(result, g.name)
	}

	return result
}
