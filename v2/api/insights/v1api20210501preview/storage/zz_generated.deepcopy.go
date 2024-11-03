//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiagnosticSetting) DeepCopyInto(out *DiagnosticSetting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiagnosticSetting.
func (in *DiagnosticSetting) DeepCopy() *DiagnosticSetting {
	if in == nil {
		return nil
	}
	out := new(DiagnosticSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiagnosticSetting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiagnosticSettingList) DeepCopyInto(out *DiagnosticSettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DiagnosticSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiagnosticSettingList.
func (in *DiagnosticSettingList) DeepCopy() *DiagnosticSettingList {
	if in == nil {
		return nil
	}
	out := new(DiagnosticSettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiagnosticSettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiagnosticSetting_STATUS) DeepCopyInto(out *DiagnosticSetting_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventHubAuthorizationRuleId != nil {
		in, out := &in.EventHubAuthorizationRuleId, &out.EventHubAuthorizationRuleId
		*out = new(string)
		**out = **in
	}
	if in.EventHubName != nil {
		in, out := &in.EventHubName, &out.EventHubName
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.LogAnalyticsDestinationType != nil {
		in, out := &in.LogAnalyticsDestinationType, &out.LogAnalyticsDestinationType
		*out = new(string)
		**out = **in
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		*out = make([]LogSettings_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MarketplacePartnerId != nil {
		in, out := &in.MarketplacePartnerId, &out.MarketplacePartnerId
		*out = new(string)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]MetricSettings_STATUS, len(*in))
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
	if in.ServiceBusRuleId != nil {
		in, out := &in.ServiceBusRuleId, &out.ServiceBusRuleId
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountId != nil {
		in, out := &in.StorageAccountId, &out.StorageAccountId
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
	if in.WorkspaceId != nil {
		in, out := &in.WorkspaceId, &out.WorkspaceId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiagnosticSetting_STATUS.
func (in *DiagnosticSetting_STATUS) DeepCopy() *DiagnosticSetting_STATUS {
	if in == nil {
		return nil
	}
	out := new(DiagnosticSetting_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiagnosticSetting_Spec) DeepCopyInto(out *DiagnosticSetting_Spec) {
	*out = *in
	if in.EventHubAuthorizationRuleReference != nil {
		in, out := &in.EventHubAuthorizationRuleReference, &out.EventHubAuthorizationRuleReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.EventHubName != nil {
		in, out := &in.EventHubName, &out.EventHubName
		*out = new(string)
		**out = **in
	}
	if in.LogAnalyticsDestinationType != nil {
		in, out := &in.LogAnalyticsDestinationType, &out.LogAnalyticsDestinationType
		*out = new(string)
		**out = **in
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		*out = make([]LogSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MarketplacePartnerReference != nil {
		in, out := &in.MarketplacePartnerReference, &out.MarketplacePartnerReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]MetricSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.ArbitraryOwnerReference)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceBusRuleId != nil {
		in, out := &in.ServiceBusRuleId, &out.ServiceBusRuleId
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountReference != nil {
		in, out := &in.StorageAccountReference, &out.StorageAccountReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.WorkspaceReference != nil {
		in, out := &in.WorkspaceReference, &out.WorkspaceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiagnosticSetting_Spec.
func (in *DiagnosticSetting_Spec) DeepCopy() *DiagnosticSetting_Spec {
	if in == nil {
		return nil
	}
	out := new(DiagnosticSetting_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSettings) DeepCopyInto(out *LogSettings) {
	*out = *in
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = new(string)
		**out = **in
	}
	if in.CategoryGroup != nil {
		in, out := &in.CategoryGroup, &out.CategoryGroup
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = new(RetentionPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSettings.
func (in *LogSettings) DeepCopy() *LogSettings {
	if in == nil {
		return nil
	}
	out := new(LogSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSettings_STATUS) DeepCopyInto(out *LogSettings_STATUS) {
	*out = *in
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = new(string)
		**out = **in
	}
	if in.CategoryGroup != nil {
		in, out := &in.CategoryGroup, &out.CategoryGroup
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = new(RetentionPolicy_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSettings_STATUS.
func (in *LogSettings_STATUS) DeepCopy() *LogSettings_STATUS {
	if in == nil {
		return nil
	}
	out := new(LogSettings_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSettings) DeepCopyInto(out *MetricSettings) {
	*out = *in
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = new(RetentionPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeGrain != nil {
		in, out := &in.TimeGrain, &out.TimeGrain
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSettings.
func (in *MetricSettings) DeepCopy() *MetricSettings {
	if in == nil {
		return nil
	}
	out := new(MetricSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSettings_STATUS) DeepCopyInto(out *MetricSettings_STATUS) {
	*out = *in
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = new(RetentionPolicy_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeGrain != nil {
		in, out := &in.TimeGrain, &out.TimeGrain
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSettings_STATUS.
func (in *MetricSettings_STATUS) DeepCopy() *MetricSettings_STATUS {
	if in == nil {
		return nil
	}
	out := new(MetricSettings_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicy) DeepCopyInto(out *RetentionPolicy) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicy.
func (in *RetentionPolicy) DeepCopy() *RetentionPolicy {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicy_STATUS) DeepCopyInto(out *RetentionPolicy_STATUS) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicy_STATUS.
func (in *RetentionPolicy_STATUS) DeepCopy() *RetentionPolicy_STATUS {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicy_STATUS)
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
