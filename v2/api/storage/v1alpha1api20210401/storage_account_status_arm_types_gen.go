// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

// Deprecated version of StorageAccount_STATUS. Use v1beta20210401.StorageAccount_STATUS instead
type StorageAccount_STATUSARM struct {
	ExtendedLocation *ExtendedLocation_STATUSARM         `json:"extendedLocation,omitempty"`
	Id               *string                             `json:"id,omitempty"`
	Identity         *Identity_STATUSARM                 `json:"identity,omitempty"`
	Kind             *StorageAccountSTATUSKind           `json:"kind,omitempty"`
	Location         *string                             `json:"location,omitempty"`
	Name             *string                             `json:"name,omitempty"`
	Properties       *StorageAccountProperties_STATUSARM `json:"properties,omitempty"`
	Sku              *Sku_STATUSARM                      `json:"sku,omitempty"`
	Tags             map[string]string                   `json:"tags,omitempty"`
	Type             *string                             `json:"type,omitempty"`
}

// Deprecated version of ExtendedLocation_STATUS. Use v1beta20210401.ExtendedLocation_STATUS instead
type ExtendedLocation_STATUSARM struct {
	Name *string                      `json:"name,omitempty"`
	Type *ExtendedLocationType_STATUS `json:"type,omitempty"`
}

// Deprecated version of Identity_STATUS. Use v1beta20210401.Identity_STATUS instead
type Identity_STATUSARM struct {
	PrincipalId            *string                                   `json:"principalId,omitempty"`
	TenantId               *string                                   `json:"tenantId,omitempty"`
	Type                   *IdentitySTATUSType                       `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUSARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of StorageAccountProperties_STATUS. Use v1beta20210401.StorageAccountProperties_STATUS instead
type StorageAccountProperties_STATUSARM struct {
	AccessTier                            *StorageAccountPropertiesSTATUSAccessTier                 `json:"accessTier,omitempty"`
	AllowBlobPublicAccess                 *bool                                                     `json:"allowBlobPublicAccess,omitempty"`
	AllowCrossTenantReplication           *bool                                                     `json:"allowCrossTenantReplication,omitempty"`
	AllowSharedKeyAccess                  *bool                                                     `json:"allowSharedKeyAccess,omitempty"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication_STATUSARM          `json:"azureFilesIdentityBasedAuthentication,omitempty"`
	BlobRestoreStatus                     *BlobRestoreStatus_STATUSARM                              `json:"blobRestoreStatus,omitempty"`
	CreationTime                          *string                                                   `json:"creationTime,omitempty"`
	CustomDomain                          *CustomDomain_STATUSARM                                   `json:"customDomain,omitempty"`
	Encryption                            *Encryption_STATUSARM                                     `json:"encryption,omitempty"`
	FailoverInProgress                    *bool                                                     `json:"failoverInProgress,omitempty"`
	GeoReplicationStats                   *GeoReplicationStats_STATUSARM                            `json:"geoReplicationStats,omitempty"`
	IsHnsEnabled                          *bool                                                     `json:"isHnsEnabled,omitempty"`
	IsNfsV3Enabled                        *bool                                                     `json:"isNfsV3Enabled,omitempty"`
	KeyCreationTime                       *KeyCreationTime_STATUSARM                                `json:"keyCreationTime,omitempty"`
	KeyPolicy                             *KeyPolicy_STATUSARM                                      `json:"keyPolicy,omitempty"`
	LargeFileSharesState                  *StorageAccountPropertiesSTATUSLargeFileSharesState       `json:"largeFileSharesState,omitempty"`
	LastGeoFailoverTime                   *string                                                   `json:"lastGeoFailoverTime,omitempty"`
	MinimumTlsVersion                     *StorageAccountPropertiesSTATUSMinimumTlsVersion          `json:"minimumTlsVersion,omitempty"`
	NetworkAcls                           *NetworkRuleSet_STATUSARM                                 `json:"networkAcls,omitempty"`
	PrimaryEndpoints                      *Endpoints_STATUSARM                                      `json:"primaryEndpoints,omitempty"`
	PrimaryLocation                       *string                                                   `json:"primaryLocation,omitempty"`
	PrivateEndpointConnections            []PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                     *StorageAccountPropertiesSTATUSProvisioningState          `json:"provisioningState,omitempty"`
	RoutingPreference                     *RoutingPreference_STATUSARM                              `json:"routingPreference,omitempty"`
	SasPolicy                             *SasPolicy_STATUSARM                                      `json:"sasPolicy,omitempty"`
	SecondaryEndpoints                    *Endpoints_STATUSARM                                      `json:"secondaryEndpoints,omitempty"`
	SecondaryLocation                     *string                                                   `json:"secondaryLocation,omitempty"`
	StatusOfPrimary                       *StorageAccountPropertiesSTATUSStatusOfPrimary            `json:"statusOfPrimary,omitempty"`
	StatusOfSecondary                     *StorageAccountPropertiesSTATUSStatusOfSecondary          `json:"statusOfSecondary,omitempty"`
	SupportsHttpsTrafficOnly              *bool                                                     `json:"supportsHttpsTrafficOnly,omitempty"`
}

