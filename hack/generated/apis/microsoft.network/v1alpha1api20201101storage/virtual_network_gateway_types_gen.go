// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

import (
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.VirtualNetworkGateway
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworkGateways
type VirtualNetworkGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGateways_Spec  `json:"spec,omitempty"`
	Status            VirtualNetworkGateway_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualNetworkGateway{}

// GetConditions returns the conditions of the resource
func (virtualNetworkGateway *VirtualNetworkGateway) GetConditions() conditions.Conditions {
	return virtualNetworkGateway.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (virtualNetworkGateway *VirtualNetworkGateway) SetConditions(conditions conditions.Conditions) {
	virtualNetworkGateway.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &VirtualNetworkGateway{}

// AzureName returns the Azure name of the resource
func (virtualNetworkGateway *VirtualNetworkGateway) AzureName() string {
	return virtualNetworkGateway.Spec.AzureName
}

// GetSpec returns the specification of this resource
func (virtualNetworkGateway *VirtualNetworkGateway) GetSpec() genruntime.ConvertibleSpec {
	return &virtualNetworkGateway.Spec
}

// GetStatus returns the status of this resource
func (virtualNetworkGateway *VirtualNetworkGateway) GetStatus() genruntime.ConvertibleStatus {
	return &virtualNetworkGateway.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworkGateways"
func (virtualNetworkGateway *VirtualNetworkGateway) GetType() string {
	return "Microsoft.Network/virtualNetworkGateways"
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (virtualNetworkGateway *VirtualNetworkGateway) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(virtualNetworkGateway.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: virtualNetworkGateway.Namespace, Name: virtualNetworkGateway.Spec.Owner.Name}
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (virtualNetworkGateway *VirtualNetworkGateway) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: virtualNetworkGateway.Spec.OriginalVersion,
		Kind:    "VirtualNetworkGateway",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.VirtualNetworkGateway
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworkGateways
type VirtualNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkGateway `json:"items"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkGateway_Status
//Generated from:
type VirtualNetworkGateway_Status struct {
	ActiveActive                   *bool                                         `json:"activeActive,omitempty"`
	BgpSettings                    *BgpSettings_Status                           `json:"bgpSettings,omitempty"`
	Conditions                     []conditions.Condition                        `json:"conditions,omitempty"`
	CustomRoutes                   *AddressSpace_Status                          `json:"customRoutes,omitempty"`
	EnableBgp                      *bool                                         `json:"enableBgp,omitempty"`
	EnableDnsForwarding            *bool                                         `json:"enableDnsForwarding,omitempty"`
	EnablePrivateIpAddress         *bool                                         `json:"enablePrivateIpAddress,omitempty"`
	Etag                           *string                                       `json:"etag,omitempty"`
	ExtendedLocation               *ExtendedLocation_Status                      `json:"extendedLocation,omitempty"`
	GatewayDefaultSite             *SubResource_Status                           `json:"gatewayDefaultSite,omitempty"`
	GatewayType                    *string                                       `json:"gatewayType,omitempty"`
	Id                             *string                                       `json:"id,omitempty"`
	InboundDnsForwardingEndpoint   *string                                       `json:"inboundDnsForwardingEndpoint,omitempty"`
	IpConfigurations               []VirtualNetworkGatewayIPConfiguration_Status `json:"ipConfigurations,omitempty"`
	Location                       *string                                       `json:"location,omitempty"`
	Name                           *string                                       `json:"name,omitempty"`
	PropertyBag                    genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	ProvisioningState              *string                                       `json:"provisioningState,omitempty"`
	ResourceGuid                   *string                                       `json:"resourceGuid,omitempty"`
	Sku                            *VirtualNetworkGatewaySku_Status              `json:"sku,omitempty"`
	Tags                           map[string]string                             `json:"tags,omitempty"`
	Type                           *string                                       `json:"type,omitempty"`
	VNetExtendedLocationResourceId *string                                       `json:"vNetExtendedLocationResourceId,omitempty"`
	VpnClientConfiguration         *VpnClientConfiguration_Status                `json:"vpnClientConfiguration,omitempty"`
	VpnGatewayGeneration           *string                                       `json:"vpnGatewayGeneration,omitempty"`
	VpnType                        *string                                       `json:"vpnType,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VirtualNetworkGateway_Status{}

// ConvertStatusFrom populates our VirtualNetworkGateway_Status from the provided source
func (virtualNetworkGatewayStatus *VirtualNetworkGateway_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == virtualNetworkGatewayStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(virtualNetworkGatewayStatus)
}

// ConvertStatusTo populates the provided destination from our VirtualNetworkGateway_Status
func (virtualNetworkGatewayStatus *VirtualNetworkGateway_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == virtualNetworkGatewayStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(virtualNetworkGatewayStatus)
}

//Storage version of v1alpha1api20201101.VirtualNetworkGateways_Spec
type VirtualNetworkGateways_Spec struct {
	ActiveActive *bool `json:"activeActive,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName              string                                                    `json:"azureName"`
	BgpSettings            *BgpSettings                                              `json:"bgpSettings,omitempty"`
	CustomRoutes           *AddressSpace                                             `json:"customRoutes,omitempty"`
	EnableBgp              *bool                                                     `json:"enableBgp,omitempty"`
	EnableDnsForwarding    *bool                                                     `json:"enableDnsForwarding,omitempty"`
	EnablePrivateIpAddress *bool                                                     `json:"enablePrivateIpAddress,omitempty"`
	GatewayDefaultSite     *SubResource                                              `json:"gatewayDefaultSite,omitempty"`
	GatewayType            *string                                                   `json:"gatewayType,omitempty"`
	IpConfigurations       []VirtualNetworkGateways_Spec_Properties_IpConfigurations `json:"ipConfigurations,omitempty"`
	Location               *string                                                   `json:"location,omitempty"`
	OriginalVersion        string                                                    `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"microsoft.resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Sku         *VirtualNetworkGatewaySku         `json:"sku,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`

	//VNetExtendedLocationResourceReference: MAS FIJI customer vnet resource id.
	//VirtualNetworkGateway of type local gateway is associated with the customer vnet.
	VNetExtendedLocationResourceReference *genruntime.ResourceReference                                  `armReference:"VNetExtendedLocationResourceId" json:"vNetExtendedLocationResourceReference,omitempty"`
	VirtualNetworkExtendedLocation        *ExtendedLocation                                              `json:"virtualNetworkExtendedLocation,omitempty"`
	VpnClientConfiguration                *VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration `json:"vpnClientConfiguration,omitempty"`
	VpnGatewayGeneration                  *string                                                        `json:"vpnGatewayGeneration,omitempty"`
	VpnType                               *string                                                        `json:"vpnType,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualNetworkGateways_Spec{}

// ConvertSpecFrom populates our VirtualNetworkGateways_Spec from the provided source
func (virtualNetworkGatewaysSpec *VirtualNetworkGateways_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == virtualNetworkGatewaysSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(virtualNetworkGatewaysSpec)
}

// ConvertSpecTo populates the provided destination from our VirtualNetworkGateways_Spec
func (virtualNetworkGatewaysSpec *VirtualNetworkGateways_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == virtualNetworkGatewaysSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(virtualNetworkGatewaysSpec)
}

//Storage version of v1alpha1api20201101.BgpSettings
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/BgpSettings
type BgpSettings struct {
	Asn                 *uint32                            `json:"asn,omitempty"`
	BgpPeeringAddress   *string                            `json:"bgpPeeringAddress,omitempty"`
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddress `json:"bgpPeeringAddresses,omitempty"`
	PeerWeight          *int                               `json:"peerWeight,omitempty"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.BgpSettings_Status
//Generated from:
type BgpSettings_Status struct {
	Asn                 *uint32                                   `json:"asn,omitempty"`
	BgpPeeringAddress   *string                                   `json:"bgpPeeringAddress,omitempty"`
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddress_Status `json:"bgpPeeringAddresses,omitempty"`
	PeerWeight          *int                                      `json:"peerWeight,omitempty"`
	PropertyBag         genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkGatewayIPConfiguration_Status
//Generated from:
type VirtualNetworkGatewayIPConfiguration_Status struct {
	Etag                      *string                `json:"etag,omitempty"`
	Id                        *string                `json:"id,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	PrivateIPAddress          *string                `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *string                `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                `json:"provisioningState,omitempty"`
	PublicIPAddress           *SubResource_Status    `json:"publicIPAddress,omitempty"`
	Subnet                    *SubResource_Status    `json:"subnet,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkGatewaySku
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkGatewaySku
type VirtualNetworkGatewaySku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkGatewaySku_Status
//Generated from:
type VirtualNetworkGatewaySku_Status struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkGateways_Spec_Properties_IpConfigurations
type VirtualNetworkGateways_Spec_Properties_IpConfigurations struct {
	Name                      *string                `json:"name,omitempty"`
	PrivateIPAllocationMethod *string                `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicIPAddress           *SubResource           `json:"publicIPAddress,omitempty"`
	Subnet                    *SubResource           `json:"subnet,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration
type VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration struct {
	AadAudience                  *string                                                                                      `json:"aadAudience,omitempty"`
	AadIssuer                    *string                                                                                      `json:"aadIssuer,omitempty"`
	AadTenant                    *string                                                                                      `json:"aadTenant,omitempty"`
	PropertyBag                  genruntime.PropertyBag                                                                       `json:"$propertyBag,omitempty"`
	RadiusServerAddress          *string                                                                                      `json:"radiusServerAddress,omitempty"`
	RadiusServerSecret           *string                                                                                      `json:"radiusServerSecret,omitempty"`
	RadiusServers                []RadiusServer                                                                               `json:"radiusServers,omitempty"`
	VpnAuthenticationTypes       []string                                                                                     `json:"vpnAuthenticationTypes,omitempty"`
	VpnClientAddressPool         *AddressSpace                                                                                `json:"vpnClientAddressPool,omitempty"`
	VpnClientIpsecPolicies       []IpsecPolicy                                                                                `json:"vpnClientIpsecPolicies,omitempty"`
	VpnClientProtocols           []string                                                                                     `json:"vpnClientProtocols,omitempty"`
	VpnClientRevokedCertificates []VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificates `json:"vpnClientRevokedCertificates,omitempty"`
	VpnClientRootCertificates    []VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificates    `json:"vpnClientRootCertificates,omitempty"`
}

//Storage version of v1alpha1api20201101.VpnClientConfiguration_Status
//Generated from:
type VpnClientConfiguration_Status struct {
	AadAudience                  *string                              `json:"aadAudience,omitempty"`
	AadIssuer                    *string                              `json:"aadIssuer,omitempty"`
	AadTenant                    *string                              `json:"aadTenant,omitempty"`
	PropertyBag                  genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	RadiusServerAddress          *string                              `json:"radiusServerAddress,omitempty"`
	RadiusServerSecret           *string                              `json:"radiusServerSecret,omitempty"`
	RadiusServers                []RadiusServer_Status                `json:"radiusServers,omitempty"`
	VpnAuthenticationTypes       []string                             `json:"vpnAuthenticationTypes,omitempty"`
	VpnClientAddressPool         *AddressSpace_Status                 `json:"vpnClientAddressPool,omitempty"`
	VpnClientIpsecPolicies       []IpsecPolicy_Status                 `json:"vpnClientIpsecPolicies,omitempty"`
	VpnClientProtocols           []string                             `json:"vpnClientProtocols,omitempty"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificate_Status `json:"vpnClientRevokedCertificates,omitempty"`
	VpnClientRootCertificates    []VpnClientRootCertificate_Status    `json:"vpnClientRootCertificates,omitempty"`
}

//Storage version of v1alpha1api20201101.IPConfigurationBgpPeeringAddress
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IPConfigurationBgpPeeringAddress
type IPConfigurationBgpPeeringAddress struct {
	CustomBgpIpAddresses []string               `json:"customBgpIpAddresses,omitempty"`
	IpconfigurationId    *string                `json:"ipconfigurationId,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.IPConfigurationBgpPeeringAddress_Status
//Generated from:
type IPConfigurationBgpPeeringAddress_Status struct {
	CustomBgpIpAddresses  []string               `json:"customBgpIpAddresses,omitempty"`
	DefaultBgpIpAddresses []string               `json:"defaultBgpIpAddresses,omitempty"`
	IpconfigurationId     *string                `json:"ipconfigurationId,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TunnelIpAddresses     []string               `json:"tunnelIpAddresses,omitempty"`
}

//Storage version of v1alpha1api20201101.IpsecPolicy
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpsecPolicy
type IpsecPolicy struct {
	DhGroup             *string                `json:"dhGroup,omitempty"`
	IkeEncryption       *string                `json:"ikeEncryption,omitempty"`
	IkeIntegrity        *string                `json:"ikeIntegrity,omitempty"`
	IpsecEncryption     *string                `json:"ipsecEncryption,omitempty"`
	IpsecIntegrity      *string                `json:"ipsecIntegrity,omitempty"`
	PfsGroup            *string                `json:"pfsGroup,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SaDataSizeKilobytes *int                   `json:"saDataSizeKilobytes,omitempty"`
	SaLifeTimeSeconds   *int                   `json:"saLifeTimeSeconds,omitempty"`
}

//Storage version of v1alpha1api20201101.IpsecPolicy_Status
//Generated from:
type IpsecPolicy_Status struct {
	DhGroup             *string                `json:"dhGroup,omitempty"`
	IkeEncryption       *string                `json:"ikeEncryption,omitempty"`
	IkeIntegrity        *string                `json:"ikeIntegrity,omitempty"`
	IpsecEncryption     *string                `json:"ipsecEncryption,omitempty"`
	IpsecIntegrity      *string                `json:"ipsecIntegrity,omitempty"`
	PfsGroup            *string                `json:"pfsGroup,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SaDataSizeKilobytes *int                   `json:"saDataSizeKilobytes,omitempty"`
	SaLifeTimeSeconds   *int                   `json:"saLifeTimeSeconds,omitempty"`
}

//Storage version of v1alpha1api20201101.RadiusServer
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/RadiusServer
type RadiusServer struct {
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RadiusServerAddress *string                `json:"radiusServerAddress,omitempty"`
	RadiusServerScore   *int                   `json:"radiusServerScore,omitempty"`
	RadiusServerSecret  *string                `json:"radiusServerSecret,omitempty"`
}

//Storage version of v1alpha1api20201101.RadiusServer_Status
//Generated from:
type RadiusServer_Status struct {
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RadiusServerAddress *string                `json:"radiusServerAddress,omitempty"`
	RadiusServerScore   *int                   `json:"radiusServerScore,omitempty"`
	RadiusServerSecret  *string                `json:"radiusServerSecret,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificates
type VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificates struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Thumbprint  *string                `json:"thumbprint,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificates
type VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificates struct {
	Name           *string                `json:"name,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicCertData *string                `json:"publicCertData,omitempty"`
}

//Storage version of v1alpha1api20201101.VpnClientRevokedCertificate_Status
//Generated from:
type VpnClientRevokedCertificate_Status struct {
	Etag              *string                `json:"etag,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Thumbprint        *string                `json:"thumbprint,omitempty"`
}

//Storage version of v1alpha1api20201101.VpnClientRootCertificate_Status
//Generated from:
type VpnClientRootCertificate_Status struct {
	Etag              *string                `json:"etag,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	PublicCertData    *string                `json:"publicCertData,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VirtualNetworkGateway{}, &VirtualNetworkGatewayList{})
}
