//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20200930storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreationData) DeepCopyInto(out *CreationData) {
	*out = *in
	if in.CreateOption != nil {
		in, out := &in.CreateOption, &out.CreateOption
		*out = new(string)
		**out = **in
	}
	if in.GalleryImageReference != nil {
		in, out := &in.GalleryImageReference, &out.GalleryImageReference
		*out = new(ImageDiskReference)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageReference != nil {
		in, out := &in.ImageReference, &out.ImageReference
		*out = new(ImageDiskReference)
		(*in).DeepCopyInto(*out)
	}
	if in.LogicalSectorSize != nil {
		in, out := &in.LogicalSectorSize, &out.LogicalSectorSize
		*out = new(int)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourceResourceReference != nil {
		in, out := &in.SourceResourceReference, &out.SourceResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.SourceUri != nil {
		in, out := &in.SourceUri, &out.SourceUri
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountId != nil {
		in, out := &in.StorageAccountId, &out.StorageAccountId
		*out = new(string)
		**out = **in
	}
	if in.UploadSizeBytes != nil {
		in, out := &in.UploadSizeBytes, &out.UploadSizeBytes
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreationData.
func (in *CreationData) DeepCopy() *CreationData {
	if in == nil {
		return nil
	}
	out := new(CreationData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreationData_Status) DeepCopyInto(out *CreationData_Status) {
	*out = *in
	if in.CreateOption != nil {
		in, out := &in.CreateOption, &out.CreateOption
		*out = new(string)
		**out = **in
	}
	if in.GalleryImageReference != nil {
		in, out := &in.GalleryImageReference, &out.GalleryImageReference
		*out = new(ImageDiskReference_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageReference != nil {
		in, out := &in.ImageReference, &out.ImageReference
		*out = new(ImageDiskReference_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.LogicalSectorSize != nil {
		in, out := &in.LogicalSectorSize, &out.LogicalSectorSize
		*out = new(int)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourceResourceId != nil {
		in, out := &in.SourceResourceId, &out.SourceResourceId
		*out = new(string)
		**out = **in
	}
	if in.SourceUniqueId != nil {
		in, out := &in.SourceUniqueId, &out.SourceUniqueId
		*out = new(string)
		**out = **in
	}
	if in.SourceUri != nil {
		in, out := &in.SourceUri, &out.SourceUri
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountId != nil {
		in, out := &in.StorageAccountId, &out.StorageAccountId
		*out = new(string)
		**out = **in
	}
	if in.UploadSizeBytes != nil {
		in, out := &in.UploadSizeBytes, &out.UploadSizeBytes
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreationData_Status.
func (in *CreationData_Status) DeepCopy() *CreationData_Status {
	if in == nil {
		return nil
	}
	out := new(CreationData_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Disk) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskList) DeepCopyInto(out *DiskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Disk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskList.
func (in *DiskList) DeepCopy() *DiskList {
	if in == nil {
		return nil
	}
	out := new(DiskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSku) DeepCopyInto(out *DiskSku) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSku.
func (in *DiskSku) DeepCopy() *DiskSku {
	if in == nil {
		return nil
	}
	out := new(DiskSku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSku_Status) DeepCopyInto(out *DiskSku_Status) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSku_Status.
func (in *DiskSku_Status) DeepCopy() *DiskSku_Status {
	if in == nil {
		return nil
	}
	out := new(DiskSku_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk_Status) DeepCopyInto(out *Disk_Status) {
	*out = *in
	if in.BurstingEnabled != nil {
		in, out := &in.BurstingEnabled, &out.BurstingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreationData != nil {
		in, out := &in.CreationData, &out.CreationData
		*out = new(CreationData_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskAccessId != nil {
		in, out := &in.DiskAccessId, &out.DiskAccessId
		*out = new(string)
		**out = **in
	}
	if in.DiskIOPSReadOnly != nil {
		in, out := &in.DiskIOPSReadOnly, &out.DiskIOPSReadOnly
		*out = new(int)
		**out = **in
	}
	if in.DiskIOPSReadWrite != nil {
		in, out := &in.DiskIOPSReadWrite, &out.DiskIOPSReadWrite
		*out = new(int)
		**out = **in
	}
	if in.DiskMBpsReadOnly != nil {
		in, out := &in.DiskMBpsReadOnly, &out.DiskMBpsReadOnly
		*out = new(int)
		**out = **in
	}
	if in.DiskMBpsReadWrite != nil {
		in, out := &in.DiskMBpsReadWrite, &out.DiskMBpsReadWrite
		*out = new(int)
		**out = **in
	}
	if in.DiskSizeBytes != nil {
		in, out := &in.DiskSizeBytes, &out.DiskSizeBytes
		*out = new(int)
		**out = **in
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.DiskState != nil {
		in, out := &in.DiskState, &out.DiskState
		*out = new(string)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(Encryption_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.EncryptionSettingsCollection != nil {
		in, out := &in.EncryptionSettingsCollection, &out.EncryptionSettingsCollection
		*out = new(EncryptionSettingsCollection_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtendedLocation != nil {
		in, out := &in.ExtendedLocation, &out.ExtendedLocation
		*out = new(ExtendedLocation_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.HyperVGeneration != nil {
		in, out := &in.HyperVGeneration, &out.HyperVGeneration
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ManagedBy != nil {
		in, out := &in.ManagedBy, &out.ManagedBy
		*out = new(string)
		**out = **in
	}
	if in.ManagedByExtended != nil {
		in, out := &in.ManagedByExtended, &out.ManagedByExtended
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MaxShares != nil {
		in, out := &in.MaxShares, &out.MaxShares
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkAccessPolicy != nil {
		in, out := &in.NetworkAccessPolicy, &out.NetworkAccessPolicy
		*out = new(string)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.PurchasePlan != nil {
		in, out := &in.PurchasePlan, &out.PurchasePlan
		*out = new(PurchasePlan_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.ShareInfo != nil {
		in, out := &in.ShareInfo, &out.ShareInfo
		*out = make([]ShareInfoElement_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(DiskSku_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UniqueId != nil {
		in, out := &in.UniqueId, &out.UniqueId
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk_Status.
func (in *Disk_Status) DeepCopy() *Disk_Status {
	if in == nil {
		return nil
	}
	out := new(Disk_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disks_Spec) DeepCopyInto(out *Disks_Spec) {
	*out = *in
	if in.BurstingEnabled != nil {
		in, out := &in.BurstingEnabled, &out.BurstingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CreationData != nil {
		in, out := &in.CreationData, &out.CreationData
		*out = new(CreationData)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskAccessReference != nil {
		in, out := &in.DiskAccessReference, &out.DiskAccessReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.DiskIOPSReadOnly != nil {
		in, out := &in.DiskIOPSReadOnly, &out.DiskIOPSReadOnly
		*out = new(int)
		**out = **in
	}
	if in.DiskIOPSReadWrite != nil {
		in, out := &in.DiskIOPSReadWrite, &out.DiskIOPSReadWrite
		*out = new(int)
		**out = **in
	}
	if in.DiskMBpsReadOnly != nil {
		in, out := &in.DiskMBpsReadOnly, &out.DiskMBpsReadOnly
		*out = new(int)
		**out = **in
	}
	if in.DiskMBpsReadWrite != nil {
		in, out := &in.DiskMBpsReadWrite, &out.DiskMBpsReadWrite
		*out = new(int)
		**out = **in
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(Encryption)
		(*in).DeepCopyInto(*out)
	}
	if in.EncryptionSettingsCollection != nil {
		in, out := &in.EncryptionSettingsCollection, &out.EncryptionSettingsCollection
		*out = new(EncryptionSettingsCollection)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtendedLocation != nil {
		in, out := &in.ExtendedLocation, &out.ExtendedLocation
		*out = new(ExtendedLocation)
		(*in).DeepCopyInto(*out)
	}
	if in.HyperVGeneration != nil {
		in, out := &in.HyperVGeneration, &out.HyperVGeneration
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaxShares != nil {
		in, out := &in.MaxShares, &out.MaxShares
		*out = new(int)
		**out = **in
	}
	if in.NetworkAccessPolicy != nil {
		in, out := &in.NetworkAccessPolicy, &out.NetworkAccessPolicy
		*out = new(string)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(string)
		**out = **in
	}
	out.Owner = in.Owner
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PurchasePlan != nil {
		in, out := &in.PurchasePlan, &out.PurchasePlan
		*out = new(PurchasePlan)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(DiskSku)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disks_Spec.
func (in *Disks_Spec) DeepCopy() *Disks_Spec {
	if in == nil {
		return nil
	}
	out := new(Disks_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Encryption) DeepCopyInto(out *Encryption) {
	*out = *in
	if in.DiskEncryptionSetReference != nil {
		in, out := &in.DiskEncryptionSetReference, &out.DiskEncryptionSetReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Encryption.
func (in *Encryption) DeepCopy() *Encryption {
	if in == nil {
		return nil
	}
	out := new(Encryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSettingsCollection) DeepCopyInto(out *EncryptionSettingsCollection) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionSettings != nil {
		in, out := &in.EncryptionSettings, &out.EncryptionSettings
		*out = make([]EncryptionSettingsElement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EncryptionSettingsVersion != nil {
		in, out := &in.EncryptionSettingsVersion, &out.EncryptionSettingsVersion
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSettingsCollection.
func (in *EncryptionSettingsCollection) DeepCopy() *EncryptionSettingsCollection {
	if in == nil {
		return nil
	}
	out := new(EncryptionSettingsCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSettingsCollection_Status) DeepCopyInto(out *EncryptionSettingsCollection_Status) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionSettings != nil {
		in, out := &in.EncryptionSettings, &out.EncryptionSettings
		*out = make([]EncryptionSettingsElement_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EncryptionSettingsVersion != nil {
		in, out := &in.EncryptionSettingsVersion, &out.EncryptionSettingsVersion
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSettingsCollection_Status.
func (in *EncryptionSettingsCollection_Status) DeepCopy() *EncryptionSettingsCollection_Status {
	if in == nil {
		return nil
	}
	out := new(EncryptionSettingsCollection_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSettingsElement) DeepCopyInto(out *EncryptionSettingsElement) {
	*out = *in
	if in.DiskEncryptionKey != nil {
		in, out := &in.DiskEncryptionKey, &out.DiskEncryptionKey
		*out = new(KeyVaultAndSecretReference)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyEncryptionKey != nil {
		in, out := &in.KeyEncryptionKey, &out.KeyEncryptionKey
		*out = new(KeyVaultAndKeyReference)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSettingsElement.
func (in *EncryptionSettingsElement) DeepCopy() *EncryptionSettingsElement {
	if in == nil {
		return nil
	}
	out := new(EncryptionSettingsElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSettingsElement_Status) DeepCopyInto(out *EncryptionSettingsElement_Status) {
	*out = *in
	if in.DiskEncryptionKey != nil {
		in, out := &in.DiskEncryptionKey, &out.DiskEncryptionKey
		*out = new(KeyVaultAndSecretReference_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyEncryptionKey != nil {
		in, out := &in.KeyEncryptionKey, &out.KeyEncryptionKey
		*out = new(KeyVaultAndKeyReference_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSettingsElement_Status.
func (in *EncryptionSettingsElement_Status) DeepCopy() *EncryptionSettingsElement_Status {
	if in == nil {
		return nil
	}
	out := new(EncryptionSettingsElement_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Encryption_Status) DeepCopyInto(out *Encryption_Status) {
	*out = *in
	if in.DiskEncryptionSetId != nil {
		in, out := &in.DiskEncryptionSetId, &out.DiskEncryptionSetId
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Encryption_Status.
func (in *Encryption_Status) DeepCopy() *Encryption_Status {
	if in == nil {
		return nil
	}
	out := new(Encryption_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedLocation) DeepCopyInto(out *ExtendedLocation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedLocation.
func (in *ExtendedLocation) DeepCopy() *ExtendedLocation {
	if in == nil {
		return nil
	}
	out := new(ExtendedLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedLocation_Status) DeepCopyInto(out *ExtendedLocation_Status) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedLocation_Status.
func (in *ExtendedLocation_Status) DeepCopy() *ExtendedLocation_Status {
	if in == nil {
		return nil
	}
	out := new(ExtendedLocation_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDiskReference) DeepCopyInto(out *ImageDiskReference) {
	*out = *in
	if in.Lun != nil {
		in, out := &in.Lun, &out.Lun
		*out = new(int)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Reference = in.Reference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDiskReference.
func (in *ImageDiskReference) DeepCopy() *ImageDiskReference {
	if in == nil {
		return nil
	}
	out := new(ImageDiskReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDiskReference_Status) DeepCopyInto(out *ImageDiskReference_Status) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Lun != nil {
		in, out := &in.Lun, &out.Lun
		*out = new(int)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDiskReference_Status.
func (in *ImageDiskReference_Status) DeepCopy() *ImageDiskReference_Status {
	if in == nil {
		return nil
	}
	out := new(ImageDiskReference_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultAndKeyReference) DeepCopyInto(out *KeyVaultAndKeyReference) {
	*out = *in
	if in.KeyUrl != nil {
		in, out := &in.KeyUrl, &out.KeyUrl
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourceVault != nil {
		in, out := &in.SourceVault, &out.SourceVault
		*out = new(SourceVault)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultAndKeyReference.
func (in *KeyVaultAndKeyReference) DeepCopy() *KeyVaultAndKeyReference {
	if in == nil {
		return nil
	}
	out := new(KeyVaultAndKeyReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultAndKeyReference_Status) DeepCopyInto(out *KeyVaultAndKeyReference_Status) {
	*out = *in
	if in.KeyUrl != nil {
		in, out := &in.KeyUrl, &out.KeyUrl
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourceVault != nil {
		in, out := &in.SourceVault, &out.SourceVault
		*out = new(SourceVault_Status)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultAndKeyReference_Status.
func (in *KeyVaultAndKeyReference_Status) DeepCopy() *KeyVaultAndKeyReference_Status {
	if in == nil {
		return nil
	}
	out := new(KeyVaultAndKeyReference_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultAndSecretReference) DeepCopyInto(out *KeyVaultAndSecretReference) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecretUrl != nil {
		in, out := &in.SecretUrl, &out.SecretUrl
		*out = new(string)
		**out = **in
	}
	if in.SourceVault != nil {
		in, out := &in.SourceVault, &out.SourceVault
		*out = new(SourceVault)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultAndSecretReference.
func (in *KeyVaultAndSecretReference) DeepCopy() *KeyVaultAndSecretReference {
	if in == nil {
		return nil
	}
	out := new(KeyVaultAndSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultAndSecretReference_Status) DeepCopyInto(out *KeyVaultAndSecretReference_Status) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecretUrl != nil {
		in, out := &in.SecretUrl, &out.SecretUrl
		*out = new(string)
		**out = **in
	}
	if in.SourceVault != nil {
		in, out := &in.SourceVault, &out.SourceVault
		*out = new(SourceVault_Status)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultAndSecretReference_Status.
func (in *KeyVaultAndSecretReference_Status) DeepCopy() *KeyVaultAndSecretReference_Status {
	if in == nil {
		return nil
	}
	out := new(KeyVaultAndSecretReference_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurchasePlan) DeepCopyInto(out *PurchasePlan) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Product != nil {
		in, out := &in.Product, &out.Product
		*out = new(string)
		**out = **in
	}
	if in.PromotionCode != nil {
		in, out := &in.PromotionCode, &out.PromotionCode
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurchasePlan.
func (in *PurchasePlan) DeepCopy() *PurchasePlan {
	if in == nil {
		return nil
	}
	out := new(PurchasePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurchasePlan_Status) DeepCopyInto(out *PurchasePlan_Status) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Product != nil {
		in, out := &in.Product, &out.Product
		*out = new(string)
		**out = **in
	}
	if in.PromotionCode != nil {
		in, out := &in.PromotionCode, &out.PromotionCode
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurchasePlan_Status.
func (in *PurchasePlan_Status) DeepCopy() *PurchasePlan_Status {
	if in == nil {
		return nil
	}
	out := new(PurchasePlan_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareInfoElement_Status) DeepCopyInto(out *ShareInfoElement_Status) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VmUri != nil {
		in, out := &in.VmUri, &out.VmUri
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareInfoElement_Status.
func (in *ShareInfoElement_Status) DeepCopy() *ShareInfoElement_Status {
	if in == nil {
		return nil
	}
	out := new(ShareInfoElement_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceVault) DeepCopyInto(out *SourceVault) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceVault.
func (in *SourceVault) DeepCopy() *SourceVault {
	if in == nil {
		return nil
	}
	out := new(SourceVault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceVault_Status) DeepCopyInto(out *SourceVault_Status) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceVault_Status.
func (in *SourceVault_Status) DeepCopy() *SourceVault_Status {
	if in == nil {
		return nil
	}
	out := new(SourceVault_Status)
	in.DeepCopyInto(out)
	return out
}
