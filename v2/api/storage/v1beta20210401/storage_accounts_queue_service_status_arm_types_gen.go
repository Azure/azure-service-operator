// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

// Deprecated version of StorageAccounts_QueueService_STATUS. Use v1api20210401.StorageAccounts_QueueService_STATUS instead
type StorageAccounts_QueueService_STATUS_ARM struct {
	Id         *string                                             `json:"id"`
	Name       *string                                             `json:"name"`
	Properties *StorageAccounts_QueueService_Properties_STATUS_ARM `json:"properties"`
	Type       *string                                             `json:"type"`
}

// Deprecated version of StorageAccounts_QueueService_Properties_STATUS. Use v1api20210401.StorageAccounts_QueueService_Properties_STATUS instead
type StorageAccounts_QueueService_Properties_STATUS_ARM struct {
	Cors *CorsRules_STATUS_ARM `json:"cors"`
}
