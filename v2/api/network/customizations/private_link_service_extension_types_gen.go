// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220701 "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type PrivateLinkServiceExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *PrivateLinkServiceExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20220701.PrivateLinkService{},
		&storage.PrivateLinkService{}}
}
