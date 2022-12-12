// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccount_Spec_ARM struct {
	// ExtendedLocation: Optional. Set the extended location of the resource. If not set, the storage account will be created
	// in Azure main region. Otherwise it will be created in the specified extended location
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`

	// Identity: The identity of the resource.
	Identity *Identity_ARM `json:"identity,omitempty"`

	// Kind: Required. Indicates the type of storage account.
	Kind *StorageAccount_Kind_Spec `json:"kind,omitempty"`

	// Location: Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure
	// Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is
	// created, but if an identical geo region is specified on update, the request will succeed.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The parameters used to create the storage account.
	Properties *StorageAccountPropertiesCreateParameters_ARM `json:"properties,omitempty"`

	// Sku: Required. Gets or sets the SKU name.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping
	// this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key
	// with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccount_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (account StorageAccount_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (account *StorageAccount_Spec_ARM) GetName() string {
	return account.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts"
func (account *StorageAccount_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts"
}

type ExtendedLocation_ARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType `json:"type,omitempty"`
}

type Identity_ARM struct {
	// Type: The identity type.
	Type *Identity_Type `json:"type,omitempty"`
}

type Sku_ARM struct {
	Name *SkuName `json:"name,omitempty"`
	Tier *Tier    `json:"tier,omitempty"`
}

// +kubebuilder:validation:Enum={"BlobStorage","BlockBlobStorage","FileStorage","Storage","StorageV2"}
type StorageAccount_Kind_Spec string

const (
	StorageAccount_Kind_Spec_BlobStorage      = StorageAccount_Kind_Spec("BlobStorage")
	StorageAccount_Kind_Spec_BlockBlobStorage = StorageAccount_Kind_Spec("BlockBlobStorage")
	StorageAccount_Kind_Spec_FileStorage      = StorageAccount_Kind_Spec("FileStorage")
	StorageAccount_Kind_Spec_Storage          = StorageAccount_Kind_Spec("Storage")
	StorageAccount_Kind_Spec_StorageV2        = StorageAccount_Kind_Spec("StorageV2")
)

type StorageAccountPropertiesCreateParameters_ARM struct {
	// AccessTier: Required for storage accounts where kind = BlobStorage. The access tier used for billing.
	AccessTier *StorageAccountPropertiesCreateParameters_AccessTier `json:"accessTier,omitempty"`

	// AllowBlobPublicAccess: Allow or disallow public access to all blobs or containers in the storage account. The default
	// interpretation is true for this property.
	AllowBlobPublicAccess *bool `json:"allowBlobPublicAccess,omitempty"`

	// AllowCrossTenantReplication: Allow or disallow cross AAD tenant object replication. The default interpretation is true
	// for this property.
	AllowCrossTenantReplication *bool `json:"allowCrossTenantReplication,omitempty"`

	// AllowSharedKeyAccess: Indicates whether the storage account permits requests to be authorized with the account access
	// key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure
	// Active Directory (Azure AD). The default value is null, which is equivalent to true.
	AllowSharedKeyAccess *bool `json:"allowSharedKeyAccess,omitempty"`

	// AzureFilesIdentityBasedAuthentication: Provides the identity based authentication settings for Azure Files.
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication_ARM `json:"azureFilesIdentityBasedAuthentication,omitempty"`

	// CustomDomain: User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported
	// per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name
	// property.
	CustomDomain *CustomDomain_ARM `json:"customDomain,omitempty"`

	// Encryption: Not applicable. Azure Storage encryption is enabled for all storage accounts and cannot be disabled.
	Encryption *Encryption_ARM `json:"encryption,omitempty"`

	// IsHnsEnabled: Account HierarchicalNamespace enabled if sets to true.
	IsHnsEnabled *bool `json:"isHnsEnabled,omitempty"`

	// IsNfsV3Enabled: NFS 3.0 protocol support enabled if set to true.
	IsNfsV3Enabled *bool `json:"isNfsV3Enabled,omitempty"`

	// KeyPolicy: KeyPolicy assigned to the storage account.
	KeyPolicy *KeyPolicy_ARM `json:"keyPolicy,omitempty"`

	// LargeFileSharesState: Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
	LargeFileSharesState *StorageAccountPropertiesCreateParameters_LargeFileSharesState `json:"largeFileSharesState,omitempty"`

	// MinimumTlsVersion: Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS
	// 1.0 for this property.
	MinimumTlsVersion *StorageAccountPropertiesCreateParameters_MinimumTlsVersion `json:"minimumTlsVersion,omitempty"`

	// NetworkAcls: Network rule set
	NetworkAcls *NetworkRuleSet_ARM `json:"networkAcls,omitempty"`

	// RoutingPreference: Maintains information about the network routing choice opted by the user for data transfer
	RoutingPreference *RoutingPreference_ARM `json:"routingPreference,omitempty"`

	// SasPolicy: SasPolicy assigned to the storage account.
	SasPolicy *SasPolicy_ARM `json:"sasPolicy,omitempty"`

	// SupportsHttpsTrafficOnly: Allows https traffic only to storage service if sets to true. The default value is true since
	// API version 2019-04-01.
	SupportsHttpsTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`
}

