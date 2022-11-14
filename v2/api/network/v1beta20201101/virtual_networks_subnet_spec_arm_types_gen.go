// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type VirtualNetworks_Subnet_Spec_ARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`

	// Properties: Properties of the subnet.
	Properties *SubnetPropertiesFormat_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
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

type SubnetPropertiesFormat_ARM struct {
	// AddressPrefix: The address prefix for the subnet.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// AddressPrefixes: List of address prefixes for the subnet.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`

	// ApplicationGatewayIpConfigurations: Application gateway IP configurations of virtual network resource.
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_ARM `json:"applicationGatewayIpConfigurations,omitempty"`

	// Delegations: An array of references to the delegations on the subnet.
	Delegations []Delegation_ARM `json:"delegations,omitempty"`

	// IpAllocations: Array of IpAllocation which reference this subnet.
	IpAllocations []SubResource_ARM `json:"ipAllocations,omitempty"`

	// NatGateway: Nat gateway associated with this subnet.
	NatGateway *SubResource_ARM `json:"natGateway,omitempty"`

	// NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupSpec_ARM `json:"networkSecurityGroup,omitempty"`

	// PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on private end point in the subnet.
	PrivateEndpointNetworkPolicies *SubnetPropertiesFormat_PrivateEndpointNetworkPolicies `json:"privateEndpointNetworkPolicies,omitempty"`

	// PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on private link service in the subnet.
	PrivateLinkServiceNetworkPolicies *SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies `json:"privateLinkServiceNetworkPolicies,omitempty"`

	// RouteTable: The reference to the RouteTable resource.
	RouteTable *RouteTableSpec_ARM `json:"routeTable,omitempty"`

	// ServiceEndpointPolicies: An array of service endpoint policies.
	ServiceEndpointPolicies []ServiceEndpointPolicySpec_ARM `json:"serviceEndpointPolicies,omitempty"`

	// ServiceEndpoints: An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormat_ARM `json:"serviceEndpoints,omitempty"`
}

type ApplicationGatewayIPConfiguration_ARM struct {
	Id *string `json:"id,omitempty"`

	// Name: Name of the IP configuration that is unique within an Application Gateway.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the application gateway IP configuration.
	Properties *ApplicationGatewayIPConfigurationPropertiesFormat_ARM `json:"properties,omitempty"`
}

type Delegation_ARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a subnet. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the subnet.
	Properties *ServiceDelegationPropertiesFormat_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type RouteTableSpec_ARM struct {
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Properties: Properties of the route table.
	Properties *RouteTablePropertiesFormat_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

type ServiceEndpointPolicySpec_ARM struct {
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Properties: Properties of the service end point policy.
	Properties *ServiceEndpointPolicyPropertiesFormat_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

type ServiceEndpointPropertiesFormat_ARM struct {
	// Locations: A list of locations.
	Locations []string `json:"locations,omitempty"`

	// Service: The type of the endpoint service.
	Service *string `json:"service,omitempty"`
}

type ApplicationGatewayIPConfigurationPropertiesFormat_ARM struct {
	// Subnet: Reference to the subnet resource. A subnet from where application gateway gets its private address.
	Subnet *SubResource_ARM `json:"subnet,omitempty"`
}

type ServiceDelegationPropertiesFormat_ARM struct {
	// ServiceName: The name of the service to whom the subnet should be delegated (e.g. Microsoft.Sql/servers).
	ServiceName *string `json:"serviceName,omitempty"`
}

type ServiceEndpointPolicyPropertiesFormat_ARM struct {
	// ServiceEndpointPolicyDefinitions: A collection of service endpoint policy definitions of the service endpoint policy.
	ServiceEndpointPolicyDefinitions []ServiceEndpointPolicyDefinition_ARM `json:"serviceEndpointPolicyDefinitions,omitempty"`
}

type ServiceEndpointPolicyDefinition_ARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the service endpoint policy definition.
	Properties *ServiceEndpointPolicyDefinitionPropertiesFormat_ARM `json:"properties,omitempty"`
}

type ServiceEndpointPolicyDefinitionPropertiesFormat_ARM struct {
	// Description: A description for this rule. Restricted to 140 chars.
	Description *string `json:"description,omitempty"`

	// Service: Service endpoint name.
	Service *string `json:"service,omitempty"`

	// ServiceResources: A list of service resources.
	ServiceResources []string `json:"serviceResources,omitempty"`
}
