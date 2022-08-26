// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespaces_eventhub__spec_arm_types_gen.go
type NamespacesEventhub_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`
	Name      string `json:"name,omitempty"`

	// Properties: Properties supplied to the Create Or Update Event Hub operation.
	Properties *NamespacesEventhub_Spec_PropertiesARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesEventhub_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (eventhub NamespacesEventhub_SpecARM) GetAPIVersion() string {
========
type Namespaces_Eventhubs_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Properties: Properties supplied to the Create Or Update Event Hub operation.
	Properties *Namespaces_Eventhubs_Spec_PropertiesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespaces_Eventhubs_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (eventhubs Namespaces_Eventhubs_SpecARM) GetAPIVersion() string {
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_eventhubs_spec_arm_types_gen.go
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespaces_eventhub__spec_arm_types_gen.go
func (eventhub *NamespacesEventhub_SpecARM) GetName() string {
	return eventhub.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs"
func (eventhub *NamespacesEventhub_SpecARM) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs"
}

type NamespacesEventhub_Spec_PropertiesARM struct {
	// CaptureDescription: Properties of capture description
	CaptureDescription *CaptureDescriptionARM `json:"captureDescription,omitempty"`
========
func (eventhubs *Namespaces_Eventhubs_SpecARM) GetName() string {
	return eventhubs.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs"
func (eventhubs *Namespaces_Eventhubs_SpecARM) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs"
}

type Namespaces_Eventhubs_Spec_PropertiesARM struct {
	// CaptureDescription: Properties to configure capture description for eventhub
	CaptureDescription *Namespaces_Eventhubs_Spec_Properties_CaptureDescriptionARM `json:"captureDescription,omitempty"`
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_eventhubs_spec_arm_types_gen.go

	// MessageRetentionInDays: Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays *int `json:"messageRetentionInDays,omitempty"`

	// PartitionCount: Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount *int `json:"partitionCount,omitempty"`

	// Status: Enumerates the possible values for the status of the Event Hub.
	Status *NamespacesEventhub_Spec_Properties_Status `json:"status,omitempty"`
}

<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespaces_eventhub__spec_arm_types_gen.go
type CaptureDescriptionARM struct {
	// Destination: Properties of Destination where capture will be stored. (Storage Account, Blob Names)
	Destination *DestinationARM `json:"destination,omitempty"`
========
type Namespaces_Eventhubs_Spec_Properties_CaptureDescriptionARM struct {
	// Destination: Capture storage details for capture description
	Destination *Namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationARM `json:"destination,omitempty"`
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_eventhubs_spec_arm_types_gen.go

	// Enabled: A value that indicates whether capture description is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Encoding: Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be
<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespaces_eventhub__spec_arm_types_gen.go
	// deprecated in New API Version
	Encoding *CaptureDescription_Encoding `json:"encoding,omitempty"`
========
	// deprecated in New API Version.
	Encoding *Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Encoding `json:"encoding,omitempty"`
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_eventhubs_spec_arm_types_gen.go

	// IntervalInSeconds: The time window allows you to set the frequency with which the capture to Azure Blobs will happen,
	// value should between 60 to 900 seconds
	IntervalInSeconds *int `json:"intervalInSeconds,omitempty"`

	// SizeLimitInBytes: The size window defines the amount of data built up in your Event Hub before an capture operation,
	// value should be between 10485760 to 524288000 bytes
	SizeLimitInBytes *int `json:"sizeLimitInBytes,omitempty"`

	// SkipEmptyArchives: A value that indicates whether to Skip Empty Archives
	SkipEmptyArchives *bool `json:"skipEmptyArchives,omitempty"`
}

<<<<<<<< HEAD:v2/api/eventhub/v1beta20211101/namespaces_eventhub__spec_arm_types_gen.go
type DestinationARM struct {
========
type Namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationARM struct {
>>>>>>>> main:v2/api/eventhub/v1beta20211101/namespaces_eventhubs_spec_arm_types_gen.go
	// Name: Name for capture destination
	Name *string `json:"name,omitempty"`

	// Properties: Properties describing the storage account, blob container and archive name format for capture destination
	Properties *Destination_PropertiesARM `json:"properties,omitempty"`
}

type Destination_PropertiesARM struct {
	// ArchiveNameFormat: Blob naming convention for archive, e.g.
	// {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters
	// (Namespace,EventHub .. etc) are mandatory irrespective of order
	ArchiveNameFormat *string `json:"archiveNameFormat,omitempty"`

	// BlobContainer: Blob container Name
	BlobContainer *string `json:"blobContainer,omitempty"`

	// DataLakeAccountName: The Azure Data Lake Store name for the captured events
	DataLakeAccountName *string `json:"dataLakeAccountName,omitempty"`

	// DataLakeFolderPath: The destination folder path for the captured events
	DataLakeFolderPath *string `json:"dataLakeFolderPath,omitempty"`

	// DataLakeSubscriptionId: Subscription Id of Azure Data Lake Store
	DataLakeSubscriptionId   *string `json:"dataLakeSubscriptionId,omitempty"`
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
}
