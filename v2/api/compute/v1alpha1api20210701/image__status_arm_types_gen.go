// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210701

type Image_StatusARM struct {
	//ExtendedLocation: The extended location of the Image.
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`

	//Id: Resource Id
	Id *string `json:"id,omitempty"`

	//Location: Resource location
	Location *string `json:"location,omitempty"`

	//Name: Resource name
	Name       *string                    `json:"name,omitempty"`
	Properties *ImageProperties_StatusARM `json:"properties,omitempty"`

	//Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type
	Type *string `json:"type,omitempty"`
}

type ExtendedLocation_StatusARM struct {
	//Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	//Type: The type of the extended location.
	Type *ExtendedLocationType_Status `json:"type,omitempty"`
}

type ImageProperties_StatusARM struct {
	//HyperVGeneration: Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version
	//2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource
	//like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed
	//resource.
	HyperVGeneration *HyperVGenerationType_Status `json:"hyperVGeneration,omitempty"`

	//ProvisioningState: The provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	//SourceVirtualMachine: The source virtual machine from which Image is created.
	SourceVirtualMachine *SubResource_StatusARM `json:"sourceVirtualMachine,omitempty"`

	//StorageProfile: Specifies the storage settings for the virtual machine disks.
	StorageProfile *ImageStorageProfile_StatusARM `json:"storageProfile,omitempty"`
}

type ExtendedLocationType_Status string

const ExtendedLocationType_StatusEdgeZone = ExtendedLocationType_Status("EdgeZone")

type ImageStorageProfile_StatusARM struct {
	//DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
	//For more information about disks, see [About disks and VHDs for Azure virtual
	//machines](https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview).
	DataDisks []ImageDataDisk_StatusARM `json:"dataDisks,omitempty"`

	//OsDisk: Specifies information about the operating system disk used by the virtual machine.
	//For more information about disks, see [About disks and VHDs for Azure virtual
	//machines](https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview).
	OsDisk *ImageOSDisk_StatusARM `json:"osDisk,omitempty"`

	//ZoneResilient: Specifies whether an image is zone resilient or not. Default is false. Zone resilient images can be
	//created only in regions that provide Zone Redundant Storage (ZRS).
	ZoneResilient *bool `json:"zoneResilient,omitempty"`
}

type SubResource_StatusARM struct {
	//Id: Resource Id
	Id *string `json:"id,omitempty"`
}

type ImageDataDisk_StatusARM struct {
	//BlobUri: The Virtual Hard Disk.
	BlobUri *string `json:"blobUri,omitempty"`

	//Caching: Specifies the caching requirements.
	//Possible values are:
	//None
	//ReadOnly
	//ReadWrite
	//Default: None for Standard storage. ReadOnly for Premium storage
	Caching *ImageDataDiskStatusCaching `json:"caching,omitempty"`

	//DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed image disk.
	DiskEncryptionSet *SubResource_StatusARM `json:"diskEncryptionSet,omitempty"`

	//DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
	//disk in a virtual machine image.
	//This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	//Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
	//therefore must be unique for each data disk attached to a VM.
	Lun *int `json:"lun,omitempty"`

	//ManagedDisk: The managedDisk.
	ManagedDisk *SubResource_StatusARM `json:"managedDisk,omitempty"`

	//Snapshot: The snapshot.
	Snapshot *SubResource_StatusARM `json:"snapshot,omitempty"`

	//StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
	//data disks, it cannot be used with OS Disk.
	StorageAccountType *StorageAccountType_Status `json:"storageAccountType,omitempty"`
}

type ImageOSDisk_StatusARM struct {
	//BlobUri: The Virtual Hard Disk.
	BlobUri *string `json:"blobUri,omitempty"`

	//Caching: Specifies the caching requirements.
	//Possible values are:
	//None
	//ReadOnly
	//ReadWrite
	//Default: None for Standard storage. ReadOnly for Premium storage
	Caching *ImageOSDiskStatusCaching `json:"caching,omitempty"`

	//DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed image disk.
	DiskEncryptionSet *SubResource_StatusARM `json:"diskEncryptionSet,omitempty"`

	//DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
	//disk in a virtual machine image.
	//This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	//ManagedDisk: The managedDisk.
	ManagedDisk *SubResource_StatusARM `json:"managedDisk,omitempty"`

	//OsState: The OS State.
	OsState *ImageOSDiskStatusOsState `json:"osState,omitempty"`

	//OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from a
	//custom image.
	//Possible values are:
	//Windows
	//Linux
	OsType *ImageOSDiskStatusOsType `json:"osType,omitempty"`

	//Snapshot: The snapshot.
	Snapshot *SubResource_StatusARM `json:"snapshot,omitempty"`

	//StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
	//data disks, it cannot be used with OS Disk.
	StorageAccountType *StorageAccountType_Status `json:"storageAccountType,omitempty"`
}
