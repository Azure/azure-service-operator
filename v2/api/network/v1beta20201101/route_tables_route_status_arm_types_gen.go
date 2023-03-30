// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

// Deprecated version of RouteTables_Route_STATUS. Use v1api20201101.RouteTables_Route_STATUS instead
type RouteTables_Route_STATUS_ARM struct {
	Etag       *string                           `json:"etag,omitempty"`
	Id         *string                           `json:"id,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Properties *RoutePropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                           `json:"type,omitempty"`
}

// Deprecated version of RoutePropertiesFormat_STATUS. Use v1api20201101.RoutePropertiesFormat_STATUS instead
type RoutePropertiesFormat_STATUS_ARM struct {
	AddressPrefix     *string                   `json:"addressPrefix,omitempty"`
	HasBgpOverride    *bool                     `json:"hasBgpOverride,omitempty"`
	NextHopIpAddress  *string                   `json:"nextHopIpAddress,omitempty"`
	NextHopType       *RouteNextHopType_STATUS  `json:"nextHopType,omitempty"`
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}
