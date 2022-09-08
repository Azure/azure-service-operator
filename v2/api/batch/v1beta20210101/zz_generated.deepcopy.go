//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta20210101

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStorageBaseProperties) DeepCopyInto(out *AutoStorageBaseProperties) {
	*out = *in
	if in.StorageAccountReference != nil {
		in, out := &in.StorageAccountReference, &out.StorageAccountReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStorageBaseProperties.
func (in *AutoStorageBaseProperties) DeepCopy() *AutoStorageBaseProperties {
	if in == nil {
		return nil
	}
	out := new(AutoStorageBaseProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStorageBasePropertiesARM) DeepCopyInto(out *AutoStorageBasePropertiesARM) {
	*out = *in
	if in.StorageAccountId != nil {
		in, out := &in.StorageAccountId, &out.StorageAccountId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStorageBasePropertiesARM.
func (in *AutoStorageBasePropertiesARM) DeepCopy() *AutoStorageBasePropertiesARM {
	if in == nil {
		return nil
	}
	out := new(AutoStorageBasePropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStorageProperties_STATUS) DeepCopyInto(out *AutoStorageProperties_STATUS) {
	*out = *in
	if in.LastKeySync != nil {
		in, out := &in.LastKeySync, &out.LastKeySync
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountId != nil {
		in, out := &in.StorageAccountId, &out.StorageAccountId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStorageProperties_STATUS.
func (in *AutoStorageProperties_STATUS) DeepCopy() *AutoStorageProperties_STATUS {
	if in == nil {
		return nil
	}
	out := new(AutoStorageProperties_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStorageProperties_STATUSARM) DeepCopyInto(out *AutoStorageProperties_STATUSARM) {
	*out = *in
	if in.LastKeySync != nil {
		in, out := &in.LastKeySync, &out.LastKeySync
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountId != nil {
		in, out := &in.StorageAccountId, &out.StorageAccountId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStorageProperties_STATUSARM.
func (in *AutoStorageProperties_STATUSARM) DeepCopy() *AutoStorageProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(AutoStorageProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccount) DeepCopyInto(out *BatchAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccount.
func (in *BatchAccount) DeepCopy() *BatchAccount {
	if in == nil {
		return nil
	}
	out := new(BatchAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountCreatePropertiesARM) DeepCopyInto(out *BatchAccountCreatePropertiesARM) {
	*out = *in
	if in.AutoStorage != nil {
		in, out := &in.AutoStorage, &out.AutoStorage
		*out = new(AutoStorageBasePropertiesARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionPropertiesARM)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(KeyVaultReferenceARM)
		(*in).DeepCopyInto(*out)
	}
	if in.PoolAllocationMode != nil {
		in, out := &in.PoolAllocationMode, &out.PoolAllocationMode
		*out = new(BatchAccountCreateProperties_PoolAllocationMode)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(BatchAccountCreateProperties_PublicNetworkAccess)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountCreatePropertiesARM.
func (in *BatchAccountCreatePropertiesARM) DeepCopy() *BatchAccountCreatePropertiesARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccountCreatePropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity) DeepCopyInto(out *BatchAccountIdentity) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(BatchAccountIdentity_Type)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity.
func (in *BatchAccountIdentity) DeepCopy() *BatchAccountIdentity {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentityARM) DeepCopyInto(out *BatchAccountIdentityARM) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(BatchAccountIdentity_Type)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentityARM.
func (in *BatchAccountIdentityARM) DeepCopy() *BatchAccountIdentityARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentityARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity_STATUS) DeepCopyInto(out *BatchAccountIdentity_STATUS) {
	*out = *in
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(BatchAccountIdentity_STATUS_Type)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]BatchAccountIdentity_STATUS_UserAssignedIdentities, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_STATUS.
func (in *BatchAccountIdentity_STATUS) DeepCopy() *BatchAccountIdentity_STATUS {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity_STATUSARM) DeepCopyInto(out *BatchAccountIdentity_STATUSARM) {
	*out = *in
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(BatchAccountIdentity_STATUS_Type)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]BatchAccountIdentity_STATUS_UserAssignedIdentitiesARM, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_STATUSARM.
func (in *BatchAccountIdentity_STATUSARM) DeepCopy() *BatchAccountIdentity_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity_STATUS_UserAssignedIdentities) DeepCopyInto(out *BatchAccountIdentity_STATUS_UserAssignedIdentities) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_STATUS_UserAssignedIdentities.
func (in *BatchAccountIdentity_STATUS_UserAssignedIdentities) DeepCopy() *BatchAccountIdentity_STATUS_UserAssignedIdentities {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_STATUS_UserAssignedIdentities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity_STATUS_UserAssignedIdentitiesARM) DeepCopyInto(out *BatchAccountIdentity_STATUS_UserAssignedIdentitiesARM) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_STATUS_UserAssignedIdentitiesARM.
func (in *BatchAccountIdentity_STATUS_UserAssignedIdentitiesARM) DeepCopy() *BatchAccountIdentity_STATUS_UserAssignedIdentitiesARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_STATUS_UserAssignedIdentitiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountList) DeepCopyInto(out *BatchAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BatchAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountList.
func (in *BatchAccountList) DeepCopy() *BatchAccountList {
	if in == nil {
		return nil
	}
	out := new(BatchAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountProperties_STATUSARM) DeepCopyInto(out *BatchAccountProperties_STATUSARM) {
	*out = *in
	if in.AccountEndpoint != nil {
		in, out := &in.AccountEndpoint, &out.AccountEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ActiveJobAndJobScheduleQuota != nil {
		in, out := &in.ActiveJobAndJobScheduleQuota, &out.ActiveJobAndJobScheduleQuota
		*out = new(int)
		**out = **in
	}
	if in.AutoStorage != nil {
		in, out := &in.AutoStorage, &out.AutoStorage
		*out = new(AutoStorageProperties_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.DedicatedCoreQuota != nil {
		in, out := &in.DedicatedCoreQuota, &out.DedicatedCoreQuota
		*out = new(int)
		**out = **in
	}
	if in.DedicatedCoreQuotaPerVMFamily != nil {
		in, out := &in.DedicatedCoreQuotaPerVMFamily, &out.DedicatedCoreQuotaPerVMFamily
		*out = make([]VirtualMachineFamilyCoreQuota_STATUSARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DedicatedCoreQuotaPerVMFamilyEnforced != nil {
		in, out := &in.DedicatedCoreQuotaPerVMFamilyEnforced, &out.DedicatedCoreQuotaPerVMFamilyEnforced
		*out = new(bool)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionProperties_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(KeyVaultReference_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.LowPriorityCoreQuota != nil {
		in, out := &in.LowPriorityCoreQuota, &out.LowPriorityCoreQuota
		*out = new(int)
		**out = **in
	}
	if in.PoolAllocationMode != nil {
		in, out := &in.PoolAllocationMode, &out.PoolAllocationMode
		*out = new(PoolAllocationMode_STATUS)
		**out = **in
	}
	if in.PoolQuota != nil {
		in, out := &in.PoolQuota, &out.PoolQuota
		*out = new(int)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_STATUSARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(BatchAccountProperties_STATUS_ProvisioningState)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(PublicNetworkAccessType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountProperties_STATUSARM.
func (in *BatchAccountProperties_STATUSARM) DeepCopy() *BatchAccountProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccountProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccount_STATUS) DeepCopyInto(out *BatchAccount_STATUS) {
	*out = *in
	if in.AccountEndpoint != nil {
		in, out := &in.AccountEndpoint, &out.AccountEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ActiveJobAndJobScheduleQuota != nil {
		in, out := &in.ActiveJobAndJobScheduleQuota, &out.ActiveJobAndJobScheduleQuota
		*out = new(int)
		**out = **in
	}
	if in.AutoStorage != nil {
		in, out := &in.AutoStorage, &out.AutoStorage
		*out = new(AutoStorageProperties_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DedicatedCoreQuota != nil {
		in, out := &in.DedicatedCoreQuota, &out.DedicatedCoreQuota
		*out = new(int)
		**out = **in
	}
	if in.DedicatedCoreQuotaPerVMFamily != nil {
		in, out := &in.DedicatedCoreQuotaPerVMFamily, &out.DedicatedCoreQuotaPerVMFamily
		*out = make([]VirtualMachineFamilyCoreQuota_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DedicatedCoreQuotaPerVMFamilyEnforced != nil {
		in, out := &in.DedicatedCoreQuotaPerVMFamilyEnforced, &out.DedicatedCoreQuotaPerVMFamilyEnforced
		*out = new(bool)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionProperties_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(BatchAccountIdentity_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(KeyVaultReference_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LowPriorityCoreQuota != nil {
		in, out := &in.LowPriorityCoreQuota, &out.LowPriorityCoreQuota
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PoolAllocationMode != nil {
		in, out := &in.PoolAllocationMode, &out.PoolAllocationMode
		*out = new(PoolAllocationMode_STATUS)
		**out = **in
	}
	if in.PoolQuota != nil {
		in, out := &in.PoolQuota, &out.PoolQuota
		*out = new(int)
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
		*out = new(BatchAccountProperties_STATUS_ProvisioningState)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(PublicNetworkAccessType_STATUS)
		**out = **in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccount_STATUS.
func (in *BatchAccount_STATUS) DeepCopy() *BatchAccount_STATUS {
	if in == nil {
		return nil
	}
	out := new(BatchAccount_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccount_STATUSARM) DeepCopyInto(out *BatchAccount_STATUSARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(BatchAccountIdentity_STATUSARM)
		(*in).DeepCopyInto(*out)
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
		*out = new(BatchAccountProperties_STATUSARM)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccount_STATUSARM.
func (in *BatchAccount_STATUSARM) DeepCopy() *BatchAccount_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccount_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccount_Spec) DeepCopyInto(out *BatchAccount_Spec) {
	*out = *in
	if in.AutoStorage != nil {
		in, out := &in.AutoStorage, &out.AutoStorage
		*out = new(AutoStorageBaseProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(BatchAccountIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(KeyVaultReference)
		(*in).DeepCopyInto(*out)
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
	if in.PoolAllocationMode != nil {
		in, out := &in.PoolAllocationMode, &out.PoolAllocationMode
		*out = new(BatchAccountCreateProperties_PoolAllocationMode)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(BatchAccountCreateProperties_PublicNetworkAccess)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccount_Spec.
func (in *BatchAccount_Spec) DeepCopy() *BatchAccount_Spec {
	if in == nil {
		return nil
	}
	out := new(BatchAccount_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccount_SpecARM) DeepCopyInto(out *BatchAccount_SpecARM) {
	*out = *in
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(BatchAccountIdentityARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(BatchAccountCreatePropertiesARM)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccount_SpecARM.
func (in *BatchAccount_SpecARM) DeepCopy() *BatchAccount_SpecARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccount_SpecARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionProperties) DeepCopyInto(out *EncryptionProperties) {
	*out = *in
	if in.KeySource != nil {
		in, out := &in.KeySource, &out.KeySource
		*out = new(EncryptionProperties_KeySource)
		**out = **in
	}
	if in.KeyVaultProperties != nil {
		in, out := &in.KeyVaultProperties, &out.KeyVaultProperties
		*out = new(KeyVaultProperties)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionProperties.
func (in *EncryptionProperties) DeepCopy() *EncryptionProperties {
	if in == nil {
		return nil
	}
	out := new(EncryptionProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionPropertiesARM) DeepCopyInto(out *EncryptionPropertiesARM) {
	*out = *in
	if in.KeySource != nil {
		in, out := &in.KeySource, &out.KeySource
		*out = new(EncryptionProperties_KeySource)
		**out = **in
	}
	if in.KeyVaultProperties != nil {
		in, out := &in.KeyVaultProperties, &out.KeyVaultProperties
		*out = new(KeyVaultPropertiesARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionPropertiesARM.
func (in *EncryptionPropertiesARM) DeepCopy() *EncryptionPropertiesARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionPropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionProperties_STATUS) DeepCopyInto(out *EncryptionProperties_STATUS) {
	*out = *in
	if in.KeySource != nil {
		in, out := &in.KeySource, &out.KeySource
		*out = new(EncryptionProperties_STATUS_KeySource)
		**out = **in
	}
	if in.KeyVaultProperties != nil {
		in, out := &in.KeyVaultProperties, &out.KeyVaultProperties
		*out = new(KeyVaultProperties_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionProperties_STATUS.
func (in *EncryptionProperties_STATUS) DeepCopy() *EncryptionProperties_STATUS {
	if in == nil {
		return nil
	}
	out := new(EncryptionProperties_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionProperties_STATUSARM) DeepCopyInto(out *EncryptionProperties_STATUSARM) {
	*out = *in
	if in.KeySource != nil {
		in, out := &in.KeySource, &out.KeySource
		*out = new(EncryptionProperties_STATUS_KeySource)
		**out = **in
	}
	if in.KeyVaultProperties != nil {
		in, out := &in.KeyVaultProperties, &out.KeyVaultProperties
		*out = new(KeyVaultProperties_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionProperties_STATUSARM.
func (in *EncryptionProperties_STATUSARM) DeepCopy() *EncryptionProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultProperties) DeepCopyInto(out *KeyVaultProperties) {
	*out = *in
	if in.KeyIdentifier != nil {
		in, out := &in.KeyIdentifier, &out.KeyIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultProperties.
func (in *KeyVaultProperties) DeepCopy() *KeyVaultProperties {
	if in == nil {
		return nil
	}
	out := new(KeyVaultProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultPropertiesARM) DeepCopyInto(out *KeyVaultPropertiesARM) {
	*out = *in
	if in.KeyIdentifier != nil {
		in, out := &in.KeyIdentifier, &out.KeyIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultPropertiesARM.
func (in *KeyVaultPropertiesARM) DeepCopy() *KeyVaultPropertiesARM {
	if in == nil {
		return nil
	}
	out := new(KeyVaultPropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultProperties_STATUS) DeepCopyInto(out *KeyVaultProperties_STATUS) {
	*out = *in
	if in.KeyIdentifier != nil {
		in, out := &in.KeyIdentifier, &out.KeyIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultProperties_STATUS.
func (in *KeyVaultProperties_STATUS) DeepCopy() *KeyVaultProperties_STATUS {
	if in == nil {
		return nil
	}
	out := new(KeyVaultProperties_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultProperties_STATUSARM) DeepCopyInto(out *KeyVaultProperties_STATUSARM) {
	*out = *in
	if in.KeyIdentifier != nil {
		in, out := &in.KeyIdentifier, &out.KeyIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultProperties_STATUSARM.
func (in *KeyVaultProperties_STATUSARM) DeepCopy() *KeyVaultProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(KeyVaultProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultReference) DeepCopyInto(out *KeyVaultReference) {
	*out = *in
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultReference.
func (in *KeyVaultReference) DeepCopy() *KeyVaultReference {
	if in == nil {
		return nil
	}
	out := new(KeyVaultReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultReferenceARM) DeepCopyInto(out *KeyVaultReferenceARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultReferenceARM.
func (in *KeyVaultReferenceARM) DeepCopy() *KeyVaultReferenceARM {
	if in == nil {
		return nil
	}
	out := new(KeyVaultReferenceARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultReference_STATUS) DeepCopyInto(out *KeyVaultReference_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultReference_STATUS.
func (in *KeyVaultReference_STATUS) DeepCopy() *KeyVaultReference_STATUS {
	if in == nil {
		return nil
	}
	out := new(KeyVaultReference_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultReference_STATUSARM) DeepCopyInto(out *KeyVaultReference_STATUSARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultReference_STATUSARM.
func (in *KeyVaultReference_STATUSARM) DeepCopy() *KeyVaultReference_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(KeyVaultReference_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnectionProperties_STATUSARM) DeepCopyInto(out *PrivateEndpointConnectionProperties_STATUSARM) {
	*out = *in
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(PrivateEndpoint_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateLinkServiceConnectionState != nil {
		in, out := &in.PrivateLinkServiceConnectionState, &out.PrivateLinkServiceConnectionState
		*out = new(PrivateLinkServiceConnectionState_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(PrivateEndpointConnectionProperties_STATUS_ProvisioningState)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnectionProperties_STATUSARM.
func (in *PrivateEndpointConnectionProperties_STATUSARM) DeepCopy() *PrivateEndpointConnectionProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnectionProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_STATUS) DeepCopyInto(out *PrivateEndpointConnection_STATUS) {
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(PrivateEndpoint_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateLinkServiceConnectionState != nil {
		in, out := &in.PrivateLinkServiceConnectionState, &out.PrivateLinkServiceConnectionState
		*out = new(PrivateLinkServiceConnectionState_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(PrivateEndpointConnectionProperties_STATUS_ProvisioningState)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
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
func (in *PrivateEndpointConnection_STATUSARM) DeepCopyInto(out *PrivateEndpointConnection_STATUSARM) {
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(PrivateEndpointConnectionProperties_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_STATUSARM.
func (in *PrivateEndpointConnection_STATUSARM) DeepCopy() *PrivateEndpointConnection_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpoint_STATUS) DeepCopyInto(out *PrivateEndpoint_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpoint_STATUS.
func (in *PrivateEndpoint_STATUS) DeepCopy() *PrivateEndpoint_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpoint_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpoint_STATUSARM) DeepCopyInto(out *PrivateEndpoint_STATUSARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpoint_STATUSARM.
func (in *PrivateEndpoint_STATUSARM) DeepCopy() *PrivateEndpoint_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpoint_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkServiceConnectionState_STATUS) DeepCopyInto(out *PrivateLinkServiceConnectionState_STATUS) {
	*out = *in
	if in.ActionRequired != nil {
		in, out := &in.ActionRequired, &out.ActionRequired
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(PrivateLinkServiceConnectionStatus_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateLinkServiceConnectionState_STATUS.
func (in *PrivateLinkServiceConnectionState_STATUS) DeepCopy() *PrivateLinkServiceConnectionState_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateLinkServiceConnectionState_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkServiceConnectionState_STATUSARM) DeepCopyInto(out *PrivateLinkServiceConnectionState_STATUSARM) {
	*out = *in
	if in.ActionRequired != nil {
		in, out := &in.ActionRequired, &out.ActionRequired
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(PrivateLinkServiceConnectionStatus_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateLinkServiceConnectionState_STATUSARM.
func (in *PrivateLinkServiceConnectionState_STATUSARM) DeepCopy() *PrivateLinkServiceConnectionState_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(PrivateLinkServiceConnectionState_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineFamilyCoreQuota_STATUS) DeepCopyInto(out *VirtualMachineFamilyCoreQuota_STATUS) {
	*out = *in
	if in.CoreQuota != nil {
		in, out := &in.CoreQuota, &out.CoreQuota
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineFamilyCoreQuota_STATUS.
func (in *VirtualMachineFamilyCoreQuota_STATUS) DeepCopy() *VirtualMachineFamilyCoreQuota_STATUS {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineFamilyCoreQuota_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineFamilyCoreQuota_STATUSARM) DeepCopyInto(out *VirtualMachineFamilyCoreQuota_STATUSARM) {
	*out = *in
	if in.CoreQuota != nil {
		in, out := &in.CoreQuota, &out.CoreQuota
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineFamilyCoreQuota_STATUSARM.
func (in *VirtualMachineFamilyCoreQuota_STATUSARM) DeepCopy() *VirtualMachineFamilyCoreQuota_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineFamilyCoreQuota_STATUSARM)
	in.DeepCopyInto(out)
	return out
}
