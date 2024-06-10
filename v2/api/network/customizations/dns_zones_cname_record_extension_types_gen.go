// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20180501 "github.com/Azure/azure-service-operator/v2/api/network/v1api20180501"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20180501/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type DnsZonesCNAMERecordExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *DnsZonesCNAMERecordExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20180501.DnsZonesCNAMERecord{},
		&storage.DnsZonesCNAMERecord{}}
}
