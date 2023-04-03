// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

// Deprecated version of DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS. Use v1api20210515.DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS instead
type DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS_ARM struct {
	Id         *string                                         `json:"id,omitempty"`
	Location   *string                                         `json:"location,omitempty"`
	Name       *string                                         `json:"name,omitempty"`
	Properties *SqlUserDefinedFunctionGetProperties_STATUS_ARM `json:"properties,omitempty"`
	Tags       map[string]string                               `json:"tags,omitempty"`
	Type       *string                                         `json:"type,omitempty"`
}

// Deprecated version of SqlUserDefinedFunctionGetProperties_STATUS. Use v1api20210515.SqlUserDefinedFunctionGetProperties_STATUS instead
type SqlUserDefinedFunctionGetProperties_STATUS_ARM struct {
	Resource *SqlUserDefinedFunctionGetProperties_Resource_STATUS_ARM `json:"resource,omitempty"`
}

// Deprecated version of SqlUserDefinedFunctionGetProperties_Resource_STATUS. Use v1api20210515.SqlUserDefinedFunctionGetProperties_Resource_STATUS instead
type SqlUserDefinedFunctionGetProperties_Resource_STATUS_ARM struct {
	Body *string  `json:"body,omitempty"`
	Etag *string  `json:"_etag,omitempty"`
	Id   *string  `json:"id,omitempty"`
	Rid  *string  `json:"_rid,omitempty"`
	Ts   *float64 `json:"_ts,omitempty"`
}
