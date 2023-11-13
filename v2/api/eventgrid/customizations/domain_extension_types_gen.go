// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601/storage"
	v1beta20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601"
	v1beta20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type DomainExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *DomainExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20200601.Domain{},
		&v20200601s.Domain{},
		&v1beta20200601.Domain{},
		&v1beta20200601s.Domain{}}
}
