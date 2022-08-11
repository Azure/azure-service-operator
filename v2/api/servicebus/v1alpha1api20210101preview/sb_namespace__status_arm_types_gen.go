// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

// Deprecated version of SBNamespace_Status. Use v1beta20210101preview.SBNamespace_Status instead
type SBNamespace_StatusARM struct {
	Id         *string                          `json:"id,omitempty"`
	Identity   *Identity_StatusARM              `json:"identity,omitempty"`
	Location   *string                          `json:"location,omitempty"`
	Name       *string                          `json:"name,omitempty"`
	Properties *SBNamespaceProperties_StatusARM `json:"properties,omitempty"`
	Sku        *SBSku_StatusARM                 `json:"sku,omitempty"`
	SystemData *SystemData_StatusARM            `json:"systemData,omitempty"`
	Tags       map[string]string                `json:"tags,omitempty"`
	Type       *string                          `json:"type,omitempty"`
}

// Deprecated version of Identity_Status. Use v1beta20210101preview.Identity_Status instead
type Identity_StatusARM struct {
	PrincipalId            *string                              `json:"principalId,omitempty"`
	TenantId               *string                              `json:"tenantId,omitempty"`
	Type                   *IdentityStatusType                  `json:"type,omitempty"`
	UserAssignedIdentities map[string]DictionaryValue_StatusARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of SBNamespaceProperties_Status. Use v1beta20210101preview.SBNamespaceProperties_Status instead
type SBNamespaceProperties_StatusARM struct {
	CreatedAt                  *string                                                   `json:"createdAt,omitempty"`
	Encryption                 *Encryption_StatusARM                                     `json:"encryption,omitempty"`
	MetricId                   *string                                                   `json:"metricId,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *string                                                   `json:"provisioningState,omitempty"`
	ServiceBusEndpoint         *string                                                   `json:"serviceBusEndpoint,omitempty"`
	Status                     *string                                                   `json:"status,omitempty"`
	UpdatedAt                  *string                                                   `json:"updatedAt,omitempty"`
	ZoneRedundant              *bool                                                     `json:"zoneRedundant,omitempty"`
}

// Deprecated version of SBSku_Status. Use v1beta20210101preview.SBSku_Status instead
type SBSku_StatusARM struct {
	Capacity *int             `json:"capacity,omitempty"`
	Name     *SBSkuStatusName `json:"name,omitempty"`
	Tier     *SBSkuStatusTier `json:"tier,omitempty"`
}

// Deprecated version of SystemData_Status. Use v1beta20210101preview.SystemData_Status instead
type SystemData_StatusARM struct {
	CreatedAt          *string                             `json:"createdAt,omitempty"`
	CreatedBy          *string                             `json:"createdBy,omitempty"`
	CreatedByType      *SystemDataStatusCreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                             `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                             `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemDataStatusLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

// Deprecated version of DictionaryValue_Status. Use v1beta20210101preview.DictionaryValue_Status instead
type DictionaryValue_StatusARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of Encryption_Status. Use v1beta20210101preview.Encryption_Status instead
type Encryption_StatusARM struct {
	KeySource                       *EncryptionStatusKeySource     `json:"keySource,omitempty"`
	KeyVaultProperties              []KeyVaultProperties_StatusARM `json:"keyVaultProperties,omitempty"`
	RequireInfrastructureEncryption *bool                          `json:"requireInfrastructureEncryption,omitempty"`
}

// Deprecated version of IdentityStatusType. Use v1beta20210101preview.IdentityStatusType instead
type IdentityStatusType string

const (
	IdentityStatusType_None                       = IdentityStatusType("None")
	IdentityStatusType_SystemAssigned             = IdentityStatusType("SystemAssigned")
	IdentityStatusType_SystemAssignedUserAssigned = IdentityStatusType("SystemAssigned, UserAssigned")
	IdentityStatusType_UserAssigned               = IdentityStatusType("UserAssigned")
)

// Deprecated version of PrivateEndpointConnection_Status_SubResourceEmbedded. Use v1beta20210101preview.PrivateEndpointConnection_Status_SubResourceEmbedded instead
type PrivateEndpointConnection_Status_SubResourceEmbeddedARM struct {
	Id         *string               `json:"id,omitempty"`
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`
}

// Deprecated version of SBSkuStatusName. Use v1beta20210101preview.SBSkuStatusName instead
type SBSkuStatusName string

const (
	SBSkuStatusName_Basic    = SBSkuStatusName("Basic")
	SBSkuStatusName_Premium  = SBSkuStatusName("Premium")
	SBSkuStatusName_Standard = SBSkuStatusName("Standard")
)

// Deprecated version of SBSkuStatusTier. Use v1beta20210101preview.SBSkuStatusTier instead
type SBSkuStatusTier string

const (
	SBSkuStatusTier_Basic    = SBSkuStatusTier("Basic")
	SBSkuStatusTier_Premium  = SBSkuStatusTier("Premium")
	SBSkuStatusTier_Standard = SBSkuStatusTier("Standard")
)

// Deprecated version of SystemDataStatusCreatedByType. Use v1beta20210101preview.SystemDataStatusCreatedByType instead
type SystemDataStatusCreatedByType string

const (
	SystemDataStatusCreatedByType_Application     = SystemDataStatusCreatedByType("Application")
	SystemDataStatusCreatedByType_Key             = SystemDataStatusCreatedByType("Key")
	SystemDataStatusCreatedByType_ManagedIdentity = SystemDataStatusCreatedByType("ManagedIdentity")
	SystemDataStatusCreatedByType_User            = SystemDataStatusCreatedByType("User")
)

// Deprecated version of SystemDataStatusLastModifiedByType. Use v1beta20210101preview.SystemDataStatusLastModifiedByType
// instead
type SystemDataStatusLastModifiedByType string

const (
	SystemDataStatusLastModifiedByType_Application     = SystemDataStatusLastModifiedByType("Application")
	SystemDataStatusLastModifiedByType_Key             = SystemDataStatusLastModifiedByType("Key")
	SystemDataStatusLastModifiedByType_ManagedIdentity = SystemDataStatusLastModifiedByType("ManagedIdentity")
	SystemDataStatusLastModifiedByType_User            = SystemDataStatusLastModifiedByType("User")
)

// Deprecated version of KeyVaultProperties_Status. Use v1beta20210101preview.KeyVaultProperties_Status instead
type KeyVaultProperties_StatusARM struct {
	Identity    *UserAssignedIdentityProperties_StatusARM `json:"identity,omitempty"`
	KeyName     *string                                   `json:"keyName,omitempty"`
	KeyVaultUri *string                                   `json:"keyVaultUri,omitempty"`
	KeyVersion  *string                                   `json:"keyVersion,omitempty"`
}

// Deprecated version of UserAssignedIdentityProperties_Status. Use v1beta20210101preview.UserAssignedIdentityProperties_Status instead
type UserAssignedIdentityProperties_StatusARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
