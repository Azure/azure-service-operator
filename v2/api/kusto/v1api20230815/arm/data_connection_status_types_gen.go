// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "encoding/json"

type DataConnection_STATUS struct {
	// CosmosDb: Mutually exclusive with all other properties
	CosmosDb *CosmosDbDataConnection_STATUS `json:"cosmosDb,omitempty"`

	// EventGrid: Mutually exclusive with all other properties
	EventGrid *EventGridDataConnection_STATUS `json:"eventGrid,omitempty"`

	// EventHub: Mutually exclusive with all other properties
	EventHub *EventHubDataConnection_STATUS `json:"eventHub,omitempty"`

	// IotHub: Mutually exclusive with all other properties
	IotHub *IotHubDataConnection_STATUS `json:"iotHub,omitempty"`

	// Name: The name of the resource
	Name *string `conversion:"pushtoleaf" json:"name,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because DataConnection_STATUS represents a discriminated union (JSON OneOf)
func (connection DataConnection_STATUS) MarshalJSON() ([]byte, error) {
	if connection.CosmosDb != nil {
		return json.Marshal(connection.CosmosDb)
	}

	if connection.EventGrid != nil {
		return json.Marshal(connection.EventGrid)
	}

	if connection.EventHub != nil {
		return json.Marshal(connection.EventHub)
	}

	if connection.IotHub != nil {
		return json.Marshal(connection.IotHub)
	}

	return nil, nil
}

// UnmarshalJSON unmarshals the DataConnection_STATUS
func (connection *DataConnection_STATUS) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["kind"]
	if discriminator == "CosmosDb" {
		connection.CosmosDb = &CosmosDbDataConnection_STATUS{}
		return json.Unmarshal(data, connection.CosmosDb)
	}
	if discriminator == "EventGrid" {
		connection.EventGrid = &EventGridDataConnection_STATUS{}
		return json.Unmarshal(data, connection.EventGrid)
	}
	if discriminator == "EventHub" {
		connection.EventHub = &EventHubDataConnection_STATUS{}
		return json.Unmarshal(data, connection.EventHub)
	}
	if discriminator == "IotHub" {
		connection.IotHub = &IotHubDataConnection_STATUS{}
		return json.Unmarshal(data, connection.IotHub)
	}

	// No error
	return nil
}

type CosmosDbDataConnection_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Kind: Kind of the endpoint for the data connection
	Kind CosmosDbDataConnection_Kind_STATUS `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `conversion:"noarmconversion" json:"name,omitempty"`

	// Properties: The properties of the CosmosDb data connection.
	Properties *CosmosDbDataConnectionProperties_STATUS `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type EventGridDataConnection_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Kind: Kind of the endpoint for the data connection
	Kind EventGridDataConnection_Kind_STATUS `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `conversion:"noarmconversion" json:"name,omitempty"`

	// Properties: The properties of the Event Grid data connection.
	Properties *EventGridConnectionProperties_STATUS `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type EventHubDataConnection_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Kind: Kind of the endpoint for the data connection
	Kind EventHubDataConnection_Kind_STATUS `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `conversion:"noarmconversion" json:"name,omitempty"`

	// Properties: The Event Hub data connection properties to validate.
	Properties *EventHubConnectionProperties_STATUS `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type IotHubDataConnection_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Kind: Kind of the endpoint for the data connection
	Kind IotHubDataConnection_Kind_STATUS `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `conversion:"noarmconversion" json:"name,omitempty"`

	// Properties: The Iot Hub data connection properties.
	Properties *IotHubConnectionProperties_STATUS `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type CosmosDbDataConnection_Kind_STATUS string

const CosmosDbDataConnection_Kind_STATUS_CosmosDb = CosmosDbDataConnection_Kind_STATUS("CosmosDb")

// Mapping from string to CosmosDbDataConnection_Kind_STATUS
var cosmosDbDataConnection_Kind_STATUS_Values = map[string]CosmosDbDataConnection_Kind_STATUS{
	"cosmosdb": CosmosDbDataConnection_Kind_STATUS_CosmosDb,
}

// Class representing the Kusto CosmosDb data connection properties.
type CosmosDbDataConnectionProperties_STATUS struct {
	// CosmosDbAccountResourceId: The resource ID of the Cosmos DB account used to create the data connection.
	CosmosDbAccountResourceId *string `json:"cosmosDbAccountResourceId,omitempty"`

	// CosmosDbContainer: The name of an existing container in the Cosmos DB database.
	CosmosDbContainer *string `json:"cosmosDbContainer,omitempty"`

	// CosmosDbDatabase: The name of an existing database in the Cosmos DB account.
	CosmosDbDatabase *string `json:"cosmosDbDatabase,omitempty"`

	// ManagedIdentityObjectId: The object ID of the managed identity resource.
	ManagedIdentityObjectId *string `json:"managedIdentityObjectId,omitempty"`

	// ManagedIdentityResourceId: The resource ID of a managed system or user-assigned identity. The identity is used to
	// authenticate with Cosmos DB.
	ManagedIdentityResourceId *string `json:"managedIdentityResourceId,omitempty"`

	// MappingRuleName: The name of an existing mapping rule to use when ingesting the retrieved data.
	MappingRuleName *string `json:"mappingRuleName,omitempty"`

	// ProvisioningState: The provisioned state of the resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// RetrievalStartDate: Optional. If defined, the data connection retrieves Cosmos DB documents created or updated after the
	// specified retrieval start date.
	RetrievalStartDate *string `json:"retrievalStartDate,omitempty"`

	// TableName: The case-sensitive name of the existing target table in your cluster. Retrieved data is ingested into this
	// table.
	TableName *string `json:"tableName,omitempty"`
}

// Class representing the Kusto event grid connection properties.
type EventGridConnectionProperties_STATUS struct {
	// BlobStorageEventType: The name of blob storage event type to process.
	BlobStorageEventType *BlobStorageEventType_STATUS `json:"blobStorageEventType,omitempty"`

	// ConsumerGroup: The event hub consumer group.
	ConsumerGroup *string `json:"consumerGroup,omitempty"`

	// DataFormat: The data format of the message. Optionally the data format can be added to each message.
	DataFormat *EventGridDataFormat_STATUS `json:"dataFormat,omitempty"`

	// DatabaseRouting: Indication for database routing information from the data connection, by default only database routing
	// information is allowed
	DatabaseRouting *EventGridConnectionProperties_DatabaseRouting_STATUS `json:"databaseRouting,omitempty"`

	// EventGridResourceId: The resource ID of the event grid that is subscribed to the storage account events.
	EventGridResourceId *string `json:"eventGridResourceId,omitempty"`

	// EventHubResourceId: The resource ID where the event grid is configured to send events.
	EventHubResourceId *string `json:"eventHubResourceId,omitempty"`

	// IgnoreFirstRecord: A Boolean value that, if set to true, indicates that ingestion should ignore the first record of
	// every file
	IgnoreFirstRecord *bool `json:"ignoreFirstRecord,omitempty"`

	// ManagedIdentityObjectId: The object ID of managedIdentityResourceId
	ManagedIdentityObjectId *string `json:"managedIdentityObjectId,omitempty"`

	// ManagedIdentityResourceId: The resource ID of a managed identity (system or user assigned) to be used to authenticate
	// with event hub and storage account.
	ManagedIdentityResourceId *string `json:"managedIdentityResourceId,omitempty"`

	// MappingRuleName: The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each
	// message.
	MappingRuleName *string `json:"mappingRuleName,omitempty"`

	// ProvisioningState: The provisioned state of the resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// StorageAccountResourceId: The resource ID of the storage account where the data resides.
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`

	// TableName: The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName *string `json:"tableName,omitempty"`
}

