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

// CRDKustomizeFile captures the structure of the kustomization.yaml file we want to write
// All file paths specified as relative paths from the location of kustomize.yaml
//
// The final file output looks like this:
//
// resources:
// - bases/microsoft.authorization.azure.com_roleassignments.yaml
// - bases/microsoft.batch.azure.com_batchaccounts.yaml
// ...
// patches:
// - patches/webhook-conversion-microsoft.authorization.azure.com_roleassignments.yaml
// - patches/webhook-conversion-microsoft.batch.azure.com_batchaccounts.yaml
// ...
type CRDKustomizeFile struct {
	Resources      []string `yaml:"resources"`      // File paths to resource CRDs
	Patches        []string `yaml:"patches"`        // File paths to patch files used to enable conversions
	Configurations []string `yaml:"configurations"` // File paths to configuration files
}

// NewCRDKustomizeFile creates a new CRDKustomizeFile ready to populate
func NewCRDKustomizeFile() *CRDKustomizeFile {
	return &CRDKustomizeFile{}
}

// AddResource adds a filepath specifying another CRD definition to include
func (k *CRDKustomizeFile) AddResource(resourceFilePath string) {
	k.Resources = append(k.Resources, resourceFilePath)
}

// AddPatch adds a filepath specifying another CRD patch to include
func (k *CRDKustomizeFile) AddPatch(patchFilePath string) {
	k.Patches = append(k.Patches, patchFilePath)
}

func (k *CRDKustomizeFile) AddConfiguration(configFilePath string) {
	k.Configurations = append(k.Configurations, configFilePath)
}

// Save writes the kustomize configuration to the specified file path
func (k *CRDKustomizeFile) Save(destination string) error {
	data, err := yaml.Marshal(*k)
	if err != nil {
		return errors.Wrap(err, "serializing to yaml")
	}

	err = ioutil.WriteFile(destination, data, 0644) // #nosec G306
	if err != nil {
		return errors.Wrapf(err, "writing to %s", destination)
	}

	return nil
}
