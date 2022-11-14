// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

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
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (namespace *Namespace_Spec_ARM) GetName() string {
	return namespace.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces"
func (namespace *Namespace_Spec_ARM) GetType() string {
	return "Microsoft.ServiceBus/namespaces"
}

type Identity_ARM struct {
	// Type: Type of managed service identity.
	Type *Identity_Type `json:"type,omitempty"`
}

type SBNamespaceProperties_ARM struct {
	// Encryption: Properties of BYOK Encryption description
	Encryption *Encryption_ARM `json:"encryption,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnection_ARM `json:"privateEndpointConnections,omitempty"`

	// ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

type SBSku_ARM struct {
	// Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity *int `json:"capacity,omitempty"`

	// Name: Name of this SKU.
	Name *SBSku_Name `json:"name,omitempty"`

	// Tier: The billing tier of this particular SKU.
	Tier *SBSku_Tier `json:"tier,omitempty"`
}

type Encryption_ARM struct {
	// KeySource: Enumerates the possible value of keySource for Encryption
	KeySource *Encryption_KeySource `json:"keySource,omitempty"`

	// KeyVaultProperties: Properties of KeyVault
	KeyVaultProperties []KeyVaultProperties_ARM `json:"keyVaultProperties,omitempty"`

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

type PrivateEndpointConnection_ARM struct {
	// Properties: Properties of the PrivateEndpointConnection.
	Properties *PrivateEndpointConnectionProperties_ARM `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SBSku_Name string

const (
	SBSku_Name_Basic    = SBSku_Name("Basic")
	SBSku_Name_Premium  = SBSku_Name("Premium")
	SBSku_Name_Standard = SBSku_Name("Standard")
)

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SBSku_Tier string

const (
	SBSku_Tier_Basic    = SBSku_Tier("Basic")
	SBSku_Tier_Premium  = SBSku_Tier("Premium")
	SBSku_Tier_Standard = SBSku_Tier("Standard")
)

type KeyVaultProperties_ARM struct {
	Identity *UserAssignedIdentityProperties_ARM `json:"identity,omitempty"`

	// KeyName: Name of the Key from KeyVault
	KeyName *string `json:"keyName,omitempty"`

	// KeyVaultUri: Uri of KeyVault
	KeyVaultUri *string `json:"keyVaultUri,omitempty"`

	// KeyVersion: Version of KeyVault
	KeyVersion *string `json:"keyVersion,omitempty"`
}

type PrivateEndpointConnectionProperties_ARM struct {
	// PrivateEndpoint: The Private Endpoint resource for this Connection.
	PrivateEndpoint *PrivateEndpoint_ARM `json:"privateEndpoint,omitempty"`

	// PrivateLinkServiceConnectionState: Details about the state of the connection.
	PrivateLinkServiceConnectionState *ConnectionState_ARM `json:"privateLinkServiceConnectionState,omitempty"`

	// ProvisioningState: Provisioning state of the Private Endpoint Connection.
	ProvisioningState *PrivateEndpointConnectionProperties_ProvisioningState `json:"provisioningState,omitempty"`
}

type ConnectionState_ARM struct {
	// Description: Description of the connection state.
	Description *string `json:"description,omitempty"`

	// Status: Status of the connection.
	Status *ConnectionState_Status `json:"status,omitempty"`
}

type PrivateEndpoint_ARM struct {
	// Id: The ARM identifier for Private Endpoint.
	Id *string `json:"id,omitempty"`
}

type UserAssignedIdentityProperties_ARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
