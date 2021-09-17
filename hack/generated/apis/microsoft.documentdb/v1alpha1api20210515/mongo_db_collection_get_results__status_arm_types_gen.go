// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

//Generated from:
type MongoDBCollectionGetResults_StatusARM struct {
	//Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	//Properties: The properties of an Azure Cosmos DB MongoDB collection
	Properties *MongoDBCollectionGetProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                         `json:"tags,omitempty"`

	//Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

//Generated from:
type MongoDBCollectionGetProperties_StatusARM struct {
	Options  *OptionsResource_StatusARM                         `json:"options,omitempty"`
	Resource *MongoDBCollectionGetProperties_Status_ResourceARM `json:"resource,omitempty"`
}

type MongoDBCollectionGetProperties_Status_ResourceARM struct {
	//AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	//Etag: A system generated property representing the resource etag required for
	//optimistic concurrency control.
	Etag *string `json:"_etag,omitempty"`

	//Id: Name of the Cosmos DB MongoDB collection
	Id string `json:"id"`

	//Indexes: List of index keys
	Indexes []MongoIndex_StatusARM `json:"indexes,omitempty"`

	//Rid: A system generated property. A unique identifier.
	Rid *string `json:"_rid,omitempty"`

	//ShardKey: A key-value pair of shard keys to be applied for the request.
	ShardKey map[string]string `json:"shardKey,omitempty"`

	//Ts: A system generated property that denotes the last updated timestamp of the
	//resource.
	Ts *float64 `json:"_ts,omitempty"`
}

//Generated from:
type OptionsResource_StatusARM struct {
	//AutoscaleSettings: Specifies the Autoscale settings.
	AutoscaleSettings *AutoscaleSettings_StatusARM `json:"autoscaleSettings,omitempty"`

	//Throughput: Value of the Cosmos DB resource throughput or autoscaleSettings. Use
	//the ThroughputSetting resource when retrieving offer details.
	Throughput *int `json:"throughput,omitempty"`
}

//Generated from:
type AutoscaleSettings_StatusARM struct {
	//MaxThroughput: Represents maximum throughput, the resource can scale up to.
	MaxThroughput *int `json:"maxThroughput,omitempty"`
}

//Generated from:
type MongoIndex_StatusARM struct {
	//Key: Cosmos DB MongoDB collection index keys
	Key *MongoIndexKeys_StatusARM `json:"key,omitempty"`

	//Options: Cosmos DB MongoDB collection index key options
	Options *MongoIndexOptions_StatusARM `json:"options,omitempty"`
}

//Generated from:
type MongoIndexKeys_StatusARM struct {
	//Keys: List of keys for each MongoDB collection in the Azure Cosmos DB service
	Keys []string `json:"keys,omitempty"`
}

//Generated from:
type MongoIndexOptions_StatusARM struct {
	//ExpireAfterSeconds: Expire after seconds
	ExpireAfterSeconds *int `json:"expireAfterSeconds,omitempty"`

	//Unique: Is unique or not
	Unique *bool `json:"unique,omitempty"`
}
