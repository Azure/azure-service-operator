//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta20180601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationList) DeepCopyInto(out *ConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Configuration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationList.
func (in *ConfigurationList) DeepCopy() *ConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration_STATUS) DeepCopyInto(out *Configuration_STATUS) {
	*out = *in
	if in.AllowedValues != nil {
		in, out := &in.AllowedValues, &out.AllowedValues
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataType != nil {
		in, out := &in.DataType, &out.DataType
		*out = new(string)
		**out = **in
	}
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
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
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration_STATUS.
func (in *Configuration_STATUS) DeepCopy() *Configuration_STATUS {
	if in == nil {
		return nil
	}
	out := new(Configuration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Database) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseList) DeepCopyInto(out *DatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Database, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseList.
func (in *DatabaseList) DeepCopy() *DatabaseList {
	if in == nil {
		return nil
	}
	out := new(DatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database_STATUS) DeepCopyInto(out *Database_STATUS) {
	*out = *in
	if in.Charset != nil {
		in, out := &in.Charset, &out.Charset
		*out = new(string)
		**out = **in
	}
	if in.Collation != nil {
		in, out := &in.Collation, &out.Collation
		*out = new(string)
		**out = **in
	}
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database_STATUS.
func (in *Database_STATUS) DeepCopy() *Database_STATUS {
	if in == nil {
		return nil
	}
	out := new(Database_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointProperty_STATUS) DeepCopyInto(out *PrivateEndpointProperty_STATUS) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointProperty_STATUS.
func (in *PrivateEndpointProperty_STATUS) DeepCopy() *PrivateEndpointProperty_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointProperty_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Server) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerList) DeepCopyInto(out *ServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Server, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerList.
func (in *ServerList) DeepCopy() *ServerList {
	if in == nil {
		return nil
	}
	out := new(ServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerOperatorSecrets) DeepCopyInto(out *ServerOperatorSecrets) {
	*out = *in
	if in.FullyQualifiedDomainName != nil {
		in, out := &in.FullyQualifiedDomainName, &out.FullyQualifiedDomainName
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerOperatorSecrets.
func (in *ServerOperatorSecrets) DeepCopy() *ServerOperatorSecrets {
	if in == nil {
		return nil
	}
	out := new(ServerOperatorSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerOperatorSpec) DeepCopyInto(out *ServerOperatorSpec) {
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
		*out = new(ServerOperatorSecrets)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerOperatorSpec.
func (in *ServerOperatorSpec) DeepCopy() *ServerOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(ServerOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPrivateEndpointConnectionProperties_STATUS) DeepCopyInto(out *ServerPrivateEndpointConnectionProperties_STATUS) {
	*out = *in
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(PrivateEndpointProperty_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateLinkServiceConnectionState != nil {
		in, out := &in.PrivateLinkServiceConnectionState, &out.PrivateLinkServiceConnectionState
		*out = new(ServerPrivateLinkServiceConnectionStateProperty_STATUS)
		(*in).DeepCopyInto(*out)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPrivateEndpointConnectionProperties_STATUS.
func (in *ServerPrivateEndpointConnectionProperties_STATUS) DeepCopy() *ServerPrivateEndpointConnectionProperties_STATUS {
	if in == nil {
		return nil
	}
	out := new(ServerPrivateEndpointConnectionProperties_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPrivateEndpointConnection_STATUS) DeepCopyInto(out *ServerPrivateEndpointConnection_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ServerPrivateEndpointConnectionProperties_STATUS)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPrivateEndpointConnection_STATUS.
func (in *ServerPrivateEndpointConnection_STATUS) DeepCopy() *ServerPrivateEndpointConnection_STATUS {
	if in == nil {
		return nil
	}
	out := new(ServerPrivateEndpointConnection_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPrivateLinkServiceConnectionStateProperty_STATUS) DeepCopyInto(out *ServerPrivateLinkServiceConnectionStateProperty_STATUS) {
	*out = *in
	if in.ActionsRequired != nil {
		in, out := &in.ActionsRequired, &out.ActionsRequired
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPrivateLinkServiceConnectionStateProperty_STATUS.
func (in *ServerPrivateLinkServiceConnectionStateProperty_STATUS) DeepCopy() *ServerPrivateLinkServiceConnectionStateProperty_STATUS {
	if in == nil {
		return nil
	}
	out := new(ServerPrivateLinkServiceConnectionStateProperty_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPropertiesForCreate) DeepCopyInto(out *ServerPropertiesForCreate) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServerPropertiesForDefaultCreate != nil {
		in, out := &in.ServerPropertiesForDefaultCreate, &out.ServerPropertiesForDefaultCreate
		*out = new(ServerPropertiesForDefaultCreate)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerPropertiesForGeoRestore != nil {
		in, out := &in.ServerPropertiesForGeoRestore, &out.ServerPropertiesForGeoRestore
		*out = new(ServerPropertiesForGeoRestore)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerPropertiesForReplica != nil {
		in, out := &in.ServerPropertiesForReplica, &out.ServerPropertiesForReplica
		*out = new(ServerPropertiesForReplica)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerPropertiesForRestore != nil {
		in, out := &in.ServerPropertiesForRestore, &out.ServerPropertiesForRestore
		*out = new(ServerPropertiesForRestore)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPropertiesForCreate.
func (in *ServerPropertiesForCreate) DeepCopy() *ServerPropertiesForCreate {
	if in == nil {
		return nil
	}
	out := new(ServerPropertiesForCreate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPropertiesForDefaultCreate) DeepCopyInto(out *ServerPropertiesForDefaultCreate) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.AdministratorLoginPassword != nil {
		in, out := &in.AdministratorLoginPassword, &out.AdministratorLoginPassword
		*out = new(genruntime.SecretReference)
		**out = **in
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.MinimalTlsVersion != nil {
		in, out := &in.MinimalTlsVersion, &out.MinimalTlsVersion
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
	if in.SslEnforcement != nil {
		in, out := &in.SslEnforcement, &out.SslEnforcement
		*out = new(string)
		**out = **in
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(StorageProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPropertiesForDefaultCreate.
func (in *ServerPropertiesForDefaultCreate) DeepCopy() *ServerPropertiesForDefaultCreate {
	if in == nil {
		return nil
	}
	out := new(ServerPropertiesForDefaultCreate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPropertiesForGeoRestore) DeepCopyInto(out *ServerPropertiesForGeoRestore) {
	*out = *in
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.MinimalTlsVersion != nil {
		in, out := &in.MinimalTlsVersion, &out.MinimalTlsVersion
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
	if in.SourceServerId != nil {
		in, out := &in.SourceServerId, &out.SourceServerId
		*out = new(string)
		**out = **in
	}
	if in.SslEnforcement != nil {
		in, out := &in.SslEnforcement, &out.SslEnforcement
		*out = new(string)
		**out = **in
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(StorageProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPropertiesForGeoRestore.
func (in *ServerPropertiesForGeoRestore) DeepCopy() *ServerPropertiesForGeoRestore {
	if in == nil {
		return nil
	}
	out := new(ServerPropertiesForGeoRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPropertiesForReplica) DeepCopyInto(out *ServerPropertiesForReplica) {
	*out = *in
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.MinimalTlsVersion != nil {
		in, out := &in.MinimalTlsVersion, &out.MinimalTlsVersion
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
	if in.SourceServerId != nil {
		in, out := &in.SourceServerId, &out.SourceServerId
		*out = new(string)
		**out = **in
	}
	if in.SslEnforcement != nil {
		in, out := &in.SslEnforcement, &out.SslEnforcement
		*out = new(string)
		**out = **in
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(StorageProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPropertiesForReplica.
func (in *ServerPropertiesForReplica) DeepCopy() *ServerPropertiesForReplica {
	if in == nil {
		return nil
	}
	out := new(ServerPropertiesForReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPropertiesForRestore) DeepCopyInto(out *ServerPropertiesForRestore) {
	*out = *in
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.MinimalTlsVersion != nil {
		in, out := &in.MinimalTlsVersion, &out.MinimalTlsVersion
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
	if in.RestorePointInTime != nil {
		in, out := &in.RestorePointInTime, &out.RestorePointInTime
		*out = new(string)
		**out = **in
	}
	if in.SourceServerId != nil {
		in, out := &in.SourceServerId, &out.SourceServerId
		*out = new(string)
		**out = **in
	}
	if in.SslEnforcement != nil {
		in, out := &in.SslEnforcement, &out.SslEnforcement
		*out = new(string)
		**out = **in
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(StorageProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPropertiesForRestore.
func (in *ServerPropertiesForRestore) DeepCopy() *ServerPropertiesForRestore {
	if in == nil {
		return nil
	}
	out := new(ServerPropertiesForRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server_STATUS) DeepCopyInto(out *Server_STATUS) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EarliestRestoreDate != nil {
		in, out := &in.EarliestRestoreDate, &out.EarliestRestoreDate
		*out = new(string)
		**out = **in
	}
	if in.FullyQualifiedDomainName != nil {
		in, out := &in.FullyQualifiedDomainName, &out.FullyQualifiedDomainName
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
	if in.MasterServerId != nil {
		in, out := &in.MasterServerId, &out.MasterServerId
		*out = new(string)
		**out = **in
	}
	if in.MinimalTlsVersion != nil {
		in, out := &in.MinimalTlsVersion, &out.MinimalTlsVersion
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
		*out = make([]ServerPrivateEndpointConnection_STATUS, len(*in))
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
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.ReplicaCapacity != nil {
		in, out := &in.ReplicaCapacity, &out.ReplicaCapacity
		*out = new(int)
		**out = **in
	}
	if in.ReplicationRole != nil {
		in, out := &in.ReplicationRole, &out.ReplicationRole
		*out = new(string)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.SslEnforcement != nil {
		in, out := &in.SslEnforcement, &out.SslEnforcement
		*out = new(string)
		**out = **in
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(StorageProfile_STATUS)
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
	if in.UserVisibleState != nil {
		in, out := &in.UserVisibleState, &out.UserVisibleState
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server_STATUS.
func (in *Server_STATUS) DeepCopy() *Server_STATUS {
	if in == nil {
		return nil
	}
	out := new(Server_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServersConfigurations_Spec) DeepCopyInto(out *ServersConfigurations_Spec) {
	*out = *in
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
	if in.Source != nil {
		in, out := &in.Source, &out.Source
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
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServersConfigurations_Spec.
func (in *ServersConfigurations_Spec) DeepCopy() *ServersConfigurations_Spec {
	if in == nil {
		return nil
	}
	out := new(ServersConfigurations_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServersDatabases_Spec) DeepCopyInto(out *ServersDatabases_Spec) {
	*out = *in
	if in.Charset != nil {
		in, out := &in.Charset, &out.Charset
		*out = new(string)
		**out = **in
	}
	if in.Collation != nil {
		in, out := &in.Collation, &out.Collation
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
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServersDatabases_Spec.
func (in *ServersDatabases_Spec) DeepCopy() *ServersDatabases_Spec {
	if in == nil {
		return nil
	}
	out := new(ServersDatabases_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Servers_Spec) DeepCopyInto(out *Servers_Spec) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(ServerOperatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ServerPropertiesForCreate)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Servers_Spec.
func (in *Servers_Spec) DeepCopy() *Servers_Spec {
	if in == nil {
		return nil
	}
	out := new(Servers_Spec)
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
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
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
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
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
func (in *StorageProfile) DeepCopyInto(out *StorageProfile) {
	*out = *in
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(int)
		**out = **in
	}
	if in.GeoRedundantBackup != nil {
		in, out := &in.GeoRedundantBackup, &out.GeoRedundantBackup
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
	if in.StorageAutogrow != nil {
		in, out := &in.StorageAutogrow, &out.StorageAutogrow
		*out = new(string)
		**out = **in
	}
	if in.StorageMB != nil {
		in, out := &in.StorageMB, &out.StorageMB
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfile.
func (in *StorageProfile) DeepCopy() *StorageProfile {
	if in == nil {
		return nil
	}
	out := new(StorageProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfile_STATUS) DeepCopyInto(out *StorageProfile_STATUS) {
	*out = *in
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(int)
		**out = **in
	}
	if in.GeoRedundantBackup != nil {
		in, out := &in.GeoRedundantBackup, &out.GeoRedundantBackup
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
	if in.StorageAutogrow != nil {
		in, out := &in.StorageAutogrow, &out.StorageAutogrow
		*out = new(string)
		**out = **in
	}
	if in.StorageMB != nil {
		in, out := &in.StorageMB, &out.StorageMB
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfile_STATUS.
func (in *StorageProfile_STATUS) DeepCopy() *StorageProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(StorageProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}
