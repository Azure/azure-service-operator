//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20210701

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedLocation) DeepCopyInto(out *ExtendedLocation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ExtendedLocationType)
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
func (in *ExtendedLocationARM) DeepCopyInto(out *ExtendedLocationARM) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ExtendedLocationType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedLocationARM.
func (in *ExtendedLocationARM) DeepCopy() *ExtendedLocationARM {
	if in == nil {
		return nil
	}
	out := new(ExtendedLocationARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedLocation_STATUS) DeepCopyInto(out *ExtendedLocation_STATUS) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ExtendedLocationType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedLocation_STATUS.
func (in *ExtendedLocation_STATUS) DeepCopy() *ExtendedLocation_STATUS {
	if in == nil {
		return nil
	}
	out := new(ExtendedLocation_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedLocation_STATUSARM) DeepCopyInto(out *ExtendedLocation_STATUSARM) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ExtendedLocationType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedLocation_STATUSARM.
func (in *ExtendedLocation_STATUSARM) DeepCopy() *ExtendedLocation_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(ExtendedLocation_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDataDisk) DeepCopyInto(out *ImageDataDisk) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(ImageDataDisk_Caching)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.Lun != nil {
		in, out := &in.Lun, &out.Lun
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(StorageAccountType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDataDisk.
func (in *ImageDataDisk) DeepCopy() *ImageDataDisk {
	if in == nil {
		return nil
	}
	out := new(ImageDataDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDataDiskARM) DeepCopyInto(out *ImageDataDiskARM) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(ImageDataDisk_Caching)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResourceARM)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.Lun != nil {
		in, out := &in.Lun, &out.Lun
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResourceARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResourceARM)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(StorageAccountType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDataDiskARM.
func (in *ImageDataDiskARM) DeepCopy() *ImageDataDiskARM {
	if in == nil {
		return nil
	}
	out := new(ImageDataDiskARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDataDisk_STATUS) DeepCopyInto(out *ImageDataDisk_STATUS) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(ImageDataDisk_Caching_STATUS)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.Lun != nil {
		in, out := &in.Lun, &out.Lun
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(StorageAccountType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDataDisk_STATUS.
func (in *ImageDataDisk_STATUS) DeepCopy() *ImageDataDisk_STATUS {
	if in == nil {
		return nil
	}
	out := new(ImageDataDisk_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDataDisk_STATUSARM) DeepCopyInto(out *ImageDataDisk_STATUSARM) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(ImageDataDisk_Caching_STATUS)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.Lun != nil {
		in, out := &in.Lun, &out.Lun
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(StorageAccountType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDataDisk_STATUSARM.
func (in *ImageDataDisk_STATUSARM) DeepCopy() *ImageDataDisk_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(ImageDataDisk_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageList) DeepCopyInto(out *ImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Image, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageList.
func (in *ImageList) DeepCopy() *ImageList {
	if in == nil {
		return nil
	}
	out := new(ImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageOSDisk) DeepCopyInto(out *ImageOSDisk) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(ImageOSDisk_Caching)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.OsState != nil {
		in, out := &in.OsState, &out.OsState
		*out = new(ImageOSDisk_OsState)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(ImageOSDisk_OsType)
		**out = **in
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(StorageAccountType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageOSDisk.
func (in *ImageOSDisk) DeepCopy() *ImageOSDisk {
	if in == nil {
		return nil
	}
	out := new(ImageOSDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageOSDiskARM) DeepCopyInto(out *ImageOSDiskARM) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(ImageOSDisk_Caching)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResourceARM)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResourceARM)
		(*in).DeepCopyInto(*out)
	}
	if in.OsState != nil {
		in, out := &in.OsState, &out.OsState
		*out = new(ImageOSDisk_OsState)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(ImageOSDisk_OsType)
		**out = **in
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResourceARM)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(StorageAccountType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageOSDiskARM.
func (in *ImageOSDiskARM) DeepCopy() *ImageOSDiskARM {
	if in == nil {
		return nil
	}
	out := new(ImageOSDiskARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageOSDisk_STATUS) DeepCopyInto(out *ImageOSDisk_STATUS) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(ImageOSDisk_Caching_STATUS)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.OsState != nil {
		in, out := &in.OsState, &out.OsState
		*out = new(ImageOSDisk_OsState_STATUS)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(ImageOSDisk_OsType_STATUS)
		**out = **in
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(StorageAccountType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageOSDisk_STATUS.
func (in *ImageOSDisk_STATUS) DeepCopy() *ImageOSDisk_STATUS {
	if in == nil {
		return nil
	}
	out := new(ImageOSDisk_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageOSDisk_STATUSARM) DeepCopyInto(out *ImageOSDisk_STATUSARM) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(ImageOSDisk_Caching_STATUS)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.OsState != nil {
		in, out := &in.OsState, &out.OsState
		*out = new(ImageOSDisk_OsState_STATUS)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(ImageOSDisk_OsType_STATUS)
		**out = **in
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(StorageAccountType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageOSDisk_STATUSARM.
func (in *ImageOSDisk_STATUSARM) DeepCopy() *ImageOSDisk_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(ImageOSDisk_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePropertiesARM) DeepCopyInto(out *ImagePropertiesARM) {
	*out = *in
	if in.HyperVGeneration != nil {
		in, out := &in.HyperVGeneration, &out.HyperVGeneration
		*out = new(HyperVGenerationType)
		**out = **in
	}
	if in.SourceVirtualMachine != nil {
		in, out := &in.SourceVirtualMachine, &out.SourceVirtualMachine
		*out = new(SubResourceARM)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(ImageStorageProfileARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePropertiesARM.
func (in *ImagePropertiesARM) DeepCopy() *ImagePropertiesARM {
	if in == nil {
		return nil
	}
	out := new(ImagePropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageProperties_STATUSARM) DeepCopyInto(out *ImageProperties_STATUSARM) {
	*out = *in
	if in.HyperVGeneration != nil {
		in, out := &in.HyperVGeneration, &out.HyperVGeneration
		*out = new(HyperVGenerationType_STATUS)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.SourceVirtualMachine != nil {
		in, out := &in.SourceVirtualMachine, &out.SourceVirtualMachine
		*out = new(SubResource_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(ImageStorageProfile_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageProperties_STATUSARM.
func (in *ImageProperties_STATUSARM) DeepCopy() *ImageProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(ImageProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStorageProfile) DeepCopyInto(out *ImageStorageProfile) {
	*out = *in
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]ImageDataDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OsDisk != nil {
		in, out := &in.OsDisk, &out.OsDisk
		*out = new(ImageOSDisk)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneResilient != nil {
		in, out := &in.ZoneResilient, &out.ZoneResilient
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStorageProfile.
func (in *ImageStorageProfile) DeepCopy() *ImageStorageProfile {
	if in == nil {
		return nil
	}
	out := new(ImageStorageProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStorageProfileARM) DeepCopyInto(out *ImageStorageProfileARM) {
	*out = *in
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]ImageDataDiskARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OsDisk != nil {
		in, out := &in.OsDisk, &out.OsDisk
		*out = new(ImageOSDiskARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneResilient != nil {
		in, out := &in.ZoneResilient, &out.ZoneResilient
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStorageProfileARM.
func (in *ImageStorageProfileARM) DeepCopy() *ImageStorageProfileARM {
	if in == nil {
		return nil
	}
	out := new(ImageStorageProfileARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStorageProfile_STATUS) DeepCopyInto(out *ImageStorageProfile_STATUS) {
	*out = *in
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]ImageDataDisk_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OsDisk != nil {
		in, out := &in.OsDisk, &out.OsDisk
		*out = new(ImageOSDisk_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneResilient != nil {
		in, out := &in.ZoneResilient, &out.ZoneResilient
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStorageProfile_STATUS.
func (in *ImageStorageProfile_STATUS) DeepCopy() *ImageStorageProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(ImageStorageProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStorageProfile_STATUSARM) DeepCopyInto(out *ImageStorageProfile_STATUSARM) {
	*out = *in
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]ImageDataDisk_STATUSARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OsDisk != nil {
		in, out := &in.OsDisk, &out.OsDisk
		*out = new(ImageOSDisk_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneResilient != nil {
		in, out := &in.ZoneResilient, &out.ZoneResilient
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStorageProfile_STATUSARM.
func (in *ImageStorageProfile_STATUSARM) DeepCopy() *ImageStorageProfile_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(ImageStorageProfile_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image_STATUS) DeepCopyInto(out *Image_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtendedLocation != nil {
		in, out := &in.ExtendedLocation, &out.ExtendedLocation
		*out = new(ExtendedLocation_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.HyperVGeneration != nil {
		in, out := &in.HyperVGeneration, &out.HyperVGeneration
		*out = new(HyperVGenerationType_STATUS)
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.SourceVirtualMachine != nil {
		in, out := &in.SourceVirtualMachine, &out.SourceVirtualMachine
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(ImageStorageProfile_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image_STATUS.
func (in *Image_STATUS) DeepCopy() *Image_STATUS {
	if in == nil {
		return nil
	}
	out := new(Image_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image_STATUSARM) DeepCopyInto(out *Image_STATUSARM) {
	*out = *in
	if in.ExtendedLocation != nil {
		in, out := &in.ExtendedLocation, &out.ExtendedLocation
		*out = new(ExtendedLocation_STATUSARM)
		(*in).DeepCopyInto(*out)
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ImageProperties_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image_STATUSARM.
func (in *Image_STATUSARM) DeepCopy() *Image_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(Image_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image_Spec) DeepCopyInto(out *Image_Spec) {
	*out = *in
	if in.ExtendedLocation != nil {
		in, out := &in.ExtendedLocation, &out.ExtendedLocation
		*out = new(ExtendedLocation)
		(*in).DeepCopyInto(*out)
	}
	if in.HyperVGeneration != nil {
		in, out := &in.HyperVGeneration, &out.HyperVGeneration
		*out = new(HyperVGenerationType)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.SourceVirtualMachine != nil {
		in, out := &in.SourceVirtualMachine, &out.SourceVirtualMachine
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(ImageStorageProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image_Spec.
func (in *Image_Spec) DeepCopy() *Image_Spec {
	if in == nil {
		return nil
	}
	out := new(Image_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image_SpecARM) DeepCopyInto(out *Image_SpecARM) {
	*out = *in
	if in.ExtendedLocation != nil {
		in, out := &in.ExtendedLocation, &out.ExtendedLocation
		*out = new(ExtendedLocationARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ImagePropertiesARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image_SpecARM.
func (in *Image_SpecARM) DeepCopy() *Image_SpecARM {
	if in == nil {
		return nil
	}
	out := new(Image_SpecARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubResource) DeepCopyInto(out *SubResource) {
	*out = *in
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResource.
func (in *SubResource) DeepCopy() *SubResource {
	if in == nil {
		return nil
	}
	out := new(SubResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubResourceARM) DeepCopyInto(out *SubResourceARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResourceARM.
func (in *SubResourceARM) DeepCopy() *SubResourceARM {
	if in == nil {
		return nil
	}
	out := new(SubResourceARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubResource_STATUS) DeepCopyInto(out *SubResource_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResource_STATUS.
func (in *SubResource_STATUS) DeepCopy() *SubResource_STATUS {
	if in == nil {
		return nil
	}
	out := new(SubResource_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubResource_STATUSARM) DeepCopyInto(out *SubResource_STATUSARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResource_STATUSARM.
func (in *SubResource_STATUSARM) DeepCopy() *SubResource_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(SubResource_STATUSARM)
	in.DeepCopyInto(out)
	return out
}
