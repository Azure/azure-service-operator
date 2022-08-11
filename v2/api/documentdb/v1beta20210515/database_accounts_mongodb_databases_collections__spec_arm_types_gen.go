// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccountsMongodbDatabasesCollections_SpecARM struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: Cosmos DB collection name.
	Name string `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB MongoDB collection.
	Properties *MongoDBCollectionCreateUpdatePropertiesARM `json:"properties,omitempty"`

	// Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
	// resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
	// greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
	// type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph",
	// "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsMongodbDatabasesCollections_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (collections DatabaseAccountsMongodbDatabasesCollections_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (collections *DatabaseAccountsMongodbDatabasesCollections_SpecARM) GetName() string {
	return collections.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
func (collections *DatabaseAccountsMongodbDatabasesCollections_SpecARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionCreateUpdateProperties
type MongoDBCollectionCreateUpdatePropertiesARM struct {
	// Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are "If-Match",
	// "If-None-Match", "Session-Token" and "Throughput"
	Options *CreateUpdateOptionsARM `json:"options,omitempty"`

	// Resource: Cosmos DB MongoDB collection resource object
	Resource *MongoDBCollectionResourceARM `json:"resource,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CreateUpdateOptions
type CreateUpdateOptionsARM struct {
	AutoscaleSettings *AutoscaleSettingsARM `json:"autoscaleSettings,omitempty"`

	// Throughput: Request Units per second. For example, "throughput": 10000.
	Throughput *int `json:"throughput,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionResource
type MongoDBCollectionResourceARM struct {
	// AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	// Id: Name of the Cosmos DB MongoDB collection
	Id *string `json:"id,omitempty"`

	// Indexes: List of index keys
	Indexes []MongoIndexARM `json:"indexes,omitempty"`

	// ShardKey: The shard key and partition kind pair, only support "Hash" partition kind
	ShardKey map[string]string `json:"shardKey,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettings
type AutoscaleSettingsARM struct {
	// MaxThroughput: Represents maximum throughput, the resource can scale up to.
	MaxThroughput *int `json:"maxThroughput,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndex
type MongoIndexARM struct {
	// Key: Cosmos DB MongoDB collection resource object
	Key *MongoIndexKeysARM `json:"key,omitempty"`

	// Options: Cosmos DB MongoDB collection index options
	Options *MongoIndexOptionsARM `json:"options,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexKeys
type MongoIndexKeysARM struct {
	// Keys: List of keys for each MongoDB collection in the Azure Cosmos DB service
	Keys []string `json:"keys,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexOptions
type MongoIndexOptionsARM struct {
	// ExpireAfterSeconds: Expire after seconds
	ExpireAfterSeconds *int `json:"expireAfterSeconds,omitempty"`

	// Unique: Is unique or not
	Unique *bool `json:"unique,omitempty"`
}
