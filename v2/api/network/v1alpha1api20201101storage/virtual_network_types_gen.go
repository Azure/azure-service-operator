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

// +kubebuilder:rbac:groups=network.azure.com,resources=virtualnetworks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={virtualnetworks/status,virtualnetworks/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.VirtualNetwork
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks
type VirtualNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworks_Spec  `json:"spec,omitempty"`
	Status            VirtualNetwork_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualNetwork{}

// GetConditions returns the conditions of the resource
func (network *VirtualNetwork) GetConditions() conditions.Conditions {
	return network.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (network *VirtualNetwork) SetConditions(conditions conditions.Conditions) {
	network.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &VirtualNetwork{}

// AzureName returns the Azure name of the resource
func (network *VirtualNetwork) AzureName() string {
	return network.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (network VirtualNetwork) GetAPIVersion() string {
	return "2020-11-01"
}

// GetResourceKind returns the kind of the resource
func (network *VirtualNetwork) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (network *VirtualNetwork) GetSpec() genruntime.ConvertibleSpec {
	return &network.Spec
}

// GetStatus returns the status of this resource
func (network *VirtualNetwork) GetStatus() genruntime.ConvertibleStatus {
	return &network.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks"
func (network *VirtualNetwork) GetType() string {
	return "Microsoft.Network/virtualNetworks"
}

// NewEmptyStatus returns a new empty (blank) status
func (network *VirtualNetwork) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &VirtualNetwork_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (network *VirtualNetwork) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(network.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  network.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (network *VirtualNetwork) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*VirtualNetwork_Status); ok {
		network.Status = *st
		return nil
	}

	// Convert status to required version
	var st VirtualNetwork_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	network.Status = st
	return nil
}

// Hub marks that this VirtualNetwork is the hub type for conversion
func (network *VirtualNetwork) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (network *VirtualNetwork) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: network.Spec.OriginalVersion,
		Kind:    "VirtualNetwork",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.VirtualNetwork
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks
type VirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetwork `json:"items"`
}

//Storage version of v1alpha1api20201101.VirtualNetwork_Status
type VirtualNetwork_Status struct {
	AddressSpace           *AddressSpace_Status                               `json:"addressSpace,omitempty"`
	BgpCommunities         *VirtualNetworkBgpCommunities_Status               `json:"bgpCommunities,omitempty"`
	Conditions             []conditions.Condition                             `json:"conditions,omitempty"`
	DdosProtectionPlan     *SubResource_Status                                `json:"ddosProtectionPlan,omitempty"`
	DhcpOptions            *DhcpOptions_Status                                `json:"dhcpOptions,omitempty"`
	EnableDdosProtection   *bool                                              `json:"enableDdosProtection,omitempty"`
	EnableVmProtection     *bool                                              `json:"enableVmProtection,omitempty"`
	Etag                   *string                                            `json:"etag,omitempty"`
	ExtendedLocation       *ExtendedLocation_Status                           `json:"extendedLocation,omitempty"`
	Id                     *string                                            `json:"id,omitempty"`
	IpAllocations          []SubResource_Status                               `json:"ipAllocations,omitempty"`
	Location               *string                                            `json:"location,omitempty"`
	Name                   *string                                            `json:"name,omitempty"`
	PropertyBag            genruntime.PropertyBag                             `json:"$propertyBag,omitempty"`
	ProvisioningState      *string                                            `json:"provisioningState,omitempty"`
	ResourceGuid           *string                                            `json:"resourceGuid,omitempty"`
	Subnets                []Subnet_Status_VirtualNetwork_SubResourceEmbedded `json:"subnets,omitempty"`
	Tags                   map[string]string                                  `json:"tags,omitempty"`
	Type                   *string                                            `json:"type,omitempty"`
	VirtualNetworkPeerings []VirtualNetworkPeering_Status_SubResourceEmbedded `json:"virtualNetworkPeerings,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VirtualNetwork_Status{}

// ConvertStatusFrom populates our VirtualNetwork_Status from the provided source
func (network *VirtualNetwork_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == network {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(network)
}

// ConvertStatusTo populates the provided destination from our VirtualNetwork_Status
func (network *VirtualNetwork_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == network {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(network)
}

//Storage version of v1alpha1api20201101.VirtualNetworks_Spec
type VirtualNetworks_Spec struct {
	AddressSpace *AddressSpace `json:"addressSpace,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName            string                        `json:"azureName"`
	BgpCommunities       *VirtualNetworkBgpCommunities `json:"bgpCommunities,omitempty"`
	DdosProtectionPlan   *SubResource                  `json:"ddosProtectionPlan,omitempty"`
	DhcpOptions          *DhcpOptions                  `json:"dhcpOptions,omitempty"`
	EnableDdosProtection *bool                         `json:"enableDdosProtection,omitempty"`
	EnableVmProtection   *bool                         `json:"enableVmProtection,omitempty"`
	ExtendedLocation     *ExtendedLocation             `json:"extendedLocation,omitempty"`
	IpAllocations        []SubResource                 `json:"ipAllocations,omitempty"`
	Location             *string                       `json:"location,omitempty"`
	OriginalVersion      string                        `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference         `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	Subnets     []VirtualNetworks_Spec_Properties_Subnets `json:"subnets,omitempty"`
	Tags        map[string]string                         `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualNetworks_Spec{}

// ConvertSpecFrom populates our VirtualNetworks_Spec from the provided source
func (networks *VirtualNetworks_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == networks {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(networks)
}

// ConvertSpecTo populates the provided destination from our VirtualNetworks_Spec
func (networks *VirtualNetworks_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == networks {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(networks)
}

//Storage version of v1alpha1api20201101.AddressSpace
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/AddressSpace
type AddressSpace struct {
	AddressPrefixes []string               `json:"addressPrefixes,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.AddressSpace_Status
type AddressSpace_Status struct {
	AddressPrefixes []string               `json:"addressPrefixes,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.DhcpOptions
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DhcpOptions
type DhcpOptions struct {
	DnsServers  []string               `json:"dnsServers,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.DhcpOptions_Status
type DhcpOptions_Status struct {
	DnsServers  []string               `json:"dnsServers,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.Subnet_Status_VirtualNetwork_SubResourceEmbedded
type Subnet_Status_VirtualNetwork_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkBgpCommunities
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkBgpCommunities
type VirtualNetworkBgpCommunities struct {
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VirtualNetworkCommunity *string                `json:"virtualNetworkCommunity,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkBgpCommunities_Status
type VirtualNetworkBgpCommunities_Status struct {
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RegionalCommunity       *string                `json:"regionalCommunity,omitempty"`
	VirtualNetworkCommunity *string                `json:"virtualNetworkCommunity,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkPeering_Status_SubResourceEmbedded
type VirtualNetworkPeering_Status_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworks_Spec_Properties_Subnets
type VirtualNetworks_Spec_Properties_Subnets struct {
	AddressPrefix                     *string                                                          `json:"addressPrefix,omitempty"`
	AddressPrefixes                   []string                                                         `json:"addressPrefixes,omitempty"`
	Delegations                       []VirtualNetworks_Spec_Properties_Subnets_Properties_Delegations `json:"delegations,omitempty"`
	IpAllocations                     []SubResource                                                    `json:"ipAllocations,omitempty"`
	Name                              *string                                                          `json:"name,omitempty"`
	NatGateway                        *SubResource                                                     `json:"natGateway,omitempty"`
	NetworkSecurityGroup              *SubResource                                                     `json:"networkSecurityGroup,omitempty"`
	PrivateEndpointNetworkPolicies    *string                                                          `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateLinkServiceNetworkPolicies *string                                                          `json:"privateLinkServiceNetworkPolicies,omitempty"`
	PropertyBag                       genruntime.PropertyBag                                           `json:"$propertyBag,omitempty"`
	RouteTable                        *SubResource                                                     `json:"routeTable,omitempty"`
	ServiceEndpointPolicies           []SubResource                                                    `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                  []ServiceEndpointPropertiesFormat                                `json:"serviceEndpoints,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworks_Spec_Properties_Subnets_Properties_Delegations
type VirtualNetworks_Spec_Properties_Subnets_Properties_Delegations struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServiceName *string                `json:"serviceName,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VirtualNetwork{}, &VirtualNetworkList{})
}
