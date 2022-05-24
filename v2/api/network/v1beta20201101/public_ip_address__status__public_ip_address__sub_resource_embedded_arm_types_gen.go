// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

type PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Public IP address properties.
	Properties *PublicIPAddressPropertiesFormat_StatusARM `json:"properties,omitempty"`

	// Sku: The public IP address SKU.
	Sku *PublicIPAddressSku_StatusARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type PublicIPAddressPropertiesFormat_StatusARM struct {
	// DdosSettings: The DDoS protection custom policy associated with the public IP address.
	DdosSettings *DdosSettings_StatusARM `json:"ddosSettings,omitempty"`

	// DnsSettings: The FQDN of the DNS record associated with the public IP address.
	DnsSettings *PublicIPAddressDnsSettings_StatusARM `json:"dnsSettings,omitempty"`

	// IdleTimeoutInMinutes: The idle timeout of the public IP address.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// IpAddress: The IP address associated with the public IP address resource.
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpConfiguration: The IP configuration associated with the public IP address.
	IpConfiguration *IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM `json:"ipConfiguration,omitempty"`

	// IpTags: The list of tags associated with the public IP address.
	IpTags []IpTag_StatusARM `json:"ipTags,omitempty"`

	// MigrationPhase: Migration phase of Public IP Address.
	MigrationPhase *string `json:"migrationPhase,omitempty"`

	// NatGateway: The NatGateway for the Public IP address.
	NatGateway *NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM `json:"natGateway,omitempty"`

	// ProvisioningState: The provisioning state of the public IP address resource.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// PublicIPAddressVersion: The public IP address version.
	PublicIPAddressVersion *string `json:"publicIPAddressVersion,omitempty"`

	// PublicIPAllocationMethod: The public IP address allocation method.
	PublicIPAllocationMethod *string `json:"publicIPAllocationMethod,omitempty"`

	// PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated from.
	PublicIPPrefix *SubResource_StatusARM `json:"publicIPPrefix,omitempty"`

	// ResourceGuid: The resource GUID property of the public IP address resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`
}

type PublicIPAddressSku_StatusARM struct {
	// Name: Name of a public IP address SKU.
	Name *string `json:"name,omitempty"`

	// Tier: Tier of a public IP address SKU.
	Tier *string `json:"tier,omitempty"`
}

type DdosSettings_StatusARM struct {
	// DdosCustomPolicy: The DDoS custom policy associated with the public IP.
	DdosCustomPolicy *SubResource_StatusARM `json:"ddosCustomPolicy,omitempty"`

	// ProtectedIP: Enables DDoS protection on the public IP.
	ProtectedIP *bool `json:"protectedIP,omitempty"`

	// ProtectionCoverage: The DDoS protection policy customizability of the public IP. Only standard coverage will have the
	// ability to be customized.
	ProtectionCoverage *string `json:"protectionCoverage,omitempty"`
}

type IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the IP configuration.
	Properties *IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM `json:"properties,omitempty"`
}

type IpTag_StatusARM struct {
	// IpTagType: The IP tag type. Example: FirstPartyUsage.
	IpTagType *string `json:"ipTagType,omitempty"`

	// Tag: The value of the IP tag associated with the public IP. Example: SQL.
	Tag *string `json:"tag,omitempty"`
}

type NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Sku: The nat gateway SKU.
	Sku *NatGatewaySku_StatusARM `json:"sku,omitempty"`

	// Zones: A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones []string `json:"zones,omitempty"`
}

type PublicIPAddressDnsSettings_StatusARM struct {
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

type IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM struct {
	// PrivateIPAddress: The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	// PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *string `json:"privateIPAllocationMethod,omitempty"`

	// ProvisioningState: The provisioning state of the IP configuration resource.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// Subnet: The reference to the subnet resource.
	Subnet *Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM `json:"subnet,omitempty"`
}

type NatGatewaySku_StatusARM struct {
	// Name: Name of Nat Gateway SKU.
	Name *string `json:"name,omitempty"`
}

type Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}
