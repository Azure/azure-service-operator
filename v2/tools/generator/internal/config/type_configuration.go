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
type TypeConfiguration struct {
	name                     string
	properties               map[string]*PropertyConfiguration
	nameInNextVersion        configurable[string]
	export                   configurable[bool]
	exportAs                 configurable[string]
	azureGeneratedSecrets    configurable[[]string]
	azureGeneratedConfigs    configurable[[]string]
	supportedFrom            configurable[string]
	isResource               configurable[bool]
	resourceEmbeddedInParent configurable[string]
	importable               configurable[bool]
	advisor                  *TypoAdvisor
}

const (
	azureGeneratedSecretsTag    = "$azureGeneratedSecrets"    // A set of strings specifying which secrets are generated by Azure
	azureGeneratedConfigsTag    = "$azureGeneratedConfigs"    // A set of strings specifying which config maps are generated by Azure
	exportTag                   = "$export"                   // Boolean specifying whether a resource type is exported
	exportAsTag                 = "$exportAs"                 // String specifying the name to use for a type (implies $export: true)
	importableTag               = "$importable"               // Boolean specifying whether a resource type is importable via asoctl (defaults to true)
	isResourceTag               = "$isResource"               // Boolean specifying whether a particular type is a resource or not.
	nameInNextVersionTag        = "$nameInNextVersion"        // String specifying a type or property name change in the next version
	supportedFromTag            = "$supportedFrom"            // Label specifying the first ASO release supporting the resource
	resourceEmbeddedInParentTag = "$resourceEmbeddedInParent" // String specifying resource name of parent
)

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

