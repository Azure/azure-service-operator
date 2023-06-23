/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/util/typo"
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
type ObjectModelConfiguration struct {
	groups      map[string]*GroupConfiguration // nested configuration for individual groups
	typoAdvisor *typo.Advisor
}

// NewObjectModelConfiguration returns a new (empty) ObjectModelConfiguration
func NewObjectModelConfiguration() *ObjectModelConfiguration {
	return &ObjectModelConfiguration{
		groups:      make(map[string]*GroupConfiguration),
		typoAdvisor: typo.NewAdvisor(),
	}
}

// IsEmpty returns true if we have no configuration at all, false if we have some groups configured.
func (omc *ObjectModelConfiguration) IsEmpty() bool {
	return len(omc.groups) == 0
}

// IsGroupConfigured returns true if we have any configuration for the specified group, false otherwise.
func (omc *ObjectModelConfiguration) IsGroupConfigured(pkg astmodel.PackageReference) bool {
	var result bool
	visitor := newSingleGroupConfigurationVisitor(pkg, func(configuration *GroupConfiguration) error {
		result = true
		return nil
	})

	err := visitor.Visit(omc)
	if err != nil {
		if IsNotConfiguredError(err) {
			// No configuration for this package, we're not expecting any types
			return false
		}

		// Some other error, we'll assume we're expecting types
		return true
	}

	return result
}

// LookupNameInNextVersion checks whether we have an alternative name for the specified type, returning the name if
// found. Returns a NotConfiguredError if no rename is available.
func (omc *ObjectModelConfiguration) LookupNameInNextVersion(name astmodel.TypeName) (string, error) {
	var newName string
	visitor := newSingleTypeConfigurationVisitor(
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
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyNameInNextVersionConsumed()
		})
	return visitor.Visit(omc)
}

// LookupExport checks to see whether a specified type is configured for export, returning the value if found. Returns a
// NotConfiguredError if no export is configured.
func (omc *ObjectModelConfiguration) LookupExport(name astmodel.TypeName) (bool, error) {
	var export bool
	visitor := newSingleTypeConfigurationVisitor(
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
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyExportConsumed()
		})
	return visitor.Visit(omc)
}

// LookupExportAs checks to see whether a specified type is configured for export with an alternative name, returning the
// name if found. Returns a NotConfiguredError if no export is configured.
func (omc *ObjectModelConfiguration) LookupExportAs(name astmodel.TypeName) (string, error) {
	var exportAs string
	typeVisitor := newSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			ea, err := configuration.LookupExportAs()
			exportAs = ea
			return err
		})
	err := typeVisitor.Visit(omc)
	if err != nil {
		// No need to wrap this error, it already has all the details
		return "", err
	}

	// Add an alias so that any existing configuration can be found via the new name
	versionVisitor := newSingleVersionConfigurationVisitor(
		name.PackageReference,
		func(configuration *VersionConfiguration) error {
			return configuration.addTypeAlias(name.Name(), exportAs)
		})
	err = versionVisitor.Visit(omc)
	if err != nil {
		// No need to wrap this error, it already has all the details
		return "", err
	}

	return exportAs, nil
}

// VerifyExportAsConsumed returns an error if our configured export name was not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifyExportAsConsumed() error {
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyExportAsConsumed()
		})
	return visitor.Visit(omc)
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
// Returns true or false if configured, or a NotConfiguredError if not.
func (omc *ObjectModelConfiguration) ARMReference(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	var result bool
	visitor := newSinglePropertyConfigurationVisitor(
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
	visitor := newEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			return configuration.VerifyARMReferenceConsumed()
		})
	return visitor.Visit(omc)
}

// AzureGeneratedSecrets looks up a type to determine if it has any Azure generated secrets
func (omc *ObjectModelConfiguration) AzureGeneratedSecrets(name astmodel.TypeName) ([]string, error) {
	var result []string
	visitor := newSingleTypeConfigurationVisitor(
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
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyAzureGeneratedSecretsConsumed()
		})
	return visitor.Visit(omc)
}

