// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type LoadBalancer_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// ExtendedLocation: The extended location of the load balancer.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`
	Id               *string              `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of load balancer.
	Properties *LoadBalancerPropertiesFormatARM `json:"properties,omitempty"`

	// Sku: The load balancer SKU.
	Sku *LoadBalancerSkuARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &LoadBalancer_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (balancer LoadBalancer_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (balancer *LoadBalancer_SpecARM) GetName() string {
	return balancer.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/loadBalancers"
func (balancer *LoadBalancer_SpecARM) GetType() string {
	return "Microsoft.Network/loadBalancers"
}

type ExtendedLocationARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType `json:"type,omitempty"`
}

type LoadBalancerPropertiesFormatARM struct {
	// BackendAddressPools: Collection of backend address pools used by a load balancer.
	BackendAddressPools []BackendAddressPool_LoadBalancer_SubResourceEmbeddedARM `json:"backendAddressPools,omitempty"`

	// FrontendIPConfigurations: Object representing the frontend IPs to be used for the load balancer.
	FrontendIPConfigurations []FrontendIPConfiguration_LoadBalancer_SubResourceEmbeddedARM `json:"frontendIPConfigurations,omitempty"`

	// InboundNatPools: Defines an external port range for inbound NAT to a single backend port on NICs associated with a load
	// balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external
	// port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat
	// rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual
	// virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools []InboundNatPoolARM `json:"inboundNatPools,omitempty"`

	// InboundNatRules: Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load
	// balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine
	// scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to
	// reference individual inbound NAT rules.
	InboundNatRules []InboundNatRule_LoadBalancer_SubResourceEmbeddedARM `json:"inboundNatRules,omitempty"`

	// LoadBalancingRules: Object collection representing the load balancing rules Gets the provisioning.
	LoadBalancingRules []LoadBalancingRuleARM `json:"loadBalancingRules,omitempty"`

	// OutboundRules: The outbound rules.
	OutboundRules []OutboundRuleARM `json:"outboundRules,omitempty"`

	// Probes: Collection of probe objects used in the load balancer.
	Probes []ProbeARM `json:"probes,omitempty"`
}

type LoadBalancerSkuARM struct {
	// Name: Name of a load balancer SKU.
	Name *LoadBalancerSku_Name `json:"name,omitempty"`

	// Tier: Tier of a load balancer SKU.
	Tier *LoadBalancerSku_Tier `json:"tier,omitempty"`
}

type BackendAddressPool_LoadBalancer_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationType_EdgeZone = ExtendedLocationType("EdgeZone")

type FrontendIPConfiguration_LoadBalancer_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within the set of frontend IP configurations used by the load balancer.
	// This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the load balancer probe.
	Properties *FrontendIPConfigurationPropertiesFormat_LoadBalancer_SubResourceEmbeddedARM `json:"properties,omitempty"`

	// Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type InboundNatPoolARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within the set of inbound NAT pools used by the load balancer. This name
	// can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of load balancer inbound nat pool.
	Properties *InboundNatPoolPropertiesFormatARM `json:"properties,omitempty"`
}

type InboundNatRule_LoadBalancer_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","Standard"}
type LoadBalancerSku_Name string

const (
	LoadBalancerSku_Name_Basic    = LoadBalancerSku_Name("Basic")
	LoadBalancerSku_Name_Standard = LoadBalancerSku_Name("Standard")
)

// +kubebuilder:validation:Enum={"Global","Regional"}
type LoadBalancerSku_Tier string

const (
	LoadBalancerSku_Tier_Global   = LoadBalancerSku_Tier("Global")
	LoadBalancerSku_Tier_Regional = LoadBalancerSku_Tier("Regional")
)

type LoadBalancingRuleARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within the set of load balancing rules used by the load balancer. This
	// name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of load balancer load balancing rule.
	Properties *LoadBalancingRulePropertiesFormatARM `json:"properties,omitempty"`
}

type OutboundRuleARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within the set of outbound rules used by the load balancer. This name can
	// be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of load balancer outbound rule.
	Properties *OutboundRulePropertiesFormatARM `json:"properties,omitempty"`
}

type ProbeARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within the set of probes used by the load balancer. This name can be used
	// to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of load balancer probe.
	Properties *ProbePropertiesFormatARM `json:"properties,omitempty"`
}

type FrontendIPConfigurationPropertiesFormat_LoadBalancer_SubResourceEmbeddedARM struct {
	// PrivateIPAddress: The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	// PrivateIPAddressVersion: Whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.
	PrivateIPAddressVersion *IPVersion `json:"privateIPAddressVersion,omitempty"`

	// PrivateIPAllocationMethod: The Private IP allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`

	// PublicIPAddress: The reference to the Public IP resource.
	PublicIPAddress *PublicIPAddressSpecARM `json:"publicIPAddress,omitempty"`

	// PublicIPPrefix: The reference to the Public IP Prefix resource.
	PublicIPPrefix *SubResourceARM `json:"publicIPPrefix,omitempty"`

	// Subnet: The reference to the subnet resource.
	Subnet *Subnet_LoadBalancer_SubResourceEmbeddedARM `json:"subnet,omitempty"`
}

