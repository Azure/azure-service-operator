// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

type Eventhub_StatusARM struct {
	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Properties: Properties supplied to the Create Or Update Event Hub operation.
	Properties *Eventhub_Status_PropertiesARM `json:"properties,omitempty"`

	//SystemData: The system meta data relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

type Eventhub_Status_PropertiesARM struct {
	//CaptureDescription: Properties of capture description
	CaptureDescription *CaptureDescription_StatusARM `json:"captureDescription,omitempty"`

	//CreatedAt: Exact time the Event Hub was created.
	CreatedAt *string `json:"createdAt,omitempty"`

	//MessageRetentionInDays: Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays *int `json:"messageRetentionInDays,omitempty"`

	//PartitionCount: Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount *int `json:"partitionCount,omitempty"`

	//PartitionIds: Current number of shards on the Event Hub.
	PartitionIds []string `json:"partitionIds,omitempty"`

	//Status: Enumerates the possible values for the status of the Event Hub.
	Status *EventhubStatusPropertiesStatus `json:"status,omitempty"`

	//UpdatedAt: The exact time the message was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type CaptureDescription_StatusARM struct {
	//Destination: Properties of Destination where capture will be stored. (Storage Account, Blob Names)
	Destination *Destination_StatusARM `json:"destination,omitempty"`

	//Enabled: A value that indicates whether capture description is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	//Encoding: Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be
	//deprecated in New API Version
	Encoding *CaptureDescriptionStatusEncoding `json:"encoding,omitempty"`

	//IntervalInSeconds: The time window allows you to set the frequency with which the capture to Azure Blobs will happen,
	//value should between 60 to 900 seconds
	IntervalInSeconds *int `json:"intervalInSeconds,omitempty"`

	//SizeLimitInBytes: The size window defines the amount of data built up in your Event Hub before an capture operation,
	//value should be between 10485760 to 524288000 bytes
	SizeLimitInBytes *int `json:"sizeLimitInBytes,omitempty"`

	//SkipEmptyArchives: A value that indicates whether to Skip Empty Archives
	SkipEmptyArchives *bool `json:"skipEmptyArchives,omitempty"`
}

type EventhubStatusPropertiesStatus string

const (
	EventhubStatusPropertiesStatusActive          = EventhubStatusPropertiesStatus("Active")
	EventhubStatusPropertiesStatusCreating        = EventhubStatusPropertiesStatus("Creating")
	EventhubStatusPropertiesStatusDeleting        = EventhubStatusPropertiesStatus("Deleting")
	EventhubStatusPropertiesStatusDisabled        = EventhubStatusPropertiesStatus("Disabled")
	EventhubStatusPropertiesStatusReceiveDisabled = EventhubStatusPropertiesStatus("ReceiveDisabled")
	EventhubStatusPropertiesStatusRenaming        = EventhubStatusPropertiesStatus("Renaming")
	EventhubStatusPropertiesStatusRestoring       = EventhubStatusPropertiesStatus("Restoring")
	EventhubStatusPropertiesStatusSendDisabled    = EventhubStatusPropertiesStatus("SendDisabled")
	EventhubStatusPropertiesStatusUnknown         = EventhubStatusPropertiesStatus("Unknown")
)

type CaptureDescriptionStatusEncoding string

const (
	CaptureDescriptionStatusEncodingAvro        = CaptureDescriptionStatusEncoding("Avro")
	CaptureDescriptionStatusEncodingAvroDeflate = CaptureDescriptionStatusEncoding("AvroDeflate")
)

type Destination_StatusARM struct {
	//Name: Name for capture destination
	Name *string `json:"name,omitempty"`

	//Properties: Properties describing the storage account, blob container and archive name format for capture destination
	Properties *Destination_Status_PropertiesARM `json:"properties,omitempty"`
}

type Destination_Status_PropertiesARM struct {
	//ArchiveNameFormat: Blob naming convention for archive, e.g.
	//{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters
	//(Namespace,EventHub .. etc) are mandatory irrespective of order
	ArchiveNameFormat *string `json:"archiveNameFormat,omitempty"`

	//BlobContainer: Blob container Name
	BlobContainer *string `json:"blobContainer,omitempty"`

	//DataLakeAccountName: The Azure Data Lake Store name for the captured events
	DataLakeAccountName *string `json:"dataLakeAccountName,omitempty"`

	//DataLakeFolderPath: The destination folder path for the captured events
	DataLakeFolderPath *string `json:"dataLakeFolderPath,omitempty"`

	//DataLakeSubscriptionId: Subscription Id of Azure Data Lake Store
	DataLakeSubscriptionId *string `json:"dataLakeSubscriptionId,omitempty"`

	//StorageAccountResourceId: Resource id of the storage account to be used to create the blobs
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
}
