/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
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

	// Group access fields here (alphabetical, please)
	PayloadType propertyAccess[PayloadType]

	// Type access fields here (alphabetical, please)
	AzureGeneratedSecrets    typeAccess[[]string]
	DefaultAzureName         typeAccess[bool]
	Export                   typeAccess[bool]
	ExportAs                 typeAccess[string]
	GeneratedConfigs         typeAccess[map[string]string]
	Importable               typeAccess[bool]
	IsResource               typeAccess[bool]
	ManualConfigs            typeAccess[[]string]
	RenameTo                 typeAccess[string]
	ResourceEmbeddedInParent typeAccess[string]
	OperatorSpecProperties   typeAccess[[]OperatorSpecPropertyConfiguration]
	SupportedFrom            typeAccess[string]
	TypeNameInNextVersion    typeAccess[string]

	// Property access fields here (alphabetical, please)
	ARMReference                   propertyAccess[bool]
	ImportConfigMapMode            propertyAccess[ImportConfigMapMode]
	IsSecret                       propertyAccess[bool]
	PropertyNameInNextVersion      propertyAccess[string]
	RenamePropertyTo               propertyAccess[string]
	ResourceLifecycleOwnedByParent propertyAccess[string]
}

// NewObjectModelConfiguration returns a new (empty) ObjectModelConfiguration
func NewObjectModelConfiguration() *ObjectModelConfiguration {
	result := &ObjectModelConfiguration{
		groups:      make(map[string]*GroupConfiguration),
		typoAdvisor: typo.NewAdvisor(),
	}

	// Initialize group access fields here (alphabetical, please)
	// Initialize multi-level access fields here (alphabetical, please)
	result.PayloadType = makeGroupAccess[PayloadType](
		result,
		func(c *GroupConfiguration) *configurable[PayloadType] { return &c.PayloadType },
	).withTypeOverride(
		func(c *TypeConfiguration) *configurable[PayloadType] { return &c.PayloadType },
	).withPropertyOverride(
		func(c *PropertyConfiguration) *configurable[PayloadType] { return &c.PayloadType },
	)

	// Initialize type access fields here (alphabetical, please)
	result.AzureGeneratedSecrets = makeTypeAccess[[]string](
		result, func(c *TypeConfiguration) *configurable[[]string] { return &c.AzureGeneratedSecrets })
	result.DefaultAzureName = makeTypeAccess[bool](
		result, func(c *TypeConfiguration) *configurable[bool] { return &c.DefaultAzureName })
	result.Export = makeTypeAccess[bool](
		result, func(c *TypeConfiguration) *configurable[bool] { return &c.Export })
	result.ExportAs = makeTypeAccess[string](
		result, func(c *TypeConfiguration) *configurable[string] { return &c.ExportAs })
	result.GeneratedConfigs = makeTypeAccess[map[string]string](
		result, func(c *TypeConfiguration) *configurable[map[string]string] { return &c.GeneratedConfigs })
	result.Importable = makeTypeAccess[bool](
		result, func(c *TypeConfiguration) *configurable[bool] { return &c.Importable })
	result.IsResource = makeTypeAccess[bool](
		result, func(c *TypeConfiguration) *configurable[bool] { return &c.IsResource })
	result.ManualConfigs = makeTypeAccess[[]string](
		result, func(c *TypeConfiguration) *configurable[[]string] { return &c.ManualConfigs })
	result.OperatorSpecProperties = makeTypeAccess[[]OperatorSpecPropertyConfiguration](
		result, func(c *TypeConfiguration) *configurable[[]OperatorSpecPropertyConfiguration] {
			return &c.OperatorSpecProperties
		})
	result.RenameTo = makeTypeAccess[string](
		result, func(c *TypeConfiguration) *configurable[string] { return &c.RenameTo })
	result.ResourceEmbeddedInParent = makeTypeAccess[string](
		result, func(c *TypeConfiguration) *configurable[string] { return &c.ResourceEmbeddedInParent })
	result.SupportedFrom = makeTypeAccess[string](
		result, func(c *TypeConfiguration) *configurable[string] { return &c.SupportedFrom })
	result.TypeNameInNextVersion = makeTypeAccess[string](
		result, func(c *TypeConfiguration) *configurable[string] { return &c.NameInNextVersion })

	// Initialize property access fields here (alphabetical, please)
	result.ARMReference = makePropertyAccess[bool](
		result, func(c *PropertyConfiguration) *configurable[bool] { return &c.ARMReference })
	result.ImportConfigMapMode = makePropertyAccess[ImportConfigMapMode](
		result, func(c *PropertyConfiguration) *configurable[ImportConfigMapMode] { return &c.ImportConfigMapMode })
	result.IsSecret = makePropertyAccess[bool](
		result, func(c *PropertyConfiguration) *configurable[bool] { return &c.IsSecret })
	result.PropertyNameInNextVersion = makePropertyAccess[string](
		result, func(c *PropertyConfiguration) *configurable[string] { return &c.NameInNextVersion })
	result.RenamePropertyTo = makePropertyAccess[string](
		result, func(c *PropertyConfiguration) *configurable[string] { return &c.RenameTo })
	result.ResourceLifecycleOwnedByParent = makePropertyAccess[string](
		result, func(c *PropertyConfiguration) *configurable[string] { return &c.ResourceLifecycleOwnedByParent })

	return result
}

