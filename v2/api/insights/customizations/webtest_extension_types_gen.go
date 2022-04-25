// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	alpha20180501p "github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20180501preview"
	alpha20180501ps "github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20180501previewstorage"
	v20180501p "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20180501preview"
	v20180501ps "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20180501previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type WebtestExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *WebtestExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&alpha20180501p.Webtest{},
		&alpha20180501ps.Webtest{},
		&v20180501p.Webtest{},
		&v20180501ps.Webtest{}}
}
