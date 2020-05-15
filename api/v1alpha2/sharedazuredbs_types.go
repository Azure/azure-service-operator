// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
)

type AzureDBsSQLSku struct {
	// Name - The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.
	Name string `json:"name,omitempty"`
	// Tier - The tier of the particular SKU, e.g. Basic. Possible values include: 'Basic', 'GeneralPurpose', 'MemoryOptimized'
	Tier SkuTier `json:"tier,omitempty"`
	// Capacity - The scale up/out capacity, representing server's compute units.
	Capacity int32 `json:"capacity,omitempty"`
	// Size - The size code, to be interpreted by resource as appropriate.
	Size string `json:"size,omitempty"`
	// Family - The family of hardware.
	Family string `json:"family,omitempty"`
}

// ServerVersion enumerates the values for server version.
type ServerVersion string

const (
	// NineFullStopFive ...
	NineFullStopFive ServerVersion = "9.5"
	// NineFullStopSix ...
	NineFullStopSix ServerVersion = "9.6"
	// OneOne ...
	OneOne ServerVersion = "11"
	// OneZero ...
	OneZero ServerVersion = "10"
	// OneZeroFullStopTwo ...
	OneZeroFullStopTwo ServerVersion = "10.2"
	// OneZeroFullStopZero ...
	OneZeroFullStopZero ServerVersion = "10.0"
)

type SkuTier string

const (
	// Basic ...
	PSQLBasic SkuTier = "Basic"
	// GeneralPurpose ...
	PSQLGeneralPurpose SkuTier = "GeneralPurpose"
	// MemoryOptimized ...
	PSQLMemoryOptimized SkuTier = "MemoryOptimized"
)

type SslEnforcementEnum string

const (
	// SslEnforcementEnumDisabled ...
	SslEnforcementEnumDisabled SslEnforcementEnum = "Disabled"
	// SslEnforcementEnumEnabled ...
	SslEnforcementEnumEnabled SslEnforcementEnum = "Enabled"
)

type MySQLStorageProfile struct {
	// BackupRetentionDays - Backup retention days for the server.
	BackupRetentionDays *int32 `json:"backupRetentionDays,omitempty"`
	// GeoRedundantBackup - Enable Geo-redundant or not for server backup. Possible values include: 'Enabled', 'Disabled'
	GeoRedundantBackup mysql.GeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
	// StorageMB - Max storage allowed for a server.
	StorageMB *int32 `json:"storageMB,omitempty"`
	// StorageAutogrow - Enable Storage Auto Grow. Possible values include: 'StorageAutogrowEnabled', 'StorageAutogrowDisabled'
	StorageAutogrow mysql.StorageAutogrow `json:"storageAutogrow,omitempty"`
}

type PSQLStorageProfile struct {
	// BackupRetentionDays - Backup retention days for the server.
	BackupRetentionDays *int32 `json:"backupRetentionDays,omitempty"`
	// GeoRedundantBackup - Enable Geo-redundant or not for server backup. Possible values include: 'Enabled', 'Disabled'
	GeoRedundantBackup psql.GeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
	// StorageMB - Max storage allowed for a server.
	StorageMB *int32 `json:"storageMB,omitempty"`
	// StorageAutogrow - Enable Storage Auto Grow. Possible values include: 'StorageAutogrowEnabled', 'StorageAutogrowDisabled'
	StorageAutogrow psql.StorageAutogrow `json:"storageAutogrow,omitempty"`
}

type ReplicaProperties struct {
	SourceServerId string `json:"sourceServerId,omitempty"`
}
