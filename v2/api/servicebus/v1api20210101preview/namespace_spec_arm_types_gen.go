// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Namespace_Spec_ARM struct {
	// Identity: Properties of BYOK Identity description
	Identity *Identity_ARM `json:"identity,omitempty"`

	// Location: The Geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the namespace.
	Properties *SBNamespaceProperties_ARM `json:"properties,omitempty"`

	// Sku: Properties of SKU
	Sku *SBSku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespace_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (namespace Namespace_Spec_ARM) GetAPIVersion() string {
	return "2021-01-01-preview"
}

// GetName returns the Name of the resource
func (namespace *Namespace_Spec_ARM) GetName() string {
	return namespace.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces"
func (namespace *Namespace_Spec_ARM) GetType() string {
	return "Microsoft.ServiceBus/namespaces"
}

// Properties to configure User Assigned Identities for Bring your Own Keys
type Identity_ARM struct {
	// Type: Type of managed service identity.
	Type                   *Identity_Type_ARM                         `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Properties of the namespace.
type SBNamespaceProperties_ARM struct {
	// Encryption: Properties of BYOK Encryption description
	Encryption *Encryption_ARM `json:"encryption,omitempty"`

	// ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

// SKU of the namespace.
type SBSku_ARM struct {
	// Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity *int `json:"capacity,omitempty"`

	// Name: Name of this SKU.
	Name *SBSku_Name_ARM `json:"name,omitempty"`

	// Tier: The billing tier of this particular SKU.
	Tier *SBSku_Tier_ARM `json:"tier,omitempty"`
}

// Properties to configure Encryption
type Encryption_ARM struct {
	// KeySource: Enumerates the possible value of keySource for Encryption
	KeySource *Encryption_KeySource_ARM `json:"keySource,omitempty"`

	// KeyVaultProperties: Properties of KeyVault
	KeyVaultProperties []KeyVaultProperties_ARM `json:"keyVaultProperties,omitempty"`

	// RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type Identity_Type_ARM string

const (
	Identity_Type_ARM_None                       = Identity_Type_ARM("None")
	Identity_Type_ARM_SystemAssigned             = Identity_Type_ARM("SystemAssigned")
	Identity_Type_ARM_SystemAssignedUserAssigned = Identity_Type_ARM("SystemAssigned, UserAssigned")
	Identity_Type_ARM_UserAssigned               = Identity_Type_ARM("UserAssigned")
)

// Mapping from string to Identity_Type_ARM
var identity_Type_ARM_Values = map[string]Identity_Type_ARM{
	"none":                         Identity_Type_ARM_None,
	"systemassigned":               Identity_Type_ARM_SystemAssigned,
	"systemassigned, userassigned": Identity_Type_ARM_SystemAssignedUserAssigned,
	"userassigned":                 Identity_Type_ARM_UserAssigned,
}

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SBSku_Name_ARM string

const (
	SBSku_Name_ARM_Basic    = SBSku_Name_ARM("Basic")
	SBSku_Name_ARM_Premium  = SBSku_Name_ARM("Premium")
	SBSku_Name_ARM_Standard = SBSku_Name_ARM("Standard")
)

// Mapping from string to SBSku_Name_ARM
var sBSku_Name_ARM_Values = map[string]SBSku_Name_ARM{
	"basic":    SBSku_Name_ARM_Basic,
	"premium":  SBSku_Name_ARM_Premium,
	"standard": SBSku_Name_ARM_Standard,
}

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SBSku_Tier_ARM string

const (
	SBSku_Tier_ARM_Basic    = SBSku_Tier_ARM("Basic")
	SBSku_Tier_ARM_Premium  = SBSku_Tier_ARM("Premium")
	SBSku_Tier_ARM_Standard = SBSku_Tier_ARM("Standard")
)

// Mapping from string to SBSku_Tier_ARM
var sBSku_Tier_ARM_Values = map[string]SBSku_Tier_ARM{
	"basic":    SBSku_Tier_ARM_Basic,
	"premium":  SBSku_Tier_ARM_Premium,
	"standard": SBSku_Tier_ARM_Standard,
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// +kubebuilder:validation:Enum={"Microsoft.KeyVault"}
type Encryption_KeySource_ARM string

const Encryption_KeySource_ARM_MicrosoftKeyVault = Encryption_KeySource_ARM("Microsoft.KeyVault")

// Mapping from string to Encryption_KeySource_ARM
var encryption_KeySource_ARM_Values = map[string]Encryption_KeySource_ARM{
	"microsoft.keyvault": Encryption_KeySource_ARM_MicrosoftKeyVault,
}

// Properties to configure keyVault Properties
type KeyVaultProperties_ARM struct {
	Identity *UserAssignedIdentityProperties_ARM `json:"identity,omitempty"`

	// KeyName: Name of the Key from KeyVault
	KeyName *string `json:"keyName,omitempty"`

	// KeyVaultUri: Uri of KeyVault
	KeyVaultUri *string `json:"keyVaultUri,omitempty"`

	// KeyVersion: Version of KeyVault
	KeyVersion *string `json:"keyVersion,omitempty"`
}

type UserAssignedIdentityProperties_ARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
