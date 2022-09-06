// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventhub.azure.com,resources=namespaceseventhubs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventhub.azure.com,resources={namespaceseventhubs/status,namespaceseventhubs/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20211101.NamespacesEventhub
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs
type NamespacesEventhub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Eventhubs_Spec `json:"spec,omitempty"`
	Status            Eventhub_STATUS           `json:"status,omitempty"`
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

var _ genruntime.KubernetesResource = &NamespacesEventhub{}

// AzureName returns the Azure name of the resource
func (eventhub *NamespacesEventhub) AzureName() string {
	return eventhub.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (eventhub NamespacesEventhub) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (eventhub *NamespacesEventhub) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
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
	return &Eventhub_STATUS{}
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
	if st, ok := status.(*Eventhub_STATUS); ok {
		eventhub.Status = *st
		return nil
	}

	// Convert status to required version
	var st Eventhub_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	eventhub.Status = st
	return nil
}

// Hub marks that this NamespacesEventhub is the hub type for conversion
func (eventhub *NamespacesEventhub) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (eventhub *NamespacesEventhub) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: eventhub.Spec.OriginalVersion,
		Kind:    "NamespacesEventhub",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20211101.NamespacesEventhub
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs
type NamespacesEventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhub `json:"items"`
}

// Storage version of v1beta20211101.Eventhub_STATUS
type Eventhub_STATUS struct {
	CaptureDescription     *CaptureDescription_STATUS `json:"captureDescription,omitempty"`
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
	SystemData             *SystemData_STATUS         `json:"systemData,omitempty"`
	Type                   *string                    `json:"type,omitempty"`
	UpdatedAt              *string                    `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Eventhub_STATUS{}

// ConvertStatusFrom populates our Eventhub_STATUS from the provided source
func (eventhub *Eventhub_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == eventhub {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(eventhub)
}

// ConvertStatusTo populates the provided destination from our Eventhub_STATUS
func (eventhub *Eventhub_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == eventhub {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(eventhub)
}

// Storage version of v1beta20211101.Namespaces_Eventhubs_Spec
type Namespaces_Eventhubs_Spec struct {
	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName              string                                                   `json:"azureName,omitempty"`
	CaptureDescription     *Namespaces_Eventhubs_Spec_Properties_CaptureDescription `json:"captureDescription,omitempty"`
	Location               *string                                                  `json:"location,omitempty"`
	MessageRetentionInDays *int                                                     `json:"messageRetentionInDays,omitempty"`
	OriginalVersion        string                                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/Namespace resource
	Owner          *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PartitionCount *int                               `json:"partitionCount,omitempty"`
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags           map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Namespaces_Eventhubs_Spec{}

// ConvertSpecFrom populates our Namespaces_Eventhubs_Spec from the provided source
func (eventhubs *Namespaces_Eventhubs_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == eventhubs {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(eventhubs)
}

// ConvertSpecTo populates the provided destination from our Namespaces_Eventhubs_Spec
func (eventhubs *Namespaces_Eventhubs_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == eventhubs {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(eventhubs)
}

// Storage version of v1beta20211101.CaptureDescription_STATUS
type CaptureDescription_STATUS struct {
	Destination       *Destination_STATUS    `json:"destination,omitempty"`
	Enabled           *bool                  `json:"enabled,omitempty"`
	Encoding          *string                `json:"encoding,omitempty"`
	IntervalInSeconds *int                   `json:"intervalInSeconds,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SizeLimitInBytes  *int                   `json:"sizeLimitInBytes,omitempty"`
	SkipEmptyArchives *bool                  `json:"skipEmptyArchives,omitempty"`
}

// Storage version of v1beta20211101.Namespaces_Eventhubs_Spec_Properties_CaptureDescription
type Namespaces_Eventhubs_Spec_Properties_CaptureDescription struct {
	Destination       *Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination `json:"destination,omitempty"`
	Enabled           *bool                                                                `json:"enabled,omitempty"`
	Encoding          *string                                                              `json:"encoding,omitempty"`
	IntervalInSeconds *int                                                                 `json:"intervalInSeconds,omitempty"`
	PropertyBag       genruntime.PropertyBag                                               `json:"$propertyBag,omitempty"`
	SizeLimitInBytes  *int                                                                 `json:"sizeLimitInBytes,omitempty"`
	SkipEmptyArchives *bool                                                                `json:"skipEmptyArchives,omitempty"`
}

// Storage version of v1beta20211101.Destination_STATUS
type Destination_STATUS struct {
	ArchiveNameFormat        *string                `json:"archiveNameFormat,omitempty"`
	BlobContainer            *string                `json:"blobContainer,omitempty"`
	DataLakeAccountName      *string                `json:"dataLakeAccountName,omitempty"`
	DataLakeFolderPath       *string                `json:"dataLakeFolderPath,omitempty"`
	DataLakeSubscriptionId   *string                `json:"dataLakeSubscriptionId,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageAccountResourceId *string                `json:"storageAccountResourceId,omitempty"`
}

// Storage version of v1beta20211101.Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination
type Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination struct {
	ArchiveNameFormat      *string                `json:"archiveNameFormat,omitempty"`
	BlobContainer          *string                `json:"blobContainer,omitempty"`
	DataLakeAccountName    *string                `json:"dataLakeAccountName,omitempty"`
	DataLakeFolderPath     *string                `json:"dataLakeFolderPath,omitempty"`
	DataLakeSubscriptionId *string                `json:"dataLakeSubscriptionId,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// StorageAccountResourceReference: Resource id of the storage account to be used to create the blobs
	StorageAccountResourceReference *genruntime.ResourceReference `armReference:"StorageAccountResourceId" json:"storageAccountResourceReference,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NamespacesEventhub{}, &NamespacesEventhubList{})
}
