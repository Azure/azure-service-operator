// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200930storage

import (
	v20201201s "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=compute.azure.com,resources=disks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=compute.azure.com,resources={disks/status,disks/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20200930.Disk
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/resourceDefinitions/disks
type Disk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Disks_Spec  `json:"spec,omitempty"`
	Status            Disk_STATUS `json:"status,omitempty"`
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

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-09-30"
func (disk Disk) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (disk *Disk) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
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

// NewEmptyStatus returns a new empty (blank) status
func (disk *Disk) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Disk_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (disk *Disk) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(disk.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  disk.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (disk *Disk) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Disk_STATUS); ok {
		disk.Status = *st
		return nil
	}

	// Convert status to required version
	var st Disk_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	disk.Status = st
	return nil
}

// Hub marks that this Disk is the hub type for conversion
func (disk *Disk) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (disk *Disk) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: disk.Spec.OriginalVersion,
		Kind:    "Disk",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20200930.Disk
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/resourceDefinitions/disks
type DiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Disk `json:"items"`
}

// Storage version of v1beta20200930.APIVersion
// +kubebuilder:validation:Enum={"2020-09-30"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-09-30")

// Storage version of v1beta20200930.Disk_STATUS
type Disk_STATUS struct {
	BurstingEnabled              *bool                                `json:"burstingEnabled,omitempty"`
	Conditions                   []conditions.Condition               `json:"conditions,omitempty"`
	CreationData                 *CreationData_STATUS                 `json:"creationData,omitempty"`
	DiskAccessId                 *string                              `json:"diskAccessId,omitempty"`
	DiskIOPSReadOnly             *int                                 `json:"diskIOPSReadOnly,omitempty"`
	DiskIOPSReadWrite            *int                                 `json:"diskIOPSReadWrite,omitempty"`
	DiskMBpsReadOnly             *int                                 `json:"diskMBpsReadOnly,omitempty"`
	DiskMBpsReadWrite            *int                                 `json:"diskMBpsReadWrite,omitempty"`
	DiskSizeBytes                *int                                 `json:"diskSizeBytes,omitempty"`
	DiskSizeGB                   *int                                 `json:"diskSizeGB,omitempty"`
	DiskState                    *string                              `json:"diskState,omitempty"`
	Encryption                   *Encryption_STATUS                   `json:"encryption,omitempty"`
	EncryptionSettingsCollection *EncryptionSettingsCollection_STATUS `json:"encryptionSettingsCollection,omitempty"`
	ExtendedLocation             *ExtendedLocation_STATUS             `json:"extendedLocation,omitempty"`
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
	PurchasePlan                 *PurchasePlan_STATUS                 `json:"purchasePlan,omitempty"`
	ShareInfo                    []ShareInfoElement_STATUS            `json:"shareInfo,omitempty"`
	Sku                          *DiskSku_STATUS                      `json:"sku,omitempty"`
	Tags                         map[string]string                    `json:"tags,omitempty"`
	Tier                         *string                              `json:"tier,omitempty"`
	TimeCreated                  *string                              `json:"timeCreated,omitempty"`
	Type                         *string                              `json:"type,omitempty"`
	UniqueId                     *string                              `json:"uniqueId,omitempty"`
	Zones                        []string                             `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Disk_STATUS{}

// ConvertStatusFrom populates our Disk_STATUS from the provided source
func (disk *Disk_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == disk {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(disk)
}

// ConvertStatusTo populates the provided destination from our Disk_STATUS
func (disk *Disk_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == disk {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(disk)
}

// Storage version of v1beta20200930.Disks_Spec
type Disks_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string        `json:"azureName,omitempty"`
	BurstingEnabled *bool         `json:"burstingEnabled,omitempty"`
	CreationData    *CreationData `json:"creationData,omitempty"`

	// DiskAccessReference: ARM id of the DiskAccess resource for using private endpoints on disks.
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
	OriginalVersion              string                        `json:"originalVersion,omitempty"`
	OsType                       *string                       `json:"osType,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner        *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag  genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PurchasePlan *PurchasePlan                      `json:"purchasePlan,omitempty"`
	Sku          *DiskSku                           `json:"sku,omitempty"`
	Tags         map[string]string                  `json:"tags,omitempty"`
	Tier         *string                            `json:"tier,omitempty"`
	Zones        []string                           `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Disks_Spec{}

// ConvertSpecFrom populates our Disks_Spec from the provided source
func (disks *Disks_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == disks {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(disks)
}

// ConvertSpecTo populates the provided destination from our Disks_Spec
func (disks *Disks_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == disks {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(disks)
}

// Storage version of v1beta20200930.CreationData
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/CreationData
type CreationData struct {
	CreateOption          *string                `json:"createOption,omitempty"`
	GalleryImageReference *ImageDiskReference    `json:"galleryImageReference,omitempty"`
	ImageReference        *ImageDiskReference    `json:"imageReference,omitempty"`
	LogicalSectorSize     *int                   `json:"logicalSectorSize,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// SourceResourceReference: If createOption is Copy, this is the ARM id of the source snapshot or disk.
	SourceResourceReference *genruntime.ResourceReference `armReference:"SourceResourceId" json:"sourceResourceReference,omitempty"`
	SourceUri               *string                       `json:"sourceUri,omitempty"`
	StorageAccountId        *string                       `json:"storageAccountId,omitempty"`
	UploadSizeBytes         *int                          `json:"uploadSizeBytes,omitempty"`
}

