// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230630

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServer_Spec_ARM struct {
	// Identity: The cmk identity for the server.
	Identity *MySQLServerIdentity_ARM `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties_ARM `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *MySQLServerSku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServer_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-06-30"
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

// Properties to configure Identity for Bring your Own Keys
type MySQLServerIdentity_ARM struct {
	// Type: Type of managed service identity.
	Type                   *MySQLServerIdentity_Type                  `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Billing information related properties of a server.
type MySQLServerSku_ARM struct {
	// Name: The name of the sku, e.g. Standard_D32s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. GeneralPurpose.
	Tier *MySQLServerSku_Tier `json:"tier,omitempty"`
}

// The properties of a server.
type ServerProperties_ARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The password of the administrator login (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// AvailabilityZone: availability Zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup related properties of a server.
	Backup *Backup_ARM `json:"backup,omitempty"`

	// CreateMode: The mode to create a new MySQL server.
	CreateMode *ServerProperties_CreateMode `json:"createMode,omitempty"`

	// DataEncryption: The Data Encryption for CMK.
	DataEncryption *DataEncryption_ARM `json:"dataEncryption,omitempty"`

	// HighAvailability: High availability related properties of a server.
	HighAvailability *HighAvailability_ARM `json:"highAvailability,omitempty"`

	// ImportSourceProperties: Source properties for import from storage.
	ImportSourceProperties *ImportSourceProperties_ARM `json:"importSourceProperties,omitempty"`

	// MaintenanceWindow: Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow_ARM `json:"maintenanceWindow,omitempty"`

	// Network: Network related properties of a server.
	Network *Network_ARM `json:"network,omitempty"`

	// ReplicationRole: The replication role.
	ReplicationRole *ReplicationRole `json:"replicationRole,omitempty"`

	// RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime     *string `json:"restorePointInTime,omitempty"`
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// Storage: Storage related properties of a server.
	Storage *Storage_ARM `json:"storage,omitempty"`

	// Version: Server version.
	Version *ServerVersion `json:"version,omitempty"`
}

// Storage Profile properties of a server
type Backup_ARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// GeoRedundantBackup: Whether or not geo redundant backup is enabled.
	GeoRedundantBackup *EnableStatusEnum `json:"geoRedundantBackup,omitempty"`
}

// The date encryption for cmk.
type DataEncryption_ARM struct {
	// GeoBackupKeyURI: Geo backup key uri as key vault can't cross region, need cmk in same region as geo backup
	GeoBackupKeyURI                 *string `json:"geoBackupKeyURI,omitempty"`
	GeoBackupUserAssignedIdentityId *string `json:"geoBackupUserAssignedIdentityId,omitempty"`

	// PrimaryKeyURI: Primary key uri
	PrimaryKeyURI                 *string `json:"primaryKeyURI,omitempty"`
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// Type: The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
	Type *DataEncryption_Type `json:"type,omitempty"`
}

// Network related properties of a server
type HighAvailability_ARM struct {
	// Mode: High availability mode for a server.
	Mode *HighAvailability_Mode `json:"mode,omitempty"`

	// StandbyAvailabilityZone: Availability zone of the standby server.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`
}

// Import source related properties.
type ImportSourceProperties_ARM struct {
	// DataDirPath: Relative path of data directory in storage.
	DataDirPath *string `json:"dataDirPath,omitempty"`

	// SasToken: Sas token for accessing source storage. Read and list permissions are required for sas token.
	SasToken *string `json:"sasToken,omitempty"`

	// StorageType: Storage type of import source.
	StorageType *ImportSourceProperties_StorageType `json:"storageType,omitempty"`

	// StorageUrl: Uri of the import source storage.
	StorageUrl *string `json:"storageUrl,omitempty"`
}

// Maintenance window of a server.
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

// +kubebuilder:validation:Enum={"UserAssigned"}
type MySQLServerIdentity_Type string

const MySQLServerIdentity_Type_UserAssigned = MySQLServerIdentity_Type("UserAssigned")

// Mapping from string to MySQLServerIdentity_Type
var mySQLServerIdentity_Type_Values = map[string]MySQLServerIdentity_Type{
	"userassigned": MySQLServerIdentity_Type_UserAssigned,
}

// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type MySQLServerSku_Tier string

const (
	MySQLServerSku_Tier_Burstable       = MySQLServerSku_Tier("Burstable")
	MySQLServerSku_Tier_GeneralPurpose  = MySQLServerSku_Tier("GeneralPurpose")
	MySQLServerSku_Tier_MemoryOptimized = MySQLServerSku_Tier("MemoryOptimized")
)

// Mapping from string to MySQLServerSku_Tier
var mySQLServerSku_Tier_Values = map[string]MySQLServerSku_Tier{
	"burstable":       MySQLServerSku_Tier_Burstable,
	"generalpurpose":  MySQLServerSku_Tier_GeneralPurpose,
	"memoryoptimized": MySQLServerSku_Tier_MemoryOptimized,
}

// Network related properties of a server
type Network_ARM struct {
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneResourceId  *string `json:"privateDnsZoneResourceId,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is 'Disabled' when server
	// has VNet integration.
	PublicNetworkAccess *EnableStatusEnum `json:"publicNetworkAccess,omitempty"`
}

// Storage Profile properties of a server
type Storage_ARM struct {
	// AutoGrow: Enable Storage Auto Grow or not.
	AutoGrow *EnableStatusEnum `json:"autoGrow,omitempty"`

	// AutoIoScaling: Enable IO Auto Scaling or not.
	AutoIoScaling *EnableStatusEnum `json:"autoIoScaling,omitempty"`

	// Iops: Storage IOPS for a server.
	Iops *int `json:"iops,omitempty"`

	// LogOnDisk: Enable Log On Disk or not.
	LogOnDisk *EnableStatusEnum `json:"logOnDisk,omitempty"`

	// StorageSizeGB: Max storage size allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}
