// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220501 "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1api20220501"
	storage "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1api20220501/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ConfigurationStoreExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ConfigurationStoreExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20220501.ConfigurationStore{},
		&storage.ConfigurationStore{}}
}
