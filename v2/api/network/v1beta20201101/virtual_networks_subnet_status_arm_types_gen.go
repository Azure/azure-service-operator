// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

<<<<<<<< HEAD:v2/api/network/v1beta20201101/virtual_networks_subnet_status_arm_types_gen.go
type VirtualNetworksSubnet_STATUSARM struct {
========
type Subnet_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM struct {
>>>>>>>> main:v2/api/network/v1beta20201101/subnet_status_virtual_networks_subnet_sub_resource_embedded_arm_types_gen.go
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the subnet.
	Properties *SubnetPropertiesFormat_STATUSARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type SubnetPropertiesFormat_STATUSARM struct {
	// AddressPrefix: The address prefix for the subnet.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// AddressPrefixes: List of address prefixes for the subnet.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`

	// ApplicationGatewayIpConfigurations: Application gateway IP configurations of virtual network resource.
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_STATUSARM `json:"applicationGatewayIpConfigurations,omitempty"`

	// Delegations: An array of references to the delegations on the subnet.
	Delegations []Delegation_STATUSARM `json:"delegations,omitempty"`

	// IpAllocations: Array of IpAllocation which reference this subnet.
	IpAllocations []SubResource_STATUSARM `json:"ipAllocations,omitempty"`

	// IpConfigurationProfiles: Array of IP configuration profiles which reference this subnet.
	IpConfigurationProfiles []IPConfigurationProfile_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM `json:"ipConfigurationProfiles,omitempty"`

	// IpConfigurations: An array of references to the network interface IP configurations using subnet.
	IpConfigurations []IPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM `json:"ipConfigurations,omitempty"`

	// NatGateway: Nat gateway associated with this subnet.
	NatGateway *SubResource_STATUSARM `json:"natGateway,omitempty"`

	// NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroup_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM `json:"networkSecurityGroup,omitempty"`

	// PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on private end point in the subnet.
<<<<<<<< HEAD:v2/api/network/v1beta20201101/virtual_networks_subnet_status_arm_types_gen.go
	PrivateEndpointNetworkPolicies *SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_STATUS `json:"privateEndpointNetworkPolicies,omitempty"`
========
	PrivateEndpointNetworkPolicies *SubnetPropertiesFormat_STATUS_PrivateEndpointNetworkPolicies `json:"privateEndpointNetworkPolicies,omitempty"`
>>>>>>>> main:v2/api/network/v1beta20201101/subnet_status_virtual_networks_subnet_sub_resource_embedded_arm_types_gen.go

	// PrivateEndpoints: An array of references to private endpoints.
	PrivateEndpoints []PrivateEndpoint_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM `json:"privateEndpoints,omitempty"`

	// PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on private link service in the subnet.
<<<<<<<< HEAD:v2/api/network/v1beta20201101/virtual_networks_subnet_status_arm_types_gen.go
	PrivateLinkServiceNetworkPolicies *SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_STATUS `json:"privateLinkServiceNetworkPolicies,omitempty"`
========
	PrivateLinkServiceNetworkPolicies *SubnetPropertiesFormat_STATUS_PrivateLinkServiceNetworkPolicies `json:"privateLinkServiceNetworkPolicies,omitempty"`
>>>>>>>> main:v2/api/network/v1beta20201101/subnet_status_virtual_networks_subnet_sub_resource_embedded_arm_types_gen.go

	// ProvisioningState: The provisioning state of the subnet resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Purpose: A read-only string identifying the intention of use for this subnet based on delegations and other user-defined
	// properties.
	Purpose *string `json:"purpose,omitempty"`

	// ResourceNavigationLinks: An array of references to the external resources using subnet.
	ResourceNavigationLinks []ResourceNavigationLink_STATUSARM `json:"resourceNavigationLinks,omitempty"`

	// RouteTable: The reference to the RouteTable resource.
	RouteTable *RouteTable_STATUS_SubResourceEmbeddedARM `json:"routeTable,omitempty"`

	// ServiceAssociationLinks: An array of references to services injecting into this subnet.
	ServiceAssociationLinks []ServiceAssociationLink_STATUSARM `json:"serviceAssociationLinks,omitempty"`

	// ServiceEndpointPolicies: An array of service endpoint policies.
	ServiceEndpointPolicies []ServiceEndpointPolicy_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM `json:"serviceEndpointPolicies,omitempty"`

	// ServiceEndpoints: An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormat_STATUSARM `json:"serviceEndpoints,omitempty"`
}

type ApplicationGatewayIPConfiguration_STATUSARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Name of the IP configuration that is unique within an Application Gateway.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the application gateway IP configuration.
	Properties *ApplicationGatewayIPConfigurationPropertiesFormat_STATUSARM `json:"properties,omitempty"`

	// Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type Delegation_STATUSARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a subnet. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the subnet.
	Properties *ServiceDelegationPropertiesFormat_STATUSARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type IPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the IP configuration.
	Properties *IPConfigurationPropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM `json:"properties,omitempty"`
}

type IPConfigurationProfile_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the IP configuration profile.
	Properties *IPConfigurationProfilePropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM `json:"properties,omitempty"`

	// Type: Sub Resource type.
	Type *string `json:"type,omitempty"`
}

