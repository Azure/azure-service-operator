// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200930

// Disk resource.
type Disk_STATUS_ARM struct {
	// ExtendedLocation: The extended location where the disk will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`

	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// ManagedBy: A relative URI containing the ID of the VM that has the disk attached.
	ManagedBy *string `json:"managedBy,omitempty"`

	// ManagedByExtended: List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be
	// set to a value greater than one for disks to allow attaching them to multiple VMs.
	ManagedByExtended []string `json:"managedByExtended,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Properties: Disk resource properties.
	Properties *DiskProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
	Sku *DiskSku_STATUS_ARM `json:"sku,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`

	// Zones: The Logical zone list for Disk.
	Zones []string `json:"zones,omitempty"`
}

// Disk resource properties.
type DiskProperties_STATUS_ARM struct {
	// BurstingEnabled: Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is
	// disabled by default. Does not apply to Ultra disks.
	BurstingEnabled *bool `json:"burstingEnabled,omitempty"`

	// CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData *CreationData_STATUS_ARM `json:"creationData,omitempty"`

	// DiskAccessId: ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `json:"diskAccessId,omitempty"`

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

	// DiskSizeBytes: The size of the disk in bytes. This field is read only.
	DiskSizeBytes *int `json:"diskSizeBytes,omitempty"`

	// DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
	// create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
	// allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// DiskState: The state of the disk.
	DiskState *DiskState_STATUS `json:"diskState,omitempty"`

	// Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption_STATUS_ARM `json:"encryption,omitempty"`

	// EncryptionSettingsCollection: Encryption settings collection used for Azure Disk Encryption, can contain multiple
	// encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection_STATUS_ARM `json:"encryptionSettingsCollection,omitempty"`

	// HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *DiskProperties_HyperVGeneration_STATUS `json:"hyperVGeneration,omitempty"`

	// MaxShares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a
	// disk that can be mounted on multiple VMs at the same time.
	MaxShares *int `json:"maxShares,omitempty"`

	// NetworkAccessPolicy: Policy for accessing the disk via network.
	NetworkAccessPolicy *NetworkAccessPolicy_STATUS `json:"networkAccessPolicy,omitempty"`

	// OsType: The Operating System type.
	OsType *DiskProperties_OsType_STATUS `json:"osType,omitempty"`

	// ProvisioningState: The disk provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// PurchasePlan: Purchase plan information for the the image from which the OS disk was created. E.g. - {name:
	// 2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}
	PurchasePlan *PurchasePlan_STATUS_ARM `json:"purchasePlan,omitempty"`

	// ShareInfo: Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than
	// one for disks to allow attaching them to multiple VMs.
	ShareInfo []ShareInfoElement_STATUS_ARM `json:"shareInfo,omitempty"`

	// Tier: Performance tier of the disk (e.g, P4, S10) as described here:
	// https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
	Tier *string `json:"tier,omitempty"`

	// TimeCreated: The time when the disk was created.
	TimeCreated *string `json:"timeCreated,omitempty"`

	// UniqueId: Unique Guid identifying the resource.
	UniqueId *string `json:"uniqueId,omitempty"`
}

// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
type DiskSku_STATUS_ARM struct {
	// Name: The sku name.
	Name *DiskSku_Name_STATUS `json:"name,omitempty"`

	// Tier: The sku tier.
	Tier *string `json:"tier,omitempty"`
}

// The complex type of the extended location.
type ExtendedLocation_STATUS_ARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType_STATUS `json:"type,omitempty"`
}