type EventGridDataConnection_Kind_STATUS string

const EventGridDataConnection_Kind_STATUS_EventGrid = EventGridDataConnection_Kind_STATUS("EventGrid")

// Mapping from string to EventGridDataConnection_Kind_STATUS
var eventGridDataConnection_Kind_STATUS_Values = map[string]EventGridDataConnection_Kind_STATUS{
	"eventgrid": EventGridDataConnection_Kind_STATUS_EventGrid,
}

// Class representing the Kusto event hub connection properties.
type EventHubConnectionProperties_STATUS struct {
	// Compression: The event hub messages compression type
	Compression *Compression_STATUS `json:"compression,omitempty"`

	// ConsumerGroup: The event hub consumer group.
	ConsumerGroup *string `json:"consumerGroup,omitempty"`

	// DataFormat: The data format of the message. Optionally the data format can be added to each message.
	DataFormat *EventHubDataFormat_STATUS `json:"dataFormat,omitempty"`

	// DatabaseRouting: Indication for database routing information from the data connection, by default only database routing
	// information is allowed
	DatabaseRouting *EventHubConnectionProperties_DatabaseRouting_STATUS `json:"databaseRouting,omitempty"`

	// EventHubResourceId: The resource ID of the event hub to be used to create a data connection.
	EventHubResourceId *string `json:"eventHubResourceId,omitempty"`

	// EventSystemProperties: System properties of the event hub
	EventSystemProperties []string `json:"eventSystemProperties,omitempty"`

	// ManagedIdentityObjectId: The object ID of the managedIdentityResourceId
	ManagedIdentityObjectId *string `json:"managedIdentityObjectId,omitempty"`

	// ManagedIdentityResourceId: The resource ID of a managed identity (system or user assigned) to be used to authenticate
	// with event hub.
	ManagedIdentityResourceId *string `json:"managedIdentityResourceId,omitempty"`

	// MappingRuleName: The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each
	// message.
	MappingRuleName *string `json:"mappingRuleName,omitempty"`

	// ProvisioningState: The provisioned state of the resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// RetrievalStartDate: When defined, the data connection retrieves existing Event hub events created since the Retrieval
	// start date. It can only retrieve events retained by the Event hub, based on its retention period.
	RetrievalStartDate *string `json:"retrievalStartDate,omitempty"`

	// TableName: The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName *string `json:"tableName,omitempty"`
}

