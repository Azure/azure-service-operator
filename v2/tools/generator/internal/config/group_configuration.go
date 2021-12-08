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
	name     string
	versions map[string]*VersionConfiguration
}

// NewGroupConfiguration returns a new (empty) GroupConfiguration
func NewGroupConfiguration(name string) *GroupConfiguration {
	return &GroupConfiguration{
		name:     name,
		versions: make(map[string]*VersionConfiguration),
	}
}

// TypeRename looks up a rename for the specified type, returning the new name and true if found, or empty string
// and false if not.
func (gc *GroupConfiguration) TypeRename(name astmodel.TypeName) (string, error) {
	version, err := gc.findVersion(name)
	if err != nil {
		return "", err
	}

	rename, err := version.TypeRename(name.Name())
	if err != nil {
		return "", errors.Wrapf(
			err,
			"configuration of group %s",
			gc.name)
	}

	return rename, nil
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (gc *GroupConfiguration) ARMReference(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	version, err := gc.findVersion(name)
	if err != nil {
		return false, err
	}

	armReference, err := version.ARMReference(name.Name(), property)
	if err != nil {
		return false, errors.Wrapf(
			err,
			"configuration of group %s",
			gc.name)
	}

	return armReference, nil
}

// FindUnusedARMReferences returns a slice listing any unused ARMReference configuration
func (gc *GroupConfiguration) FindUnusedARMReferences() []string {
	var result []string

	// All our versions are listed twice, under two different keys, so we hedge against processing them multiple times
	versionsSeen := astmodel.MakeStringSet()
	for _, vc := range gc.versions {
		if versionsSeen.Contains(vc.name) {
			continue
		}

		versionsSeen.Add(vc.name)
		for _, s := range vc.FindUnusedARMReferences() {
			msg := fmt.Sprintf("group %s %s", gc.name, s)
			result = append(result, msg)
		}
	}

	return result
}

// Add includes configuration for the specified version as a part of this group configuration
// In addition to indexing by the name of the version, we also index by the local-package-name of the version, so we can
// do lookups via TypeName. All indexing is lower-case to allow case-insensitive lookups (this makes our configuration
// more forgiving).
func (gc *GroupConfiguration) Add(version *VersionConfiguration) *GroupConfiguration {
	pkg := astmodel.CreateLocalPackageNameFromVersion(version.name)
	gc.versions[strings.ToLower(version.name)] = version
	gc.versions[strings.ToLower(pkg)] = version
	return gc
}

// findVersion uses the provided TypeName to work out which nested VersionConfiguration should be used
func (gc *GroupConfiguration) findVersion(name astmodel.TypeName) (*VersionConfiguration, error) {
	v := strings.ToLower(name.PackageReference.PackageName())
	if version, ok := gc.versions[v]; ok {
		return version, nil
	}

	return nil, errors.Errorf(
		"configuration of group %s has no detail for version %s (configured versions are: %s)",
		gc.name,
		name,
		strings.Join(gc.configuredVersions(), ", "))

}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (gc *GroupConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	gc.versions = make(map[string]*VersionConfiguration)
	var lastId string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = c.Value
			continue
		}

		// Handle nested version metadata
		if c.Kind == yaml.MappingNode {
			v := NewVersionConfiguration(lastId)
			err := c.Decode(&v)
			if err != nil {
				return errors.Wrapf(err, "decoding yaml for %q", lastId)
			}

			gc.Add(v)
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"group configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}

// configuredVersions returns a sorted slice containing all the versions configured in this group
func (gc *GroupConfiguration) configuredVersions() []string {
	var result []string

	// All our versions are listed twice, under two different keys, so we hedge against processing them multiple times
	versionsSeen := astmodel.MakeStringSet()
	for _, v := range gc.versions {
		if versionsSeen.Contains(v.name) {
			continue
		}

		// Use the actual names of the versions, not the lower-cased keys of the map
		result = append(result, v.name)
		versionsSeen.Add(v.name)
	}

	sort.Strings(result)
	return result
}
