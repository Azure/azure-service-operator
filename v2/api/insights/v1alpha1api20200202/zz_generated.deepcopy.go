//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20200202

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInsightsComponentPropertiesARM) DeepCopyInto(out *ApplicationInsightsComponentPropertiesARM) {
	*out = *in
	if in.ApplicationType != nil {
		in, out := &in.ApplicationType, &out.ApplicationType
		*out = new(ApplicationInsightsComponentPropertiesApplicationType)
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
	if in.FlowType != nil {
		in, out := &in.FlowType, &out.FlowType
		*out = new(ApplicationInsightsComponentPropertiesFlowType)
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
		*out = new(ApplicationInsightsComponentPropertiesIngestionMode)
		**out = **in
	}
	if in.PublicNetworkAccessForIngestion != nil {
		in, out := &in.PublicNetworkAccessForIngestion, &out.PublicNetworkAccessForIngestion
		*out = new(ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion)
		**out = **in
	}
	if in.PublicNetworkAccessForQuery != nil {
		in, out := &in.PublicNetworkAccessForQuery, &out.PublicNetworkAccessForQuery
		*out = new(ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery)
		**out = **in
	}
	if in.RequestSource != nil {
		in, out := &in.RequestSource, &out.RequestSource
		*out = new(ApplicationInsightsComponentPropertiesRequestSource)
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
	if in.WorkspaceResourceId != nil {
		in, out := &in.WorkspaceResourceId, &out.WorkspaceResourceId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInsightsComponentPropertiesARM.
func (in *ApplicationInsightsComponentPropertiesARM) DeepCopy() *ApplicationInsightsComponentPropertiesARM {
	if in == nil {
		return nil
	}
	out := new(ApplicationInsightsComponentPropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInsightsComponentProperties_STATUSARM) DeepCopyInto(out *ApplicationInsightsComponentProperties_STATUSARM) {
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
	if in.ApplicationType != nil {
		in, out := &in.ApplicationType, &out.ApplicationType
		*out = new(ApplicationInsightsComponentPropertiesSTATUSApplicationType)
		**out = **in
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
	if in.FlowType != nil {
		in, out := &in.FlowType, &out.FlowType
		*out = new(ApplicationInsightsComponentPropertiesSTATUSFlowType)
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
	if in.ImmediatePurgeDataOn30Days != nil {
		in, out := &in.ImmediatePurgeDataOn30Days, &out.ImmediatePurgeDataOn30Days
		*out = new(bool)
		**out = **in
	}
	if in.IngestionMode != nil {
		in, out := &in.IngestionMode, &out.IngestionMode
		*out = new(ApplicationInsightsComponentPropertiesSTATUSIngestionMode)
		**out = **in
	}
	if in.InstrumentationKey != nil {
		in, out := &in.InstrumentationKey, &out.InstrumentationKey
		*out = new(string)
		**out = **in
	}
	if in.LaMigrationDate != nil {
		in, out := &in.LaMigrationDate, &out.LaMigrationDate
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
		*out = make([]PrivateLinkScopedResource_STATUSARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
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
	if in.RequestSource != nil {
		in, out := &in.RequestSource, &out.RequestSource
		*out = new(ApplicationInsightsComponentPropertiesSTATUSRequestSource)
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
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceResourceId != nil {
		in, out := &in.WorkspaceResourceId, &out.WorkspaceResourceId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInsightsComponentProperties_STATUSARM.
func (in *ApplicationInsightsComponentProperties_STATUSARM) DeepCopy() *ApplicationInsightsComponentProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(ApplicationInsightsComponentProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInsightsComponent_STATUS) DeepCopyInto(out *ApplicationInsightsComponent_STATUS) {
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
	if in.ApplicationType != nil {
		in, out := &in.ApplicationType, &out.ApplicationType
		*out = new(ApplicationInsightsComponentPropertiesSTATUSApplicationType)
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
	if in.FlowType != nil {
		in, out := &in.FlowType, &out.FlowType
		*out = new(ApplicationInsightsComponentPropertiesSTATUSFlowType)
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
		*out = new(ApplicationInsightsComponentPropertiesSTATUSIngestionMode)
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
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
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
	if in.RequestSource != nil {
		in, out := &in.RequestSource, &out.RequestSource
		*out = new(ApplicationInsightsComponentPropertiesSTATUSRequestSource)
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
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInsightsComponent_STATUS.
func (in *ApplicationInsightsComponent_STATUS) DeepCopy() *ApplicationInsightsComponent_STATUS {
	if in == nil {
		return nil
	}
	out := new(ApplicationInsightsComponent_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInsightsComponent_STATUSARM) DeepCopyInto(out *ApplicationInsightsComponent_STATUSARM) {
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ApplicationInsightsComponentProperties_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInsightsComponent_STATUSARM.
func (in *ApplicationInsightsComponent_STATUSARM) DeepCopy() *ApplicationInsightsComponent_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(ApplicationInsightsComponent_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

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
func (in *Components_Spec) DeepCopyInto(out *Components_Spec) {
	*out = *in
	if in.ApplicationType != nil {
		in, out := &in.ApplicationType, &out.ApplicationType
		*out = new(ApplicationInsightsComponentPropertiesApplicationType)
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
	if in.FlowType != nil {
		in, out := &in.FlowType, &out.FlowType
		*out = new(ApplicationInsightsComponentPropertiesFlowType)
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
		*out = new(ApplicationInsightsComponentPropertiesIngestionMode)
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
	if in.PublicNetworkAccessForIngestion != nil {
		in, out := &in.PublicNetworkAccessForIngestion, &out.PublicNetworkAccessForIngestion
		*out = new(ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion)
		**out = **in
	}
	if in.PublicNetworkAccessForQuery != nil {
		in, out := &in.PublicNetworkAccessForQuery, &out.PublicNetworkAccessForQuery
		*out = new(ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery)
		**out = **in
	}
	if in.RequestSource != nil {
		in, out := &in.RequestSource, &out.RequestSource
		*out = new(ApplicationInsightsComponentPropertiesRequestSource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Components_Spec.
func (in *Components_Spec) DeepCopy() *Components_Spec {
	if in == nil {
		return nil
	}
	out := new(Components_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Components_SpecARM) DeepCopyInto(out *Components_SpecARM) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
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
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ApplicationInsightsComponentPropertiesARM)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Components_SpecARM.
func (in *Components_SpecARM) DeepCopy() *Components_SpecARM {
	if in == nil {
		return nil
	}
	out := new(Components_SpecARM)
	in.DeepCopyInto(out)
	return out
}

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
