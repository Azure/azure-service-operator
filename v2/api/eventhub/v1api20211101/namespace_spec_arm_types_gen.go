// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Namespace_Spec_ARM struct {
	// Identity: Properties of BYOK Identity description
	Identity *Identity_ARM `json:"identity,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Namespace properties supplied for create namespace operation.
	Properties *Namespace_Properties_Spec_ARM `json:"properties,omitempty"`

	// Sku: Properties of sku resource
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespace_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (namespace Namespace_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (namespace *Namespace_Spec_ARM) GetName() string {
	return namespace.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces"
func (namespace *Namespace_Spec_ARM) GetType() string {
	return "Microsoft.EventHub/namespaces"
}

// Properties to configure Identity for Bring your Own Keys
type Identity_ARM struct {
	// Type: Type of managed service identity.
	Type                   *Identity_Type_ARM                         `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

type Namespace_Properties_Spec_ARM struct {
	// AlternateName: Alternate name specified when alias and namespace names are same.
	AlternateName *string `json:"alternateName,omitempty"`
	ClusterArmId  *string `json:"clusterArmId,omitempty"`

	// DisableLocalAuth: This property disables SAS authentication for the Event Hubs namespace.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// Encryption: Properties of BYOK Encryption description
	Encryption *Encryption_ARM `json:"encryption,omitempty"`

	// IsAutoInflateEnabled: Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled *bool `json:"isAutoInflateEnabled,omitempty"`

	// KafkaEnabled: Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled *bool `json:"kafkaEnabled,omitempty"`

	// MaximumThroughputUnits: Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20
	// throughput units. ( '0' if AutoInflateEnabled = true)
	MaximumThroughputUnits *int `json:"maximumThroughputUnits,omitempty"`

	// ZoneRedundant: Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

// SKU parameters supplied to the create namespace operation
type Sku_ARM struct {
	// Capacity: The Event Hubs throughput units for Basic or Standard tiers, where value should be 0 to 20 throughput units.
	// The Event Hubs premium units for Premium tier, where value should be 0 to 10 premium units.
	Capacity *int `json:"capacity,omitempty"`

	// Name: Name of this SKU.
	Name *Sku_Name_ARM `json:"name,omitempty"`

	// Tier: The billing tier of this particular SKU.
	Tier *Sku_Tier_ARM `json:"tier,omitempty"`
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
type Sku_Name_ARM string

const (
	Sku_Name_ARM_Basic    = Sku_Name_ARM("Basic")
	Sku_Name_ARM_Premium  = Sku_Name_ARM("Premium")
	Sku_Name_ARM_Standard = Sku_Name_ARM("Standard")
)

// Mapping from string to Sku_Name_ARM
var sku_Name_ARM_Values = map[string]Sku_Name_ARM{
	"basic":    Sku_Name_ARM_Basic,
	"premium":  Sku_Name_ARM_Premium,
	"standard": Sku_Name_ARM_Standard,
}

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type Sku_Tier_ARM string

const (
	Sku_Tier_ARM_Basic    = Sku_Tier_ARM("Basic")
	Sku_Tier_ARM_Premium  = Sku_Tier_ARM("Premium")
	Sku_Tier_ARM_Standard = Sku_Tier_ARM("Standard")
)

// Mapping from string to Sku_Tier_ARM
var sku_Tier_ARM_Values = map[string]Sku_Tier_ARM{
	"basic":    Sku_Tier_ARM_Basic,
	"premium":  Sku_Tier_ARM_Premium,
	"standard": Sku_Tier_ARM_Standard,
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

	// KeyVersion: Key Version
	KeyVersion *string `json:"keyVersion,omitempty"`
}

type UserAssignedIdentityProperties_ARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
