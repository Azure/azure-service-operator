// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of VirtualNetworkGateways_Spec. Use v1beta20201101.VirtualNetworkGateways_Spec instead
type VirtualNetworkGateways_SpecARM struct {
	Location   *string                                    `json:"location,omitempty"`
	Name       string                                     `json:"name,omitempty"`
	Properties *VirtualNetworkGateways_Spec_PropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                          `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworkGateways_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (gateways VirtualNetworkGateways_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (gateways *VirtualNetworkGateways_SpecARM) GetName() string {
	return gateways.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworkGateways"
func (gateways *VirtualNetworkGateways_SpecARM) GetType() string {
	return "Microsoft.Network/virtualNetworkGateways"
}

// Deprecated version of VirtualNetworkGateways_Spec_Properties. Use v1beta20201101.VirtualNetworkGateways_Spec_Properties instead
type VirtualNetworkGateways_Spec_PropertiesARM struct {
	ActiveActive                   *bool                                                             `json:"activeActive,omitempty"`
	BgpSettings                    *BgpSettingsARM                                                   `json:"bgpSettings,omitempty"`
	CustomRoutes                   *AddressSpaceARM                                                  `json:"customRoutes,omitempty"`
	EnableBgp                      *bool                                                             `json:"enableBgp,omitempty"`
	EnableDnsForwarding            *bool                                                             `json:"enableDnsForwarding,omitempty"`
	EnablePrivateIpAddress         *bool                                                             `json:"enablePrivateIpAddress,omitempty"`
	GatewayDefaultSite             *SubResourceARM                                                   `json:"gatewayDefaultSite,omitempty"`
	GatewayType                    *VirtualNetworkGatewaysSpecPropertiesGatewayType                  `json:"gatewayType,omitempty"`
	IpConfigurations               []VirtualNetworkGateways_Spec_Properties_IpConfigurationsARM      `json:"ipConfigurations,omitempty"`
	Sku                            *VirtualNetworkGatewaySkuARM                                      `json:"sku,omitempty"`
	VNetExtendedLocationResourceId *string                                                           `json:"vNetExtendedLocationResourceId,omitempty"`
	VirtualNetworkExtendedLocation *ExtendedLocationARM                                              `json:"virtualNetworkExtendedLocation,omitempty"`
	VpnClientConfiguration         *VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM `json:"vpnClientConfiguration,omitempty"`
	VpnGatewayGeneration           *VirtualNetworkGatewaysSpecPropertiesVpnGatewayGeneration         `json:"vpnGatewayGeneration,omitempty"`
	VpnType                        *VirtualNetworkGatewaysSpecPropertiesVpnType                      `json:"vpnType,omitempty"`
}

// Deprecated version of AddressSpace. Use v1beta20201101.AddressSpace instead
type AddressSpaceARM struct {
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`
}

// Deprecated version of BgpSettings. Use v1beta20201101.BgpSettings instead
type BgpSettingsARM struct {
	Asn                 *uint32                               `json:"asn,omitempty"`
	BgpPeeringAddress   *string                               `json:"bgpPeeringAddress,omitempty"`
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddressARM `json:"bgpPeeringAddresses,omitempty"`
	PeerWeight          *int                                  `json:"peerWeight,omitempty"`
}

// Deprecated version of VirtualNetworkGatewaySku. Use v1beta20201101.VirtualNetworkGatewaySku instead
type VirtualNetworkGatewaySkuARM struct {
	Name *VirtualNetworkGatewaySkuName `json:"name,omitempty"`
	Tier *VirtualNetworkGatewaySkuTier `json:"tier,omitempty"`
}

