// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101storage"
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
// Storage version of v1beta20211101.NamespacesEventhubsConsumerGroup
// Deprecated version of NamespacesEventhubsConsumerGroup. Use v1api20211101.NamespacesEventhubsConsumerGroup instead
type NamespacesEventhubsConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Eventhubs_Consumergroup_Spec   `json:"spec,omitempty"`
	Status            Namespaces_Eventhubs_Consumergroup_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesEventhubsConsumerGroup{}

// GetConditions returns the conditions of the resource
func (group *NamespacesEventhubsConsumerGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *NamespacesEventhubsConsumerGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesEventhubsConsumerGroup{}

// ConvertFrom populates our NamespacesEventhubsConsumerGroup from the provided hub NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.NamespacesEventhubsConsumerGroup)
	if !ok {
		return fmt.Errorf("expected eventhub/v1api20211101storage/NamespacesEventhubsConsumerGroup but received %T instead", hub)
	}

	return group.AssignProperties_From_NamespacesEventhubsConsumerGroup(source)
}

// ConvertTo populates the provided hub NamespacesEventhubsConsumerGroup from our NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.NamespacesEventhubsConsumerGroup)
	if !ok {
		return fmt.Errorf("expected eventhub/v1api20211101storage/NamespacesEventhubsConsumerGroup but received %T instead", hub)
	}

	return group.AssignProperties_To_NamespacesEventhubsConsumerGroup(destination)
}

var _ genruntime.KubernetesResource = &NamespacesEventhubsConsumerGroup{}

