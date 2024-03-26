// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230701

// Resource information with extended details.
type Vault_STATUS_ARM struct {
	// Id: Fully qualified identifier of the key vault resource.
	Id *string `json:"id,omitempty"`

	// Location: Azure location of the key vault resource.
	Location *string `json:"location,omitempty"`

	// Name: Name of the key vault resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the vault
	Properties *VaultProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: System metadata for the key vault.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: Tags assigned to the key vault resource.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type of the key vault resource.
	Type *string `json:"type,omitempty"`
}

// Metadata pertaining to creation and last modification of the key vault resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of the key vault resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the key vault resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the key vault resource.
	CreatedByType *IdentityType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of the key vault resource last modification (UTC).
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the key vault resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the key vault resource.
	LastModifiedByType *IdentityType_STATUS `json:"lastModifiedByType,omitempty"`
}

// Properties of the vault
type VaultProperties_STATUS_ARM struct {
	// AccessPolicies: An array of 0 to 1024 identities that have access to the key vault. All identities in the array must use
	// the same tenant ID as the key vault's tenant ID. When `createMode` is set to `recover`, access policies are not
	// required. Otherwise, access policies are required.
	AccessPolicies []AccessPolicyEntry_STATUS_ARM `json:"accessPolicies,omitempty"`

	// CreateMode: The vault's create mode to indicate whether the vault need to be recovered or not.
	CreateMode *VaultProperties_CreateMode_STATUS `json:"createMode,omitempty"`

	// EnablePurgeProtection: Property specifying whether protection against purge is enabled for this vault. Setting this
	// property to true activates protection against purge for this vault and its content - only the Key Vault service may
	// initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this
	// functionality is irreversible - that is, the property does not accept false as its value.
	EnablePurgeProtection *bool `json:"enablePurgeProtection,omitempty"`

	// EnableRbacAuthorization: Property that controls how data actions are authorized. When true, the key vault will use Role
	// Based Access Control (RBAC) for authorization of data actions, and the access policies specified in vault properties
	// will be  ignored. When false, the key vault will use the access policies specified in vault properties, and any policy
	// stored on Azure Resource Manager will be ignored. If null or not specified, the vault is created with the default value
	// of false. Note that management actions are always authorized with RBAC.
	EnableRbacAuthorization *bool `json:"enableRbacAuthorization,omitempty"`

	// EnableSoftDelete: Property to specify whether the 'soft delete' functionality is enabled for this key vault. If it's not
	// set to any value(true or false) when creating new key vault, it will be set to true by default. Once set to true, it
	// cannot be reverted to false.
	EnableSoftDelete *bool `json:"enableSoftDelete,omitempty"`

	// EnabledForDeployment: Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored
	// as secrets from the key vault.
	EnabledForDeployment *bool `json:"enabledForDeployment,omitempty"`

	// EnabledForDiskEncryption: Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the
	// vault and unwrap keys.
	EnabledForDiskEncryption *bool `json:"enabledForDiskEncryption,omitempty"`

	// EnabledForTemplateDeployment: Property to specify whether Azure Resource Manager is permitted to retrieve secrets from
	// the key vault.
	EnabledForTemplateDeployment *bool `json:"enabledForTemplateDeployment,omitempty"`

	// HsmPoolResourceId: The resource id of HSM Pool.
	HsmPoolResourceId *string `json:"hsmPoolResourceId,omitempty"`

	// NetworkAcls: Rules governing the accessibility of the key vault from specific network locations.
	NetworkAcls *NetworkRuleSet_STATUS_ARM `json:"networkAcls,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections associated with the key vault.
	PrivateEndpointConnections []PrivateEndpointConnectionItem_STATUS_ARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Provisioning state of the vault.
	ProvisioningState *VaultProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Property to specify whether the vault will accept traffic from public internet. If set to
	// 'disabled' all traffic except private endpoint traffic and that that originates from trusted services will be blocked.
	// This will override the set firewall rules, meaning that even if the firewall rules are present we will not honor the
	// rules.
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty"`

	// Sku: SKU details
	Sku *Sku_STATUS_ARM `json:"sku,omitempty"`

	// SoftDeleteRetentionInDays: softDelete data retention days. It accepts >=7 and <=90.
	SoftDeleteRetentionInDays *int `json:"softDeleteRetentionInDays,omitempty"`

	// TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId *string `json:"tenantId,omitempty"`

	// VaultUri: The URI of the vault for performing operations on keys and secrets.
	VaultUri *string `json:"vaultUri,omitempty"`
}

