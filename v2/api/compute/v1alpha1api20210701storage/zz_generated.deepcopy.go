//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20210701storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionSetParameters) DeepCopyInto(out *DiskEncryptionSetParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionSetParameters.
func (in *DiskEncryptionSetParameters) DeepCopy() *DiskEncryptionSetParameters {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionSetParameters)
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
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(DiskEncryptionSetParameters)
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(string)
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
func (in *ImageDataDisk_Status) DeepCopyInto(out *ImageDataDisk_Status) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource_Status)
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
		*out = new(SubResource_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDataDisk_Status.
func (in *ImageDataDisk_Status) DeepCopy() *ImageDataDisk_Status {
	if in == nil {
		return nil
	}
	out := new(ImageDataDisk_Status)
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
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(DiskEncryptionSetParameters)
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
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(string)
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
func (in *ImageOSDisk_Status) DeepCopyInto(out *ImageOSDisk_Status) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.OsState != nil {
		in, out := &in.OsState, &out.OsState
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
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageOSDisk_Status.
func (in *ImageOSDisk_Status) DeepCopy() *ImageOSDisk_Status {
	if in == nil {
		return nil
	}
	out := new(ImageOSDisk_Status)
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
func (in *ImageStorageProfile_Status) DeepCopyInto(out *ImageStorageProfile_Status) {
	*out = *in
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]ImageDataDisk_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OsDisk != nil {
		in, out := &in.OsDisk, &out.OsDisk
		*out = new(ImageOSDisk_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ZoneResilient != nil {
		in, out := &in.ZoneResilient, &out.ZoneResilient
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStorageProfile_Status.
func (in *ImageStorageProfile_Status) DeepCopy() *ImageStorageProfile_Status {
	if in == nil {
		return nil
	}
	out := new(ImageStorageProfile_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image_Status) DeepCopyInto(out *Image_Status) {
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
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.SourceVirtualMachine != nil {
		in, out := &in.SourceVirtualMachine, &out.SourceVirtualMachine
		*out = new(SubResource_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(ImageStorageProfile_Status)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image_Status.
func (in *Image_Status) DeepCopy() *Image_Status {
	if in == nil {
		return nil
	}
	out := new(Image_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Images_Spec) DeepCopyInto(out *Images_Spec) {
	*out = *in
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
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Images_Spec.
func (in *Images_Spec) DeepCopy() *Images_Spec {
	if in == nil {
		return nil
	}
	out := new(Images_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubResource) DeepCopyInto(out *SubResource) {
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
func (in *SubResource_Status) DeepCopyInto(out *SubResource_Status) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResource_Status.
func (in *SubResource_Status) DeepCopy() *SubResource_Status {
	if in == nil {
		return nil
	}
	out := new(SubResource_Status)
	in.DeepCopyInto(out)
	return out
}
