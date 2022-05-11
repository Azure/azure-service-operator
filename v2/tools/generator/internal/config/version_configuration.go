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
	name    string
	types   map[string]*TypeConfiguration
	advisor *TypoAdvisor
}

// NewVersionConfiguration returns a new (empty) VersionConfiguration
func NewVersionConfiguration(name string) *VersionConfiguration {
	return &VersionConfiguration{
		name:    name,
		types:   make(map[string]*TypeConfiguration),
		advisor: NewTypoAdvisor(),
	}
}

// addType includes configuration for the specified type as a part of this version configuration
func (vc *VersionConfiguration) addType(name string, tc *TypeConfiguration) {
	// Indexed by lowercase name of the type to allow case-insensitive lookups
	vc.types[strings.ToLower(name)] = tc
}

// visitType invokes the provided visitor on the specified type if present.
// Returns a NotConfiguredError if the type is not found; otherwise whatever error is returned by the visitor.
func (vc *VersionConfiguration) visitType(
	typeName string,
	visitor *configurationVisitor,
) error {
	tc, err := vc.findType(typeName)
	if err != nil {
		return err
	}

	return visitor.visitType(tc)
}

// visitTypes invokes the provided visitor on all nested types.
func (vc *VersionConfiguration) visitTypes(visitor *configurationVisitor) error {
	var errs []error
	for _, tc := range vc.types {
		err := visitor.visitType(tc)
		err = vc.advisor.Wrapf(err, tc.name, "type %s not seen", tc.name)
		errs = append(errs, err)
	}

	// Both errors.Wrapf() and kerrors.NewAggregate() return nil if nothing went wrong
	return errors.Wrapf(
		kerrors.NewAggregate(errs),
		"version %s",
		vc.name)
}

// findType uses the provided name to work out which nested TypeConfiguration should be used
func (vc *VersionConfiguration) findType(name string) (*TypeConfiguration, error) {
	vc.advisor.AddTerm(name)
	n := strings.ToLower(name)
	if t, ok := vc.types[n]; ok {
		return t, nil
	}

	msg := fmt.Sprintf("configuration of version %s has no detail for type %s",
		vc.name,
		name)
	return nil, NewNotConfiguredError(msg).WithOptions("types", vc.configuredTypes())
}

// addTypeAlias adds an alias for the specified type, so it may be found with an alternative name
// We use this when $exportedAs is used to rename a type so that any remaining configuration is still accessible under
// the new name.
func (vc *VersionConfiguration) addTypeAlias(name string, alias string) error {
	// Lookup the existing configuration for 'name'
	tc, err := vc.findType(name)
	if err != nil {
		return errors.Wrapf(err, "unable to create type alias %s", alias)
	}

	// Make sure we don't already have a conflicting configuration for 'alias'
	other, err := vc.findType(alias)
	if err != nil {
		// A NotConfiguredError is good, anything else we return
		if !IsNotConfiguredError(err) {
			return errors.Wrapf(err, "unable to create type alias %s", alias)
		}
	} else {
		// if it's already aliased, it's ok if it's the same config, otherwise return an error
		if other != tc {
			return errors.Errorf(
				"unable to create type alias %s for %s because that would conflict with existing configuration",
				alias,
				name)
		}
	}

	// Add the alias as another route to the existing configuration
	vc.addType(alias, tc)
	return nil
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

			vc.addType(lastId, tc)
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