// An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key
// vault's tenant ID.
type AccessPolicyEntry_STATUS_ARM struct {
	// ApplicationId:  Application ID of the client making request on behalf of a principal
	ApplicationId *string `json:"applicationId,omitempty"`

	// ObjectId: The object ID of a user, service principal or security group in the Azure Active Directory tenant for the
	// vault. The object ID must be unique for the list of access policies.
	ObjectId *string `json:"objectId,omitempty"`

	// Permissions: Permissions the identity has for keys, secrets and certificates.
	Permissions *Permissions_STATUS_ARM `json:"permissions,omitempty"`

	// TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId *string `json:"tenantId,omitempty"`
}

// The type of identity.
type IdentityType_STATUS string

const (
	IdentityType_STATUS_Application     = IdentityType_STATUS("Application")
	IdentityType_STATUS_Key             = IdentityType_STATUS("Key")
	IdentityType_STATUS_ManagedIdentity = IdentityType_STATUS("ManagedIdentity")
	IdentityType_STATUS_User            = IdentityType_STATUS("User")
)

// Mapping from string to IdentityType_STATUS
var identityType_STATUS_Values = map[string]IdentityType_STATUS{
	"application":     IdentityType_STATUS_Application,
	"key":             IdentityType_STATUS_Key,
	"managedidentity": IdentityType_STATUS_ManagedIdentity,
	"user":            IdentityType_STATUS_User,
}

// A set of rules governing the network accessibility of a vault.
type NetworkRuleSet_STATUS_ARM struct {
	// Bypass: Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'.  If not specified the
	// default is 'AzureServices'.
	Bypass *NetworkRuleSet_Bypass_STATUS `json:"bypass,omitempty"`

	// DefaultAction: The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after
	// the bypass property has been evaluated.
	DefaultAction *NetworkRuleSet_DefaultAction_STATUS `json:"defaultAction,omitempty"`

	// IpRules: The list of IP address rules.
	IpRules []IPRule_STATUS_ARM `json:"ipRules,omitempty"`

	// VirtualNetworkRules: The list of virtual network rules.
	VirtualNetworkRules []VirtualNetworkRule_STATUS_ARM `json:"virtualNetworkRules,omitempty"`
}

// Private endpoint connection item.
type PrivateEndpointConnectionItem_STATUS_ARM struct {
	// Etag: Modified whenever there is a change in the state of private endpoint connection.
	Etag *string `json:"etag,omitempty"`

	// Id: Id of private endpoint connection.
	Id *string `json:"id,omitempty"`

	// Properties: Private endpoint connection properties.
	Properties *PrivateEndpointConnectionProperties_STATUS_ARM `json:"properties,omitempty"`
}

// SKU details
type Sku_STATUS_ARM struct {
	// Family: SKU family name
	Family *Sku_Family_STATUS `json:"family,omitempty"`

	// Name: SKU name to specify whether the key vault is a standard vault or a premium vault.
	Name *Sku_Name_STATUS `json:"name,omitempty"`
}

type VaultProperties_CreateMode_STATUS string

const (
	VaultProperties_CreateMode_STATUS_CreateOrRecover = VaultProperties_CreateMode_STATUS("createOrRecover")
	VaultProperties_CreateMode_STATUS_Default         = VaultProperties_CreateMode_STATUS("default")
	VaultProperties_CreateMode_STATUS_PurgeThenCreate = VaultProperties_CreateMode_STATUS("purgeThenCreate")
	VaultProperties_CreateMode_STATUS_Recover         = VaultProperties_CreateMode_STATUS("recover")
)

// Mapping from string to VaultProperties_CreateMode_STATUS
var vaultProperties_CreateMode_STATUS_Values = map[string]VaultProperties_CreateMode_STATUS{
	"createorrecover": VaultProperties_CreateMode_STATUS_CreateOrRecover,
	"default":         VaultProperties_CreateMode_STATUS_Default,
	"purgethencreate": VaultProperties_CreateMode_STATUS_PurgeThenCreate,
	"recover":         VaultProperties_CreateMode_STATUS_Recover,
}

