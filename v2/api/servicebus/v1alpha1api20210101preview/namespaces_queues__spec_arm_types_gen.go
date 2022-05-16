// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of NamespacesQueues_Spec. Use v1beta20210101preview.NamespacesQueues_Spec instead
type NamespacesQueues_SpecARM struct {
	Location   *string               `json:"location,omitempty"`
	Name       string                `json:"name,omitempty"`
	Properties *SBQueuePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string     `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesQueues_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (queues NamespacesQueues_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (queues NamespacesQueues_SpecARM) GetName() string {
	return queues.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/queues"
func (queues NamespacesQueues_SpecARM) GetType() string {
	return "Microsoft.ServiceBus/namespaces/queues"
}

// Deprecated version of SBQueueProperties. Use v1beta20210101preview.SBQueueProperties instead
type SBQueuePropertiesARM struct {
	AutoDeleteOnIdle                    *string `json:"autoDeleteOnIdle,omitempty"`
	DeadLetteringOnMessageExpiration    *bool   `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool   `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool   `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool   `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string `json:"forwardTo,omitempty"`
	LockDuration                        *string `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int    `json:"maxDeliveryCount,omitempty"`
	MaxSizeInMegabytes                  *int    `json:"maxSizeInMegabytes,omitempty"`
	RequiresDuplicateDetection          *bool   `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession                     *bool   `json:"requiresSession,omitempty"`
}
