// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

//Deprecated version of SBQueue_Status. Use v1beta20210101preview.SBQueue_Status instead
type SBQueue_StatusARM struct {
	Id         *string                      `json:"id,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	Properties *SBQueueProperties_StatusARM `json:"properties,omitempty"`
	SystemData *SystemData_StatusARM        `json:"systemData,omitempty"`
	Type       *string                      `json:"type,omitempty"`
}

//Deprecated version of SBQueueProperties_Status. Use v1beta20210101preview.SBQueueProperties_Status instead
type SBQueueProperties_StatusARM struct {
	AccessedAt                          *string                        `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                        `json:"autoDeleteOnIdle,omitempty"`
	CountDetails                        *MessageCountDetails_StatusARM `json:"countDetails,omitempty"`
	CreatedAt                           *string                        `json:"createdAt,omitempty"`
	DeadLetteringOnMessageExpiration    *bool                          `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string                        `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                        `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                          `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                          `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                          `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string                        `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string                        `json:"forwardTo,omitempty"`
	LockDuration                        *string                        `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int                           `json:"maxDeliveryCount,omitempty"`
	MaxSizeInMegabytes                  *int                           `json:"maxSizeInMegabytes,omitempty"`
	MessageCount                        *int                           `json:"messageCount,omitempty"`
	RequiresDuplicateDetection          *bool                          `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession                     *bool                          `json:"requiresSession,omitempty"`
	SizeInBytes                         *int                           `json:"sizeInBytes,omitempty"`
	Status                              *EntityStatus_Status           `json:"status,omitempty"`
	UpdatedAt                           *string                        `json:"updatedAt,omitempty"`
}

//Deprecated version of MessageCountDetails_Status. Use v1beta20210101preview.MessageCountDetails_Status instead
type MessageCountDetails_StatusARM struct {
	ActiveMessageCount             *int `json:"activeMessageCount,omitempty"`
	DeadLetterMessageCount         *int `json:"deadLetterMessageCount,omitempty"`
	ScheduledMessageCount          *int `json:"scheduledMessageCount,omitempty"`
	TransferDeadLetterMessageCount *int `json:"transferDeadLetterMessageCount,omitempty"`
	TransferMessageCount           *int `json:"transferMessageCount,omitempty"`
}
