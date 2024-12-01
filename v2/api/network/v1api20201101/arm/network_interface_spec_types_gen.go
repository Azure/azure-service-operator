// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NetworkInterface_Spec struct {
	// ExtendedLocation: The extended location of the network interface.
	ExtendedLocation *ExtendedLocation `json:"extendedLocation,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the network interface.
	Properties *NetworkInterfacePropertiesFormat `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NetworkInterface_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (networkInterface NetworkInterface_Spec) GetAPIVersion() string {
	return "2020-11-01"
}

// GetName returns the Name of the resource
func (networkInterface *NetworkInterface_Spec) GetName() string {
	return networkInterface.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkInterfaces"
func (networkInterface *NetworkInterface_Spec) GetType() string {
	return "Microsoft.Network/networkInterfaces"
}

// NetworkInterface properties.
type NetworkInterfacePropertiesFormat struct {
	// DnsSettings: The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettings `json:"dnsSettings,omitempty"`

	// EnableAcceleratedNetworking: If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty"`

	// EnableIPForwarding: Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding *bool `json:"enableIPForwarding,omitempty"`

	// IpConfigurations: A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterfaceIPConfiguration_NetworkInterface_SubResourceEmbedded `json:"ipConfigurations,omitempty"`

	// NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupSpec_NetworkInterface_SubResourceEmbedded `json:"networkSecurityGroup,omitempty"`

	// NicType: Type of Network Interface resource.
	NicType *NetworkInterfacePropertiesFormat_NicType `json:"nicType,omitempty"`

	// PrivateLinkService: Privatelinkservice of the network interface resource.
	PrivateLinkService *PrivateLinkServiceSpec `json:"privateLinkService,omitempty"`
}

// DNS settings of a network interface.
type NetworkInterfaceDnsSettings struct {
	// DnsServers: List of DNS servers IP addresses. Use 'AzureProvidedDNS' to switch to azure provided DNS resolution.
	// 'AzureProvidedDNS' value cannot be combined with other IPs, it must be the only value in dnsServers collection.
	DnsServers []string `json:"dnsServers,omitempty"`

	// InternalDnsNameLabel: Relative DNS name for this NIC used for internal communications between VMs in the same virtual
	// network.
	InternalDnsNameLabel *string `json:"internalDnsNameLabel,omitempty"`
}

// IPConfiguration in a network interface.
type NetworkInterfaceIPConfiguration_NetworkInterface_SubResourceEmbedded struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Network interface IP configuration properties.
	Properties *NetworkInterfaceIPConfigurationPropertiesFormat `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"Elastic","Standard"}
type NetworkInterfacePropertiesFormat_NicType string

const (
	NetworkInterfacePropertiesFormat_NicType_Elastic  = NetworkInterfacePropertiesFormat_NicType("Elastic")
	NetworkInterfacePropertiesFormat_NicType_Standard = NetworkInterfacePropertiesFormat_NicType("Standard")
)

// Mapping from string to NetworkInterfacePropertiesFormat_NicType
var networkInterfacePropertiesFormat_NicType_Values = map[string]NetworkInterfacePropertiesFormat_NicType{
	"elastic":  NetworkInterfacePropertiesFormat_NicType_Elastic,
	"standard": NetworkInterfacePropertiesFormat_NicType_Standard,
}

// NetworkSecurityGroup resource.
type NetworkSecurityGroupSpec_NetworkInterface_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// Private link service resource.
type PrivateLinkServiceSpec struct {
	Id *string `json:"id,omitempty"`
}

// Properties of IP configuration.
type NetworkInterfaceIPConfigurationPropertiesFormat struct {
	// ApplicationGatewayBackendAddressPools: The reference to ApplicationGatewayBackendAddressPool resource.
	ApplicationGatewayBackendAddressPools []ApplicationGatewayBackendAddressPool_NetworkInterface_SubResourceEmbedded `json:"applicationGatewayBackendAddressPools,omitempty"`

	// ApplicationSecurityGroups: Application security groups in which the IP configuration is included.
	ApplicationSecurityGroups []ApplicationSecurityGroupSpec_NetworkInterface_SubResourceEmbedded `json:"applicationSecurityGroups,omitempty"`

	// LoadBalancerBackendAddressPools: The reference to LoadBalancerBackendAddressPool resource.
	LoadBalancerBackendAddressPools []BackendAddressPool_NetworkInterface_SubResourceEmbedded `json:"loadBalancerBackendAddressPools,omitempty"`

	// LoadBalancerInboundNatRules: A list of references of LoadBalancerInboundNatRules.
	LoadBalancerInboundNatRules []InboundNatRule_NetworkInterface_SubResourceEmbedded `json:"loadBalancerInboundNatRules,omitempty"`

	// Primary: Whether this is a primary customer address on the network interface.
	Primary *bool `json:"primary,omitempty"`

	// PrivateIPAddress: Private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	// PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.
	PrivateIPAddressVersion *IPVersion `json:"privateIPAddressVersion,omitempty"`

	// PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`

	// PublicIPAddress: Public IP address bound to the IP configuration.
	PublicIPAddress *PublicIPAddressSpec_NetworkInterface_SubResourceEmbedded `json:"publicIPAddress,omitempty"`

	// Subnet: Subnet bound to the IP configuration.
	Subnet *Subnet_NetworkInterface_SubResourceEmbedded `json:"subnet,omitempty"`

	// VirtualNetworkTaps: The reference to Virtual Network Taps.
	VirtualNetworkTaps []VirtualNetworkTapSpec_NetworkInterface_SubResourceEmbedded `json:"virtualNetworkTaps,omitempty"`
}

// Backend Address Pool of an application gateway.
type ApplicationGatewayBackendAddressPool_NetworkInterface_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// An application security group in a resource group.
type ApplicationSecurityGroupSpec_NetworkInterface_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// Pool of backend IP addresses.
type BackendAddressPool_NetworkInterface_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// Inbound NAT rule of the load balancer.
type InboundNatRule_NetworkInterface_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// Public IP address resource.
type PublicIPAddressSpec_NetworkInterface_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// Subnet in a virtual network resource.
type Subnet_NetworkInterface_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// Virtual Network Tap resource.
type VirtualNetworkTapSpec_NetworkInterface_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}
