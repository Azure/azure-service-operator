// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101storage

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
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
// Storage version of v1alpha1api20211101.NamespacesEventhub
// Deprecated version of NamespacesEventhub. Use v1beta20211101.NamespacesEventhub instead
type NamespacesEventhub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesEventhubs_Spec `json:"spec,omitempty"`
	Status            Eventhub_Status          `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesEventhub{}

// GetConditions returns the conditions of the resource
func (eventhub *NamespacesEventhub) GetConditions() conditions.Conditions {
	return eventhub.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (eventhub *NamespacesEventhub) SetConditions(conditions conditions.Conditions) {
	eventhub.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesEventhub{}

// ConvertFrom populates our NamespacesEventhub from the provided hub NamespacesEventhub
func (eventhub *NamespacesEventhub) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.NamespacesEventhub)
	if !ok {
		return fmt.Errorf("expected eventhub/v1beta20211101storage/NamespacesEventhub but received %T instead", hub)
	}

	return eventhub.AssignPropertiesFromNamespacesEventhub(source)
}

// ConvertTo populates the provided hub NamespacesEventhub from our NamespacesEventhub
func (eventhub *NamespacesEventhub) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.NamespacesEventhub)
	if !ok {
		return fmt.Errorf("expected eventhub/v1beta20211101storage/NamespacesEventhub but received %T instead", hub)
	}

	return eventhub.AssignPropertiesToNamespacesEventhub(destination)
}

var _ genruntime.KubernetesResource = &NamespacesEventhub{}

// AzureName returns the Azure name of the resource
func (eventhub *NamespacesEventhub) AzureName() string {
	return eventhub.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (eventhub NamespacesEventhub) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (eventhub *NamespacesEventhub) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (eventhub *NamespacesEventhub) GetSpec() genruntime.ConvertibleSpec {
	return &eventhub.Spec
}

// GetStatus returns the status of this resource
func (eventhub *NamespacesEventhub) GetStatus() genruntime.ConvertibleStatus {
	return &eventhub.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs"
func (eventhub *NamespacesEventhub) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs"
}

// NewEmptyStatus returns a new empty (blank) status
func (eventhub *NamespacesEventhub) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Eventhub_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (eventhub *NamespacesEventhub) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(eventhub.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  eventhub.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (eventhub *NamespacesEventhub) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Eventhub_Status); ok {
		eventhub.Status = *st
		return nil
	}

	// Convert status to required version
	var st Eventhub_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	eventhub.Status = st
	return nil
}

// AssignPropertiesFromNamespacesEventhub populates our NamespacesEventhub from the provided source NamespacesEventhub
func (eventhub *NamespacesEventhub) AssignPropertiesFromNamespacesEventhub(source *v20211101s.NamespacesEventhub) error {

	// ObjectMeta
	eventhub.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesEventhubs_Spec
	err := spec.AssignPropertiesFromNamespacesEventhubsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesEventhubsSpec() to populate field Spec")
	}
	eventhub.Spec = spec

	// Status
	var status Eventhub_Status
	err = status.AssignPropertiesFromEventhubStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromEventhubStatus() to populate field Status")
	}
	eventhub.Status = status

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhub populates the provided destination NamespacesEventhub from our NamespacesEventhub
func (eventhub *NamespacesEventhub) AssignPropertiesToNamespacesEventhub(destination *v20211101s.NamespacesEventhub) error {

	// ObjectMeta
	destination.ObjectMeta = *eventhub.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.NamespacesEventhubs_Spec
	err := eventhub.Spec.AssignPropertiesToNamespacesEventhubsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesEventhubsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Eventhub_Status
	err = eventhub.Status.AssignPropertiesToEventhubStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToEventhubStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (eventhub *NamespacesEventhub) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: eventhub.Spec.OriginalVersion,
		Kind:    "NamespacesEventhub",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20211101.NamespacesEventhub
// Deprecated version of NamespacesEventhub. Use v1beta20211101.NamespacesEventhub instead
type NamespacesEventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhub `json:"items"`
}

// Storage version of v1alpha1api20211101.Eventhub_Status
// Deprecated version of Eventhub_Status. Use v1beta20211101.Eventhub_Status instead
type Eventhub_Status struct {
	CaptureDescription     *CaptureDescription_Status `json:"captureDescription,omitempty"`
	Conditions             []conditions.Condition     `json:"conditions,omitempty"`
	CreatedAt              *string                    `json:"createdAt,omitempty"`
	Id                     *string                    `json:"id,omitempty"`
	Location               *string                    `json:"location,omitempty"`
	MessageRetentionInDays *int                       `json:"messageRetentionInDays,omitempty"`
	Name                   *string                    `json:"name,omitempty"`
	PartitionCount         *int                       `json:"partitionCount,omitempty"`
	PartitionIds           []string                   `json:"partitionIds,omitempty"`
	PropertyBag            genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Status                 *string                    `json:"status,omitempty"`
	SystemData             *SystemData_Status         `json:"systemData,omitempty"`
	Type                   *string                    `json:"type,omitempty"`
	UpdatedAt              *string                    `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Eventhub_Status{}

