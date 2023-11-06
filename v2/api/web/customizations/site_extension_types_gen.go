// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220301 "github.com/Azure/azure-service-operator/v2/api/web/v1api20220301"
	v20220301s "github.com/Azure/azure-service-operator/v2/api/web/v1api20220301/storage"
	v1beta20220301 "github.com/Azure/azure-service-operator/v2/api/web/v1beta20220301"
	v1beta20220301s "github.com/Azure/azure-service-operator/v2/api/web/v1beta20220301/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type SiteExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *SiteExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20220301.Site{},
		&v20220301s.Site{},
		&v1beta20220301.Site{},
		&v1beta20220301s.Site{}}
}
