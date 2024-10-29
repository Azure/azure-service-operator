// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type VirtualNetworkGateway_Spec struct {
	// ExtendedLocation: The extended location of type local virtual network gateway.
	ExtendedLocation *ExtendedLocation `json:"extendedLocation,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the virtual network gateway.
	Properties *VirtualNetworkGatewayPropertiesFormat `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworkGateway_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (gateway VirtualNetworkGateway_Spec) GetAPIVersion() string {
	return "2020-11-01"
}

// GetName returns the Name of the resource
func (gateway *VirtualNetworkGateway_Spec) GetName() string {
	return gateway.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworkGateways"
func (gateway *VirtualNetworkGateway_Spec) GetType() string {
	return "Microsoft.Network/virtualNetworkGateways"
}

// VirtualNetworkGateway properties.
type VirtualNetworkGatewayPropertiesFormat struct {
	// ActiveActive: ActiveActive flag.
	ActiveActive *bool `json:"activeActive,omitempty"`

	// BgpSettings: Virtual network gateway's BGP speaker settings.
	BgpSettings *BgpSettings `json:"bgpSettings,omitempty"`

	// CustomRoutes: The reference to the address space resource which represents the custom routes address space specified by
	// the customer for virtual network gateway and VpnClient.
	CustomRoutes *AddressSpace `json:"customRoutes,omitempty"`

	// EnableBgp: Whether BGP is enabled for this virtual network gateway or not.
	EnableBgp *bool `json:"enableBgp,omitempty"`

	// EnableDnsForwarding: Whether dns forwarding is enabled or not.
	EnableDnsForwarding *bool `json:"enableDnsForwarding,omitempty"`

	// EnablePrivateIpAddress: Whether private IP needs to be enabled on this gateway for connections or not.
	EnablePrivateIpAddress *bool `json:"enablePrivateIpAddress,omitempty"`

	// GatewayDefaultSite: The reference to the LocalNetworkGateway resource which represents local network site having default
	// routes. Assign Null value in case of removing existing default site setting.
	GatewayDefaultSite *SubResource `json:"gatewayDefaultSite,omitempty"`

	// GatewayType: The type of this virtual network gateway.
	GatewayType *VirtualNetworkGatewayPropertiesFormat_GatewayType `json:"gatewayType,omitempty"`

	// IpConfigurations: IP configurations for virtual network gateway.
	IpConfigurations []VirtualNetworkGatewayIPConfiguration `json:"ipConfigurations,omitempty"`

	// Sku: The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network
	// gateway.
	Sku                            *VirtualNetworkGatewaySku `json:"sku,omitempty"`
	VNetExtendedLocationResourceId *string                   `json:"vNetExtendedLocationResourceId,omitempty"`

	// VpnClientConfiguration: The reference to the VpnClientConfiguration resource which represents the P2S VpnClient
	// configurations.
	VpnClientConfiguration *VpnClientConfiguration `json:"vpnClientConfiguration,omitempty"`

	// VpnGatewayGeneration: The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
	VpnGatewayGeneration *VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration `json:"vpnGatewayGeneration,omitempty"`

	// VpnType: The type of this virtual network gateway.
	VpnType *VirtualNetworkGatewayPropertiesFormat_VpnType `json:"vpnType,omitempty"`
}

// AddressSpace contains an array of IP address ranges that can be used by subnets of the virtual network.
type AddressSpace struct {
	// AddressPrefixes: A list of address blocks reserved for this virtual network in CIDR notation.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`
}

// BGP settings details.
type BgpSettings struct {
	// Asn: The BGP speaker's ASN.
	Asn *uint32 `json:"asn,omitempty"`

	// BgpPeeringAddress: The BGP peering address and BGP identifier of this BGP speaker.
	BgpPeeringAddress *string `json:"bgpPeeringAddress,omitempty"`

	// BgpPeeringAddresses: BGP peering address with IP configuration ID for virtual network gateway.
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddress `json:"bgpPeeringAddresses,omitempty"`

	// PeerWeight: The weight added to routes learned from this BGP speaker.
	PeerWeight *int `json:"peerWeight,omitempty"`
}

// IP configuration for virtual network gateway.
type VirtualNetworkGatewayIPConfiguration struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the virtual network gateway ip configuration.
	Properties *VirtualNetworkGatewayIPConfigurationPropertiesFormat `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"ExpressRoute","LocalGateway","Vpn"}
type VirtualNetworkGatewayPropertiesFormat_GatewayType string

const (
	VirtualNetworkGatewayPropertiesFormat_GatewayType_ExpressRoute = VirtualNetworkGatewayPropertiesFormat_GatewayType("ExpressRoute")
	VirtualNetworkGatewayPropertiesFormat_GatewayType_LocalGateway = VirtualNetworkGatewayPropertiesFormat_GatewayType("LocalGateway")
	VirtualNetworkGatewayPropertiesFormat_GatewayType_Vpn          = VirtualNetworkGatewayPropertiesFormat_GatewayType("Vpn")
)

// Mapping from string to VirtualNetworkGatewayPropertiesFormat_GatewayType
var virtualNetworkGatewayPropertiesFormat_GatewayType_Values = map[string]VirtualNetworkGatewayPropertiesFormat_GatewayType{
	"expressroute": VirtualNetworkGatewayPropertiesFormat_GatewayType_ExpressRoute,
	"localgateway": VirtualNetworkGatewayPropertiesFormat_GatewayType_LocalGateway,
	"vpn":          VirtualNetworkGatewayPropertiesFormat_GatewayType_Vpn,
}

// +kubebuilder:validation:Enum={"Generation1","Generation2","None"}
type VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration string

const (
	VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration_Generation1 = VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration("Generation1")
	VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration_Generation2 = VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration("Generation2")
	VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration_None        = VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration("None")
)

// Mapping from string to VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration
var virtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration_Values = map[string]VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration{
	"generation1": VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration_Generation1,
	"generation2": VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration_Generation2,
	"none":        VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration_None,
}

// +kubebuilder:validation:Enum={"PolicyBased","RouteBased"}
type VirtualNetworkGatewayPropertiesFormat_VpnType string

const (
	VirtualNetworkGatewayPropertiesFormat_VpnType_PolicyBased = VirtualNetworkGatewayPropertiesFormat_VpnType("PolicyBased")
	VirtualNetworkGatewayPropertiesFormat_VpnType_RouteBased  = VirtualNetworkGatewayPropertiesFormat_VpnType("RouteBased")
)

// Mapping from string to VirtualNetworkGatewayPropertiesFormat_VpnType
var virtualNetworkGatewayPropertiesFormat_VpnType_Values = map[string]VirtualNetworkGatewayPropertiesFormat_VpnType{
	"policybased": VirtualNetworkGatewayPropertiesFormat_VpnType_PolicyBased,
	"routebased":  VirtualNetworkGatewayPropertiesFormat_VpnType_RouteBased,
}

// VirtualNetworkGatewaySku details.
type VirtualNetworkGatewaySku struct {
	// Name: Gateway SKU name.
	Name *VirtualNetworkGatewaySku_Name `json:"name,omitempty"`

	// Tier: Gateway SKU tier.
	Tier *VirtualNetworkGatewaySku_Tier `json:"tier,omitempty"`
}

// VpnClientConfiguration for P2S client.
type VpnClientConfiguration struct {
	// AadAudience: The AADAudience property of the VirtualNetworkGateway resource for vpn client connection used for AAD
	// authentication.
	AadAudience *string `json:"aadAudience,omitempty"`

	// AadIssuer: The AADIssuer property of the VirtualNetworkGateway resource for vpn client connection used for AAD
	// authentication.
	AadIssuer *string `json:"aadIssuer,omitempty"`

	// AadTenant: The AADTenant property of the VirtualNetworkGateway resource for vpn client connection used for AAD
	// authentication.
	AadTenant *string `json:"aadTenant,omitempty"`

	// RadiusServerAddress: The radius server address property of the VirtualNetworkGateway resource for vpn client connection.
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty"`

	// RadiusServerSecret: The radius secret property of the VirtualNetworkGateway resource for vpn client connection.
	RadiusServerSecret *string `json:"radiusServerSecret,omitempty"`

	// RadiusServers: The radiusServers property for multiple radius server configuration.
	RadiusServers []RadiusServer `json:"radiusServers,omitempty"`

	// VpnAuthenticationTypes: VPN authentication types for the virtual network gateway..
	VpnAuthenticationTypes []VpnClientConfiguration_VpnAuthenticationTypes `json:"vpnAuthenticationTypes,omitempty"`

	// VpnClientAddressPool: The reference to the address space resource which represents Address space for P2S VpnClient.
	VpnClientAddressPool *AddressSpace `json:"vpnClientAddressPool,omitempty"`

	// VpnClientIpsecPolicies: VpnClientIpsecPolicies for virtual network gateway P2S client.
	VpnClientIpsecPolicies []IpsecPolicy `json:"vpnClientIpsecPolicies,omitempty"`

	// VpnClientProtocols: VpnClientProtocols for Virtual network gateway.
	VpnClientProtocols []VpnClientConfiguration_VpnClientProtocols `json:"vpnClientProtocols,omitempty"`

	// VpnClientRevokedCertificates: VpnClientRevokedCertificate for Virtual network gateway.
	VpnClientRevokedCertificates []VpnClientRevokedCertificate `json:"vpnClientRevokedCertificates,omitempty"`

	// VpnClientRootCertificates: VpnClientRootCertificate for virtual network gateway.
	VpnClientRootCertificates []VpnClientRootCertificate `json:"vpnClientRootCertificates,omitempty"`
}

