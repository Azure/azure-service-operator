// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930storage

import (
	"fmt"
	alpha20201201s "github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20201201storage"
	alpha20210701s "github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20210701storage"
	v20200930s "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20200930storage"
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
// Storage version of v1alpha1api20200930.Snapshot
// Deprecated version of Snapshot. Use v1beta20200930.Snapshot instead
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Snapshots_Spec  `json:"spec,omitempty"`
	Status            Snapshot_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Snapshot{}

// GetConditions returns the conditions of the resource
func (snapshot *Snapshot) GetConditions() conditions.Conditions {
	return snapshot.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (snapshot *Snapshot) SetConditions(conditions conditions.Conditions) {
	snapshot.Status.Conditions = conditions
}

var _ conversion.Convertible = &Snapshot{}

// ConvertFrom populates our Snapshot from the provided hub Snapshot
func (snapshot *Snapshot) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20200930s.Snapshot)
	if !ok {
		return fmt.Errorf("expected compute/v1beta20200930storage/Snapshot but received %T instead", hub)
	}

	return snapshot.AssignPropertiesFromSnapshot(source)
}

// ConvertTo populates the provided hub Snapshot from our Snapshot
func (snapshot *Snapshot) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20200930s.Snapshot)
	if !ok {
		return fmt.Errorf("expected compute/v1beta20200930storage/Snapshot but received %T instead", hub)
	}

	return snapshot.AssignPropertiesToSnapshot(destination)
}

var _ genruntime.KubernetesResource = &Snapshot{}

// AzureName returns the Azure name of the resource
func (snapshot *Snapshot) AzureName() string {
	return snapshot.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-09-30"
func (snapshot Snapshot) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceScope returns the scope of the resource
func (snapshot *Snapshot) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (snapshot *Snapshot) GetSpec() genruntime.ConvertibleSpec {
	return &snapshot.Spec
}

// GetStatus returns the status of this resource
func (snapshot *Snapshot) GetStatus() genruntime.ConvertibleStatus {
	return &snapshot.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/snapshots"
func (snapshot *Snapshot) GetType() string {
	return "Microsoft.Compute/snapshots"
}

// NewEmptyStatus returns a new empty (blank) status
func (snapshot *Snapshot) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Snapshot_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (snapshot *Snapshot) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(snapshot.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  snapshot.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (snapshot *Snapshot) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Snapshot_Status); ok {
		snapshot.Status = *st
		return nil
	}

	// Convert status to required version
	var st Snapshot_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	snapshot.Status = st
	return nil
}

// AssignPropertiesFromSnapshot populates our Snapshot from the provided source Snapshot
func (snapshot *Snapshot) AssignPropertiesFromSnapshot(source *v20200930s.Snapshot) error {

	// ObjectMeta
	snapshot.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Snapshots_Spec
	err := spec.AssignPropertiesFromSnapshotsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSnapshotsSpec() to populate field Spec")
	}
	snapshot.Spec = spec

	// Status
	var status Snapshot_Status
	err = status.AssignPropertiesFromSnapshotStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSnapshotStatus() to populate field Status")
	}
	snapshot.Status = status

	// No error
	return nil
}

