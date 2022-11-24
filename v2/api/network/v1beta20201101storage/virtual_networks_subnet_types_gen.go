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

// +kubebuilder:rbac:groups=network.azure.com,resources=virtualnetworkssubnets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={virtualnetworkssubnets/status,virtualnetworkssubnets/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201101.VirtualNetworksSubnet
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/virtualNetwork.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
type VirtualNetworksSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworks_Subnet_Spec   `json:"spec,omitempty"`
	Status            VirtualNetworks_Subnet_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualNetworksSubnet{}

// GetConditions returns the conditions of the resource
func (subnet *VirtualNetworksSubnet) GetConditions() conditions.Conditions {
	return subnet.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (subnet *VirtualNetworksSubnet) SetConditions(conditions conditions.Conditions) {
	subnet.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &VirtualNetworksSubnet{}

// AzureName returns the Azure name of the resource
func (subnet *VirtualNetworksSubnet) AzureName() string {
	return subnet.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (subnet VirtualNetworksSubnet) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (subnet *VirtualNetworksSubnet) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (subnet *VirtualNetworksSubnet) GetSpec() genruntime.ConvertibleSpec {
	return &subnet.Spec
}

// GetStatus returns the status of this resource
func (subnet *VirtualNetworksSubnet) GetStatus() genruntime.ConvertibleStatus {
	return &subnet.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/subnets"
func (subnet *VirtualNetworksSubnet) GetType() string {
	return "Microsoft.Network/virtualNetworks/subnets"
}

// NewEmptyStatus returns a new empty (blank) status
func (subnet *VirtualNetworksSubnet) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &VirtualNetworks_Subnet_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (subnet *VirtualNetworksSubnet) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(subnet.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  subnet.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (subnet *VirtualNetworksSubnet) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*VirtualNetworks_Subnet_STATUS); ok {
		subnet.Status = *st
		return nil
	}

	// Convert status to required version
	var st VirtualNetworks_Subnet_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	subnet.Status = st
	return nil
}

// Hub marks that this VirtualNetworksSubnet is the hub type for conversion
func (subnet *VirtualNetworksSubnet) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (subnet *VirtualNetworksSubnet) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: subnet.Spec.OriginalVersion,
		Kind:    "VirtualNetworksSubnet",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201101.VirtualNetworksSubnet
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/virtualNetwork.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
type VirtualNetworksSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworksSubnet `json:"items"`
}

// Storage version of v1beta20201101.VirtualNetworks_Subnet_Spec
type VirtualNetworks_Subnet_Spec struct {
	AddressPrefix                      *string                                                                        `json:"addressPrefix,omitempty"`
	AddressPrefixes                    []string                                                                       `json:"addressPrefixes,omitempty"`
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded `json:"applicationGatewayIpConfigurations,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName            string                                                               `json:"azureName,omitempty"`
	Delegations          []Delegation_VirtualNetworks_Subnet_SubResourceEmbedded              `json:"delegations,omitempty"`
	IpAllocations        []SubResource                                                        `json:"ipAllocations,omitempty"`
	NatGateway           *SubResource                                                         `json:"natGateway,omitempty"`
	NetworkSecurityGroup *NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded `json:"networkSecurityGroup,omitempty"`
	OriginalVersion      string                                                               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/VirtualNetwork resource
	Owner                             *genruntime.KnownResourceReference                                     `group:"network.azure.com" json:"owner,omitempty" kind:"VirtualNetwork"`
	PrivateEndpointNetworkPolicies    *string                                                                `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateLinkServiceNetworkPolicies *string                                                                `json:"privateLinkServiceNetworkPolicies,omitempty"`
	PropertyBag                       genruntime.PropertyBag                                                 `json:"$propertyBag,omitempty"`
	RouteTable                        *RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded             `json:"routeTable,omitempty"`
	ServiceEndpointPolicies           []ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                  []ServiceEndpointPropertiesFormat                                      `json:"serviceEndpoints,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualNetworks_Subnet_Spec{}

// ConvertSpecFrom populates our VirtualNetworks_Subnet_Spec from the provided source
func (subnet *VirtualNetworks_Subnet_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == subnet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(subnet)
}

// ConvertSpecTo populates the provided destination from our VirtualNetworks_Subnet_Spec
func (subnet *VirtualNetworks_Subnet_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == subnet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(subnet)
}

// Storage version of v1beta20201101.VirtualNetworks_Subnet_STATUS
type VirtualNetworks_Subnet_STATUS struct {
	AddressPrefix                      *string                                                                               `json:"addressPrefix,omitempty"`
	AddressPrefixes                    []string                                                                              `json:"addressPrefixes,omitempty"`
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded `json:"applicationGatewayIpConfigurations,omitempty"`
	Conditions                         []conditions.Condition                                                                `json:"conditions,omitempty"`
	Delegations                        []Delegation_STATUS                                                                   `json:"delegations,omitempty"`
	Etag                               *string                                                                               `json:"etag,omitempty"`
	Id                                 *string                                                                               `json:"id,omitempty"`
	IpAllocations                      []SubResource_STATUS                                                                  `json:"ipAllocations,omitempty"`
	IpConfigurationProfiles            []IPConfigurationProfile_STATUS                                                       `json:"ipConfigurationProfiles,omitempty"`
	IpConfigurations                   []IPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded                   `json:"ipConfigurations,omitempty"`
	Name                               *string                                                                               `json:"name,omitempty"`
	NatGateway                         *SubResource_STATUS                                                                   `json:"natGateway,omitempty"`
	NetworkSecurityGroup               *NetworkSecurityGroup_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded               `json:"networkSecurityGroup,omitempty"`
	PrivateEndpointNetworkPolicies     *string                                                                               `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateEndpoints                   []PrivateEndpoint_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded                   `json:"privateEndpoints,omitempty"`
	PrivateLinkServiceNetworkPolicies  *string                                                                               `json:"privateLinkServiceNetworkPolicies,omitempty"`
	PropertyBag                        genruntime.PropertyBag                                                                `json:"$propertyBag,omitempty"`
	ProvisioningState                  *string                                                                               `json:"provisioningState,omitempty"`
	Purpose                            *string                                                                               `json:"purpose,omitempty"`
	ResourceNavigationLinks            []ResourceNavigationLink_STATUS                                                       `json:"resourceNavigationLinks,omitempty"`
	RouteTable                         *RouteTable_STATUS_SubResourceEmbedded                                                `json:"routeTable,omitempty"`
	ServiceAssociationLinks            []ServiceAssociationLink_STATUS                                                       `json:"serviceAssociationLinks,omitempty"`
	ServiceEndpointPolicies            []ServiceEndpointPolicy_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded             `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                   []ServiceEndpointPropertiesFormat_STATUS                                              `json:"serviceEndpoints,omitempty"`
	Type                               *string                                                                               `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VirtualNetworks_Subnet_STATUS{}

// ConvertStatusFrom populates our VirtualNetworks_Subnet_STATUS from the provided source
func (subnet *VirtualNetworks_Subnet_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == subnet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(subnet)
}

// ConvertStatusTo populates the provided destination from our VirtualNetworks_Subnet_STATUS
func (subnet *VirtualNetworks_Subnet_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == subnet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(subnet)
}

// Storage version of v1beta20201101.ApplicationGatewayIPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded
type ApplicationGatewayIPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded
type ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20201101.Delegation_STATUS
type Delegation_STATUS struct {
	Actions           []string               `json:"actions,omitempty"`
	Etag              *string                `json:"etag,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	ServiceName       *string                `json:"serviceName,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

// Storage version of v1beta20201101.Delegation_VirtualNetworks_Subnet_SubResourceEmbedded
type Delegation_VirtualNetworks_Subnet_SubResourceEmbedded struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServiceName *string                `json:"serviceName,omitempty"`
}

// Storage version of v1beta20201101.IPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded
type IPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.IPConfigurationProfile_STATUS
type IPConfigurationProfile_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.NetworkSecurityGroup_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded
type NetworkSecurityGroup_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded
type NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20201101.PrivateEndpoint_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded
type PrivateEndpoint_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.ResourceNavigationLink_STATUS
type ResourceNavigationLink_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.RouteTable_STATUS_SubResourceEmbedded
type RouteTable_STATUS_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded
type RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20201101.ServiceAssociationLink_STATUS
type ServiceAssociationLink_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.ServiceEndpointPolicy_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded
type ServiceEndpointPolicy_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded
type ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20201101.ServiceEndpointPropertiesFormat
type ServiceEndpointPropertiesFormat struct {
	Locations   []string               `json:"locations,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Service     *string                `json:"service,omitempty"`
}

// Storage version of v1beta20201101.ServiceEndpointPropertiesFormat_STATUS
type ServiceEndpointPropertiesFormat_STATUS struct {
	Locations         []string               `json:"locations,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Service           *string                `json:"service,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VirtualNetworksSubnet{}, &VirtualNetworksSubnetList{})
}
