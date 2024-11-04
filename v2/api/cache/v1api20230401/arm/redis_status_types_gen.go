// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type Redis_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: The identity of the resource.
	Identity *ManagedServiceIdentity_STATUS `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Redis cache properties.
	Properties *RedisProperties_STATUS `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity_STATUS struct {
	// PrincipalId: The service principal ID of the system assigned identity. This property will only be provided for a system
	// assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type                   *ManagedServiceIdentityType_STATUS     `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Properties of the redis cache.
type RedisProperties_STATUS struct {
	// EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`

	// HostName: Redis host name.
	HostName *string `json:"hostName,omitempty"`

	// Instances: List of the Redis instances associated with the cache
	Instances []RedisInstanceDetails_STATUS `json:"instances,omitempty"`

	// LinkedServers: List of the linked servers associated with the cache
	LinkedServers []RedisLinkedServer_STATUS `json:"linkedServers,omitempty"`

	// MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1',
	// '1.2')
	MinimumTlsVersion *RedisProperties_MinimumTlsVersion_STATUS `json:"minimumTlsVersion,omitempty"`

	// Port: Redis non-SSL port.
	Port *int `json:"port,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connection associated with the specified redis cache
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Redis instance provisioning status.
	ProvisioningState *RedisProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional, but if passed
	// in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is
	// 'Enabled'. Note: This setting is important for caches with private endpoints. It has *no effect* on caches that are
	// joined to, or injected into, a virtual network subnet.
	PublicNetworkAccess *RedisProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// RedisConfiguration: All Redis Settings. Few possible keys:
	// rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
	// etc.
	RedisConfiguration *RedisProperties_RedisConfiguration_STATUS `json:"redisConfiguration,omitempty"`

	// RedisVersion: Redis version. This should be in the form 'major[.minor]' (only 'major' is required) or the value 'latest'
	// which refers to the latest stable Redis version that is available. Supported versions: 4.0, 6.0 (latest). Default value
	// is 'latest'.
	RedisVersion *string `json:"redisVersion,omitempty"`

	// ReplicasPerMaster: The number of replicas to be created per primary.
	ReplicasPerMaster *int `json:"replicasPerMaster,omitempty"`

	// ReplicasPerPrimary: The number of replicas to be created per primary.
	ReplicasPerPrimary *int `json:"replicasPerPrimary,omitempty"`

	// ShardCount: The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `json:"shardCount,omitempty"`

	// Sku: The SKU of the Redis cache to deploy.
	Sku *Sku_STATUS `json:"sku,omitempty"`

	// SslPort: Redis SSL port.
	SslPort *int `json:"sslPort,omitempty"`

	// StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
	// Network; auto assigned by default.
	StaticIP *string `json:"staticIP,omitempty"`

	// SubnetId: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId *string `json:"subnetId,omitempty"`

	// TenantSettings: A dictionary of tenant settings
	TenantSettings map[string]string `json:"tenantSettings,omitempty"`
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType_STATUS string

const (
	ManagedServiceIdentityType_STATUS_None                       = ManagedServiceIdentityType_STATUS("None")
	ManagedServiceIdentityType_STATUS_SystemAssigned             = ManagedServiceIdentityType_STATUS("SystemAssigned")
	ManagedServiceIdentityType_STATUS_SystemAssignedUserAssigned = ManagedServiceIdentityType_STATUS("SystemAssigned, UserAssigned")
	ManagedServiceIdentityType_STATUS_UserAssigned               = ManagedServiceIdentityType_STATUS("UserAssigned")
)

// Mapping from string to ManagedServiceIdentityType_STATUS
var managedServiceIdentityType_STATUS_Values = map[string]ManagedServiceIdentityType_STATUS{
	"none":                         ManagedServiceIdentityType_STATUS_None,
	"systemassigned":               ManagedServiceIdentityType_STATUS_SystemAssigned,
	"systemassigned, userassigned": ManagedServiceIdentityType_STATUS_SystemAssignedUserAssigned,
	"userassigned":                 ManagedServiceIdentityType_STATUS_UserAssigned,
}

// The Private Endpoint Connection resource.
type PrivateEndpointConnection_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`
}

// Details of single instance of redis.
type RedisInstanceDetails_STATUS struct {
	// IsMaster: Specifies whether the instance is a primary node.
	IsMaster *bool `json:"isMaster,omitempty"`

	// IsPrimary: Specifies whether the instance is a primary node.
	IsPrimary *bool `json:"isPrimary,omitempty"`

	// NonSslPort: If enableNonSslPort is true, provides Redis instance Non-SSL port.
	NonSslPort *int `json:"nonSslPort,omitempty"`

	// ShardId: If clustering is enabled, the Shard ID of Redis Instance
	ShardId *int `json:"shardId,omitempty"`

	// SslPort: Redis instance SSL port.
	SslPort *int `json:"sslPort,omitempty"`

	// Zone: If the Cache uses availability zones, specifies availability zone where this instance is located.
	Zone *string `json:"zone,omitempty"`
}

// Linked server Id
type RedisLinkedServer_STATUS struct {
	// Id: Linked server Id.
	Id *string `json:"id,omitempty"`
}

type RedisProperties_MinimumTlsVersion_STATUS string

const (
	RedisProperties_MinimumTlsVersion_STATUS_10 = RedisProperties_MinimumTlsVersion_STATUS("1.0")
	RedisProperties_MinimumTlsVersion_STATUS_11 = RedisProperties_MinimumTlsVersion_STATUS("1.1")
	RedisProperties_MinimumTlsVersion_STATUS_12 = RedisProperties_MinimumTlsVersion_STATUS("1.2")
)

// Mapping from string to RedisProperties_MinimumTlsVersion_STATUS
var redisProperties_MinimumTlsVersion_STATUS_Values = map[string]RedisProperties_MinimumTlsVersion_STATUS{
	"1.0": RedisProperties_MinimumTlsVersion_STATUS_10,
	"1.1": RedisProperties_MinimumTlsVersion_STATUS_11,
	"1.2": RedisProperties_MinimumTlsVersion_STATUS_12,
}

type RedisProperties_ProvisioningState_STATUS string

const (
	RedisProperties_ProvisioningState_STATUS_Creating               = RedisProperties_ProvisioningState_STATUS("Creating")
	RedisProperties_ProvisioningState_STATUS_Deleting               = RedisProperties_ProvisioningState_STATUS("Deleting")
	RedisProperties_ProvisioningState_STATUS_Disabled               = RedisProperties_ProvisioningState_STATUS("Disabled")
	RedisProperties_ProvisioningState_STATUS_Failed                 = RedisProperties_ProvisioningState_STATUS("Failed")
	RedisProperties_ProvisioningState_STATUS_Linking                = RedisProperties_ProvisioningState_STATUS("Linking")
	RedisProperties_ProvisioningState_STATUS_Provisioning           = RedisProperties_ProvisioningState_STATUS("Provisioning")
	RedisProperties_ProvisioningState_STATUS_RecoveringScaleFailure = RedisProperties_ProvisioningState_STATUS("RecoveringScaleFailure")
	RedisProperties_ProvisioningState_STATUS_Scaling                = RedisProperties_ProvisioningState_STATUS("Scaling")
	RedisProperties_ProvisioningState_STATUS_Succeeded              = RedisProperties_ProvisioningState_STATUS("Succeeded")
	RedisProperties_ProvisioningState_STATUS_Unlinking              = RedisProperties_ProvisioningState_STATUS("Unlinking")
	RedisProperties_ProvisioningState_STATUS_Unprovisioning         = RedisProperties_ProvisioningState_STATUS("Unprovisioning")
	RedisProperties_ProvisioningState_STATUS_Updating               = RedisProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to RedisProperties_ProvisioningState_STATUS
var redisProperties_ProvisioningState_STATUS_Values = map[string]RedisProperties_ProvisioningState_STATUS{
	"creating":               RedisProperties_ProvisioningState_STATUS_Creating,
	"deleting":               RedisProperties_ProvisioningState_STATUS_Deleting,
	"disabled":               RedisProperties_ProvisioningState_STATUS_Disabled,
	"failed":                 RedisProperties_ProvisioningState_STATUS_Failed,
	"linking":                RedisProperties_ProvisioningState_STATUS_Linking,
	"provisioning":           RedisProperties_ProvisioningState_STATUS_Provisioning,
	"recoveringscalefailure": RedisProperties_ProvisioningState_STATUS_RecoveringScaleFailure,
	"scaling":                RedisProperties_ProvisioningState_STATUS_Scaling,
	"succeeded":              RedisProperties_ProvisioningState_STATUS_Succeeded,
	"unlinking":              RedisProperties_ProvisioningState_STATUS_Unlinking,
	"unprovisioning":         RedisProperties_ProvisioningState_STATUS_Unprovisioning,
	"updating":               RedisProperties_ProvisioningState_STATUS_Updating,
}

type RedisProperties_PublicNetworkAccess_STATUS string

const (
	RedisProperties_PublicNetworkAccess_STATUS_Disabled = RedisProperties_PublicNetworkAccess_STATUS("Disabled")
	RedisProperties_PublicNetworkAccess_STATUS_Enabled  = RedisProperties_PublicNetworkAccess_STATUS("Enabled")
)

// Mapping from string to RedisProperties_PublicNetworkAccess_STATUS
var redisProperties_PublicNetworkAccess_STATUS_Values = map[string]RedisProperties_PublicNetworkAccess_STATUS{
	"disabled": RedisProperties_PublicNetworkAccess_STATUS_Disabled,
	"enabled":  RedisProperties_PublicNetworkAccess_STATUS_Enabled,
}

type RedisProperties_RedisConfiguration_STATUS struct {
	// AofBackupEnabled: Specifies whether the aof backup is enabled
	AofBackupEnabled *string `json:"aof-backup-enabled,omitempty"`

	// AofStorageConnectionString0: First storage account connection string
	AofStorageConnectionString0 *string `json:"aof-storage-connection-string-0,omitempty"`

	// AofStorageConnectionString1: Second storage account connection string
	AofStorageConnectionString1 *string `json:"aof-storage-connection-string-1,omitempty"`

	// Authnotrequired: Specifies whether the authentication is disabled. Setting this property is highly discouraged from
	// security point of view.
	Authnotrequired *string `json:"authnotrequired,omitempty"`

	// Maxclients: The max clients config
	Maxclients *string `json:"maxclients,omitempty"`

	// MaxfragmentationmemoryReserved: Value in megabytes reserved for fragmentation per shard
	MaxfragmentationmemoryReserved *string `json:"maxfragmentationmemory-reserved,omitempty"`

	// MaxmemoryDelta: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryDelta *string `json:"maxmemory-delta,omitempty"`

	// MaxmemoryPolicy: The eviction strategy used when your data won't fit within its memory limit.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	// MaxmemoryReserved: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryReserved *string `json:"maxmemory-reserved,omitempty"`

	// PreferredDataArchiveAuthMethod: Preferred auth method to communicate to storage account used for data archive, specify
	// SAS or ManagedIdentity, default value is SAS
	PreferredDataArchiveAuthMethod *string `json:"preferred-data-archive-auth-method,omitempty"`

	// PreferredDataPersistenceAuthMethod: Preferred auth method to communicate to storage account used for data persistence,
	// specify SAS or ManagedIdentity, default value is SAS
	PreferredDataPersistenceAuthMethod *string `json:"preferred-data-persistence-auth-method,omitempty"`

	// RdbBackupEnabled: Specifies whether the rdb backup is enabled
	RdbBackupEnabled *string `json:"rdb-backup-enabled,omitempty"`

	// RdbBackupFrequency: Specifies the frequency for creating rdb backup in minutes. Valid values: (15, 30, 60, 360, 720,
	// 1440)
	RdbBackupFrequency *string `json:"rdb-backup-frequency,omitempty"`

	// RdbBackupMaxSnapshotCount: Specifies the maximum number of snapshots for rdb backup
	RdbBackupMaxSnapshotCount *string `json:"rdb-backup-max-snapshot-count,omitempty"`

	// RdbStorageConnectionString: The storage account connection string for storing rdb file
	RdbStorageConnectionString *string `json:"rdb-storage-connection-string,omitempty"`

	// StorageSubscriptionId: SubscriptionId of the storage account for persistence (aof/rdb) using ManagedIdentity.
	StorageSubscriptionId *string `json:"storage-subscription-id,omitempty"`

	// ZonalConfiguration: Zonal Configuration
	ZonalConfiguration *string `json:"zonal-configuration,omitempty"`
}

// SKU parameters supplied to the create Redis operation.
type Sku_STATUS struct {
	// Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
	// P (Premium) family (1, 2, 3, 4).
	Capacity *int `json:"capacity,omitempty"`

	// Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family *Sku_Family_STATUS `json:"family,omitempty"`

	// Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
	Name *Sku_Name_STATUS `json:"name,omitempty"`
}

// User assigned identity properties
type UserAssignedIdentity_STATUS struct {
	// ClientId: The client ID of the assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal ID of the assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

type Sku_Family_STATUS string

const (
	Sku_Family_STATUS_C = Sku_Family_STATUS("C")
	Sku_Family_STATUS_P = Sku_Family_STATUS("P")
)

// Mapping from string to Sku_Family_STATUS
var sku_Family_STATUS_Values = map[string]Sku_Family_STATUS{
	"c": Sku_Family_STATUS_C,
	"p": Sku_Family_STATUS_P,
}

type Sku_Name_STATUS string

const (
	Sku_Name_STATUS_Basic    = Sku_Name_STATUS("Basic")
	Sku_Name_STATUS_Premium  = Sku_Name_STATUS("Premium")
	Sku_Name_STATUS_Standard = Sku_Name_STATUS("Standard")
)

// Mapping from string to Sku_Name_STATUS
var sku_Name_STATUS_Values = map[string]Sku_Name_STATUS{
	"basic":    Sku_Name_STATUS_Basic,
	"premium":  Sku_Name_STATUS_Premium,
	"standard": Sku_Name_STATUS_Standard,
}