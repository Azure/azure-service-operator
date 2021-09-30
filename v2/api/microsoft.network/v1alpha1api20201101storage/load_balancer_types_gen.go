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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.LoadBalancer
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/loadBalancers
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancers_Spec  `json:"spec,omitempty"`
	Status            LoadBalancer_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &LoadBalancer{}

// GetConditions returns the conditions of the resource
func (loadBalancer *LoadBalancer) GetConditions() conditions.Conditions {
	return loadBalancer.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (loadBalancer *LoadBalancer) SetConditions(conditions conditions.Conditions) {
	loadBalancer.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &LoadBalancer{}

// AzureName returns the Azure name of the resource
func (loadBalancer *LoadBalancer) AzureName() string {
	return loadBalancer.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (loadBalancer *LoadBalancer) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (loadBalancer *LoadBalancer) GetSpec() genruntime.ConvertibleSpec {
	return &loadBalancer.Spec
}

// GetStatus returns the status of this resource
func (loadBalancer *LoadBalancer) GetStatus() genruntime.ConvertibleStatus {
	return &loadBalancer.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/loadBalancers"
func (loadBalancer *LoadBalancer) GetType() string {
	return "Microsoft.Network/loadBalancers"
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (loadBalancer *LoadBalancer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(loadBalancer.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: loadBalancer.Namespace, Name: loadBalancer.Spec.Owner.Name}
}

// SetStatus sets the status of this resource
func (loadBalancer *LoadBalancer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*LoadBalancer_Status); ok {
		loadBalancer.Status = *st
		return nil
	}

	// Convert status to required version
	var st LoadBalancer_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	loadBalancer.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (loadBalancer *LoadBalancer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: loadBalancer.Spec.OriginalVersion,
		Kind:    "LoadBalancer",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.LoadBalancer
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/loadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

//Storage version of v1alpha1api20201101.LoadBalancer_Status
//Generated from:
type LoadBalancer_Status struct {
	BackendAddressPools      []BackendAddressPool_Status_LoadBalancer_SubResourceEmbedded      `json:"backendAddressPools,omitempty"`
	Conditions               []conditions.Condition                                            `json:"conditions,omitempty"`
	Etag                     *string                                                           `json:"etag,omitempty"`
	ExtendedLocation         *ExtendedLocation_Status                                          `json:"extendedLocation,omitempty"`
	FrontendIPConfigurations []FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded `json:"frontendIPConfigurations,omitempty"`
	Id                       *string                                                           `json:"id,omitempty"`
	InboundNatPools          []InboundNatPool_Status                                           `json:"inboundNatPools,omitempty"`
	InboundNatRules          []InboundNatRule_Status_LoadBalancer_SubResourceEmbedded          `json:"inboundNatRules,omitempty"`
	LoadBalancingRules       []LoadBalancingRule_Status                                        `json:"loadBalancingRules,omitempty"`
	Location                 *string                                                           `json:"location,omitempty"`
	Name                     *string                                                           `json:"name,omitempty"`
	OutboundRules            []OutboundRule_Status                                             `json:"outboundRules,omitempty"`
	Probes                   []Probe_Status                                                    `json:"probes,omitempty"`
	PropertyBag              genruntime.PropertyBag                                            `json:"$propertyBag,omitempty"`
	ProvisioningState        *string                                                           `json:"provisioningState,omitempty"`
	ResourceGuid             *string                                                           `json:"resourceGuid,omitempty"`
	Sku                      *LoadBalancerSku_Status                                           `json:"sku,omitempty"`
	Tags                     map[string]string                                                 `json:"tags,omitempty"`
	Type                     *string                                                           `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &LoadBalancer_Status{}

// ConvertStatusFrom populates our LoadBalancer_Status from the provided source
func (loadBalancerStatus *LoadBalancer_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == loadBalancerStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(loadBalancerStatus)
}

// ConvertStatusTo populates the provided destination from our LoadBalancer_Status
func (loadBalancerStatus *LoadBalancer_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == loadBalancerStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(loadBalancerStatus)
}

//Storage version of v1alpha1api20201101.LoadBalancers_Spec
type LoadBalancers_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                string                                                   `json:"azureName"`
	BackendAddressPools      []LoadBalancers_Spec_Properties_BackendAddressPools      `json:"backendAddressPools,omitempty"`
	ExtendedLocation         *ExtendedLocation                                        `json:"extendedLocation,omitempty"`
	FrontendIPConfigurations []LoadBalancers_Spec_Properties_FrontendIPConfigurations `json:"frontendIPConfigurations,omitempty"`
	InboundNatPools          []LoadBalancers_Spec_Properties_InboundNatPools          `json:"inboundNatPools,omitempty"`
	LoadBalancingRules       []LoadBalancers_Spec_Properties_LoadBalancingRules       `json:"loadBalancingRules,omitempty"`
	Location                 *string                                                  `json:"location,omitempty"`
	OriginalVersion          string                                                   `json:"originalVersion"`
	OutboundRules            []LoadBalancers_Spec_Properties_OutboundRules            `json:"outboundRules,omitempty"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference      `group:"microsoft.resources.azure.com" json:"owner" kind:"ResourceGroup"`
	Probes      []LoadBalancers_Spec_Properties_Probes `json:"probes,omitempty"`
	PropertyBag genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	Sku         *LoadBalancerSku                       `json:"sku,omitempty"`
	Tags        map[string]string                      `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &LoadBalancers_Spec{}

// ConvertSpecFrom populates our LoadBalancers_Spec from the provided source
func (loadBalancersSpec *LoadBalancers_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == loadBalancersSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(loadBalancersSpec)
}

// ConvertSpecTo populates the provided destination from our LoadBalancers_Spec
func (loadBalancersSpec *LoadBalancers_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == loadBalancersSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(loadBalancersSpec)
}

//Storage version of v1alpha1api20201101.BackendAddressPool_Status_LoadBalancer_SubResourceEmbedded
//Generated from:
type BackendAddressPool_Status_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.ExtendedLocation
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ExtendedLocation
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.ExtendedLocation_Status
//Generated from:
type ExtendedLocation_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded
//Generated from:
type FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded struct {
	Etag                      *string                                                  `json:"etag,omitempty"`
	Id                        *string                                                  `json:"id,omitempty"`
	InboundNatPools           []SubResource_Status                                     `json:"inboundNatPools,omitempty"`
	InboundNatRules           []SubResource_Status                                     `json:"inboundNatRules,omitempty"`
	LoadBalancingRules        []SubResource_Status                                     `json:"loadBalancingRules,omitempty"`
	Name                      *string                                                  `json:"name,omitempty"`
	OutboundRules             []SubResource_Status                                     `json:"outboundRules,omitempty"`
	PrivateIPAddress          *string                                                  `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *string                                                  `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *string                                                  `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag                                   `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                                                  `json:"provisioningState,omitempty"`
	PublicIPAddress           *PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *SubResource_Status                                      `json:"publicIPPrefix,omitempty"`
	Subnet                    *Subnet_Status_LoadBalancer_SubResourceEmbedded          `json:"subnet,omitempty"`
	Type                      *string                                                  `json:"type,omitempty"`
	Zones                     []string                                                 `json:"zones,omitempty"`
}

//Storage version of v1alpha1api20201101.InboundNatPool_Status
//Generated from:
type InboundNatPool_Status struct {
	BackendPort             *int                   `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	Etag                    *string                `json:"etag,omitempty"`
	FrontendIPConfiguration *SubResource_Status    `json:"frontendIPConfiguration,omitempty"`
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

//Storage version of v1alpha1api20201101.InboundNatRule_Status_LoadBalancer_SubResourceEmbedded
//Generated from:
type InboundNatRule_Status_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancerSku
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/LoadBalancerSku
type LoadBalancerSku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancerSku_Status
//Generated from:
type LoadBalancerSku_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancers_Spec_Properties_BackendAddressPools
type LoadBalancers_Spec_Properties_BackendAddressPools struct {
	LoadBalancerBackendAddresses []LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses `json:"loadBalancerBackendAddresses,omitempty"`
	Location                     *string                                                                                     `json:"location,omitempty"`
	Name                         *string                                                                                     `json:"name,omitempty"`
	PropertyBag                  genruntime.PropertyBag                                                                      `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurations
type LoadBalancers_Spec_Properties_FrontendIPConfigurations struct {
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

//Storage version of v1alpha1api20201101.LoadBalancers_Spec_Properties_InboundNatPools
type LoadBalancers_Spec_Properties_InboundNatPools struct {
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

//Storage version of v1alpha1api20201101.LoadBalancers_Spec_Properties_LoadBalancingRules
type LoadBalancers_Spec_Properties_LoadBalancingRules struct {
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

//Storage version of v1alpha1api20201101.LoadBalancers_Spec_Properties_OutboundRules
type LoadBalancers_Spec_Properties_OutboundRules struct {
	AllocatedOutboundPorts   *int                   `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       *SubResource           `json:"backendAddressPool,omitempty"`
	EnableTcpReset           *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfigurations []SubResource          `json:"frontendIPConfigurations,omitempty"`
	IdleTimeoutInMinutes     *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                 *string                `json:"protocol,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancers_Spec_Properties_Probes
type LoadBalancers_Spec_Properties_Probes struct {
	IntervalInSeconds *int                   `json:"intervalInSeconds,omitempty"`
	Name              *string                `json:"name,omitempty"`
	NumberOfProbes    *int                   `json:"numberOfProbes,omitempty"`
	Port              *int                   `json:"port,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol          *string                `json:"protocol,omitempty"`
	RequestPath       *string                `json:"requestPath,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancingRule_Status
//Generated from:
type LoadBalancingRule_Status struct {
	BackendAddressPool      *SubResource_Status    `json:"backendAddressPool,omitempty"`
	BackendPort             *int                   `json:"backendPort,omitempty"`
	DisableOutboundSnat     *bool                  `json:"disableOutboundSnat,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	Etag                    *string                `json:"etag,omitempty"`
	FrontendIPConfiguration *SubResource_Status    `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                   `json:"frontendPort,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	LoadDistribution        *string                `json:"loadDistribution,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	Probe                   *SubResource_Status    `json:"probe,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
	ProvisioningState       *string                `json:"provisioningState,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.OutboundRule_Status
//Generated from:
type OutboundRule_Status struct {
	AllocatedOutboundPorts   *int                   `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       *SubResource_Status    `json:"backendAddressPool,omitempty"`
	EnableTcpReset           *bool                  `json:"enableTcpReset,omitempty"`
	Etag                     *string                `json:"etag,omitempty"`
	FrontendIPConfigurations []SubResource_Status   `json:"frontendIPConfigurations,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes     *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                 *string                `json:"protocol,omitempty"`
	ProvisioningState        *string                `json:"provisioningState,omitempty"`
	Type                     *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.Probe_Status
//Generated from:
type Probe_Status struct {
	Etag               *string                `json:"etag,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	IntervalInSeconds  *int                   `json:"intervalInSeconds,omitempty"`
	LoadBalancingRules []SubResource_Status   `json:"loadBalancingRules,omitempty"`
	Name               *string                `json:"name,omitempty"`
	NumberOfProbes     *int                   `json:"numberOfProbes,omitempty"`
	Port               *int                   `json:"port,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol           *string                `json:"protocol,omitempty"`
	ProvisioningState  *string                `json:"provisioningState,omitempty"`
	RequestPath        *string                `json:"requestPath,omitempty"`
	Type               *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses
type LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses struct {
	IpAddress                           *string                `json:"ipAddress,omitempty"`
	LoadBalancerFrontendIPConfiguration *SubResource           `json:"loadBalancerFrontendIPConfiguration,omitempty"`
	Name                                *string                `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subnet                              *SubResource           `json:"subnet,omitempty"`
	VirtualNetwork                      *SubResource           `json:"virtualNetwork,omitempty"`
}

//Storage version of v1alpha1api20201101.PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded
//Generated from:
type PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_Status   `json:"extendedLocation,omitempty"`
	Id               *string                    `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Sku              *PublicIPAddressSku_Status `json:"sku,omitempty"`
	Zones            []string                   `json:"zones,omitempty"`
}

//Storage version of v1alpha1api20201101.Subnet_Status_LoadBalancer_SubResourceEmbedded
//Generated from:
type Subnet_Status_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
