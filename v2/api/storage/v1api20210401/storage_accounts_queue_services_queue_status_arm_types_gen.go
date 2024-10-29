// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210401

type StorageAccountsQueueServicesQueue_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Queue resource properties.
	Properties *QueueProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type QueueProperties_STATUS_ARM struct {
	// ApproximateMessageCount: Integer indicating an approximate number of messages in the queue. This number is not lower
	// than the actual number of messages in the queue, but could be higher.
	ApproximateMessageCount *int `json:"approximateMessageCount,omitempty"`

	// Metadata: A name-value pair that represents queue metadata.
	Metadata map[string]string `json:"metadata"`
}
