// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20231001 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001"
	v20231001s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
	v20240402p "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240402preview"
	v20240402ps "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240402preview/storage"
	v20240901 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901"
	v20240901s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type TrustedAccessRoleBindingExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *TrustedAccessRoleBindingExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20231001.TrustedAccessRoleBinding{},
		&v20231001s.TrustedAccessRoleBinding{},
		&v20240402p.TrustedAccessRoleBinding{},
		&v20240402ps.TrustedAccessRoleBinding{},
		&v20240901.TrustedAccessRoleBinding{},
		&v20240901s.TrustedAccessRoleBinding{}}
}
