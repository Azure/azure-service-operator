// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101preview

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/Servers.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Server_Spec   `json:"spec,omitempty"`
	Status            Server_STATUS `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/Servers.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// +kubebuilder:validation:Enum={"2020-11-01-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-11-01-preview")

// Mapping from string to APIVersion
var APIVersion_Cache = map[string]APIVersion{
	"2020-11-01-preview": APIVersion_Value,
}

type Server_Spec struct {
	v1.ResourceSpec `json:",inline,omitempty"`
	ForProvider     ServerParameters `json:"forProvider,omitempty"`
}

type Server_STATUS struct {
	v1.ResourceStatus `json:",inline,omitempty"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

type ServerObservation struct {
	// AdministratorLogin: Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// Administrators: The Azure Active Directory identity of the server.
	Administrators *ServerExternalAdministrator_STATUS `json:"administrators,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of the server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Identity: The Azure Active Directory identity of the server.
	Identity *ResourceIdentity_STATUS `json:"identity,omitempty"`

	// KeyId: A CMK URI of the key to use for encryption.
	KeyId *string `json:"keyId,omitempty"`

	// Kind: Kind of sql server. This is metadata used for the Azure portal experience.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// MinimalTlsVersion: Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `json:"minimalTlsVersion,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// PrimaryUserAssignedIdentityId: The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections on a server
	PrivateEndpointConnections []ServerPrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this server.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *ServerProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// State: The state of the server.
	State *string `json:"state,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Version: The version of the server.
	Version *string `json:"version,omitempty"`

	// WorkspaceFeature: Whether or not existing server has a workspace created and if it allows connection from workspace
	WorkspaceFeature *ServerProperties_WorkspaceFeature_STATUS `json:"workspaceFeature,omitempty"`
}

type ServerParameters struct {
	// AdministratorLogin: Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// Administrators: The Azure Active Directory identity of the server.
	Administrators *ServerExternalAdministrator `json:"administrators,omitempty"`

	// Identity: The Azure Active Directory identity of the server.
	Identity *ResourceIdentity `json:"identity,omitempty"`

	// KeyId: A CMK URI of the key to use for encryption.
	KeyId *string `json:"keyId,omitempty"`

	// +kubebuilder:validation:Required
	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// MinimalTlsVersion: Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `json:"minimalTlsVersion,omitempty"`
	Name              string  `json:"name,omitempty"`

	// PrimaryUserAssignedIdentityId: The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this server.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess       *ServerProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
	ResourceGroupName         string                                `json:"resourceGroupName,omitempty"`
	ResourceGroupNameRef      *v1.Reference                         `json:"resourceGroupNameRef,omitempty"`
	ResourceGroupNameSelector *v1.Selector                          `json:"resourceGroupNameSelector,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Version: The version of the server.
	Version *string `json:"version,omitempty"`
}

// Azure Active Directory identity configuration for a resource.
type ResourceIdentity struct {
	// Type: The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active
	// Directory principal for the resource.
	Type *ResourceIdentity_Type `json:"type,omitempty"`
}

