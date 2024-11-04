// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersAdministrator_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of an administrator.
	Properties *AdministratorProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersAdministrator_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-01"
func (administrator FlexibleServersAdministrator_Spec) GetAPIVersion() string {
	return "2022-01-01"
}

// GetName returns the Name of the resource
func (administrator *FlexibleServersAdministrator_Spec) GetName() string {
	return administrator.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/administrators"
func (administrator *FlexibleServersAdministrator_Spec) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/administrators"
}

// The properties of an administrator.
type AdministratorProperties struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType  *AdministratorProperties_AdministratorType `json:"administratorType,omitempty"`
	IdentityResourceId *string                                    `json:"identityResourceId,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty" optionalConfigMapPair:"Sid"`

	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty" optionalConfigMapPair:"TenantId"`
}

// +kubebuilder:validation:Enum={"ActiveDirectory"}
type AdministratorProperties_AdministratorType string

const AdministratorProperties_AdministratorType_ActiveDirectory = AdministratorProperties_AdministratorType("ActiveDirectory")

// Mapping from string to AdministratorProperties_AdministratorType
var administratorProperties_AdministratorType_Values = map[string]AdministratorProperties_AdministratorType{
	"activedirectory": AdministratorProperties_AdministratorType_ActiveDirectory,
}