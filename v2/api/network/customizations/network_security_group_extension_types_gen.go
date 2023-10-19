// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101/storage"
	v1beta20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	v1beta20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type NetworkSecurityGroupExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *NetworkSecurityGroupExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20201101.NetworkSecurityGroup{},
		&v20201101s.NetworkSecurityGroup{},
		&v1beta20201101.NetworkSecurityGroup{},
		&v1beta20201101s.NetworkSecurityGroup{}}
}
