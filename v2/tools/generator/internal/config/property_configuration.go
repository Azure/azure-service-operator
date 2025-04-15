/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"strings"

	"github.com/rotisserie/eris"
	"gopkg.in/yaml.v3"
)

// PropertyConfiguration contains additional information about a specific property and forms part of a hierarchy
// containing information to supplement the schema and swagger sources consumed by the generator.
//
// ┌──────────────────────────┐       ┌────────────────────┐       ┌──────────────────────┐       ┌───────────────────┐       ╔═══════════════════════╗
// │                          │       │                    │       │                      │       │                   │       ║                       ║
// │ ObjectModelConfiguration │───────│ GroupConfiguration │───────│ VersionConfiguration │───────│ TypeConfiguration │───────║ PropertyConfiguration ║
// │                          │1  1..n│                    │1  1..n│                      │1  1..n│                   │1  1..n║                       ║
// └──────────────────────────┘       └────────────────────┘       └──────────────────────┘       └───────────────────┘       ╚═══════════════════════╝
type PropertyConfiguration struct {
	name string
	// Configurable properties here (alphabetical, please)
	Description                    configurable[string]              // Specify a description override for this property
	ImportConfigMapMode            configurable[ImportConfigMapMode] // The config map mode
	IsSecret                       configurable[bool]                // Specify whether this property is a secret
	NameInNextVersion              configurable[string]              // Name this property has in the next version
	PayloadType                    configurable[PayloadType]         // Specify how this property should be serialized for ARM
	ReferenceType                  configurable[ReferenceType]       // Specify whether this property is an ARM reference or some other kind
	RenameTo                       configurable[string]              // Name this property should be renamed to
	ResourceLifecycleOwnedByParent configurable[string]              // Name of the parent resource which owns the lifecycle of the sub-resource.
}

type ImportConfigMapMode string

const (
	ImportConfigMapModeOptional = "optional"
	ImportConfigMapModeRequired = "required"
)

type ReferenceType string

const (
	ReferenceTypeARM    = ReferenceType("arm")    // An ARM reference
	ReferenceTypeSimple = ReferenceType("simple") // A simple reference requiring no special handling
)

// Tags used in yaml files to specify configurable properties. Alphabetical please.
const (
	descriptionTag                    = "$description"                    // String overriding the properties default description
	exportAsConfigMapPropertyNameTag  = "$exportAsConfigMapPropertyName"  // String specifying the name of the property set to export this property as a config map.
	importConfigMapModeTag            = "$importConfigMapMode"            // string specifying the ImportConfigMapMode mode
	isSecretTag                       = "$isSecret"                       // Bool specifying whether a property contains a secret
	referenceTypeTag                  = "$referenceType"                  // String specifying what kind of reference we have
	renamePropertyToTag               = "$renameTo"                       // String specifying the name this property should be renamed to
	resourceLifecycleOwnedByParentTag = "$resourceLifecycleOwnedByParent" // String specifying whether a property represents a subresource whose lifecycle is owned by the parent resource (and what that parent resource is)
)

