/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package kustomization

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type ResourceDefinition struct {
	Metadata ResourceDefinitionMetadata `yaml:"metadata"`
}

func LoadResourceDefinition(filePath string) (*ResourceDefinition, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "reading resource definition from %s", filePath)
	}

	result := &ResourceDefinition{}
	err = yaml.Unmarshal(fileBytes, result)
	if err != nil {
		return nil, errors.Wrapf(err, "unmarshalling resource definition from %s", filePath)
	}

	return result, nil
}

func (d *ResourceDefinition) Name() string {
	return d.Metadata.Name
}

type ResourceDefinitionMetadata struct {
	Name string `yaml:"name"`
}
