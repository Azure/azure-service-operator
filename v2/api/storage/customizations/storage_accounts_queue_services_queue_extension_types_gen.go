// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	storagev1alpha1api20210401 "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401"
	"github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401storage"
	storagev1beta20210401 "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"
	"github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type StorageAccountsQueueServicesQueueExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *StorageAccountsQueueServicesQueueExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&storagev1alpha1api20210401.StorageAccountsQueueServicesQueue{},
		&v1alpha1api20210401storage.StorageAccountsQueueServicesQueue{},
		&storagev1beta20210401.StorageAccountsQueueServicesQueue{},
		&v1beta20210401storage.StorageAccountsQueueServicesQueue{}}
}