// Deprecated version of VirtualNetworkGateways_Spec_Properties_IpConfigurations. Use v1beta20201101.VirtualNetworkGateways_Spec_Properties_IpConfigurations instead
type VirtualNetworkGateways_Spec_Properties_IpConfigurationsARM struct {
	Name       *string                                                  `json:"name,omitempty"`
	Properties *VirtualNetworkGatewayIPConfigurationPropertiesFormatARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration. Use v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration instead
type VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM struct {
	AadAudience                  *string                                                                                         `json:"aadAudience,omitempty"`
	AadIssuer                    *string                                                                                         `json:"aadIssuer,omitempty"`
	AadTenant                    *string                                                                                         `json:"aadTenant,omitempty"`
	RadiusServerAddress          *string                                                                                         `json:"radiusServerAddress,omitempty"`
	RadiusServerSecret           *string                                                                                         `json:"radiusServerSecret,omitempty"`
	RadiusServers                []RadiusServerARM                                                                               `json:"radiusServers,omitempty"`
	VpnAuthenticationTypes       []VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnAuthenticationTypes              `json:"vpnAuthenticationTypes,omitempty"`
	VpnClientAddressPool         *AddressSpaceARM                                                                                `json:"vpnClientAddressPool,omitempty"`
	VpnClientIpsecPolicies       []IpsecPolicyARM                                                                                `json:"vpnClientIpsecPolicies,omitempty"`
	VpnClientProtocols           []VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnClientProtocols                  `json:"vpnClientProtocols,omitempty"`
	VpnClientRevokedCertificates []VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificatesARM `json:"vpnClientRevokedCertificates,omitempty"`
	VpnClientRootCertificates    []VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificatesARM    `json:"vpnClientRootCertificates,omitempty"`
}

// Deprecated version of IPConfigurationBgpPeeringAddress. Use v1beta20201101.IPConfigurationBgpPeeringAddress instead
type IPConfigurationBgpPeeringAddressARM struct {
	CustomBgpIpAddresses []string `json:"customBgpIpAddresses,omitempty"`
	IpconfigurationId    *string  `json:"ipconfigurationId,omitempty"`
}

// Deprecated version of IpsecPolicy. Use v1beta20201101.IpsecPolicy instead
type IpsecPolicyARM struct {
	DhGroup             *IpsecPolicyDhGroup         `json:"dhGroup,omitempty"`
	IkeEncryption       *IpsecPolicyIkeEncryption   `json:"ikeEncryption,omitempty"`
	IkeIntegrity        *IpsecPolicyIkeIntegrity    `json:"ikeIntegrity,omitempty"`
	IpsecEncryption     *IpsecPolicyIpsecEncryption `json:"ipsecEncryption,omitempty"`
	IpsecIntegrity      *IpsecPolicyIpsecIntegrity  `json:"ipsecIntegrity,omitempty"`
	PfsGroup            *IpsecPolicyPfsGroup        `json:"pfsGroup,omitempty"`
	SaDataSizeKilobytes *int                        `json:"saDataSizeKilobytes,omitempty"`
	SaLifeTimeSeconds   *int                        `json:"saLifeTimeSeconds,omitempty"`
}

// Deprecated version of RadiusServer. Use v1beta20201101.RadiusServer instead
type RadiusServerARM struct {
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty"`
	RadiusServerScore   *int    `json:"radiusServerScore,omitempty"`
	RadiusServerSecret  *string `json:"radiusServerSecret,omitempty"`
}

// Deprecated version of VirtualNetworkGatewayIPConfigurationPropertiesFormat. Use v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormat instead
type VirtualNetworkGatewayIPConfigurationPropertiesFormatARM struct {
	PrivateIPAllocationMethod *VirtualNetworkGatewayIPConfigurationPropertiesFormatPrivateIPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`
	PublicIPAddress           *SubResourceARM                                                                `json:"publicIPAddress,omitempty"`
	Subnet                    *SubResourceARM                                                                `json:"subnet,omitempty"`
}

// Deprecated version of VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificates. Use v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificates instead
type VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificatesARM struct {
	Name       *string                                         `json:"name,omitempty"`
	Properties *VpnClientRevokedCertificatePropertiesFormatARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificates. Use v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificates instead
type VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificatesARM struct {
	Name       *string                                      `json:"name,omitempty"`
	Properties *VpnClientRootCertificatePropertiesFormatARM `json:"properties,omitempty"`
}

// Deprecated version of VpnClientRevokedCertificatePropertiesFormat. Use v1beta20201101.VpnClientRevokedCertificatePropertiesFormat instead
type VpnClientRevokedCertificatePropertiesFormatARM struct {
	Thumbprint *string `json:"thumbprint,omitempty"`
}

// Deprecated version of VpnClientRootCertificatePropertiesFormat. Use v1beta20201101.VpnClientRootCertificatePropertiesFormat instead
type VpnClientRootCertificatePropertiesFormatARM struct {
	PublicCertData *string `json:"publicCertData,omitempty"`
}