// GeneratedConfigs looks up a type to determine if it has any generated configs
func (omc *ObjectModelConfiguration) GeneratedConfigs(name astmodel.TypeName) (map[string]string, error) {
	var result map[string]string
	visitor := newSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			var err error
			result, err = configuration.GeneratedConfigs()
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// VerifyGeneratedConfigsConsumed returns an error if generated configs were not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifyGeneratedConfigsConsumed() error {
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyGeneratedConfigsConsumed()
		})
	return visitor.Visit(omc)
}

// ManualConfigs looks up a type to determine if it has any manual configs
func (omc *ObjectModelConfiguration) ManualConfigs(name astmodel.TypeName) ([]string, error) {
	var result []string
	visitor := newSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			var err error
			result, err = configuration.ManualConfigs()
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// VerifyManualConfigsConsumed returns an error if manual configs were not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifyManualConfigsConsumed() error {
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyManualConfigsConsumed()
		})
	return visitor.Visit(omc)
}

// IsSecret looks up a property to determine whether it is a secret.
func (omc *ObjectModelConfiguration) IsSecret(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	var result bool
	visitor := newSinglePropertyConfigurationVisitor(
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
	visitor := newEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			return configuration.VerifyIsSecretConsumed()
		})
	return visitor.Visit(omc)
}

// ResourceLifecycleOwnedByParent returns the name of the parent resource if the property represents a subresource whose resource lifecycle is owned by the parent resource.
// An empty string + error is returned if this is not configured for the property in question
func (omc *ObjectModelConfiguration) ResourceLifecycleOwnedByParent(name astmodel.TypeName, property astmodel.PropertyName) (string, error) {
	var result string
	visitor := newSinglePropertyConfigurationVisitor(
		name,
		property,
		func(configuration *PropertyConfiguration) error {
			resourceLifecycleOwnedByParent, err := configuration.ResourceLifecycleOwnedByParent()
			result = resourceLifecycleOwnedByParent
			return err
		})

	err := visitor.Visit(omc)
	if err != nil {
		return "", err
	}

	return result, nil
}

// MarkResourceLifecycleOwnedByParentUnconsumed marks all ResourceLifecycleOwnedByParent as unconsumed
func (omc *ObjectModelConfiguration) MarkResourceLifecycleOwnedByParentUnconsumed() error {
	visitor := newEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			configuration.ClearResourceLifecycleOwnedByParentConsumed()
			return nil
		})
	return visitor.Visit(omc)
}

// VerifyResourceLifecycleOwnedByParentConsumed returns an error if any ResourceLifecycleOwnedByParent configuration
// was not consumed
func (omc *ObjectModelConfiguration) VerifyResourceLifecycleOwnedByParentConsumed() error {
	visitor := newEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			return configuration.VerifyResourceLifecycleOwnedByParentConsumed()
		})
	return visitor.Visit(omc)
}

// LookupSupportedFrom checks to see whether a specified type has its first ASO release configured, returning either
// that release or a NotConfiguredError.
func (omc *ObjectModelConfiguration) LookupSupportedFrom(name astmodel.TypeName) (string, error) {
	var result string
	visitor := newSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			var err error
			result, err = configuration.LookupSupportedFrom()
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return "", err
	}

	return result, nil
}

// VerifySupportedFromConsumed returns an error if any configured supportedFrom tag was not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifySupportedFromConsumed() error {
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifySupportedFromConsumed()
		})
	return visitor.Visit(omc)
}

// ImportConfigMapMode looks up a property to determine its ImportConfigMapMode.
// Returns the ImportConfigMapMode, or a NotConfiguredError if not configured.
func (omc *ObjectModelConfiguration) ImportConfigMapMode(name astmodel.TypeName, property astmodel.PropertyName) (ImportConfigMapMode, error) {
	var result ImportConfigMapMode
	visitor := newSinglePropertyConfigurationVisitor(
		name,
		property,
		func(configuration *PropertyConfiguration) error {
			mode, err := configuration.ImportConfigMapMode()
			result = mode
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return "", err
	}

	return result, nil
}

// VerifyImportConfigMapModeConsumed returns an error if any ImportConfigMapMode configuration was not consumed
func (omc *ObjectModelConfiguration) VerifyImportConfigMapModeConsumed() error {
	visitor := newEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			return configuration.VerifyImportConfigMapModeConsumed()
		})
	return visitor.Visit(omc)
}

