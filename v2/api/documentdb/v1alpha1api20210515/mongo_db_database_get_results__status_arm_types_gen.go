// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

// Deprecated version of MongoDBDatabaseGetResults_Status. Use v1beta20210515.MongoDBDatabaseGetResults_Status instead
type MongoDBDatabaseGetResults_StatusARM struct {
	Id         *string                                 `json:"id,omitempty"`
	Location   *string                                 `json:"location,omitempty"`
	Name       *string                                 `json:"name,omitempty"`
	Properties *MongoDBDatabaseGetProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                       `json:"tags,omitempty"`
	Type       *string                                 `json:"type,omitempty"`
}

// Deprecated version of MongoDBDatabaseGetProperties_Status. Use v1beta20210515.MongoDBDatabaseGetProperties_Status instead
type MongoDBDatabaseGetProperties_StatusARM struct {
	Options  *OptionsResource_StatusARM                       `json:"options,omitempty"`
	Resource *MongoDBDatabaseGetProperties_Status_ResourceARM `json:"resource,omitempty"`
}

// Deprecated version of MongoDBDatabaseGetProperties_Status_Resource. Use v1beta20210515.MongoDBDatabaseGetProperties_Status_Resource instead
type MongoDBDatabaseGetProperties_Status_ResourceARM struct {
	Etag *string  `json:"_etag,omitempty"`
	Id   *string  `json:"id,omitempty"`
	Rid  *string  `json:"_rid,omitempty"`
	Ts   *float64 `json:"_ts,omitempty"`
}