type VaultProperties_ProvisioningState_STATUS string

const (
	VaultProperties_ProvisioningState_STATUS_RegisteringDns = VaultProperties_ProvisioningState_STATUS("RegisteringDns")
	VaultProperties_ProvisioningState_STATUS_Succeeded      = VaultProperties_ProvisioningState_STATUS("Succeeded")
)

// Mapping from string to VaultProperties_ProvisioningState_STATUS
var vaultProperties_ProvisioningState_STATUS_Values = map[string]VaultProperties_ProvisioningState_STATUS{
	"registeringdns": VaultProperties_ProvisioningState_STATUS_RegisteringDns,
	"succeeded":      VaultProperties_ProvisioningState_STATUS_Succeeded,
}

// A rule governing the accessibility of a vault from a specific ip address or ip range.
type IPRule_STATUS_ARM struct {
	// Value: An IPv4 address range in CIDR notation, such as '124.56.78.91' (simple IP address) or '124.56.78.0/24' (all
	// addresses that start with 124.56.78).
	Value *string `json:"value,omitempty"`
}

type NetworkRuleSet_Bypass_STATUS string

const (
	NetworkRuleSet_Bypass_STATUS_AzureServices = NetworkRuleSet_Bypass_STATUS("AzureServices")
	NetworkRuleSet_Bypass_STATUS_None          = NetworkRuleSet_Bypass_STATUS("None")
)

// Mapping from string to NetworkRuleSet_Bypass_STATUS
var networkRuleSet_Bypass_STATUS_Values = map[string]NetworkRuleSet_Bypass_STATUS{
	"azureservices": NetworkRuleSet_Bypass_STATUS_AzureServices,
	"none":          NetworkRuleSet_Bypass_STATUS_None,
}

type NetworkRuleSet_DefaultAction_STATUS string

const (
	NetworkRuleSet_DefaultAction_STATUS_Allow = NetworkRuleSet_DefaultAction_STATUS("Allow")
	NetworkRuleSet_DefaultAction_STATUS_Deny  = NetworkRuleSet_DefaultAction_STATUS("Deny")
)

// Mapping from string to NetworkRuleSet_DefaultAction_STATUS
var networkRuleSet_DefaultAction_STATUS_Values = map[string]NetworkRuleSet_DefaultAction_STATUS{
	"allow": NetworkRuleSet_DefaultAction_STATUS_Allow,
	"deny":  NetworkRuleSet_DefaultAction_STATUS_Deny,
}

// Permissions the identity has for keys, secrets, certificates and storage.
type Permissions_STATUS_ARM struct {
	// Certificates: Permissions to certificates
	Certificates []Permissions_Certificates_STATUS `json:"certificates,omitempty"`

	// Keys: Permissions to keys
	Keys []Permissions_Keys_STATUS `json:"keys,omitempty"`

	// Secrets: Permissions to secrets
	Secrets []Permissions_Secrets_STATUS `json:"secrets,omitempty"`

	// Storage: Permissions to storage accounts
	Storage []Permissions_Storage_STATUS `json:"storage,omitempty"`
}