type EventHubDataConnection_Kind_STATUS string

const EventHubDataConnection_Kind_STATUS_EventHub = EventHubDataConnection_Kind_STATUS("EventHub")

// Mapping from string to EventHubDataConnection_Kind_STATUS
var eventHubDataConnection_Kind_STATUS_Values = map[string]EventHubDataConnection_Kind_STATUS{
	"eventhub": EventHubDataConnection_Kind_STATUS_EventHub,
}

// Class representing the Kusto Iot hub connection properties.
type IotHubConnectionProperties_STATUS struct {
	// ConsumerGroup: The iot hub consumer group.
	ConsumerGroup *string `json:"consumerGroup,omitempty"`

	// DataFormat: The data format of the message. Optionally the data format can be added to each message.
	DataFormat *IotHubDataFormat_STATUS `json:"dataFormat,omitempty"`

	// DatabaseRouting: Indication for database routing information from the data connection, by default only database routing
	// information is allowed
	DatabaseRouting *IotHubConnectionProperties_DatabaseRouting_STATUS `json:"databaseRouting,omitempty"`

	// EventSystemProperties: System properties of the iot hub
	EventSystemProperties []string `json:"eventSystemProperties,omitempty"`

	// IotHubResourceId: The resource ID of the Iot hub to be used to create a data connection.
	IotHubResourceId *string `json:"iotHubResourceId,omitempty"`

	// MappingRuleName: The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each
	// message.
	MappingRuleName *string `json:"mappingRuleName,omitempty"`

	// ProvisioningState: The provisioned state of the resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// RetrievalStartDate: When defined, the data connection retrieves existing Event hub events created since the Retrieval
	// start date. It can only retrieve events retained by the Event hub, based on its retention period.
	RetrievalStartDate *string `json:"retrievalStartDate,omitempty"`

	// SharedAccessPolicyName: The name of the share access policy
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty"`

	// TableName: The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName *string `json:"tableName,omitempty"`
}

type IotHubDataConnection_Kind_STATUS string

const IotHubDataConnection_Kind_STATUS_IotHub = IotHubDataConnection_Kind_STATUS("IotHub")

// Mapping from string to IotHubDataConnection_Kind_STATUS
var iotHubDataConnection_Kind_STATUS_Values = map[string]IotHubDataConnection_Kind_STATUS{
	"iothub": IotHubDataConnection_Kind_STATUS_IotHub,
}

// The name of blob storage event type to process.
type BlobStorageEventType_STATUS string

const (
	BlobStorageEventType_STATUS_MicrosoftStorageBlobCreated = BlobStorageEventType_STATUS("Microsoft.Storage.BlobCreated")
	BlobStorageEventType_STATUS_MicrosoftStorageBlobRenamed = BlobStorageEventType_STATUS("Microsoft.Storage.BlobRenamed")
)