// Data used when creating a disk.
type CreationData_STATUS_ARM struct {
	// CreateOption: This enumerates the possible sources of a disk's creation.
	CreateOption *CreationData_CreateOption_STATUS `json:"createOption,omitempty"`

	// GalleryImageReference: Required if creating from a Gallery Image. The id of the ImageDiskReference will be the ARM id of
	// the shared galley image version from which to create a disk.
	GalleryImageReference *ImageDiskReference_STATUS_ARM `json:"galleryImageReference,omitempty"`

	// ImageReference: Disk source information.
	ImageReference *ImageDiskReference_STATUS_ARM `json:"imageReference,omitempty"`

	// LogicalSectorSize: Logical sector size in bytes for Ultra disks. Supported values are 512 ad 4096. 4096 is the default.
	LogicalSectorSize *int `json:"logicalSectorSize,omitempty"`

	// SourceResourceId: If createOption is Copy, this is the ARM id of the source snapshot or disk.
	SourceResourceId *string `json:"sourceResourceId,omitempty"`

	// SourceUniqueId: If this field is set, this is the unique id identifying the source of this resource.
	SourceUniqueId *string `json:"sourceUniqueId,omitempty"`

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

type DiskSku_Name_STATUS string

const (
	DiskSku_Name_STATUS_Premium_LRS     = DiskSku_Name_STATUS("Premium_LRS")
	DiskSku_Name_STATUS_StandardSSD_LRS = DiskSku_Name_STATUS("StandardSSD_LRS")
	DiskSku_Name_STATUS_Standard_LRS    = DiskSku_Name_STATUS("Standard_LRS")
	DiskSku_Name_STATUS_UltraSSD_LRS    = DiskSku_Name_STATUS("UltraSSD_LRS")
)

// Encryption at rest settings for disk or snapshot
type Encryption_STATUS_ARM struct {
	// DiskEncryptionSetId: ResourceId of the disk encryption set to use for enabling encryption at rest.
	DiskEncryptionSetId *string `json:"diskEncryptionSetId,omitempty"`

	// Type: The type of key used to encrypt the data of the disk.
	Type *EncryptionType_STATUS `json:"type,omitempty"`
}

// Encryption settings for disk or snapshot
type EncryptionSettingsCollection_STATUS_ARM struct {
	// Enabled: Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set
	// this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is
	// null in the request object, the existing settings remain unchanged.
	Enabled *bool `json:"enabled,omitempty"`

	// EncryptionSettings: A collection of encryption settings, one for each disk volume.
	EncryptionSettings []EncryptionSettingsElement_STATUS_ARM `json:"encryptionSettings,omitempty"`

	// EncryptionSettingsVersion: Describes what type of encryption is used for the disks. Once this field is set, it cannot be
	// overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.
	EncryptionSettingsVersion *string `json:"encryptionSettingsVersion,omitempty"`
}

// The type of extendedLocation.
type ExtendedLocationType_STATUS string

const ExtendedLocationType_STATUS_EdgeZone = ExtendedLocationType_STATUS("EdgeZone")

// Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.
type PurchasePlan_STATUS_ARM struct {
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

type ShareInfoElement_STATUS_ARM struct {
	// VmUri: A relative URI containing the ID of the VM that has the disk attached.
	VmUri *string `json:"vmUri,omitempty"`
}

// Encryption settings for one disk volume.
type EncryptionSettingsElement_STATUS_ARM struct {
	// DiskEncryptionKey: Key Vault Secret Url and vault id of the disk encryption key
	DiskEncryptionKey *KeyVaultAndSecretReference_STATUS_ARM `json:"diskEncryptionKey,omitempty"`

	// KeyEncryptionKey: Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when
	// provided is used to unwrap the disk encryption key.
	KeyEncryptionKey *KeyVaultAndKeyReference_STATUS_ARM `json:"keyEncryptionKey,omitempty"`
}

// The source image used for creating the disk.
type ImageDiskReference_STATUS_ARM struct {
	// Id: A relative uri containing either a Platform Image Repository or user image reference.
	Id *string `json:"id,omitempty"`

	// Lun: If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the
	// image to use. For OS disks, this field is null.
	Lun *int `json:"lun,omitempty"`
}

// Key Vault Key Url and vault id of KeK, KeK is optional and when provided is used to unwrap the encryptionKey
type KeyVaultAndKeyReference_STATUS_ARM struct {
	// KeyUrl: Url pointing to a key or secret in KeyVault
	KeyUrl *string `json:"keyUrl,omitempty"`

	// SourceVault: Resource id of the KeyVault containing the key or secret
	SourceVault *SourceVault_STATUS_ARM `json:"sourceVault,omitempty"`
}

// Key Vault Secret Url and vault id of the encryption key
type KeyVaultAndSecretReference_STATUS_ARM struct {
	// SecretUrl: Url pointing to a key or secret in KeyVault
	SecretUrl *string `json:"secretUrl,omitempty"`

	// SourceVault: Resource id of the KeyVault containing the key or secret
	SourceVault *SourceVault_STATUS_ARM `json:"sourceVault,omitempty"`
}

// The vault id is an Azure Resource Manager Resource id in the form
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}
type SourceVault_STATUS_ARM struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`
}
