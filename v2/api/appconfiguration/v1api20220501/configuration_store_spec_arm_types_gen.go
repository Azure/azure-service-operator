// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ConfigurationStore_Spec_ARM struct {
	// Identity: The managed identity information, if configured.
	Identity *ResourceIdentity_ARM `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The properties of a configuration store.
	Properties *ConfigurationStoreProperties_ARM `json:"properties,omitempty"`

	// Sku: The sku of the configuration store.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// SystemData: Resource system metadata.
	SystemData *SystemData_ARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ConfigurationStore_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-05-01"
func (store ConfigurationStore_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (store *ConfigurationStore_Spec_ARM) GetName() string {
	return store.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.AppConfiguration/configurationStores"
func (store *ConfigurationStore_Spec_ARM) GetType() string {
	return "Microsoft.AppConfiguration/configurationStores"
}

// The properties of a configuration store.
type ConfigurationStoreProperties_ARM struct {
	// CreateMode: Indicates whether the configuration store need to be recovered.
	CreateMode *ConfigurationStoreProperties_CreateMode `json:"createMode,omitempty"`

	// DisableLocalAuth: Disables all authentication methods other than AAD authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// EnablePurgeProtection: Property specifying whether protection against purge is enabled for this configuration store.
	EnablePurgeProtection *bool `json:"enablePurgeProtection,omitempty"`

	// Encryption: The encryption settings of the configuration store.
	Encryption *EncryptionProperties_ARM `json:"encryption,omitempty"`

	// PublicNetworkAccess: Control permission for data plane traffic coming from public networks while private endpoint is
	// enabled.
	PublicNetworkAccess *ConfigurationStoreProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// SoftDeleteRetentionInDays: The amount of time in days that the configuration store will be retained when it is soft
	// deleted.
	SoftDeleteRetentionInDays *int `json:"softDeleteRetentionInDays,omitempty"`
}

// An identity that can be associated with a resource.
type ResourceIdentity_ARM struct {
	// Type: The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created
	// identity and a set of user-assigned identities. The type 'None' will remove any identities.
	Type                   *ResourceIdentity_Type                     `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Describes a configuration store SKU.
type Sku_ARM struct {
	// Name: The SKU name of the configuration store.
	Name *string `json:"name,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

// The encryption settings for a configuration store.
type EncryptionProperties_ARM struct {
	// KeyVaultProperties: Key vault properties.
	KeyVaultProperties *KeyVaultProperties_ARM `json:"keyVaultProperties,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type ResourceIdentity_Type string

const (
	ResourceIdentity_Type_None                       = ResourceIdentity_Type("None")
	ResourceIdentity_Type_SystemAssigned             = ResourceIdentity_Type("SystemAssigned")
	ResourceIdentity_Type_SystemAssignedUserAssigned = ResourceIdentity_Type("SystemAssigned, UserAssigned")
	ResourceIdentity_Type_UserAssigned               = ResourceIdentity_Type("UserAssigned")
)

// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_CreatedByType string

const (
	SystemData_CreatedByType_Application     = SystemData_CreatedByType("Application")
	SystemData_CreatedByType_Key             = SystemData_CreatedByType("Key")
	SystemData_CreatedByType_ManagedIdentity = SystemData_CreatedByType("ManagedIdentity")
	SystemData_CreatedByType_User            = SystemData_CreatedByType("User")
)

// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_LastModifiedByType string

const (
	SystemData_LastModifiedByType_Application     = SystemData_LastModifiedByType("Application")
	SystemData_LastModifiedByType_Key             = SystemData_LastModifiedByType("Key")
	SystemData_LastModifiedByType_ManagedIdentity = SystemData_LastModifiedByType("ManagedIdentity")
	SystemData_LastModifiedByType_User            = SystemData_LastModifiedByType("User")
)

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// Settings concerning key vault encryption for a configuration store.
type KeyVaultProperties_ARM struct {
	// IdentityClientId: The client id of the identity which will be used to access key vault.
	IdentityClientId *string `json:"identityClientId,omitempty"`

	// KeyIdentifier: The URI of the key vault key used to encrypt data.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}
