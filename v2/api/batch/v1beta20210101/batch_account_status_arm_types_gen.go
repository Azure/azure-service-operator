// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101

// Deprecated version of BatchAccount_STATUS. Use v1api20210101.BatchAccount_STATUS instead
type BatchAccount_STATUS_ARM struct {
	Id         *string                            `json:"id,omitempty"`
	Identity   *BatchAccountIdentity_STATUS_ARM   `json:"identity,omitempty"`
	Location   *string                            `json:"location,omitempty"`
	Name       *string                            `json:"name,omitempty"`
	Properties *BatchAccountProperties_STATUS_ARM `json:"properties,omitempty"`
	Tags       map[string]string                  `json:"tags,omitempty"`
	Type       *string                            `json:"type,omitempty"`
}

// Deprecated version of BatchAccountIdentity_STATUS. Use v1api20210101.BatchAccountIdentity_STATUS instead
type BatchAccountIdentity_STATUS_ARM struct {
	PrincipalId            *string                                                           `json:"principalId,omitempty"`
	TenantId               *string                                                           `json:"tenantId,omitempty"`
	Type                   *BatchAccountIdentity_Type_STATUS                                 `json:"type,omitempty"`
	UserAssignedIdentities map[string]BatchAccountIdentity_UserAssignedIdentities_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of BatchAccountProperties_STATUS. Use v1api20210101.BatchAccountProperties_STATUS instead
type BatchAccountProperties_STATUS_ARM struct {
	AccountEndpoint                       *string                                          `json:"accountEndpoint,omitempty"`
	ActiveJobAndJobScheduleQuota          *int                                             `json:"activeJobAndJobScheduleQuota,omitempty"`
	AutoStorage                           *AutoStorageProperties_STATUS_ARM                `json:"autoStorage,omitempty"`
	DedicatedCoreQuota                    *int                                             `json:"dedicatedCoreQuota,omitempty"`
	DedicatedCoreQuotaPerVMFamily         []VirtualMachineFamilyCoreQuota_STATUS_ARM       `json:"dedicatedCoreQuotaPerVMFamily,omitempty"`
	DedicatedCoreQuotaPerVMFamilyEnforced *bool                                            `json:"dedicatedCoreQuotaPerVMFamilyEnforced,omitempty"`
	Encryption                            *EncryptionProperties_STATUS_ARM                 `json:"encryption,omitempty"`
	KeyVaultReference                     *KeyVaultReference_STATUS_ARM                    `json:"keyVaultReference,omitempty"`
	LowPriorityCoreQuota                  *int                                             `json:"lowPriorityCoreQuota,omitempty"`
	PoolAllocationMode                    *PoolAllocationMode_STATUS                       `json:"poolAllocationMode,omitempty"`
	PoolQuota                             *int                                             `json:"poolQuota,omitempty"`
	PrivateEndpointConnections            []PrivateEndpointConnection_STATUS_ARM           `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                     *BatchAccountProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`
	PublicNetworkAccess                   *PublicNetworkAccessType_STATUS                  `json:"publicNetworkAccess,omitempty"`
}

// Deprecated version of AutoStorageProperties_STATUS. Use v1api20210101.AutoStorageProperties_STATUS instead
type AutoStorageProperties_STATUS_ARM struct {
	LastKeySync      *string `json:"lastKeySync,omitempty"`
	StorageAccountId *string `json:"storageAccountId,omitempty"`
}

// Deprecated version of BatchAccountIdentity_Type_STATUS. Use v1api20210101.BatchAccountIdentity_Type_STATUS instead
type BatchAccountIdentity_Type_STATUS string

const (
	BatchAccountIdentity_Type_STATUS_None           = BatchAccountIdentity_Type_STATUS("None")
	BatchAccountIdentity_Type_STATUS_SystemAssigned = BatchAccountIdentity_Type_STATUS("SystemAssigned")
	BatchAccountIdentity_Type_STATUS_UserAssigned   = BatchAccountIdentity_Type_STATUS("UserAssigned")
)

// Deprecated version of BatchAccountIdentity_UserAssignedIdentities_STATUS. Use v1api20210101.BatchAccountIdentity_UserAssignedIdentities_STATUS instead
type BatchAccountIdentity_UserAssignedIdentities_STATUS_ARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of EncryptionProperties_STATUS. Use v1api20210101.EncryptionProperties_STATUS instead
type EncryptionProperties_STATUS_ARM struct {
	KeySource          *EncryptionProperties_KeySource_STATUS `json:"keySource,omitempty"`
	KeyVaultProperties *KeyVaultProperties_STATUS_ARM         `json:"keyVaultProperties,omitempty"`
}

// Deprecated version of KeyVaultReference_STATUS. Use v1api20210101.KeyVaultReference_STATUS instead
type KeyVaultReference_STATUS_ARM struct {
	Id  *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
}

// Deprecated version of PrivateEndpointConnection_STATUS. Use v1api20210101.PrivateEndpointConnection_STATUS instead
type PrivateEndpointConnection_STATUS_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of VirtualMachineFamilyCoreQuota_STATUS. Use v1api20210101.VirtualMachineFamilyCoreQuota_STATUS instead
type VirtualMachineFamilyCoreQuota_STATUS_ARM struct {
	CoreQuota *int    `json:"coreQuota,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// Deprecated version of KeyVaultProperties_STATUS. Use v1api20210101.KeyVaultProperties_STATUS instead
type KeyVaultProperties_STATUS_ARM struct {
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}
