// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

// Deprecated version of Namespace_STATUS. Use v1beta20210101preview.Namespace_STATUS instead
type Namespace_STATUS_ARM struct {
	Id         *string                           `json:"id,omitempty"`
	Identity   *Identity_STATUS_ARM              `json:"identity,omitempty"`
	Location   *string                           `json:"location,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Properties *SBNamespaceProperties_STATUS_ARM `json:"properties,omitempty"`
	Sku        *SBSku_STATUS_ARM                 `json:"sku,omitempty"`
	SystemData *SystemData_STATUS_ARM            `json:"systemData,omitempty"`
	Tags       map[string]string                 `json:"tags,omitempty"`
	Type       *string                           `json:"type,omitempty"`
}

// Deprecated version of Identity_STATUS. Use v1beta20210101preview.Identity_STATUS instead
type Identity_STATUS_ARM struct {
	PrincipalId            *string                               `json:"principalId,omitempty"`
	TenantId               *string                               `json:"tenantId,omitempty"`
	Type                   *Identity_Type_STATUS                 `json:"type,omitempty"`
	UserAssignedIdentities map[string]DictionaryValue_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of SBNamespaceProperties_STATUS. Use v1beta20210101preview.SBNamespaceProperties_STATUS instead
type SBNamespaceProperties_STATUS_ARM struct {
	CreatedAt                  *string                                `json:"createdAt,omitempty"`
	Encryption                 *Encryption_STATUS_ARM                 `json:"encryption,omitempty"`
	MetricId                   *string                                `json:"metricId,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *string                                `json:"provisioningState,omitempty"`
	ServiceBusEndpoint         *string                                `json:"serviceBusEndpoint,omitempty"`
	Status                     *string                                `json:"status,omitempty"`
	UpdatedAt                  *string                                `json:"updatedAt,omitempty"`
	ZoneRedundant              *bool                                  `json:"zoneRedundant,omitempty"`
}

// Deprecated version of SBSku_STATUS. Use v1beta20210101preview.SBSku_STATUS instead
type SBSku_STATUS_ARM struct {
	Capacity *int               `json:"capacity,omitempty"`
	Name     *SBSku_Name_STATUS `json:"name,omitempty"`
	Tier     *SBSku_Tier_STATUS `json:"tier,omitempty"`
}

// Deprecated version of SystemData_STATUS. Use v1beta20210101preview.SystemData_STATUS instead
type SystemData_STATUS_ARM struct {
	CreatedAt          *string                               `json:"createdAt,omitempty"`
	CreatedBy          *string                               `json:"createdBy,omitempty"`
	CreatedByType      *SystemData_CreatedByType_STATUS      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                               `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                               `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// Deprecated version of DictionaryValue_STATUS. Use v1beta20210101preview.DictionaryValue_STATUS instead
type DictionaryValue_STATUS_ARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of Encryption_STATUS. Use v1beta20210101preview.Encryption_STATUS instead
type Encryption_STATUS_ARM struct {
	KeySource                       *Encryption_KeySource_STATUS    `json:"keySource,omitempty"`
	KeyVaultProperties              []KeyVaultProperties_STATUS_ARM `json:"keyVaultProperties,omitempty"`
	RequireInfrastructureEncryption *bool                           `json:"requireInfrastructureEncryption,omitempty"`
}

// Deprecated version of Identity_Type_STATUS. Use v1beta20210101preview.Identity_Type_STATUS instead
type Identity_Type_STATUS string

const (
	Identity_Type_STATUS_None                       = Identity_Type_STATUS("None")
	Identity_Type_STATUS_SystemAssigned             = Identity_Type_STATUS("SystemAssigned")
	Identity_Type_STATUS_SystemAssignedUserAssigned = Identity_Type_STATUS("SystemAssigned, UserAssigned")
	Identity_Type_STATUS_UserAssigned               = Identity_Type_STATUS("UserAssigned")
)

// Deprecated version of PrivateEndpointConnection_STATUS. Use v1beta20210101preview.PrivateEndpointConnection_STATUS instead
type PrivateEndpointConnection_STATUS_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of SBSku_Name_STATUS. Use v1beta20210101preview.SBSku_Name_STATUS instead
type SBSku_Name_STATUS string

const (
	SBSku_Name_STATUS_Basic    = SBSku_Name_STATUS("Basic")
	SBSku_Name_STATUS_Premium  = SBSku_Name_STATUS("Premium")
	SBSku_Name_STATUS_Standard = SBSku_Name_STATUS("Standard")
)

// Deprecated version of SBSku_Tier_STATUS. Use v1beta20210101preview.SBSku_Tier_STATUS instead
type SBSku_Tier_STATUS string

const (
	SBSku_Tier_STATUS_Basic    = SBSku_Tier_STATUS("Basic")
	SBSku_Tier_STATUS_Premium  = SBSku_Tier_STATUS("Premium")
	SBSku_Tier_STATUS_Standard = SBSku_Tier_STATUS("Standard")
)

// Deprecated version of SystemData_CreatedByType_STATUS. Use v1beta20210101preview.SystemData_CreatedByType_STATUS instead
type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Deprecated version of SystemData_LastModifiedByType_STATUS. Use
// v1beta20210101preview.SystemData_LastModifiedByType_STATUS instead
type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Deprecated version of KeyVaultProperties_STATUS. Use v1beta20210101preview.KeyVaultProperties_STATUS instead
type KeyVaultProperties_STATUS_ARM struct {
	Identity    *UserAssignedIdentityProperties_STATUS_ARM `json:"identity,omitempty"`
	KeyName     *string                                    `json:"keyName,omitempty"`
	KeyVaultUri *string                                    `json:"keyVaultUri,omitempty"`
	KeyVersion  *string                                    `json:"keyVersion,omitempty"`
}

// Deprecated version of UserAssignedIdentityProperties_STATUS. Use v1beta20210101preview.UserAssignedIdentityProperties_STATUS instead
type UserAssignedIdentityProperties_STATUS_ARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
