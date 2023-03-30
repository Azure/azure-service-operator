//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20200202storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Component) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentList) DeepCopyInto(out *ComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Component, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentList.
func (in *ComponentList) DeepCopy() *ComponentList {
	if in == nil {
		return nil
	}
	out := new(ComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component_STATUS) DeepCopyInto(out *Component_STATUS) {
	*out = *in
	if in.AppId != nil {
		in, out := &in.AppId, &out.AppId
		*out = new(string)
		**out = **in
	}
	if in.ApplicationId != nil {
		in, out := &in.ApplicationId, &out.ApplicationId
		*out = new(string)
		**out = **in
	}
	if in.Application_Type != nil {
		in, out := &in.Application_Type, &out.Application_Type
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
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(string)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = new(string)
		**out = **in
	}
	if in.DisableIpMasking != nil {
		in, out := &in.DisableIpMasking, &out.DisableIpMasking
		*out = new(bool)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Flow_Type != nil {
		in, out := &in.Flow_Type, &out.Flow_Type
		*out = new(string)
		**out = **in
	}
	if in.ForceCustomerStorageForProfiler != nil {
		in, out := &in.ForceCustomerStorageForProfiler, &out.ForceCustomerStorageForProfiler
		*out = new(bool)
		**out = **in
	}
	if in.HockeyAppId != nil {
		in, out := &in.HockeyAppId, &out.HockeyAppId
		*out = new(string)
		**out = **in
	}
	if in.HockeyAppToken != nil {
		in, out := &in.HockeyAppToken, &out.HockeyAppToken
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.ImmediatePurgeDataOn30Days != nil {
		in, out := &in.ImmediatePurgeDataOn30Days, &out.ImmediatePurgeDataOn30Days
		*out = new(bool)
		**out = **in
	}
	if in.IngestionMode != nil {
		in, out := &in.IngestionMode, &out.IngestionMode
		*out = new(string)
		**out = **in
	}
	if in.InstrumentationKey != nil {
		in, out := &in.InstrumentationKey, &out.InstrumentationKey
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.LaMigrationDate != nil {
		in, out := &in.LaMigrationDate, &out.LaMigrationDate
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
	if in.PrivateLinkScopedResources != nil {
		in, out := &in.PrivateLinkScopedResources, &out.PrivateLinkScopedResources
		*out = make([]PrivateLinkScopedResource_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PropertiesName != nil {
		in, out := &in.PropertiesName, &out.PropertiesName
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
	if in.PublicNetworkAccessForIngestion != nil {
		in, out := &in.PublicNetworkAccessForIngestion, &out.PublicNetworkAccessForIngestion
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessForQuery != nil {
		in, out := &in.PublicNetworkAccessForQuery, &out.PublicNetworkAccessForQuery
		*out = new(string)
		**out = **in
	}
	if in.Request_Source != nil {
		in, out := &in.Request_Source, &out.Request_Source
		*out = new(string)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int)
		**out = **in
	}
	if in.SamplingPercentage != nil {
		in, out := &in.SamplingPercentage, &out.SamplingPercentage
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceResourceId != nil {
		in, out := &in.WorkspaceResourceId, &out.WorkspaceResourceId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component_STATUS.
func (in *Component_STATUS) DeepCopy() *Component_STATUS {
	if in == nil {
		return nil
	}
	out := new(Component_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component_Spec) DeepCopyInto(out *Component_Spec) {
	*out = *in
	if in.Application_Type != nil {
		in, out := &in.Application_Type, &out.Application_Type
		*out = new(string)
		**out = **in
	}
	if in.DisableIpMasking != nil {
		in, out := &in.DisableIpMasking, &out.DisableIpMasking
		*out = new(bool)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Flow_Type != nil {
		in, out := &in.Flow_Type, &out.Flow_Type
		*out = new(string)
		**out = **in
	}
	if in.ForceCustomerStorageForProfiler != nil {
		in, out := &in.ForceCustomerStorageForProfiler, &out.ForceCustomerStorageForProfiler
		*out = new(bool)
		**out = **in
	}
	if in.HockeyAppId != nil {
		in, out := &in.HockeyAppId, &out.HockeyAppId
		*out = new(string)
		**out = **in
	}
	if in.ImmediatePurgeDataOn30Days != nil {
		in, out := &in.ImmediatePurgeDataOn30Days, &out.ImmediatePurgeDataOn30Days
		*out = new(bool)
		**out = **in
	}
	if in.IngestionMode != nil {
		in, out := &in.IngestionMode, &out.IngestionMode
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
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
	if in.PublicNetworkAccessForIngestion != nil {
		in, out := &in.PublicNetworkAccessForIngestion, &out.PublicNetworkAccessForIngestion
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessForQuery != nil {
		in, out := &in.PublicNetworkAccessForQuery, &out.PublicNetworkAccessForQuery
		*out = new(string)
		**out = **in
	}
	if in.Request_Source != nil {
		in, out := &in.Request_Source, &out.Request_Source
		*out = new(string)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int)
		**out = **in
	}
	if in.SamplingPercentage != nil {
		in, out := &in.SamplingPercentage, &out.SamplingPercentage
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.WorkspaceResourceReference != nil {
		in, out := &in.WorkspaceResourceReference, &out.WorkspaceResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component_Spec.
func (in *Component_Spec) DeepCopy() *Component_Spec {
	if in == nil {
		return nil
	}
	out := new(Component_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkScopedResource_STATUS) DeepCopyInto(out *PrivateLinkScopedResource_STATUS) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ResourceId != nil {
		in, out := &in.ResourceId, &out.ResourceId
		*out = new(string)
		**out = **in
	}
	if in.ScopeId != nil {
		in, out := &in.ScopeId, &out.ScopeId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateLinkScopedResource_STATUS.
func (in *PrivateLinkScopedResource_STATUS) DeepCopy() *PrivateLinkScopedResource_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateLinkScopedResource_STATUS)
	in.DeepCopyInto(out)
	return out
}
