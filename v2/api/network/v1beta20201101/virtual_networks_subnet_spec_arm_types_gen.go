// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of VirtualNetworks_Subnet_Spec. Use v1api20201101.VirtualNetworks_Subnet_Spec instead
type VirtualNetworks_Subnet_Spec_ARM struct {
	Name       string                                                                 `json:"name,omitempty"`
	Properties *SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworks_Subnet_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (subnet VirtualNetworks_Subnet_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (subnet *VirtualNetworks_Subnet_Spec_ARM) GetName() string {
	return subnet.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/subnets"
func (subnet *VirtualNetworks_Subnet_Spec_ARM) GetType() string {
	return "Microsoft.Network/virtualNetworks/subnets"
}

// Deprecated version of SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1api20201101.SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded instead
type SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	AddressPrefix                      *string                                                                            `json:"addressPrefix,omitempty"`
	AddressPrefixes                    []string                                                                           `json:"addressPrefixes,omitempty"`
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded_ARM `json:"applicationGatewayIpConfigurations,omitempty"`
	Delegations                        []Delegation_ARM                                                                   `json:"delegations,omitempty"`
	IpAllocations                      []SubResource_ARM                                                                  `json:"ipAllocations,omitempty"`
	NatGateway                         *SubResource_ARM                                                                   `json:"natGateway,omitempty"`
	NetworkSecurityGroup               *NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded_ARM           `json:"networkSecurityGroup,omitempty"`
	PrivateEndpointNetworkPolicies     *SubnetPropertiesFormat_PrivateEndpointNetworkPolicies                             `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateLinkServiceNetworkPolicies  *SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies                          `json:"privateLinkServiceNetworkPolicies,omitempty"`
	RouteTable                         *RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded_ARM                     `json:"routeTable,omitempty"`
	ServiceEndpointPolicies            []ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded_ARM         `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                   []ServiceEndpointPropertiesFormat_ARM                                              `json:"serviceEndpoints,omitempty"`
}

// Deprecated version of ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1api20201101.ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded instead
type ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of Delegation. Use v1api20201101.Delegation instead
type Delegation_ARM struct {
	Name       *string                                `json:"name,omitempty"`
	Properties *ServiceDelegationPropertiesFormat_ARM `json:"properties,omitempty"`
}

// Deprecated version of NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1api20201101.NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded instead
type NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1api20201101.RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded instead
type RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1api20201101.ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded instead
type ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of ServiceEndpointPropertiesFormat. Use v1api20201101.ServiceEndpointPropertiesFormat instead
type ServiceEndpointPropertiesFormat_ARM struct {
	Locations []string `json:"locations,omitempty"`
	Service   *string  `json:"service,omitempty"`
}

// Deprecated version of ServiceDelegationPropertiesFormat. Use v1api20201101.ServiceDelegationPropertiesFormat instead
type ServiceDelegationPropertiesFormat_ARM struct {
	ServiceName *string `json:"serviceName,omitempty"`
}
