// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccounts_FileService_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of File services in storage account.
	Properties *StorageAccounts_FileService_Properties_Spec_ARM `json:"properties"`
}

var _ genruntime.ARMResourceSpec = &StorageAccounts_FileService_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
func (service StorageAccounts_FileService_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (service *StorageAccounts_FileService_Spec_ARM) GetName() string {
	return service.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/fileServices"
func (service *StorageAccounts_FileService_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/fileServices"
}

type StorageAccounts_FileService_Properties_Spec_ARM struct {
	// Cors: Specifies CORS rules for the File service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// File service.
	Cors *CorsRules_ARM `json:"cors"`

	// ProtocolSettings: Protocol settings for file service
	ProtocolSettings *ProtocolSettings_ARM `json:"protocolSettings"`

	// ShareDeleteRetentionPolicy: The file service properties for share soft delete.
	ShareDeleteRetentionPolicy *DeleteRetentionPolicy_ARM `json:"shareDeleteRetentionPolicy"`
}

// Protocol settings for file service
type ProtocolSettings_ARM struct {
	// Smb: Setting for SMB protocol
	Smb *SmbSetting_ARM `json:"smb"`
}

// Setting for SMB protocol
type SmbSetting_ARM struct {
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
	Multichannel *Multichannel_ARM `json:"multichannel"`

	// Versions: SMB protocol versions supported by server. Valid values are SMB2.1, SMB3.0, SMB3.1.1. Should be passed as a
	// string with delimiter ';'.
	Versions *string `json:"versions"`
}

// Multichannel setting. Applies to Premium FileStorage only.
type Multichannel_ARM struct {
	// Enabled: Indicates whether multichannel is enabled
	Enabled *bool `json:"enabled"`
}
