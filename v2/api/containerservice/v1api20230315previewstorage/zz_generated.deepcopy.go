//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20230315previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorAdditionalInfo_STATUS) DeepCopyInto(out *ErrorAdditionalInfo_STATUS) {
	*out = *in
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorAdditionalInfo_STATUS.
func (in *ErrorAdditionalInfo_STATUS) DeepCopy() *ErrorAdditionalInfo_STATUS {
	if in == nil {
		return nil
	}
	out := new(ErrorAdditionalInfo_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorDetail_STATUS) DeepCopyInto(out *ErrorDetail_STATUS) {
	*out = *in
	if in.AdditionalInfo != nil {
		in, out := &in.AdditionalInfo, &out.AdditionalInfo
		*out = make([]ErrorAdditionalInfo_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]ErrorDetail_STATUS_Unrolled, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
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
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorDetail_STATUS.
func (in *ErrorDetail_STATUS) DeepCopy() *ErrorDetail_STATUS {
	if in == nil {
		return nil
	}
	out := new(ErrorDetail_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorDetail_STATUS_Unrolled) DeepCopyInto(out *ErrorDetail_STATUS_Unrolled) {
	*out = *in
	if in.AdditionalInfo != nil {
		in, out := &in.AdditionalInfo, &out.AdditionalInfo
		*out = make([]ErrorAdditionalInfo_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
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
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorDetail_STATUS_Unrolled.
func (in *ErrorDetail_STATUS_Unrolled) DeepCopy() *ErrorDetail_STATUS_Unrolled {
	if in == nil {
		return nil
	}
	out := new(ErrorDetail_STATUS_Unrolled)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleet) DeepCopyInto(out *Fleet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleet.
func (in *Fleet) DeepCopy() *Fleet {
	if in == nil {
		return nil
	}
	out := new(Fleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Fleet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetHubProfile) DeepCopyInto(out *FleetHubProfile) {
	*out = *in
	if in.DnsPrefix != nil {
		in, out := &in.DnsPrefix, &out.DnsPrefix
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetHubProfile.
func (in *FleetHubProfile) DeepCopy() *FleetHubProfile {
	if in == nil {
		return nil
	}
	out := new(FleetHubProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetHubProfile_STATUS) DeepCopyInto(out *FleetHubProfile_STATUS) {
	*out = *in
	if in.DnsPrefix != nil {
		in, out := &in.DnsPrefix, &out.DnsPrefix
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetHubProfile_STATUS.
func (in *FleetHubProfile_STATUS) DeepCopy() *FleetHubProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(FleetHubProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetList) DeepCopyInto(out *FleetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Fleet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetList.
func (in *FleetList) DeepCopy() *FleetList {
	if in == nil {
		return nil
	}
	out := new(FleetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetOperatorSecrets) DeepCopyInto(out *FleetOperatorSecrets) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UserCredentials != nil {
		in, out := &in.UserCredentials, &out.UserCredentials
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetOperatorSecrets.
func (in *FleetOperatorSecrets) DeepCopy() *FleetOperatorSecrets {
	if in == nil {
		return nil
	}
	out := new(FleetOperatorSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetOperatorSpec) DeepCopyInto(out *FleetOperatorSpec) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(FleetOperatorSecrets)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetOperatorSpec.
func (in *FleetOperatorSpec) DeepCopy() *FleetOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(FleetOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleet_STATUS) DeepCopyInto(out *Fleet_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ETag != nil {
		in, out := &in.ETag, &out.ETag
		*out = new(string)
		**out = **in
	}
	if in.HubProfile != nil {
		in, out := &in.HubProfile, &out.HubProfile
		*out = new(FleetHubProfile_STATUS)
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
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleet_STATUS.
func (in *Fleet_STATUS) DeepCopy() *Fleet_STATUS {
	if in == nil {
		return nil
	}
	out := new(Fleet_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleet_Spec) DeepCopyInto(out *Fleet_Spec) {
	*out = *in
	if in.HubProfile != nil {
		in, out := &in.HubProfile, &out.HubProfile
		*out = new(FleetHubProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(FleetOperatorSpec)
		(*in).DeepCopyInto(*out)
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
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleet_Spec.
func (in *Fleet_Spec) DeepCopy() *Fleet_Spec {
	if in == nil {
		return nil
	}
	out := new(Fleet_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetsMember) DeepCopyInto(out *FleetsMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetsMember.
func (in *FleetsMember) DeepCopy() *FleetsMember {
	if in == nil {
		return nil
	}
	out := new(FleetsMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetsMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetsMemberList) DeepCopyInto(out *FleetsMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FleetsMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetsMemberList.
func (in *FleetsMemberList) DeepCopy() *FleetsMemberList {
	if in == nil {
		return nil
	}
	out := new(FleetsMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetsMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetsUpdateRun) DeepCopyInto(out *FleetsUpdateRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetsUpdateRun.
func (in *FleetsUpdateRun) DeepCopy() *FleetsUpdateRun {
	if in == nil {
		return nil
	}
	out := new(FleetsUpdateRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetsUpdateRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetsUpdateRunList) DeepCopyInto(out *FleetsUpdateRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FleetsUpdateRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetsUpdateRunList.
func (in *FleetsUpdateRunList) DeepCopy() *FleetsUpdateRunList {
	if in == nil {
		return nil
	}
	out := new(FleetsUpdateRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetsUpdateRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleets_Member_STATUS) DeepCopyInto(out *Fleets_Member_STATUS) {
	*out = *in
	if in.ClusterResourceId != nil {
		in, out := &in.ClusterResourceId, &out.ClusterResourceId
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ETag != nil {
		in, out := &in.ETag, &out.ETag
		*out = new(string)
		**out = **in
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
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
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleets_Member_STATUS.
func (in *Fleets_Member_STATUS) DeepCopy() *Fleets_Member_STATUS {
	if in == nil {
		return nil
	}
	out := new(Fleets_Member_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleets_Member_Spec) DeepCopyInto(out *Fleets_Member_Spec) {
	*out = *in
	if in.ClusterResourceReference != nil {
		in, out := &in.ClusterResourceReference, &out.ClusterResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleets_Member_Spec.
func (in *Fleets_Member_Spec) DeepCopy() *Fleets_Member_Spec {
	if in == nil {
		return nil
	}
	out := new(Fleets_Member_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleets_UpdateRun_STATUS) DeepCopyInto(out *Fleets_UpdateRun_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ETag != nil {
		in, out := &in.ETag, &out.ETag
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.ManagedClusterUpdate != nil {
		in, out := &in.ManagedClusterUpdate, &out.ManagedClusterUpdate
		*out = new(ManagedClusterUpdate_STATUS)
		(*in).DeepCopyInto(*out)
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
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(UpdateRunStatus_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(UpdateRunStrategy_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleets_UpdateRun_STATUS.
func (in *Fleets_UpdateRun_STATUS) DeepCopy() *Fleets_UpdateRun_STATUS {
	if in == nil {
		return nil
	}
	out := new(Fleets_UpdateRun_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleets_UpdateRun_Spec) DeepCopyInto(out *Fleets_UpdateRun_Spec) {
	*out = *in
	if in.ManagedClusterUpdate != nil {
		in, out := &in.ManagedClusterUpdate, &out.ManagedClusterUpdate
		*out = new(ManagedClusterUpdate)
		(*in).DeepCopyInto(*out)
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
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(UpdateRunStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleets_UpdateRun_Spec.
func (in *Fleets_UpdateRun_Spec) DeepCopy() *Fleets_UpdateRun_Spec {
	if in == nil {
		return nil
	}
	out := new(Fleets_UpdateRun_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterUpdate) DeepCopyInto(out *ManagedClusterUpdate) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(ManagedClusterUpgradeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterUpdate.
func (in *ManagedClusterUpdate) DeepCopy() *ManagedClusterUpdate {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterUpdate_STATUS) DeepCopyInto(out *ManagedClusterUpdate_STATUS) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(ManagedClusterUpgradeSpec_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterUpdate_STATUS.
func (in *ManagedClusterUpdate_STATUS) DeepCopy() *ManagedClusterUpdate_STATUS {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterUpdate_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterUpgradeSpec) DeepCopyInto(out *ManagedClusterUpgradeSpec) {
	*out = *in
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterUpgradeSpec.
func (in *ManagedClusterUpgradeSpec) DeepCopy() *ManagedClusterUpgradeSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterUpgradeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterUpgradeSpec_STATUS) DeepCopyInto(out *ManagedClusterUpgradeSpec_STATUS) {
	*out = *in
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterUpgradeSpec_STATUS.
func (in *ManagedClusterUpgradeSpec_STATUS) DeepCopy() *ManagedClusterUpgradeSpec_STATUS {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterUpgradeSpec_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberUpdateStatus_STATUS) DeepCopyInto(out *MemberUpdateStatus_STATUS) {
	*out = *in
	if in.ClusterResourceId != nil {
		in, out := &in.ClusterResourceId, &out.ClusterResourceId
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OperationId != nil {
		in, out := &in.OperationId, &out.OperationId
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
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(UpdateStatus_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberUpdateStatus_STATUS.
func (in *MemberUpdateStatus_STATUS) DeepCopy() *MemberUpdateStatus_STATUS {
	if in == nil {
		return nil
	}
	out := new(MemberUpdateStatus_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_STATUS) DeepCopyInto(out *SystemData_STATUS) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedByType != nil {
		in, out := &in.CreatedByType, &out.CreatedByType
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedAt != nil {
		in, out := &in.LastModifiedAt, &out.LastModifiedAt
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedByType != nil {
		in, out := &in.LastModifiedByType, &out.LastModifiedByType
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS.
func (in *SystemData_STATUS) DeepCopy() *SystemData_STATUS {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateGroup) DeepCopyInto(out *UpdateGroup) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateGroup.
func (in *UpdateGroup) DeepCopy() *UpdateGroup {
	if in == nil {
		return nil
	}
	out := new(UpdateGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateGroupStatus_STATUS) DeepCopyInto(out *UpdateGroupStatus_STATUS) {
	*out = *in
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]MemberUpdateStatus_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(UpdateStatus_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateGroupStatus_STATUS.
func (in *UpdateGroupStatus_STATUS) DeepCopy() *UpdateGroupStatus_STATUS {
	if in == nil {
		return nil
	}
	out := new(UpdateGroupStatus_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateGroup_STATUS) DeepCopyInto(out *UpdateGroup_STATUS) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateGroup_STATUS.
func (in *UpdateGroup_STATUS) DeepCopy() *UpdateGroup_STATUS {
	if in == nil {
		return nil
	}
	out := new(UpdateGroup_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRunStatus_STATUS) DeepCopyInto(out *UpdateRunStatus_STATUS) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]UpdateStageStatus_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(UpdateStatus_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRunStatus_STATUS.
func (in *UpdateRunStatus_STATUS) DeepCopy() *UpdateRunStatus_STATUS {
	if in == nil {
		return nil
	}
	out := new(UpdateRunStatus_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRunStrategy) DeepCopyInto(out *UpdateRunStrategy) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]UpdateStage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRunStrategy.
func (in *UpdateRunStrategy) DeepCopy() *UpdateRunStrategy {
	if in == nil {
		return nil
	}
	out := new(UpdateRunStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRunStrategy_STATUS) DeepCopyInto(out *UpdateRunStrategy_STATUS) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]UpdateStage_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRunStrategy_STATUS.
func (in *UpdateRunStrategy_STATUS) DeepCopy() *UpdateRunStrategy_STATUS {
	if in == nil {
		return nil
	}
	out := new(UpdateRunStrategy_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateStage) DeepCopyInto(out *UpdateStage) {
	*out = *in
	if in.AfterStageWaitInSeconds != nil {
		in, out := &in.AfterStageWaitInSeconds, &out.AfterStageWaitInSeconds
		*out = new(int)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]UpdateGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateStage.
func (in *UpdateStage) DeepCopy() *UpdateStage {
	if in == nil {
		return nil
	}
	out := new(UpdateStage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateStageStatus_STATUS) DeepCopyInto(out *UpdateStageStatus_STATUS) {
	*out = *in
	if in.AfterStageWaitStatus != nil {
		in, out := &in.AfterStageWaitStatus, &out.AfterStageWaitStatus
		*out = new(WaitStatus_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]UpdateGroupStatus_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(UpdateStatus_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateStageStatus_STATUS.
func (in *UpdateStageStatus_STATUS) DeepCopy() *UpdateStageStatus_STATUS {
	if in == nil {
		return nil
	}
	out := new(UpdateStageStatus_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateStage_STATUS) DeepCopyInto(out *UpdateStage_STATUS) {
	*out = *in
	if in.AfterStageWaitInSeconds != nil {
		in, out := &in.AfterStageWaitInSeconds, &out.AfterStageWaitInSeconds
		*out = new(int)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]UpdateGroup_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateStage_STATUS.
func (in *UpdateStage_STATUS) DeepCopy() *UpdateStage_STATUS {
	if in == nil {
		return nil
	}
	out := new(UpdateStage_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateStatus_STATUS) DeepCopyInto(out *UpdateStatus_STATUS) {
	*out = *in
	if in.CompletedTime != nil {
		in, out := &in.CompletedTime, &out.CompletedTime
		*out = new(string)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(ErrorDetail_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateStatus_STATUS.
func (in *UpdateStatus_STATUS) DeepCopy() *UpdateStatus_STATUS {
	if in == nil {
		return nil
	}
	out := new(UpdateStatus_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaitStatus_STATUS) DeepCopyInto(out *WaitStatus_STATUS) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(UpdateStatus_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.WaitDurationInSeconds != nil {
		in, out := &in.WaitDurationInSeconds, &out.WaitDurationInSeconds
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaitStatus_STATUS.
func (in *WaitStatus_STATUS) DeepCopy() *WaitStatus_STATUS {
	if in == nil {
		return nil
	}
	out := new(WaitStatus_STATUS)
	in.DeepCopyInto(out)
	return out
}