// IsEmpty returns true if we have no configuration at all, false if we have some groups configured.
func (omc *ObjectModelConfiguration) IsEmpty() bool {
	return len(omc.groups) == 0
}

// IsGroupConfigured returns true if we have any configuration for the specified group, false otherwise.
func (omc *ObjectModelConfiguration) IsGroupConfigured(pkg astmodel.InternalPackageReference) bool {
	var result bool
	visitor := newSingleGroupConfigurationVisitor(pkg, func(configuration *GroupConfiguration) error {
		result = true
		return nil
	})

	err := visitor.visit(omc)
	if err != nil {
		// For any error, we'll assume we're expecting the group
		return true
	}

	return result
}

// IsTypeConfigured returns true if we have any configuration for the specified type, false otherwise.
func (omc *ObjectModelConfiguration) IsTypeConfigured(name astmodel.InternalTypeName) bool {
	var result bool
	visitor := newSingleTypeConfigurationVisitor(name, func(configuration *TypeConfiguration) error {
		result = true
		return nil
	})

	err := visitor.visit(omc)
	if err != nil {
		// For any error, we'll assume we're expecting the type
		return true
	}

	return result
}

// AddTypeAlias adds a type alias for the specified type name,
// allowing configuration related to the type to be accessed via the new name.
func (omc *ObjectModelConfiguration) AddTypeAlias(name astmodel.InternalTypeName, alias string) {
	versionVisitor := newSingleVersionConfigurationVisitor(
		name.InternalPackageReference(),
		func(configuration *VersionConfiguration) error {
			return configuration.addTypeAlias(name.Name(), alias)
		})

	err := versionVisitor.visit(omc)
	if err != nil {
		// Should never have an error in this case, but if we do make sure we know
		panic(err)
	}
}

var VersionRegex = regexp.MustCompile(`^v\d\d?$`)

// FindHandCraftedTypeNames returns the set of type-names that are hand-crafted.
// These are identified by having `v<n>` as their version.
func (omc *ObjectModelConfiguration) FindHandCraftedTypeNames(localPath string) (astmodel.InternalTypeNameSet, error) {
	result := astmodel.NewInternalTypeNameSet()
	var currentGroup string
	var currentPackage astmodel.InternalPackageReference

	// Collect the names of hand-crafted types
	typeVisitor := newEveryTypeConfigurationVisitor(
		func(typeConfig *TypeConfiguration) error {
			name := astmodel.MakeInternalTypeName(currentPackage, typeConfig.name)
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

	err := groupVisitor.visit(omc)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find hand-crafted packages")
	}

	return result, nil
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
	ref astmodel.InternalPackageReference,
	visitor *configurationVisitor,
) error {
	group := omc.findGroup(ref)
	if group == nil {
		return nil
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
func (omc *ObjectModelConfiguration) findGroup(ref astmodel.InternalPackageReference) *GroupConfiguration {
	group := ref.Group()

	if omc == nil || omc.groups == nil {
		return nil
	}

	omc.typoAdvisor.AddTerm(group)
	if g, ok := omc.groups[group]; ok {
		return g
	}

	return nil
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
	ref astmodel.InternalPackageReference,
	action func(configuration *GroupConfiguration) error,
) error {
	groupName := ref.Group()
	grp := omc.findGroup(ref)
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
	ref astmodel.InternalPackageReference,
	action func(configuration *VersionConfiguration) error,
) error {
	_, version := ref.GroupVersion()
	return omc.ModifyGroup(
		ref,
		func(configuration *GroupConfiguration) error {
			ver := configuration.findVersion(ref)
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
	name astmodel.InternalTypeName,
	action func(typeConfiguration *TypeConfiguration) error,
) error {
	return omc.ModifyVersion(
		name.InternalPackageReference(),
		func(versionConfiguration *VersionConfiguration) error {
			typeName := name.Name()
			typ := versionConfiguration.findType(typeName)
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
	typeName astmodel.InternalTypeName,
	property astmodel.PropertyName,
	action func(propertyConfiguration *PropertyConfiguration) error,
) error {
	return omc.ModifyType(
		typeName,
		func(typeConfiguration *TypeConfiguration) error {
			prop := typeConfiguration.findProperty(property)
			if prop == nil {
				name := property.String()
				prop = NewPropertyConfiguration(name)
				typeConfiguration.addProperty(name, prop)
			}

			return action(prop)
		})
}
