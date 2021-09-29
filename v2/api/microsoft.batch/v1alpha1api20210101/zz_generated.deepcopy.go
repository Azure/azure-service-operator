//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20210101

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStorageBaseProperties) DeepCopyInto(out *AutoStorageBaseProperties) {
	*out = *in
	out.StorageAccountReference = in.StorageAccountReference
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
func (in *AutoStorageProperties_Status) DeepCopyInto(out *AutoStorageProperties_Status) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStorageProperties_Status.
func (in *AutoStorageProperties_Status) DeepCopy() *AutoStorageProperties_Status {
	if in == nil {
		return nil
	}
	out := new(AutoStorageProperties_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStorageProperties_StatusARM) DeepCopyInto(out *AutoStorageProperties_StatusARM) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStorageProperties_StatusARM.
func (in *AutoStorageProperties_StatusARM) DeepCopy() *AutoStorageProperties_StatusARM {
	if in == nil {
		return nil
	}
	out := new(AutoStorageProperties_StatusARM)
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
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionPropertiesARM)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(KeyVaultReferenceARM)
		**out = **in
	}
	if in.PoolAllocationMode != nil {
		in, out := &in.PoolAllocationMode, &out.PoolAllocationMode
		*out = new(BatchAccountCreatePropertiesPoolAllocationMode)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(BatchAccountCreatePropertiesPublicNetworkAccess)
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
func (in *BatchAccountIdentity_Status) DeepCopyInto(out *BatchAccountIdentity_Status) {
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
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]BatchAccountIdentity_Status_UserAssignedIdentities, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_Status.
func (in *BatchAccountIdentity_Status) DeepCopy() *BatchAccountIdentity_Status {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity_StatusARM) DeepCopyInto(out *BatchAccountIdentity_StatusARM) {
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
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]BatchAccountIdentity_Status_UserAssignedIdentitiesARM, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_StatusARM.
func (in *BatchAccountIdentity_StatusARM) DeepCopy() *BatchAccountIdentity_StatusARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity_Status_UserAssignedIdentities) DeepCopyInto(out *BatchAccountIdentity_Status_UserAssignedIdentities) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_Status_UserAssignedIdentities.
func (in *BatchAccountIdentity_Status_UserAssignedIdentities) DeepCopy() *BatchAccountIdentity_Status_UserAssignedIdentities {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_Status_UserAssignedIdentities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity_Status_UserAssignedIdentitiesARM) DeepCopyInto(out *BatchAccountIdentity_Status_UserAssignedIdentitiesARM) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_Status_UserAssignedIdentitiesARM.
func (in *BatchAccountIdentity_Status_UserAssignedIdentitiesARM) DeepCopy() *BatchAccountIdentity_Status_UserAssignedIdentitiesARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_Status_UserAssignedIdentitiesARM)
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
func (in *BatchAccountProperties_StatusARM) DeepCopyInto(out *BatchAccountProperties_StatusARM) {
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
		*out = new(AutoStorageProperties_StatusARM)
		**out = **in
	}
	if in.DedicatedCoreQuota != nil {
		in, out := &in.DedicatedCoreQuota, &out.DedicatedCoreQuota
		*out = new(int)
		**out = **in
	}
	if in.DedicatedCoreQuotaPerVMFamily != nil {
		in, out := &in.DedicatedCoreQuotaPerVMFamily, &out.DedicatedCoreQuotaPerVMFamily
		*out = make([]VirtualMachineFamilyCoreQuota_StatusARM, len(*in))
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
		*out = new(EncryptionProperties_StatusARM)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(KeyVaultReference_StatusARM)
		**out = **in
	}
	if in.LowPriorityCoreQuota != nil {
		in, out := &in.LowPriorityCoreQuota, &out.LowPriorityCoreQuota
		*out = new(int)
		**out = **in
	}
	if in.PoolAllocationMode != nil {
		in, out := &in.PoolAllocationMode, &out.PoolAllocationMode
		*out = new(PoolAllocationMode_Status)
		**out = **in
	}
	if in.PoolQuota != nil {
		in, out := &in.PoolQuota, &out.PoolQuota
		*out = new(int)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_StatusARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(BatchAccountPropertiesStatusProvisioningState)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(PublicNetworkAccessType_Status)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountProperties_StatusARM.
func (in *BatchAccountProperties_StatusARM) DeepCopy() *BatchAccountProperties_StatusARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccountProperties_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccount_Status) DeepCopyInto(out *BatchAccount_Status) {
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
		*out = new(AutoStorageProperties_Status)
		**out = **in
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
		*out = make([]VirtualMachineFamilyCoreQuota_Status, len(*in))
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
		*out = new(EncryptionProperties_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(BatchAccountIdentity_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(KeyVaultReference_Status)
		**out = **in
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
		*out = new(PoolAllocationMode_Status)
		**out = **in
	}
	if in.PoolQuota != nil {
		in, out := &in.PoolQuota, &out.PoolQuota
		*out = new(int)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(BatchAccountPropertiesStatusProvisioningState)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(PublicNetworkAccessType_Status)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccount_Status.
func (in *BatchAccount_Status) DeepCopy() *BatchAccount_Status {
	if in == nil {
		return nil
	}
	out := new(BatchAccount_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccount_StatusARM) DeepCopyInto(out *BatchAccount_StatusARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(BatchAccountIdentity_StatusARM)
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
		*out = new(BatchAccountProperties_StatusARM)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccount_StatusARM.
func (in *BatchAccount_StatusARM) DeepCopy() *BatchAccount_StatusARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccount_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccounts_Spec) DeepCopyInto(out *BatchAccounts_Spec) {
	*out = *in
	if in.AutoStorage != nil {
		in, out := &in.AutoStorage, &out.AutoStorage
		*out = new(AutoStorageBaseProperties)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(BatchAccountIdentity)
		**out = **in
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(KeyVaultReference)
		**out = **in
	}
	out.Owner = in.Owner
	if in.PoolAllocationMode != nil {
		in, out := &in.PoolAllocationMode, &out.PoolAllocationMode
		*out = new(BatchAccountCreatePropertiesPoolAllocationMode)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(BatchAccountCreatePropertiesPublicNetworkAccess)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccounts_Spec.
func (in *BatchAccounts_Spec) DeepCopy() *BatchAccounts_Spec {
	if in == nil {
		return nil
	}
	out := new(BatchAccounts_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccounts_SpecARM) DeepCopyInto(out *BatchAccounts_SpecARM) {
	*out = *in
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(BatchAccountIdentityARM)
		**out = **in
	}
	in.Properties.DeepCopyInto(&out.Properties)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccounts_SpecARM.
func (in *BatchAccounts_SpecARM) DeepCopy() *BatchAccounts_SpecARM {
	if in == nil {
		return nil
	}
	out := new(BatchAccounts_SpecARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionProperties) DeepCopyInto(out *EncryptionProperties) {
	*out = *in
	if in.KeySource != nil {
		in, out := &in.KeySource, &out.KeySource
		*out = new(EncryptionPropertiesKeySource)
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
		*out = new(EncryptionPropertiesKeySource)
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
func (in *EncryptionProperties_Status) DeepCopyInto(out *EncryptionProperties_Status) {
	*out = *in
	if in.KeySource != nil {
		in, out := &in.KeySource, &out.KeySource
		*out = new(EncryptionPropertiesStatusKeySource)
		**out = **in
	}
	if in.KeyVaultProperties != nil {
		in, out := &in.KeyVaultProperties, &out.KeyVaultProperties
		*out = new(KeyVaultProperties_Status)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionProperties_Status.
func (in *EncryptionProperties_Status) DeepCopy() *EncryptionProperties_Status {
	if in == nil {
		return nil
	}
	out := new(EncryptionProperties_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionProperties_StatusARM) DeepCopyInto(out *EncryptionProperties_StatusARM) {
	*out = *in
	if in.KeySource != nil {
		in, out := &in.KeySource, &out.KeySource
		*out = new(EncryptionPropertiesStatusKeySource)
		**out = **in
	}
	if in.KeyVaultProperties != nil {
		in, out := &in.KeyVaultProperties, &out.KeyVaultProperties
		*out = new(KeyVaultProperties_StatusARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionProperties_StatusARM.
func (in *EncryptionProperties_StatusARM) DeepCopy() *EncryptionProperties_StatusARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionProperties_StatusARM)
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
func (in *KeyVaultProperties_Status) DeepCopyInto(out *KeyVaultProperties_Status) {
	*out = *in
	if in.KeyIdentifier != nil {
		in, out := &in.KeyIdentifier, &out.KeyIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultProperties_Status.
func (in *KeyVaultProperties_Status) DeepCopy() *KeyVaultProperties_Status {
	if in == nil {
		return nil
	}
	out := new(KeyVaultProperties_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultProperties_StatusARM) DeepCopyInto(out *KeyVaultProperties_StatusARM) {
	*out = *in
	if in.KeyIdentifier != nil {
		in, out := &in.KeyIdentifier, &out.KeyIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultProperties_StatusARM.
func (in *KeyVaultProperties_StatusARM) DeepCopy() *KeyVaultProperties_StatusARM {
	if in == nil {
		return nil
	}
	out := new(KeyVaultProperties_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultReference) DeepCopyInto(out *KeyVaultReference) {
	*out = *in
	out.Reference = in.Reference
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
func (in *KeyVaultReference_Status) DeepCopyInto(out *KeyVaultReference_Status) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultReference_Status.
func (in *KeyVaultReference_Status) DeepCopy() *KeyVaultReference_Status {
	if in == nil {
		return nil
	}
	out := new(KeyVaultReference_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultReference_StatusARM) DeepCopyInto(out *KeyVaultReference_StatusARM) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultReference_StatusARM.
func (in *KeyVaultReference_StatusARM) DeepCopy() *KeyVaultReference_StatusARM {
	if in == nil {
		return nil
	}
	out := new(KeyVaultReference_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnectionProperties_StatusARM) DeepCopyInto(out *PrivateEndpointConnectionProperties_StatusARM) {
	*out = *in
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(PrivateEndpoint_StatusARM)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateLinkServiceConnectionState != nil {
		in, out := &in.PrivateLinkServiceConnectionState, &out.PrivateLinkServiceConnectionState
		*out = new(PrivateLinkServiceConnectionState_StatusARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(PrivateEndpointConnectionPropertiesStatusProvisioningState)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnectionProperties_StatusARM.
func (in *PrivateEndpointConnectionProperties_StatusARM) DeepCopy() *PrivateEndpointConnectionProperties_StatusARM {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnectionProperties_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_Status) DeepCopyInto(out *PrivateEndpointConnection_Status) {
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
		*out = new(PrivateEndpoint_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateLinkServiceConnectionState != nil {
		in, out := &in.PrivateLinkServiceConnectionState, &out.PrivateLinkServiceConnectionState
		*out = new(PrivateLinkServiceConnectionState_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(PrivateEndpointConnectionPropertiesStatusProvisioningState)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_Status.
func (in *PrivateEndpointConnection_Status) DeepCopy() *PrivateEndpointConnection_Status {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_StatusARM) DeepCopyInto(out *PrivateEndpointConnection_StatusARM) {
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
		*out = new(PrivateEndpointConnectionProperties_StatusARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_StatusARM.
func (in *PrivateEndpointConnection_StatusARM) DeepCopy() *PrivateEndpointConnection_StatusARM {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpoint_Status) DeepCopyInto(out *PrivateEndpoint_Status) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpoint_Status.
func (in *PrivateEndpoint_Status) DeepCopy() *PrivateEndpoint_Status {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpoint_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpoint_StatusARM) DeepCopyInto(out *PrivateEndpoint_StatusARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpoint_StatusARM.
func (in *PrivateEndpoint_StatusARM) DeepCopy() *PrivateEndpoint_StatusARM {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpoint_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkServiceConnectionState_Status) DeepCopyInto(out *PrivateLinkServiceConnectionState_Status) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateLinkServiceConnectionState_Status.
func (in *PrivateLinkServiceConnectionState_Status) DeepCopy() *PrivateLinkServiceConnectionState_Status {
	if in == nil {
		return nil
	}
	out := new(PrivateLinkServiceConnectionState_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkServiceConnectionState_StatusARM) DeepCopyInto(out *PrivateLinkServiceConnectionState_StatusARM) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateLinkServiceConnectionState_StatusARM.
func (in *PrivateLinkServiceConnectionState_StatusARM) DeepCopy() *PrivateLinkServiceConnectionState_StatusARM {
	if in == nil {
		return nil
	}
	out := new(PrivateLinkServiceConnectionState_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineFamilyCoreQuota_Status) DeepCopyInto(out *VirtualMachineFamilyCoreQuota_Status) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineFamilyCoreQuota_Status.
func (in *VirtualMachineFamilyCoreQuota_Status) DeepCopy() *VirtualMachineFamilyCoreQuota_Status {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineFamilyCoreQuota_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineFamilyCoreQuota_StatusARM) DeepCopyInto(out *VirtualMachineFamilyCoreQuota_StatusARM) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineFamilyCoreQuota_StatusARM.
func (in *VirtualMachineFamilyCoreQuota_StatusARM) DeepCopy() *VirtualMachineFamilyCoreQuota_StatusARM {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineFamilyCoreQuota_StatusARM)
	in.DeepCopyInto(out)
	return out
}