// Properties of IPConfigurationBgpPeeringAddress.
type IPConfigurationBgpPeeringAddress struct {
	// CustomBgpIpAddresses: The list of custom BGP peering addresses which belong to IP configuration.
	CustomBgpIpAddresses []string `json:"customBgpIpAddresses,omitempty"`

	// IpconfigurationId: The ID of IP configuration which belongs to gateway.
	IpconfigurationId *string `json:"ipconfigurationId,omitempty"`
}

// An IPSec Policy configuration for a virtual network gateway connection.
type IpsecPolicy struct {
	// DhGroup: The DH Group used in IKE Phase 1 for initial SA.
	DhGroup *DhGroup `json:"dhGroup,omitempty"`

	// IkeEncryption: The IKE encryption algorithm (IKE phase 2).
	IkeEncryption *IkeEncryption `json:"ikeEncryption,omitempty"`

	// IkeIntegrity: The IKE integrity algorithm (IKE phase 2).
	IkeIntegrity *IkeIntegrity `json:"ikeIntegrity,omitempty"`

	// IpsecEncryption: The IPSec encryption algorithm (IKE phase 1).
	IpsecEncryption *IpsecEncryption `json:"ipsecEncryption,omitempty"`

	// IpsecIntegrity: The IPSec integrity algorithm (IKE phase 1).
	IpsecIntegrity *IpsecIntegrity `json:"ipsecIntegrity,omitempty"`

	// PfsGroup: The Pfs Group used in IKE Phase 2 for new child SA.
	PfsGroup *PfsGroup `json:"pfsGroup,omitempty"`

	// SaDataSizeKilobytes: The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site
	// to site VPN tunnel.
	SaDataSizeKilobytes *int `json:"saDataSizeKilobytes,omitempty"`

	// SaLifeTimeSeconds: The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site
	// to site VPN tunnel.
	SaLifeTimeSeconds *int `json:"saLifeTimeSeconds,omitempty"`
}

