// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Disk_Spec. Use v1beta20200930.Disk_Spec instead
type Disk_Spec_ARM struct {
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`
	Location         *string               `json:"location,omitempty"`
	Name             string                `json:"name,omitempty"`
	Properties       *DiskProperties_ARM   `json:"properties,omitempty"`
	Sku              *DiskSku_ARM          `json:"sku,omitempty"`
	Tags             map[string]string     `json:"tags,omitempty"`
	Zones            []string              `json:"zones,omitempty"`
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

// Deprecated version of DiskProperties. Use v1beta20200930.DiskProperties instead
type DiskProperties_ARM struct {
	BurstingEnabled              *bool                             `json:"burstingEnabled,omitempty"`
	CreationData                 *CreationData_ARM                 `json:"creationData,omitempty"`
	DiskAccessId                 *string                           `json:"diskAccessId,omitempty"`
	DiskIOPSReadOnly             *int                              `json:"diskIOPSReadOnly,omitempty"`
	DiskIOPSReadWrite            *int                              `json:"diskIOPSReadWrite,omitempty"`
	DiskMBpsReadOnly             *int                              `json:"diskMBpsReadOnly,omitempty"`
	DiskMBpsReadWrite            *int                              `json:"diskMBpsReadWrite,omitempty"`
	DiskSizeGB                   *int                              `json:"diskSizeGB,omitempty"`
	Encryption                   *Encryption_ARM                   `json:"encryption,omitempty"`
	EncryptionSettingsCollection *EncryptionSettingsCollection_ARM `json:"encryptionSettingsCollection,omitempty"`
	HyperVGeneration             *DiskProperties_HyperVGeneration  `json:"hyperVGeneration,omitempty"`
	MaxShares                    *int                              `json:"maxShares,omitempty"`
	NetworkAccessPolicy          *NetworkAccessPolicy              `json:"networkAccessPolicy,omitempty"`
	OsType                       *DiskProperties_OsType            `json:"osType,omitempty"`
	PurchasePlan                 *PurchasePlan_ARM                 `json:"purchasePlan,omitempty"`
	Tier                         *string                           `json:"tier,omitempty"`
}

// Deprecated version of DiskSku. Use v1beta20200930.DiskSku instead
type DiskSku_ARM struct {
	Name *DiskSku_Name `json:"name,omitempty"`
}

// Deprecated version of ExtendedLocation. Use v1beta20200930.ExtendedLocation instead
type ExtendedLocation_ARM struct {
	Name *string               `json:"name,omitempty"`
	Type *ExtendedLocationType `json:"type,omitempty"`
}

// Deprecated version of CreationData. Use v1beta20200930.CreationData instead
type CreationData_ARM struct {
	CreateOption          *CreationData_CreateOption `json:"createOption,omitempty"`
	GalleryImageReference *ImageDiskReference_ARM    `json:"galleryImageReference,omitempty"`
	ImageReference        *ImageDiskReference_ARM    `json:"imageReference,omitempty"`
	LogicalSectorSize     *int                       `json:"logicalSectorSize,omitempty"`
	SourceResourceId      *string                    `json:"sourceResourceId,omitempty"`
	SourceUri             *string                    `json:"sourceUri,omitempty"`
	StorageAccountId      *string                    `json:"storageAccountId,omitempty"`
	UploadSizeBytes       *int                       `json:"uploadSizeBytes,omitempty"`
}

// Deprecated version of DiskSku_Name. Use v1beta20200930.DiskSku_Name instead
// +kubebuilder:validation:Enum={"Premium_LRS","StandardSSD_LRS","Standard_LRS","UltraSSD_LRS"}
type DiskSku_Name string

const (
	DiskSku_Name_Premium_LRS     = DiskSku_Name("Premium_LRS")
	DiskSku_Name_StandardSSD_LRS = DiskSku_Name("StandardSSD_LRS")
	DiskSku_Name_Standard_LRS    = DiskSku_Name("Standard_LRS")
	DiskSku_Name_UltraSSD_LRS    = DiskSku_Name("UltraSSD_LRS")
)

// Deprecated version of Encryption. Use v1beta20200930.Encryption instead
type Encryption_ARM struct {
	DiskEncryptionSetId *string         `json:"diskEncryptionSetId,omitempty"`
	Type                *EncryptionType `json:"type,omitempty"`
}

// Deprecated version of EncryptionSettingsCollection. Use v1beta20200930.EncryptionSettingsCollection instead
type EncryptionSettingsCollection_ARM struct {
	Enabled                   *bool                           `json:"enabled,omitempty"`
	EncryptionSettings        []EncryptionSettingsElement_ARM `json:"encryptionSettings,omitempty"`
	EncryptionSettingsVersion *string                         `json:"encryptionSettingsVersion,omitempty"`
}

// Deprecated version of ExtendedLocationType. Use v1beta20200930.ExtendedLocationType instead
// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationType_EdgeZone = ExtendedLocationType("EdgeZone")

// Deprecated version of PurchasePlan. Use v1beta20200930.PurchasePlan instead
type PurchasePlan_ARM struct {
	Name          *string `json:"name,omitempty"`
	Product       *string `json:"product,omitempty"`
	PromotionCode *string `json:"promotionCode,omitempty"`
	Publisher     *string `json:"publisher,omitempty"`
}

// Deprecated version of EncryptionSettingsElement. Use v1beta20200930.EncryptionSettingsElement instead
type EncryptionSettingsElement_ARM struct {
	DiskEncryptionKey *KeyVaultAndSecretReference_ARM `json:"diskEncryptionKey,omitempty"`
	KeyEncryptionKey  *KeyVaultAndKeyReference_ARM    `json:"keyEncryptionKey,omitempty"`
}

// Deprecated version of ImageDiskReference. Use v1beta20200930.ImageDiskReference instead
type ImageDiskReference_ARM struct {
	Id  *string `json:"id,omitempty"`
	Lun *int    `json:"lun,omitempty"`
}

// Deprecated version of KeyVaultAndKeyReference. Use v1beta20200930.KeyVaultAndKeyReference instead
type KeyVaultAndKeyReference_ARM struct {
	KeyUrl      *string          `json:"keyUrl,omitempty"`
	SourceVault *SourceVault_ARM `json:"sourceVault,omitempty"`
}

// Deprecated version of KeyVaultAndSecretReference. Use v1beta20200930.KeyVaultAndSecretReference instead
type KeyVaultAndSecretReference_ARM struct {
	SecretUrl   *string          `json:"secretUrl,omitempty"`
	SourceVault *SourceVault_ARM `json:"sourceVault,omitempty"`
}

// Deprecated version of SourceVault. Use v1beta20200930.SourceVault instead
type SourceVault_ARM struct {
	Id *string `json:"id,omitempty"`
}
