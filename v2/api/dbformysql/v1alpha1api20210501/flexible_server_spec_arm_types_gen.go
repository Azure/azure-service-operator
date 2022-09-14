// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of FlexibleServer_Spec. Use v1beta20210501.FlexibleServer_Spec instead
type FlexibleServer_Spec_ARM struct {
	Identity   *Identity_ARM         `json:"identity,omitempty"`
	Location   *string               `json:"location,omitempty"`
	Name       string                `json:"name,omitempty"`
	Properties *ServerProperties_ARM `json:"properties,omitempty"`
	Sku        *Sku_ARM              `json:"sku,omitempty"`
	Tags       map[string]string     `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServer_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (server FlexibleServer_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (server *FlexibleServer_Spec_ARM) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers"
func (server *FlexibleServer_Spec_ARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers"
}

// Deprecated version of Identity. Use v1beta20210501.Identity instead
type Identity_ARM struct {
	Type *Identity_Type `json:"type,omitempty"`
}

// Deprecated version of ServerProperties. Use v1beta20210501.ServerProperties instead
type ServerProperties_ARM struct {
	AdministratorLogin         *string                           `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *string                           `json:"administratorLoginPassword,omitempty"`
	AvailabilityZone           *string                           `json:"availabilityZone,omitempty"`
	Backup                     *Backup_ARM                       `json:"backup,omitempty"`
	CreateMode                 *ServerProperties_CreateMode      `json:"createMode,omitempty"`
	DataEncryption             *DataEncryption_ARM               `json:"dataEncryption,omitempty"`
	HighAvailability           *HighAvailability_ARM             `json:"highAvailability,omitempty"`
	MaintenanceWindow          *MaintenanceWindow_ARM            `json:"maintenanceWindow,omitempty"`
	Network                    *Network_ARM                      `json:"network,omitempty"`
	ReplicationRole            *ServerProperties_ReplicationRole `json:"replicationRole,omitempty"`
	RestorePointInTime         *string                           `json:"restorePointInTime,omitempty"`
	SourceServerResourceId     *string                           `json:"sourceServerResourceId,omitempty"`
	Storage                    *Storage_ARM                      `json:"storage,omitempty"`
	Version                    *ServerProperties_Version         `json:"version,omitempty"`
}

// Deprecated version of Sku. Use v1beta20210501.Sku instead
type Sku_ARM struct {
	Name *string   `json:"name,omitempty"`
	Tier *Sku_Tier `json:"tier,omitempty"`
}

// Deprecated version of Backup. Use v1beta20210501.Backup instead
type Backup_ARM struct {
	BackupRetentionDays *int                       `json:"backupRetentionDays,omitempty"`
	GeoRedundantBackup  *Backup_GeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
}

// Deprecated version of DataEncryption. Use v1beta20210501.DataEncryption instead
type DataEncryption_ARM struct {
	GeoBackupKeyURI                 *string              `json:"geoBackupKeyURI,omitempty"`
	GeoBackupUserAssignedIdentityId *string              `json:"geoBackupUserAssignedIdentityId,omitempty"`
	PrimaryKeyURI                   *string              `json:"primaryKeyURI,omitempty"`
	PrimaryUserAssignedIdentityId   *string              `json:"primaryUserAssignedIdentityId,omitempty"`
	Type                            *DataEncryption_Type `json:"type,omitempty"`
}

// Deprecated version of HighAvailability. Use v1beta20210501.HighAvailability instead
type HighAvailability_ARM struct {
	Mode                    *HighAvailability_Mode `json:"mode,omitempty"`
	StandbyAvailabilityZone *string                `json:"standbyAvailabilityZone,omitempty"`
}

// Deprecated version of Identity_Type. Use v1beta20210501.Identity_Type instead
// +kubebuilder:validation:Enum={"UserAssigned"}
type Identity_Type string

const Identity_Type_UserAssigned = Identity_Type("UserAssigned")

// Deprecated version of MaintenanceWindow. Use v1beta20210501.MaintenanceWindow instead
type MaintenanceWindow_ARM struct {
	CustomWindow *string `json:"customWindow,omitempty"`
	DayOfWeek    *int    `json:"dayOfWeek,omitempty"`
	StartHour    *int    `json:"startHour,omitempty"`
	StartMinute  *int    `json:"startMinute,omitempty"`
}

// Deprecated version of Network. Use v1beta20210501.Network instead
type Network_ARM struct {
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneResourceId  *string `json:"privateDnsZoneResourceId,omitempty"`
}

// Deprecated version of Sku_Tier. Use v1beta20210501.Sku_Tier instead
// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type Sku_Tier string

const (
	Sku_Tier_Burstable       = Sku_Tier("Burstable")
	Sku_Tier_GeneralPurpose  = Sku_Tier("GeneralPurpose")
	Sku_Tier_MemoryOptimized = Sku_Tier("MemoryOptimized")
)

// Deprecated version of Storage. Use v1beta20210501.Storage instead
type Storage_ARM struct {
	AutoGrow      *Storage_AutoGrow `json:"autoGrow,omitempty"`
	Iops          *int              `json:"iops,omitempty"`
	StorageSizeGB *int              `json:"storageSizeGB,omitempty"`
}
