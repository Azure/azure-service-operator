// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

<<<<<<<< HEAD:v2/api/servicebus/v1alpha1api20210101preview/namespaces_queue__spec_arm_types_gen.go
// Deprecated version of NamespacesQueue_Spec. Use v1beta20210101preview.NamespacesQueue_Spec instead
type NamespacesQueue_SpecARM struct {
	AzureName  string                `json:"azureName,omitempty"`
========
// Deprecated version of Namespaces_Queues_Spec. Use v1beta20210101preview.Namespaces_Queues_Spec instead
type Namespaces_Queues_SpecARM struct {
	Location   *string               `json:"location,omitempty"`
>>>>>>>> main:v2/api/servicebus/v1alpha1api20210101preview/namespaces_queues_spec_arm_types_gen.go
	Name       string                `json:"name,omitempty"`
	Properties *SBQueuePropertiesARM `json:"properties,omitempty"`
}

<<<<<<<< HEAD:v2/api/servicebus/v1alpha1api20210101preview/namespaces_queue__spec_arm_types_gen.go
var _ genruntime.ARMResourceSpec = &NamespacesQueue_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (queue NamespacesQueue_SpecARM) GetAPIVersion() string {
========
var _ genruntime.ARMResourceSpec = &Namespaces_Queues_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (queues Namespaces_Queues_SpecARM) GetAPIVersion() string {
>>>>>>>> main:v2/api/servicebus/v1alpha1api20210101preview/namespaces_queues_spec_arm_types_gen.go
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
<<<<<<<< HEAD:v2/api/servicebus/v1alpha1api20210101preview/namespaces_queue__spec_arm_types_gen.go
func (queue *NamespacesQueue_SpecARM) GetName() string {
	return queue.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/queues"
func (queue *NamespacesQueue_SpecARM) GetType() string {
========
func (queues *Namespaces_Queues_SpecARM) GetName() string {
	return queues.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/queues"
func (queues *Namespaces_Queues_SpecARM) GetType() string {
>>>>>>>> main:v2/api/servicebus/v1alpha1api20210101preview/namespaces_queues_spec_arm_types_gen.go
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
