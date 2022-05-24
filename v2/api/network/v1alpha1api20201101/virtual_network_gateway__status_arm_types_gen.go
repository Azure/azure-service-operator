// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

// Deprecated version of VirtualNetworkGateway_Status. Use v1beta20201101.VirtualNetworkGateway_Status instead
type VirtualNetworkGateway_StatusARM struct {
	Etag             *string                                          `json:"etag,omitempty"`
	ExtendedLocation *ExtendedLocation_StatusARM                      `json:"extendedLocation,omitempty"`
	Id               *string                                          `json:"id,omitempty"`
	Location         *string                                          `json:"location,omitempty"`
	Name             *string                                          `json:"name,omitempty"`
	Properties       *VirtualNetworkGatewayPropertiesFormat_StatusARM `json:"properties,omitempty"`
	Tags             map[string]string                                `json:"tags,omitempty"`
	Type             *string                                          `json:"type,omitempty"`
}

// Deprecated version of VirtualNetworkGatewayPropertiesFormat_Status. Use v1beta20201101.VirtualNetworkGatewayPropertiesFormat_Status instead
type VirtualNetworkGatewayPropertiesFormat_StatusARM struct {
	ActiveActive                   *bool                                            `json:"activeActive,omitempty"`
	BgpSettings                    *BgpSettings_StatusARM                           `json:"bgpSettings,omitempty"`
	CustomRoutes                   *AddressSpace_StatusARM                          `json:"customRoutes,omitempty"`
	EnableBgp                      *bool                                            `json:"enableBgp,omitempty"`
	EnableDnsForwarding            *bool                                            `json:"enableDnsForwarding,omitempty"`
	EnablePrivateIpAddress         *bool                                            `json:"enablePrivateIpAddress,omitempty"`
	GatewayDefaultSite             *SubResource_StatusARM                           `json:"gatewayDefaultSite,omitempty"`
	GatewayType                    *string                                          `json:"gatewayType,omitempty"`
	InboundDnsForwardingEndpoint   *string                                          `json:"inboundDnsForwardingEndpoint,omitempty"`
	IpConfigurations               []VirtualNetworkGatewayIPConfiguration_StatusARM `json:"ipConfigurations,omitempty"`
	ProvisioningState              *string                                          `json:"provisioningState,omitempty"`
	ResourceGuid                   *string                                          `json:"resourceGuid,omitempty"`
	Sku                            *VirtualNetworkGatewaySku_StatusARM              `json:"sku,omitempty"`
	VNetExtendedLocationResourceId *string                                          `json:"vNetExtendedLocationResourceId,omitempty"`
	VpnClientConfiguration         *VpnClientConfiguration_StatusARM                `json:"vpnClientConfiguration,omitempty"`
	VpnGatewayGeneration           *string                                          `json:"vpnGatewayGeneration,omitempty"`
	VpnType                        *string                                          `json:"vpnType,omitempty"`
}

// Deprecated version of AddressSpace_Status. Use v1beta20201101.AddressSpace_Status instead
type AddressSpace_StatusARM struct {
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`
}

// Deprecated version of BgpSettings_Status. Use v1beta20201101.BgpSettings_Status instead
type BgpSettings_StatusARM struct {
	Asn                 *uint32                                      `json:"asn,omitempty"`
	BgpPeeringAddress   *string                                      `json:"bgpPeeringAddress,omitempty"`
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddress_StatusARM `json:"bgpPeeringAddresses,omitempty"`
	PeerWeight          *int                                         `json:"peerWeight,omitempty"`
}

// Deprecated version of VirtualNetworkGatewayIPConfiguration_Status. Use v1beta20201101.VirtualNetworkGatewayIPConfiguration_Status instead
type VirtualNetworkGatewayIPConfiguration_StatusARM struct {
	Etag       *string                                                         `json:"etag,omitempty"`
	Id         *string                                                         `json:"id,omitempty"`
	Name       *string                                                         `json:"name,omitempty"`
	Properties *VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualNetworkGatewaySku_Status. Use v1beta20201101.VirtualNetworkGatewaySku_Status instead
type VirtualNetworkGatewaySku_StatusARM struct {
	Capacity *int    `json:"capacity,omitempty"`
	Name     *string `json:"name,omitempty"`
	Tier     *string `json:"tier,omitempty"`
}

// Deprecated version of VpnClientConfiguration_Status. Use v1beta20201101.VpnClientConfiguration_Status instead
type VpnClientConfiguration_StatusARM struct {
	AadAudience                  *string                                 `json:"aadAudience,omitempty"`
	AadIssuer                    *string                                 `json:"aadIssuer,omitempty"`
	AadTenant                    *string                                 `json:"aadTenant,omitempty"`
	RadiusServerAddress          *string                                 `json:"radiusServerAddress,omitempty"`
	RadiusServerSecret           *string                                 `json:"radiusServerSecret,omitempty"`
	RadiusServers                []RadiusServer_StatusARM                `json:"radiusServers,omitempty"`
	VpnAuthenticationTypes       []string                                `json:"vpnAuthenticationTypes,omitempty"`
	VpnClientAddressPool         *AddressSpace_StatusARM                 `json:"vpnClientAddressPool,omitempty"`
	VpnClientIpsecPolicies       []IpsecPolicy_StatusARM                 `json:"vpnClientIpsecPolicies,omitempty"`
	VpnClientProtocols           []string                                `json:"vpnClientProtocols,omitempty"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificate_StatusARM `json:"vpnClientRevokedCertificates,omitempty"`
	VpnClientRootCertificates    []VpnClientRootCertificate_StatusARM    `json:"vpnClientRootCertificates,omitempty"`
}