// Storage version of v1beta20200930.CreationData_STATUS
type CreationData_STATUS struct {
	CreateOption          *string                    `json:"createOption,omitempty"`
	GalleryImageReference *ImageDiskReference_STATUS `json:"galleryImageReference,omitempty"`
	ImageReference        *ImageDiskReference_STATUS `json:"imageReference,omitempty"`
	LogicalSectorSize     *int                       `json:"logicalSectorSize,omitempty"`
	PropertyBag           genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	SourceResourceId      *string                    `json:"sourceResourceId,omitempty"`
	SourceUniqueId        *string                    `json:"sourceUniqueId,omitempty"`
	SourceUri             *string                    `json:"sourceUri,omitempty"`
	StorageAccountId      *string                    `json:"storageAccountId,omitempty"`
	UploadSizeBytes       *int                       `json:"uploadSizeBytes,omitempty"`
}

// Storage version of v1beta20200930.DiskSku
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/DiskSku
type DiskSku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200930.DiskSku_STATUS
type DiskSku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20200930.Encryption
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/Encryption
type Encryption struct {
	// DiskEncryptionSetReference: ResourceId of the disk encryption set to use for enabling encryption at rest.
	DiskEncryptionSetReference *genruntime.ResourceReference `armReference:"DiskEncryptionSetId" json:"diskEncryptionSetReference,omitempty"`
	PropertyBag                genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                       *string                       `json:"type,omitempty"`
}

// Storage version of v1beta20200930.Encryption_STATUS
type Encryption_STATUS struct {
	DiskEncryptionSetId *string                `json:"diskEncryptionSetId,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                *string                `json:"type,omitempty"`
}

// Storage version of v1beta20200930.EncryptionSettingsCollection
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsCollection
type EncryptionSettingsCollection struct {
	Enabled                   *bool                       `json:"enabled,omitempty"`
	EncryptionSettings        []EncryptionSettingsElement `json:"encryptionSettings,omitempty"`
	EncryptionSettingsVersion *string                     `json:"encryptionSettingsVersion,omitempty"`
	PropertyBag               genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200930.EncryptionSettingsCollection_STATUS
type EncryptionSettingsCollection_STATUS struct {
	Enabled                   *bool                              `json:"enabled,omitempty"`
	EncryptionSettings        []EncryptionSettingsElement_STATUS `json:"encryptionSettings,omitempty"`
	EncryptionSettingsVersion *string                            `json:"encryptionSettingsVersion,omitempty"`
	PropertyBag               genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200930.ExtendedLocation
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ExtendedLocation
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// AssignPropertiesFromExtendedLocation populates our ExtendedLocation from the provided source ExtendedLocation
func (location *ExtendedLocation) AssignPropertiesFromExtendedLocation(source *v20201201s.ExtendedLocation) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Name
	location.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	location.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		location.PropertyBag = propertyBag
	} else {
		location.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToExtendedLocation populates the provided destination ExtendedLocation from our ExtendedLocation
func (location *ExtendedLocation) AssignPropertiesToExtendedLocation(destination *v20201201s.ExtendedLocation) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(location.PropertyBag)

	// Name
	destination.Name = genruntime.ClonePointerToString(location.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(location.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1beta20200930.ExtendedLocation_STATUS
type ExtendedLocation_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// AssignPropertiesFromExtendedLocationSTATUS populates our ExtendedLocation_STATUS from the provided source ExtendedLocation_STATUS
func (location *ExtendedLocation_STATUS) AssignPropertiesFromExtendedLocationSTATUS(source *v20201201s.ExtendedLocation_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Name
	location.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	location.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		location.PropertyBag = propertyBag
	} else {
		location.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToExtendedLocationSTATUS populates the provided destination ExtendedLocation_STATUS from our ExtendedLocation_STATUS
func (location *ExtendedLocation_STATUS) AssignPropertiesToExtendedLocationSTATUS(destination *v20201201s.ExtendedLocation_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(location.PropertyBag)

	// Name
	destination.Name = genruntime.ClonePointerToString(location.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(location.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1beta20200930.PurchasePlan
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/PurchasePlan
type PurchasePlan struct {
	Name          *string                `json:"name,omitempty"`
	Product       *string                `json:"product,omitempty"`
	PromotionCode *string                `json:"promotionCode,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Publisher     *string                `json:"publisher,omitempty"`
}

