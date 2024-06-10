// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20230202p "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"
	storage "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type TrustedAccessRoleBindingExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *TrustedAccessRoleBindingExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20230202p.TrustedAccessRoleBinding{},
		&storage.TrustedAccessRoleBinding{}}
}
