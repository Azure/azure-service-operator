// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Servers_ElasticPool_Spec_ARM struct {
	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *ElasticPoolProperties_ARM `json:"properties,omitempty"`

	// Sku: The elastic pool SKU.
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition,
	// family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation`
	// REST API or the following command:
	// ```azurecli
	// az sql elastic-pool list-editions -l <location> -o table
	// ````
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_ElasticPool_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (pool Servers_ElasticPool_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (pool *Servers_ElasticPool_Spec_ARM) GetName() string {
	return pool.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/elasticPools"
func (pool *Servers_ElasticPool_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers/elasticPools"
}

// Properties of an elastic pool
type ElasticPoolProperties_ARM struct {
	// HighAvailabilityReplicaCount: The number of secondary replicas associated with the elastic pool that are used to provide
	// high availability. Applicable only to Hyperscale elastic pools.
	HighAvailabilityReplicaCount *int `json:"highAvailabilityReplicaCount,omitempty"`

	// LicenseType: The license type to apply for this elastic pool.
	LicenseType *ElasticPoolProperties_LicenseType `json:"licenseType,omitempty"`

	// MaintenanceConfigurationId: Maintenance configuration id assigned to the elastic pool. This configuration defines the
	// period when the maintenance updates will will occur.
	MaintenanceConfigurationId *string `json:"maintenanceConfigurationId,omitempty"`

	// MaxSizeBytes: The storage limit for the database elastic pool in bytes.
	MaxSizeBytes *int `json:"maxSizeBytes,omitempty"`

	// MinCapacity: Minimal capacity that serverless pool will not shrink below, if not paused
	MinCapacity *float64 `json:"minCapacity,omitempty"`

	// PerDatabaseSettings: The per database settings for the elastic pool.
	PerDatabaseSettings *ElasticPoolPerDatabaseSettings_ARM `json:"perDatabaseSettings,omitempty"`

	// ZoneRedundant: Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be
	// spread across multiple availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

// Per database settings of an elastic pool.
type ElasticPoolPerDatabaseSettings_ARM struct {
	// MaxCapacity: The maximum capacity any one database can consume.
	MaxCapacity *float64 `json:"maxCapacity,omitempty"`

	// MinCapacity: The minimum capacity all databases are guaranteed.
	MinCapacity *float64 `json:"minCapacity,omitempty"`
}