// Azure Active Directory identity configuration for a resource.
type ResourceIdentity_STATUS struct {
	// PrincipalId: The Azure Active Directory principal id.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The Azure Active Directory tenant id.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active
	// Directory principal for the resource.
	Type *ResourceIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: The resource ids of the user assigned identities to use
	UserAssignedIdentities map[string]UserIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Properties of a active directory administrator.
type ServerExternalAdministrator struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType *ServerExternalAdministrator_AdministratorType `json:"administratorType,omitempty"`

	// AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// PrincipalType: Principal Type of the sever administrator.
	PrincipalType *ServerExternalAdministrator_PrincipalType `json:"principalType,omitempty"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

// Properties of a active directory administrator.
type ServerExternalAdministrator_STATUS struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType *ServerExternalAdministrator_AdministratorType_STATUS `json:"administratorType,omitempty"`

	// AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// PrincipalType: Principal Type of the sever administrator.
	PrincipalType *ServerExternalAdministrator_PrincipalType_STATUS `json:"principalType,omitempty"`

	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

// A private endpoint connection under a server
type ServerPrivateEndpointConnection_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Properties: Private endpoint connection properties
	Properties *PrivateEndpointConnectionProperties_STATUS `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerProperties_PublicNetworkAccess string

const (
	ServerProperties_PublicNetworkAccess_Disabled = ServerProperties_PublicNetworkAccess("Disabled")
	ServerProperties_PublicNetworkAccess_Enabled  = ServerProperties_PublicNetworkAccess("Enabled")
)

// Mapping from string to ServerProperties_PublicNetworkAccess
var ServerProperties_PublicNetworkAccess_Cache = map[string]ServerProperties_PublicNetworkAccess{
	"disabled": ServerProperties_PublicNetworkAccess_Disabled,
	"enabled":  ServerProperties_PublicNetworkAccess_Enabled,
}

type ServerProperties_PublicNetworkAccess_STATUS string

const (
	ServerProperties_PublicNetworkAccess_STATUS_Disabled = ServerProperties_PublicNetworkAccess_STATUS("Disabled")
	ServerProperties_PublicNetworkAccess_STATUS_Enabled  = ServerProperties_PublicNetworkAccess_STATUS("Enabled")
)

// Mapping from string to ServerProperties_PublicNetworkAccess_STATUS
var ServerProperties_PublicNetworkAccess_STATUS_Cache = map[string]ServerProperties_PublicNetworkAccess_STATUS{
	"disabled": ServerProperties_PublicNetworkAccess_STATUS_Disabled,
	"enabled":  ServerProperties_PublicNetworkAccess_STATUS_Enabled,
}

type ServerProperties_WorkspaceFeature_STATUS string

const (
	ServerProperties_WorkspaceFeature_STATUS_Connected    = ServerProperties_WorkspaceFeature_STATUS("Connected")
	ServerProperties_WorkspaceFeature_STATUS_Disconnected = ServerProperties_WorkspaceFeature_STATUS("Disconnected")
)

// Mapping from string to ServerProperties_WorkspaceFeature_STATUS
var ServerProperties_WorkspaceFeature_STATUS_Cache = map[string]ServerProperties_WorkspaceFeature_STATUS{
	"connected":    ServerProperties_WorkspaceFeature_STATUS_Connected,
	"disconnected": ServerProperties_WorkspaceFeature_STATUS_Disconnected,
}

// Properties of a private endpoint connection.
type PrivateEndpointConnectionProperties_STATUS struct {
	// PrivateEndpoint: Private endpoint which the connection belongs to.
	PrivateEndpoint *PrivateEndpointProperty_STATUS `json:"privateEndpoint,omitempty"`

	// PrivateLinkServiceConnectionState: Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty_STATUS `json:"privateLinkServiceConnectionState,omitempty"`

	// ProvisioningState: State of the private endpoint connection.
	ProvisioningState *PrivateEndpointConnectionProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ResourceIdentity_Type string

const (
	ResourceIdentity_Type_None                       = ResourceIdentity_Type("None")
	ResourceIdentity_Type_SystemAssigned             = ResourceIdentity_Type("SystemAssigned")
	ResourceIdentity_Type_SystemAssignedUserAssigned = ResourceIdentity_Type("SystemAssigned,UserAssigned")
	ResourceIdentity_Type_UserAssigned               = ResourceIdentity_Type("UserAssigned")
)

// Mapping from string to ResourceIdentity_Type
var ResourceIdentity_Type_Cache = map[string]ResourceIdentity_Type{
	"none":                        ResourceIdentity_Type_None,
	"systemassigned":              ResourceIdentity_Type_SystemAssigned,
	"systemassigned,userassigned": ResourceIdentity_Type_SystemAssignedUserAssigned,
	"userassigned":                ResourceIdentity_Type_UserAssigned,
}

type ResourceIdentity_Type_STATUS string

const (
	ResourceIdentity_Type_STATUS_None                       = ResourceIdentity_Type_STATUS("None")
	ResourceIdentity_Type_STATUS_SystemAssigned             = ResourceIdentity_Type_STATUS("SystemAssigned")
	ResourceIdentity_Type_STATUS_SystemAssignedUserAssigned = ResourceIdentity_Type_STATUS("SystemAssigned,UserAssigned")
	ResourceIdentity_Type_STATUS_UserAssigned               = ResourceIdentity_Type_STATUS("UserAssigned")
)

// Mapping from string to ResourceIdentity_Type_STATUS
var ResourceIdentity_Type_STATUS_Cache = map[string]ResourceIdentity_Type_STATUS{
	"none":                        ResourceIdentity_Type_STATUS_None,
	"systemassigned":              ResourceIdentity_Type_STATUS_SystemAssigned,
	"systemassigned,userassigned": ResourceIdentity_Type_STATUS_SystemAssignedUserAssigned,
	"userassigned":                ResourceIdentity_Type_STATUS_UserAssigned,
}

// +kubebuilder:validation:Enum={"ActiveDirectory"}
type ServerExternalAdministrator_AdministratorType string

const ServerExternalAdministrator_AdministratorType_ActiveDirectory = ServerExternalAdministrator_AdministratorType("ActiveDirectory")

// Mapping from string to ServerExternalAdministrator_AdministratorType
var ServerExternalAdministrator_AdministratorType_Cache = map[string]ServerExternalAdministrator_AdministratorType{
	"activedirectory": ServerExternalAdministrator_AdministratorType_ActiveDirectory,
}

type ServerExternalAdministrator_AdministratorType_STATUS string

const ServerExternalAdministrator_AdministratorType_STATUS_ActiveDirectory = ServerExternalAdministrator_AdministratorType_STATUS("ActiveDirectory")

// Mapping from string to ServerExternalAdministrator_AdministratorType_STATUS
var ServerExternalAdministrator_AdministratorType_STATUS_Cache = map[string]ServerExternalAdministrator_AdministratorType_STATUS{
	"activedirectory": ServerExternalAdministrator_AdministratorType_STATUS_ActiveDirectory,
}

// +kubebuilder:validation:Enum={"Application","Group","User"}
type ServerExternalAdministrator_PrincipalType string

const (
	ServerExternalAdministrator_PrincipalType_Application = ServerExternalAdministrator_PrincipalType("Application")
	ServerExternalAdministrator_PrincipalType_Group       = ServerExternalAdministrator_PrincipalType("Group")
	ServerExternalAdministrator_PrincipalType_User        = ServerExternalAdministrator_PrincipalType("User")
)

// Mapping from string to ServerExternalAdministrator_PrincipalType
var ServerExternalAdministrator_PrincipalType_Cache = map[string]ServerExternalAdministrator_PrincipalType{
	"application": ServerExternalAdministrator_PrincipalType_Application,
	"group":       ServerExternalAdministrator_PrincipalType_Group,
	"user":        ServerExternalAdministrator_PrincipalType_User,
}

type ServerExternalAdministrator_PrincipalType_STATUS string

const (
	ServerExternalAdministrator_PrincipalType_STATUS_Application = ServerExternalAdministrator_PrincipalType_STATUS("Application")
	ServerExternalAdministrator_PrincipalType_STATUS_Group       = ServerExternalAdministrator_PrincipalType_STATUS("Group")
	ServerExternalAdministrator_PrincipalType_STATUS_User        = ServerExternalAdministrator_PrincipalType_STATUS("User")
)

// Mapping from string to ServerExternalAdministrator_PrincipalType_STATUS
var ServerExternalAdministrator_PrincipalType_STATUS_Cache = map[string]ServerExternalAdministrator_PrincipalType_STATUS{
	"application": ServerExternalAdministrator_PrincipalType_STATUS_Application,
	"group":       ServerExternalAdministrator_PrincipalType_STATUS_Group,
	"user":        ServerExternalAdministrator_PrincipalType_STATUS_User,
}

// Azure Active Directory identity configuration for a resource.
type UserIdentity_STATUS struct {
	// ClientId: The Azure Active Directory client id.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The Azure Active Directory principal id.
	PrincipalId *string `json:"principalId,omitempty"`
}

type PrivateEndpointConnectionProperties_ProvisioningState_STATUS string

const (
	PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Approving = PrivateEndpointConnectionProperties_ProvisioningState_STATUS("Approving")
	PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Dropping  = PrivateEndpointConnectionProperties_ProvisioningState_STATUS("Dropping")
	PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Failed    = PrivateEndpointConnectionProperties_ProvisioningState_STATUS("Failed")
	PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Ready     = PrivateEndpointConnectionProperties_ProvisioningState_STATUS("Ready")
	PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Rejecting = PrivateEndpointConnectionProperties_ProvisioningState_STATUS("Rejecting")
)

// Mapping from string to PrivateEndpointConnectionProperties_ProvisioningState_STATUS
var PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Cache = map[string]PrivateEndpointConnectionProperties_ProvisioningState_STATUS{
	"approving": PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Approving,
	"dropping":  PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Dropping,
	"failed":    PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Failed,
	"ready":     PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Ready,
	"rejecting": PrivateEndpointConnectionProperties_ProvisioningState_STATUS_Rejecting,
}

type PrivateEndpointProperty_STATUS struct {
	// Id: Resource id of the private endpoint.
	Id *string `json:"id,omitempty"`
}

type PrivateLinkServiceConnectionStateProperty_STATUS struct {
	// ActionsRequired: The actions required for private link service connection.
	ActionsRequired *PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS `json:"actionsRequired,omitempty"`

	// Description: The private link service connection description.
	Description *string `json:"description,omitempty"`

	// Status: The private link service connection status.
	Status *PrivateLinkServiceConnectionStateProperty_Status_STATUS `json:"status,omitempty"`
}

type PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS string

const PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS_None = PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS("None")

// Mapping from string to PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS
var PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS_Cache = map[string]PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS{
	"none": PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS_None,
}

type PrivateLinkServiceConnectionStateProperty_Status_STATUS string

const (
	PrivateLinkServiceConnectionStateProperty_Status_STATUS_Approved     = PrivateLinkServiceConnectionStateProperty_Status_STATUS("Approved")
	PrivateLinkServiceConnectionStateProperty_Status_STATUS_Disconnected = PrivateLinkServiceConnectionStateProperty_Status_STATUS("Disconnected")
	PrivateLinkServiceConnectionStateProperty_Status_STATUS_Pending      = PrivateLinkServiceConnectionStateProperty_Status_STATUS("Pending")
	PrivateLinkServiceConnectionStateProperty_Status_STATUS_Rejected     = PrivateLinkServiceConnectionStateProperty_Status_STATUS("Rejected")
)

// Mapping from string to PrivateLinkServiceConnectionStateProperty_Status_STATUS
var PrivateLinkServiceConnectionStateProperty_Status_STATUS_Cache = map[string]PrivateLinkServiceConnectionStateProperty_Status_STATUS{
	"approved":     PrivateLinkServiceConnectionStateProperty_Status_STATUS_Approved,
	"disconnected": PrivateLinkServiceConnectionStateProperty_Status_STATUS_Disconnected,
	"pending":      PrivateLinkServiceConnectionStateProperty_Status_STATUS_Pending,
	"rejected":     PrivateLinkServiceConnectionStateProperty_Status_STATUS_Rejected,
}

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