// Mapping from string to BlobStorageEventType_STATUS
var blobStorageEventType_STATUS_Values = map[string]BlobStorageEventType_STATUS{
	"microsoft.storage.blobcreated": BlobStorageEventType_STATUS_MicrosoftStorageBlobCreated,
	"microsoft.storage.blobrenamed": BlobStorageEventType_STATUS_MicrosoftStorageBlobRenamed,
}

// The compression type
type Compression_STATUS string

const (
	Compression_STATUS_GZip = Compression_STATUS("GZip")
	Compression_STATUS_None = Compression_STATUS("None")
)

// Mapping from string to Compression_STATUS
var compression_STATUS_Values = map[string]Compression_STATUS{
	"gzip": Compression_STATUS_GZip,
	"none": Compression_STATUS_None,
}

type EventGridConnectionProperties_DatabaseRouting_STATUS string

const (
	EventGridConnectionProperties_DatabaseRouting_STATUS_Multi  = EventGridConnectionProperties_DatabaseRouting_STATUS("Multi")
	EventGridConnectionProperties_DatabaseRouting_STATUS_Single = EventGridConnectionProperties_DatabaseRouting_STATUS("Single")
)

// Mapping from string to EventGridConnectionProperties_DatabaseRouting_STATUS
var eventGridConnectionProperties_DatabaseRouting_STATUS_Values = map[string]EventGridConnectionProperties_DatabaseRouting_STATUS{
	"multi":  EventGridConnectionProperties_DatabaseRouting_STATUS_Multi,
	"single": EventGridConnectionProperties_DatabaseRouting_STATUS_Single,
}

// The data format of the message. Optionally the data format can be added to each message.
type EventGridDataFormat_STATUS string

const (
	EventGridDataFormat_STATUS_APACHEAVRO = EventGridDataFormat_STATUS("APACHEAVRO")
	EventGridDataFormat_STATUS_AVRO       = EventGridDataFormat_STATUS("AVRO")
	EventGridDataFormat_STATUS_CSV        = EventGridDataFormat_STATUS("CSV")
	EventGridDataFormat_STATUS_JSON       = EventGridDataFormat_STATUS("JSON")
	EventGridDataFormat_STATUS_MULTIJSON  = EventGridDataFormat_STATUS("MULTIJSON")
	EventGridDataFormat_STATUS_ORC        = EventGridDataFormat_STATUS("ORC")
	EventGridDataFormat_STATUS_PARQUET    = EventGridDataFormat_STATUS("PARQUET")
	EventGridDataFormat_STATUS_PSV        = EventGridDataFormat_STATUS("PSV")
	EventGridDataFormat_STATUS_RAW        = EventGridDataFormat_STATUS("RAW")
	EventGridDataFormat_STATUS_SCSV       = EventGridDataFormat_STATUS("SCSV")
	EventGridDataFormat_STATUS_SINGLEJSON = EventGridDataFormat_STATUS("SINGLEJSON")
	EventGridDataFormat_STATUS_SOHSV      = EventGridDataFormat_STATUS("SOHSV")
	EventGridDataFormat_STATUS_TSV        = EventGridDataFormat_STATUS("TSV")
	EventGridDataFormat_STATUS_TSVE       = EventGridDataFormat_STATUS("TSVE")
	EventGridDataFormat_STATUS_TXT        = EventGridDataFormat_STATUS("TXT")
	EventGridDataFormat_STATUS_W3CLOGFILE = EventGridDataFormat_STATUS("W3CLOGFILE")
)

// Mapping from string to EventGridDataFormat_STATUS
var eventGridDataFormat_STATUS_Values = map[string]EventGridDataFormat_STATUS{
	"apacheavro": EventGridDataFormat_STATUS_APACHEAVRO,
	"avro":       EventGridDataFormat_STATUS_AVRO,
	"csv":        EventGridDataFormat_STATUS_CSV,
	"json":       EventGridDataFormat_STATUS_JSON,
	"multijson":  EventGridDataFormat_STATUS_MULTIJSON,
	"orc":        EventGridDataFormat_STATUS_ORC,
	"parquet":    EventGridDataFormat_STATUS_PARQUET,
	"psv":        EventGridDataFormat_STATUS_PSV,
	"raw":        EventGridDataFormat_STATUS_RAW,
	"scsv":       EventGridDataFormat_STATUS_SCSV,
	"singlejson": EventGridDataFormat_STATUS_SINGLEJSON,
	"sohsv":      EventGridDataFormat_STATUS_SOHSV,
	"tsv":        EventGridDataFormat_STATUS_TSV,
	"tsve":       EventGridDataFormat_STATUS_TSVE,
	"txt":        EventGridDataFormat_STATUS_TXT,
	"w3clogfile": EventGridDataFormat_STATUS_W3CLOGFILE,
}

