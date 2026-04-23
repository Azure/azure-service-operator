/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"os"

	"github.com/rotisserie/eris"
	"gopkg.in/yaml.v3"
)

// GroupConfigurationFile represents a per-group configuration file containing
// TypeFilters, TypeTransformers, and GroupModelConfiguration for a single group.
type GroupConfigurationFile struct {
	// Filters used to control which types from this group are included
	TypeFilters []*TypeFilter `yaml:"typeFilters,omitempty"`
	// Transformers used to remap types in this group
	Transformers []*TypeTransformer `yaml:"typeTransformers,omitempty"`
	// GroupModelConfiguration contains version/type/property configuration for this group.
	// This field is not directly YAML-decoded; it's handled via custom UnmarshalYAML
	// so that we can properly initialize GroupConfiguration with the group name.
	GroupModelConfiguration *GroupConfiguration `yaml:"-"`
	// groupModelNode stores the raw YAML for deferred decoding of GroupModelConfiguration
	groupModelNode *yaml.Node
}

// UnmarshalYAML implements custom YAML unmarshalling for GroupConfigurationFile.
// We need custom handling to capture the groupModelConfiguration node for deferred decoding.
func (gcf *GroupConfigurationFile) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return eris.Errorf("expected mapping node for group configuration file, but found %s", value.Tag)
	}

	var lastKey string
	for i, n := range value.Content {
		if i%2 == 0 {
			lastKey = n.Value
			continue
		}

		switch lastKey {
		case "typeFilters":
			if err := n.Decode(&gcf.TypeFilters); err != nil {
				return eris.Wrap(err, "decoding typeFilters")
			}
		case "typeTransformers":
			if err := n.Decode(&gcf.Transformers); err != nil {
				return eris.Wrap(err, "decoding typeTransformers")
			}
		case "groupModelConfiguration":
			gcf.groupModelNode = n
		default:
			return eris.Errorf("unexpected field %q in group configuration file", lastKey)
		}
	}

	return nil
}

// loadGroupConfigurationFile loads a per-group configuration file from the given path.
// The groupName parameter is used to initialize the GroupConfiguration with the correct name.
func loadGroupConfigurationFile(path string, groupName string) (*GroupConfigurationFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, eris.Wrapf(err, "opening group configuration file %q", path)
	}
	defer f.Close()

	result := &GroupConfigurationFile{}
	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(result)
	if err != nil {
		return nil, eris.Wrapf(err, "group configuration file %q is not valid YAML", path)
	}

	// Decode the groupModelConfiguration with a properly initialized GroupConfiguration
	if result.groupModelNode != nil {
		gc := NewGroupConfiguration(groupName)
		if err := result.groupModelNode.Decode(gc); err != nil {
			return nil, eris.Wrapf(err, "decoding groupModelConfiguration in %q", path)
		}

		result.GroupModelConfiguration = gc
	}

	return result, nil
}
