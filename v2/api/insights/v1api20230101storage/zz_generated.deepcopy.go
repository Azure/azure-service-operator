//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20230101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionGroup) DeepCopyInto(out *ActionGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionGroup.
func (in *ActionGroup) DeepCopy() *ActionGroup {
	if in == nil {
		return nil
	}
	out := new(ActionGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActionGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionGroupList) DeepCopyInto(out *ActionGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActionGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionGroupList.
func (in *ActionGroupList) DeepCopy() *ActionGroupList {
	if in == nil {
		return nil
	}
	out := new(ActionGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActionGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionGroupResource_STATUS) DeepCopyInto(out *ActionGroupResource_STATUS) {
	*out = *in
	if in.ArmRoleReceivers != nil {
		in, out := &in.ArmRoleReceivers, &out.ArmRoleReceivers
		*out = make([]ArmRoleReceiver_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutomationRunbookReceivers != nil {
		in, out := &in.AutomationRunbookReceivers, &out.AutomationRunbookReceivers
		*out = make([]AutomationRunbookReceiver_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AzureAppPushReceivers != nil {
		in, out := &in.AzureAppPushReceivers, &out.AzureAppPushReceivers
		*out = make([]AzureAppPushReceiver_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AzureFunctionReceivers != nil {
		in, out := &in.AzureFunctionReceivers, &out.AzureFunctionReceivers
		*out = make([]AzureFunctionReceiver_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EmailReceivers != nil {
		in, out := &in.EmailReceivers, &out.EmailReceivers
		*out = make([]EmailReceiver_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventHubReceivers != nil {
		in, out := &in.EventHubReceivers, &out.EventHubReceivers
		*out = make([]EventHubReceiver_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GroupShortName != nil {
		in, out := &in.GroupShortName, &out.GroupShortName
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.ItsmReceivers != nil {
		in, out := &in.ItsmReceivers, &out.ItsmReceivers
		*out = make([]ItsmReceiver_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LogicAppReceivers != nil {
		in, out := &in.LogicAppReceivers, &out.LogicAppReceivers
		*out = make([]LogicAppReceiver_STATUS, len(*in))
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
	if in.SmsReceivers != nil {
		in, out := &in.SmsReceivers, &out.SmsReceivers
		*out = make([]SmsReceiver_STATUS, len(*in))
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.VoiceReceivers != nil {
		in, out := &in.VoiceReceivers, &out.VoiceReceivers
		*out = make([]VoiceReceiver_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebhookReceivers != nil {
		in, out := &in.WebhookReceivers, &out.WebhookReceivers
		*out = make([]WebhookReceiver_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionGroupResource_STATUS.
func (in *ActionGroupResource_STATUS) DeepCopy() *ActionGroupResource_STATUS {
	if in == nil {
		return nil
	}
	out := new(ActionGroupResource_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionGroup_Spec) DeepCopyInto(out *ActionGroup_Spec) {
	*out = *in
	if in.ArmRoleReceivers != nil {
		in, out := &in.ArmRoleReceivers, &out.ArmRoleReceivers
		*out = make([]ArmRoleReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutomationRunbookReceivers != nil {
		in, out := &in.AutomationRunbookReceivers, &out.AutomationRunbookReceivers
		*out = make([]AutomationRunbookReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AzureAppPushReceivers != nil {
		in, out := &in.AzureAppPushReceivers, &out.AzureAppPushReceivers
		*out = make([]AzureAppPushReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AzureFunctionReceivers != nil {
		in, out := &in.AzureFunctionReceivers, &out.AzureFunctionReceivers
		*out = make([]AzureFunctionReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EmailReceivers != nil {
		in, out := &in.EmailReceivers, &out.EmailReceivers
		*out = make([]EmailReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventHubReceivers != nil {
		in, out := &in.EventHubReceivers, &out.EventHubReceivers
		*out = make([]EventHubReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GroupShortName != nil {
		in, out := &in.GroupShortName, &out.GroupShortName
		*out = new(string)
		**out = **in
	}
	if in.ItsmReceivers != nil {
		in, out := &in.ItsmReceivers, &out.ItsmReceivers
		*out = make([]ItsmReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LogicAppReceivers != nil {
		in, out := &in.LogicAppReceivers, &out.LogicAppReceivers
		*out = make([]LogicAppReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.SmsReceivers != nil {
		in, out := &in.SmsReceivers, &out.SmsReceivers
		*out = make([]SmsReceiver, len(*in))
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
	if in.VoiceReceivers != nil {
		in, out := &in.VoiceReceivers, &out.VoiceReceivers
		*out = make([]VoiceReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebhookReceivers != nil {
		in, out := &in.WebhookReceivers, &out.WebhookReceivers
		*out = make([]WebhookReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionGroup_Spec.
func (in *ActionGroup_Spec) DeepCopy() *ActionGroup_Spec {
	if in == nil {
		return nil
	}
	out := new(ActionGroup_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmRoleReceiver) DeepCopyInto(out *ArmRoleReceiver) {
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
	if in.RoleId != nil {
		in, out := &in.RoleId, &out.RoleId
		*out = new(string)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmRoleReceiver.
func (in *ArmRoleReceiver) DeepCopy() *ArmRoleReceiver {
	if in == nil {
		return nil
	}
	out := new(ArmRoleReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmRoleReceiver_STATUS) DeepCopyInto(out *ArmRoleReceiver_STATUS) {
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
	if in.RoleId != nil {
		in, out := &in.RoleId, &out.RoleId
		*out = new(string)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmRoleReceiver_STATUS.
func (in *ArmRoleReceiver_STATUS) DeepCopy() *ArmRoleReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(ArmRoleReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutomationRunbookReceiver) DeepCopyInto(out *AutomationRunbookReceiver) {
	*out = *in
	if in.AutomationAccountId != nil {
		in, out := &in.AutomationAccountId, &out.AutomationAccountId
		*out = new(string)
		**out = **in
	}
	if in.IsGlobalRunbook != nil {
		in, out := &in.IsGlobalRunbook, &out.IsGlobalRunbook
		*out = new(bool)
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
	if in.RunbookName != nil {
		in, out := &in.RunbookName, &out.RunbookName
		*out = new(string)
		**out = **in
	}
	if in.ServiceUri != nil {
		in, out := &in.ServiceUri, &out.ServiceUri
		*out = new(string)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
	if in.WebhookResourceReference != nil {
		in, out := &in.WebhookResourceReference, &out.WebhookResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutomationRunbookReceiver.
func (in *AutomationRunbookReceiver) DeepCopy() *AutomationRunbookReceiver {
	if in == nil {
		return nil
	}
	out := new(AutomationRunbookReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutomationRunbookReceiver_STATUS) DeepCopyInto(out *AutomationRunbookReceiver_STATUS) {
	*out = *in
	if in.AutomationAccountId != nil {
		in, out := &in.AutomationAccountId, &out.AutomationAccountId
		*out = new(string)
		**out = **in
	}
	if in.IsGlobalRunbook != nil {
		in, out := &in.IsGlobalRunbook, &out.IsGlobalRunbook
		*out = new(bool)
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
	if in.RunbookName != nil {
		in, out := &in.RunbookName, &out.RunbookName
		*out = new(string)
		**out = **in
	}
	if in.ServiceUri != nil {
		in, out := &in.ServiceUri, &out.ServiceUri
		*out = new(string)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
	if in.WebhookResourceId != nil {
		in, out := &in.WebhookResourceId, &out.WebhookResourceId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutomationRunbookReceiver_STATUS.
func (in *AutomationRunbookReceiver_STATUS) DeepCopy() *AutomationRunbookReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(AutomationRunbookReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAppPushReceiver) DeepCopyInto(out *AzureAppPushReceiver) {
	*out = *in
	if in.EmailAddress != nil {
		in, out := &in.EmailAddress, &out.EmailAddress
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAppPushReceiver.
func (in *AzureAppPushReceiver) DeepCopy() *AzureAppPushReceiver {
	if in == nil {
		return nil
	}
	out := new(AzureAppPushReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAppPushReceiver_STATUS) DeepCopyInto(out *AzureAppPushReceiver_STATUS) {
	*out = *in
	if in.EmailAddress != nil {
		in, out := &in.EmailAddress, &out.EmailAddress
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAppPushReceiver_STATUS.
func (in *AzureAppPushReceiver_STATUS) DeepCopy() *AzureAppPushReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(AzureAppPushReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFunctionReceiver) DeepCopyInto(out *AzureFunctionReceiver) {
	*out = *in
	if in.FunctionAppResourceReference != nil {
		in, out := &in.FunctionAppResourceReference, &out.FunctionAppResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
		*out = new(string)
		**out = **in
	}
	if in.HttpTriggerUrl != nil {
		in, out := &in.HttpTriggerUrl, &out.HttpTriggerUrl
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
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFunctionReceiver.
func (in *AzureFunctionReceiver) DeepCopy() *AzureFunctionReceiver {
	if in == nil {
		return nil
	}
	out := new(AzureFunctionReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFunctionReceiver_STATUS) DeepCopyInto(out *AzureFunctionReceiver_STATUS) {
	*out = *in
	if in.FunctionAppResourceId != nil {
		in, out := &in.FunctionAppResourceId, &out.FunctionAppResourceId
		*out = new(string)
		**out = **in
	}
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
		*out = new(string)
		**out = **in
	}
	if in.HttpTriggerUrl != nil {
		in, out := &in.HttpTriggerUrl, &out.HttpTriggerUrl
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
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFunctionReceiver_STATUS.
func (in *AzureFunctionReceiver_STATUS) DeepCopy() *AzureFunctionReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(AzureFunctionReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailReceiver) DeepCopyInto(out *EmailReceiver) {
	*out = *in
	if in.EmailAddress != nil {
		in, out := &in.EmailAddress, &out.EmailAddress
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
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailReceiver.
func (in *EmailReceiver) DeepCopy() *EmailReceiver {
	if in == nil {
		return nil
	}
	out := new(EmailReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailReceiver_STATUS) DeepCopyInto(out *EmailReceiver_STATUS) {
	*out = *in
	if in.EmailAddress != nil {
		in, out := &in.EmailAddress, &out.EmailAddress
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
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailReceiver_STATUS.
func (in *EmailReceiver_STATUS) DeepCopy() *EmailReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(EmailReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubReceiver) DeepCopyInto(out *EventHubReceiver) {
	*out = *in
	if in.EventHubName != nil {
		in, out := &in.EventHubName, &out.EventHubName
		*out = new(string)
		**out = **in
	}
	if in.EventHubNameSpace != nil {
		in, out := &in.EventHubNameSpace, &out.EventHubNameSpace
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
	if in.SubscriptionId != nil {
		in, out := &in.SubscriptionId, &out.SubscriptionId
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubReceiver.
func (in *EventHubReceiver) DeepCopy() *EventHubReceiver {
	if in == nil {
		return nil
	}
	out := new(EventHubReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubReceiver_STATUS) DeepCopyInto(out *EventHubReceiver_STATUS) {
	*out = *in
	if in.EventHubName != nil {
		in, out := &in.EventHubName, &out.EventHubName
		*out = new(string)
		**out = **in
	}
	if in.EventHubNameSpace != nil {
		in, out := &in.EventHubNameSpace, &out.EventHubNameSpace
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
	if in.SubscriptionId != nil {
		in, out := &in.SubscriptionId, &out.SubscriptionId
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubReceiver_STATUS.
func (in *EventHubReceiver_STATUS) DeepCopy() *EventHubReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(EventHubReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ItsmReceiver) DeepCopyInto(out *ItsmReceiver) {
	*out = *in
	if in.ConnectionId != nil {
		in, out := &in.ConnectionId, &out.ConnectionId
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
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TicketConfiguration != nil {
		in, out := &in.TicketConfiguration, &out.TicketConfiguration
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceId != nil {
		in, out := &in.WorkspaceId, &out.WorkspaceId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ItsmReceiver.
func (in *ItsmReceiver) DeepCopy() *ItsmReceiver {
	if in == nil {
		return nil
	}
	out := new(ItsmReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ItsmReceiver_STATUS) DeepCopyInto(out *ItsmReceiver_STATUS) {
	*out = *in
	if in.ConnectionId != nil {
		in, out := &in.ConnectionId, &out.ConnectionId
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
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TicketConfiguration != nil {
		in, out := &in.TicketConfiguration, &out.TicketConfiguration
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceId != nil {
		in, out := &in.WorkspaceId, &out.WorkspaceId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ItsmReceiver_STATUS.
func (in *ItsmReceiver_STATUS) DeepCopy() *ItsmReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(ItsmReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicAppReceiver) DeepCopyInto(out *LogicAppReceiver) {
	*out = *in
	if in.CallbackUrl != nil {
		in, out := &in.CallbackUrl, &out.CallbackUrl
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
	if in.ResourceReference != nil {
		in, out := &in.ResourceReference, &out.ResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicAppReceiver.
func (in *LogicAppReceiver) DeepCopy() *LogicAppReceiver {
	if in == nil {
		return nil
	}
	out := new(LogicAppReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicAppReceiver_STATUS) DeepCopyInto(out *LogicAppReceiver_STATUS) {
	*out = *in
	if in.CallbackUrl != nil {
		in, out := &in.CallbackUrl, &out.CallbackUrl
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
	if in.ResourceId != nil {
		in, out := &in.ResourceId, &out.ResourceId
		*out = new(string)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicAppReceiver_STATUS.
func (in *LogicAppReceiver_STATUS) DeepCopy() *LogicAppReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(LogicAppReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmsReceiver) DeepCopyInto(out *SmsReceiver) {
	*out = *in
	if in.CountryCode != nil {
		in, out := &in.CountryCode, &out.CountryCode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PhoneNumber != nil {
		in, out := &in.PhoneNumber, &out.PhoneNumber
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmsReceiver.
func (in *SmsReceiver) DeepCopy() *SmsReceiver {
	if in == nil {
		return nil
	}
	out := new(SmsReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmsReceiver_STATUS) DeepCopyInto(out *SmsReceiver_STATUS) {
	*out = *in
	if in.CountryCode != nil {
		in, out := &in.CountryCode, &out.CountryCode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PhoneNumber != nil {
		in, out := &in.PhoneNumber, &out.PhoneNumber
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
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmsReceiver_STATUS.
func (in *SmsReceiver_STATUS) DeepCopy() *SmsReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(SmsReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceReceiver) DeepCopyInto(out *VoiceReceiver) {
	*out = *in
	if in.CountryCode != nil {
		in, out := &in.CountryCode, &out.CountryCode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PhoneNumber != nil {
		in, out := &in.PhoneNumber, &out.PhoneNumber
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceReceiver.
func (in *VoiceReceiver) DeepCopy() *VoiceReceiver {
	if in == nil {
		return nil
	}
	out := new(VoiceReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VoiceReceiver_STATUS) DeepCopyInto(out *VoiceReceiver_STATUS) {
	*out = *in
	if in.CountryCode != nil {
		in, out := &in.CountryCode, &out.CountryCode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PhoneNumber != nil {
		in, out := &in.PhoneNumber, &out.PhoneNumber
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VoiceReceiver_STATUS.
func (in *VoiceReceiver_STATUS) DeepCopy() *VoiceReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(VoiceReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookReceiver) DeepCopyInto(out *WebhookReceiver) {
	*out = *in
	if in.IdentifierUri != nil {
		in, out := &in.IdentifierUri, &out.IdentifierUri
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObjectId != nil {
		in, out := &in.ObjectId, &out.ObjectId
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
	if in.ServiceUri != nil {
		in, out := &in.ServiceUri, &out.ServiceUri
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.UseAadAuth != nil {
		in, out := &in.UseAadAuth, &out.UseAadAuth
		*out = new(bool)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookReceiver.
func (in *WebhookReceiver) DeepCopy() *WebhookReceiver {
	if in == nil {
		return nil
	}
	out := new(WebhookReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookReceiver_STATUS) DeepCopyInto(out *WebhookReceiver_STATUS) {
	*out = *in
	if in.IdentifierUri != nil {
		in, out := &in.IdentifierUri, &out.IdentifierUri
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObjectId != nil {
		in, out := &in.ObjectId, &out.ObjectId
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
	if in.ServiceUri != nil {
		in, out := &in.ServiceUri, &out.ServiceUri
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.UseAadAuth != nil {
		in, out := &in.UseAadAuth, &out.UseAadAuth
		*out = new(bool)
		**out = **in
	}
	if in.UseCommonAlertSchema != nil {
		in, out := &in.UseCommonAlertSchema, &out.UseCommonAlertSchema
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookReceiver_STATUS.
func (in *WebhookReceiver_STATUS) DeepCopy() *WebhookReceiver_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebhookReceiver_STATUS)
	in.DeepCopyInto(out)
	return out
}
