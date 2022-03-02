// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Snapshots_SpecARM struct {
	//ExtendedLocation: The complex type of the extended location.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported
	//characters for the name are a-z, A-Z, 0-9 and _. The max name length is 80 characters.
	Name string `json:"name,omitempty"`

	//Properties: Snapshot resource properties.
	Properties *SnapshotPropertiesARM `json:"properties,omitempty"`

	//Sku: The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for
	//incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
	Sku *SnapshotSkuARM `json:"sku,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Snapshots_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-09-30"
func (snapshots Snapshots_SpecARM) GetAPIVersion() string {
	return "2020-09-30"
}

// GetName returns the Name of the resource
func (snapshots Snapshots_SpecARM) GetName() string {
	return snapshots.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/snapshots"
func (snapshots Snapshots_SpecARM) GetType() string {
	return "Microsoft.Compute/snapshots"
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SnapshotProperties
type SnapshotPropertiesARM struct {
	//CreationData: Data used when creating a disk.
	CreationData *CreationDataARM `json:"creationData,omitempty"`
	DiskAccessId *string          `json:"diskAccessId,omitempty"`

	//DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
	//create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
	//allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	//DiskState: The state of the snapshot.
	DiskState *SnapshotPropertiesDiskState `json:"diskState,omitempty"`

	//Encryption: Encryption at rest settings for disk or snapshot
	Encryption *EncryptionARM `json:"encryption,omitempty"`

	//EncryptionSettingsCollection: Encryption settings for disk or snapshot
	EncryptionSettingsCollection *EncryptionSettingsCollectionARM `json:"encryptionSettingsCollection,omitempty"`

	//HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *SnapshotPropertiesHyperVGeneration `json:"hyperVGeneration,omitempty"`

	//Incremental: Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full
	//snapshots and can be diffed.
	Incremental         *bool                                  `json:"incremental,omitempty"`
	NetworkAccessPolicy *SnapshotPropertiesNetworkAccessPolicy `json:"networkAccessPolicy,omitempty"`

	//OsType: The Operating System type.
	OsType *SnapshotPropertiesOsType `json:"osType,omitempty"`

	//PurchasePlan: Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.
	PurchasePlan *PurchasePlanARM `json:"purchasePlan,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SnapshotSku
type SnapshotSkuARM struct {
	//Name: The sku name.
	Name *SnapshotSkuName `json:"name,omitempty"`
}

// +kubebuilder:validation:Enum={"Premium_LRS","Standard_LRS","Standard_ZRS"}
type SnapshotSkuName string

const (
	SnapshotSkuNamePremiumLRS  = SnapshotSkuName("Premium_LRS")
	SnapshotSkuNameStandardLRS = SnapshotSkuName("Standard_LRS")
	SnapshotSkuNameStandardZRS = SnapshotSkuName("Standard_ZRS")
)
