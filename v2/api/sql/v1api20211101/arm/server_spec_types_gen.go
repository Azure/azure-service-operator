// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Server_Spec struct {
	// Identity: The Azure Active Directory identity of the server.
	Identity *ResourceIdentity `json:"identity,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *ServerProperties `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Server_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (server Server_Spec) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (server *Server_Spec) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers"
func (server *Server_Spec) GetType() string {
	return "Microsoft.Sql/servers"
}

// Azure Active Directory identity configuration for a resource.
type ResourceIdentity struct {
	// Type: The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active
	// Directory principal for the resource.
	Type                   *ResourceIdentity_Type                 `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// The properties of a server.
type ServerProperties struct {
	// AdministratorLogin: Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// Administrators: The Azure Active Directory administrator of the server.
	Administrators *ServerExternalAdministrator `json:"administrators,omitempty"`

	// FederatedClientId: The Client id used for cross tenant CMK scenario
	FederatedClientId *string `json:"federatedClientId,omitempty"`

	// KeyId: A CMK URI of the key to use for encryption.
	KeyId *string `json:"keyId,omitempty"`

	// MinimalTlsVersion: Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion             *string `json:"minimalTlsVersion,omitempty"`
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this server.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *ServerProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// RestrictOutboundNetworkAccess: Whether or not to restrict outbound network access for this server.  Value is optional
	// but if passed in, must be 'Enabled' or 'Disabled'
	RestrictOutboundNetworkAccess *ServerProperties_RestrictOutboundNetworkAccess `json:"restrictOutboundNetworkAccess,omitempty"`

	// Version: The version of the server.
	Version *string `json:"version,omitempty"`
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
var resourceIdentity_Type_Values = map[string]ResourceIdentity_Type{
	"none":                        ResourceIdentity_Type_None,
	"systemassigned":              ResourceIdentity_Type_SystemAssigned,
	"systemassigned,userassigned": ResourceIdentity_Type_SystemAssignedUserAssigned,
	"userassigned":                ResourceIdentity_Type_UserAssigned,
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

	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerProperties_PublicNetworkAccess string

const (
	ServerProperties_PublicNetworkAccess_Disabled = ServerProperties_PublicNetworkAccess("Disabled")
	ServerProperties_PublicNetworkAccess_Enabled  = ServerProperties_PublicNetworkAccess("Enabled")
)

// Mapping from string to ServerProperties_PublicNetworkAccess
var serverProperties_PublicNetworkAccess_Values = map[string]ServerProperties_PublicNetworkAccess{
	"disabled": ServerProperties_PublicNetworkAccess_Disabled,
	"enabled":  ServerProperties_PublicNetworkAccess_Enabled,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerProperties_RestrictOutboundNetworkAccess string

const (
	ServerProperties_RestrictOutboundNetworkAccess_Disabled = ServerProperties_RestrictOutboundNetworkAccess("Disabled")
	ServerProperties_RestrictOutboundNetworkAccess_Enabled  = ServerProperties_RestrictOutboundNetworkAccess("Enabled")
)

// Mapping from string to ServerProperties_RestrictOutboundNetworkAccess
var serverProperties_RestrictOutboundNetworkAccess_Values = map[string]ServerProperties_RestrictOutboundNetworkAccess{
	"disabled": ServerProperties_RestrictOutboundNetworkAccess_Disabled,
	"enabled":  ServerProperties_RestrictOutboundNetworkAccess_Enabled,
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
}

// +kubebuilder:validation:Enum={"ActiveDirectory"}
type ServerExternalAdministrator_AdministratorType string

const ServerExternalAdministrator_AdministratorType_ActiveDirectory = ServerExternalAdministrator_AdministratorType("ActiveDirectory")

// Mapping from string to ServerExternalAdministrator_AdministratorType
var serverExternalAdministrator_AdministratorType_Values = map[string]ServerExternalAdministrator_AdministratorType{
	"activedirectory": ServerExternalAdministrator_AdministratorType_ActiveDirectory,
}

// +kubebuilder:validation:Enum={"Application","Group","User"}
type ServerExternalAdministrator_PrincipalType string

const (
	ServerExternalAdministrator_PrincipalType_Application = ServerExternalAdministrator_PrincipalType("Application")
	ServerExternalAdministrator_PrincipalType_Group       = ServerExternalAdministrator_PrincipalType("Group")
	ServerExternalAdministrator_PrincipalType_User        = ServerExternalAdministrator_PrincipalType("User")
)

// Mapping from string to ServerExternalAdministrator_PrincipalType
var serverExternalAdministrator_PrincipalType_Values = map[string]ServerExternalAdministrator_PrincipalType{
	"application": ServerExternalAdministrator_PrincipalType_Application,
	"group":       ServerExternalAdministrator_PrincipalType_Group,
	"user":        ServerExternalAdministrator_PrincipalType_User,
}