// Storage version of v1beta20200930.PurchasePlan_STATUS
type PurchasePlan_STATUS struct {
	Name          *string                `json:"name,omitempty"`
	Product       *string                `json:"product,omitempty"`
	PromotionCode *string                `json:"promotionCode,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Publisher     *string                `json:"publisher,omitempty"`
}

// Storage version of v1beta20200930.ShareInfoElement_STATUS
type ShareInfoElement_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VmUri       *string                `json:"vmUri,omitempty"`
}

// Storage version of v1beta20200930.EncryptionSettingsElement
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsElement
type EncryptionSettingsElement struct {
	DiskEncryptionKey *KeyVaultAndSecretReference `json:"diskEncryptionKey,omitempty"`
	KeyEncryptionKey  *KeyVaultAndKeyReference    `json:"keyEncryptionKey,omitempty"`
	PropertyBag       genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200930.EncryptionSettingsElement_STATUS
type EncryptionSettingsElement_STATUS struct {
	DiskEncryptionKey *KeyVaultAndSecretReference_STATUS `json:"diskEncryptionKey,omitempty"`
	KeyEncryptionKey  *KeyVaultAndKeyReference_STATUS    `json:"keyEncryptionKey,omitempty"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200930.ImageDiskReference
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ImageDiskReference
type ImageDiskReference struct {
	Lun         *int                   `json:"lun,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: A relative uri containing either a Platform Image Repository or user image reference.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20200930.ImageDiskReference_STATUS
type ImageDiskReference_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	Lun         *int                   `json:"lun,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200930.KeyVaultAndKeyReference
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndKeyReference
type KeyVaultAndKeyReference struct {
	KeyUrl      *string                `json:"keyUrl,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceVault *SourceVault           `json:"sourceVault,omitempty"`
}

// Storage version of v1beta20200930.KeyVaultAndKeyReference_STATUS
type KeyVaultAndKeyReference_STATUS struct {
	KeyUrl      *string                `json:"keyUrl,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceVault *SourceVault_STATUS    `json:"sourceVault,omitempty"`
}

// Storage version of v1beta20200930.KeyVaultAndSecretReference
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndSecretReference
type KeyVaultAndSecretReference struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SecretUrl   *string                `json:"secretUrl,omitempty"`
	SourceVault *SourceVault           `json:"sourceVault,omitempty"`
}

// Storage version of v1beta20200930.KeyVaultAndSecretReference_STATUS
type KeyVaultAndSecretReference_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SecretUrl   *string                `json:"secretUrl,omitempty"`
	SourceVault *SourceVault_STATUS    `json:"sourceVault,omitempty"`
}

// Storage version of v1beta20200930.SourceVault
// Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SourceVault
type SourceVault struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource Id
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20200930.SourceVault_STATUS
type SourceVault_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Disk{}, &DiskList{})
}
