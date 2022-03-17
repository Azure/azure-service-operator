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

// Add includes configuration for the specified type as a part of this version configuration
func (vc *VersionConfiguration) add(tc *TypeConfiguration) {
	// Indexed by lowercase name of the type to allow case-insensitive lookups
	vc.types[strings.ToLower(tc.name)] = tc
}

// visitType invokes the provided visitor on the specified type if present.
// Returns a NotConfiguredError if the type is not found; otherwise whatever error is returned by the visitor.
func (vc *VersionConfiguration) visitType(
	typeName astmodel.TypeName,
	visitor *configurationVisitor,
) error {
	tc, err := vc.findType(typeName.Name())
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
	n := strings.ToLower(name)
	if t, ok := vc.types[n]; ok {
		vc.advisor.AddTerm(t.name)
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

			vc.add(tc)
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
