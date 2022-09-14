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

// +kubebuilder:rbac:groups=network.azure.com,resources=loadbalancers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={loadbalancers/status,loadbalancers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201101.LoadBalancer
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/loadBalancers
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancer_Spec   `json:"spec,omitempty"`
	Status            LoadBalancer_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &LoadBalancer{}

// GetConditions returns the conditions of the resource
func (balancer *LoadBalancer) GetConditions() conditions.Conditions {
	return balancer.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (balancer *LoadBalancer) SetConditions(conditions conditions.Conditions) {
	balancer.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &LoadBalancer{}

// AzureName returns the Azure name of the resource
func (balancer *LoadBalancer) AzureName() string {
	return balancer.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (balancer LoadBalancer) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (balancer *LoadBalancer) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (balancer *LoadBalancer) GetSpec() genruntime.ConvertibleSpec {
	return &balancer.Spec
}

// GetStatus returns the status of this resource
func (balancer *LoadBalancer) GetStatus() genruntime.ConvertibleStatus {
	return &balancer.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/loadBalancers"
func (balancer *LoadBalancer) GetType() string {
	return "Microsoft.Network/loadBalancers"
}

// NewEmptyStatus returns a new empty (blank) status
func (balancer *LoadBalancer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &LoadBalancer_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (balancer *LoadBalancer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(balancer.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  balancer.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (balancer *LoadBalancer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*LoadBalancer_STATUS); ok {
		balancer.Status = *st
		return nil
	}

	// Convert status to required version
	var st LoadBalancer_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	balancer.Status = st
	return nil
}

// Hub marks that this LoadBalancer is the hub type for conversion
func (balancer *LoadBalancer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (balancer *LoadBalancer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: balancer.Spec.OriginalVersion,
		Kind:    "LoadBalancer",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201101.LoadBalancer
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/loadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

// Storage version of v1beta20201101.APIVersion
// +kubebuilder:validation:Enum={"2020-11-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-11-01")

// Storage version of v1beta20201101.LoadBalancer_Spec
type LoadBalancer_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                string                                                  `json:"azureName,omitempty"`
	BackendAddressPools      []LoadBalancer_Properties_BackendAddressPools_Spec      `json:"backendAddressPools,omitempty"`
	ExtendedLocation         *ExtendedLocation                                       `json:"extendedLocation,omitempty"`
	FrontendIPConfigurations []LoadBalancer_Properties_FrontendIPConfigurations_Spec `json:"frontendIPConfigurations,omitempty"`
	InboundNatPools          []LoadBalancer_Properties_InboundNatPools_Spec          `json:"inboundNatPools,omitempty"`
	LoadBalancingRules       []LoadBalancer_Properties_LoadBalancingRules_Spec       `json:"loadBalancingRules,omitempty"`
	Location                 *string                                                 `json:"location,omitempty"`
	OriginalVersion          string                                                  `json:"originalVersion,omitempty"`
	OutboundRules            []LoadBalancer_Properties_OutboundRules_Spec            `json:"outboundRules,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference    `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	Probes      []LoadBalancer_Properties_Probes_Spec `json:"probes,omitempty"`
	PropertyBag genruntime.PropertyBag                `json:"$propertyBag,omitempty"`
	Sku         *LoadBalancerSku                      `json:"sku,omitempty"`
	Tags        map[string]string                     `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &LoadBalancer_Spec{}

// ConvertSpecFrom populates our LoadBalancer_Spec from the provided source
func (balancer *LoadBalancer_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == balancer {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(balancer)
}

// ConvertSpecTo populates the provided destination from our LoadBalancer_Spec
func (balancer *LoadBalancer_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == balancer {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(balancer)
}

// Storage version of v1beta20201101.LoadBalancer_STATUS
type LoadBalancer_STATUS struct {
	BackendAddressPools      []BackendAddressPool_STATUS_LoadBalancer_SubResourceEmbedded      `json:"backendAddressPools,omitempty"`
	Conditions               []conditions.Condition                                            `json:"conditions,omitempty"`
	Etag                     *string                                                           `json:"etag,omitempty"`
	ExtendedLocation         *ExtendedLocation_STATUS                                          `json:"extendedLocation,omitempty"`
	FrontendIPConfigurations []FrontendIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded `json:"frontendIPConfigurations,omitempty"`
	Id                       *string                                                           `json:"id,omitempty"`
	InboundNatPools          []InboundNatPool_STATUS                                           `json:"inboundNatPools,omitempty"`
	InboundNatRules          []InboundNatRule_STATUS_LoadBalancer_SubResourceEmbedded          `json:"inboundNatRules,omitempty"`
	LoadBalancingRules       []LoadBalancingRule_STATUS                                        `json:"loadBalancingRules,omitempty"`
	Location                 *string                                                           `json:"location,omitempty"`
	Name                     *string                                                           `json:"name,omitempty"`
	OutboundRules            []OutboundRule_STATUS                                             `json:"outboundRules,omitempty"`
	Probes                   []Probe_STATUS                                                    `json:"probes,omitempty"`
	PropertyBag              genruntime.PropertyBag                                            `json:"$propertyBag,omitempty"`
	ProvisioningState        *string                                                           `json:"provisioningState,omitempty"`
	ResourceGuid             *string                                                           `json:"resourceGuid,omitempty"`
	Sku                      *LoadBalancerSku_STATUS                                           `json:"sku,omitempty"`
	Tags                     map[string]string                                                 `json:"tags,omitempty"`
	Type                     *string                                                           `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &LoadBalancer_STATUS{}

// ConvertStatusFrom populates our LoadBalancer_STATUS from the provided source
func (balancer *LoadBalancer_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == balancer {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(balancer)
}

// ConvertStatusTo populates the provided destination from our LoadBalancer_STATUS
func (balancer *LoadBalancer_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == balancer {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(balancer)
}

// Storage version of v1beta20201101.BackendAddressPool_STATUS_LoadBalancer_SubResourceEmbedded
type BackendAddressPool_STATUS_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.ExtendedLocation
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ExtendedLocation
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20201101.ExtendedLocation_STATUS
type ExtendedLocation_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20201101.FrontendIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded
type FrontendIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded struct {
	Etag                      *string                                                  `json:"etag,omitempty"`
	Id                        *string                                                  `json:"id,omitempty"`
	InboundNatPools           []SubResource_STATUS                                     `json:"inboundNatPools,omitempty"`
	InboundNatRules           []SubResource_STATUS                                     `json:"inboundNatRules,omitempty"`
	LoadBalancingRules        []SubResource_STATUS                                     `json:"loadBalancingRules,omitempty"`
	Name                      *string                                                  `json:"name,omitempty"`
	OutboundRules             []SubResource_STATUS                                     `json:"outboundRules,omitempty"`
	PrivateIPAddress          *string                                                  `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *string                                                  `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *string                                                  `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag                                   `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                                                  `json:"provisioningState,omitempty"`
	PublicIPAddress           *PublicIPAddress_STATUS_LoadBalancer_SubResourceEmbedded `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *SubResource_STATUS                                      `json:"publicIPPrefix,omitempty"`
	Subnet                    *Subnet_STATUS_LoadBalancer_SubResourceEmbedded          `json:"subnet,omitempty"`
	Type                      *string                                                  `json:"type,omitempty"`
	Zones                     []string                                                 `json:"zones,omitempty"`
}

// Storage version of v1beta20201101.InboundNatPool_STATUS
type InboundNatPool_STATUS struct {
	BackendPort             *int                   `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	Etag                    *string                `json:"etag,omitempty"`
	FrontendIPConfiguration *SubResource_STATUS    `json:"frontendIPConfiguration,omitempty"`
	FrontendPortRangeEnd    *int                   `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int                   `json:"frontendPortRangeStart,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
	ProvisioningState       *string                `json:"provisioningState,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

// Storage version of v1beta20201101.InboundNatRule_STATUS_LoadBalancer_SubResourceEmbedded
type InboundNatRule_STATUS_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.LoadBalancer_Properties_BackendAddressPools_Spec
type LoadBalancer_Properties_BackendAddressPools_Spec struct {
	LoadBalancerBackendAddresses []LoadBalancer_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses_Spec `json:"loadBalancerBackendAddresses,omitempty"`
	Location                     *string                                                                                    `json:"location,omitempty"`
	Name                         *string                                                                                    `json:"name,omitempty"`
	PropertyBag                  genruntime.PropertyBag                                                                     `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.LoadBalancer_Properties_FrontendIPConfigurations_Spec
type LoadBalancer_Properties_FrontendIPConfigurations_Spec struct {
	Name                      *string                `json:"name,omitempty"`
	PrivateIPAddress          *string                `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *string                `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *string                `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicIPAddress           *SubResource           `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *SubResource           `json:"publicIPPrefix,omitempty"`
	Subnet                    *SubResource           `json:"subnet,omitempty"`
	Zones                     []string               `json:"zones,omitempty"`
}

// Storage version of v1beta20201101.LoadBalancer_Properties_InboundNatPools_Spec
type LoadBalancer_Properties_InboundNatPools_Spec struct {
	BackendPort             *int                   `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResource           `json:"frontendIPConfiguration,omitempty"`
	FrontendPortRangeEnd    *int                   `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int                   `json:"frontendPortRangeStart,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
}

// Storage version of v1beta20201101.LoadBalancer_Properties_LoadBalancingRules_Spec
type LoadBalancer_Properties_LoadBalancingRules_Spec struct {
	BackendAddressPool      *SubResource           `json:"backendAddressPool,omitempty"`
	BackendPort             *int                   `json:"backendPort,omitempty"`
	DisableOutboundSnat     *bool                  `json:"disableOutboundSnat,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResource           `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                   `json:"frontendPort,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	LoadDistribution        *string                `json:"loadDistribution,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	Probe                   *SubResource           `json:"probe,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
}

// Storage version of v1beta20201101.LoadBalancer_Properties_OutboundRules_Spec
type LoadBalancer_Properties_OutboundRules_Spec struct {
	AllocatedOutboundPorts   *int                   `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       *SubResource           `json:"backendAddressPool,omitempty"`
	EnableTcpReset           *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfigurations []SubResource          `json:"frontendIPConfigurations,omitempty"`
	IdleTimeoutInMinutes     *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                 *string                `json:"protocol,omitempty"`
}

// Storage version of v1beta20201101.LoadBalancer_Properties_Probes_Spec
type LoadBalancer_Properties_Probes_Spec struct {
	IntervalInSeconds *int                   `json:"intervalInSeconds,omitempty"`
	Name              *string                `json:"name,omitempty"`
	NumberOfProbes    *int                   `json:"numberOfProbes,omitempty"`
	Port              *int                   `json:"port,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol          *string                `json:"protocol,omitempty"`
	RequestPath       *string                `json:"requestPath,omitempty"`
}

// Storage version of v1beta20201101.LoadBalancerSku
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/LoadBalancerSku
type LoadBalancerSku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20201101.LoadBalancerSku_STATUS
type LoadBalancerSku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20201101.LoadBalancingRule_STATUS
type LoadBalancingRule_STATUS struct {
	BackendAddressPool      *SubResource_STATUS    `json:"backendAddressPool,omitempty"`
	BackendPort             *int                   `json:"backendPort,omitempty"`
	DisableOutboundSnat     *bool                  `json:"disableOutboundSnat,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	Etag                    *string                `json:"etag,omitempty"`
	FrontendIPConfiguration *SubResource_STATUS    `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                   `json:"frontendPort,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	LoadDistribution        *string                `json:"loadDistribution,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	Probe                   *SubResource_STATUS    `json:"probe,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
	ProvisioningState       *string                `json:"provisioningState,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

// Storage version of v1beta20201101.OutboundRule_STATUS
type OutboundRule_STATUS struct {
	AllocatedOutboundPorts   *int                   `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       *SubResource_STATUS    `json:"backendAddressPool,omitempty"`
	EnableTcpReset           *bool                  `json:"enableTcpReset,omitempty"`
	Etag                     *string                `json:"etag,omitempty"`
	FrontendIPConfigurations []SubResource_STATUS   `json:"frontendIPConfigurations,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes     *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                 *string                `json:"protocol,omitempty"`
	ProvisioningState        *string                `json:"provisioningState,omitempty"`
	Type                     *string                `json:"type,omitempty"`
}

// Storage version of v1beta20201101.Probe_STATUS
type Probe_STATUS struct {
	Etag               *string                `json:"etag,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	IntervalInSeconds  *int                   `json:"intervalInSeconds,omitempty"`
	LoadBalancingRules []SubResource_STATUS   `json:"loadBalancingRules,omitempty"`
	Name               *string                `json:"name,omitempty"`
	NumberOfProbes     *int                   `json:"numberOfProbes,omitempty"`
	Port               *int                   `json:"port,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol           *string                `json:"protocol,omitempty"`
	ProvisioningState  *string                `json:"provisioningState,omitempty"`
	RequestPath        *string                `json:"requestPath,omitempty"`
	Type               *string                `json:"type,omitempty"`
}

// Storage version of v1beta20201101.LoadBalancer_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses_Spec
type LoadBalancer_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses_Spec struct {
	IpAddress                           *string                `json:"ipAddress,omitempty"`
	LoadBalancerFrontendIPConfiguration *SubResource           `json:"loadBalancerFrontendIPConfiguration,omitempty"`
	Name                                *string                `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subnet                              *SubResource           `json:"subnet,omitempty"`
	VirtualNetwork                      *SubResource           `json:"virtualNetwork,omitempty"`
}

// Storage version of v1beta20201101.PublicIPAddress_STATUS_LoadBalancer_SubResourceEmbedded
type PublicIPAddress_STATUS_LoadBalancer_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_STATUS   `json:"extendedLocation,omitempty"`
	Id               *string                    `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Sku              *PublicIPAddressSku_STATUS `json:"sku,omitempty"`
	Zones            []string                   `json:"zones,omitempty"`
}

// Storage version of v1beta20201101.Subnet_STATUS_LoadBalancer_SubResourceEmbedded
type Subnet_STATUS_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
