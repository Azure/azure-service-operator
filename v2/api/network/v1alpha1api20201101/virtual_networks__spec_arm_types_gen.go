// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of VirtualNetworks_Spec. Use v1beta20201101.VirtualNetworks_Spec instead
type VirtualNetworks_SpecARM struct {
	ExtendedLocation *ExtendedLocationARM                `json:"extendedLocation,omitempty"`
	Location         *string                             `json:"location,omitempty"`
	Name             string                              `json:"name,omitempty"`
	Properties       *VirtualNetworks_Spec_PropertiesARM `json:"properties,omitempty"`
	Tags             map[string]string                   `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworks_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (networks VirtualNetworks_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (networks *VirtualNetworks_SpecARM) GetName() string {
	return networks.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks"
func (networks *VirtualNetworks_SpecARM) GetType() string {
	return "Microsoft.Network/virtualNetworks"
}

// Deprecated version of VirtualNetworks_Spec_Properties. Use v1beta20201101.VirtualNetworks_Spec_Properties instead
type VirtualNetworks_Spec_PropertiesARM struct {
	AddressSpace           *AddressSpaceARM                                            `json:"addressSpace,omitempty"`
	BgpCommunities         *VirtualNetworkBgpCommunitiesARM                            `json:"bgpCommunities,omitempty"`
	DdosProtectionPlan     *SubResourceARM                                             `json:"ddosProtectionPlan,omitempty"`
	DhcpOptions            *DhcpOptionsARM                                             `json:"dhcpOptions,omitempty"`
	EnableDdosProtection   *bool                                                       `json:"enableDdosProtection,omitempty"`
	EnableVmProtection     *bool                                                       `json:"enableVmProtection,omitempty"`
	IpAllocations          []SubResourceARM                                            `json:"ipAllocations,omitempty"`
	Subnets                []VirtualNetworks_Spec_Properties_SubnetsARM                `json:"subnets,omitempty"`
	VirtualNetworkPeerings []VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM `json:"virtualNetworkPeerings,omitempty"`
}

// Deprecated version of DhcpOptions. Use v1beta20201101.DhcpOptions instead
type DhcpOptionsARM struct {
	DnsServers []string `json:"dnsServers,omitempty"`
}

// Deprecated version of VirtualNetworks_Spec_Properties_Subnets. Use v1beta20201101.VirtualNetworks_Spec_Properties_Subnets instead
type VirtualNetworks_Spec_Properties_SubnetsARM struct {
	Name       *string                                                `json:"name,omitempty"`
	Properties *VirtualNetworks_Spec_Properties_Subnets_PropertiesARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualNetworks_Spec_Properties_VirtualNetworkPeerings. Use v1beta20201101.VirtualNetworks_Spec_Properties_VirtualNetworkPeerings instead
type VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM struct {
	Name       *string                                   `json:"name,omitempty"`
	Properties *VirtualNetworkPeeringPropertiesFormatARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualNetworks_Spec_Properties_Subnets_Properties. Use v1beta20201101.VirtualNetworks_Spec_Properties_Subnets_Properties instead
type VirtualNetworks_Spec_Properties_Subnets_PropertiesARM struct {
	AddressPrefix                     *string                                                             `json:"addressPrefix,omitempty"`
	AddressPrefixes                   []string                                                            `json:"addressPrefixes,omitempty"`
	Delegations                       []VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM `json:"delegations,omitempty"`
	IpAllocations                     []SubResourceARM                                                    `json:"ipAllocations,omitempty"`
	NatGateway                        *SubResourceARM                                                     `json:"natGateway,omitempty"`
	NetworkSecurityGroup              *SubResourceARM                                                     `json:"networkSecurityGroup,omitempty"`
	PrivateEndpointNetworkPolicies    *string                                                             `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateLinkServiceNetworkPolicies *string                                                             `json:"privateLinkServiceNetworkPolicies,omitempty"`
	RouteTable                        *SubResourceARM                                                     `json:"routeTable,omitempty"`
	ServiceEndpointPolicies           []SubResourceARM                                                    `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                  []ServiceEndpointPropertiesFormatARM                                `json:"serviceEndpoints,omitempty"`
}

// Deprecated version of VirtualNetworks_Spec_Properties_Subnets_Properties_Delegations. Use v1beta20201101.VirtualNetworks_Spec_Properties_Subnets_Properties_Delegations instead
type VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM struct {
	Name       *string                               `json:"name,omitempty"`
	Properties *ServiceDelegationPropertiesFormatARM `json:"properties,omitempty"`
}
