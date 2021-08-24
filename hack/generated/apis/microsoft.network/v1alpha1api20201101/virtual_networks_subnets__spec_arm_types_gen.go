// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"

type VirtualNetworksSubnets_SpecARM struct {
	//APIVersion: API Version of the resource type, optional when apiProfile is used
	//on the template
	APIVersion VirtualNetworksSubnetsSpecAPIVersion `json:"apiVersion"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: Name of the resource
	Name string `json:"name"`

	//Properties: Properties of the subnet.
	Properties VirtualNetworksSubnets_Spec_PropertiesARM `json:"properties"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type
	Type VirtualNetworksSubnetsSpecType `json:"type"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworksSubnets_SpecARM{}

// GetAPIVersion returns the APIVersion of the resource
func (virtualNetworksSubnetsSpecARM VirtualNetworksSubnets_SpecARM) GetAPIVersion() string {
	return string(virtualNetworksSubnetsSpecARM.APIVersion)
}

// GetName returns the Name of the resource
func (virtualNetworksSubnetsSpecARM VirtualNetworksSubnets_SpecARM) GetName() string {
	return virtualNetworksSubnetsSpecARM.Name
}

// GetType returns the Type of the resource
func (virtualNetworksSubnetsSpecARM VirtualNetworksSubnets_SpecARM) GetType() string {
	return string(virtualNetworksSubnetsSpecARM.Type)
}

// +kubebuilder:validation:Enum={"2020-11-01"}
type VirtualNetworksSubnetsSpecAPIVersion string

const VirtualNetworksSubnetsSpecAPIVersion20201101 = VirtualNetworksSubnetsSpecAPIVersion("2020-11-01")

// +kubebuilder:validation:Enum={"Microsoft.Network/virtualNetworks/subnets"}
type VirtualNetworksSubnetsSpecType string

const VirtualNetworksSubnetsSpecTypeMicrosoftNetworkVirtualNetworksSubnets = VirtualNetworksSubnetsSpecType("Microsoft.Network/virtualNetworks/subnets")

type VirtualNetworksSubnets_Spec_PropertiesARM struct {
	//AddressPrefix: The address prefix for the subnet.
	AddressPrefix string `json:"addressPrefix"`

	//AddressPrefixes: List of address prefixes for the subnet.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`

	//Delegations: An array of references to the delegations on the subnet.
	Delegations []VirtualNetworksSubnets_Spec_Properties_DelegationsARM `json:"delegations,omitempty"`

	//IpAllocations: Array of IpAllocation which reference this subnet.
	IpAllocations []SubResourceARM `json:"ipAllocations,omitempty"`

	//NatGateway: Nat gateway associated with this subnet.
	NatGateway *SubResourceARM `json:"natGateway,omitempty"`

	//NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *SubResourceARM `json:"networkSecurityGroup,omitempty"`

	//PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on
	//private end point in the subnet.
	PrivateEndpointNetworkPolicies *string `json:"privateEndpointNetworkPolicies,omitempty"`

	//PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on
	//private link service in the subnet.
	PrivateLinkServiceNetworkPolicies *string `json:"privateLinkServiceNetworkPolicies,omitempty"`

	//RouteTable: The reference to the RouteTable resource.
	RouteTable *SubResourceARM `json:"routeTable,omitempty"`

	//ServiceEndpointPolicies: An array of service endpoint policies.
	ServiceEndpointPolicies []SubResourceARM `json:"serviceEndpointPolicies,omitempty"`

	//ServiceEndpoints: An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormatARM `json:"serviceEndpoints,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ServiceEndpointPropertiesFormat
type ServiceEndpointPropertiesFormatARM struct {
	//Locations: A list of locations.
	Locations []string `json:"locations,omitempty"`

	//Service: The type of the endpoint service.
	Service *string `json:"service,omitempty"`
}

type VirtualNetworksSubnets_Spec_Properties_DelegationsARM struct {
	//Name: The name of the resource that is unique within a subnet. This name can be
	//used to access the resource.
	Name string `json:"name"`

	//Properties: Properties of the subnet.
	Properties *ServiceDelegationPropertiesFormatARM `json:"properties,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ServiceDelegationPropertiesFormat
type ServiceDelegationPropertiesFormatARM struct {
	//ServiceName: The name of the service to whom the subnet should be delegated
	//(e.g. Microsoft.Sql/servers).
	ServiceName *string `json:"serviceName,omitempty"`
}