type EventHubConnectionProperties_DatabaseRouting_STATUS string

const (
	EventHubConnectionProperties_DatabaseRouting_STATUS_Multi  = EventHubConnectionProperties_DatabaseRouting_STATUS("Multi")
	EventHubConnectionProperties_DatabaseRouting_STATUS_Single = EventHubConnectionProperties_DatabaseRouting_STATUS("Single")
)

// Mapping from string to EventHubConnectionProperties_DatabaseRouting_STATUS
var eventHubConnectionProperties_DatabaseRouting_STATUS_Values = map[string]EventHubConnectionProperties_DatabaseRouting_STATUS{
	"multi":  EventHubConnectionProperties_DatabaseRouting_STATUS_Multi,
	"single": EventHubConnectionProperties_DatabaseRouting_STATUS_Single,
}

// The data format of the message. Optionally the data format can be added to each message.
type EventHubDataFormat_STATUS string

const (
	EventHubDataFormat_STATUS_APACHEAVRO = EventHubDataFormat_STATUS("APACHEAVRO")
	EventHubDataFormat_STATUS_AVRO       = EventHubDataFormat_STATUS("AVRO")
	EventHubDataFormat_STATUS_CSV        = EventHubDataFormat_STATUS("CSV")
	EventHubDataFormat_STATUS_JSON       = EventHubDataFormat_STATUS("JSON")
	EventHubDataFormat_STATUS_MULTIJSON  = EventHubDataFormat_STATUS("MULTIJSON")
	EventHubDataFormat_STATUS_ORC        = EventHubDataFormat_STATUS("ORC")
	EventHubDataFormat_STATUS_PARQUET    = EventHubDataFormat_STATUS("PARQUET")
	EventHubDataFormat_STATUS_PSV        = EventHubDataFormat_STATUS("PSV")
	EventHubDataFormat_STATUS_RAW        = EventHubDataFormat_STATUS("RAW")
	EventHubDataFormat_STATUS_SCSV       = EventHubDataFormat_STATUS("SCSV")
	EventHubDataFormat_STATUS_SINGLEJSON = EventHubDataFormat_STATUS("SINGLEJSON")
	EventHubDataFormat_STATUS_SOHSV      = EventHubDataFormat_STATUS("SOHSV")
	EventHubDataFormat_STATUS_TSV        = EventHubDataFormat_STATUS("TSV")
	EventHubDataFormat_STATUS_TSVE       = EventHubDataFormat_STATUS("TSVE")
	EventHubDataFormat_STATUS_TXT        = EventHubDataFormat_STATUS("TXT")
	EventHubDataFormat_STATUS_W3CLOGFILE = EventHubDataFormat_STATUS("W3CLOGFILE")
)

// Mapping from string to EventHubDataFormat_STATUS
var eventHubDataFormat_STATUS_Values = map[string]EventHubDataFormat_STATUS{
	"apacheavro": EventHubDataFormat_STATUS_APACHEAVRO,
	"avro":       EventHubDataFormat_STATUS_AVRO,
	"csv":        EventHubDataFormat_STATUS_CSV,
	"json":       EventHubDataFormat_STATUS_JSON,
	"multijson":  EventHubDataFormat_STATUS_MULTIJSON,
	"orc":        EventHubDataFormat_STATUS_ORC,
	"parquet":    EventHubDataFormat_STATUS_PARQUET,
	"psv":        EventHubDataFormat_STATUS_PSV,
	"raw":        EventHubDataFormat_STATUS_RAW,
	"scsv":       EventHubDataFormat_STATUS_SCSV,
	"singlejson": EventHubDataFormat_STATUS_SINGLEJSON,
	"sohsv":      EventHubDataFormat_STATUS_SOHSV,
	"tsv":        EventHubDataFormat_STATUS_TSV,
	"tsve":       EventHubDataFormat_STATUS_TSVE,
	"txt":        EventHubDataFormat_STATUS_TXT,
	"w3clogfile": EventHubDataFormat_STATUS_W3CLOGFILE,
}

