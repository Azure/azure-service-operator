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
	name                      string
	nameInNextVersion         *string
	nameInNextVersionConsumed bool
	armReference              *bool
	armReferenceConsumed      bool
	isSecret                  *bool
	isSecretConsumed          bool
}

const (
	armReferenceTag      = "$armReference"
	nameInNextVersionTag = "$nameInNextVersion"
	isSecretTag          = "$isSecret"
)

// NewPropertyConfiguration returns a new (empty) property configuration
func NewPropertyConfiguration(name string) *PropertyConfiguration {
	return &PropertyConfiguration{
		name: name,
	}
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (pc *PropertyConfiguration) ARMReference() (bool, error) {
	if pc.armReference == nil {
		msg := fmt.Sprintf(armReferenceTag+" not specified for property %s", pc.name)
		return false, NewNotConfiguredError(msg)
	}

	pc.armReferenceConsumed = true
	return *pc.armReference, nil
}

// SetARMReference configures specifies whether this property is an ARM reference or not
func (pc *PropertyConfiguration) SetARMReference(isARMRef bool) *PropertyConfiguration {
	pc.armReference = &isARMRef
	return pc
}

// VerifyARMReferenceConsumed returns an error if our configuration as an ARM reference was not consumed.
func (pc *PropertyConfiguration) VerifyARMReferenceConsumed() error {
	if pc.armReference != nil && !pc.armReferenceConsumed {
		return errors.Errorf("property %s: "+armReferenceTag+": %t not consumed", pc.name, *pc.armReference)
	}

	return nil
}

// IsSecret looks up a property to determine if it's a secret
func (pc *PropertyConfiguration) IsSecret() (bool, error) {
	if pc.isSecret == nil {
		return false, errors.Errorf(isSecretTag+" not specified for property %s", pc.name)
	}

	pc.isSecretConsumed = true
	return *pc.isSecret, nil
}

// SetIsSecret marks this property as a secret
func (pc *PropertyConfiguration) SetIsSecret(isSecret bool) *PropertyConfiguration {
	pc.isSecret = &isSecret
	return pc
}

// VerifyIsSecretConsumed returns an error if our configuration as a secret was not consumed.
func (pc *PropertyConfiguration) VerifyIsSecretConsumed() error {
	if pc.isSecret != nil && !pc.isSecretConsumed {
		return errors.Errorf("property %s: "+isSecretTag+": %t not consumed", pc.name, *pc.isSecret)
	}

	return nil
}

// PropertyRename looks up a property to determine whether it is being renamed in the next version
func (pc *PropertyConfiguration) PropertyRename() (string, error) {
	if pc.nameInNextVersion == nil {
		msg := fmt.Sprintf(nameInNextVersionTag+" not specified for property %s", pc.name)
		return "", NewNotConfiguredError(msg)
	}

	pc.nameInNextVersionConsumed = true
	return *pc.nameInNextVersion, nil
}

// SetRenamedTo configures what this property is renamed to in the next version of the containing type
func (pc *PropertyConfiguration) SetRenamedTo(renamedTo string) *PropertyConfiguration {
	pc.nameInNextVersion = &renamedTo
	return pc
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

		if strings.EqualFold(lastId, nameInNextVersionTag) && c.Kind == yaml.ScalarNode {
			pc.SetRenamedTo(c.Value)
			continue
		}

		if strings.EqualFold(lastId, isSecretTag) && c.Kind == yaml.ScalarNode {
			var isSecret bool
			err := c.Decode(&isSecret)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", isSecretTag)
			}

			pc.SetIsSecret(isSecret)
			continue
		}

		if strings.EqualFold(lastId, armReferenceTag) && c.Kind == yaml.ScalarNode {
			var isARMRef bool
			err := c.Decode(&isARMRef)
			if err != nil {
				return errors.Wrapf(err, "decoding %s", armReferenceTag)
			}

			pc.SetARMReference(isARMRef)
			continue
		}

		// No handler for this value, return an error
		return errors.Errorf(
			"property configuration, unexpected yaml value %s: %s (line %d col %d)", lastId, c.Value, c.Line, c.Column)
	}

	return nil
}
