//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta20220501

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStore) DeepCopyInto(out *ConfigurationStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStore.
func (in *ConfigurationStore) DeepCopy() *ConfigurationStore {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStoreList) DeepCopyInto(out *ConfigurationStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigurationStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStoreList.
func (in *ConfigurationStoreList) DeepCopy() *ConfigurationStoreList {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStoreOperatorSecrets) DeepCopyInto(out *ConfigurationStoreOperatorSecrets) {
	*out = *in
	if in.PrimaryConnectionString != nil {
		in, out := &in.PrimaryConnectionString, &out.PrimaryConnectionString
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.PrimaryKey != nil {
		in, out := &in.PrimaryKey, &out.PrimaryKey
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.PrimaryKeyID != nil {
		in, out := &in.PrimaryKeyID, &out.PrimaryKeyID
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.PrimaryReadOnlyConnectionString != nil {
		in, out := &in.PrimaryReadOnlyConnectionString, &out.PrimaryReadOnlyConnectionString
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.PrimaryReadOnlyKey != nil {
		in, out := &in.PrimaryReadOnlyKey, &out.PrimaryReadOnlyKey
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.PrimaryReadOnlyKeyID != nil {
		in, out := &in.PrimaryReadOnlyKeyID, &out.PrimaryReadOnlyKeyID
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.SecondaryConnectionString != nil {
		in, out := &in.SecondaryConnectionString, &out.SecondaryConnectionString
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.SecondaryKey != nil {
		in, out := &in.SecondaryKey, &out.SecondaryKey
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.SecondaryKeyID != nil {
		in, out := &in.SecondaryKeyID, &out.SecondaryKeyID
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.SecondaryReadOnlyConnectionString != nil {
		in, out := &in.SecondaryReadOnlyConnectionString, &out.SecondaryReadOnlyConnectionString
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.SecondaryReadOnlyKey != nil {
		in, out := &in.SecondaryReadOnlyKey, &out.SecondaryReadOnlyKey
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.SecondaryReadOnlyKeyID != nil {
		in, out := &in.SecondaryReadOnlyKeyID, &out.SecondaryReadOnlyKeyID
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStoreOperatorSecrets.
func (in *ConfigurationStoreOperatorSecrets) DeepCopy() *ConfigurationStoreOperatorSecrets {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStoreOperatorSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStoreOperatorSpec) DeepCopyInto(out *ConfigurationStoreOperatorSpec) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(ConfigurationStoreOperatorSecrets)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStoreOperatorSpec.
func (in *ConfigurationStoreOperatorSpec) DeepCopy() *ConfigurationStoreOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStoreOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStoreProperties_ARM) DeepCopyInto(out *ConfigurationStoreProperties_ARM) {
	*out = *in
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(ConfigurationStoreProperties_CreateMode)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnablePurgeProtection != nil {
		in, out := &in.EnablePurgeProtection, &out.EnablePurgeProtection
		*out = new(bool)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionProperties_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(ConfigurationStoreProperties_PublicNetworkAccess)
		**out = **in
	}
	if in.SoftDeleteRetentionInDays != nil {
		in, out := &in.SoftDeleteRetentionInDays, &out.SoftDeleteRetentionInDays
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStoreProperties_ARM.
func (in *ConfigurationStoreProperties_ARM) DeepCopy() *ConfigurationStoreProperties_ARM {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStoreProperties_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStoreProperties_STATUS_ARM) DeepCopyInto(out *ConfigurationStoreProperties_STATUS_ARM) {
	*out = *in
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(ConfigurationStoreProperties_CreateMode_STATUS)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = new(string)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnablePurgeProtection != nil {
		in, out := &in.EnablePurgeProtection, &out.EnablePurgeProtection
		*out = new(bool)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionProperties_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnectionReference_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ConfigurationStoreProperties_ProvisioningState_STATUS)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(ConfigurationStoreProperties_PublicNetworkAccess_STATUS)
		**out = **in
	}
	if in.SoftDeleteRetentionInDays != nil {
		in, out := &in.SoftDeleteRetentionInDays, &out.SoftDeleteRetentionInDays
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStoreProperties_STATUS_ARM.
func (in *ConfigurationStoreProperties_STATUS_ARM) DeepCopy() *ConfigurationStoreProperties_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStoreProperties_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStore_STATUS) DeepCopyInto(out *ConfigurationStore_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(ConfigurationStoreProperties_CreateMode_STATUS)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = new(string)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnablePurgeProtection != nil {
		in, out := &in.EnablePurgeProtection, &out.EnablePurgeProtection
		*out = new(bool)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionProperties_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(ResourceIdentity_STATUS)
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
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnectionReference_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ConfigurationStoreProperties_ProvisioningState_STATUS)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(ConfigurationStoreProperties_PublicNetworkAccess_STATUS)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.SoftDeleteRetentionInDays != nil {
		in, out := &in.SoftDeleteRetentionInDays, &out.SoftDeleteRetentionInDays
		*out = new(int)
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStore_STATUS.
func (in *ConfigurationStore_STATUS) DeepCopy() *ConfigurationStore_STATUS {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStore_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStore_STATUS_ARM) DeepCopyInto(out *ConfigurationStore_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(ResourceIdentity_STATUS_ARM)
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
		*out = new(ConfigurationStoreProperties_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS_ARM)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStore_STATUS_ARM.
func (in *ConfigurationStore_STATUS_ARM) DeepCopy() *ConfigurationStore_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStore_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStore_Spec) DeepCopyInto(out *ConfigurationStore_Spec) {
	*out = *in
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(ConfigurationStoreProperties_CreateMode)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EnablePurgeProtection != nil {
		in, out := &in.EnablePurgeProtection, &out.EnablePurgeProtection
		*out = new(bool)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(ResourceIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(ConfigurationStoreOperatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(ConfigurationStoreProperties_PublicNetworkAccess)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.SoftDeleteRetentionInDays != nil {
		in, out := &in.SoftDeleteRetentionInDays, &out.SoftDeleteRetentionInDays
		*out = new(int)
		**out = **in
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStore_Spec.
func (in *ConfigurationStore_Spec) DeepCopy() *ConfigurationStore_Spec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStore_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStore_Spec_ARM) DeepCopyInto(out *ConfigurationStore_Spec_ARM) {
	*out = *in
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(ResourceIdentity_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ConfigurationStoreProperties_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_ARM)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStore_Spec_ARM.
func (in *ConfigurationStore_Spec_ARM) DeepCopy() *ConfigurationStore_Spec_ARM {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStore_Spec_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionProperties) DeepCopyInto(out *EncryptionProperties) {
	*out = *in
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
func (in *EncryptionProperties_ARM) DeepCopyInto(out *EncryptionProperties_ARM) {
	*out = *in
	if in.KeyVaultProperties != nil {
		in, out := &in.KeyVaultProperties, &out.KeyVaultProperties
		*out = new(KeyVaultProperties_ARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionProperties_ARM.
func (in *EncryptionProperties_ARM) DeepCopy() *EncryptionProperties_ARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionProperties_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionProperties_STATUS) DeepCopyInto(out *EncryptionProperties_STATUS) {
	*out = *in
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
func (in *EncryptionProperties_STATUS_ARM) DeepCopyInto(out *EncryptionProperties_STATUS_ARM) {
	*out = *in
	if in.KeyVaultProperties != nil {
		in, out := &in.KeyVaultProperties, &out.KeyVaultProperties
		*out = new(KeyVaultProperties_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionProperties_STATUS_ARM.
func (in *EncryptionProperties_STATUS_ARM) DeepCopy() *EncryptionProperties_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionProperties_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultProperties) DeepCopyInto(out *KeyVaultProperties) {
	*out = *in
	if in.IdentityClientId != nil {
		in, out := &in.IdentityClientId, &out.IdentityClientId
		*out = new(string)
		**out = **in
	}
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
func (in *KeyVaultProperties_ARM) DeepCopyInto(out *KeyVaultProperties_ARM) {
	*out = *in
	if in.IdentityClientId != nil {
		in, out := &in.IdentityClientId, &out.IdentityClientId
		*out = new(string)
		**out = **in
	}
	if in.KeyIdentifier != nil {
		in, out := &in.KeyIdentifier, &out.KeyIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultProperties_ARM.
func (in *KeyVaultProperties_ARM) DeepCopy() *KeyVaultProperties_ARM {
	if in == nil {
		return nil
	}
	out := new(KeyVaultProperties_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultProperties_STATUS) DeepCopyInto(out *KeyVaultProperties_STATUS) {
	*out = *in
	if in.IdentityClientId != nil {
		in, out := &in.IdentityClientId, &out.IdentityClientId
		*out = new(string)
		**out = **in
	}
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
func (in *KeyVaultProperties_STATUS_ARM) DeepCopyInto(out *KeyVaultProperties_STATUS_ARM) {
	*out = *in
	if in.IdentityClientId != nil {
		in, out := &in.IdentityClientId, &out.IdentityClientId
		*out = new(string)
		**out = **in
	}
	if in.KeyIdentifier != nil {
		in, out := &in.KeyIdentifier, &out.KeyIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultProperties_STATUS_ARM.
func (in *KeyVaultProperties_STATUS_ARM) DeepCopy() *KeyVaultProperties_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(KeyVaultProperties_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnectionReference_STATUS) DeepCopyInto(out *PrivateEndpointConnectionReference_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnectionReference_STATUS.
func (in *PrivateEndpointConnectionReference_STATUS) DeepCopy() *PrivateEndpointConnectionReference_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnectionReference_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnectionReference_STATUS_ARM) DeepCopyInto(out *PrivateEndpointConnectionReference_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnectionReference_STATUS_ARM.
func (in *PrivateEndpointConnectionReference_STATUS_ARM) DeepCopy() *PrivateEndpointConnectionReference_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnectionReference_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceIdentity) DeepCopyInto(out *ResourceIdentity) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ResourceIdentity_Type)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make([]UserAssignedIdentityDetails, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceIdentity.
func (in *ResourceIdentity) DeepCopy() *ResourceIdentity {
	if in == nil {
		return nil
	}
	out := new(ResourceIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceIdentity_ARM) DeepCopyInto(out *ResourceIdentity_ARM) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ResourceIdentity_Type)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]UserAssignedIdentityDetails_ARM, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceIdentity_ARM.
func (in *ResourceIdentity_ARM) DeepCopy() *ResourceIdentity_ARM {
	if in == nil {
		return nil
	}
	out := new(ResourceIdentity_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceIdentity_STATUS) DeepCopyInto(out *ResourceIdentity_STATUS) {
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
		*out = new(ResourceIdentity_Type_STATUS)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]UserIdentity_STATUS, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceIdentity_STATUS.
func (in *ResourceIdentity_STATUS) DeepCopy() *ResourceIdentity_STATUS {
	if in == nil {
		return nil
	}
	out := new(ResourceIdentity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceIdentity_STATUS_ARM) DeepCopyInto(out *ResourceIdentity_STATUS_ARM) {
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
		*out = new(ResourceIdentity_Type_STATUS)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]UserIdentity_STATUS_ARM, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceIdentity_STATUS_ARM.
func (in *ResourceIdentity_STATUS_ARM) DeepCopy() *ResourceIdentity_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(ResourceIdentity_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku) DeepCopyInto(out *Sku) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
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
func (in *Sku_ARM) DeepCopyInto(out *Sku_ARM) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_ARM.
func (in *Sku_ARM) DeepCopy() *Sku_ARM {
	if in == nil {
		return nil
	}
	out := new(Sku_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_STATUS) DeepCopyInto(out *Sku_STATUS) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_STATUS_ARM) DeepCopyInto(out *Sku_STATUS_ARM) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_STATUS_ARM.
func (in *Sku_STATUS_ARM) DeepCopy() *Sku_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(Sku_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData) DeepCopyInto(out *SystemData) {
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
		*out = new(SystemData_CreatedByType)
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
		*out = new(SystemData_LastModifiedByType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData.
func (in *SystemData) DeepCopy() *SystemData {
	if in == nil {
		return nil
	}
	out := new(SystemData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_ARM) DeepCopyInto(out *SystemData_ARM) {
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
		*out = new(SystemData_CreatedByType)
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
		*out = new(SystemData_LastModifiedByType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_ARM.
func (in *SystemData_ARM) DeepCopy() *SystemData_ARM {
	if in == nil {
		return nil
	}
	out := new(SystemData_ARM)
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
func (in *SystemData_STATUS_ARM) DeepCopyInto(out *SystemData_STATUS_ARM) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS_ARM.
func (in *SystemData_STATUS_ARM) DeepCopy() *SystemData_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAssignedIdentityDetails) DeepCopyInto(out *UserAssignedIdentityDetails) {
	*out = *in
	out.Reference = in.Reference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAssignedIdentityDetails.
func (in *UserAssignedIdentityDetails) DeepCopy() *UserAssignedIdentityDetails {
	if in == nil {
		return nil
	}
	out := new(UserAssignedIdentityDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAssignedIdentityDetails_ARM) DeepCopyInto(out *UserAssignedIdentityDetails_ARM) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAssignedIdentityDetails_ARM.
func (in *UserAssignedIdentityDetails_ARM) DeepCopy() *UserAssignedIdentityDetails_ARM {
	if in == nil {
		return nil
	}
	out := new(UserAssignedIdentityDetails_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserIdentity_STATUS) DeepCopyInto(out *UserIdentity_STATUS) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserIdentity_STATUS.
func (in *UserIdentity_STATUS) DeepCopy() *UserIdentity_STATUS {
	if in == nil {
		return nil
	}
	out := new(UserIdentity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserIdentity_STATUS_ARM) DeepCopyInto(out *UserIdentity_STATUS_ARM) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserIdentity_STATUS_ARM.
func (in *UserIdentity_STATUS_ARM) DeepCopy() *UserIdentity_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(UserIdentity_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}
