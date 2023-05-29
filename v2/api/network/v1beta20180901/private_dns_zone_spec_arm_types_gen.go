// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180901

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of PrivateDnsZone_Spec. Use v1api20180901.PrivateDnsZone_Spec instead
type PrivateDnsZone_Spec_ARM struct {
	Etag     *string           `json:"etag,omitempty"`
	Location *string           `json:"location,omitempty"`
	Name     string            `json:"name,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PrivateDnsZone_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-09-01"
func (zone PrivateDnsZone_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (zone *PrivateDnsZone_Spec_ARM) GetName() string {
	return zone.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateDnsZones"
func (zone *PrivateDnsZone_Spec_ARM) GetType() string {
	return "Microsoft.Network/privateDnsZones"
}