// Radius Server Settings.
type RadiusServer struct {
	// RadiusServerAddress: The address of this radius server.
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty"`

	// RadiusServerScore: The initial score assigned to this radius server.
	RadiusServerScore *int `json:"radiusServerScore,omitempty"`

	// RadiusServerSecret: The secret used for this radius server.
	RadiusServerSecret *string `json:"radiusServerSecret,omitempty"`
}

// Properties of VirtualNetworkGatewayIPConfiguration.
type VirtualNetworkGatewayIPConfigurationPropertiesFormat struct {
	// PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`

	// PublicIPAddress: The reference to the public IP resource.
	PublicIPAddress *SubResource `json:"publicIPAddress,omitempty"`

	// Subnet: The reference to the subnet resource.
	Subnet *SubResource `json:"subnet,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","ErGw1AZ","ErGw2AZ","ErGw3AZ","HighPerformance","Standard","UltraPerformance","VpnGw1","VpnGw1AZ","VpnGw2","VpnGw2AZ","VpnGw3","VpnGw3AZ","VpnGw4","VpnGw4AZ","VpnGw5","VpnGw5AZ"}
type VirtualNetworkGatewaySku_Name string

const (
	VirtualNetworkGatewaySku_Name_Basic            = VirtualNetworkGatewaySku_Name("Basic")
	VirtualNetworkGatewaySku_Name_ErGw1AZ          = VirtualNetworkGatewaySku_Name("ErGw1AZ")
	VirtualNetworkGatewaySku_Name_ErGw2AZ          = VirtualNetworkGatewaySku_Name("ErGw2AZ")
	VirtualNetworkGatewaySku_Name_ErGw3AZ          = VirtualNetworkGatewaySku_Name("ErGw3AZ")
	VirtualNetworkGatewaySku_Name_HighPerformance  = VirtualNetworkGatewaySku_Name("HighPerformance")
	VirtualNetworkGatewaySku_Name_Standard         = VirtualNetworkGatewaySku_Name("Standard")
	VirtualNetworkGatewaySku_Name_UltraPerformance = VirtualNetworkGatewaySku_Name("UltraPerformance")
	VirtualNetworkGatewaySku_Name_VpnGw1           = VirtualNetworkGatewaySku_Name("VpnGw1")
	VirtualNetworkGatewaySku_Name_VpnGw1AZ         = VirtualNetworkGatewaySku_Name("VpnGw1AZ")
	VirtualNetworkGatewaySku_Name_VpnGw2           = VirtualNetworkGatewaySku_Name("VpnGw2")
	VirtualNetworkGatewaySku_Name_VpnGw2AZ         = VirtualNetworkGatewaySku_Name("VpnGw2AZ")
	VirtualNetworkGatewaySku_Name_VpnGw3           = VirtualNetworkGatewaySku_Name("VpnGw3")
	VirtualNetworkGatewaySku_Name_VpnGw3AZ         = VirtualNetworkGatewaySku_Name("VpnGw3AZ")
	VirtualNetworkGatewaySku_Name_VpnGw4           = VirtualNetworkGatewaySku_Name("VpnGw4")
	VirtualNetworkGatewaySku_Name_VpnGw4AZ         = VirtualNetworkGatewaySku_Name("VpnGw4AZ")
	VirtualNetworkGatewaySku_Name_VpnGw5           = VirtualNetworkGatewaySku_Name("VpnGw5")
	VirtualNetworkGatewaySku_Name_VpnGw5AZ         = VirtualNetworkGatewaySku_Name("VpnGw5AZ")
)

// Mapping from string to VirtualNetworkGatewaySku_Name
var virtualNetworkGatewaySku_Name_Values = map[string]VirtualNetworkGatewaySku_Name{
	"basic":            VirtualNetworkGatewaySku_Name_Basic,
	"ergw1az":          VirtualNetworkGatewaySku_Name_ErGw1AZ,
	"ergw2az":          VirtualNetworkGatewaySku_Name_ErGw2AZ,
	"ergw3az":          VirtualNetworkGatewaySku_Name_ErGw3AZ,
	"highperformance":  VirtualNetworkGatewaySku_Name_HighPerformance,
	"standard":         VirtualNetworkGatewaySku_Name_Standard,
	"ultraperformance": VirtualNetworkGatewaySku_Name_UltraPerformance,
	"vpngw1":           VirtualNetworkGatewaySku_Name_VpnGw1,
	"vpngw1az":         VirtualNetworkGatewaySku_Name_VpnGw1AZ,
	"vpngw2":           VirtualNetworkGatewaySku_Name_VpnGw2,
	"vpngw2az":         VirtualNetworkGatewaySku_Name_VpnGw2AZ,
	"vpngw3":           VirtualNetworkGatewaySku_Name_VpnGw3,
	"vpngw3az":         VirtualNetworkGatewaySku_Name_VpnGw3AZ,
	"vpngw4":           VirtualNetworkGatewaySku_Name_VpnGw4,
	"vpngw4az":         VirtualNetworkGatewaySku_Name_VpnGw4AZ,
	"vpngw5":           VirtualNetworkGatewaySku_Name_VpnGw5,
	"vpngw5az":         VirtualNetworkGatewaySku_Name_VpnGw5AZ,
}

// +kubebuilder:validation:Enum={"Basic","ErGw1AZ","ErGw2AZ","ErGw3AZ","HighPerformance","Standard","UltraPerformance","VpnGw1","VpnGw1AZ","VpnGw2","VpnGw2AZ","VpnGw3","VpnGw3AZ","VpnGw4","VpnGw4AZ","VpnGw5","VpnGw5AZ"}
type VirtualNetworkGatewaySku_Tier string

const (
	VirtualNetworkGatewaySku_Tier_Basic            = VirtualNetworkGatewaySku_Tier("Basic")
	VirtualNetworkGatewaySku_Tier_ErGw1AZ          = VirtualNetworkGatewaySku_Tier("ErGw1AZ")
	VirtualNetworkGatewaySku_Tier_ErGw2AZ          = VirtualNetworkGatewaySku_Tier("ErGw2AZ")
	VirtualNetworkGatewaySku_Tier_ErGw3AZ          = VirtualNetworkGatewaySku_Tier("ErGw3AZ")
	VirtualNetworkGatewaySku_Tier_HighPerformance  = VirtualNetworkGatewaySku_Tier("HighPerformance")
	VirtualNetworkGatewaySku_Tier_Standard         = VirtualNetworkGatewaySku_Tier("Standard")
	VirtualNetworkGatewaySku_Tier_UltraPerformance = VirtualNetworkGatewaySku_Tier("UltraPerformance")
	VirtualNetworkGatewaySku_Tier_VpnGw1           = VirtualNetworkGatewaySku_Tier("VpnGw1")
	VirtualNetworkGatewaySku_Tier_VpnGw1AZ         = VirtualNetworkGatewaySku_Tier("VpnGw1AZ")
	VirtualNetworkGatewaySku_Tier_VpnGw2           = VirtualNetworkGatewaySku_Tier("VpnGw2")
	VirtualNetworkGatewaySku_Tier_VpnGw2AZ         = VirtualNetworkGatewaySku_Tier("VpnGw2AZ")
	VirtualNetworkGatewaySku_Tier_VpnGw3           = VirtualNetworkGatewaySku_Tier("VpnGw3")
	VirtualNetworkGatewaySku_Tier_VpnGw3AZ         = VirtualNetworkGatewaySku_Tier("VpnGw3AZ")
	VirtualNetworkGatewaySku_Tier_VpnGw4           = VirtualNetworkGatewaySku_Tier("VpnGw4")
	VirtualNetworkGatewaySku_Tier_VpnGw4AZ         = VirtualNetworkGatewaySku_Tier("VpnGw4AZ")
	VirtualNetworkGatewaySku_Tier_VpnGw5           = VirtualNetworkGatewaySku_Tier("VpnGw5")
	VirtualNetworkGatewaySku_Tier_VpnGw5AZ         = VirtualNetworkGatewaySku_Tier("VpnGw5AZ")
)

// Mapping from string to VirtualNetworkGatewaySku_Tier
var virtualNetworkGatewaySku_Tier_Values = map[string]VirtualNetworkGatewaySku_Tier{
	"basic":            VirtualNetworkGatewaySku_Tier_Basic,
	"ergw1az":          VirtualNetworkGatewaySku_Tier_ErGw1AZ,
	"ergw2az":          VirtualNetworkGatewaySku_Tier_ErGw2AZ,
	"ergw3az":          VirtualNetworkGatewaySku_Tier_ErGw3AZ,
	"highperformance":  VirtualNetworkGatewaySku_Tier_HighPerformance,
	"standard":         VirtualNetworkGatewaySku_Tier_Standard,
	"ultraperformance": VirtualNetworkGatewaySku_Tier_UltraPerformance,
	"vpngw1":           VirtualNetworkGatewaySku_Tier_VpnGw1,
	"vpngw1az":         VirtualNetworkGatewaySku_Tier_VpnGw1AZ,
	"vpngw2":           VirtualNetworkGatewaySku_Tier_VpnGw2,
	"vpngw2az":         VirtualNetworkGatewaySku_Tier_VpnGw2AZ,
	"vpngw3":           VirtualNetworkGatewaySku_Tier_VpnGw3,
	"vpngw3az":         VirtualNetworkGatewaySku_Tier_VpnGw3AZ,
	"vpngw4":           VirtualNetworkGatewaySku_Tier_VpnGw4,
	"vpngw4az":         VirtualNetworkGatewaySku_Tier_VpnGw4AZ,
	"vpngw5":           VirtualNetworkGatewaySku_Tier_VpnGw5,
	"vpngw5az":         VirtualNetworkGatewaySku_Tier_VpnGw5AZ,
}

// +kubebuilder:validation:Enum={"AAD","Certificate","Radius"}
type VpnClientConfiguration_VpnAuthenticationTypes string

const (
	VpnClientConfiguration_VpnAuthenticationTypes_AAD         = VpnClientConfiguration_VpnAuthenticationTypes("AAD")
	VpnClientConfiguration_VpnAuthenticationTypes_Certificate = VpnClientConfiguration_VpnAuthenticationTypes("Certificate")
	VpnClientConfiguration_VpnAuthenticationTypes_Radius      = VpnClientConfiguration_VpnAuthenticationTypes("Radius")
)

// Mapping from string to VpnClientConfiguration_VpnAuthenticationTypes
var vpnClientConfiguration_VpnAuthenticationTypes_Values = map[string]VpnClientConfiguration_VpnAuthenticationTypes{
	"aad":         VpnClientConfiguration_VpnAuthenticationTypes_AAD,
	"certificate": VpnClientConfiguration_VpnAuthenticationTypes_Certificate,
	"radius":      VpnClientConfiguration_VpnAuthenticationTypes_Radius,
}

// +kubebuilder:validation:Enum={"IkeV2","OpenVPN","SSTP"}
type VpnClientConfiguration_VpnClientProtocols string

const (
	VpnClientConfiguration_VpnClientProtocols_IkeV2   = VpnClientConfiguration_VpnClientProtocols("IkeV2")
	VpnClientConfiguration_VpnClientProtocols_OpenVPN = VpnClientConfiguration_VpnClientProtocols("OpenVPN")
	VpnClientConfiguration_VpnClientProtocols_SSTP    = VpnClientConfiguration_VpnClientProtocols("SSTP")
)

// Mapping from string to VpnClientConfiguration_VpnClientProtocols
var vpnClientConfiguration_VpnClientProtocols_Values = map[string]VpnClientConfiguration_VpnClientProtocols{
	"ikev2":   VpnClientConfiguration_VpnClientProtocols_IkeV2,
	"openvpn": VpnClientConfiguration_VpnClientProtocols_OpenVPN,
	"sstp":    VpnClientConfiguration_VpnClientProtocols_SSTP,
}

// VPN client revoked certificate of virtual network gateway.
type VpnClientRevokedCertificate struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the vpn client revoked certificate.
	Properties *VpnClientRevokedCertificatePropertiesFormat `json:"properties,omitempty"`
}

// VPN client root certificate of virtual network gateway.
type VpnClientRootCertificate struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the vpn client root certificate.
	Properties *VpnClientRootCertificatePropertiesFormat `json:"properties,omitempty"`
}

// The DH Groups used in IKE Phase 1 for initial SA.
// +kubebuilder:validation:Enum={"DHGroup1","DHGroup14","DHGroup2","DHGroup2048","DHGroup24","ECP256","ECP384","None"}
type DhGroup string

const (
	DhGroup_DHGroup1    = DhGroup("DHGroup1")
	DhGroup_DHGroup14   = DhGroup("DHGroup14")
	DhGroup_DHGroup2    = DhGroup("DHGroup2")
	DhGroup_DHGroup2048 = DhGroup("DHGroup2048")
	DhGroup_DHGroup24   = DhGroup("DHGroup24")
	DhGroup_ECP256      = DhGroup("ECP256")
	DhGroup_ECP384      = DhGroup("ECP384")
	DhGroup_None        = DhGroup("None")
)

// Mapping from string to DhGroup
var dhGroup_Values = map[string]DhGroup{
	"dhgroup1":    DhGroup_DHGroup1,
	"dhgroup14":   DhGroup_DHGroup14,
	"dhgroup2":    DhGroup_DHGroup2,
	"dhgroup2048": DhGroup_DHGroup2048,
	"dhgroup24":   DhGroup_DHGroup24,
	"ecp256":      DhGroup_ECP256,
	"ecp384":      DhGroup_ECP384,
	"none":        DhGroup_None,
}

// The IKE encryption algorithm (IKE phase 2).
// +kubebuilder:validation:Enum={"AES128","AES192","AES256","DES","DES3","GCMAES128","GCMAES256"}
type IkeEncryption string

const (
	IkeEncryption_AES128    = IkeEncryption("AES128")
	IkeEncryption_AES192    = IkeEncryption("AES192")
	IkeEncryption_AES256    = IkeEncryption("AES256")
	IkeEncryption_DES       = IkeEncryption("DES")
	IkeEncryption_DES3      = IkeEncryption("DES3")
	IkeEncryption_GCMAES128 = IkeEncryption("GCMAES128")
	IkeEncryption_GCMAES256 = IkeEncryption("GCMAES256")
)

// Mapping from string to IkeEncryption
var ikeEncryption_Values = map[string]IkeEncryption{
	"aes128":    IkeEncryption_AES128,
	"aes192":    IkeEncryption_AES192,
	"aes256":    IkeEncryption_AES256,
	"des":       IkeEncryption_DES,
	"des3":      IkeEncryption_DES3,
	"gcmaes128": IkeEncryption_GCMAES128,
	"gcmaes256": IkeEncryption_GCMAES256,
}

// The IKE integrity algorithm (IKE phase 2).
// +kubebuilder:validation:Enum={"GCMAES128","GCMAES256","MD5","SHA1","SHA256","SHA384"}
type IkeIntegrity string

const (
	IkeIntegrity_GCMAES128 = IkeIntegrity("GCMAES128")
	IkeIntegrity_GCMAES256 = IkeIntegrity("GCMAES256")
	IkeIntegrity_MD5       = IkeIntegrity("MD5")
	IkeIntegrity_SHA1      = IkeIntegrity("SHA1")
	IkeIntegrity_SHA256    = IkeIntegrity("SHA256")
	IkeIntegrity_SHA384    = IkeIntegrity("SHA384")
)

// Mapping from string to IkeIntegrity
var ikeIntegrity_Values = map[string]IkeIntegrity{
	"gcmaes128": IkeIntegrity_GCMAES128,
	"gcmaes256": IkeIntegrity_GCMAES256,
	"md5":       IkeIntegrity_MD5,
	"sha1":      IkeIntegrity_SHA1,
	"sha256":    IkeIntegrity_SHA256,
	"sha384":    IkeIntegrity_SHA384,
}

// The IPSec encryption algorithm (IKE phase 1).
// +kubebuilder:validation:Enum={"AES128","AES192","AES256","DES","DES3","GCMAES128","GCMAES192","GCMAES256","None"}
type IpsecEncryption string

const (
	IpsecEncryption_AES128    = IpsecEncryption("AES128")
	IpsecEncryption_AES192    = IpsecEncryption("AES192")
	IpsecEncryption_AES256    = IpsecEncryption("AES256")
	IpsecEncryption_DES       = IpsecEncryption("DES")
	IpsecEncryption_DES3      = IpsecEncryption("DES3")
	IpsecEncryption_GCMAES128 = IpsecEncryption("GCMAES128")
	IpsecEncryption_GCMAES192 = IpsecEncryption("GCMAES192")
	IpsecEncryption_GCMAES256 = IpsecEncryption("GCMAES256")
	IpsecEncryption_None      = IpsecEncryption("None")
)

// Mapping from string to IpsecEncryption
var ipsecEncryption_Values = map[string]IpsecEncryption{
	"aes128":    IpsecEncryption_AES128,
	"aes192":    IpsecEncryption_AES192,
	"aes256":    IpsecEncryption_AES256,
	"des":       IpsecEncryption_DES,
	"des3":      IpsecEncryption_DES3,
	"gcmaes128": IpsecEncryption_GCMAES128,
	"gcmaes192": IpsecEncryption_GCMAES192,
	"gcmaes256": IpsecEncryption_GCMAES256,
	"none":      IpsecEncryption_None,
}

// The IPSec integrity algorithm (IKE phase 1).
// +kubebuilder:validation:Enum={"GCMAES128","GCMAES192","GCMAES256","MD5","SHA1","SHA256"}
type IpsecIntegrity string

const (
	IpsecIntegrity_GCMAES128 = IpsecIntegrity("GCMAES128")
	IpsecIntegrity_GCMAES192 = IpsecIntegrity("GCMAES192")
	IpsecIntegrity_GCMAES256 = IpsecIntegrity("GCMAES256")
	IpsecIntegrity_MD5       = IpsecIntegrity("MD5")
	IpsecIntegrity_SHA1      = IpsecIntegrity("SHA1")
	IpsecIntegrity_SHA256    = IpsecIntegrity("SHA256")
)

// Mapping from string to IpsecIntegrity
var ipsecIntegrity_Values = map[string]IpsecIntegrity{
	"gcmaes128": IpsecIntegrity_GCMAES128,
	"gcmaes192": IpsecIntegrity_GCMAES192,
	"gcmaes256": IpsecIntegrity_GCMAES256,
	"md5":       IpsecIntegrity_MD5,
	"sha1":      IpsecIntegrity_SHA1,
	"sha256":    IpsecIntegrity_SHA256,
}

// The Pfs Groups used in IKE Phase 2 for new child SA.
// +kubebuilder:validation:Enum={"ECP256","ECP384","None","PFS1","PFS14","PFS2","PFS2048","PFS24","PFSMM"}
type PfsGroup string

const (
	PfsGroup_ECP256  = PfsGroup("ECP256")
	PfsGroup_ECP384  = PfsGroup("ECP384")
	PfsGroup_None    = PfsGroup("None")
	PfsGroup_PFS1    = PfsGroup("PFS1")
	PfsGroup_PFS14   = PfsGroup("PFS14")
	PfsGroup_PFS2    = PfsGroup("PFS2")
	PfsGroup_PFS2048 = PfsGroup("PFS2048")
	PfsGroup_PFS24   = PfsGroup("PFS24")
	PfsGroup_PFSMM   = PfsGroup("PFSMM")
)

// Mapping from string to PfsGroup
var pfsGroup_Values = map[string]PfsGroup{
	"ecp256":  PfsGroup_ECP256,
	"ecp384":  PfsGroup_ECP384,
	"none":    PfsGroup_None,
	"pfs1":    PfsGroup_PFS1,
	"pfs14":   PfsGroup_PFS14,
	"pfs2":    PfsGroup_PFS2,
	"pfs2048": PfsGroup_PFS2048,
	"pfs24":   PfsGroup_PFS24,
	"pfsmm":   PfsGroup_PFSMM,
}

// Properties of the revoked VPN client certificate of virtual network gateway.
type VpnClientRevokedCertificatePropertiesFormat struct {
	// Thumbprint: The revoked VPN client certificate thumbprint.
	Thumbprint *string `json:"thumbprint,omitempty"`
}

// Properties of SSL certificates of application gateway.
type VpnClientRootCertificatePropertiesFormat struct {
	// PublicCertData: The certificate public data.
	PublicCertData *string `json:"publicCertData,omitempty"`
}
