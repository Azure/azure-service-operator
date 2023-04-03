// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	v1api20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101storage"
	v20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type VirtualNetworkGatewayExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *VirtualNetworkGatewayExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20201101.VirtualNetworkGateway{},
		&v1api20201101s.VirtualNetworkGateway{},
		&v20201101.VirtualNetworkGateway{},
		&v20201101s.VirtualNetworkGateway{}}
}
