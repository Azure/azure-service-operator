// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of PublicIPAddresses_Spec. Use v1beta20201101.PublicIPAddresses_Spec instead
type PublicIPAddresses_Spec_ARM struct {
	ExtendedLocation *ExtendedLocation_ARM                `json:"extendedLocation,omitempty"`
	Location         *string                              `json:"location,omitempty"`
	Name             string                               `json:"name,omitempty"`
	Properties       *PublicIPAddressPropertiesFormat_ARM `json:"properties,omitempty"`
	Sku              *PublicIPAddressSku_ARM              `json:"sku,omitempty"`
	Tags             map[string]string                    `json:"tags,omitempty"`
	Zones            []string                             `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PublicIPAddresses_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (addresses PublicIPAddresses_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (addresses *PublicIPAddresses_Spec_ARM) GetName() string {
	return addresses.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/publicIPAddresses"
func (addresses *PublicIPAddresses_Spec_ARM) GetType() string {
	return "Microsoft.Network/publicIPAddresses"
}

// Deprecated version of PublicIPAddressPropertiesFormat. Use v1beta20201101.PublicIPAddressPropertiesFormat instead
type PublicIPAddressPropertiesFormat_ARM struct {
	DdosSettings             *DdosSettings_ARM                                         `json:"ddosSettings,omitempty"`
	DnsSettings              *PublicIPAddressDnsSettings_ARM                           `json:"dnsSettings,omitempty"`
	IdleTimeoutInMinutes     *int                                                      `json:"idleTimeoutInMinutes,omitempty"`
	IpAddress                *string                                                   `json:"ipAddress,omitempty"`
	IpTags                   []IpTag_ARM                                               `json:"ipTags,omitempty"`
	PublicIPAddressVersion   *PublicIPAddressPropertiesFormat_PublicIPAddressVersion   `json:"publicIPAddressVersion,omitempty"`
	PublicIPAllocationMethod *PublicIPAddressPropertiesFormat_PublicIPAllocationMethod `json:"publicIPAllocationMethod,omitempty"`
	PublicIPPrefix           *SubResource_ARM                                          `json:"publicIPPrefix,omitempty"`
}

// Deprecated version of PublicIPAddressSku. Use v1beta20201101.PublicIPAddressSku instead
type PublicIPAddressSku_ARM struct {
	Name *PublicIPAddressSku_Name `json:"name,omitempty"`
	Tier *PublicIPAddressSku_Tier `json:"tier,omitempty"`
}

// Deprecated version of DdosSettings. Use v1beta20201101.DdosSettings instead
type DdosSettings_ARM struct {
	DdosCustomPolicy   *SubResource_ARM                 `json:"ddosCustomPolicy,omitempty"`
	ProtectedIP        *bool                            `json:"protectedIP,omitempty"`
	ProtectionCoverage *DdosSettings_ProtectionCoverage `json:"protectionCoverage,omitempty"`
}

// Deprecated version of IpTag. Use v1beta20201101.IpTag instead
type IpTag_ARM struct {
	IpTagType *string `json:"ipTagType,omitempty"`
	Tag       *string `json:"tag,omitempty"`
}

// Deprecated version of PublicIPAddressDnsSettings. Use v1beta20201101.PublicIPAddressDnsSettings instead
type PublicIPAddressDnsSettings_ARM struct {
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`
	Fqdn            *string `json:"fqdn,omitempty"`
	ReverseFqdn     *string `json:"reverseFqdn,omitempty"`
}

// Deprecated version of PublicIPAddressSku_Name. Use v1beta20201101.PublicIPAddressSku_Name instead
// +kubebuilder:validation:Enum={"Basic","Standard"}
type PublicIPAddressSku_Name string

const (
	PublicIPAddressSku_Name_Basic    = PublicIPAddressSku_Name("Basic")
	PublicIPAddressSku_Name_Standard = PublicIPAddressSku_Name("Standard")
)

// Deprecated version of PublicIPAddressSku_Tier. Use v1beta20201101.PublicIPAddressSku_Tier instead
// +kubebuilder:validation:Enum={"Global","Regional"}
type PublicIPAddressSku_Tier string

const (
	PublicIPAddressSku_Tier_Global   = PublicIPAddressSku_Tier("Global")
	PublicIPAddressSku_Tier_Regional = PublicIPAddressSku_Tier("Regional")
)
