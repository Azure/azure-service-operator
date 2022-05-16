// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccountsblobservicescontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccountsblobservicescontainers/status,storageaccountsblobservicescontainers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210401.StorageAccountsBlobServicesContainer
// Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts_blobServices_containers
type StorageAccountsBlobServicesContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsBlobServicesContainers_Spec `json:"spec,omitempty"`
	Status            BlobContainer_Status                       `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsBlobServicesContainer{}

// GetConditions returns the conditions of the resource
func (container *StorageAccountsBlobServicesContainer) GetConditions() conditions.Conditions {
	return container.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (container *StorageAccountsBlobServicesContainer) SetConditions(conditions conditions.Conditions) {
	container.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &StorageAccountsBlobServicesContainer{}

// AzureName returns the Azure name of the resource
func (container *StorageAccountsBlobServicesContainer) AzureName() string {
	return container.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (container StorageAccountsBlobServicesContainer) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (container *StorageAccountsBlobServicesContainer) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (container *StorageAccountsBlobServicesContainer) GetSpec() genruntime.ConvertibleSpec {
	return &container.Spec
}

// GetStatus returns the status of this resource
func (container *StorageAccountsBlobServicesContainer) GetStatus() genruntime.ConvertibleStatus {
	return &container.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices/containers"
func (container *StorageAccountsBlobServicesContainer) GetType() string {
	return "Microsoft.Storage/storageAccounts/blobServices/containers"
}

// NewEmptyStatus returns a new empty (blank) status
func (container *StorageAccountsBlobServicesContainer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &BlobContainer_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (container *StorageAccountsBlobServicesContainer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(container.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  container.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (container *StorageAccountsBlobServicesContainer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*BlobContainer_Status); ok {
		container.Status = *st
		return nil
	}

	// Convert status to required version
	var st BlobContainer_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	container.Status = st
	return nil
}

// Hub marks that this StorageAccountsBlobServicesContainer is the hub type for conversion
func (container *StorageAccountsBlobServicesContainer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (container *StorageAccountsBlobServicesContainer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: container.Spec.OriginalVersion,
		Kind:    "StorageAccountsBlobServicesContainer",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210401.StorageAccountsBlobServicesContainer
// Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts_blobServices_containers
type StorageAccountsBlobServicesContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsBlobServicesContainer `json:"items"`
}

// Storage version of v1beta20210401.BlobContainer_Status
type BlobContainer_Status struct {
	Conditions                     []conditions.Condition                 `json:"conditions,omitempty"`
	DefaultEncryptionScope         *string                                `json:"defaultEncryptionScope,omitempty"`
	Deleted                        *bool                                  `json:"deleted,omitempty"`
	DeletedTime                    *string                                `json:"deletedTime,omitempty"`
	DenyEncryptionScopeOverride    *bool                                  `json:"denyEncryptionScopeOverride,omitempty"`
	Etag                           *string                                `json:"etag,omitempty"`
	HasImmutabilityPolicy          *bool                                  `json:"hasImmutabilityPolicy,omitempty"`
	HasLegalHold                   *bool                                  `json:"hasLegalHold,omitempty"`
	Id                             *string                                `json:"id,omitempty"`
	ImmutabilityPolicy             *ImmutabilityPolicyProperties_Status   `json:"immutabilityPolicy,omitempty"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning_Status `json:"immutableStorageWithVersioning,omitempty"`
	LastModifiedTime               *string                                `json:"lastModifiedTime,omitempty"`
	LeaseDuration                  *string                                `json:"leaseDuration,omitempty"`
	LeaseState                     *string                                `json:"leaseState,omitempty"`
	LeaseStatus                    *string                                `json:"leaseStatus,omitempty"`
	LegalHold                      *LegalHoldProperties_Status            `json:"legalHold,omitempty"`
	Metadata                       map[string]string                      `json:"metadata,omitempty"`
	Name                           *string                                `json:"name,omitempty"`
	PropertyBag                    genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	PublicAccess                   *string                                `json:"publicAccess,omitempty"`
	RemainingRetentionDays         *int                                   `json:"remainingRetentionDays,omitempty"`
	Type                           *string                                `json:"type,omitempty"`
	Version                        *string                                `json:"version,omitempty"`
}

