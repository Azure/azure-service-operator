// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	v20180901s "github.com/Azure/azure-service-operator/v2/api/network/v1api20180901/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20180901.PrivateDnsZone
// Deprecated version of PrivateDnsZone. Use v1api20180901.PrivateDnsZone instead
type PrivateDnsZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDnsZone_Spec   `json:"spec,omitempty"`
	Status            PrivateDnsZone_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PrivateDnsZone{}

// GetConditions returns the conditions of the resource
func (zone *PrivateDnsZone) GetConditions() conditions.Conditions {
	return zone.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (zone *PrivateDnsZone) SetConditions(conditions conditions.Conditions) {
	zone.Status.Conditions = conditions
}

var _ conversion.Convertible = &PrivateDnsZone{}

// ConvertFrom populates our PrivateDnsZone from the provided hub PrivateDnsZone
func (zone *PrivateDnsZone) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20180901s.PrivateDnsZone)
	if !ok {
		return fmt.Errorf("expected network/v1api20180901/storage/PrivateDnsZone but received %T instead", hub)
	}

	return zone.AssignProperties_From_PrivateDnsZone(source)
}

// ConvertTo populates the provided hub PrivateDnsZone from our PrivateDnsZone
func (zone *PrivateDnsZone) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20180901s.PrivateDnsZone)
	if !ok {
		return fmt.Errorf("expected network/v1api20180901/storage/PrivateDnsZone but received %T instead", hub)
	}

	return zone.AssignProperties_To_PrivateDnsZone(destination)
}

var _ genruntime.KubernetesResource = &PrivateDnsZone{}

// AzureName returns the Azure name of the resource
func (zone *PrivateDnsZone) AzureName() string {
	return zone.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-09-01"
func (zone PrivateDnsZone) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (zone *PrivateDnsZone) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (zone *PrivateDnsZone) GetSpec() genruntime.ConvertibleSpec {
	return &zone.Spec
}

// GetStatus returns the status of this resource
func (zone *PrivateDnsZone) GetStatus() genruntime.ConvertibleStatus {
	return &zone.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (zone *PrivateDnsZone) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateDnsZones"
func (zone *PrivateDnsZone) GetType() string {
	return "Microsoft.Network/privateDnsZones"
}

// NewEmptyStatus returns a new empty (blank) status
func (zone *PrivateDnsZone) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PrivateDnsZone_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (zone *PrivateDnsZone) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(zone.Spec)
	return zone.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (zone *PrivateDnsZone) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PrivateDnsZone_STATUS); ok {
		zone.Status = *st
		return nil
	}

	// Convert status to required version
	var st PrivateDnsZone_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	zone.Status = st
	return nil
}

