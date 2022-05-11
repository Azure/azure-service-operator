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
	name                  string
	properties            map[string]*PropertyConfiguration
	nameInNextVersion     configurableString
	export                configurableBool
	exportAs              configurableString
	azureGeneratedSecrets configurableStringSlice
	advisor               *TypoAdvisor
}

const azureGeneratedSecretsTag = "$azureGeneratedSecrets"

func NewTypeConfiguration(name string) *TypeConfiguration {
	return &TypeConfiguration{
		name:       name,
		properties: make(map[string]*PropertyConfiguration),
		advisor:    NewTypoAdvisor(),
	}
}

// LookupNameInNextVersion checks to see whether the name of this type in the next version is configured, returning
// either that name or a NotConfiguredError.
func (tc *TypeConfiguration) LookupNameInNextVersion() (string, error) {
	name, ok := tc.nameInNextVersion.read()
	if !ok {
		msg := fmt.Sprintf(nameInNextVersionTag+" not specified for type %s", tc.name)
		return "", NewNotConfiguredError(msg)
	}

	return name, nil
}

// VerifyNameInNextVersionConsumed returns an error if our configured rename was not used, nil otherwise.
func (tc *TypeConfiguration) VerifyNameInNextVersionConsumed() error {
	if tc.nameInNextVersion.isUnconsumed() {
		v, _ := tc.nameInNextVersion.read()
		return errors.Errorf("type %s: "+nameInNextVersionTag+": %s not consumed", tc.name, v)
	}

	return nil
}

// WriteNameInNextVersion sets the $nameInNextVersion for testing purposes
func (tc *TypeConfiguration) WriteNameInNextVersion(name string) {
	tc.nameInNextVersion.write(name)
}

// LookupExport checks to see whether this type is configured for export, returning either that value or a
// NotConfiguredError.
func (tc *TypeConfiguration) LookupExport() (bool, error) {
	v, ok := tc.export.read()
	if !ok {
		msg := fmt.Sprintf(exportTag+" not specified for type %s", tc.name)
		return false, NewNotConfiguredError(msg)
	}

	return v, nil
}

// VerifyExportConsumed returns an error if our configured export flag was not used, nil otherwise.
func (tc *TypeConfiguration) VerifyExportConsumed() error {
	if tc.export.isUnconsumed() {
		v, _ := tc.export.read()
		return errors.Errorf("type %s: "+exportTag+": %t not consumed", tc.name, v)
	}

	return nil
}

// LookupExportAs checks to see whether this type has a custom name configured for export, returning either that name
// or a NotConfiguredError.
func (tc *TypeConfiguration) LookupExportAs() (string, error) {
	v, ok := tc.exportAs.read()
	if !ok {
		msg := fmt.Sprintf(exportAsTag+" not specified for type %s", tc.name)
		return "", NewNotConfiguredError(msg)
	}

	return v, nil
}

// VerifyExportAsConsumed returns an error if our configured export name was not used, nil otherwise.
func (tc *TypeConfiguration) VerifyExportAsConsumed() error {
	if tc.exportAs.isUnconsumed() {
		v, _ := tc.exportAs.read()
		return errors.Errorf("type %s: "+exportAsTag+": %s not consumed", tc.name, v)
	}

	return nil
}

// SetAzureGeneratedSecrets sets the list of Azure Generated secrets this type supports
func (tc *TypeConfiguration) SetAzureGeneratedSecrets(secrets []string) *TypeConfiguration {
	tc.azureGeneratedSecrets.write(secrets)
	return tc
}

// AzureGeneratedSecrets gets the list of Azure Generated secrets this type supports
func (tc *TypeConfiguration) AzureGeneratedSecrets() ([]string, error) {
	v, ok := tc.azureGeneratedSecrets.read()
	if !ok {
		msg := fmt.Sprintf("%s not specified for type %s", azureGeneratedSecretsTag, tc.name)
		return nil, NewNotConfiguredError(msg)
	}

	return v, nil
}

// VerifyAzureGeneratedSecretsConsumed returns an error if our configured azureGeneratedSecrets were not used,
// nil otherwise.
func (tc *TypeConfiguration) VerifyAzureGeneratedSecretsConsumed() error {
	if tc.azureGeneratedSecrets.isUnconsumed() {
		return errors.Errorf("type %s: "+azureGeneratedSecretsTag+": not consumed", tc.name)
	}

	return nil
}

// Add includes configuration for the specified property as a part of this type configuration
func (tc *TypeConfiguration) addProperty(name string, property *PropertyConfiguration) {
	// Indexed by lowercase name of the property to allow case-insensitive lookups
	tc.properties[strings.ToLower(name)] = property
}

// visitProperty invokes the provided visitor on the specified property if present.
// Returns a NotConfiguredError if the property is not found; otherwise whatever error is returned by the visitor.
func (tc *TypeConfiguration) visitProperty(
	property astmodel.PropertyName,
	visitor *configurationVisitor,
) error {
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
		err := visitor.visitProperty(pc)
		err = tc.advisor.Wrapf(err, pc.name, "property %s not seen", pc.name)
		errs = append(errs, err)
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
	tc.advisor.AddTerm(string(property))
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

			tc.addProperty(lastId, p)
			continue
		}

		// $nameInNextVersion: <string>
		if strings.EqualFold(lastId, nameInNextVersionTag) && c.Kind == yaml.ScalarNode {
			tc.nameInNextVersion.write(c.Value)
			continue
		}

		// $export: <bool>
		if strings.EqualFold(lastId, exportTag) && c.Kind == yaml.ScalarNode {
			var export bool
			err := c.Decode(&export)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", exportTag)
			}

			tc.export.write(export)
			continue
		}

		// $exportAs: <string>
		if strings.EqualFold(lastId, exportAsTag) && c.Kind == yaml.ScalarNode {
			tc.exportAs.write(c.Value)
			continue
		}

		// $azureGeneratedSecrets:
		// - secret1
		// - secret2
		if strings.EqualFold(lastId, azureGeneratedSecretsTag) && c.Kind == yaml.SequenceNode {
			var azureGeneratedSecrets []string
			for _, content := range c.Content {
				azureGeneratedSecrets = append(azureGeneratedSecrets, content.Value)
			}
			tc.SetAzureGeneratedSecrets(azureGeneratedSecrets)
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
