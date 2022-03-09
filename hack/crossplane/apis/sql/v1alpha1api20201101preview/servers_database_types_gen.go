// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101preview

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversdatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversdatabases/status,serversdatabases/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
//Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/resourceDefinitions/servers_databases
type ServersDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServersDatabases_Spec `json:"spec,omitempty"`
	Status            Database_Status       `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/resourceDefinitions/servers_databases
type ServersDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabase `json:"items"`
}

type Database_Status struct {
	v1alpha1.ResourceStatus `json:",inline,omitempty"`
	AtProvider              DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:validation:Enum={"2020-11-01-preview"}
type ServersDatabasesSpecAPIVersion string

const ServersDatabasesSpecAPIVersion20201101Preview = ServersDatabasesSpecAPIVersion("2020-11-01-preview")

type ServersDatabases_Spec struct {
	v1alpha1.ResourceSpec `json:",inline,omitempty"`
	ForProvider           ServersDatabasesParameters `json:"forProvider,omitempty"`
}

type DatabaseObservation struct {
	//AutoPauseDelay: Time in minutes after which database is automatically paused. A value of -1 means that automatic pause
	//is disabled
	AutoPauseDelay *int `json:"autoPauseDelay,omitempty"`

	//CatalogCollation: Collation of the metadata catalog.
	CatalogCollation *DatabasePropertiesStatusCatalogCollation `json:"catalogCollation,omitempty"`

	//Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	//CreateMode: Specifies the mode of database creation.
	//Default: regular database creation.
	//Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the
	//source database.
	//Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the
	//resource ID of the existing primary database.
	//PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId
	//must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable
	//database resource ID to restore.
	//Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If
	//sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise
	//sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored.
	//restorePointInTime may also be specified to restore from an earlier point in time.
	//RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault.
	//recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode *DatabasePropertiesStatusCreateMode `json:"createMode,omitempty"`

	//CreationDate: The creation date of the database (ISO8601 format).
	CreationDate *string `json:"creationDate,omitempty"`

	//CurrentBackupStorageRedundancy: The storage account type used to store backups for this database.
	CurrentBackupStorageRedundancy *DatabasePropertiesStatusCurrentBackupStorageRedundancy `json:"currentBackupStorageRedundancy,omitempty"`

	//CurrentServiceObjectiveName: The current service level objective name of the database.
	CurrentServiceObjectiveName *string `json:"currentServiceObjectiveName,omitempty"`

	//CurrentSku: The name and tier of the SKU.
	CurrentSku *Sku_Status `json:"currentSku,omitempty"`

	//DatabaseId: The ID of the database.
	DatabaseId *string `json:"databaseId,omitempty"`

	//DefaultSecondaryLocation: The default secondary region for this database.
	DefaultSecondaryLocation *string `json:"defaultSecondaryLocation,omitempty"`

	//EarliestRestoreDate: This records the earliest start date and time that restore is available for this database (ISO8601
	//format).
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`

	//ElasticPoolId: The resource identifier of the elastic pool containing this database.
	ElasticPoolId *string `json:"elasticPoolId,omitempty"`

	//FailoverGroupId: Failover Group resource identifier that this database belongs to.
	FailoverGroupId *string `json:"failoverGroupId,omitempty"`

	//HighAvailabilityReplicaCount: The number of secondary replicas associated with the database that are used to provide
	//high availability.
	HighAvailabilityReplicaCount *int `json:"highAvailabilityReplicaCount,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Kind: Kind of database. This is metadata used for the Azure portal experience.
	Kind *string `json:"kind,omitempty"`

	//LicenseType: The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you
	//have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType *DatabasePropertiesStatusLicenseType `json:"licenseType,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//LongTermRetentionBackupResourceId: The resource identifier of the long term retention backup associated with create
	//operation of this database.
	LongTermRetentionBackupResourceId *string `json:"longTermRetentionBackupResourceId,omitempty"`

	//MaintenanceConfigurationId: Maintenance configuration id assigned to the database. This configuration defines the period
	//when the maintenance updates will occur.
	MaintenanceConfigurationId *string `json:"maintenanceConfigurationId,omitempty"`

	//ManagedBy: Resource that manages the database.
	ManagedBy *string `json:"managedBy,omitempty"`

	//MaxLogSizeBytes: The max log size for this database.
	MaxLogSizeBytes *int `json:"maxLogSizeBytes,omitempty"`

	//MaxSizeBytes: The max size of the database expressed in bytes.
	MaxSizeBytes *int `json:"maxSizeBytes,omitempty"`

	//MinCapacity: Minimal capacity that database will always have allocated, if not paused
	MinCapacity *float64 `json:"minCapacity,omitempty"`

	//Name: Resource name.
	Name *string `json:"name,omitempty"`

	//PausedDate: The date when database was paused by user configuration or action(ISO8601 format). Null if the database is
	//ready.
	PausedDate *string `json:"pausedDate,omitempty"`

	//ReadScale: The state of read-only routing. If enabled, connections that have application intent set to readonly in their
	//connection string may be routed to a readonly secondary replica in the same region.
	ReadScale *DatabasePropertiesStatusReadScale `json:"readScale,omitempty"`

	//RecoverableDatabaseId: The resource identifier of the recoverable database associated with create operation of this
	//database.
	RecoverableDatabaseId *string `json:"recoverableDatabaseId,omitempty"`

	//RecoveryServicesRecoveryPointId: The resource identifier of the recovery point associated with create operation of this
	//database.
	RecoveryServicesRecoveryPointId *string `json:"recoveryServicesRecoveryPointId,omitempty"`

	//RequestedBackupStorageRedundancy: The storage account type to be used to store backups for this database.
	RequestedBackupStorageRedundancy *DatabasePropertiesStatusRequestedBackupStorageRedundancy `json:"requestedBackupStorageRedundancy,omitempty"`

	//RequestedServiceObjectiveName: The requested service level objective name of the database.
	RequestedServiceObjectiveName *string `json:"requestedServiceObjectiveName,omitempty"`

	//RestorableDroppedDatabaseId: The resource identifier of the restorable dropped database associated with create operation
	//of this database.
	RestorableDroppedDatabaseId *string `json:"restorableDroppedDatabaseId,omitempty"`

	//RestorePointInTime: Specifies the point in time (ISO8601 format) of the source database that will be restored to create
	//the new database.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	//ResumedDate: The date when database was resumed by user action or database login (ISO8601 format). Null if the database
	//is paused.
	ResumedDate *string `json:"resumedDate,omitempty"`

	//SampleName: The name of the sample schema to apply when creating this database.
	SampleName *DatabasePropertiesStatusSampleName `json:"sampleName,omitempty"`

	//SecondaryType: The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
	SecondaryType *DatabasePropertiesStatusSecondaryType `json:"secondaryType,omitempty"`

	//Sku: The database SKU.
	//The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition,
	//family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation`
	//REST API or one of the following commands:
	//```azurecli
	//az sql db list-editions -l <location> -o table
	//````
	//```powershell
	//Get-AzSqlServerServiceObjective -Location <location>
	//````
	Sku *Sku_Status `json:"sku,omitempty"`

	//SourceDatabaseDeletionDate: Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `json:"sourceDatabaseDeletionDate,omitempty"`

	//SourceDatabaseId: The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId *string `json:"sourceDatabaseId,omitempty"`

	//Status: The status of the database.
	Status *DatabasePropertiesStatusStatus `json:"status,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`

	//ZoneRedundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread
	//across multiple availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

