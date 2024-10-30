// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type TrafficManagerProfilesExternalEndpoint_STATUS struct {
	// Id: Fully qualified resource Id for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of the Traffic Manager endpoint.
	Properties *EndpointProperties_STATUS `json:"properties,omitempty"`

	// Type: The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `json:"type,omitempty"`
}