// Deprecated version of StorageAccountSTATUSKind. Use v1beta20210401.StorageAccountSTATUSKind instead
type StorageAccountSTATUSKind string

const (
	StorageAccountSTATUSKind_BlobStorage      = StorageAccountSTATUSKind("BlobStorage")
	StorageAccountSTATUSKind_BlockBlobStorage = StorageAccountSTATUSKind("BlockBlobStorage")
	StorageAccountSTATUSKind_FileStorage      = StorageAccountSTATUSKind("FileStorage")
	StorageAccountSTATUSKind_Storage          = StorageAccountSTATUSKind("Storage")
	StorageAccountSTATUSKind_StorageV2        = StorageAccountSTATUSKind("StorageV2")
)

// Deprecated version of AzureFilesIdentityBasedAuthentication_STATUS. Use v1beta20210401.AzureFilesIdentityBasedAuthentication_STATUS instead
type AzureFilesIdentityBasedAuthentication_STATUSARM struct {
	ActiveDirectoryProperties *ActiveDirectoryProperties_STATUSARM                                `json:"activeDirectoryProperties,omitempty"`
	DefaultSharePermission    *AzureFilesIdentityBasedAuthenticationSTATUSDefaultSharePermission  `json:"defaultSharePermission,omitempty"`
	DirectoryServiceOptions   *AzureFilesIdentityBasedAuthenticationSTATUSDirectoryServiceOptions `json:"directoryServiceOptions,omitempty"`
}

// Deprecated version of BlobRestoreStatus_STATUS. Use v1beta20210401.BlobRestoreStatus_STATUS instead
type BlobRestoreStatus_STATUSARM struct {
	FailureReason *string                          `json:"failureReason,omitempty"`
	Parameters    *BlobRestoreParameters_STATUSARM `json:"parameters,omitempty"`
	RestoreId     *string                          `json:"restoreId,omitempty"`
	Status        *BlobRestoreStatusSTATUSStatus   `json:"status,omitempty"`
}

// Deprecated version of CustomDomain_STATUS. Use v1beta20210401.CustomDomain_STATUS instead
type CustomDomain_STATUSARM struct {
	Name             *string `json:"name,omitempty"`
	UseSubDomainName *bool   `json:"useSubDomainName,omitempty"`
}

// Deprecated version of Encryption_STATUS. Use v1beta20210401.Encryption_STATUS instead
type Encryption_STATUSARM struct {
	Identity                        *EncryptionIdentity_STATUSARM `json:"identity,omitempty"`
	KeySource                       *EncryptionSTATUSKeySource    `json:"keySource,omitempty"`
	Keyvaultproperties              *KeyVaultProperties_STATUSARM `json:"keyvaultproperties,omitempty"`
	RequireInfrastructureEncryption *bool                         `json:"requireInfrastructureEncryption,omitempty"`
	Services                        *EncryptionServices_STATUSARM `json:"services,omitempty"`
}

// Deprecated version of Endpoints_STATUS. Use v1beta20210401.Endpoints_STATUS instead
type Endpoints_STATUSARM struct {
	Blob               *string                                     `json:"blob,omitempty"`
	Dfs                *string                                     `json:"dfs,omitempty"`
	File               *string                                     `json:"file,omitempty"`
	InternetEndpoints  *StorageAccountInternetEndpoints_STATUSARM  `json:"internetEndpoints,omitempty"`
	MicrosoftEndpoints *StorageAccountMicrosoftEndpoints_STATUSARM `json:"microsoftEndpoints,omitempty"`
	Queue              *string                                     `json:"queue,omitempty"`
	Table              *string                                     `json:"table,omitempty"`
	Web                *string                                     `json:"web,omitempty"`
}

