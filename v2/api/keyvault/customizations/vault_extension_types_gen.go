// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210401p "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20210401preview"
	v20210401ps "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20210401previewstorage"
	v1beta20210401p "github.com/Azure/azure-service-operator/v2/api/keyvault/v1beta20210401preview"
	v1beta20210401ps "github.com/Azure/azure-service-operator/v2/api/keyvault/v1beta20210401previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type VaultExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *VaultExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210401p.Vault{},
		&v20210401ps.Vault{},
		&v1beta20210401p.Vault{},
		&v1beta20210401ps.Vault{}}
}
