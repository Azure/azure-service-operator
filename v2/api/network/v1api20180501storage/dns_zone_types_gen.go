// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501storage

import (
	v20200601s "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=dnszones,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={dnszones/status,dnszones/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20180501.DnsZone
// Generator information:
// - Generated from: /dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}
type DnsZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsZone_Spec   `json:"spec,omitempty"`
	Status            DnsZone_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsZone{}

// GetConditions returns the conditions of the resource
func (zone *DnsZone) GetConditions() conditions.Conditions {
	return zone.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (zone *DnsZone) SetConditions(conditions conditions.Conditions) {
	zone.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &DnsZone{}

// AzureName returns the Azure name of the resource
func (zone *DnsZone) AzureName() string {
	return zone.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-05-01"
func (zone DnsZone) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (zone *DnsZone) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (zone *DnsZone) GetSpec() genruntime.ConvertibleSpec {
	return &zone.Spec
}

// GetStatus returns the status of this resource
func (zone *DnsZone) GetStatus() genruntime.ConvertibleStatus {
	return &zone.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsZones"
func (zone *DnsZone) GetType() string {
	return "Microsoft.Network/dnsZones"
}

// NewEmptyStatus returns a new empty (blank) status
func (zone *DnsZone) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsZone_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (zone *DnsZone) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(zone.Spec)
	return zone.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (zone *DnsZone) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsZone_STATUS); ok {
		zone.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsZone_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	zone.Status = st
	return nil
}

// Hub marks that this DnsZone is the hub type for conversion
func (zone *DnsZone) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (zone *DnsZone) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: zone.Spec.OriginalVersion,
		Kind:    "DnsZone",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20180501.DnsZone
// Generator information:
// - Generated from: /dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}
type DnsZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsZone `json:"items"`
}

// Storage version of v1api20180501.APIVersion
// +kubebuilder:validation:Enum={"2018-05-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2018-05-01")

// Storage version of v1api20180501.DnsZone_Spec
type DnsZone_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag                 genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RegistrationVirtualNetworks []SubResource                      `json:"registrationVirtualNetworks,omitempty"`
	ResolutionVirtualNetworks   []SubResource                      `json:"resolutionVirtualNetworks,omitempty"`
	Tags                        map[string]string                  `json:"tags,omitempty"`
	ZoneType                    *string                            `json:"zoneType,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DnsZone_Spec{}

// ConvertSpecFrom populates our DnsZone_Spec from the provided source
func (zone *DnsZone_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == zone {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(zone)
}

// ConvertSpecTo populates the provided destination from our DnsZone_Spec
func (zone *DnsZone_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == zone {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(zone)
}

// Storage version of v1api20180501.DnsZone_STATUS
type DnsZone_STATUS struct {
	Conditions                     []conditions.Condition `json:"conditions,omitempty"`
	Etag                           *string                `json:"etag,omitempty"`
	Id                             *string                `json:"id,omitempty"`
	Location                       *string                `json:"location,omitempty"`
	MaxNumberOfRecordSets          *int                   `json:"maxNumberOfRecordSets,omitempty"`
	MaxNumberOfRecordsPerRecordSet *int                   `json:"maxNumberOfRecordsPerRecordSet,omitempty"`
	Name                           *string                `json:"name,omitempty"`
	NameServers                    []string               `json:"nameServers,omitempty"`
	NumberOfRecordSets             *int                   `json:"numberOfRecordSets,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RegistrationVirtualNetworks    []SubResource_STATUS   `json:"registrationVirtualNetworks,omitempty"`
	ResolutionVirtualNetworks      []SubResource_STATUS   `json:"resolutionVirtualNetworks,omitempty"`
	Tags                           map[string]string      `json:"tags,omitempty"`
	Type                           *string                `json:"type,omitempty"`
	ZoneType                       *string                `json:"zoneType,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DnsZone_STATUS{}

// ConvertStatusFrom populates our DnsZone_STATUS from the provided source
func (zone *DnsZone_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == zone {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(zone)
}

// ConvertStatusTo populates the provided destination from our DnsZone_STATUS
func (zone *DnsZone_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == zone {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(zone)
}

// Storage version of v1api20180501.SubResource
// A reference to a another resource
type SubResource struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource Id.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// AssignProperties_From_SubResource populates our SubResource from the provided source SubResource
func (resource *SubResource) AssignProperties_From_SubResource(source *v20200601s.SubResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Reference
	if source.Reference != nil {
		reference := source.Reference.Copy()
		resource.Reference = &reference
	} else {
		resource.Reference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SubResource populates the provided destination SubResource from our SubResource
func (resource *SubResource) AssignProperties_To_SubResource(destination *v20200601s.SubResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Reference
	if resource.Reference != nil {
		reference := resource.Reference.Copy()
		destination.Reference = &reference
	} else {
		destination.Reference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20180501.SubResource_STATUS
// A reference to a another resource
type SubResource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_SubResource_STATUS populates our SubResource_STATUS from the provided source SubResource_STATUS
func (resource *SubResource_STATUS) AssignProperties_From_SubResource_STATUS(source *v20200601s.SubResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource_STATUS); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SubResource_STATUS populates the provided destination SubResource_STATUS from our SubResource_STATUS
func (resource *SubResource_STATUS) AssignProperties_To_SubResource_STATUS(destination *v20200601s.SubResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource_STATUS); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSubResource interface {
	AssignPropertiesFrom(src *v20200601s.SubResource) error
	AssignPropertiesTo(dst *v20200601s.SubResource) error
}

type augmentConversionForSubResource_STATUS interface {
	AssignPropertiesFrom(src *v20200601s.SubResource_STATUS) error
	AssignPropertiesTo(dst *v20200601s.SubResource_STATUS) error
}

func init() {
	SchemeBuilder.Register(&DnsZone{}, &DnsZoneList{})
}
