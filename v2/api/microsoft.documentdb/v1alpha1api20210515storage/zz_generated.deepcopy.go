//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticalStorageConfiguration) DeepCopyInto(out *AnalyticalStorageConfiguration) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SchemaType != nil {
		in, out := &in.SchemaType, &out.SchemaType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticalStorageConfiguration.
func (in *AnalyticalStorageConfiguration) DeepCopy() *AnalyticalStorageConfiguration {
	if in == nil {
		return nil
	}
	out := new(AnalyticalStorageConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticalStorageConfiguration_Status) DeepCopyInto(out *AnalyticalStorageConfiguration_Status) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SchemaType != nil {
		in, out := &in.SchemaType, &out.SchemaType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticalStorageConfiguration_Status.
func (in *AnalyticalStorageConfiguration_Status) DeepCopy() *AnalyticalStorageConfiguration_Status {
	if in == nil {
		return nil
	}
	out := new(AnalyticalStorageConfiguration_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiProperties) DeepCopyInto(out *ApiProperties) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServerVersion != nil {
		in, out := &in.ServerVersion, &out.ServerVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiProperties.
func (in *ApiProperties) DeepCopy() *ApiProperties {
	if in == nil {
		return nil
	}
	out := new(ApiProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiProperties_Status) DeepCopyInto(out *ApiProperties_Status) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServerVersion != nil {
		in, out := &in.ServerVersion, &out.ServerVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiProperties_Status.
func (in *ApiProperties_Status) DeepCopy() *ApiProperties_Status {
	if in == nil {
		return nil
	}
	out := new(ApiProperties_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicy) DeepCopyInto(out *BackupPolicy) {
	*out = *in
	if in.ContinuousModeBackupPolicy != nil {
		in, out := &in.ContinuousModeBackupPolicy, &out.ContinuousModeBackupPolicy
		*out = new(ContinuousModeBackupPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.PeriodicModeBackupPolicy != nil {
		in, out := &in.PeriodicModeBackupPolicy, &out.PeriodicModeBackupPolicy
		*out = new(PeriodicModeBackupPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicy.
func (in *BackupPolicy) DeepCopy() *BackupPolicy {
	if in == nil {
		return nil
	}
	out := new(BackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicy_Status) DeepCopyInto(out *BackupPolicy_Status) {
	*out = *in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicy_Status.
func (in *BackupPolicy_Status) DeepCopy() *BackupPolicy_Status {
	if in == nil {
		return nil
	}
	out := new(BackupPolicy_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Capability) DeepCopyInto(out *Capability) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Capability.
func (in *Capability) DeepCopy() *Capability {
	if in == nil {
		return nil
	}
	out := new(Capability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Capability_Status) DeepCopyInto(out *Capability_Status) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Capability_Status.
func (in *Capability_Status) DeepCopy() *Capability_Status {
	if in == nil {
		return nil
	}
	out := new(Capability_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsistencyPolicy) DeepCopyInto(out *ConsistencyPolicy) {
	*out = *in
	if in.DefaultConsistencyLevel != nil {
		in, out := &in.DefaultConsistencyLevel, &out.DefaultConsistencyLevel
		*out = new(string)
		**out = **in
	}
	if in.MaxIntervalInSeconds != nil {
		in, out := &in.MaxIntervalInSeconds, &out.MaxIntervalInSeconds
		*out = new(int)
		**out = **in
	}
	if in.MaxStalenessPrefix != nil {
		in, out := &in.MaxStalenessPrefix, &out.MaxStalenessPrefix
		*out = new(int)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsistencyPolicy.
func (in *ConsistencyPolicy) DeepCopy() *ConsistencyPolicy {
	if in == nil {
		return nil
	}
	out := new(ConsistencyPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsistencyPolicy_Status) DeepCopyInto(out *ConsistencyPolicy_Status) {
	*out = *in
	if in.DefaultConsistencyLevel != nil {
		in, out := &in.DefaultConsistencyLevel, &out.DefaultConsistencyLevel
		*out = new(string)
		**out = **in
	}
	if in.MaxIntervalInSeconds != nil {
		in, out := &in.MaxIntervalInSeconds, &out.MaxIntervalInSeconds
		*out = new(int)
		**out = **in
	}
	if in.MaxStalenessPrefix != nil {
		in, out := &in.MaxStalenessPrefix, &out.MaxStalenessPrefix
		*out = new(int)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsistencyPolicy_Status.
func (in *ConsistencyPolicy_Status) DeepCopy() *ConsistencyPolicy_Status {
	if in == nil {
		return nil
	}
	out := new(ConsistencyPolicy_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContinuousModeBackupPolicy) DeepCopyInto(out *ContinuousModeBackupPolicy) {
	*out = *in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContinuousModeBackupPolicy.
func (in *ContinuousModeBackupPolicy) DeepCopy() *ContinuousModeBackupPolicy {
	if in == nil {
		return nil
	}
	out := new(ContinuousModeBackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsPolicy) DeepCopyInto(out *CorsPolicy) {
	*out = *in
	if in.AllowedHeaders != nil {
		in, out := &in.AllowedHeaders, &out.AllowedHeaders
		*out = new(string)
		**out = **in
	}
	if in.AllowedMethods != nil {
		in, out := &in.AllowedMethods, &out.AllowedMethods
		*out = new(string)
		**out = **in
	}
	if in.AllowedOrigins != nil {
		in, out := &in.AllowedOrigins, &out.AllowedOrigins
		*out = new(string)
		**out = **in
	}
	if in.ExposedHeaders != nil {
		in, out := &in.ExposedHeaders, &out.ExposedHeaders
		*out = new(string)
		**out = **in
	}
	if in.MaxAgeInSeconds != nil {
		in, out := &in.MaxAgeInSeconds, &out.MaxAgeInSeconds
		*out = new(int)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsPolicy.
func (in *CorsPolicy) DeepCopy() *CorsPolicy {
	if in == nil {
		return nil
	}
	out := new(CorsPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsPolicy_Status) DeepCopyInto(out *CorsPolicy_Status) {
	*out = *in
	if in.AllowedHeaders != nil {
		in, out := &in.AllowedHeaders, &out.AllowedHeaders
		*out = new(string)
		**out = **in
	}
	if in.AllowedMethods != nil {
		in, out := &in.AllowedMethods, &out.AllowedMethods
		*out = new(string)
		**out = **in
	}
	if in.AllowedOrigins != nil {
		in, out := &in.AllowedOrigins, &out.AllowedOrigins
		*out = new(string)
		**out = **in
	}
	if in.ExposedHeaders != nil {
		in, out := &in.ExposedHeaders, &out.ExposedHeaders
		*out = new(string)
		**out = **in
	}
	if in.MaxAgeInSeconds != nil {
		in, out := &in.MaxAgeInSeconds, &out.MaxAgeInSeconds
		*out = new(int)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsPolicy_Status.
func (in *CorsPolicy_Status) DeepCopy() *CorsPolicy_Status {
	if in == nil {
		return nil
	}
	out := new(CorsPolicy_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseAccount) DeepCopyInto(out *DatabaseAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseAccount.
func (in *DatabaseAccount) DeepCopy() *DatabaseAccount {
	if in == nil {
		return nil
	}
	out := new(DatabaseAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseAccountGetResults_Status) DeepCopyInto(out *DatabaseAccountGetResults_Status) {
	*out = *in
	if in.AnalyticalStorageConfiguration != nil {
		in, out := &in.AnalyticalStorageConfiguration, &out.AnalyticalStorageConfiguration
		*out = new(AnalyticalStorageConfiguration_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiProperties != nil {
		in, out := &in.ApiProperties, &out.ApiProperties
		*out = new(ApiProperties_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupPolicy != nil {
		in, out := &in.BackupPolicy, &out.BackupPolicy
		*out = new(BackupPolicy_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]Capability_Status, len(*in))
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
	if in.ConnectorOffer != nil {
		in, out := &in.ConnectorOffer, &out.ConnectorOffer
		*out = new(string)
		**out = **in
	}
	if in.ConsistencyPolicy != nil {
		in, out := &in.ConsistencyPolicy, &out.ConsistencyPolicy
		*out = new(ConsistencyPolicy_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = make([]CorsPolicy_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DatabaseAccountOfferType != nil {
		in, out := &in.DatabaseAccountOfferType, &out.DatabaseAccountOfferType
		*out = new(string)
		**out = **in
	}
	if in.DefaultIdentity != nil {
		in, out := &in.DefaultIdentity, &out.DefaultIdentity
		*out = new(string)
		**out = **in
	}
	if in.DisableKeyBasedMetadataWriteAccess != nil {
		in, out := &in.DisableKeyBasedMetadataWriteAccess, &out.DisableKeyBasedMetadataWriteAccess
		*out = new(bool)
		**out = **in
	}
	if in.DocumentEndpoint != nil {
		in, out := &in.DocumentEndpoint, &out.DocumentEndpoint
		*out = new(string)
		**out = **in
	}
	if in.EnableAnalyticalStorage != nil {
		in, out := &in.EnableAnalyticalStorage, &out.EnableAnalyticalStorage
		*out = new(bool)
		**out = **in
	}
	if in.EnableAutomaticFailover != nil {
		in, out := &in.EnableAutomaticFailover, &out.EnableAutomaticFailover
		*out = new(bool)
		**out = **in
	}
	if in.EnableCassandraConnector != nil {
		in, out := &in.EnableCassandraConnector, &out.EnableCassandraConnector
		*out = new(bool)
		**out = **in
	}
	if in.EnableFreeTier != nil {
		in, out := &in.EnableFreeTier, &out.EnableFreeTier
		*out = new(bool)
		**out = **in
	}
	if in.EnableMultipleWriteLocations != nil {
		in, out := &in.EnableMultipleWriteLocations, &out.EnableMultipleWriteLocations
		*out = new(bool)
		**out = **in
	}
	if in.FailoverPolicies != nil {
		in, out := &in.FailoverPolicies, &out.FailoverPolicies
		*out = make([]FailoverPolicy_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(ManagedServiceIdentity_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.IpRules != nil {
		in, out := &in.IpRules, &out.IpRules
		*out = make([]IpAddressOrRange_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsVirtualNetworkFilterEnabled != nil {
		in, out := &in.IsVirtualNetworkFilterEnabled, &out.IsVirtualNetworkFilterEnabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyVaultKeyUri != nil {
		in, out := &in.KeyVaultKeyUri, &out.KeyVaultKeyUri
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
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]Location_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkAclBypass != nil {
		in, out := &in.NetworkAclBypass, &out.NetworkAclBypass
		*out = new(string)
		**out = **in
	}
	if in.NetworkAclBypassResourceIds != nil {
		in, out := &in.NetworkAclBypassResourceIds, &out.NetworkAclBypassResourceIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_Status_SubResourceEmbedded, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.ReadLocations != nil {
		in, out := &in.ReadLocations, &out.ReadLocations
		*out = make([]Location_Status, len(*in))
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
	if in.VirtualNetworkRules != nil {
		in, out := &in.VirtualNetworkRules, &out.VirtualNetworkRules
		*out = make([]VirtualNetworkRule_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WriteLocations != nil {
		in, out := &in.WriteLocations, &out.WriteLocations
		*out = make([]Location_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseAccountGetResults_Status.
func (in *DatabaseAccountGetResults_Status) DeepCopy() *DatabaseAccountGetResults_Status {
	if in == nil {
		return nil
	}
	out := new(DatabaseAccountGetResults_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseAccountList) DeepCopyInto(out *DatabaseAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseAccountList.
func (in *DatabaseAccountList) DeepCopy() *DatabaseAccountList {
	if in == nil {
		return nil
	}
	out := new(DatabaseAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseAccounts_Spec) DeepCopyInto(out *DatabaseAccounts_Spec) {
	*out = *in
	if in.AnalyticalStorageConfiguration != nil {
		in, out := &in.AnalyticalStorageConfiguration, &out.AnalyticalStorageConfiguration
		*out = new(AnalyticalStorageConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiProperties != nil {
		in, out := &in.ApiProperties, &out.ApiProperties
		*out = new(ApiProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupPolicy != nil {
		in, out := &in.BackupPolicy, &out.BackupPolicy
		*out = new(BackupPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]Capability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConnectorOffer != nil {
		in, out := &in.ConnectorOffer, &out.ConnectorOffer
		*out = new(string)
		**out = **in
	}
	if in.ConsistencyPolicy != nil {
		in, out := &in.ConsistencyPolicy, &out.ConsistencyPolicy
		*out = new(ConsistencyPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = make([]CorsPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DatabaseAccountOfferType != nil {
		in, out := &in.DatabaseAccountOfferType, &out.DatabaseAccountOfferType
		*out = new(string)
		**out = **in
	}
	if in.DefaultIdentity != nil {
		in, out := &in.DefaultIdentity, &out.DefaultIdentity
		*out = new(string)
		**out = **in
	}
	if in.DisableKeyBasedMetadataWriteAccess != nil {
		in, out := &in.DisableKeyBasedMetadataWriteAccess, &out.DisableKeyBasedMetadataWriteAccess
		*out = new(bool)
		**out = **in
	}
	if in.EnableAnalyticalStorage != nil {
		in, out := &in.EnableAnalyticalStorage, &out.EnableAnalyticalStorage
		*out = new(bool)
		**out = **in
	}
	if in.EnableAutomaticFailover != nil {
		in, out := &in.EnableAutomaticFailover, &out.EnableAutomaticFailover
		*out = new(bool)
		**out = **in
	}
	if in.EnableCassandraConnector != nil {
		in, out := &in.EnableCassandraConnector, &out.EnableCassandraConnector
		*out = new(bool)
		**out = **in
	}
	if in.EnableFreeTier != nil {
		in, out := &in.EnableFreeTier, &out.EnableFreeTier
		*out = new(bool)
		**out = **in
	}
	if in.EnableMultipleWriteLocations != nil {
		in, out := &in.EnableMultipleWriteLocations, &out.EnableMultipleWriteLocations
		*out = new(bool)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(ManagedServiceIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.IpRules != nil {
		in, out := &in.IpRules, &out.IpRules
		*out = make([]IpAddressOrRange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsVirtualNetworkFilterEnabled != nil {
		in, out := &in.IsVirtualNetworkFilterEnabled, &out.IsVirtualNetworkFilterEnabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyVaultKeyUri != nil {
		in, out := &in.KeyVaultKeyUri, &out.KeyVaultKeyUri
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
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]Location, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAclBypass != nil {
		in, out := &in.NetworkAclBypass, &out.NetworkAclBypass
		*out = new(string)
		**out = **in
	}
	if in.NetworkAclBypassResourceIds != nil {
		in, out := &in.NetworkAclBypassResourceIds, &out.NetworkAclBypassResourceIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Owner = in.Owner
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
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VirtualNetworkRules != nil {
		in, out := &in.VirtualNetworkRules, &out.VirtualNetworkRules
		*out = make([]VirtualNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseAccounts_Spec.
func (in *DatabaseAccounts_Spec) DeepCopy() *DatabaseAccounts_Spec {
	if in == nil {
		return nil
	}
	out := new(DatabaseAccounts_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailoverPolicy_Status) DeepCopyInto(out *FailoverPolicy_Status) {
	*out = *in
	if in.FailoverPriority != nil {
		in, out := &in.FailoverPriority, &out.FailoverPriority
		*out = new(int)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.LocationName != nil {
		in, out := &in.LocationName, &out.LocationName
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailoverPolicy_Status.
func (in *FailoverPolicy_Status) DeepCopy() *FailoverPolicy_Status {
	if in == nil {
		return nil
	}
	out := new(FailoverPolicy_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpAddressOrRange) DeepCopyInto(out *IpAddressOrRange) {
	*out = *in
	if in.IpAddressOrRange != nil {
		in, out := &in.IpAddressOrRange, &out.IpAddressOrRange
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpAddressOrRange.
func (in *IpAddressOrRange) DeepCopy() *IpAddressOrRange {
	if in == nil {
		return nil
	}
	out := new(IpAddressOrRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpAddressOrRange_Status) DeepCopyInto(out *IpAddressOrRange_Status) {
	*out = *in
	if in.IpAddressOrRange != nil {
		in, out := &in.IpAddressOrRange, &out.IpAddressOrRange
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpAddressOrRange_Status.
func (in *IpAddressOrRange_Status) DeepCopy() *IpAddressOrRange_Status {
	if in == nil {
		return nil
	}
	out := new(IpAddressOrRange_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Location) DeepCopyInto(out *Location) {
	*out = *in
	if in.FailoverPriority != nil {
		in, out := &in.FailoverPriority, &out.FailoverPriority
		*out = new(int)
		**out = **in
	}
	if in.IsZoneRedundant != nil {
		in, out := &in.IsZoneRedundant, &out.IsZoneRedundant
		*out = new(bool)
		**out = **in
	}
	if in.LocationName != nil {
		in, out := &in.LocationName, &out.LocationName
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Location.
func (in *Location) DeepCopy() *Location {
	if in == nil {
		return nil
	}
	out := new(Location)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Location_Status) DeepCopyInto(out *Location_Status) {
	*out = *in
	if in.DocumentEndpoint != nil {
		in, out := &in.DocumentEndpoint, &out.DocumentEndpoint
		*out = new(string)
		**out = **in
	}
	if in.FailoverPriority != nil {
		in, out := &in.FailoverPriority, &out.FailoverPriority
		*out = new(int)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.IsZoneRedundant != nil {
		in, out := &in.IsZoneRedundant, &out.IsZoneRedundant
		*out = new(bool)
		**out = **in
	}
	if in.LocationName != nil {
		in, out := &in.LocationName, &out.LocationName
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Location_Status.
func (in *Location_Status) DeepCopy() *Location_Status {
	if in == nil {
		return nil
	}
	out := new(Location_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedServiceIdentity) DeepCopyInto(out *ManagedServiceIdentity) {
	*out = *in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedServiceIdentity.
func (in *ManagedServiceIdentity) DeepCopy() *ManagedServiceIdentity {
	if in == nil {
		return nil
	}
	out := new(ManagedServiceIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedServiceIdentity_Status) DeepCopyInto(out *ManagedServiceIdentity_Status) {
	*out = *in
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
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
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]ManagedServiceIdentity_Status_UserAssignedIdentities, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedServiceIdentity_Status.
func (in *ManagedServiceIdentity_Status) DeepCopy() *ManagedServiceIdentity_Status {
	if in == nil {
		return nil
	}
	out := new(ManagedServiceIdentity_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedServiceIdentity_Status_UserAssignedIdentities) DeepCopyInto(out *ManagedServiceIdentity_Status_UserAssignedIdentities) {
	*out = *in
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedServiceIdentity_Status_UserAssignedIdentities.
func (in *ManagedServiceIdentity_Status_UserAssignedIdentities) DeepCopy() *ManagedServiceIdentity_Status_UserAssignedIdentities {
	if in == nil {
		return nil
	}
	out := new(ManagedServiceIdentity_Status_UserAssignedIdentities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeriodicModeBackupPolicy) DeepCopyInto(out *PeriodicModeBackupPolicy) {
	*out = *in
	if in.PeriodicModeProperties != nil {
		in, out := &in.PeriodicModeProperties, &out.PeriodicModeProperties
		*out = new(PeriodicModeProperties)
		(*in).DeepCopyInto(*out)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeriodicModeBackupPolicy.
func (in *PeriodicModeBackupPolicy) DeepCopy() *PeriodicModeBackupPolicy {
	if in == nil {
		return nil
	}
	out := new(PeriodicModeBackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeriodicModeProperties) DeepCopyInto(out *PeriodicModeProperties) {
	*out = *in
	if in.BackupIntervalInMinutes != nil {
		in, out := &in.BackupIntervalInMinutes, &out.BackupIntervalInMinutes
		*out = new(int)
		**out = **in
	}
	if in.BackupRetentionIntervalInHours != nil {
		in, out := &in.BackupRetentionIntervalInHours, &out.BackupRetentionIntervalInHours
		*out = new(int)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeriodicModeProperties.
func (in *PeriodicModeProperties) DeepCopy() *PeriodicModeProperties {
	if in == nil {
		return nil
	}
	out := new(PeriodicModeProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_Status_SubResourceEmbedded) DeepCopyInto(out *PrivateEndpointConnection_Status_SubResourceEmbedded) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_Status_SubResourceEmbedded.
func (in *PrivateEndpointConnection_Status_SubResourceEmbedded) DeepCopy() *PrivateEndpointConnection_Status_SubResourceEmbedded {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_Status_SubResourceEmbedded)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRule) DeepCopyInto(out *VirtualNetworkRule) {
	*out = *in
	if in.IgnoreMissingVNetServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVNetServiceEndpoint, &out.IgnoreMissingVNetServiceEndpoint
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
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRule.
func (in *VirtualNetworkRule) DeepCopy() *VirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRule_Status) DeepCopyInto(out *VirtualNetworkRule_Status) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.IgnoreMissingVNetServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVNetServiceEndpoint, &out.IgnoreMissingVNetServiceEndpoint
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRule_Status.
func (in *VirtualNetworkRule_Status) DeepCopy() *VirtualNetworkRule_Status {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRule_Status)
	in.DeepCopyInto(out)
	return out
}
