// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=servicebus.azure.com,resources=namespacesqueues,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=servicebus.azure.com,resources={namespacesqueues/status,namespacesqueues/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210101preview.NamespacesQueue
// Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_queues
type NamespacesQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesQueues_Spec `json:"spec,omitempty"`
	Status            SBQueue_Status        `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesQueue{}

// GetConditions returns the conditions of the resource
func (queue *NamespacesQueue) GetConditions() conditions.Conditions {
	return queue.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (queue *NamespacesQueue) SetConditions(conditions conditions.Conditions) {
	queue.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NamespacesQueue{}

// AzureName returns the Azure name of the resource
func (queue *NamespacesQueue) AzureName() string {
	return queue.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (queue NamespacesQueue) GetAPIVersion() string {
	return "2021-01-01-preview"
}

// GetResourceKind returns the kind of the resource
func (queue *NamespacesQueue) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (queue *NamespacesQueue) GetSpec() genruntime.ConvertibleSpec {
	return &queue.Spec
}

// GetStatus returns the status of this resource
func (queue *NamespacesQueue) GetStatus() genruntime.ConvertibleStatus {
	return &queue.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/queues"
func (queue *NamespacesQueue) GetType() string {
	return "Microsoft.ServiceBus/namespaces/queues"
}

// NewEmptyStatus returns a new empty (blank) status
func (queue *NamespacesQueue) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SBQueue_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (queue *NamespacesQueue) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(queue.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  queue.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (queue *NamespacesQueue) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SBQueue_Status); ok {
		queue.Status = *st
		return nil
	}

	// Convert status to required version
	var st SBQueue_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	queue.Status = st
	return nil
}

// Hub marks that this NamespacesQueue is the hub type for conversion
func (queue *NamespacesQueue) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (queue *NamespacesQueue) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: queue.Spec.OriginalVersion,
		Kind:    "NamespacesQueue",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210101preview.NamespacesQueue
// Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_queues
type NamespacesQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesQueue `json:"items"`
}

// Storage version of v1beta20210101preview.NamespacesQueues_Spec
type NamespacesQueues_Spec struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                           string  `json:"azureName,omitempty"`
	DeadLetteringOnMessageExpiration    *bool   `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool   `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool   `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool   `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string `json:"forwardTo,omitempty"`
	Location                            *string `json:"location,omitempty"`
	LockDuration                        *string `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int    `json:"maxDeliveryCount,omitempty"`
	MaxSizeInMegabytes                  *int    `json:"maxSizeInMegabytes,omitempty"`
	OriginalVersion                     string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a servicebus.azure.com/Namespace resource
	Owner                      *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection *bool                              `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession            *bool                              `json:"requiresSession,omitempty"`
	Tags                       map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesQueues_Spec{}

// ConvertSpecFrom populates our NamespacesQueues_Spec from the provided source
func (queues *NamespacesQueues_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == queues {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(queues)
}

// ConvertSpecTo populates the provided destination from our NamespacesQueues_Spec
func (queues *NamespacesQueues_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == queues {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(queues)
}

// Storage version of v1beta20210101preview.SBQueue_Status
type SBQueue_Status struct {
	AccessedAt                          *string                     `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                     `json:"autoDeleteOnIdle,omitempty"`
	Conditions                          []conditions.Condition      `json:"conditions,omitempty"`
	CountDetails                        *MessageCountDetails_Status `json:"countDetails,omitempty"`
	CreatedAt                           *string                     `json:"createdAt,omitempty"`
	DeadLetteringOnMessageExpiration    *bool                       `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string                     `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                     `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                       `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                       `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                       `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string                     `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string                     `json:"forwardTo,omitempty"`
	Id                                  *string                     `json:"id,omitempty"`
	LockDuration                        *string                     `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int                        `json:"maxDeliveryCount,omitempty"`
	MaxSizeInMegabytes                  *int                        `json:"maxSizeInMegabytes,omitempty"`
	MessageCount                        *int                        `json:"messageCount,omitempty"`
	Name                                *string                     `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection          *bool                       `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession                     *bool                       `json:"requiresSession,omitempty"`
	SizeInBytes                         *int                        `json:"sizeInBytes,omitempty"`
	Status                              *string                     `json:"status,omitempty"`
	SystemData                          *SystemData_Status          `json:"systemData,omitempty"`
	Type                                *string                     `json:"type,omitempty"`
	UpdatedAt                           *string                     `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SBQueue_Status{}

// ConvertStatusFrom populates our SBQueue_Status from the provided source
func (queue *SBQueue_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == queue {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(queue)
}

// ConvertStatusTo populates the provided destination from our SBQueue_Status
func (queue *SBQueue_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == queue {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(queue)
}

// Storage version of v1beta20210101preview.MessageCountDetails_Status
type MessageCountDetails_Status struct {
	ActiveMessageCount             *int                   `json:"activeMessageCount,omitempty"`
	DeadLetterMessageCount         *int                   `json:"deadLetterMessageCount,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScheduledMessageCount          *int                   `json:"scheduledMessageCount,omitempty"`
	TransferDeadLetterMessageCount *int                   `json:"transferDeadLetterMessageCount,omitempty"`
	TransferMessageCount           *int                   `json:"transferMessageCount,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NamespacesQueue{}, &NamespacesQueueList{})
}