type InboundNatPoolPropertiesFormatARM struct {
	// BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.
	BackendPort *int `json:"backendPort,omitempty"`

	// EnableFloatingIP: Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL
	// AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
	// This setting can't be changed after you create the endpoint.
	EnableFloatingIP *bool `json:"enableFloatingIP,omitempty"`

	// EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
	// element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `json:"enableTcpReset,omitempty"`

	// FrontendIPConfiguration: A reference to frontend IP addresses.
	FrontendIPConfiguration *SubResourceARM `json:"frontendIPConfiguration,omitempty"`

	// FrontendPortRangeEnd: The last port number in the range of external ports that will be used to provide Inbound Nat to
	// NICs associated with a load balancer. Acceptable values range between 1 and 65535.
	FrontendPortRangeEnd *int `json:"frontendPortRangeEnd,omitempty"`

	// FrontendPortRangeStart: The first port number in the range of external ports that will be used to provide Inbound Nat to
	// NICs associated with a load balancer. Acceptable values range between 1 and 65534.
	FrontendPortRangeStart *int `json:"frontendPortRangeStart,omitempty"`

	// IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
	// default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// Protocol: The reference to the transport protocol used by the inbound NAT pool.
	Protocol *TransportProtocol `json:"protocol,omitempty"`
}

type LoadBalancingRulePropertiesFormatARM struct {
	// BackendAddressPool: A reference to a pool of DIPs. Inbound traffic is randomly load balanced across IPs in the backend
	// IPs.
	BackendAddressPool *SubResourceARM `json:"backendAddressPool,omitempty"`

	// BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 0 and 65535. Note
	// that value 0 enables "Any Port".
	BackendPort *int `json:"backendPort,omitempty"`

	// DisableOutboundSnat: Configures SNAT for the VMs in the backend pool to use the publicIP address specified in the
	// frontend of the load balancing rule.
	DisableOutboundSnat *bool `json:"disableOutboundSnat,omitempty"`

	// EnableFloatingIP: Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL
	// AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
	// This setting can't be changed after you create the endpoint.
	EnableFloatingIP *bool `json:"enableFloatingIP,omitempty"`

	// EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
	// element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `json:"enableTcpReset,omitempty"`

	// FrontendIPConfiguration: A reference to frontend IP addresses.
	FrontendIPConfiguration *SubResourceARM `json:"frontendIPConfiguration,omitempty"`

	// FrontendPort: The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer.
	// Acceptable values are between 0 and 65534. Note that value 0 enables "Any Port".
	FrontendPort *int `json:"frontendPort,omitempty"`

	// IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
	// default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// LoadDistribution: The load distribution policy for this rule.
	LoadDistribution *LoadBalancingRulePropertiesFormat_LoadDistribution `json:"loadDistribution,omitempty"`

	// Probe: The reference to the load balancer probe used by the load balancing rule.
	Probe *SubResourceARM `json:"probe,omitempty"`

	// Protocol: The reference to the transport protocol used by the load balancing rule.
	Protocol *TransportProtocol `json:"protocol,omitempty"`
}

