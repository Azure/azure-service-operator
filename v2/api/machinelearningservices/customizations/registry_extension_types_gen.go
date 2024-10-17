// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20240401 "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20240401"
	storage "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20240401/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type RegistryExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *RegistryExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20240401.Registry{},
		&storage.Registry{}}
}
