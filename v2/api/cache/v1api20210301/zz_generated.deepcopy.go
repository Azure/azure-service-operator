//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20210301

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module_STATUS) DeepCopyInto(out *Module_STATUS) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module_STATUS.
func (in *Module_STATUS) DeepCopy() *Module_STATUS {
	if in == nil {
		return nil
	}
	out := new(Module_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence) DeepCopyInto(out *Persistence) {
	*out = *in
	if in.AofEnabled != nil {
		in, out := &in.AofEnabled, &out.AofEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofFrequency != nil {
		in, out := &in.AofFrequency, &out.AofFrequency
		*out = new(Persistence_AofFrequency)
		**out = **in
	}
	if in.RdbEnabled != nil {
		in, out := &in.RdbEnabled, &out.RdbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbFrequency != nil {
		in, out := &in.RdbFrequency, &out.RdbFrequency
		*out = new(Persistence_RdbFrequency)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence.
func (in *Persistence) DeepCopy() *Persistence {
	if in == nil {
		return nil
	}
	out := new(Persistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence_STATUS) DeepCopyInto(out *Persistence_STATUS) {
	*out = *in
	if in.AofEnabled != nil {
		in, out := &in.AofEnabled, &out.AofEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofFrequency != nil {
		in, out := &in.AofFrequency, &out.AofFrequency
		*out = new(Persistence_AofFrequency_STATUS)
		**out = **in
	}
	if in.RdbEnabled != nil {
		in, out := &in.RdbEnabled, &out.RdbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbFrequency != nil {
		in, out := &in.RdbFrequency, &out.RdbFrequency
		*out = new(Persistence_RdbFrequency_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence_STATUS.
func (in *Persistence_STATUS) DeepCopy() *Persistence_STATUS {
	if in == nil {
		return nil
	}
	out := new(Persistence_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_STATUS) DeepCopyInto(out *PrivateEndpointConnection_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_STATUS.
func (in *PrivateEndpointConnection_STATUS) DeepCopy() *PrivateEndpointConnection_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise) DeepCopyInto(out *RedisEnterprise) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise.
func (in *RedisEnterprise) DeepCopy() *RedisEnterprise {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterprise) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabase) DeepCopyInto(out *RedisEnterpriseDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabase.
func (in *RedisEnterpriseDatabase) DeepCopy() *RedisEnterpriseDatabase {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabaseList) DeepCopyInto(out *RedisEnterpriseDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisEnterpriseDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabaseList.
func (in *RedisEnterpriseDatabaseList) DeepCopy() *RedisEnterpriseDatabaseList {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabase_STATUS) DeepCopyInto(out *RedisEnterpriseDatabase_STATUS) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(DatabaseProperties_ClientProtocol_STATUS)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(DatabaseProperties_ClusteringPolicy_STATUS)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(DatabaseProperties_EvictionPolicy_STATUS)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS)
		**out = **in
	}
	if in.ResourceState != nil {
		in, out := &in.ResourceState, &out.ResourceState
		*out = new(ResourceState_STATUS)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabase_STATUS.
func (in *RedisEnterpriseDatabase_STATUS) DeepCopy() *RedisEnterpriseDatabase_STATUS {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabase_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabase_Spec) DeepCopyInto(out *RedisEnterpriseDatabase_Spec) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(DatabaseProperties_ClientProtocol)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(DatabaseProperties_ClusteringPolicy)
		**out = **in
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(DatabaseProperties_EvictionPolicy)
		**out = **in
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabase_Spec.
func (in *RedisEnterpriseDatabase_Spec) DeepCopy() *RedisEnterpriseDatabase_Spec {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabase_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseList) DeepCopyInto(out *RedisEnterpriseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisEnterprise, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseList.
func (in *RedisEnterpriseList) DeepCopy() *RedisEnterpriseList {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_STATUS) DeepCopyInto(out *RedisEnterprise_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
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
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(ClusterProperties_MinimumTlsVersion_STATUS)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS)
		**out = **in
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ResourceState != nil {
		in, out := &in.ResourceState, &out.ResourceState
		*out = new(ResourceState_STATUS)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS)
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
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_STATUS.
func (in *RedisEnterprise_STATUS) DeepCopy() *RedisEnterprise_STATUS {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_Spec) DeepCopyInto(out *RedisEnterprise_Spec) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(ClusterProperties_MinimumTlsVersion)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_Spec.
func (in *RedisEnterprise_Spec) DeepCopy() *RedisEnterprise_Spec {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_Spec)
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(Sku_Name)
		**out = **in
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(Sku_Name_STATUS)
		**out = **in
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
