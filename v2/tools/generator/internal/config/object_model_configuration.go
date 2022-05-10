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
	groups      map[string]*GroupConfiguration // nested configuration for individual groups
	typoAdvisor *TypoAdvisor
}

// NewObjectModelConfiguration returns a new (empty) ObjectModelConfiguration
func NewObjectModelConfiguration() *ObjectModelConfiguration {
	return &ObjectModelConfiguration{
		groups:      make(map[string]*GroupConfiguration),
		typoAdvisor: NewTypoAdvisor(),
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

// AzureGeneratedSecrets looks up a type to determine if it has any Azure generated secrets
func (omc *ObjectModelConfiguration) AzureGeneratedSecrets(name astmodel.TypeName) ([]string, error) {
	var result []string
	visitor := NewSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			var err error
			result, err = configuration.AzureGeneratedSecrets()
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// VerifyAzureGeneratedSecretsConsumed returns an error if Azure generated secrets were not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifyAzureGeneratedSecretsConsumed() error {
	visitor := NewEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyAzureGeneratedSecretsConsumed()
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

// IsResourceLifecycleOwnedByParent returns true if the property represents a subresource whose resource lifecycle is owned by the parent resource.
// False is returned if the property does not represent a subresource whose lifecycle is owned by the parent.
func (omc *ObjectModelConfiguration) IsResourceLifecycleOwnedByParent(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	var result bool
	visitor := NewSinglePropertyConfigurationVisitor(
		name,
		property,
		func(configuration *PropertyConfiguration) error {
			isResourceLifecycleOwnedByParent, err := configuration.IsResourceLifecycleOwnedByParent()
			result = isResourceLifecycleOwnedByParent
			return err
		})

	err := visitor.Visit(omc)
	if err != nil {
		return false, err
	}

	return result, nil
}

// MarkIsResourceLifecycleOwnedByParentUnconsumed marks all IsResourceLifecycleOwnedByParent as unconsumed
func (omc *ObjectModelConfiguration) MarkIsResourceLifecycleOwnedByParentUnconsumed() error {
	visitor := NewEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			configuration.ClearResourceLifecycleOwnedByParentConsumed()
			return nil
		})
	return visitor.Visit(omc)
}

// VerifyIsResourceLifecycleOwnedByParentConsumed returns an error if any IsResourceLifecycleOwnedByParent configuration
// was not consumed
func (omc *ObjectModelConfiguration) VerifyIsResourceLifecycleOwnedByParentConsumed() error {
	visitor := NewEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			return configuration.VerifyIsResourceLifecycleOwnedByParentConsumed()
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
	ref astmodel.PackageReference,
	visitor *configurationVisitor,
) error {
	group, err := omc.findGroup(ref)
	if err != nil {
		return err
	}

	return visitor.visitGroup(group)
}

// visitGroups invokes the provided visitor on all nested groups.
func (omc *ObjectModelConfiguration) visitGroups(visitor *configurationVisitor) error {
	var errs []error
	for _, gc := range omc.groups {
		err := visitor.visitGroup(gc)
		err = omc.typoAdvisor.Wrapf(err, gc.name, "group %s not seen", gc.name)
		errs = append(errs, err)
	}

	// kerrors.NewAggregate() returns nil if nothing went wrong
	return kerrors.NewAggregate(errs)
}

// findGroup uses the provided TypeName to work out which nested GroupConfiguration should be used
func (omc *ObjectModelConfiguration) findGroup(ref astmodel.PackageReference) (*GroupConfiguration, error) {
	group, _, ok := ref.GroupVersion()
	if !ok {
		return nil, errors.Errorf(
			"external package reference %s not supported",
			ref)
	}

	if omc == nil || omc.groups == nil {
		msg := fmt.Sprintf("no configuration for group %s", group)
		return nil, NewNotConfiguredError(msg)
	}

	omc.typoAdvisor.AddTerm(group)
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

// ModifyGroup allows the configuration of a specific group to be modified.
// If configuration for that group doesn't exist, it will be created.
// While intended for test use, this isn't in a _test.go file as we want to use it from tests in multiple packages.
func (omc *ObjectModelConfiguration) ModifyGroup(
	ref astmodel.PackageReference,
	action func(configuration *GroupConfiguration) error) error {
	group, _, ok := ref.GroupVersion()
	if !ok {
		return errors.Errorf(
			"external package reference %s not supported",
			ref)
	}

	grp, err := omc.findGroup(ref)
	if err != nil && !IsNotConfiguredError(err) {
		return errors.Wrapf(err, "configuring group %s", group)
	}

	if grp == nil {
		grp = NewGroupConfiguration(group)
		omc.add(grp)
	}

	return action(grp)
}

// ModifyVersion allows the configuration of a specific version to be modified.
// If configuration for that version doesn't exist, it will be created.
// While intended for test use, this isn't in a _test.go file as we want to use it from tests in multiple packages.
func (omc *ObjectModelConfiguration) ModifyVersion(
	ref astmodel.PackageReference,
	action func(configuration *VersionConfiguration) error) error {
	_, version, ok := ref.GroupVersion()
	if !ok {
		return errors.Errorf(
			"external package reference %s not supported",
			ref)
	}

	return omc.ModifyGroup(
		ref,
		func(configuration *GroupConfiguration) error {
			ver, err := configuration.findVersion(ref)
			if err != nil && !IsNotConfiguredError(err) {
				return errors.Wrapf(err, "configuring version %s", version)
			}

			if ver == nil {
				ver = NewVersionConfiguration(version)
				configuration.add(ver)
			}

			return action(ver)
		})
}

// ModifyType allows the configuration of a specific type to be modified.
// If configuration for that type doesn't exist, it will be created.
// While intended for test use, this isn't in a _test.go file as we want to use it from tests in multiple packages.
func (omc *ObjectModelConfiguration) ModifyType(
	name astmodel.TypeName,
	action func(typeConfiguration *TypeConfiguration) error) error {
	return omc.ModifyVersion(
		name.PackageReference,
		func(versionConfiguration *VersionConfiguration) error {
			typ, err := versionConfiguration.findType(name.Name())
			if err != nil && !IsNotConfiguredError(err) {
				return errors.Wrapf(err, "configuring type %s", name.Name())
			}

			if typ == nil {
				typ = NewTypeConfiguration(name.Name())
				versionConfiguration.add(typ)
			}

			return action(typ)
		})
}

// ModifyProperty allows the configuration of a specific property to be modified.
// If configuration for that property doesn't exist, it will be created.
// While intended for test use, this isn't in a _test.go file as we want to use it from tests in multiple packages.
func (omc *ObjectModelConfiguration) ModifyProperty(
	typeName astmodel.TypeName,
	property astmodel.PropertyName,
	action func(propertyConfiguration *PropertyConfiguration) error) error {
	return omc.ModifyType(
		typeName,
		func(typeConfiguration *TypeConfiguration) error {
			prop, err := typeConfiguration.findProperty(property)
			if err != nil && !IsNotConfiguredError(err) {
				return errors.Wrapf(err, "configuring property %s", property)
			}

			if prop == nil {
				prop = NewPropertyConfiguration(property.String())
				typeConfiguration.add(prop)
			}

			return action(prop)
		})
}
