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
	kerrors "k8s.io/apimachinery/pkg/util/errors"

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

// Add includes configuration for the specified version as a part of this group configuration
// In addition to indexing by the name of the version, we also index by the local-package-name and storage-package-name
// of the version, so we can do lookups via TypeName. All indexing is lower-case to allow case-insensitive lookups (this
// makes our configuration more forgiving).
func (gc *GroupConfiguration) add(version *VersionConfiguration) {
	// Convert version.name into a package version
	// We do this by constructing a local package reference because this avoids replicating the logic here and risking
	// inconsistency if things are changed in the future.
	local := astmodel.MakeLocalPackageReference("prefix", "group", astmodel.GeneratorVersionPrefix, version.name)
	_, lv, ok := local.GroupVersion()
	if !ok {
		msg := fmt.Sprintf("local package reference %s unexpectedly failed to return GroupVersion()", local)
		panic(msg)
	}

	// Convert version.name into a storage package version
	// We do this by constructing a storage package reference for reasons similar to above.
	storage := astmodel.MakeStoragePackageReference(local)
	_, sv, ok := storage.GroupVersion()
	if !ok {
		msg := fmt.Sprintf("storage package reference %s unexpectedly failed to return GroupVersion()", storage)
		panic(msg)
	}

	gc.versions[strings.ToLower(version.name)] = version
	gc.versions[strings.ToLower(lv)] = version
	gc.versions[strings.ToLower(sv)] = version
}

// visitVersion invokes the provided visitor on the specified version if present.
// Returns a NotConfiguredError if the version is not found; otherwise whatever error is returned by the visitor.
func (gc *GroupConfiguration) visitVersion(
	name astmodel.TypeName,
	visitor *configurationVisitor) error {

	vc, err := gc.findVersion(name)
	if err != nil {
		return err
	}

	return visitor.visitVersion(vc)
}

// visitVersions invokes the provided visitor on all versions.
func (gc *GroupConfiguration) visitVersions(visitor *configurationVisitor) error {
	var errs []error

	// All our versions are listed under multiple keys, so we hedge against processing them multiple times
	versionsSeen := astmodel.MakeStringSet()
	for _, v := range gc.versions {
		if versionsSeen.Contains(v.name) {
			continue
		}

		errs = append(errs, visitor.visitVersion(v))
		versionsSeen.Add(v.name)
	}

	// Both errors.Wrapf() and kerrors.NewAggregate() return nil if nothing went wrong
	return errors.Wrapf(
		kerrors.NewAggregate(errs),
		"group %s",
		gc.name)
}

// findVersion uses the provided TypeName to work out which nested VersionConfiguration should be used
func (gc *GroupConfiguration) findVersion(name astmodel.TypeName) (*VersionConfiguration, error) {
	ref := name.PackageReference
	if s, ok := ref.(astmodel.StoragePackageReference); ok {
		// If we have a storage package reference, need to unwrap the actual local package reference
		// as all our configuration is based on API versions, not storage versions
		ref = s.Local()
	}

	v := strings.ToLower(ref.PackageName())
	if version, ok := gc.versions[v]; ok {
		return version, nil
	}

	msg := fmt.Sprintf(
		"configuration of group %s has no detail for version %s",
		gc.name,
		ref.PackageName())
	return nil, NewNotConfiguredError(msg).WithOptions("versions", gc.configuredVersions())
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

			gc.add(v)
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

	return result
}
