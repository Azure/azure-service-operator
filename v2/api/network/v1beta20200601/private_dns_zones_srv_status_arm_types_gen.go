// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

type PrivateDnsZones_SRV_STATUS_ARM struct {
	// Etag: The ETag of the record set.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource Id for the resource. Example -
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	Id *string `json:"id,omitempty"`

	// Name: The name of the record set.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of the record set.
	Properties *RecordSetProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `json:"type,omitempty"`
}
