/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"strings"

	"github.com/rotisserie/eris"
	"gopkg.in/yaml.v3"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/util/typo"
)

// VersionConfiguration contains additional information about a specific version of a group and forms part of a
// hierarchy containing information to supplement the schema and swagger sources consumed by the generator.
//
// ┌──────────────────────────┐       ┌────────────────────┐       ╔══════════════════════╗       ┌───────────────────┐       ┌───────────────────────┐
// │                          │       │                    │       ║                      ║       │                   │       │                       │
// │ ObjectModelConfiguration │───────│ GroupConfiguration │───────║ VersionConfiguration ║───────│ TypeConfiguration │───────│ PropertyConfiguration │
// │                          │1  1..n│                    │1  1..n║                      ║1  1..n│                   │1  1..n│                       │
// └──────────────────────────┘       └────────────────────┘       ╚══════════════════════╝       └───────────────────┘       └───────────────────────┘
type VersionConfiguration struct {
	name    string
	types   map[string]*TypeConfiguration
	advisor *typo.Advisor
}

// NewVersionConfiguration returns a new (empty) VersionConfiguration
func NewVersionConfiguration(name string) *VersionConfiguration {
	return &VersionConfiguration{
		name:    name,
		types:   make(map[string]*TypeConfiguration),
		advisor: typo.NewAdvisor(),
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
	tc := vc.findType(typeName)
	if tc == nil {
		return nil
	}

	err := visitor.visitType(tc)
	if err != nil {
		return eris.Wrapf(err, "configuration of version %s", vc.name)
	}

	return nil
}

// visitTypes invokes the provided visitor on all nested types.
func (vc *VersionConfiguration) visitTypes(visitor *configurationVisitor) error {
	errs := make([]error, 0, len(vc.types))
	for _, tc := range vc.types {
		err := visitor.visitType(tc)
		err = vc.advisor.Wrapf(err, tc.name, "type %s not seen", tc.name)
		errs = append(errs, err)
	}

	// Both errors.Wrapf() and kerrors.NewAggregate() return nil if nothing went wrong
	return eris.Wrapf(
		kerrors.NewAggregate(errs),
		"version %s",
		vc.name)
}

// findType uses the provided name to work out which nested TypeConfiguration should be used
func (vc *VersionConfiguration) findType(name string) *TypeConfiguration {
	vc.advisor.AddTerm(name)
	n := strings.ToLower(name)
	if t, ok := vc.types[n]; ok {
		return t
	}

	return nil
}

// addTypeAlias adds an alias for the specified type, so it may be found with an alternative name
// We use this when $exportedAs is used to rename a type so that any remaining configuration is still accessible under
// the new name.
func (vc *VersionConfiguration) addTypeAlias(name string, alias string) error {
	// Lookup the existing configuration for 'name'
	tc := vc.findType(name)
	if tc == nil {
		return eris.Errorf("unable to create type alias %s", alias)
	}

	// Make sure we don't already have a conflicting configuration for 'alias'
	// if it's already aliased, it's ok if it's the same config, otherwise return an error
	other := vc.findType(alias)
	if other != nil && other != tc {
		return eris.Errorf(
			"unable to create type alias %s for %s because that would conflict with existing configuration",
			alias,
			name)
	}

	// Add the alias as another route to the existing configuration
	vc.addType(alias, tc)
	return nil
}

// UnmarshalYAML populates our instance from the YAML.
// Nested objects are handled by a *pair* of nodes (!!), one for the id and one for the content of the object
func (vc *VersionConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return eris.New("expected mapping")
	}

	vc.types = make(map[string]*TypeConfiguration)
	var lastID string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastID = c.Value
			continue
		}

		// Handle nested kind metadata
		if c.Kind == yaml.MappingNode {
			tc := NewTypeConfiguration(lastID)
			err := c.Decode(&tc)
			if err != nil {
				return eris.Wrapf(err, "decoding yaml for %q", lastID)
			}

			vc.addType(lastID, tc)
			continue
		}

		// No handler for this value, return an error
		return eris.Errorf(
			"version configuration, unexpected yaml value %s: %s (line %d col %d)", lastID, c.Value, c.Line, c.Column)

	}

	return nil
}
