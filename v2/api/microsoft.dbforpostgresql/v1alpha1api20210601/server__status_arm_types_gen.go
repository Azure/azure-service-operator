// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

//Generated from:
type Server_StatusARM struct {
	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Identity: The Azure Active Directory identity of the server.
	Identity *Identity_StatusARM `json:"identity,omitempty"`

	//Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the server.
	Properties *ServerProperties_StatusARM `json:"properties,omitempty"`

	//Sku: The SKU (pricing tier) of the server.
	Sku *Sku_StatusARM `json:"sku,omitempty"`

	//SystemData: The system metadata relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
	//"Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

//Generated from:
type Identity_StatusARM struct {
	//PrincipalId: The principal ID of resource identity.
	PrincipalId *string `json:"principalId,omitempty"`

	//TenantId: The tenant ID of resource.
	TenantId *string `json:"tenantId,omitempty"`

	//Type: The identity type.
	Type *IdentityStatusType `json:"type,omitempty"`
}

//Generated from:
type ServerProperties_StatusARM struct {
	//AdministratorLogin: The administrator's login name of a server. Can only be
	//specified when the server is being created (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	//AdministratorLoginPassword: The administrator login password (required for
	//server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	//AvailabilityZone: availability zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	//Backup: Backup properties of a server.
	Backup *Backup_StatusARM `json:"backup,omitempty"`

	//CreateMode: The mode to create a new PostgreSQL server.
	CreateMode *ServerPropertiesStatusCreateMode `json:"createMode,omitempty"`

	//FullyQualifiedDomainName: The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	//HighAvailability: High availability properties of a server.
	HighAvailability *HighAvailability_StatusARM `json:"highAvailability,omitempty"`

	//MaintenanceWindow: Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindow_StatusARM `json:"maintenanceWindow,omitempty"`

	//MinorVersion: The minor version of the server.
	MinorVersion *string `json:"minorVersion,omitempty"`

	//Network: Network properties of a server.
	Network *Network_StatusARM `json:"network,omitempty"`

	//PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the
	//time to restore from. It's required when 'createMode' is 'PointInTimeRestore'.
	PointInTimeUTC *string `json:"pointInTimeUTC,omitempty"`

	//SourceServerResourceId: The source server resource ID to restore from. It's
	//required when 'createMode' is 'PointInTimeRestore'.
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	//State: A state of a server that is visible to user.
	State *ServerPropertiesStatusState `json:"state,omitempty"`

	//Storage: Storage properties of a server.
	Storage *Storage_StatusARM `json:"storage,omitempty"`

	//Tags: Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`

	//Version: PostgreSQL Server version.
	Version *ServerVersion_Status `json:"version,omitempty"`
}

//Generated from:
type Sku_StatusARM struct {
	//Name: The name of the sku, typically, tier + family + cores, e.g.
	//Standard_D4s_v3.
	Name string `json:"name"`

	//Tier: The tier of the particular SKU, e.g. Burstable.
	Tier SkuStatusTier `json:"tier"`
}

//Generated from:
type Backup_StatusARM struct {
	//BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	//EarliestRestoreDate: The earliest restore point time (ISO8601 format) for server.
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`

	//GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled
	//on the server.
	GeoRedundantBackup *BackupStatusGeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
}

//Generated from:
type HighAvailability_StatusARM struct {
	//Mode: The HA mode for the server.
	Mode *HighAvailabilityStatusMode `json:"mode,omitempty"`

	//StandbyAvailabilityZone: availability zone information of the standby.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`

	//State: A state of a HA server that is visible to user.
	State *HighAvailabilityStatusState `json:"state,omitempty"`
}

type IdentityStatusType string

const IdentityStatusTypeSystemAssigned = IdentityStatusType("SystemAssigned")

//Generated from:
type MaintenanceWindow_StatusARM struct {
	//CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	//DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	//StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	//StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

//Generated from:
type Network_StatusARM struct {
	//DelegatedSubnetResourceId: delegated subnet arm resource id.
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`

	//PrivateDnsZoneArmResourceId: private dns zone arm resource id.
	PrivateDnsZoneArmResourceId *string `json:"privateDnsZoneArmResourceId,omitempty"`

	//PublicNetworkAccess: public network access is enabled or not
	PublicNetworkAccess *NetworkStatusPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

type SkuStatusTier string

const (
	SkuStatusTierBurstable       = SkuStatusTier("Burstable")
	SkuStatusTierGeneralPurpose  = SkuStatusTier("GeneralPurpose")
	SkuStatusTierMemoryOptimized = SkuStatusTier("MemoryOptimized")
)

//Generated from:
type Storage_StatusARM struct {
	//StorageSizeGB: Max storage allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}