// Properties of the private endpoint connection resource.
type PrivateEndpointConnectionProperties_STATUS_ARM struct {
	// PrivateEndpoint: Properties of the private endpoint object.
	PrivateEndpoint *PrivateEndpoint_STATUS_ARM `json:"privateEndpoint,omitempty"`

	// PrivateLinkServiceConnectionState: Approval state of the private link connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState_STATUS_ARM `json:"privateLinkServiceConnectionState,omitempty"`

	// ProvisioningState: Provisioning state of the private endpoint connection.
	ProvisioningState *PrivateEndpointConnectionProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

type Sku_Family_STATUS string

const Sku_Family_STATUS_A = Sku_Family_STATUS("A")

// Mapping from string to Sku_Family_STATUS
var sku_Family_STATUS_Values = map[string]Sku_Family_STATUS{
	"a": Sku_Family_STATUS_A,
}

type Sku_Name_STATUS string

const (
	Sku_Name_STATUS_Premium  = Sku_Name_STATUS("premium")
	Sku_Name_STATUS_Standard = Sku_Name_STATUS("standard")
)

// Mapping from string to Sku_Name_STATUS
var sku_Name_STATUS_Values = map[string]Sku_Name_STATUS{
	"premium":  Sku_Name_STATUS_Premium,
	"standard": Sku_Name_STATUS_Standard,
}

// A rule governing the accessibility of a vault from a specific virtual network.
type VirtualNetworkRule_STATUS_ARM struct {
	// Id: Full resource id of a vnet subnet, such as
	// '/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1'.
	Id *string `json:"id,omitempty"`

	// IgnoreMissingVnetServiceEndpoint: Property to specify whether NRP will ignore the check if parent subnet has
	// serviceEndpoints configured.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty"`
}

type Permissions_Certificates_STATUS string

const (
	Permissions_Certificates_STATUS_All            = Permissions_Certificates_STATUS("all")
	Permissions_Certificates_STATUS_Backup         = Permissions_Certificates_STATUS("backup")
	Permissions_Certificates_STATUS_Create         = Permissions_Certificates_STATUS("create")
	Permissions_Certificates_STATUS_Delete         = Permissions_Certificates_STATUS("delete")
	Permissions_Certificates_STATUS_Deleteissuers  = Permissions_Certificates_STATUS("deleteissuers")
	Permissions_Certificates_STATUS_Get            = Permissions_Certificates_STATUS("get")
	Permissions_Certificates_STATUS_Getissuers     = Permissions_Certificates_STATUS("getissuers")
	Permissions_Certificates_STATUS_Import         = Permissions_Certificates_STATUS("import")
	Permissions_Certificates_STATUS_List           = Permissions_Certificates_STATUS("list")
	Permissions_Certificates_STATUS_Listissuers    = Permissions_Certificates_STATUS("listissuers")
	Permissions_Certificates_STATUS_Managecontacts = Permissions_Certificates_STATUS("managecontacts")
	Permissions_Certificates_STATUS_Manageissuers  = Permissions_Certificates_STATUS("manageissuers")
	Permissions_Certificates_STATUS_Purge          = Permissions_Certificates_STATUS("purge")
	Permissions_Certificates_STATUS_Recover        = Permissions_Certificates_STATUS("recover")
	Permissions_Certificates_STATUS_Restore        = Permissions_Certificates_STATUS("restore")
	Permissions_Certificates_STATUS_Setissuers     = Permissions_Certificates_STATUS("setissuers")
	Permissions_Certificates_STATUS_Update         = Permissions_Certificates_STATUS("update")
)

// Mapping from string to Permissions_Certificates_STATUS
var permissions_Certificates_STATUS_Values = map[string]Permissions_Certificates_STATUS{
	"all":            Permissions_Certificates_STATUS_All,
	"backup":         Permissions_Certificates_STATUS_Backup,
	"create":         Permissions_Certificates_STATUS_Create,
	"delete":         Permissions_Certificates_STATUS_Delete,
	"deleteissuers":  Permissions_Certificates_STATUS_Deleteissuers,
	"get":            Permissions_Certificates_STATUS_Get,
	"getissuers":     Permissions_Certificates_STATUS_Getissuers,
	"import":         Permissions_Certificates_STATUS_Import,
	"list":           Permissions_Certificates_STATUS_List,
	"listissuers":    Permissions_Certificates_STATUS_Listissuers,
	"managecontacts": Permissions_Certificates_STATUS_Managecontacts,
	"manageissuers":  Permissions_Certificates_STATUS_Manageissuers,
	"purge":          Permissions_Certificates_STATUS_Purge,
	"recover":        Permissions_Certificates_STATUS_Recover,
	"restore":        Permissions_Certificates_STATUS_Restore,
	"setissuers":     Permissions_Certificates_STATUS_Setissuers,
	"update":         Permissions_Certificates_STATUS_Update,
}

type Permissions_Keys_STATUS string

const (
	Permissions_Keys_STATUS_All               = Permissions_Keys_STATUS("all")
	Permissions_Keys_STATUS_Backup            = Permissions_Keys_STATUS("backup")
	Permissions_Keys_STATUS_Create            = Permissions_Keys_STATUS("create")
	Permissions_Keys_STATUS_Decrypt           = Permissions_Keys_STATUS("decrypt")
	Permissions_Keys_STATUS_Delete            = Permissions_Keys_STATUS("delete")
	Permissions_Keys_STATUS_Encrypt           = Permissions_Keys_STATUS("encrypt")
	Permissions_Keys_STATUS_Get               = Permissions_Keys_STATUS("get")
	Permissions_Keys_STATUS_Getrotationpolicy = Permissions_Keys_STATUS("getrotationpolicy")
	Permissions_Keys_STATUS_Import            = Permissions_Keys_STATUS("import")
	Permissions_Keys_STATUS_List              = Permissions_Keys_STATUS("list")
	Permissions_Keys_STATUS_Purge             = Permissions_Keys_STATUS("purge")
	Permissions_Keys_STATUS_Recover           = Permissions_Keys_STATUS("recover")
	Permissions_Keys_STATUS_Release           = Permissions_Keys_STATUS("release")
	Permissions_Keys_STATUS_Restore           = Permissions_Keys_STATUS("restore")
	Permissions_Keys_STATUS_Rotate            = Permissions_Keys_STATUS("rotate")
	Permissions_Keys_STATUS_Setrotationpolicy = Permissions_Keys_STATUS("setrotationpolicy")
	Permissions_Keys_STATUS_Sign              = Permissions_Keys_STATUS("sign")
	Permissions_Keys_STATUS_UnwrapKey         = Permissions_Keys_STATUS("unwrapKey")
	Permissions_Keys_STATUS_Update            = Permissions_Keys_STATUS("update")
	Permissions_Keys_STATUS_Verify            = Permissions_Keys_STATUS("verify")
	Permissions_Keys_STATUS_WrapKey           = Permissions_Keys_STATUS("wrapKey")
)

// Mapping from string to Permissions_Keys_STATUS
var permissions_Keys_STATUS_Values = map[string]Permissions_Keys_STATUS{
	"all":               Permissions_Keys_STATUS_All,
	"backup":            Permissions_Keys_STATUS_Backup,
	"create":            Permissions_Keys_STATUS_Create,
	"decrypt":           Permissions_Keys_STATUS_Decrypt,
	"delete":            Permissions_Keys_STATUS_Delete,
	"encrypt":           Permissions_Keys_STATUS_Encrypt,
	"get":               Permissions_Keys_STATUS_Get,
	"getrotationpolicy": Permissions_Keys_STATUS_Getrotationpolicy,
	"import":            Permissions_Keys_STATUS_Import,
	"list":              Permissions_Keys_STATUS_List,
	"purge":             Permissions_Keys_STATUS_Purge,
	"recover":           Permissions_Keys_STATUS_Recover,
	"release":           Permissions_Keys_STATUS_Release,
	"restore":           Permissions_Keys_STATUS_Restore,
	"rotate":            Permissions_Keys_STATUS_Rotate,
	"setrotationpolicy": Permissions_Keys_STATUS_Setrotationpolicy,
	"sign":              Permissions_Keys_STATUS_Sign,
	"unwrapkey":         Permissions_Keys_STATUS_UnwrapKey,
	"update":            Permissions_Keys_STATUS_Update,
	"verify":            Permissions_Keys_STATUS_Verify,
	"wrapkey":           Permissions_Keys_STATUS_WrapKey,
}

type Permissions_Secrets_STATUS string

const (
	Permissions_Secrets_STATUS_All     = Permissions_Secrets_STATUS("all")
	Permissions_Secrets_STATUS_Backup  = Permissions_Secrets_STATUS("backup")
	Permissions_Secrets_STATUS_Delete  = Permissions_Secrets_STATUS("delete")
	Permissions_Secrets_STATUS_Get     = Permissions_Secrets_STATUS("get")
	Permissions_Secrets_STATUS_List    = Permissions_Secrets_STATUS("list")
	Permissions_Secrets_STATUS_Purge   = Permissions_Secrets_STATUS("purge")
	Permissions_Secrets_STATUS_Recover = Permissions_Secrets_STATUS("recover")
	Permissions_Secrets_STATUS_Restore = Permissions_Secrets_STATUS("restore")
	Permissions_Secrets_STATUS_Set     = Permissions_Secrets_STATUS("set")
)

// Mapping from string to Permissions_Secrets_STATUS
var permissions_Secrets_STATUS_Values = map[string]Permissions_Secrets_STATUS{
	"all":     Permissions_Secrets_STATUS_All,
	"backup":  Permissions_Secrets_STATUS_Backup,
	"delete":  Permissions_Secrets_STATUS_Delete,
	"get":     Permissions_Secrets_STATUS_Get,
	"list":    Permissions_Secrets_STATUS_List,
	"purge":   Permissions_Secrets_STATUS_Purge,
	"recover": Permissions_Secrets_STATUS_Recover,
	"restore": Permissions_Secrets_STATUS_Restore,
	"set":     Permissions_Secrets_STATUS_Set,
}

type Permissions_Storage_STATUS string

const (
	Permissions_Storage_STATUS_All           = Permissions_Storage_STATUS("all")
	Permissions_Storage_STATUS_Backup        = Permissions_Storage_STATUS("backup")
	Permissions_Storage_STATUS_Delete        = Permissions_Storage_STATUS("delete")
	Permissions_Storage_STATUS_Deletesas     = Permissions_Storage_STATUS("deletesas")
	Permissions_Storage_STATUS_Get           = Permissions_Storage_STATUS("get")
	Permissions_Storage_STATUS_Getsas        = Permissions_Storage_STATUS("getsas")
	Permissions_Storage_STATUS_List          = Permissions_Storage_STATUS("list")
	Permissions_Storage_STATUS_Listsas       = Permissions_Storage_STATUS("listsas")
	Permissions_Storage_STATUS_Purge         = Permissions_Storage_STATUS("purge")
	Permissions_Storage_STATUS_Recover       = Permissions_Storage_STATUS("recover")
	Permissions_Storage_STATUS_Regeneratekey = Permissions_Storage_STATUS("regeneratekey")
	Permissions_Storage_STATUS_Restore       = Permissions_Storage_STATUS("restore")
	Permissions_Storage_STATUS_Set           = Permissions_Storage_STATUS("set")
	Permissions_Storage_STATUS_Setsas        = Permissions_Storage_STATUS("setsas")
	Permissions_Storage_STATUS_Update        = Permissions_Storage_STATUS("update")
)

// Mapping from string to Permissions_Storage_STATUS
var permissions_Storage_STATUS_Values = map[string]Permissions_Storage_STATUS{
	"all":           Permissions_Storage_STATUS_All,
	"backup":        Permissions_Storage_STATUS_Backup,
	"delete":        Permissions_Storage_STATUS_Delete,
	"deletesas":     Permissions_Storage_STATUS_Deletesas,
	"get":           Permissions_Storage_STATUS_Get,
	"getsas":        Permissions_Storage_STATUS_Getsas,
	"list":          Permissions_Storage_STATUS_List,
	"listsas":       Permissions_Storage_STATUS_Listsas,
	"purge":         Permissions_Storage_STATUS_Purge,
	"recover":       Permissions_Storage_STATUS_Recover,
	"regeneratekey": Permissions_Storage_STATUS_Regeneratekey,
	"restore":       Permissions_Storage_STATUS_Restore,
	"set":           Permissions_Storage_STATUS_Set,
	"setsas":        Permissions_Storage_STATUS_Setsas,
	"update":        Permissions_Storage_STATUS_Update,
}

// Private endpoint object properties.
type PrivateEndpoint_STATUS_ARM struct {
	// Id: Full identifier of the private endpoint resource.
	Id *string `json:"id,omitempty"`
}

// An object that represents the approval state of the private link connection.
type PrivateLinkServiceConnectionState_STATUS_ARM struct {
	// ActionsRequired: A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *PrivateLinkServiceConnectionState_ActionsRequired_STATUS `json:"actionsRequired,omitempty"`

	// Description: The reason for approval or rejection.
	Description *string `json:"description,omitempty"`

	// Status: Indicates whether the connection has been approved, rejected or removed by the key vault owner.
	Status *PrivateEndpointServiceConnectionStatus_STATUS `json:"status,omitempty"`
}