type ServersDatabasesParameters struct {
	//AutoPauseDelay: Time in minutes after which database is automatically paused. A value of -1 means that automatic pause
	//is disabled
	AutoPauseDelay *int `json:"autoPauseDelay,omitempty"`

	//CatalogCollation: Collation of the metadata catalog.
	CatalogCollation *DatabasePropertiesCatalogCollation `json:"catalogCollation,omitempty"`

	//Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	//CreateMode: Specifies the mode of database creation.
	//Default: regular database creation.
	//Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the
	//source database.
	//Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the
	//resource ID of the existing primary database.
	//PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId
	//must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable
	//database resource ID to restore.
	//Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If
	//sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise
	//sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored.
	//restorePointInTime may also be specified to restore from an earlier point in time.
	//RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault.
	//recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode *DatabasePropertiesCreateMode `json:"createMode,omitempty"`

	//ElasticPoolId: The resource identifier of the elastic pool containing this database.
	ElasticPoolId *string `json:"elasticPoolId,omitempty"`

	//HighAvailabilityReplicaCount: The number of secondary replicas associated with the database that are used to provide
	//high availability.
	HighAvailabilityReplicaCount *int `json:"highAvailabilityReplicaCount,omitempty"`

	//LicenseType: The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you
	//have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType *DatabasePropertiesLicenseType `json:"licenseType,omitempty"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//LongTermRetentionBackupResourceId: The resource identifier of the long term retention backup associated with create
	//operation of this database.
	LongTermRetentionBackupResourceId *string `json:"longTermRetentionBackupResourceId,omitempty"`

	//MaintenanceConfigurationId: Maintenance configuration id assigned to the database. This configuration defines the period
	//when the maintenance updates will occur.
	MaintenanceConfigurationId *string `json:"maintenanceConfigurationId,omitempty"`

	//MaxSizeBytes: The max size of the database expressed in bytes.
	MaxSizeBytes *int `json:"maxSizeBytes,omitempty"`

	//MinCapacity: Minimal capacity that database will always have allocated, if not paused
	MinCapacity *float64 `json:"minCapacity,omitempty"`

	// +kubebuilder:validation:Required
	//Name: The name of the database.
	Name string `json:"name,omitempty"`

	//ReadScale: The state of read-only routing. If enabled, connections that have application intent set to readonly in their
	//connection string may be routed to a readonly secondary replica in the same region.
	ReadScale *DatabasePropertiesReadScale `json:"readScale,omitempty"`

	//RecoverableDatabaseId: The resource identifier of the recoverable database associated with create operation of this
	//database.
	RecoverableDatabaseId *string `json:"recoverableDatabaseId,omitempty"`

	//RecoveryServicesRecoveryPointId: The resource identifier of the recovery point associated with create operation of this
	//database.
	RecoveryServicesRecoveryPointId *string `json:"recoveryServicesRecoveryPointId,omitempty"`

	//RequestedBackupStorageRedundancy: The storage account type to be used to store backups for this database.
	RequestedBackupStorageRedundancy *DatabasePropertiesRequestedBackupStorageRedundancy `json:"requestedBackupStorageRedundancy,omitempty"`
	ResourceGroupName                string                                              `json:"resourceGroupName,omitempty"`
	ResourceGroupNameRef             *v1alpha1.Reference                                 `json:"resourceGroupNameRef,omitempty"`
	ResourceGroupNameSelector        *v1alpha1.Selector                                  `json:"resourceGroupNameSelector,omitempty"`

	//RestorableDroppedDatabaseId: The resource identifier of the restorable dropped database associated with create operation
	//of this database.
	RestorableDroppedDatabaseId *string `json:"restorableDroppedDatabaseId,omitempty"`

	//RestorePointInTime: Specifies the point in time (ISO8601 format) of the source database that will be restored to create
	//the new database.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	//SampleName: The name of the sample schema to apply when creating this database.
	SampleName *DatabasePropertiesSampleName `json:"sampleName,omitempty"`

	//SecondaryType: The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
	SecondaryType      *DatabasePropertiesSecondaryType `json:"secondaryType,omitempty"`
	ServerName         string                           `json:"serverName,omitempty"`
	ServerNameRef      *v1alpha1.Reference              `json:"serverNameRef,omitempty"`
	ServerNameSelector *v1alpha1.Selector               `json:"serverNameSelector,omitempty"`

	//Sku: An ARM Resource SKU.
	Sku *Sku `json:"sku,omitempty"`

	//SourceDatabaseDeletionDate: Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `json:"sourceDatabaseDeletionDate,omitempty"`

	//SourceDatabaseId: The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId *string `json:"sourceDatabaseId,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	//ZoneRedundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread
	//across multiple availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

// +kubebuilder:validation:Enum={"DATABASE_DEFAULT","SQL_Latin1_General_CP1_CI_AS"}
type DatabasePropertiesCatalogCollation string

const (
	DatabasePropertiesCatalogCollationDATABASEDEFAULT         = DatabasePropertiesCatalogCollation("DATABASE_DEFAULT")
	DatabasePropertiesCatalogCollationSQLLatin1GeneralCP1CIAS = DatabasePropertiesCatalogCollation("SQL_Latin1_General_CP1_CI_AS")
)

// +kubebuilder:validation:Enum={"Copy","Default","OnlineSecondary","PointInTimeRestore","Recovery","Restore","RestoreExternalBackup","RestoreExternalBackupSecondary","RestoreLongTermRetentionBackup","Secondary"}
type DatabasePropertiesCreateMode string

const (
	DatabasePropertiesCreateModeCopy                           = DatabasePropertiesCreateMode("Copy")
	DatabasePropertiesCreateModeDefault                        = DatabasePropertiesCreateMode("Default")
	DatabasePropertiesCreateModeOnlineSecondary                = DatabasePropertiesCreateMode("OnlineSecondary")
	DatabasePropertiesCreateModePointInTimeRestore             = DatabasePropertiesCreateMode("PointInTimeRestore")
	DatabasePropertiesCreateModeRecovery                       = DatabasePropertiesCreateMode("Recovery")
	DatabasePropertiesCreateModeRestore                        = DatabasePropertiesCreateMode("Restore")
	DatabasePropertiesCreateModeRestoreExternalBackup          = DatabasePropertiesCreateMode("RestoreExternalBackup")
	DatabasePropertiesCreateModeRestoreExternalBackupSecondary = DatabasePropertiesCreateMode("RestoreExternalBackupSecondary")
	DatabasePropertiesCreateModeRestoreLongTermRetentionBackup = DatabasePropertiesCreateMode("RestoreLongTermRetentionBackup")
	DatabasePropertiesCreateModeSecondary                      = DatabasePropertiesCreateMode("Secondary")
)

// +kubebuilder:validation:Enum={"BasePrice","LicenseIncluded"}
type DatabasePropertiesLicenseType string

const (
	DatabasePropertiesLicenseTypeBasePrice       = DatabasePropertiesLicenseType("BasePrice")
	DatabasePropertiesLicenseTypeLicenseIncluded = DatabasePropertiesLicenseType("LicenseIncluded")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type DatabasePropertiesReadScale string

const (
	DatabasePropertiesReadScaleDisabled = DatabasePropertiesReadScale("Disabled")
	DatabasePropertiesReadScaleEnabled  = DatabasePropertiesReadScale("Enabled")
)

// +kubebuilder:validation:Enum={"Geo","Local","Zone"}
type DatabasePropertiesRequestedBackupStorageRedundancy string

const (
	DatabasePropertiesRequestedBackupStorageRedundancyGeo   = DatabasePropertiesRequestedBackupStorageRedundancy("Geo")
	DatabasePropertiesRequestedBackupStorageRedundancyLocal = DatabasePropertiesRequestedBackupStorageRedundancy("Local")
	DatabasePropertiesRequestedBackupStorageRedundancyZone  = DatabasePropertiesRequestedBackupStorageRedundancy("Zone")
)

// +kubebuilder:validation:Enum={"AdventureWorksLT","WideWorldImportersFull","WideWorldImportersStd"}
type DatabasePropertiesSampleName string

const (
	DatabasePropertiesSampleNameAdventureWorksLT       = DatabasePropertiesSampleName("AdventureWorksLT")
	DatabasePropertiesSampleNameWideWorldImportersFull = DatabasePropertiesSampleName("WideWorldImportersFull")
	DatabasePropertiesSampleNameWideWorldImportersStd  = DatabasePropertiesSampleName("WideWorldImportersStd")
)

// +kubebuilder:validation:Enum={"Geo","Named"}
type DatabasePropertiesSecondaryType string

const (
	DatabasePropertiesSecondaryTypeGeo   = DatabasePropertiesSecondaryType("Geo")
	DatabasePropertiesSecondaryTypeNamed = DatabasePropertiesSecondaryType("Named")
)

type DatabasePropertiesStatusCatalogCollation string

const (
	DatabasePropertiesStatusCatalogCollationDATABASEDEFAULT         = DatabasePropertiesStatusCatalogCollation("DATABASE_DEFAULT")
	DatabasePropertiesStatusCatalogCollationSQLLatin1GeneralCP1CIAS = DatabasePropertiesStatusCatalogCollation("SQL_Latin1_General_CP1_CI_AS")
)

type DatabasePropertiesStatusCreateMode string

const (
	DatabasePropertiesStatusCreateModeCopy                           = DatabasePropertiesStatusCreateMode("Copy")
	DatabasePropertiesStatusCreateModeDefault                        = DatabasePropertiesStatusCreateMode("Default")
	DatabasePropertiesStatusCreateModeOnlineSecondary                = DatabasePropertiesStatusCreateMode("OnlineSecondary")
	DatabasePropertiesStatusCreateModePointInTimeRestore             = DatabasePropertiesStatusCreateMode("PointInTimeRestore")
	DatabasePropertiesStatusCreateModeRecovery                       = DatabasePropertiesStatusCreateMode("Recovery")
	DatabasePropertiesStatusCreateModeRestore                        = DatabasePropertiesStatusCreateMode("Restore")
	DatabasePropertiesStatusCreateModeRestoreExternalBackup          = DatabasePropertiesStatusCreateMode("RestoreExternalBackup")
	DatabasePropertiesStatusCreateModeRestoreExternalBackupSecondary = DatabasePropertiesStatusCreateMode("RestoreExternalBackupSecondary")
	DatabasePropertiesStatusCreateModeRestoreLongTermRetentionBackup = DatabasePropertiesStatusCreateMode("RestoreLongTermRetentionBackup")
	DatabasePropertiesStatusCreateModeSecondary                      = DatabasePropertiesStatusCreateMode("Secondary")
)

type DatabasePropertiesStatusCurrentBackupStorageRedundancy string

const (
	DatabasePropertiesStatusCurrentBackupStorageRedundancyGeo   = DatabasePropertiesStatusCurrentBackupStorageRedundancy("Geo")
	DatabasePropertiesStatusCurrentBackupStorageRedundancyLocal = DatabasePropertiesStatusCurrentBackupStorageRedundancy("Local")
	DatabasePropertiesStatusCurrentBackupStorageRedundancyZone  = DatabasePropertiesStatusCurrentBackupStorageRedundancy("Zone")
)

type DatabasePropertiesStatusLicenseType string

const (
	DatabasePropertiesStatusLicenseTypeBasePrice       = DatabasePropertiesStatusLicenseType("BasePrice")
	DatabasePropertiesStatusLicenseTypeLicenseIncluded = DatabasePropertiesStatusLicenseType("LicenseIncluded")
)

type DatabasePropertiesStatusReadScale string

const (
	DatabasePropertiesStatusReadScaleDisabled = DatabasePropertiesStatusReadScale("Disabled")
	DatabasePropertiesStatusReadScaleEnabled  = DatabasePropertiesStatusReadScale("Enabled")
)

type DatabasePropertiesStatusRequestedBackupStorageRedundancy string

const (
	DatabasePropertiesStatusRequestedBackupStorageRedundancyGeo   = DatabasePropertiesStatusRequestedBackupStorageRedundancy("Geo")
	DatabasePropertiesStatusRequestedBackupStorageRedundancyLocal = DatabasePropertiesStatusRequestedBackupStorageRedundancy("Local")
	DatabasePropertiesStatusRequestedBackupStorageRedundancyZone  = DatabasePropertiesStatusRequestedBackupStorageRedundancy("Zone")
)

type DatabasePropertiesStatusSampleName string

const (
	DatabasePropertiesStatusSampleNameAdventureWorksLT       = DatabasePropertiesStatusSampleName("AdventureWorksLT")
	DatabasePropertiesStatusSampleNameWideWorldImportersFull = DatabasePropertiesStatusSampleName("WideWorldImportersFull")
	DatabasePropertiesStatusSampleNameWideWorldImportersStd  = DatabasePropertiesStatusSampleName("WideWorldImportersStd")
)

type DatabasePropertiesStatusSecondaryType string

const (
	DatabasePropertiesStatusSecondaryTypeGeo   = DatabasePropertiesStatusSecondaryType("Geo")
	DatabasePropertiesStatusSecondaryTypeNamed = DatabasePropertiesStatusSecondaryType("Named")
)

type DatabasePropertiesStatusStatus string

const (
	DatabasePropertiesStatusStatusAutoClosed                        = DatabasePropertiesStatusStatus("AutoClosed")
	DatabasePropertiesStatusStatusCopying                           = DatabasePropertiesStatusStatus("Copying")
	DatabasePropertiesStatusStatusCreating                          = DatabasePropertiesStatusStatus("Creating")
	DatabasePropertiesStatusStatusDisabled                          = DatabasePropertiesStatusStatus("Disabled")
	DatabasePropertiesStatusStatusEmergencyMode                     = DatabasePropertiesStatusStatus("EmergencyMode")
	DatabasePropertiesStatusStatusInaccessible                      = DatabasePropertiesStatusStatus("Inaccessible")
	DatabasePropertiesStatusStatusOffline                           = DatabasePropertiesStatusStatus("Offline")
	DatabasePropertiesStatusStatusOfflineChangingDwPerformanceTiers = DatabasePropertiesStatusStatus("OfflineChangingDwPerformanceTiers")
	DatabasePropertiesStatusStatusOfflineSecondary                  = DatabasePropertiesStatusStatus("OfflineSecondary")
	DatabasePropertiesStatusStatusOnline                            = DatabasePropertiesStatusStatus("Online")
	DatabasePropertiesStatusStatusOnlineChangingDwPerformanceTiers  = DatabasePropertiesStatusStatus("OnlineChangingDwPerformanceTiers")
	DatabasePropertiesStatusStatusPaused                            = DatabasePropertiesStatusStatus("Paused")
	DatabasePropertiesStatusStatusPausing                           = DatabasePropertiesStatusStatus("Pausing")
	DatabasePropertiesStatusStatusRecovering                        = DatabasePropertiesStatusStatus("Recovering")
	DatabasePropertiesStatusStatusRecoveryPending                   = DatabasePropertiesStatusStatus("RecoveryPending")
	DatabasePropertiesStatusStatusRestoring                         = DatabasePropertiesStatusStatus("Restoring")
	DatabasePropertiesStatusStatusResuming                          = DatabasePropertiesStatusStatus("Resuming")
	DatabasePropertiesStatusStatusScaling                           = DatabasePropertiesStatusStatus("Scaling")
	DatabasePropertiesStatusStatusShutdown                          = DatabasePropertiesStatusStatus("Shutdown")
	DatabasePropertiesStatusStatusStandby                           = DatabasePropertiesStatusStatus("Standby")
	DatabasePropertiesStatusStatusSuspect                           = DatabasePropertiesStatusStatus("Suspect")
)

//Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/definitions/Sku
type Sku struct {
	//Capacity: Capacity of the particular SKU.
	Capacity *int `json:"capacity,omitempty"`

	//Family: If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`

	// +kubebuilder:validation:Required
	//Name: The name of the SKU, typically, a letter + Number code, e.g. P3.
	Name *string `json:"name,omitempty"`

	//Size: Size of the particular SKU
	Size *string `json:"size,omitempty"`

	//Tier: The tier or edition of the particular SKU, e.g. Basic, Premium.
	Tier *string `json:"tier,omitempty"`
}

type Sku_Status struct {
	//Capacity: Capacity of the particular SKU.
	Capacity *int `json:"capacity,omitempty"`

	//Family: If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`

	//Name: The name of the SKU, typically, a letter + Number code, e.g. P3.
	Name *string `json:"name,omitempty"`

	//Size: Size of the particular SKU
	Size *string `json:"size,omitempty"`

	//Tier: The tier or edition of the particular SKU, e.g. Basic, Premium.
	Tier *string `json:"tier,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ServersDatabase{}, &ServersDatabaseList{})
}
