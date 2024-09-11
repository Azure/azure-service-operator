// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Server_Spec_ARM struct {
	// Identity: The Azure Active Directory identity of the server.
	Identity *ResourceIdentity_ARM `json:"identity,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *ServerProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Server_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (server Server_Spec_ARM) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (server *Server_Spec_ARM) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers"
func (server *Server_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers"
}

// Azure Active Directory identity configuration for a resource.
type ResourceIdentity_ARM struct {
	// Type: The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active
	// Directory principal for the resource.
	Type                   *ResourceIdentity_Type_ARM                 `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// The properties of a server.
type ServerProperties_ARM struct {
	// AdministratorLogin: Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// Administrators: The Azure Active Directory administrator of the server.
	Administrators *ServerExternalAdministrator_ARM `json:"administrators,omitempty"`

	// FederatedClientId: The Client id used for cross tenant CMK scenario
	FederatedClientId *string `json:"federatedClientId,omitempty"`

	// KeyId: A CMK URI of the key to use for encryption.
	KeyId *string `json:"keyId,omitempty"`

	// MinimalTlsVersion: Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion             *string `json:"minimalTlsVersion,omitempty"`
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this server.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *ServerProperties_PublicNetworkAccess_ARM `json:"publicNetworkAccess,omitempty"`

	// RestrictOutboundNetworkAccess: Whether or not to restrict outbound network access for this server.  Value is optional
	// but if passed in, must be 'Enabled' or 'Disabled'
	RestrictOutboundNetworkAccess *ServerProperties_RestrictOutboundNetworkAccess_ARM `json:"restrictOutboundNetworkAccess,omitempty"`

	// Version: The version of the server.
	Version *string `json:"version,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ResourceIdentity_Type_ARM string

const (
	ResourceIdentity_Type_ARM_None                       = ResourceIdentity_Type_ARM("None")
	ResourceIdentity_Type_ARM_SystemAssigned             = ResourceIdentity_Type_ARM("SystemAssigned")
	ResourceIdentity_Type_ARM_SystemAssignedUserAssigned = ResourceIdentity_Type_ARM("SystemAssigned,UserAssigned")
	ResourceIdentity_Type_ARM_UserAssigned               = ResourceIdentity_Type_ARM("UserAssigned")
)

// Mapping from string to ResourceIdentity_Type_ARM
var resourceIdentity_Type_ARM_Values = map[string]ResourceIdentity_Type_ARM{
	"none":                        ResourceIdentity_Type_ARM_None,
	"systemassigned":              ResourceIdentity_Type_ARM_SystemAssigned,
	"systemassigned,userassigned": ResourceIdentity_Type_ARM_SystemAssignedUserAssigned,
	"userassigned":                ResourceIdentity_Type_ARM_UserAssigned,
}

// Properties of a active directory administrator.
type ServerExternalAdministrator_ARM struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType *ServerExternalAdministrator_AdministratorType_ARM `json:"administratorType,omitempty"`

	// AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// PrincipalType: Principal Type of the sever administrator.
	PrincipalType *ServerExternalAdministrator_PrincipalType_ARM `json:"principalType,omitempty"`

	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerProperties_PublicNetworkAccess_ARM string

const (
	ServerProperties_PublicNetworkAccess_ARM_Disabled = ServerProperties_PublicNetworkAccess_ARM("Disabled")
	ServerProperties_PublicNetworkAccess_ARM_Enabled  = ServerProperties_PublicNetworkAccess_ARM("Enabled")
)

// Mapping from string to ServerProperties_PublicNetworkAccess_ARM
var serverProperties_PublicNetworkAccess_ARM_Values = map[string]ServerProperties_PublicNetworkAccess_ARM{
	"disabled": ServerProperties_PublicNetworkAccess_ARM_Disabled,
	"enabled":  ServerProperties_PublicNetworkAccess_ARM_Enabled,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerProperties_RestrictOutboundNetworkAccess_ARM string

const (
	ServerProperties_RestrictOutboundNetworkAccess_ARM_Disabled = ServerProperties_RestrictOutboundNetworkAccess_ARM("Disabled")
	ServerProperties_RestrictOutboundNetworkAccess_ARM_Enabled  = ServerProperties_RestrictOutboundNetworkAccess_ARM("Enabled")
)

// Mapping from string to ServerProperties_RestrictOutboundNetworkAccess_ARM
var serverProperties_RestrictOutboundNetworkAccess_ARM_Values = map[string]ServerProperties_RestrictOutboundNetworkAccess_ARM{
	"disabled": ServerProperties_RestrictOutboundNetworkAccess_ARM_Disabled,
	"enabled":  ServerProperties_RestrictOutboundNetworkAccess_ARM_Enabled,
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// +kubebuilder:validation:Enum={"ActiveDirectory"}
type ServerExternalAdministrator_AdministratorType_ARM string

const ServerExternalAdministrator_AdministratorType_ARM_ActiveDirectory = ServerExternalAdministrator_AdministratorType_ARM("ActiveDirectory")

// Mapping from string to ServerExternalAdministrator_AdministratorType_ARM
var serverExternalAdministrator_AdministratorType_ARM_Values = map[string]ServerExternalAdministrator_AdministratorType_ARM{
	"activedirectory": ServerExternalAdministrator_AdministratorType_ARM_ActiveDirectory,
}

// +kubebuilder:validation:Enum={"Application","Group","User"}
type ServerExternalAdministrator_PrincipalType_ARM string

const (
	ServerExternalAdministrator_PrincipalType_ARM_Application = ServerExternalAdministrator_PrincipalType_ARM("Application")
	ServerExternalAdministrator_PrincipalType_ARM_Group       = ServerExternalAdministrator_PrincipalType_ARM("Group")
	ServerExternalAdministrator_PrincipalType_ARM_User        = ServerExternalAdministrator_PrincipalType_ARM("User")
)

// Mapping from string to ServerExternalAdministrator_PrincipalType_ARM
var serverExternalAdministrator_PrincipalType_ARM_Values = map[string]ServerExternalAdministrator_PrincipalType_ARM{
	"application": ServerExternalAdministrator_PrincipalType_ARM_Application,
	"group":       ServerExternalAdministrator_PrincipalType_ARM_Group,
	"user":        ServerExternalAdministrator_PrincipalType_ARM_User,
}