// LookupResourceEmbeddedInParent checks to see whether a specified type is labelled as a resource that is embedded
// inside its parent. Returns a NotConfiguredError if no $resourceEmbeddedInParent flag is configured.
func (omc *ObjectModelConfiguration) LookupResourceEmbeddedInParent(name astmodel.TypeName) (string, error) {
	var embeddedInParent string
	visitor := newSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			embedded, err := configuration.LookupResourceEmbeddedInParent()
			embeddedInParent = embedded
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return "", err
	}

	return embeddedInParent, nil
}

// VerifyResourceEmbeddedInParentConsumed returns an error if our configured $resourceEmbeddedInParent flag was not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifyResourceEmbeddedInParentConsumed() error {
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyResourceEmbeddedInParentConsumed()
		})
	return visitor.Visit(omc)
}

// LookupIsResource checks to see whether a specified type is labelled as a resource.
// Returns a NotConfiguredError if no $resourceEmbeddedInParent flag is configured.
func (omc *ObjectModelConfiguration) LookupIsResource(name astmodel.TypeName) (bool, error) {
	var isResource bool
	visitor := newSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			is, err := configuration.LookupIsResource()
			isResource = is
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return false, err
	}

	return isResource, nil
}

// VerifyIsResourceConsumed returns an error if our configured $isResource flag was not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifyIsResourceConsumed() error {
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyIsResourceConsumed()
		})
	return visitor.Visit(omc)
}

// LookupImportable checks to see whether a specified type is labelled as importable.
// Returns a NotConfiguredError if no $importable flag is configured.
func (omc *ObjectModelConfiguration) LookupImportable(name astmodel.TypeName) (bool, error) {
	var importable bool
	visitor := newSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			im, err := configuration.LookupImportable()
			importable = im
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return false, err
	}

	return importable, nil
}

// VerifyImportableConsumed returns an error if our configured $importable flag was not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifyImportableConsumed() error {
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyImportableConsumed()
		})
	return visitor.Visit(omc)
}

// LookupDefaultAzureName checks to see whether a specified type should default the Azure Name property.
// Returns a NotConfiguredError if no $defaultAzureName flag is configured.
func (omc *ObjectModelConfiguration) LookupDefaultAzureName(name astmodel.TypeName) (bool, error) {
	var defaultAzureName bool
	visitor := newSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			defAzure, err := configuration.LookupDefaultAzureName()
			defaultAzureName = defAzure
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return false, err
	}

	return defaultAzureName, nil
}

// VerifyDefaultAzureNameConsumed returns an error if our configured $defaultAzureName flag was not used, nil otherwise.
func (omc *ObjectModelConfiguration) VerifyDefaultAzureNameConsumed() error {
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			return configuration.VerifyDefaultAzureNameConsumed()
		})
	return visitor.Visit(omc)
}

var VersionRegex = regexp.MustCompile(`^v\d\d?$`)

// FindHandCraftedTypeNames returns the set of typenames that are hand-crafted.
// These are identified by having `v<n>` as their version.
func (omc *ObjectModelConfiguration) FindHandCraftedTypeNames(localPath string) (astmodel.TypeNameSet, error) {
	result := make(astmodel.TypeNameSet)
	var currentGroup string
	var currentPackage astmodel.PackageReference

	// Collect the names of hand-crafted types
	typeVisitor := newEveryTypeConfigurationVisitor(
		func(typeConfig *TypeConfiguration) error {
			name := astmodel.MakeTypeName(currentPackage, typeConfig.name)
			result.Add(name)
			return nil
		})

	// Collect hand-crafted versions as we see them.
	// They look like v<n> where n is a small number.
	versionVisitor := newEveryVersionConfigurationVisitor(
		func(verConfig *VersionConfiguration) error {
			if VersionRegex.MatchString(verConfig.name) {
				currentPackage = astmodel.MakeLocalPackageReference(
					localPath,
					currentGroup,
					"", // no prefix needed (or wanted!) for v1
					verConfig.name)
				return verConfig.visitTypes(typeVisitor)
			}

			return nil
		})

	// Look inside each group for hand-crafted versions
	groupVisitor := newEveryGroupConfigurationVisitor(
		func(groupConfig *GroupConfiguration) error {
			currentGroup = groupConfig.name
			return groupConfig.visitVersions(versionVisitor)
		})

	err := groupVisitor.Visit(omc)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find hand-crafted packages")
	}

	return result, nil
}

