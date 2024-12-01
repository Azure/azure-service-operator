// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// The configuration store along with all resource properties. The Configuration Store will have all information to begin
// utilizing it.
type ConfigurationStore_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: The managed identity information, if configured.
	Identity *ResourceIdentity_STATUS `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of a configuration store.
	Properties *ConfigurationStoreProperties_STATUS `json:"properties,omitempty"`

	// Sku: The sku of the configuration store.
	Sku *Sku_STATUS `json:"sku,omitempty"`

	// SystemData: Resource system metadata.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of a configuration store.
type ConfigurationStoreProperties_STATUS struct {
	// CreateMode: Indicates whether the configuration store need to be recovered.
	CreateMode *ConfigurationStoreProperties_CreateMode_STATUS `json:"createMode,omitempty"`

	// CreationDate: The creation date of configuration store.
	CreationDate *string `json:"creationDate,omitempty"`

	// DisableLocalAuth: Disables all authentication methods other than AAD authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// EnablePurgeProtection: Property specifying whether protection against purge is enabled for this configuration store.
	EnablePurgeProtection *bool `json:"enablePurgeProtection,omitempty"`

	// Encryption: The encryption settings of the configuration store.
	Encryption *EncryptionProperties_STATUS `json:"encryption,omitempty"`

	// Endpoint: The DNS endpoint where the configuration store API will be available.
	Endpoint *string `json:"endpoint,omitempty"`

	// PrivateEndpointConnections: The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections []PrivateEndpointConnectionReference_STATUS `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: The provisioning state of the configuration store.
	ProvisioningState *ConfigurationStoreProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Control permission for data plane traffic coming from public networks while private endpoint is
	// enabled.
	PublicNetworkAccess *ConfigurationStoreProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// SoftDeleteRetentionInDays: The amount of time in days that the configuration store will be retained when it is soft
	// deleted.
	SoftDeleteRetentionInDays *int `json:"softDeleteRetentionInDays,omitempty"`
}

