// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Deprecated version of DatabaseAccount_Spec. Use v1api20210515.DatabaseAccount_Spec instead
type DatabaseAccount_Spec_ARM struct {
	Identity   *ManagedServiceIdentity_ARM                `json:"identity,omitempty"`
	Kind       *DatabaseAccount_Kind_Spec                 `json:"kind,omitempty"`
	Location   *string                                    `json:"location,omitempty"`
	Name       string                                     `json:"name,omitempty"`
	Properties *DatabaseAccountCreateUpdateProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                          `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccount_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (account DatabaseAccount_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (account *DatabaseAccount_Spec_ARM) GetName() string {
	return account.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts"
func (account *DatabaseAccount_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts"
}

// Deprecated version of DatabaseAccount_Kind_Spec. Use v1api20210515.DatabaseAccount_Kind_Spec instead
// +kubebuilder:validation:Enum={"GlobalDocumentDB","MongoDB","Parse"}
type DatabaseAccount_Kind_Spec string

const (
	DatabaseAccount_Kind_Spec_GlobalDocumentDB = DatabaseAccount_Kind_Spec("GlobalDocumentDB")
	DatabaseAccount_Kind_Spec_MongoDB          = DatabaseAccount_Kind_Spec("MongoDB")
	DatabaseAccount_Kind_Spec_Parse            = DatabaseAccount_Kind_Spec("Parse")
)

// Deprecated version of DatabaseAccountCreateUpdateProperties. Use v1api20210515.DatabaseAccountCreateUpdateProperties instead
type DatabaseAccountCreateUpdateProperties_ARM struct {
	AnalyticalStorageConfiguration     *AnalyticalStorageConfiguration_ARM `json:"analyticalStorageConfiguration,omitempty"`
	ApiProperties                      *ApiProperties_ARM                  `json:"apiProperties,omitempty"`
	BackupPolicy                       *BackupPolicy_ARM                   `json:"backupPolicy,omitempty"`
	Capabilities                       []Capability_ARM                    `json:"capabilities,omitempty"`
	ConnectorOffer                     *ConnectorOffer                     `json:"connectorOffer,omitempty"`
	ConsistencyPolicy                  *ConsistencyPolicy_ARM              `json:"consistencyPolicy,omitempty"`
	Cors                               []CorsPolicy_ARM                    `json:"cors,omitempty"`
	DatabaseAccountOfferType           *DatabaseAccountOfferType           `json:"databaseAccountOfferType,omitempty"`
	DefaultIdentity                    *string                             `json:"defaultIdentity,omitempty"`
	DisableKeyBasedMetadataWriteAccess *bool                               `json:"disableKeyBasedMetadataWriteAccess,omitempty"`
	EnableAnalyticalStorage            *bool                               `json:"enableAnalyticalStorage,omitempty"`
	EnableAutomaticFailover            *bool                               `json:"enableAutomaticFailover,omitempty"`
	EnableCassandraConnector           *bool                               `json:"enableCassandraConnector,omitempty"`
	EnableFreeTier                     *bool                               `json:"enableFreeTier,omitempty"`
	EnableMultipleWriteLocations       *bool                               `json:"enableMultipleWriteLocations,omitempty"`
	IpRules                            []IpAddressOrRange_ARM              `json:"ipRules,omitempty"`
	IsVirtualNetworkFilterEnabled      *bool                               `json:"isVirtualNetworkFilterEnabled,omitempty"`
	KeyVaultKeyUri                     *string                             `json:"keyVaultKeyUri,omitempty"`
	Locations                          []Location_ARM                      `json:"locations,omitempty"`
	NetworkAclBypass                   *NetworkAclBypass                   `json:"networkAclBypass,omitempty"`
	NetworkAclBypassResourceIds        []string                            `json:"networkAclBypassResourceIds,omitempty"`
	PublicNetworkAccess                *PublicNetworkAccess                `json:"publicNetworkAccess,omitempty"`
	VirtualNetworkRules                []VirtualNetworkRule_ARM            `json:"virtualNetworkRules,omitempty"`
}

// Deprecated version of ManagedServiceIdentity. Use v1api20210515.ManagedServiceIdentity instead
type ManagedServiceIdentity_ARM struct {
	Type                   *ManagedServiceIdentity_Type               `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of AnalyticalStorageConfiguration. Use v1api20210515.AnalyticalStorageConfiguration instead
type AnalyticalStorageConfiguration_ARM struct {
	SchemaType *AnalyticalStorageSchemaType `json:"schemaType,omitempty"`
}

// Deprecated version of ApiProperties. Use v1api20210515.ApiProperties instead
type ApiProperties_ARM struct {
	ServerVersion *ApiProperties_ServerVersion `json:"serverVersion,omitempty"`
}

// Deprecated version of BackupPolicy. Use v1api20210515.BackupPolicy instead
type BackupPolicy_ARM struct {
	Continuous *ContinuousModeBackupPolicy_ARM `json:"continuous,omitempty"`
	Periodic   *PeriodicModeBackupPolicy_ARM   `json:"periodic,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupPolicy_ARM represents a discriminated union (JSON OneOf)
func (policy BackupPolicy_ARM) MarshalJSON() ([]byte, error) {
	if policy.Continuous != nil {
		return json.Marshal(policy.Continuous)
	}
	if policy.Periodic != nil {
		return json.Marshal(policy.Periodic)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupPolicy_ARM
func (policy *BackupPolicy_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "Continuous" {
		policy.Continuous = &ContinuousModeBackupPolicy_ARM{}
		return json.Unmarshal(data, policy.Continuous)
	}
	if discriminator == "Periodic" {
		policy.Periodic = &PeriodicModeBackupPolicy_ARM{}
		return json.Unmarshal(data, policy.Periodic)
	}

	// No error
	return nil
}

// Deprecated version of Capability. Use v1api20210515.Capability instead
type Capability_ARM struct {
	Name *string `json:"name,omitempty"`
}

// Deprecated version of ConsistencyPolicy. Use v1api20210515.ConsistencyPolicy instead
type ConsistencyPolicy_ARM struct {
	DefaultConsistencyLevel *ConsistencyPolicy_DefaultConsistencyLevel `json:"defaultConsistencyLevel,omitempty"`
	MaxIntervalInSeconds    *int                                       `json:"maxIntervalInSeconds,omitempty"`
	MaxStalenessPrefix      *int                                       `json:"maxStalenessPrefix,omitempty"`
}

// Deprecated version of CorsPolicy. Use v1api20210515.CorsPolicy instead
type CorsPolicy_ARM struct {
	AllowedHeaders  *string `json:"allowedHeaders,omitempty"`
	AllowedMethods  *string `json:"allowedMethods,omitempty"`
	AllowedOrigins  *string `json:"allowedOrigins,omitempty"`
	ExposedHeaders  *string `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int    `json:"maxAgeInSeconds,omitempty"`
}

// Deprecated version of IpAddressOrRange. Use v1api20210515.IpAddressOrRange instead
type IpAddressOrRange_ARM struct {
	IpAddressOrRange *string `json:"ipAddressOrRange,omitempty"`
}

// Deprecated version of Location. Use v1api20210515.Location instead
type Location_ARM struct {
	FailoverPriority *int    `json:"failoverPriority,omitempty"`
	IsZoneRedundant  *bool   `json:"isZoneRedundant,omitempty"`
	LocationName     *string `json:"locationName,omitempty"`
}

// Deprecated version of ManagedServiceIdentity_Type. Use v1api20210515.ManagedServiceIdentity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ManagedServiceIdentity_Type string

const (
	ManagedServiceIdentity_Type_None                       = ManagedServiceIdentity_Type("None")
	ManagedServiceIdentity_Type_SystemAssigned             = ManagedServiceIdentity_Type("SystemAssigned")
	ManagedServiceIdentity_Type_SystemAssignedUserAssigned = ManagedServiceIdentity_Type("SystemAssigned,UserAssigned")
	ManagedServiceIdentity_Type_UserAssigned               = ManagedServiceIdentity_Type("UserAssigned")
)

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// Deprecated version of VirtualNetworkRule. Use v1api20210515.VirtualNetworkRule instead
type VirtualNetworkRule_ARM struct {
	Id                               *string `json:"id,omitempty"`
	IgnoreMissingVNetServiceEndpoint *bool   `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
}

// Deprecated version of ContinuousModeBackupPolicy. Use v1api20210515.ContinuousModeBackupPolicy instead
type ContinuousModeBackupPolicy_ARM struct {
	Type ContinuousModeBackupPolicy_Type `json:"type,omitempty"`
}

// Deprecated version of PeriodicModeBackupPolicy. Use v1api20210515.PeriodicModeBackupPolicy instead
type PeriodicModeBackupPolicy_ARM struct {
	PeriodicModeProperties *PeriodicModeProperties_ARM   `json:"periodicModeProperties,omitempty"`
	Type                   PeriodicModeBackupPolicy_Type `json:"type,omitempty"`
}

// Deprecated version of PeriodicModeProperties. Use v1api20210515.PeriodicModeProperties instead
type PeriodicModeProperties_ARM struct {
	BackupIntervalInMinutes        *int `json:"backupIntervalInMinutes,omitempty"`
	BackupRetentionIntervalInHours *int `json:"backupRetentionIntervalInHours,omitempty"`
}