// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccountsTableService_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a storage account’s Table service.
	Properties *StorageAccounts_TableService_Properties_Spec `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsTableService_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
func (service StorageAccountsTableService_Spec) GetAPIVersion() string {
	return "2022-09-01"
}

// GetName returns the Name of the resource
func (service *StorageAccountsTableService_Spec) GetName() string {
	return service.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/tableServices"
func (service *StorageAccountsTableService_Spec) GetType() string {
	return "Microsoft.Storage/storageAccounts/tableServices"
}

type StorageAccounts_TableService_Properties_Spec struct {
	// Cors: Specifies CORS rules for the Table service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Table service.
	Cors *CorsRules `json:"cors,omitempty"`
}
