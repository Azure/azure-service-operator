// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231115

type DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS_ARM struct {
	// Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of an Azure Cosmos DB trigger
	Properties *SqlTriggerGetProperties_STATUS_ARM `json:"properties,omitempty"`
	Tags       map[string]string                   `json:"tags,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

// The properties of an Azure Cosmos DB trigger
type SqlTriggerGetProperties_STATUS_ARM struct {
	Resource *SqlTriggerGetProperties_Resource_STATUS_ARM `json:"resource,omitempty"`
}

type SqlTriggerGetProperties_Resource_STATUS_ARM struct {
	// Body: Body of the Trigger
	Body *string `json:"body,omitempty"`

	// Etag: A system generated property representing the resource etag required for optimistic concurrency control.
	Etag *string `json:"_etag,omitempty"`

	// Id: Name of the Cosmos DB SQL trigger
	Id *string `json:"id,omitempty"`

	// Rid: A system generated property. A unique identifier.
	Rid *string `json:"_rid,omitempty"`

	// TriggerOperation: The operation the trigger is associated with
	TriggerOperation *SqlTriggerGetProperties_Resource_TriggerOperation_STATUS `json:"triggerOperation,omitempty"`

	// TriggerType: Type of the Trigger
	TriggerType *SqlTriggerGetProperties_Resource_TriggerType_STATUS `json:"triggerType,omitempty"`

	// Ts: A system generated property that denotes the last updated timestamp of the resource.
	Ts *float64 `json:"_ts,omitempty"`
}

type SqlTriggerGetProperties_Resource_TriggerOperation_STATUS string

const (
	SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_All     = SqlTriggerGetProperties_Resource_TriggerOperation_STATUS("All")
	SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Create  = SqlTriggerGetProperties_Resource_TriggerOperation_STATUS("Create")
	SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Delete  = SqlTriggerGetProperties_Resource_TriggerOperation_STATUS("Delete")
	SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Replace = SqlTriggerGetProperties_Resource_TriggerOperation_STATUS("Replace")
	SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Update  = SqlTriggerGetProperties_Resource_TriggerOperation_STATUS("Update")
)

// Mapping from string to SqlTriggerGetProperties_Resource_TriggerOperation_STATUS
var sqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Values = map[string]SqlTriggerGetProperties_Resource_TriggerOperation_STATUS{
	"all":     SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_All,
	"create":  SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Create,
	"delete":  SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Delete,
	"replace": SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Replace,
	"update":  SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Update,
}

type SqlTriggerGetProperties_Resource_TriggerType_STATUS string

const (
	SqlTriggerGetProperties_Resource_TriggerType_STATUS_Post = SqlTriggerGetProperties_Resource_TriggerType_STATUS("Post")
	SqlTriggerGetProperties_Resource_TriggerType_STATUS_Pre  = SqlTriggerGetProperties_Resource_TriggerType_STATUS("Pre")
)

// Mapping from string to SqlTriggerGetProperties_Resource_TriggerType_STATUS
var sqlTriggerGetProperties_Resource_TriggerType_STATUS_Values = map[string]SqlTriggerGetProperties_Resource_TriggerType_STATUS{
	"post": SqlTriggerGetProperties_Resource_TriggerType_STATUS_Post,
	"pre":  SqlTriggerGetProperties_Resource_TriggerType_STATUS_Pre,
}