// ConvertStatusFrom populates our Eventhub_Status from the provided source
func (eventhub *Eventhub_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Eventhub_Status)
	if ok {
		// Populate our instance from source
		return eventhub.AssignPropertiesFromEventhubStatus(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Eventhub_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = eventhub.AssignPropertiesFromEventhubStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Eventhub_Status
func (eventhub *Eventhub_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Eventhub_Status)
	if ok {
		// Populate destination from our instance
		return eventhub.AssignPropertiesToEventhubStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Eventhub_Status{}
	err := eventhub.AssignPropertiesToEventhubStatus(dst)
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

// AssignPropertiesFromEventhubStatus populates our Eventhub_Status from the provided source Eventhub_Status
func (eventhub *Eventhub_Status) AssignPropertiesFromEventhubStatus(source *v20211101s.Eventhub_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// CaptureDescription
	if source.CaptureDescription != nil {
		var captureDescription CaptureDescription_Status
		err := captureDescription.AssignPropertiesFromCaptureDescriptionStatus(source.CaptureDescription)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCaptureDescriptionStatus() to populate field CaptureDescription")
		}
		eventhub.CaptureDescription = &captureDescription
	} else {
		eventhub.CaptureDescription = nil
	}

	// Conditions
	eventhub.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreatedAt
	eventhub.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// Id
	eventhub.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	eventhub.Location = genruntime.ClonePointerToString(source.Location)

	// MessageRetentionInDays
	eventhub.MessageRetentionInDays = genruntime.ClonePointerToInt(source.MessageRetentionInDays)

	// Name
	eventhub.Name = genruntime.ClonePointerToString(source.Name)

	// PartitionCount
	eventhub.PartitionCount = genruntime.ClonePointerToInt(source.PartitionCount)

	// PartitionIds
	eventhub.PartitionIds = genruntime.CloneSliceOfString(source.PartitionIds)

	// Status
	eventhub.Status = genruntime.ClonePointerToString(source.Status)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		eventhub.SystemData = &systemDatum
	} else {
		eventhub.SystemData = nil
	}

	// Type
	eventhub.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedAt
	eventhub.UpdatedAt = genruntime.ClonePointerToString(source.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		eventhub.PropertyBag = propertyBag
	} else {
		eventhub.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToEventhubStatus populates the provided destination Eventhub_Status from our Eventhub_Status
func (eventhub *Eventhub_Status) AssignPropertiesToEventhubStatus(destination *v20211101s.Eventhub_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(eventhub.PropertyBag)

	// CaptureDescription
	if eventhub.CaptureDescription != nil {
		var captureDescription v20211101s.CaptureDescription_Status
		err := eventhub.CaptureDescription.AssignPropertiesToCaptureDescriptionStatus(&captureDescription)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCaptureDescriptionStatus() to populate field CaptureDescription")
		}
		destination.CaptureDescription = &captureDescription
	} else {
		destination.CaptureDescription = nil
	}

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(eventhub.Conditions)

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(eventhub.CreatedAt)

	// Id
	destination.Id = genruntime.ClonePointerToString(eventhub.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(eventhub.Location)

	// MessageRetentionInDays
	destination.MessageRetentionInDays = genruntime.ClonePointerToInt(eventhub.MessageRetentionInDays)

	// Name
	destination.Name = genruntime.ClonePointerToString(eventhub.Name)

	// PartitionCount
	destination.PartitionCount = genruntime.ClonePointerToInt(eventhub.PartitionCount)

	// PartitionIds
	destination.PartitionIds = genruntime.CloneSliceOfString(eventhub.PartitionIds)

	// Status
	destination.Status = genruntime.ClonePointerToString(eventhub.Status)

	// SystemData
	if eventhub.SystemData != nil {
		var systemDatum v20211101s.SystemData_Status
		err := eventhub.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(eventhub.Type)

	// UpdatedAt
	destination.UpdatedAt = genruntime.ClonePointerToString(eventhub.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20211101.NamespacesEventhubs_Spec
type NamespacesEventhubs_Spec struct {
	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName              string                                                  `json:"azureName,omitempty"`
	CaptureDescription     *NamespacesEventhubs_Spec_Properties_CaptureDescription `json:"captureDescription,omitempty"`
	Location               *string                                                 `json:"location,omitempty"`
	MessageRetentionInDays *int                                                    `json:"messageRetentionInDays,omitempty"`
	OriginalVersion        string                                                  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/Namespace resource
	Owner          *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PartitionCount *int                               `json:"partitionCount,omitempty"`
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags           map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesEventhubs_Spec{}

// ConvertSpecFrom populates our NamespacesEventhubs_Spec from the provided source
func (eventhubs *NamespacesEventhubs_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.NamespacesEventhubs_Spec)
	if ok {
		// Populate our instance from source
		return eventhubs.AssignPropertiesFromNamespacesEventhubsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.NamespacesEventhubs_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = eventhubs.AssignPropertiesFromNamespacesEventhubsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesEventhubs_Spec
func (eventhubs *NamespacesEventhubs_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.NamespacesEventhubs_Spec)
	if ok {
		// Populate destination from our instance
		return eventhubs.AssignPropertiesToNamespacesEventhubsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.NamespacesEventhubs_Spec{}
	err := eventhubs.AssignPropertiesToNamespacesEventhubsSpec(dst)
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

// AssignPropertiesFromNamespacesEventhubsSpec populates our NamespacesEventhubs_Spec from the provided source NamespacesEventhubs_Spec
func (eventhubs *NamespacesEventhubs_Spec) AssignPropertiesFromNamespacesEventhubsSpec(source *v20211101s.NamespacesEventhubs_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	eventhubs.AzureName = source.AzureName

	// CaptureDescription
	if source.CaptureDescription != nil {
		var captureDescription NamespacesEventhubs_Spec_Properties_CaptureDescription
		err := captureDescription.AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescription(source.CaptureDescription)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescription() to populate field CaptureDescription")
		}
		eventhubs.CaptureDescription = &captureDescription
	} else {
		eventhubs.CaptureDescription = nil
	}

	// Location
	eventhubs.Location = genruntime.ClonePointerToString(source.Location)

	// MessageRetentionInDays
	eventhubs.MessageRetentionInDays = genruntime.ClonePointerToInt(source.MessageRetentionInDays)

	// OriginalVersion
	eventhubs.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		eventhubs.Owner = &owner
	} else {
		eventhubs.Owner = nil
	}

	// PartitionCount
	eventhubs.PartitionCount = genruntime.ClonePointerToInt(source.PartitionCount)

	// Tags
	eventhubs.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		eventhubs.PropertyBag = propertyBag
	} else {
		eventhubs.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsSpec populates the provided destination NamespacesEventhubs_Spec from our NamespacesEventhubs_Spec
func (eventhubs *NamespacesEventhubs_Spec) AssignPropertiesToNamespacesEventhubsSpec(destination *v20211101s.NamespacesEventhubs_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(eventhubs.PropertyBag)

	// AzureName
	destination.AzureName = eventhubs.AzureName

	// CaptureDescription
	if eventhubs.CaptureDescription != nil {
		var captureDescription v20211101s.NamespacesEventhubs_Spec_Properties_CaptureDescription
		err := eventhubs.CaptureDescription.AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescription(&captureDescription)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescription() to populate field CaptureDescription")
		}
		destination.CaptureDescription = &captureDescription
	} else {
		destination.CaptureDescription = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(eventhubs.Location)

	// MessageRetentionInDays
	destination.MessageRetentionInDays = genruntime.ClonePointerToInt(eventhubs.MessageRetentionInDays)

	// OriginalVersion
	destination.OriginalVersion = eventhubs.OriginalVersion

	// Owner
	if eventhubs.Owner != nil {
		owner := eventhubs.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// PartitionCount
	destination.PartitionCount = genruntime.ClonePointerToInt(eventhubs.PartitionCount)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(eventhubs.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20211101.CaptureDescription_Status
// Deprecated version of CaptureDescription_Status. Use v1beta20211101.CaptureDescription_Status instead
type CaptureDescription_Status struct {
	Destination       *Destination_Status    `json:"destination,omitempty"`
	Enabled           *bool                  `json:"enabled,omitempty"`
	Encoding          *string                `json:"encoding,omitempty"`
	IntervalInSeconds *int                   `json:"intervalInSeconds,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SizeLimitInBytes  *int                   `json:"sizeLimitInBytes,omitempty"`
	SkipEmptyArchives *bool                  `json:"skipEmptyArchives,omitempty"`
}

// AssignPropertiesFromCaptureDescriptionStatus populates our CaptureDescription_Status from the provided source CaptureDescription_Status
func (description *CaptureDescription_Status) AssignPropertiesFromCaptureDescriptionStatus(source *v20211101s.CaptureDescription_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Destination
	if source.Destination != nil {
		var destination Destination_Status
		err := destination.AssignPropertiesFromDestinationStatus(source.Destination)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromDestinationStatus() to populate field Destination")
		}
		description.Destination = &destination
	} else {
		description.Destination = nil
	}

	// Enabled
	if source.Enabled != nil {
		enabled := *source.Enabled
		description.Enabled = &enabled
	} else {
		description.Enabled = nil
	}

	// Encoding
	description.Encoding = genruntime.ClonePointerToString(source.Encoding)

	// IntervalInSeconds
	description.IntervalInSeconds = genruntime.ClonePointerToInt(source.IntervalInSeconds)

	// SizeLimitInBytes
	description.SizeLimitInBytes = genruntime.ClonePointerToInt(source.SizeLimitInBytes)

	// SkipEmptyArchives
	if source.SkipEmptyArchives != nil {
		skipEmptyArchive := *source.SkipEmptyArchives
		description.SkipEmptyArchives = &skipEmptyArchive
	} else {
		description.SkipEmptyArchives = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		description.PropertyBag = propertyBag
	} else {
		description.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToCaptureDescriptionStatus populates the provided destination CaptureDescription_Status from our CaptureDescription_Status
func (description *CaptureDescription_Status) AssignPropertiesToCaptureDescriptionStatus(destination *v20211101s.CaptureDescription_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(description.PropertyBag)

	// Destination
	if description.Destination != nil {
		var destinationLocal v20211101s.Destination_Status
		err := description.Destination.AssignPropertiesToDestinationStatus(&destinationLocal)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToDestinationStatus() to populate field Destination")
		}
		destination.Destination = &destinationLocal
	} else {
		destination.Destination = nil
	}

	// Enabled
	if description.Enabled != nil {
		enabled := *description.Enabled
		destination.Enabled = &enabled
	} else {
		destination.Enabled = nil
	}

	// Encoding
	destination.Encoding = genruntime.ClonePointerToString(description.Encoding)

	// IntervalInSeconds
	destination.IntervalInSeconds = genruntime.ClonePointerToInt(description.IntervalInSeconds)

	// SizeLimitInBytes
	destination.SizeLimitInBytes = genruntime.ClonePointerToInt(description.SizeLimitInBytes)

	// SkipEmptyArchives
	if description.SkipEmptyArchives != nil {
		skipEmptyArchive := *description.SkipEmptyArchives
		destination.SkipEmptyArchives = &skipEmptyArchive
	} else {
		destination.SkipEmptyArchives = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription
// Deprecated version of NamespacesEventhubs_Spec_Properties_CaptureDescription. Use v1beta20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription instead
type NamespacesEventhubs_Spec_Properties_CaptureDescription struct {
	Destination       *NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination `json:"destination,omitempty"`
	Enabled           *bool                                                               `json:"enabled,omitempty"`
	Encoding          *string                                                             `json:"encoding,omitempty"`
	IntervalInSeconds *int                                                                `json:"intervalInSeconds,omitempty"`
	PropertyBag       genruntime.PropertyBag                                              `json:"$propertyBag,omitempty"`
	SizeLimitInBytes  *int                                                                `json:"sizeLimitInBytes,omitempty"`
	SkipEmptyArchives *bool                                                               `json:"skipEmptyArchives,omitempty"`
}

// AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescription populates our NamespacesEventhubs_Spec_Properties_CaptureDescription from the provided source NamespacesEventhubs_Spec_Properties_CaptureDescription
func (description *NamespacesEventhubs_Spec_Properties_CaptureDescription) AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescription(source *v20211101s.NamespacesEventhubs_Spec_Properties_CaptureDescription) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Destination
	if source.Destination != nil {
		var destination NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination
		err := destination.AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination(source.Destination)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination() to populate field Destination")
		}
		description.Destination = &destination
	} else {
		description.Destination = nil
	}

	// Enabled
	if source.Enabled != nil {
		enabled := *source.Enabled
		description.Enabled = &enabled
	} else {
		description.Enabled = nil
	}

	// Encoding
	description.Encoding = genruntime.ClonePointerToString(source.Encoding)

	// IntervalInSeconds
	description.IntervalInSeconds = genruntime.ClonePointerToInt(source.IntervalInSeconds)

	// SizeLimitInBytes
	description.SizeLimitInBytes = genruntime.ClonePointerToInt(source.SizeLimitInBytes)

	// SkipEmptyArchives
	if source.SkipEmptyArchives != nil {
		skipEmptyArchive := *source.SkipEmptyArchives
		description.SkipEmptyArchives = &skipEmptyArchive
	} else {
		description.SkipEmptyArchives = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		description.PropertyBag = propertyBag
	} else {
		description.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescription populates the provided destination NamespacesEventhubs_Spec_Properties_CaptureDescription from our NamespacesEventhubs_Spec_Properties_CaptureDescription
func (description *NamespacesEventhubs_Spec_Properties_CaptureDescription) AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescription(destination *v20211101s.NamespacesEventhubs_Spec_Properties_CaptureDescription) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(description.PropertyBag)

	// Destination
	if description.Destination != nil {
		var destinationLocal v20211101s.NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination
		err := description.Destination.AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination(&destinationLocal)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination() to populate field Destination")
		}
		destination.Destination = &destinationLocal
	} else {
		destination.Destination = nil
	}

	// Enabled
	if description.Enabled != nil {
		enabled := *description.Enabled
		destination.Enabled = &enabled
	} else {
		destination.Enabled = nil
	}

	// Encoding
	destination.Encoding = genruntime.ClonePointerToString(description.Encoding)

	// IntervalInSeconds
	destination.IntervalInSeconds = genruntime.ClonePointerToInt(description.IntervalInSeconds)

	// SizeLimitInBytes
	destination.SizeLimitInBytes = genruntime.ClonePointerToInt(description.SizeLimitInBytes)

	// SkipEmptyArchives
	if description.SkipEmptyArchives != nil {
		skipEmptyArchive := *description.SkipEmptyArchives
		destination.SkipEmptyArchives = &skipEmptyArchive
	} else {
		destination.SkipEmptyArchives = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20211101.Destination_Status
// Deprecated version of Destination_Status. Use v1beta20211101.Destination_Status instead
type Destination_Status struct {
	ArchiveNameFormat        *string                `json:"archiveNameFormat,omitempty"`
	BlobContainer            *string                `json:"blobContainer,omitempty"`
	DataLakeAccountName      *string                `json:"dataLakeAccountName,omitempty"`
	DataLakeFolderPath       *string                `json:"dataLakeFolderPath,omitempty"`
	DataLakeSubscriptionId   *string                `json:"dataLakeSubscriptionId,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageAccountResourceId *string                `json:"storageAccountResourceId,omitempty"`
}

// AssignPropertiesFromDestinationStatus populates our Destination_Status from the provided source Destination_Status
func (destination *Destination_Status) AssignPropertiesFromDestinationStatus(source *v20211101s.Destination_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ArchiveNameFormat
	destination.ArchiveNameFormat = genruntime.ClonePointerToString(source.ArchiveNameFormat)

	// BlobContainer
	destination.BlobContainer = genruntime.ClonePointerToString(source.BlobContainer)

	// DataLakeAccountName
	destination.DataLakeAccountName = genruntime.ClonePointerToString(source.DataLakeAccountName)

	// DataLakeFolderPath
	destination.DataLakeFolderPath = genruntime.ClonePointerToString(source.DataLakeFolderPath)

	// DataLakeSubscriptionId
	destination.DataLakeSubscriptionId = genruntime.ClonePointerToString(source.DataLakeSubscriptionId)

	// Name
	destination.Name = genruntime.ClonePointerToString(source.Name)

	// StorageAccountResourceId
	destination.StorageAccountResourceId = genruntime.ClonePointerToString(source.StorageAccountResourceId)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDestinationStatus populates the provided destination Destination_Status from our Destination_Status
func (destination *Destination_Status) AssignPropertiesToDestinationStatus(target *v20211101s.Destination_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(destination.PropertyBag)

	// ArchiveNameFormat
	target.ArchiveNameFormat = genruntime.ClonePointerToString(destination.ArchiveNameFormat)

	// BlobContainer
	target.BlobContainer = genruntime.ClonePointerToString(destination.BlobContainer)

	// DataLakeAccountName
	target.DataLakeAccountName = genruntime.ClonePointerToString(destination.DataLakeAccountName)

	// DataLakeFolderPath
	target.DataLakeFolderPath = genruntime.ClonePointerToString(destination.DataLakeFolderPath)

	// DataLakeSubscriptionId
	target.DataLakeSubscriptionId = genruntime.ClonePointerToString(destination.DataLakeSubscriptionId)

	// Name
	target.Name = genruntime.ClonePointerToString(destination.Name)

	// StorageAccountResourceId
	target.StorageAccountResourceId = genruntime.ClonePointerToString(destination.StorageAccountResourceId)

	// Update the property bag
	if len(propertyBag) > 0 {
		target.PropertyBag = propertyBag
	} else {
		target.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination
// Deprecated version of NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination. Use v1beta20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination instead
type NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination struct {
	ArchiveNameFormat               *string                       `json:"archiveNameFormat,omitempty"`
	BlobContainer                   *string                       `json:"blobContainer,omitempty"`
	DataLakeAccountName             *string                       `json:"dataLakeAccountName,omitempty"`
	DataLakeFolderPath              *string                       `json:"dataLakeFolderPath,omitempty"`
	DataLakeSubscriptionId          *string                       `json:"dataLakeSubscriptionId,omitempty"`
	Name                            *string                       `json:"name,omitempty"`
	PropertyBag                     genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	StorageAccountResourceReference *genruntime.ResourceReference `armReference:"StorageAccountResourceId" json:"storageAccountResourceReference,omitempty"`
}

// AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination populates our NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination from the provided source NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination
func (destination *NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination) AssignPropertiesFromNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination(source *v20211101s.NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ArchiveNameFormat
	destination.ArchiveNameFormat = genruntime.ClonePointerToString(source.ArchiveNameFormat)

	// BlobContainer
	destination.BlobContainer = genruntime.ClonePointerToString(source.BlobContainer)

	// DataLakeAccountName
	destination.DataLakeAccountName = genruntime.ClonePointerToString(source.DataLakeAccountName)

	// DataLakeFolderPath
	destination.DataLakeFolderPath = genruntime.ClonePointerToString(source.DataLakeFolderPath)

	// DataLakeSubscriptionId
	destination.DataLakeSubscriptionId = genruntime.ClonePointerToString(source.DataLakeSubscriptionId)

	// Name
	destination.Name = genruntime.ClonePointerToString(source.Name)

	// StorageAccountResourceReference
	if source.StorageAccountResourceReference != nil {
		storageAccountResourceReference := source.StorageAccountResourceReference.Copy()
		destination.StorageAccountResourceReference = &storageAccountResourceReference
	} else {
		destination.StorageAccountResourceReference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination populates the provided destination NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination from our NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination
func (destination *NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination) AssignPropertiesToNamespacesEventhubsSpecPropertiesCaptureDescriptionDestination(target *v20211101s.NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(destination.PropertyBag)

	// ArchiveNameFormat
	target.ArchiveNameFormat = genruntime.ClonePointerToString(destination.ArchiveNameFormat)

	// BlobContainer
	target.BlobContainer = genruntime.ClonePointerToString(destination.BlobContainer)

	// DataLakeAccountName
	target.DataLakeAccountName = genruntime.ClonePointerToString(destination.DataLakeAccountName)

	// DataLakeFolderPath
	target.DataLakeFolderPath = genruntime.ClonePointerToString(destination.DataLakeFolderPath)

	// DataLakeSubscriptionId
	target.DataLakeSubscriptionId = genruntime.ClonePointerToString(destination.DataLakeSubscriptionId)

	// Name
	target.Name = genruntime.ClonePointerToString(destination.Name)

	// StorageAccountResourceReference
	if destination.StorageAccountResourceReference != nil {
		storageAccountResourceReference := destination.StorageAccountResourceReference.Copy()
		target.StorageAccountResourceReference = &storageAccountResourceReference
	} else {
		target.StorageAccountResourceReference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		target.PropertyBag = propertyBag
	} else {
		target.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&NamespacesEventhub{}, &NamespacesEventhubList{})
}
