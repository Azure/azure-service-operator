// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101/storage"
	v20240301 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	v20240301s "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type VirtualNetworksSubnetExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *VirtualNetworksSubnetExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20201101.VirtualNetworksSubnet{},
		&v20201101s.VirtualNetworksSubnet{},
		&v20240301.VirtualNetworksSubnet{},
		&v20240301s.VirtualNetworksSubnet{}}
}
