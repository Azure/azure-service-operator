// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// Public IP address resource.
type PublicIPAddress_STATUS struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocation_STATUS `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Public IP address properties.
	Properties *PublicIPAddressPropertiesFormat_STATUS `json:"properties,omitempty"`

	// Sku: The public IP address SKU.
	Sku *PublicIPAddressSku_STATUS `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

// Public IP address properties.
type PublicIPAddressPropertiesFormat_STATUS struct {
	// DdosSettings: The DDoS protection custom policy associated with the public IP address.
	DdosSettings *DdosSettings_STATUS `json:"ddosSettings,omitempty"`

	// DeleteOption: Specify what happens to the public IP address when the VM using it is deleted
	DeleteOption *PublicIPAddressPropertiesFormat_DeleteOption_STATUS `json:"deleteOption,omitempty"`

	// DnsSettings: The FQDN of the DNS record associated with the public IP address.
	DnsSettings *PublicIPAddressDnsSettings_STATUS `json:"dnsSettings,omitempty"`

	// IdleTimeoutInMinutes: The idle timeout of the public IP address.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// IpAddress: The IP address associated with the public IP address resource.
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpConfiguration: The IP configuration associated with the public IP address.
	IpConfiguration *IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded `json:"ipConfiguration,omitempty"`

	// IpTags: The list of tags associated with the public IP address.
	IpTags []IpTag_STATUS `json:"ipTags,omitempty"`

	// MigrationPhase: Migration phase of Public IP Address.
	MigrationPhase *PublicIPAddressPropertiesFormat_MigrationPhase_STATUS `json:"migrationPhase,omitempty"`

	// NatGateway: The NatGateway for the Public IP address.
	NatGateway *NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded `json:"natGateway,omitempty"`

	// ProvisioningState: The provisioning state of the public IP address resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicIPAddressVersion: The public IP address version.
	PublicIPAddressVersion *IPVersion_STATUS `json:"publicIPAddressVersion,omitempty"`

	// PublicIPAllocationMethod: The public IP address allocation method.
	PublicIPAllocationMethod *IPAllocationMethod_STATUS `json:"publicIPAllocationMethod,omitempty"`

	// PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated from.
	PublicIPPrefix *SubResource_STATUS `json:"publicIPPrefix,omitempty"`

	// ResourceGuid: The resource GUID property of the public IP address resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`
}

// SKU of a public IP address.
type PublicIPAddressSku_STATUS struct {
	// Name: Name of a public IP address SKU.
	Name *PublicIPAddressSku_Name_STATUS `json:"name,omitempty"`

	// Tier: Tier of a public IP address SKU.
	Tier *PublicIPAddressSku_Tier_STATUS `json:"tier,omitempty"`
}

// Contains the DDoS protection settings of the public IP.
type DdosSettings_STATUS struct {
	// DdosProtectionPlan: The DDoS protection plan associated with the public IP. Can only be set if ProtectionMode is Enabled
	DdosProtectionPlan *SubResource_STATUS `json:"ddosProtectionPlan,omitempty"`

	// ProtectionMode: The DDoS protection mode of the public IP
	ProtectionMode *DdosSettings_ProtectionMode_STATUS `json:"protectionMode,omitempty"`
}

// IP address allocation method.
type IPAllocationMethod_STATUS string

const (
	IPAllocationMethod_STATUS_Dynamic = IPAllocationMethod_STATUS("Dynamic")
	IPAllocationMethod_STATUS_Static  = IPAllocationMethod_STATUS("Static")
)

// Mapping from string to IPAllocationMethod_STATUS
var iPAllocationMethod_STATUS_Values = map[string]IPAllocationMethod_STATUS{
	"dynamic": IPAllocationMethod_STATUS_Dynamic,
	"static":  IPAllocationMethod_STATUS_Static,
}

// IP configuration.
type IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Contains the IpTag associated with the object.
type IpTag_STATUS struct {
	// IpTagType: The IP tag type. Example: FirstPartyUsage.
	IpTagType *string `json:"ipTagType,omitempty"`

	// Tag: The value of the IP tag associated with the public IP. Example: SQL.
	Tag *string `json:"tag,omitempty"`
}

// IP address version.
type IPVersion_STATUS string

const (
	IPVersion_STATUS_IPv4 = IPVersion_STATUS("IPv4")
	IPVersion_STATUS_IPv6 = IPVersion_STATUS("IPv6")
)

// Mapping from string to IPVersion_STATUS
var iPVersion_STATUS_Values = map[string]IPVersion_STATUS{
	"ipv4": IPVersion_STATUS_IPv4,
	"ipv6": IPVersion_STATUS_IPv6,
}

// Nat Gateway resource.
type NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Contains FQDN of the DNS record associated with the public IP address.
type PublicIPAddressDnsSettings_STATUS struct {
	// DomainNameLabel: The domain name label. The concatenation of the domain name label and the regionalized DNS zone make up
	// the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS
	// record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`

	// DomainNameLabelScope: The domain name label scope. If a domain name label and a domain name label scope are specified,
	// an A DNS record is created for the public IP in the Microsoft Azure DNS system with a hashed value includes in FQDN.
	DomainNameLabelScope *PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS `json:"domainNameLabelScope,omitempty"`

	// Fqdn: The Fully Qualified Domain Name of the A DNS record associated with the public IP. This is the concatenation of
	// the domainNameLabel and the regionalized DNS zone.
	Fqdn *string `json:"fqdn,omitempty"`

	// ReverseFqdn: The reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If
	// the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain
	// to the reverse FQDN.
	ReverseFqdn *string `json:"reverseFqdn,omitempty"`
}

