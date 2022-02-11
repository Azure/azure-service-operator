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

// LookupNameInNextVersion checks whether we have an alternative name for the specified type, returning the name if
// found. Returns a NotConfiguredError if no rename is available.
func (omc *ObjectModelConfiguration) LookupNameInNextVersion(name astmodel.TypeName) (string, error) {
	var newName string
	visitor := NewSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			n, err := configuration.LookupNameInNextVersion()
			newName = n
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return "", err
	}

	return newName, nil
}

// VerifyNameInNextVersionConsumed returns an error if any configured type renames were not consumed
func (omc *ObjectModelConfiguration) VerifyNameInNextVersionConsumed() error {
	visitor := NewEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyNameInNextVersionConsumed()
		})
	return visitor.Visit(omc)
}

// LookupExport checks to see whether a specified type is configured for export, returning the value if found. Returns a
// NotConfiguredError if no export is configured.
func (omc *ObjectModelConfiguration) LookupExport(name astmodel.TypeName) (bool, error) {
	var export bool
	visitor := NewSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			ex, err := configuration.LookupExport()
			export = ex
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return false, err
	}

	return export, nil
}

// VerifyExportConsumed returns an error if our configured export flag was not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifyExportConsumed() error {
	visitor := NewEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyExportConsumed()
		})
	return visitor.Visit(omc)
}

// LookupExportAs checks to see whether a specified type is configured for export with an alternative name, returning the
// name if found. Returns a NotConfiguredError if no export is configured.
func (omc *ObjectModelConfiguration) LookupExportAs(name astmodel.TypeName) (string, error) {
	var exportAs string
	visitor := NewSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			ea, err := configuration.LookupExportAs()
			exportAs = ea
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return "", err
	}

	return exportAs, nil
}

// VerifyExportAsConsumed returns an error if our configured export name was not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifyExportAsConsumed() error {
	visitor := NewEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyExportAsConsumed()
		})
	return visitor.Visit(omc)
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
// Returns true or false if configured, or a NotConfiguredError if not.
func (omc *ObjectModelConfiguration) ARMReference(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	var result bool
	visitor := NewSinglePropertyConfigurationVisitor(
		name,
		property,
		func(configuration *PropertyConfiguration) error {
			isArmReference, err := configuration.ARMReference()
			result = isArmReference
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return false, err
	}

	return result, nil
}

// VerifyARMReferencesConsumed returns an error if any ARM Reference configuration was not consumed
func (omc *ObjectModelConfiguration) VerifyARMReferencesConsumed() error {
	visitor := NewEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			return configuration.VerifyARMReferenceConsumed()
		})
	return visitor.Visit(omc)
}

// IsSecret looks up a property to determine whether it is a secret.
func (omc *ObjectModelConfiguration) IsSecret(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	var result bool
	visitor := NewSinglePropertyConfigurationVisitor(
		name,
		property,
		func(configuration *PropertyConfiguration) error {
			isSecret, err := configuration.IsSecret()
			result = isSecret
			return err
		})

	err := visitor.Visit(omc)
	if err != nil {
		return false, err
	}

	return result, nil
}

// VerifyIsSecretConsumed returns an error if any IsSecret configuration was not consumed
func (omc *ObjectModelConfiguration) VerifyIsSecretConsumed() error {
	visitor := NewEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			return configuration.VerifyIsSecretConsumed()
		})
	return visitor.Visit(omc)
}

// Add includes the provided GroupConfiguration in this model configuration
func (omc *ObjectModelConfiguration) add(group *GroupConfiguration) {
	if omc.groups == nil {
		// Initialize the map just-in-time
		omc.groups = make(map[string]*GroupConfiguration)
	}

	// store the group name using lowercase,
	// so we can do case-insensitive lookups later
	omc.groups[strings.ToLower(group.name)] = group
}

// visitGroup invokes the provided visitor on the specified group if present.
// Returns a NotConfiguredError if the group is not found; otherwise whatever error is returned by the visitor.
func (omc *ObjectModelConfiguration) visitGroup(
	name astmodel.TypeName,
	visitor *configurationVisitor) error {

	group, err := omc.findGroup(name)
	if err != nil {
		return err
	}

	return visitor.visitGroup(group)
}

// visitGroups invokes the provided visitor on all nested groups.
func (omc *ObjectModelConfiguration) visitGroups(visitor *configurationVisitor) error {
	var errs []error
	for _, gc := range omc.groups {
		errs = append(errs, visitor.visitGroup(gc))
	}

	// kerrors.NewAggregate() returns nil if nothing went wrong
	return kerrors.NewAggregate(errs)
}

// findGroup uses the provided TypeName to work out which nested GroupConfiguration should be used
func (omc *ObjectModelConfiguration) findGroup(name astmodel.TypeName) (*GroupConfiguration, error) {
	group, _, ok := name.PackageReference.GroupVersion()
	if !ok {
		return nil, errors.Errorf(
			"external package reference %s not supported",
			name.PackageReference)
	}

	if omc == nil || omc.groups == nil {
		msg := fmt.Sprintf("no configuration for group %s", group)
		return nil, NewNotConfiguredError(msg)
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

			omc.add(g)
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

// CreateTestObjectModelConfiguration builds up a new configuration for a particular type, returning both the top level
// configuration and the type configuration.
// While intended only for test use, this isn't in a _test.go file as we want to use it from tests in multiple packages.
func CreateTestObjectModelConfiguration(name astmodel.TypeName) (*ObjectModelConfiguration, *TypeConfiguration) {
	group, version, ok := name.PackageReference.GroupVersion()
	if !ok {
		panic(fmt.Sprintf("unexpected external package reference for resource name %s", name))
	}

	typeConfig := NewTypeConfiguration(name.Name())

	versionConfig := NewVersionConfiguration(version)
	versionConfig.add(typeConfig)

	groupConfig := NewGroupConfiguration(group)
	groupConfig.add(versionConfig)

	objectModelConfig := NewObjectModelConfiguration()
	objectModelConfig.add(groupConfig)

	return objectModelConfig, typeConfig
}

// CreateTestObjectModelConfigurationForRename builds up a new configuring for testing rename of a particular type.
// While intended only for test use, this isn't in a _test.go file as we want to use it from tests in multiple packages.
func CreateTestObjectModelConfigurationForRename(name astmodel.TypeName, newName string) *ObjectModelConfiguration {
	omc, tc := CreateTestObjectModelConfiguration(name)
	tc.nameInNextVersion.write(newName)
	return omc
}
