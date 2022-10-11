// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of StorageAccounts_QueueService_Spec. Use v1beta20210401.StorageAccounts_QueueService_Spec instead
type StorageAccounts_QueueService_Spec_ARM struct {
	Location   *string                               `json:"location,omitempty"`
	Name       string                                `json:"name,omitempty"`
	Properties *QueueServicePropertiesProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                     `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccounts_QueueService_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (service StorageAccounts_QueueService_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (service *StorageAccounts_QueueService_Spec_ARM) GetName() string {
	return service.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices"
func (service *StorageAccounts_QueueService_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices"
}

// Deprecated version of QueueServicePropertiesProperties. Use v1beta20210401.QueueServicePropertiesProperties instead
type QueueServicePropertiesProperties_ARM struct {
	Cors *CorsRules_ARM `json:"cors,omitempty"`
}
