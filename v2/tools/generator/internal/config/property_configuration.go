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
	name                           string
	nameInNextVersion              configurable[string] // Name this property has in the next version
	armReference                   configurable[bool]   // Specify whether this property is an ARM reference
	isSecret                       configurable[bool]   // Specify whether this property is a secret
	resourceLifecycleOwnedByParent configurable[string]
	exportAsConfigMapPropertyName  configurable[string]              // The name of the exportAsConfigMap property.
	importConfigMapMode            configurable[ImportConfigMapMode] // The config map mode
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
	importConfigMapModeTag            = "$importConfigMapMode"            // string specifying the importConfigMapMode mode
)

// NewPropertyConfiguration returns a new (empty) property configuration
func NewPropertyConfiguration(name string) *PropertyConfiguration {
	return &PropertyConfiguration{
		name: name,
	}
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (pc *PropertyConfiguration) ARMReference() (bool, error) {
	armReference, ok := pc.armReference.read()
	if !ok {
		msg := fmt.Sprintf(armReferenceTag+" not specified for property %s", pc.name)
		return false, NewNotConfiguredError(msg)
	}

	return armReference, nil
}

// VerifyARMReferenceConsumed returns an error if our configuration as an ARM reference was not consumed.
func (pc *PropertyConfiguration) VerifyARMReferenceConsumed() error {
	if pc.armReference.isUnconsumed() {
		v, _ := pc.armReference.read()
		return errors.Errorf("property %s: "+armReferenceTag+": %t not consumed", pc.name, v)
	}

	return nil
}

// SetARMReference sets the ARM reference property
func (pc *PropertyConfiguration) SetARMReference(value bool) {
	pc.armReference.write(value)
}

// IsSecret looks up a property to determine if it's a secret
func (pc *PropertyConfiguration) IsSecret() (bool, error) {
	isSecret, ok := pc.isSecret.read()
	if !ok {
		msg := fmt.Sprintf(isSecretTag+" not specified for property %s", pc.name)
		return false, NewNotConfiguredError(msg)
	}

	return isSecret, nil
}

// VerifyIsSecretConsumed returns an error if our configuration as a secret was not consumed.
func (pc *PropertyConfiguration) VerifyIsSecretConsumed() error {
	if pc.isSecret.isUnconsumed() {
		v, _ := pc.isSecret.read()
		return errors.Errorf("property %s: "+isSecretTag+": %t not consumed", pc.name, v)
	}

	return nil
}

// NameInNextVersion looks up a property to determine whether it is being renamed in the next version
func (pc *PropertyConfiguration) NameInNextVersion() (string, error) {
	name, ok := pc.nameInNextVersion.read()
	if !ok {
		msg := fmt.Sprintf(nameInNextVersionTag+" not specified for property %s", pc.name)
		return "", NewNotConfiguredError(msg)
	}

	return name, nil
}

func (pc *PropertyConfiguration) VerifyRenamedInNextVersionConsumed() error {
	if pc.nameInNextVersion.isUnconsumed() {
		v, _ := pc.nameInNextVersion.read()
		return errors.Errorf("property %s: "+nameInNextVersionTag+": %s not consumed", pc.name, v)
	}

	return nil
}

// ResourceLifecycleOwnedByParent looks up a property to determine what parent owns this resource lifecycle
func (pc *PropertyConfiguration) ResourceLifecycleOwnedByParent() (string, error) {
	resourceLifecycleOwnedByParent, ok := pc.resourceLifecycleOwnedByParent.read()
	if !ok {
		msg := fmt.Sprintf(resourceLifecycleOwnedByParentTag+" not specified for property %s", pc.name)
		return "", NewNotConfiguredError(msg)
	}

	return resourceLifecycleOwnedByParent, nil
}

// ClearResourceLifecycleOwnedByParentConsumed clears the consumed bit for this flag so that it can be reused
func (pc *PropertyConfiguration) ClearResourceLifecycleOwnedByParentConsumed() {
	pc.resourceLifecycleOwnedByParent.markUnconsumed()
}

// VerifyResourceLifecycleOwnedByParentConsumed returns an error if our configuration has the
// $resourceLifecycleOwnedByParent flag, and was not consumed.
func (pc *PropertyConfiguration) VerifyResourceLifecycleOwnedByParentConsumed() error {
	if pc.resourceLifecycleOwnedByParent.isUnconsumed() {
		v, _ := pc.resourceLifecycleOwnedByParent.read()
		return errors.Errorf("property %s: "+resourceLifecycleOwnedByParentTag+": %s not consumed", pc.name, v)
	}

	return nil
}

// ExportAsConfigMapPropertyName looks up a property to determine if it should support being exported to a configMap
func (pc *PropertyConfiguration) ExportAsConfigMapPropertyName() (string, error) {
	val, ok := pc.exportAsConfigMapPropertyName.read()
	if !ok {
		msg := fmt.Sprintf(exportAsConfigMapPropertyNameTag+" not specified for property %s", pc.name)
		return "", NewNotConfiguredError(msg)
	}

	return val, nil
}

// VerifyExportAsConfigMapPropertyNameConsumed returns an error if the config has the exportAsConfigMapPropertyName flag set and
// it was not consumed
func (pc *PropertyConfiguration) VerifyExportAsConfigMapPropertyNameConsumed() error {
	if pc.exportAsConfigMapPropertyName.isUnconsumed() {
		v, _ := pc.exportAsConfigMapPropertyName.read()
		return errors.Errorf("property %s: "+exportAsConfigMapPropertyNameTag+": %s not consumed", pc.name, v)
	}

	return nil
}

// SetExportAsConfigMapPropertyName sets the configmap property name of this property
func (pc *PropertyConfiguration) SetExportAsConfigMapPropertyName(name string) *PropertyConfiguration {
	pc.exportAsConfigMapPropertyName.write(name)
	return pc
}

// SetImportConfigMapMode sets the import configMap mode
func (pc *PropertyConfiguration) SetImportConfigMapMode(mode ImportConfigMapMode) *PropertyConfiguration {
	pc.importConfigMapMode.write(mode)
	return pc
}

// ImportConfigMapMode looks up a property to determine its ImportConfigMapMode
func (pc *PropertyConfiguration) ImportConfigMapMode() (ImportConfigMapMode, error) {
	mode, ok := pc.importConfigMapMode.read()
	if !ok {
		msg := fmt.Sprintf(importConfigMapModeTag+" not specified for property %s", pc.name)
		return mode, NewNotConfiguredError(msg)
	}

	return mode, nil
}

// VerifyImportConfigMapModeConsumed returns an error if our configuration had the importConfigMapMode set and was not consumed.
func (pc *PropertyConfiguration) VerifyImportConfigMapModeConsumed() error {
	if pc.importConfigMapMode.isUnconsumed() {
		v, _ := pc.importConfigMapMode.read()
		return errors.Errorf("property %s: "+importConfigMapModeTag+": %s not consumed", pc.name, v)
	}

	return nil
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
			pc.nameInNextVersion.write(c.Value)
			continue
		}

		// $isSecret: <bool>
		if strings.EqualFold(lastId, isSecretTag) && c.Kind == yaml.ScalarNode {
			var isSecret bool
			err := c.Decode(&isSecret)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", isSecretTag)
			}

			pc.isSecret.write(isSecret)
			continue
		}

		// $resourceLifecycleOwnedByParent: string
		if strings.EqualFold(lastId, resourceLifecycleOwnedByParentTag) && c.Kind == yaml.ScalarNode {
			var resourceLifecycleOwnedByParent string
			err := c.Decode(&resourceLifecycleOwnedByParent)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", resourceLifecycleOwnedByParentTag)
			}

			pc.resourceLifecycleOwnedByParent.write(resourceLifecycleOwnedByParent)
			continue
		}

		// $armReference: <bool>
		if strings.EqualFold(lastId, armReferenceTag) && c.Kind == yaml.ScalarNode {
			var isARMRef bool
			err := c.Decode(&isARMRef)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", armReferenceTag)
			}

			pc.armReference.write(isARMRef)
			continue
		}

		// $exportAsConfigMapPropertyName: <string>
		if strings.EqualFold(lastId, exportAsConfigMapPropertyNameTag) && c.Kind == yaml.ScalarNode {
			var exportAsConfigMapPropertyName string
			err := c.Decode(&exportAsConfigMapPropertyName)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", exportAsConfigMapPropertyNameTag)
			}

			pc.SetExportAsConfigMapPropertyName(exportAsConfigMapPropertyName)
			continue
		}

		// $importConfigMapMode: <string>
		if strings.EqualFold(lastId, importConfigMapModeTag) && c.Kind == yaml.ScalarNode {
			switch c.Value {
			case ImportConfigMapModeOptional:
				pc.importConfigMapMode.write(ImportConfigMapModeOptional)
			case ImportConfigMapModeRequired:
				pc.importConfigMapMode.write(ImportConfigMapModeRequired)
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
