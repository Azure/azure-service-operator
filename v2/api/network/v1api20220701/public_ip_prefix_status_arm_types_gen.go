// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

// Public IP prefix resource.
type PublicIPPrefix_STATUS_ARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Public IP prefix properties.
	Properties *PublicIPPrefixPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The public IP prefix SKU.
	Sku *PublicIPPrefixSku_STATUS_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

// Public IP prefix properties.
type PublicIPPrefixPropertiesFormat_STATUS_ARM struct {
	// CustomIPPrefix: The customIpPrefix that this prefix is associated with.
	CustomIPPrefix *PublicIpPrefixSubResource_STATUS_ARM `json:"customIPPrefix,omitempty"`

	// IpPrefix: The allocated Prefix.
	IpPrefix *string `json:"ipPrefix,omitempty"`

	// IpTags: The list of tags associated with the public IP prefix.
	IpTags []IpTag_STATUS_ARM `json:"ipTags,omitempty"`

	// LoadBalancerFrontendIpConfiguration: The reference to load balancer frontend IP configuration associated with the public
	// IP prefix.
	LoadBalancerFrontendIpConfiguration *PublicIpPrefixSubResource_STATUS_ARM `json:"loadBalancerFrontendIpConfiguration,omitempty"`

	// NatGateway: NatGateway of Public IP Prefix.
	NatGateway *NatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded_ARM `json:"natGateway,omitempty"`

	// PrefixLength: The Length of the Public IP Prefix.
	PrefixLength *int `json:"prefixLength,omitempty"`

	// ProvisioningState: The provisioning state of the public IP prefix resource.
	ProvisioningState *PublicIpPrefixProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`

	// PublicIPAddressVersion: The public IP address version.
	PublicIPAddressVersion *IPVersion_STATUS_ARM `json:"publicIPAddressVersion,omitempty"`

	// PublicIPAddresses: The list of all referenced PublicIPAddresses.
	PublicIPAddresses []ReferencedPublicIpAddress_STATUS_ARM `json:"publicIPAddresses,omitempty"`

	// ResourceGuid: The resource GUID property of the public IP prefix resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`
}

// SKU of a public IP prefix.
type PublicIPPrefixSku_STATUS_ARM struct {
	// Name: Name of a public IP prefix SKU.
	Name *PublicIPPrefixSku_Name_STATUS_ARM `json:"name,omitempty"`

	// Tier: Tier of a public IP prefix SKU.
	Tier *PublicIPPrefixSku_Tier_STATUS_ARM `json:"tier,omitempty"`
}

// Contains the IpTag associated with the object.
type IpTag_STATUS_ARM struct {
	// IpTagType: The IP tag type. Example: FirstPartyUsage.
	IpTagType *string `json:"ipTagType,omitempty"`

	// Tag: The value of the IP tag associated with the public IP. Example: SQL.
	Tag *string `json:"tag,omitempty"`
}

// IP address version.
type IPVersion_STATUS_ARM string

const (
	IPVersion_STATUS_ARM_IPv4 = IPVersion_STATUS_ARM("IPv4")
	IPVersion_STATUS_ARM_IPv6 = IPVersion_STATUS_ARM("IPv6")
)

// Mapping from string to IPVersion_STATUS_ARM
var iPVersion_STATUS_ARM_Values = map[string]IPVersion_STATUS_ARM{
	"ipv4": IPVersion_STATUS_ARM_IPv4,
	"ipv6": IPVersion_STATUS_ARM_IPv6,
}

// Nat Gateway resource.
type NatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// The current provisioning state.
type PublicIpPrefixProvisioningState_STATUS_ARM string

const (
	PublicIpPrefixProvisioningState_STATUS_ARM_Deleting  = PublicIpPrefixProvisioningState_STATUS_ARM("Deleting")
	PublicIpPrefixProvisioningState_STATUS_ARM_Failed    = PublicIpPrefixProvisioningState_STATUS_ARM("Failed")
	PublicIpPrefixProvisioningState_STATUS_ARM_Succeeded = PublicIpPrefixProvisioningState_STATUS_ARM("Succeeded")
	PublicIpPrefixProvisioningState_STATUS_ARM_Updating  = PublicIpPrefixProvisioningState_STATUS_ARM("Updating")
)

// Mapping from string to PublicIpPrefixProvisioningState_STATUS_ARM
var publicIpPrefixProvisioningState_STATUS_ARM_Values = map[string]PublicIpPrefixProvisioningState_STATUS_ARM{
	"deleting":  PublicIpPrefixProvisioningState_STATUS_ARM_Deleting,
	"failed":    PublicIpPrefixProvisioningState_STATUS_ARM_Failed,
	"succeeded": PublicIpPrefixProvisioningState_STATUS_ARM_Succeeded,
	"updating":  PublicIpPrefixProvisioningState_STATUS_ARM_Updating,
}

type PublicIPPrefixSku_Name_STATUS_ARM string

const PublicIPPrefixSku_Name_STATUS_ARM_Standard = PublicIPPrefixSku_Name_STATUS_ARM("Standard")

// Mapping from string to PublicIPPrefixSku_Name_STATUS_ARM
var publicIPPrefixSku_Name_STATUS_ARM_Values = map[string]PublicIPPrefixSku_Name_STATUS_ARM{
	"standard": PublicIPPrefixSku_Name_STATUS_ARM_Standard,
}

type PublicIPPrefixSku_Tier_STATUS_ARM string

const (
	PublicIPPrefixSku_Tier_STATUS_ARM_Global   = PublicIPPrefixSku_Tier_STATUS_ARM("Global")
	PublicIPPrefixSku_Tier_STATUS_ARM_Regional = PublicIPPrefixSku_Tier_STATUS_ARM("Regional")
)

// Mapping from string to PublicIPPrefixSku_Tier_STATUS_ARM
var publicIPPrefixSku_Tier_STATUS_ARM_Values = map[string]PublicIPPrefixSku_Tier_STATUS_ARM{
	"global":   PublicIPPrefixSku_Tier_STATUS_ARM_Global,
	"regional": PublicIPPrefixSku_Tier_STATUS_ARM_Regional,
}

// Reference to another subresource.
type PublicIpPrefixSubResource_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Reference to a public IP address.
type ReferencedPublicIpAddress_STATUS_ARM struct {
	// Id: The PublicIPAddress Reference.
	Id *string `json:"id,omitempty"`
}