// An identity that can be associated with a resource.
type ResourceIdentity_STATUS struct {
	// PrincipalId: The principal id of the identity. This property will only be provided for a system-assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant id associated with the resource's identity. This property will only be provided for a
	// system-assigned identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created
	// identity and a set of user-assigned identities. The type 'None' will remove any identities.
	Type *ResourceIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user-assigned identities associated with the resource. The user-assigned identity
	// dictionary keys will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]UserIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Describes a configuration store SKU.
type Sku_STATUS struct {
	// Name: The SKU name of the configuration store.
	Name *string `json:"name,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

type ConfigurationStoreProperties_CreateMode_STATUS string

const (
	ConfigurationStoreProperties_CreateMode_STATUS_Default = ConfigurationStoreProperties_CreateMode_STATUS("Default")
	ConfigurationStoreProperties_CreateMode_STATUS_Recover = ConfigurationStoreProperties_CreateMode_STATUS("Recover")
)

// Mapping from string to ConfigurationStoreProperties_CreateMode_STATUS
var configurationStoreProperties_CreateMode_STATUS_Values = map[string]ConfigurationStoreProperties_CreateMode_STATUS{
	"default": ConfigurationStoreProperties_CreateMode_STATUS_Default,
	"recover": ConfigurationStoreProperties_CreateMode_STATUS_Recover,
}

type ConfigurationStoreProperties_ProvisioningState_STATUS string

const (
	ConfigurationStoreProperties_ProvisioningState_STATUS_Canceled  = ConfigurationStoreProperties_ProvisioningState_STATUS("Canceled")
	ConfigurationStoreProperties_ProvisioningState_STATUS_Creating  = ConfigurationStoreProperties_ProvisioningState_STATUS("Creating")
	ConfigurationStoreProperties_ProvisioningState_STATUS_Deleting  = ConfigurationStoreProperties_ProvisioningState_STATUS("Deleting")
	ConfigurationStoreProperties_ProvisioningState_STATUS_Failed    = ConfigurationStoreProperties_ProvisioningState_STATUS("Failed")
	ConfigurationStoreProperties_ProvisioningState_STATUS_Succeeded = ConfigurationStoreProperties_ProvisioningState_STATUS("Succeeded")
	ConfigurationStoreProperties_ProvisioningState_STATUS_Updating  = ConfigurationStoreProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to ConfigurationStoreProperties_ProvisioningState_STATUS
var configurationStoreProperties_ProvisioningState_STATUS_Values = map[string]ConfigurationStoreProperties_ProvisioningState_STATUS{
	"canceled":  ConfigurationStoreProperties_ProvisioningState_STATUS_Canceled,
	"creating":  ConfigurationStoreProperties_ProvisioningState_STATUS_Creating,
	"deleting":  ConfigurationStoreProperties_ProvisioningState_STATUS_Deleting,
	"failed":    ConfigurationStoreProperties_ProvisioningState_STATUS_Failed,
	"succeeded": ConfigurationStoreProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  ConfigurationStoreProperties_ProvisioningState_STATUS_Updating,
}

type ConfigurationStoreProperties_PublicNetworkAccess_STATUS string

const (
	ConfigurationStoreProperties_PublicNetworkAccess_STATUS_Disabled = ConfigurationStoreProperties_PublicNetworkAccess_STATUS("Disabled")
	ConfigurationStoreProperties_PublicNetworkAccess_STATUS_Enabled  = ConfigurationStoreProperties_PublicNetworkAccess_STATUS("Enabled")
)

// Mapping from string to ConfigurationStoreProperties_PublicNetworkAccess_STATUS
var configurationStoreProperties_PublicNetworkAccess_STATUS_Values = map[string]ConfigurationStoreProperties_PublicNetworkAccess_STATUS{
	"disabled": ConfigurationStoreProperties_PublicNetworkAccess_STATUS_Disabled,
	"enabled":  ConfigurationStoreProperties_PublicNetworkAccess_STATUS_Enabled,
}

// The encryption settings for a configuration store.
type EncryptionProperties_STATUS struct {
	// KeyVaultProperties: Key vault properties.
	KeyVaultProperties *KeyVaultProperties_STATUS `json:"keyVaultProperties,omitempty"`
}

// A reference to a related private endpoint connection.
type PrivateEndpointConnectionReference_STATUS struct {
	// Id: The resource ID.
	Id *string `json:"id,omitempty"`
}

type ResourceIdentity_Type_STATUS string

const (
	ResourceIdentity_Type_STATUS_None                       = ResourceIdentity_Type_STATUS("None")
	ResourceIdentity_Type_STATUS_SystemAssigned             = ResourceIdentity_Type_STATUS("SystemAssigned")
	ResourceIdentity_Type_STATUS_SystemAssignedUserAssigned = ResourceIdentity_Type_STATUS("SystemAssigned, UserAssigned")
	ResourceIdentity_Type_STATUS_UserAssigned               = ResourceIdentity_Type_STATUS("UserAssigned")
)

// Mapping from string to ResourceIdentity_Type_STATUS
var resourceIdentity_Type_STATUS_Values = map[string]ResourceIdentity_Type_STATUS{
	"none":                         ResourceIdentity_Type_STATUS_None,
	"systemassigned":               ResourceIdentity_Type_STATUS_SystemAssigned,
	"systemassigned, userassigned": ResourceIdentity_Type_STATUS_SystemAssignedUserAssigned,
	"userassigned":                 ResourceIdentity_Type_STATUS_UserAssigned,
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}

// A resource identity that is managed by the user of the service.
type UserIdentity_STATUS struct {
	// ClientId: The client ID of the user-assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal ID of the user-assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

// Settings concerning key vault encryption for a configuration store.
type KeyVaultProperties_STATUS struct {
	// IdentityClientId: The client id of the identity which will be used to access key vault.
	IdentityClientId *string `json:"identityClientId,omitempty"`

	// KeyIdentifier: The URI of the key vault key used to encrypt data.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}
