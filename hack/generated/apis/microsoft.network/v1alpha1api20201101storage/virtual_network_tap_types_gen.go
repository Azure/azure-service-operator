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
//Storage version of v1alpha1api20201101.VirtualNetworkTap
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworkTaps
type VirtualNetworkTap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkTaps_Spec                                        `json:"spec,omitempty"`
	Status            VirtualNetworkTap_Status_VirtualNetworkTap_SubResourceEmbedded `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualNetworkTap{}

// GetConditions returns the conditions of the resource
func (virtualNetworkTap *VirtualNetworkTap) GetConditions() conditions.Conditions {
	return virtualNetworkTap.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (virtualNetworkTap *VirtualNetworkTap) SetConditions(conditions conditions.Conditions) {
	virtualNetworkTap.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &VirtualNetworkTap{}

// AzureName returns the Azure name of the resource
func (virtualNetworkTap *VirtualNetworkTap) AzureName() string {
	return virtualNetworkTap.Spec.AzureName
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (virtualNetworkTap *VirtualNetworkTap) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(virtualNetworkTap.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: virtualNetworkTap.Namespace, Name: virtualNetworkTap.Spec.Owner.Name}
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (virtualNetworkTap *VirtualNetworkTap) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: virtualNetworkTap.Spec.OriginalVersion,
		Kind:    "VirtualNetworkTap",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.VirtualNetworkTap
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworkTaps
type VirtualNetworkTapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkTap `json:"items"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkTap_Status_VirtualNetworkTap_SubResourceEmbedded
//Generated from:
type VirtualNetworkTap_Status_VirtualNetworkTap_SubResourceEmbedded struct {
	Conditions                                     []conditions.Condition                                                          `json:"conditions,omitempty"`
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfiguration_Status_VirtualNetworkTap_SubResourceEmbedded           `json:"destinationLoadBalancerFrontEndIPConfiguration,omitempty"`
	DestinationNetworkInterfaceIPConfiguration     *NetworkInterfaceIPConfiguration_Status_VirtualNetworkTap_SubResourceEmbedded   `json:"destinationNetworkInterfaceIPConfiguration,omitempty"`
	DestinationPort                                *int                                                                            `json:"destinationPort,omitempty"`
	Etag                                           *string                                                                         `json:"etag,omitempty"`
	Id                                             *string                                                                         `json:"id,omitempty"`
	Location                                       *string                                                                         `json:"location,omitempty"`
	Name                                           *string                                                                         `json:"name,omitempty"`
	NetworkInterfaceTapConfigurations              []NetworkInterfaceTapConfiguration_Status_VirtualNetworkTap_SubResourceEmbedded `json:"networkInterfaceTapConfigurations,omitempty"`
	PropertyBag                                    genruntime.PropertyBag                                                          `json:"$propertyBag,omitempty"`
	ProvisioningState                              *string                                                                         `json:"provisioningState,omitempty"`
	ResourceGuid                                   *string                                                                         `json:"resourceGuid,omitempty"`
	Tags                                           map[string]string                                                               `json:"tags,omitempty"`
	Type                                           *string                                                                         `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VirtualNetworkTap_Status_VirtualNetworkTap_SubResourceEmbedded{}

// ConvertStatusFrom populates our VirtualNetworkTap_Status_VirtualNetworkTap_SubResourceEmbedded from the provided source
func (virtualNetworkTapStatusVirtualNetworkTapSubResourceEmbedded *VirtualNetworkTap_Status_VirtualNetworkTap_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == virtualNetworkTapStatusVirtualNetworkTapSubResourceEmbedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(virtualNetworkTapStatusVirtualNetworkTapSubResourceEmbedded)
}

// ConvertStatusTo populates the provided destination from our VirtualNetworkTap_Status_VirtualNetworkTap_SubResourceEmbedded
func (virtualNetworkTapStatusVirtualNetworkTapSubResourceEmbedded *VirtualNetworkTap_Status_VirtualNetworkTap_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == virtualNetworkTapStatusVirtualNetworkTapSubResourceEmbedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(virtualNetworkTapStatusVirtualNetworkTapSubResourceEmbedded)
}

//Storage version of v1alpha1api20201101.VirtualNetworkTaps_Spec
type VirtualNetworkTaps_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                                      string       `json:"azureName"`
	DestinationLoadBalancerFrontEndIPConfiguration *SubResource `json:"destinationLoadBalancerFrontEndIPConfiguration,omitempty"`
	DestinationNetworkInterfaceIPConfiguration     *SubResource `json:"destinationNetworkInterfaceIPConfiguration,omitempty"`
	DestinationPort                                *int         `json:"destinationPort,omitempty"`
	Location                                       *string      `json:"location,omitempty"`
	OriginalVersion                                string       `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"microsoft.resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualNetworkTaps_Spec{}

// ConvertSpecFrom populates our VirtualNetworkTaps_Spec from the provided source
func (virtualNetworkTapsSpec *VirtualNetworkTaps_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == virtualNetworkTapsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(virtualNetworkTapsSpec)
}

// ConvertSpecTo populates the provided destination from our VirtualNetworkTaps_Spec
func (virtualNetworkTapsSpec *VirtualNetworkTaps_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == virtualNetworkTapsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(virtualNetworkTapsSpec)
}

//Storage version of v1alpha1api20201101.FrontendIPConfiguration_Status_VirtualNetworkTap_SubResourceEmbedded
//Generated from:
type FrontendIPConfiguration_Status_VirtualNetworkTap_SubResourceEmbedded struct {
	Etag                      *string                                                       `json:"etag,omitempty"`
	Id                        *string                                                       `json:"id,omitempty"`
	InboundNatPools           []SubResource_Status                                          `json:"inboundNatPools,omitempty"`
	InboundNatRules           []SubResource_Status                                          `json:"inboundNatRules,omitempty"`
	LoadBalancingRules        []SubResource_Status                                          `json:"loadBalancingRules,omitempty"`
	Name                      *string                                                       `json:"name,omitempty"`
	OutboundRules             []SubResource_Status                                          `json:"outboundRules,omitempty"`
	PrivateIPAddress          *string                                                       `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *string                                                       `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *string                                                       `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag                                        `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                                                       `json:"provisioningState,omitempty"`
	PublicIPAddress           *PublicIPAddress_Status_VirtualNetworkTap_SubResourceEmbedded `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *SubResource_Status                                           `json:"publicIPPrefix,omitempty"`
	Subnet                    *Subnet_Status_VirtualNetworkTap_SubResourceEmbedded          `json:"subnet,omitempty"`
	Type                      *string                                                       `json:"type,omitempty"`
	Zones                     []string                                                      `json:"zones,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkInterfaceIPConfiguration_Status_VirtualNetworkTap_SubResourceEmbedded
//Generated from:
type NetworkInterfaceIPConfiguration_Status_VirtualNetworkTap_SubResourceEmbedded struct {
	ApplicationGatewayBackendAddressPools []ApplicationGatewayBackendAddressPool_Status_VirtualNetworkTap_SubResourceEmbedded `json:"applicationGatewayBackendAddressPools,omitempty"`
	ApplicationSecurityGroups             []ApplicationSecurityGroup_Status_VirtualNetworkTap_SubResourceEmbedded             `json:"applicationSecurityGroups,omitempty"`
	Etag                                  *string                                                                             `json:"etag,omitempty"`
	Id                                    *string                                                                             `json:"id,omitempty"`
	LoadBalancerBackendAddressPools       []BackendAddressPool_Status_VirtualNetworkTap_SubResourceEmbedded                   `json:"loadBalancerBackendAddressPools,omitempty"`
	LoadBalancerInboundNatRules           []InboundNatRule_Status_VirtualNetworkTap_SubResourceEmbedded                       `json:"loadBalancerInboundNatRules,omitempty"`
	Name                                  *string                                                                             `json:"name,omitempty"`
	Primary                               *bool                                                                               `json:"primary,omitempty"`
	PrivateIPAddress                      *string                                                                             `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion               *string                                                                             `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod             *string                                                                             `json:"privateIPAllocationMethod,omitempty"`
	PrivateLinkConnectionProperties       *NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status              `json:"privateLinkConnectionProperties,omitempty"`
	PropertyBag                           genruntime.PropertyBag                                                              `json:"$propertyBag,omitempty"`
	ProvisioningState                     *string                                                                             `json:"provisioningState,omitempty"`
	PublicIPAddress                       *PublicIPAddress_Status_VirtualNetworkTap_SubResourceEmbedded                       `json:"publicIPAddress,omitempty"`
	Subnet                                *Subnet_Status_VirtualNetworkTap_SubResourceEmbedded                                `json:"subnet,omitempty"`
	Type                                  *string                                                                             `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkInterfaceTapConfiguration_Status_VirtualNetworkTap_SubResourceEmbedded
//Generated from:
type NetworkInterfaceTapConfiguration_Status_VirtualNetworkTap_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.ApplicationGatewayBackendAddressPool_Status_VirtualNetworkTap_SubResourceEmbedded
//Generated from:
type ApplicationGatewayBackendAddressPool_Status_VirtualNetworkTap_SubResourceEmbedded struct {
	BackendAddresses  []ApplicationGatewayBackendAddress_Status `json:"backendAddresses,omitempty"`
	Etag              *string                                   `json:"etag,omitempty"`
	Id                *string                                   `json:"id,omitempty"`
	Name              *string                                   `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	ProvisioningState *string                                   `json:"provisioningState,omitempty"`
	Type              *string                                   `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.ApplicationSecurityGroup_Status_VirtualNetworkTap_SubResourceEmbedded
//Generated from:
type ApplicationSecurityGroup_Status_VirtualNetworkTap_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.BackendAddressPool_Status_VirtualNetworkTap_SubResourceEmbedded
//Generated from:
type BackendAddressPool_Status_VirtualNetworkTap_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.InboundNatRule_Status_VirtualNetworkTap_SubResourceEmbedded
//Generated from:
type InboundNatRule_Status_VirtualNetworkTap_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status
//Generated from:
type NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status struct {
	Fqdns              []string               `json:"fqdns,omitempty"`
	GroupId            *string                `json:"groupId,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequiredMemberName *string                `json:"requiredMemberName,omitempty"`
}

//Storage version of v1alpha1api20201101.PublicIPAddress_Status_VirtualNetworkTap_SubResourceEmbedded
//Generated from:
type PublicIPAddress_Status_VirtualNetworkTap_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_Status   `json:"extendedLocation,omitempty"`
	Id               *string                    `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Sku              *PublicIPAddressSku_Status `json:"sku,omitempty"`
	Zones            []string                   `json:"zones,omitempty"`
}

//Storage version of v1alpha1api20201101.Subnet_Status_VirtualNetworkTap_SubResourceEmbedded
//Generated from:
type Subnet_Status_VirtualNetworkTap_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.ApplicationGatewayBackendAddress_Status
//Generated from:
type ApplicationGatewayBackendAddress_Status struct {
	Fqdn        *string                `json:"fqdn,omitempty"`
	IpAddress   *string                `json:"ipAddress,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VirtualNetworkTap{}, &VirtualNetworkTapList{})
}
