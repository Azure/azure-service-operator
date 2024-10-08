// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual
// machine. If SourceImage is provided, the destination virtual hard drive must not exist.
type Image_STATUS struct {
	// ExtendedLocation: The extended location of the Image.
	ExtendedLocation *ExtendedLocation_STATUS `json:"extendedLocation,omitempty"`

	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Properties: Describes the properties of an Image.
	Properties *ImageProperties_STATUS `json:"properties,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

// The complex type of the extended location.
type ExtendedLocation_STATUS struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType_STATUS `json:"type,omitempty"`
}

// Describes the properties of an Image.
type ImageProperties_STATUS struct {
	// HyperVGeneration: Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version
	// 2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource
	// like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed
	// resource.
	HyperVGeneration *HyperVGenerationType_STATUS `json:"hyperVGeneration,omitempty"`

	// ProvisioningState: The provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// SourceVirtualMachine: The source virtual machine from which Image is created.
	SourceVirtualMachine *SubResource_STATUS `json:"sourceVirtualMachine,omitempty"`

	// StorageProfile: Specifies the storage settings for the virtual machine disks.
	StorageProfile *ImageStorageProfile_STATUS `json:"storageProfile,omitempty"`
}

// The type of extendedLocation.
type ExtendedLocationType_STATUS string

const ExtendedLocationType_STATUS_EdgeZone = ExtendedLocationType_STATUS("EdgeZone")

// Mapping from string to ExtendedLocationType_STATUS
var extendedLocationType_STATUS_Values = map[string]ExtendedLocationType_STATUS{
	"edgezone": ExtendedLocationType_STATUS_EdgeZone,
}

// Specifies the HyperVGeneration Type
type HyperVGenerationType_STATUS string

const (
	HyperVGenerationType_STATUS_V1 = HyperVGenerationType_STATUS("V1")
	HyperVGenerationType_STATUS_V2 = HyperVGenerationType_STATUS("V2")
)

// Mapping from string to HyperVGenerationType_STATUS
var hyperVGenerationType_STATUS_Values = map[string]HyperVGenerationType_STATUS{
	"v1": HyperVGenerationType_STATUS_V1,
	"v2": HyperVGenerationType_STATUS_V2,
}

// Describes a storage profile.
type ImageStorageProfile_STATUS struct {
	// DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
	// For more information about disks, see [About disks and VHDs for Azure virtual
	// machines](https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview).
	DataDisks []ImageDataDisk_STATUS `json:"dataDisks,omitempty"`

	// OsDisk: Specifies information about the operating system disk used by the virtual machine.
	// For more information about disks, see [About disks and VHDs for Azure virtual
	// machines](https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview).
	OsDisk *ImageOSDisk_STATUS `json:"osDisk,omitempty"`

	// ZoneResilient: Specifies whether an image is zone resilient or not. Default is false. Zone resilient images can be
	// created only in regions that provide Zone Redundant Storage (ZRS).
	ZoneResilient *bool `json:"zoneResilient,omitempty"`
}

type SubResource_STATUS struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`
}

// Describes a data disk.
type ImageDataDisk_STATUS struct {
	// BlobUri: The Virtual Hard Disk.
	BlobUri *string `json:"blobUri,omitempty"`

	// Caching: Specifies the caching requirements.
	// Possible values are:
	// None
	// ReadOnly
	// ReadWrite
	// Default: None for Standard storage. ReadOnly for Premium storage
	Caching *ImageDataDisk_Caching_STATUS `json:"caching,omitempty"`

	// DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed image disk.
	DiskEncryptionSet *SubResource_STATUS `json:"diskEncryptionSet,omitempty"`

	// DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
	// disk in a virtual machine image.
	// This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
	// therefore must be unique for each data disk attached to a VM.
	Lun *int `json:"lun,omitempty"`

	// ManagedDisk: The managedDisk.
	ManagedDisk *SubResource_STATUS `json:"managedDisk,omitempty"`

	// Snapshot: The snapshot.
	Snapshot *SubResource_STATUS `json:"snapshot,omitempty"`

	// StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
	// data disks, it cannot be used with OS Disk.
	StorageAccountType *StorageAccountType_STATUS `json:"storageAccountType,omitempty"`
}

// Describes an Operating System disk.
type ImageOSDisk_STATUS struct {
	// BlobUri: The Virtual Hard Disk.
	BlobUri *string `json:"blobUri,omitempty"`

	// Caching: Specifies the caching requirements.
	// Possible values are:
	// None
	// ReadOnly
	// ReadWrite
	// Default: None for Standard storage. ReadOnly for Premium storage
	Caching *ImageOSDisk_Caching_STATUS `json:"caching,omitempty"`

	// DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed image disk.
	DiskEncryptionSet *SubResource_STATUS `json:"diskEncryptionSet,omitempty"`

	// DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
	// disk in a virtual machine image.
	// This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// ManagedDisk: The managedDisk.
	ManagedDisk *SubResource_STATUS `json:"managedDisk,omitempty"`

	// OsState: The OS State.
	OsState *ImageOSDisk_OsState_STATUS `json:"osState,omitempty"`

	// OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from a
	// custom image.
	// Possible values are:
	// Windows
	// Linux
	OsType *ImageOSDisk_OsType_STATUS `json:"osType,omitempty"`

	// Snapshot: The snapshot.
	Snapshot *SubResource_STATUS `json:"snapshot,omitempty"`

	// StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
	// data disks, it cannot be used with OS Disk.
	StorageAccountType *StorageAccountType_STATUS `json:"storageAccountType,omitempty"`
}

