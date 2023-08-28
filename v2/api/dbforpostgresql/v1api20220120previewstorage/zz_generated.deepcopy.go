//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20220120previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
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
	if in.FullyQualifiedDomainName != nil {
		in, out := &in.FullyQualifiedDomainName, &out.FullyQualifiedDomainName
		*out = new(genruntime.ConfigMapDestination)
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
func (in *FlexibleServerOperatorSpec) DeepCopyInto(out *FlexibleServerOperatorSpec) {
	*out = *in
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = new(FlexibleServerOperatorConfigMaps)
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
		*out = new(string)
		**out = **in
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
	if in.MinorVersion != nil {
		in, out := &in.MinorVersion, &out.MinorVersion
		*out = new(string)
		**out = **in
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
	if in.PointInTimeUTC != nil {
		in, out := &in.PointInTimeUTC, &out.PointInTimeUTC
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
		*out = new(string)
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
		*out = new(string)
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
		*out = new(string)
		**out = **in
	}
	if in.HighAvailability != nil {
		in, out := &in.HighAvailability, &out.HighAvailability
		*out = new(HighAvailability)
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
	if in.PointInTimeUTC != nil {
		in, out := &in.PointInTimeUTC, &out.PointInTimeUTC
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
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceServerResourceReference != nil {
		in, out := &in.SourceServerResourceReference, &out.SourceServerResourceReference
		*out = new(genruntime.ResourceReference)
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
		*out = new(string)
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
func (in *FlexibleServersConfiguration) DeepCopyInto(out *FlexibleServersConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersConfiguration.
func (in *FlexibleServersConfiguration) DeepCopy() *FlexibleServersConfiguration {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServersConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersConfigurationList) DeepCopyInto(out *FlexibleServersConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlexibleServersConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersConfigurationList.
func (in *FlexibleServersConfigurationList) DeepCopy() *FlexibleServersConfigurationList {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServersConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
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
func (in *FlexibleServers_Configuration_STATUS) DeepCopyInto(out *FlexibleServers_Configuration_STATUS) {
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
	if in.DocumentationLink != nil {
		in, out := &in.DocumentationLink, &out.DocumentationLink
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.IsConfigPendingRestart != nil {
		in, out := &in.IsConfigPendingRestart, &out.IsConfigPendingRestart
		*out = new(bool)
		**out = **in
	}
	if in.IsDynamicConfig != nil {
		in, out := &in.IsDynamicConfig, &out.IsDynamicConfig
		*out = new(bool)
		**out = **in
	}
	if in.IsReadOnly != nil {
		in, out := &in.IsReadOnly, &out.IsReadOnly
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
	if in.Source != nil {
		in, out := &in.Source, &out.Source
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
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServers_Configuration_STATUS.
func (in *FlexibleServers_Configuration_STATUS) DeepCopy() *FlexibleServers_Configuration_STATUS {
	if in == nil {
		return nil
	}
	out := new(FlexibleServers_Configuration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServers_Configuration_Spec) DeepCopyInto(out *FlexibleServers_Configuration_Spec) {
	*out = *in
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
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServers_Configuration_Spec.
func (in *FlexibleServers_Configuration_Spec) DeepCopy() *FlexibleServers_Configuration_Spec {
	if in == nil {
		return nil
	}
	out := new(FlexibleServers_Configuration_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServers_Database_STATUS) DeepCopyInto(out *FlexibleServers_Database_STATUS) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServers_Database_STATUS.
func (in *FlexibleServers_Database_STATUS) DeepCopy() *FlexibleServers_Database_STATUS {
	if in == nil {
		return nil
	}
	out := new(FlexibleServers_Database_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServers_Database_Spec) DeepCopyInto(out *FlexibleServers_Database_Spec) {
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServers_Database_Spec.
func (in *FlexibleServers_Database_Spec) DeepCopy() *FlexibleServers_Database_Spec {
	if in == nil {
		return nil
	}
	out := new(FlexibleServers_Database_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServers_FirewallRule_STATUS) DeepCopyInto(out *FlexibleServers_FirewallRule_STATUS) {
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServers_FirewallRule_STATUS.
func (in *FlexibleServers_FirewallRule_STATUS) DeepCopy() *FlexibleServers_FirewallRule_STATUS {
	if in == nil {
		return nil
	}
	out := new(FlexibleServers_FirewallRule_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServers_FirewallRule_Spec) DeepCopyInto(out *FlexibleServers_FirewallRule_Spec) {
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StartIpAddress != nil {
		in, out := &in.StartIpAddress, &out.StartIpAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServers_FirewallRule_Spec.
func (in *FlexibleServers_FirewallRule_Spec) DeepCopy() *FlexibleServers_FirewallRule_Spec {
	if in == nil {
		return nil
	}
	out := new(FlexibleServers_FirewallRule_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailability) DeepCopyInto(out *HighAvailability) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
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
	if in.StandbyAvailabilityZone != nil {
		in, out := &in.StandbyAvailabilityZone, &out.StandbyAvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
	if in.PrivateDnsZoneArmResourceReference != nil {
		in, out := &in.PrivateDnsZoneArmResourceReference, &out.PrivateDnsZoneArmResourceReference
		*out = new(genruntime.ResourceReference)
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
	if in.PrivateDnsZoneArmResourceId != nil {
		in, out := &in.PrivateDnsZoneArmResourceId, &out.PrivateDnsZoneArmResourceId
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StorageSizeGB != nil {
		in, out := &in.StorageSizeGB, &out.StorageSizeGB
		*out = new(int)
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
		*out = new(string)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS.
func (in *SystemData_STATUS) DeepCopy() *SystemData_STATUS {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS)
	in.DeepCopyInto(out)
	return out
}
