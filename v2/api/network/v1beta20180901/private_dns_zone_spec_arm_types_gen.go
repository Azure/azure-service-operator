// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180901

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type PrivateDnsZone_Spec_ARM struct {
	// Etag: The ETag of the Private DNS zone.
	Etag *string `json:"etag,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
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
