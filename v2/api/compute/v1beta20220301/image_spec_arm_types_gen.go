// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Image_Spec_ARM struct {
	// ExtendedLocation: The complex type of the extended location.
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: The name of the image.
	Name string `json:"name,omitempty"`

	// Properties: Describes the properties of an Image.
	Properties *ImageProperties_ARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Image_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-03-01"
func (image Image_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (image *Image_Spec_ARM) GetName() string {
	return image.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/images"
func (image *Image_Spec_ARM) GetType() string {
	return "Microsoft.Compute/images"
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Compute.json#/definitions/ExtendedLocation
type ExtendedLocation_ARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocation_Type `json:"type,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Compute.json#/definitions/ImageProperties
type ImageProperties_ARM struct {
	// HyperVGeneration: Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version
	// 2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource
	// like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed
	// resource.
	HyperVGeneration     *ImageProperties_HyperVGeneration `json:"hyperVGeneration,omitempty"`
	SourceVirtualMachine *SubResource_ARM                  `json:"sourceVirtualMachine,omitempty"`

	// StorageProfile: Describes a storage profile.
	StorageProfile *ImageStorageProfile_ARM `json:"storageProfile,omitempty"`
}

// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocation_Type string

const ExtendedLocation_Type_EdgeZone = ExtendedLocation_Type("EdgeZone")

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Compute.json#/definitions/ImageStorageProfile
type ImageStorageProfile_ARM struct {
	// DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
	// For more information about disks, see [About disks and VHDs for Azure virtual
	// machines](https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview).
	DataDisks []ImageDataDisk_ARM `json:"dataDisks,omitempty"`

	// OsDisk: Describes an Operating System disk.
	OsDisk *ImageOSDisk_ARM `json:"osDisk,omitempty"`

	// ZoneResilient: Specifies whether an image is zone resilient or not. Default is false. Zone resilient images can be
	// created only in regions that provide Zone Redundant Storage (ZRS).
	ZoneResilient *bool `json:"zoneResilient,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Compute.json#/definitions/SubResource
type SubResource_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Compute.json#/definitions/ImageDataDisk
type ImageDataDisk_ARM struct {
	// BlobUri: The Virtual Hard Disk.
	BlobUri *string `json:"blobUri,omitempty"`

	// Caching: Specifies the caching requirements.
	// Possible values are:
	// None
	// ReadOnly
	// ReadWrite
	// Default: None for Standard storage. ReadOnly for Premium storage.
	Caching *ImageDataDisk_Caching `json:"caching,omitempty"`

	// DiskEncryptionSet: Describes the parameter of customer managed disk encryption set resource id that can be specified for
	// disk.
	// NOTE: The disk encryption set resource id can only be specified for managed disk. Please refer
	// https://aka.ms/mdssewithcmkoverview for more details.
	DiskEncryptionSet *DiskEncryptionSetParameters_ARM `json:"diskEncryptionSet,omitempty"`

	// DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
	// disk in a virtual machine image.
	// This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
	// therefore must be unique for each data disk attached to a VM.
	Lun         *int             `json:"lun,omitempty"`
	ManagedDisk *SubResource_ARM `json:"managedDisk,omitempty"`
	Snapshot    *SubResource_ARM `json:"snapshot,omitempty"`

	// StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
	// data disks, it cannot be used with OS Disk.
	StorageAccountType *ImageDataDisk_StorageAccountType `json:"storageAccountType,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Compute.json#/definitions/ImageOSDisk
type ImageOSDisk_ARM struct {
	// BlobUri: The Virtual Hard Disk.
	BlobUri *string `json:"blobUri,omitempty"`

	// Caching: Specifies the caching requirements.
	// Possible values are:
	// None
	// ReadOnly
	// ReadWrite
	// Default: None for Standard storage. ReadOnly for Premium storage.
	Caching *ImageOSDisk_Caching `json:"caching,omitempty"`

	// DiskEncryptionSet: Describes the parameter of customer managed disk encryption set resource id that can be specified for
	// disk.
	// NOTE: The disk encryption set resource id can only be specified for managed disk. Please refer
	// https://aka.ms/mdssewithcmkoverview for more details.
	DiskEncryptionSet *DiskEncryptionSetParameters_ARM `json:"diskEncryptionSet,omitempty"`

	// DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
	// disk in a virtual machine image.
	// This value cannot be larger than 1023 GB
	DiskSizeGB  *int             `json:"diskSizeGB,omitempty"`
	ManagedDisk *SubResource_ARM `json:"managedDisk,omitempty"`

	// OsState: The OS State. For managed images, use Generalized.
	OsState *ImageOSDisk_OsState `json:"osState,omitempty"`

	// OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from a
	// custom image.
	// Possible values are:
	// Windows
	// Linux.
	OsType   *ImageOSDisk_OsType `json:"osType,omitempty"`
	Snapshot *SubResource_ARM    `json:"snapshot,omitempty"`

	// StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
	// data disks, it cannot be used with OS Disk.
	StorageAccountType *ImageOSDisk_StorageAccountType `json:"storageAccountType,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Compute.json#/definitions/DiskEncryptionSetParameters
type DiskEncryptionSetParameters_ARM struct {
	Id *string `json:"id,omitempty"`
}
