// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210501 "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20210501"
	v20210501s "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20210501/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FlexibleServerExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FlexibleServerExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210501.FlexibleServer{},
		&v20210501s.FlexibleServer{}}
}
