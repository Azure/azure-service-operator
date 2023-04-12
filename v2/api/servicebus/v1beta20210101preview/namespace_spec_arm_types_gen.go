// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Namespace_Spec. Use v1api20210101preview.Namespace_Spec instead
type Namespace_Spec_ARM struct {
	Identity   *Identity_ARM              `json:"identity,omitempty"`
	Location   *string                    `json:"location,omitempty"`
	Name       string                     `json:"name,omitempty"`
	Properties *SBNamespaceProperties_ARM `json:"properties,omitempty"`
	Sku        *SBSku_ARM                 `json:"sku,omitempty"`
	Tags       map[string]string          `json:"tags,omitempty"`
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

// Deprecated version of Identity. Use v1api20210101preview.Identity instead
type Identity_ARM struct {
	Type                   *Identity_Type                             `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of SBNamespaceProperties. Use v1api20210101preview.SBNamespaceProperties instead
type SBNamespaceProperties_ARM struct {
	Encryption    *Encryption_ARM `json:"encryption,omitempty"`
	ZoneRedundant *bool           `json:"zoneRedundant,omitempty"`
}

// Deprecated version of SBSku. Use v1api20210101preview.SBSku instead
type SBSku_ARM struct {
	Capacity *int        `json:"capacity,omitempty"`
	Name     *SBSku_Name `json:"name,omitempty"`
	Tier     *SBSku_Tier `json:"tier,omitempty"`
}

// Deprecated version of Encryption. Use v1api20210101preview.Encryption instead
type Encryption_ARM struct {
	KeySource                       *Encryption_KeySource    `json:"keySource,omitempty"`
	KeyVaultProperties              []KeyVaultProperties_ARM `json:"keyVaultProperties,omitempty"`
	RequireInfrastructureEncryption *bool                    `json:"requireInfrastructureEncryption,omitempty"`
}

// Deprecated version of Identity_Type. Use v1api20210101preview.Identity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type Identity_Type string

const (
	Identity_Type_None                       = Identity_Type("None")
	Identity_Type_SystemAssigned             = Identity_Type("SystemAssigned")
	Identity_Type_SystemAssignedUserAssigned = Identity_Type("SystemAssigned, UserAssigned")
	Identity_Type_UserAssigned               = Identity_Type("UserAssigned")
)

// Deprecated version of SBSku_Name. Use v1api20210101preview.SBSku_Name instead
// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SBSku_Name string

const (
	SBSku_Name_Basic    = SBSku_Name("Basic")
	SBSku_Name_Premium  = SBSku_Name("Premium")
	SBSku_Name_Standard = SBSku_Name("Standard")
)

// Deprecated version of SBSku_Tier. Use v1api20210101preview.SBSku_Tier instead
// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SBSku_Tier string

const (
	SBSku_Tier_Basic    = SBSku_Tier("Basic")
	SBSku_Tier_Premium  = SBSku_Tier("Premium")
	SBSku_Tier_Standard = SBSku_Tier("Standard")
)

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// Deprecated version of KeyVaultProperties. Use v1api20210101preview.KeyVaultProperties instead
type KeyVaultProperties_ARM struct {
	Identity    *UserAssignedIdentityProperties_ARM `json:"identity,omitempty"`
	KeyName     *string                             `json:"keyName,omitempty"`
	KeyVaultUri *string                             `json:"keyVaultUri,omitempty"`
	KeyVersion  *string                             `json:"keyVersion,omitempty"`
}

// Deprecated version of UserAssignedIdentityProperties. Use v1api20210101preview.UserAssignedIdentityProperties instead
type UserAssignedIdentityProperties_ARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
