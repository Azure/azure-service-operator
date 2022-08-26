// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210901

// Deprecated version of Registry_STATUS. Use v1beta20210901.Registry_STATUS instead
type Registry_STATUSARM struct {
	Id         *string                       `json:"id,omitempty"`
	Identity   *IdentityProperties_STATUSARM `json:"identity,omitempty"`
	Location   *string                       `json:"location,omitempty"`
	Name       *string                       `json:"name,omitempty"`
	Properties *RegistryProperties_STATUSARM `json:"properties,omitempty"`
	Sku        *Sku_STATUSARM                `json:"sku,omitempty"`
	SystemData *SystemData_STATUSARM         `json:"systemData,omitempty"`
	Tags       map[string]string             `json:"tags,omitempty"`
	Type       *string                       `json:"type,omitempty"`
}

// Deprecated version of IdentityProperties_STATUS. Use v1beta20210901.IdentityProperties_STATUS instead
type IdentityProperties_STATUSARM struct {
	PrincipalId            *string                                     `json:"principalId,omitempty"`
	TenantId               *string                                     `json:"tenantId,omitempty"`
<<<<<<< HEAD
	Type                   *IdentityProperties_Type_STATUS             `json:"type,omitempty"`
=======
	Type                   *IdentityProperties_STATUS_Type             `json:"type,omitempty"`
>>>>>>> main
	UserAssignedIdentities map[string]UserIdentityProperties_STATUSARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of RegistryProperties_STATUS. Use v1beta20210901.RegistryProperties_STATUS instead
type RegistryProperties_STATUSARM struct {
<<<<<<< HEAD
	AdminUserEnabled           *bool                                               `json:"adminUserEnabled,omitempty"`
	CreationDate               *string                                             `json:"creationDate,omitempty"`
	DataEndpointEnabled        *bool                                               `json:"dataEndpointEnabled,omitempty"`
	DataEndpointHostNames      []string                                            `json:"dataEndpointHostNames,omitempty"`
	Encryption                 *EncryptionProperty_STATUSARM                       `json:"encryption,omitempty"`
	LoginServer                *string                                             `json:"loginServer,omitempty"`
	NetworkRuleBypassOptions   *RegistryProperties_NetworkRuleBypassOptions_STATUS `json:"networkRuleBypassOptions,omitempty"`
	NetworkRuleSet             *NetworkRuleSet_STATUSARM                           `json:"networkRuleSet,omitempty"`
	Policies                   *Policies_STATUSARM                                 `json:"policies,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUSARM               `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *RegistryProperties_ProvisioningState_STATUS        `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *RegistryProperties_PublicNetworkAccess_STATUS      `json:"publicNetworkAccess,omitempty"`
	Status                     *Status_STATUSARM                                   `json:"status,omitempty"`
	ZoneRedundancy             *RegistryProperties_ZoneRedundancy_STATUS           `json:"zoneRedundancy,omitempty"`
=======
	AdminUserEnabled           *bool                                                     `json:"adminUserEnabled,omitempty"`
	CreationDate               *string                                                   `json:"creationDate,omitempty"`
	DataEndpointEnabled        *bool                                                     `json:"dataEndpointEnabled,omitempty"`
	DataEndpointHostNames      []string                                                  `json:"dataEndpointHostNames,omitempty"`
	Encryption                 *EncryptionProperty_STATUSARM                             `json:"encryption,omitempty"`
	LoginServer                *string                                                   `json:"loginServer,omitempty"`
	NetworkRuleBypassOptions   *RegistryProperties_STATUS_NetworkRuleBypassOptions       `json:"networkRuleBypassOptions,omitempty"`
	NetworkRuleSet             *NetworkRuleSet_STATUSARM                                 `json:"networkRuleSet,omitempty"`
	Policies                   *Policies_STATUSARM                                       `json:"policies,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *RegistryProperties_STATUS_ProvisioningState              `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *RegistryProperties_STATUS_PublicNetworkAccess            `json:"publicNetworkAccess,omitempty"`
	Status                     *Status_STATUSARM                                         `json:"status,omitempty"`
	ZoneRedundancy             *RegistryProperties_STATUS_ZoneRedundancy                 `json:"zoneRedundancy,omitempty"`
>>>>>>> main
}

// Deprecated version of Sku_STATUS. Use v1beta20210901.Sku_STATUS instead
type Sku_STATUSARM struct {
<<<<<<< HEAD
	Name *Sku_Name_STATUS `json:"name,omitempty"`
	Tier *Sku_Tier_STATUS `json:"tier,omitempty"`
=======
	Name *Sku_STATUS_Name `json:"name,omitempty"`
	Tier *Sku_STATUS_Tier `json:"tier,omitempty"`
>>>>>>> main
}

// Deprecated version of SystemData_STATUS. Use v1beta20210901.SystemData_STATUS instead
type SystemData_STATUSARM struct {
	CreatedAt          *string                               `json:"createdAt,omitempty"`
	CreatedBy          *string                               `json:"createdBy,omitempty"`
<<<<<<< HEAD
	CreatedByType      *SystemData_CreatedByType_STATUS      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                               `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                               `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
=======
	CreatedByType      *SystemData_STATUS_CreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                               `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                               `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemData_STATUS_LastModifiedByType `json:"lastModifiedByType,omitempty"`
>>>>>>> main
}

// Deprecated version of EncryptionProperty_STATUS. Use v1beta20210901.EncryptionProperty_STATUS instead
type EncryptionProperty_STATUSARM struct {
	KeyVaultProperties *KeyVaultProperties_STATUSARM     `json:"keyVaultProperties,omitempty"`
<<<<<<< HEAD
	Status             *EncryptionProperty_Status_STATUS `json:"status,omitempty"`
}

// Deprecated version of IdentityProperties_Type_STATUS. Use v1beta20210901.IdentityProperties_Type_STATUS instead
type IdentityProperties_Type_STATUS string

const (
	IdentityProperties_Type_None_STATUS                       = IdentityProperties_Type_STATUS("None")
	IdentityProperties_Type_SystemAssigned_STATUS             = IdentityProperties_Type_STATUS("SystemAssigned")
	IdentityProperties_Type_SystemAssignedUserAssigned_STATUS = IdentityProperties_Type_STATUS("SystemAssigned, UserAssigned")
	IdentityProperties_Type_UserAssigned_STATUS               = IdentityProperties_Type_STATUS("UserAssigned")
=======
	Status             *EncryptionProperty_STATUS_Status `json:"status,omitempty"`
}

// Deprecated version of IdentityProperties_STATUS_Type. Use v1beta20210901.IdentityProperties_STATUS_Type instead
type IdentityProperties_STATUS_Type string

const (
	IdentityProperties_STATUS_Type_None                       = IdentityProperties_STATUS_Type("None")
	IdentityProperties_STATUS_Type_SystemAssigned             = IdentityProperties_STATUS_Type("SystemAssigned")
	IdentityProperties_STATUS_Type_SystemAssignedUserAssigned = IdentityProperties_STATUS_Type("SystemAssigned, UserAssigned")
	IdentityProperties_STATUS_Type_UserAssigned               = IdentityProperties_STATUS_Type("UserAssigned")
>>>>>>> main
)

// Deprecated version of NetworkRuleSet_STATUS. Use v1beta20210901.NetworkRuleSet_STATUS instead
type NetworkRuleSet_STATUSARM struct {
<<<<<<< HEAD
	DefaultAction *NetworkRuleSet_DefaultAction_STATUS `json:"defaultAction,omitempty"`
=======
	DefaultAction *NetworkRuleSet_STATUS_DefaultAction `json:"defaultAction,omitempty"`
>>>>>>> main
	IpRules       []IPRule_STATUSARM                   `json:"ipRules,omitempty"`
}

// Deprecated version of Policies_STATUS. Use v1beta20210901.Policies_STATUS instead
type Policies_STATUSARM struct {
	ExportPolicy     *ExportPolicy_STATUSARM     `json:"exportPolicy,omitempty"`
	QuarantinePolicy *QuarantinePolicy_STATUSARM `json:"quarantinePolicy,omitempty"`
	RetentionPolicy  *RetentionPolicy_STATUSARM  `json:"retentionPolicy,omitempty"`
	TrustPolicy      *TrustPolicy_STATUSARM      `json:"trustPolicy,omitempty"`
}

// Deprecated version of PrivateEndpointConnection_STATUS. Use v1beta20210901.PrivateEndpointConnection_STATUS instead
type PrivateEndpointConnection_STATUSARM struct {
	Id         *string               `json:"id,omitempty"`
	SystemData *SystemData_STATUSARM `json:"systemData,omitempty"`
}

<<<<<<< HEAD
// Deprecated version of Sku_Name_STATUS. Use v1beta20210901.Sku_Name_STATUS instead
type Sku_Name_STATUS string

const (
	Sku_Name_Basic_STATUS    = Sku_Name_STATUS("Basic")
	Sku_Name_Classic_STATUS  = Sku_Name_STATUS("Classic")
	Sku_Name_Premium_STATUS  = Sku_Name_STATUS("Premium")
	Sku_Name_Standard_STATUS = Sku_Name_STATUS("Standard")
)

// Deprecated version of Sku_Tier_STATUS. Use v1beta20210901.Sku_Tier_STATUS instead
type Sku_Tier_STATUS string

const (
	Sku_Tier_Basic_STATUS    = Sku_Tier_STATUS("Basic")
	Sku_Tier_Classic_STATUS  = Sku_Tier_STATUS("Classic")
	Sku_Tier_Premium_STATUS  = Sku_Tier_STATUS("Premium")
	Sku_Tier_Standard_STATUS = Sku_Tier_STATUS("Standard")
=======
// Deprecated version of Sku_STATUS_Name. Use v1beta20210901.Sku_STATUS_Name instead
type Sku_STATUS_Name string

const (
	Sku_STATUS_Name_Basic    = Sku_STATUS_Name("Basic")
	Sku_STATUS_Name_Classic  = Sku_STATUS_Name("Classic")
	Sku_STATUS_Name_Premium  = Sku_STATUS_Name("Premium")
	Sku_STATUS_Name_Standard = Sku_STATUS_Name("Standard")
)

// Deprecated version of Sku_STATUS_Tier. Use v1beta20210901.Sku_STATUS_Tier instead
type Sku_STATUS_Tier string

const (
	Sku_STATUS_Tier_Basic    = Sku_STATUS_Tier("Basic")
	Sku_STATUS_Tier_Classic  = Sku_STATUS_Tier("Classic")
	Sku_STATUS_Tier_Premium  = Sku_STATUS_Tier("Premium")
	Sku_STATUS_Tier_Standard = Sku_STATUS_Tier("Standard")
>>>>>>> main
)

// Deprecated version of Status_STATUS. Use v1beta20210901.Status_STATUS instead
type Status_STATUSARM struct {
	DisplayStatus *string `json:"displayStatus,omitempty"`
	Message       *string `json:"message,omitempty"`
	Timestamp     *string `json:"timestamp,omitempty"`
}

<<<<<<< HEAD
// Deprecated version of SystemData_CreatedByType_STATUS. Use v1beta20210901.SystemData_CreatedByType_STATUS instead
type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_Application_STATUS     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_Key_STATUS             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_ManagedIdentity_STATUS = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_User_STATUS            = SystemData_CreatedByType_STATUS("User")
)

// Deprecated version of SystemData_LastModifiedByType_STATUS. Use v1beta20210901.SystemData_LastModifiedByType_STATUS
// instead
type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_Application_STATUS     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_Key_STATUS             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_ManagedIdentity_STATUS = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_User_STATUS            = SystemData_LastModifiedByType_STATUS("User")
=======
// Deprecated version of SystemData_STATUS_CreatedByType. Use v1beta20210901.SystemData_STATUS_CreatedByType instead
type SystemData_STATUS_CreatedByType string

const (
	SystemData_STATUS_CreatedByType_Application     = SystemData_STATUS_CreatedByType("Application")
	SystemData_STATUS_CreatedByType_Key             = SystemData_STATUS_CreatedByType("Key")
	SystemData_STATUS_CreatedByType_ManagedIdentity = SystemData_STATUS_CreatedByType("ManagedIdentity")
	SystemData_STATUS_CreatedByType_User            = SystemData_STATUS_CreatedByType("User")
)

// Deprecated version of SystemData_STATUS_LastModifiedByType. Use v1beta20210901.SystemData_STATUS_LastModifiedByType
// instead
type SystemData_STATUS_LastModifiedByType string

const (
	SystemData_STATUS_LastModifiedByType_Application     = SystemData_STATUS_LastModifiedByType("Application")
	SystemData_STATUS_LastModifiedByType_Key             = SystemData_STATUS_LastModifiedByType("Key")
	SystemData_STATUS_LastModifiedByType_ManagedIdentity = SystemData_STATUS_LastModifiedByType("ManagedIdentity")
	SystemData_STATUS_LastModifiedByType_User            = SystemData_STATUS_LastModifiedByType("User")
>>>>>>> main
)

// Deprecated version of UserIdentityProperties_STATUS. Use v1beta20210901.UserIdentityProperties_STATUS instead
type UserIdentityProperties_STATUSARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of ExportPolicy_STATUS. Use v1beta20210901.ExportPolicy_STATUS instead
type ExportPolicy_STATUSARM struct {
<<<<<<< HEAD
	Status *ExportPolicy_Status_STATUS `json:"status,omitempty"`
=======
	Status *ExportPolicy_STATUS_Status `json:"status,omitempty"`
>>>>>>> main
}

// Deprecated version of IPRule_STATUS. Use v1beta20210901.IPRule_STATUS instead
type IPRule_STATUSARM struct {
<<<<<<< HEAD
	Action *IPRule_Action_STATUS `json:"action,omitempty"`
=======
	Action *IPRule_STATUS_Action `json:"action,omitempty"`
>>>>>>> main
	Value  *string               `json:"value,omitempty"`
}

// Deprecated version of KeyVaultProperties_STATUS. Use v1beta20210901.KeyVaultProperties_STATUS instead
type KeyVaultProperties_STATUSARM struct {
	Identity                 *string `json:"identity,omitempty"`
	KeyIdentifier            *string `json:"keyIdentifier,omitempty"`
	KeyRotationEnabled       *bool   `json:"keyRotationEnabled,omitempty"`
	LastKeyRotationTimestamp *string `json:"lastKeyRotationTimestamp,omitempty"`
	VersionedKeyIdentifier   *string `json:"versionedKeyIdentifier,omitempty"`
}

// Deprecated version of QuarantinePolicy_STATUS. Use v1beta20210901.QuarantinePolicy_STATUS instead
type QuarantinePolicy_STATUSARM struct {
<<<<<<< HEAD
	Status *QuarantinePolicy_Status_STATUS `json:"status,omitempty"`
=======
	Status *QuarantinePolicy_STATUS_Status `json:"status,omitempty"`
>>>>>>> main
}

// Deprecated version of RetentionPolicy_STATUS. Use v1beta20210901.RetentionPolicy_STATUS instead
type RetentionPolicy_STATUSARM struct {
	Days            *int                           `json:"days,omitempty"`
	LastUpdatedTime *string                        `json:"lastUpdatedTime,omitempty"`
<<<<<<< HEAD
	Status          *RetentionPolicy_Status_STATUS `json:"status,omitempty"`
=======
	Status          *RetentionPolicy_STATUS_Status `json:"status,omitempty"`
>>>>>>> main
}

// Deprecated version of TrustPolicy_STATUS. Use v1beta20210901.TrustPolicy_STATUS instead
type TrustPolicy_STATUSARM struct {
<<<<<<< HEAD
	Status *TrustPolicy_Status_STATUS `json:"status,omitempty"`
	Type   *TrustPolicy_Type_STATUS   `json:"type,omitempty"`
=======
	Status *TrustPolicy_STATUS_Status `json:"status,omitempty"`
	Type   *TrustPolicy_STATUS_Type   `json:"type,omitempty"`
>>>>>>> main
}
