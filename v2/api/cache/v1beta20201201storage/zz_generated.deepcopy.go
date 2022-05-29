//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta20201201storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis) DeepCopyInto(out *Redis) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis.
func (in *Redis) DeepCopy() *Redis {
	if in == nil {
		return nil
	}
	out := new(Redis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Redis) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCreateProperties_RedisConfiguration_STATUS) DeepCopyInto(out *RedisCreateProperties_RedisConfiguration_STATUS) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AofStorageConnectionString0 != nil {
		in, out := &in.AofStorageConnectionString0, &out.AofStorageConnectionString0
		*out = new(string)
		**out = **in
	}
	if in.AofStorageConnectionString1 != nil {
		in, out := &in.AofStorageConnectionString1, &out.AofStorageConnectionString1
		*out = new(string)
		**out = **in
	}
	if in.Maxclients != nil {
		in, out := &in.Maxclients, &out.Maxclients
		*out = new(string)
		**out = **in
	}
	if in.MaxfragmentationmemoryReserved != nil {
		in, out := &in.MaxfragmentationmemoryReserved, &out.MaxfragmentationmemoryReserved
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryDelta != nil {
		in, out := &in.MaxmemoryDelta, &out.MaxmemoryDelta
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryPolicy != nil {
		in, out := &in.MaxmemoryPolicy, &out.MaxmemoryPolicy
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryReserved != nil {
		in, out := &in.MaxmemoryReserved, &out.MaxmemoryReserved
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
	if in.RdbBackupEnabled != nil {
		in, out := &in.RdbBackupEnabled, &out.RdbBackupEnabled
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupFrequency != nil {
		in, out := &in.RdbBackupFrequency, &out.RdbBackupFrequency
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupMaxSnapshotCount != nil {
		in, out := &in.RdbBackupMaxSnapshotCount, &out.RdbBackupMaxSnapshotCount
		*out = new(string)
		**out = **in
	}
	if in.RdbStorageConnectionString != nil {
		in, out := &in.RdbStorageConnectionString, &out.RdbStorageConnectionString
		*out = new(string)
		**out = **in
	}
	if in.ZonalConfiguration != nil {
		in, out := &in.ZonalConfiguration, &out.ZonalConfiguration
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCreateProperties_RedisConfiguration_STATUS.
func (in *RedisCreateProperties_RedisConfiguration_STATUS) DeepCopy() *RedisCreateProperties_RedisConfiguration_STATUS {
	if in == nil {
		return nil
	}
	out := new(RedisCreateProperties_RedisConfiguration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRule) DeepCopyInto(out *RedisFirewallRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRule.
func (in *RedisFirewallRule) DeepCopy() *RedisFirewallRule {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisFirewallRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRuleList) DeepCopyInto(out *RedisFirewallRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisFirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRuleList.
func (in *RedisFirewallRuleList) DeepCopy() *RedisFirewallRuleList {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisFirewallRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRule_STATUS) DeepCopyInto(out *RedisFirewallRule_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EndIP != nil {
		in, out := &in.EndIP, &out.EndIP
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
	if in.StartIP != nil {
		in, out := &in.StartIP, &out.StartIP
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRule_STATUS.
func (in *RedisFirewallRule_STATUS) DeepCopy() *RedisFirewallRule_STATUS {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRule_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRule_Spec) DeepCopyInto(out *RedisFirewallRule_Spec) {
	*out = *in
	if in.EndIP != nil {
		in, out := &in.EndIP, &out.EndIP
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
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
	if in.StartIP != nil {
		in, out := &in.StartIP, &out.StartIP
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRule_Spec.
func (in *RedisFirewallRule_Spec) DeepCopy() *RedisFirewallRule_Spec {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRule_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServer) DeepCopyInto(out *RedisLinkedServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServer.
func (in *RedisLinkedServer) DeepCopy() *RedisLinkedServer {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisLinkedServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServerList) DeepCopyInto(out *RedisLinkedServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisLinkedServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServerList.
func (in *RedisLinkedServerList) DeepCopy() *RedisLinkedServerList {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisLinkedServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServer_STATUS) DeepCopyInto(out *RedisLinkedServer_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.LinkedRedisCacheId != nil {
		in, out := &in.LinkedRedisCacheId, &out.LinkedRedisCacheId
		*out = new(string)
		**out = **in
	}
	if in.LinkedRedisCacheLocation != nil {
		in, out := &in.LinkedRedisCacheLocation, &out.LinkedRedisCacheLocation
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
	if in.ServerRole != nil {
		in, out := &in.ServerRole, &out.ServerRole
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServer_STATUS.
func (in *RedisLinkedServer_STATUS) DeepCopy() *RedisLinkedServer_STATUS {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServer_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServer_Spec) DeepCopyInto(out *RedisLinkedServer_Spec) {
	*out = *in
	if in.LinkedRedisCacheLocation != nil {
		in, out := &in.LinkedRedisCacheLocation, &out.LinkedRedisCacheLocation
		*out = new(string)
		**out = **in
	}
	if in.LinkedRedisCacheReference != nil {
		in, out := &in.LinkedRedisCacheReference, &out.LinkedRedisCacheReference
		*out = new(genruntime.ResourceReference)
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
	if in.ServerRole != nil {
		in, out := &in.ServerRole, &out.ServerRole
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServer_Spec.
func (in *RedisLinkedServer_Spec) DeepCopy() *RedisLinkedServer_Spec {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServer_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisList) DeepCopyInto(out *RedisList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Redis, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisList.
func (in *RedisList) DeepCopy() *RedisList {
	if in == nil {
		return nil
	}
	out := new(RedisList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisOperatorSecrets) DeepCopyInto(out *RedisOperatorSecrets) {
	*out = *in
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.PrimaryKey != nil {
		in, out := &in.PrimaryKey, &out.PrimaryKey
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SSLPort != nil {
		in, out := &in.SSLPort, &out.SSLPort
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.SecondaryKey != nil {
		in, out := &in.SecondaryKey, &out.SecondaryKey
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisOperatorSecrets.
func (in *RedisOperatorSecrets) DeepCopy() *RedisOperatorSecrets {
	if in == nil {
		return nil
	}
	out := new(RedisOperatorSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisOperatorSpec) DeepCopyInto(out *RedisOperatorSpec) {
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
		*out = new(RedisOperatorSecrets)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisOperatorSpec.
func (in *RedisOperatorSpec) DeepCopy() *RedisOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(RedisOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisPatchSchedule) DeepCopyInto(out *RedisPatchSchedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisPatchSchedule.
func (in *RedisPatchSchedule) DeepCopy() *RedisPatchSchedule {
	if in == nil {
		return nil
	}
	out := new(RedisPatchSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisPatchSchedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisPatchScheduleList) DeepCopyInto(out *RedisPatchScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisPatchSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisPatchScheduleList.
func (in *RedisPatchScheduleList) DeepCopy() *RedisPatchScheduleList {
	if in == nil {
		return nil
	}
	out := new(RedisPatchScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisPatchScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisPatchSchedule_STATUS) DeepCopyInto(out *RedisPatchSchedule_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.ScheduleEntries != nil {
		in, out := &in.ScheduleEntries, &out.ScheduleEntries
		*out = make([]ScheduleEntry_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisPatchSchedule_STATUS.
func (in *RedisPatchSchedule_STATUS) DeepCopy() *RedisPatchSchedule_STATUS {
	if in == nil {
		return nil
	}
	out := new(RedisPatchSchedule_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisPatchSchedule_Spec) DeepCopyInto(out *RedisPatchSchedule_Spec) {
	*out = *in
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
	if in.ScheduleEntries != nil {
		in, out := &in.ScheduleEntries, &out.ScheduleEntries
		*out = make([]ScheduleEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisPatchSchedule_Spec.
func (in *RedisPatchSchedule_Spec) DeepCopy() *RedisPatchSchedule_Spec {
	if in == nil {
		return nil
	}
	out := new(RedisPatchSchedule_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis_STATUS) DeepCopyInto(out *Redis_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableNonSslPort != nil {
		in, out := &in.EnableNonSslPort, &out.EnableNonSslPort
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
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
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
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = new(RedisCreateProperties_RedisConfiguration_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(int)
		**out = **in
	}
	if in.ReplicasPerPrimary != nil {
		in, out := &in.ReplicasPerPrimary, &out.ReplicasPerPrimary
		*out = new(int)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.StaticIP != nil {
		in, out := &in.StaticIP, &out.StaticIP
		*out = new(string)
		**out = **in
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
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
	if in.TenantSettings != nil {
		in, out := &in.TenantSettings, &out.TenantSettings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis_STATUS.
func (in *Redis_STATUS) DeepCopy() *Redis_STATUS {
	if in == nil {
		return nil
	}
	out := new(Redis_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis_Spec) DeepCopyInto(out *Redis_Spec) {
	*out = *in
	if in.EnableNonSslPort != nil {
		in, out := &in.EnableNonSslPort, &out.EnableNonSslPort
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(string)
		**out = **in
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(RedisOperatorSpec)
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
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(int)
		**out = **in
	}
	if in.ReplicasPerPrimary != nil {
		in, out := &in.ReplicasPerPrimary, &out.ReplicasPerPrimary
		*out = new(int)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.StaticIP != nil {
		in, out := &in.StaticIP, &out.StaticIP
		*out = new(string)
		**out = **in
	}
	if in.SubnetReference != nil {
		in, out := &in.SubnetReference, &out.SubnetReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TenantSettings != nil {
		in, out := &in.TenantSettings, &out.TenantSettings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis_Spec.
func (in *Redis_Spec) DeepCopy() *Redis_Spec {
	if in == nil {
		return nil
	}
	out := new(Redis_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleEntry) DeepCopyInto(out *ScheduleEntry) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
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
	if in.StartHourUtc != nil {
		in, out := &in.StartHourUtc, &out.StartHourUtc
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleEntry.
func (in *ScheduleEntry) DeepCopy() *ScheduleEntry {
	if in == nil {
		return nil
	}
	out := new(ScheduleEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleEntry_STATUS) DeepCopyInto(out *ScheduleEntry_STATUS) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
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
	if in.StartHourUtc != nil {
		in, out := &in.StartHourUtc, &out.StartHourUtc
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleEntry_STATUS.
func (in *ScheduleEntry_STATUS) DeepCopy() *ScheduleEntry_STATUS {
	if in == nil {
		return nil
	}
	out := new(ScheduleEntry_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku) DeepCopyInto(out *Sku) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku.
func (in *Sku) DeepCopy() *Sku {
	if in == nil {
		return nil
	}
	out := new(Sku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_STATUS) DeepCopyInto(out *Sku_STATUS) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_STATUS.
func (in *Sku_STATUS) DeepCopy() *Sku_STATUS {
	if in == nil {
		return nil
	}
	out := new(Sku_STATUS)
	in.DeepCopyInto(out)
	return out
}