type AzureFilesIdentityBasedAuthentication_ARM struct {
	// ActiveDirectoryProperties: Required if choose AD.
	ActiveDirectoryProperties *ActiveDirectoryProperties_ARM `json:"activeDirectoryProperties,omitempty"`

	// DefaultSharePermission: Default share permission for users using Kerberos authentication if RBAC role is not assigned.
	DefaultSharePermission *AzureFilesIdentityBasedAuthentication_DefaultSharePermission `json:"defaultSharePermission,omitempty"`

	// DirectoryServiceOptions: Indicates the directory service used.
	DirectoryServiceOptions *AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions `json:"directoryServiceOptions,omitempty"`
}

type CustomDomain_ARM struct {
	// Name: Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
	Name *string `json:"name,omitempty"`

	// UseSubDomainName: Indicates whether indirect CName validation is enabled. Default value is false. This should only be
	// set on updates.
	UseSubDomainName *bool `json:"useSubDomainName,omitempty"`
}

type Encryption_ARM struct {
	// Identity: The identity to be used with service-side encryption at rest.
	Identity *EncryptionIdentity_ARM `json:"identity,omitempty"`

	// KeySource: The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage,
	// Microsoft.Keyvault
	KeySource *Encryption_KeySource `json:"keySource,omitempty"`

	// Keyvaultproperties: Properties provided by key vault.
	Keyvaultproperties *KeyVaultProperties_ARM `json:"keyvaultproperties,omitempty"`

	// RequireInfrastructureEncryption: A boolean indicating whether or not the service applies a secondary layer of encryption
	// with platform managed keys for data at rest.
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`

	// Services: List of services which support encryption.
	Services *EncryptionServices_ARM `json:"services,omitempty"`
}

// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationType_EdgeZone = ExtendedLocationType("EdgeZone")

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type Identity_Type string

const (
	Identity_Type_None                       = Identity_Type("None")
	Identity_Type_SystemAssigned             = Identity_Type("SystemAssigned")
	Identity_Type_SystemAssignedUserAssigned = Identity_Type("SystemAssigned,UserAssigned")
	Identity_Type_UserAssigned               = Identity_Type("UserAssigned")
)

type KeyPolicy_ARM struct {
	// KeyExpirationPeriodInDays: The key expiration period in days.
	KeyExpirationPeriodInDays *int `json:"keyExpirationPeriodInDays,omitempty"`
}

type NetworkRuleSet_ARM struct {
	// Bypass: Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of
	// Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics.
	Bypass *NetworkRuleSet_Bypass `json:"bypass,omitempty"`

	// DefaultAction: Specifies the default action of allow or deny when no other rules match.
	DefaultAction *NetworkRuleSet_DefaultAction `json:"defaultAction,omitempty"`

	// IpRules: Sets the IP ACL rules
	IpRules []IPRule_ARM `json:"ipRules,omitempty"`

	// ResourceAccessRules: Sets the resource access rules
	ResourceAccessRules []ResourceAccessRule_ARM `json:"resourceAccessRules,omitempty"`

	// VirtualNetworkRules: Sets the virtual network rules
	VirtualNetworkRules []VirtualNetworkRule_ARM `json:"virtualNetworkRules,omitempty"`
}

type RoutingPreference_ARM struct {
	// PublishInternetEndpoints: A boolean flag which indicates whether internet routing storage endpoints are to be published
	PublishInternetEndpoints *bool `json:"publishInternetEndpoints,omitempty"`

	// PublishMicrosoftEndpoints: A boolean flag which indicates whether microsoft routing storage endpoints are to be published
	PublishMicrosoftEndpoints *bool `json:"publishMicrosoftEndpoints,omitempty"`

	// RoutingChoice: Routing Choice defines the kind of network routing opted by the user.
	RoutingChoice *RoutingPreference_RoutingChoice `json:"routingChoice,omitempty"`
}

type SasPolicy_ARM struct {
	// ExpirationAction: The SAS expiration action. Can only be Log.
	ExpirationAction *SasPolicy_ExpirationAction `json:"expirationAction,omitempty"`

	// SasExpirationPeriod: The SAS expiration period, DD.HH:MM:SS.
	SasExpirationPeriod *string `json:"sasExpirationPeriod,omitempty"`
}

// +kubebuilder:validation:Enum={"Premium_LRS","Premium_ZRS","Standard_GRS","Standard_GZRS","Standard_LRS","Standard_RAGRS","Standard_RAGZRS","Standard_ZRS"}
type SkuName string

const (
	SkuName_Premium_LRS     = SkuName("Premium_LRS")
	SkuName_Premium_ZRS     = SkuName("Premium_ZRS")
	SkuName_Standard_GRS    = SkuName("Standard_GRS")
	SkuName_Standard_GZRS   = SkuName("Standard_GZRS")
	SkuName_Standard_LRS    = SkuName("Standard_LRS")
	SkuName_Standard_RAGRS  = SkuName("Standard_RAGRS")
	SkuName_Standard_RAGZRS = SkuName("Standard_RAGZRS")
	SkuName_Standard_ZRS    = SkuName("Standard_ZRS")
)

// +kubebuilder:validation:Enum={"Premium","Standard"}
type Tier string

const (
	Tier_Premium  = Tier("Premium")
	Tier_Standard = Tier("Standard")
)

type ActiveDirectoryProperties_ARM struct {
	// AzureStorageSid: Specifies the security identifier (SID) for Azure Storage.
	AzureStorageSid *string `json:"azureStorageSid,omitempty"`

	// DomainGuid: Specifies the domain GUID.
	DomainGuid *string `json:"domainGuid,omitempty"`

	// DomainName: Specifies the primary domain that the AD DNS server is authoritative for.
	DomainName *string `json:"domainName,omitempty"`

	// DomainSid: Specifies the security identifier (SID).
	DomainSid *string `json:"domainSid,omitempty"`

	// ForestName: Specifies the Active Directory forest to get.
	ForestName *string `json:"forestName,omitempty"`

	// NetBiosDomainName: Specifies the NetBIOS domain name.
	NetBiosDomainName *string `json:"netBiosDomainName,omitempty"`
}

type EncryptionIdentity_ARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

type EncryptionServices_ARM struct {
	// Blob: The encryption function of the blob storage service.
	Blob *EncryptionService_ARM `json:"blob,omitempty"`

	// File: The encryption function of the file storage service.
	File *EncryptionService_ARM `json:"file,omitempty"`

	// Queue: The encryption function of the queue storage service.
	Queue *EncryptionService_ARM `json:"queue,omitempty"`

	// Table: The encryption function of the table storage service.
	Table *EncryptionService_ARM `json:"table,omitempty"`
}

type IPRule_ARM struct {
	// Action: The action of IP ACL rule.
	Action *IPRule_Action `json:"action,omitempty"`

	// Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	Value *string `json:"value,omitempty"`
}

type KeyVaultProperties_ARM struct {
	// Keyname: The name of KeyVault key.
	Keyname *string `json:"keyname,omitempty"`

	// Keyvaulturi: The Uri of KeyVault.
	Keyvaulturi *string `json:"keyvaulturi,omitempty"`

	// Keyversion: The version of KeyVault key.
	Keyversion *string `json:"keyversion,omitempty"`
}

type ResourceAccessRule_ARM struct {
	ResourceId *string `json:"resourceId,omitempty"`

	// TenantId: Tenant Id
	TenantId *string `json:"tenantId,omitempty"`
}

type VirtualNetworkRule_ARM struct {
	// Action: The action of virtual network rule.
	Action *VirtualNetworkRule_Action `json:"action,omitempty"`
	Id     *string                    `json:"id,omitempty"`

	// State: Gets the state of virtual network rule.
	State *VirtualNetworkRule_State `json:"state,omitempty"`
}

type EncryptionService_ARM struct {
	// Enabled: A boolean indicating whether or not the service encrypts the data as it is stored.
	Enabled *bool `json:"enabled,omitempty"`

	// KeyType: Encryption key type to be used for the encryption service. 'Account' key type implies that an account-scoped
	// encryption key will be used. 'Service' key type implies that a default service key is used.
	KeyType *EncryptionService_KeyType `json:"keyType,omitempty"`
}
