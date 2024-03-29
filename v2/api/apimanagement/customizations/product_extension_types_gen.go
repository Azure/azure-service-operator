// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220801 "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801"
	v20220801s "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	v20230501p "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview"
	v20230501ps "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ProductExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ProductExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20220801.Product{},
		&v20220801s.Product{},
		&v20230501p.Product{},
		&v20230501ps.Product{}}
}
