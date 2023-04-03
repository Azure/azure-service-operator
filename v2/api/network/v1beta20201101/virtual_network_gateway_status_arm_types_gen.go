// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

// Deprecated version of VirtualNetworkGateway_STATUS. Use v1api20201101.VirtualNetworkGateway_STATUS instead
type VirtualNetworkGateway_STATUS_ARM struct {
	Etag             *string                                           `json:"etag,omitempty"`
	ExtendedLocation *ExtendedLocation_STATUS_ARM                      `json:"extendedLocation,omitempty"`
	Id               *string                                           `json:"id,omitempty"`
	Location         *string                                           `json:"location,omitempty"`
	Name             *string                                           `json:"name,omitempty"`
	Properties       *VirtualNetworkGatewayPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Tags             map[string]string                                 `json:"tags,omitempty"`
	Type             *string                                           `json:"type,omitempty"`
}

// Deprecated version of VirtualNetworkGatewayPropertiesFormat_STATUS. Use v1api20201101.VirtualNetworkGatewayPropertiesFormat_STATUS instead
type VirtualNetworkGatewayPropertiesFormat_STATUS_ARM struct {
	ActiveActive                   *bool                                                              `json:"activeActive,omitempty"`
	BgpSettings                    *BgpSettings_STATUS_ARM                                            `json:"bgpSettings,omitempty"`
	CustomRoutes                   *AddressSpace_STATUS_ARM                                           `json:"customRoutes,omitempty"`
	EnableBgp                      *bool                                                              `json:"enableBgp,omitempty"`
	EnableDnsForwarding            *bool                                                              `json:"enableDnsForwarding,omitempty"`
	EnablePrivateIpAddress         *bool                                                              `json:"enablePrivateIpAddress,omitempty"`
	GatewayDefaultSite             *SubResource_STATUS_ARM                                            `json:"gatewayDefaultSite,omitempty"`
	GatewayType                    *VirtualNetworkGatewayPropertiesFormat_GatewayType_STATUS          `json:"gatewayType,omitempty"`
	InboundDnsForwardingEndpoint   *string                                                            `json:"inboundDnsForwardingEndpoint,omitempty"`
	IpConfigurations               []VirtualNetworkGatewayIPConfiguration_STATUS_ARM                  `json:"ipConfigurations,omitempty"`
	ProvisioningState              *ProvisioningState_STATUS                                          `json:"provisioningState,omitempty"`
	ResourceGuid                   *string                                                            `json:"resourceGuid,omitempty"`
	Sku                            *VirtualNetworkGatewaySku_STATUS_ARM                               `json:"sku,omitempty"`
	VNetExtendedLocationResourceId *string                                                            `json:"vNetExtendedLocationResourceId,omitempty"`
	VpnClientConfiguration         *VpnClientConfiguration_STATUS_ARM                                 `json:"vpnClientConfiguration,omitempty"`
	VpnGatewayGeneration           *VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration_STATUS `json:"vpnGatewayGeneration,omitempty"`
	VpnType                        *VirtualNetworkGatewayPropertiesFormat_VpnType_STATUS              `json:"vpnType,omitempty"`
}

// Deprecated version of AddressSpace_STATUS. Use v1api20201101.AddressSpace_STATUS instead
type AddressSpace_STATUS_ARM struct {
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`
}

// Deprecated version of BgpSettings_STATUS. Use v1api20201101.BgpSettings_STATUS instead
type BgpSettings_STATUS_ARM struct {
	Asn                 *uint32                                       `json:"asn,omitempty"`
	BgpPeeringAddress   *string                                       `json:"bgpPeeringAddress,omitempty"`
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddress_STATUS_ARM `json:"bgpPeeringAddresses,omitempty"`
	PeerWeight          *int                                          `json:"peerWeight,omitempty"`
}

// Deprecated version of VirtualNetworkGatewayIPConfiguration_STATUS. Use v1api20201101.VirtualNetworkGatewayIPConfiguration_STATUS instead
type VirtualNetworkGatewayIPConfiguration_STATUS_ARM struct {
	Etag       *string                                                          `json:"etag,omitempty"`
	Id         *string                                                          `json:"id,omitempty"`
	Name       *string                                                          `json:"name,omitempty"`
	Properties *VirtualNetworkGatewayIPConfigurationPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualNetworkGatewaySku_STATUS. Use v1api20201101.VirtualNetworkGatewaySku_STATUS instead
type VirtualNetworkGatewaySku_STATUS_ARM struct {
	Capacity *int                                  `json:"capacity,omitempty"`
	Name     *VirtualNetworkGatewaySku_Name_STATUS `json:"name,omitempty"`
	Tier     *VirtualNetworkGatewaySku_Tier_STATUS `json:"tier,omitempty"`
}

// Deprecated version of VpnClientConfiguration_STATUS. Use v1api20201101.VpnClientConfiguration_STATUS instead
type VpnClientConfiguration_STATUS_ARM struct {
	AadAudience                  *string                                                `json:"aadAudience,omitempty"`
	AadIssuer                    *string                                                `json:"aadIssuer,omitempty"`
	AadTenant                    *string                                                `json:"aadTenant,omitempty"`
	RadiusServerAddress          *string                                                `json:"radiusServerAddress,omitempty"`
	RadiusServerSecret           *string                                                `json:"radiusServerSecret,omitempty"`
	RadiusServers                []RadiusServer_STATUS_ARM                              `json:"radiusServers,omitempty"`
	VpnAuthenticationTypes       []VpnClientConfiguration_VpnAuthenticationTypes_STATUS `json:"vpnAuthenticationTypes,omitempty"`
	VpnClientAddressPool         *AddressSpace_STATUS_ARM                               `json:"vpnClientAddressPool,omitempty"`
	VpnClientIpsecPolicies       []IpsecPolicy_STATUS_ARM                               `json:"vpnClientIpsecPolicies,omitempty"`
	VpnClientProtocols           []VpnClientConfiguration_VpnClientProtocols_STATUS     `json:"vpnClientProtocols,omitempty"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificate_STATUS_ARM               `json:"vpnClientRevokedCertificates,omitempty"`
	VpnClientRootCertificates    []VpnClientRootCertificate_STATUS_ARM                  `json:"vpnClientRootCertificates,omitempty"`
}

