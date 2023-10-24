// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Trafficmanagerprofiles_AzureEndpoint_Spec_ARM struct {
	// Name: The name of the resource
	Name string `json:"name,omitempty"`

	// Properties: The properties of the Traffic Manager endpoint.
	Properties *EndpointProperties_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Trafficmanagerprofiles_AzureEndpoint_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-04-01"
func (endpoint Trafficmanagerprofiles_AzureEndpoint_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (endpoint *Trafficmanagerprofiles_AzureEndpoint_Spec_ARM) GetName() string {
	return endpoint.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/trafficmanagerprofiles/AzureEndpoints"
func (endpoint *Trafficmanagerprofiles_AzureEndpoint_Spec_ARM) GetType() string {
	return "Microsoft.Network/trafficmanagerprofiles/AzureEndpoints"
}

// Class representing a Traffic Manager endpoint properties.
type EndpointProperties_ARM struct {
	// AlwaysServe: If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in
	// the traffic routing method.
	AlwaysServe *EndpointProperties_AlwaysServe `json:"alwaysServe,omitempty"`

	// CustomHeaders: List of custom headers.
	CustomHeaders []EndpointProperties_CustomHeaders_ARM `json:"customHeaders,omitempty"`

	// EndpointLocation: Specifies the location of the external or nested endpoints when using the 'Performance' traffic
	// routing method.
	EndpointLocation *string `json:"endpointLocation,omitempty"`

	// EndpointMonitorStatus: The monitoring status of the endpoint.
	EndpointMonitorStatus *EndpointProperties_EndpointMonitorStatus `json:"endpointMonitorStatus,omitempty"`

	// EndpointStatus: The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included
	// in the traffic routing method.
	EndpointStatus *EndpointProperties_EndpointStatus `json:"endpointStatus,omitempty"`

	// GeoMapping: The list of countries/regions mapped to this endpoint when using the 'Geographic' traffic routing method.
	// Please consult Traffic Manager Geographic documentation for a full list of accepted values.
	GeoMapping []string `json:"geoMapping,omitempty"`

	// MinChildEndpoints: The minimum number of endpoints that must be available in the child profile in order for the parent
	// profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints *int `json:"minChildEndpoints,omitempty"`

	// MinChildEndpointsIPv4: The minimum number of IPv4 (DNS record type A) endpoints that must be available in the child
	// profile in order for the parent profile to be considered available. Only applicable to endpoint of type
	// 'NestedEndpoints'.
	MinChildEndpointsIPv4 *int `json:"minChildEndpointsIPv4,omitempty"`

	// MinChildEndpointsIPv6: The minimum number of IPv6 (DNS record type AAAA) endpoints that must be available in the child
	// profile in order for the parent profile to be considered available. Only applicable to endpoint of type
	// 'NestedEndpoints'.
	MinChildEndpointsIPv6 *int `json:"minChildEndpointsIPv6,omitempty"`

	// Priority: The priority of this endpoint when using the 'Priority' traffic routing method. Possible values are from 1 to
	// 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all
	// endpoints, and no two endpoints can share the same priority value.
	Priority *int `json:"priority,omitempty"`

	// Subnets: The list of subnets, IP addresses, and/or address ranges mapped to this endpoint when using the 'Subnet'
	// traffic routing method. An empty list will match all ranges not covered by other endpoints.
	Subnets []EndpointProperties_Subnets_ARM `json:"subnets,omitempty"`

	// Target: The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses
	// to direct traffic to this endpoint.
	Target           *string `json:"target,omitempty"`
	TargetResourceId *string `json:"targetResourceId,omitempty"`

	// Weight: The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight *int `json:"weight,omitempty"`
}

type EndpointProperties_CustomHeaders_ARM struct {
	// Name: Header name.
	Name *string `json:"name,omitempty"`

	// Value: Header value.
	Value *string `json:"value,omitempty"`
}

type EndpointProperties_Subnets_ARM struct {
	// First: First address in the subnet.
	First *string `json:"first,omitempty"`

	// Last: Last address in the subnet.
	Last *string `json:"last,omitempty"`

	// Scope: Block size (number of leading bits in the subnet mask).
	Scope *int `json:"scope,omitempty"`
}