type OutboundRulePropertiesFormatARM struct {
	// AllocatedOutboundPorts: The number of outbound ports to be used for NAT.
	AllocatedOutboundPorts *int `json:"allocatedOutboundPorts,omitempty"`

	// BackendAddressPool: A reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend
	// IPs.
	BackendAddressPool *SubResourceARM `json:"backendAddressPool,omitempty"`

	// EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
	// element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `json:"enableTcpReset,omitempty"`

	// FrontendIPConfigurations: The Frontend IP addresses of the load balancer.
	FrontendIPConfigurations []SubResourceARM `json:"frontendIPConfigurations,omitempty"`

	// IdleTimeoutInMinutes: The timeout for the TCP idle connection.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// Protocol: The protocol for the outbound rule in load balancer.
	Protocol *OutboundRulePropertiesFormat_Protocol `json:"protocol,omitempty"`
}

type ProbePropertiesFormatARM struct {
	// IntervalInSeconds: The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the
	// interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking
	// the instance out of rotation. The default value is 15, the minimum value is 5.
	IntervalInSeconds *int `json:"intervalInSeconds,omitempty"`

	// NumberOfProbes: The number of probes where if no response, will result in stopping further traffic from being delivered
	// to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used
	// in Azure.
	NumberOfProbes *int `json:"numberOfProbes,omitempty"`

	// Port: The port for communicating the probe. Possible values range from 1 to 65535, inclusive.
	Port *int `json:"port,omitempty"`

	// Protocol: The protocol of the end point. If 'Tcp' is specified, a received ACK is required for the probe to be
	// successful. If 'Http' or 'Https' is specified, a 200 OK response from the specifies URI is required for the probe to be
	// successful.
	Protocol *ProbePropertiesFormat_Protocol `json:"protocol,omitempty"`

	// RequestPath: The URI used for requesting health status from the VM. Path is required if a protocol is set to http.
	// Otherwise, it is not allowed. There is no default value.
	RequestPath *string `json:"requestPath,omitempty"`
}

type PublicIPAddressSpecARM struct {
	// ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`
	Id               *string              `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Properties: Public IP address properties.
	Properties *PublicIPAddressPropertiesFormatARM `json:"properties,omitempty"`

	// Sku: The public IP address SKU.
	Sku *PublicIPAddressSkuARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type Subnet_LoadBalancer_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

type PublicIPAddressPropertiesFormatARM struct {
	// DdosSettings: The DDoS protection custom policy associated with the public IP address.
	DdosSettings *DdosSettingsARM `json:"ddosSettings,omitempty"`

	// DnsSettings: The FQDN of the DNS record associated with the public IP address.
	DnsSettings *PublicIPAddressDnsSettingsARM `json:"dnsSettings,omitempty"`

	// IdleTimeoutInMinutes: The idle timeout of the public IP address.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// IpAddress: The IP address associated with the public IP address resource.
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpTags: The list of tags associated with the public IP address.
	IpTags []IpTagARM `json:"ipTags,omitempty"`

	// LinkedPublicIPAddress: The linked public IP address of the public IP address resource.
	LinkedPublicIPAddress *PublicIPAddressSpecARM `json:"linkedPublicIPAddress,omitempty"`

	// MigrationPhase: Migration phase of Public IP Address.
	MigrationPhase *PublicIPAddressPropertiesFormat_MigrationPhase `json:"migrationPhase,omitempty"`

	// NatGateway: The NatGateway for the Public IP address.
	NatGateway *NatGatewaySpecARM `json:"natGateway,omitempty"`

	// PublicIPAddressVersion: The public IP address version.
	PublicIPAddressVersion *IPVersion `json:"publicIPAddressVersion,omitempty"`

	// PublicIPAllocationMethod: The public IP address allocation method.
	PublicIPAllocationMethod *IPAllocationMethod `json:"publicIPAllocationMethod,omitempty"`

	// PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated from.
	PublicIPPrefix *SubResourceARM `json:"publicIPPrefix,omitempty"`

	// ServicePublicIPAddress: The service public IP address of the public IP address resource.
	ServicePublicIPAddress *PublicIPAddressSpecARM `json:"servicePublicIPAddress,omitempty"`
}
