// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

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
//Storage version of v1alpha1api20201101.VirtualNetworksVirtualNetworkPeering
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks_virtualNetworkPeerings
type VirtualNetworksVirtualNetworkPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworksVirtualNetworkPeerings_Spec `json:"spec,omitempty"`
	Status            VirtualNetworkPeering_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualNetworksVirtualNetworkPeering{}

// GetConditions returns the conditions of the resource
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) GetConditions() conditions.Conditions {
	return virtualNetworksVirtualNetworkPeering.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) SetConditions(conditions conditions.Conditions) {
	virtualNetworksVirtualNetworkPeering.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &VirtualNetworksVirtualNetworkPeering{}

// AzureName returns the Azure name of the resource
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) AzureName() string {
	return virtualNetworksVirtualNetworkPeering.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (virtualNetworksVirtualNetworkPeering VirtualNetworksVirtualNetworkPeering) GetAPIVersion() string {
	return "2020-11-01"
}

// GetResourceKind returns the kind of the resource
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) GetSpec() genruntime.ConvertibleSpec {
	return &virtualNetworksVirtualNetworkPeering.Spec
}

// GetStatus returns the status of this resource
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) GetStatus() genruntime.ConvertibleStatus {
	return &virtualNetworksVirtualNetworkPeering.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) GetType() string {
	return "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
}

// NewEmptyStatus returns a new empty (blank) status
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &VirtualNetworkPeering_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(virtualNetworksVirtualNetworkPeering.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  virtualNetworksVirtualNetworkPeering.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*VirtualNetworkPeering_Status); ok {
		virtualNetworksVirtualNetworkPeering.Status = *st
		return nil
	}

	// Convert status to required version
	var st VirtualNetworkPeering_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	virtualNetworksVirtualNetworkPeering.Status = st
	return nil
}

// Hub marks that this VirtualNetworksVirtualNetworkPeering is the hub type for conversion
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (virtualNetworksVirtualNetworkPeering *VirtualNetworksVirtualNetworkPeering) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: virtualNetworksVirtualNetworkPeering.Spec.OriginalVersion,
		Kind:    "VirtualNetworksVirtualNetworkPeering",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.VirtualNetworksVirtualNetworkPeering
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks_virtualNetworkPeerings
type VirtualNetworksVirtualNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworksVirtualNetworkPeering `json:"items"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkPeering_Status
type VirtualNetworkPeering_Status struct {
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
	RemoteAddressSpace        *AddressSpace_Status                 `json:"remoteAddressSpace,omitempty"`
	RemoteBgpCommunities      *VirtualNetworkBgpCommunities_Status `json:"remoteBgpCommunities,omitempty"`
	RemoteVirtualNetwork      *SubResource_Status                  `json:"remoteVirtualNetwork,omitempty"`
	ResourceGuid              *string                              `json:"resourceGuid,omitempty"`
	Type                      *string                              `json:"type,omitempty"`
	UseRemoteGateways         *bool                                `json:"useRemoteGateways,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VirtualNetworkPeering_Status{}

// ConvertStatusFrom populates our VirtualNetworkPeering_Status from the provided source
func (virtualNetworkPeeringStatus *VirtualNetworkPeering_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == virtualNetworkPeeringStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(virtualNetworkPeeringStatus)
}

// ConvertStatusTo populates the provided destination from our VirtualNetworkPeering_Status
func (virtualNetworkPeeringStatus *VirtualNetworkPeering_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == virtualNetworkPeeringStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(virtualNetworkPeeringStatus)
}

//Storage version of v1alpha1api20201101.VirtualNetworksVirtualNetworkPeerings_Spec
type VirtualNetworksVirtualNetworkPeerings_Spec struct {
	AllowForwardedTraffic     *bool `json:"allowForwardedTraffic,omitempty"`
	AllowGatewayTransit       *bool `json:"allowGatewayTransit,omitempty"`
	AllowVirtualNetworkAccess *bool `json:"allowVirtualNetworkAccess,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string  `json:"azureName"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner                genruntime.KnownResourceReference `group:"network.azure.com" json:"owner" kind:"VirtualNetwork"`
	PeeringState         *string                           `json:"peeringState,omitempty"`
	PropertyBag          genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	RemoteAddressSpace   *AddressSpace                     `json:"remoteAddressSpace,omitempty"`
	RemoteBgpCommunities *VirtualNetworkBgpCommunities     `json:"remoteBgpCommunities,omitempty"`
	RemoteVirtualNetwork *SubResource                      `json:"remoteVirtualNetwork,omitempty"`
	Tags                 map[string]string                 `json:"tags,omitempty"`
	UseRemoteGateways    *bool                             `json:"useRemoteGateways,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualNetworksVirtualNetworkPeerings_Spec{}

// ConvertSpecFrom populates our VirtualNetworksVirtualNetworkPeerings_Spec from the provided source
func (virtualNetworksVirtualNetworkPeeringsSpec *VirtualNetworksVirtualNetworkPeerings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == virtualNetworksVirtualNetworkPeeringsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(virtualNetworksVirtualNetworkPeeringsSpec)
}

// ConvertSpecTo populates the provided destination from our VirtualNetworksVirtualNetworkPeerings_Spec
func (virtualNetworksVirtualNetworkPeeringsSpec *VirtualNetworksVirtualNetworkPeerings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == virtualNetworksVirtualNetworkPeeringsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(virtualNetworksVirtualNetworkPeeringsSpec)
}

func init() {
	SchemeBuilder.Register(&VirtualNetworksVirtualNetworkPeering{}, &VirtualNetworksVirtualNetworkPeeringList{})
}
