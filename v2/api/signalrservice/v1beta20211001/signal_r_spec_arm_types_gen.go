// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of SignalR_Spec. Use v1api20211001.SignalR_Spec instead
type SignalR_Spec_ARM struct {
	Identity   *ManagedIdentity_ARM   `json:"identity,omitempty"`
	Kind       *ServiceKind           `json:"kind,omitempty"`
	Location   *string                `json:"location,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Properties *SignalRProperties_ARM `json:"properties,omitempty"`
	Sku        *ResourceSku_ARM       `json:"sku,omitempty"`
	Tags       map[string]string      `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &SignalR_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-10-01"
func (signalR SignalR_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (signalR *SignalR_Spec_ARM) GetName() string {
	return signalR.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.SignalRService/signalR"
func (signalR *SignalR_Spec_ARM) GetType() string {
	return "Microsoft.SignalRService/signalR"
}

// Deprecated version of ManagedIdentity. Use v1api20211001.ManagedIdentity instead
type ManagedIdentity_ARM struct {
	Type *ManagedIdentityType `json:"type,omitempty"`
}

// Deprecated version of ResourceSku. Use v1api20211001.ResourceSku instead
type ResourceSku_ARM struct {
	Capacity *int            `json:"capacity,omitempty"`
	Name     *string         `json:"name,omitempty"`
	Tier     *SignalRSkuTier `json:"tier,omitempty"`
}

// Deprecated version of ServiceKind. Use v1api20211001.ServiceKind instead
// +kubebuilder:validation:Enum={"RawWebSockets","SignalR"}
type ServiceKind string

const (
	ServiceKind_RawWebSockets = ServiceKind("RawWebSockets")
	ServiceKind_SignalR       = ServiceKind("SignalR")
)

// Deprecated version of SignalRProperties. Use v1api20211001.SignalRProperties instead
type SignalRProperties_ARM struct {
	Cors                     *SignalRCorsSettings_ARM        `json:"cors,omitempty"`
	DisableAadAuth           *bool                           `json:"disableAadAuth,omitempty"`
	DisableLocalAuth         *bool                           `json:"disableLocalAuth,omitempty"`
	Features                 []SignalRFeature_ARM            `json:"features,omitempty"`
	NetworkACLs              *SignalRNetworkACLs_ARM         `json:"networkACLs,omitempty"`
	PublicNetworkAccess      *string                         `json:"publicNetworkAccess,omitempty"`
	ResourceLogConfiguration *ResourceLogConfiguration_ARM   `json:"resourceLogConfiguration,omitempty"`
	Tls                      *SignalRTlsSettings_ARM         `json:"tls,omitempty"`
	Upstream                 *ServerlessUpstreamSettings_ARM `json:"upstream,omitempty"`
}

// Deprecated version of ManagedIdentityType. Use v1api20211001.ManagedIdentityType instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","UserAssigned"}
type ManagedIdentityType string

const (
	ManagedIdentityType_None           = ManagedIdentityType("None")
	ManagedIdentityType_SystemAssigned = ManagedIdentityType("SystemAssigned")
	ManagedIdentityType_UserAssigned   = ManagedIdentityType("UserAssigned")
)

// Deprecated version of ResourceLogConfiguration. Use v1api20211001.ResourceLogConfiguration instead
type ResourceLogConfiguration_ARM struct {
	Categories []ResourceLogCategory_ARM `json:"categories,omitempty"`
}

// Deprecated version of ServerlessUpstreamSettings. Use v1api20211001.ServerlessUpstreamSettings instead
type ServerlessUpstreamSettings_ARM struct {
	Templates []UpstreamTemplate_ARM `json:"templates,omitempty"`
}

// Deprecated version of SignalRCorsSettings. Use v1api20211001.SignalRCorsSettings instead
type SignalRCorsSettings_ARM struct {
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
}

// Deprecated version of SignalRFeature. Use v1api20211001.SignalRFeature instead
type SignalRFeature_ARM struct {
	Flag       *FeatureFlags     `json:"flag,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Value      *string           `json:"value,omitempty"`
}

// Deprecated version of SignalRNetworkACLs. Use v1api20211001.SignalRNetworkACLs instead
type SignalRNetworkACLs_ARM struct {
	DefaultAction    *ACLAction               `json:"defaultAction,omitempty"`
	PrivateEndpoints []PrivateEndpointACL_ARM `json:"privateEndpoints,omitempty"`
	PublicNetwork    *NetworkACL_ARM          `json:"publicNetwork,omitempty"`
}

// Deprecated version of SignalRSkuTier. Use v1api20211001.SignalRSkuTier instead
// +kubebuilder:validation:Enum={"Basic","Free","Premium","Standard"}
type SignalRSkuTier string

const (
	SignalRSkuTier_Basic    = SignalRSkuTier("Basic")
	SignalRSkuTier_Free     = SignalRSkuTier("Free")
	SignalRSkuTier_Premium  = SignalRSkuTier("Premium")
	SignalRSkuTier_Standard = SignalRSkuTier("Standard")
)

// Deprecated version of SignalRTlsSettings. Use v1api20211001.SignalRTlsSettings instead
type SignalRTlsSettings_ARM struct {
	ClientCertEnabled *bool `json:"clientCertEnabled,omitempty"`
}

// Deprecated version of NetworkACL. Use v1api20211001.NetworkACL instead
type NetworkACL_ARM struct {
	Allow []SignalRRequestType `json:"allow,omitempty"`
	Deny  []SignalRRequestType `json:"deny,omitempty"`
}

// Deprecated version of PrivateEndpointACL. Use v1api20211001.PrivateEndpointACL instead
type PrivateEndpointACL_ARM struct {
	Allow []SignalRRequestType `json:"allow,omitempty"`
	Deny  []SignalRRequestType `json:"deny,omitempty"`
	Name  *string              `json:"name,omitempty"`
}

// Deprecated version of ResourceLogCategory. Use v1api20211001.ResourceLogCategory instead
type ResourceLogCategory_ARM struct {
	Enabled *string `json:"enabled,omitempty"`
	Name    *string `json:"name,omitempty"`
}

// Deprecated version of UpstreamTemplate. Use v1api20211001.UpstreamTemplate instead
type UpstreamTemplate_ARM struct {
	Auth            *UpstreamAuthSettings_ARM `json:"auth,omitempty"`
	CategoryPattern *string                   `json:"categoryPattern,omitempty"`
	EventPattern    *string                   `json:"eventPattern,omitempty"`
	HubPattern      *string                   `json:"hubPattern,omitempty"`
	UrlTemplate     *string                   `json:"urlTemplate,omitempty"`
}

// Deprecated version of UpstreamAuthSettings. Use v1api20211001.UpstreamAuthSettings instead
type UpstreamAuthSettings_ARM struct {
	ManagedIdentity *ManagedIdentitySettings_ARM `json:"managedIdentity,omitempty"`
	Type            *UpstreamAuthType            `json:"type,omitempty"`
}

// Deprecated version of ManagedIdentitySettings. Use v1api20211001.ManagedIdentitySettings instead
type ManagedIdentitySettings_ARM struct {
	Resource *string `json:"resource,omitempty"`
}
