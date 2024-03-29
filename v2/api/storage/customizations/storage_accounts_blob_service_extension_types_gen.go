// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210401 "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	v20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401/storage"
	v20220901 "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901"
	v20220901s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/storage"
	v20230101 "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101"
	v20230101s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type StorageAccountsBlobServiceExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *StorageAccountsBlobServiceExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210401.StorageAccountsBlobService{},
		&v20210401s.StorageAccountsBlobService{},
		&v20220901.StorageAccountsBlobService{},
		&v20220901s.StorageAccountsBlobService{},
		&v20230101.StorageAccountsBlobService{},
		&v20230101s.StorageAccountsBlobService{}}
}