// Deprecated version of IPConfigurationBgpPeeringAddress_STATUS. Use v1api20201101.IPConfigurationBgpPeeringAddress_STATUS instead
type IPConfigurationBgpPeeringAddress_STATUS_ARM struct {
	CustomBgpIpAddresses  []string `json:"customBgpIpAddresses,omitempty"`
	DefaultBgpIpAddresses []string `json:"defaultBgpIpAddresses,omitempty"`
	IpconfigurationId     *string  `json:"ipconfigurationId,omitempty"`
	TunnelIpAddresses     []string `json:"tunnelIpAddresses,omitempty"`
}

// Deprecated version of IpsecPolicy_STATUS. Use v1api20201101.IpsecPolicy_STATUS instead
type IpsecPolicy_STATUS_ARM struct {
	DhGroup             *DhGroup_STATUS         `json:"dhGroup,omitempty"`
	IkeEncryption       *IkeEncryption_STATUS   `json:"ikeEncryption,omitempty"`
	IkeIntegrity        *IkeIntegrity_STATUS    `json:"ikeIntegrity,omitempty"`
	IpsecEncryption     *IpsecEncryption_STATUS `json:"ipsecEncryption,omitempty"`
	IpsecIntegrity      *IpsecIntegrity_STATUS  `json:"ipsecIntegrity,omitempty"`
	PfsGroup            *PfsGroup_STATUS        `json:"pfsGroup,omitempty"`
	SaDataSizeKilobytes *int                    `json:"saDataSizeKilobytes,omitempty"`
	SaLifeTimeSeconds   *int                    `json:"saLifeTimeSeconds,omitempty"`
}

// Deprecated version of RadiusServer_STATUS. Use v1api20201101.RadiusServer_STATUS instead
type RadiusServer_STATUS_ARM struct {
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty"`
	RadiusServerScore   *int    `json:"radiusServerScore,omitempty"`
	RadiusServerSecret  *string `json:"radiusServerSecret,omitempty"`
}

// Deprecated version of VirtualNetworkGatewayIPConfigurationPropertiesFormat_STATUS. Use v1api20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormat_STATUS instead
type VirtualNetworkGatewayIPConfigurationPropertiesFormat_STATUS_ARM struct {
	PrivateIPAddress          *string                    `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *IPAllocationMethod_STATUS `json:"privateIPAllocationMethod,omitempty"`
	ProvisioningState         *ProvisioningState_STATUS  `json:"provisioningState,omitempty"`
	PublicIPAddress           *SubResource_STATUS_ARM    `json:"publicIPAddress,omitempty"`
	Subnet                    *SubResource_STATUS_ARM    `json:"subnet,omitempty"`
}

// Deprecated version of VpnClientRevokedCertificate_STATUS. Use v1api20201101.VpnClientRevokedCertificate_STATUS instead
type VpnClientRevokedCertificate_STATUS_ARM struct {
	Etag       *string                                                 `json:"etag,omitempty"`
	Id         *string                                                 `json:"id,omitempty"`
	Name       *string                                                 `json:"name,omitempty"`
	Properties *VpnClientRevokedCertificatePropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
}

// Deprecated version of VpnClientRootCertificate_STATUS. Use v1api20201101.VpnClientRootCertificate_STATUS instead
type VpnClientRootCertificate_STATUS_ARM struct {
	Etag       *string                                              `json:"etag,omitempty"`
	Id         *string                                              `json:"id,omitempty"`
	Name       *string                                              `json:"name,omitempty"`
	Properties *VpnClientRootCertificatePropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
}

// Deprecated version of VpnClientRevokedCertificatePropertiesFormat_STATUS. Use v1api20201101.VpnClientRevokedCertificatePropertiesFormat_STATUS instead
type VpnClientRevokedCertificatePropertiesFormat_STATUS_ARM struct {
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
	Thumbprint        *string                   `json:"thumbprint,omitempty"`
}

// Deprecated version of VpnClientRootCertificatePropertiesFormat_STATUS. Use v1api20201101.VpnClientRootCertificatePropertiesFormat_STATUS instead
type VpnClientRootCertificatePropertiesFormat_STATUS_ARM struct {
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
	PublicCertData    *string                   `json:"publicCertData,omitempty"`
}
