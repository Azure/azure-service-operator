// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301

//Deprecated version of Database_Status. Use v1beta20210301.Database_Status instead
type Database_StatusARM struct {
	Id         *string                       `json:"id,omitempty"`
	Name       *string                       `json:"name,omitempty"`
	Properties *DatabaseProperties_StatusARM `json:"properties,omitempty"`
	Type       *string                       `json:"type,omitempty"`
}

//Deprecated version of DatabaseProperties_Status. Use v1beta20210301.DatabaseProperties_Status instead
type DatabaseProperties_StatusARM struct {
	ClientProtocol    *DatabasePropertiesStatusClientProtocol   `json:"clientProtocol,omitempty"`
	ClusteringPolicy  *DatabasePropertiesStatusClusteringPolicy `json:"clusteringPolicy,omitempty"`
	EvictionPolicy    *DatabasePropertiesStatusEvictionPolicy   `json:"evictionPolicy,omitempty"`
	Modules           []Module_StatusARM                        `json:"modules,omitempty"`
	Persistence       *Persistence_StatusARM                    `json:"persistence,omitempty"`
	Port              *int                                      `json:"port,omitempty"`
	ProvisioningState *ProvisioningState_Status                 `json:"provisioningState,omitempty"`
	ResourceState     *ResourceState_Status                     `json:"resourceState,omitempty"`
}

//Deprecated version of DatabasePropertiesStatusClientProtocol. Use v1beta20210301.DatabasePropertiesStatusClientProtocol
//instead
type DatabasePropertiesStatusClientProtocol string

const (
	DatabasePropertiesStatusClientProtocolEncrypted = DatabasePropertiesStatusClientProtocol("Encrypted")
	DatabasePropertiesStatusClientProtocolPlaintext = DatabasePropertiesStatusClientProtocol("Plaintext")
)

//Deprecated version of DatabasePropertiesStatusClusteringPolicy. Use
//v1beta20210301.DatabasePropertiesStatusClusteringPolicy instead
type DatabasePropertiesStatusClusteringPolicy string

const (
	DatabasePropertiesStatusClusteringPolicyEnterpriseCluster = DatabasePropertiesStatusClusteringPolicy("EnterpriseCluster")
	DatabasePropertiesStatusClusteringPolicyOSSCluster        = DatabasePropertiesStatusClusteringPolicy("OSSCluster")
)

//Deprecated version of DatabasePropertiesStatusEvictionPolicy. Use v1beta20210301.DatabasePropertiesStatusEvictionPolicy
//instead
type DatabasePropertiesStatusEvictionPolicy string

const (
	DatabasePropertiesStatusEvictionPolicyAllKeysLFU     = DatabasePropertiesStatusEvictionPolicy("AllKeysLFU")
	DatabasePropertiesStatusEvictionPolicyAllKeysLRU     = DatabasePropertiesStatusEvictionPolicy("AllKeysLRU")
	DatabasePropertiesStatusEvictionPolicyAllKeysRandom  = DatabasePropertiesStatusEvictionPolicy("AllKeysRandom")
	DatabasePropertiesStatusEvictionPolicyNoEviction     = DatabasePropertiesStatusEvictionPolicy("NoEviction")
	DatabasePropertiesStatusEvictionPolicyVolatileLFU    = DatabasePropertiesStatusEvictionPolicy("VolatileLFU")
	DatabasePropertiesStatusEvictionPolicyVolatileLRU    = DatabasePropertiesStatusEvictionPolicy("VolatileLRU")
	DatabasePropertiesStatusEvictionPolicyVolatileRandom = DatabasePropertiesStatusEvictionPolicy("VolatileRandom")
	DatabasePropertiesStatusEvictionPolicyVolatileTTL    = DatabasePropertiesStatusEvictionPolicy("VolatileTTL")
)

//Deprecated version of Module_Status. Use v1beta20210301.Module_Status instead
type Module_StatusARM struct {
	Args    *string `json:"args,omitempty"`
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

//Deprecated version of Persistence_Status. Use v1beta20210301.Persistence_Status instead
type Persistence_StatusARM struct {
	AofEnabled   *bool                          `json:"aofEnabled,omitempty"`
	AofFrequency *PersistenceStatusAofFrequency `json:"aofFrequency,omitempty"`
	RdbEnabled   *bool                          `json:"rdbEnabled,omitempty"`
	RdbFrequency *PersistenceStatusRdbFrequency `json:"rdbFrequency,omitempty"`
}

//Deprecated version of PersistenceStatusAofFrequency. Use v1beta20210301.PersistenceStatusAofFrequency instead
type PersistenceStatusAofFrequency string

const (
	PersistenceStatusAofFrequency1S     = PersistenceStatusAofFrequency("1s")
	PersistenceStatusAofFrequencyAlways = PersistenceStatusAofFrequency("always")
)

//Deprecated version of PersistenceStatusRdbFrequency. Use v1beta20210301.PersistenceStatusRdbFrequency instead
type PersistenceStatusRdbFrequency string

const (
	PersistenceStatusRdbFrequency12H = PersistenceStatusRdbFrequency("12h")
	PersistenceStatusRdbFrequency1H  = PersistenceStatusRdbFrequency("1h")
	PersistenceStatusRdbFrequency6H  = PersistenceStatusRdbFrequency("6h")
)
