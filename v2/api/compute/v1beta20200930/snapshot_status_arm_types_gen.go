// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200930

type Snapshot_STATUSARM struct {
	// ExtendedLocation: The extended location where the snapshot will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation_STATUSARM `json:"extendedLocation,omitempty"`

	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// ManagedBy: Unused. Always Null.
	ManagedBy *string `json:"managedBy,omitempty"`

	// Name: Resource name
	Name       *string                       `json:"name,omitempty"`
	Properties *SnapshotProperties_STATUSARM `json:"properties,omitempty"`
	Sku        *SnapshotSku_STATUSARM        `json:"sku,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

type SnapshotProperties_STATUSARM struct {
	// CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData *CreationData_STATUSARM `json:"creationData,omitempty"`

	// DiskAccessId: ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `json:"diskAccessId,omitempty"`

	// DiskSizeBytes: The size of the disk in bytes. This field is read only.
	DiskSizeBytes *int `json:"diskSizeBytes,omitempty"`

	// DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
	// create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
	// allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// DiskState: The state of the snapshot.
	DiskState *DiskState_STATUS `json:"diskState,omitempty"`

	// Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption_STATUSARM `json:"encryption,omitempty"`

	// EncryptionSettingsCollection: Encryption settings collection used be Azure Disk Encryption, can contain multiple
	// encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection_STATUSARM `json:"encryptionSettingsCollection,omitempty"`

	// HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *SnapshotProperties_HyperVGeneration_STATUS `json:"hyperVGeneration,omitempty"`

	// Incremental: Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full
	// snapshots and can be diffed.
	Incremental         *bool                       `json:"incremental,omitempty"`
	NetworkAccessPolicy *NetworkAccessPolicy_STATUS `json:"networkAccessPolicy,omitempty"`

	// OsType: The Operating System type.
	OsType *SnapshotProperties_OsType_STATUS `json:"osType,omitempty"`

	// ProvisioningState: The disk provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// PurchasePlan: Purchase plan information for the image from which the source disk for the snapshot was originally created.
	PurchasePlan *PurchasePlan_STATUSARM `json:"purchasePlan,omitempty"`

	// TimeCreated: The time when the snapshot was created.
	TimeCreated *string `json:"timeCreated,omitempty"`

	// UniqueId: Unique Guid identifying the resource.
	UniqueId *string `json:"uniqueId,omitempty"`
}

type SnapshotSku_STATUSARM struct {
	// Name: The sku name.
	Name *SnapshotSku_Name_STATUS `json:"name,omitempty"`

	// Tier: The sku tier.
	Tier *string `json:"tier,omitempty"`
}

type SnapshotSku_Name_STATUS string

const (
	SnapshotSku_Name_Premium_LRS_STATUS  = SnapshotSku_Name_STATUS("Premium_LRS")
	SnapshotSku_Name_Standard_LRS_STATUS = SnapshotSku_Name_STATUS("Standard_LRS")
	SnapshotSku_Name_Standard_ZRS_STATUS = SnapshotSku_Name_STATUS("Standard_ZRS")
)
