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

// ConversionPatchFile specifies a fragment of YAML we create to configure storage conversions via Kustomize
//
// The final file looks like this:
//
// apiVersion: apiextensions.k8s.io/v1
// kind: CustomResourceDefinition
// metadata:
//
//	name: roleassignments.microsoft.authorization.azure.com
//	annotations:
//	    cert-manager.io/inject-ca-from: $(CERTIFICATE_NAMESPACE)/$(CERTIFICATE_NAME)
//
// spec:
//
//	preserveUnknownFields: false
//	conversion:
//	    strategy: Webhook
//	    webhook:
//	        conversionReviewVersions:
//	            - v1beta1
//	        clientConfig:
//	            service:
//	                namespace: system
//	                name: webhook-service
//	                path: /convert
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
			Annotations: map[string]string{
				certManagerInjectKey: certManagerInjectValue,
			},
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
							Name:      "webhook-service",
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

const (
	// This annotation instructs the cert-manager webhooks to inject
	// the CA bundle from a certificate into the client config of this
	// CRD's conversion webhook.
	certManagerInjectKey = "cert-manager.io/inject-ca-from"

	// The values here will be substituted by kusomization vars when
	// this kustomize directory is built.
	certManagerInjectValue = "$(CERTIFICATE_NAMESPACE)/$(CERTIFICATE_NAME)"
)

type conversionPatchMetadata struct {
	Name        string            `yaml:"name"`
	Annotations map[string]string `yaml:"annotations"`
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
