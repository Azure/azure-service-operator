// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of VirtualNetworksSubnets_Spec. Use v1beta20201101.VirtualNetworksSubnets_Spec instead
type VirtualNetworksSubnets_SpecARM struct {
	Name       string                                     `json:"name,omitempty"`
	Properties *VirtualNetworksSubnets_Spec_PropertiesARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworksSubnets_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (subnets VirtualNetworksSubnets_SpecARM) GetAPIVersion() string {
	return "2020-11-01"
}

// GetName returns the Name of the resource
func (subnets VirtualNetworksSubnets_SpecARM) GetName() string {
	return subnets.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/subnets"
func (subnets VirtualNetworksSubnets_SpecARM) GetType() string {
	return "Microsoft.Network/virtualNetworks/subnets"
}

// Deprecated version of VirtualNetworksSubnets_Spec_Properties. Use v1beta20201101.VirtualNetworksSubnets_Spec_Properties instead
type VirtualNetworksSubnets_Spec_PropertiesARM struct {
	AddressPrefix                     *string                                                 `json:"addressPrefix,omitempty"`
	AddressPrefixes                   []string                                                `json:"addressPrefixes,omitempty"`
	Delegations                       []VirtualNetworksSubnets_Spec_Properties_DelegationsARM `json:"delegations,omitempty"`
	IpAllocations                     []SubResourceARM                                        `json:"ipAllocations,omitempty"`
	NatGateway                        *SubResourceARM                                         `json:"natGateway,omitempty"`
	NetworkSecurityGroup              *SubResourceARM                                         `json:"networkSecurityGroup,omitempty"`
	PrivateEndpointNetworkPolicies    *string                                                 `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateLinkServiceNetworkPolicies *string                                                 `json:"privateLinkServiceNetworkPolicies,omitempty"`
	RouteTable                        *SubResourceARM                                         `json:"routeTable,omitempty"`
	ServiceEndpointPolicies           []SubResourceARM                                        `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                  []ServiceEndpointPropertiesFormatARM                    `json:"serviceEndpoints,omitempty"`
}

// Deprecated version of ServiceEndpointPropertiesFormat. Use v1beta20201101.ServiceEndpointPropertiesFormat instead
type ServiceEndpointPropertiesFormatARM struct {
	Locations []string `json:"locations,omitempty"`
	Service   *string  `json:"service,omitempty"`
}

// Deprecated version of VirtualNetworksSubnets_Spec_Properties_Delegations. Use v1beta20201101.VirtualNetworksSubnets_Spec_Properties_Delegations instead
type VirtualNetworksSubnets_Spec_Properties_DelegationsARM struct {
	Name       *string                               `json:"name,omitempty"`
	Properties *ServiceDelegationPropertiesFormatARM `json:"properties,omitempty"`
}

// Deprecated version of ServiceDelegationPropertiesFormat. Use v1beta20201101.ServiceDelegationPropertiesFormat instead
type ServiceDelegationPropertiesFormatARM struct {
	ServiceName *string `json:"serviceName,omitempty"`
}
