// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisEnterpriseDatabase_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Other properties of the database.
	Properties *DatabaseProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisEnterpriseDatabase_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (database RedisEnterpriseDatabase_Spec_ARM) GetAPIVersion() string {
	return "2021-03-01"
}

// GetName returns the Name of the resource
func (database *RedisEnterpriseDatabase_Spec_ARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise/databases"
func (database *RedisEnterpriseDatabase_Spec_ARM) GetType() string {
	return "Microsoft.Cache/redisEnterprise/databases"
}

// Properties of RedisEnterprise databases, as opposed to general resource properties like location, tags
type DatabaseProperties_ARM struct {
	// ClientProtocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is
	// TLS-encrypted.
	ClientProtocol *DatabaseProperties_ClientProtocol_ARM `json:"clientProtocol,omitempty"`

	// ClusteringPolicy: Clustering policy - default is OSSCluster. Specified at create time.
	ClusteringPolicy *DatabaseProperties_ClusteringPolicy_ARM `json:"clusteringPolicy,omitempty"`

	// EvictionPolicy: Redis eviction policy - default is VolatileLRU
	EvictionPolicy *DatabaseProperties_EvictionPolicy_ARM `json:"evictionPolicy,omitempty"`

	// Modules: Optional set of redis modules to enable in this database - modules can only be added at creation time.
	Modules []Module_ARM `json:"modules,omitempty"`

	// Persistence: Persistence settings
	Persistence *Persistence_ARM `json:"persistence,omitempty"`

	// Port: TCP port of the database endpoint. Specified at create time. Defaults to an available port.
	Port *int `json:"port,omitempty"`
}

// +kubebuilder:validation:Enum={"Encrypted","Plaintext"}
type DatabaseProperties_ClientProtocol_ARM string

const (
	DatabaseProperties_ClientProtocol_ARM_Encrypted = DatabaseProperties_ClientProtocol_ARM("Encrypted")
	DatabaseProperties_ClientProtocol_ARM_Plaintext = DatabaseProperties_ClientProtocol_ARM("Plaintext")
)

// Mapping from string to DatabaseProperties_ClientProtocol_ARM
var databaseProperties_ClientProtocol_ARM_Values = map[string]DatabaseProperties_ClientProtocol_ARM{
	"encrypted": DatabaseProperties_ClientProtocol_ARM_Encrypted,
	"plaintext": DatabaseProperties_ClientProtocol_ARM_Plaintext,
}

// +kubebuilder:validation:Enum={"EnterpriseCluster","OSSCluster"}
type DatabaseProperties_ClusteringPolicy_ARM string

const (
	DatabaseProperties_ClusteringPolicy_ARM_EnterpriseCluster = DatabaseProperties_ClusteringPolicy_ARM("EnterpriseCluster")
	DatabaseProperties_ClusteringPolicy_ARM_OSSCluster        = DatabaseProperties_ClusteringPolicy_ARM("OSSCluster")
)

// Mapping from string to DatabaseProperties_ClusteringPolicy_ARM
var databaseProperties_ClusteringPolicy_ARM_Values = map[string]DatabaseProperties_ClusteringPolicy_ARM{
	"enterprisecluster": DatabaseProperties_ClusteringPolicy_ARM_EnterpriseCluster,
	"osscluster":        DatabaseProperties_ClusteringPolicy_ARM_OSSCluster,
}

// +kubebuilder:validation:Enum={"AllKeysLFU","AllKeysLRU","AllKeysRandom","NoEviction","VolatileLFU","VolatileLRU","VolatileRandom","VolatileTTL"}
type DatabaseProperties_EvictionPolicy_ARM string

const (
	DatabaseProperties_EvictionPolicy_ARM_AllKeysLFU     = DatabaseProperties_EvictionPolicy_ARM("AllKeysLFU")
	DatabaseProperties_EvictionPolicy_ARM_AllKeysLRU     = DatabaseProperties_EvictionPolicy_ARM("AllKeysLRU")
	DatabaseProperties_EvictionPolicy_ARM_AllKeysRandom  = DatabaseProperties_EvictionPolicy_ARM("AllKeysRandom")
	DatabaseProperties_EvictionPolicy_ARM_NoEviction     = DatabaseProperties_EvictionPolicy_ARM("NoEviction")
	DatabaseProperties_EvictionPolicy_ARM_VolatileLFU    = DatabaseProperties_EvictionPolicy_ARM("VolatileLFU")
	DatabaseProperties_EvictionPolicy_ARM_VolatileLRU    = DatabaseProperties_EvictionPolicy_ARM("VolatileLRU")
	DatabaseProperties_EvictionPolicy_ARM_VolatileRandom = DatabaseProperties_EvictionPolicy_ARM("VolatileRandom")
	DatabaseProperties_EvictionPolicy_ARM_VolatileTTL    = DatabaseProperties_EvictionPolicy_ARM("VolatileTTL")
)

// Mapping from string to DatabaseProperties_EvictionPolicy_ARM
var databaseProperties_EvictionPolicy_ARM_Values = map[string]DatabaseProperties_EvictionPolicy_ARM{
	"allkeyslfu":     DatabaseProperties_EvictionPolicy_ARM_AllKeysLFU,
	"allkeyslru":     DatabaseProperties_EvictionPolicy_ARM_AllKeysLRU,
	"allkeysrandom":  DatabaseProperties_EvictionPolicy_ARM_AllKeysRandom,
	"noeviction":     DatabaseProperties_EvictionPolicy_ARM_NoEviction,
	"volatilelfu":    DatabaseProperties_EvictionPolicy_ARM_VolatileLFU,
	"volatilelru":    DatabaseProperties_EvictionPolicy_ARM_VolatileLRU,
	"volatilerandom": DatabaseProperties_EvictionPolicy_ARM_VolatileRandom,
	"volatilettl":    DatabaseProperties_EvictionPolicy_ARM_VolatileTTL,
}

// Specifies configuration of a redis module
type Module_ARM struct {
	// Args: Configuration options for the module, e.g. 'ERROR_RATE 0.00 INITIAL_SIZE 400'.
	Args *string `json:"args,omitempty"`

	// Name: The name of the module, e.g. 'RedisBloom', 'RediSearch', 'RedisTimeSeries'
	Name *string `json:"name,omitempty"`
}

// Persistence-related configuration for the RedisEnterprise database
type Persistence_ARM struct {
	// AofEnabled: Sets whether AOF is enabled.
	AofEnabled *bool `json:"aofEnabled,omitempty"`

	// AofFrequency: Sets the frequency at which data is written to disk.
	AofFrequency *Persistence_AofFrequency_ARM `json:"aofFrequency,omitempty"`

	// RdbEnabled: Sets whether RDB is enabled.
	RdbEnabled *bool `json:"rdbEnabled,omitempty"`

	// RdbFrequency: Sets the frequency at which a snapshot of the database is created.
	RdbFrequency *Persistence_RdbFrequency_ARM `json:"rdbFrequency,omitempty"`
}

// +kubebuilder:validation:Enum={"1s","always"}
type Persistence_AofFrequency_ARM string

const (
	Persistence_AofFrequency_ARM_1S     = Persistence_AofFrequency_ARM("1s")
	Persistence_AofFrequency_ARM_Always = Persistence_AofFrequency_ARM("always")
)

// Mapping from string to Persistence_AofFrequency_ARM
var persistence_AofFrequency_ARM_Values = map[string]Persistence_AofFrequency_ARM{
	"1s":     Persistence_AofFrequency_ARM_1S,
	"always": Persistence_AofFrequency_ARM_Always,
}

// +kubebuilder:validation:Enum={"12h","1h","6h"}
type Persistence_RdbFrequency_ARM string

const (
	Persistence_RdbFrequency_ARM_12H = Persistence_RdbFrequency_ARM("12h")
	Persistence_RdbFrequency_ARM_1H  = Persistence_RdbFrequency_ARM("1h")
	Persistence_RdbFrequency_ARM_6H  = Persistence_RdbFrequency_ARM("6h")
)

// Mapping from string to Persistence_RdbFrequency_ARM
var persistence_RdbFrequency_ARM_Values = map[string]Persistence_RdbFrequency_ARM{
	"12h": Persistence_RdbFrequency_ARM_12H,
	"1h":  Persistence_RdbFrequency_ARM_1H,
	"6h":  Persistence_RdbFrequency_ARM_6H,
}
