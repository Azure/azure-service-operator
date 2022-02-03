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
	name              string
	renamedTo         *string
	renamedToConsumed bool
	properties        map[string]*PropertyConfiguration
}

func NewTypeConfiguration(name string) *TypeConfiguration {
	return &TypeConfiguration{
		name:       name,
		properties: make(map[string]*PropertyConfiguration),
	}
}

// TypeRename returns a new name (and true) if one is configured for this type, or empty string and false if not.
func (tc *TypeConfiguration) TypeRename() (string, error) {
	if tc.renamedTo == nil {
		msg := fmt.Sprintf(renamedToTag+" not specified for type %s", tc.name)
		return "", NewNotConfiguredError(msg)
	}

	tc.renamedToConsumed = true
	return *tc.renamedTo, nil
}

// SetTypeRename sets the name this type is renamed to
func (tc *TypeConfiguration) SetTypeRename(renameTo string) *TypeConfiguration {
	tc.renamedTo = &renameTo
	return tc
}

// VerifyTypeRenameConsumed returns an error if our configured rename was not used, nil otherwise.
func (tc *TypeConfiguration) VerifyTypeRenameConsumed() error {
	if tc.renamedTo != nil && !tc.renamedToConsumed {
		return errors.Errorf("type %s: "+renamedToTag+": %s not consumed", tc.name, *tc.renamedTo)
	}

	return nil
}

// Add includes configuration for the specified property as a part of this type configuration
func (tc *TypeConfiguration) Add(property *PropertyConfiguration) *TypeConfiguration {
	// Indexed by lowercase name of the property to allow case-insensitive lookups
	tc.properties[strings.ToLower(property.name)] = property
	return tc
}

// visitProperty invokes the provided visitor on the specified property if present.
// Returns a NotConfiguredError if the property is not found; otherwise whatever error is returned by the visitor.
func (tc *TypeConfiguration) visitProperty(
	property astmodel.PropertyName,
	visitor *configurationVisitor) error {

	pc, err := tc.findProperty(property)
	if err != nil {
		return err
	}

	return visitor.visitProperty(pc)
}

// visitProperties invokes the provided visitor on all properties.
func (tc *TypeConfiguration) visitProperties(visitor *configurationVisitor) error {
	var errs []error
	for _, pc := range tc.properties {
		errs = append(errs, visitor.visitProperty(pc))
	}

	// Both errors.Wrapf() and kerrors.NewAggregate() return nil if nothing went wrong
	return errors.Wrapf(
		kerrors.NewAggregate(errs),
		"type %s",
		tc.name)
}

// findProperty uses the provided property name to work out which nested PropertyConfiguration should be used
// either returns the requested property configuration, or an error saying that it couldn't be found
func (tc *TypeConfiguration) findProperty(property astmodel.PropertyName) (*PropertyConfiguration, error) {
	// Store the property id using lowercase,
	// so we can do case-insensitive lookups later
	p := strings.ToLower(string(property))
	if pc, ok := tc.properties[p]; ok {
		return pc, nil
	}

	msg := fmt.Sprintf(
		"configuration of type %s has no detail for property %s",
		tc.name,
		property)
	return nil, NewNotConfiguredError(msg).WithOptions("properties", tc.configuredProperties())
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

		if strings.EqualFold(lastId, renamedToTag) && c.Kind == yaml.ScalarNode {
			tc.SetTypeRename(c.Value)
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"type configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}

// configuredProperties returns a sorted slice containing all the properties configured on this type
func (tc *TypeConfiguration) configuredProperties() []string {
	var result []string
	for _, c := range tc.properties {
		// Use the actual names of the properties, not the lower-cased keys of the map
		result = append(result, c.name)
	}

	return result
}
