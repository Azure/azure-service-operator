// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type PublicIPAddress_Spec struct {
	// ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocation `json:"extendedLocation,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Public IP address properties.
	Properties *PublicIPAddressPropertiesFormat `json:"properties,omitempty"`

	// Sku: The public IP address SKU.
	Sku *PublicIPAddressSku `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PublicIPAddress_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (address PublicIPAddress_Spec) GetAPIVersion() string {
	return "2020-11-01"
}

// GetName returns the Name of the resource
func (address *PublicIPAddress_Spec) GetName() string {
	return address.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/publicIPAddresses"
func (address *PublicIPAddress_Spec) GetType() string {
	return "Microsoft.Network/publicIPAddresses"
}

// Public IP address properties.
type PublicIPAddressPropertiesFormat struct {
	// DdosSettings: The DDoS protection custom policy associated with the public IP address.
	DdosSettings *DdosSettings `json:"ddosSettings,omitempty"`

	// DnsSettings: The FQDN of the DNS record associated with the public IP address.
	DnsSettings *PublicIPAddressDnsSettings `json:"dnsSettings,omitempty"`

	// IdleTimeoutInMinutes: The idle timeout of the public IP address.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// IpAddress: The IP address associated with the public IP address resource.
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpTags: The list of tags associated with the public IP address.
	IpTags []IpTag `json:"ipTags,omitempty"`

	// LinkedPublicIPAddress: The linked public IP address of the public IP address resource.
	LinkedPublicIPAddress *PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded `json:"linkedPublicIPAddress,omitempty"`

	// NatGateway: The NatGateway for the Public IP address.
	NatGateway *NatGatewaySpec_PublicIPAddress_SubResourceEmbedded `json:"natGateway,omitempty"`

	// PublicIPAddressVersion: The public IP address version.
	PublicIPAddressVersion *IPVersion `json:"publicIPAddressVersion,omitempty"`

	// PublicIPAllocationMethod: The public IP address allocation method.
	PublicIPAllocationMethod *IPAllocationMethod `json:"publicIPAllocationMethod,omitempty"`

	// PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated from.
	PublicIPPrefix *SubResource `json:"publicIPPrefix,omitempty"`

	// ServicePublicIPAddress: The service public IP address of the public IP address resource.
	ServicePublicIPAddress *PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded `json:"servicePublicIPAddress,omitempty"`
}

// SKU of a public IP address.
type PublicIPAddressSku struct {
	// Name: Name of a public IP address SKU.
	Name *PublicIPAddressSku_Name `json:"name,omitempty"`

	// Tier: Tier of a public IP address SKU.
	Tier *PublicIPAddressSku_Tier `json:"tier,omitempty"`
}

// Contains the DDoS protection settings of the public IP.
type DdosSettings struct {
	// DdosCustomPolicy: The DDoS custom policy associated with the public IP.
	DdosCustomPolicy *SubResource `json:"ddosCustomPolicy,omitempty"`

	// ProtectedIP: Enables DDoS protection on the public IP.
	ProtectedIP *bool `json:"protectedIP,omitempty"`

	// ProtectionCoverage: The DDoS protection policy customizability of the public IP. Only standard coverage will have the
	// ability to be customized.
	ProtectionCoverage *DdosSettings_ProtectionCoverage `json:"protectionCoverage,omitempty"`
}

// IP address allocation method.
// +kubebuilder:validation:Enum={"Dynamic","Static"}
type IPAllocationMethod string

const (
	IPAllocationMethod_Dynamic = IPAllocationMethod("Dynamic")
	IPAllocationMethod_Static  = IPAllocationMethod("Static")
)

// Mapping from string to IPAllocationMethod
var iPAllocationMethod_Values = map[string]IPAllocationMethod{
	"dynamic": IPAllocationMethod_Dynamic,
	"static":  IPAllocationMethod_Static,
}

// Contains the IpTag associated with the object.
type IpTag struct {
	// IpTagType: The IP tag type. Example: FirstPartyUsage.
	IpTagType *string `json:"ipTagType,omitempty"`

	// Tag: The value of the IP tag associated with the public IP. Example: SQL.
	Tag *string `json:"tag,omitempty"`
}

// IP address version.
// +kubebuilder:validation:Enum={"IPv4","IPv6"}
type IPVersion string

const (
	IPVersion_IPv4 = IPVersion("IPv4")
	IPVersion_IPv6 = IPVersion("IPv6")
)

// Mapping from string to IPVersion
var iPVersion_Values = map[string]IPVersion{
	"ipv4": IPVersion_IPv4,
	"ipv6": IPVersion_IPv6,
}

// Nat Gateway resource.
type NatGatewaySpec_PublicIPAddress_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// Contains FQDN of the DNS record associated with the public IP address.
type PublicIPAddressDnsSettings struct {
	// DomainNameLabel: The domain name label. The concatenation of the domain name label and the regionalized DNS zone make up
	// the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS
	// record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`

	// Fqdn: The Fully Qualified Domain Name of the A DNS record associated with the public IP. This is the concatenation of
	// the domainNameLabel and the regionalized DNS zone.
	Fqdn *string `json:"fqdn,omitempty"`

	// ReverseFqdn: The reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If
	// the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain
	// to the reverse FQDN.
	ReverseFqdn *string `json:"reverseFqdn,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","Standard"}
type PublicIPAddressSku_Name string

const (
	PublicIPAddressSku_Name_Basic    = PublicIPAddressSku_Name("Basic")
	PublicIPAddressSku_Name_Standard = PublicIPAddressSku_Name("Standard")
)

// Mapping from string to PublicIPAddressSku_Name
var publicIPAddressSku_Name_Values = map[string]PublicIPAddressSku_Name{
	"basic":    PublicIPAddressSku_Name_Basic,
	"standard": PublicIPAddressSku_Name_Standard,
}

// +kubebuilder:validation:Enum={"Global","Regional"}
type PublicIPAddressSku_Tier string

const (
	PublicIPAddressSku_Tier_Global   = PublicIPAddressSku_Tier("Global")
	PublicIPAddressSku_Tier_Regional = PublicIPAddressSku_Tier("Regional")
)

// Mapping from string to PublicIPAddressSku_Tier
var publicIPAddressSku_Tier_Values = map[string]PublicIPAddressSku_Tier{
	"global":   PublicIPAddressSku_Tier_Global,
	"regional": PublicIPAddressSku_Tier_Regional,
}

// Public IP address resource.
type PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","Standard"}
type DdosSettings_ProtectionCoverage string

const (
	DdosSettings_ProtectionCoverage_Basic    = DdosSettings_ProtectionCoverage("Basic")
	DdosSettings_ProtectionCoverage_Standard = DdosSettings_ProtectionCoverage("Standard")
)

// Mapping from string to DdosSettings_ProtectionCoverage
var ddosSettings_ProtectionCoverage_Values = map[string]DdosSettings_ProtectionCoverage{
	"basic":    DdosSettings_ProtectionCoverage_Basic,
	"standard": DdosSettings_ProtectionCoverage_Standard,
}
