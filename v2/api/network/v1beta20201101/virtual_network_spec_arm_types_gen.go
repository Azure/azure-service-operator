// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type VirtualNetwork_Spec_ARM struct {
	// ExtendedLocation: The extended location of the virtual network.
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the virtual network.
	Properties *VirtualNetworkPropertiesFormat_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetwork_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (network VirtualNetwork_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (network *VirtualNetwork_Spec_ARM) GetName() string {
	return network.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks"
func (network *VirtualNetwork_Spec_ARM) GetType() string {
	return "Microsoft.Network/virtualNetworks"
}

// Properties of the virtual network.
type VirtualNetworkPropertiesFormat_ARM struct {
	// AddressSpace: The AddressSpace that contains an array of IP address ranges that can be used by subnets.
	AddressSpace *AddressSpace_ARM `json:"addressSpace,omitempty"`

	// BgpCommunities: Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
	BgpCommunities *VirtualNetworkBgpCommunities_ARM `json:"bgpCommunities,omitempty"`

	// DdosProtectionPlan: The DDoS protection plan associated with the virtual network.
	DdosProtectionPlan *SubResource_ARM `json:"ddosProtectionPlan,omitempty"`

	// DhcpOptions: The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
	DhcpOptions *DhcpOptions_ARM `json:"dhcpOptions,omitempty"`

	// EnableDdosProtection: Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It
	// requires a DDoS protection plan associated with the resource.
	EnableDdosProtection *bool `json:"enableDdosProtection,omitempty"`

	// EnableVmProtection: Indicates if VM protection is enabled for all the subnets in the virtual network.
	EnableVmProtection *bool `json:"enableVmProtection,omitempty"`

	// IpAllocations: Array of IpAllocation which reference this VNET.
	IpAllocations []SubResource_ARM `json:"ipAllocations,omitempty"`

	// Subnets: A list of subnets in a Virtual Network.
	Subnets []Subnet_VirtualNetwork_SubResourceEmbedded_ARM `json:"subnets,omitempty"`

	// VirtualNetworkPeerings: A list of peerings in a Virtual Network.
	VirtualNetworkPeerings []VirtualNetworkPeering_ARM `json:"virtualNetworkPeerings,omitempty"`
}

// DhcpOptions contains an array of DNS servers available to VMs deployed in the virtual network. Standard DHCP option for
// a subnet overrides VNET DHCP options.
type DhcpOptions_ARM struct {
	// DnsServers: The list of DNS servers IP addresses.
	DnsServers []string `json:"dnsServers,omitempty"`
}

// Subnet in a virtual network resource.
type Subnet_VirtualNetwork_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the subnet.
	Properties *SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
type VirtualNetworkBgpCommunities_ARM struct {
	// VirtualNetworkCommunity: The BGP community associated with the virtual network.
	VirtualNetworkCommunity *string `json:"virtualNetworkCommunity,omitempty"`
}

// Peerings in a virtual network resource.
type VirtualNetworkPeering_ARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the virtual network peering.
	Properties *VirtualNetworkPeeringPropertiesFormat_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of the subnet.
type SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM struct {
	// AddressPrefix: The address prefix for the subnet.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// AddressPrefixes: List of address prefixes for the subnet.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`

	// ApplicationGatewayIpConfigurations: Application gateway IP configurations of virtual network resource.
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM `json:"applicationGatewayIpConfigurations,omitempty"`

	// Delegations: An array of references to the delegations on the subnet.
	Delegations []Delegation_ARM `json:"delegations,omitempty"`

	// IpAllocations: Array of IpAllocation which reference this subnet.
	IpAllocations []SubResource_ARM `json:"ipAllocations,omitempty"`

	// NatGateway: Nat gateway associated with this subnet.
	NatGateway *SubResource_ARM `json:"natGateway,omitempty"`

	// NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM `json:"networkSecurityGroup,omitempty"`

	// PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on private end point in the subnet.
	PrivateEndpointNetworkPolicies *SubnetPropertiesFormat_PrivateEndpointNetworkPolicies `json:"privateEndpointNetworkPolicies,omitempty"`

	// PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on private link service in the subnet.
	PrivateLinkServiceNetworkPolicies *SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies `json:"privateLinkServiceNetworkPolicies,omitempty"`

	// RouteTable: The reference to the RouteTable resource.
	RouteTable *RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM `json:"routeTable,omitempty"`

	// ServiceEndpointPolicies: An array of service endpoint policies.
	ServiceEndpointPolicies []ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM `json:"serviceEndpointPolicies,omitempty"`

	// ServiceEndpoints: An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormat_ARM `json:"serviceEndpoints,omitempty"`
}

// IP configuration of an application gateway. Currently 1 public and 1 private IP configuration is allowed.
type ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// NetworkSecurityGroup resource.
type NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Route table resource.
type RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Service End point policy resource.
type ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}
