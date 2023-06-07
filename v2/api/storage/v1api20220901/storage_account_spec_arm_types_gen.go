// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccount_Spec_ARM struct {
	// ExtendedLocation: Optional. Set the extended location of the resource. If not set, the storage account will be created
	// in Azure main region. Otherwise it will be created in the specified extended location
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation"`

	// Identity: The identity of the resource.
	Identity *Identity_ARM `json:"identity"`

	// Kind: Required. Indicates the type of storage account.
	Kind *StorageAccount_Kind_Spec `json:"kind"`

	// Location: Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure
	// Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is
	// created, but if an identical geo region is specified on update, the request will succeed.
	Location *string `json:"location"`
	Name     string  `json:"name,omitempty"`

	// Properties: The parameters used to create the storage account.
	Properties *StorageAccountPropertiesCreateParameters_ARM `json:"properties"`

	// Sku: Required. Gets or sets the SKU name.
	Sku *Sku_ARM `json:"sku"`

	// Tags: Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping
	// this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key
	// with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags map[string]string `json:"tags"`
}

var _ genruntime.ARMResourceSpec = &StorageAccount_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
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

// The complex type of the extended location.
type ExtendedLocation_ARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType `json:"type"`
}

// Identity for the resource.
type Identity_ARM struct {
	// Type: The identity type.
	Type                   *Identity_Type                             `json:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// The SKU of the storage account.
type Sku_ARM struct {
	// Name: The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called
	//  accountType.
	Name *SkuName `json:"name"`

	// Tier: The SKU tier. This is based on the SKU name.
	Tier *Tier `json:"tier"`
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

// The parameters used to create the storage account.
type StorageAccountPropertiesCreateParameters_ARM struct {
	// AccessTier: Required for storage accounts where kind = BlobStorage. The access tier is used for billing. The 'Premium'
	// access tier is the default value for premium block blobs storage account type and it cannot be changed for the premium
	// block blobs storage account type.
	AccessTier *StorageAccountPropertiesCreateParameters_AccessTier `json:"accessTier"`

	// AllowBlobPublicAccess: Allow or disallow public access to all blobs or containers in the storage account. The default
	// interpretation is true for this property.
	AllowBlobPublicAccess *bool `json:"allowBlobPublicAccess"`

	// AllowCrossTenantReplication: Allow or disallow cross AAD tenant object replication. The default interpretation is true
	// for this property.
	AllowCrossTenantReplication *bool `json:"allowCrossTenantReplication"`

	// AllowSharedKeyAccess: Indicates whether the storage account permits requests to be authorized with the account access
	// key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure
	// Active Directory (Azure AD). The default value is null, which is equivalent to true.
	AllowSharedKeyAccess *bool `json:"allowSharedKeyAccess"`

	// AllowedCopyScope: Restrict copy to and from Storage Accounts within an AAD tenant or with Private Links to the same VNet.
	AllowedCopyScope *StorageAccountPropertiesCreateParameters_AllowedCopyScope `json:"allowedCopyScope"`

	// AzureFilesIdentityBasedAuthentication: Provides the identity based authentication settings for Azure Files.
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication_ARM `json:"azureFilesIdentityBasedAuthentication"`

	// CustomDomain: User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported
	// per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name
	// property.
	CustomDomain *CustomDomain_ARM `json:"customDomain"`

	// DefaultToOAuthAuthentication: A boolean flag which indicates whether the default authentication is OAuth or not. The
	// default interpretation is false for this property.
	DefaultToOAuthAuthentication *bool `json:"defaultToOAuthAuthentication"`

	// DnsEndpointType: Allows you to specify the type of endpoint. Set this to AzureDNSZone to create a large number of
	// accounts in a single subscription, which creates accounts in an Azure DNS Zone and the endpoint URL will have an
	// alphanumeric DNS Zone identifier.
	DnsEndpointType *StorageAccountPropertiesCreateParameters_DnsEndpointType `json:"dnsEndpointType"`

	// Encryption: Encryption settings to be used for server-side encryption for the storage account.
	Encryption *Encryption_ARM `json:"encryption"`

	// ImmutableStorageWithVersioning: The property is immutable and can only be set to true at the account creation time. When
	// set to true, it enables object level immutability for all the new containers in the account by default.
	ImmutableStorageWithVersioning *ImmutableStorageAccount_ARM `json:"immutableStorageWithVersioning"`

	// IsHnsEnabled: Account HierarchicalNamespace enabled if sets to true.
	IsHnsEnabled *bool `json:"isHnsEnabled"`

	// IsLocalUserEnabled: Enables local users feature, if set to true
	IsLocalUserEnabled *bool `json:"isLocalUserEnabled"`

	// IsNfsV3Enabled: NFS 3.0 protocol support enabled if set to true.
	IsNfsV3Enabled *bool `json:"isNfsV3Enabled"`

	// IsSftpEnabled: Enables Secure File Transfer Protocol, if set to true
	IsSftpEnabled *bool `json:"isSftpEnabled"`

	// KeyPolicy: KeyPolicy assigned to the storage account.
	KeyPolicy *KeyPolicy_ARM `json:"keyPolicy"`

	// LargeFileSharesState: Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
	LargeFileSharesState *StorageAccountPropertiesCreateParameters_LargeFileSharesState `json:"largeFileSharesState"`

	// MinimumTlsVersion: Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS
	// 1.0 for this property.
	MinimumTlsVersion *StorageAccountPropertiesCreateParameters_MinimumTlsVersion `json:"minimumTlsVersion"`

	// NetworkAcls: Network rule set
	NetworkAcls *NetworkRuleSet_ARM `json:"networkAcls"`

	// PublicNetworkAccess: Allow or disallow public network access to Storage Account. Value is optional but if passed in,
	// must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *StorageAccountPropertiesCreateParameters_PublicNetworkAccess `json:"publicNetworkAccess"`

	// RoutingPreference: Maintains information about the network routing choice opted by the user for data transfer
	RoutingPreference *RoutingPreference_ARM `json:"routingPreference"`

	// SasPolicy: SasPolicy assigned to the storage account.
	SasPolicy *SasPolicy_ARM `json:"sasPolicy"`

	// SupportsHttpsTrafficOnly: Allows https traffic only to storage service if sets to true. The default value is true since
	// API version 2019-04-01.
	SupportsHttpsTrafficOnly *bool `json:"supportsHttpsTrafficOnly"`
}

// Settings for Azure Files identity based authentication.
type AzureFilesIdentityBasedAuthentication_ARM struct {
	// ActiveDirectoryProperties: Required if directoryServiceOptions are AD, optional if they are AADKERB.
	ActiveDirectoryProperties *ActiveDirectoryProperties_ARM `json:"activeDirectoryProperties"`

	// DefaultSharePermission: Default share permission for users using Kerberos authentication if RBAC role is not assigned.
	DefaultSharePermission *AzureFilesIdentityBasedAuthentication_DefaultSharePermission `json:"defaultSharePermission"`

	// DirectoryServiceOptions: Indicates the directory service used. Note that this enum may be extended in the future.
	DirectoryServiceOptions *AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions `json:"directoryServiceOptions"`
}

// The custom domain assigned to this storage account. This can be set via Update.
type CustomDomain_ARM struct {
	// Name: Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
	Name *string `json:"name"`

	// UseSubDomainName: Indicates whether indirect CName validation is enabled. Default value is false. This should only be
	// set on updates.
	UseSubDomainName *bool `json:"useSubDomainName"`
}

// The encryption settings on the storage account.
type Encryption_ARM struct {
	// Identity: The identity to be used with service-side encryption at rest.
	Identity *EncryptionIdentity_ARM `json:"identity"`

	// KeySource: The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage,
	// Microsoft.Keyvault
	KeySource *Encryption_KeySource `json:"keySource"`

	// Keyvaultproperties: Properties provided by key vault.
	Keyvaultproperties *KeyVaultProperties_ARM `json:"keyvaultproperties"`

	// RequireInfrastructureEncryption: A boolean indicating whether or not the service applies a secondary layer of encryption
	// with platform managed keys for data at rest.
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption"`

	// Services: List of services which support encryption.
	Services *EncryptionServices_ARM `json:"services"`
}

