// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20211101 "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	v1api20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ServersAdvancedThreatProtectionSettingExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ServersAdvancedThreatProtectionSettingExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20211101.ServersAdvancedThreatProtectionSetting{},
		&v1api20211101s.ServersAdvancedThreatProtectionSetting{}}
}