// LookupPayloadType checks to see whether a specified type has a configured payload type
func (omc *ObjectModelConfiguration) LookupPayloadType(name astmodel.TypeName) (PayloadType, error) {
	var payloadType PayloadType
	visitor := newSingleGroupConfigurationVisitor(
		name.PackageReference,
		func(configuration *GroupConfiguration) error {
			pt, err := configuration.LookupPayloadType()
			payloadType = pt
			return err
		})
	err := visitor.Visit(omc)
	if err != nil {
		return "", err
	}

	return payloadType, nil
}

func (omc *ObjectModelConfiguration) VerifyPayloadTypeConsumed() error {
	visitor := newEveryGroupConfigurationVisitor(
		func(configuration *GroupConfiguration) error {
			return configuration.VerifyPayloadTypeConsumed()
		})
	return visitor.Visit(omc)
}

// addGroup includes the provided GroupConfiguration in this model configuration
func (omc *ObjectModelConfiguration) addGroup(name string, group *GroupConfiguration) {
	if omc.groups == nil {
		// Initialize the map just-in-time
		omc.groups = make(map[string]*GroupConfiguration)
	}

	// store the group name using lowercase,
	// so we can do case-insensitive lookups later
	omc.groups[strings.ToLower(name)] = group
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
	errs := make([]error, 0, len(omc.groups))
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
	group, _ := ref.GroupVersion()

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

			omc.addGroup(lastId, g)
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
	result := make([]string, 0, len(omc.groups))
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
	action func(configuration *GroupConfiguration) error,
) error {
	groupName, _ := ref.GroupVersion()
	grp, err := omc.findGroup(ref)
	if err != nil && !IsNotConfiguredError(err) {
		return errors.Wrapf(err, "configuring groupName %s", groupName)
	}

	if grp == nil {
		grp = NewGroupConfiguration(groupName)
		omc.addGroup(groupName, grp)
	}

	return action(grp)
}

// ModifyVersion allows the configuration of a specific version to be modified.
// If configuration for that version doesn't exist, it will be created.
// While intended for test use, this isn't in a _test.go file as we want to use it from tests in multiple packages.
func (omc *ObjectModelConfiguration) ModifyVersion(
	ref astmodel.PackageReference,
	action func(configuration *VersionConfiguration) error,
) error {
	_, version := ref.GroupVersion()
	return omc.ModifyGroup(
		ref,
		func(configuration *GroupConfiguration) error {
			ver, err := configuration.findVersion(ref)
			if err != nil && !IsNotConfiguredError(err) {
				return errors.Wrapf(err, "configuring version %s", version)
			}

			if ver == nil {
				ver = NewVersionConfiguration(version)
				configuration.addVersion(version, ver)
			}

			return action(ver)
		})
}

// ModifyType allows the configuration of a specific type to be modified.
// If configuration for that type doesn't exist, it will be created.
// While intended for test use, this isn't in a _test.go file as we want to use it from tests in multiple packages.
func (omc *ObjectModelConfiguration) ModifyType(
	name astmodel.TypeName,
	action func(typeConfiguration *TypeConfiguration) error,
) error {
	return omc.ModifyVersion(
		name.PackageReference,
		func(versionConfiguration *VersionConfiguration) error {
			typeName := name.Name()
			typ, err := versionConfiguration.findType(typeName)
			if err != nil && !IsNotConfiguredError(err) {
				return errors.Wrapf(err, "configuring type %s", typeName)
			}

			if typ == nil {
				typ = NewTypeConfiguration(typeName)
				versionConfiguration.addType(typeName, typ)
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
	action func(propertyConfiguration *PropertyConfiguration) error,
) error {
	return omc.ModifyType(
		typeName,
		func(typeConfiguration *TypeConfiguration) error {
			prop, err := typeConfiguration.findProperty(property)
			if err != nil && !IsNotConfiguredError(err) {
				return errors.Wrapf(err, "configuring property %s", property)
			}

			if prop == nil {
				name := property.String()
				prop = NewPropertyConfiguration(name)
				typeConfiguration.addProperty(name, prop)
			}

			return action(prop)
		})
}
