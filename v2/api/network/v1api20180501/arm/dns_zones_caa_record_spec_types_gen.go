// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DnsZonesCAARecord_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of the record set.
	Properties *RecordSetProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DnsZonesCAARecord_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-05-01"
func (record DnsZonesCAARecord_Spec) GetAPIVersion() string {
	return "2018-05-01"
}

// GetName returns the Name of the resource
func (record *DnsZonesCAARecord_Spec) GetName() string {
	return record.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsZones/CAA"
func (record *DnsZonesCAARecord_Spec) GetType() string {
	return "Microsoft.Network/dnsZones/CAA"
}
