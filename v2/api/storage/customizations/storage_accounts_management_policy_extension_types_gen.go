// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	alpha20210401 "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401"
	alpha20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401storage"
	v20210401 "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"
	v20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type StorageAccountsManagementPolicyExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *StorageAccountsManagementPolicyExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&alpha20210401.StorageAccountsManagementPolicy{},
		&alpha20210401s.StorageAccountsManagementPolicy{},
		&v20210401.StorageAccountsManagementPolicy{},
		&v20210401s.StorageAccountsManagementPolicy{}}
}
