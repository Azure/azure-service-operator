// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20200601 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20200601"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20200601storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type PrivateDnsZonesRecordsetTypeAExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *PrivateDnsZonesRecordsetTypeAExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20200601.PrivateDnsZonesRecordsetTypeA{},
		&v20200601s.PrivateDnsZonesRecordsetTypeA{}}
}
