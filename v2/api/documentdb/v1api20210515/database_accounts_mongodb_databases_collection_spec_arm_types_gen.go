// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB MongoDB collection.
	Properties *MongoDBCollectionCreateUpdateProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                            `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (collection DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (collection *DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM) GetName() string {
	return collection.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
func (collection *DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
}

// Properties to create and update Azure Cosmos DB MongoDB collection.
type MongoDBCollectionCreateUpdateProperties_ARM struct {
	// Options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the
	// request.
	Options *CreateUpdateOptions_ARM `json:"options,omitempty"`

	// Resource: The standard JSON format of a MongoDB collection
	Resource *MongoDBCollectionResource_ARM `json:"resource,omitempty"`
}

// Cosmos DB MongoDB collection resource object
type MongoDBCollectionResource_ARM struct {
	// AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	// Id: Name of the Cosmos DB MongoDB collection
	Id *string `json:"id,omitempty"`

	// Indexes: List of index keys
	Indexes []MongoIndex_ARM `json:"indexes,omitempty"`

	// ShardKey: A key-value pair of shard keys to be applied for the request.
	ShardKey map[string]string `json:"shardKey,omitempty"`
}

// Cosmos DB MongoDB collection index key
type MongoIndex_ARM struct {
	// Key: Cosmos DB MongoDB collection index keys
	Key *MongoIndexKeys_ARM `json:"key,omitempty"`

	// Options: Cosmos DB MongoDB collection index key options
	Options *MongoIndexOptions_ARM `json:"options,omitempty"`
}

// Cosmos DB MongoDB collection resource object
type MongoIndexKeys_ARM struct {
	// Keys: List of keys for each MongoDB collection in the Azure Cosmos DB service
	Keys []string `json:"keys,omitempty"`
}

// Cosmos DB MongoDB collection index options
type MongoIndexOptions_ARM struct {
	// ExpireAfterSeconds: Expire after seconds
	ExpireAfterSeconds *int `json:"expireAfterSeconds,omitempty"`

	// Unique: Is unique or not
	Unique *bool `json:"unique,omitempty"`
}
