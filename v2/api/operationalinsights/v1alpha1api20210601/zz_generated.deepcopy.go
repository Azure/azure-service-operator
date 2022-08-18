//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20210601

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkScopedResource_STATUS) DeepCopyInto(out *PrivateLinkScopedResource_STATUS) {
	*out = *in
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkScopedResource_STATUSARM) DeepCopyInto(out *PrivateLinkScopedResource_STATUSARM) {
	*out = *in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateLinkScopedResource_STATUSARM.
func (in *PrivateLinkScopedResource_STATUSARM) DeepCopy() *PrivateLinkScopedResource_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(PrivateLinkScopedResource_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCapping) DeepCopyInto(out *WorkspaceCapping) {
	*out = *in
	if in.DailyQuotaGb != nil {
		in, out := &in.DailyQuotaGb, &out.DailyQuotaGb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCapping.
func (in *WorkspaceCapping) DeepCopy() *WorkspaceCapping {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCappingARM) DeepCopyInto(out *WorkspaceCappingARM) {
	*out = *in
	if in.DailyQuotaGb != nil {
		in, out := &in.DailyQuotaGb, &out.DailyQuotaGb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCappingARM.
func (in *WorkspaceCappingARM) DeepCopy() *WorkspaceCappingARM {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCappingARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCapping_STATUS) DeepCopyInto(out *WorkspaceCapping_STATUS) {
	*out = *in
	if in.DailyQuotaGb != nil {
		in, out := &in.DailyQuotaGb, &out.DailyQuotaGb
		*out = new(float64)
		**out = **in
	}
	if in.DataIngestionStatus != nil {
		in, out := &in.DataIngestionStatus, &out.DataIngestionStatus
		*out = new(WorkspaceCappingSTATUSDataIngestionStatus)
		**out = **in
	}
	if in.QuotaNextResetTime != nil {
		in, out := &in.QuotaNextResetTime, &out.QuotaNextResetTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCapping_STATUS.
func (in *WorkspaceCapping_STATUS) DeepCopy() *WorkspaceCapping_STATUS {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCapping_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCapping_STATUSARM) DeepCopyInto(out *WorkspaceCapping_STATUSARM) {
	*out = *in
	if in.DailyQuotaGb != nil {
		in, out := &in.DailyQuotaGb, &out.DailyQuotaGb
		*out = new(float64)
		**out = **in
	}
	if in.DataIngestionStatus != nil {
		in, out := &in.DataIngestionStatus, &out.DataIngestionStatus
		*out = new(WorkspaceCappingSTATUSDataIngestionStatus)
		**out = **in
	}
	if in.QuotaNextResetTime != nil {
		in, out := &in.QuotaNextResetTime, &out.QuotaNextResetTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCapping_STATUSARM.
func (in *WorkspaceCapping_STATUSARM) DeepCopy() *WorkspaceCapping_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCapping_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceFeatures) DeepCopyInto(out *WorkspaceFeatures) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ClusterResourceReference != nil {
		in, out := &in.ClusterResourceReference, &out.ClusterResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnableDataExport != nil {
		in, out := &in.EnableDataExport, &out.EnableDataExport
		*out = new(bool)
		**out = **in
	}
	if in.EnableLogAccessUsingOnlyResourcePermissions != nil {
		in, out := &in.EnableLogAccessUsingOnlyResourcePermissions, &out.EnableLogAccessUsingOnlyResourcePermissions
		*out = new(bool)
		**out = **in
	}
	if in.ImmediatePurgeDataOn30Days != nil {
		in, out := &in.ImmediatePurgeDataOn30Days, &out.ImmediatePurgeDataOn30Days
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceFeatures.
func (in *WorkspaceFeatures) DeepCopy() *WorkspaceFeatures {
	if in == nil {
		return nil
	}
	out := new(WorkspaceFeatures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceFeaturesARM) DeepCopyInto(out *WorkspaceFeaturesARM) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ClusterResourceId != nil {
		in, out := &in.ClusterResourceId, &out.ClusterResourceId
		*out = new(string)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnableDataExport != nil {
		in, out := &in.EnableDataExport, &out.EnableDataExport
		*out = new(bool)
		**out = **in
	}
	if in.EnableLogAccessUsingOnlyResourcePermissions != nil {
		in, out := &in.EnableLogAccessUsingOnlyResourcePermissions, &out.EnableLogAccessUsingOnlyResourcePermissions
		*out = new(bool)
		**out = **in
	}
	if in.ImmediatePurgeDataOn30Days != nil {
		in, out := &in.ImmediatePurgeDataOn30Days, &out.ImmediatePurgeDataOn30Days
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceFeaturesARM.
func (in *WorkspaceFeaturesARM) DeepCopy() *WorkspaceFeaturesARM {
	if in == nil {
		return nil
	}
	out := new(WorkspaceFeaturesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceFeatures_STATUS) DeepCopyInto(out *WorkspaceFeatures_STATUS) {
	*out = *in
	if in.ClusterResourceId != nil {
		in, out := &in.ClusterResourceId, &out.ClusterResourceId
		*out = new(string)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnableDataExport != nil {
		in, out := &in.EnableDataExport, &out.EnableDataExport
		*out = new(bool)
		**out = **in
	}
	if in.EnableLogAccessUsingOnlyResourcePermissions != nil {
		in, out := &in.EnableLogAccessUsingOnlyResourcePermissions, &out.EnableLogAccessUsingOnlyResourcePermissions
		*out = new(bool)
		**out = **in
	}
	if in.ImmediatePurgeDataOn30Days != nil {
		in, out := &in.ImmediatePurgeDataOn30Days, &out.ImmediatePurgeDataOn30Days
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceFeatures_STATUS.
func (in *WorkspaceFeatures_STATUS) DeepCopy() *WorkspaceFeatures_STATUS {
	if in == nil {
		return nil
	}
	out := new(WorkspaceFeatures_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceFeatures_STATUSARM) DeepCopyInto(out *WorkspaceFeatures_STATUSARM) {
	*out = *in
	if in.ClusterResourceId != nil {
		in, out := &in.ClusterResourceId, &out.ClusterResourceId
		*out = new(string)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnableDataExport != nil {
		in, out := &in.EnableDataExport, &out.EnableDataExport
		*out = new(bool)
		**out = **in
	}
	if in.EnableLogAccessUsingOnlyResourcePermissions != nil {
		in, out := &in.EnableLogAccessUsingOnlyResourcePermissions, &out.EnableLogAccessUsingOnlyResourcePermissions
		*out = new(bool)
		**out = **in
	}
	if in.ImmediatePurgeDataOn30Days != nil {
		in, out := &in.ImmediatePurgeDataOn30Days, &out.ImmediatePurgeDataOn30Days
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceFeatures_STATUSARM.
func (in *WorkspaceFeatures_STATUSARM) DeepCopy() *WorkspaceFeatures_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(WorkspaceFeatures_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceList) DeepCopyInto(out *WorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceList.
func (in *WorkspaceList) DeepCopy() *WorkspaceList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacePropertiesARM) DeepCopyInto(out *WorkspacePropertiesARM) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(WorkspaceFeaturesARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ForceCmkForQuery != nil {
		in, out := &in.ForceCmkForQuery, &out.ForceCmkForQuery
		*out = new(bool)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(WorkspacePropertiesProvisioningState)
		**out = **in
	}
	if in.PublicNetworkAccessForIngestion != nil {
		in, out := &in.PublicNetworkAccessForIngestion, &out.PublicNetworkAccessForIngestion
		*out = new(WorkspacePropertiesPublicNetworkAccessForIngestion)
		**out = **in
	}
	if in.PublicNetworkAccessForQuery != nil {
		in, out := &in.PublicNetworkAccessForQuery, &out.PublicNetworkAccessForQuery
		*out = new(WorkspacePropertiesPublicNetworkAccessForQuery)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(WorkspaceSkuARM)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceCapping != nil {
		in, out := &in.WorkspaceCapping, &out.WorkspaceCapping
		*out = new(WorkspaceCappingARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacePropertiesARM.
func (in *WorkspacePropertiesARM) DeepCopy() *WorkspacePropertiesARM {
	if in == nil {
		return nil
	}
	out := new(WorkspacePropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceProperties_STATUSARM) DeepCopyInto(out *WorkspaceProperties_STATUSARM) {
	*out = *in
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = new(string)
		**out = **in
	}
	if in.CustomerId != nil {
		in, out := &in.CustomerId, &out.CustomerId
		*out = new(string)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(WorkspaceFeatures_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ForceCmkForQuery != nil {
		in, out := &in.ForceCmkForQuery, &out.ForceCmkForQuery
		*out = new(bool)
		**out = **in
	}
	if in.ModifiedDate != nil {
		in, out := &in.ModifiedDate, &out.ModifiedDate
		*out = new(string)
		**out = **in
	}
	if in.PrivateLinkScopedResources != nil {
		in, out := &in.PrivateLinkScopedResources, &out.PrivateLinkScopedResources
		*out = make([]PrivateLinkScopedResource_STATUSARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(WorkspacePropertiesSTATUSProvisioningState)
		**out = **in
	}
	if in.PublicNetworkAccessForIngestion != nil {
		in, out := &in.PublicNetworkAccessForIngestion, &out.PublicNetworkAccessForIngestion
		*out = new(PublicNetworkAccessType_STATUS)
		**out = **in
	}
	if in.PublicNetworkAccessForQuery != nil {
		in, out := &in.PublicNetworkAccessForQuery, &out.PublicNetworkAccessForQuery
		*out = new(PublicNetworkAccessType_STATUS)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(WorkspaceSku_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceCapping != nil {
		in, out := &in.WorkspaceCapping, &out.WorkspaceCapping
		*out = new(WorkspaceCapping_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceProperties_STATUSARM.
func (in *WorkspaceProperties_STATUSARM) DeepCopy() *WorkspaceProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(WorkspaceProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSku) DeepCopyInto(out *WorkspaceSku) {
	*out = *in
	if in.CapacityReservationLevel != nil {
		in, out := &in.CapacityReservationLevel, &out.CapacityReservationLevel
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(WorkspaceSkuName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSku.
func (in *WorkspaceSku) DeepCopy() *WorkspaceSku {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSkuARM) DeepCopyInto(out *WorkspaceSkuARM) {
	*out = *in
	if in.CapacityReservationLevel != nil {
		in, out := &in.CapacityReservationLevel, &out.CapacityReservationLevel
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(WorkspaceSkuName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSkuARM.
func (in *WorkspaceSkuARM) DeepCopy() *WorkspaceSkuARM {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSkuARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSku_STATUS) DeepCopyInto(out *WorkspaceSku_STATUS) {
	*out = *in
	if in.CapacityReservationLevel != nil {
		in, out := &in.CapacityReservationLevel, &out.CapacityReservationLevel
		*out = new(WorkspaceSkuSTATUSCapacityReservationLevel)
		**out = **in
	}
	if in.LastSkuUpdate != nil {
		in, out := &in.LastSkuUpdate, &out.LastSkuUpdate
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(WorkspaceSkuSTATUSName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSku_STATUS.
func (in *WorkspaceSku_STATUS) DeepCopy() *WorkspaceSku_STATUS {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSku_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSku_STATUSARM) DeepCopyInto(out *WorkspaceSku_STATUSARM) {
	*out = *in
	if in.CapacityReservationLevel != nil {
		in, out := &in.CapacityReservationLevel, &out.CapacityReservationLevel
		*out = new(WorkspaceSkuSTATUSCapacityReservationLevel)
		**out = **in
	}
	if in.LastSkuUpdate != nil {
		in, out := &in.LastSkuUpdate, &out.LastSkuUpdate
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(WorkspaceSkuSTATUSName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSku_STATUSARM.
func (in *WorkspaceSku_STATUSARM) DeepCopy() *WorkspaceSku_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSku_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace_STATUS) DeepCopyInto(out *Workspace_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = new(string)
		**out = **in
	}
	if in.CustomerId != nil {
		in, out := &in.CustomerId, &out.CustomerId
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(WorkspaceFeatures_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ForceCmkForQuery != nil {
		in, out := &in.ForceCmkForQuery, &out.ForceCmkForQuery
		*out = new(bool)
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
	if in.ModifiedDate != nil {
		in, out := &in.ModifiedDate, &out.ModifiedDate
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
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(WorkspacePropertiesSTATUSProvisioningState)
		**out = **in
	}
	if in.PublicNetworkAccessForIngestion != nil {
		in, out := &in.PublicNetworkAccessForIngestion, &out.PublicNetworkAccessForIngestion
		*out = new(PublicNetworkAccessType_STATUS)
		**out = **in
	}
	if in.PublicNetworkAccessForQuery != nil {
		in, out := &in.PublicNetworkAccessForQuery, &out.PublicNetworkAccessForQuery
		*out = new(PublicNetworkAccessType_STATUS)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(WorkspaceSku_STATUS)
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
	if in.WorkspaceCapping != nil {
		in, out := &in.WorkspaceCapping, &out.WorkspaceCapping
		*out = new(WorkspaceCapping_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace_STATUS.
func (in *Workspace_STATUS) DeepCopy() *Workspace_STATUS {
	if in == nil {
		return nil
	}
	out := new(Workspace_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace_STATUSARM) DeepCopyInto(out *Workspace_STATUSARM) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
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
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(WorkspaceProperties_STATUSARM)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace_STATUSARM.
func (in *Workspace_STATUSARM) DeepCopy() *Workspace_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(Workspace_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspaces_Spec) DeepCopyInto(out *Workspaces_Spec) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(WorkspaceFeatures)
		(*in).DeepCopyInto(*out)
	}
	if in.ForceCmkForQuery != nil {
		in, out := &in.ForceCmkForQuery, &out.ForceCmkForQuery
		*out = new(bool)
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
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(WorkspacePropertiesProvisioningState)
		**out = **in
	}
	if in.PublicNetworkAccessForIngestion != nil {
		in, out := &in.PublicNetworkAccessForIngestion, &out.PublicNetworkAccessForIngestion
		*out = new(WorkspacePropertiesPublicNetworkAccessForIngestion)
		**out = **in
	}
	if in.PublicNetworkAccessForQuery != nil {
		in, out := &in.PublicNetworkAccessForQuery, &out.PublicNetworkAccessForQuery
		*out = new(WorkspacePropertiesPublicNetworkAccessForQuery)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(WorkspaceSku)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.WorkspaceCapping != nil {
		in, out := &in.WorkspaceCapping, &out.WorkspaceCapping
		*out = new(WorkspaceCapping)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspaces_Spec.
func (in *Workspaces_Spec) DeepCopy() *Workspaces_Spec {
	if in == nil {
		return nil
	}
	out := new(Workspaces_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspaces_SpecARM) DeepCopyInto(out *Workspaces_SpecARM) {
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
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(WorkspacePropertiesARM)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspaces_SpecARM.
func (in *Workspaces_SpecARM) DeepCopy() *Workspaces_SpecARM {
	if in == nil {
		return nil
	}
	out := new(Workspaces_SpecARM)
	in.DeepCopyInto(out)
	return out
}
