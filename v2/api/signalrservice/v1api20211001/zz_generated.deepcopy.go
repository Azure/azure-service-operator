//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20211001

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedIdentity) DeepCopyInto(out *ManagedIdentity) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ManagedIdentityType)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make([]UserAssignedIdentityDetails, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedIdentity.
func (in *ManagedIdentity) DeepCopy() *ManagedIdentity {
	if in == nil {
		return nil
	}
	out := new(ManagedIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedIdentitySettings) DeepCopyInto(out *ManagedIdentitySettings) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedIdentitySettings.
func (in *ManagedIdentitySettings) DeepCopy() *ManagedIdentitySettings {
	if in == nil {
		return nil
	}
	out := new(ManagedIdentitySettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedIdentitySettings_STATUS) DeepCopyInto(out *ManagedIdentitySettings_STATUS) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedIdentitySettings_STATUS.
func (in *ManagedIdentitySettings_STATUS) DeepCopy() *ManagedIdentitySettings_STATUS {
	if in == nil {
		return nil
	}
	out := new(ManagedIdentitySettings_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedIdentity_STATUS) DeepCopyInto(out *ManagedIdentity_STATUS) {
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
		*out = new(ManagedIdentityType_STATUS)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]UserAssignedIdentityProperty_STATUS, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedIdentity_STATUS.
func (in *ManagedIdentity_STATUS) DeepCopy() *ManagedIdentity_STATUS {
	if in == nil {
		return nil
	}
	out := new(ManagedIdentity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkACL) DeepCopyInto(out *NetworkACL) {
	*out = *in
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = make([]SignalRRequestType, len(*in))
		copy(*out, *in)
	}
	if in.Deny != nil {
		in, out := &in.Deny, &out.Deny
		*out = make([]SignalRRequestType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkACL.
func (in *NetworkACL) DeepCopy() *NetworkACL {
	if in == nil {
		return nil
	}
	out := new(NetworkACL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkACL_STATUS) DeepCopyInto(out *NetworkACL_STATUS) {
	*out = *in
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = make([]SignalRRequestType_STATUS, len(*in))
		copy(*out, *in)
	}
	if in.Deny != nil {
		in, out := &in.Deny, &out.Deny
		*out = make([]SignalRRequestType_STATUS, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkACL_STATUS.
func (in *NetworkACL_STATUS) DeepCopy() *NetworkACL_STATUS {
	if in == nil {
		return nil
	}
	out := new(NetworkACL_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointACL) DeepCopyInto(out *PrivateEndpointACL) {
	*out = *in
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = make([]SignalRRequestType, len(*in))
		copy(*out, *in)
	}
	if in.Deny != nil {
		in, out := &in.Deny, &out.Deny
		*out = make([]SignalRRequestType, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointACL.
func (in *PrivateEndpointACL) DeepCopy() *PrivateEndpointACL {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointACL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointACL_STATUS) DeepCopyInto(out *PrivateEndpointACL_STATUS) {
	*out = *in
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = make([]SignalRRequestType_STATUS, len(*in))
		copy(*out, *in)
	}
	if in.Deny != nil {
		in, out := &in.Deny, &out.Deny
		*out = make([]SignalRRequestType_STATUS, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointACL_STATUS.
func (in *PrivateEndpointACL_STATUS) DeepCopy() *PrivateEndpointACL_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointACL_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded) DeepCopyInto(out *PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded.
func (in *PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded) DeepCopy() *PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLogCategory) DeepCopyInto(out *ResourceLogCategory) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLogCategory.
func (in *ResourceLogCategory) DeepCopy() *ResourceLogCategory {
	if in == nil {
		return nil
	}
	out := new(ResourceLogCategory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLogCategory_STATUS) DeepCopyInto(out *ResourceLogCategory_STATUS) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLogCategory_STATUS.
func (in *ResourceLogCategory_STATUS) DeepCopy() *ResourceLogCategory_STATUS {
	if in == nil {
		return nil
	}
	out := new(ResourceLogCategory_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLogConfiguration) DeepCopyInto(out *ResourceLogConfiguration) {
	*out = *in
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]ResourceLogCategory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLogConfiguration.
func (in *ResourceLogConfiguration) DeepCopy() *ResourceLogConfiguration {
	if in == nil {
		return nil
	}
	out := new(ResourceLogConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLogConfiguration_STATUS) DeepCopyInto(out *ResourceLogConfiguration_STATUS) {
	*out = *in
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]ResourceLogCategory_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLogConfiguration_STATUS.
func (in *ResourceLogConfiguration_STATUS) DeepCopy() *ResourceLogConfiguration_STATUS {
	if in == nil {
		return nil
	}
	out := new(ResourceLogConfiguration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSku) DeepCopyInto(out *ResourceSku) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(SignalRSkuTier)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSku.
func (in *ResourceSku) DeepCopy() *ResourceSku {
	if in == nil {
		return nil
	}
	out := new(ResourceSku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSku_STATUS) DeepCopyInto(out *ResourceSku_STATUS) {
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
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(SignalRSkuTier_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSku_STATUS.
func (in *ResourceSku_STATUS) DeepCopy() *ResourceSku_STATUS {
	if in == nil {
		return nil
	}
	out := new(ResourceSku_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessUpstreamSettings) DeepCopyInto(out *ServerlessUpstreamSettings) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]UpstreamTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessUpstreamSettings.
func (in *ServerlessUpstreamSettings) DeepCopy() *ServerlessUpstreamSettings {
	if in == nil {
		return nil
	}
	out := new(ServerlessUpstreamSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessUpstreamSettings_STATUS) DeepCopyInto(out *ServerlessUpstreamSettings_STATUS) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]UpstreamTemplate_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessUpstreamSettings_STATUS.
func (in *ServerlessUpstreamSettings_STATUS) DeepCopy() *ServerlessUpstreamSettings_STATUS {
	if in == nil {
		return nil
	}
	out := new(ServerlessUpstreamSettings_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded) DeepCopyInto(out *SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded.
func (in *SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded) DeepCopy() *SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded {
	if in == nil {
		return nil
	}
	out := new(SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalR) DeepCopyInto(out *SignalR) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalR.
func (in *SignalR) DeepCopy() *SignalR {
	if in == nil {
		return nil
	}
	out := new(SignalR)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SignalR) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalRCorsSettings) DeepCopyInto(out *SignalRCorsSettings) {
	*out = *in
	if in.AllowedOrigins != nil {
		in, out := &in.AllowedOrigins, &out.AllowedOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalRCorsSettings.
func (in *SignalRCorsSettings) DeepCopy() *SignalRCorsSettings {
	if in == nil {
		return nil
	}
	out := new(SignalRCorsSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalRCorsSettings_STATUS) DeepCopyInto(out *SignalRCorsSettings_STATUS) {
	*out = *in
	if in.AllowedOrigins != nil {
		in, out := &in.AllowedOrigins, &out.AllowedOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalRCorsSettings_STATUS.
func (in *SignalRCorsSettings_STATUS) DeepCopy() *SignalRCorsSettings_STATUS {
	if in == nil {
		return nil
	}
	out := new(SignalRCorsSettings_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalRFeature) DeepCopyInto(out *SignalRFeature) {
	*out = *in
	if in.Flag != nil {
		in, out := &in.Flag, &out.Flag
		*out = new(FeatureFlags)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalRFeature.
func (in *SignalRFeature) DeepCopy() *SignalRFeature {
	if in == nil {
		return nil
	}
	out := new(SignalRFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalRFeature_STATUS) DeepCopyInto(out *SignalRFeature_STATUS) {
	*out = *in
	if in.Flag != nil {
		in, out := &in.Flag, &out.Flag
		*out = new(FeatureFlags_STATUS)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalRFeature_STATUS.
func (in *SignalRFeature_STATUS) DeepCopy() *SignalRFeature_STATUS {
	if in == nil {
		return nil
	}
	out := new(SignalRFeature_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalRList) DeepCopyInto(out *SignalRList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SignalR, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalRList.
func (in *SignalRList) DeepCopy() *SignalRList {
	if in == nil {
		return nil
	}
	out := new(SignalRList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SignalRList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalRNetworkACLs) DeepCopyInto(out *SignalRNetworkACLs) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(ACLAction)
		**out = **in
	}
	if in.PrivateEndpoints != nil {
		in, out := &in.PrivateEndpoints, &out.PrivateEndpoints
		*out = make([]PrivateEndpointACL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicNetwork != nil {
		in, out := &in.PublicNetwork, &out.PublicNetwork
		*out = new(NetworkACL)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalRNetworkACLs.
func (in *SignalRNetworkACLs) DeepCopy() *SignalRNetworkACLs {
	if in == nil {
		return nil
	}
	out := new(SignalRNetworkACLs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalRNetworkACLs_STATUS) DeepCopyInto(out *SignalRNetworkACLs_STATUS) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(ACLAction_STATUS)
		**out = **in
	}
	if in.PrivateEndpoints != nil {
		in, out := &in.PrivateEndpoints, &out.PrivateEndpoints
		*out = make([]PrivateEndpointACL_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicNetwork != nil {
		in, out := &in.PublicNetwork, &out.PublicNetwork
		*out = new(NetworkACL_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalRNetworkACLs_STATUS.
func (in *SignalRNetworkACLs_STATUS) DeepCopy() *SignalRNetworkACLs_STATUS {
	if in == nil {
		return nil
	}
	out := new(SignalRNetworkACLs_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalROperatorSecrets) DeepCopyInto(out *SignalROperatorSecrets) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalROperatorSecrets.
func (in *SignalROperatorSecrets) DeepCopy() *SignalROperatorSecrets {
	if in == nil {
		return nil
	}
	out := new(SignalROperatorSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalROperatorSpec) DeepCopyInto(out *SignalROperatorSpec) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(SignalROperatorSecrets)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalROperatorSpec.
func (in *SignalROperatorSpec) DeepCopy() *SignalROperatorSpec {
	if in == nil {
		return nil
	}
	out := new(SignalROperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalRTlsSettings) DeepCopyInto(out *SignalRTlsSettings) {
	*out = *in
	if in.ClientCertEnabled != nil {
		in, out := &in.ClientCertEnabled, &out.ClientCertEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalRTlsSettings.
func (in *SignalRTlsSettings) DeepCopy() *SignalRTlsSettings {
	if in == nil {
		return nil
	}
	out := new(SignalRTlsSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalRTlsSettings_STATUS) DeepCopyInto(out *SignalRTlsSettings_STATUS) {
	*out = *in
	if in.ClientCertEnabled != nil {
		in, out := &in.ClientCertEnabled, &out.ClientCertEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalRTlsSettings_STATUS.
func (in *SignalRTlsSettings_STATUS) DeepCopy() *SignalRTlsSettings_STATUS {
	if in == nil {
		return nil
	}
	out := new(SignalRTlsSettings_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalR_STATUS) DeepCopyInto(out *SignalR_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = new(SignalRCorsSettings_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.DisableAadAuth != nil {
		in, out := &in.DisableAadAuth, &out.DisableAadAuth
		*out = new(bool)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.ExternalIP != nil {
		in, out := &in.ExternalIP, &out.ExternalIP
		*out = new(string)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]SignalRFeature_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.HostNamePrefix != nil {
		in, out := &in.HostNamePrefix, &out.HostNamePrefix
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
		*out = new(ManagedIdentity_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(ServiceKind_STATUS)
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
	if in.NetworkACLs != nil {
		in, out := &in.NetworkACLs, &out.NetworkACLs
		*out = new(SignalRNetworkACLs_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.PublicPort != nil {
		in, out := &in.PublicPort, &out.PublicPort
		*out = new(int)
		**out = **in
	}
	if in.ResourceLogConfiguration != nil {
		in, out := &in.ResourceLogConfiguration, &out.ResourceLogConfiguration
		*out = new(ResourceLogConfiguration_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerPort != nil {
		in, out := &in.ServerPort, &out.ServerPort
		*out = new(int)
		**out = **in
	}
	if in.SharedPrivateLinkResources != nil {
		in, out := &in.SharedPrivateLinkResources, &out.SharedPrivateLinkResources
		*out = make([]SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(ResourceSku_STATUS)
		(*in).DeepCopyInto(*out)
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
	if in.Tls != nil {
		in, out := &in.Tls, &out.Tls
		*out = new(SignalRTlsSettings_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		*out = new(ServerlessUpstreamSettings_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalR_STATUS.
func (in *SignalR_STATUS) DeepCopy() *SignalR_STATUS {
	if in == nil {
		return nil
	}
	out := new(SignalR_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalR_Spec) DeepCopyInto(out *SignalR_Spec) {
	*out = *in
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = new(SignalRCorsSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.DisableAadAuth != nil {
		in, out := &in.DisableAadAuth, &out.DisableAadAuth
		*out = new(bool)
		**out = **in
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]SignalRFeature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(ManagedIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(ServiceKind)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.NetworkACLs != nil {
		in, out := &in.NetworkACLs, &out.NetworkACLs
		*out = new(SignalRNetworkACLs)
		(*in).DeepCopyInto(*out)
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(SignalROperatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.ResourceLogConfiguration != nil {
		in, out := &in.ResourceLogConfiguration, &out.ResourceLogConfiguration
		*out = new(ResourceLogConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(ResourceSku)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tls != nil {
		in, out := &in.Tls, &out.Tls
		*out = new(SignalRTlsSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		*out = new(ServerlessUpstreamSettings)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalR_Spec.
func (in *SignalR_Spec) DeepCopy() *SignalR_Spec {
	if in == nil {
		return nil
	}
	out := new(SignalR_Spec)
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
func (in *UpstreamAuthSettings) DeepCopyInto(out *UpstreamAuthSettings) {
	*out = *in
	if in.ManagedIdentity != nil {
		in, out := &in.ManagedIdentity, &out.ManagedIdentity
		*out = new(ManagedIdentitySettings)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(UpstreamAuthType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamAuthSettings.
func (in *UpstreamAuthSettings) DeepCopy() *UpstreamAuthSettings {
	if in == nil {
		return nil
	}
	out := new(UpstreamAuthSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamAuthSettings_STATUS) DeepCopyInto(out *UpstreamAuthSettings_STATUS) {
	*out = *in
	if in.ManagedIdentity != nil {
		in, out := &in.ManagedIdentity, &out.ManagedIdentity
		*out = new(ManagedIdentitySettings_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(UpstreamAuthType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamAuthSettings_STATUS.
func (in *UpstreamAuthSettings_STATUS) DeepCopy() *UpstreamAuthSettings_STATUS {
	if in == nil {
		return nil
	}
	out := new(UpstreamAuthSettings_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamTemplate) DeepCopyInto(out *UpstreamTemplate) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(UpstreamAuthSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.CategoryPattern != nil {
		in, out := &in.CategoryPattern, &out.CategoryPattern
		*out = new(string)
		**out = **in
	}
	if in.EventPattern != nil {
		in, out := &in.EventPattern, &out.EventPattern
		*out = new(string)
		**out = **in
	}
	if in.HubPattern != nil {
		in, out := &in.HubPattern, &out.HubPattern
		*out = new(string)
		**out = **in
	}
	if in.UrlTemplate != nil {
		in, out := &in.UrlTemplate, &out.UrlTemplate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamTemplate.
func (in *UpstreamTemplate) DeepCopy() *UpstreamTemplate {
	if in == nil {
		return nil
	}
	out := new(UpstreamTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamTemplate_STATUS) DeepCopyInto(out *UpstreamTemplate_STATUS) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(UpstreamAuthSettings_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.CategoryPattern != nil {
		in, out := &in.CategoryPattern, &out.CategoryPattern
		*out = new(string)
		**out = **in
	}
	if in.EventPattern != nil {
		in, out := &in.EventPattern, &out.EventPattern
		*out = new(string)
		**out = **in
	}
	if in.HubPattern != nil {
		in, out := &in.HubPattern, &out.HubPattern
		*out = new(string)
		**out = **in
	}
	if in.UrlTemplate != nil {
		in, out := &in.UrlTemplate, &out.UrlTemplate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamTemplate_STATUS.
func (in *UpstreamTemplate_STATUS) DeepCopy() *UpstreamTemplate_STATUS {
	if in == nil {
		return nil
	}
	out := new(UpstreamTemplate_STATUS)
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
func (in *UserAssignedIdentityProperty_STATUS) DeepCopyInto(out *UserAssignedIdentityProperty_STATUS) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAssignedIdentityProperty_STATUS.
func (in *UserAssignedIdentityProperty_STATUS) DeepCopy() *UserAssignedIdentityProperty_STATUS {
	if in == nil {
		return nil
	}
	out := new(UserAssignedIdentityProperty_STATUS)
	in.DeepCopyInto(out)
	return out
}
