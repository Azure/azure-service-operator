// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type MongodbDatabaseCollection_Spec struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB MongoDB collection.
	Properties *MongoDBCollectionCreateUpdateProperties `json:"properties,omitempty"`
	Tags       map[string]string                        `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &MongodbDatabaseCollection_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (collection MongodbDatabaseCollection_Spec) GetAPIVersion() string {
	return "2021-05-15"
}

// GetName returns the Name of the resource
func (collection *MongodbDatabaseCollection_Spec) GetName() string {
	return collection.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
func (collection *MongodbDatabaseCollection_Spec) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
}

// Properties to create and update Azure Cosmos DB MongoDB collection.
type MongoDBCollectionCreateUpdateProperties struct {
	// Options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the
	// request.
	Options *CreateUpdateOptions `json:"options,omitempty"`

	// Resource: The standard JSON format of a MongoDB collection
	Resource *MongoDBCollectionResource `json:"resource,omitempty"`
}

// CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are "If-Match",
// "If-None-Match", "Session-Token" and "Throughput"
type CreateUpdateOptions struct {
	// AutoscaleSettings: Specifies the Autoscale settings.
	AutoscaleSettings *AutoscaleSettings `json:"autoscaleSettings,omitempty"`

	// Throughput: Request Units per second. For example, "throughput": 10000.
	Throughput *int `json:"throughput,omitempty"`
}

// Cosmos DB MongoDB collection resource object
type MongoDBCollectionResource struct {
	// AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	// Id: Name of the Cosmos DB MongoDB collection
	Id *string `json:"id,omitempty"`

	// Indexes: List of index keys
	Indexes []MongoIndex `json:"indexes,omitempty"`

	// ShardKey: A key-value pair of shard keys to be applied for the request.
	ShardKey map[string]string `json:"shardKey,omitempty"`
}

type AutoscaleSettings struct {
	// MaxThroughput: Represents maximum throughput, the resource can scale up to.
	MaxThroughput *int `json:"maxThroughput,omitempty"`
}

// Cosmos DB MongoDB collection index key
type MongoIndex struct {
	// Key: Cosmos DB MongoDB collection index keys
	Key *MongoIndexKeys `json:"key,omitempty"`

	// Options: Cosmos DB MongoDB collection index key options
	Options *MongoIndexOptions `json:"options,omitempty"`
}

// Cosmos DB MongoDB collection resource object
type MongoIndexKeys struct {
	// Keys: List of keys for each MongoDB collection in the Azure Cosmos DB service
	Keys []string `json:"keys,omitempty"`
}

// Cosmos DB MongoDB collection index options
type MongoIndexOptions struct {
	// ExpireAfterSeconds: Expire after seconds
	ExpireAfterSeconds *int `json:"expireAfterSeconds,omitempty"`

	// Unique: Is unique or not
	Unique *bool `json:"unique,omitempty"`
}
