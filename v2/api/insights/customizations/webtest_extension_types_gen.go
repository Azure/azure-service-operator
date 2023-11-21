// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20180501p "github.com/Azure/azure-service-operator/v2/api/insights/v1api20180501preview"
	v20180501ps "github.com/Azure/azure-service-operator/v2/api/insights/v1api20180501preview/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type WebtestExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *WebtestExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20180501p.Webtest{},
		&v20180501ps.Webtest{}}
}
