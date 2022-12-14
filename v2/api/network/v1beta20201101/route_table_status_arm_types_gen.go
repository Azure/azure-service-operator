// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

// Route table resource.
type RouteTable_STATUS_ARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the route table.
	Properties *RouteTablePropertiesFormat_STATUS_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Route Table resource.
type RouteTablePropertiesFormat_STATUS_ARM struct {
	// DisableBgpRoutePropagation: Whether to disable the routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation *bool `json:"disableBgpRoutePropagation,omitempty"`

	// ProvisioningState: The provisioning state of the route table resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceGuid: The resource GUID property of the route table.
	ResourceGuid *string `json:"resourceGuid,omitempty"`
}
