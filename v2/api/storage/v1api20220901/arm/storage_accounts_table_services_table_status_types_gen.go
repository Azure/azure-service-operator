// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type StorageAccountsTableServicesTable_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Table resource properties.
	Properties *TableProperties_STATUS `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type TableProperties_STATUS struct {
	// SignedIdentifiers: List of stored access policies specified on the table.
	SignedIdentifiers []TableSignedIdentifier_STATUS `json:"signedIdentifiers"`

	// TableName: Table name under the specified account
	TableName *string `json:"tableName,omitempty"`
}

// Object to set Table Access Policy.
type TableSignedIdentifier_STATUS struct {
	// AccessPolicy: Access policy
	AccessPolicy *TableAccessPolicy_STATUS `json:"accessPolicy,omitempty"`

	// Id: unique-64-character-value of the stored access policy.
	Id *string `json:"id,omitempty"`
}

// Table Access Policy Properties Object.
type TableAccessPolicy_STATUS struct {
	// ExpiryTime: Expiry time of the access policy
	ExpiryTime *string `json:"expiryTime,omitempty"`

	// Permission: Required. List of abbreviated permissions. Supported permission values include 'r','a','u','d'
	Permission *string `json:"permission,omitempty"`

	// StartTime: Start time of the access policy
	StartTime *string `json:"startTime,omitempty"`
}
