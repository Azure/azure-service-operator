// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type MongodbDatabase_STATUS struct {
	// Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of an Azure Cosmos DB MongoDB database
	Properties *MongoDBDatabaseGetProperties_STATUS `json:"properties,omitempty"`
	Tags       map[string]string                    `json:"tags,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

// The properties of an Azure Cosmos DB MongoDB database
type MongoDBDatabaseGetProperties_STATUS struct {
	// Options: Cosmos DB options resource object
	Options  *OptionsResource_STATUS                       `json:"options,omitempty"`
	Resource *MongoDBDatabaseGetProperties_Resource_STATUS `json:"resource,omitempty"`
}

type MongoDBDatabaseGetProperties_Resource_STATUS struct {
	// CreateMode: Enum to indicate the mode of resource creation.
	CreateMode *CreateMode_STATUS `json:"createMode,omitempty"`

	// Etag: A system generated property representing the resource etag required for optimistic concurrency control.
	Etag *string `json:"_etag,omitempty"`

	// Id: Name of the Cosmos DB MongoDB database
	Id *string `json:"id,omitempty"`

	// RestoreParameters: Parameters to indicate the information about the restore
	RestoreParameters *RestoreParametersBase_STATUS `json:"restoreParameters,omitempty"`

	// Rid: A system generated property. A unique identifier.
	Rid *string `json:"_rid,omitempty"`

	// Ts: A system generated property that denotes the last updated timestamp of the resource.
	Ts *float64 `json:"_ts,omitempty"`
}
