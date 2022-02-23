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

// +kubebuilder:rbac:groups=network.azure.com,resources=publicipaddresses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={publicipaddresses/status,publicipaddresses/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.PublicIPAddress
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/publicIPAddresses
type PublicIPAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicIPAddresses_Spec                                     `json:"spec,omitempty"`
	Status            PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PublicIPAddress{}

// GetConditions returns the conditions of the resource
func (address *PublicIPAddress) GetConditions() conditions.Conditions {
	return address.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (address *PublicIPAddress) SetConditions(conditions conditions.Conditions) {
	address.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &PublicIPAddress{}

// AzureName returns the Azure name of the resource
func (address *PublicIPAddress) AzureName() string {
	return address.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (address PublicIPAddress) GetAPIVersion() string {
	return "2020-11-01"
}

// GetResourceKind returns the kind of the resource
func (address *PublicIPAddress) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (address *PublicIPAddress) GetSpec() genruntime.ConvertibleSpec {
	return &address.Spec
}

// GetStatus returns the status of this resource
func (address *PublicIPAddress) GetStatus() genruntime.ConvertibleStatus {
	return &address.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/publicIPAddresses"
func (address *PublicIPAddress) GetType() string {
	return "Microsoft.Network/publicIPAddresses"
}

// NewEmptyStatus returns a new empty (blank) status
func (address *PublicIPAddress) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (address *PublicIPAddress) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(address.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  address.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (address *PublicIPAddress) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded); ok {
		address.Status = *st
		return nil
	}

	// Convert status to required version
	var st PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	address.Status = st
	return nil
}

// Hub marks that this PublicIPAddress is the hub type for conversion
func (address *PublicIPAddress) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (address *PublicIPAddress) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: address.Spec.OriginalVersion,
		Kind:    "PublicIPAddress",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.PublicIPAddress
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/publicIPAddresses
type PublicIPAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicIPAddress `json:"items"`
}

//Storage version of v1alpha1api20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded
type PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded struct {
	Conditions               []conditions.Condition                                      `json:"conditions,omitempty"`
	DdosSettings             *DdosSettings_Status                                        `json:"ddosSettings,omitempty"`
	DnsSettings              *PublicIPAddressDnsSettings_Status                          `json:"dnsSettings,omitempty"`
	Etag                     *string                                                     `json:"etag,omitempty"`
	ExtendedLocation         *ExtendedLocation_Status                                    `json:"extendedLocation,omitempty"`
	Id                       *string                                                     `json:"id,omitempty"`
	IdleTimeoutInMinutes     *int                                                        `json:"idleTimeoutInMinutes,omitempty"`
	IpAddress                *string                                                     `json:"ipAddress,omitempty"`
	IpConfiguration          *IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded `json:"ipConfiguration,omitempty"`
	IpTags                   []IpTag_Status                                              `json:"ipTags,omitempty"`
	Location                 *string                                                     `json:"location,omitempty"`
	MigrationPhase           *string                                                     `json:"migrationPhase,omitempty"`
	Name                     *string                                                     `json:"name,omitempty"`
	NatGateway               *NatGateway_Status_PublicIPAddress_SubResourceEmbedded      `json:"natGateway,omitempty"`
	PropertyBag              genruntime.PropertyBag                                      `json:"$propertyBag,omitempty"`
	ProvisioningState        *string                                                     `json:"provisioningState,omitempty"`
	PublicIPAddressVersion   *string                                                     `json:"publicIPAddressVersion,omitempty"`
	PublicIPAllocationMethod *string                                                     `json:"publicIPAllocationMethod,omitempty"`
	PublicIPPrefix           *SubResource_Status                                         `json:"publicIPPrefix,omitempty"`
	ResourceGuid             *string                                                     `json:"resourceGuid,omitempty"`
	Sku                      *PublicIPAddressSku_Status                                  `json:"sku,omitempty"`
	Tags                     map[string]string                                           `json:"tags,omitempty"`
	Type                     *string                                                     `json:"type,omitempty"`
	Zones                    []string                                                    `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded{}

// ConvertStatusFrom populates our PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded from the provided source
func (embedded *PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == embedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(embedded)
}

// ConvertStatusTo populates the provided destination from our PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded
func (embedded *PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == embedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(embedded)
}

//Storage version of v1alpha1api20201101.PublicIPAddresses_Spec
type PublicIPAddresses_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName            string                      `json:"azureName"`
	DdosSettings         *DdosSettings               `json:"ddosSettings,omitempty"`
	DnsSettings          *PublicIPAddressDnsSettings `json:"dnsSettings,omitempty"`
	ExtendedLocation     *ExtendedLocation           `json:"extendedLocation,omitempty"`
	IdleTimeoutInMinutes *int                        `json:"idleTimeoutInMinutes,omitempty"`
	IpAddress            *string                     `json:"ipAddress,omitempty"`
	IpTags               []IpTag                     `json:"ipTags,omitempty"`
	Location             *string                     `json:"location,omitempty"`
	OriginalVersion      string                      `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner                    genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag              genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	PublicIPAddressVersion   *string                           `json:"publicIPAddressVersion,omitempty"`
	PublicIPAllocationMethod *string                           `json:"publicIPAllocationMethod,omitempty"`
	PublicIPPrefix           *SubResource                      `json:"publicIPPrefix,omitempty"`
	Sku                      *PublicIPAddressSku               `json:"sku,omitempty"`
	Tags                     map[string]string                 `json:"tags,omitempty"`
	Zones                    []string                          `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &PublicIPAddresses_Spec{}

// ConvertSpecFrom populates our PublicIPAddresses_Spec from the provided source
func (addresses *PublicIPAddresses_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == addresses {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(addresses)
}

// ConvertSpecTo populates the provided destination from our PublicIPAddresses_Spec
func (addresses *PublicIPAddresses_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == addresses {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(addresses)
}

//Storage version of v1alpha1api20201101.DdosSettings
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DdosSettings
type DdosSettings struct {
	DdosCustomPolicy   *SubResource           `json:"ddosCustomPolicy,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProtectedIP        *bool                  `json:"protectedIP,omitempty"`
	ProtectionCoverage *string                `json:"protectionCoverage,omitempty"`
}

//Storage version of v1alpha1api20201101.DdosSettings_Status
type DdosSettings_Status struct {
	DdosCustomPolicy   *SubResource_Status    `json:"ddosCustomPolicy,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProtectedIP        *bool                  `json:"protectedIP,omitempty"`
	ProtectionCoverage *string                `json:"protectionCoverage,omitempty"`
}

//Storage version of v1alpha1api20201101.IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded
type IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded struct {
	Etag                      *string                                            `json:"etag,omitempty"`
	Id                        *string                                            `json:"id,omitempty"`
	Name                      *string                                            `json:"name,omitempty"`
	PrivateIPAddress          *string                                            `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *string                                            `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag                             `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                                            `json:"provisioningState,omitempty"`
	Subnet                    *Subnet_Status_PublicIPAddress_SubResourceEmbedded `json:"subnet,omitempty"`
}

//Storage version of v1alpha1api20201101.IpTag
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpTag
type IpTag struct {
	IpTagType   *string                `json:"ipTagType,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tag         *string                `json:"tag,omitempty"`
}

//Storage version of v1alpha1api20201101.IpTag_Status
type IpTag_Status struct {
	IpTagType   *string                `json:"ipTagType,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tag         *string                `json:"tag,omitempty"`
}

//Storage version of v1alpha1api20201101.NatGateway_Status_PublicIPAddress_SubResourceEmbedded
type NatGateway_Status_PublicIPAddress_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Sku         *NatGatewaySku_Status  `json:"sku,omitempty"`
	Zones       []string               `json:"zones,omitempty"`
}

//Storage version of v1alpha1api20201101.PublicIPAddressDnsSettings
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressDnsSettings
type PublicIPAddressDnsSettings struct {
	DomainNameLabel *string                `json:"domainNameLabel,omitempty"`
	Fqdn            *string                `json:"fqdn,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReverseFqdn     *string                `json:"reverseFqdn,omitempty"`
}

//Storage version of v1alpha1api20201101.PublicIPAddressDnsSettings_Status
type PublicIPAddressDnsSettings_Status struct {
	DomainNameLabel *string                `json:"domainNameLabel,omitempty"`
	Fqdn            *string                `json:"fqdn,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReverseFqdn     *string                `json:"reverseFqdn,omitempty"`
}

//Storage version of v1alpha1api20201101.PublicIPAddressSku
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressSku
type PublicIPAddressSku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20201101.PublicIPAddressSku_Status
type PublicIPAddressSku_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20201101.NatGatewaySku_Status
type NatGatewaySku_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.Subnet_Status_PublicIPAddress_SubResourceEmbedded
type Subnet_Status_PublicIPAddress_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&PublicIPAddress{}, &PublicIPAddressList{})
}
