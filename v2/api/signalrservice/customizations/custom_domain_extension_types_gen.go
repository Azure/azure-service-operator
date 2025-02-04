// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20240301 "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1api20240301"
	storage "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1api20240301/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type CustomDomainExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *CustomDomainExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20240301.CustomDomain{},
		&storage.CustomDomain{}}
}