// Deprecated version of IPConfigurationBgpPeeringAddress_Status. Use v1beta20201101.IPConfigurationBgpPeeringAddress_Status instead
type IPConfigurationBgpPeeringAddress_StatusARM struct {
	CustomBgpIpAddresses  []string `json:"customBgpIpAddresses,omitempty"`
	DefaultBgpIpAddresses []string `json:"defaultBgpIpAddresses,omitempty"`
	IpconfigurationId     *string  `json:"ipconfigurationId,omitempty"`
	TunnelIpAddresses     []string `json:"tunnelIpAddresses,omitempty"`
}

// Deprecated version of IpsecPolicy_Status. Use v1beta20201101.IpsecPolicy_Status instead
type IpsecPolicy_StatusARM struct {
	DhGroup             *string `json:"dhGroup,omitempty"`
	IkeEncryption       *string `json:"ikeEncryption,omitempty"`
	IkeIntegrity        *string `json:"ikeIntegrity,omitempty"`
	IpsecEncryption     *string `json:"ipsecEncryption,omitempty"`
	IpsecIntegrity      *string `json:"ipsecIntegrity,omitempty"`
	PfsGroup            *string `json:"pfsGroup,omitempty"`
	SaDataSizeKilobytes *int    `json:"saDataSizeKilobytes,omitempty"`
	SaLifeTimeSeconds   *int    `json:"saLifeTimeSeconds,omitempty"`
}

// Deprecated version of RadiusServer_Status. Use v1beta20201101.RadiusServer_Status instead
type RadiusServer_StatusARM struct {
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty"`
	RadiusServerScore   *int    `json:"radiusServerScore,omitempty"`
	RadiusServerSecret  *string `json:"radiusServerSecret,omitempty"`
}

// Deprecated version of VirtualNetworkGatewayIPConfigurationPropertiesFormat_Status. Use v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormat_Status instead
type VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM struct {
	PrivateIPAddress          *string                `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *string                `json:"privateIPAllocationMethod,omitempty"`
	ProvisioningState         *string                `json:"provisioningState,omitempty"`
	PublicIPAddress           *SubResource_StatusARM `json:"publicIPAddress,omitempty"`
	Subnet                    *SubResource_StatusARM `json:"subnet,omitempty"`
}

// Deprecated version of VpnClientRevokedCertificate_Status. Use v1beta20201101.VpnClientRevokedCertificate_Status instead
type VpnClientRevokedCertificate_StatusARM struct {
	Etag       *string                                                `json:"etag,omitempty"`
	Id         *string                                                `json:"id,omitempty"`
	Name       *string                                                `json:"name,omitempty"`
	Properties *VpnClientRevokedCertificatePropertiesFormat_StatusARM `json:"properties,omitempty"`
}

// Deprecated version of VpnClientRootCertificate_Status. Use v1beta20201101.VpnClientRootCertificate_Status instead
type VpnClientRootCertificate_StatusARM struct {
	Etag       *string                                             `json:"etag,omitempty"`
	Id         *string                                             `json:"id,omitempty"`
	Name       *string                                             `json:"name,omitempty"`
	Properties *VpnClientRootCertificatePropertiesFormat_StatusARM `json:"properties,omitempty"`
}

// Deprecated version of VpnClientRevokedCertificatePropertiesFormat_Status. Use v1beta20201101.VpnClientRevokedCertificatePropertiesFormat_Status instead
type VpnClientRevokedCertificatePropertiesFormat_StatusARM struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`
	Thumbprint        *string `json:"thumbprint,omitempty"`
}

// Deprecated version of VpnClientRootCertificatePropertiesFormat_Status. Use v1beta20201101.VpnClientRootCertificatePropertiesFormat_Status instead
type VpnClientRootCertificatePropertiesFormat_StatusARM struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`
	PublicCertData    *string `json:"publicCertData,omitempty"`
}
