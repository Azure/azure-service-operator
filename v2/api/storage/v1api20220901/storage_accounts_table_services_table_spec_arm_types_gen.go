// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccounts_TableServices_Table_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Table resource properties.
	Properties *TableProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccounts_TableServices_Table_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
func (table StorageAccounts_TableServices_Table_Spec_ARM) GetAPIVersion() string {
	return "2022-09-01"
}

// GetName returns the Name of the resource
func (table *StorageAccounts_TableServices_Table_Spec_ARM) GetName() string {
	return table.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/tableServices/tables"
func (table *StorageAccounts_TableServices_Table_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/tableServices/tables"
}

type TableProperties_ARM struct {
	// SignedIdentifiers: List of stored access policies specified on the table.
	SignedIdentifiers []TableSignedIdentifier_ARM `json:"signedIdentifiers"`
}

// Object to set Table Access Policy.
type TableSignedIdentifier_ARM struct {
	// AccessPolicy: Access policy
	AccessPolicy *TableAccessPolicy_ARM `json:"accessPolicy,omitempty"`
	Id           *string                `json:"id,omitempty"`
}

// Table Access Policy Properties Object.
type TableAccessPolicy_ARM struct {
	// ExpiryTime: Expiry time of the access policy
	ExpiryTime *string `json:"expiryTime,omitempty"`

	// Permission: Required. List of abbreviated permissions. Supported permission values include 'r','a','u','d'
	Permission *string `json:"permission,omitempty"`

	// StartTime: Start time of the access policy
	StartTime *string `json:"startTime,omitempty"`
}