// The type of extendedLocation.
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

// This property enables and defines account-level immutability. Enabling the feature auto-enables Blob Versioning.
type ImmutableStorageAccount_ARM struct {
	// Enabled: A boolean flag which enables account-level immutability. All the containers under such an account have
	// object-level immutability enabled by default.
	Enabled *bool `json:"enabled"`

	// ImmutabilityPolicy: Specifies the default account-level immutability policy which is inherited and applied to objects
	// that do not possess an explicit immutability policy at the object level. The object-level immutability policy has higher
	// precedence than the container-level immutability policy, which has a higher precedence than the account-level
	// immutability policy.
	ImmutabilityPolicy *AccountImmutabilityPolicyProperties_ARM `json:"immutabilityPolicy"`
}

// KeyPolicy assigned to the storage account.
type KeyPolicy_ARM struct {
	// KeyExpirationPeriodInDays: The key expiration period in days.
	KeyExpirationPeriodInDays *int `json:"keyExpirationPeriodInDays"`
}

// Network rule set
type NetworkRuleSet_ARM struct {
	// Bypass: Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of
	// Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics.
	Bypass *NetworkRuleSet_Bypass `json:"bypass"`

	// DefaultAction: Specifies the default action of allow or deny when no other rules match.
	DefaultAction *NetworkRuleSet_DefaultAction `json:"defaultAction"`

	// IpRules: Sets the IP ACL rules
	IpRules []IPRule_ARM `json:"ipRules"`

	// ResourceAccessRules: Sets the resource access rules
	ResourceAccessRules []ResourceAccessRule_ARM `json:"resourceAccessRules"`

	// VirtualNetworkRules: Sets the virtual network rules
	VirtualNetworkRules []VirtualNetworkRule_ARM `json:"virtualNetworkRules"`
}

