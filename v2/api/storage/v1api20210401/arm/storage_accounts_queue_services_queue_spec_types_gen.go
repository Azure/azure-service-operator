// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccountsQueueServicesQueue_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Queue resource properties.
	Properties *QueueProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsQueueServicesQueue_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (queue StorageAccountsQueueServicesQueue_Spec) GetAPIVersion() string {
	return "2021-04-01"
}

// GetName returns the Name of the resource
func (queue *StorageAccountsQueueServicesQueue_Spec) GetName() string {
	return queue.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices/queues"
func (queue *StorageAccountsQueueServicesQueue_Spec) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices/queues"
}

type QueueProperties struct {
	// Metadata: A name-value pair that represents queue metadata.
	Metadata map[string]string `json:"metadata"`
}
