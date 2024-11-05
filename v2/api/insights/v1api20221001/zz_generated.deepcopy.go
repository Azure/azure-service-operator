//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20221001

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscaleNotification) DeepCopyInto(out *AutoscaleNotification) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(EmailNotification)
		(*in).DeepCopyInto(*out)
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(AutoscaleNotification_Operation)
		**out = **in
	}
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]WebhookNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscaleNotification.
func (in *AutoscaleNotification) DeepCopy() *AutoscaleNotification {
	if in == nil {
		return nil
	}
	out := new(AutoscaleNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscaleNotification_STATUS) DeepCopyInto(out *AutoscaleNotification_STATUS) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(EmailNotification_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(AutoscaleNotification_Operation_STATUS)
		**out = **in
	}
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]WebhookNotification_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscaleNotification_STATUS.
func (in *AutoscaleNotification_STATUS) DeepCopy() *AutoscaleNotification_STATUS {
	if in == nil {
		return nil
	}
	out := new(AutoscaleNotification_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscaleProfile) DeepCopyInto(out *AutoscaleProfile) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(ScaleCapacity)
		(*in).DeepCopyInto(*out)
	}
	if in.FixedDate != nil {
		in, out := &in.FixedDate, &out.FixedDate
		*out = new(TimeWindow)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Recurrence != nil {
		in, out := &in.Recurrence, &out.Recurrence
		*out = new(Recurrence)
		(*in).DeepCopyInto(*out)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]ScaleRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscaleProfile.
func (in *AutoscaleProfile) DeepCopy() *AutoscaleProfile {
	if in == nil {
		return nil
	}
	out := new(AutoscaleProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscaleProfile_STATUS) DeepCopyInto(out *AutoscaleProfile_STATUS) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(ScaleCapacity_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.FixedDate != nil {
		in, out := &in.FixedDate, &out.FixedDate
		*out = new(TimeWindow_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Recurrence != nil {
		in, out := &in.Recurrence, &out.Recurrence
		*out = new(Recurrence_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]ScaleRule_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscaleProfile_STATUS.
func (in *AutoscaleProfile_STATUS) DeepCopy() *AutoscaleProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(AutoscaleProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscaleSetting) DeepCopyInto(out *AutoscaleSetting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscaleSetting.
func (in *AutoscaleSetting) DeepCopy() *AutoscaleSetting {
	if in == nil {
		return nil
	}
	out := new(AutoscaleSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoscaleSetting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscaleSettingList) DeepCopyInto(out *AutoscaleSettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutoscaleSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscaleSettingList.
func (in *AutoscaleSettingList) DeepCopy() *AutoscaleSettingList {
	if in == nil {
		return nil
	}
	out := new(AutoscaleSettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoscaleSettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscaleSetting_Spec) DeepCopyInto(out *AutoscaleSetting_Spec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
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
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]AutoscaleNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.PredictiveAutoscalePolicy != nil {
		in, out := &in.PredictiveAutoscalePolicy, &out.PredictiveAutoscalePolicy
		*out = new(PredictiveAutoscalePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]AutoscaleProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TargetResourceLocation != nil {
		in, out := &in.TargetResourceLocation, &out.TargetResourceLocation
		*out = new(string)
		**out = **in
	}
	if in.TargetResourceUriReference != nil {
		in, out := &in.TargetResourceUriReference, &out.TargetResourceUriReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscaleSetting_Spec.
func (in *AutoscaleSetting_Spec) DeepCopy() *AutoscaleSetting_Spec {
	if in == nil {
		return nil
	}
	out := new(AutoscaleSetting_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Autoscalesetting_STATUS) DeepCopyInto(out *Autoscalesetting_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]AutoscaleNotification_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PredictiveAutoscalePolicy != nil {
		in, out := &in.PredictiveAutoscalePolicy, &out.PredictiveAutoscalePolicy
		*out = new(PredictiveAutoscalePolicy_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]AutoscaleProfile_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PropertiesName != nil {
		in, out := &in.PropertiesName, &out.PropertiesName
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
	if in.TargetResourceLocation != nil {
		in, out := &in.TargetResourceLocation, &out.TargetResourceLocation
		*out = new(string)
		**out = **in
	}
	if in.TargetResourceUri != nil {
		in, out := &in.TargetResourceUri, &out.TargetResourceUri
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Autoscalesetting_STATUS.
func (in *Autoscalesetting_STATUS) DeepCopy() *Autoscalesetting_STATUS {
	if in == nil {
		return nil
	}
	out := new(Autoscalesetting_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailNotification) DeepCopyInto(out *EmailNotification) {
	*out = *in
	if in.CustomEmails != nil {
		in, out := &in.CustomEmails, &out.CustomEmails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SendToSubscriptionAdministrator != nil {
		in, out := &in.SendToSubscriptionAdministrator, &out.SendToSubscriptionAdministrator
		*out = new(bool)
		**out = **in
	}
	if in.SendToSubscriptionCoAdministrators != nil {
		in, out := &in.SendToSubscriptionCoAdministrators, &out.SendToSubscriptionCoAdministrators
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailNotification.
func (in *EmailNotification) DeepCopy() *EmailNotification {
	if in == nil {
		return nil
	}
	out := new(EmailNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailNotification_STATUS) DeepCopyInto(out *EmailNotification_STATUS) {
	*out = *in
	if in.CustomEmails != nil {
		in, out := &in.CustomEmails, &out.CustomEmails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SendToSubscriptionAdministrator != nil {
		in, out := &in.SendToSubscriptionAdministrator, &out.SendToSubscriptionAdministrator
		*out = new(bool)
		**out = **in
	}
	if in.SendToSubscriptionCoAdministrators != nil {
		in, out := &in.SendToSubscriptionCoAdministrators, &out.SendToSubscriptionCoAdministrators
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailNotification_STATUS.
func (in *EmailNotification_STATUS) DeepCopy() *EmailNotification_STATUS {
	if in == nil {
		return nil
	}
	out := new(EmailNotification_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTrigger) DeepCopyInto(out *MetricTrigger) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]ScaleRuleMetricDimension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DividePerInstance != nil {
		in, out := &in.DividePerInstance, &out.DividePerInstance
		*out = new(bool)
		**out = **in
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.MetricNamespace != nil {
		in, out := &in.MetricNamespace, &out.MetricNamespace
		*out = new(string)
		**out = **in
	}
	if in.MetricResourceLocation != nil {
		in, out := &in.MetricResourceLocation, &out.MetricResourceLocation
		*out = new(string)
		**out = **in
	}
	if in.MetricResourceUriReference != nil {
		in, out := &in.MetricResourceUriReference, &out.MetricResourceUriReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(MetricTrigger_Operator)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(MetricTrigger_Statistic)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.TimeAggregation != nil {
		in, out := &in.TimeAggregation, &out.TimeAggregation
		*out = new(MetricTrigger_TimeAggregation)
		**out = **in
	}
	if in.TimeGrain != nil {
		in, out := &in.TimeGrain, &out.TimeGrain
		*out = new(string)
		**out = **in
	}
	if in.TimeWindow != nil {
		in, out := &in.TimeWindow, &out.TimeWindow
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTrigger.
func (in *MetricTrigger) DeepCopy() *MetricTrigger {
	if in == nil {
		return nil
	}
	out := new(MetricTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTrigger_STATUS) DeepCopyInto(out *MetricTrigger_STATUS) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]ScaleRuleMetricDimension_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DividePerInstance != nil {
		in, out := &in.DividePerInstance, &out.DividePerInstance
		*out = new(bool)
		**out = **in
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.MetricNamespace != nil {
		in, out := &in.MetricNamespace, &out.MetricNamespace
		*out = new(string)
		**out = **in
	}
	if in.MetricResourceLocation != nil {
		in, out := &in.MetricResourceLocation, &out.MetricResourceLocation
		*out = new(string)
		**out = **in
	}
	if in.MetricResourceUri != nil {
		in, out := &in.MetricResourceUri, &out.MetricResourceUri
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(MetricTrigger_Operator_STATUS)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(MetricTrigger_Statistic_STATUS)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.TimeAggregation != nil {
		in, out := &in.TimeAggregation, &out.TimeAggregation
		*out = new(MetricTrigger_TimeAggregation_STATUS)
		**out = **in
	}
	if in.TimeGrain != nil {
		in, out := &in.TimeGrain, &out.TimeGrain
		*out = new(string)
		**out = **in
	}
	if in.TimeWindow != nil {
		in, out := &in.TimeWindow, &out.TimeWindow
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTrigger_STATUS.
func (in *MetricTrigger_STATUS) DeepCopy() *MetricTrigger_STATUS {
	if in == nil {
		return nil
	}
	out := new(MetricTrigger_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictiveAutoscalePolicy) DeepCopyInto(out *PredictiveAutoscalePolicy) {
	*out = *in
	if in.ScaleLookAheadTime != nil {
		in, out := &in.ScaleLookAheadTime, &out.ScaleLookAheadTime
		*out = new(string)
		**out = **in
	}
	if in.ScaleMode != nil {
		in, out := &in.ScaleMode, &out.ScaleMode
		*out = new(PredictiveAutoscalePolicy_ScaleMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictiveAutoscalePolicy.
func (in *PredictiveAutoscalePolicy) DeepCopy() *PredictiveAutoscalePolicy {
	if in == nil {
		return nil
	}
	out := new(PredictiveAutoscalePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictiveAutoscalePolicy_STATUS) DeepCopyInto(out *PredictiveAutoscalePolicy_STATUS) {
	*out = *in
	if in.ScaleLookAheadTime != nil {
		in, out := &in.ScaleLookAheadTime, &out.ScaleLookAheadTime
		*out = new(string)
		**out = **in
	}
	if in.ScaleMode != nil {
		in, out := &in.ScaleMode, &out.ScaleMode
		*out = new(PredictiveAutoscalePolicy_ScaleMode_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictiveAutoscalePolicy_STATUS.
func (in *PredictiveAutoscalePolicy_STATUS) DeepCopy() *PredictiveAutoscalePolicy_STATUS {
	if in == nil {
		return nil
	}
	out := new(PredictiveAutoscalePolicy_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Recurrence) DeepCopyInto(out *Recurrence) {
	*out = *in
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(Recurrence_Frequency)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(RecurrentSchedule)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Recurrence.
func (in *Recurrence) DeepCopy() *Recurrence {
	if in == nil {
		return nil
	}
	out := new(Recurrence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Recurrence_STATUS) DeepCopyInto(out *Recurrence_STATUS) {
	*out = *in
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(Recurrence_Frequency_STATUS)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(RecurrentSchedule_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Recurrence_STATUS.
func (in *Recurrence_STATUS) DeepCopy() *Recurrence_STATUS {
	if in == nil {
		return nil
	}
	out := new(Recurrence_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecurrentSchedule) DeepCopyInto(out *RecurrentSchedule) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecurrentSchedule.
func (in *RecurrentSchedule) DeepCopy() *RecurrentSchedule {
	if in == nil {
		return nil
	}
	out := new(RecurrentSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecurrentSchedule_STATUS) DeepCopyInto(out *RecurrentSchedule_STATUS) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecurrentSchedule_STATUS.
func (in *RecurrentSchedule_STATUS) DeepCopy() *RecurrentSchedule_STATUS {
	if in == nil {
		return nil
	}
	out := new(RecurrentSchedule_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleAction) DeepCopyInto(out *ScaleAction) {
	*out = *in
	if in.Cooldown != nil {
		in, out := &in.Cooldown, &out.Cooldown
		*out = new(string)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(ScaleAction_Direction)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ScaleAction_Type)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleAction.
func (in *ScaleAction) DeepCopy() *ScaleAction {
	if in == nil {
		return nil
	}
	out := new(ScaleAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleAction_STATUS) DeepCopyInto(out *ScaleAction_STATUS) {
	*out = *in
	if in.Cooldown != nil {
		in, out := &in.Cooldown, &out.Cooldown
		*out = new(string)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(ScaleAction_Direction_STATUS)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ScaleAction_Type_STATUS)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleAction_STATUS.
func (in *ScaleAction_STATUS) DeepCopy() *ScaleAction_STATUS {
	if in == nil {
		return nil
	}
	out := new(ScaleAction_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleCapacity) DeepCopyInto(out *ScaleCapacity) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(string)
		**out = **in
	}
	if in.Maximum != nil {
		in, out := &in.Maximum, &out.Maximum
		*out = new(string)
		**out = **in
	}
	if in.Minimum != nil {
		in, out := &in.Minimum, &out.Minimum
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleCapacity.
func (in *ScaleCapacity) DeepCopy() *ScaleCapacity {
	if in == nil {
		return nil
	}
	out := new(ScaleCapacity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleCapacity_STATUS) DeepCopyInto(out *ScaleCapacity_STATUS) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(string)
		**out = **in
	}
	if in.Maximum != nil {
		in, out := &in.Maximum, &out.Maximum
		*out = new(string)
		**out = **in
	}
	if in.Minimum != nil {
		in, out := &in.Minimum, &out.Minimum
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleCapacity_STATUS.
func (in *ScaleCapacity_STATUS) DeepCopy() *ScaleCapacity_STATUS {
	if in == nil {
		return nil
	}
	out := new(ScaleCapacity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleRule) DeepCopyInto(out *ScaleRule) {
	*out = *in
	if in.MetricTrigger != nil {
		in, out := &in.MetricTrigger, &out.MetricTrigger
		*out = new(MetricTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaleAction != nil {
		in, out := &in.ScaleAction, &out.ScaleAction
		*out = new(ScaleAction)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleRule.
func (in *ScaleRule) DeepCopy() *ScaleRule {
	if in == nil {
		return nil
	}
	out := new(ScaleRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleRuleMetricDimension) DeepCopyInto(out *ScaleRuleMetricDimension) {
	*out = *in
	if in.DimensionName != nil {
		in, out := &in.DimensionName, &out.DimensionName
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(ScaleRuleMetricDimension_Operator)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleRuleMetricDimension.
func (in *ScaleRuleMetricDimension) DeepCopy() *ScaleRuleMetricDimension {
	if in == nil {
		return nil
	}
	out := new(ScaleRuleMetricDimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleRuleMetricDimension_STATUS) DeepCopyInto(out *ScaleRuleMetricDimension_STATUS) {
	*out = *in
	if in.DimensionName != nil {
		in, out := &in.DimensionName, &out.DimensionName
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(ScaleRuleMetricDimension_Operator_STATUS)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleRuleMetricDimension_STATUS.
func (in *ScaleRuleMetricDimension_STATUS) DeepCopy() *ScaleRuleMetricDimension_STATUS {
	if in == nil {
		return nil
	}
	out := new(ScaleRuleMetricDimension_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleRule_STATUS) DeepCopyInto(out *ScaleRule_STATUS) {
	*out = *in
	if in.MetricTrigger != nil {
		in, out := &in.MetricTrigger, &out.MetricTrigger
		*out = new(MetricTrigger_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaleAction != nil {
		in, out := &in.ScaleAction, &out.ScaleAction
		*out = new(ScaleAction_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleRule_STATUS.
func (in *ScaleRule_STATUS) DeepCopy() *ScaleRule_STATUS {
	if in == nil {
		return nil
	}
	out := new(ScaleRule_STATUS)
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
		*out = new(SystemData_CreatedByType_STATUS)
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
		*out = new(SystemData_LastModifiedByType_STATUS)
		**out = **in
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
func (in *TimeWindow) DeepCopyInto(out *TimeWindow) {
	*out = *in
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = new(string)
		**out = **in
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeWindow.
func (in *TimeWindow) DeepCopy() *TimeWindow {
	if in == nil {
		return nil
	}
	out := new(TimeWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeWindow_STATUS) DeepCopyInto(out *TimeWindow_STATUS) {
	*out = *in
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = new(string)
		**out = **in
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeWindow_STATUS.
func (in *TimeWindow_STATUS) DeepCopy() *TimeWindow_STATUS {
	if in == nil {
		return nil
	}
	out := new(TimeWindow_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookNotification) DeepCopyInto(out *WebhookNotification) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceUri != nil {
		in, out := &in.ServiceUri, &out.ServiceUri
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookNotification.
func (in *WebhookNotification) DeepCopy() *WebhookNotification {
	if in == nil {
		return nil
	}
	out := new(WebhookNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookNotification_STATUS) DeepCopyInto(out *WebhookNotification_STATUS) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceUri != nil {
		in, out := &in.ServiceUri, &out.ServiceUri
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookNotification_STATUS.
func (in *WebhookNotification_STATUS) DeepCopy() *WebhookNotification_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebhookNotification_STATUS)
	in.DeepCopyInto(out)
	return out
}
