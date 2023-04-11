// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210401storage

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
// Storage version of v1api20210401.StorageAccountsBlobServicesContainer
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/blob.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}
type StorageAccountsBlobServicesContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccounts_BlobServices_Container_Spec   `json:"spec,omitempty"`
	Status            StorageAccounts_BlobServices_Container_STATUS `json:"status,omitempty"`
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
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (container *StorageAccountsBlobServicesContainer) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
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
	return &StorageAccounts_BlobServices_Container_STATUS{}
}

// Owner returns the ResourceReference of the owner
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
	if st, ok := status.(*StorageAccounts_BlobServices_Container_STATUS); ok {
		container.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccounts_BlobServices_Container_STATUS
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
// Storage version of v1api20210401.StorageAccountsBlobServicesContainer
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/blob.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}
type StorageAccountsBlobServicesContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsBlobServicesContainer `json:"items"`
}

// Storage version of v1api20210401.StorageAccounts_BlobServices_Container_Spec
type StorageAccounts_BlobServices_Container_Spec struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:MinLength=3
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                      string                          `json:"azureName,omitempty"`
	DefaultEncryptionScope         *string                         `json:"defaultEncryptionScope,omitempty"`
	DenyEncryptionScopeOverride    *bool                           `json:"denyEncryptionScopeOverride,omitempty"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning `json:"immutableStorageWithVersioning,omitempty"`
	Metadata                       map[string]string               `json:"metadata,omitempty"`
	OriginalVersion                string                          `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccountsBlobService resource
	Owner        *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccountsBlobService"`
	PropertyBag  genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicAccess *string                            `json:"publicAccess,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccounts_BlobServices_Container_Spec{}

// ConvertSpecFrom populates our StorageAccounts_BlobServices_Container_Spec from the provided source
func (container *StorageAccounts_BlobServices_Container_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == container {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(container)
}

// ConvertSpecTo populates the provided destination from our StorageAccounts_BlobServices_Container_Spec
func (container *StorageAccounts_BlobServices_Container_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == container {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(container)
}

// Storage version of v1api20210401.StorageAccounts_BlobServices_Container_STATUS
type StorageAccounts_BlobServices_Container_STATUS struct {
	Conditions                     []conditions.Condition                 `json:"conditions,omitempty"`
	DefaultEncryptionScope         *string                                `json:"defaultEncryptionScope,omitempty"`
	Deleted                        *bool                                  `json:"deleted,omitempty"`
	DeletedTime                    *string                                `json:"deletedTime,omitempty"`
	DenyEncryptionScopeOverride    *bool                                  `json:"denyEncryptionScopeOverride,omitempty"`
	Etag                           *string                                `json:"etag,omitempty"`
	HasImmutabilityPolicy          *bool                                  `json:"hasImmutabilityPolicy,omitempty"`
	HasLegalHold                   *bool                                  `json:"hasLegalHold,omitempty"`
	Id                             *string                                `json:"id,omitempty"`
	ImmutabilityPolicy             *ImmutabilityPolicyProperties_STATUS   `json:"immutabilityPolicy,omitempty"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning_STATUS `json:"immutableStorageWithVersioning,omitempty"`
	LastModifiedTime               *string                                `json:"lastModifiedTime,omitempty"`
	LeaseDuration                  *string                                `json:"leaseDuration,omitempty"`
	LeaseState                     *string                                `json:"leaseState,omitempty"`
	LeaseStatus                    *string                                `json:"leaseStatus,omitempty"`
	LegalHold                      *LegalHoldProperties_STATUS            `json:"legalHold,omitempty"`
	Metadata                       map[string]string                      `json:"metadata,omitempty"`
	Name                           *string                                `json:"name,omitempty"`
	PropertyBag                    genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	PublicAccess                   *string                                `json:"publicAccess,omitempty"`
	RemainingRetentionDays         *int                                   `json:"remainingRetentionDays,omitempty"`
	Type                           *string                                `json:"type,omitempty"`
	Version                        *string                                `json:"version,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccounts_BlobServices_Container_STATUS{}

// ConvertStatusFrom populates our StorageAccounts_BlobServices_Container_STATUS from the provided source
func (container *StorageAccounts_BlobServices_Container_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == container {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(container)
}

// ConvertStatusTo populates the provided destination from our StorageAccounts_BlobServices_Container_STATUS
func (container *StorageAccounts_BlobServices_Container_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == container {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(container)
}

// Storage version of v1api20210401.ImmutabilityPolicyProperties_STATUS
// The properties of an ImmutabilityPolicy of a blob container.
type ImmutabilityPolicyProperties_STATUS struct {
	AllowProtectedAppendWrites            *bool                          `json:"allowProtectedAppendWrites,omitempty"`
	Etag                                  *string                        `json:"etag,omitempty"`
	ImmutabilityPeriodSinceCreationInDays *int                           `json:"immutabilityPeriodSinceCreationInDays,omitempty"`
	PropertyBag                           genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	State                                 *string                        `json:"state,omitempty"`
	UpdateHistory                         []UpdateHistoryProperty_STATUS `json:"updateHistory,omitempty"`
}

// Storage version of v1api20210401.ImmutableStorageWithVersioning
// Object level immutability properties of the container.
type ImmutableStorageWithVersioning struct {
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210401.ImmutableStorageWithVersioning_STATUS
// Object level immutability properties of the container.
type ImmutableStorageWithVersioning_STATUS struct {
	Enabled        *bool                  `json:"enabled,omitempty"`
	MigrationState *string                `json:"migrationState,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TimeStamp      *string                `json:"timeStamp,omitempty"`
}

// Storage version of v1api20210401.LegalHoldProperties_STATUS
// The LegalHold property of a blob container.
type LegalHoldProperties_STATUS struct {
	HasLegalHold *bool                  `json:"hasLegalHold,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tags         []TagProperty_STATUS   `json:"tags,omitempty"`
}

// Storage version of v1api20210401.TagProperty_STATUS
// A tag of the LegalHold of a blob container.
type TagProperty_STATUS struct {
	ObjectIdentifier *string                `json:"objectIdentifier,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tag              *string                `json:"tag,omitempty"`
	TenantId         *string                `json:"tenantId,omitempty"`
	Timestamp        *string                `json:"timestamp,omitempty"`
	Upn              *string                `json:"upn,omitempty"`
}

// Storage version of v1api20210401.UpdateHistoryProperty_STATUS
// An update history of the ImmutabilityPolicy of a blob container.
type UpdateHistoryProperty_STATUS struct {
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
