// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of FlexibleServers_Spec. Use v1beta20210501.FlexibleServers_Spec instead
type FlexibleServers_SpecARM struct {
	Location   *string              `json:"location,omitempty"`
	Name       string               `json:"name,omitempty"`
	Properties *ServerPropertiesARM `json:"properties,omitempty"`
	Sku        *SkuARM              `json:"sku,omitempty"`
	Tags       map[string]string    `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (servers FlexibleServers_SpecARM) GetAPIVersion() string {
	return "2021-05-01"
}

// GetName returns the Name of the resource
func (servers FlexibleServers_SpecARM) GetName() string {
	return servers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers"
func (servers FlexibleServers_SpecARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers"
}

// Deprecated version of ServerProperties. Use v1beta20210501.ServerProperties instead
type ServerPropertiesARM struct {
	AdministratorLogin         *string                          `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *string                          `json:"administratorLoginPassword,omitempty"`
	AvailabilityZone           *string                          `json:"availabilityZone,omitempty"`
	Backup                     *BackupARM                       `json:"backup,omitempty"`
	CreateMode                 *ServerPropertiesCreateMode      `json:"createMode,omitempty"`
	HighAvailability           *HighAvailabilityARM             `json:"highAvailability,omitempty"`
	MaintenanceWindow          *MaintenanceWindowARM            `json:"maintenanceWindow,omitempty"`
	Network                    *NetworkARM                      `json:"network,omitempty"`
	ReplicationRole            *ServerPropertiesReplicationRole `json:"replicationRole,omitempty"`
	RestorePointInTime         *string                          `json:"restorePointInTime,omitempty"`
	SourceServerResourceId     *string                          `json:"sourceServerResourceId,omitempty"`
	Storage                    *StorageARM                      `json:"storage,omitempty"`
	Version                    *ServerPropertiesVersion         `json:"version,omitempty"`
}

// Deprecated version of Sku. Use v1beta20210501.Sku instead
type SkuARM struct {
	Name *string  `json:"name,omitempty"`
	Tier *SkuTier `json:"tier,omitempty"`
}

// Deprecated version of Backup. Use v1beta20210501.Backup instead
type BackupARM struct {
	BackupRetentionDays *int                      `json:"backupRetentionDays,omitempty"`
	GeoRedundantBackup  *BackupGeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
}

// Deprecated version of HighAvailability. Use v1beta20210501.HighAvailability instead
type HighAvailabilityARM struct {
	Mode                    *HighAvailabilityMode `json:"mode,omitempty"`
	StandbyAvailabilityZone *string               `json:"standbyAvailabilityZone,omitempty"`
}

// Deprecated version of MaintenanceWindow. Use v1beta20210501.MaintenanceWindow instead
type MaintenanceWindowARM struct {
	CustomWindow *string `json:"customWindow,omitempty"`
	DayOfWeek    *int    `json:"dayOfWeek,omitempty"`
	StartHour    *int    `json:"startHour,omitempty"`
	StartMinute  *int    `json:"startMinute,omitempty"`
}

// Deprecated version of Network. Use v1beta20210501.Network instead
type NetworkARM struct {
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneResourceId  *string `json:"privateDnsZoneResourceId,omitempty"`
}

// Deprecated version of SkuTier. Use v1beta20210501.SkuTier instead
// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type SkuTier string

const (
	SkuTierBurstable       = SkuTier("Burstable")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

// Deprecated version of Storage. Use v1beta20210501.Storage instead
type StorageARM struct {
	AutoGrow      *StorageAutoGrow `json:"autoGrow,omitempty"`
	Iops          *int             `json:"iops,omitempty"`
	StorageSizeGB *int             `json:"storageSizeGB,omitempty"`
}