type NetworkSecurityGroup_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type PrivateEndpoint_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM struct {
	// ExtendedLocation: The extended location of the load balancer.
	ExtendedLocation *ExtendedLocation_STATUSARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type ResourceNavigationLink_STATUSARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource navigation link identifier.
	Id *string `json:"id,omitempty"`

	// Name: Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Resource navigation link properties format.
	Properties *ResourceNavigationLinkFormat_STATUSARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type RouteTable_STATUS_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type ServiceAssociationLink_STATUSARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Resource navigation link properties format.
	Properties *ServiceAssociationLinkPropertiesFormat_STATUSARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type ServiceEndpointPolicy_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Kind: Kind of service endpoint policy. This is metadata used for the Azure portal experience.
	Kind *string `json:"kind,omitempty"`
}

type ServiceEndpointPropertiesFormat_STATUSARM struct {
	// Locations: A list of locations.
	Locations []string `json:"locations,omitempty"`

	// ProvisioningState: The provisioning state of the service endpoint resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Service: The type of the endpoint service.
	Service *string `json:"service,omitempty"`
}

<<<<<<<< HEAD:v2/api/network/v1beta20201101/virtual_networks_subnet_status_arm_types_gen.go
========
type SubnetPropertiesFormat_STATUS_PrivateEndpointNetworkPolicies string

const (
	SubnetPropertiesFormat_STATUS_PrivateEndpointNetworkPolicies_Disabled = SubnetPropertiesFormat_STATUS_PrivateEndpointNetworkPolicies("Disabled")
	SubnetPropertiesFormat_STATUS_PrivateEndpointNetworkPolicies_Enabled  = SubnetPropertiesFormat_STATUS_PrivateEndpointNetworkPolicies("Enabled")
)

type SubnetPropertiesFormat_STATUS_PrivateLinkServiceNetworkPolicies string

const (
	SubnetPropertiesFormat_STATUS_PrivateLinkServiceNetworkPolicies_Disabled = SubnetPropertiesFormat_STATUS_PrivateLinkServiceNetworkPolicies("Disabled")
	SubnetPropertiesFormat_STATUS_PrivateLinkServiceNetworkPolicies_Enabled  = SubnetPropertiesFormat_STATUS_PrivateLinkServiceNetworkPolicies("Enabled")
)

>>>>>>>> main:v2/api/network/v1beta20201101/subnet_status_virtual_networks_subnet_sub_resource_embedded_arm_types_gen.go
type ApplicationGatewayIPConfigurationPropertiesFormat_STATUSARM struct {
	// ProvisioningState: The provisioning state of the application gateway IP configuration resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Subnet: Reference to the subnet resource. A subnet from where application gateway gets its private address.
	Subnet *SubResource_STATUSARM `json:"subnet,omitempty"`
}

type IPConfigurationProfilePropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM struct {
	// ProvisioningState: The provisioning state of the IP configuration profile resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Subnet: The reference to the subnet resource to create a container network interface ip configuration.
	Subnet *Subnet_STATUS_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"subnet,omitempty"`
}

type IPConfigurationPropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM struct {
	// PrivateIPAddress: The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	// PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod_STATUS `json:"privateIPAllocationMethod,omitempty"`

	// ProvisioningState: The provisioning state of the IP configuration resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicIPAddress: The reference to the public IP resource.
<<<<<<<< HEAD:v2/api/network/v1beta20201101/virtual_networks_subnet_status_arm_types_gen.go
	PublicIPAddress *PublicIPAddress_STATUS_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"publicIPAddress,omitempty"`

	// Subnet: The reference to the subnet resource.
	Subnet *Subnet_STATUS_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"subnet,omitempty"`
========
	PublicIPAddress *PublicIPAddress_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM `json:"publicIPAddress,omitempty"`
>>>>>>>> main:v2/api/network/v1beta20201101/subnet_status_virtual_networks_subnet_sub_resource_embedded_arm_types_gen.go
}

type ResourceNavigationLinkFormat_STATUSARM struct {
	// Link: Link to the external resource.
	Link *string `json:"link,omitempty"`

	// LinkedResourceType: Resource type of the linked resource.
	LinkedResourceType *string `json:"linkedResourceType,omitempty"`

	// ProvisioningState: The provisioning state of the resource navigation link resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

type ServiceAssociationLinkPropertiesFormat_STATUSARM struct {
	// AllowDelete: If true, the resource can be deleted.
	AllowDelete *bool `json:"allowDelete,omitempty"`

	// Link: Link to the external resource.
	Link *string `json:"link,omitempty"`

	// LinkedResourceType: Resource type of the linked resource.
	LinkedResourceType *string `json:"linkedResourceType,omitempty"`

	// Locations: A list of locations.
	Locations []string `json:"locations,omitempty"`

	// ProvisioningState: The provisioning state of the service association link resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

type ServiceDelegationPropertiesFormat_STATUSARM struct {
	// Actions: The actions permitted to the service upon delegation.
	Actions []string `json:"actions,omitempty"`

	// ProvisioningState: The provisioning state of the service delegation resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ServiceName: The name of the service to whom the subnet should be delegated (e.g. Microsoft.Sql/servers).
	ServiceName *string `json:"serviceName,omitempty"`
}

type PublicIPAddress_STATUS_VirtualNetworks_Subnet_SubResourceEmbeddedARM struct {
	// ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocation_STATUSARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Sku: The public IP address SKU.
	Sku *PublicIPAddressSku_STATUSARM `json:"sku,omitempty"`

	// Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type Subnet_STATUS_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}
