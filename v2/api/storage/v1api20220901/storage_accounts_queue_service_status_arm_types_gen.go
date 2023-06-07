// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

type StorageAccounts_QueueService_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id"`

	// Name: The name of the resource
	Name *string `json:"name"`

	// Properties: The properties of a storage account’s Queue service.
	Properties *StorageAccounts_QueueService_Properties_STATUS_ARM `json:"properties"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type"`
}

type StorageAccounts_QueueService_Properties_STATUS_ARM struct {
	// Cors: Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Queue service.
	Cors *CorsRules_STATUS_ARM `json:"cors"`
}
