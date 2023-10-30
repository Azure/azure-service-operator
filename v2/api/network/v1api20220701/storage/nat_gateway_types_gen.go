// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=natgateways,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={natgateways/status,natgateways/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220701.NatGateway
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2022-07-01/natGateway.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}
type NatGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NatGateway_Spec   `json:"spec,omitempty"`
	Status            NatGateway_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NatGateway{}

// GetConditions returns the conditions of the resource
func (gateway *NatGateway) GetConditions() conditions.Conditions {
	return gateway.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (gateway *NatGateway) SetConditions(conditions conditions.Conditions) {
	gateway.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NatGateway{}

// AzureName returns the Azure name of the resource
func (gateway *NatGateway) AzureName() string {
	return gateway.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (gateway NatGateway) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (gateway *NatGateway) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (gateway *NatGateway) GetSpec() genruntime.ConvertibleSpec {
	return &gateway.Spec
}

// GetStatus returns the status of this resource
func (gateway *NatGateway) GetStatus() genruntime.ConvertibleStatus {
	return &gateway.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/natGateways"
func (gateway *NatGateway) GetType() string {
	return "Microsoft.Network/natGateways"
}

// NewEmptyStatus returns a new empty (blank) status
func (gateway *NatGateway) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NatGateway_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (gateway *NatGateway) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(gateway.Spec)
	return gateway.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (gateway *NatGateway) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NatGateway_STATUS); ok {
		gateway.Status = *st
		return nil
	}

	// Convert status to required version
	var st NatGateway_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	gateway.Status = st
	return nil
}

// Hub marks that this NatGateway is the hub type for conversion
func (gateway *NatGateway) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (gateway *NatGateway) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: gateway.Spec.OriginalVersion,
		Kind:    "NatGateway",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220701.NatGateway
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2022-07-01/natGateway.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}
type NatGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NatGateway `json:"items"`
}

// Storage version of v1api20220701.NatGateway_Spec
type NatGateway_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName            string  `json:"azureName,omitempty"`
	IdleTimeoutInMinutes *int    `json:"idleTimeoutInMinutes,omitempty"`
	Location             *string `json:"location,omitempty"`
	OriginalVersion      string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner             *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicIpAddresses []ApplicationGatewaySubResource    `json:"publicIpAddresses,omitempty"`
	PublicIpPrefixes  []ApplicationGatewaySubResource    `json:"publicIpPrefixes,omitempty"`
	Sku               *NatGatewaySku                     `json:"sku,omitempty"`
	Tags              map[string]string                  `json:"tags,omitempty"`
	Zones             []string                           `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NatGateway_Spec{}

// ConvertSpecFrom populates our NatGateway_Spec from the provided source
func (gateway *NatGateway_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == gateway {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(gateway)
}

// ConvertSpecTo populates the provided destination from our NatGateway_Spec
func (gateway *NatGateway_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == gateway {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(gateway)
}

// Storage version of v1api20220701.NatGateway_STATUS
// Nat Gateway resource.
type NatGateway_STATUS struct {
	Conditions           []conditions.Condition                 `json:"conditions,omitempty"`
	Etag                 *string                                `json:"etag,omitempty"`
	Id                   *string                                `json:"id,omitempty"`
	IdleTimeoutInMinutes *int                                   `json:"idleTimeoutInMinutes,omitempty"`
	Location             *string                                `json:"location,omitempty"`
	Name                 *string                                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	ProvisioningState    *string                                `json:"provisioningState,omitempty"`
	PublicIpAddresses    []ApplicationGatewaySubResource_STATUS `json:"publicIpAddresses,omitempty"`
	PublicIpPrefixes     []ApplicationGatewaySubResource_STATUS `json:"publicIpPrefixes,omitempty"`
	ResourceGuid         *string                                `json:"resourceGuid,omitempty"`
	Sku                  *NatGatewaySku_STATUS                  `json:"sku,omitempty"`
	Subnets              []ApplicationGatewaySubResource_STATUS `json:"subnets,omitempty"`
	Tags                 map[string]string                      `json:"tags,omitempty"`
	Type                 *string                                `json:"type,omitempty"`
	Zones                []string                               `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NatGateway_STATUS{}

// ConvertStatusFrom populates our NatGateway_STATUS from the provided source
func (gateway *NatGateway_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == gateway {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(gateway)
}

// ConvertStatusTo populates the provided destination from our NatGateway_STATUS
func (gateway *NatGateway_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == gateway {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(gateway)
}

// Storage version of v1api20220701.NatGatewaySku
// SKU of nat gateway.
type NatGatewaySku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220701.NatGatewaySku_STATUS
// SKU of nat gateway.
type NatGatewaySku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NatGateway{}, &NatGatewayList{})
}