// Routing preference defines the type of network, either microsoft or internet routing to be used to deliver the user
// data, the default option is microsoft routing
type RoutingPreference_ARM struct {
	// PublishInternetEndpoints: A boolean flag which indicates whether internet routing storage endpoints are to be published
	PublishInternetEndpoints *bool `json:"publishInternetEndpoints"`

	// PublishMicrosoftEndpoints: A boolean flag which indicates whether microsoft routing storage endpoints are to be published
	PublishMicrosoftEndpoints *bool `json:"publishMicrosoftEndpoints"`

	// RoutingChoice: Routing Choice defines the kind of network routing opted by the user.
	RoutingChoice *RoutingPreference_RoutingChoice `json:"routingChoice"`
}

// SasPolicy assigned to the storage account.
type SasPolicy_ARM struct {
	// ExpirationAction: The SAS expiration action. Can only be Log.
	ExpirationAction *SasPolicy_ExpirationAction `json:"expirationAction"`

	// SasExpirationPeriod: The SAS expiration period, DD.HH:MM:SS.
	SasExpirationPeriod *string `json:"sasExpirationPeriod"`
}

// The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called
// accountType.
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

// The SKU tier. This is based on the SKU name.
// +kubebuilder:validation:Enum={"Premium","Standard"}
type Tier string

const (
	Tier_Premium  = Tier("Premium")
	Tier_Standard = Tier("Standard")
)

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// This defines account-level immutability policy properties.
type AccountImmutabilityPolicyProperties_ARM struct {
	// AllowProtectedAppendWrites: This property can only be changed for disabled and unlocked time-based retention policies.
	// When enabled, new blocks can be written to an append blob while maintaining immutability protection and compliance. Only
	// new blocks can be added and any existing blocks cannot be modified or deleted.
	AllowProtectedAppendWrites *bool `json:"allowProtectedAppendWrites"`

	// ImmutabilityPeriodSinceCreationInDays: The immutability period for the blobs in the container since the policy creation,
	// in days.
	ImmutabilityPeriodSinceCreationInDays *int `json:"immutabilityPeriodSinceCreationInDays"`

	// State: The ImmutabilityPolicy state defines the mode of the policy. Disabled state disables the policy, Unlocked state
	// allows increase and decrease of immutability retention time and also allows toggling allowProtectedAppendWrites
	// property, Locked state only allows the increase of the immutability retention time. A policy can only be created in a
	// Disabled or Unlocked state and can be toggled between the two states. Only a policy in an Unlocked state can transition
	// to a Locked state which cannot be reverted.
	State *AccountImmutabilityPolicyProperties_State `json:"state"`
}

