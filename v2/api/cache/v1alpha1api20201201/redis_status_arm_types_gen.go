// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

// Deprecated version of Redis_STATUS. Use v1beta20201201.Redis_STATUS instead
type Redis_STATUS_ARM struct {
	Id         *string                     `json:"id,omitempty"`
	Location   *string                     `json:"location,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Properties *RedisProperties_STATUS_ARM `json:"properties,omitempty"`
	Tags       map[string]string           `json:"tags,omitempty"`
	Type       *string                     `json:"type,omitempty"`
	Zones      []string                    `json:"zones,omitempty"`
}

// Deprecated version of RedisProperties_STATUS. Use v1beta20201201.RedisProperties_STATUS instead
type RedisProperties_STATUS_ARM struct {
	EnableNonSslPort           *bool                                          `json:"enableNonSslPort,omitempty"`
	HostName                   *string                                        `json:"hostName,omitempty"`
	Instances                  []RedisInstanceDetails_STATUS_ARM              `json:"instances,omitempty"`
	LinkedServers              []RedisLinkedServer_STATUS_ARM                 `json:"linkedServers,omitempty"`
	MinimumTlsVersion          *RedisProperties_MinimumTlsVersion_STATUS      `json:"minimumTlsVersion,omitempty"`
	Port                       *int                                           `json:"port,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM         `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *RedisProperties_ProvisioningState_STATUS      `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *RedisProperties_PublicNetworkAccess_STATUS    `json:"publicNetworkAccess,omitempty"`
	RedisConfiguration         *RedisProperties_RedisConfiguration_STATUS_ARM `json:"redisConfiguration,omitempty"`
	RedisVersion               *string                                        `json:"redisVersion,omitempty"`
	ReplicasPerMaster          *int                                           `json:"replicasPerMaster,omitempty"`
	ReplicasPerPrimary         *int                                           `json:"replicasPerPrimary,omitempty"`
	ShardCount                 *int                                           `json:"shardCount,omitempty"`
	Sku                        *Sku_STATUS_ARM                                `json:"sku,omitempty"`
	SslPort                    *int                                           `json:"sslPort,omitempty"`
	StaticIP                   *string                                        `json:"staticIP,omitempty"`
	SubnetId                   *string                                        `json:"subnetId,omitempty"`
	TenantSettings             map[string]string                              `json:"tenantSettings,omitempty"`
}

// Deprecated version of PrivateEndpointConnection_STATUS. Use v1beta20201201.PrivateEndpointConnection_STATUS instead
type PrivateEndpointConnection_STATUS_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of RedisInstanceDetails_STATUS. Use v1beta20201201.RedisInstanceDetails_STATUS instead
type RedisInstanceDetails_STATUS_ARM struct {
	IsMaster   *bool   `json:"isMaster,omitempty"`
	IsPrimary  *bool   `json:"isPrimary,omitempty"`
	NonSslPort *int    `json:"nonSslPort,omitempty"`
	ShardId    *int    `json:"shardId,omitempty"`
	SslPort    *int    `json:"sslPort,omitempty"`
	Zone       *string `json:"zone,omitempty"`
}

// Deprecated version of RedisLinkedServer_STATUS. Use v1beta20201201.RedisLinkedServer_STATUS instead
type RedisLinkedServer_STATUS_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of RedisProperties_RedisConfiguration_STATUS. Use v1beta20201201.RedisProperties_RedisConfiguration_STATUS instead
type RedisProperties_RedisConfiguration_STATUS_ARM struct {
	AdditionalProperties           map[string]string `json:"additionalProperties,omitempty"`
	AofBackupEnabled               *string           `json:"aof-backup-enabled,omitempty"`
	AofStorageConnectionString0    *string           `json:"aof-storage-connection-string-0,omitempty"`
	AofStorageConnectionString1    *string           `json:"aof-storage-connection-string-1,omitempty"`
	Authnotrequired                *string           `json:"authnotrequired,omitempty"`
	Maxclients                     *string           `json:"maxclients,omitempty"`
	MaxfragmentationmemoryReserved *string           `json:"maxfragmentationmemory-reserved,omitempty"`
	MaxmemoryDelta                 *string           `json:"maxmemory-delta,omitempty"`
	MaxmemoryPolicy                *string           `json:"maxmemory-policy,omitempty"`
	MaxmemoryReserved              *string           `json:"maxmemory-reserved,omitempty"`
	RdbBackupEnabled               *string           `json:"rdb-backup-enabled,omitempty"`
	RdbBackupFrequency             *string           `json:"rdb-backup-frequency,omitempty"`
	RdbBackupMaxSnapshotCount      *string           `json:"rdb-backup-max-snapshot-count,omitempty"`
	RdbStorageConnectionString     *string           `json:"rdb-storage-connection-string,omitempty"`
	ZonalConfiguration             *string           `json:"zonal-configuration,omitempty"`
}

// Deprecated version of Sku_STATUS. Use v1beta20201201.Sku_STATUS instead
type Sku_STATUS_ARM struct {
	Capacity *int               `json:"capacity,omitempty"`
	Family   *Sku_Family_STATUS `json:"family,omitempty"`
	Name     *Sku_Name_STATUS   `json:"name,omitempty"`
}