// AzureName returns the Azure name of the resource
func (group *NamespacesEventhubsConsumerGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (group NamespacesEventhubsConsumerGroup) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (group *NamespacesEventhubsConsumerGroup) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (group *NamespacesEventhubsConsumerGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *NamespacesEventhubsConsumerGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/consumergroups"
func (group *NamespacesEventhubsConsumerGroup) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/consumergroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *NamespacesEventhubsConsumerGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Namespaces_Eventhubs_Consumergroup_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (group *NamespacesEventhubsConsumerGroup) Owner() *genruntime.ResourceReference {
	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return &genruntime.ResourceReference{
		Group: ownerGroup,
		Kind:  ownerKind,
		Name:  group.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (group *NamespacesEventhubsConsumerGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Namespaces_Eventhubs_Consumergroup_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st Namespaces_Eventhubs_Consumergroup_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// AssignProperties_From_NamespacesEventhubsConsumerGroup populates our NamespacesEventhubsConsumerGroup from the provided source NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) AssignProperties_From_NamespacesEventhubsConsumerGroup(source *v20211101s.NamespacesEventhubsConsumerGroup) error {

	// ObjectMeta
	group.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Namespaces_Eventhubs_Consumergroup_Spec
	err := spec.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec() to populate field Spec")
	}
	group.Spec = spec

	// Status
	var status Namespaces_Eventhubs_Consumergroup_STATUS
	err = status.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS() to populate field Status")
	}
	group.Status = status

	// Invoke the augmentConversionForNamespacesEventhubsConsumerGroup interface (if implemented) to customize the conversion
	var groupAsAny any = group
	if augmentedGroup, ok := groupAsAny.(augmentConversionForNamespacesEventhubsConsumerGroup); ok {
		err := augmentedGroup.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesEventhubsConsumerGroup populates the provided destination NamespacesEventhubsConsumerGroup from our NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) AssignProperties_To_NamespacesEventhubsConsumerGroup(destination *v20211101s.NamespacesEventhubsConsumerGroup) error {

	// ObjectMeta
	destination.ObjectMeta = *group.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Namespaces_Eventhubs_Consumergroup_Spec
	err := group.Spec.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Namespaces_Eventhubs_Consumergroup_STATUS
	err = group.Status.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForNamespacesEventhubsConsumerGroup interface (if implemented) to customize the conversion
	var groupAsAny any = group
	if augmentedGroup, ok := groupAsAny.(augmentConversionForNamespacesEventhubsConsumerGroup); ok {
		err := augmentedGroup.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *NamespacesEventhubsConsumerGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion,
		Kind:    "NamespacesEventhubsConsumerGroup",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20211101.NamespacesEventhubsConsumerGroup
// Deprecated version of NamespacesEventhubsConsumerGroup. Use v1api20211101.NamespacesEventhubsConsumerGroup instead
type NamespacesEventhubsConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsConsumerGroup `json:"items"`
}

type augmentConversionForNamespacesEventhubsConsumerGroup interface {
	AssignPropertiesFrom(src *v20211101s.NamespacesEventhubsConsumerGroup) error
	AssignPropertiesTo(dst *v20211101s.NamespacesEventhubsConsumerGroup) error
}

// Storage version of v1beta20211101.Namespaces_Eventhubs_Consumergroup_Spec
type Namespaces_Eventhubs_Consumergroup_Spec struct {
	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/NamespacesEventhub resource
	Owner        *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"NamespacesEventhub"`
	PropertyBag  genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	UserMetadata *string                            `json:"userMetadata,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Namespaces_Eventhubs_Consumergroup_Spec{}

// ConvertSpecFrom populates our Namespaces_Eventhubs_Consumergroup_Spec from the provided source
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Namespaces_Eventhubs_Consumergroup_Spec)
	if ok {
		// Populate our instance from source
		return consumergroup.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Namespaces_Eventhubs_Consumergroup_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = consumergroup.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Namespaces_Eventhubs_Consumergroup_Spec
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Namespaces_Eventhubs_Consumergroup_Spec)
	if ok {
		// Populate destination from our instance
		return consumergroup.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Namespaces_Eventhubs_Consumergroup_Spec{}
	err := consumergroup.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(dst)
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

// AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec populates our Namespaces_Eventhubs_Consumergroup_Spec from the provided source Namespaces_Eventhubs_Consumergroup_Spec
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec(source *v20211101s.Namespaces_Eventhubs_Consumergroup_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	consumergroup.AzureName = source.AzureName

	// OriginalVersion
	consumergroup.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		consumergroup.Owner = &owner
	} else {
		consumergroup.Owner = nil
	}

	// UserMetadata
	consumergroup.UserMetadata = genruntime.ClonePointerToString(source.UserMetadata)

	// Update the property bag
	if len(propertyBag) > 0 {
		consumergroup.PropertyBag = propertyBag
	} else {
		consumergroup.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespaces_Eventhubs_Consumergroup_Spec interface (if implemented) to customize the conversion
	var consumergroupAsAny any = consumergroup
	if augmentedConsumergroup, ok := consumergroupAsAny.(augmentConversionForNamespaces_Eventhubs_Consumergroup_Spec); ok {
		err := augmentedConsumergroup.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec populates the provided destination Namespaces_Eventhubs_Consumergroup_Spec from our Namespaces_Eventhubs_Consumergroup_Spec
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(destination *v20211101s.Namespaces_Eventhubs_Consumergroup_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(consumergroup.PropertyBag)

	// AzureName
	destination.AzureName = consumergroup.AzureName

	// OriginalVersion
	destination.OriginalVersion = consumergroup.OriginalVersion

	// Owner
	if consumergroup.Owner != nil {
		owner := consumergroup.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// UserMetadata
	destination.UserMetadata = genruntime.ClonePointerToString(consumergroup.UserMetadata)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespaces_Eventhubs_Consumergroup_Spec interface (if implemented) to customize the conversion
	var consumergroupAsAny any = consumergroup
	if augmentedConsumergroup, ok := consumergroupAsAny.(augmentConversionForNamespaces_Eventhubs_Consumergroup_Spec); ok {
		err := augmentedConsumergroup.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20211101.Namespaces_Eventhubs_Consumergroup_STATUS
// Deprecated version of Namespaces_Eventhubs_Consumergroup_STATUS. Use v1api20211101.Namespaces_Eventhubs_Consumergroup_STATUS instead
type Namespaces_Eventhubs_Consumergroup_STATUS struct {
	Conditions   []conditions.Condition `json:"conditions,omitempty"`
	CreatedAt    *string                `json:"createdAt,omitempty"`
	Id           *string                `json:"id,omitempty"`
	Location     *string                `json:"location,omitempty"`
	Name         *string                `json:"name,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData   *SystemData_STATUS     `json:"systemData,omitempty"`
	Type         *string                `json:"type,omitempty"`
	UpdatedAt    *string                `json:"updatedAt,omitempty"`
	UserMetadata *string                `json:"userMetadata,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Namespaces_Eventhubs_Consumergroup_STATUS{}

// ConvertStatusFrom populates our Namespaces_Eventhubs_Consumergroup_STATUS from the provided source
func (consumergroup *Namespaces_Eventhubs_Consumergroup_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Namespaces_Eventhubs_Consumergroup_STATUS)
	if ok {
		// Populate our instance from source
		return consumergroup.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Namespaces_Eventhubs_Consumergroup_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = consumergroup.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Namespaces_Eventhubs_Consumergroup_STATUS
func (consumergroup *Namespaces_Eventhubs_Consumergroup_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Namespaces_Eventhubs_Consumergroup_STATUS)
	if ok {
		// Populate destination from our instance
		return consumergroup.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Namespaces_Eventhubs_Consumergroup_STATUS{}
	err := consumergroup.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS(dst)
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

// AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS populates our Namespaces_Eventhubs_Consumergroup_STATUS from the provided source Namespaces_Eventhubs_Consumergroup_STATUS
func (consumergroup *Namespaces_Eventhubs_Consumergroup_STATUS) AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS(source *v20211101s.Namespaces_Eventhubs_Consumergroup_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	consumergroup.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreatedAt
	consumergroup.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// Id
	consumergroup.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	consumergroup.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	consumergroup.Name = genruntime.ClonePointerToString(source.Name)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		consumergroup.SystemData = &systemDatum
	} else {
		consumergroup.SystemData = nil
	}

	// Type
	consumergroup.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedAt
	consumergroup.UpdatedAt = genruntime.ClonePointerToString(source.UpdatedAt)

	// UserMetadata
	consumergroup.UserMetadata = genruntime.ClonePointerToString(source.UserMetadata)

	// Update the property bag
	if len(propertyBag) > 0 {
		consumergroup.PropertyBag = propertyBag
	} else {
		consumergroup.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespaces_Eventhubs_Consumergroup_STATUS interface (if implemented) to customize the conversion
	var consumergroupAsAny any = consumergroup
	if augmentedConsumergroup, ok := consumergroupAsAny.(augmentConversionForNamespaces_Eventhubs_Consumergroup_STATUS); ok {
		err := augmentedConsumergroup.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS populates the provided destination Namespaces_Eventhubs_Consumergroup_STATUS from our Namespaces_Eventhubs_Consumergroup_STATUS
func (consumergroup *Namespaces_Eventhubs_Consumergroup_STATUS) AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS(destination *v20211101s.Namespaces_Eventhubs_Consumergroup_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(consumergroup.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(consumergroup.Conditions)

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(consumergroup.CreatedAt)

	// Id
	destination.Id = genruntime.ClonePointerToString(consumergroup.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(consumergroup.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(consumergroup.Name)

	// SystemData
	if consumergroup.SystemData != nil {
		var systemDatum v20211101s.SystemData_STATUS
		err := consumergroup.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(consumergroup.Type)

	// UpdatedAt
	destination.UpdatedAt = genruntime.ClonePointerToString(consumergroup.UpdatedAt)

	// UserMetadata
	destination.UserMetadata = genruntime.ClonePointerToString(consumergroup.UserMetadata)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespaces_Eventhubs_Consumergroup_STATUS interface (if implemented) to customize the conversion
	var consumergroupAsAny any = consumergroup
	if augmentedConsumergroup, ok := consumergroupAsAny.(augmentConversionForNamespaces_Eventhubs_Consumergroup_STATUS); ok {
		err := augmentedConsumergroup.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForNamespaces_Eventhubs_Consumergroup_Spec interface {
	AssignPropertiesFrom(src *v20211101s.Namespaces_Eventhubs_Consumergroup_Spec) error
	AssignPropertiesTo(dst *v20211101s.Namespaces_Eventhubs_Consumergroup_Spec) error
}

type augmentConversionForNamespaces_Eventhubs_Consumergroup_STATUS interface {
	AssignPropertiesFrom(src *v20211101s.Namespaces_Eventhubs_Consumergroup_STATUS) error
	AssignPropertiesTo(dst *v20211101s.Namespaces_Eventhubs_Consumergroup_STATUS) error
}

func init() {
	SchemeBuilder.Register(&NamespacesEventhubsConsumerGroup{}, &NamespacesEventhubsConsumerGroupList{})
}
