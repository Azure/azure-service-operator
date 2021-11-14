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

type ConversionPatchFile struct {
	ApiVersion string                  `yaml:"apiVersion"`
	Kind       string                  `yaml:"kind"`
	Metadata   conversionPatchMetadata `yaml:"metadata"`
	Spec       conversionPatchSpec     `yaml:"spec"`
}

func NewConversionPatchFile(resourceName string) *ConversionPatchFile {
	return &ConversionPatchFile{
		ApiVersion: "apiextensions.k8s.io/v1",
		Kind:       "CustomResourceDefinition",
		Metadata: conversionPatchMetadata{
			Name: resourceName,
		},
		Spec: conversionPatchSpec{
			PreserveUnknownFields: false,
			Conversion: conversionPatchConversion{
				Strategy: "Webhook",
				Webhook: conversionPatchWebhook{
					ConversionReviewVersions: []string{
						"v1beta1",
					},
					ClientConfig: conversionPatchClientConfig{
						Service: conversionPatchServiceConfig{
							Namespace: "system",
							Name:      "webook-service",
							Path:      "/convert",
						},
					},
				},
			},
		},
	}
}

// Save writes the patch to the specified file path
func (p *ConversionPatchFile) Save(destination string) error {
	data, err := yaml.Marshal(*p)
	if err != nil {
		return errors.Wrap(err, "serializing to yaml")
	}

	err = ioutil.WriteFile(destination, data, 0644) // #nosec G306
	if err != nil {
		return errors.Wrapf(err, "writing to %s", destination)
	}

	return nil
}

type conversionPatchMetadata struct {
	Name string `yaml:"name"`
}

type conversionPatchSpec struct {
	PreserveUnknownFields bool                      `yaml:"preserveUnknownFields"`
	Conversion            conversionPatchConversion `yaml:"conversion"`
}

type conversionPatchConversion struct {
	Strategy string                 `yaml:"strategy"`
	Webhook  conversionPatchWebhook `yaml:"webhook"`
}

type conversionPatchWebhook struct {
	ConversionReviewVersions []string                    `yaml:"conversionReviewVersions"`
	ClientConfig             conversionPatchClientConfig `yaml:"clientConfig"`
}

type conversionPatchClientConfig struct {
	Service conversionPatchServiceConfig `yaml:"service"`
}

type conversionPatchServiceConfig struct {
	Namespace string `yaml:"namespace"`
	Name      string `yaml:"name"`
	Path      string `yaml:"path"`
}
