// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

// Deprecated version of Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded. Use v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded instead
type Subnet_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	Etag       *string                           `json:"etag,omitempty"`
	Id         *string                           `json:"id,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Properties *SubnetPropertiesFormat_StatusARM `json:"properties,omitempty"`
	Type       *string                           `json:"type,omitempty"`
}

// Deprecated version of SubnetPropertiesFormat_Status. Use v1beta20201101.SubnetPropertiesFormat_Status instead
type SubnetPropertiesFormat_StatusARM struct {
	AddressPrefix                      *string                                                                      `json:"addressPrefix,omitempty"`
	AddressPrefixes                    []string                                                                     `json:"addressPrefixes,omitempty"`
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_StatusARM                                `json:"applicationGatewayIpConfigurations,omitempty"`
	Delegations                        []Delegation_StatusARM                                                       `json:"delegations,omitempty"`
	IpAllocations                      []SubResource_StatusARM                                                      `json:"ipAllocations,omitempty"`
	IpConfigurationProfiles            []IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"ipConfigurationProfiles,omitempty"`
	IpConfigurations                   []IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM        `json:"ipConfigurations,omitempty"`
	NatGateway                         *SubResource_StatusARM                                                       `json:"natGateway,omitempty"`
	NetworkSecurityGroup               *NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM    `json:"networkSecurityGroup,omitempty"`
	PrivateEndpointNetworkPolicies     *string                                                                      `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateEndpoints                   []PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM        `json:"privateEndpoints,omitempty"`
	PrivateLinkServiceNetworkPolicies  *string                                                                      `json:"privateLinkServiceNetworkPolicies,omitempty"`
	ProvisioningState                  *string                                                                      `json:"provisioningState,omitempty"`
	Purpose                            *string                                                                      `json:"purpose,omitempty"`
	ResourceNavigationLinks            []ResourceNavigationLink_StatusARM                                           `json:"resourceNavigationLinks,omitempty"`
	RouteTable                         *RouteTable_Status_SubResourceEmbeddedARM                                    `json:"routeTable,omitempty"`
	ServiceAssociationLinks            []ServiceAssociationLink_StatusARM                                           `json:"serviceAssociationLinks,omitempty"`
	ServiceEndpointPolicies            []ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM  `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                   []ServiceEndpointPropertiesFormat_StatusARM                                  `json:"serviceEndpoints,omitempty"`
}

// Deprecated version of ApplicationGatewayIPConfiguration_Status. Use v1beta20201101.ApplicationGatewayIPConfiguration_Status instead
type ApplicationGatewayIPConfiguration_StatusARM struct {
	Etag       *string                                                      `json:"etag,omitempty"`
	Id         *string                                                      `json:"id,omitempty"`
	Name       *string                                                      `json:"name,omitempty"`
	Properties *ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM `json:"properties,omitempty"`
	Type       *string                                                      `json:"type,omitempty"`
}

// Deprecated version of Delegation_Status. Use v1beta20201101.Delegation_Status instead
type Delegation_StatusARM struct {
	Etag       *string                                      `json:"etag,omitempty"`
	Id         *string                                      `json:"id,omitempty"`
	Name       *string                                      `json:"name,omitempty"`
	Properties *ServiceDelegationPropertiesFormat_StatusARM `json:"properties,omitempty"`
	Type       *string                                      `json:"type,omitempty"`
}

// Deprecated version of IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded. Use v1beta20201101.IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded instead
type IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	Etag       *string                                                                                     `json:"etag,omitempty"`
	Id         *string                                                                                     `json:"id,omitempty"`
	Name       *string                                                                                     `json:"name,omitempty"`
	Properties *IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"properties,omitempty"`
	Type       *string                                                                                     `json:"type,omitempty"`
}

// Deprecated version of IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded. Use v1beta20201101.IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded instead
type IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	Etag       *string                                                                              `json:"etag,omitempty"`
	Id         *string                                                                              `json:"id,omitempty"`
	Name       *string                                                                              `json:"name,omitempty"`
	Properties *IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"properties,omitempty"`
}

