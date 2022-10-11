// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type SignalR_Spec_ARM struct {
	// Identity: A class represent managed identities used for request and response
	Identity *ManagedIdentity_ARM `json:"identity,omitempty"`
	Kind     *SignalR_Kind_Spec   `json:"kind,omitempty"`

	// Location: The GEO location of the resource. e.g. West US | East US | North Central US | South Central US.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource.
	Name string `json:"name,omitempty"`

	// Properties: A class that describes the properties of the resource
	Properties *SignalRProperties_ARM `json:"properties,omitempty"`

	// Sku: The billing information of the resource.
	Sku *ResourceSku_ARM `json:"sku,omitempty"`

	// Tags: Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]string `json:"tags,omitempty"`
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

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentity
type ManagedIdentity_ARM struct {
	Type *ManagedIdentity_Type `json:"type,omitempty"`

	// UserAssignedIdentities: Get or set the user assigned identities
	UserAssignedIdentities map[string]v1.JSON `json:"userAssignedIdentities,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceSku
type ResourceSku_ARM struct {
	// Capacity: Optional, integer. The unit count of the resource. 1 by default.
	// If present, following values are allowed:
	// Free: 1
	// Standard: 1,2,5,10,20,50,100
	Capacity *int `json:"capacity,omitempty"`

	// Name: The name of the SKU. Required.
	// Allowed values: Standard_S1, Free_F1
	Name *string           `json:"name,omitempty"`
	Tier *ResourceSku_Tier `json:"tier,omitempty"`
}

// +kubebuilder:validation:Enum={"RawWebSockets","SignalR"}
type SignalR_Kind_Spec string

const (
	SignalR_Kind_Spec_RawWebSockets = SignalR_Kind_Spec("RawWebSockets")
	SignalR_Kind_Spec_SignalR       = SignalR_Kind_Spec("SignalR")
)

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRProperties
type SignalRProperties_ARM struct {
	// Cors: Cross-Origin Resource Sharing (CORS) settings.
	Cors *SignalRCorsSettings_ARM `json:"cors,omitempty"`

	// DisableAadAuth: DisableLocalAuth
	// Enable or disable aad auth
	// When set as true, connection with AuthType=aad won't work.
	DisableAadAuth *bool `json:"disableAadAuth,omitempty"`

	// DisableLocalAuth: DisableLocalAuth
	// Enable or disable local auth with AccessKey
	// When set as true, connection with AccessKey=xxx won't work.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// Features: List of the featureFlags.
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, its globally default value will be used
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features []SignalRFeature_ARM `json:"features,omitempty"`

	// NetworkACLs: Network ACLs for the resource
	NetworkACLs *SignalRNetworkACLs_ARM `json:"networkACLs,omitempty"`

	// PublicNetworkAccess: Enable or disable public network access. Default to "Enabled".
	// When it's Enabled, network ACLs still apply.
	// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty"`

	// ResourceLogConfiguration: Resource log configuration of a Microsoft.SignalRService resource.
	ResourceLogConfiguration *ResourceLogConfiguration_ARM `json:"resourceLogConfiguration,omitempty"`

	// Tls: TLS settings for the resource
	Tls *SignalRTlsSettings_ARM `json:"tls,omitempty"`

	// Upstream: The settings for the Upstream when the service is in server-less mode.
	Upstream *ServerlessUpstreamSettings_ARM `json:"upstream,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","UserAssigned"}
type ManagedIdentity_Type string

const (
	ManagedIdentity_Type_None           = ManagedIdentity_Type("None")
	ManagedIdentity_Type_SystemAssigned = ManagedIdentity_Type("SystemAssigned")
	ManagedIdentity_Type_UserAssigned   = ManagedIdentity_Type("UserAssigned")
)

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogConfiguration
type ResourceLogConfiguration_ARM struct {
	// Categories: Gets or sets the list of category configurations.
	Categories []ResourceLogCategory_ARM `json:"categories,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","Free","Premium","Standard"}
type ResourceSku_Tier string

