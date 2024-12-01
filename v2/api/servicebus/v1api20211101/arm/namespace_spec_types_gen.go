// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Namespace_Spec struct {
	// Identity: Properties of BYOK Identity description
	Identity *Identity `json:"identity,omitempty"`

	// Location: The Geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the namespace.
	Properties *SBNamespaceProperties `json:"properties,omitempty"`

	// Sku: Properties of SKU
	Sku *SBSku `json:"sku,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespace_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (namespace Namespace_Spec) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (namespace *Namespace_Spec) GetName() string {
	return namespace.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces"
func (namespace *Namespace_Spec) GetType() string {
	return "Microsoft.ServiceBus/namespaces"
}

// Properties to configure User Assigned Identities for Bring your Own Keys
type Identity struct {
	// Type: Type of managed service identity.
	Type                   *Identity_Type                         `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Properties of the namespace.
type SBNamespaceProperties struct {
	// AlternateName: Alternate name for namespace
	AlternateName *string `json:"alternateName,omitempty"`

	// DisableLocalAuth: This property disables SAS authentication for the Service Bus namespace.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// Encryption: Properties of BYOK Encryption description
	Encryption *Encryption `json:"encryption,omitempty"`

	// ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

// SKU of the namespace.
type SBSku struct {
	// Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity *int `json:"capacity,omitempty"`

	// Name: Name of this SKU.
	Name *SBSku_Name `json:"name,omitempty"`

	// Tier: The billing tier of this particular SKU.
	Tier *SBSku_Tier `json:"tier,omitempty"`
}

// Properties to configure Encryption
type Encryption struct {
	// KeySource: Enumerates the possible value of keySource for Encryption
	KeySource *Encryption_KeySource `json:"keySource,omitempty"`

	// KeyVaultProperties: Properties of KeyVault
	KeyVaultProperties []KeyVaultProperties `json:"keyVaultProperties,omitempty"`

	// RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type Identity_Type string

const (
	Identity_Type_None                       = Identity_Type("None")
	Identity_Type_SystemAssigned             = Identity_Type("SystemAssigned")
	Identity_Type_SystemAssignedUserAssigned = Identity_Type("SystemAssigned, UserAssigned")
	Identity_Type_UserAssigned               = Identity_Type("UserAssigned")
)

// Mapping from string to Identity_Type
var identity_Type_Values = map[string]Identity_Type{
	"none":                         Identity_Type_None,
	"systemassigned":               Identity_Type_SystemAssigned,
	"systemassigned, userassigned": Identity_Type_SystemAssignedUserAssigned,
	"userassigned":                 Identity_Type_UserAssigned,
}

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SBSku_Name string

const (
	SBSku_Name_Basic    = SBSku_Name("Basic")
	SBSku_Name_Premium  = SBSku_Name("Premium")
	SBSku_Name_Standard = SBSku_Name("Standard")
)

// Mapping from string to SBSku_Name
var sBSku_Name_Values = map[string]SBSku_Name{
	"basic":    SBSku_Name_Basic,
	"premium":  SBSku_Name_Premium,
	"standard": SBSku_Name_Standard,
}

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SBSku_Tier string

const (
	SBSku_Tier_Basic    = SBSku_Tier("Basic")
	SBSku_Tier_Premium  = SBSku_Tier("Premium")
	SBSku_Tier_Standard = SBSku_Tier("Standard")
)

// Mapping from string to SBSku_Tier
var sBSku_Tier_Values = map[string]SBSku_Tier{
	"basic":    SBSku_Tier_Basic,
	"premium":  SBSku_Tier_Premium,
	"standard": SBSku_Tier_Standard,
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
}

// +kubebuilder:validation:Enum={"Microsoft.KeyVault"}
type Encryption_KeySource string

const Encryption_KeySource_MicrosoftKeyVault = Encryption_KeySource("Microsoft.KeyVault")

// Mapping from string to Encryption_KeySource
var encryption_KeySource_Values = map[string]Encryption_KeySource{
	"microsoft.keyvault": Encryption_KeySource_MicrosoftKeyVault,
}

// Properties to configure keyVault Properties
type KeyVaultProperties struct {
	Identity *UserAssignedIdentityProperties `json:"identity,omitempty"`

	// KeyName: Name of the Key from KeyVault
	KeyName *string `json:"keyName,omitempty"`

	// KeyVaultUri: Uri of KeyVault
	KeyVaultUri *string `json:"keyVaultUri,omitempty"`

	// KeyVersion: Version of KeyVault
	KeyVersion *string `json:"keyVersion,omitempty"`
}

type UserAssignedIdentityProperties struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
