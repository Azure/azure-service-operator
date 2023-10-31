// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210601 "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20210601"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20210601/storage"
	v1beta20210601 "github.com/Azure/azure-service-operator/v2/api/cdn/v1beta20210601"
	v1beta20210601s "github.com/Azure/azure-service-operator/v2/api/cdn/v1beta20210601/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ProfileExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ProfileExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210601.Profile{},
		&v20210601s.Profile{},
		&v1beta20210601.Profile{},
		&v1beta20210601s.Profile{}}
}
