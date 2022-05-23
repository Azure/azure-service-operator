// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type BatchAccounts_SpecARM struct {
	// Identity: The identity of the Batch account, if configured. This is only used when the user specifies
	// 'Microsoft.KeyVault' as their Batch account encryption configuration.
	Identity *BatchAccountIdentityARM `json:"identity,omitempty"`

	// Location: The region in which to create the account.
	Location *string `json:"location,omitempty"`

	// Name: A name for the Batch account which must be unique within the region. Batch account names must be between 3 and 24
	// characters in length and must use only numbers and lowercase letters. This name is used as part of the DNS name that is
	// used to access the Batch service in the region in which the account is created. For example:
	// http://accountname.region.batch.azure.com/.
	Name string `json:"name,omitempty"`

	// Properties: The properties of a Batch account.
	Properties *BatchAccountCreatePropertiesARM `json:"properties,omitempty"`

	// Tags: The user-specified tags associated with the account.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &BatchAccounts_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01"
func (accounts BatchAccounts_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (accounts *BatchAccounts_SpecARM) GetName() string {
	return accounts.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Batch/batchAccounts"
func (accounts *BatchAccounts_SpecARM) GetType() string {
	return "Microsoft.Batch/batchAccounts"
}

// Generated from: https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/BatchAccountCreateProperties
type BatchAccountCreatePropertiesARM struct {
	// AutoStorage: The properties related to the auto-storage account.
	AutoStorage *AutoStorageBasePropertiesARM `json:"autoStorage,omitempty"`

	// Encryption: Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using
	// a Microsoft managed key. For additional control, a customer-managed key can be used instead.
	Encryption *EncryptionPropertiesARM `json:"encryption,omitempty"`

	// KeyVaultReference: Identifies the Azure key vault associated with a Batch account.
	KeyVaultReference *KeyVaultReferenceARM `json:"keyVaultReference,omitempty"`

	// PoolAllocationMode: The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the
	// mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is
	// UserSubscription, clients must use Azure Active Directory. The default is BatchService.
	PoolAllocationMode *BatchAccountCreatePropertiesPoolAllocationMode `json:"poolAllocationMode,omitempty"`

	// PublicNetworkAccess: If not specified, the default value is 'enabled'.
	PublicNetworkAccess *BatchAccountCreatePropertiesPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/BatchAccountIdentity
type BatchAccountIdentityARM struct {
	// Type: The type of identity used for the Batch account.
	Type *BatchAccountIdentityType `json:"type,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/AutoStorageBaseProperties
type AutoStorageBasePropertiesARM struct {
	StorageAccountId *string `json:"storageAccountId,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","UserAssigned"}
type BatchAccountIdentityType string

const (
	BatchAccountIdentityTypeNone           = BatchAccountIdentityType("None")
	BatchAccountIdentityTypeSystemAssigned = BatchAccountIdentityType("SystemAssigned")
	BatchAccountIdentityTypeUserAssigned   = BatchAccountIdentityType("UserAssigned")
)

// Generated from: https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/EncryptionProperties
type EncryptionPropertiesARM struct {
	// KeySource: Type of the key source.
	KeySource *EncryptionPropertiesKeySource `json:"keySource,omitempty"`

	// KeyVaultProperties: KeyVault configuration when using an encryption KeySource of Microsoft.KeyVault.
	KeyVaultProperties *KeyVaultPropertiesARM `json:"keyVaultProperties,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/KeyVaultReference
type KeyVaultReferenceARM struct {
	Id *string `json:"id,omitempty"`

	// Url: The URL of the Azure key vault associated with the Batch account.
	Url *string `json:"url,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/KeyVaultProperties
type KeyVaultPropertiesARM struct {
	// KeyIdentifier: Full path to the versioned secret. Example
	// https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053. To be usable the following
	// prerequisites must be met:
	// The Batch Account has a System Assigned identity
	// The account identity has been granted Key/Get, Key/Unwrap and Key/Wrap permissions
	// The KeyVault has soft-delete and purge protection enabled
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}