type ImageDataDisk_Caching_STATUS string

const (
	ImageDataDisk_Caching_STATUS_None      = ImageDataDisk_Caching_STATUS("None")
	ImageDataDisk_Caching_STATUS_ReadOnly  = ImageDataDisk_Caching_STATUS("ReadOnly")
	ImageDataDisk_Caching_STATUS_ReadWrite = ImageDataDisk_Caching_STATUS("ReadWrite")
)

// Mapping from string to ImageDataDisk_Caching_STATUS
var imageDataDisk_Caching_STATUS_Values = map[string]ImageDataDisk_Caching_STATUS{
	"none":      ImageDataDisk_Caching_STATUS_None,
	"readonly":  ImageDataDisk_Caching_STATUS_ReadOnly,
	"readwrite": ImageDataDisk_Caching_STATUS_ReadWrite,
}

type ImageOSDisk_Caching_STATUS string

const (
	ImageOSDisk_Caching_STATUS_None      = ImageOSDisk_Caching_STATUS("None")
	ImageOSDisk_Caching_STATUS_ReadOnly  = ImageOSDisk_Caching_STATUS("ReadOnly")
	ImageOSDisk_Caching_STATUS_ReadWrite = ImageOSDisk_Caching_STATUS("ReadWrite")
)

// Mapping from string to ImageOSDisk_Caching_STATUS
var imageOSDisk_Caching_STATUS_Values = map[string]ImageOSDisk_Caching_STATUS{
	"none":      ImageOSDisk_Caching_STATUS_None,
	"readonly":  ImageOSDisk_Caching_STATUS_ReadOnly,
	"readwrite": ImageOSDisk_Caching_STATUS_ReadWrite,
}

type ImageOSDisk_OsState_STATUS string

const (
	ImageOSDisk_OsState_STATUS_Generalized = ImageOSDisk_OsState_STATUS("Generalized")
	ImageOSDisk_OsState_STATUS_Specialized = ImageOSDisk_OsState_STATUS("Specialized")
)

// Mapping from string to ImageOSDisk_OsState_STATUS
var imageOSDisk_OsState_STATUS_Values = map[string]ImageOSDisk_OsState_STATUS{
	"generalized": ImageOSDisk_OsState_STATUS_Generalized,
	"specialized": ImageOSDisk_OsState_STATUS_Specialized,
}

type ImageOSDisk_OsType_STATUS string

const (
	ImageOSDisk_OsType_STATUS_Linux   = ImageOSDisk_OsType_STATUS("Linux")
	ImageOSDisk_OsType_STATUS_Windows = ImageOSDisk_OsType_STATUS("Windows")
)

// Mapping from string to ImageOSDisk_OsType_STATUS
var imageOSDisk_OsType_STATUS_Values = map[string]ImageOSDisk_OsType_STATUS{
	"linux":   ImageOSDisk_OsType_STATUS_Linux,
	"windows": ImageOSDisk_OsType_STATUS_Windows,
}

// Specifies the storage account type for the managed disk. Managed OS disk storage account type can only be set when you
// create the scale set. NOTE: UltraSSD_LRS can only be used with data disks. It cannot be used with OS Disk. Standard_LRS
// uses Standard HDD. StandardSSD_LRS uses Standard SSD. Premium_LRS uses Premium SSD. UltraSSD_LRS uses Ultra disk.
// Premium_ZRS uses Premium SSD zone redundant storage. StandardSSD_ZRS uses Standard SSD zone redundant storage. For more
// information regarding disks supported for Windows Virtual Machines, refer to
// https://docs.microsoft.com/azure/virtual-machines/windows/disks-types and, for Linux Virtual Machines, refer to
// https://docs.microsoft.com/azure/virtual-machines/linux/disks-types
type StorageAccountType_STATUS string

const (
	StorageAccountType_STATUS_Premium_LRS     = StorageAccountType_STATUS("Premium_LRS")
	StorageAccountType_STATUS_Premium_ZRS     = StorageAccountType_STATUS("Premium_ZRS")
	StorageAccountType_STATUS_StandardSSD_LRS = StorageAccountType_STATUS("StandardSSD_LRS")
	StorageAccountType_STATUS_StandardSSD_ZRS = StorageAccountType_STATUS("StandardSSD_ZRS")
	StorageAccountType_STATUS_Standard_LRS    = StorageAccountType_STATUS("Standard_LRS")
	StorageAccountType_STATUS_UltraSSD_LRS    = StorageAccountType_STATUS("UltraSSD_LRS")
)

// Mapping from string to StorageAccountType_STATUS
var storageAccountType_STATUS_Values = map[string]StorageAccountType_STATUS{
	"premium_lrs":     StorageAccountType_STATUS_Premium_LRS,
	"premium_zrs":     StorageAccountType_STATUS_Premium_ZRS,
	"standardssd_lrs": StorageAccountType_STATUS_StandardSSD_LRS,
	"standardssd_zrs": StorageAccountType_STATUS_StandardSSD_ZRS,
	"standard_lrs":    StorageAccountType_STATUS_Standard_LRS,
	"ultrassd_lrs":    StorageAccountType_STATUS_UltraSSD_LRS,
}
