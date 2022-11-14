// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200930

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Disk_Spec_ARM struct {
	// ExtendedLocation: The extended location where the disk will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`

	// Location: Resource location
	Location   *string             `json:"location,omitempty"`
	Name       string              `json:"name,omitempty"`
	Properties *DiskProperties_ARM `json:"properties,omitempty"`
	Sku        *DiskSku_ARM        `json:"sku,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Zones: The Logical zone list for Disk.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Disk_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-09-30"
func (disk Disk_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (disk *Disk_Spec_ARM) GetName() string {
	return disk.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/disks"
func (disk *Disk_Spec_ARM) GetType() string {
	return "Microsoft.Compute/disks"
}

type DiskProperties_ARM struct {
	// BurstingEnabled: Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is
	// disabled by default. Does not apply to Ultra disks.
	BurstingEnabled *bool `json:"burstingEnabled,omitempty"`

	// CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData *CreationData_ARM `json:"creationData,omitempty"`
	DiskAccessId *string           `json:"diskAccessId,omitempty"`

	// DiskIOPSReadOnly: The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One
	// operation can transfer between 4k and 256k bytes.
	DiskIOPSReadOnly *int `json:"diskIOPSReadOnly,omitempty"`

	// DiskIOPSReadWrite: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can
	// transfer between 4k and 256k bytes.
	DiskIOPSReadWrite *int `json:"diskIOPSReadWrite,omitempty"`

	// DiskMBpsReadOnly: The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly.
	// MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadOnly *int `json:"diskMBpsReadOnly,omitempty"`

	// DiskMBpsReadWrite: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes
	// per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadWrite *int `json:"diskMBpsReadWrite,omitempty"`

	// DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
	// create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
	// allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption_ARM `json:"encryption,omitempty"`

	// EncryptionSettingsCollection: Encryption settings collection used for Azure Disk Encryption, can contain multiple
	// encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection_ARM `json:"encryptionSettingsCollection,omitempty"`

	// HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *DiskProperties_HyperVGeneration `json:"hyperVGeneration,omitempty"`

	// MaxShares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a
	// disk that can be mounted on multiple VMs at the same time.
	MaxShares           *int                 `json:"maxShares,omitempty"`
	NetworkAccessPolicy *NetworkAccessPolicy `json:"networkAccessPolicy,omitempty"`

	// OsType: The Operating System type.
	OsType *DiskProperties_OsType `json:"osType,omitempty"`

	// PurchasePlan: Purchase plan information for the the image from which the OS disk was created. E.g. - {name:
	// 2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}
	PurchasePlan *PurchasePlan_ARM `json:"purchasePlan,omitempty"`

	// Tier: Performance tier of the disk (e.g, P4, S10) as described here:
	// https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
	Tier *string `json:"tier,omitempty"`
}

type DiskSku_ARM struct {
	// Name: The sku name.
	Name *DiskSku_Name `json:"name,omitempty"`
}

type ExtendedLocation_ARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType `json:"type,omitempty"`
}

type CreationData_ARM struct {
	// CreateOption: This enumerates the possible sources of a disk's creation.
	CreateOption *CreationData_CreateOption `json:"createOption,omitempty"`

	// GalleryImageReference: Required if creating from a Gallery Image. The id of the ImageDiskReference will be the ARM id of
	// the shared galley image version from which to create a disk.
	GalleryImageReference *ImageDiskReference_ARM `json:"galleryImageReference,omitempty"`

	// ImageReference: Disk source information.
	ImageReference *ImageDiskReference_ARM `json:"imageReference,omitempty"`

	// LogicalSectorSize: Logical sector size in bytes for Ultra disks. Supported values are 512 ad 4096. 4096 is the default.
	LogicalSectorSize *int    `json:"logicalSectorSize,omitempty"`
	SourceResourceId  *string `json:"sourceResourceId,omitempty"`

	// SourceUri: If createOption is Import, this is the URI of a blob to be imported into a managed disk.
	SourceUri *string `json:"sourceUri,omitempty"`

	// StorageAccountId: Required if createOption is Import. The Azure Resource Manager identifier of the storage account
	// containing the blob to import as a disk.
	StorageAccountId *string `json:"storageAccountId,omitempty"`

	// UploadSizeBytes: If createOption is Upload, this is the size of the contents of the upload including the VHD footer.
	// This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512
	// bytes for the VHD footer).
	UploadSizeBytes *int `json:"uploadSizeBytes,omitempty"`
}

// +kubebuilder:validation:Enum={"Premium_LRS","StandardSSD_LRS","Standard_LRS","UltraSSD_LRS"}
type DiskSku_Name string

const (
	DiskSku_Name_Premium_LRS     = DiskSku_Name("Premium_LRS")
	DiskSku_Name_StandardSSD_LRS = DiskSku_Name("StandardSSD_LRS")
	DiskSku_Name_Standard_LRS    = DiskSku_Name("Standard_LRS")
	DiskSku_Name_UltraSSD_LRS    = DiskSku_Name("UltraSSD_LRS")
)

type Encryption_ARM struct {
	DiskEncryptionSetId *string         `json:"diskEncryptionSetId,omitempty"`
	Type                *EncryptionType `json:"type,omitempty"`
}

type EncryptionSettingsCollection_ARM struct {
	// Enabled: Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set
	// this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is
	// null in the request object, the existing settings remain unchanged.
	Enabled *bool `json:"enabled,omitempty"`

	// EncryptionSettings: A collection of encryption settings, one for each disk volume.
	EncryptionSettings []EncryptionSettingsElement_ARM `json:"encryptionSettings,omitempty"`

	// EncryptionSettingsVersion: Describes what type of encryption is used for the disks. Once this field is set, it cannot be
	// overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.
	EncryptionSettingsVersion *string `json:"encryptionSettingsVersion,omitempty"`
}

// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationType_EdgeZone = ExtendedLocationType("EdgeZone")

type PurchasePlan_ARM struct {
	// Name: The plan ID.
	Name *string `json:"name,omitempty"`

	// Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
	// imageReference element.
	Product *string `json:"product,omitempty"`

	// PromotionCode: The Offer Promotion Code.
	PromotionCode *string `json:"promotionCode,omitempty"`

	// Publisher: The publisher ID.
	Publisher *string `json:"publisher,omitempty"`
}

type EncryptionSettingsElement_ARM struct {
	// DiskEncryptionKey: Key Vault Secret Url and vault id of the disk encryption key
	DiskEncryptionKey *KeyVaultAndSecretReference_ARM `json:"diskEncryptionKey,omitempty"`

	// KeyEncryptionKey: Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when
	// provided is used to unwrap the disk encryption key.
	KeyEncryptionKey *KeyVaultAndKeyReference_ARM `json:"keyEncryptionKey,omitempty"`
}

type ImageDiskReference_ARM struct {
	Id *string `json:"id,omitempty"`

	// Lun: If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the
	// image to use. For OS disks, this field is null.
	Lun *int `json:"lun,omitempty"`
}

type KeyVaultAndKeyReference_ARM struct {
	// KeyUrl: Url pointing to a key or secret in KeyVault
	KeyUrl *string `json:"keyUrl,omitempty"`

	// SourceVault: Resource id of the KeyVault containing the key or secret
	SourceVault *SourceVault_ARM `json:"sourceVault,omitempty"`
}

type KeyVaultAndSecretReference_ARM struct {
	// SecretUrl: Url pointing to a key or secret in KeyVault
	SecretUrl *string `json:"secretUrl,omitempty"`

	// SourceVault: Resource id of the KeyVault containing the key or secret
	SourceVault *SourceVault_ARM `json:"sourceVault,omitempty"`
}

type SourceVault_ARM struct {
	Id *string `json:"id,omitempty"`
}