// Deprecated version of NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbedded. Use v1beta20201101.NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbedded instead
type NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded. Use v1beta20201101.PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded instead
type PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`
	Id               *string                     `json:"id,omitempty"`
}

// Deprecated version of ResourceNavigationLink_Status. Use v1beta20201101.ResourceNavigationLink_Status instead
type ResourceNavigationLink_StatusARM struct {
	Etag       *string                                 `json:"etag,omitempty"`
	Id         *string                                 `json:"id,omitempty"`
	Name       *string                                 `json:"name,omitempty"`
	Properties *ResourceNavigationLinkFormat_StatusARM `json:"properties,omitempty"`
	Type       *string                                 `json:"type,omitempty"`
}

// Deprecated version of RouteTable_Status_SubResourceEmbedded. Use v1beta20201101.RouteTable_Status_SubResourceEmbedded instead
type RouteTable_Status_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of ServiceAssociationLink_Status. Use v1beta20201101.ServiceAssociationLink_Status instead
type ServiceAssociationLink_StatusARM struct {
	Etag       *string                                           `json:"etag,omitempty"`
	Id         *string                                           `json:"id,omitempty"`
	Name       *string                                           `json:"name,omitempty"`
	Properties *ServiceAssociationLinkPropertiesFormat_StatusARM `json:"properties,omitempty"`
	Type       *string                                           `json:"type,omitempty"`
}

// Deprecated version of ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbedded. Use v1beta20201101.ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbedded instead
type ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	Id   *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
}

// Deprecated version of ServiceEndpointPropertiesFormat_Status. Use v1beta20201101.ServiceEndpointPropertiesFormat_Status instead
type ServiceEndpointPropertiesFormat_StatusARM struct {
	Locations         []string `json:"locations,omitempty"`
	ProvisioningState *string  `json:"provisioningState,omitempty"`
	Service           *string  `json:"service,omitempty"`
}

// Deprecated version of ApplicationGatewayIPConfigurationPropertiesFormat_Status. Use v1beta20201101.ApplicationGatewayIPConfigurationPropertiesFormat_Status instead
type ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM struct {
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Subnet            *SubResource_StatusARM `json:"subnet,omitempty"`
}

// Deprecated version of IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbedded. Use v1beta20201101.IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbedded instead
type IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// Deprecated version of IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbedded. Use v1beta20201101.IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbedded instead
type IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	PrivateIPAddress          *string                                                              `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *string                                                              `json:"privateIPAllocationMethod,omitempty"`
	ProvisioningState         *string                                                              `json:"provisioningState,omitempty"`
	PublicIPAddress           *PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"publicIPAddress,omitempty"`
}

// Deprecated version of ResourceNavigationLinkFormat_Status. Use v1beta20201101.ResourceNavigationLinkFormat_Status instead
type ResourceNavigationLinkFormat_StatusARM struct {
	Link               *string `json:"link,omitempty"`
	LinkedResourceType *string `json:"linkedResourceType,omitempty"`
	ProvisioningState  *string `json:"provisioningState,omitempty"`
}

// Deprecated version of ServiceAssociationLinkPropertiesFormat_Status. Use v1beta20201101.ServiceAssociationLinkPropertiesFormat_Status instead
type ServiceAssociationLinkPropertiesFormat_StatusARM struct {
	AllowDelete        *bool    `json:"allowDelete,omitempty"`
	Link               *string  `json:"link,omitempty"`
	LinkedResourceType *string  `json:"linkedResourceType,omitempty"`
	Locations          []string `json:"locations,omitempty"`
	ProvisioningState  *string  `json:"provisioningState,omitempty"`
}

// Deprecated version of ServiceDelegationPropertiesFormat_Status. Use v1beta20201101.ServiceDelegationPropertiesFormat_Status instead
type ServiceDelegationPropertiesFormat_StatusARM struct {
	Actions           []string `json:"actions,omitempty"`
	ProvisioningState *string  `json:"provisioningState,omitempty"`
	ServiceName       *string  `json:"serviceName,omitempty"`
}

// Deprecated version of PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded. Use v1beta20201101.PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded instead
type PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	ExtendedLocation *ExtendedLocation_StatusARM   `json:"extendedLocation,omitempty"`
	Id               *string                       `json:"id,omitempty"`
	Sku              *PublicIPAddressSku_StatusARM `json:"sku,omitempty"`
	Zones            []string                      `json:"zones,omitempty"`
}
