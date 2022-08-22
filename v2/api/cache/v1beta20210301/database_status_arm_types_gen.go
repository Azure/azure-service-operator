// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301

type Database_STATUSARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Other properties of the database.
	Properties *DatabaseProperties_STATUSARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type DatabaseProperties_STATUSARM struct {
	// ClientProtocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is
	// TLS-encrypted.
	ClientProtocol *DatabasePropertiesSTATUSClientProtocol `json:"clientProtocol,omitempty"`

	// ClusteringPolicy: Clustering policy - default is OSSCluster. Specified at create time.
	ClusteringPolicy *DatabasePropertiesSTATUSClusteringPolicy `json:"clusteringPolicy,omitempty"`

	// EvictionPolicy: Redis eviction policy - default is VolatileLRU
	EvictionPolicy *DatabasePropertiesSTATUSEvictionPolicy `json:"evictionPolicy,omitempty"`

	// Modules: Optional set of redis modules to enable in this database - modules can only be added at creation time.
	Modules []Module_STATUSARM `json:"modules,omitempty"`

	// Persistence: Persistence settings
	Persistence *Persistence_STATUSARM `json:"persistence,omitempty"`

	// Port: TCP port of the database endpoint. Specified at create time. Defaults to an available port.
	Port *int `json:"port,omitempty"`

	// ProvisioningState: Current provisioning status of the database
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceState: Current resource status of the database
	ResourceState *ResourceState_STATUS `json:"resourceState,omitempty"`
}

type DatabasePropertiesSTATUSClientProtocol string

const (
	DatabasePropertiesSTATUSClientProtocol_Encrypted = DatabasePropertiesSTATUSClientProtocol("Encrypted")
	DatabasePropertiesSTATUSClientProtocol_Plaintext = DatabasePropertiesSTATUSClientProtocol("Plaintext")
)

type DatabasePropertiesSTATUSClusteringPolicy string

const (
	DatabasePropertiesSTATUSClusteringPolicy_EnterpriseCluster = DatabasePropertiesSTATUSClusteringPolicy("EnterpriseCluster")
	DatabasePropertiesSTATUSClusteringPolicy_OSSCluster        = DatabasePropertiesSTATUSClusteringPolicy("OSSCluster")
)

type DatabasePropertiesSTATUSEvictionPolicy string

const (
	DatabasePropertiesSTATUSEvictionPolicy_AllKeysLFU     = DatabasePropertiesSTATUSEvictionPolicy("AllKeysLFU")
	DatabasePropertiesSTATUSEvictionPolicy_AllKeysLRU     = DatabasePropertiesSTATUSEvictionPolicy("AllKeysLRU")
	DatabasePropertiesSTATUSEvictionPolicy_AllKeysRandom  = DatabasePropertiesSTATUSEvictionPolicy("AllKeysRandom")
	DatabasePropertiesSTATUSEvictionPolicy_NoEviction     = DatabasePropertiesSTATUSEvictionPolicy("NoEviction")
	DatabasePropertiesSTATUSEvictionPolicy_VolatileLFU    = DatabasePropertiesSTATUSEvictionPolicy("VolatileLFU")
	DatabasePropertiesSTATUSEvictionPolicy_VolatileLRU    = DatabasePropertiesSTATUSEvictionPolicy("VolatileLRU")
	DatabasePropertiesSTATUSEvictionPolicy_VolatileRandom = DatabasePropertiesSTATUSEvictionPolicy("VolatileRandom")
	DatabasePropertiesSTATUSEvictionPolicy_VolatileTTL    = DatabasePropertiesSTATUSEvictionPolicy("VolatileTTL")
)

type Module_STATUSARM struct {
	// Args: Configuration options for the module, e.g. 'ERROR_RATE 0.00 INITIAL_SIZE 400'.
	Args *string `json:"args,omitempty"`

	// Name: The name of the module, e.g. 'RedisBloom', 'RediSearch', 'RedisTimeSeries'
	Name *string `json:"name,omitempty"`

	// Version: The version of the module, e.g. '1.0'.
	Version *string `json:"version,omitempty"`
}

type Persistence_STATUSARM struct {
	// AofEnabled: Sets whether AOF is enabled.
	AofEnabled *bool `json:"aofEnabled,omitempty"`

	// AofFrequency: Sets the frequency at which data is written to disk.
	AofFrequency *PersistenceSTATUSAofFrequency `json:"aofFrequency,omitempty"`

	// RdbEnabled: Sets whether RDB is enabled.
	RdbEnabled *bool `json:"rdbEnabled,omitempty"`

	// RdbFrequency: Sets the frequency at which a snapshot of the database is created.
	RdbFrequency *PersistenceSTATUSRdbFrequency `json:"rdbFrequency,omitempty"`
}

type PersistenceSTATUSAofFrequency string

const (
	PersistenceSTATUSAofFrequency_1S     = PersistenceSTATUSAofFrequency("1s")
	PersistenceSTATUSAofFrequency_Always = PersistenceSTATUSAofFrequency("always")
)

type PersistenceSTATUSRdbFrequency string

const (
	PersistenceSTATUSRdbFrequency_12H = PersistenceSTATUSRdbFrequency("12h")
	PersistenceSTATUSRdbFrequency_1H  = PersistenceSTATUSRdbFrequency("1h")
	PersistenceSTATUSRdbFrequency_6H  = PersistenceSTATUSRdbFrequency("6h")
)
