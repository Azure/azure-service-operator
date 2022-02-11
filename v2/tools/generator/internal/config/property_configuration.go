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
//
type PropertyConfiguration struct {
	name              string
	nameInNextVersion configurableString // Name this property has in the next version
	armReference      configurableBool   // Specify whether this property is an ARM reference
	isSecret          configurableBool   // Specify whether this property is a secret
}

const (
	armReferenceTag      = "$armReference"      // Bool specifying whether a property is an ARM reference
	exportTag            = "$export"            // Boolean specifying whether a resource type is exported
	exportAsTag          = "$exportAs"          // String specifying the name to use for a type (implies $export: true)
	nameInNextVersionTag = "$nameInNextVersion" // String specifying a type or property name change in the next version
	isSecretTag          = "$isSecret"          // Bool specifying whether a property contains a secret
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

// IsSecret looks up a property to determine if it's a secret
func (pc *PropertyConfiguration) IsSecret() (bool, error) {
	isSecret, ok := pc.isSecret.read()
	if !ok {
		return false, errors.Errorf(isSecretTag+" not specified for property %s", pc.name)
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

		// No handler for this value, return an error
		return errors.Errorf(
			"property configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}
