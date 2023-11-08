// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20211001 "github.com/Azure/azure-service-operator/v2/api/subscription/v1api20211001"
	v20211001s "github.com/Azure/azure-service-operator/v2/api/subscription/v1api20211001/storage"
	v1beta20211001 "github.com/Azure/azure-service-operator/v2/api/subscription/v1beta20211001"
	v1beta20211001s "github.com/Azure/azure-service-operator/v2/api/subscription/v1beta20211001/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type AliasExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *AliasExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20211001.Alias{},
		&v20211001s.Alias{},
		&v1beta20211001.Alias{},
		&v1beta20211001s.Alias{}}
}