// Deprecated version of ExtendedLocationType_STATUS. Use v1beta20210401.ExtendedLocationType_STATUS instead
type ExtendedLocationType_STATUS string

const ExtendedLocationType_STATUS_EdgeZone = ExtendedLocationType_STATUS("EdgeZone")

// Deprecated version of GeoReplicationStats_STATUS. Use v1beta20210401.GeoReplicationStats_STATUS instead
type GeoReplicationStats_STATUSARM struct {
	CanFailover  *bool                            `json:"canFailover,omitempty"`
	LastSyncTime *string                          `json:"lastSyncTime,omitempty"`
	Status       *GeoReplicationStatsSTATUSStatus `json:"status,omitempty"`
}

// Deprecated version of IdentitySTATUSType. Use v1beta20210401.IdentitySTATUSType instead
type IdentitySTATUSType string

const (
	IdentitySTATUSType_None                       = IdentitySTATUSType("None")
	IdentitySTATUSType_SystemAssigned             = IdentitySTATUSType("SystemAssigned")
	IdentitySTATUSType_SystemAssignedUserAssigned = IdentitySTATUSType("SystemAssigned,UserAssigned")
	IdentitySTATUSType_UserAssigned               = IdentitySTATUSType("UserAssigned")
)

// Deprecated version of KeyCreationTime_STATUS. Use v1beta20210401.KeyCreationTime_STATUS instead
type KeyCreationTime_STATUSARM struct {
	Key1 *string `json:"key1,omitempty"`
	Key2 *string `json:"key2,omitempty"`
}

// Deprecated version of KeyPolicy_STATUS. Use v1beta20210401.KeyPolicy_STATUS instead
type KeyPolicy_STATUSARM struct {
	KeyExpirationPeriodInDays *int `json:"keyExpirationPeriodInDays,omitempty"`
}

