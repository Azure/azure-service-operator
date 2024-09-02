// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210701 "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20210701"
	v20210701s "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20210701/storage"
	v20240401 "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20240401"
	v20240401s "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20240401/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type WorkspacesComputeExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *WorkspacesComputeExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210701.WorkspacesCompute{},
		&v20210701s.WorkspacesCompute{},
		&v20240401.WorkspacesCompute{},
		&v20240401s.WorkspacesCompute{}}
}
