//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20210501

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(int)
		**out = **in
	}
	if in.GeoRedundantBackup != nil {
		in, out := &in.GeoRedundantBackup, &out.GeoRedundantBackup
		*out = new(EnableStatusEnum)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup_STATUS) DeepCopyInto(out *Backup_STATUS) {
	*out = *in
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(int)
		**out = **in
	}
	if in.EarliestRestoreDate != nil {
		in, out := &in.EarliestRestoreDate, &out.EarliestRestoreDate
		*out = new(string)
		**out = **in
	}
	if in.GeoRedundantBackup != nil {
		in, out := &in.GeoRedundantBackup, &out.GeoRedundantBackup
		*out = new(EnableStatusEnum_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup_STATUS.
func (in *Backup_STATUS) DeepCopy() *Backup_STATUS {
	if in == nil {
		return nil
	}
	out := new(Backup_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataEncryption) DeepCopyInto(out *DataEncryption) {
	*out = *in
	if in.GeoBackupKeyURI != nil {
		in, out := &in.GeoBackupKeyURI, &out.GeoBackupKeyURI
		*out = new(string)
		**out = **in
	}
	if in.GeoBackupUserAssignedIdentityReference != nil {
		in, out := &in.GeoBackupUserAssignedIdentityReference, &out.GeoBackupUserAssignedIdentityReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.PrimaryKeyURI != nil {
		in, out := &in.PrimaryKeyURI, &out.PrimaryKeyURI
		*out = new(string)
		**out = **in
	}
	if in.PrimaryUserAssignedIdentityReference != nil {
		in, out := &in.PrimaryUserAssignedIdentityReference, &out.PrimaryUserAssignedIdentityReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(DataEncryption_Type)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataEncryption.
func (in *DataEncryption) DeepCopy() *DataEncryption {
	if in == nil {
		return nil
	}
	out := new(DataEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataEncryption_STATUS) DeepCopyInto(out *DataEncryption_STATUS) {
	*out = *in
	if in.GeoBackupKeyURI != nil {
		in, out := &in.GeoBackupKeyURI, &out.GeoBackupKeyURI
		*out = new(string)
		**out = **in
	}
	if in.GeoBackupUserAssignedIdentityId != nil {
		in, out := &in.GeoBackupUserAssignedIdentityId, &out.GeoBackupUserAssignedIdentityId
		*out = new(string)
		**out = **in
	}
	if in.PrimaryKeyURI != nil {
		in, out := &in.PrimaryKeyURI, &out.PrimaryKeyURI
		*out = new(string)
		**out = **in
	}
	if in.PrimaryUserAssignedIdentityId != nil {
		in, out := &in.PrimaryUserAssignedIdentityId, &out.PrimaryUserAssignedIdentityId
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(DataEncryption_Type_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataEncryption_STATUS.
func (in *DataEncryption_STATUS) DeepCopy() *DataEncryption_STATUS {
	if in == nil {
		return nil
	}
	out := new(DataEncryption_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServer) DeepCopyInto(out *FlexibleServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServer.
func (in *FlexibleServer) DeepCopy() *FlexibleServer {
	if in == nil {
		return nil
	}
	out := new(FlexibleServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerList) DeepCopyInto(out *FlexibleServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlexibleServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerList.
func (in *FlexibleServerList) DeepCopy() *FlexibleServerList {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerOperatorConfigMaps) DeepCopyInto(out *FlexibleServerOperatorConfigMaps) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(genruntime.ConfigMapDestination)
		**out = **in
	}
	if in.FullyQualifiedDomainName != nil {
		in, out := &in.FullyQualifiedDomainName, &out.FullyQualifiedDomainName
		*out = new(genruntime.ConfigMapDestination)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerOperatorConfigMaps.
func (in *FlexibleServerOperatorConfigMaps) DeepCopy() *FlexibleServerOperatorConfigMaps {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerOperatorConfigMaps)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerOperatorSecrets) DeepCopyInto(out *FlexibleServerOperatorSecrets) {
	*out = *in
	if in.FullyQualifiedDomainName != nil {
		in, out := &in.FullyQualifiedDomainName, &out.FullyQualifiedDomainName
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerOperatorSecrets.
func (in *FlexibleServerOperatorSecrets) DeepCopy() *FlexibleServerOperatorSecrets {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerOperatorSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerOperatorSpec) DeepCopyInto(out *FlexibleServerOperatorSpec) {
	*out = *in
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = new(FlexibleServerOperatorConfigMaps)
		(*in).DeepCopyInto(*out)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(FlexibleServerOperatorSecrets)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerOperatorSpec.
func (in *FlexibleServerOperatorSpec) DeepCopy() *FlexibleServerOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServer_STATUS) DeepCopyInto(out *FlexibleServer_STATUS) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(Backup_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(ServerProperties_CreateMode_STATUS)
		**out = **in
	}
	if in.DataEncryption != nil {
		in, out := &in.DataEncryption, &out.DataEncryption
		*out = new(DataEncryption_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.FullyQualifiedDomainName != nil {
		in, out := &in.FullyQualifiedDomainName, &out.FullyQualifiedDomainName
		*out = new(string)
		**out = **in
	}
	if in.HighAvailability != nil {
		in, out := &in.HighAvailability, &out.HighAvailability
		*out = new(HighAvailability_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(Identity_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindow_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(Network_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ReplicaCapacity != nil {
		in, out := &in.ReplicaCapacity, &out.ReplicaCapacity
		*out = new(int)
		**out = **in
	}
	if in.ReplicationRole != nil {
		in, out := &in.ReplicationRole, &out.ReplicationRole
		*out = new(ReplicationRole_STATUS)
		**out = **in
	}
	if in.RestorePointInTime != nil {
		in, out := &in.RestorePointInTime, &out.RestorePointInTime
		*out = new(string)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceServerResourceId != nil {
		in, out := &in.SourceServerResourceId, &out.SourceServerResourceId
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ServerProperties_State_STATUS)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(Storage_STATUS)
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(ServerVersion_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServer_STATUS.
func (in *FlexibleServer_STATUS) DeepCopy() *FlexibleServer_STATUS {
	if in == nil {
		return nil
	}
	out := new(FlexibleServer_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServer_Spec) DeepCopyInto(out *FlexibleServer_Spec) {
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
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(Backup)
		(*in).DeepCopyInto(*out)
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(ServerProperties_CreateMode)
		**out = **in
	}
	if in.DataEncryption != nil {
		in, out := &in.DataEncryption, &out.DataEncryption
		*out = new(DataEncryption)
		(*in).DeepCopyInto(*out)
	}
	if in.HighAvailability != nil {
		in, out := &in.HighAvailability, &out.HighAvailability
		*out = new(HighAvailability)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(Identity)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindow)
		(*in).DeepCopyInto(*out)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(Network)
		(*in).DeepCopyInto(*out)
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(FlexibleServerOperatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.ReplicationRole != nil {
		in, out := &in.ReplicationRole, &out.ReplicationRole
		*out = new(ReplicationRole)
		**out = **in
	}
	if in.RestorePointInTime != nil {
		in, out := &in.RestorePointInTime, &out.RestorePointInTime
		*out = new(string)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceServerResourceId != nil {
		in, out := &in.SourceServerResourceId, &out.SourceServerResourceId
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(Storage)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(ServerVersion)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServer_Spec.
func (in *FlexibleServer_Spec) DeepCopy() *FlexibleServer_Spec {
	if in == nil {
		return nil
	}
	out := new(FlexibleServer_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersDatabase) DeepCopyInto(out *FlexibleServersDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersDatabase.
func (in *FlexibleServersDatabase) DeepCopy() *FlexibleServersDatabase {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServersDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersDatabaseList) DeepCopyInto(out *FlexibleServersDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlexibleServersDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersDatabaseList.
func (in *FlexibleServersDatabaseList) DeepCopy() *FlexibleServersDatabaseList {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServersDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersDatabase_STATUS) DeepCopyInto(out *FlexibleServersDatabase_STATUS) {
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
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersDatabase_STATUS.
func (in *FlexibleServersDatabase_STATUS) DeepCopy() *FlexibleServersDatabase_STATUS {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersDatabase_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersDatabase_Spec) DeepCopyInto(out *FlexibleServersDatabase_Spec) {
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
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersDatabase_Spec.
func (in *FlexibleServersDatabase_Spec) DeepCopy() *FlexibleServersDatabase_Spec {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersDatabase_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersFirewallRule) DeepCopyInto(out *FlexibleServersFirewallRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersFirewallRule.
func (in *FlexibleServersFirewallRule) DeepCopy() *FlexibleServersFirewallRule {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersFirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServersFirewallRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersFirewallRuleList) DeepCopyInto(out *FlexibleServersFirewallRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlexibleServersFirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersFirewallRuleList.
func (in *FlexibleServersFirewallRuleList) DeepCopy() *FlexibleServersFirewallRuleList {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersFirewallRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServersFirewallRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersFirewallRule_STATUS) DeepCopyInto(out *FlexibleServersFirewallRule_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EndIpAddress != nil {
		in, out := &in.EndIpAddress, &out.EndIpAddress
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
	if in.StartIpAddress != nil {
		in, out := &in.StartIpAddress, &out.StartIpAddress
		*out = new(string)
		**out = **in
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersFirewallRule_STATUS.
func (in *FlexibleServersFirewallRule_STATUS) DeepCopy() *FlexibleServersFirewallRule_STATUS {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersFirewallRule_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersFirewallRule_Spec) DeepCopyInto(out *FlexibleServersFirewallRule_Spec) {
	*out = *in
	if in.EndIpAddress != nil {
		in, out := &in.EndIpAddress, &out.EndIpAddress
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.StartIpAddress != nil {
		in, out := &in.StartIpAddress, &out.StartIpAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersFirewallRule_Spec.
func (in *FlexibleServersFirewallRule_Spec) DeepCopy() *FlexibleServersFirewallRule_Spec {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersFirewallRule_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailability) DeepCopyInto(out *HighAvailability) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(HighAvailability_Mode)
		**out = **in
	}
	if in.StandbyAvailabilityZone != nil {
		in, out := &in.StandbyAvailabilityZone, &out.StandbyAvailabilityZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HighAvailability.
func (in *HighAvailability) DeepCopy() *HighAvailability {
	if in == nil {
		return nil
	}
	out := new(HighAvailability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailability_STATUS) DeepCopyInto(out *HighAvailability_STATUS) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(HighAvailability_Mode_STATUS)
		**out = **in
	}
	if in.StandbyAvailabilityZone != nil {
		in, out := &in.StandbyAvailabilityZone, &out.StandbyAvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(HighAvailability_State_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HighAvailability_STATUS.
func (in *HighAvailability_STATUS) DeepCopy() *HighAvailability_STATUS {
	if in == nil {
		return nil
	}
	out := new(HighAvailability_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity) DeepCopyInto(out *Identity) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(Identity_Type)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make([]UserAssignedIdentityDetails, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity.
func (in *Identity) DeepCopy() *Identity {
	if in == nil {
		return nil
	}
	out := new(Identity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_STATUS) DeepCopyInto(out *Identity_STATUS) {
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
		*out = new(Identity_Type_STATUS)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_STATUS.
func (in *Identity_STATUS) DeepCopy() *Identity_STATUS {
	if in == nil {
		return nil
	}
	out := new(Identity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindow) DeepCopyInto(out *MaintenanceWindow) {
	*out = *in
	if in.CustomWindow != nil {
		in, out := &in.CustomWindow, &out.CustomWindow
		*out = new(string)
		**out = **in
	}
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(int)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(int)
		**out = **in
	}
	if in.StartMinute != nil {
		in, out := &in.StartMinute, &out.StartMinute
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindow.
func (in *MaintenanceWindow) DeepCopy() *MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindow_STATUS) DeepCopyInto(out *MaintenanceWindow_STATUS) {
	*out = *in
	if in.CustomWindow != nil {
		in, out := &in.CustomWindow, &out.CustomWindow
		*out = new(string)
		**out = **in
	}
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(int)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(int)
		**out = **in
	}
	if in.StartMinute != nil {
		in, out := &in.StartMinute, &out.StartMinute
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindow_STATUS.
func (in *MaintenanceWindow_STATUS) DeepCopy() *MaintenanceWindow_STATUS {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindow_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.DelegatedSubnetResourceReference != nil {
		in, out := &in.DelegatedSubnetResourceReference, &out.DelegatedSubnetResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.PrivateDnsZoneResourceReference != nil {
		in, out := &in.PrivateDnsZoneResourceReference, &out.PrivateDnsZoneResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network_STATUS) DeepCopyInto(out *Network_STATUS) {
	*out = *in
	if in.DelegatedSubnetResourceId != nil {
		in, out := &in.DelegatedSubnetResourceId, &out.DelegatedSubnetResourceId
		*out = new(string)
		**out = **in
	}
	if in.PrivateDnsZoneResourceId != nil {
		in, out := &in.PrivateDnsZoneResourceId, &out.PrivateDnsZoneResourceId
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(EnableStatusEnum_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network_STATUS.
func (in *Network_STATUS) DeepCopy() *Network_STATUS {
	if in == nil {
		return nil
	}
	out := new(Network_STATUS)
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
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(Sku_Tier)
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(Sku_Tier_STATUS)
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
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.AutoGrow != nil {
		in, out := &in.AutoGrow, &out.AutoGrow
		*out = new(EnableStatusEnum)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int)
		**out = **in
	}
	if in.StorageSizeGB != nil {
		in, out := &in.StorageSizeGB, &out.StorageSizeGB
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage_STATUS) DeepCopyInto(out *Storage_STATUS) {
	*out = *in
	if in.AutoGrow != nil {
		in, out := &in.AutoGrow, &out.AutoGrow
		*out = new(EnableStatusEnum_STATUS)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int)
		**out = **in
	}
	if in.StorageSizeGB != nil {
		in, out := &in.StorageSizeGB, &out.StorageSizeGB
		*out = new(int)
		**out = **in
	}
	if in.StorageSku != nil {
		in, out := &in.StorageSku, &out.StorageSku
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage_STATUS.
func (in *Storage_STATUS) DeepCopy() *Storage_STATUS {
	if in == nil {
		return nil
	}
	out := new(Storage_STATUS)
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
