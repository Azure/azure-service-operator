// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930storage

import (
	"github.com/Azure/azure-service-operator/hack/generated/apis/microsoft.compute/v1alpha1api20201201storage"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20200930.Disk
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/resourceDefinitions/disks
type Disk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Disks_Spec  `json:"spec,omitempty"`
	Status            Disk_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Disk{}

// GetConditions returns the conditions of the resource
func (disk *Disk) GetConditions() conditions.Conditions {
	return disk.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (disk *Disk) SetConditions(conditions conditions.Conditions) {
	disk.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Disk{}

// AzureName returns the Azure name of the resource
func (disk *Disk) AzureName() string {
	return disk.Spec.AzureName
}

// GetSpec returns the specification of this resource
func (disk *Disk) GetSpec() genruntime.ConvertibleSpec {
	return &disk.Spec
}

// GetStatus returns the status of this resource
func (disk *Disk) GetStatus() genruntime.ConvertibleStatus {
	return &disk.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/disks"
func (disk *Disk) GetType() string {
	return "Microsoft.Compute/disks"
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (disk *Disk) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(disk.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: disk.Namespace, Name: disk.Spec.Owner.Name}
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (disk *Disk) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: disk.Spec.OriginalVersion,
		Kind:    "Disk",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20200930.Disk
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/resourceDefinitions/disks
type DiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Disk `json:"items"`
}

//Storage version of v1alpha1api20200930.Disk_Status
//Generated from:
type Disk_Status struct {
	BurstingEnabled              *bool                                `json:"burstingEnabled,omitempty"`
	Conditions                   []conditions.Condition               `json:"conditions,omitempty"`
	CreationData                 *CreationData_Status                 `json:"creationData,omitempty"`
	DiskAccessId                 *string                              `json:"diskAccessId,omitempty"`
	DiskIOPSReadOnly             *int                                 `json:"diskIOPSReadOnly,omitempty"`
	DiskIOPSReadWrite            *int                                 `json:"diskIOPSReadWrite,omitempty"`
	DiskMBpsReadOnly             *int                                 `json:"diskMBpsReadOnly,omitempty"`
	DiskMBpsReadWrite            *int                                 `json:"diskMBpsReadWrite,omitempty"`
	DiskSizeBytes                *int                                 `json:"diskSizeBytes,omitempty"`
	DiskSizeGB                   *int                                 `json:"diskSizeGB,omitempty"`
	DiskState                    *string                              `json:"diskState,omitempty"`
	Encryption                   *Encryption_Status                   `json:"encryption,omitempty"`
	EncryptionSettingsCollection *EncryptionSettingsCollection_Status `json:"encryptionSettingsCollection,omitempty"`
	ExtendedLocation             *ExtendedLocation_Status             `json:"extendedLocation,omitempty"`
	HyperVGeneration             *string                              `json:"hyperVGeneration,omitempty"`
	Id                           *string                              `json:"id,omitempty"`
	Location                     *string                              `json:"location,omitempty"`
	ManagedBy                    *string                              `json:"managedBy,omitempty"`
	ManagedByExtended            []string                             `json:"managedByExtended,omitempty"`
	MaxShares                    *int                                 `json:"maxShares,omitempty"`
	Name                         *string                              `json:"name,omitempty"`
	NetworkAccessPolicy          *string                              `json:"networkAccessPolicy,omitempty"`
	OsType                       *string                              `json:"osType,omitempty"`
	PropertyBag                  genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                              `json:"provisioningState,omitempty"`
	PurchasePlan                 *PurchasePlan_Status                 `json:"purchasePlan,omitempty"`
	ShareInfo                    []ShareInfoElement_Status            `json:"shareInfo,omitempty"`
	Sku                          *DiskSku_Status                      `json:"sku,omitempty"`
	Tags                         map[string]string                    `json:"tags,omitempty"`
	Tier                         *string                              `json:"tier,omitempty"`
	TimeCreated                  *string                              `json:"timeCreated,omitempty"`
	Type                         *string                              `json:"type,omitempty"`
	UniqueId                     *string                              `json:"uniqueId,omitempty"`
	Zones                        []string                             `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Disk_Status{}

// ConvertStatusFrom populates our Disk_Status from the provided source
func (diskStatus *Disk_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == diskStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(diskStatus)
}

// ConvertStatusTo populates the provided destination from our Disk_Status
func (diskStatus *Disk_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == diskStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(diskStatus)
}

//Storage version of v1alpha1api20200930.Disks_Spec
type Disks_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string        `json:"azureName"`
	BurstingEnabled *bool         `json:"burstingEnabled,omitempty"`
	CreationData    *CreationData `json:"creationData,omitempty"`

	//DiskAccessReference: ARM id of the DiskAccess resource for using private
	//endpoints on disks.
	DiskAccessReference          *genruntime.ResourceReference `armReference:"DiskAccessId" json:"diskAccessReference,omitempty"`
	DiskIOPSReadOnly             *int                          `json:"diskIOPSReadOnly,omitempty"`
	DiskIOPSReadWrite            *int                          `json:"diskIOPSReadWrite,omitempty"`
	DiskMBpsReadOnly             *int                          `json:"diskMBpsReadOnly,omitempty"`
	DiskMBpsReadWrite            *int                          `json:"diskMBpsReadWrite,omitempty"`
	DiskSizeGB                   *int                          `json:"diskSizeGB,omitempty"`
	Encryption                   *Encryption                   `json:"encryption,omitempty"`
	EncryptionSettingsCollection *EncryptionSettingsCollection `json:"encryptionSettingsCollection,omitempty"`
	ExtendedLocation             *ExtendedLocation             `json:"extendedLocation,omitempty"`
	HyperVGeneration             *string                       `json:"hyperVGeneration,omitempty"`
	Location                     *string                       `json:"location,omitempty"`
	MaxShares                    *int                          `json:"maxShares,omitempty"`
	NetworkAccessPolicy          *string                       `json:"networkAccessPolicy,omitempty"`
	OriginalVersion              string                        `json:"originalVersion"`
	OsType                       *string                       `json:"osType,omitempty"`

	// +kubebuilder:validation:Required
	Owner        genruntime.KnownResourceReference `group:"microsoft.resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag  genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	PurchasePlan *PurchasePlan                     `json:"purchasePlan,omitempty"`
	Sku          *DiskSku                          `json:"sku,omitempty"`
	Tags         map[string]string                 `json:"tags,omitempty"`
	Tier         *string                           `json:"tier,omitempty"`
	Zones        []string                          `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Disks_Spec{}

// ConvertSpecFrom populates our Disks_Spec from the provided source
func (disksSpec *Disks_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == disksSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(disksSpec)
}

// ConvertSpecTo populates the provided destination from our Disks_Spec
func (disksSpec *Disks_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == disksSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(disksSpec)
}

//Storage version of v1alpha1api20200930.CreationData
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/CreationData
type CreationData struct {
	CreateOption          *string                `json:"createOption,omitempty"`
	GalleryImageReference *ImageDiskReference    `json:"galleryImageReference,omitempty"`
	ImageReference        *ImageDiskReference    `json:"imageReference,omitempty"`
	LogicalSectorSize     *int                   `json:"logicalSectorSize,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//SourceResourceReference: If createOption is Copy, this is the ARM id of the
	//source snapshot or disk.
	SourceResourceReference *genruntime.ResourceReference `armReference:"SourceResourceId" json:"sourceResourceReference,omitempty"`
	SourceUri               *string                       `json:"sourceUri,omitempty"`
	StorageAccountId        *string                       `json:"storageAccountId,omitempty"`
	UploadSizeBytes         *int                          `json:"uploadSizeBytes,omitempty"`
}

//Storage version of v1alpha1api20200930.CreationData_Status
//Generated from:
type CreationData_Status struct {
	CreateOption          *string                    `json:"createOption,omitempty"`
	GalleryImageReference *ImageDiskReference_Status `json:"galleryImageReference,omitempty"`
	ImageReference        *ImageDiskReference_Status `json:"imageReference,omitempty"`
	LogicalSectorSize     *int                       `json:"logicalSectorSize,omitempty"`
	PropertyBag           genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	SourceResourceId      *string                    `json:"sourceResourceId,omitempty"`
	SourceUniqueId        *string                    `json:"sourceUniqueId,omitempty"`
	SourceUri             *string                    `json:"sourceUri,omitempty"`
	StorageAccountId      *string                    `json:"storageAccountId,omitempty"`
	UploadSizeBytes       *int                       `json:"uploadSizeBytes,omitempty"`
}

//Storage version of v1alpha1api20200930.DiskSku
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/DiskSku
type DiskSku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200930.DiskSku_Status
//Generated from:
type DiskSku_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20200930.Encryption
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/Encryption
type Encryption struct {
	DiskEncryptionSetId *string                `json:"diskEncryptionSetId,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20200930.EncryptionSettingsCollection
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsCollection
type EncryptionSettingsCollection struct {
	Enabled                   *bool                       `json:"enabled,omitempty"`
	EncryptionSettings        []EncryptionSettingsElement `json:"encryptionSettings,omitempty"`
	EncryptionSettingsVersion *string                     `json:"encryptionSettingsVersion,omitempty"`
	PropertyBag               genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200930.EncryptionSettingsCollection_Status
//Generated from:
type EncryptionSettingsCollection_Status struct {
	Enabled                   *bool                              `json:"enabled,omitempty"`
	EncryptionSettings        []EncryptionSettingsElement_Status `json:"encryptionSettings,omitempty"`
	EncryptionSettingsVersion *string                            `json:"encryptionSettingsVersion,omitempty"`
	PropertyBag               genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200930.Encryption_Status
//Generated from:
type Encryption_Status struct {
	DiskEncryptionSetId *string                `json:"diskEncryptionSetId,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20200930.ExtendedLocation
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ExtendedLocation
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// AssignPropertiesFromExtendedLocation populates our ExtendedLocation from the provided source ExtendedLocation
func (extendedLocation *ExtendedLocation) AssignPropertiesFromExtendedLocation(source *v1alpha1api20201201storage.ExtendedLocation) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Name
	if source.Name != nil {
		name := *source.Name
		extendedLocation.Name = &name
	} else {
		extendedLocation.Name = nil
	}

	// Type
	if source.Type != nil {
		typeVar := *source.Type
		extendedLocation.Type = &typeVar
	} else {
		extendedLocation.Type = nil
	}

	// Update the property bag
	extendedLocation.PropertyBag = propertyBag

	// No error
	return nil
}

// AssignPropertiesToExtendedLocation populates the provided destination ExtendedLocation from our ExtendedLocation
func (extendedLocation *ExtendedLocation) AssignPropertiesToExtendedLocation(destination *v1alpha1api20201201storage.ExtendedLocation) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(extendedLocation.PropertyBag)

	// Name
	if extendedLocation.Name != nil {
		name := *extendedLocation.Name
		destination.Name = &name
	} else {
		destination.Name = nil
	}

	// Type
	if extendedLocation.Type != nil {
		typeVar := *extendedLocation.Type
		destination.Type = &typeVar
	} else {
		destination.Type = nil
	}

	// Update the property bag
	destination.PropertyBag = propertyBag

	// No error
	return nil
}

//Storage version of v1alpha1api20200930.ExtendedLocation_Status
//Generated from:
type ExtendedLocation_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// AssignPropertiesFromExtendedLocationStatus populates our ExtendedLocation_Status from the provided source ExtendedLocation_Status
func (extendedLocationStatus *ExtendedLocation_Status) AssignPropertiesFromExtendedLocationStatus(source *v1alpha1api20201201storage.ExtendedLocation_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Name
	if source.Name != nil {
		name := *source.Name
		extendedLocationStatus.Name = &name
	} else {
		extendedLocationStatus.Name = nil
	}

	// Type
	if source.Type != nil {
		typeVar := *source.Type
		extendedLocationStatus.Type = &typeVar
	} else {
		extendedLocationStatus.Type = nil
	}

	// Update the property bag
	extendedLocationStatus.PropertyBag = propertyBag

	// No error
	return nil
}

// AssignPropertiesToExtendedLocationStatus populates the provided destination ExtendedLocation_Status from our ExtendedLocation_Status
func (extendedLocationStatus *ExtendedLocation_Status) AssignPropertiesToExtendedLocationStatus(destination *v1alpha1api20201201storage.ExtendedLocation_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(extendedLocationStatus.PropertyBag)

	// Name
	if extendedLocationStatus.Name != nil {
		name := *extendedLocationStatus.Name
		destination.Name = &name
	} else {
		destination.Name = nil
	}

	// Type
	if extendedLocationStatus.Type != nil {
		typeVar := *extendedLocationStatus.Type
		destination.Type = &typeVar
	} else {
		destination.Type = nil
	}

	// Update the property bag
	destination.PropertyBag = propertyBag

	// No error
	return nil
}

//Storage version of v1alpha1api20200930.PurchasePlan
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/PurchasePlan
type PurchasePlan struct {
	Name          *string                `json:"name,omitempty"`
	Product       *string                `json:"product,omitempty"`
	PromotionCode *string                `json:"promotionCode,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Publisher     *string                `json:"publisher,omitempty"`
}

//Storage version of v1alpha1api20200930.PurchasePlan_Status
//Generated from:
type PurchasePlan_Status struct {
	Name          *string                `json:"name,omitempty"`
	Product       *string                `json:"product,omitempty"`
	PromotionCode *string                `json:"promotionCode,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Publisher     *string                `json:"publisher,omitempty"`
}

//Storage version of v1alpha1api20200930.ShareInfoElement_Status
//Generated from:
type ShareInfoElement_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VmUri       *string                `json:"vmUri,omitempty"`
}

//Storage version of v1alpha1api20200930.EncryptionSettingsElement
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsElement
type EncryptionSettingsElement struct {
	DiskEncryptionKey *KeyVaultAndSecretReference `json:"diskEncryptionKey,omitempty"`
	KeyEncryptionKey  *KeyVaultAndKeyReference    `json:"keyEncryptionKey,omitempty"`
	PropertyBag       genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200930.EncryptionSettingsElement_Status
//Generated from:
type EncryptionSettingsElement_Status struct {
	DiskEncryptionKey *KeyVaultAndSecretReference_Status `json:"diskEncryptionKey,omitempty"`
	KeyEncryptionKey  *KeyVaultAndKeyReference_Status    `json:"keyEncryptionKey,omitempty"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200930.ImageDiskReference
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ImageDiskReference
type ImageDiskReference struct {
	Lun         *int                   `json:"lun,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	//Reference: A relative uri containing either a Platform Image Repository or user
	//image reference.
	Reference genruntime.ResourceReference `armReference:"Id" json:"reference"`
}

//Storage version of v1alpha1api20200930.ImageDiskReference_Status
//Generated from:
type ImageDiskReference_Status struct {
	Id          *string                `json:"id,omitempty"`
	Lun         *int                   `json:"lun,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200930.KeyVaultAndKeyReference
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndKeyReference
type KeyVaultAndKeyReference struct {
	KeyUrl      *string                `json:"keyUrl,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceVault *SourceVault           `json:"sourceVault,omitempty"`
}

//Storage version of v1alpha1api20200930.KeyVaultAndKeyReference_Status
//Generated from:
type KeyVaultAndKeyReference_Status struct {
	KeyUrl      *string                `json:"keyUrl,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceVault *SourceVault_Status    `json:"sourceVault,omitempty"`
}

//Storage version of v1alpha1api20200930.KeyVaultAndSecretReference
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndSecretReference
type KeyVaultAndSecretReference struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SecretUrl   *string                `json:"secretUrl,omitempty"`
	SourceVault *SourceVault           `json:"sourceVault,omitempty"`
}

//Storage version of v1alpha1api20200930.KeyVaultAndSecretReference_Status
//Generated from:
type KeyVaultAndSecretReference_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SecretUrl   *string                `json:"secretUrl,omitempty"`
	SourceVault *SourceVault_Status    `json:"sourceVault,omitempty"`
}

//Storage version of v1alpha1api20200930.SourceVault
//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SourceVault
type SourceVault struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//Reference: Resource Id
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

//Storage version of v1alpha1api20200930.SourceVault_Status
//Generated from:
type SourceVault_Status struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Disk{}, &DiskList{})
}