// SetNameInNextVersion sets the configured $nameInNextVersion for this type
func (tc *TypeConfiguration) SetNameInNextVersion(name string) {
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

// SetAzureGeneratedConfigs sets the list of Azure Generated configmaps this type supports
func (tc *TypeConfiguration) SetAzureGeneratedConfigs(configMaps []string) *TypeConfiguration {
	tc.azureGeneratedConfigs.write(configMaps)
	return tc
}

// AzureGeneratedConfigs gets the list of Azure Generated config maps this type supports
func (tc *TypeConfiguration) AzureGeneratedConfigs() ([]string, error) {
	v, ok := tc.azureGeneratedConfigs.read()
	if !ok {
		msg := fmt.Sprintf("%s not specified for type %s", azureGeneratedConfigsTag, tc.name)
		return nil, NewNotConfiguredError(msg)
	}

	return v, nil
}

// VerifyAzureGeneratedConfigsConsumed returns an error if our configured azureGeneratedConfigs were not used,
// nil otherwise.
func (tc *TypeConfiguration) VerifyAzureGeneratedConfigsConsumed() error {
	if tc.azureGeneratedConfigs.isUnconsumed() {
		return errors.Errorf("type %s: "+azureGeneratedConfigsTag+": not consumed", tc.name)
	}

	return nil
}

// LookupSupportedFrom checks to see whether this type has its first ASO release configured, returning either that
// release or a NotConfiguredError.
func (tc *TypeConfiguration) LookupSupportedFrom() (string, error) {
	v, ok := tc.supportedFrom.read()
	if !ok {
		msg := fmt.Sprintf(supportedFromTag+" not specified for type %s", tc.name)
		return "", NewNotConfiguredError(msg)
	}

	return v, nil
}

// VerifySupportedFromConsumed returns an error if our configured supportedFrom tag was not used, nil otherwise.
func (tc *TypeConfiguration) VerifySupportedFromConsumed() error {
	if tc.supportedFrom.isUnconsumed() {
		v, _ := tc.supportedFrom.read()
		return errors.Errorf("type %s: "+supportedFromTag+": %s not consumed", tc.name, v)
	}

	return nil
}

// SetSupportedFrom sets the configured $supportedFrom for this type
func (tc *TypeConfiguration) SetSupportedFrom(from string) {
	tc.supportedFrom.write(from)
}

// LookupResourceEmbeddedInParent checks to see whether this type is a resource embedded in its parent
func (tc *TypeConfiguration) LookupResourceEmbeddedInParent() (string, error) {
	v, ok := tc.resourceEmbeddedInParent.read()
	if !ok {
		msg := fmt.Sprintf(resourceEmbeddedInParentTag+" not specified for type %s", tc.name)
		return "", NewNotConfiguredError(msg)
	}

	return v, nil
}

// VerifyResourceEmbeddedInParentConsumed returns an error if our configured isResource flag was not used, nil otherwise.
func (tc *TypeConfiguration) VerifyResourceEmbeddedInParentConsumed() error {
	if tc.resourceEmbeddedInParent.isUnconsumed() {
		v, _ := tc.export.read()
		return errors.Errorf("type %s: "+resourceEmbeddedInParentTag+": %t not consumed", tc.name, v)
	}

	return nil
}

// LookupIsResource checks to see whether this type is a resource embedded in its parent
func (tc *TypeConfiguration) LookupIsResource() (bool, error) {
	v, ok := tc.isResource.read()
	if !ok {
		msg := fmt.Sprintf(isResourceTag+" not specified for type %s", tc.name)
		return false, NewNotConfiguredError(msg)
	}

	return v, nil
}

// VerifyIsResourceConsumed returns an error if our configured isResource flag was not used, nil otherwise.
func (tc *TypeConfiguration) VerifyIsResourceConsumed() error {
	if tc.isResource.isUnconsumed() {
		v, _ := tc.export.read()
		return errors.Errorf("type %s: "+isResourceTag+": %t not consumed", tc.name, v)
	}

	return nil
}

// LookupImportable checks to see whether this resource type is importable via asoctl
func (tc *TypeConfiguration) LookupImportable() (bool, error) {
	v, ok := tc.importable.read()
	if !ok {
		msg := fmt.Sprintf(importableTag+" not specified for type %s", tc.name)
		return false, NewNotConfiguredError(msg)
	}

	return v, nil
}

// VerifyImportable consumed returns an error if our configured importable flag was not used, nil otherwise.
func (tc *TypeConfiguration) VerifyImportableConsumed() error {
	if tc.importable.isUnconsumed() {
		v, _ := tc.importable.read()
		return errors.Errorf("type %s: "+importableTag+": %t not consumed", tc.name, v)
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

	err = visitor.visitProperty(pc)
	if err != nil {
		return errors.Wrapf(err, "configuration of type %s", tc.name)
	}

	return nil
}

// visitProperties invokes the provided visitor on all properties.
func (tc *TypeConfiguration) visitProperties(visitor *configurationVisitor) error {
	errs := make([]error, 0, len(tc.properties))
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
				if content.Kind == yaml.ScalarNode {
					azureGeneratedSecrets = append(azureGeneratedSecrets, content.Value)
				} else {
					return errors.Errorf(
						"unexpected yam value for %s (line %d col %d)",
						azureGeneratedSecretsTag,
						content.Line,
						content.Column)
				}
			}

			tc.SetAzureGeneratedSecrets(azureGeneratedSecrets)
			continue
		}

		// $supportedFrom
		if strings.EqualFold(lastId, supportedFromTag) && c.Kind == yaml.ScalarNode {
			tc.supportedFrom.write(c.Value)
			continue
		}

		// $resourceEmbeddedInParent: <string>
		if strings.EqualFold(lastId, resourceEmbeddedInParentTag) && c.Kind == yaml.ScalarNode {
			var resourceEmbeddedInParent string
			err := c.Decode(&resourceEmbeddedInParent)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", resourceEmbeddedInParentTag)
			}

			tc.resourceEmbeddedInParent.write(resourceEmbeddedInParent)
			continue
		}

		// $isResource: <bool>
		if strings.EqualFold(lastId, isResourceTag) && c.Kind == yaml.ScalarNode {
			var isResource bool
			err := c.Decode(&isResource)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", isResourceTag)
			}

			tc.isResource.write(isResource)
			continue
		}

		// $importable: <bool>
		if strings.EqualFold(lastId, importableTag) && c.Kind == yaml.ScalarNode {
			var importable bool
			err := c.Decode(&importable)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", importableTag)
			}

			tc.importable.write(importable)
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
	result := make([]string, 0, len(tc.properties))
	for _, c := range tc.properties {
		// Use the actual names of the properties, not the lower-cased keys of the map
		result = append(result, c.name)
	}

	return result
}