// AssignProperties_From_PrivateDnsZone populates our PrivateDnsZone from the provided source PrivateDnsZone
func (zone *PrivateDnsZone) AssignProperties_From_PrivateDnsZone(source *v20180901s.PrivateDnsZone) error {

	// ObjectMeta
	zone.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec PrivateDnsZone_Spec
	err := spec.AssignProperties_From_PrivateDnsZone_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_PrivateDnsZone_Spec() to populate field Spec")
	}
	zone.Spec = spec

	// Status
	var status PrivateDnsZone_STATUS
	err = status.AssignProperties_From_PrivateDnsZone_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_PrivateDnsZone_STATUS() to populate field Status")
	}
	zone.Status = status

	// Invoke the augmentConversionForPrivateDnsZone interface (if implemented) to customize the conversion
	var zoneAsAny any = zone
	if augmentedZone, ok := zoneAsAny.(augmentConversionForPrivateDnsZone); ok {
		err := augmentedZone.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_PrivateDnsZone populates the provided destination PrivateDnsZone from our PrivateDnsZone
func (zone *PrivateDnsZone) AssignProperties_To_PrivateDnsZone(destination *v20180901s.PrivateDnsZone) error {

	// ObjectMeta
	destination.ObjectMeta = *zone.ObjectMeta.DeepCopy()

	// Spec
	var spec v20180901s.PrivateDnsZone_Spec
	err := zone.Spec.AssignProperties_To_PrivateDnsZone_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_PrivateDnsZone_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20180901s.PrivateDnsZone_STATUS
	err = zone.Status.AssignProperties_To_PrivateDnsZone_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_PrivateDnsZone_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForPrivateDnsZone interface (if implemented) to customize the conversion
	var zoneAsAny any = zone
	if augmentedZone, ok := zoneAsAny.(augmentConversionForPrivateDnsZone); ok {
		err := augmentedZone.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (zone *PrivateDnsZone) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: zone.Spec.OriginalVersion,
		Kind:    "PrivateDnsZone",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20180901.PrivateDnsZone
// Deprecated version of PrivateDnsZone. Use v1api20180901.PrivateDnsZone instead
type PrivateDnsZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDnsZone `json:"items"`
}

// Storage version of v1beta20180901.APIVersion
// Deprecated version of APIVersion. Use v1api20180901.APIVersion instead
// +kubebuilder:validation:Enum={"2018-09-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2018-09-01")

type augmentConversionForPrivateDnsZone interface {
	AssignPropertiesFrom(src *v20180901s.PrivateDnsZone) error
	AssignPropertiesTo(dst *v20180901s.PrivateDnsZone) error
}

// Storage version of v1beta20180901.PrivateDnsZone_Spec
type PrivateDnsZone_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Etag            *string `json:"etag,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &PrivateDnsZone_Spec{}

// ConvertSpecFrom populates our PrivateDnsZone_Spec from the provided source
func (zone *PrivateDnsZone_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20180901s.PrivateDnsZone_Spec)
	if ok {
		// Populate our instance from source
		return zone.AssignProperties_From_PrivateDnsZone_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20180901s.PrivateDnsZone_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = zone.AssignProperties_From_PrivateDnsZone_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our PrivateDnsZone_Spec
func (zone *PrivateDnsZone_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20180901s.PrivateDnsZone_Spec)
	if ok {
		// Populate destination from our instance
		return zone.AssignProperties_To_PrivateDnsZone_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20180901s.PrivateDnsZone_Spec{}
	err := zone.AssignProperties_To_PrivateDnsZone_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_PrivateDnsZone_Spec populates our PrivateDnsZone_Spec from the provided source PrivateDnsZone_Spec
func (zone *PrivateDnsZone_Spec) AssignProperties_From_PrivateDnsZone_Spec(source *v20180901s.PrivateDnsZone_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	zone.AzureName = source.AzureName

	// Etag
	zone.Etag = genruntime.ClonePointerToString(source.Etag)

	// Location
	zone.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	zone.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		zone.Owner = &owner
	} else {
		zone.Owner = nil
	}

	// Tags
	zone.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		zone.PropertyBag = propertyBag
	} else {
		zone.PropertyBag = nil
	}

	// Invoke the augmentConversionForPrivateDnsZone_Spec interface (if implemented) to customize the conversion
	var zoneAsAny any = zone
	if augmentedZone, ok := zoneAsAny.(augmentConversionForPrivateDnsZone_Spec); ok {
		err := augmentedZone.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_PrivateDnsZone_Spec populates the provided destination PrivateDnsZone_Spec from our PrivateDnsZone_Spec
func (zone *PrivateDnsZone_Spec) AssignProperties_To_PrivateDnsZone_Spec(destination *v20180901s.PrivateDnsZone_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(zone.PropertyBag)

	// AzureName
	destination.AzureName = zone.AzureName

	// Etag
	destination.Etag = genruntime.ClonePointerToString(zone.Etag)

	// Location
	destination.Location = genruntime.ClonePointerToString(zone.Location)

	// OriginalVersion
	destination.OriginalVersion = zone.OriginalVersion

	// Owner
	if zone.Owner != nil {
		owner := zone.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(zone.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForPrivateDnsZone_Spec interface (if implemented) to customize the conversion
	var zoneAsAny any = zone
	if augmentedZone, ok := zoneAsAny.(augmentConversionForPrivateDnsZone_Spec); ok {
		err := augmentedZone.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20180901.PrivateDnsZone_STATUS
// Deprecated version of PrivateDnsZone_STATUS. Use v1api20180901.PrivateDnsZone_STATUS instead
type PrivateDnsZone_STATUS struct {
	Conditions                                     []conditions.Condition `json:"conditions,omitempty"`
	Etag                                           *string                `json:"etag,omitempty"`
	Id                                             *string                `json:"id,omitempty"`
	Location                                       *string                `json:"location,omitempty"`
	MaxNumberOfRecordSets                          *int                   `json:"maxNumberOfRecordSets,omitempty"`
	MaxNumberOfVirtualNetworkLinks                 *int                   `json:"maxNumberOfVirtualNetworkLinks,omitempty"`
	MaxNumberOfVirtualNetworkLinksWithRegistration *int                   `json:"maxNumberOfVirtualNetworkLinksWithRegistration,omitempty"`
	Name                                           *string                `json:"name,omitempty"`
	NumberOfRecordSets                             *int                   `json:"numberOfRecordSets,omitempty"`
	NumberOfVirtualNetworkLinks                    *int                   `json:"numberOfVirtualNetworkLinks,omitempty"`
	NumberOfVirtualNetworkLinksWithRegistration    *int                   `json:"numberOfVirtualNetworkLinksWithRegistration,omitempty"`
	PropertyBag                                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState                              *string                `json:"provisioningState,omitempty"`
	Tags                                           map[string]string      `json:"tags,omitempty"`
	Type                                           *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &PrivateDnsZone_STATUS{}

// ConvertStatusFrom populates our PrivateDnsZone_STATUS from the provided source
func (zone *PrivateDnsZone_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20180901s.PrivateDnsZone_STATUS)
	if ok {
		// Populate our instance from source
		return zone.AssignProperties_From_PrivateDnsZone_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20180901s.PrivateDnsZone_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = zone.AssignProperties_From_PrivateDnsZone_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our PrivateDnsZone_STATUS
func (zone *PrivateDnsZone_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20180901s.PrivateDnsZone_STATUS)
	if ok {
		// Populate destination from our instance
		return zone.AssignProperties_To_PrivateDnsZone_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20180901s.PrivateDnsZone_STATUS{}
	err := zone.AssignProperties_To_PrivateDnsZone_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_PrivateDnsZone_STATUS populates our PrivateDnsZone_STATUS from the provided source PrivateDnsZone_STATUS
func (zone *PrivateDnsZone_STATUS) AssignProperties_From_PrivateDnsZone_STATUS(source *v20180901s.PrivateDnsZone_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	zone.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Etag
	zone.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	zone.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	zone.Location = genruntime.ClonePointerToString(source.Location)

	// MaxNumberOfRecordSets
	zone.MaxNumberOfRecordSets = genruntime.ClonePointerToInt(source.MaxNumberOfRecordSets)

	// MaxNumberOfVirtualNetworkLinks
	zone.MaxNumberOfVirtualNetworkLinks = genruntime.ClonePointerToInt(source.MaxNumberOfVirtualNetworkLinks)

	// MaxNumberOfVirtualNetworkLinksWithRegistration
	zone.MaxNumberOfVirtualNetworkLinksWithRegistration = genruntime.ClonePointerToInt(source.MaxNumberOfVirtualNetworkLinksWithRegistration)

	// Name
	zone.Name = genruntime.ClonePointerToString(source.Name)

	// NumberOfRecordSets
	zone.NumberOfRecordSets = genruntime.ClonePointerToInt(source.NumberOfRecordSets)

	// NumberOfVirtualNetworkLinks
	zone.NumberOfVirtualNetworkLinks = genruntime.ClonePointerToInt(source.NumberOfVirtualNetworkLinks)

	// NumberOfVirtualNetworkLinksWithRegistration
	zone.NumberOfVirtualNetworkLinksWithRegistration = genruntime.ClonePointerToInt(source.NumberOfVirtualNetworkLinksWithRegistration)

	// ProvisioningState
	zone.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// Tags
	zone.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	zone.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		zone.PropertyBag = propertyBag
	} else {
		zone.PropertyBag = nil
	}

	// Invoke the augmentConversionForPrivateDnsZone_STATUS interface (if implemented) to customize the conversion
	var zoneAsAny any = zone
	if augmentedZone, ok := zoneAsAny.(augmentConversionForPrivateDnsZone_STATUS); ok {
		err := augmentedZone.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_PrivateDnsZone_STATUS populates the provided destination PrivateDnsZone_STATUS from our PrivateDnsZone_STATUS
func (zone *PrivateDnsZone_STATUS) AssignProperties_To_PrivateDnsZone_STATUS(destination *v20180901s.PrivateDnsZone_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(zone.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(zone.Conditions)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(zone.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(zone.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(zone.Location)

	// MaxNumberOfRecordSets
	destination.MaxNumberOfRecordSets = genruntime.ClonePointerToInt(zone.MaxNumberOfRecordSets)

	// MaxNumberOfVirtualNetworkLinks
	destination.MaxNumberOfVirtualNetworkLinks = genruntime.ClonePointerToInt(zone.MaxNumberOfVirtualNetworkLinks)

	// MaxNumberOfVirtualNetworkLinksWithRegistration
	destination.MaxNumberOfVirtualNetworkLinksWithRegistration = genruntime.ClonePointerToInt(zone.MaxNumberOfVirtualNetworkLinksWithRegistration)

	// Name
	destination.Name = genruntime.ClonePointerToString(zone.Name)

	// NumberOfRecordSets
	destination.NumberOfRecordSets = genruntime.ClonePointerToInt(zone.NumberOfRecordSets)

	// NumberOfVirtualNetworkLinks
	destination.NumberOfVirtualNetworkLinks = genruntime.ClonePointerToInt(zone.NumberOfVirtualNetworkLinks)

	// NumberOfVirtualNetworkLinksWithRegistration
	destination.NumberOfVirtualNetworkLinksWithRegistration = genruntime.ClonePointerToInt(zone.NumberOfVirtualNetworkLinksWithRegistration)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(zone.ProvisioningState)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(zone.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(zone.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForPrivateDnsZone_STATUS interface (if implemented) to customize the conversion
	var zoneAsAny any = zone
	if augmentedZone, ok := zoneAsAny.(augmentConversionForPrivateDnsZone_STATUS); ok {
		err := augmentedZone.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForPrivateDnsZone_Spec interface {
	AssignPropertiesFrom(src *v20180901s.PrivateDnsZone_Spec) error
	AssignPropertiesTo(dst *v20180901s.PrivateDnsZone_Spec) error
}

type augmentConversionForPrivateDnsZone_STATUS interface {
	AssignPropertiesFrom(src *v20180901s.PrivateDnsZone_STATUS) error
	AssignPropertiesTo(dst *v20180901s.PrivateDnsZone_STATUS) error
}

func init() {
	SchemeBuilder.Register(&PrivateDnsZone{}, &PrivateDnsZoneList{})
}