const (
	ResourceSku_Tier_Basic    = ResourceSku_Tier("Basic")
	ResourceSku_Tier_Free     = ResourceSku_Tier("Free")
	ResourceSku_Tier_Premium  = ResourceSku_Tier("Premium")
	ResourceSku_Tier_Standard = ResourceSku_Tier("Standard")
)

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ServerlessUpstreamSettings
type ServerlessUpstreamSettings_ARM struct {
	// Templates: Gets or sets the list of Upstream URL templates. Order matters, and the first matching template takes effects.
	Templates []UpstreamTemplate_ARM `json:"templates,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRCorsSettings
type SignalRCorsSettings_ARM struct {
	// AllowedOrigins: Gets or sets the list of origins that should be allowed to make cross-origin calls (for example:
	// http://example.com:12345). Use "*" to allow all. If omitted, allow all by default.
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRFeature
type SignalRFeature_ARM struct {
	Flag *SignalRFeature_Flag `json:"flag,omitempty"`

	// Properties: Optional properties related to this feature.
	Properties map[string]string `json:"properties,omitempty"`

	// Value: Value of the feature flag. See Azure SignalR service document https://docs.microsoft.com/azure/azure-signalr/ for
	// allowed values.
	Value *string `json:"value,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRNetworkACLs
type SignalRNetworkACLs_ARM struct {
	DefaultAction *SignalRNetworkACLs_DefaultAction `json:"defaultAction,omitempty"`

	// PrivateEndpoints: ACLs for requests from private endpoints
	PrivateEndpoints []PrivateEndpointACL_ARM `json:"privateEndpoints,omitempty"`

	// PublicNetwork: Network ACL
	PublicNetwork *NetworkACL_ARM `json:"publicNetwork,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRTlsSettings
type SignalRTlsSettings_ARM struct {
	// ClientCertEnabled: Request client certificate during TLS handshake if enabled
	ClientCertEnabled *bool `json:"clientCertEnabled,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/NetworkACL
type NetworkACL_ARM struct {
	// Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Allow []NetworkACL_Allow `json:"allow,omitempty"`

	// Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Deny []NetworkACL_Deny `json:"deny,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/PrivateEndpointACL
type PrivateEndpointACL_ARM struct {
	// Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Allow []PrivateEndpointACL_Allow `json:"allow,omitempty"`

	// Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Deny []PrivateEndpointACL_Deny `json:"deny,omitempty"`

	// Name: Name of the private endpoint connection
	Name *string `json:"name,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogCategory
type ResourceLogCategory_ARM struct {
	// Enabled: Indicates whether or the resource log category is enabled.
	// Available values: true, false.
	// Case insensitive.
	Enabled *string `json:"enabled,omitempty"`

	// Name: Gets or sets the resource log category's name.
	// Available values: ConnectivityLogs, MessagingLogs.
	// Case insensitive.
	Name *string `json:"name,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamTemplate
type UpstreamTemplate_ARM struct {
	// Auth: Upstream auth settings. If not set, no auth is used for upstream messages.
	Auth *UpstreamAuthSettings_ARM `json:"auth,omitempty"`

	// CategoryPattern: Gets or sets the matching pattern for category names. If not set, it matches any category.
	// There are 3 kind of patterns supported:
	// 1. "*", it to matches any category name
	// 2. Combine multiple categories with ",", for example "connections,messages", it matches category "connections" and
	// "messages"
	// 3. The single category name, for example, "connections", it matches the category "connections"
	CategoryPattern *string `json:"categoryPattern,omitempty"`

	// EventPattern: Gets or sets the matching pattern for event names. If not set, it matches any event.
	// There are 3 kind of patterns supported:
	// 1. "*", it to matches any event name
	// 2. Combine multiple events with ",", for example "connect,disconnect", it matches event "connect" and "disconnect"
	// 3. The single event name, for example, "connect", it matches "connect"
	EventPattern *string `json:"eventPattern,omitempty"`

	// HubPattern: Gets or sets the matching pattern for hub names. If not set, it matches any hub.
	// There are 3 kind of patterns supported:
	// 1. "*", it to matches any hub name
	// 2. Combine multiple hubs with ",", for example "hub1,hub2", it matches "hub1" and "hub2"
	// 3. The single hub name, for example, "hub1", it matches "hub1"
	HubPattern *string `json:"hubPattern,omitempty"`

	// UrlTemplate: Gets or sets the Upstream URL template. You can use 3 predefined parameters {hub}, {category} {event}
	// inside the template, the value of the Upstream URL is dynamically calculated when the client request comes in.
	// For example, if the urlTemplate is `http://example.com/{hub}/api/{event}`, with a client request from hub `chat`
	// connects, it will first POST to this URL: `http://example.com/chat/api/connect`.
	UrlTemplate *string `json:"urlTemplate,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamAuthSettings
type UpstreamAuthSettings_ARM struct {
	// ManagedIdentity: Managed identity settings for upstream.
	ManagedIdentity *ManagedIdentitySettings_ARM `json:"managedIdentity,omitempty"`
	Type            *UpstreamAuthSettings_Type   `json:"type,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentitySettings
type ManagedIdentitySettings_ARM struct {
	// Resource: The Resource indicating the App ID URI of the target resource.
	// It also appears in the aud (audience) claim of the issued token.
	Resource *string `json:"resource,omitempty"`
}
