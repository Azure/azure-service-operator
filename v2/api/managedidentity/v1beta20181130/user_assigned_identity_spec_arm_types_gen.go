// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20181130

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of UserAssignedIdentity_Spec. Use v1api20181130.UserAssignedIdentity_Spec instead
type UserAssignedIdentity_Spec_ARM struct {
	Location *string           `json:"location,omitempty"`
	Name     string            `json:"name,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &UserAssignedIdentity_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-11-30"
func (identity UserAssignedIdentity_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (identity *UserAssignedIdentity_Spec_ARM) GetName() string {
	return identity.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities"
func (identity *UserAssignedIdentity_Spec_ARM) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities"
}