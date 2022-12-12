// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=virtualnetworksvirtualnetworkpeerings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={virtualnetworksvirtualnetworkpeerings/status,virtualnetworksvirtualnetworkpeerings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201101.VirtualNetworksVirtualNetworkPeering
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/virtualNetwork.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}
type VirtualNetworksVirtualNetworkPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworks_VirtualNetworkPeering_Spec   `json:"spec,omitempty"`
	Status            VirtualNetworks_VirtualNetworkPeering_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualNetworksVirtualNetworkPeering{}

// GetConditions returns the conditions of the resource
func (peering *VirtualNetworksVirtualNetworkPeering) GetConditions() conditions.Conditions {
	return peering.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (peering *VirtualNetworksVirtualNetworkPeering) SetConditions(conditions conditions.Conditions) {
	peering.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &VirtualNetworksVirtualNetworkPeering{}

// AzureName returns the Azure name of the resource
func (peering *VirtualNetworksVirtualNetworkPeering) AzureName() string {
	return peering.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (peering VirtualNetworksVirtualNetworkPeering) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (peering *VirtualNetworksVirtualNetworkPeering) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (peering *VirtualNetworksVirtualNetworkPeering) GetSpec() genruntime.ConvertibleSpec {
	return &peering.Spec
}

// GetStatus returns the status of this resource
func (peering *VirtualNetworksVirtualNetworkPeering) GetStatus() genruntime.ConvertibleStatus {
	return &peering.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
func (peering *VirtualNetworksVirtualNetworkPeering) GetType() string {
	return "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
}

// NewEmptyStatus returns a new empty (blank) status
func (peering *VirtualNetworksVirtualNetworkPeering) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &VirtualNetworks_VirtualNetworkPeering_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (peering *VirtualNetworksVirtualNetworkPeering) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(peering.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  peering.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (peering *VirtualNetworksVirtualNetworkPeering) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*VirtualNetworks_VirtualNetworkPeering_STATUS); ok {
		peering.Status = *st
		return nil
	}

	// Convert status to required version
	var st VirtualNetworks_VirtualNetworkPeering_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	peering.Status = st
	return nil
}

// Hub marks that this VirtualNetworksVirtualNetworkPeering is the hub type for conversion
func (peering *VirtualNetworksVirtualNetworkPeering) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (peering *VirtualNetworksVirtualNetworkPeering) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: peering.Spec.OriginalVersion,
		Kind:    "VirtualNetworksVirtualNetworkPeering",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201101.VirtualNetworksVirtualNetworkPeering
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/virtualNetwork.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}
type VirtualNetworksVirtualNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworksVirtualNetworkPeering `json:"items"`
}

// Storage version of v1beta20201101.VirtualNetworks_VirtualNetworkPeering_Spec
type VirtualNetworks_VirtualNetworkPeering_Spec struct {
	AllowForwardedTraffic     *bool `json:"allowForwardedTraffic,omitempty"`
	AllowGatewayTransit       *bool `json:"allowGatewayTransit,omitempty"`
	AllowVirtualNetworkAccess *bool `json:"allowVirtualNetworkAccess,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                 string `json:"azureName,omitempty"`
	DoNotVerifyRemoteGateways *bool  `json:"doNotVerifyRemoteGateways,omitempty"`
	OriginalVersion           string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/VirtualNetwork resource
	Owner                *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"VirtualNetwork"`
	PeeringState         *string                            `json:"peeringState,omitempty"`
	PropertyBag          genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RemoteAddressSpace   *AddressSpace                      `json:"remoteAddressSpace,omitempty"`
	RemoteBgpCommunities *VirtualNetworkBgpCommunities      `json:"remoteBgpCommunities,omitempty"`
	RemoteVirtualNetwork *SubResource                       `json:"remoteVirtualNetwork,omitempty"`
	UseRemoteGateways    *bool                              `json:"useRemoteGateways,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualNetworks_VirtualNetworkPeering_Spec{}

// ConvertSpecFrom populates our VirtualNetworks_VirtualNetworkPeering_Spec from the provided source
func (peering *VirtualNetworks_VirtualNetworkPeering_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == peering {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(peering)
}

// ConvertSpecTo populates the provided destination from our VirtualNetworks_VirtualNetworkPeering_Spec
func (peering *VirtualNetworks_VirtualNetworkPeering_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == peering {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(peering)
}

// Storage version of v1beta20201101.VirtualNetworks_VirtualNetworkPeering_STATUS
type VirtualNetworks_VirtualNetworkPeering_STATUS struct {
	AllowForwardedTraffic     *bool                                `json:"allowForwardedTraffic,omitempty"`
	AllowGatewayTransit       *bool                                `json:"allowGatewayTransit,omitempty"`
	AllowVirtualNetworkAccess *bool                                `json:"allowVirtualNetworkAccess,omitempty"`
	Conditions                []conditions.Condition               `json:"conditions,omitempty"`
	DoNotVerifyRemoteGateways *bool                                `json:"doNotVerifyRemoteGateways,omitempty"`
	Etag                      *string                              `json:"etag,omitempty"`
	Id                        *string                              `json:"id,omitempty"`
	Name                      *string                              `json:"name,omitempty"`
	PeeringState              *string                              `json:"peeringState,omitempty"`
	PropertyBag               genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                              `json:"provisioningState,omitempty"`
	RemoteAddressSpace        *AddressSpace_STATUS                 `json:"remoteAddressSpace,omitempty"`
	RemoteBgpCommunities      *VirtualNetworkBgpCommunities_STATUS `json:"remoteBgpCommunities,omitempty"`
	RemoteVirtualNetwork      *SubResource_STATUS                  `json:"remoteVirtualNetwork,omitempty"`
	ResourceGuid              *string                              `json:"resourceGuid,omitempty"`
	Type                      *string                              `json:"type,omitempty"`
	UseRemoteGateways         *bool                                `json:"useRemoteGateways,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VirtualNetworks_VirtualNetworkPeering_STATUS{}

// ConvertStatusFrom populates our VirtualNetworks_VirtualNetworkPeering_STATUS from the provided source
func (peering *VirtualNetworks_VirtualNetworkPeering_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == peering {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(peering)
}

// ConvertStatusTo populates the provided destination from our VirtualNetworks_VirtualNetworkPeering_STATUS
func (peering *VirtualNetworks_VirtualNetworkPeering_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == peering {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(peering)
}

func init() {
	SchemeBuilder.Register(&VirtualNetworksVirtualNetworkPeering{}, &VirtualNetworksVirtualNetworkPeeringList{})
}