// Settings properties for Active Directory (AD).
type ActiveDirectoryProperties_ARM struct {
	// AccountType: Specifies the Active Directory account type for Azure Storage.
	AccountType *ActiveDirectoryProperties_AccountType `json:"accountType"`

	// AzureStorageSid: Specifies the security identifier (SID) for Azure Storage.
	AzureStorageSid *string `json:"azureStorageSid"`

	// DomainGuid: Specifies the domain GUID.
	DomainGuid *string `json:"domainGuid"`

	// DomainName: Specifies the primary domain that the AD DNS server is authoritative for.
	DomainName *string `json:"domainName"`

	// DomainSid: Specifies the security identifier (SID).
	DomainSid *string `json:"domainSid"`

	// ForestName: Specifies the Active Directory forest to get.
	ForestName *string `json:"forestName"`

	// NetBiosDomainName: Specifies the NetBIOS domain name.
	NetBiosDomainName *string `json:"netBiosDomainName"`

	// SamAccountName: Specifies the Active Directory SAMAccountName for Azure Storage.
	SamAccountName *string `json:"samAccountName"`
}

// Encryption identity for the storage account.
type EncryptionIdentity_ARM struct {
	// FederatedIdentityClientId: ClientId of the multi-tenant application to be used in conjunction with the user-assigned
	// identity for cross-tenant customer-managed-keys server-side encryption on the storage account.
	FederatedIdentityClientId *string `json:"federatedIdentityClientId"`
	UserAssignedIdentity      *string `json:"userAssignedIdentity,omitempty"`
}

// A list of services that support encryption.
type EncryptionServices_ARM struct {
	// Blob: The encryption function of the blob storage service.
	Blob *EncryptionService_ARM `json:"blob"`

	// File: The encryption function of the file storage service.
	File *EncryptionService_ARM `json:"file"`

	// Queue: The encryption function of the queue storage service.
	Queue *EncryptionService_ARM `json:"queue"`

	// Table: The encryption function of the table storage service.
	Table *EncryptionService_ARM `json:"table"`
}

// IP rule with specific IP or IP range in CIDR format.
type IPRule_ARM struct {
	// Action: The action of IP ACL rule.
	Action *IPRule_Action `json:"action"`

	// Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	Value *string `json:"value"`
}

// Properties of key vault.
type KeyVaultProperties_ARM struct {
	// Keyname: The name of KeyVault key.
	Keyname *string `json:"keyname"`

	// Keyvaulturi: The Uri of KeyVault.
	Keyvaulturi *string `json:"keyvaulturi"`

	// Keyversion: The version of KeyVault key.
	Keyversion *string `json:"keyversion"`
}

// Resource Access Rule.
type ResourceAccessRule_ARM struct {
	ResourceId *string `json:"resourceId,omitempty"`

	// TenantId: Tenant Id
	TenantId *string `json:"tenantId"`
}

// Virtual Network rule.
type VirtualNetworkRule_ARM struct {
	// Action: The action of virtual network rule.
	Action *VirtualNetworkRule_Action `json:"action"`
	Id     *string                    `json:"id,omitempty"`

	// State: Gets the state of virtual network rule.
	State *VirtualNetworkRule_State `json:"state"`
}

// A service that allows server-side encryption to be used.
type EncryptionService_ARM struct {
	// Enabled: A boolean indicating whether or not the service encrypts the data as it is stored. Encryption at rest is
	// enabled by default today and cannot be disabled.
	Enabled *bool `json:"enabled"`

	// KeyType: Encryption key type to be used for the encryption service. 'Account' key type implies that an account-scoped
	// encryption key will be used. 'Service' key type implies that a default service key is used.
	KeyType *EncryptionService_KeyType `json:"keyType"`
}
