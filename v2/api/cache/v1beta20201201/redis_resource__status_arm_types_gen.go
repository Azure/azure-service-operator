// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

type RedisResource_StatusARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Redis cache properties.
	Properties *RedisProperties_StatusARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type RedisProperties_StatusARM struct {
	// EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`

	// HostName: Redis host name.
	HostName *string `json:"hostName,omitempty"`

	// Instances: List of the Redis instances associated with the cache
	Instances []RedisInstanceDetails_StatusARM `json:"instances,omitempty"`

	// LinkedServers: List of the linked servers associated with the cache
	LinkedServers []RedisLinkedServer_StatusARM `json:"linkedServers,omitempty"`

	// MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1',
	// '1.2')
	MinimumTlsVersion *RedisPropertiesStatusMinimumTlsVersion `json:"minimumTlsVersion,omitempty"`

	// Port: Redis non-SSL port.
	Port *int `json:"port,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connection associated with the specified redis cache
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Redis instance provisioning status.
	ProvisioningState *RedisPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is
	// 'Enabled'
	PublicNetworkAccess *RedisPropertiesStatusPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// RedisConfiguration: All Redis Settings. Few possible keys:
	// rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
	// etc.
	RedisConfiguration map[string]string `json:"redisConfiguration,omitempty"`

	// RedisVersion: Redis version. Only major version will be used in PUT/PATCH request with current valid values: (4, 6)
	RedisVersion *string `json:"redisVersion,omitempty"`

	// ReplicasPerMaster: The number of replicas to be created per primary.
	ReplicasPerMaster *int `json:"replicasPerMaster,omitempty"`

	// ReplicasPerPrimary: The number of replicas to be created per primary.
	ReplicasPerPrimary *int `json:"replicasPerPrimary,omitempty"`

	// ShardCount: The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `json:"shardCount,omitempty"`

	// Sku: The SKU of the Redis cache to deploy.
	Sku *Sku_StatusARM `json:"sku,omitempty"`

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

type PrivateEndpointConnection_Status_SubResourceEmbeddedARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`
}

type RedisInstanceDetails_StatusARM struct {
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

type RedisLinkedServer_StatusARM struct {
	// Id: Linked server Id.
	Id *string `json:"id,omitempty"`
}

type Sku_StatusARM struct {
	// Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
	// P (Premium) family (1, 2, 3, 4).
	Capacity *int `json:"capacity,omitempty"`

	// Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family *SkuStatusFamily `json:"family,omitempty"`

	// Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
	Name *SkuStatusName `json:"name,omitempty"`
}
