// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

type StorageAccounts_FileService_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id"`

	// Name: The name of the resource
	Name *string `json:"name"`

	// Properties: The properties of File services in storage account.
	Properties *StorageAccounts_FileService_Properties_STATUS_ARM `json:"properties"`

	// Sku: Sku name and tier.
	Sku *Sku_STATUS_ARM `json:"sku"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type"`
}

type StorageAccounts_FileService_Properties_STATUS_ARM struct {
	// Cors: Specifies CORS rules for the File service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// File service.
	Cors *CorsRules_STATUS_ARM `json:"cors"`

	// ProtocolSettings: Protocol settings for file service
	ProtocolSettings *ProtocolSettings_STATUS_ARM `json:"protocolSettings"`

	// ShareDeleteRetentionPolicy: The file service properties for share soft delete.
	ShareDeleteRetentionPolicy *DeleteRetentionPolicy_STATUS_ARM `json:"shareDeleteRetentionPolicy"`
}

// Protocol settings for file service
type ProtocolSettings_STATUS_ARM struct {
	// Smb: Setting for SMB protocol
	Smb *SmbSetting_STATUS_ARM `json:"smb"`
}

// Setting for SMB protocol
type SmbSetting_STATUS_ARM struct {
	// AuthenticationMethods: SMB authentication methods supported by server. Valid values are NTLMv2, Kerberos. Should be
	// passed as a string with delimiter ';'.
	AuthenticationMethods *string `json:"authenticationMethods"`

	// ChannelEncryption: SMB channel encryption supported by server. Valid values are AES-128-CCM, AES-128-GCM, AES-256-GCM.
	// Should be passed as a string with delimiter ';'.
	ChannelEncryption *string `json:"channelEncryption"`

	// KerberosTicketEncryption: Kerberos ticket encryption supported by server. Valid values are RC4-HMAC, AES-256. Should be
	// passed as a string with delimiter ';'
	KerberosTicketEncryption *string `json:"kerberosTicketEncryption"`

	// Multichannel: Multichannel setting. Applies to Premium FileStorage only.
	Multichannel *Multichannel_STATUS_ARM `json:"multichannel"`

	// Versions: SMB protocol versions supported by server. Valid values are SMB2.1, SMB3.0, SMB3.1.1. Should be passed as a
	// string with delimiter ';'.
	Versions *string `json:"versions"`
}

// Multichannel setting. Applies to Premium FileStorage only.
type Multichannel_STATUS_ARM struct {
	// Enabled: Indicates whether multichannel is enabled
	Enabled *bool `json:"enabled"`
}