// Deprecated version of NetworkRuleSet_STATUS. Use v1beta20210401.NetworkRuleSet_STATUS instead
type NetworkRuleSet_STATUSARM struct {
	Bypass              *NetworkRuleSetSTATUSBypass        `json:"bypass,omitempty"`
	DefaultAction       *NetworkRuleSetSTATUSDefaultAction `json:"defaultAction,omitempty"`
	IpRules             []IPRule_STATUSARM                 `json:"ipRules,omitempty"`
	ResourceAccessRules []ResourceAccessRule_STATUSARM     `json:"resourceAccessRules,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule_STATUSARM     `json:"virtualNetworkRules,omitempty"`
}

// Deprecated version of PrivateEndpointConnection_STATUS_SubResourceEmbedded. Use v1beta20210401.PrivateEndpointConnection_STATUS_SubResourceEmbedded instead
type PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of RoutingPreference_STATUS. Use v1beta20210401.RoutingPreference_STATUS instead
type RoutingPreference_STATUSARM struct {
	PublishInternetEndpoints  *bool                                 `json:"publishInternetEndpoints,omitempty"`
	PublishMicrosoftEndpoints *bool                                 `json:"publishMicrosoftEndpoints,omitempty"`
	RoutingChoice             *RoutingPreferenceSTATUSRoutingChoice `json:"routingChoice,omitempty"`
}

// Deprecated version of SasPolicy_STATUS. Use v1beta20210401.SasPolicy_STATUS instead
type SasPolicy_STATUSARM struct {
	ExpirationAction    *SasPolicySTATUSExpirationAction `json:"expirationAction,omitempty"`
	SasExpirationPeriod *string                          `json:"sasExpirationPeriod,omitempty"`
}

// Deprecated version of UserAssignedIdentity_STATUS. Use v1beta20210401.UserAssignedIdentity_STATUS instead
type UserAssignedIdentity_STATUSARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of ActiveDirectoryProperties_STATUS. Use v1beta20210401.ActiveDirectoryProperties_STATUS instead
type ActiveDirectoryProperties_STATUSARM struct {
	AzureStorageSid   *string `json:"azureStorageSid,omitempty"`
	DomainGuid        *string `json:"domainGuid,omitempty"`
	DomainName        *string `json:"domainName,omitempty"`
	DomainSid         *string `json:"domainSid,omitempty"`
	ForestName        *string `json:"forestName,omitempty"`
	NetBiosDomainName *string `json:"netBiosDomainName,omitempty"`
}

// Deprecated version of BlobRestoreParameters_STATUS. Use v1beta20210401.BlobRestoreParameters_STATUS instead
type BlobRestoreParameters_STATUSARM struct {
	BlobRanges    []BlobRestoreRange_STATUSARM `json:"blobRanges,omitempty"`
	TimeToRestore *string                      `json:"timeToRestore,omitempty"`
}

// Deprecated version of EncryptionIdentity_STATUS. Use v1beta20210401.EncryptionIdentity_STATUS instead
type EncryptionIdentity_STATUSARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// Deprecated version of EncryptionServices_STATUS. Use v1beta20210401.EncryptionServices_STATUS instead
type EncryptionServices_STATUSARM struct {
	Blob  *EncryptionService_STATUSARM `json:"blob,omitempty"`
	File  *EncryptionService_STATUSARM `json:"file,omitempty"`
	Queue *EncryptionService_STATUSARM `json:"queue,omitempty"`
	Table *EncryptionService_STATUSARM `json:"table,omitempty"`
}

// Deprecated version of IPRule_STATUS. Use v1beta20210401.IPRule_STATUS instead
type IPRule_STATUSARM struct {
	Action *IPRuleSTATUSAction `json:"action,omitempty"`
	Value  *string             `json:"value,omitempty"`
}

// Deprecated version of KeyVaultProperties_STATUS. Use v1beta20210401.KeyVaultProperties_STATUS instead
type KeyVaultProperties_STATUSARM struct {
	CurrentVersionedKeyIdentifier *string `json:"currentVersionedKeyIdentifier,omitempty"`
	Keyname                       *string `json:"keyname,omitempty"`
	Keyvaulturi                   *string `json:"keyvaulturi,omitempty"`
	Keyversion                    *string `json:"keyversion,omitempty"`
	LastKeyRotationTimestamp      *string `json:"lastKeyRotationTimestamp,omitempty"`
}

// Deprecated version of ResourceAccessRule_STATUS. Use v1beta20210401.ResourceAccessRule_STATUS instead
type ResourceAccessRule_STATUSARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
	TenantId   *string `json:"tenantId,omitempty"`
}

// Deprecated version of StorageAccountInternetEndpoints_STATUS. Use v1beta20210401.StorageAccountInternetEndpoints_STATUS instead
type StorageAccountInternetEndpoints_STATUSARM struct {
	Blob *string `json:"blob,omitempty"`
	Dfs  *string `json:"dfs,omitempty"`
	File *string `json:"file,omitempty"`
	Web  *string `json:"web,omitempty"`
}

// Deprecated version of StorageAccountMicrosoftEndpoints_STATUS. Use v1beta20210401.StorageAccountMicrosoftEndpoints_STATUS instead
type StorageAccountMicrosoftEndpoints_STATUSARM struct {
	Blob  *string `json:"blob,omitempty"`
	Dfs   *string `json:"dfs,omitempty"`
	File  *string `json:"file,omitempty"`
	Queue *string `json:"queue,omitempty"`
	Table *string `json:"table,omitempty"`
	Web   *string `json:"web,omitempty"`
}

// Deprecated version of VirtualNetworkRule_STATUS. Use v1beta20210401.VirtualNetworkRule_STATUS instead
type VirtualNetworkRule_STATUSARM struct {
	Action *VirtualNetworkRuleSTATUSAction `json:"action,omitempty"`
	Id     *string                         `json:"id,omitempty"`
	State  *VirtualNetworkRuleSTATUSState  `json:"state,omitempty"`
}

// Deprecated version of BlobRestoreRange_STATUS. Use v1beta20210401.BlobRestoreRange_STATUS instead
type BlobRestoreRange_STATUSARM struct {
	EndRange   *string `json:"endRange,omitempty"`
	StartRange *string `json:"startRange,omitempty"`
}

// Deprecated version of EncryptionService_STATUS. Use v1beta20210401.EncryptionService_STATUS instead
type EncryptionService_STATUSARM struct {
	Enabled         *bool                           `json:"enabled,omitempty"`
	KeyType         *EncryptionServiceSTATUSKeyType `json:"keyType,omitempty"`
	LastEnabledTime *string                         `json:"lastEnabledTime,omitempty"`
}