var _ genruntime.ConvertibleStatus = &BlobContainer_Status{}

// ConvertStatusFrom populates our BlobContainer_Status from the provided source
func (container *BlobContainer_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == container {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(container)
}

// ConvertStatusTo populates the provided destination from our BlobContainer_Status
func (container *BlobContainer_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == container {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(container)
}

// Storage version of v1beta20210401.StorageAccountsBlobServicesContainers_Spec
type StorageAccountsBlobServicesContainers_Spec struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:MinLength=3
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                      string                          `json:"azureName,omitempty"`
	DefaultEncryptionScope         *string                         `json:"defaultEncryptionScope,omitempty"`
	DenyEncryptionScopeOverride    *bool                           `json:"denyEncryptionScopeOverride,omitempty"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning `json:"immutableStorageWithVersioning,omitempty"`
	Location                       *string                         `json:"location,omitempty"`
	Metadata                       map[string]string               `json:"metadata,omitempty"`
	OriginalVersion                string                          `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccountsBlobService resource
	Owner        *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccountsBlobService"`
	PropertyBag  genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicAccess *string                            `json:"publicAccess,omitempty"`
	Tags         map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsBlobServicesContainers_Spec{}

// ConvertSpecFrom populates our StorageAccountsBlobServicesContainers_Spec from the provided source
func (containers *StorageAccountsBlobServicesContainers_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == containers {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(containers)
}

// ConvertSpecTo populates the provided destination from our StorageAccountsBlobServicesContainers_Spec
func (containers *StorageAccountsBlobServicesContainers_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == containers {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(containers)
}

// Storage version of v1beta20210401.ImmutabilityPolicyProperties_Status
type ImmutabilityPolicyProperties_Status struct {
	AllowProtectedAppendWrites            *bool                          `json:"allowProtectedAppendWrites,omitempty"`
	Etag                                  *string                        `json:"etag,omitempty"`
	ImmutabilityPeriodSinceCreationInDays *int                           `json:"immutabilityPeriodSinceCreationInDays,omitempty"`
	PropertyBag                           genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	State                                 *string                        `json:"state,omitempty"`
	UpdateHistory                         []UpdateHistoryProperty_Status `json:"updateHistory,omitempty"`
}

// Storage version of v1beta20210401.ImmutableStorageWithVersioning
// Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ImmutableStorageWithVersioning
type ImmutableStorageWithVersioning struct {
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210401.ImmutableStorageWithVersioning_Status
type ImmutableStorageWithVersioning_Status struct {
	Enabled        *bool                  `json:"enabled,omitempty"`
	MigrationState *string                `json:"migrationState,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TimeStamp      *string                `json:"timeStamp,omitempty"`
}

// Storage version of v1beta20210401.LegalHoldProperties_Status
type LegalHoldProperties_Status struct {
	HasLegalHold *bool                  `json:"hasLegalHold,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tags         []TagProperty_Status   `json:"tags,omitempty"`
}

// Storage version of v1beta20210401.TagProperty_Status
type TagProperty_Status struct {
	ObjectIdentifier *string                `json:"objectIdentifier,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tag              *string                `json:"tag,omitempty"`
	TenantId         *string                `json:"tenantId,omitempty"`
	Timestamp        *string                `json:"timestamp,omitempty"`
	Upn              *string                `json:"upn,omitempty"`
}

// Storage version of v1beta20210401.UpdateHistoryProperty_Status
type UpdateHistoryProperty_Status struct {
	ImmutabilityPeriodSinceCreationInDays *int                   `json:"immutabilityPeriodSinceCreationInDays,omitempty"`
	ObjectIdentifier                      *string                `json:"objectIdentifier,omitempty"`
	PropertyBag                           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId                              *string                `json:"tenantId,omitempty"`
	Timestamp                             *string                `json:"timestamp,omitempty"`
	Update                                *string                `json:"update,omitempty"`
	Upn                                   *string                `json:"upn,omitempty"`
}

func init() {
	SchemeBuilder.Register(&StorageAccountsBlobServicesContainer{}, &StorageAccountsBlobServicesContainerList{})
}
