// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601storage

import (
	"fmt"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601storage"
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
// Storage version of v1beta20200601.DomainsTopic
// Deprecated version of DomainsTopic. Use v1api20200601.DomainsTopic instead
type DomainsTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Domains_Topic_Spec   `json:"spec,omitempty"`
	Status            Domains_Topic_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DomainsTopic{}

// GetConditions returns the conditions of the resource
func (topic *DomainsTopic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *DomainsTopic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ conversion.Convertible = &DomainsTopic{}

// ConvertFrom populates our DomainsTopic from the provided hub DomainsTopic
func (topic *DomainsTopic) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20200601s.DomainsTopic)
	if !ok {
		return fmt.Errorf("expected eventgrid/v1api20200601storage/DomainsTopic but received %T instead", hub)
	}

	return topic.AssignProperties_From_DomainsTopic(source)
}

// ConvertTo populates the provided hub DomainsTopic from our DomainsTopic
func (topic *DomainsTopic) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20200601s.DomainsTopic)
	if !ok {
		return fmt.Errorf("expected eventgrid/v1api20200601storage/DomainsTopic but received %T instead", hub)
	}

	return topic.AssignProperties_To_DomainsTopic(destination)
}

var _ genruntime.KubernetesResource = &DomainsTopic{}

// AzureName returns the Azure name of the resource
func (topic *DomainsTopic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic DomainsTopic) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (topic *DomainsTopic) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (topic *DomainsTopic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *DomainsTopic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (topic *DomainsTopic) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *DomainsTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Domains_Topic_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (topic *DomainsTopic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return topic.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (topic *DomainsTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Domains_Topic_STATUS); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st Domains_Topic_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// AssignProperties_From_DomainsTopic populates our DomainsTopic from the provided source DomainsTopic
func (topic *DomainsTopic) AssignProperties_From_DomainsTopic(source *v20200601s.DomainsTopic) error {

	// ObjectMeta
	topic.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Domains_Topic_Spec
	err := spec.AssignProperties_From_Domains_Topic_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Domains_Topic_Spec() to populate field Spec")
	}
	topic.Spec = spec

	// Status
	var status Domains_Topic_STATUS
	err = status.AssignProperties_From_Domains_Topic_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Domains_Topic_STATUS() to populate field Status")
	}
	topic.Status = status

	// Invoke the augmentConversionForDomainsTopic interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForDomainsTopic); ok {
		err := augmentedTopic.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_DomainsTopic populates the provided destination DomainsTopic from our DomainsTopic
func (topic *DomainsTopic) AssignProperties_To_DomainsTopic(destination *v20200601s.DomainsTopic) error {

	// ObjectMeta
	destination.ObjectMeta = *topic.ObjectMeta.DeepCopy()

	// Spec
	var spec v20200601s.Domains_Topic_Spec
	err := topic.Spec.AssignProperties_To_Domains_Topic_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Domains_Topic_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20200601s.Domains_Topic_STATUS
	err = topic.Status.AssignProperties_To_Domains_Topic_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Domains_Topic_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForDomainsTopic interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForDomainsTopic); ok {
		err := augmentedTopic.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *DomainsTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion,
		Kind:    "DomainsTopic",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20200601.DomainsTopic
// Deprecated version of DomainsTopic. Use v1api20200601.DomainsTopic instead
type DomainsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainsTopic `json:"items"`
}

type augmentConversionForDomainsTopic interface {
	AssignPropertiesFrom(src *v20200601s.DomainsTopic) error
	AssignPropertiesTo(dst *v20200601s.DomainsTopic) error
}

// Storage version of v1beta20200601.Domains_Topic_Spec
type Domains_Topic_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventgrid.azure.com/Domain resource
	Owner       *genruntime.KnownResourceReference `group:"eventgrid.azure.com" json:"owner,omitempty" kind:"Domain"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Domains_Topic_Spec{}

// ConvertSpecFrom populates our Domains_Topic_Spec from the provided source
func (topic *Domains_Topic_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20200601s.Domains_Topic_Spec)
	if ok {
		// Populate our instance from source
		return topic.AssignProperties_From_Domains_Topic_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20200601s.Domains_Topic_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = topic.AssignProperties_From_Domains_Topic_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Domains_Topic_Spec
func (topic *Domains_Topic_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20200601s.Domains_Topic_Spec)
	if ok {
		// Populate destination from our instance
		return topic.AssignProperties_To_Domains_Topic_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20200601s.Domains_Topic_Spec{}
	err := topic.AssignProperties_To_Domains_Topic_Spec(dst)
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

// AssignProperties_From_Domains_Topic_Spec populates our Domains_Topic_Spec from the provided source Domains_Topic_Spec
func (topic *Domains_Topic_Spec) AssignProperties_From_Domains_Topic_Spec(source *v20200601s.Domains_Topic_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	topic.AzureName = source.AzureName

	// OriginalVersion
	topic.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		topic.Owner = &owner
	} else {
		topic.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		topic.PropertyBag = propertyBag
	} else {
		topic.PropertyBag = nil
	}

	// Invoke the augmentConversionForDomains_Topic_Spec interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForDomains_Topic_Spec); ok {
		err := augmentedTopic.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Domains_Topic_Spec populates the provided destination Domains_Topic_Spec from our Domains_Topic_Spec
func (topic *Domains_Topic_Spec) AssignProperties_To_Domains_Topic_Spec(destination *v20200601s.Domains_Topic_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(topic.PropertyBag)

	// AzureName
	destination.AzureName = topic.AzureName

	// OriginalVersion
	destination.OriginalVersion = topic.OriginalVersion

	// Owner
	if topic.Owner != nil {
		owner := topic.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForDomains_Topic_Spec interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForDomains_Topic_Spec); ok {
		err := augmentedTopic.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20200601.Domains_Topic_STATUS
// Deprecated version of Domains_Topic_STATUS. Use v1api20200601.Domains_Topic_STATUS instead
type Domains_Topic_STATUS struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	SystemData        *SystemData_STATUS     `json:"systemData,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Domains_Topic_STATUS{}

// ConvertStatusFrom populates our Domains_Topic_STATUS from the provided source
func (topic *Domains_Topic_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20200601s.Domains_Topic_STATUS)
	if ok {
		// Populate our instance from source
		return topic.AssignProperties_From_Domains_Topic_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20200601s.Domains_Topic_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = topic.AssignProperties_From_Domains_Topic_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Domains_Topic_STATUS
func (topic *Domains_Topic_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20200601s.Domains_Topic_STATUS)
	if ok {
		// Populate destination from our instance
		return topic.AssignProperties_To_Domains_Topic_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20200601s.Domains_Topic_STATUS{}
	err := topic.AssignProperties_To_Domains_Topic_STATUS(dst)
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

// AssignProperties_From_Domains_Topic_STATUS populates our Domains_Topic_STATUS from the provided source Domains_Topic_STATUS
func (topic *Domains_Topic_STATUS) AssignProperties_From_Domains_Topic_STATUS(source *v20200601s.Domains_Topic_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	topic.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	topic.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	topic.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	topic.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		topic.SystemData = &systemDatum
	} else {
		topic.SystemData = nil
	}

	// Type
	topic.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		topic.PropertyBag = propertyBag
	} else {
		topic.PropertyBag = nil
	}

	// Invoke the augmentConversionForDomains_Topic_STATUS interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForDomains_Topic_STATUS); ok {
		err := augmentedTopic.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Domains_Topic_STATUS populates the provided destination Domains_Topic_STATUS from our Domains_Topic_STATUS
func (topic *Domains_Topic_STATUS) AssignProperties_To_Domains_Topic_STATUS(destination *v20200601s.Domains_Topic_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(topic.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(topic.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(topic.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(topic.Name)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(topic.ProvisioningState)

	// SystemData
	if topic.SystemData != nil {
		var systemDatum v20200601s.SystemData_STATUS
		err := topic.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(topic.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForDomains_Topic_STATUS interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForDomains_Topic_STATUS); ok {
		err := augmentedTopic.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForDomains_Topic_Spec interface {
	AssignPropertiesFrom(src *v20200601s.Domains_Topic_Spec) error
	AssignPropertiesTo(dst *v20200601s.Domains_Topic_Spec) error
}

type augmentConversionForDomains_Topic_STATUS interface {
	AssignPropertiesFrom(src *v20200601s.Domains_Topic_STATUS) error
	AssignPropertiesTo(dst *v20200601s.Domains_Topic_STATUS) error
}

func init() {
	SchemeBuilder.Register(&DomainsTopic{}, &DomainsTopicList{})
}
