// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20230101 "github.com/Azure/azure-service-operator/v2/api/insights/v1api20230101"
	v20230101s "github.com/Azure/azure-service-operator/v2/api/insights/v1api20230101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ActionGroupExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ActionGroupExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20230101.ActionGroup{},
		&v20230101s.ActionGroup{}}
}
