// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of UserAssignedIdentities_Spec. Use v1beta20181130.UserAssignedIdentities_Spec instead
type UserAssignedIdentities_SpecARM struct {
	Location *string           `json:"location,omitempty"`
	Name     string            `json:"name,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &UserAssignedIdentities_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-11-30"
func (identities UserAssignedIdentities_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (identities *UserAssignedIdentities_SpecARM) GetName() string {
	return identities.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities"
func (identities *UserAssignedIdentities_SpecARM) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities"
}