// AssignPropertiesToSnapshot populates the provided destination Snapshot from our Snapshot
func (snapshot *Snapshot) AssignPropertiesToSnapshot(destination *v20200930s.Snapshot) error {

	// ObjectMeta
	destination.ObjectMeta = *snapshot.ObjectMeta.DeepCopy()

	// Spec
	var spec v20200930s.Snapshots_Spec
	err := snapshot.Spec.AssignPropertiesToSnapshotsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSnapshotsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20200930s.Snapshot_Status
	err = snapshot.Status.AssignPropertiesToSnapshotStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSnapshotStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (snapshot *Snapshot) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: snapshot.Spec.OriginalVersion,
		Kind:    "Snapshot",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20200930.Snapshot
// Deprecated version of Snapshot. Use v1beta20200930.Snapshot instead
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Storage version of v1alpha1api20200930.Snapshot_Status
// Deprecated version of Snapshot_Status. Use v1beta20200930.Snapshot_Status instead
type Snapshot_Status struct {
	Conditions                   []conditions.Condition               `json:"conditions,omitempty"`
	CreationData                 *CreationData_Status                 `json:"creationData,omitempty"`
	DiskAccessId                 *string                              `json:"diskAccessId,omitempty"`
	DiskSizeBytes                *int                                 `json:"diskSizeBytes,omitempty"`
	DiskSizeGB                   *int                                 `json:"diskSizeGB,omitempty"`
	DiskState                    *string                              `json:"diskState,omitempty"`
	Encryption                   *Encryption_Status                   `json:"encryption,omitempty"`
	EncryptionSettingsCollection *EncryptionSettingsCollection_Status `json:"encryptionSettingsCollection,omitempty"`
	ExtendedLocation             *ExtendedLocation_Status             `json:"extendedLocation,omitempty"`
	HyperVGeneration             *string                              `json:"hyperVGeneration,omitempty"`
	Id                           *string                              `json:"id,omitempty"`
	Incremental                  *bool                                `json:"incremental,omitempty"`
	Location                     *string                              `json:"location,omitempty"`
	ManagedBy                    *string                              `json:"managedBy,omitempty"`
	Name                         *string                              `json:"name,omitempty"`
	NetworkAccessPolicy          *string                              `json:"networkAccessPolicy,omitempty"`
	OsType                       *string                              `json:"osType,omitempty"`
	PropertyBag                  genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                              `json:"provisioningState,omitempty"`
	PurchasePlan                 *PurchasePlan_Status                 `json:"purchasePlan,omitempty"`
	Sku                          *SnapshotSku_Status                  `json:"sku,omitempty"`
	Tags                         map[string]string                    `json:"tags,omitempty"`
	TimeCreated                  *string                              `json:"timeCreated,omitempty"`
	Type                         *string                              `json:"type,omitempty"`
	UniqueId                     *string                              `json:"uniqueId,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Snapshot_Status{}

// ConvertStatusFrom populates our Snapshot_Status from the provided source
func (snapshot *Snapshot_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20200930s.Snapshot_Status)
	if ok {
		// Populate our instance from source
		return snapshot.AssignPropertiesFromSnapshotStatus(src)
	}

	// Convert to an intermediate form
	src = &v20200930s.Snapshot_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = snapshot.AssignPropertiesFromSnapshotStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Snapshot_Status
func (snapshot *Snapshot_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20200930s.Snapshot_Status)
	if ok {
		// Populate destination from our instance
		return snapshot.AssignPropertiesToSnapshotStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20200930s.Snapshot_Status{}
	err := snapshot.AssignPropertiesToSnapshotStatus(dst)
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

// AssignPropertiesFromSnapshotStatus populates our Snapshot_Status from the provided source Snapshot_Status
func (snapshot *Snapshot_Status) AssignPropertiesFromSnapshotStatus(source *v20200930s.Snapshot_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	snapshot.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreationData
	if source.CreationData != nil {
		var creationDatum CreationData_Status
		err := creationDatum.AssignPropertiesFromCreationDataStatus(source.CreationData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreationDataStatus() to populate field CreationData")
		}
		snapshot.CreationData = &creationDatum
	} else {
		snapshot.CreationData = nil
	}

	// DiskAccessId
	snapshot.DiskAccessId = genruntime.ClonePointerToString(source.DiskAccessId)

	// DiskSizeBytes
	snapshot.DiskSizeBytes = genruntime.ClonePointerToInt(source.DiskSizeBytes)

	// DiskSizeGB
	snapshot.DiskSizeGB = genruntime.ClonePointerToInt(source.DiskSizeGB)

	// DiskState
	snapshot.DiskState = genruntime.ClonePointerToString(source.DiskState)

	// Encryption
	if source.Encryption != nil {
		var encryption Encryption_Status
		err := encryption.AssignPropertiesFromEncryptionStatus(source.Encryption)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromEncryptionStatus() to populate field Encryption")
		}
		snapshot.Encryption = &encryption
	} else {
		snapshot.Encryption = nil
	}

	// EncryptionSettingsCollection
	if source.EncryptionSettingsCollection != nil {
		var encryptionSettingsCollection EncryptionSettingsCollection_Status
		err := encryptionSettingsCollection.AssignPropertiesFromEncryptionSettingsCollectionStatus(source.EncryptionSettingsCollection)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromEncryptionSettingsCollectionStatus() to populate field EncryptionSettingsCollection")
		}
		snapshot.EncryptionSettingsCollection = &encryptionSettingsCollection
	} else {
		snapshot.EncryptionSettingsCollection = nil
	}

	// ExtendedLocation
	if source.ExtendedLocation != nil {
		var extendedLocationStatusStash alpha20210701s.ExtendedLocation_Status
		err := extendedLocationStatusStash.AssignPropertiesFromExtendedLocationStatus(source.ExtendedLocation)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromExtendedLocationStatus() to populate field ExtendedLocation_StatusStash from ExtendedLocation")
		}
		var extendedLocationStatusStashLocal alpha20201201s.ExtendedLocation_Status
		err = extendedLocationStatusStashLocal.AssignPropertiesFromExtendedLocationStatus(&extendedLocationStatusStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromExtendedLocationStatus() to populate field ExtendedLocation_StatusStash")
		}
		var extendedLocation ExtendedLocation_Status
		err = extendedLocation.AssignPropertiesFromExtendedLocationStatus(&extendedLocationStatusStashLocal)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromExtendedLocationStatus() to populate field ExtendedLocation from ExtendedLocation_StatusStash")
		}
		snapshot.ExtendedLocation = &extendedLocation
	} else {
		snapshot.ExtendedLocation = nil
	}

	// HyperVGeneration
	snapshot.HyperVGeneration = genruntime.ClonePointerToString(source.HyperVGeneration)

	// Id
	snapshot.Id = genruntime.ClonePointerToString(source.Id)

	// Incremental
	if source.Incremental != nil {
		incremental := *source.Incremental
		snapshot.Incremental = &incremental
	} else {
		snapshot.Incremental = nil
	}

	// Location
	snapshot.Location = genruntime.ClonePointerToString(source.Location)

	// ManagedBy
	snapshot.ManagedBy = genruntime.ClonePointerToString(source.ManagedBy)

	// Name
	snapshot.Name = genruntime.ClonePointerToString(source.Name)

	// NetworkAccessPolicy
	snapshot.NetworkAccessPolicy = genruntime.ClonePointerToString(source.NetworkAccessPolicy)

	// OsType
	snapshot.OsType = genruntime.ClonePointerToString(source.OsType)

	// ProvisioningState
	snapshot.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// PurchasePlan
	if source.PurchasePlan != nil {
		var purchasePlan PurchasePlan_Status
		err := purchasePlan.AssignPropertiesFromPurchasePlanStatus(source.PurchasePlan)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromPurchasePlanStatus() to populate field PurchasePlan")
		}
		snapshot.PurchasePlan = &purchasePlan
	} else {
		snapshot.PurchasePlan = nil
	}

	// Sku
	if source.Sku != nil {
		var sku SnapshotSku_Status
		err := sku.AssignPropertiesFromSnapshotSkuStatus(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSnapshotSkuStatus() to populate field Sku")
		}
		snapshot.Sku = &sku
	} else {
		snapshot.Sku = nil
	}

	// Tags
	snapshot.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// TimeCreated
	snapshot.TimeCreated = genruntime.ClonePointerToString(source.TimeCreated)

	// Type
	snapshot.Type = genruntime.ClonePointerToString(source.Type)

	// UniqueId
	snapshot.UniqueId = genruntime.ClonePointerToString(source.UniqueId)

	// Update the property bag
	if len(propertyBag) > 0 {
		snapshot.PropertyBag = propertyBag
	} else {
		snapshot.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSnapshotStatus populates the provided destination Snapshot_Status from our Snapshot_Status
func (snapshot *Snapshot_Status) AssignPropertiesToSnapshotStatus(destination *v20200930s.Snapshot_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(snapshot.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(snapshot.Conditions)

	// CreationData
	if snapshot.CreationData != nil {
		var creationDatum v20200930s.CreationData_Status
		err := snapshot.CreationData.AssignPropertiesToCreationDataStatus(&creationDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreationDataStatus() to populate field CreationData")
		}
		destination.CreationData = &creationDatum
	} else {
		destination.CreationData = nil
	}

	// DiskAccessId
	destination.DiskAccessId = genruntime.ClonePointerToString(snapshot.DiskAccessId)

	// DiskSizeBytes
	destination.DiskSizeBytes = genruntime.ClonePointerToInt(snapshot.DiskSizeBytes)

	// DiskSizeGB
	destination.DiskSizeGB = genruntime.ClonePointerToInt(snapshot.DiskSizeGB)

	// DiskState
	destination.DiskState = genruntime.ClonePointerToString(snapshot.DiskState)

	// Encryption
	if snapshot.Encryption != nil {
		var encryption v20200930s.Encryption_Status
		err := snapshot.Encryption.AssignPropertiesToEncryptionStatus(&encryption)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToEncryptionStatus() to populate field Encryption")
		}
		destination.Encryption = &encryption
	} else {
		destination.Encryption = nil
	}

	// EncryptionSettingsCollection
	if snapshot.EncryptionSettingsCollection != nil {
		var encryptionSettingsCollection v20200930s.EncryptionSettingsCollection_Status
		err := snapshot.EncryptionSettingsCollection.AssignPropertiesToEncryptionSettingsCollectionStatus(&encryptionSettingsCollection)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToEncryptionSettingsCollectionStatus() to populate field EncryptionSettingsCollection")
		}
		destination.EncryptionSettingsCollection = &encryptionSettingsCollection
	} else {
		destination.EncryptionSettingsCollection = nil
	}

	// ExtendedLocation
	if snapshot.ExtendedLocation != nil {
		var extendedLocationStatusStash alpha20201201s.ExtendedLocation_Status
		err := snapshot.ExtendedLocation.AssignPropertiesToExtendedLocationStatus(&extendedLocationStatusStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToExtendedLocationStatus() to populate field ExtendedLocation_StatusStash from ExtendedLocation")
		}
		var extendedLocationStatusStashLocal alpha20210701s.ExtendedLocation_Status
		err = extendedLocationStatusStash.AssignPropertiesToExtendedLocationStatus(&extendedLocationStatusStashLocal)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToExtendedLocationStatus() to populate field ExtendedLocation_StatusStash")
		}
		var extendedLocation v20200930s.ExtendedLocation_Status
		err = extendedLocationStatusStashLocal.AssignPropertiesToExtendedLocationStatus(&extendedLocation)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToExtendedLocationStatus() to populate field ExtendedLocation from ExtendedLocation_StatusStash")
		}
		destination.ExtendedLocation = &extendedLocation
	} else {
		destination.ExtendedLocation = nil
	}

	// HyperVGeneration
	destination.HyperVGeneration = genruntime.ClonePointerToString(snapshot.HyperVGeneration)

	// Id
	destination.Id = genruntime.ClonePointerToString(snapshot.Id)

	// Incremental
	if snapshot.Incremental != nil {
		incremental := *snapshot.Incremental
		destination.Incremental = &incremental
	} else {
		destination.Incremental = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(snapshot.Location)

	// ManagedBy
	destination.ManagedBy = genruntime.ClonePointerToString(snapshot.ManagedBy)

	// Name
	destination.Name = genruntime.ClonePointerToString(snapshot.Name)

	// NetworkAccessPolicy
	destination.NetworkAccessPolicy = genruntime.ClonePointerToString(snapshot.NetworkAccessPolicy)

	// OsType
	destination.OsType = genruntime.ClonePointerToString(snapshot.OsType)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(snapshot.ProvisioningState)

	// PurchasePlan
	if snapshot.PurchasePlan != nil {
		var purchasePlan v20200930s.PurchasePlan_Status
		err := snapshot.PurchasePlan.AssignPropertiesToPurchasePlanStatus(&purchasePlan)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToPurchasePlanStatus() to populate field PurchasePlan")
		}
		destination.PurchasePlan = &purchasePlan
	} else {
		destination.PurchasePlan = nil
	}

	// Sku
	if snapshot.Sku != nil {
		var sku v20200930s.SnapshotSku_Status
		err := snapshot.Sku.AssignPropertiesToSnapshotSkuStatus(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSnapshotSkuStatus() to populate field Sku")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(snapshot.Tags)

	// TimeCreated
	destination.TimeCreated = genruntime.ClonePointerToString(snapshot.TimeCreated)

	// Type
	destination.Type = genruntime.ClonePointerToString(snapshot.Type)

	// UniqueId
	destination.UniqueId = genruntime.ClonePointerToString(snapshot.UniqueId)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20200930.Snapshots_Spec
type Snapshots_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                    string                        `json:"azureName,omitempty"`
	CreationData                 *CreationData                 `json:"creationData,omitempty"`
	DiskAccessReference          *genruntime.ResourceReference `armReference:"DiskAccessId" json:"diskAccessReference,omitempty"`
	DiskSizeGB                   *int                          `json:"diskSizeGB,omitempty"`
	DiskState                    *string                       `json:"diskState,omitempty"`
	Encryption                   *Encryption                   `json:"encryption,omitempty"`
	EncryptionSettingsCollection *EncryptionSettingsCollection `json:"encryptionSettingsCollection,omitempty"`
	ExtendedLocation             *ExtendedLocation             `json:"extendedLocation,omitempty"`
	HyperVGeneration             *string                       `json:"hyperVGeneration,omitempty"`
	Incremental                  *bool                         `json:"incremental,omitempty"`
	Location                     *string                       `json:"location,omitempty"`
	NetworkAccessPolicy          *string                       `json:"networkAccessPolicy,omitempty"`
	OriginalVersion              string                        `json:"originalVersion,omitempty"`
	OsType                       *string                       `json:"osType,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner        *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag  genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PurchasePlan *PurchasePlan                      `json:"purchasePlan,omitempty"`
	Sku          *SnapshotSku                       `json:"sku,omitempty"`
	Tags         map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Snapshots_Spec{}

// ConvertSpecFrom populates our Snapshots_Spec from the provided source
func (snapshots *Snapshots_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20200930s.Snapshots_Spec)
	if ok {
		// Populate our instance from source
		return snapshots.AssignPropertiesFromSnapshotsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20200930s.Snapshots_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = snapshots.AssignPropertiesFromSnapshotsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Snapshots_Spec
func (snapshots *Snapshots_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20200930s.Snapshots_Spec)
	if ok {
		// Populate destination from our instance
		return snapshots.AssignPropertiesToSnapshotsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20200930s.Snapshots_Spec{}
	err := snapshots.AssignPropertiesToSnapshotsSpec(dst)
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

// AssignPropertiesFromSnapshotsSpec populates our Snapshots_Spec from the provided source Snapshots_Spec
func (snapshots *Snapshots_Spec) AssignPropertiesFromSnapshotsSpec(source *v20200930s.Snapshots_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	snapshots.AzureName = source.AzureName

	// CreationData
	if source.CreationData != nil {
		var creationDatum CreationData
		err := creationDatum.AssignPropertiesFromCreationData(source.CreationData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreationData() to populate field CreationData")
		}
		snapshots.CreationData = &creationDatum
	} else {
		snapshots.CreationData = nil
	}

	// DiskAccessReference
	if source.DiskAccessReference != nil {
		diskAccessReference := source.DiskAccessReference.Copy()
		snapshots.DiskAccessReference = &diskAccessReference
	} else {
		snapshots.DiskAccessReference = nil
	}

	// DiskSizeGB
	snapshots.DiskSizeGB = genruntime.ClonePointerToInt(source.DiskSizeGB)

	// DiskState
	snapshots.DiskState = genruntime.ClonePointerToString(source.DiskState)

	// Encryption
	if source.Encryption != nil {
		var encryption Encryption
		err := encryption.AssignPropertiesFromEncryption(source.Encryption)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromEncryption() to populate field Encryption")
		}
		snapshots.Encryption = &encryption
	} else {
		snapshots.Encryption = nil
	}

	// EncryptionSettingsCollection
	if source.EncryptionSettingsCollection != nil {
		var encryptionSettingsCollection EncryptionSettingsCollection
		err := encryptionSettingsCollection.AssignPropertiesFromEncryptionSettingsCollection(source.EncryptionSettingsCollection)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromEncryptionSettingsCollection() to populate field EncryptionSettingsCollection")
		}
		snapshots.EncryptionSettingsCollection = &encryptionSettingsCollection
	} else {
		snapshots.EncryptionSettingsCollection = nil
	}

	// ExtendedLocation
	if source.ExtendedLocation != nil {
		var extendedLocationStash alpha20210701s.ExtendedLocation
		err := extendedLocationStash.AssignPropertiesFromExtendedLocation(source.ExtendedLocation)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromExtendedLocation() to populate field ExtendedLocationStash from ExtendedLocation")
		}
		var extendedLocationStashLocal alpha20201201s.ExtendedLocation
		err = extendedLocationStashLocal.AssignPropertiesFromExtendedLocation(&extendedLocationStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromExtendedLocation() to populate field ExtendedLocationStash")
		}
		var extendedLocation ExtendedLocation
		err = extendedLocation.AssignPropertiesFromExtendedLocation(&extendedLocationStashLocal)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromExtendedLocation() to populate field ExtendedLocation from ExtendedLocationStash")
		}
		snapshots.ExtendedLocation = &extendedLocation
	} else {
		snapshots.ExtendedLocation = nil
	}

	// HyperVGeneration
	snapshots.HyperVGeneration = genruntime.ClonePointerToString(source.HyperVGeneration)

	// Incremental
	if source.Incremental != nil {
		incremental := *source.Incremental
		snapshots.Incremental = &incremental
	} else {
		snapshots.Incremental = nil
	}

	// Location
	snapshots.Location = genruntime.ClonePointerToString(source.Location)

	// NetworkAccessPolicy
	snapshots.NetworkAccessPolicy = genruntime.ClonePointerToString(source.NetworkAccessPolicy)

	// OriginalVersion
	snapshots.OriginalVersion = source.OriginalVersion

	// OsType
	snapshots.OsType = genruntime.ClonePointerToString(source.OsType)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		snapshots.Owner = &owner
	} else {
		snapshots.Owner = nil
	}

	// PurchasePlan
	if source.PurchasePlan != nil {
		var purchasePlan PurchasePlan
		err := purchasePlan.AssignPropertiesFromPurchasePlan(source.PurchasePlan)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromPurchasePlan() to populate field PurchasePlan")
		}
		snapshots.PurchasePlan = &purchasePlan
	} else {
		snapshots.PurchasePlan = nil
	}

	// Sku
	if source.Sku != nil {
		var sku SnapshotSku
		err := sku.AssignPropertiesFromSnapshotSku(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSnapshotSku() to populate field Sku")
		}
		snapshots.Sku = &sku
	} else {
		snapshots.Sku = nil
	}

	// Tags
	snapshots.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		snapshots.PropertyBag = propertyBag
	} else {
		snapshots.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSnapshotsSpec populates the provided destination Snapshots_Spec from our Snapshots_Spec
func (snapshots *Snapshots_Spec) AssignPropertiesToSnapshotsSpec(destination *v20200930s.Snapshots_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(snapshots.PropertyBag)

	// AzureName
	destination.AzureName = snapshots.AzureName

	// CreationData
	if snapshots.CreationData != nil {
		var creationDatum v20200930s.CreationData
		err := snapshots.CreationData.AssignPropertiesToCreationData(&creationDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreationData() to populate field CreationData")
		}
		destination.CreationData = &creationDatum
	} else {
		destination.CreationData = nil
	}

	// DiskAccessReference
	if snapshots.DiskAccessReference != nil {
		diskAccessReference := snapshots.DiskAccessReference.Copy()
		destination.DiskAccessReference = &diskAccessReference
	} else {
		destination.DiskAccessReference = nil
	}

	// DiskSizeGB
	destination.DiskSizeGB = genruntime.ClonePointerToInt(snapshots.DiskSizeGB)

	// DiskState
	destination.DiskState = genruntime.ClonePointerToString(snapshots.DiskState)

	// Encryption
	if snapshots.Encryption != nil {
		var encryption v20200930s.Encryption
		err := snapshots.Encryption.AssignPropertiesToEncryption(&encryption)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToEncryption() to populate field Encryption")
		}
		destination.Encryption = &encryption
	} else {
		destination.Encryption = nil
	}

	// EncryptionSettingsCollection
	if snapshots.EncryptionSettingsCollection != nil {
		var encryptionSettingsCollection v20200930s.EncryptionSettingsCollection
		err := snapshots.EncryptionSettingsCollection.AssignPropertiesToEncryptionSettingsCollection(&encryptionSettingsCollection)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToEncryptionSettingsCollection() to populate field EncryptionSettingsCollection")
		}
		destination.EncryptionSettingsCollection = &encryptionSettingsCollection
	} else {
		destination.EncryptionSettingsCollection = nil
	}

	// ExtendedLocation
	if snapshots.ExtendedLocation != nil {
		var extendedLocationStash alpha20201201s.ExtendedLocation
		err := snapshots.ExtendedLocation.AssignPropertiesToExtendedLocation(&extendedLocationStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToExtendedLocation() to populate field ExtendedLocationStash from ExtendedLocation")
		}
		var extendedLocationStashLocal alpha20210701s.ExtendedLocation
		err = extendedLocationStash.AssignPropertiesToExtendedLocation(&extendedLocationStashLocal)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToExtendedLocation() to populate field ExtendedLocationStash")
		}
		var extendedLocation v20200930s.ExtendedLocation
		err = extendedLocationStashLocal.AssignPropertiesToExtendedLocation(&extendedLocation)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToExtendedLocation() to populate field ExtendedLocation from ExtendedLocationStash")
		}
		destination.ExtendedLocation = &extendedLocation
	} else {
		destination.ExtendedLocation = nil
	}

	// HyperVGeneration
	destination.HyperVGeneration = genruntime.ClonePointerToString(snapshots.HyperVGeneration)

	// Incremental
	if snapshots.Incremental != nil {
		incremental := *snapshots.Incremental
		destination.Incremental = &incremental
	} else {
		destination.Incremental = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(snapshots.Location)

	// NetworkAccessPolicy
	destination.NetworkAccessPolicy = genruntime.ClonePointerToString(snapshots.NetworkAccessPolicy)

	// OriginalVersion
	destination.OriginalVersion = snapshots.OriginalVersion

	// OsType
	destination.OsType = genruntime.ClonePointerToString(snapshots.OsType)

	// Owner
	if snapshots.Owner != nil {
		owner := snapshots.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// PurchasePlan
	if snapshots.PurchasePlan != nil {
		var purchasePlan v20200930s.PurchasePlan
		err := snapshots.PurchasePlan.AssignPropertiesToPurchasePlan(&purchasePlan)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToPurchasePlan() to populate field PurchasePlan")
		}
		destination.PurchasePlan = &purchasePlan
	} else {
		destination.PurchasePlan = nil
	}

	// Sku
	if snapshots.Sku != nil {
		var sku v20200930s.SnapshotSku
		err := snapshots.Sku.AssignPropertiesToSnapshotSku(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSnapshotSku() to populate field Sku")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(snapshots.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20200930.SnapshotSku
// Deprecated version of SnapshotSku. Use v1beta20200930.SnapshotSku instead
type SnapshotSku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromSnapshotSku populates our SnapshotSku from the provided source SnapshotSku
func (snapshotSku *SnapshotSku) AssignPropertiesFromSnapshotSku(source *v20200930s.SnapshotSku) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Name
	snapshotSku.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		snapshotSku.PropertyBag = propertyBag
	} else {
		snapshotSku.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSnapshotSku populates the provided destination SnapshotSku from our SnapshotSku
func (snapshotSku *SnapshotSku) AssignPropertiesToSnapshotSku(destination *v20200930s.SnapshotSku) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(snapshotSku.PropertyBag)

	// Name
	destination.Name = genruntime.ClonePointerToString(snapshotSku.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20200930.SnapshotSku_Status
// Deprecated version of SnapshotSku_Status. Use v1beta20200930.SnapshotSku_Status instead
type SnapshotSku_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// AssignPropertiesFromSnapshotSkuStatus populates our SnapshotSku_Status from the provided source SnapshotSku_Status
func (snapshotSku *SnapshotSku_Status) AssignPropertiesFromSnapshotSkuStatus(source *v20200930s.SnapshotSku_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Name
	snapshotSku.Name = genruntime.ClonePointerToString(source.Name)

	// Tier
	snapshotSku.Tier = genruntime.ClonePointerToString(source.Tier)

	// Update the property bag
	if len(propertyBag) > 0 {
		snapshotSku.PropertyBag = propertyBag
	} else {
		snapshotSku.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSnapshotSkuStatus populates the provided destination SnapshotSku_Status from our SnapshotSku_Status
func (snapshotSku *SnapshotSku_Status) AssignPropertiesToSnapshotSkuStatus(destination *v20200930s.SnapshotSku_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(snapshotSku.PropertyBag)

	// Name
	destination.Name = genruntime.ClonePointerToString(snapshotSku.Name)

	// Tier
	destination.Tier = genruntime.ClonePointerToString(snapshotSku.Tier)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
