// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220120preview

type Server_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *Sku_STATUS_ARM `json:"sku,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type ServerProperties_STATUS_ARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AvailabilityZone: availability zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup properties of a server.
	Backup *Backup_STATUS_ARM `json:"backup,omitempty"`

	// CreateMode: The mode to create a new PostgreSQL server.
	CreateMode *ServerProperties_CreateMode_STATUS `json:"createMode,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// HighAvailability: High availability properties of a server.
	HighAvailability *HighAvailability_STATUS_ARM `json:"highAvailability,omitempty"`

	// MaintenanceWindow: Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindow_STATUS_ARM `json:"maintenanceWindow,omitempty"`

	// MinorVersion: The minor version of the server.
	MinorVersion *string `json:"minorVersion,omitempty"`

	// Network: Network properties of a server.
	Network *Network_STATUS_ARM `json:"network,omitempty"`

	// PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when
	// 'createMode' is 'PointInTimeRestore'.
	PointInTimeUTC *string `json:"pointInTimeUTC,omitempty"`

	// SourceServerResourceId: The source server resource ID to restore from. It's required when 'createMode' is
	// 'PointInTimeRestore'.
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// State: A state of a server that is visible to user.
	State *ServerProperties_State_STATUS `json:"state,omitempty"`

	// Storage: Storage properties of a server.
	Storage *Storage_STATUS_ARM `json:"storage,omitempty"`

	// Version: PostgreSQL Server version.
	Version *ServerVersion_STATUS `json:"version,omitempty"`
}

type Sku_STATUS_ARM struct {
	// Name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. Burstable.
	Tier *Sku_Tier_STATUS `json:"tier,omitempty"`
}

type Backup_STATUS_ARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// EarliestRestoreDate: The earliest restore point time (ISO8601 format) for server.
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`

	// GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.
	GeoRedundantBackup *Backup_GeoRedundantBackup_STATUS `json:"geoRedundantBackup,omitempty"`
}

type HighAvailability_STATUS_ARM struct {
	// Mode: The HA mode for the server.
	Mode *HighAvailability_Mode_STATUS `json:"mode,omitempty"`

	// StandbyAvailabilityZone: availability zone information of the standby.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`

	// State: A state of a HA server that is visible to user.
	State *HighAvailability_State_STATUS `json:"state,omitempty"`
}

type MaintenanceWindow_STATUS_ARM struct {
	// CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	// StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	// StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

type Network_STATUS_ARM struct {
	// DelegatedSubnetResourceId: delegated subnet arm resource id.
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`

	// PrivateDnsZoneArmResourceId: private dns zone arm resource id.
	PrivateDnsZoneArmResourceId *string `json:"privateDnsZoneArmResourceId,omitempty"`

	// PublicNetworkAccess: public network access is enabled or not
	PublicNetworkAccess *Network_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`
}

type Sku_Tier_STATUS string

const (
	Sku_Tier_STATUS_Burstable       = Sku_Tier_STATUS("Burstable")
	Sku_Tier_STATUS_GeneralPurpose  = Sku_Tier_STATUS("GeneralPurpose")
	Sku_Tier_STATUS_MemoryOptimized = Sku_Tier_STATUS("MemoryOptimized")
)

type Storage_STATUS_ARM struct {
	// StorageSizeGB: Max storage allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}
