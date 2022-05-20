// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of NetworkSecurityGroups_Spec. Use v1beta20201101.NetworkSecurityGroups_Spec instead
type NetworkSecurityGroups_SpecARM struct {
	Location *string           `json:"location,omitempty"`
	Name     string            `json:"name,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NetworkSecurityGroups_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (groups NetworkSecurityGroups_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (groups *NetworkSecurityGroups_SpecARM) GetName() string {
	return groups.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkSecurityGroups"
func (groups *NetworkSecurityGroups_SpecARM) GetType() string {
	return "Microsoft.Network/networkSecurityGroups"
}
