package config

import (
	"gopkg.in/yaml.v3"
)

// VersionMetaData contains additional information about a specific version of a group
type VersionMetaData struct {
}


func (v *VersionMetaData) UnmarshalYAML(value *yaml.Node) error {
	return nil
}