type PublicIPAddressPropertiesFormat_DeleteOption_STATUS string

const (
	PublicIPAddressPropertiesFormat_DeleteOption_STATUS_Delete = PublicIPAddressPropertiesFormat_DeleteOption_STATUS("Delete")
	PublicIPAddressPropertiesFormat_DeleteOption_STATUS_Detach = PublicIPAddressPropertiesFormat_DeleteOption_STATUS("Detach")
)

// Mapping from string to PublicIPAddressPropertiesFormat_DeleteOption_STATUS
var publicIPAddressPropertiesFormat_DeleteOption_STATUS_Values = map[string]PublicIPAddressPropertiesFormat_DeleteOption_STATUS{
	"delete": PublicIPAddressPropertiesFormat_DeleteOption_STATUS_Delete,
	"detach": PublicIPAddressPropertiesFormat_DeleteOption_STATUS_Detach,
}

type PublicIPAddressPropertiesFormat_MigrationPhase_STATUS string

const (
	PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Abort     = PublicIPAddressPropertiesFormat_MigrationPhase_STATUS("Abort")
	PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Commit    = PublicIPAddressPropertiesFormat_MigrationPhase_STATUS("Commit")
	PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Committed = PublicIPAddressPropertiesFormat_MigrationPhase_STATUS("Committed")
	PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_None      = PublicIPAddressPropertiesFormat_MigrationPhase_STATUS("None")
	PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Prepare   = PublicIPAddressPropertiesFormat_MigrationPhase_STATUS("Prepare")
)

// Mapping from string to PublicIPAddressPropertiesFormat_MigrationPhase_STATUS
var publicIPAddressPropertiesFormat_MigrationPhase_STATUS_Values = map[string]PublicIPAddressPropertiesFormat_MigrationPhase_STATUS{
	"abort":     PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Abort,
	"commit":    PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Commit,
	"committed": PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Committed,
	"none":      PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_None,
	"prepare":   PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Prepare,
}

type PublicIPAddressSku_Name_STATUS string

const (
	PublicIPAddressSku_Name_STATUS_Basic    = PublicIPAddressSku_Name_STATUS("Basic")
	PublicIPAddressSku_Name_STATUS_Standard = PublicIPAddressSku_Name_STATUS("Standard")
)

// Mapping from string to PublicIPAddressSku_Name_STATUS
var publicIPAddressSku_Name_STATUS_Values = map[string]PublicIPAddressSku_Name_STATUS{
	"basic":    PublicIPAddressSku_Name_STATUS_Basic,
	"standard": PublicIPAddressSku_Name_STATUS_Standard,
}

type PublicIPAddressSku_Tier_STATUS string

const (
	PublicIPAddressSku_Tier_STATUS_Global   = PublicIPAddressSku_Tier_STATUS("Global")
	PublicIPAddressSku_Tier_STATUS_Regional = PublicIPAddressSku_Tier_STATUS("Regional")
)

// Mapping from string to PublicIPAddressSku_Tier_STATUS
var publicIPAddressSku_Tier_STATUS_Values = map[string]PublicIPAddressSku_Tier_STATUS{
	"global":   PublicIPAddressSku_Tier_STATUS_Global,
	"regional": PublicIPAddressSku_Tier_STATUS_Regional,
}

type DdosSettings_ProtectionMode_STATUS string

const (
	DdosSettings_ProtectionMode_STATUS_Disabled                = DdosSettings_ProtectionMode_STATUS("Disabled")
	DdosSettings_ProtectionMode_STATUS_Enabled                 = DdosSettings_ProtectionMode_STATUS("Enabled")
	DdosSettings_ProtectionMode_STATUS_VirtualNetworkInherited = DdosSettings_ProtectionMode_STATUS("VirtualNetworkInherited")
)

// Mapping from string to DdosSettings_ProtectionMode_STATUS
var ddosSettings_ProtectionMode_STATUS_Values = map[string]DdosSettings_ProtectionMode_STATUS{
	"disabled":                DdosSettings_ProtectionMode_STATUS_Disabled,
	"enabled":                 DdosSettings_ProtectionMode_STATUS_Enabled,
	"virtualnetworkinherited": DdosSettings_ProtectionMode_STATUS_VirtualNetworkInherited,
}

type PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS string

const (
	PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS_NoReuse            = PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS("NoReuse")
	PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS_ResourceGroupReuse = PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS("ResourceGroupReuse")
	PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS_SubscriptionReuse  = PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS("SubscriptionReuse")
	PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS_TenantReuse        = PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS("TenantReuse")
)

// Mapping from string to PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS
var publicIPAddressDnsSettings_DomainNameLabelScope_STATUS_Values = map[string]PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS{
	"noreuse":            PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS_NoReuse,
	"resourcegroupreuse": PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS_ResourceGroupReuse,
	"subscriptionreuse":  PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS_SubscriptionReuse,
	"tenantreuse":        PublicIPAddressDnsSettings_DomainNameLabelScope_STATUS_TenantReuse,
}