// NewPropertyConfiguration returns a new (empty) property configuration
func NewPropertyConfiguration(name string) *PropertyConfiguration {
	scope := "property " + name
	return &PropertyConfiguration{
		name: name,
		// Initialize configurable properties here (alphabetical, please)
		Description:                    makeConfigurable[string](descriptionTag, scope),
		ImportConfigMapMode:            makeConfigurable[ImportConfigMapMode](importConfigMapModeTag, scope),
		IsSecret:                       makeConfigurable[bool](isSecretTag, scope),
		NameInNextVersion:              makeConfigurable[string](nameInNextVersionTag, scope),
		ReferenceType:                  makeConfigurable[ReferenceType](referenceTypeTag, scope),
		RenameTo:                       makeConfigurable[string](renamePropertyToTag, scope),
		ResourceLifecycleOwnedByParent: makeConfigurable[string](resourceLifecycleOwnedByParentTag, scope),
	}
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (pc *PropertyConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return eris.New("expected mapping")
	}

	var lastID string
	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastID = strings.ToLower(c.Value)
			continue
		}

		// $nameInNextVersion: <string>
		if strings.EqualFold(lastID, nameInNextVersionTag) && c.Kind == yaml.ScalarNode {
			pc.NameInNextVersion.Set(c.Value)
			continue
		}

		// $isSecret: <bool>
		if strings.EqualFold(lastID, isSecretTag) && c.Kind == yaml.ScalarNode {
			var isSecret bool
			err := c.Decode(&isSecret)
			if err != nil {
				return eris.Wrapf(err, "decoding %s", isSecretTag)
			}

			pc.IsSecret.Set(isSecret)
			continue
		}

		// $resourceLifecycleOwnedByParent: string
		if strings.EqualFold(lastID, resourceLifecycleOwnedByParentTag) && c.Kind == yaml.ScalarNode {
			var resourceLifecycleOwnedByParent string
			err := c.Decode(&resourceLifecycleOwnedByParent)
			if err != nil {
				return eris.Wrapf(err, "decoding %s", resourceLifecycleOwnedByParentTag)
			}

			pc.ResourceLifecycleOwnedByParent.Set(resourceLifecycleOwnedByParent)
			continue
		}

		// $referenceType: <string>
		if strings.EqualFold(lastID, referenceTypeTag) && c.Kind == yaml.ScalarNode {
			switch strings.ToLower(c.Value) {
			case string(ReferenceTypeARM):
				pc.ReferenceType.Set(ReferenceTypeARM)
			case string(ReferenceTypeSimple):
				pc.ReferenceType.Set(ReferenceTypeSimple)
			default:
				return eris.Errorf("unknown %s value: %s.", referenceTypeTag, c.Value)
			}

			continue
		}

		// $ImportConfigMapMode: <string>
		if strings.EqualFold(lastID, importConfigMapModeTag) && c.Kind == yaml.ScalarNode {
			switch strings.ToLower(c.Value) {
			case ImportConfigMapModeOptional:
				pc.ImportConfigMapMode.Set(ImportConfigMapModeOptional)
			case ImportConfigMapModeRequired:
				pc.ImportConfigMapMode.Set(ImportConfigMapModeRequired)
			default:
				return eris.Errorf("unknown %s value: %s.", importConfigMapModeTag, c.Value)
			}

			continue
		}

		// renameTo: string
		if strings.EqualFold(lastID, renamePropertyToTag) && c.Kind == yaml.ScalarNode {
			var renameTo string
			err := c.Decode(&renameTo)
			if err != nil {
				return eris.Wrapf(err, "decoding %s", renamePropertyToTag)
			}

			pc.RenameTo.Set(renameTo)
			continue
		}

		// $payloadType: <string>
		if strings.EqualFold(lastID, payloadTypeTag) && c.Kind == yaml.ScalarNode {
			switch strings.ToLower(c.Value) {
			case string(OmitEmptyProperties):
				pc.PayloadType.Set(OmitEmptyProperties)
			case string(ExplicitCollections):
				pc.PayloadType.Set(ExplicitCollections)
			case string(ExplicitEmptyCollections):
				pc.PayloadType.Set(ExplicitEmptyCollections)
			case string(ExplicitProperties):
				pc.PayloadType.Set(ExplicitProperties)
			default:
				return eris.Errorf("unknown %s value: %s.", payloadTypeTag, c.Value)
			}

			continue
		}

		// description: string
		if strings.EqualFold(lastID, descriptionTag) && c.Kind == yaml.ScalarNode {
			var description string
			err := c.Decode(&description)
			if err != nil {
				return eris.Wrapf(err, "decoding %s", descriptionTag)
			}

			pc.Description.Set(description)
			continue
		}

		// No handler for this value, return an error
		return eris.Errorf(
			"property configuration, unexpected yaml value %s: %s (line %d col %d)", lastID, c.Value, c.Line, c.Column)

	}

	return nil
}
