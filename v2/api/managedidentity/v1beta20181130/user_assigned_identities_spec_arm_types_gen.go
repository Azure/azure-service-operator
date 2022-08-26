// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20181130

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type UserAssignedIdentity_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &UserAssignedIdentity_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-11-30"
func (identity UserAssignedIdentity_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (identity *UserAssignedIdentity_SpecARM) GetName() string {
	return identity.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities"
func (identity *UserAssignedIdentity_SpecARM) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities"
}
