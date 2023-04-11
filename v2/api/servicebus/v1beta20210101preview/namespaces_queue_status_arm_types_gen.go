// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

// Deprecated version of Namespaces_Queue_STATUS. Use v1api20210101preview.Namespaces_Queue_STATUS instead
type Namespaces_Queue_STATUS_ARM struct {
	Id         *string                       `json:"id,omitempty"`
	Name       *string                       `json:"name,omitempty"`
	Properties *SBQueueProperties_STATUS_ARM `json:"properties,omitempty"`
	SystemData *SystemData_STATUS_ARM        `json:"systemData,omitempty"`
	Type       *string                       `json:"type,omitempty"`
}

// Deprecated version of SBQueueProperties_STATUS. Use v1api20210101preview.SBQueueProperties_STATUS instead
type SBQueueProperties_STATUS_ARM struct {
	AccessedAt                          *string                         `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                         `json:"autoDeleteOnIdle,omitempty"`
	CountDetails                        *MessageCountDetails_STATUS_ARM `json:"countDetails,omitempty"`
	CreatedAt                           *string                         `json:"createdAt,omitempty"`
	DeadLetteringOnMessageExpiration    *bool                           `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string                         `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                         `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                           `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                           `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                           `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string                         `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string                         `json:"forwardTo,omitempty"`
	LockDuration                        *string                         `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int                            `json:"maxDeliveryCount,omitempty"`
	MaxSizeInMegabytes                  *int                            `json:"maxSizeInMegabytes,omitempty"`
	MessageCount                        *int                            `json:"messageCount,omitempty"`
	RequiresDuplicateDetection          *bool                           `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession                     *bool                           `json:"requiresSession,omitempty"`
	SizeInBytes                         *int                            `json:"sizeInBytes,omitempty"`
	Status                              *EntityStatus_STATUS            `json:"status,omitempty"`
	UpdatedAt                           *string                         `json:"updatedAt,omitempty"`
}

// Deprecated version of MessageCountDetails_STATUS. Use v1api20210101preview.MessageCountDetails_STATUS instead
type MessageCountDetails_STATUS_ARM struct {
	ActiveMessageCount             *int `json:"activeMessageCount,omitempty"`
	DeadLetterMessageCount         *int `json:"deadLetterMessageCount,omitempty"`
	ScheduledMessageCount          *int `json:"scheduledMessageCount,omitempty"`
	TransferDeadLetterMessageCount *int `json:"transferDeadLetterMessageCount,omitempty"`
	TransferMessageCount           *int `json:"transferMessageCount,omitempty"`
}
