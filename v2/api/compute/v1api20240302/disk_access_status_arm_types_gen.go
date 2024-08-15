// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240302

// disk access resource.
type DiskAccess_STATUS_ARM struct {
	// ExtendedLocation: The extended location where the disk access will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`

	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// Name: Resource name
	Name       *string                          `json:"name,omitempty"`
	Properties *DiskAccessProperties_STATUS_ARM `json:"properties,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

type DiskAccessProperties_STATUS_ARM struct {
	// PrivateEndpointConnections: A readonly collection of private endpoint connections created on the disk. Currently only
	// one endpoint connection is supported.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: The disk access resource provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// TimeCreated: The time when the disk access was created.
	TimeCreated *string `json:"timeCreated,omitempty"`
}

// The complex type of the extended location.
type ExtendedLocation_STATUS_ARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType_STATUS `json:"type,omitempty"`
}

// The type of extendedLocation.
type ExtendedLocationType_STATUS string

const ExtendedLocationType_STATUS_EdgeZone = ExtendedLocationType_STATUS("EdgeZone")

// Mapping from string to ExtendedLocationType_STATUS
var extendedLocationType_STATUS_Values = map[string]ExtendedLocationType_STATUS{
	"edgezone": ExtendedLocationType_STATUS_EdgeZone,
}

// The Private Endpoint Connection resource.
type PrivateEndpointConnection_STATUS_ARM struct {
	// Id: private endpoint connection Id
	Id *string `json:"id,omitempty"`
}