type IotHubConnectionProperties_DatabaseRouting_STATUS string

const (
	IotHubConnectionProperties_DatabaseRouting_STATUS_Multi  = IotHubConnectionProperties_DatabaseRouting_STATUS("Multi")
	IotHubConnectionProperties_DatabaseRouting_STATUS_Single = IotHubConnectionProperties_DatabaseRouting_STATUS("Single")
)

// Mapping from string to IotHubConnectionProperties_DatabaseRouting_STATUS
var iotHubConnectionProperties_DatabaseRouting_STATUS_Values = map[string]IotHubConnectionProperties_DatabaseRouting_STATUS{
	"multi":  IotHubConnectionProperties_DatabaseRouting_STATUS_Multi,
	"single": IotHubConnectionProperties_DatabaseRouting_STATUS_Single,
}

// The data format of the message. Optionally the data format can be added to each message.
type IotHubDataFormat_STATUS string

const (
	IotHubDataFormat_STATUS_APACHEAVRO = IotHubDataFormat_STATUS("APACHEAVRO")
	IotHubDataFormat_STATUS_AVRO       = IotHubDataFormat_STATUS("AVRO")
	IotHubDataFormat_STATUS_CSV        = IotHubDataFormat_STATUS("CSV")
	IotHubDataFormat_STATUS_JSON       = IotHubDataFormat_STATUS("JSON")
	IotHubDataFormat_STATUS_MULTIJSON  = IotHubDataFormat_STATUS("MULTIJSON")
	IotHubDataFormat_STATUS_ORC        = IotHubDataFormat_STATUS("ORC")
	IotHubDataFormat_STATUS_PARQUET    = IotHubDataFormat_STATUS("PARQUET")
	IotHubDataFormat_STATUS_PSV        = IotHubDataFormat_STATUS("PSV")
	IotHubDataFormat_STATUS_RAW        = IotHubDataFormat_STATUS("RAW")
	IotHubDataFormat_STATUS_SCSV       = IotHubDataFormat_STATUS("SCSV")
	IotHubDataFormat_STATUS_SINGLEJSON = IotHubDataFormat_STATUS("SINGLEJSON")
	IotHubDataFormat_STATUS_SOHSV      = IotHubDataFormat_STATUS("SOHSV")
	IotHubDataFormat_STATUS_TSV        = IotHubDataFormat_STATUS("TSV")
	IotHubDataFormat_STATUS_TSVE       = IotHubDataFormat_STATUS("TSVE")
	IotHubDataFormat_STATUS_TXT        = IotHubDataFormat_STATUS("TXT")
	IotHubDataFormat_STATUS_W3CLOGFILE = IotHubDataFormat_STATUS("W3CLOGFILE")
)

// Mapping from string to IotHubDataFormat_STATUS
var iotHubDataFormat_STATUS_Values = map[string]IotHubDataFormat_STATUS{
	"apacheavro": IotHubDataFormat_STATUS_APACHEAVRO,
	"avro":       IotHubDataFormat_STATUS_AVRO,
	"csv":        IotHubDataFormat_STATUS_CSV,
	"json":       IotHubDataFormat_STATUS_JSON,
	"multijson":  IotHubDataFormat_STATUS_MULTIJSON,
	"orc":        IotHubDataFormat_STATUS_ORC,
	"parquet":    IotHubDataFormat_STATUS_PARQUET,
	"psv":        IotHubDataFormat_STATUS_PSV,
	"raw":        IotHubDataFormat_STATUS_RAW,
	"scsv":       IotHubDataFormat_STATUS_SCSV,
	"singlejson": IotHubDataFormat_STATUS_SINGLEJSON,
	"sohsv":      IotHubDataFormat_STATUS_SOHSV,
	"tsv":        IotHubDataFormat_STATUS_TSV,
	"tsve":       IotHubDataFormat_STATUS_TSVE,
	"txt":        IotHubDataFormat_STATUS_TXT,
	"w3clogfile": IotHubDataFormat_STATUS_W3CLOGFILE,
}
