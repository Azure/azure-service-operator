// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20200601 "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601/storage"
	v20240601 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240601"
	v20240601s "github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type PrivateDnsZonesCNAMERecordExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *PrivateDnsZonesCNAMERecordExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20200601.PrivateDnsZonesCNAMERecord{},
		&v20200601s.PrivateDnsZonesCNAMERecord{},
		&v20240601.PrivateDnsZonesCNAMERecord{},
		&v20240601s.PrivateDnsZonesCNAMERecord{}}
}
