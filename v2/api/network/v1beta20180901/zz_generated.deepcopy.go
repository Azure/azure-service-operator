//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta20180901

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsZone) DeepCopyInto(out *PrivateDnsZone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsZone.
func (in *PrivateDnsZone) DeepCopy() *PrivateDnsZone {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateDnsZone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsZoneList) DeepCopyInto(out *PrivateDnsZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrivateDnsZone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsZoneList.
func (in *PrivateDnsZoneList) DeepCopy() *PrivateDnsZoneList {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateDnsZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
<<<<<<< HEAD
func (in *PrivateDnsZone_STATUS) DeepCopyInto(out *PrivateDnsZone_STATUS) {
=======
func (in *PrivateDnsZones_Spec) DeepCopyInto(out *PrivateDnsZones_Spec) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
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
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsZones_Spec.
func (in *PrivateDnsZones_Spec) DeepCopy() *PrivateDnsZones_Spec {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsZones_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsZones_SpecARM) DeepCopyInto(out *PrivateDnsZones_SpecARM) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsZones_SpecARM.
func (in *PrivateDnsZones_SpecARM) DeepCopy() *PrivateDnsZones_SpecARM {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsZones_SpecARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateZoneProperties_STATUSARM) DeepCopyInto(out *PrivateZoneProperties_STATUSARM) {
	*out = *in
	if in.MaxNumberOfRecordSets != nil {
		in, out := &in.MaxNumberOfRecordSets, &out.MaxNumberOfRecordSets
		*out = new(int)
		**out = **in
	}
	if in.MaxNumberOfVirtualNetworkLinks != nil {
		in, out := &in.MaxNumberOfVirtualNetworkLinks, &out.MaxNumberOfVirtualNetworkLinks
		*out = new(int)
		**out = **in
	}
	if in.MaxNumberOfVirtualNetworkLinksWithRegistration != nil {
		in, out := &in.MaxNumberOfVirtualNetworkLinksWithRegistration, &out.MaxNumberOfVirtualNetworkLinksWithRegistration
		*out = new(int)
		**out = **in
	}
	if in.NumberOfRecordSets != nil {
		in, out := &in.NumberOfRecordSets, &out.NumberOfRecordSets
		*out = new(int)
		**out = **in
	}
	if in.NumberOfVirtualNetworkLinks != nil {
		in, out := &in.NumberOfVirtualNetworkLinks, &out.NumberOfVirtualNetworkLinks
		*out = new(int)
		**out = **in
	}
	if in.NumberOfVirtualNetworkLinksWithRegistration != nil {
		in, out := &in.NumberOfVirtualNetworkLinksWithRegistration, &out.NumberOfVirtualNetworkLinksWithRegistration
		*out = new(int)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(PrivateZoneProperties_STATUS_ProvisioningState)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateZoneProperties_STATUSARM.
func (in *PrivateZoneProperties_STATUSARM) DeepCopy() *PrivateZoneProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(PrivateZoneProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateZone_STATUS) DeepCopyInto(out *PrivateZone_STATUS) {
>>>>>>> main
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.MaxNumberOfRecordSets != nil {
		in, out := &in.MaxNumberOfRecordSets, &out.MaxNumberOfRecordSets
		*out = new(int)
		**out = **in
	}
	if in.MaxNumberOfVirtualNetworkLinks != nil {
		in, out := &in.MaxNumberOfVirtualNetworkLinks, &out.MaxNumberOfVirtualNetworkLinks
		*out = new(int)
		**out = **in
	}
	if in.MaxNumberOfVirtualNetworkLinksWithRegistration != nil {
		in, out := &in.MaxNumberOfVirtualNetworkLinksWithRegistration, &out.MaxNumberOfVirtualNetworkLinksWithRegistration
		*out = new(int)
		**out = **in
	}
	if in.NumberOfRecordSets != nil {
		in, out := &in.NumberOfRecordSets, &out.NumberOfRecordSets
		*out = new(int)
		**out = **in
	}
	if in.NumberOfVirtualNetworkLinks != nil {
		in, out := &in.NumberOfVirtualNetworkLinks, &out.NumberOfVirtualNetworkLinks
		*out = new(int)
		**out = **in
	}
	if in.NumberOfVirtualNetworkLinksWithRegistration != nil {
		in, out := &in.NumberOfVirtualNetworkLinksWithRegistration, &out.NumberOfVirtualNetworkLinksWithRegistration
		*out = new(int)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
<<<<<<< HEAD
		*out = new(PrivateZoneProperties_ProvisioningState_STATUS)
=======
		*out = new(PrivateZoneProperties_STATUS_ProvisioningState)
>>>>>>> main
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsZone_STATUS.
func (in *PrivateDnsZone_STATUS) DeepCopy() *PrivateDnsZone_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsZone_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsZone_STATUSARM) DeepCopyInto(out *PrivateDnsZone_STATUSARM) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(PrivateZoneProperties_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsZone_STATUSARM.
func (in *PrivateDnsZone_STATUSARM) DeepCopy() *PrivateDnsZone_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsZone_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsZone_Spec) DeepCopyInto(out *PrivateDnsZone_Spec) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsZone_Spec.
func (in *PrivateDnsZone_Spec) DeepCopy() *PrivateDnsZone_Spec {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsZone_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsZone_SpecARM) DeepCopyInto(out *PrivateDnsZone_SpecARM) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsZone_SpecARM.
func (in *PrivateDnsZone_SpecARM) DeepCopy() *PrivateDnsZone_SpecARM {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsZone_SpecARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateZoneProperties_STATUSARM) DeepCopyInto(out *PrivateZoneProperties_STATUSARM) {
	*out = *in
	if in.MaxNumberOfRecordSets != nil {
		in, out := &in.MaxNumberOfRecordSets, &out.MaxNumberOfRecordSets
		*out = new(int)
		**out = **in
	}
	if in.MaxNumberOfVirtualNetworkLinks != nil {
		in, out := &in.MaxNumberOfVirtualNetworkLinks, &out.MaxNumberOfVirtualNetworkLinks
		*out = new(int)
		**out = **in
	}
	if in.MaxNumberOfVirtualNetworkLinksWithRegistration != nil {
		in, out := &in.MaxNumberOfVirtualNetworkLinksWithRegistration, &out.MaxNumberOfVirtualNetworkLinksWithRegistration
		*out = new(int)
		**out = **in
	}
	if in.NumberOfRecordSets != nil {
		in, out := &in.NumberOfRecordSets, &out.NumberOfRecordSets
		*out = new(int)
		**out = **in
	}
	if in.NumberOfVirtualNetworkLinks != nil {
		in, out := &in.NumberOfVirtualNetworkLinks, &out.NumberOfVirtualNetworkLinks
		*out = new(int)
		**out = **in
	}
	if in.NumberOfVirtualNetworkLinksWithRegistration != nil {
		in, out := &in.NumberOfVirtualNetworkLinksWithRegistration, &out.NumberOfVirtualNetworkLinksWithRegistration
		*out = new(int)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(PrivateZoneProperties_ProvisioningState_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateZoneProperties_STATUSARM.
func (in *PrivateZoneProperties_STATUSARM) DeepCopy() *PrivateZoneProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(PrivateZoneProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}
