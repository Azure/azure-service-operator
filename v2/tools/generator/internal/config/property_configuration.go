/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"strings"

	"github.com/pkg/errors"
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
	ARMReference                   configurable[bool]                // Specify whether this property is an ARM reference
	ImportConfigMapMode            configurable[ImportConfigMapMode] // The config map mode
	IsSecret                       configurable[bool]                // Specify whether this property is a secret
	NameInNextVersion              configurable[string]              // Name this property has in the next version
	ResourceLifecycleOwnedByParent configurable[string]
}

type ImportConfigMapMode string

const (
	ImportConfigMapModeOptional = "optional"
	ImportConfigMapModeRequired = "required"
)

const (
	armReferenceTag                   = "$armReference"                   // Bool specifying whether a property is an ARM reference
	isSecretTag                       = "$isSecret"                       // Bool specifying whether a property contains a secret
	resourceLifecycleOwnedByParentTag = "$resourceLifecycleOwnedByParent" // String specifying whether a property represents a subresource whose lifecycle is owned by the parent resource (and what that parent resource is)
	exportAsConfigMapPropertyNameTag  = "$exportAsConfigMapPropertyName"  // String specifying the name of the property set to export this property as a config map.
	importConfigMapModeTag            = "$importConfigMapMode"            // string specifying the ImportConfigMapMode mode
)

// NewPropertyConfiguration returns a new (empty) property configuration
func NewPropertyConfiguration(name string) *PropertyConfiguration {
	scope := "property " + name
	return &PropertyConfiguration{
		name: name,
		// Initialize configurable properties here (alphabetical, please)
		ARMReference:                   makeConfigurable[bool](armReferenceTag, scope),
		ImportConfigMapMode:            makeConfigurable[ImportConfigMapMode](importConfigMapModeTag, scope),
		IsSecret:                       makeConfigurable[bool](isSecretTag, scope),
		NameInNextVersion:              makeConfigurable[string](nameInNextVersionTag, scope),
		ResourceLifecycleOwnedByParent: makeConfigurable[string](resourceLifecycleOwnedByParentTag, scope),
	}
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (pc *PropertyConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return errors.New("expected mapping")
	}

	var lastId string
	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastId = strings.ToLower(c.Value)
			continue
		}

		// $nameInNextVersion: <string>
		if strings.EqualFold(lastId, nameInNextVersionTag) && c.Kind == yaml.ScalarNode {
			pc.NameInNextVersion.Set(c.Value)
			continue
		}

		// $isSecret: <bool>
		if strings.EqualFold(lastId, isSecretTag) && c.Kind == yaml.ScalarNode {
			var isSecret bool
			err := c.Decode(&isSecret)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", isSecretTag)
			}

			pc.IsSecret.Set(isSecret)
			continue
		}

		// $resourceLifecycleOwnedByParent: string
		if strings.EqualFold(lastId, resourceLifecycleOwnedByParentTag) && c.Kind == yaml.ScalarNode {
			var resourceLifecycleOwnedByParent string
			err := c.Decode(&resourceLifecycleOwnedByParent)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", resourceLifecycleOwnedByParentTag)
			}

			pc.ResourceLifecycleOwnedByParent.Set(resourceLifecycleOwnedByParent)
			continue
		}

		// $armReference: <bool>
		if strings.EqualFold(lastId, armReferenceTag) && c.Kind == yaml.ScalarNode {
			var isARMRef bool
			err := c.Decode(&isARMRef)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", armReferenceTag)
			}

			pc.ARMReference.Set(isARMRef)
			continue
		}

		// $ImportConfigMapMode: <string>
		if strings.EqualFold(lastId, importConfigMapModeTag) && c.Kind == yaml.ScalarNode {
			switch strings.ToLower(c.Value) {
			case ImportConfigMapModeOptional:
				pc.ImportConfigMapMode.Set(ImportConfigMapModeOptional)
			case ImportConfigMapModeRequired:
				pc.ImportConfigMapMode.Set(ImportConfigMapModeRequired)
			default:
				return errors.Errorf("unknown %s value: %s.", importConfigMapModeTag, c.Value)
			}

			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"property configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}
