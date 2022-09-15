// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

// Deprecated version of Eventhub_STATUS. Use v1beta20211101.Eventhub_STATUS instead
type Eventhub_STATUS_ARM struct {
	Id         *string                         `json:"id,omitempty"`
	Location   *string                         `json:"location,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Properties *Eventhub_Properties_STATUS_ARM `json:"properties,omitempty"`
	SystemData *SystemData_STATUS_ARM          `json:"systemData,omitempty"`
	Type       *string                         `json:"type,omitempty"`
}

// Deprecated version of Eventhub_Properties_STATUS. Use v1beta20211101.Eventhub_Properties_STATUS instead
type Eventhub_Properties_STATUS_ARM struct {
	CaptureDescription     *CaptureDescription_STATUS_ARM     `json:"captureDescription,omitempty"`
	CreatedAt              *string                            `json:"createdAt,omitempty"`
	MessageRetentionInDays *int                               `json:"messageRetentionInDays,omitempty"`
	PartitionCount         *int                               `json:"partitionCount,omitempty"`
	PartitionIds           []string                           `json:"partitionIds,omitempty"`
	Status                 *Eventhub_Properties_Status_STATUS `json:"status,omitempty"`
	UpdatedAt              *string                            `json:"updatedAt,omitempty"`
}

// Deprecated version of CaptureDescription_STATUS. Use v1beta20211101.CaptureDescription_STATUS instead
type CaptureDescription_STATUS_ARM struct {
	Destination       *Destination_STATUS_ARM             `json:"destination,omitempty"`
	Enabled           *bool                               `json:"enabled,omitempty"`
	Encoding          *CaptureDescription_Encoding_STATUS `json:"encoding,omitempty"`
	IntervalInSeconds *int                                `json:"intervalInSeconds,omitempty"`
	SizeLimitInBytes  *int                                `json:"sizeLimitInBytes,omitempty"`
	SkipEmptyArchives *bool                               `json:"skipEmptyArchives,omitempty"`
}

// Deprecated version of Eventhub_Properties_Status_STATUS. Use v1beta20211101.Eventhub_Properties_Status_STATUS instead
type Eventhub_Properties_Status_STATUS string

const (
	Eventhub_Properties_Status_STATUS_Active          = Eventhub_Properties_Status_STATUS("Active")
	Eventhub_Properties_Status_STATUS_Creating        = Eventhub_Properties_Status_STATUS("Creating")
	Eventhub_Properties_Status_STATUS_Deleting        = Eventhub_Properties_Status_STATUS("Deleting")
	Eventhub_Properties_Status_STATUS_Disabled        = Eventhub_Properties_Status_STATUS("Disabled")
	Eventhub_Properties_Status_STATUS_ReceiveDisabled = Eventhub_Properties_Status_STATUS("ReceiveDisabled")
	Eventhub_Properties_Status_STATUS_Renaming        = Eventhub_Properties_Status_STATUS("Renaming")
	Eventhub_Properties_Status_STATUS_Restoring       = Eventhub_Properties_Status_STATUS("Restoring")
	Eventhub_Properties_Status_STATUS_SendDisabled    = Eventhub_Properties_Status_STATUS("SendDisabled")
	Eventhub_Properties_Status_STATUS_Unknown         = Eventhub_Properties_Status_STATUS("Unknown")
)

// Deprecated version of CaptureDescription_Encoding_STATUS. Use v1beta20211101.CaptureDescription_Encoding_STATUS instead
type CaptureDescription_Encoding_STATUS string

const (
	CaptureDescription_Encoding_STATUS_Avro        = CaptureDescription_Encoding_STATUS("Avro")
	CaptureDescription_Encoding_STATUS_AvroDeflate = CaptureDescription_Encoding_STATUS("AvroDeflate")
)

// Deprecated version of Destination_STATUS. Use v1beta20211101.Destination_STATUS instead
type Destination_STATUS_ARM struct {
	Name       *string                            `json:"name,omitempty"`
	Properties *Destination_Properties_STATUS_ARM `json:"properties,omitempty"`
}

// Deprecated version of Destination_Properties_STATUS. Use v1beta20211101.Destination_Properties_STATUS instead
type Destination_Properties_STATUS_ARM struct {
	ArchiveNameFormat        *string `json:"archiveNameFormat,omitempty"`
	BlobContainer            *string `json:"blobContainer,omitempty"`
	DataLakeAccountName      *string `json:"dataLakeAccountName,omitempty"`
	DataLakeFolderPath       *string `json:"dataLakeFolderPath,omitempty"`
	DataLakeSubscriptionId   *string `json:"dataLakeSubscriptionId,omitempty"`
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
}
