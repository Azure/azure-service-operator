// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220120preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServer_Spec_ARM struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties_ARM `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServer_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-20-preview"
func (server FlexibleServer_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (server *FlexibleServer_Spec_ARM) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers"
func (server *FlexibleServer_Spec_ARM) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers"
}

// The properties of a server.
type ServerProperties_ARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// AvailabilityZone: availability zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup properties of a server.
	Backup *Backup_ARM `json:"backup,omitempty"`

	// CreateMode: The mode to create a new PostgreSQL server.
	CreateMode *ServerProperties_CreateMode `json:"createMode,omitempty"`

	// HighAvailability: High availability properties of a server.
	HighAvailability *HighAvailability_ARM `json:"highAvailability,omitempty"`

	// MaintenanceWindow: Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindow_ARM `json:"maintenanceWindow,omitempty"`

	// Network: Network properties of a server.
	Network *Network_ARM `json:"network,omitempty"`

	// PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when
	// 'createMode' is 'PointInTimeRestore'.
	PointInTimeUTC         *string `json:"pointInTimeUTC,omitempty"`
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// Storage: Storage properties of a server.
	Storage *Storage_ARM `json:"storage,omitempty"`

	// Version: PostgreSQL Server version.
	Version *ServerVersion `json:"version,omitempty"`
}

// Sku information related properties of a server.
type Sku_ARM struct {
	// Name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. Burstable.
	Tier *Sku_Tier `json:"tier,omitempty"`
}

// Backup properties of a server
type Backup_ARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.
	GeoRedundantBackup *Backup_GeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
}

// High availability properties of a server
type HighAvailability_ARM struct {
	// Mode: The HA mode for the server.
	Mode *HighAvailability_Mode `json:"mode,omitempty"`

	// StandbyAvailabilityZone: availability zone information of the standby.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`
}

// Maintenance window properties of a server.
type MaintenanceWindow_ARM struct {
	// CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	// StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	// StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

// Network properties of a server
type Network_ARM struct {
	DelegatedSubnetResourceId   *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneArmResourceId *string `json:"privateDnsZoneArmResourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type Sku_Tier string

const (
	Sku_Tier_Burstable       = Sku_Tier("Burstable")
	Sku_Tier_GeneralPurpose  = Sku_Tier("GeneralPurpose")
	Sku_Tier_MemoryOptimized = Sku_Tier("MemoryOptimized")
)

// Mapping from string to Sku_Tier
var sku_Tier_Values = map[string]Sku_Tier{
	"burstable":       Sku_Tier_Burstable,
	"generalpurpose":  Sku_Tier_GeneralPurpose,
	"memoryoptimized": Sku_Tier_MemoryOptimized,
}

// Storage properties of a server
type Storage_ARM struct {
	// StorageSizeGB: Max storage allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}
